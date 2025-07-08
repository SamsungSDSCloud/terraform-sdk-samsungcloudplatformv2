# \VpcV1VpcPeeringApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVpcPeering**](VpcV1VpcPeeringApiAPI.md#CreateVpcPeering) | **Post** /v1/vpc-peerings | Create VPC Peering
[**DeleteVpcPeering**](VpcV1VpcPeeringApiAPI.md#DeleteVpcPeering) | **Delete** /v1/vpc-peerings/{vpc_peering_id} | Delete VPC Peering
[**ListVpcPeeringAvailabilities**](VpcV1VpcPeeringApiAPI.md#ListVpcPeeringAvailabilities) | **Get** /v1/vpc-peerings/peering-availabilities | List VPC Peering Availabilities
[**ListVpcPeerings**](VpcV1VpcPeeringApiAPI.md#ListVpcPeerings) | **Get** /v1/vpc-peerings | List VPC Peering
[**SetVpcPeering**](VpcV1VpcPeeringApiAPI.md#SetVpcPeering) | **Put** /v1/vpc-peerings/{vpc_peering_id} | Set VPC Peering
[**ShowVpcPeering**](VpcV1VpcPeeringApiAPI.md#ShowVpcPeering) | **Get** /v1/vpc-peerings/{vpc_peering_id} | Show VPC Peering



## CreateVpcPeering

> VpcPeeringShowResponse CreateVpcPeering(ctx).VpcPeeringCreateRequest(vpcPeeringCreateRequest).Execute()

Create VPC Peering



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
	vpcPeeringCreateRequest := *openapiclient.NewVpcPeeringCreateRequest("ApproverVpcId_example", "Name_example", "RequesterVpcId_example") // VpcPeeringCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcPeeringApiAPI.CreateVpcPeering(context.Background()).VpcPeeringCreateRequest(vpcPeeringCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcPeeringApiAPI.CreateVpcPeering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpcPeering`: VpcPeeringShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcPeeringApiAPI.CreateVpcPeering`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpcPeeringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpcPeeringCreateRequest** | [**VpcPeeringCreateRequest**](VpcPeeringCreateRequest.md) |  | 

### Return type

[**VpcPeeringShowResponse**](VpcPeeringShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpcPeering

> DeleteVpcPeering(ctx, vpcPeeringId).Execute()

Delete VPC Peering



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1VpcPeeringApiAPI.DeleteVpcPeering(context.Background(), vpcPeeringId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcPeeringApiAPI.DeleteVpcPeering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcPeeringId** | **string** | VPC Peering ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpcPeeringRequest struct via the builder pattern


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


## ListVpcPeeringAvailabilities

> VpcPeeringAvailabilityResponse ListVpcPeeringAvailabilities(ctx).RequesterVpcId(requesterVpcId).ApproverVpcId(approverVpcId).ApproverVpcAccountId(approverVpcAccountId).Execute()

List VPC Peering Availabilities



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
	requesterVpcId := "requesterVpcId_example" // string | Requester VPC ID
	approverVpcId := "approverVpcId_example" // string | Approver VPC ID (optional)
	approverVpcAccountId := "approverVpcAccountId_example" // string | Approver VPC Account ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcPeeringApiAPI.ListVpcPeeringAvailabilities(context.Background()).RequesterVpcId(requesterVpcId).ApproverVpcId(approverVpcId).ApproverVpcAccountId(approverVpcAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcPeeringApiAPI.ListVpcPeeringAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpcPeeringAvailabilities`: VpcPeeringAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcPeeringApiAPI.ListVpcPeeringAvailabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVpcPeeringAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requesterVpcId** | **string** | Requester VPC ID | 
 **approverVpcId** | **string** | Approver VPC ID | 
 **approverVpcAccountId** | **string** | Approver VPC Account ID | 

### Return type

[**VpcPeeringAvailabilityResponse**](VpcPeeringAvailabilityResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVpcPeerings

> VpcPeeringListResponse ListVpcPeerings(ctx).Size(size).Page(page).Sort(sort).Id(id).Name(name).RequesterVpcId(requesterVpcId).RequesterVpcName(requesterVpcName).ApproverVpcId(approverVpcId).ApproverVpcName(approverVpcName).AccountType(accountType).State(state).Execute()

List VPC Peering



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
	id := "id_example" // string | VPC Peering ID (optional)
	name := "name_example" // string | VPC Peering Name (optional)
	requesterVpcId := "requesterVpcId_example" // string | Requester VPC ID (optional)
	requesterVpcName := "requesterVpcName_example" // string | Requester VPC Name (optional)
	approverVpcId := "approverVpcId_example" // string | Approver VPC ID (optional)
	approverVpcName := "approverVpcName_example" // string | Approver VPC Name (optional)
	accountType := openapiclient.VpcPeeringAccountType("SAME") // VpcPeeringAccountType | Account Type (optional)
	state := openapiclient.VpcPeeringState("CREATING") // VpcPeeringState | State (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcPeeringApiAPI.ListVpcPeerings(context.Background()).Size(size).Page(page).Sort(sort).Id(id).Name(name).RequesterVpcId(requesterVpcId).RequesterVpcName(requesterVpcName).ApproverVpcId(approverVpcId).ApproverVpcName(approverVpcName).AccountType(accountType).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcPeeringApiAPI.ListVpcPeerings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpcPeerings`: VpcPeeringListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcPeeringApiAPI.ListVpcPeerings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVpcPeeringsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **id** | **string** | VPC Peering ID | 
 **name** | **string** | VPC Peering Name | 
 **requesterVpcId** | **string** | Requester VPC ID | 
 **requesterVpcName** | **string** | Requester VPC Name | 
 **approverVpcId** | **string** | Approver VPC ID | 
 **approverVpcName** | **string** | Approver VPC Name | 
 **accountType** | [**VpcPeeringAccountType**](VpcPeeringAccountType.md) | Account Type | 
 **state** | [**VpcPeeringState**](VpcPeeringState.md) | State | 

### Return type

[**VpcPeeringListResponse**](VpcPeeringListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVpcPeering

> VpcPeeringShowResponse SetVpcPeering(ctx, vpcPeeringId).VpcPeeringSetRequest(vpcPeeringSetRequest).Execute()

Set VPC Peering



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
	vpcPeeringSetRequest := *openapiclient.NewVpcPeeringSetRequest() // VpcPeeringSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcPeeringApiAPI.SetVpcPeering(context.Background(), vpcPeeringId).VpcPeeringSetRequest(vpcPeeringSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcPeeringApiAPI.SetVpcPeering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVpcPeering`: VpcPeeringShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcPeeringApiAPI.SetVpcPeering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcPeeringId** | **string** | VPC Peering ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVpcPeeringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpcPeeringSetRequest** | [**VpcPeeringSetRequest**](VpcPeeringSetRequest.md) |  | 

### Return type

[**VpcPeeringShowResponse**](VpcPeeringShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVpcPeering

> VpcPeeringShowResponse ShowVpcPeering(ctx, vpcPeeringId).Execute()

Show VPC Peering



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcPeeringApiAPI.ShowVpcPeering(context.Background(), vpcPeeringId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcPeeringApiAPI.ShowVpcPeering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVpcPeering`: VpcPeeringShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcPeeringApiAPI.ShowVpcPeering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcPeeringId** | **string** | VPC Peering ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVpcPeeringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VpcPeeringShowResponse**](VpcPeeringShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

