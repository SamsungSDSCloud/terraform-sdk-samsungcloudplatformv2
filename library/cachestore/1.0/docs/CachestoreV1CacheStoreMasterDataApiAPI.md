# \CachestoreV1CacheStoreMasterDataApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CachestoreListEngineVersionProperties**](CachestoreV1CacheStoreMasterDataApiAPI.md#CachestoreListEngineVersionProperties) | **Get** /v1/engine-versions/{dbaas_engine_version_id}/properties | List Engine Version Properties
[**CachestoreListEngineVersions**](CachestoreV1CacheStoreMasterDataApiAPI.md#CachestoreListEngineVersions) | **Get** /v1/engine-versions | List Engine Versions
[**CachestoreListParameterGroups**](CachestoreV1CacheStoreMasterDataApiAPI.md#CachestoreListParameterGroups) | **Get** /v1/parameter-groups | List Parameter Groups
[**CachestoreListParameters**](CachestoreV1CacheStoreMasterDataApiAPI.md#CachestoreListParameters) | **Get** /v1/parameters | List Parameters
[**CachestoreListServerTypes**](CachestoreV1CacheStoreMasterDataApiAPI.md#CachestoreListServerTypes) | **Get** /v1/server-types | List Server Types
[**CachestoreShowRequest**](CachestoreV1CacheStoreMasterDataApiAPI.md#CachestoreShowRequest) | **Get** /v1/requests/{request_id} | Show Request



## CachestoreListEngineVersionProperties

> EnginePropertyListResponse CachestoreListEngineVersionProperties(ctx, dbaasEngineVersionId).Execute()

List Engine Version Properties



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
	dbaasEngineVersionId := "dbaasEngineVersionId_example" // string | DBaaS engine version ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListEngineVersionProperties(context.Background(), dbaasEngineVersionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListEngineVersionProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreListEngineVersionProperties`: EnginePropertyListResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListEngineVersionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbaasEngineVersionId** | **string** | DBaaS engine version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreListEngineVersionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnginePropertyListResponse**](EnginePropertyListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CachestoreListEngineVersions

> EngineListResponse CachestoreListEngineVersions(ctx).Id(id).ProductImageType(productImageType).EosIncluded(eosIncluded).Execute()

List Engine Versions



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
	id := "id_example" // string | DBaaS engine version ID (optional)
	productImageType := openapiclient.ProductImageType("PostgreSQL Community") // ProductImageType | Product image type (optional)
	eosIncluded := true // bool | EoS included (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListEngineVersions(context.Background()).Id(id).ProductImageType(productImageType).EosIncluded(eosIncluded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListEngineVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreListEngineVersions`: EngineListResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListEngineVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreListEngineVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | DBaaS engine version ID | 
 **productImageType** | [**ProductImageType**](ProductImageType.md) | Product image type | 
 **eosIncluded** | **bool** | EoS included | [default to false]

### Return type

[**EngineListResponse**](EngineListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CachestoreListParameterGroups

> ParameterGroupListResponse CachestoreListParameterGroups(ctx).MajorVersion(majorVersion).ProductImageType(productImageType).Sort(sort).Execute()

List Parameter Groups



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
	majorVersion := "majorVersion_example" // string | Software major version (optional)
	productImageType := "productImageType_example" // string | Product image type (optional)
	sort := "created_at:desc" // string | sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListParameterGroups(context.Background()).MajorVersion(majorVersion).ProductImageType(productImageType).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListParameterGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreListParameterGroups`: ParameterGroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListParameterGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreListParameterGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **majorVersion** | **string** | Software major version | 
 **productImageType** | **string** | Product image type | 
 **sort** | **string** | sort | 

### Return type

[**ParameterGroupListResponse**](ParameterGroupListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CachestoreListParameters

> ParameterPageResponse CachestoreListParameters(ctx).DbaasParameterGroupId(dbaasParameterGroupId).Size(size).Page(page).Sort(sort).Execute()

List Parameters



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
	dbaasParameterGroupId := "dbaasParameterGroupId_example" // string | Parameter group ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListParameters(context.Background()).DbaasParameterGroupId(dbaasParameterGroupId).Size(size).Page(page).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreListParameters`: ParameterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreListParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbaasParameterGroupId** | **string** | Parameter group ID | 
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 

### Return type

[**ParameterPageResponse**](ParameterPageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CachestoreListServerTypes

> ServerTypeListResponse CachestoreListServerTypes(ctx).Name(name).ProductImageType(productImageType).Execute()

List Server Types



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
	name := "name_example" // string | Server type name (optional)
	productImageType := "productImageType_example" // string | Product image type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListServerTypes(context.Background()).Name(name).ProductImageType(productImageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListServerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreListServerTypes`: ServerTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreMasterDataApiAPI.CachestoreListServerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreListServerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Server type name | 
 **productImageType** | **string** | Product image type | 

### Return type

[**ServerTypeListResponse**](ServerTypeListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CachestoreShowRequest

> RequestStateResponse CachestoreShowRequest(ctx, requestId).Execute()

Show Request



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
	requestId := "requestId_example" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreMasterDataApiAPI.CachestoreShowRequest(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreMasterDataApiAPI.CachestoreShowRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreShowRequest`: RequestStateResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreMasterDataApiAPI.CachestoreShowRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreShowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestStateResponse**](RequestStateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

