# \SkeV1NodepoolsApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodepool**](SkeV1NodepoolsApiAPI.md#CreateNodepool) | **Post** /v1/nodepools | Create Nodepool
[**DeleteNodepool**](SkeV1NodepoolsApiAPI.md#DeleteNodepool) | **Delete** /v1/nodepools/{nodepool_id} | Delete Nodepool
[**ListNodepoolNodes**](SkeV1NodepoolsApiAPI.md#ListNodepoolNodes) | **Get** /v1/nodepools/{nodepool_id}/nodes | List Nodepool Nodes
[**ListNodepools**](SkeV1NodepoolsApiAPI.md#ListNodepools) | **Get** /v1/clusters/{cluster_id}/nodepools | List Nodepools
[**SetNodepool**](SkeV1NodepoolsApiAPI.md#SetNodepool) | **Put** /v1/nodepools/{nodepool_id} | Set Nodepool
[**SetNodepoolUpgrade**](SkeV1NodepoolsApiAPI.md#SetNodepoolUpgrade) | **Put** /v1/nodepools/{nodepool_id}/upgrade | Set Nodepool Upgrade
[**ShowNodepool**](SkeV1NodepoolsApiAPI.md#ShowNodepool) | **Get** /v1/nodepools/{nodepool_id} | Show Nodepool



## CreateNodepool

> NodepoolShowResponse CreateNodepool(ctx).NodepoolCreateRequest(nodepoolCreateRequest).Execute()

Create Nodepool



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
	nodepoolCreateRequest := *openapiclient.NewNodepoolCreateRequest("ClusterId_example", "ImageOs_example", "ImageOsVersion_example", false, false, "KeypairName_example", "KubernetesVersion_example", "Name_example", "ServerTypeId_example", int32(123), "VolumeTypeName_example") // NodepoolCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1NodepoolsApiAPI.CreateNodepool(context.Background()).NodepoolCreateRequest(nodepoolCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1NodepoolsApiAPI.CreateNodepool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNodepool`: NodepoolShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1NodepoolsApiAPI.CreateNodepool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodepoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodepoolCreateRequest** | [**NodepoolCreateRequest**](NodepoolCreateRequest.md) |  | 

### Return type

[**NodepoolShowResponse**](NodepoolShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNodepool

> AsyncResponse DeleteNodepool(ctx, nodepoolId).Execute()

Delete Nodepool



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
	nodepoolId := "nodepoolId_example" // string | Nodepool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1NodepoolsApiAPI.DeleteNodepool(context.Background(), nodepoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1NodepoolsApiAPI.DeleteNodepool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNodepool`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1NodepoolsApiAPI.DeleteNodepool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodepoolId** | **string** | Nodepool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodepoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodepoolNodes

> NodeListInNodepoolResponse ListNodepoolNodes(ctx, nodepoolId).Size(size).Page(page).Sort(sort).NodeName(nodeName).Execute()

List Nodepool Nodes



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
	nodepoolId := "nodepoolId_example" // string | Nodepool ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	nodeName := "nodeName_example" // string | Node Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1NodepoolsApiAPI.ListNodepoolNodes(context.Background(), nodepoolId).Size(size).Page(page).Sort(sort).NodeName(nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1NodepoolsApiAPI.ListNodepoolNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNodepoolNodes`: NodeListInNodepoolResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1NodepoolsApiAPI.ListNodepoolNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodepoolId** | **string** | Nodepool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodepoolNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **nodeName** | **string** | Node Name | 

### Return type

[**NodeListInNodepoolResponse**](NodeListInNodepoolResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodepools

> NodepoolListResponse ListNodepools(ctx, clusterId).Size(size).Page(page).Sort(sort).NodepoolName(nodepoolName).Execute()

List Nodepools



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
	clusterId := "10a599e031e749b7b260868f441e862b" // string | Cluster ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	nodepoolName := "nodepoolName_example" // string | Nodepool Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1NodepoolsApiAPI.ListNodepools(context.Background(), clusterId).Size(size).Page(page).Sort(sort).NodepoolName(nodepoolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1NodepoolsApiAPI.ListNodepools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNodepools`: NodepoolListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1NodepoolsApiAPI.ListNodepools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodepoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **nodepoolName** | **string** | Nodepool Name | 

### Return type

[**NodepoolListResponse**](NodepoolListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNodepool

> AsyncResponse SetNodepool(ctx, nodepoolId).NodepoolUpdateRequest(nodepoolUpdateRequest).Execute()

Set Nodepool



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
	nodepoolId := "nodepoolId_example" // string | Nodepool ID
	nodepoolUpdateRequest := *openapiclient.NewNodepoolUpdateRequest() // NodepoolUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1NodepoolsApiAPI.SetNodepool(context.Background(), nodepoolId).NodepoolUpdateRequest(nodepoolUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1NodepoolsApiAPI.SetNodepool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNodepool`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1NodepoolsApiAPI.SetNodepool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodepoolId** | **string** | Nodepool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNodepoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodepoolUpdateRequest** | [**NodepoolUpdateRequest**](NodepoolUpdateRequest.md) |  | 

### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNodepoolUpgrade

> AsyncResponse SetNodepoolUpgrade(ctx, nodepoolId).Execute()

Set Nodepool Upgrade



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
	nodepoolId := "nodepoolId_example" // string | Nodepool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1NodepoolsApiAPI.SetNodepoolUpgrade(context.Background(), nodepoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1NodepoolsApiAPI.SetNodepoolUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNodepoolUpgrade`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1NodepoolsApiAPI.SetNodepoolUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodepoolId** | **string** | Nodepool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNodepoolUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowNodepool

> NodepoolShowResponse ShowNodepool(ctx, nodepoolId).Execute()

Show Nodepool



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
	nodepoolId := "nodepoolId_example" // string | Nodepool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1NodepoolsApiAPI.ShowNodepool(context.Background(), nodepoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1NodepoolsApiAPI.ShowNodepool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowNodepool`: NodepoolShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1NodepoolsApiAPI.ShowNodepool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodepoolId** | **string** | Nodepool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowNodepoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NodepoolShowResponse**](NodepoolShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

