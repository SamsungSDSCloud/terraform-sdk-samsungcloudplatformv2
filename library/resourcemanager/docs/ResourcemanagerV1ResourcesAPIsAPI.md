# \ResourcemanagerV1ResourcesAPIsAPI

All URIs are relative to *https://resourcemanager.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowResource**](ResourcemanagerV1ResourcesAPIsAPI.md#ShowResource) | **Get** /v1/resources/{srn} | ShowResource
[**ShowResourceByComponents**](ResourcemanagerV1ResourcesAPIsAPI.md#ShowResourceByComponents) | **Get** /v1/resources/{region}/{service_type}/{resource_type}/{resource_identifier} | ShowResourceByComponents



## ShowResource

> ResourceShowResponse ShowResource(ctx, srn).Execute()

ShowResource



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
	srn := "srn_example" // string | SRN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1ResourcesAPIsAPI.ShowResource(context.Background(), srn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1ResourcesAPIsAPI.ShowResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowResource`: ResourceShowResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1ResourcesAPIsAPI.ShowResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srn** | **string** | SRN | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceShowResponse**](ResourceShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowResourceByComponents

> ResourceShowResponse ShowResourceByComponents(ctx, region, serviceType, resourceType, resourceIdentifier).Execute()

ShowResourceByComponents



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
	region := "region_example" // string | 리전
	serviceType := "serviceType_example" // string | 서비스 유형
	resourceType := "resourceType_example" // string | 자원 유형
	resourceIdentifier := "resourceIdentifier_example" // string | 자원 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1ResourcesAPIsAPI.ShowResourceByComponents(context.Background(), region, serviceType, resourceType, resourceIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1ResourcesAPIsAPI.ShowResourceByComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowResourceByComponents`: ResourceShowResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1ResourcesAPIsAPI.ShowResourceByComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | 리전 | 
**serviceType** | **string** | 서비스 유형 | 
**resourceType** | **string** | 자원 유형 | 
**resourceIdentifier** | **string** | 자원 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowResourceByComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ResourceShowResponse**](ResourceShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

