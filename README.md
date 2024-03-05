![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

# YuanJing OpenAPI SDK for Go

## Requirements
- It's necessary for you to make sure your system have installed Go environment which version greater than 1.15.0.

## Installation
If you use `go mod` to manage your dependence, you can use the following command:
```
go get github.com/aliyun/alibabacloud-yjopenapi-go-client 1.1.20240306
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
 | *AdaptApi* | **CreateAndSubmitAll** | *CreateAndSubmitAllForms*  | *AdaptCreateAndSubmitAllResult* | createAndSubmitAll |
 | *AdaptApi* | **QueryRequestById** | *QueryRequestByIdForms*  | *AdaptQueryRequestByIdResult* | queryRequestById |
 | *ConsoleAdminApi* | **ActivateDeployment** | *ActivateDeploymentForms*  | *ConsoleAdminActivateDeploymentResult* |  |
 | *ConsoleAdminApi* | **AdaptGameVersion** | *AdaptGameVersionForms*  | *ConsoleAdminAdaptGameVersionResult* |  |
 | *ConsoleAdminApi* | **AddGameToProject** | *AddGameToProjectForms*  | *ConsoleAdminAddGameToProjectResult* |  |
 | *ConsoleAdminApi* | **BatchUpdateDispatchConfig** | *BatchUpdateDispatchConfigForms*  | *ConsoleAdminBatchUpdateDispatchConfigResult* | 批量更新游戏各自调度配置 |
 | *ConsoleAdminApi* | **CreateGame** | *CreateGameForms*  | *ConsoleAdminCreateGameResult* |  |
 | *ConsoleAdminApi* | **CreateOrder** | *CreateOrderForms*  | *ConsoleAdminCreateOrderResult* | 订单下单 |
 | *ConsoleAdminApi* | **CreateProject** | *CreateProjectForms*  | *ConsoleAdminCreateProjectResult* |  |
 | *ConsoleAdminApi* | **DeleteGame** | *DeleteGameForms*  | *ConsoleAdminDeleteGameResult* |  |
 | *ConsoleAdminApi* | **DeleteGameVersion** | *DeleteGameVersionForms*  | *ConsoleAdminDeleteGameVersionResult* |  |
 | *ConsoleAdminApi* | **DeleteProject** | *DeleteProjectForms*  | *ConsoleAdminDeleteProjectResult* |  |
 | *ConsoleAdminApi* | **GetBillFlowInfo** | *GetBillFlowInfoForms*  | *ConsoleAdminGetBillFlowInfoResult* |  |
 | *ConsoleAdminApi* | **GetGameVersion** | *GetGameVersionForms*  | *ConsoleAdminGetGameVersionResult* |  |
 | *ConsoleAdminApi* | **GetGameVersionProgress** | *GetGameVersionProgressForms*  | *ConsoleAdminGetGameVersionProgressResult* |  |
 | *ConsoleAdminApi* | **GetOrder** | *GetOrderForms*  | *ConsoleAdminGetOrderResult* | 查询订单 |
 | *ConsoleAdminApi* | **ListActivateableInstances** | *ListActivateableInstancesForms*  | *ConsoleAdminListActivateableInstancesResult* |  |
 | *ConsoleAdminApi* | **ListActivatedInstances** | *ListActivatedInstancesForms*  | *ConsoleAdminListActivatedInstancesResult* |  |
 | *ConsoleAdminApi* | **ListControllersOfGame** | *ListControllersOfGameForms*  | *ConsoleAdminListControllersOfGameResult* |  |
 | *ConsoleAdminApi* | **ListDeployableInstances** | *ListDeployableInstancesForms*  | *ConsoleAdminListDeployableInstancesResult* |  |
 | *ConsoleAdminApi* | **ListGameDeployDetailsOfProject** | *ListGameDeployDetailsOfProjectForms*  | *ConsoleAdminListGameDeployDetailsOfProjectResult* | 获取项目下游戏部署版本信息。 |
 | *ConsoleAdminApi* | **ListGameVersions** | *ListGameVersionsForms*  | *ConsoleAdminListGameVersionsResult* |  |
 | *ConsoleAdminApi* | **ListGames** | *ListGamesForms*  | *ConsoleAdminListGamesResult* |  |
 | *ConsoleAdminApi* | **ListInstancesOfProject** | *ListInstancesOfProjectForms*  | *ConsoleAdminListInstancesOfProjectResult* | 分页获取项目中的实例 |
 | *ConsoleAdminApi* | **ListProjects** | *ListProjectsForms*  | *ConsoleAdminListProjectsResult* |  |
 | *ConsoleAdminApi* | **ListVersionDeployInstances** | *ListVersionDeployInstancesForms*  | *ConsoleAdminListVersionDeployInstancesResult* | 获取项目下游戏版本的部署实例信息。 |
 | *ConsoleAdminApi* | **QueryAdaptResultByVersionId** | *QueryAdaptResultByVersionIdForms*  | *ConsoleAdminQueryAdaptResultByVersionIdResult* | 查询适配结果 |
 | *ConsoleAdminApi* | **RecommendSpecification** | *RecommendSpecificationForms*  | *ConsoleAdminRecommendSpecificationResult* |  |
 | *ConsoleAdminApi* | **RemoveGameFromProject** | *RemoveGameFromProjectForms*  | *ConsoleAdminRemoveGameFromProjectResult* |  |
 | *ConsoleAdminApi* | **SubmitDeployment** | *SubmitDeploymentForms*  | *ConsoleAdminSubmitDeploymentResult* |  |
 | *ConsoleAdminApi* | **UploadGameVersionByDownload** | *UploadGameVersionByDownloadForms*  | *ConsoleAdminUploadGameVersionByDownloadResult* |  |
 | *DispatchApi* | **BatchStopGame** | *BatchStopGameForms*  | *BatchStopGameResult* |  |
 | *DispatchApi* | **CancelGameHang** | *CancelGameHangForms*  | *CancelGameHangResult* | 取消游戏挂机 |
 | *DispatchApi* | **ClientNotify** | *ClientNotifyForms*  | *ClientNotifyResult* | clientNotify |
 | *DispatchApi* | **GameNotify** | *GameNotifyForms*  | *GameNotifyResult* | 游戏通知接口 |
 | *DispatchApi* | **GetGameConcurrency** | *GetGameConcurrencyForms*  | *GetGameConcurrencyResult* | 调用GetGameConcurrency获取游戏当前并发数 |
 | *DispatchApi* | **GetStock** | *GetStockForms*  | *GetStockResult* | 调用GetStock获取游戏当前库存 |
 | *DispatchApi* | **GetStopGameToken** | *GetStopGameTokenForms*  | *GetStopGameTokenResult* | 全量踢下线获取token |
 | *DispatchApi* | **ListGameServerIp** | *ListGameServerIpForms*  | *ListGameServerIpResult* |  |
 | *DispatchApi* | **QueryGameHang** | *QueryGameHangForms*  | *QueryGameHangResult* | 查询游戏挂机状态 |
 | *DispatchApi* | **QuerySessionStatus** | *QuerySessionStatusForms*  | *QuerySessionStatusResult* | 查询会话当前状态 |
 | *DispatchApi* | **SetGameAlive** | *SetGameAliveForms*  | *SetGameAliveResult* | 设置游戏可运行时长 |
 | *DispatchApi* | **SetGameHang** | *SetGameHangForms*  | *SetGameHangResult* | 设置游戏挂机 |
 | *DispatchApi* | **StopGame** | *StopGameForms*  | *StopGameResult* | 服务端发起，停止某个用户的某个游戏的某个会话 |
 | *DispatchApi* | **StopPreopenContainer** | *StopPreopenContainerForms*  | *StopPreopenContainerResult* | 停止预开容器 |
 | *DispatchApi* | **TryToGetSlot** | *TryToGetSlotForms*  | *TryToGetSlotResult* | 为用户调度分配游戏容器，容器一旦分配成功会被锁住，一段时间内不再分配给其他用户，过期释放。 |
 | *DispatchApi* | **TryToGetSlots** | *TryToGetSlotsForms*  | *TryToGetSlotsResult* | tryToGetSlots |
 | *DispatchApi* | **UpdatePreopenStrategy** | *UpdatePreopenStrategyForms*  | *UpdatePreopenStrategyResult* | 更新预开预起策略 |
 | *LiveApi* | **QueryStatus** | *QueryStatusForms*  | *LiveQueryStatusResult* |  |
 | *LiveApi* | **StartGameLive** | *StartGameLiveForms*  | *LiveStartGameLiveResult* |  |
 | *LiveApi* | **StopGameLive** | *StopGameLiveForms*  | *LiveStopGameLiveResult* |  |
 | *MultiplayApi* | **Close** | *CloseForms*  | *MultiplayCloseResult* |  |
 | *MultiplayApi* | **Init** | *InitForms*  | *MultiplayInitResult* |  |
 | *MultiplayApi* | **Join** | *JoinForms*  | *MultiplayJoinResult* |  |
 | *MultiplayApi* | **Leave** | *LeaveForms*  | *MultiplayLeaveResult* |  |
 | *MultiplayApi* | **Modify** | *ModifyForms*  | *MultiplayModifyResult* |  |
 | *MultiplayApi* | **Query** | *QueryForms*  | *MultiplayQueryResult* |  |
 | *TokenApi* | **GetPair** |   | *GetPairResult* | 获取临时安全令牌(二元组) |
 | *TokenApi* | **GetTriple** |   | *GetTripleResult* | 获取临时安全令牌 |
 | *UsercontrollerApi* | **DeleteGameArchive** | *DeleteGameArchiveForms*  | *UsercontrollerDeleteGameArchiveResult* |  |
 | *UsercontrollerApi* | **GetGameTrialSurplusDuration** | *GetGameTrialSurplusDurationForms*  | *UsercontrollerGetGameTrialSurplusDurationResult* |  |
 | *UsercontrollerApi* | **GetUserGameArchive** | *GetUserGameArchiveForms*  | *UsercontrollerGetUserGameArchiveResult* |  |
 | *UsercontrollerApi* | **ListLatestGameArchive** | *ListLatestGameArchiveForms*  | *UsercontrollerListLatestGameArchiveResult* |  |
 | *UsercontrollerApi* | **RestoreGameArchive** | *RestoreGameArchiveForms*  | *UsercontrollerRestoreGameArchiveResult* |  |
 | *UsercontrollerApi* | **UpdateGameArchiveTagStatus** | *UpdateGameArchiveTagStatusForms*  | *UsercontrollerUpdateGameArchiveTagStatusResult* |  |

## License
[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2009-present, Alibaba Cloud All rights reserved.


