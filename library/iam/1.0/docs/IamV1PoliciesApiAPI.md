# \IamV1PoliciesApiAPI

All URIs are relative to *https://scp-iam-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicy**](IamV1PoliciesApiAPI.md#CreatePolicy) | **Post** /v1/policies | Create a policy
[**DeletePolicies**](IamV1PoliciesApiAPI.md#DeletePolicies) | **Delete** /v1/policies/bulk | Remove policies
[**DeletePolicy**](IamV1PoliciesApiAPI.md#DeletePolicy) | **Delete** /v1/policies/{policy_id} | Delete the policy
[**ListPolicy**](IamV1PoliciesApiAPI.md#ListPolicy) | **Get** /v1/policies | Get policy list
[**ListPolicyBinding**](IamV1PoliciesApiAPI.md#ListPolicyBinding) | **Get** /v1/policies/{policy_id}/bindings | Get a specific policy binding list
[**QueryPolicy**](IamV1PoliciesApiAPI.md#QueryPolicy) | **Post** /v1/policies/list | Query policy list
[**SetPolicy**](IamV1PoliciesApiAPI.md#SetPolicy) | **Put** /v1/policies/{policy_id} | Set the policy
[**SetPolicyGroupBinding**](IamV1PoliciesApiAPI.md#SetPolicyGroupBinding) | **Put** /v1/policies/{policy_id}/bindings | Set the policy group binding
[**ShowPolicy**](IamV1PoliciesApiAPI.md#ShowPolicy) | **Get** /v1/policies/{policy_id} | Show the policy



## CreatePolicy

> PolicyShowResponse CreatePolicy(ctx).PolicyCreateRequest(policyCreateRequest).Execute()

Create a policy



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
	policyCreateRequest := *openapiclient.NewPolicyCreateRequest("ViewerAccess", *openapiclient.NewPolicyVersionCreateRequest(*openapiclient.NewIamPolicyDocument([]openapiclient.Statement{*openapiclient.NewStatement("Allow")}, "2024-07-01"))) // PolicyCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1PoliciesApiAPI.CreatePolicy(context.Background()).PolicyCreateRequest(policyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1PoliciesApiAPI.CreatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicy`: PolicyShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1PoliciesApiAPI.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyCreateRequest** | [**PolicyCreateRequest**](PolicyCreateRequest.md) |  | 

### Return type

[**PolicyShowResponse**](PolicyShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicies

> DeletePolicies(ctx).ListPolicyRemoveRequest(listPolicyRemoveRequest).Execute()

Remove policies



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
	listPolicyRemoveRequest := *openapiclient.NewListPolicyRemoveRequest() // ListPolicyRemoveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1PoliciesApiAPI.DeletePolicies(context.Background()).ListPolicyRemoveRequest(listPolicyRemoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1PoliciesApiAPI.DeletePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listPolicyRemoveRequest** | [**ListPolicyRemoveRequest**](ListPolicyRemoveRequest.md) |  | 

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


## DeletePolicy

> DeletePolicy(ctx, policyId).Execute()

Delete the policy



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
	policyId := "37f2e31ff86b415698d7e8eeafab445d" // string | Policy ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1PoliciesApiAPI.DeletePolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1PoliciesApiAPI.DeletePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


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


## ListPolicy

> PolicyPageResponse ListPolicy(ctx).Size(size).Page(page).Sort(sort).Id(id).PolicyName(policyName).PolicyType(policyType).ServiceType(serviceType).CreatorName(creatorName).CreatorEmail(creatorEmail).ModifierName(modifierName).ModifierEmail(modifierEmail).ExcludeGroupId(excludeGroupId).ExcludeUserId(excludeUserId).Execute()

Get policy list



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
	id := "37f2e31ff86b415698d7e8eeafab445d" // string | Policy ID (optional)
	policyName := "ViewerAccess" // string | Policy Name (optional)
	policyType := *openapiclient.NewPolicyType1() // PolicyType1 | Policy Type (optional)
	serviceType := *openapiclient.NewServiceType() // ServiceType | Service Type (optional)
	creatorName := "test1" // string | Policy 생성자 성, 이름 (optional)
	creatorEmail := "test@test.com" // string | Policy 생성자 Email (optional)
	modifierName := "test1" // string | Policy 수정자 성, 이름 (optional)
	modifierEmail := "test@test.com" // string | Policy 수정자 Email (optional)
	excludeGroupId := "a946662dc4314dac93da413a32457459" // string | 제외할 Group ID (optional)
	excludeUserId := "excludeUserId_example" // string | 제외할 User ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1PoliciesApiAPI.ListPolicy(context.Background()).Size(size).Page(page).Sort(sort).Id(id).PolicyName(policyName).PolicyType(policyType).ServiceType(serviceType).CreatorName(creatorName).CreatorEmail(creatorEmail).ModifierName(modifierName).ModifierEmail(modifierEmail).ExcludeGroupId(excludeGroupId).ExcludeUserId(excludeUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1PoliciesApiAPI.ListPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPolicy`: PolicyPageResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1PoliciesApiAPI.ListPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **id** | **string** | Policy ID | 
 **policyName** | **string** | Policy Name | 
 **policyType** | [**PolicyType1**](PolicyType1.md) | Policy Type | 
 **serviceType** | [**ServiceType**](ServiceType.md) | Service Type | 
 **creatorName** | **string** | Policy 생성자 성, 이름 | 
 **creatorEmail** | **string** | Policy 생성자 Email | 
 **modifierName** | **string** | Policy 수정자 성, 이름 | 
 **modifierEmail** | **string** | Policy 수정자 Email | 
 **excludeGroupId** | **string** | 제외할 Group ID | 
 **excludeUserId** | **string** | 제외할 User ID | 

### Return type

[**PolicyPageResponse**](PolicyPageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicyBinding

> PolicyBindingPageResponse ListPolicyBinding(ctx, policyId).Size(size).Page(page).Sort(sort).IdentityId(identityId).IdentityType(identityType).Name(name).Execute()

Get a specific policy binding list



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
	policyId := "37f2e31ff86b415698d7e8eeafab445d" // string | Policy ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	identityId := "f39c460fade34fecb05ede8f904b24b7" // string | Identity ID (optional)
	identityType := "GROUP" // string | Identity Type (optional)
	name := "TestGroup" // string | Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1PoliciesApiAPI.ListPolicyBinding(context.Background(), policyId).Size(size).Page(page).Sort(sort).IdentityId(identityId).IdentityType(identityType).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1PoliciesApiAPI.ListPolicyBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPolicyBinding`: PolicyBindingPageResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1PoliciesApiAPI.ListPolicyBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **identityId** | **string** | Identity ID | 
 **identityType** | **string** | Identity Type | 
 **name** | **string** | Name | 

### Return type

[**PolicyBindingPageResponse**](PolicyBindingPageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryPolicy

> PolicyPageResponse QueryPolicy(ctx).PolicyQueryRequest(policyQueryRequest).Size(size).Page(page).Sort(sort).Id(id).PolicyName(policyName).PolicyType(policyType).ServiceType(serviceType).CreatorName(creatorName).CreatorEmail(creatorEmail).ModifierName(modifierName).ModifierEmail(modifierEmail).ExcludeGroupId(excludeGroupId).ExcludeUserId(excludeUserId).Execute()

Query policy list



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
	policyQueryRequest := *openapiclient.NewPolicyQueryRequest() // PolicyQueryRequest | 
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	id := "37f2e31ff86b415698d7e8eeafab445d" // string | Policy ID (optional)
	policyName := "ViewerAccess" // string | Policy Name (optional)
	policyType := *openapiclient.NewPolicyType2() // PolicyType2 | Policy Type (optional)
	serviceType := *openapiclient.NewServiceType1() // ServiceType1 | Service Type (optional)
	creatorName := "test1" // string | Policy 생성자 성, 이름 (optional)
	creatorEmail := "test@test.com" // string | Policy 생성자 Email (optional)
	modifierName := "test1" // string | Policy 수정자 성, 이름 (optional)
	modifierEmail := "test@test.com" // string | Policy 수정자 Email (optional)
	excludeGroupId := "a946662dc4314dac93da413a32457459" // string | 제외할 Group ID (optional)
	excludeUserId := "excludeUserId_example" // string | 제외할 User ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1PoliciesApiAPI.QueryPolicy(context.Background()).PolicyQueryRequest(policyQueryRequest).Size(size).Page(page).Sort(sort).Id(id).PolicyName(policyName).PolicyType(policyType).ServiceType(serviceType).CreatorName(creatorName).CreatorEmail(creatorEmail).ModifierName(modifierName).ModifierEmail(modifierEmail).ExcludeGroupId(excludeGroupId).ExcludeUserId(excludeUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1PoliciesApiAPI.QueryPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryPolicy`: PolicyPageResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1PoliciesApiAPI.QueryPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyQueryRequest** | [**PolicyQueryRequest**](PolicyQueryRequest.md) |  | 
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **id** | **string** | Policy ID | 
 **policyName** | **string** | Policy Name | 
 **policyType** | [**PolicyType2**](PolicyType2.md) | Policy Type | 
 **serviceType** | [**ServiceType1**](ServiceType1.md) | Service Type | 
 **creatorName** | **string** | Policy 생성자 성, 이름 | 
 **creatorEmail** | **string** | Policy 생성자 Email | 
 **modifierName** | **string** | Policy 수정자 성, 이름 | 
 **modifierEmail** | **string** | Policy 수정자 Email | 
 **excludeGroupId** | **string** | 제외할 Group ID | 
 **excludeUserId** | **string** | 제외할 User ID | 

### Return type

[**PolicyPageResponse**](PolicyPageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPolicy

> PolicyShowResponse SetPolicy(ctx, policyId).PolicySetRequest(policySetRequest).Execute()

Set the policy



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
	policyId := "37f2e31ff86b415698d7e8eeafab445d" // string | Policy ID
	policySetRequest := *openapiclient.NewPolicySetRequest() // PolicySetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1PoliciesApiAPI.SetPolicy(context.Background(), policyId).PolicySetRequest(policySetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1PoliciesApiAPI.SetPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPolicy`: PolicyShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1PoliciesApiAPI.SetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policySetRequest** | [**PolicySetRequest**](PolicySetRequest.md) |  | 

### Return type

[**PolicyShowResponse**](PolicyShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPolicyGroupBinding

> PolicyBindingPageResponse SetPolicyGroupBinding(ctx, policyId).PolicyBindingSetRequest(policyBindingSetRequest).Execute()

Set the policy group binding



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
	policyId := "37f2e31ff86b415698d7e8eeafab445d" // string | Policy ID
	policyBindingSetRequest := *openapiclient.NewPolicyBindingSetRequest("GROUP") // PolicyBindingSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1PoliciesApiAPI.SetPolicyGroupBinding(context.Background(), policyId).PolicyBindingSetRequest(policyBindingSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1PoliciesApiAPI.SetPolicyGroupBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPolicyGroupBinding`: PolicyBindingPageResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1PoliciesApiAPI.SetPolicyGroupBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPolicyGroupBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyBindingSetRequest** | [**PolicyBindingSetRequest**](PolicyBindingSetRequest.md) |  | 

### Return type

[**PolicyBindingPageResponse**](PolicyBindingPageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPolicy

> PolicyShowResponse ShowPolicy(ctx, policyId).Execute()

Show the policy



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
	policyId := "37f2e31ff86b415698d7e8eeafab445d" // string | Policy ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1PoliciesApiAPI.ShowPolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1PoliciesApiAPI.ShowPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowPolicy`: PolicyShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1PoliciesApiAPI.ShowPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyShowResponse**](PolicyShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

