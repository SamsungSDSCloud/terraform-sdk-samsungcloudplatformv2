# \SecurityGroupV1SecurityGroupRuleApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityGroupRule**](SecurityGroupV1SecurityGroupRuleApiAPI.md#CreateSecurityGroupRule) | **Post** /v1/security-group-rules | Create Security Group Rule
[**CreateSecurityGroupRules**](SecurityGroupV1SecurityGroupRuleApiAPI.md#CreateSecurityGroupRules) | **Post** /v1/security-group-rules/bulk | Create Security Group Rules
[**DeleteSecurityGroupRule**](SecurityGroupV1SecurityGroupRuleApiAPI.md#DeleteSecurityGroupRule) | **Delete** /v1/security-group-rules/{security_group_rule_id} | Delete Security Group Rule
[**DownloadSecurityGroupRules**](SecurityGroupV1SecurityGroupRuleApiAPI.md#DownloadSecurityGroupRules) | **Get** /v1/security-group-rules/{security_group_id}/downloads | Download Security Group Rules
[**ListSecurityGroupRules**](SecurityGroupV1SecurityGroupRuleApiAPI.md#ListSecurityGroupRules) | **Get** /v1/security-group-rules | List Security Group Rules
[**ShowSecurityGroupRule**](SecurityGroupV1SecurityGroupRuleApiAPI.md#ShowSecurityGroupRule) | **Get** /v1/security-group-rules/{security_group_rule_id} | Show Security Group Rule



## CreateSecurityGroupRule

> SecurityGroupRuleShowResponse CreateSecurityGroupRule(ctx).SecurityGroupRuleCreateRequest(securityGroupRuleCreateRequest).Execute()

Create Security Group Rule



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
	securityGroupRuleCreateRequest := *openapiclient.NewSecurityGroupRuleCreateRequest("Direction_example", "Protocol_example", "SecurityGroupId_example") // SecurityGroupRuleCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupV1SecurityGroupRuleApiAPI.CreateSecurityGroupRule(context.Background()).SecurityGroupRuleCreateRequest(securityGroupRuleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupV1SecurityGroupRuleApiAPI.CreateSecurityGroupRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityGroupRule`: SecurityGroupRuleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupV1SecurityGroupRuleApiAPI.CreateSecurityGroupRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityGroupRuleCreateRequest** | [**SecurityGroupRuleCreateRequest**](SecurityGroupRuleCreateRequest.md) |  | 

### Return type

[**SecurityGroupRuleShowResponse**](SecurityGroupRuleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecurityGroupRules

> SecurityGroupRulesCreateResponse CreateSecurityGroupRules(ctx).SecurityGroupRulesCreateRequest(securityGroupRulesCreateRequest).Execute()

Create Security Group Rules



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
	securityGroupRulesCreateRequest := *openapiclient.NewSecurityGroupRulesCreateRequest("SecurityGroupId_example", []openapiclient.SecurityGroupRuleCreateRequestParam{*openapiclient.NewSecurityGroupRuleCreateRequestParam("Direction_example", "Protocol_example")}) // SecurityGroupRulesCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupV1SecurityGroupRuleApiAPI.CreateSecurityGroupRules(context.Background()).SecurityGroupRulesCreateRequest(securityGroupRulesCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupV1SecurityGroupRuleApiAPI.CreateSecurityGroupRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityGroupRules`: SecurityGroupRulesCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupV1SecurityGroupRuleApiAPI.CreateSecurityGroupRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityGroupRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityGroupRulesCreateRequest** | [**SecurityGroupRulesCreateRequest**](SecurityGroupRulesCreateRequest.md) |  | 

### Return type

[**SecurityGroupRulesCreateResponse**](SecurityGroupRulesCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityGroupRule

> DeleteSecurityGroupRule(ctx, securityGroupRuleId).Execute()

Delete Security Group Rule



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
	securityGroupRuleId := "securityGroupRuleId_example" // string | Security Group Rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityGroupV1SecurityGroupRuleApiAPI.DeleteSecurityGroupRule(context.Background(), securityGroupRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupV1SecurityGroupRuleApiAPI.DeleteSecurityGroupRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupRuleId** | **string** | Security Group Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityGroupRuleRequest struct via the builder pattern


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


## DownloadSecurityGroupRules

> DownloadSecurityGroupRules(ctx, securityGroupId).Execute()

Download Security Group Rules



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
	r, err := apiClient.SecurityGroupV1SecurityGroupRuleApiAPI.DownloadSecurityGroupRules(context.Background(), securityGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupV1SecurityGroupRuleApiAPI.DownloadSecurityGroupRules``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDownloadSecurityGroupRulesRequest struct via the builder pattern


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


## ListSecurityGroupRules

> SecurityGroupRuleListResponse ListSecurityGroupRules(ctx).SecurityGroupId(securityGroupId).Size(size).Page(page).Sort(sort).Id(id).RemoteIpPrefix(remoteIpPrefix).RemoteGroupId(remoteGroupId).Description(description).Direction(direction).Service(service).Execute()

List Security Group Rules



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
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	id := "id_example" // string | Security Group Rule ID (optional)
	remoteIpPrefix := "remoteIpPrefix_example" // string | Security Group Rule Remote IP Address (optional)
	remoteGroupId := "remoteGroupId_example" // string | Security Group Rule Remote Group ID (optional)
	description := "description_example" // string | Security Group Rule Description (optional)
	direction := "direction_example" // string | Security Group Rule Direction (Ingress, Egress) (optional)
	service := "service_example" // string | Security Group Rule Service (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupV1SecurityGroupRuleApiAPI.ListSecurityGroupRules(context.Background()).SecurityGroupId(securityGroupId).Size(size).Page(page).Sort(sort).Id(id).RemoteIpPrefix(remoteIpPrefix).RemoteGroupId(remoteGroupId).Description(description).Direction(direction).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupV1SecurityGroupRuleApiAPI.ListSecurityGroupRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityGroupRules`: SecurityGroupRuleListResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupV1SecurityGroupRuleApiAPI.ListSecurityGroupRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityGroupRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityGroupId** | **string** | Security Group ID | 
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **id** | **string** | Security Group Rule ID | 
 **remoteIpPrefix** | **string** | Security Group Rule Remote IP Address | 
 **remoteGroupId** | **string** | Security Group Rule Remote Group ID | 
 **description** | **string** | Security Group Rule Description | 
 **direction** | **string** | Security Group Rule Direction (Ingress, Egress) | 
 **service** | **string** | Security Group Rule Service | 

### Return type

[**SecurityGroupRuleListResponse**](SecurityGroupRuleListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowSecurityGroupRule

> SecurityGroupRuleShowResponse ShowSecurityGroupRule(ctx, securityGroupRuleId).Fields(fields).Execute()

Show Security Group Rule



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
	securityGroupRuleId := "securityGroupRuleId_example" // string | Security Group Rule ID
	fields := *openapiclient.NewFields() // Fields | Field Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupV1SecurityGroupRuleApiAPI.ShowSecurityGroupRule(context.Background(), securityGroupRuleId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupV1SecurityGroupRuleApiAPI.ShowSecurityGroupRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowSecurityGroupRule`: SecurityGroupRuleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupV1SecurityGroupRuleApiAPI.ShowSecurityGroupRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupRuleId** | **string** | Security Group Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowSecurityGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**Fields**](Fields.md) | Field Name | 

### Return type

[**SecurityGroupRuleShowResponse**](SecurityGroupRuleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

