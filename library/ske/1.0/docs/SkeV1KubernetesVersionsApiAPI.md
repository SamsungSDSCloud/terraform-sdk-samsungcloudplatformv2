# \SkeV1KubernetesVersionsApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListKubernetesVersions**](SkeV1KubernetesVersionsApiAPI.md#ListKubernetesVersions) | **Get** /v1/kubernetes-versions | List Kubernetes Versions



## ListKubernetesVersions

> KubernetesVersionListResponse ListKubernetesVersions(ctx).Execute()

List Kubernetes Versions



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
	resp, r, err := apiClient.SkeV1KubernetesVersionsApiAPI.ListKubernetesVersions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1KubernetesVersionsApiAPI.ListKubernetesVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesVersions`: KubernetesVersionListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1KubernetesVersionsApiAPI.ListKubernetesVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesVersionsRequest struct via the builder pattern


### Return type

[**KubernetesVersionListResponse**](KubernetesVersionListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

