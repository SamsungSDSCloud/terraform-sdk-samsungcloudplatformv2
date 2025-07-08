# \SkeV1JobsApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteJob**](SkeV1JobsApiAPI.md#DeleteJob) | **Delete** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/jobs | Delete Job
[**ListJobs**](SkeV1JobsApiAPI.md#ListJobs) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/jobs | List Jobs
[**ShowJob**](SkeV1JobsApiAPI.md#ShowJob) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/jobs/{job_name} | Show Job



## DeleteJob

> DeleteJob(ctx, clusterId, namespaceName).Execute()

Delete Job



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
	clusterId := "clusterId_example" // string | Cluster ID
	namespaceName := "namespaceName_example" // string | Namespace Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SkeV1JobsApiAPI.DeleteJob(context.Background(), clusterId, namespaceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1JobsApiAPI.DeleteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobRequest struct via the builder pattern


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


## ListJobs

> JobListResponse ListJobs(ctx, clusterId, namespaceName).Execute()

List Jobs



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
	clusterId := "clusterId_example" // string | Cluster ID
	namespaceName := "namespaceName_example" // string | Namespace Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1JobsApiAPI.ListJobs(context.Background(), clusterId, namespaceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1JobsApiAPI.ListJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListJobs`: JobListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1JobsApiAPI.ListJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**JobListResponse**](JobListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowJob

> JobShowResponse ShowJob(ctx, clusterId, namespaceName, jobName).Execute()

Show Job



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
	clusterId := "clusterId_example" // string | Cluster ID
	namespaceName := "namespaceName_example" // string | Namespace Name
	jobName := "jobName_example" // string | Job Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1JobsApiAPI.ShowJob(context.Background(), clusterId, namespaceName, jobName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1JobsApiAPI.ShowJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowJob`: JobShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1JobsApiAPI.ShowJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 
**jobName** | **string** | Job Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**JobShowResponse**](JobShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

