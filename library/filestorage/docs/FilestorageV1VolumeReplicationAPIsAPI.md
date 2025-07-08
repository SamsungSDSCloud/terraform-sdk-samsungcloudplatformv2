# \FilestorageV1VolumeReplicationAPIsAPI

All URIs are relative to *https://filestorage-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolumeReplication**](FilestorageV1VolumeReplicationAPIsAPI.md#CreateVolumeReplication) | **Post** /v1/replications | CreateVolumeReplication
[**DeleteVolumeReplication**](FilestorageV1VolumeReplicationAPIsAPI.md#DeleteVolumeReplication) | **Delete** /v1/replications/{replication_id} | DeleteVolumeReplication
[**ListVolumeReplicationRegion**](FilestorageV1VolumeReplicationAPIsAPI.md#ListVolumeReplicationRegion) | **Get** /v1/replications/regions | ListVolumeReplicationRegion
[**ListVolumeReplications**](FilestorageV1VolumeReplicationAPIsAPI.md#ListVolumeReplications) | **Get** /v1/replications | ListVolumeReplications
[**SetVolumeReplication**](FilestorageV1VolumeReplicationAPIsAPI.md#SetVolumeReplication) | **Put** /v1/replications/{replication_id} | SetVolumeReplication
[**ShowVolumeReplication**](FilestorageV1VolumeReplicationAPIsAPI.md#ShowVolumeReplication) | **Get** /v1/replications/{replication_id} | ShowVolumeReplication



## CreateVolumeReplication

> ReplicationCreateResponse CreateVolumeReplication(ctx).ReplicationCreateRequest(replicationCreateRequest).Execute()

CreateVolumeReplication



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
	replicationCreateRequest := *openapiclient.NewReplicationCreateRequest("my_volume", "kr-west1", "5min", "replication", "bfdbabf2-04d9-4e8b-a205-020f8e6da438") // ReplicationCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1VolumeReplicationAPIsAPI.CreateVolumeReplication(context.Background()).ReplicationCreateRequest(replicationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeReplicationAPIsAPI.CreateVolumeReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolumeReplication`: ReplicationCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1VolumeReplicationAPIsAPI.CreateVolumeReplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replicationCreateRequest** | [**ReplicationCreateRequest**](ReplicationCreateRequest.md) |  | 

### Return type

[**ReplicationCreateResponse**](ReplicationCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolumeReplication

> DeleteVolumeReplication(ctx, replicationId).VolumeId(volumeId).Execute()

DeleteVolumeReplication



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
	replicationId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Replication ID
	volumeId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilestorageV1VolumeReplicationAPIsAPI.DeleteVolumeReplication(context.Background(), replicationId).VolumeId(volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeReplicationAPIsAPI.DeleteVolumeReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicationId** | **string** | Replication ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeId** | **string** | Volume ID | 

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


## ListVolumeReplicationRegion

> ReplicationRegionListResponse ListVolumeReplicationRegion(ctx).TypeName(typeName).SourceRegionName(sourceRegionName).ReplicationType(replicationType).Execute()

ListVolumeReplicationRegion



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
	typeName := "HDD" // string | Volume Type Name
	sourceRegionName := "kr-west1" // string | Region
	replicationType := "replication" // string | Replication Type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1VolumeReplicationAPIsAPI.ListVolumeReplicationRegion(context.Background()).TypeName(typeName).SourceRegionName(sourceRegionName).ReplicationType(replicationType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeReplicationAPIsAPI.ListVolumeReplicationRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumeReplicationRegion`: ReplicationRegionListResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1VolumeReplicationAPIsAPI.ListVolumeReplicationRegion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeReplicationRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typeName** | **string** | Volume Type Name | 
 **sourceRegionName** | **string** | Region | 
 **replicationType** | **string** | Replication Type | 

### Return type

[**ReplicationRegionListResponse**](ReplicationRegionListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumeReplications

> ReplicationListResponse ListVolumeReplications(ctx).VolumeId(volumeId).Execute()

ListVolumeReplications



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
	volumeId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1VolumeReplicationAPIsAPI.ListVolumeReplications(context.Background()).VolumeId(volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeReplicationAPIsAPI.ListVolumeReplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumeReplications`: ReplicationListResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1VolumeReplicationAPIsAPI.ListVolumeReplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeReplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | **string** | Volume ID | 

### Return type

[**ReplicationListResponse**](ReplicationListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVolumeReplication

> SetVolumeReplication(ctx, replicationId).VolumeId(volumeId).ReplicationUpdateRequest(replicationUpdateRequest).Execute()

SetVolumeReplication



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
	replicationId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Replication ID
	volumeId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Volume ID
	replicationUpdateRequest := *openapiclient.NewReplicationUpdateRequest("policy") // ReplicationUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilestorageV1VolumeReplicationAPIsAPI.SetVolumeReplication(context.Background(), replicationId).VolumeId(volumeId).ReplicationUpdateRequest(replicationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeReplicationAPIsAPI.SetVolumeReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicationId** | **string** | Replication ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVolumeReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeId** | **string** | Volume ID | 
 **replicationUpdateRequest** | [**ReplicationUpdateRequest**](ReplicationUpdateRequest.md) |  | 

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


## ShowVolumeReplication

> ReplicationShowResponse ShowVolumeReplication(ctx, replicationId).VolumeId(volumeId).Execute()

ShowVolumeReplication



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
	replicationId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Replication ID
	volumeId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1VolumeReplicationAPIsAPI.ShowVolumeReplication(context.Background(), replicationId).VolumeId(volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeReplicationAPIsAPI.ShowVolumeReplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVolumeReplication`: ReplicationShowResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1VolumeReplicationAPIsAPI.ShowVolumeReplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicationId** | **string** | Replication ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVolumeReplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeId** | **string** | Volume ID | 

### Return type

[**ReplicationShowResponse**](ReplicationShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

