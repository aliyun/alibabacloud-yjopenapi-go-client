![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

# YuanJing OpenAPI SDK for Go

## Requirements
- It's necessary for you to make sure your system have installed Go environment which version greater than 1.15.0.

## Installation
If you use `go mod` to manage your dependence, you can use the following command:
```
go get github.com/aliyun/alibabacloud-yjopenapi-go-client v1.0.20221015-beta
```

## Usage
```
import (
    "github.com/aliyun/alibabacloud-yjopenapi-go-client/client/api"
    "github.com/aliyun/alibabacloud-yjopenapi-go-client/client/model"
)

configuration := api.DefaultConfiguration
configuration.Host = "host"
configuration.AccessKey = "Your Access Key"
configuration.SecretKey = "Your Secret Key"

client := api.NewAPIClient(configuration)

// {{Api}},{{Method}},{{Param}} is placeholder, take a look at Explain Of Usage Placeholder
result, response, error := client.{{Api}}.{{Method}}(&model.{{Params}}{})

// OpenAPI TraceId
traceId := response.Header.Get(client.Trace_Id)
// OpenAPI Status Code
statusCode := response.Header.Get(client.Result_Status)

// OpenAPI result
_ := result
```

## Explain Of Usage Placeholder

| Api | Method | Params | Result | Description |
| ------------ | ------------- | ------------- | ------------- | ------------- |
 | *InteractiveApi* | **GetParty** | *GetPartyForms*  | *InteractiveGetPartyResult* | 获取派对 |
 | *InteractiveApi* | **GetPartyStatus** | *GetPartyStatusForms*  | *InteractiveGetPartyStatusResult* | 查询派对游戏状态 |
 | *InteractiveApi* | **JoinParty** | *JoinPartyForms*  | *InteractiveJoinPartyResult* | 加入分配席位 |
 | *InteractiveApi* | **KickOutUser** | *KickOutUserForms*  | *InteractiveKickOutUserResult* | 踢出派对 |
 | *InteractiveApi* | **ModifySeats** | *ModifySeatsForms*  | *InteractiveModifySeatsResult* | 修改席位 |
 | *InteractiveApi* | **ShutDownParty** | *ShutDownPartyForms*  | *InteractiveShutDownPartyResult* | 关闭派对 |
 | *TokenApi* | **GetTriple** |   | *GetTripleResult* | 获取临时安全令牌 |

## License
[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2009-present, Alibaba Cloud All rights reserved.


