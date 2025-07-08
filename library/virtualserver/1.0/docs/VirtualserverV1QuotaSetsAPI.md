# \VirtualserverV1QuotaSetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowQuotaSet**](VirtualserverV1QuotaSetsAPI.md#ShowQuotaSet) | **Get** /v1/quota-sets | Show Server Quota Set
[**ShowVolumeQuotaSet**](VirtualserverV1QuotaSetsAPI.md#ShowVolumeQuotaSet) | **Get** /v1/volumes/quota-sets | Show Volume Quota Set



## ShowQuotaSet

> ServerQuotaSet ShowQuotaSet(ctx).Execute()

Show Server Quota Set



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
	resp, r, err := apiClient.VirtualserverV1QuotaSetsAPI.ShowQuotaSet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1QuotaSetsAPI.ShowQuotaSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowQuotaSet`: ServerQuotaSet
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1QuotaSetsAPI.ShowQuotaSet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowQuotaSetRequest struct via the builder pattern


### Return type

[**ServerQuotaSet**](ServerQuotaSet.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVolumeQuotaSet

> VolumeQuotaSet ShowVolumeQuotaSet(ctx).Execute()

Show Volume Quota Set



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
	resp, r, err := apiClient.VirtualserverV1QuotaSetsAPI.ShowVolumeQuotaSet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1QuotaSetsAPI.ShowVolumeQuotaSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVolumeQuotaSet`: VolumeQuotaSet
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1QuotaSetsAPI.ShowVolumeQuotaSet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowVolumeQuotaSetRequest struct via the builder pattern


### Return type

[**VolumeQuotaSet**](VolumeQuotaSet.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

