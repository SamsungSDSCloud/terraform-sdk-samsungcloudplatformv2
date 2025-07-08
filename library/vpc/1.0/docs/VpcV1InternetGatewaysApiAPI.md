# \VpcV1InternetGatewaysApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInternetGateway**](VpcV1InternetGatewaysApiAPI.md#CreateInternetGateway) | **Post** /v1/internet-gateways | Create Internet Gateway
[**DeleteInternetGateway**](VpcV1InternetGatewaysApiAPI.md#DeleteInternetGateway) | **Delete** /v1/internet-gateways/{internet_gateway_id} | Delete Internet Gateway
[**ListCreatableTypes**](VpcV1InternetGatewaysApiAPI.md#ListCreatableTypes) | **Get** /v1/internet-gateways/creatable-types | ListCreatableTypes
[**ListCreatableVpcs**](VpcV1InternetGatewaysApiAPI.md#ListCreatableVpcs) | **Get** /v1/internet-gateways/creatable-vpcs | ListCreatableVpcs
[**ListInternetGateways**](VpcV1InternetGatewaysApiAPI.md#ListInternetGateways) | **Get** /v1/internet-gateways | List Internet Gateways
[**SetInternetGateway**](VpcV1InternetGatewaysApiAPI.md#SetInternetGateway) | **Put** /v1/internet-gateways/{internet_gateway_id} | Set Internet Gateway
[**ShowInternetGateway**](VpcV1InternetGatewaysApiAPI.md#ShowInternetGateway) | **Get** /v1/internet-gateways/{internet_gateway_id} | Show Internet Gateway



## CreateInternetGateway

> InternetGatewayShowResponse CreateInternetGateway(ctx).InternetGatewayCreateRequest(internetGatewayCreateRequest).Execute()

Create Internet Gateway



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
	internetGatewayCreateRequest := *openapiclient.NewInternetGatewayCreateRequest(openapiclient.InternetGatewayType("IGW"), "VpcId_example") // InternetGatewayCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1InternetGatewaysApiAPI.CreateInternetGateway(context.Background()).InternetGatewayCreateRequest(internetGatewayCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1InternetGatewaysApiAPI.CreateInternetGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInternetGateway`: InternetGatewayShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1InternetGatewaysApiAPI.CreateInternetGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInternetGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **internetGatewayCreateRequest** | [**InternetGatewayCreateRequest**](InternetGatewayCreateRequest.md) |  | 

### Return type

[**InternetGatewayShowResponse**](InternetGatewayShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInternetGateway

> DeleteInternetGateway(ctx, internetGatewayId).Execute()

Delete Internet Gateway



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
	internetGatewayId := "internetGatewayId_example" // string | Internet Gateway ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1InternetGatewaysApiAPI.DeleteInternetGateway(context.Background(), internetGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1InternetGatewaysApiAPI.DeleteInternetGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internetGatewayId** | **string** | Internet Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInternetGatewayRequest struct via the builder pattern


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


## ListCreatableTypes

> CreatableTypeListResponse ListCreatableTypes(ctx).Execute()

ListCreatableTypes



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
	resp, r, err := apiClient.VpcV1InternetGatewaysApiAPI.ListCreatableTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1InternetGatewaysApiAPI.ListCreatableTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCreatableTypes`: CreatableTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1InternetGatewaysApiAPI.ListCreatableTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCreatableTypesRequest struct via the builder pattern


### Return type

[**CreatableTypeListResponse**](CreatableTypeListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCreatableVpcs

> CreatableVpcListResponse ListCreatableVpcs(ctx).VpcId(vpcId).Execute()

ListCreatableVpcs



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
	vpcId := "vpcId_example" // string | VPC Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1InternetGatewaysApiAPI.ListCreatableVpcs(context.Background()).VpcId(vpcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1InternetGatewaysApiAPI.ListCreatableVpcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCreatableVpcs`: CreatableVpcListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1InternetGatewaysApiAPI.ListCreatableVpcs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCreatableVpcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpcId** | **string** | VPC Id | 

### Return type

[**CreatableVpcListResponse**](CreatableVpcListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInternetGateways

> InternetGatewayListResponse ListInternetGateways(ctx).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).Name(name).Type_(type_).State(state).VpcId(vpcId).VpcName(vpcName).Execute()

List Internet Gateways



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
	id := "id_example" // string | Internet Gateway ID (optional)
	name := "name_example" // string | Internet Gateway Name (optional)
	type_ := openapiclient.InternetGatewayType("IGW") // InternetGatewayType | Internet Gateway Type (optional) (default to "")
	state := "state_example" // string | State (optional) (default to "")
	vpcId := "vpcId_example" // string | VPC Id (optional)
	vpcName := "vpcName_example" // string | VPC Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1InternetGatewaysApiAPI.ListInternetGateways(context.Background()).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).Name(name).Type_(type_).State(state).VpcId(vpcId).VpcName(vpcName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1InternetGatewaysApiAPI.ListInternetGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInternetGateways`: InternetGatewayListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1InternetGatewaysApiAPI.ListInternetGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInternetGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 
 **id** | **string** | Internet Gateway ID | 
 **name** | **string** | Internet Gateway Name | 
 **type_** | [**InternetGatewayType**](InternetGatewayType.md) | Internet Gateway Type | [default to &quot;&quot;]
 **state** | **string** | State | [default to &quot;&quot;]
 **vpcId** | **string** | VPC Id | 
 **vpcName** | **string** | VPC Name | 

### Return type

[**InternetGatewayListResponse**](InternetGatewayListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInternetGateway

> InternetGatewayShowResponse SetInternetGateway(ctx, internetGatewayId).InternetGatewaySetRequest(internetGatewaySetRequest).Execute()

Set Internet Gateway



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
	internetGatewayId := "internetGatewayId_example" // string | Internet Gateway ID
	internetGatewaySetRequest := *openapiclient.NewInternetGatewaySetRequest() // InternetGatewaySetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1InternetGatewaysApiAPI.SetInternetGateway(context.Background(), internetGatewayId).InternetGatewaySetRequest(internetGatewaySetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1InternetGatewaysApiAPI.SetInternetGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetInternetGateway`: InternetGatewayShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1InternetGatewaysApiAPI.SetInternetGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internetGatewayId** | **string** | Internet Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetInternetGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **internetGatewaySetRequest** | [**InternetGatewaySetRequest**](InternetGatewaySetRequest.md) |  | 

### Return type

[**InternetGatewayShowResponse**](InternetGatewayShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowInternetGateway

> InternetGatewayShowResponse ShowInternetGateway(ctx, internetGatewayId).Execute()

Show Internet Gateway



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
	internetGatewayId := "internetGatewayId_example" // string | Internet Gateway ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1InternetGatewaysApiAPI.ShowInternetGateway(context.Background(), internetGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1InternetGatewaysApiAPI.ShowInternetGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowInternetGateway`: InternetGatewayShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1InternetGatewaysApiAPI.ShowInternetGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internetGatewayId** | **string** | Internet Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowInternetGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternetGatewayShowResponse**](InternetGatewayShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

