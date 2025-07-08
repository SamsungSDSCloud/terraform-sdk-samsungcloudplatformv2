# \BaremetalBlockstorageV1VolumeReplicationV1APIsAPI

All URIs are relative to *http://70.50.166.119:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolumeReplication**](BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.md#CreateVolumeReplication) | **Post** /v1/volumes/{volume_id}/replications | Create Volume Replication
[**DeleteVolumeReplication**](BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.md#DeleteVolumeReplication) | **Delete** /v1/volumes/{volume_id}/replications | Delete Volume Replication
[**SetVolumeReplicationCycle**](BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.md#SetVolumeReplicationCycle) | **Put** /v1/volumes/{volume_id}/replications/cycle | Set Volume Replication Cycle
[**SetVolumeReplicationPolicy**](BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.md#SetVolumeReplicationPolicy) | **Put** /v1/volumes/{volume_id}/replications/policy | Set Volume Replication Policy
[**ShowVolumeReplication**](BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.md#ShowVolumeReplication) | **Get** /v1/volumes/{volume_id}/replications | Show Volume Replication



## CreateVolumeReplication

> AsyncResponse CreateVolumeReplication(ctx, volumeId).VolumeReplicationCreateRequest(volumeReplicationCreateRequest).Execute()

Create Volume Replication



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id
	volumeReplicationCreateRequest := *openapiclient.NewVolumeReplicationCreateRequest(openapiclient.ReplicationCycle("5MIN"), "replica-bs-01", "kr-west1") // VolumeReplicationCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.CreateVolumeReplication(context.Background(), volumeId).VolumeReplicationCreateRequest(volumeReplicationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.CreateVolumeReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolumeReplication`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.CreateVolumeReplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeReplicationCreateRequest** | [**VolumeReplicationCreateRequest**](VolumeReplicationCreateRequest.md) |  | 

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


## DeleteVolumeReplication

> AsyncResponse DeleteVolumeReplication(ctx, volumeId).Execute()

Delete Volume Replication



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.DeleteVolumeReplication(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.DeleteVolumeReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolumeReplication`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.DeleteVolumeReplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVolumeReplicationCycle

> VolumeReplicationCycleResponse SetVolumeReplicationCycle(ctx, volumeId).VolumeReplicationCycleRequest(volumeReplicationCycleRequest).Execute()

Set Volume Replication Cycle



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id
	volumeReplicationCycleRequest := *openapiclient.NewVolumeReplicationCycleRequest(openapiclient.ReplicationCycle("5MIN")) // VolumeReplicationCycleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.SetVolumeReplicationCycle(context.Background(), volumeId).VolumeReplicationCycleRequest(volumeReplicationCycleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.SetVolumeReplicationCycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVolumeReplicationCycle`: VolumeReplicationCycleResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.SetVolumeReplicationCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVolumeReplicationCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeReplicationCycleRequest** | [**VolumeReplicationCycleRequest**](VolumeReplicationCycleRequest.md) |  | 

### Return type

[**VolumeReplicationCycleResponse**](VolumeReplicationCycleResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVolumeReplicationPolicy

> VolumeReplicationPolicyResponse SetVolumeReplicationPolicy(ctx, volumeId).VolumeReplicationPolicyRequest(volumeReplicationPolicyRequest).Execute()

Set Volume Replication Policy



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id
	volumeReplicationPolicyRequest := *openapiclient.NewVolumeReplicationPolicyRequest(openapiclient.ReplicationPolicyMode("RESYNC")) // VolumeReplicationPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.SetVolumeReplicationPolicy(context.Background(), volumeId).VolumeReplicationPolicyRequest(volumeReplicationPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.SetVolumeReplicationPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVolumeReplicationPolicy`: VolumeReplicationPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.SetVolumeReplicationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVolumeReplicationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeReplicationPolicyRequest** | [**VolumeReplicationPolicyRequest**](VolumeReplicationPolicyRequest.md) |  | 

### Return type

[**VolumeReplicationPolicyResponse**](VolumeReplicationPolicyResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVolumeReplication

> VolumeReplicationResponse ShowVolumeReplication(ctx, volumeId).Execute()

Show Volume Replication



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.ShowVolumeReplication(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.ShowVolumeReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVolumeReplication`: VolumeReplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeReplicationV1APIsAPI.ShowVolumeReplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVolumeReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeReplicationResponse**](VolumeReplicationResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

