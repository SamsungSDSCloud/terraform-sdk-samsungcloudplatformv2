# \EpasV1EpasMasterDataApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EpasListEngineVersionProperties**](EpasV1EpasMasterDataApiAPI.md#EpasListEngineVersionProperties) | **Get** /v1/engine-versions/{dbaas_engine_version_id}/properties | List Engine Version Properties
[**EpasListEngineVersions**](EpasV1EpasMasterDataApiAPI.md#EpasListEngineVersions) | **Get** /v1/engine-versions | List Engine Versions
[**EpasListParameterGroups**](EpasV1EpasMasterDataApiAPI.md#EpasListParameterGroups) | **Get** /v1/parameter-groups | List Parameter Groups
[**EpasListParameters**](EpasV1EpasMasterDataApiAPI.md#EpasListParameters) | **Get** /v1/parameters | List Parameters
[**EpasListServerTypes**](EpasV1EpasMasterDataApiAPI.md#EpasListServerTypes) | **Get** /v1/server-types | List Server Types
[**EpasShowRequest**](EpasV1EpasMasterDataApiAPI.md#EpasShowRequest) | **Get** /v1/requests/{request_id} | Show Request



## EpasListEngineVersionProperties

> EnginePropertyListResponse EpasListEngineVersionProperties(ctx, dbaasEngineVersionId).Execute()

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
	resp, r, err := apiClient.EpasV1EpasMasterDataApiAPI.EpasListEngineVersionProperties(context.Background(), dbaasEngineVersionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasMasterDataApiAPI.EpasListEngineVersionProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasListEngineVersionProperties`: EnginePropertyListResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasMasterDataApiAPI.EpasListEngineVersionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbaasEngineVersionId** | **string** | DBaaS engine version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasListEngineVersionPropertiesRequest struct via the builder pattern


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


## EpasListEngineVersions

> EngineListResponse EpasListEngineVersions(ctx).Id(id).ProductImageType(productImageType).EosIncluded(eosIncluded).Execute()

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
	resp, r, err := apiClient.EpasV1EpasMasterDataApiAPI.EpasListEngineVersions(context.Background()).Id(id).ProductImageType(productImageType).EosIncluded(eosIncluded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasMasterDataApiAPI.EpasListEngineVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasListEngineVersions`: EngineListResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasMasterDataApiAPI.EpasListEngineVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEpasListEngineVersionsRequest struct via the builder pattern


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


## EpasListParameterGroups

> ParameterGroupListResponse EpasListParameterGroups(ctx).MajorVersion(majorVersion).ProductImageType(productImageType).Sort(sort).Execute()

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
	resp, r, err := apiClient.EpasV1EpasMasterDataApiAPI.EpasListParameterGroups(context.Background()).MajorVersion(majorVersion).ProductImageType(productImageType).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasMasterDataApiAPI.EpasListParameterGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasListParameterGroups`: ParameterGroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasMasterDataApiAPI.EpasListParameterGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEpasListParameterGroupsRequest struct via the builder pattern


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


## EpasListParameters

> ParameterPageResponse EpasListParameters(ctx).DbaasParameterGroupId(dbaasParameterGroupId).Size(size).Page(page).Sort(sort).Execute()

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
	resp, r, err := apiClient.EpasV1EpasMasterDataApiAPI.EpasListParameters(context.Background()).DbaasParameterGroupId(dbaasParameterGroupId).Size(size).Page(page).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasMasterDataApiAPI.EpasListParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasListParameters`: ParameterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasMasterDataApiAPI.EpasListParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEpasListParametersRequest struct via the builder pattern


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


## EpasListServerTypes

> ServerTypeListResponse EpasListServerTypes(ctx).Name(name).ProductImageType(productImageType).Execute()

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
	resp, r, err := apiClient.EpasV1EpasMasterDataApiAPI.EpasListServerTypes(context.Background()).Name(name).ProductImageType(productImageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasMasterDataApiAPI.EpasListServerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasListServerTypes`: ServerTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasMasterDataApiAPI.EpasListServerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEpasListServerTypesRequest struct via the builder pattern


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


## EpasShowRequest

> RequestStateResponse EpasShowRequest(ctx, requestId).Execute()

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
	resp, r, err := apiClient.EpasV1EpasMasterDataApiAPI.EpasShowRequest(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasMasterDataApiAPI.EpasShowRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasShowRequest`: RequestStateResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasMasterDataApiAPI.EpasShowRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasShowRequestRequest struct via the builder pattern


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

