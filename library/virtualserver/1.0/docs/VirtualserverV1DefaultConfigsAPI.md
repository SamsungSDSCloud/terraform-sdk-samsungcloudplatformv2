# \VirtualserverV1DefaultConfigsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDefaultConfigs**](VirtualserverV1DefaultConfigsAPI.md#ListDefaultConfigs) | **Get** /v1/default-configs | List Default Configs



## ListDefaultConfigs

> DefaultConfigListResponse ListDefaultConfigs(ctx).Name(name).Category(category).Execute()

List Default Configs



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
	name := "key" // string | Name (optional)
	category := "AUTO_SCALING_GROUP" // string | Category (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1DefaultConfigsAPI.ListDefaultConfigs(context.Background()).Name(name).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1DefaultConfigsAPI.ListDefaultConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDefaultConfigs`: DefaultConfigListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1DefaultConfigsAPI.ListDefaultConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDefaultConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name | 
 **category** | **string** | Category | 

### Return type

[**DefaultConfigListResponse**](DefaultConfigListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

