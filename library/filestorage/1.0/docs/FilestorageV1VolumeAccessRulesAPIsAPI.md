# \FilestorageV1VolumeAccessRulesAPIsAPI

All URIs are relative to *https://filestorage-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccessRules**](FilestorageV1VolumeAccessRulesAPIsAPI.md#ListAccessRules) | **Get** /v1/volumes/{volume_id}/access-rules | ListAccessRules
[**SetAccessRule**](FilestorageV1VolumeAccessRulesAPIsAPI.md#SetAccessRule) | **Put** /v1/volumes/{volume_id}/access-rules | SetAccessRule



## ListAccessRules

> VolumeObjectAccessRuleListResponse ListAccessRules(ctx, volumeId).Execute()

ListAccessRules



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
	volumeId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1VolumeAccessRulesAPIsAPI.ListAccessRules(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeAccessRulesAPIsAPI.ListAccessRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccessRules`: VolumeObjectAccessRuleListResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1VolumeAccessRulesAPIsAPI.ListAccessRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeObjectAccessRuleListResponse**](VolumeObjectAccessRuleListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAccessRule

> AccessRuleResponse SetAccessRule(ctx, volumeId).AccessRuleRequest(accessRuleRequest).Execute()

SetAccessRule



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
	volumeId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Volume ID
	accessRuleRequest := *openapiclient.NewAccessRuleRequest("add", "43fq3347-02q4-4aa8-ccf9-affe4917bb6f", "VM") // AccessRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1VolumeAccessRulesAPIsAPI.SetAccessRule(context.Background(), volumeId).AccessRuleRequest(accessRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeAccessRulesAPIsAPI.SetAccessRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAccessRule`: AccessRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1VolumeAccessRulesAPIsAPI.SetAccessRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAccessRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessRuleRequest** | [**AccessRuleRequest**](AccessRuleRequest.md) |  | 

### Return type

[**AccessRuleResponse**](AccessRuleResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

