![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

# YuanJing OpenAPI SDK for Go

## Requirements
- It's necessary for you to make sure your system have installed Go environment which version greater than 1.15.0.

## Installation
If you use `go mod` to manage your dependence, you can use the following command:
```
go get github.com/aliyun/alibabacloud-yjopenapi-go-client 1.0.20230410
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
 | *ConsoleAdminApi* | **ActivateDeployment** | *ActivateDeploymentForms*  | *ConsoleAdminActivateDeploymentResult* | 激活已部署成功的游戏版本的部署 |
 | *ConsoleAdminApi* | **AdaptGameVersion** | *AdaptGameVersionForms*  | *ConsoleAdminAdaptGameVersionResult* | 发起游戏版本适配 |
 | *ConsoleAdminApi* | **AddGameToProject** | *AddGameToProjectForms*  | *ConsoleAdminAddGameToProjectResult* | 将游戏添加到项目 |
 | *ConsoleAdminApi* | **CreateGame** | *CreateGameForms*  | *ConsoleAdminCreateGameResult* | 创建游戏 |
 | *ConsoleAdminApi* | **CreateProject** | *CreateProjectForms*  | *ConsoleAdminCreateProjectResult* | 创建项目 |
 | *ConsoleAdminApi* | **DeleteGame** | *DeleteGameForms*  | *ConsoleAdminDeleteGameResult* | 删除指定的游戏 |
 | *ConsoleAdminApi* | **DeleteGameVersion** | *DeleteGameVersionForms*  | *ConsoleAdminDeleteGameVersionResult* | 发起游戏版本适配 |
 | *ConsoleAdminApi* | **DeleteProject** | *DeleteProjectForms*  | *ConsoleAdminDeleteProjectResult* | 删除指定的项目 |
 | *ConsoleAdminApi* | **GetGameVersion** | *GetGameVersionForms*  | *ConsoleAdminGetGameVersionResult* | 获取单个游戏版本信息 |
 | *ConsoleAdminApi* | **GetGameVersionProgress** | *GetGameVersionProgressForms*  | *ConsoleAdminGetGameVersionProgressResult* | 查询版本处理进度（包含上传、适配、部署） |
 | *ConsoleAdminApi* | **ListActivateableInstances** | *ListActivateableInstancesForms*  | *ConsoleAdminListActivateableInstancesResult* | 指定项目和游戏版本，获取可激活且可调度的实例及调度配置 |
 | *ConsoleAdminApi* | **ListActivatedInstances** | *ListActivatedInstancesForms*  | *ConsoleAdminListActivatedInstancesResult* | 指定项目和游戏，获取已激活版本的可调度实例及调度配置 |
 | *ConsoleAdminApi* | **ListControllersOfGame** | *ListControllersOfGameForms*  | *ConsoleAdminListControllersOfGameResult* | 获取单个游戏关联的控制器列表 |
 | *ConsoleAdminApi* | **ListDeployableInstances** | *ListDeployableInstancesForms*  | *ConsoleAdminListDeployableInstancesResult* | 指定项目和游戏版本，获取可以部署的实例 |
 | *ConsoleAdminApi* | **ListGameVersions** | *ListGameVersionsForms*  | *ConsoleAdminListGameVersionsResult* | 分页获取游戏版本列表 |
 | *ConsoleAdminApi* | **ListGames** | *ListGamesForms*  | *ConsoleAdminListGamesResult* | 分页获取游戏列表 |
 | *ConsoleAdminApi* | **ListProjects** | *ListProjectsForms*  | *ConsoleAdminListProjectsResult* | 分页获取项目列表 |
 | *ConsoleAdminApi* | **RemoveGameFromProject** | *RemoveGameFromProjectForms*  | *ConsoleAdminRemoveGameFromProjectResult* | 将游戏移出项目 |
 | *ConsoleAdminApi* | **SubmitDeployment** | *SubmitDeploymentForms*  | *ConsoleAdminSubmitDeploymentResult* | 提交游戏版本的部署请求 |
 | *ConsoleAdminApi* | **UploadGameVersionByDownload** | *UploadGameVersionByDownloadForms*  | *ConsoleAdminUploadGameVersionByDownloadResult* | 一键上传：文件上传接口，用远程下载的方式生成新版本 |
 | *DispatchApi* | **BatchStopGame** | *BatchStopGameForms*  | *BatchStopGameResult* | 游戏下全量踢下线，异步接口 |
 | *DispatchApi* | **CancelGameHang** | *CancelGameHangForms*  | *CancelGameHangResult* | 取消游戏挂机 |
 | *DispatchApi* | **GetGameConcurrency** | *GetGameConcurrencyForms*  | *GetGameConcurrencyResult* | 调用GetGameConcurrency获取游戏当前并发数 |
 | *DispatchApi* | **GetStock** | *GetStockForms*  | *GetStockResult* | 调用GetStock获取游戏当前库存 |
 | *DispatchApi* | **GetStopGameToken** | *GetStopGameTokenForms*  | *GetStopGameTokenResult* | 全量踢下线获取token |
 | *DispatchApi* | **ListGameServerIp** | *ListGameServerIpForms*  | *ListGameServerIpResult* | 获取自己租户下的游戏服务器ip列表 |
 | *DispatchApi* | **QueryGameHang** | *QueryGameHangForms*  | *QueryGameHangResult* | 查询游戏挂机状态 |
 | *DispatchApi* | **QuerySessionStatus** | *QuerySessionStatusForms*  | *QuerySessionStatusResult* | 查询会话当前状态 |
 | *DispatchApi* | **SetGameAlive** | *SetGameAliveForms*  | *SetGameAliveResult* | 设置游戏可运行时长 |
 | *DispatchApi* | **SetGameHang** | *SetGameHangForms*  | *SetGameHangResult* | 设置游戏挂机 |
 | *DispatchApi* | **StopGame** | *StopGameForms*  | *StopGameResult* | 服务端发起，停止某个用户的某个游戏的某个会话 |
 | *DispatchApi* | **TryToGetSlot** | *TryToGetSlotForms*  | *TryToGetSlotResult* | 为用户调度分配游戏容器，容器一旦分配成功会被锁住，一段时间内不再分配给其他用户，过期释放。 |
 | *InteractiveApi* | **GetParty** | *GetPartyForms*  | *InteractiveGetPartyResult* | 获取派对 |
 | *InteractiveApi* | **GetPartyStatus** | *GetPartyStatusForms*  | *InteractiveGetPartyStatusResult* | 查询派对游戏状态 |
 | *InteractiveApi* | **JoinParty** | *JoinPartyForms*  | *InteractiveJoinPartyResult* | 加入分配席位 |
 | *InteractiveApi* | **KickOutUser** | *KickOutUserForms*  | *InteractiveKickOutUserResult* | 踢出派对 |
 | *InteractiveApi* | **ModifySeats** | *ModifySeatsForms*  | *InteractiveModifySeatsResult* | 修改席位 |
 | *InteractiveApi* | **ShutDownParty** | *ShutDownPartyForms*  | *InteractiveShutDownPartyResult* | 关闭派对 |
 | *LiveApi* | **QueryStatus** | *QueryStatusForms*  | *LiveQueryStatusResult* | 查询推流状态 |
 | *LiveApi* | **StartGameLive** | *StartGameLiveForms*  | *LiveStartGameLiveResult* | 开始直播推流 |
 | *LiveApi* | **StopGameLive** | *StopGameLiveForms*  | *LiveStopGameLiveResult* | 结束直播推流 |
 | *MultiplayApi* | **Close** | *CloseForms*  | *MultiplayCloseResult* | 关闭联机 |
 | *MultiplayApi* | **Init** | *InitForms*  | *MultiplayInitResult* | 初始化联机 |
 | *MultiplayApi* | **Join** | *JoinForms*  | *MultiplayJoinResult* | 加入联机 |
 | *MultiplayApi* | **Leave** | *LeaveForms*  | *MultiplayLeaveResult* | 离开联机 |
 | *MultiplayApi* | **Modify** | *ModifyForms*  | *MultiplayModifyResult* | 修改联机 |
 | *MultiplayApi* | **Query** | *QueryForms*  | *MultiplayQueryResult* | 离开联机 |
 | *TokenApi* | **GetTriple** |   | *GetTripleResult* | 获取临时安全令牌 |
 | *UsercontrollerApi* | **DeleteGameArchive** | *DeleteGameArchiveForms*  | *UsercontollerDeleteGameArchiveResult* | 根据存档id删除存档纪录 |
 | *UsercontrollerApi* | **GetGameTrialSurplusDuration** | *GetGameTrialSurplusDurationForms*  | *UsercontollerGetGameTrialSurplusDurationResult* | 查询剩余试玩游戏时长 |
 | *UsercontrollerApi* | **ListLatestGameArchive** | *ListLatestGameArchiveForms*  | *UsercontollerListLatestGameArchiveResult* | 查询用户正常状态的最新存档纪录，按照存档时间倒序 |
 | *UsercontrollerApi* | **RestoreGameArchive** | *RestoreGameArchiveForms*  | *UsercontollerRestoreGameArchiveResult* | 将指定的存档ID恢复为最新存档 |
 | *UsercontrollerApi* | **UpdateGameArchiveTagStatus** | *UpdateGameArchiveTagStatusForms*  | *UsercontollerUpdateGameArchiveTagStatusResult* | 更新存档打标状态 |

## License
[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2009-present, Alibaba Cloud All rights reserved.


