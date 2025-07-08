# \MariadbV1MariadbMasterDataApiAPI

All URIs are relative to *https://scp-dbaas-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariadbListEngineVersionProperties**](MariadbV1MariadbMasterDataApiAPI.md#MariadbListEngineVersionProperties) | **Get** /v1/engine-versions/{dbaas_engine_version_id}/properties | List Engine Version Properties
[**MariadbListEngineVersions**](MariadbV1MariadbMasterDataApiAPI.md#MariadbListEngineVersions) | **Get** /v1/engine-versions | List Engine Versions
[**MariadbListParameterGroups**](MariadbV1MariadbMasterDataApiAPI.md#MariadbListParameterGroups) | **Get** /v1/parameter-groups | List Parameter Groups
[**MariadbListParameters**](MariadbV1MariadbMasterDataApiAPI.md#MariadbListParameters) | **Get** /v1/parameters | List Parameters
[**MariadbListServerTypes**](MariadbV1MariadbMasterDataApiAPI.md#MariadbListServerTypes) | **Get** /v1/server-types | List Server Types
[**MariadbShowRequest**](MariadbV1MariadbMasterDataApiAPI.md#MariadbShowRequest) | **Get** /v1/requests/{request_id} | Show Request



## MariadbListEngineVersionProperties

> EnginePropertyListResponse MariadbListEngineVersionProperties(ctx, dbaasEngineVersionId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbMasterDataApiAPI.MariadbListEngineVersionProperties(context.Background(), dbaasEngineVersionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbMasterDataApiAPI.MariadbListEngineVersionProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbListEngineVersionProperties`: EnginePropertyListResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbMasterDataApiAPI.MariadbListEngineVersionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbaasEngineVersionId** | **string** | DBaaS engine version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbListEngineVersionPropertiesRequest struct via the builder pattern


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


## MariadbListEngineVersions

> EngineListResponse MariadbListEngineVersions(ctx).Id(id).ProductImageType(productImageType).EosIncluded(eosIncluded).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbMasterDataApiAPI.MariadbListEngineVersions(context.Background()).Id(id).ProductImageType(productImageType).EosIncluded(eosIncluded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbMasterDataApiAPI.MariadbListEngineVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbListEngineVersions`: EngineListResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbMasterDataApiAPI.MariadbListEngineVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMariadbListEngineVersionsRequest struct via the builder pattern


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


## MariadbListParameterGroups

> ParameterGroupListResponse MariadbListParameterGroups(ctx).MajorVersion(majorVersion).Sort(sort).Execute()

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
	sort := "created_at:desc" // string | sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MariadbV1MariadbMasterDataApiAPI.MariadbListParameterGroups(context.Background()).MajorVersion(majorVersion).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbMasterDataApiAPI.MariadbListParameterGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbListParameterGroups`: ParameterGroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbMasterDataApiAPI.MariadbListParameterGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMariadbListParameterGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **majorVersion** | **string** | Software major version | 
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


## MariadbListParameters

> ParameterPageResponse MariadbListParameters(ctx).DbaasParameterGroupId(dbaasParameterGroupId).Size(size).Page(page).Sort(sort).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbMasterDataApiAPI.MariadbListParameters(context.Background()).DbaasParameterGroupId(dbaasParameterGroupId).Size(size).Page(page).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbMasterDataApiAPI.MariadbListParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbListParameters`: ParameterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbMasterDataApiAPI.MariadbListParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMariadbListParametersRequest struct via the builder pattern


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


## MariadbListServerTypes

> ServerTypeListResponse MariadbListServerTypes(ctx).Name(name).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MariadbV1MariadbMasterDataApiAPI.MariadbListServerTypes(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbMasterDataApiAPI.MariadbListServerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbListServerTypes`: ServerTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbMasterDataApiAPI.MariadbListServerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMariadbListServerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Server type name | 

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


## MariadbShowRequest

> RequestStateResponse MariadbShowRequest(ctx, requestId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbMasterDataApiAPI.MariadbShowRequest(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbMasterDataApiAPI.MariadbShowRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbShowRequest`: RequestStateResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbMasterDataApiAPI.MariadbShowRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbShowRequestRequest struct via the builder pattern


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

