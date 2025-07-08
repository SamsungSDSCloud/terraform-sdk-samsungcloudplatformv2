# \MysqlV1MysqlMasterDataApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlListEngineVersionProperties**](MysqlV1MysqlMasterDataApiAPI.md#MysqlListEngineVersionProperties) | **Get** /v1/engine-versions/{dbaas_engine_version_id}/properties | List Engine Version Properties
[**MysqlListEngineVersions**](MysqlV1MysqlMasterDataApiAPI.md#MysqlListEngineVersions) | **Get** /v1/engine-versions | List Engine Versions
[**MysqlListParameterGroups**](MysqlV1MysqlMasterDataApiAPI.md#MysqlListParameterGroups) | **Get** /v1/parameter-groups | List Parameter Groups
[**MysqlListParameters**](MysqlV1MysqlMasterDataApiAPI.md#MysqlListParameters) | **Get** /v1/parameters | List Parameters
[**MysqlListServerTypes**](MysqlV1MysqlMasterDataApiAPI.md#MysqlListServerTypes) | **Get** /v1/server-types | List Server Types
[**MysqlShowRequest**](MysqlV1MysqlMasterDataApiAPI.md#MysqlShowRequest) | **Get** /v1/requests/{request_id} | Show Request



## MysqlListEngineVersionProperties

> EnginePropertyListResponse MysqlListEngineVersionProperties(ctx, dbaasEngineVersionId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlMasterDataApiAPI.MysqlListEngineVersionProperties(context.Background(), dbaasEngineVersionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlMasterDataApiAPI.MysqlListEngineVersionProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlListEngineVersionProperties`: EnginePropertyListResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlMasterDataApiAPI.MysqlListEngineVersionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbaasEngineVersionId** | **string** | DBaaS engine version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlListEngineVersionPropertiesRequest struct via the builder pattern


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


## MysqlListEngineVersions

> EngineListResponse MysqlListEngineVersions(ctx).Id(id).ProductImageType(productImageType).EosIncluded(eosIncluded).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlMasterDataApiAPI.MysqlListEngineVersions(context.Background()).Id(id).ProductImageType(productImageType).EosIncluded(eosIncluded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlMasterDataApiAPI.MysqlListEngineVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlListEngineVersions`: EngineListResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlMasterDataApiAPI.MysqlListEngineVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMysqlListEngineVersionsRequest struct via the builder pattern


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


## MysqlListParameterGroups

> ParameterGroupListResponse MysqlListParameterGroups(ctx).MajorVersion(majorVersion).ProductImageType(productImageType).Sort(sort).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlMasterDataApiAPI.MysqlListParameterGroups(context.Background()).MajorVersion(majorVersion).ProductImageType(productImageType).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlMasterDataApiAPI.MysqlListParameterGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlListParameterGroups`: ParameterGroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlMasterDataApiAPI.MysqlListParameterGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMysqlListParameterGroupsRequest struct via the builder pattern


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


## MysqlListParameters

> ParameterPageResponse MysqlListParameters(ctx).DbaasParameterGroupId(dbaasParameterGroupId).Size(size).Page(page).Sort(sort).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlMasterDataApiAPI.MysqlListParameters(context.Background()).DbaasParameterGroupId(dbaasParameterGroupId).Size(size).Page(page).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlMasterDataApiAPI.MysqlListParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlListParameters`: ParameterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlMasterDataApiAPI.MysqlListParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMysqlListParametersRequest struct via the builder pattern


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


## MysqlListServerTypes

> ServerTypeListResponse MysqlListServerTypes(ctx).Name(name).ProductImageType(productImageType).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlMasterDataApiAPI.MysqlListServerTypes(context.Background()).Name(name).ProductImageType(productImageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlMasterDataApiAPI.MysqlListServerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlListServerTypes`: ServerTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlMasterDataApiAPI.MysqlListServerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMysqlListServerTypesRequest struct via the builder pattern


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


## MysqlShowRequest

> RequestStateResponse MysqlShowRequest(ctx, requestId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlMasterDataApiAPI.MysqlShowRequest(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlMasterDataApiAPI.MysqlShowRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlShowRequest`: RequestStateResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlMasterDataApiAPI.MysqlShowRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlShowRequestRequest struct via the builder pattern


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

