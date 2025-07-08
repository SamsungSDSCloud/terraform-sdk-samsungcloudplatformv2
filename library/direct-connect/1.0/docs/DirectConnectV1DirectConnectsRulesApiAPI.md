# \DirectConnectV1DirectConnectsRulesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckDestinationCidrAvailabilities**](DirectConnectV1DirectConnectsRulesApiAPI.md#CheckDestinationCidrAvailabilities) | **Get** /v1/direct-connects/{direct_connect_id}/routing-rules/destination-cidr-availabilities | CheckDestinationCidrAvailabilities
[**CreateRoutingRule**](DirectConnectV1DirectConnectsRulesApiAPI.md#CreateRoutingRule) | **Post** /v1/direct-connects/{direct_connect_id}/routing-rules | Create Routing Rule
[**DeleteRoutingRule**](DirectConnectV1DirectConnectsRulesApiAPI.md#DeleteRoutingRule) | **Delete** /v1/direct-connects/{direct_connect_id}/routing-rules/{routing_rule_id} | Delete Routing Rule
[**ListRoutingRules**](DirectConnectV1DirectConnectsRulesApiAPI.md#ListRoutingRules) | **Get** /v1/direct-connects/{direct_connect_id}/routing-rules | List Routing Rules



## CheckDestinationCidrAvailabilities

> RoutingRuleCidrAvailabilityResponse CheckDestinationCidrAvailabilities(ctx, directConnectId).DestinationType(destinationType).DestinationCidr(destinationCidr).Execute()

CheckDestinationCidrAvailabilities



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
	destinationType := openapiclient.DirectConnectRoutingRuleDestinationType("VPC") // DirectConnectRoutingRuleDestinationType | Destination Type
	destinationCidr := "destinationCidr_example" // string | Destination CIDR

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectConnectV1DirectConnectsRulesApiAPI.CheckDestinationCidrAvailabilities(context.Background(), directConnectId).DestinationType(destinationType).DestinationCidr(destinationCidr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectConnectV1DirectConnectsRulesApiAPI.CheckDestinationCidrAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckDestinationCidrAvailabilities`: RoutingRuleCidrAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectConnectV1DirectConnectsRulesApiAPI.CheckDestinationCidrAvailabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directConnectId** | **string** | Direct Connect ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckDestinationCidrAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationType** | [**DirectConnectRoutingRuleDestinationType**](DirectConnectRoutingRuleDestinationType.md) | Destination Type | 
 **destinationCidr** | **string** | Destination CIDR | 

### Return type

[**RoutingRuleCidrAvailabilityResponse**](RoutingRuleCidrAvailabilityResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoutingRule

> RoutingRuleShowResponse CreateRoutingRule(ctx, directConnectId).RoutingRuleCreateRequest(routingRuleCreateRequest).Execute()

Create Routing Rule



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
	routingRuleCreateRequest := *openapiclient.NewRoutingRuleCreateRequest("DestinationCidr_example", openapiclient.DirectConnectRoutingRuleDestinationType("VPC")) // RoutingRuleCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectConnectV1DirectConnectsRulesApiAPI.CreateRoutingRule(context.Background(), directConnectId).RoutingRuleCreateRequest(routingRuleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectConnectV1DirectConnectsRulesApiAPI.CreateRoutingRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoutingRule`: RoutingRuleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectConnectV1DirectConnectsRulesApiAPI.CreateRoutingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directConnectId** | **string** | Direct Connect ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoutingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routingRuleCreateRequest** | [**RoutingRuleCreateRequest**](RoutingRuleCreateRequest.md) |  | 

### Return type

[**RoutingRuleShowResponse**](RoutingRuleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoutingRule

> DeleteRoutingRule(ctx, directConnectId, routingRuleId).Execute()

Delete Routing Rule



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
	routingRuleId := "routingRuleId_example" // string | Routing Rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DirectConnectV1DirectConnectsRulesApiAPI.DeleteRoutingRule(context.Background(), directConnectId, routingRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectConnectV1DirectConnectsRulesApiAPI.DeleteRoutingRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directConnectId** | **string** | Direct Connect ID | 
**routingRuleId** | **string** | Routing Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoutingRuleRequest struct via the builder pattern


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


## ListRoutingRules

> RoutingRuleListResponse ListRoutingRules(ctx, directConnectId).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).DestinationType(destinationType).DestinationCidr(destinationCidr).State(state).Execute()

List Routing Rules



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
	withCount := "true" // string | with count (optional)
	limit := int32(20) // int32 | limit (optional)
	marker := "607e0938521643b5b4b266f343fae693" // string | marker (optional)
	sort := "created_at:desc" // string | sort (optional)
	id := "id_example" // string | Routing Rule ID (optional)
	destinationType := openapiclient.DirectConnectRoutingRuleDestinationType("VPC") // DirectConnectRoutingRuleDestinationType | Destination Type (optional)
	destinationCidr := "destinationCidr_example" // string | Destination CIDR (optional)
	state := openapiclient.RoutingRuleState("CREATING") // RoutingRuleState | State (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectConnectV1DirectConnectsRulesApiAPI.ListRoutingRules(context.Background(), directConnectId).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).DestinationType(destinationType).DestinationCidr(destinationCidr).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectConnectV1DirectConnectsRulesApiAPI.ListRoutingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoutingRules`: RoutingRuleListResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectConnectV1DirectConnectsRulesApiAPI.ListRoutingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directConnectId** | **string** | Direct Connect ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoutingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 
 **id** | **string** | Routing Rule ID | 
 **destinationType** | [**DirectConnectRoutingRuleDestinationType**](DirectConnectRoutingRuleDestinationType.md) | Destination Type | 
 **destinationCidr** | **string** | Destination CIDR | 
 **state** | [**RoutingRuleState**](RoutingRuleState.md) | State | 

### Return type

[**RoutingRuleListResponse**](RoutingRuleListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

