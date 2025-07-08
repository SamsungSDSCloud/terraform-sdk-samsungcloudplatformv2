# \VpcV1VpcPeeringRuleApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVpcPeeringRule**](VpcV1VpcPeeringRuleApiAPI.md#CreateVpcPeeringRule) | **Post** /v1/vpc-peerings/{vpc_peering_id}/routing-rules | Create Vpc Peering Rule
[**DeleteVpcPeeringRule**](VpcV1VpcPeeringRuleApiAPI.md#DeleteVpcPeeringRule) | **Delete** /v1/vpc-peerings/{vpc_peering_id}/routing-rules/{routing_rule_id} | Delete VPC Peering Rule
[**ListVpcPeeringRoutingRuleAvailabilities**](VpcV1VpcPeeringRuleApiAPI.md#ListVpcPeeringRoutingRuleAvailabilities) | **Get** /v1/vpc-peerings/{vpc_peering_id}/routing-rules/destination-cidr-availabilities | List VPC Peering Routing Rule Availabilities
[**ListVpcPeeringRules**](VpcV1VpcPeeringRuleApiAPI.md#ListVpcPeeringRules) | **Get** /v1/vpc-peerings/{vpc_peering_id}/routing-rules | List VPC Peering Rules



## CreateVpcPeeringRule

> VpcPeeringRuleShowResponse CreateVpcPeeringRule(ctx, vpcPeeringId).VpcPeeringRuleCreateRequest(vpcPeeringRuleCreateRequest).Execute()

Create Vpc Peering Rule



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
	vpcPeeringId := "vpcPeeringId_example" // string | VPC Peering ID
	vpcPeeringRuleCreateRequest := *openapiclient.NewVpcPeeringRuleCreateRequest("DestinationCidr_example", openapiclient.VpcPeeringRuleDestinationVpcType("REQUESTER_VPC")) // VpcPeeringRuleCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcPeeringRuleApiAPI.CreateVpcPeeringRule(context.Background(), vpcPeeringId).VpcPeeringRuleCreateRequest(vpcPeeringRuleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcPeeringRuleApiAPI.CreateVpcPeeringRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpcPeeringRule`: VpcPeeringRuleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcPeeringRuleApiAPI.CreateVpcPeeringRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcPeeringId** | **string** | VPC Peering ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpcPeeringRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpcPeeringRuleCreateRequest** | [**VpcPeeringRuleCreateRequest**](VpcPeeringRuleCreateRequest.md) |  | 

### Return type

[**VpcPeeringRuleShowResponse**](VpcPeeringRuleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpcPeeringRule

> DeleteVpcPeeringRule(ctx, vpcPeeringId, routingRuleId).Execute()

Delete VPC Peering Rule



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
	vpcPeeringId := "vpcPeeringId_example" // string | VPC Peering ID
	routingRuleId := "routingRuleId_example" // string | Routing Rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1VpcPeeringRuleApiAPI.DeleteVpcPeeringRule(context.Background(), vpcPeeringId, routingRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcPeeringRuleApiAPI.DeleteVpcPeeringRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcPeeringId** | **string** | VPC Peering ID | 
**routingRuleId** | **string** | Routing Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpcPeeringRuleRequest struct via the builder pattern


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


## ListVpcPeeringRoutingRuleAvailabilities

> VpcPeeringRuleAvailabilityResponse ListVpcPeeringRoutingRuleAvailabilities(ctx, vpcPeeringId).DestinationVpcType(destinationVpcType).DestinationCidr(destinationCidr).Execute()

List VPC Peering Routing Rule Availabilities



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
	vpcPeeringId := "vpcPeeringId_example" // string | VPC Peering ID
	destinationVpcType := openapiclient.VpcPeeringRuleDestinationVpcType("REQUESTER_VPC") // VpcPeeringRuleDestinationVpcType | Destination VPC Type
	destinationCidr := "destinationCidr_example" // string | Destination CIDR

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcPeeringRuleApiAPI.ListVpcPeeringRoutingRuleAvailabilities(context.Background(), vpcPeeringId).DestinationVpcType(destinationVpcType).DestinationCidr(destinationCidr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcPeeringRuleApiAPI.ListVpcPeeringRoutingRuleAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpcPeeringRoutingRuleAvailabilities`: VpcPeeringRuleAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcPeeringRuleApiAPI.ListVpcPeeringRoutingRuleAvailabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcPeeringId** | **string** | VPC Peering ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVpcPeeringRoutingRuleAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationVpcType** | [**VpcPeeringRuleDestinationVpcType**](VpcPeeringRuleDestinationVpcType.md) | Destination VPC Type | 
 **destinationCidr** | **string** | Destination CIDR | 

### Return type

[**VpcPeeringRuleAvailabilityResponse**](VpcPeeringRuleAvailabilityResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVpcPeeringRules

> VpcPeeringRuleListResponse ListVpcPeeringRules(ctx, vpcPeeringId).Size(size).Page(page).Sort(sort).Id(id).Name(name).SourceVpcId(sourceVpcId).SourceVpcType(sourceVpcType).DestinationVpcId(destinationVpcId).DestinationVpcType(destinationVpcType).DestinationCidr(destinationCidr).State(state).Execute()

List VPC Peering Rules



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
	vpcPeeringId := "vpcPeeringId_example" // string | VPC Peering ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	id := "id_example" // string | VPC Peering Rule ID (optional)
	name := "name_example" // string | VPC Peering Name (optional)
	sourceVpcId := "sourceVpcId_example" // string | Source VPC ID (optional)
	sourceVpcType := openapiclient.VpcPeeringRuleDestinationVpcType("REQUESTER_VPC") // VpcPeeringRuleDestinationVpcType | Source VPC Type (optional)
	destinationVpcId := "destinationVpcId_example" // string | Destination VPC ID (optional)
	destinationVpcType := openapiclient.VpcPeeringRuleDestinationVpcType("REQUESTER_VPC") // VpcPeeringRuleDestinationVpcType | Destination VPC Type (optional)
	destinationCidr := "destinationCidr_example" // string | Destination CIDR (optional)
	state := openapiclient.VpcPeeringState("CREATING") // VpcPeeringState | State (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcPeeringRuleApiAPI.ListVpcPeeringRules(context.Background(), vpcPeeringId).Size(size).Page(page).Sort(sort).Id(id).Name(name).SourceVpcId(sourceVpcId).SourceVpcType(sourceVpcType).DestinationVpcId(destinationVpcId).DestinationVpcType(destinationVpcType).DestinationCidr(destinationCidr).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcPeeringRuleApiAPI.ListVpcPeeringRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpcPeeringRules`: VpcPeeringRuleListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcPeeringRuleApiAPI.ListVpcPeeringRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcPeeringId** | **string** | VPC Peering ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVpcPeeringRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **id** | **string** | VPC Peering Rule ID | 
 **name** | **string** | VPC Peering Name | 
 **sourceVpcId** | **string** | Source VPC ID | 
 **sourceVpcType** | [**VpcPeeringRuleDestinationVpcType**](VpcPeeringRuleDestinationVpcType.md) | Source VPC Type | 
 **destinationVpcId** | **string** | Destination VPC ID | 
 **destinationVpcType** | [**VpcPeeringRuleDestinationVpcType**](VpcPeeringRuleDestinationVpcType.md) | Destination VPC Type | 
 **destinationCidr** | **string** | Destination CIDR | 
 **state** | [**VpcPeeringState**](VpcPeeringState.md) | State | 

### Return type

[**VpcPeeringRuleListResponse**](VpcPeeringRuleListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

