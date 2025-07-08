# \VirtualserverV1VolumeTransferApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptVolumeTransfer**](VirtualserverV1VolumeTransferApiAPI.md#AcceptVolumeTransfer) | **Post** /v1/volume-transfer/{transfer_id}/accept | Accept a volume transfer
[**CreateVolumeTransfer**](VirtualserverV1VolumeTransferApiAPI.md#CreateVolumeTransfer) | **Post** /v1/volume-transfer | Create a volume transfer
[**DeleteVolumeTransfer**](VirtualserverV1VolumeTransferApiAPI.md#DeleteVolumeTransfer) | **Delete** /v1/volume-transfer/{transfer_id} | Delete a volume transfer
[**ListVolumeTransfer**](VirtualserverV1VolumeTransferApiAPI.md#ListVolumeTransfer) | **Get** /v1/volume-transfer | List volume transfers
[**ShowVolumeTransfer**](VirtualserverV1VolumeTransferApiAPI.md#ShowVolumeTransfer) | **Get** /v1/volume-transfer/{transfer_id} | Show volume transfer detail



## AcceptVolumeTransfer

> VolumeTransferAcceptResponse AcceptVolumeTransfer(ctx, transferId).VolumeTransferAccessRequest(volumeTransferAccessRequest).Execute()

Accept a volume transfer



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
	transferId := "cceed636-1f1e-4bb0-b85c-4e5b9e0bf790" // string | Transfer ID
	volumeTransferAccessRequest := *openapiclient.NewVolumeTransferAccessRequest("9a7ad649f1cd75e7") // VolumeTransferAccessRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumeTransferApiAPI.AcceptVolumeTransfer(context.Background(), transferId).VolumeTransferAccessRequest(volumeTransferAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumeTransferApiAPI.AcceptVolumeTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptVolumeTransfer`: VolumeTransferAcceptResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumeTransferApiAPI.AcceptVolumeTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | Transfer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptVolumeTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeTransferAccessRequest** | [**VolumeTransferAccessRequest**](VolumeTransferAccessRequest.md) |  | 

### Return type

[**VolumeTransferAcceptResponse**](VolumeTransferAcceptResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolumeTransfer

> VolumeTransferCreateResponse CreateVolumeTransfer(ctx).VolumeTransferRequest(volumeTransferRequest).Execute()

Create a volume transfer



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
	volumeTransferRequest := *openapiclient.NewVolumeTransferRequest("760fee21-42b7-4717-adb7-511b31f8af74") // VolumeTransferRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumeTransferApiAPI.CreateVolumeTransfer(context.Background()).VolumeTransferRequest(volumeTransferRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumeTransferApiAPI.CreateVolumeTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolumeTransfer`: VolumeTransferCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumeTransferApiAPI.CreateVolumeTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeTransferRequest** | [**VolumeTransferRequest**](VolumeTransferRequest.md) |  | 

### Return type

[**VolumeTransferCreateResponse**](VolumeTransferCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolumeTransfer

> DeleteVolumeTransfer(ctx, transferId).Execute()

Delete a volume transfer



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
	transferId := "cceed636-1f1e-4bb0-b85c-4e5b9e0bf790" // string | Transfer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1VolumeTransferApiAPI.DeleteVolumeTransfer(context.Background(), transferId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumeTransferApiAPI.DeleteVolumeTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | Transfer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeTransferRequest struct via the builder pattern


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


## ListVolumeTransfer

> VolumeTransferListResponse ListVolumeTransfer(ctx).Execute()

List volume transfers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumeTransferApiAPI.ListVolumeTransfer(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumeTransferApiAPI.ListVolumeTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumeTransfer`: VolumeTransferListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumeTransferApiAPI.ListVolumeTransfer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeTransferRequest struct via the builder pattern


### Return type

[**VolumeTransferListResponse**](VolumeTransferListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVolumeTransfer

> VolumeTransferShowResponse ShowVolumeTransfer(ctx, transferId).Execute()

Show volume transfer detail



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
	transferId := "cceed636-1f1e-4bb0-b85c-4e5b9e0bf790" // string | Transfer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumeTransferApiAPI.ShowVolumeTransfer(context.Background(), transferId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumeTransferApiAPI.ShowVolumeTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVolumeTransfer`: VolumeTransferShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumeTransferApiAPI.ShowVolumeTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | Transfer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVolumeTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeTransferShowResponse**](VolumeTransferShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

