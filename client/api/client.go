// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package api

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	Trace_Id      string = "Traceid"
	Result_Status string = "Result-Status"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")

	DefaultConfiguration = NewConfiguration()
)

// APIClient manages communication with the  API v1.0.20221015-beta
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	InteractiveApi *InteractiveApiService

	TokenApi *TokenApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.InteractiveApi = (*InteractiveApiService)(&c.common)
	c.TokenApi = (*TokenApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	//	var delimiter string

	//	switch collectionFormat {
	//	case "pipes":
	//		delimiter = "|"
	//	case "ssv":
	//		delimiter = " "
	//	case "tsv":
	//		delimiter = "\t"
	//	case "csv":
	//		delimiter = ","
	//	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		//		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
		if marshal, err := json.Marshal(obj); err != nil {
			return ""
		} else {
			return string(marshal)
		}
	}
	if reflect.TypeOf(obj).Kind() == reflect.Struct {
		if marshal, err := json.Marshal(obj); err != nil {
			return ""
		} else {
			return string(marshal)
		}
	}
	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.cfg.HTTPClient.Do(request)
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	path string,
	method string,
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// add form parameters
	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", "cgw-client/1.0.0/go")

	// prepare sign headers
	prepareSignHeader(queryParams, formParams, localVarRequest, c.cfg)

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if strings.Contains(contentType, "application/json") {
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

func prepareSignHeader(queryParams url.Values, formParams url.Values, localVarRequest *http.Request, cfg *Configuration) {
	accessKey := cfg.AccessKey
	nonce := randNonce(16)
	version := cfg.SignatureVersion
	timestamp := time.Now().UTC().Format(time.RFC3339)
	method := cfg.SignatureMethod

	values := url.Values{}
	values.Add("AccessKey", accessKey)
	values.Add("SignatureNonce", nonce)
	values.Add("SignatureVersion", version)
	values.Add("SignatureMethod", method)
	values.Add("Timestamp", timestamp)
	if len(queryParams) > 0 {
		for s, i := range queryParams {
			for _, s2 := range i {
				values.Add(s, s2)
			}
		}
	}
	if len(formParams) > 0 {
		for s, i := range formParams {
			for _, s2 := range i {
				values.Add(s, s2)
			}
		}
	}

	sign := md5Hex([]byte(getSignRaw(values, localVarRequest.Method, cfg.SecretKey)))
	localVarRequest.Header.Add("Signature", sign)
	localVarRequest.Header.Add("Accesskey", accessKey)
	localVarRequest.Header.Add("Signaturenonce", nonce)
	localVarRequest.Header.Add("Signatureversion", version)
	localVarRequest.Header.Add("Timestamp", timestamp)
	localVarRequest.Header.Add("Signaturemethod", method)
}

func md5Hex(data []byte) string {
	md5Sum := md5.Sum(data)
	return hex.EncodeToString(md5Sum[:])
}

func getSignRaw(values url.Values, method string, secret string) string {
	return secret + "&" + keySignInput(values, method)
}

func keySignInput(values url.Values, method string) string {
	keys := make([]string, 0)
	for s, _ := range values {
		keys = append(keys, s)
	}
	sort.Strings(keys)

	stringToSign := method + "&" + encodeURLComponent("/")

	strs := make([]string, 0)
	for _, k := range keys {
		strs = append(strs, k+"="+encodeURLComponent(values.Get(k)))
	}
	return stringToSign + "&" + encodeURLComponent(strings.Join(strs, "&"))
}

func encodeURLComponent(str string) string {
	escape := url.QueryEscape(str)
	escape = strings.ReplaceAll(escape, "+", "%20")
	escape = strings.ReplaceAll(escape, "*", "%2A")
	escape = strings.ReplaceAll(escape, "%7E", "~")
	return escape
}

func randNonce(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// GenericError Provides access to the body, error and model on returned errors.
type GenericError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericError) Model() interface{} {
	return e.model
}

type Configuration struct {
	Host             string            `json:"host"`
	Scheme           string            `json:"scheme"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	AccessKey        string            `json:"accessKey"`
	SecretKey        string            `json:"secretKey"`
	SignatureVersion string            `json:"signatureVersion"`
	SignatureMethod  string            `json:"signatureMethod"`
	HTTPClient       *http.Client
}

func NewConfiguration() *Configuration {
	return &Configuration{
		Scheme:           "https",
		SignatureVersion: "1.0",
		SignatureMethod:  "MD5",
		HTTPClient:       http.DefaultClient,
	}
}
