# \BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI

All URIs are relative to *http://70.50.166.119:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolumeGroupReplication**](BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.md#CreateVolumeGroupReplication) | **Post** /v1/volume-groups/{volume_group_id}/replications | Create Volume Group Replication
[**DeleteVolumeGroupReplication**](BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.md#DeleteVolumeGroupReplication) | **Delete** /v1/volume-groups/{volume_group_id}/replications | Delete Volume Group Replication
[**SetVolumeGroupReplicationCycle**](BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.md#SetVolumeGroupReplicationCycle) | **Put** /v1/volume-groups/{volume_group_id}/replications/cycle | Set Volume Group Replication Cycle
[**SetVolumeGroupReplicationPolicy**](BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.md#SetVolumeGroupReplicationPolicy) | **Put** /v1/volume-groups/{volume_group_id}/replications/policy | Set Volume Group Replication Policy
[**ShowVolumeGroupReplication**](BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.md#ShowVolumeGroupReplication) | **Get** /v1/volume-groups/{volume_group_id}/replications | Show Volume Group Replication



## CreateVolumeGroupReplication

> AsyncResponse CreateVolumeGroupReplication(ctx, volumeGroupId).VolumeGroupReplicationCreateRequest(volumeGroupReplicationCreateRequest).Execute()

Create Volume Group Replication



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id
	volumeGroupReplicationCreateRequest := *openapiclient.NewVolumeGroupReplicationCreateRequest(openapiclient.ReplicationCycle("5MIN"), "replica-vg-01", "kr-west2", "dr") // VolumeGroupReplicationCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.CreateVolumeGroupReplication(context.Background(), volumeGroupId).VolumeGroupReplicationCreateRequest(volumeGroupReplicationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.CreateVolumeGroupReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolumeGroupReplication`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.CreateVolumeGroupReplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeGroupReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeGroupReplicationCreateRequest** | [**VolumeGroupReplicationCreateRequest**](VolumeGroupReplicationCreateRequest.md) |  | 

### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolumeGroupReplication

> VolumeGroupReplicationDeleteResponse DeleteVolumeGroupReplication(ctx, volumeGroupId).Execute()

Delete Volume Group Replication



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.DeleteVolumeGroupReplication(context.Background(), volumeGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.DeleteVolumeGroupReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolumeGroupReplication`: VolumeGroupReplicationDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.DeleteVolumeGroupReplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeGroupReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeGroupReplicationDeleteResponse**](VolumeGroupReplicationDeleteResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVolumeGroupReplicationCycle

> VolumeGroupReplicationUpdateCycleResponse SetVolumeGroupReplicationCycle(ctx, volumeGroupId).VolumeGroupReplicationUpdateCycleRequest(volumeGroupReplicationUpdateCycleRequest).Execute()

Set Volume Group Replication Cycle



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id
	volumeGroupReplicationUpdateCycleRequest := *openapiclient.NewVolumeGroupReplicationUpdateCycleRequest(openapiclient.ReplicationCycle("5MIN")) // VolumeGroupReplicationUpdateCycleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.SetVolumeGroupReplicationCycle(context.Background(), volumeGroupId).VolumeGroupReplicationUpdateCycleRequest(volumeGroupReplicationUpdateCycleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.SetVolumeGroupReplicationCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVolumeGroupReplicationCycle`: VolumeGroupReplicationUpdateCycleResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.SetVolumeGroupReplicationCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVolumeGroupReplicationCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeGroupReplicationUpdateCycleRequest** | [**VolumeGroupReplicationUpdateCycleRequest**](VolumeGroupReplicationUpdateCycleRequest.md) |  | 

### Return type

[**VolumeGroupReplicationUpdateCycleResponse**](VolumeGroupReplicationUpdateCycleResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVolumeGroupReplicationPolicy

> VolumeGroupReplicationUpdatePolicyResponse SetVolumeGroupReplicationPolicy(ctx, volumeGroupId).VolumeGroupReplicationUpdatePolicyRequest(volumeGroupReplicationUpdatePolicyRequest).Execute()

Set Volume Group Replication Policy



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id
	volumeGroupReplicationUpdatePolicyRequest := *openapiclient.NewVolumeGroupReplicationUpdatePolicyRequest(openapiclient.ReplicationPolicyMode("RESYNC")) // VolumeGroupReplicationUpdatePolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.SetVolumeGroupReplicationPolicy(context.Background(), volumeGroupId).VolumeGroupReplicationUpdatePolicyRequest(volumeGroupReplicationUpdatePolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.SetVolumeGroupReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVolumeGroupReplicationPolicy`: VolumeGroupReplicationUpdatePolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.SetVolumeGroupReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVolumeGroupReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeGroupReplicationUpdatePolicyRequest** | [**VolumeGroupReplicationUpdatePolicyRequest**](VolumeGroupReplicationUpdatePolicyRequest.md) |  | 

### Return type

[**VolumeGroupReplicationUpdatePolicyResponse**](VolumeGroupReplicationUpdatePolicyResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVolumeGroupReplication

> VolumeGroupReplicationSyncTabResponse ShowVolumeGroupReplication(ctx, volumeGroupId).Execute()

Show Volume Group Replication



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.ShowVolumeGroupReplication(context.Background(), volumeGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.ShowVolumeGroupReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVolumeGroupReplication`: VolumeGroupReplicationSyncTabResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupReplicationV1APIsAPI.ShowVolumeGroupReplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVolumeGroupReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeGroupReplicationSyncTabResponse**](VolumeGroupReplicationSyncTabResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

