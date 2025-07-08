# \SqlserverV1SqlserverMasterDataApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SqlserverListEngineVersionProperties**](SqlserverV1SqlserverMasterDataApiAPI.md#SqlserverListEngineVersionProperties) | **Get** /v1/engine-versions/{dbaas_engine_version_id}/properties | List Engine Version Properties
[**SqlserverListEngineVersions**](SqlserverV1SqlserverMasterDataApiAPI.md#SqlserverListEngineVersions) | **Get** /v1/engine-versions | List Engine Versions
[**SqlserverListParameterGroups**](SqlserverV1SqlserverMasterDataApiAPI.md#SqlserverListParameterGroups) | **Get** /v1/parameter-groups | List Parameter Groups
[**SqlserverListParameters**](SqlserverV1SqlserverMasterDataApiAPI.md#SqlserverListParameters) | **Get** /v1/parameters | List Parameters
[**SqlserverListServerTypes**](SqlserverV1SqlserverMasterDataApiAPI.md#SqlserverListServerTypes) | **Get** /v1/server-types | List Server Types
[**SqlserverShowRequest**](SqlserverV1SqlserverMasterDataApiAPI.md#SqlserverShowRequest) | **Get** /v1/requests/{request_id} | Show Request



## SqlserverListEngineVersionProperties

> EnginePropertyListResponse SqlserverListEngineVersionProperties(ctx, dbaasEngineVersionId).Execute()

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
	resp, r, err := apiClient.SqlserverV1SqlserverMasterDataApiAPI.SqlserverListEngineVersionProperties(context.Background(), dbaasEngineVersionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverMasterDataApiAPI.SqlserverListEngineVersionProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverListEngineVersionProperties`: EnginePropertyListResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverMasterDataApiAPI.SqlserverListEngineVersionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbaasEngineVersionId** | **string** | DBaaS engine version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverListEngineVersionPropertiesRequest struct via the builder pattern


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


## SqlserverListEngineVersions

> EngineListResponse SqlserverListEngineVersions(ctx).Id(id).ProductImageType(productImageType).EosIncluded(eosIncluded).Execute()

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
	resp, r, err := apiClient.SqlserverV1SqlserverMasterDataApiAPI.SqlserverListEngineVersions(context.Background()).Id(id).ProductImageType(productImageType).EosIncluded(eosIncluded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverMasterDataApiAPI.SqlserverListEngineVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverListEngineVersions`: EngineListResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverMasterDataApiAPI.SqlserverListEngineVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverListEngineVersionsRequest struct via the builder pattern


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


## SqlserverListParameterGroups

> ParameterGroupListResponse SqlserverListParameterGroups(ctx).MajorVersion(majorVersion).ProductImageType(productImageType).Sort(sort).Execute()

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
	resp, r, err := apiClient.SqlserverV1SqlserverMasterDataApiAPI.SqlserverListParameterGroups(context.Background()).MajorVersion(majorVersion).ProductImageType(productImageType).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverMasterDataApiAPI.SqlserverListParameterGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverListParameterGroups`: ParameterGroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverMasterDataApiAPI.SqlserverListParameterGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverListParameterGroupsRequest struct via the builder pattern


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


## SqlserverListParameters

> ParameterPageResponse SqlserverListParameters(ctx).DbaasParameterGroupId(dbaasParameterGroupId).Size(size).Page(page).Sort(sort).Execute()

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
	resp, r, err := apiClient.SqlserverV1SqlserverMasterDataApiAPI.SqlserverListParameters(context.Background()).DbaasParameterGroupId(dbaasParameterGroupId).Size(size).Page(page).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverMasterDataApiAPI.SqlserverListParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverListParameters`: ParameterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverMasterDataApiAPI.SqlserverListParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverListParametersRequest struct via the builder pattern


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


## SqlserverListServerTypes

> ServerTypeListResponse SqlserverListServerTypes(ctx).Name(name).ProductImageType(productImageType).Execute()

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
	resp, r, err := apiClient.SqlserverV1SqlserverMasterDataApiAPI.SqlserverListServerTypes(context.Background()).Name(name).ProductImageType(productImageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverMasterDataApiAPI.SqlserverListServerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverListServerTypes`: ServerTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverMasterDataApiAPI.SqlserverListServerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverListServerTypesRequest struct via the builder pattern


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


## SqlserverShowRequest

> RequestStateResponse SqlserverShowRequest(ctx, requestId).Execute()

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
	resp, r, err := apiClient.SqlserverV1SqlserverMasterDataApiAPI.SqlserverShowRequest(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverMasterDataApiAPI.SqlserverShowRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverShowRequest`: RequestStateResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverMasterDataApiAPI.SqlserverShowRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverShowRequestRequest struct via the builder pattern


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

