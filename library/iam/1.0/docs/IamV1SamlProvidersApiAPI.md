# \IamV1SamlProvidersApiAPI

All URIs are relative to *https://scp-iam-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSamlProvider**](IamV1SamlProvidersApiAPI.md#CreateSamlProvider) | **Post** /v1/saml-providers | Create a saml provider
[**DeleteSamlProviders**](IamV1SamlProvidersApiAPI.md#DeleteSamlProviders) | **Delete** /v1/saml-providers/bulk | Remove saml providers
[**ListSamlProvider**](IamV1SamlProvidersApiAPI.md#ListSamlProvider) | **Get** /v1/saml-providers | List saml providers
[**SetSamlProvider**](IamV1SamlProvidersApiAPI.md#SetSamlProvider) | **Put** /v1/saml-providers/{saml_provider_id} | Update saml provider information
[**ShowSamlProvider**](IamV1SamlProvidersApiAPI.md#ShowSamlProvider) | **Get** /v1/saml-providers/{saml_provider_id} | Get saml provider information



## CreateSamlProvider

> SamlProviderCreateResponse CreateSamlProvider(ctx).AccountId(accountId).FederationType(federationType).SamlProviderName(samlProviderName).Description(description).File(file).Tags(tags).Execute()

Create a saml provider



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
	accountId := "accountId_example" // string | Account ID
	federationType := "federationType_example" // string | 자격증명공급자 유형
	samlProviderName := "samlProviderName_example" // string | 자격증명공급자 이름
	description := "description_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File | 메타데이터 파일 (optional)
	tags := []map[string]string{map[string]string{"key": "Inner_example"}} // []map[string]string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1SamlProvidersApiAPI.CreateSamlProvider(context.Background()).AccountId(accountId).FederationType(federationType).SamlProviderName(samlProviderName).Description(description).File(file).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1SamlProvidersApiAPI.CreateSamlProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSamlProvider`: SamlProviderCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1SamlProvidersApiAPI.CreateSamlProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSamlProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Account ID | 
 **federationType** | **string** | 자격증명공급자 유형 | 
 **samlProviderName** | **string** | 자격증명공급자 이름 | 
 **description** | **string** |  | 
 **file** | ***os.File** | 메타데이터 파일 | 
 **tags** | **[]map[string]string** |  | 

### Return type

[**SamlProviderCreateResponse**](SamlProviderCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSamlProviders

> DeleteSamlProviders(ctx).ListSamlProviderRemoveRequest(listSamlProviderRemoveRequest).Execute()

Remove saml providers



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
	listSamlProviderRemoveRequest := *openapiclient.NewListSamlProviderRemoveRequest() // ListSamlProviderRemoveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1SamlProvidersApiAPI.DeleteSamlProviders(context.Background()).ListSamlProviderRemoveRequest(listSamlProviderRemoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1SamlProvidersApiAPI.DeleteSamlProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSamlProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listSamlProviderRemoveRequest** | [**ListSamlProviderRemoveRequest**](ListSamlProviderRemoveRequest.md) |  | 

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


## ListSamlProvider

> ListSamlProviderResponse ListSamlProvider(ctx).Size(size).Page(page).Sort(sort).FederationType(federationType).SamlProviderName(samlProviderName).Execute()

List saml providers



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
	federationType := "federationType_example" // string | 자격증명공급자 유형 (optional)
	samlProviderName := "samlProviderName_example" // string | 자격증명공급자 이름 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1SamlProvidersApiAPI.ListSamlProvider(context.Background()).Size(size).Page(page).Sort(sort).FederationType(federationType).SamlProviderName(samlProviderName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1SamlProvidersApiAPI.ListSamlProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSamlProvider`: ListSamlProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1SamlProvidersApiAPI.ListSamlProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSamlProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **federationType** | **string** | 자격증명공급자 유형 | 
 **samlProviderName** | **string** | 자격증명공급자 이름 | 

### Return type

[**ListSamlProviderResponse**](ListSamlProviderResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSamlProvider

> SamlProviderDetailResponse SetSamlProvider(ctx, samlProviderId).Description(description).File(file).SamlProviderName(samlProviderName).Tags(tags).Execute()

Update saml provider information



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
	samlProviderId := "samlProviderId_example" // string | 자격증명공급자 ID
	description := "description_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File | 메타데이터 파일 (optional)
	samlProviderName := "samlProviderName_example" // string |  (optional)
	tags := []map[string]string{map[string]string{"key": "Inner_example"}} // []map[string]string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1SamlProvidersApiAPI.SetSamlProvider(context.Background(), samlProviderId).Description(description).File(file).SamlProviderName(samlProviderName).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1SamlProvidersApiAPI.SetSamlProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSamlProvider`: SamlProviderDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1SamlProvidersApiAPI.SetSamlProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**samlProviderId** | **string** | 자격증명공급자 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSamlProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **string** |  | 
 **file** | ***os.File** | 메타데이터 파일 | 
 **samlProviderName** | **string** |  | 
 **tags** | **[]map[string]string** |  | 

### Return type

[**SamlProviderDetailResponse**](SamlProviderDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowSamlProvider

> SamlProviderDetailResponse ShowSamlProvider(ctx, samlProviderId).Execute()

Get saml provider information



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
	samlProviderId := "samlProviderId_example" // string | 자격증명공급자 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1SamlProvidersApiAPI.ShowSamlProvider(context.Background(), samlProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1SamlProvidersApiAPI.ShowSamlProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowSamlProvider`: SamlProviderDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1SamlProvidersApiAPI.ShowSamlProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**samlProviderId** | **string** | 자격증명공급자 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowSamlProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SamlProviderDetailResponse**](SamlProviderDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

