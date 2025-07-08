# \IamV1ResourcePoliciesApiAPI

All URIs are relative to *https://scp-iam-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteResourcePolicy**](IamV1ResourcePoliciesApiAPI.md#DeleteResourcePolicy) | **Delete** /v1/resource-policies/{srn} | Delete the resource based policy
[**SetResourcePolicy**](IamV1ResourcePoliciesApiAPI.md#SetResourcePolicy) | **Put** /v1/resource-policies/{srn} | Set the resource based policy
[**ShowResourcePolicy**](IamV1ResourcePoliciesApiAPI.md#ShowResourcePolicy) | **Get** /v1/resource-policies/{srn} | Show the resource based policy



## DeleteResourcePolicy

> DeleteResourcePolicy(ctx, srn).Execute()

Delete the resource based policy



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
	srn := "c3JuOmRldjI6OjEyMzRhYjU2N2NkNjQ3NjllOGY5ZzQ5MGhpMzA0ODkxOmtyLXdlc3QxOjp2aXJ0dWFsc2VydmVyOnZpcnR1YWwtc2VydmVyL3o5NWZhNTYxLTExdTMtNGZydy05NjJnLTNiMTIzMTIzYzQ5Ng==" // string | 삼성 리소스 자원명

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1ResourcePoliciesApiAPI.DeleteResourcePolicy(context.Background(), srn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1ResourcePoliciesApiAPI.DeleteResourcePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srn** | **string** | 삼성 리소스 자원명 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourcePolicyRequest struct via the builder pattern


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


## SetResourcePolicy

> ResourcePolicyShowResponse SetResourcePolicy(ctx, srn).ResourcePolicySetRequest(resourcePolicySetRequest).Execute()

Set the resource based policy



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
	srn := "c3JuOmRldjI6OjEyMzRhYjU2N2NkNjQ3NjllOGY5ZzQ5MGhpMzA0ODkxOmtyLXdlc3QxOjp2aXJ0dWFsc2VydmVyOnZpcnR1YWwtc2VydmVyL3o5NWZhNTYxLTExdTMtNGZydy05NjJnLTNiMTIzMTIzYzQ5Ng==" // string | 삼성 리소스 자원명
	resourcePolicySetRequest := *openapiclient.NewResourcePolicySetRequest(*openapiclient.NewPolicyVersionCreateRequest(*openapiclient.NewIamPolicyDocument([]openapiclient.Statement{*openapiclient.NewStatement("Allow")}, "2024-07-01"))) // ResourcePolicySetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1ResourcePoliciesApiAPI.SetResourcePolicy(context.Background(), srn).ResourcePolicySetRequest(resourcePolicySetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1ResourcePoliciesApiAPI.SetResourcePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetResourcePolicy`: ResourcePolicyShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1ResourcePoliciesApiAPI.SetResourcePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srn** | **string** | 삼성 리소스 자원명 | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetResourcePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourcePolicySetRequest** | [**ResourcePolicySetRequest**](ResourcePolicySetRequest.md) |  | 

### Return type

[**ResourcePolicyShowResponse**](ResourcePolicyShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowResourcePolicy

> ResourcePolicyShowResponse ShowResourcePolicy(ctx, srn).Execute()

Show the resource based policy



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
	srn := "c3JuOmRldjI6OjEyMzRhYjU2N2NkNjQ3NjllOGY5ZzQ5MGhpMzA0ODkxOmtyLXdlc3QxOjp2aXJ0dWFsc2VydmVyOnZpcnR1YWwtc2VydmVyL3o5NWZhNTYxLTExdTMtNGZydy05NjJnLTNiMTIzMTIzYzQ5Ng==" // string | 삼성 리소스 자원명

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1ResourcePoliciesApiAPI.ShowResourcePolicy(context.Background(), srn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1ResourcePoliciesApiAPI.ShowResourcePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowResourcePolicy`: ResourcePolicyShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1ResourcePoliciesApiAPI.ShowResourcePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srn** | **string** | 삼성 리소스 자원명 | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowResourcePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourcePolicyShowResponse**](ResourcePolicyShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

