# \LoadbalancerV1LbCertificatesApiAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLoadbalancerCertificates**](LoadbalancerV1LbCertificatesApiAPI.md#ListLoadbalancerCertificates) | **Get** /v1/loadbalancers/certificates | ListLoadbalancerCertificates
[**ShowLoadbalancerCertificate**](LoadbalancerV1LbCertificatesApiAPI.md#ShowLoadbalancerCertificate) | **Get** /v1/loadbalancers/certificates/{lb_certificate_id} | ShowLoadbalancerCertificate



## ListLoadbalancerCertificates

> LbCertificateListResponse ListLoadbalancerCertificates(ctx).Execute()

ListLoadbalancerCertificates



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
	resp, r, err := apiClient.LoadbalancerV1LbCertificatesApiAPI.ListLoadbalancerCertificates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LbCertificatesApiAPI.ListLoadbalancerCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLoadbalancerCertificates`: LbCertificateListResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LbCertificatesApiAPI.ListLoadbalancerCertificates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLoadbalancerCertificatesRequest struct via the builder pattern


### Return type

[**LbCertificateListResponse**](LbCertificateListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowLoadbalancerCertificate

> LbCertificateShowResponse ShowLoadbalancerCertificate(ctx, lbCertificateId).Execute()

ShowLoadbalancerCertificate



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
	lbCertificateId := "lbCertificateId_example" // string | The ID of the certificate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LbCertificatesApiAPI.ShowLoadbalancerCertificate(context.Background(), lbCertificateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LbCertificatesApiAPI.ShowLoadbalancerCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowLoadbalancerCertificate`: LbCertificateShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LbCertificatesApiAPI.ShowLoadbalancerCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lbCertificateId** | **string** | The ID of the certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowLoadbalancerCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LbCertificateShowResponse**](LbCertificateShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

