# \SecurityGroupV1SecurityGroupApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityGroup**](SecurityGroupV1SecurityGroupApiAPI.md#CreateSecurityGroup) | **Post** /v1/security-groups | Create Security Group
[**DeleteSecurityGroup**](SecurityGroupV1SecurityGroupApiAPI.md#DeleteSecurityGroup) | **Delete** /v1/security-groups/{security_group_id} | Delete Security Group
[**GetSecurityGroupQuotas**](SecurityGroupV1SecurityGroupApiAPI.md#GetSecurityGroupQuotas) | **Get** /v1/security-groups/metrics/quotas | Get Security Group Quotas
[**ListSecurityGroups**](SecurityGroupV1SecurityGroupApiAPI.md#ListSecurityGroups) | **Get** /v1/security-groups | List Security Groups
[**SetSecurityGroup**](SecurityGroupV1SecurityGroupApiAPI.md#SetSecurityGroup) | **Put** /v1/security-groups/{security_group_id} | Set Security Group
[**ShowSecurityGroup**](SecurityGroupV1SecurityGroupApiAPI.md#ShowSecurityGroup) | **Get** /v1/security-groups/{security_group_id} | Show Security Group



## CreateSecurityGroup

> SecurityGroupShowResponse CreateSecurityGroup(ctx).SecurityGroupCreateRequest(securityGroupCreateRequest).Execute()

Create Security Group



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
	securityGroupCreateRequest := *openapiclient.NewSecurityGroupCreateRequest("Name_example") // SecurityGroupCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupV1SecurityGroupApiAPI.CreateSecurityGroup(context.Background()).SecurityGroupCreateRequest(securityGroupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupV1SecurityGroupApiAPI.CreateSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityGroup`: SecurityGroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupV1SecurityGroupApiAPI.CreateSecurityGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityGroupCreateRequest** | [**SecurityGroupCreateRequest**](SecurityGroupCreateRequest.md) |  | 

### Return type

[**SecurityGroupShowResponse**](SecurityGroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityGroup

> DeleteSecurityGroup(ctx, securityGroupId).Execute()

Delete Security Group



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
	securityGroupId := "securityGroupId_example" // string | Security Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityGroupV1SecurityGroupApiAPI.DeleteSecurityGroup(context.Background(), securityGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupV1SecurityGroupApiAPI.DeleteSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | Security Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityGroupRequest struct via the builder pattern


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


## GetSecurityGroupQuotas

> SecurityGroupQuotaResponse GetSecurityGroupQuotas(ctx).SecurityGroupId(securityGroupId).Execute()

Get Security Group Quotas



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
	securityGroupId := "securityGroupId_example" // string | Security Group ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupV1SecurityGroupApiAPI.GetSecurityGroupQuotas(context.Background()).SecurityGroupId(securityGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupV1SecurityGroupApiAPI.GetSecurityGroupQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityGroupQuotas`: SecurityGroupQuotaResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupV1SecurityGroupApiAPI.GetSecurityGroupQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityGroupQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityGroupId** | **string** | Security Group ID | 

### Return type

[**SecurityGroupQuotaResponse**](SecurityGroupQuotaResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityGroups

> SecurityGroupListResponse ListSecurityGroups(ctx).Size(size).Page(page).Sort(sort).Id(id).Name(name).Execute()

List Security Groups



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
	id := "id_example" // string | Security Group ID (optional)
	name := "name_example" // string | Security Group Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupV1SecurityGroupApiAPI.ListSecurityGroups(context.Background()).Size(size).Page(page).Sort(sort).Id(id).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupV1SecurityGroupApiAPI.ListSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityGroups`: SecurityGroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupV1SecurityGroupApiAPI.ListSecurityGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **id** | **string** | Security Group ID | 
 **name** | **string** | Security Group Name | 

### Return type

[**SecurityGroupListResponse**](SecurityGroupListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSecurityGroup

> SetSecurityGroup(ctx, securityGroupId).SecurityGroupSetRequest(securityGroupSetRequest).Execute()

Set Security Group



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
	securityGroupId := "securityGroupId_example" // string | Security Group ID
	securityGroupSetRequest := *openapiclient.NewSecurityGroupSetRequest() // SecurityGroupSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityGroupV1SecurityGroupApiAPI.SetSecurityGroup(context.Background(), securityGroupId).SecurityGroupSetRequest(securityGroupSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupV1SecurityGroupApiAPI.SetSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | Security Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityGroupSetRequest** | [**SecurityGroupSetRequest**](SecurityGroupSetRequest.md) |  | 

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


## ShowSecurityGroup

> SecurityGroupShowResponse ShowSecurityGroup(ctx, securityGroupId).Execute()

Show Security Group



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
	securityGroupId := "securityGroupId_example" // string | Security Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupV1SecurityGroupApiAPI.ShowSecurityGroup(context.Background(), securityGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupV1SecurityGroupApiAPI.ShowSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowSecurityGroup`: SecurityGroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupV1SecurityGroupApiAPI.ShowSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | Security Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityGroupShowResponse**](SecurityGroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

