# \VirtualserverV1KeypairsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKeypair**](VirtualserverV1KeypairsAPI.md#CreateKeypair) | **Post** /v1/keypairs | Create Keypair
[**DeleteKeypair**](VirtualserverV1KeypairsAPI.md#DeleteKeypair) | **Delete** /v1/keypairs/{keypair_name} | Delete Keypair
[**ListKeypairs**](VirtualserverV1KeypairsAPI.md#ListKeypairs) | **Get** /v1/keypairs | List Keypairs
[**ShowKeypair**](VirtualserverV1KeypairsAPI.md#ShowKeypair) | **Get** /v1/keypairs/{keypair_name} | Show Keypair



## CreateKeypair

> KeypairCreateResponse CreateKeypair(ctx).KeypairCreateRequest(keypairCreateRequest).Execute()

Create Keypair



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
	keypairCreateRequest := *openapiclient.NewKeypairCreateRequest("keypair-d20a3d59-9433-4b79-8726-20b431d89c78") // KeypairCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1KeypairsAPI.CreateKeypair(context.Background()).KeypairCreateRequest(keypairCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1KeypairsAPI.CreateKeypair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKeypair`: KeypairCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1KeypairsAPI.CreateKeypair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeypairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keypairCreateRequest** | [**KeypairCreateRequest**](KeypairCreateRequest.md) |  | 

### Return type

[**KeypairCreateResponse**](KeypairCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKeypair

> DeleteKeypair(ctx, keypairName).Execute()

Delete Keypair



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
	keypairName := "keypair-d20a3d59-9433-4b79-8726-20b431d89c78" // string | Keypair name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1KeypairsAPI.DeleteKeypair(context.Background(), keypairName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1KeypairsAPI.DeleteKeypair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keypairName** | **string** | Keypair name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeypairRequest struct via the builder pattern


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


## ListKeypairs

> KeypairListResponse ListKeypairs(ctx).Limit(limit).Marker(marker).Execute()

List Keypairs



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
	limit := int32(20) // int32 | Limit (optional)
	marker := "keypair-5d935425-31d5-48a7-a0f1-e76e9813f2c3" // string | Marker (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1KeypairsAPI.ListKeypairs(context.Background()).Limit(limit).Marker(marker).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1KeypairsAPI.ListKeypairs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeypairs`: KeypairListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1KeypairsAPI.ListKeypairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKeypairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit | 
 **marker** | **string** | Marker | 

### Return type

[**KeypairListResponse**](KeypairListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowKeypair

> KeypairShowResponse ShowKeypair(ctx, keypairName).Execute()

Show Keypair



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
	keypairName := "keypair-d20a3d59-9433-4b79-8726-20b431d89c78" // string | Keypair name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1KeypairsAPI.ShowKeypair(context.Background(), keypairName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1KeypairsAPI.ShowKeypair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowKeypair`: KeypairShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1KeypairsAPI.ShowKeypair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keypairName** | **string** | Keypair name | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowKeypairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KeypairShowResponse**](KeypairShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

