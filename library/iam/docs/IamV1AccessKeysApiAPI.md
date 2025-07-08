# \IamV1AccessKeysApiAPI

All URIs are relative to *https://iam.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessKeyCreate**](IamV1AccessKeysApiAPI.md#AccessKeyCreate) | **Post** /v1/access-keys | Create an access key
[**AccessKeyDelete**](IamV1AccessKeysApiAPI.md#AccessKeyDelete) | **Delete** /v1/access-keys/{access_key_id} | Remove the access key
[**AccessKeyList**](IamV1AccessKeysApiAPI.md#AccessKeyList) | **Get** /v1/access-keys | List access keys
[**AccessKeyRefresh**](IamV1AccessKeysApiAPI.md#AccessKeyRefresh) | **Post** /v1/access-keys/refresh | Refresh access key token
[**AccessKeySendTemporaryOtp**](IamV1AccessKeysApiAPI.md#AccessKeySendTemporaryOtp) | **Post** /v1/access-keys/send-otp | Send Temporary Access Key OTP
[**AccessKeyShow**](IamV1AccessKeysApiAPI.md#AccessKeyShow) | **Get** /v1/access-keys/{access_key_id} | Get an access key



## AccessKeyCreate

> AccessKeyResponse AccessKeyCreate(ctx).AccessKeyCreateRequest(accessKeyCreateRequest).Execute()

Create an access key



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
	accessKeyCreateRequest := *openapiclient.NewAccessKeyCreateRequest(openapiclient.AccessKeyTypeEnum("PERMANENT")) // AccessKeyCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1AccessKeysApiAPI.AccessKeyCreate(context.Background()).AccessKeyCreateRequest(accessKeyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1AccessKeysApiAPI.AccessKeyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessKeyCreate`: AccessKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1AccessKeysApiAPI.AccessKeyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessKeyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessKeyCreateRequest** | [**AccessKeyCreateRequest**](AccessKeyCreateRequest.md) |  | 

### Return type

[**AccessKeyResponse**](AccessKeyResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessKeyDelete

> AccessKeyDelete(ctx, accessKeyId).Execute()

Remove the access key



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
	accessKeyId := "accessKeyId_example" // string | Access key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1AccessKeysApiAPI.AccessKeyDelete(context.Background(), accessKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1AccessKeysApiAPI.AccessKeyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessKeyId** | **string** | Access key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessKeyDeleteRequest struct via the builder pattern


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


## AccessKeyList

> ListAccessKeyResponse AccessKeyList(ctx).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).AccountId(accountId).AccessKeyType(accessKeyType).ParentAccessKeyId(parentAccessKeyId).Execute()

List access keys



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
	withCount := "true" // string | with count (optional)
	limit := int32(20) // int32 | limit (optional)
	marker := "607e0938521643b5b4b266f343fae693" // string | marker (optional)
	sort := "created_at:desc" // string | sort (optional)
	accountId := "accountId_example" // string | Project ID (optional)
	accessKeyType := "accessKeyType_example" // string | Access key type (optional)
	parentAccessKeyId := "parentAccessKeyId_example" // string | Parent access key ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1AccessKeysApiAPI.AccessKeyList(context.Background()).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).AccountId(accountId).AccessKeyType(accessKeyType).ParentAccessKeyId(parentAccessKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1AccessKeysApiAPI.AccessKeyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessKeyList`: ListAccessKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1AccessKeysApiAPI.AccessKeyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessKeyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 
 **accountId** | **string** | Project ID | 
 **accessKeyType** | **string** | Access key type | 
 **parentAccessKeyId** | **string** | Parent access key ID | 

### Return type

[**ListAccessKeyResponse**](ListAccessKeyResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessKeyRefresh

> AccessKeyResponse AccessKeyRefresh(ctx).RefreshAccessKeyRequest(refreshAccessKeyRequest).Execute()

Refresh access key token



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
	refreshAccessKeyRequest := *openapiclient.NewRefreshAccessKeyRequest() // RefreshAccessKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1AccessKeysApiAPI.AccessKeyRefresh(context.Background()).RefreshAccessKeyRequest(refreshAccessKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1AccessKeysApiAPI.AccessKeyRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessKeyRefresh`: AccessKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1AccessKeysApiAPI.AccessKeyRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessKeyRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshAccessKeyRequest** | [**RefreshAccessKeyRequest**](RefreshAccessKeyRequest.md) |  | 

### Return type

[**AccessKeyResponse**](AccessKeyResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessKeySendTemporaryOtp

> AccessKeySendTemporaryOtp(ctx).AccessKeyOtpRequest(accessKeyOtpRequest).Execute()

Send Temporary Access Key OTP



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
	accessKeyOtpRequest := *openapiclient.NewAccessKeyOtpRequest() // AccessKeyOtpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamV1AccessKeysApiAPI.AccessKeySendTemporaryOtp(context.Background()).AccessKeyOtpRequest(accessKeyOtpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1AccessKeysApiAPI.AccessKeySendTemporaryOtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessKeySendTemporaryOtpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessKeyOtpRequest** | [**AccessKeyOtpRequest**](AccessKeyOtpRequest.md) |  | 

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


## AccessKeyShow

> AccessKeyResponse AccessKeyShow(ctx, accessKeyId).Execute()

Get an access key



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
	accessKeyId := "accessKeyId_example" // string | Access key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1AccessKeysApiAPI.AccessKeyShow(context.Background(), accessKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1AccessKeysApiAPI.AccessKeyShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessKeyShow`: AccessKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1AccessKeysApiAPI.AccessKeyShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessKeyId** | **string** | Access key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessKeyShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessKeyResponse**](AccessKeyResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

