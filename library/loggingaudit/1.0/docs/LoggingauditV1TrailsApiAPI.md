# \LoggingauditV1TrailsApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrail**](LoggingauditV1TrailsApiAPI.md#CreateTrail) | **Post** /v1/trails | CreateTrail
[**DeleteTrail**](LoggingauditV1TrailsApiAPI.md#DeleteTrail) | **Delete** /v1/trails/{trail_id} | DeleteTrail
[**ListTrails**](LoggingauditV1TrailsApiAPI.md#ListTrails) | **Get** /v1/trails | listTrails
[**SetTrail**](LoggingauditV1TrailsApiAPI.md#SetTrail) | **Put** /v1/trails/{trail_id} | SetTrail
[**ShowTrail**](LoggingauditV1TrailsApiAPI.md#ShowTrail) | **Get** /v1/trails/{trail_id} | ShowTrail
[**StartTrail**](LoggingauditV1TrailsApiAPI.md#StartTrail) | **Post** /v1/trails/{trail_id}/start | StartTrail
[**StopTrail**](LoggingauditV1TrailsApiAPI.md#StopTrail) | **Post** /v1/trails/{trail_id}/stop | StopTrail



## CreateTrail

> TrailShowResponse CreateTrail(ctx).TrailCreateRequest(trailCreateRequest).Execute()

CreateTrail



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
	trailCreateRequest := *openapiclient.NewTrailCreateRequest("TestBucket00", "RegionOne", "TestTrail01", "JSON") // TrailCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoggingauditV1TrailsApiAPI.CreateTrail(context.Background()).TrailCreateRequest(trailCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoggingauditV1TrailsApiAPI.CreateTrail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTrail`: TrailShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoggingauditV1TrailsApiAPI.CreateTrail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trailCreateRequest** | [**TrailCreateRequest**](TrailCreateRequest.md) |  | 

### Return type

[**TrailShowResponse**](TrailShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrail

> DeleteTrail(ctx, trailId).Execute()

DeleteTrail



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
	trailId := "trailId_example" // string | Trail ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoggingauditV1TrailsApiAPI.DeleteTrail(context.Background(), trailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoggingauditV1TrailsApiAPI.DeleteTrail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trailId** | **string** | Trail ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrailRequest struct via the builder pattern


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


## ListTrails

> TrailListResponse ListTrails(ctx).Size(size).Page(page).Sort(sort).TrailName(trailName).BucketName(bucketName).State(state).ResourceType(resourceType).Execute()

listTrails



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
	trailName := "trailName_example" // string | Trail name (optional)
	bucketName := "bucketName_example" // string | Bucket name (optional)
	state := "state_example" // string | Trail state (optional)
	resourceType := "resourceType_example" // string | Resource type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoggingauditV1TrailsApiAPI.ListTrails(context.Background()).Size(size).Page(page).Sort(sort).TrailName(trailName).BucketName(bucketName).State(state).ResourceType(resourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoggingauditV1TrailsApiAPI.ListTrails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrails`: TrailListResponse
	fmt.Fprintf(os.Stdout, "Response from `LoggingauditV1TrailsApiAPI.ListTrails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **trailName** | **string** | Trail name | 
 **bucketName** | **string** | Bucket name | 
 **state** | **string** | Trail state | 
 **resourceType** | **string** | Resource type | 

### Return type

[**TrailListResponse**](TrailListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTrail

> TrailShowResponse SetTrail(ctx, trailId).TrailSetRequest(trailSetRequest).Execute()

SetTrail



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
	trailId := "trailId_example" // string | Trail ID
	trailSetRequest := *openapiclient.NewTrailSetRequest() // TrailSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoggingauditV1TrailsApiAPI.SetTrail(context.Background(), trailId).TrailSetRequest(trailSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoggingauditV1TrailsApiAPI.SetTrail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTrail`: TrailShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoggingauditV1TrailsApiAPI.SetTrail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trailId** | **string** | Trail ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetTrailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trailSetRequest** | [**TrailSetRequest**](TrailSetRequest.md) |  | 

### Return type

[**TrailShowResponse**](TrailShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTrail

> TrailShowResponse ShowTrail(ctx, trailId).Execute()

ShowTrail



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
	trailId := "trailId_example" // string | Trail ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoggingauditV1TrailsApiAPI.ShowTrail(context.Background(), trailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoggingauditV1TrailsApiAPI.ShowTrail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowTrail`: TrailShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoggingauditV1TrailsApiAPI.ShowTrail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trailId** | **string** | Trail ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowTrailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrailShowResponse**](TrailShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartTrail

> TrailShowResponse StartTrail(ctx, trailId).Execute()

StartTrail



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
	trailId := "trailId_example" // string | Trail ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoggingauditV1TrailsApiAPI.StartTrail(context.Background(), trailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoggingauditV1TrailsApiAPI.StartTrail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartTrail`: TrailShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoggingauditV1TrailsApiAPI.StartTrail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trailId** | **string** | Trail ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartTrailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrailShowResponse**](TrailShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopTrail

> TrailShowResponse StopTrail(ctx, trailId).Execute()

StopTrail



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
	trailId := "trailId_example" // string | Trail ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoggingauditV1TrailsApiAPI.StopTrail(context.Background(), trailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoggingauditV1TrailsApiAPI.StopTrail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopTrail`: TrailShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoggingauditV1TrailsApiAPI.StopTrail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trailId** | **string** | Trail ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopTrailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrailShowResponse**](TrailShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

