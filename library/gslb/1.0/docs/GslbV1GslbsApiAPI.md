# \GslbV1GslbsApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckGslbNameAvailability**](GslbV1GslbsApiAPI.md#CheckGslbNameAvailability) | **Get** /v1/gslbs/check-duplication | CheckGslbNameAvailability
[**CreateGslb**](GslbV1GslbsApiAPI.md#CreateGslb) | **Post** /v1/gslbs | CreateGslb
[**DeleteGslb**](GslbV1GslbsApiAPI.md#DeleteGslb) | **Delete** /v1/gslbs/{gslb_id} | DeleteGslb
[**ListGslbIpRanges**](GslbV1GslbsApiAPI.md#ListGslbIpRanges) | **Get** /v1/gslbs/ip-ranges | ListGslbIpRanges
[**ListGslbRegions**](GslbV1GslbsApiAPI.md#ListGslbRegions) | **Get** /v1/gslbs/available-regions | ListGslbRegions
[**ListGslbs**](GslbV1GslbsApiAPI.md#ListGslbs) | **Get** /v1/gslbs | ListGslbs
[**SetGslb**](GslbV1GslbsApiAPI.md#SetGslb) | **Put** /v1/gslbs/{gslb_id} | SetGslb
[**SetGslbHealthCheck**](GslbV1GslbsApiAPI.md#SetGslbHealthCheck) | **Put** /v1/gslbs/{gslb_id}/health-check | SetGslbHealthCheck
[**ShowGslb**](GslbV1GslbsApiAPI.md#ShowGslb) | **Get** /v1/gslbs/{gslb_id} | ShowGslb
[**ShowGslbDefaultDomain**](GslbV1GslbsApiAPI.md#ShowGslbDefaultDomain) | **Get** /v1/gslbs/default-domain | ShowGslbDefaultDomain



## CheckGslbNameAvailability

> GslbCheckDuplicationResponse CheckGslbNameAvailability(ctx).Name(name).Execute()

CheckGslbNameAvailability



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
	name := "name_example" // string | The Name of the gslb.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GslbV1GslbsApiAPI.CheckGslbNameAvailability(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GslbV1GslbsApiAPI.CheckGslbNameAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckGslbNameAvailability`: GslbCheckDuplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `GslbV1GslbsApiAPI.CheckGslbNameAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckGslbNameAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The Name of the gslb. | 

### Return type

[**GslbCheckDuplicationResponse**](GslbCheckDuplicationResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGslb

> GslbShowResponse CreateGslb(ctx).GslbCreateRequest(gslbCreateRequest).Execute()

CreateGslb



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
	gslbCreateRequest := *openapiclient.NewGslbCreateRequest("Algorithm_example", "EnvUsage_example", "Name_example", []openapiclient.GslbResource{*openapiclient.NewGslbResource("Destination_example", "Region_example")}) // GslbCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GslbV1GslbsApiAPI.CreateGslb(context.Background()).GslbCreateRequest(gslbCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GslbV1GslbsApiAPI.CreateGslb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGslb`: GslbShowResponse
	fmt.Fprintf(os.Stdout, "Response from `GslbV1GslbsApiAPI.CreateGslb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGslbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gslbCreateRequest** | [**GslbCreateRequest**](GslbCreateRequest.md) |  | 

### Return type

[**GslbShowResponse**](GslbShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGslb

> GslbShowResponse DeleteGslb(ctx, gslbId).Execute()

DeleteGslb



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
	gslbId := "gslbId_example" // string | The GSLB ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GslbV1GslbsApiAPI.DeleteGslb(context.Background(), gslbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GslbV1GslbsApiAPI.DeleteGslb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGslb`: GslbShowResponse
	fmt.Fprintf(os.Stdout, "Response from `GslbV1GslbsApiAPI.DeleteGslb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gslbId** | **string** | The GSLB ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGslbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GslbShowResponse**](GslbShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGslbIpRanges

> GslbIpRangeResponse ListGslbIpRanges(ctx).Execute()

ListGslbIpRanges



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
	resp, r, err := apiClient.GslbV1GslbsApiAPI.ListGslbIpRanges(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GslbV1GslbsApiAPI.ListGslbIpRanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGslbIpRanges`: GslbIpRangeResponse
	fmt.Fprintf(os.Stdout, "Response from `GslbV1GslbsApiAPI.ListGslbIpRanges`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGslbIpRangesRequest struct via the builder pattern


### Return type

[**GslbIpRangeResponse**](GslbIpRangeResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGslbRegions

> GslbConfigRegionResponse ListGslbRegions(ctx).Execute()

ListGslbRegions



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
	resp, r, err := apiClient.GslbV1GslbsApiAPI.ListGslbRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GslbV1GslbsApiAPI.ListGslbRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGslbRegions`: GslbConfigRegionResponse
	fmt.Fprintf(os.Stdout, "Response from `GslbV1GslbsApiAPI.ListGslbRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGslbRegionsRequest struct via the builder pattern


### Return type

[**GslbConfigRegionResponse**](GslbConfigRegionResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGslbs

> GslbListResponse ListGslbs(ctx).Size(size).Page(page).Sort(sort).State(state).Name(name).Execute()

ListGslbs



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
	state := "state_example" // string | The GSLB State. (optional)
	name := "name_example" // string | The Name of the gslb. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GslbV1GslbsApiAPI.ListGslbs(context.Background()).Size(size).Page(page).Sort(sort).State(state).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GslbV1GslbsApiAPI.ListGslbs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGslbs`: GslbListResponse
	fmt.Fprintf(os.Stdout, "Response from `GslbV1GslbsApiAPI.ListGslbs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGslbsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **state** | **string** | The GSLB State. | 
 **name** | **string** | The Name of the gslb. | 

### Return type

[**GslbListResponse**](GslbListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGslb

> GslbShowResponse SetGslb(ctx, gslbId).GslbSetRequest(gslbSetRequest).Execute()

SetGslb



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
	gslbId := "gslbId_example" // string | The GSLB ID.
	gslbSetRequest := *openapiclient.NewGslbSetRequest() // GslbSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GslbV1GslbsApiAPI.SetGslb(context.Background(), gslbId).GslbSetRequest(gslbSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GslbV1GslbsApiAPI.SetGslb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGslb`: GslbShowResponse
	fmt.Fprintf(os.Stdout, "Response from `GslbV1GslbsApiAPI.SetGslb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gslbId** | **string** | The GSLB ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGslbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gslbSetRequest** | [**GslbSetRequest**](GslbSetRequest.md) |  | 

### Return type

[**GslbShowResponse**](GslbShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGslbHealthCheck

> GslbShowResponse SetGslbHealthCheck(ctx, gslbId).GslbHealthCheck(gslbHealthCheck).Execute()

SetGslbHealthCheck



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
	gslbId := "gslbId_example" // string | The GSLB ID.
	gslbHealthCheck := *openapiclient.NewGslbHealthCheck("Protocol_example") // GslbHealthCheck | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GslbV1GslbsApiAPI.SetGslbHealthCheck(context.Background(), gslbId).GslbHealthCheck(gslbHealthCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GslbV1GslbsApiAPI.SetGslbHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGslbHealthCheck`: GslbShowResponse
	fmt.Fprintf(os.Stdout, "Response from `GslbV1GslbsApiAPI.SetGslbHealthCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gslbId** | **string** | The GSLB ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGslbHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gslbHealthCheck** | [**GslbHealthCheck**](GslbHealthCheck.md) |  | 

### Return type

[**GslbShowResponse**](GslbShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowGslb

> GslbShowResponse ShowGslb(ctx, gslbId).Execute()

ShowGslb



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
	gslbId := "gslbId_example" // string | The GSLB ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GslbV1GslbsApiAPI.ShowGslb(context.Background(), gslbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GslbV1GslbsApiAPI.ShowGslb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowGslb`: GslbShowResponse
	fmt.Fprintf(os.Stdout, "Response from `GslbV1GslbsApiAPI.ShowGslb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gslbId** | **string** | The GSLB ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowGslbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GslbShowResponse**](GslbShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowGslbDefaultDomain

> GslbDefaultDomainResponse ShowGslbDefaultDomain(ctx).GslbEnv(gslbEnv).Execute()

ShowGslbDefaultDomain



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
	gslbEnv := "gslbEnv_example" // string | The GSLB Environment Usage. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GslbV1GslbsApiAPI.ShowGslbDefaultDomain(context.Background()).GslbEnv(gslbEnv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GslbV1GslbsApiAPI.ShowGslbDefaultDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowGslbDefaultDomain`: GslbDefaultDomainResponse
	fmt.Fprintf(os.Stdout, "Response from `GslbV1GslbsApiAPI.ShowGslbDefaultDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowGslbDefaultDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gslbEnv** | **string** | The GSLB Environment Usage. | 

### Return type

[**GslbDefaultDomainResponse**](GslbDefaultDomainResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

