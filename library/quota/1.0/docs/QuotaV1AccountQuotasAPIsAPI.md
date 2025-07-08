# \QuotaV1AccountQuotasAPIsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccountQuota**](QuotaV1AccountQuotasAPIsAPI.md#ListAccountQuota) | **Get** /v1/account-quotas | ListAccountQuota
[**ShowAccountQuota**](QuotaV1AccountQuotasAPIsAPI.md#ShowAccountQuota) | **Get** /v1/account-quotas/{account_quota_id} | ShowAccountQuota



## ListAccountQuota

> AccountQuotaListResponse ListAccountQuota(ctx).Size(size).Page(page).Sort(sort).RequestClass(requestClass).Service(service).QuotaItem(quotaItem).Description(description).Execute()

ListAccountQuota



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
	requestClass := "Account" // string | Classification of the quota request (e.g., Account, Region) (optional)
	service := "IAM" // string | Service Name (optional)
	quotaItem := "POLICY.MAX.COUNT" // string | Quota Item Name (optional)
	description := "Virtual Server disk max size" // string | Description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotaV1AccountQuotasAPIsAPI.ListAccountQuota(context.Background()).Size(size).Page(page).Sort(sort).RequestClass(requestClass).Service(service).QuotaItem(quotaItem).Description(description).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotaV1AccountQuotasAPIsAPI.ListAccountQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccountQuota`: AccountQuotaListResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotaV1AccountQuotasAPIsAPI.ListAccountQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **requestClass** | **string** | Classification of the quota request (e.g., Account, Region) | 
 **service** | **string** | Service Name | 
 **quotaItem** | **string** | Quota Item Name | 
 **description** | **string** | Description | 

### Return type

[**AccountQuotaListResponse**](AccountQuotaListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAccountQuota

> AccountQuotaShowResponse ShowAccountQuota(ctx, accountQuotaId).Execute()

ShowAccountQuota



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
	accountQuotaId := "138c2fc8c29a449dbfa8681f8f1d78e2" // string | Account Quota ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotaV1AccountQuotasAPIsAPI.ShowAccountQuota(context.Background(), accountQuotaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotaV1AccountQuotasAPIsAPI.ShowAccountQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowAccountQuota`: AccountQuotaShowResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotaV1AccountQuotasAPIsAPI.ShowAccountQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountQuotaId** | **string** | Account Quota ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowAccountQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountQuotaShowResponse**](AccountQuotaShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

