# \IamV1GroupsApiAPI

All URIs are relative to *https://iam.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupMember**](IamV1GroupsApiAPI.md#AddGroupMember) | **Post** /v1/groups/{group_id}/members | Add a user to a group
[**AddGroupPolicyBinding**](IamV1GroupsApiAPI.md#AddGroupPolicyBinding) | **Post** /v1/groups/{group_id}/policy-bindings | Assign policies to the group
[**CreateGroup**](IamV1GroupsApiAPI.md#CreateGroup) | **Post** /v1/groups | Create a group
[**DeleteGroup**](IamV1GroupsApiAPI.md#DeleteGroup) | **Delete** /v1/groups/{group_id} | Remove the Group
[**ListGroup**](IamV1GroupsApiAPI.md#ListGroup) | **Get** /v1/groups | List Groups
[**ListGroupMember**](IamV1GroupsApiAPI.md#ListGroupMember) | **Get** /v1/groups/{group_id}/members | Group Member list
[**ListGroupPolicyBinding**](IamV1GroupsApiAPI.md#ListGroupPolicyBinding) | **Get** /v1/groups/{group_id}/policy-bindings | Show group policy
[**RemoveGroupMember**](IamV1GroupsApiAPI.md#RemoveGroupMember) | **Delete** /v1/groups/{group_id}/members/{user_id} | Remove a User from a Group
[**RemoveGroupPolicyBinding**](IamV1GroupsApiAPI.md#RemoveGroupPolicyBinding) | **Delete** /v1/groups/{group_id}/policy-bindings/{policy_id} | Remove policy assignment from group
[**SetGroup**](IamV1GroupsApiAPI.md#SetGroup) | **Put** /v1/groups/{group_id} | Modify a Group
[**ShowGroup**](IamV1GroupsApiAPI.md#ShowGroup) | **Get** /v1/groups/{group_id} | Get a Group



## AddGroupMember

> GroupMemberShowResponse AddGroupMember(ctx, groupId).GroupMemberCreateRequest(groupMemberCreateRequest).Execute()

Add a user to a group



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
	groupId := "a946662dc4314dac93da413a32457459" // string | Group ID
	groupMemberCreateRequest := *openapiclient.NewGroupMemberCreateRequest("UserId_example") // GroupMemberCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1GroupsApiAPI.AddGroupMember(context.Background(), groupId).GroupMemberCreateRequest(groupMemberCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1GroupsApiAPI.AddGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupMember`: GroupMemberShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1GroupsApiAPI.AddGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupMemberCreateRequest** | [**GroupMemberCreateRequest**](GroupMemberCreateRequest.md) |  | 

### Return type

[**GroupMemberShowResponse**](GroupMemberShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGroupPolicyBinding

> GroupPolicyResponse AddGroupPolicyBinding(ctx, groupId).GroupPolicyBindingRequest(groupPolicyBindingRequest).Execute()

Assign policies to the group



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
	groupId := "a946662dc4314dac93da413a32457459" // string | Group ID
	groupPolicyBindingRequest := *openapiclient.NewGroupPolicyBindingRequest([]string{"PolicyIds_example"}) // GroupPolicyBindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1GroupsApiAPI.AddGroupPolicyBinding(context.Background(), groupId).GroupPolicyBindingRequest(groupPolicyBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1GroupsApiAPI.AddGroupPolicyBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupPolicyBinding`: GroupPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1GroupsApiAPI.AddGroupPolicyBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupPolicyBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupPolicyBindingRequest** | [**GroupPolicyBindingRequest**](GroupPolicyBindingRequest.md) |  | 

### Return type

[**GroupPolicyResponse**](GroupPolicyResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> GroupShowResponse CreateGroup(ctx).GroupCreateRequest(groupCreateRequest).Execute()

Create a group



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
	groupCreateRequest := *openapiclient.NewGroupCreateRequest("Description_example", "Test Group") // GroupCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1GroupsApiAPI.CreateGroup(context.Background()).GroupCreateRequest(groupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1GroupsApiAPI.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroup`: GroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1GroupsApiAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupCreateRequest** | [**GroupCreateRequest**](GroupCreateRequest.md) |  | 

### Return type

[**GroupShowResponse**](GroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, groupId).Execute()

Remove the Group



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
	groupId := "a946662dc4314dac93da413a32457459" // string | Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1GroupsApiAPI.DeleteGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1GroupsApiAPI.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## ListGroup

> GroupPageResponse ListGroup(ctx).Size(size).Page(page).Sort(sort).Name(name).Types(types).Ids(ids).HasMember(hasMember).HasRole(hasRole).IsCompleted(isCompleted).CreatorName(creatorName).CreatorEmail(creatorEmail).ModifierName(modifierName).ModifierEmail(modifierEmail).ExcludeUserId(excludeUserId).ExcludePolicyId(excludePolicyId).Execute()

List Groups



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
	name := "name_example" // string | Group 이름 (optional)
	types := "types_example" // string | Group Types (optional)
	ids := "ids_example" // string | Group IDs (optional)
	hasMember := "hasMember_example" // string | Group Member 존재 여부 (optional)
	hasRole := "hasRole_example" // string | Group Role 존재 여부 (optional)
	isCompleted := "isCompleted_example" // string | Group 완결 여부 (optional)
	creatorName := "creatorName_example" // string | Group 생성자 성, 이름 (optional)
	creatorEmail := "creatorEmail_example" // string | Group 생성자 Email (optional)
	modifierName := "modifierName_example" // string | Group 수정자 성, 이름 (optional)
	modifierEmail := "modifierEmail_example" // string | Group 수정자 Email (optional)
	excludeUserId := "excludeUserId_example" // string | 제외할 사용자 ID (optional)
	excludePolicyId := "excludePolicyId_example" // string | 제외할 정책 ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1GroupsApiAPI.ListGroup(context.Background()).Size(size).Page(page).Sort(sort).Name(name).Types(types).Ids(ids).HasMember(hasMember).HasRole(hasRole).IsCompleted(isCompleted).CreatorName(creatorName).CreatorEmail(creatorEmail).ModifierName(modifierName).ModifierEmail(modifierEmail).ExcludeUserId(excludeUserId).ExcludePolicyId(excludePolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1GroupsApiAPI.ListGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroup`: GroupPageResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1GroupsApiAPI.ListGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Group 이름 | 
 **types** | **string** | Group Types | 
 **ids** | **string** | Group IDs | 
 **hasMember** | **string** | Group Member 존재 여부 | 
 **hasRole** | **string** | Group Role 존재 여부 | 
 **isCompleted** | **string** | Group 완결 여부 | 
 **creatorName** | **string** | Group 생성자 성, 이름 | 
 **creatorEmail** | **string** | Group 생성자 Email | 
 **modifierName** | **string** | Group 수정자 성, 이름 | 
 **modifierEmail** | **string** | Group 수정자 Email | 
 **excludeUserId** | **string** | 제외할 사용자 ID | 
 **excludePolicyId** | **string** | 제외할 정책 ID | 

### Return type

[**GroupPageResponse**](GroupPageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupMember

> GroupMemberPageResponse ListGroupMember(ctx, groupId).Size(size).Page(page).Sort(sort).UserName(userName).UserEmail(userEmail).CreatorName(creatorName).CreatorEmail(creatorEmail).Execute()

Group Member list



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
	groupId := "a946662dc4314dac93da413a32457459" // string | Group ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	userName := "userName_example" // string | Group Member 성, 이름 (optional)
	userEmail := "userEmail_example" // string | Group Member Email (optional)
	creatorName := "creatorName_example" // string | Group Member 생성자 성, 이름 (optional)
	creatorEmail := "creatorEmail_example" // string | Group Member 생성자 Email (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1GroupsApiAPI.ListGroupMember(context.Background(), groupId).Size(size).Page(page).Sort(sort).UserName(userName).UserEmail(userEmail).CreatorName(creatorName).CreatorEmail(creatorEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1GroupsApiAPI.ListGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupMember`: GroupMemberPageResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1GroupsApiAPI.ListGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **userName** | **string** | Group Member 성, 이름 | 
 **userEmail** | **string** | Group Member Email | 
 **creatorName** | **string** | Group Member 생성자 성, 이름 | 
 **creatorEmail** | **string** | Group Member 생성자 Email | 

### Return type

[**GroupMemberPageResponse**](GroupMemberPageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupPolicyBinding

> GroupPolicyPageResponse ListGroupPolicyBinding(ctx, groupId).Size(size).Page(page).Sort(sort).PolicyId(policyId).PolicyName(policyName).PolicyType(policyType).Execute()

Show group policy



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
	groupId := "a946662dc4314dac93da413a32457459" // string | Group ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	policyId := "37f2e31ff86b415698d7e8eeafab445d" // string | Policy Id (optional)
	policyName := "ViewerAccess" // string | Policy Name (optional)
	policyType := *openapiclient.NewPolicyType() // PolicyType | Policy Type List (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1GroupsApiAPI.ListGroupPolicyBinding(context.Background(), groupId).Size(size).Page(page).Sort(sort).PolicyId(policyId).PolicyName(policyName).PolicyType(policyType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1GroupsApiAPI.ListGroupPolicyBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupPolicyBinding`: GroupPolicyPageResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1GroupsApiAPI.ListGroupPolicyBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupPolicyBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **policyId** | **string** | Policy Id | 
 **policyName** | **string** | Policy Name | 
 **policyType** | [**PolicyType**](PolicyType.md) | Policy Type List | 

### Return type

[**GroupPolicyPageResponse**](GroupPolicyPageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGroupMember

> RemoveGroupMember(ctx, groupId, userId).GroupMemberDeleteRequest(groupMemberDeleteRequest).Execute()

Remove a User from a Group



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
	groupId := "a946662dc4314dac93da413a32457459" // string | Group ID
	userId := "a946662dc4314dac93da413a32457459" // string | User ID
	groupMemberDeleteRequest := *openapiclient.NewGroupMemberDeleteRequest() // GroupMemberDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1GroupsApiAPI.RemoveGroupMember(context.Background(), groupId, userId).GroupMemberDeleteRequest(groupMemberDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1GroupsApiAPI.RemoveGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupMemberDeleteRequest** | [**GroupMemberDeleteRequest**](GroupMemberDeleteRequest.md) |  | 

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


## RemoveGroupPolicyBinding

> RemoveGroupPolicyBinding(ctx, groupId, policyId).Execute()

Remove policy assignment from group



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
	groupId := "a946662dc4314dac93da413a32457459" // string | Group ID
	policyId := "a946662dc4314dac93da413a32457459" // string | Policy ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1GroupsApiAPI.RemoveGroupPolicyBinding(context.Background(), groupId, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1GroupsApiAPI.RemoveGroupPolicyBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupPolicyBindingRequest struct via the builder pattern


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


## SetGroup

> GroupShowResponse SetGroup(ctx, groupId).GroupSetRequest(groupSetRequest).Execute()

Modify a Group



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
	groupId := "a946662dc4314dac93da413a32457459" // string | Group ID
	groupSetRequest := *openapiclient.NewGroupSetRequest("Description_example", "Name_example") // GroupSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1GroupsApiAPI.SetGroup(context.Background(), groupId).GroupSetRequest(groupSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1GroupsApiAPI.SetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroup`: GroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1GroupsApiAPI.SetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupSetRequest** | [**GroupSetRequest**](GroupSetRequest.md) |  | 

### Return type

[**GroupShowResponse**](GroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowGroup

> GroupShowResponse ShowGroup(ctx, groupId).Execute()

Get a Group



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
	groupId := "a946662dc4314dac93da413a32457459" // string | Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1GroupsApiAPI.ShowGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1GroupsApiAPI.ShowGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowGroup`: GroupShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1GroupsApiAPI.ShowGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupShowResponse**](GroupShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

