# \LoadbalancerV1LBHealthCheckApiAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLbHealthCheck**](LoadbalancerV1LBHealthCheckApiAPI.md#CreateLbHealthCheck) | **Post** /v1/lb-health-checks | CreateLBHealthCheck
[**DeleteLbHealthCheck**](LoadbalancerV1LBHealthCheckApiAPI.md#DeleteLbHealthCheck) | **Delete** /v1/lb-health-checks/{lb_health_check_id} | DeleteLBHealthCheck
[**ListLbHealthChecks**](LoadbalancerV1LBHealthCheckApiAPI.md#ListLbHealthChecks) | **Get** /v1/lb-health-checks | ListLBHealthChecks
[**SetLbHealthCheck**](LoadbalancerV1LBHealthCheckApiAPI.md#SetLbHealthCheck) | **Put** /v1/lb-health-checks/{lb_health_check_id} | SetLBHealthCheck
[**ShowLbHealthCheck**](LoadbalancerV1LBHealthCheckApiAPI.md#ShowLbHealthCheck) | **Get** /v1/lb-health-checks/{lb_health_check_id} | ShowLBHealthCheck



## CreateLbHealthCheck

> LbHealthCheckShowResponse CreateLbHealthCheck(ctx).LbHealthCheckCreateRequest(lbHealthCheckCreateRequest).Execute()

CreateLBHealthCheck



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
	lbHealthCheckCreateRequest := *openapiclient.NewLbHealthCheckCreateRequest(*openapiclient.NewLbHealthCheckCreate("ServerGroup01", openapiclient.LbMonitorProtocol("TCP"), "60fba45cb6c811efba41ba92e4fe7200", "8acceeb6920c4fc494490d864f67f0b5")) // LbHealthCheckCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LBHealthCheckApiAPI.CreateLbHealthCheck(context.Background()).LbHealthCheckCreateRequest(lbHealthCheckCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LBHealthCheckApiAPI.CreateLbHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLbHealthCheck`: LbHealthCheckShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LBHealthCheckApiAPI.CreateLbHealthCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLbHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lbHealthCheckCreateRequest** | [**LbHealthCheckCreateRequest**](LbHealthCheckCreateRequest.md) |  | 

### Return type

[**LbHealthCheckShowResponse**](LbHealthCheckShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLbHealthCheck

> DeleteLbHealthCheck(ctx, lbHealthCheckId).Execute()

DeleteLBHealthCheck



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
	lbHealthCheckId := "e3cd678b11784734bc366148aa37580e" // string | LB Health Check ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadbalancerV1LBHealthCheckApiAPI.DeleteLbHealthCheck(context.Background(), lbHealthCheckId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LBHealthCheckApiAPI.DeleteLbHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbHealthCheckId** | **string** | LB Health Check ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLbHealthCheckRequest struct via the builder pattern


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


## ListLbHealthChecks

> LbHealthCheckListResponse ListLbHealthChecks(ctx).Size(size).Page(page).Sort(sort).Name(name).Protocol(protocol).SubnetId(subnetId).Execute()

ListLBHealthChecks



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
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	name := "Test-LbMonitor-01" // string | LB Health Check name (optional)
	protocol := *openapiclient.NewProtocol() // Protocol | LB Health Check Protocol (optional)
	subnetId := "60fba45cb6c811efba41ba92e4fe7200" // string | Service Subnet ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LBHealthCheckApiAPI.ListLbHealthChecks(context.Background()).Size(size).Page(page).Sort(sort).Name(name).Protocol(protocol).SubnetId(subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LBHealthCheckApiAPI.ListLbHealthChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLbHealthChecks`: LbHealthCheckListResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LBHealthCheckApiAPI.ListLbHealthChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLbHealthChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | LB Health Check name | 
 **protocol** | [**Protocol**](Protocol.md) | LB Health Check Protocol | 
 **subnetId** | **string** | Service Subnet ID | 

### Return type

[**LbHealthCheckListResponse**](LbHealthCheckListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLbHealthCheck

> LbHealthCheckShowResponse SetLbHealthCheck(ctx, lbHealthCheckId).LbHealthCheckSetRequest(lbHealthCheckSetRequest).Execute()

SetLBHealthCheck



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
	lbHealthCheckId := "e3cd678b11784734bc366148aa37580e" // string | LB Health Check ID
	lbHealthCheckSetRequest := *openapiclient.NewLbHealthCheckSetRequest(*openapiclient.NewLbHealthCheckSet()) // LbHealthCheckSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LBHealthCheckApiAPI.SetLbHealthCheck(context.Background(), lbHealthCheckId).LbHealthCheckSetRequest(lbHealthCheckSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LBHealthCheckApiAPI.SetLbHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetLbHealthCheck`: LbHealthCheckShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LBHealthCheckApiAPI.SetLbHealthCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbHealthCheckId** | **string** | LB Health Check ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLbHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lbHealthCheckSetRequest** | [**LbHealthCheckSetRequest**](LbHealthCheckSetRequest.md) |  | 

### Return type

[**LbHealthCheckShowResponse**](LbHealthCheckShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowLbHealthCheck

> LbHealthCheckShowResponse ShowLbHealthCheck(ctx, lbHealthCheckId).Execute()

ShowLBHealthCheck



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
	lbHealthCheckId := "e3cd678b11784734bc366148aa37580e" // string | LB Health Check ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LBHealthCheckApiAPI.ShowLbHealthCheck(context.Background(), lbHealthCheckId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LBHealthCheckApiAPI.ShowLbHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowLbHealthCheck`: LbHealthCheckShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LBHealthCheckApiAPI.ShowLbHealthCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbHealthCheckId** | **string** | LB Health Check ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowLbHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LbHealthCheckShowResponse**](LbHealthCheckShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

