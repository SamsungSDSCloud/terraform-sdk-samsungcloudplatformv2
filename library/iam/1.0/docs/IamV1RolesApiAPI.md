# \IamV1RolesApiAPI

All URIs are relative to *https://scp-iam-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRolePolicyBindings**](IamV1RolesApiAPI.md#AddRolePolicyBindings) | **Post** /v1/roles/{role_id}/policy-bindings | Attach policies to Role
[**CreateRole**](IamV1RolesApiAPI.md#CreateRole) | **Post** /v1/roles | Create role
[**DeleteBulkRole**](IamV1RolesApiAPI.md#DeleteBulkRole) | **Delete** /v1/roles/bulk | Delete bulk role
[**DeleteRole**](IamV1RolesApiAPI.md#DeleteRole) | **Delete** /v1/roles/{role_id} | Delete role
[**ListRole**](IamV1RolesApiAPI.md#ListRole) | **Get** /v1/roles | Get role list
[**ListRolePolicyBindings**](IamV1RolesApiAPI.md#ListRolePolicyBindings) | **Get** /v1/roles/{role_id}/policy-bindings | Get Role Policy-Binding list
[**RemoveBulkRolePolicyBindings**](IamV1RolesApiAPI.md#RemoveBulkRolePolicyBindings) | **Delete** /v1/roles/{role_id}/policy-bindings | Detach policies to Role
[**RemoveRolePolicyBinding**](IamV1RolesApiAPI.md#RemoveRolePolicyBinding) | **Delete** /v1/roles/{role_id}/policy-bindings/{policy_id} | Detach policy to Role
[**SetRole**](IamV1RolesApiAPI.md#SetRole) | **Put** /v1/roles/{role_id} | Set role
[**SetRoleTrustPolicy**](IamV1RolesApiAPI.md#SetRoleTrustPolicy) | **Put** /v1/roles/{role_id}/trust-policy | Set role trust policy
[**ShowRole**](IamV1RolesApiAPI.md#ShowRole) | **Get** /v1/roles/{role_id} | Get role



## AddRolePolicyBindings

> RolePolicyBindingResponse AddRolePolicyBindings(ctx, roleId).RolePolicyBindingRequest(rolePolicyBindingRequest).Execute()

Attach policies to Role



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
	roleId := "roleId_example" // string | 역할 ID
	rolePolicyBindingRequest := *openapiclient.NewRolePolicyBindingRequest([]string{"PolicyIds_example"}) // RolePolicyBindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1RolesApiAPI.AddRolePolicyBindings(context.Background(), roleId).RolePolicyBindingRequest(rolePolicyBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1RolesApiAPI.AddRolePolicyBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRolePolicyBindings`: RolePolicyBindingResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1RolesApiAPI.AddRolePolicyBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | 역할 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRolePolicyBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rolePolicyBindingRequest** | [**RolePolicyBindingRequest**](RolePolicyBindingRequest.md) |  | 

### Return type

[**RolePolicyBindingResponse**](RolePolicyBindingResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> RoleShowResponse CreateRole(ctx).RoleCreateRequest(roleCreateRequest).Execute()

Create role



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
	roleCreateRequest := *openapiclient.NewRoleCreateRequest(*openapiclient.NewPolicyDocument([]openapiclient.Statement{*openapiclient.NewStatement("Allow")}, "2024-07-01"), "Name_example", []string{"PolicyIds_example"}) // RoleCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1RolesApiAPI.CreateRole(context.Background()).RoleCreateRequest(roleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1RolesApiAPI.CreateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRole`: RoleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1RolesApiAPI.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleCreateRequest** | [**RoleCreateRequest**](RoleCreateRequest.md) |  | 

### Return type

[**RoleShowResponse**](RoleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBulkRole

> DeleteBulkRole(ctx).RoleBulkDeleteRequest(roleBulkDeleteRequest).Execute()

Delete bulk role



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
	roleBulkDeleteRequest := *openapiclient.NewRoleBulkDeleteRequest([]string{"RoleNames_example"}) // RoleBulkDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1RolesApiAPI.DeleteBulkRole(context.Background()).RoleBulkDeleteRequest(roleBulkDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1RolesApiAPI.DeleteBulkRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBulkRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleBulkDeleteRequest** | [**RoleBulkDeleteRequest**](RoleBulkDeleteRequest.md) |  | 

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


## DeleteRole

> DeleteRole(ctx, roleId).Execute()

Delete role



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
	roleId := "roleId_example" // string | 역할 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1RolesApiAPI.DeleteRole(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1RolesApiAPI.DeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | 역할 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


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


## ListRole

> RolePageResponse ListRole(ctx).Size(size).Page(page).Sort(sort).Name(name).Types(types).AccountId(accountId).ExcludePolicyId(excludePolicyId).Execute()

Get role list



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
	name := "name_example" // string | 역할명 (optional)
	types := "types_example" // string | 역할 유형 (optional)
	accountId := "accountId_example" // string | 역할이 속한 계정의 ID (optional)
	excludePolicyId := "excludePolicyId_example" // string | 예외 대상 정책 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1RolesApiAPI.ListRole(context.Background()).Size(size).Page(page).Sort(sort).Name(name).Types(types).AccountId(accountId).ExcludePolicyId(excludePolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1RolesApiAPI.ListRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRole`: RolePageResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1RolesApiAPI.ListRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | 역할명 | 
 **types** | **string** | 역할 유형 | 
 **accountId** | **string** | 역할이 속한 계정의 ID | 
 **excludePolicyId** | **string** | 예외 대상 정책 | 

### Return type

[**RolePageResponse**](RolePageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRolePolicyBindings

> RolePolicyBindingResponse ListRolePolicyBindings(ctx, roleId).Size(size).Page(page).Sort(sort).PolicyName(policyName).Execute()

Get Role Policy-Binding list



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
	roleId := "roleId_example" // string | 역할 ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	policyName := "policyName_example" // string | 역할에 연결된 정책 이름 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1RolesApiAPI.ListRolePolicyBindings(context.Background(), roleId).Size(size).Page(page).Sort(sort).PolicyName(policyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1RolesApiAPI.ListRolePolicyBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRolePolicyBindings`: RolePolicyBindingResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1RolesApiAPI.ListRolePolicyBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | 역할 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRolePolicyBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **policyName** | **string** | 역할에 연결된 정책 이름 | 

### Return type

[**RolePolicyBindingResponse**](RolePolicyBindingResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveBulkRolePolicyBindings

> RemoveBulkRolePolicyBindings(ctx, roleId).Execute()

Detach policies to Role



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
	roleId := "roleId_example" // string | 역할 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1RolesApiAPI.RemoveBulkRolePolicyBindings(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1RolesApiAPI.RemoveBulkRolePolicyBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | 역할 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveBulkRolePolicyBindingsRequest struct via the builder pattern


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


## RemoveRolePolicyBinding

> RemoveRolePolicyBinding(ctx, roleId, policyId).Execute()

Detach policy to Role



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
	roleId := "roleId_example" // string | 역할 ID
	policyId := "policyId_example" // string | 역할에 연결된 정책 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1RolesApiAPI.RemoveRolePolicyBinding(context.Background(), roleId, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1RolesApiAPI.RemoveRolePolicyBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | 역할 ID | 
**policyId** | **string** | 역할에 연결된 정책 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRolePolicyBindingRequest struct via the builder pattern


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


## SetRole

> RoleShowResponse SetRole(ctx, roleId).RoleSetRequest(roleSetRequest).Execute()

Set role



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
	roleId := "roleId_example" // string | 역할 ID
	roleSetRequest := *openapiclient.NewRoleSetRequest() // RoleSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1RolesApiAPI.SetRole(context.Background(), roleId).RoleSetRequest(roleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1RolesApiAPI.SetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRole`: RoleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1RolesApiAPI.SetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | 역할 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleSetRequest** | [**RoleSetRequest**](RoleSetRequest.md) |  | 

### Return type

[**RoleShowResponse**](RoleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRoleTrustPolicy

> RoleShowResponse SetRoleTrustPolicy(ctx, roleId).RoleTrustPolicyRequest(roleTrustPolicyRequest).Execute()

Set role trust policy



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
	roleId := "roleId_example" // string | 역할 ID
	roleTrustPolicyRequest := *openapiclient.NewRoleTrustPolicyRequest(*openapiclient.NewPolicyDocument([]openapiclient.Statement{*openapiclient.NewStatement("Allow")}, "2024-07-01")) // RoleTrustPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1RolesApiAPI.SetRoleTrustPolicy(context.Background(), roleId).RoleTrustPolicyRequest(roleTrustPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1RolesApiAPI.SetRoleTrustPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRoleTrustPolicy`: RoleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1RolesApiAPI.SetRoleTrustPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | 역할 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRoleTrustPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleTrustPolicyRequest** | [**RoleTrustPolicyRequest**](RoleTrustPolicyRequest.md) |  | 

### Return type

[**RoleShowResponse**](RoleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowRole

> RoleShowResponse ShowRole(ctx, roleId).Execute()

Get role



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
	roleId := "roleId_example" // string | 역할 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1RolesApiAPI.ShowRole(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1RolesApiAPI.ShowRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowRole`: RoleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1RolesApiAPI.ShowRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | 역할 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleShowResponse**](RoleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

