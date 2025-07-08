# \SkeV1CronJobsApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCronJob**](SkeV1CronJobsApiAPI.md#DeleteCronJob) | **Delete** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/cronjobs | Delete Cron Job
[**ListCronJobs**](SkeV1CronJobsApiAPI.md#ListCronJobs) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/cronjobs | List Cron Jobs
[**ShowCronJob**](SkeV1CronJobsApiAPI.md#ShowCronJob) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/cronjobs/{cron_job_name} | Show Cron Job



## DeleteCronJob

> DeleteCronJob(ctx, clusterId, namespaceName).Name(name).Execute()

Delete Cron Job



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
	name := *openapiclient.NewName1() // Name1 | Names

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SkeV1CronJobsApiAPI.DeleteCronJob(context.Background(), clusterId, namespaceName).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1CronJobsApiAPI.DeleteCronJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | [**Name1**](Name1.md) | Names | 

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


## ListCronJobs

> CronJobListResponse ListCronJobs(ctx, clusterId, namespaceName).Size(size).Page(page).Sort(sort).Name(name).SystemObject(systemObject).Execute()

List Cron Jobs



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
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	name := "name_example" // string | Cron Job Name (optional)
	systemObject := true // bool | System Object (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1CronJobsApiAPI.ListCronJobs(context.Background(), clusterId, namespaceName).Size(size).Page(page).Sort(sort).Name(name).SystemObject(systemObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1CronJobsApiAPI.ListCronJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCronJobs`: CronJobListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1CronJobsApiAPI.ListCronJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCronJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Cron Job Name | 
 **systemObject** | **bool** | System Object | 

### Return type

[**CronJobListResponse**](CronJobListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowCronJob

> CronJobShowResponse ShowCronJob(ctx, clusterId, namespaceName, cronJobName).Execute()

Show Cron Job



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
	cronJobName := "cronJobName_example" // string | Cron Job Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1CronJobsApiAPI.ShowCronJob(context.Background(), clusterId, namespaceName, cronJobName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1CronJobsApiAPI.ShowCronJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowCronJob`: CronJobShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1CronJobsApiAPI.ShowCronJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 
**cronJobName** | **string** | Cron Job Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CronJobShowResponse**](CronJobShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

