# \DnsV1HostedZonesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostedZone**](DnsV1HostedZonesApiAPI.md#CreateHostedZone) | **Post** /v1/hosted-zones | Create Hosted Zone
[**DeleteHostedZone**](DnsV1HostedZonesApiAPI.md#DeleteHostedZone) | **Delete** /v1/hosted-zones/{hosted_zone_id} | Delete Hosted Zone
[**ListHostedZone**](DnsV1HostedZonesApiAPI.md#ListHostedZone) | **Get** /v1/hosted-zones | List Hosted Zones
[**SetHostedZone**](DnsV1HostedZonesApiAPI.md#SetHostedZone) | **Put** /v1/hosted-zones/{hosted_zone_id} | Update Hosted Zone
[**ShowHostedZone**](DnsV1HostedZonesApiAPI.md#ShowHostedZone) | **Get** /v1/hosted-zones/{hosted_zone_id} | Show Hosted Zone



## CreateHostedZone

> HostedZoneCreateResponse CreateHostedZone(ctx).HostedZoneCreateRequest(hostedZoneCreateRequest).Execute()

Create Hosted Zone



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
	hostedZoneCreateRequest := *openapiclient.NewHostedZoneCreateRequest("Email_example", "Name_example") // HostedZoneCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1HostedZonesApiAPI.CreateHostedZone(context.Background()).HostedZoneCreateRequest(hostedZoneCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1HostedZonesApiAPI.CreateHostedZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHostedZone`: HostedZoneCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1HostedZonesApiAPI.CreateHostedZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostedZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostedZoneCreateRequest** | [**HostedZoneCreateRequest**](HostedZoneCreateRequest.md) |  | 

### Return type

[**HostedZoneCreateResponse**](HostedZoneCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHostedZone

> HostedZoneDeleteResponse DeleteHostedZone(ctx, hostedZoneId).Execute()

Delete Hosted Zone



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
	hostedZoneId := "hostedZoneId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1HostedZonesApiAPI.DeleteHostedZone(context.Background(), hostedZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1HostedZonesApiAPI.DeleteHostedZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteHostedZone`: HostedZoneDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1HostedZonesApiAPI.DeleteHostedZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostedZoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostedZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostedZoneDeleteResponse**](HostedZoneDeleteResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHostedZone

> HostedZoneListResponse ListHostedZone(ctx).Limit(limit).Marker(marker).SortDir(sortDir).SortKey(sortKey).Name(name).ExactName(exactName).Type_(type_).Email(email).Status(status).Description(description).Ttl(ttl).Execute()

List Hosted Zones



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
	limit := int32(56) // int32 | limit (optional)
	marker := "marker_example" // string | marker (optional)
	sortDir := "sortDir_example" // string | sort direction (optional)
	sortKey := "sortKey_example" // string | sort key (optional)
	name := "name_example" // string | name (optional)
	exactName := "exactName_example" // string | name (optional)
	type_ := "type__example" // string | Type of zone (optional)
	email := "email_example" // string | email (optional)
	status := "status_example" // string | The status (optional)
	description := "description_example" // string | description (optional)
	ttl := int32(56) // int32 | TTL for the zone. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1HostedZonesApiAPI.ListHostedZone(context.Background()).Limit(limit).Marker(marker).SortDir(sortDir).SortKey(sortKey).Name(name).ExactName(exactName).Type_(type_).Email(email).Status(status).Description(description).Ttl(ttl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1HostedZonesApiAPI.ListHostedZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHostedZone`: HostedZoneListResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1HostedZonesApiAPI.ListHostedZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostedZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sortDir** | **string** | sort direction | 
 **sortKey** | **string** | sort key | 
 **name** | **string** | name | 
 **exactName** | **string** | name | 
 **type_** | **string** | Type of zone | 
 **email** | **string** | email | 
 **status** | **string** | The status | 
 **description** | **string** | description | 
 **ttl** | **int32** | TTL for the zone. | 

### Return type

[**HostedZoneListResponse**](HostedZoneListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetHostedZone

> HostedZoneSetResponse SetHostedZone(ctx, hostedZoneId).HostedZoneSetRequest(hostedZoneSetRequest).Execute()

Update Hosted Zone



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
	hostedZoneId := "hostedZoneId_example" // string | ID
	hostedZoneSetRequest := *openapiclient.NewHostedZoneSetRequest() // HostedZoneSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1HostedZonesApiAPI.SetHostedZone(context.Background(), hostedZoneId).HostedZoneSetRequest(hostedZoneSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1HostedZonesApiAPI.SetHostedZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetHostedZone`: HostedZoneSetResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1HostedZonesApiAPI.SetHostedZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostedZoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetHostedZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostedZoneSetRequest** | [**HostedZoneSetRequest**](HostedZoneSetRequest.md) |  | 

### Return type

[**HostedZoneSetResponse**](HostedZoneSetResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowHostedZone

> HostedZoneShowResponse ShowHostedZone(ctx, hostedZoneId).Execute()

Show Hosted Zone



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
	hostedZoneId := "hostedZoneId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1HostedZonesApiAPI.ShowHostedZone(context.Background(), hostedZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1HostedZonesApiAPI.ShowHostedZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowHostedZone`: HostedZoneShowResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1HostedZonesApiAPI.ShowHostedZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostedZoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowHostedZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostedZoneShowResponse**](HostedZoneShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

