# \VirtualserverV1VolumeServersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachVolumeToVirtualServer**](VirtualserverV1VolumeServersApiAPI.md#AttachVolumeToVirtualServer) | **Post** /v1/volumes/{volume_id}/servers | Attach Volume to Virtual Server
[**DetachVolumeFromVirtualServer**](VirtualserverV1VolumeServersApiAPI.md#DetachVolumeFromVirtualServer) | **Delete** /v1/volumes/{volume_id}/servers/{server_id} | Detach Volume from Virtual Server



## AttachVolumeToVirtualServer

> VolumeServerAttachResponse AttachVolumeToVirtualServer(ctx, volumeId).VolumeServerAttachRequest(volumeServerAttachRequest).Execute()

Attach Volume to Virtual Server



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
	volumeId := "760fee21-42b7-4717-adb7-511b31f8af74" // string | Volume ID
	volumeServerAttachRequest := *openapiclient.NewVolumeServerAttachRequest("b13531ca-99ec-4aa5-a2dd-a85665b0a358") // VolumeServerAttachRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumeServersApiAPI.AttachVolumeToVirtualServer(context.Background(), volumeId).VolumeServerAttachRequest(volumeServerAttachRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumeServersApiAPI.AttachVolumeToVirtualServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachVolumeToVirtualServer`: VolumeServerAttachResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumeServersApiAPI.AttachVolumeToVirtualServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachVolumeToVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeServerAttachRequest** | [**VolumeServerAttachRequest**](VolumeServerAttachRequest.md) |  | 

### Return type

[**VolumeServerAttachResponse**](VolumeServerAttachResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachVolumeFromVirtualServer

> DetachVolumeFromVirtualServer(ctx, volumeId, serverId).Execute()

Detach Volume from Virtual Server



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
	volumeId := "760fee21-42b7-4717-adb7-511b31f8af74" // string | Volume ID
	serverId := "b13531ca-99ec-4aa5-a2dd-a85665b0a358" // string | Server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1VolumeServersApiAPI.DetachVolumeFromVirtualServer(context.Background(), volumeId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumeServersApiAPI.DetachVolumeFromVirtualServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume ID | 
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachVolumeFromVirtualServerRequest struct via the builder pattern


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

