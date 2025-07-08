# \VirtualserverV1ServerActionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachVirtualServerSecurityGroup**](VirtualserverV1ServerActionAPI.md#AttachVirtualServerSecurityGroup) | **Post** /v1/servers/{server_id}/security-groups | Attach Virtual Server Security Group
[**CreateVirtualServerCustomImage**](VirtualserverV1ServerActionAPI.md#CreateVirtualServerCustomImage) | **Post** /v1/servers/{server_id}/images | Create Virtual Server Custom Image
[**CreateVirtualServerDump**](VirtualserverV1ServerActionAPI.md#CreateVirtualServerDump) | **Post** /v1/servers/{server_id}/dump | Create Virtual Server Dump
[**DetachVirtualServerSecurityGroup**](VirtualserverV1ServerActionAPI.md#DetachVirtualServerSecurityGroup) | **Delete** /v1/servers/{server_id}/security-groups/{security_group_id} | Detach Virtual Server Security Group
[**ListVirtualServerSecurityGroups**](VirtualserverV1ServerActionAPI.md#ListVirtualServerSecurityGroups) | **Get** /v1/servers/{server_id}/security-groups | List Virtual Server Security Group
[**LockVirtualServer**](VirtualserverV1ServerActionAPI.md#LockVirtualServer) | **Post** /v1/servers/{server_id}/lock | Lock Virtual Server
[**RebootVirtualServer**](VirtualserverV1ServerActionAPI.md#RebootVirtualServer) | **Post** /v1/servers/{server_id}/reboot | Reboot Virtual Server
[**RebuildVirtualServer**](VirtualserverV1ServerActionAPI.md#RebuildVirtualServer) | **Post** /v1/servers/{server_id}/rebuild | Rebuild Virtual Server
[**SetVirtualServerType**](VirtualserverV1ServerActionAPI.md#SetVirtualServerType) | **Post** /v1/servers/{server_id}/server-type | Update Virtual Server type
[**ShowVirtualServerConsoleLog**](VirtualserverV1ServerActionAPI.md#ShowVirtualServerConsoleLog) | **Get** /v1/servers/{server_id}/console-log | Show Virtual Server Console Log
[**StartVirtualServer**](VirtualserverV1ServerActionAPI.md#StartVirtualServer) | **Post** /v1/servers/{server_id}/start | Start Virtual Server
[**StopVirtualServer**](VirtualserverV1ServerActionAPI.md#StopVirtualServer) | **Post** /v1/servers/{server_id}/stop | Stop Virtual Server
[**UnlockVirtualServer**](VirtualserverV1ServerActionAPI.md#UnlockVirtualServer) | **Post** /v1/servers/{server_id}/unlock | Unlock Virtual Server



## AttachVirtualServerSecurityGroup

> AttachVirtualServerSecurityGroup(ctx, serverId).ServerSecurityGroupActionRequestBody(serverSecurityGroupActionRequestBody).Execute()

Attach Virtual Server Security Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID
	serverSecurityGroupActionRequestBody := *openapiclient.NewServerSecurityGroupActionRequestBody("c09c3f05-03d9-443f-b27a-40e0f973c75f") // ServerSecurityGroupActionRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServerActionAPI.AttachVirtualServerSecurityGroup(context.Background(), serverId).ServerSecurityGroupActionRequestBody(serverSecurityGroupActionRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.AttachVirtualServerSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachVirtualServerSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverSecurityGroupActionRequestBody** | [**ServerSecurityGroupActionRequestBody**](ServerSecurityGroupActionRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVirtualServerCustomImage

> ServerCreateImageResponse CreateVirtualServerCustomImage(ctx, serverId).ServerCreateImageRequestBody(serverCreateImageRequestBody).Execute()

Create Virtual Server Custom Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID
	serverCreateImageRequestBody := *openapiclient.NewServerCreateImageRequestBody("imagename") // ServerCreateImageRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServerActionAPI.CreateVirtualServerCustomImage(context.Background(), serverId).ServerCreateImageRequestBody(serverCreateImageRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.CreateVirtualServerCustomImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVirtualServerCustomImage`: ServerCreateImageResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServerActionAPI.CreateVirtualServerCustomImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualServerCustomImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverCreateImageRequestBody** | [**ServerCreateImageRequestBody**](ServerCreateImageRequestBody.md) |  | 

### Return type

[**ServerCreateImageResponse**](ServerCreateImageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVirtualServerDump

> CreateVirtualServerDump(ctx, serverId).Execute()

Create Virtual Server Dump



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServerActionAPI.CreateVirtualServerDump(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.CreateVirtualServerDump``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualServerDumpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachVirtualServerSecurityGroup

> DetachVirtualServerSecurityGroup(ctx, serverId, securityGroupId).Execute()

Detach Virtual Server Security Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID
	securityGroupId := "c09c3f05-03d9-443f-b27a-40e0f973c75f" // string | Security Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServerActionAPI.DetachVirtualServerSecurityGroup(context.Background(), serverId, securityGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.DetachVirtualServerSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 
**securityGroupId** | **string** | Security Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachVirtualServerSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualServerSecurityGroups

> ServerSecurityGroupListResponse ListVirtualServerSecurityGroups(ctx, serverId).Execute()

List Virtual Server Security Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServerActionAPI.ListVirtualServerSecurityGroups(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.ListVirtualServerSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualServerSecurityGroups`: ServerSecurityGroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServerActionAPI.ListVirtualServerSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualServerSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerSecurityGroupListResponse**](ServerSecurityGroupListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockVirtualServer

> LockVirtualServer(ctx, serverId).Execute()

Lock Virtual Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServerActionAPI.LockVirtualServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.LockVirtualServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootVirtualServer

> RebootVirtualServer(ctx, serverId).ServerRebootRequestBody(serverRebootRequestBody).Execute()

Reboot Virtual Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID
	serverRebootRequestBody := *openapiclient.NewServerRebootRequestBody() // ServerRebootRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServerActionAPI.RebootVirtualServer(context.Background(), serverId).ServerRebootRequestBody(serverRebootRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.RebootVirtualServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverRebootRequestBody** | [**ServerRebootRequestBody**](ServerRebootRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebuildVirtualServer

> RebuildVirtualServer(ctx, serverId).Execute()

Rebuild Virtual Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServerActionAPI.RebuildVirtualServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.RebuildVirtualServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVirtualServerType

> SetVirtualServerType(ctx, serverId).ServerSetServerTypeRequestBody(serverSetServerTypeRequestBody).Execute()

Update Virtual Server type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID
	serverSetServerTypeRequestBody := *openapiclient.NewServerSetServerTypeRequestBody("s1v1m2") // ServerSetServerTypeRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServerActionAPI.SetVirtualServerType(context.Background(), serverId).ServerSetServerTypeRequestBody(serverSetServerTypeRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.SetVirtualServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVirtualServerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverSetServerTypeRequestBody** | [**ServerSetServerTypeRequestBody**](ServerSetServerTypeRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVirtualServerConsoleLog

> ServerConsoleLogResponse ShowVirtualServerConsoleLog(ctx, serverId).LineSize(lineSize).Execute()

Show Virtual Server Console Log



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID
	lineSize := int32(1) // int32 | Number of log lines to get (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServerActionAPI.ShowVirtualServerConsoleLog(context.Background(), serverId).LineSize(lineSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.ShowVirtualServerConsoleLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVirtualServerConsoleLog`: ServerConsoleLogResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServerActionAPI.ShowVirtualServerConsoleLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVirtualServerConsoleLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lineSize** | **int32** | Number of log lines to get | 

### Return type

[**ServerConsoleLogResponse**](ServerConsoleLogResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartVirtualServer

> StartVirtualServer(ctx, serverId).Execute()

Start Virtual Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServerActionAPI.StartVirtualServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.StartVirtualServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopVirtualServer

> StopVirtualServer(ctx, serverId).Execute()

Stop Virtual Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServerActionAPI.StopVirtualServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.StopVirtualServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlockVirtualServer

> UnlockVirtualServer(ctx, serverId).Execute()

Unlock Virtual Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServerActionAPI.UnlockVirtualServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerActionAPI.UnlockVirtualServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

