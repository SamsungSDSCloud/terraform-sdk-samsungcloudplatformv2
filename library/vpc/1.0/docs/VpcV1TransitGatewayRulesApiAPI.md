# \VpcV1TransitGatewayRulesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransitGatewayRule**](VpcV1TransitGatewayRulesApiAPI.md#CreateTransitGatewayRule) | **Post** /v1/transit-gateways/{transit_gateway_id}/routing-rules | Create Transit Gateway Rule
[**DeleteTransitGatewayRule**](VpcV1TransitGatewayRulesApiAPI.md#DeleteTransitGatewayRule) | **Delete** /v1/transit-gateways/{transit_gateway_id}/routing-rules/{routing_rule_id} | Delete Transit Gateway Rule
[**ListTransitGatewayRoutingRuleAvailabilities**](VpcV1TransitGatewayRulesApiAPI.md#ListTransitGatewayRoutingRuleAvailabilities) | **Get** /v1/transit-gateways/{transit_gateway_id}/routing-rules/destination-cidr-availabilities | List Transit Gateway Routing Rule Availabilities
[**ListTransitGatewayRules**](VpcV1TransitGatewayRulesApiAPI.md#ListTransitGatewayRules) | **Get** /v1/transit-gateways/{transit_gateway_id}/routing-rules | List Transit Gateway Rules



## CreateTransitGatewayRule

> TransitGatewayRuleShowResponse CreateTransitGatewayRule(ctx, transitGatewayId).TransitGatewayRuleCreateRequest(transitGatewayRuleCreateRequest).Execute()

Create Transit Gateway Rule



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
	transitGatewayId := "transitGatewayId_example" // string | Transit Gateway ID
	transitGatewayRuleCreateRequest := *openapiclient.NewTransitGatewayRuleCreateRequest("DestinationCidr_example", openapiclient.TransitGatewayRuleDestinationType("VPC"), "TgwConnectionVpcId_example") // TransitGatewayRuleCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1TransitGatewayRulesApiAPI.CreateTransitGatewayRule(context.Background(), transitGatewayId).TransitGatewayRuleCreateRequest(transitGatewayRuleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1TransitGatewayRulesApiAPI.CreateTransitGatewayRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTransitGatewayRule`: TransitGatewayRuleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1TransitGatewayRulesApiAPI.CreateTransitGatewayRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitGatewayId** | **string** | Transit Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransitGatewayRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitGatewayRuleCreateRequest** | [**TransitGatewayRuleCreateRequest**](TransitGatewayRuleCreateRequest.md) |  | 

### Return type

[**TransitGatewayRuleShowResponse**](TransitGatewayRuleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransitGatewayRule

> DeleteTransitGatewayRule(ctx, transitGatewayId, routingRuleId).Execute()

Delete Transit Gateway Rule



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
	transitGatewayId := "transitGatewayId_example" // string | Transit Gateway ID
	routingRuleId := "routingRuleId_example" // string | Routing Rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1TransitGatewayRulesApiAPI.DeleteTransitGatewayRule(context.Background(), transitGatewayId, routingRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1TransitGatewayRulesApiAPI.DeleteTransitGatewayRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitGatewayId** | **string** | Transit Gateway ID | 
**routingRuleId** | **string** | Routing Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransitGatewayRuleRequest struct via the builder pattern


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


## ListTransitGatewayRoutingRuleAvailabilities

> TransitGatewayRuleAvailabilityResponse ListTransitGatewayRoutingRuleAvailabilities(ctx, transitGatewayId).ConnectedVpcId(connectedVpcId).DestinationType(destinationType).DestinationCidr(destinationCidr).Execute()

List Transit Gateway Routing Rule Availabilities



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
	transitGatewayId := "transitGatewayId_example" // string | Transit Gateway ID
	connectedVpcId := "connectedVpcId_example" // string | VPC ID Connected to Transit Gateway.
	destinationType := openapiclient.TransitGatewayRuleDestinationType("VPC") // TransitGatewayRuleDestinationType | Destination Type
	destinationCidr := "destinationCidr_example" // string | Destination CIDR

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1TransitGatewayRulesApiAPI.ListTransitGatewayRoutingRuleAvailabilities(context.Background(), transitGatewayId).ConnectedVpcId(connectedVpcId).DestinationType(destinationType).DestinationCidr(destinationCidr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1TransitGatewayRulesApiAPI.ListTransitGatewayRoutingRuleAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransitGatewayRoutingRuleAvailabilities`: TransitGatewayRuleAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1TransitGatewayRulesApiAPI.ListTransitGatewayRoutingRuleAvailabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitGatewayId** | **string** | Transit Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransitGatewayRoutingRuleAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectedVpcId** | **string** | VPC ID Connected to Transit Gateway. | 
 **destinationType** | [**TransitGatewayRuleDestinationType**](TransitGatewayRuleDestinationType.md) | Destination Type | 
 **destinationCidr** | **string** | Destination CIDR | 

### Return type

[**TransitGatewayRuleAvailabilityResponse**](TransitGatewayRuleAvailabilityResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransitGatewayRules

> TransitGatewayRuleListResponse ListTransitGatewayRules(ctx, transitGatewayId).Size(size).Page(page).Sort(sort).Id(id).TgwConnectionVpcId(tgwConnectionVpcId).TgwConnectionVpcName(tgwConnectionVpcName).SourceType(sourceType).DestinationType(destinationType).DestinationCidr(destinationCidr).State(state).Execute()

List Transit Gateway Rules



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
	transitGatewayId := "transitGatewayId_example" // string | Transit Gateway ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	id := "id_example" // string | Routing Rule ID (optional)
	tgwConnectionVpcId := "tgwConnectionVpcId_example" // string | VPC ID Connected to Transit Gateway. (optional)
	tgwConnectionVpcName := "tgwConnectionVpcName_example" // string | VPC Name Connected to Transit Gateway. (optional)
	sourceType := openapiclient.TransitGatewayRuleDestinationType("VPC") // TransitGatewayRuleDestinationType | Source Type (optional)
	destinationType := openapiclient.TransitGatewayRuleDestinationType("VPC") // TransitGatewayRuleDestinationType | Destination Type (optional)
	destinationCidr := "destinationCidr_example" // string | Destination CIDR (optional)
	state := openapiclient.RoutingRuleState("CREATING") // RoutingRuleState | State (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1TransitGatewayRulesApiAPI.ListTransitGatewayRules(context.Background(), transitGatewayId).Size(size).Page(page).Sort(sort).Id(id).TgwConnectionVpcId(tgwConnectionVpcId).TgwConnectionVpcName(tgwConnectionVpcName).SourceType(sourceType).DestinationType(destinationType).DestinationCidr(destinationCidr).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1TransitGatewayRulesApiAPI.ListTransitGatewayRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransitGatewayRules`: TransitGatewayRuleListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1TransitGatewayRulesApiAPI.ListTransitGatewayRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitGatewayId** | **string** | Transit Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransitGatewayRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **id** | **string** | Routing Rule ID | 
 **tgwConnectionVpcId** | **string** | VPC ID Connected to Transit Gateway. | 
 **tgwConnectionVpcName** | **string** | VPC Name Connected to Transit Gateway. | 
 **sourceType** | [**TransitGatewayRuleDestinationType**](TransitGatewayRuleDestinationType.md) | Source Type | 
 **destinationType** | [**TransitGatewayRuleDestinationType**](TransitGatewayRuleDestinationType.md) | Destination Type | 
 **destinationCidr** | **string** | Destination CIDR | 
 **state** | [**RoutingRuleState**](RoutingRuleState.md) | State | 

### Return type

[**TransitGatewayRuleListResponse**](TransitGatewayRuleListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

