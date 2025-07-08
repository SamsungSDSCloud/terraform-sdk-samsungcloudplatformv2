# \DirectConnectV1DirectConnectsApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDirectConnect**](DirectConnectV1DirectConnectsApiAPI.md#CreateDirectConnect) | **Post** /v1/direct-connects | Create Direct Connect
[**DeleteDirectConnect**](DirectConnectV1DirectConnectsApiAPI.md#DeleteDirectConnect) | **Delete** /v1/direct-connects/{direct_connect_id} | Delete Direct Connect
[**ListDirectConnects**](DirectConnectV1DirectConnectsApiAPI.md#ListDirectConnects) | **Get** /v1/direct-connects | List Direct Connects
[**SetDirectConnect**](DirectConnectV1DirectConnectsApiAPI.md#SetDirectConnect) | **Put** /v1/direct-connects/{direct_connect_id} | Set Direct Connect
[**ShowDirectConnect**](DirectConnectV1DirectConnectsApiAPI.md#ShowDirectConnect) | **Get** /v1/direct-connects/{direct_connect_id} | Show Direct Connect



## CreateDirectConnect

> DirectConnectShowResponse CreateDirectConnect(ctx).DirectConnectCreateRequest(directConnectCreateRequest).Execute()

Create Direct Connect



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
	directConnectCreateRequest := *openapiclient.NewDirectConnectCreateRequest(int32(123), "Name_example", "VpcId_example") // DirectConnectCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectConnectV1DirectConnectsApiAPI.CreateDirectConnect(context.Background()).DirectConnectCreateRequest(directConnectCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectConnectV1DirectConnectsApiAPI.CreateDirectConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDirectConnect`: DirectConnectShowResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectConnectV1DirectConnectsApiAPI.CreateDirectConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDirectConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directConnectCreateRequest** | [**DirectConnectCreateRequest**](DirectConnectCreateRequest.md) |  | 

### Return type

[**DirectConnectShowResponse**](DirectConnectShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDirectConnect

> DeleteDirectConnect(ctx, directConnectId).Execute()

Delete Direct Connect



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
	directConnectId := "directConnectId_example" // string | Direct Connect ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DirectConnectV1DirectConnectsApiAPI.DeleteDirectConnect(context.Background(), directConnectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectConnectV1DirectConnectsApiAPI.DeleteDirectConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directConnectId** | **string** | Direct Connect ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDirectConnectRequest struct via the builder pattern


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


## ListDirectConnects

> DirectConnectListResponse ListDirectConnects(ctx).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).Name(name).State(state).VpcId(vpcId).VpcName(vpcName).Execute()

List Direct Connects



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
	id := "id_example" // string | Direct Connect ID (optional) (default to "")
	name := "name_example" // string | Direct Connect Name (optional) (default to "")
	state := openapiclient.DirectConnectState("CREATING") // DirectConnectState | State (optional) (default to "")
	vpcId := "vpcId_example" // string | VPC Id (optional) (default to "")
	vpcName := "vpcName_example" // string | VPC Name (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectConnectV1DirectConnectsApiAPI.ListDirectConnects(context.Background()).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).Name(name).State(state).VpcId(vpcId).VpcName(vpcName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectConnectV1DirectConnectsApiAPI.ListDirectConnects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDirectConnects`: DirectConnectListResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectConnectV1DirectConnectsApiAPI.ListDirectConnects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDirectConnectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 
 **id** | **string** | Direct Connect ID | [default to &quot;&quot;]
 **name** | **string** | Direct Connect Name | [default to &quot;&quot;]
 **state** | [**DirectConnectState**](DirectConnectState.md) | State | [default to &quot;&quot;]
 **vpcId** | **string** | VPC Id | [default to &quot;&quot;]
 **vpcName** | **string** | VPC Name | [default to &quot;&quot;]

### Return type

[**DirectConnectListResponse**](DirectConnectListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDirectConnect

> DirectConnectShowResponse SetDirectConnect(ctx, directConnectId).DirectConnectSetRequest(directConnectSetRequest).Execute()

Set Direct Connect



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
	directConnectId := "directConnectId_example" // string | Direct Connect ID
	directConnectSetRequest := *openapiclient.NewDirectConnectSetRequest("Description_example") // DirectConnectSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectConnectV1DirectConnectsApiAPI.SetDirectConnect(context.Background(), directConnectId).DirectConnectSetRequest(directConnectSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectConnectV1DirectConnectsApiAPI.SetDirectConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDirectConnect`: DirectConnectShowResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectConnectV1DirectConnectsApiAPI.SetDirectConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directConnectId** | **string** | Direct Connect ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDirectConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **directConnectSetRequest** | [**DirectConnectSetRequest**](DirectConnectSetRequest.md) |  | 

### Return type

[**DirectConnectShowResponse**](DirectConnectShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowDirectConnect

> DirectConnectShowResponse ShowDirectConnect(ctx, directConnectId).Execute()

Show Direct Connect



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
	directConnectId := "directConnectId_example" // string | Direct Connect ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectConnectV1DirectConnectsApiAPI.ShowDirectConnect(context.Background(), directConnectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectConnectV1DirectConnectsApiAPI.ShowDirectConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowDirectConnect`: DirectConnectShowResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectConnectV1DirectConnectsApiAPI.ShowDirectConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directConnectId** | **string** | Direct Connect ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowDirectConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DirectConnectShowResponse**](DirectConnectShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

