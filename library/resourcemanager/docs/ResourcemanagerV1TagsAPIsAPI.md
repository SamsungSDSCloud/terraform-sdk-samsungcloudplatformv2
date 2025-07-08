# \ResourcemanagerV1TagsAPIsAPI

All URIs are relative to *https://resourcemanager.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteComponentsTag**](ResourcemanagerV1TagsAPIsAPI.md#DeleteComponentsTag) | **Delete** /v1/tags/{region}/{service_type}/{resource_type}/{resource_identifier}/{key} | DeleteComponentsTag
[**DeleteComponentsTags**](ResourcemanagerV1TagsAPIsAPI.md#DeleteComponentsTags) | **Delete** /v1/tags/{region}/{service_type}/{resource_type}/{resource_identifier} | DeleteComponentsTags
[**DeleteResourceTag**](ResourcemanagerV1TagsAPIsAPI.md#DeleteResourceTag) | **Delete** /v1/tags/{srn}/{key} | DeleteResourceTag
[**DeleteResourceTags**](ResourcemanagerV1TagsAPIsAPI.md#DeleteResourceTags) | **Delete** /v1/tags/{srn} | DeleteResourceTags
[**DeleteTags**](ResourcemanagerV1TagsAPIsAPI.md#DeleteTags) | **Delete** /v1/tags | DeleteTags
[**ListComponentsTags**](ResourcemanagerV1TagsAPIsAPI.md#ListComponentsTags) | **Get** /v1/tags/{region}/{service_type}/{resource_type}/{resource_identifier} | ListComponentsTags
[**ListResourceTags**](ResourcemanagerV1TagsAPIsAPI.md#ListResourceTags) | **Get** /v1/tags/{srn} | ListResourceTags
[**ListTagKeys**](ResourcemanagerV1TagsAPIsAPI.md#ListTagKeys) | **Get** /v1/tags/keys | ListTagKeys
[**ListTagValues**](ResourcemanagerV1TagsAPIsAPI.md#ListTagValues) | **Get** /v1/tags/values | ListTagValues
[**ListTags**](ResourcemanagerV1TagsAPIsAPI.md#ListTags) | **Get** /v1/tags | ListTags
[**ShowComponentsTag**](ResourcemanagerV1TagsAPIsAPI.md#ShowComponentsTag) | **Get** /v1/tags/{region}/{service_type}/{resource_type}/{resource_identifier}/{key} | ShowComponentsTag
[**ShowResourceTag**](ResourcemanagerV1TagsAPIsAPI.md#ShowResourceTag) | **Get** /v1/tags/{srn}/{key} | ShowResourceTag
[**UpdateComponentsTagValue**](ResourcemanagerV1TagsAPIsAPI.md#UpdateComponentsTagValue) | **Put** /v1/tags/{region}/{service_type}/{resource_type}/{resource_identifier}/{key} | UpdateComponentsTagValue
[**UpdateComponentsTags**](ResourcemanagerV1TagsAPIsAPI.md#UpdateComponentsTags) | **Put** /v1/tags/{region}/{service_type}/{resource_type}/{resource_identifier}/bulk | UpdateComponentsTags
[**UpdateResourceTagValue**](ResourcemanagerV1TagsAPIsAPI.md#UpdateResourceTagValue) | **Put** /v1/tags/{srn}/{key} | UpdateResourceTagValue
[**UpdateResourceTags**](ResourcemanagerV1TagsAPIsAPI.md#UpdateResourceTags) | **Put** /v1/tags/{srn}/bulk | UpdateResourceTags
[**UpdateTags**](ResourcemanagerV1TagsAPIsAPI.md#UpdateTags) | **Put** /v1/tags/bulk | UpdateTags



## DeleteComponentsTag

> DeleteComponentsTag(ctx, region, serviceType, resourceType, resourceIdentifier, key).Execute()

DeleteComponentsTag



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
	key := "key_example" // string | 태그 key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.DeleteComponentsTag(context.Background(), region, serviceType, resourceType, resourceIdentifier, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.DeleteComponentsTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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
**key** | **string** | 태그 key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComponentsTagRequest struct via the builder pattern


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


## DeleteComponentsTags

> DeleteComponentsTags(ctx, region, serviceType, resourceType, resourceIdentifier).Execute()

DeleteComponentsTags



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
	r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.DeleteComponentsTags(context.Background(), region, serviceType, resourceType, resourceIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.DeleteComponentsTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

Other parameters are passed through a pointer to a apiDeleteComponentsTagsRequest struct via the builder pattern


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


## DeleteResourceTag

> DeleteResourceTag(ctx, srn, key).Execute()

DeleteResourceTag



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
	key := "key_example" // string | 태그 key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.DeleteResourceTag(context.Background(), srn, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.DeleteResourceTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srn** | **string** | SRN | 
**key** | **string** | 태그 key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceTagRequest struct via the builder pattern


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


## DeleteResourceTags

> DeleteResourceTags(ctx, srn).Execute()

DeleteResourceTags



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
	r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.DeleteResourceTags(context.Background(), srn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.DeleteResourceTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srn** | **string** | SRN | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceTagsRequest struct via the builder pattern


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


## DeleteTags

> DeleteTags(ctx).TagDeleteRequest(tagDeleteRequest).Execute()

DeleteTags



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
	tagDeleteRequest := *openapiclient.NewTagDeleteRequest(map[string][]string{"key": []string{"Inner_example"}}) // TagDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.DeleteTags(context.Background()).TagDeleteRequest(tagDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.DeleteTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagDeleteRequest** | [**TagDeleteRequest**](TagDeleteRequest.md) |  | 

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


## ListComponentsTags

> TagShowListResponse ListComponentsTags(ctx, region, serviceType, resourceType, resourceIdentifier).Size(size).Page(page).Sort(sort).Execute()

ListComponentsTags



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
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.ListComponentsTags(context.Background(), region, serviceType, resourceType, resourceIdentifier).Size(size).Page(page).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.ListComponentsTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListComponentsTags`: TagShowListResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1TagsAPIsAPI.ListComponentsTags`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListComponentsTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 

### Return type

[**TagShowListResponse**](TagShowListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceTags

> TagShowListResponse ListResourceTags(ctx, srn).Size(size).Page(page).Sort(sort).Execute()

ListResourceTags



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
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.ListResourceTags(context.Background(), srn).Size(size).Page(page).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.ListResourceTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceTags`: TagShowListResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1TagsAPIsAPI.ListResourceTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srn** | **string** | SRN | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResourceTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 

### Return type

[**TagShowListResponse**](TagShowListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTagKeys

> TagKeyResponse ListTagKeys(ctx).Key(key).Execute()

ListTagKeys



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
	key := "key_example" // string | 태그 key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.ListTagKeys(context.Background()).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.ListTagKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTagKeys`: TagKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1TagsAPIsAPI.ListTagKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTagKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | 태그 key | 

### Return type

[**TagKeyResponse**](TagKeyResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTagValues

> TagValueResponse ListTagValues(ctx).Key(key).Value(value).Execute()

ListTagValues



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
	key := "key_example" // string | 태그 key (optional)
	value := "value_example" // string | 태그 value (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.ListTagValues(context.Background()).Key(key).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.ListTagValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTagValues`: TagValueResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1TagsAPIsAPI.ListTagValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTagValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | 태그 key | 
 **value** | **string** | 태그 value | 

### Return type

[**TagValueResponse**](TagValueResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTags

> TagListResponse ListTags(ctx).Size(size).Page(page).Sort(sort).AccountId(accountId).Key(key).Value(value).ResourceIdentifier(resourceIdentifier).ResourceType(resourceType).ServiceType(serviceType).Offering(offering).Region(region).Execute()

ListTags



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
	accountId := "accountId_example" // string | 어카운트 ID (optional)
	key := "key_example" // string | 태그 key (optional)
	value := "value_example" // string | 태그 value (optional)
	resourceIdentifier := "resourceIdentifier_example" // string | 자원 ID (optional)
	resourceType := "resourceType_example" // string | 자원 유형 (optional)
	serviceType := "serviceType_example" // string | 서비스 유형 (optional)
	offering := "offering_example" // string | SCP 운영 단위 (optional)
	region := "region_example" // string | 리전 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.ListTags(context.Background()).Size(size).Page(page).Sort(sort).AccountId(accountId).Key(key).Value(value).ResourceIdentifier(resourceIdentifier).ResourceType(resourceType).ServiceType(serviceType).Offering(offering).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.ListTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTags`: TagListResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1TagsAPIsAPI.ListTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **accountId** | **string** | 어카운트 ID | 
 **key** | **string** | 태그 key | 
 **value** | **string** | 태그 value | 
 **resourceIdentifier** | **string** | 자원 ID | 
 **resourceType** | **string** | 자원 유형 | 
 **serviceType** | **string** | 서비스 유형 | 
 **offering** | **string** | SCP 운영 단위 | 
 **region** | **string** | 리전 | 

### Return type

[**TagListResponse**](TagListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowComponentsTag

> TagShowResponse ShowComponentsTag(ctx, region, serviceType, resourceType, resourceIdentifier, key).Execute()

ShowComponentsTag



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
	key := "key_example" // string | 태그 key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.ShowComponentsTag(context.Background(), region, serviceType, resourceType, resourceIdentifier, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.ShowComponentsTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowComponentsTag`: TagShowResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1TagsAPIsAPI.ShowComponentsTag`: %v\n", resp)
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
**key** | **string** | 태그 key | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowComponentsTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**TagShowResponse**](TagShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowResourceTag

> TagShowResponse ShowResourceTag(ctx, srn, key).Execute()

ShowResourceTag



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
	key := "key_example" // string | 태그 key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.ShowResourceTag(context.Background(), srn, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.ShowResourceTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowResourceTag`: TagShowResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1TagsAPIsAPI.ShowResourceTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srn** | **string** | SRN | 
**key** | **string** | 태그 key | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowResourceTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TagShowResponse**](TagShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComponentsTagValue

> TagBaseResponse UpdateComponentsTagValue(ctx, region, serviceType, resourceType, resourceIdentifier, key).TagValue(tagValue).Execute()

UpdateComponentsTagValue



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
	key := "key_example" // string | 태그 key
	tagValue := *openapiclient.NewTagValue("Value_example") // TagValue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.UpdateComponentsTagValue(context.Background(), region, serviceType, resourceType, resourceIdentifier, key).TagValue(tagValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.UpdateComponentsTagValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateComponentsTagValue`: TagBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1TagsAPIsAPI.UpdateComponentsTagValue`: %v\n", resp)
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
**key** | **string** | 태그 key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComponentsTagValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **tagValue** | [**TagValue**](TagValue.md) |  | 

### Return type

[**TagBaseResponse**](TagBaseResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComponentsTags

> TagBaseResponse UpdateComponentsTags(ctx, region, serviceType, resourceType, resourceIdentifier).TagSetRequest(tagSetRequest).Execute()

UpdateComponentsTags



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
	tagSetRequest := *openapiclient.NewTagSetRequest([]openapiclient.Tag{*openapiclient.NewTag("Key_example", "Value_example")}) // TagSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.UpdateComponentsTags(context.Background(), region, serviceType, resourceType, resourceIdentifier).TagSetRequest(tagSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.UpdateComponentsTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateComponentsTags`: TagBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1TagsAPIsAPI.UpdateComponentsTags`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateComponentsTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tagSetRequest** | [**TagSetRequest**](TagSetRequest.md) |  | 

### Return type

[**TagBaseResponse**](TagBaseResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceTagValue

> TagBaseResponse UpdateResourceTagValue(ctx, srn, key).TagValue(tagValue).Execute()

UpdateResourceTagValue



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
	key := "key_example" // string | 태그 key
	tagValue := *openapiclient.NewTagValue("Value_example") // TagValue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.UpdateResourceTagValue(context.Background(), srn, key).TagValue(tagValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.UpdateResourceTagValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResourceTagValue`: TagBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1TagsAPIsAPI.UpdateResourceTagValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srn** | **string** | SRN | 
**key** | **string** | 태그 key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceTagValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagValue** | [**TagValue**](TagValue.md) |  | 

### Return type

[**TagBaseResponse**](TagBaseResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceTags

> TagBaseResponse UpdateResourceTags(ctx, srn).TagSetRequest(tagSetRequest).Execute()

UpdateResourceTags



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
	tagSetRequest := *openapiclient.NewTagSetRequest([]openapiclient.Tag{*openapiclient.NewTag("Key_example", "Value_example")}) // TagSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.UpdateResourceTags(context.Background(), srn).TagSetRequest(tagSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.UpdateResourceTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResourceTags`: TagBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1TagsAPIsAPI.UpdateResourceTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srn** | **string** | SRN | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagSetRequest** | [**TagSetRequest**](TagSetRequest.md) |  | 

### Return type

[**TagBaseResponse**](TagBaseResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTags

> TagBaseResponse UpdateTags(ctx).TagsSetRequest(tagsSetRequest).Execute()

UpdateTags



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
	tagsSetRequest := *openapiclient.NewTagsSetRequest(map[string][]openapiclient.Tag{"key": []openapiclient.Tag{*openapiclient.NewTag("Key_example", "Value_example")}}) // TagsSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcemanagerV1TagsAPIsAPI.UpdateTags(context.Background()).TagsSetRequest(tagsSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcemanagerV1TagsAPIsAPI.UpdateTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTags`: TagBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcemanagerV1TagsAPIsAPI.UpdateTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagsSetRequest** | [**TagsSetRequest**](TagsSetRequest.md) |  | 

### Return type

[**TagBaseResponse**](TagBaseResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

