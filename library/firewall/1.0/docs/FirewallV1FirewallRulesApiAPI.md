# \FirewallV1FirewallRulesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFirewallRule**](FirewallV1FirewallRulesApiAPI.md#CreateFirewallRule) | **Post** /v1/firewalls/rules | Create Firewall Rule
[**CreateFirewallRuleBulk**](FirewallV1FirewallRulesApiAPI.md#CreateFirewallRuleBulk) | **Post** /v1/firewalls/rules/bulk | Create Firewall Rule Bulk
[**DeleteFirewallRule**](FirewallV1FirewallRulesApiAPI.md#DeleteFirewallRule) | **Delete** /v1/firewalls/rules/{firewall_rule_id} | Delete Firewall Rule
[**DeleteFirewallRuleBulk**](FirewallV1FirewallRulesApiAPI.md#DeleteFirewallRuleBulk) | **Delete** /v1/firewalls/rules/bulk | Delete Firewall Rule Bulk
[**DownloadFirewallRules**](FirewallV1FirewallRulesApiAPI.md#DownloadFirewallRules) | **Get** /v1/firewalls/rules/downloads | Download Firewall Rules
[**ListFirewallRules**](FirewallV1FirewallRulesApiAPI.md#ListFirewallRules) | **Get** /v1/firewalls/rules | List Firewall Rules
[**SetFirewallRule**](FirewallV1FirewallRulesApiAPI.md#SetFirewallRule) | **Put** /v1/firewalls/rules/{firewall_rule_id} | Set Firewall Rule
[**SetFirewallRuleBulk**](FirewallV1FirewallRulesApiAPI.md#SetFirewallRuleBulk) | **Put** /v1/firewalls/rules/bulk | Set Firewall Rule Bulk
[**SetFirewallRuleOrder**](FirewallV1FirewallRulesApiAPI.md#SetFirewallRuleOrder) | **Put** /v1/firewalls/rules/{firewall_rule_id}/order | Set Firewall Rule Order
[**ShowFirewallRule**](FirewallV1FirewallRulesApiAPI.md#ShowFirewallRule) | **Get** /v1/firewalls/rules/{firewall_rule_id} | Show Firewall Rule



## CreateFirewallRule

> FirewallRuleShowResponse CreateFirewallRule(ctx).FirewallRuleCreateSingleRequest(firewallRuleCreateSingleRequest).Execute()

Create Firewall Rule



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
	firewallRuleCreateSingleRequest := *openapiclient.NewFirewallRuleCreateSingleRequest("b156740b6335468d8354eb9ef8eddf5a", *openapiclient.NewFirewallRuleCreateRequest(openapiclient.FirewallRuleAction("ALLOW"), []string{"DestinationAddress_example"}, openapiclient.FirewallRuleDirection("INBOUND"), []openapiclient.FirewallPort{*openapiclient.NewFirewallPort(openapiclient.FirewallServiceType("TCP"))}, []string{"SourceAddress_example"}, openapiclient.FirewallStatusType("ENABLE"))) // FirewallRuleCreateSingleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallV1FirewallRulesApiAPI.CreateFirewallRule(context.Background()).FirewallRuleCreateSingleRequest(firewallRuleCreateSingleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallRulesApiAPI.CreateFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirewallRule`: FirewallRuleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallV1FirewallRulesApiAPI.CreateFirewallRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallRuleCreateSingleRequest** | [**FirewallRuleCreateSingleRequest**](FirewallRuleCreateSingleRequest.md) |  | 

### Return type

[**FirewallRuleShowResponse**](FirewallRuleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFirewallRuleBulk

> CreateFirewallRuleBulk(ctx).FirewallRuleCreateBulkRequest(firewallRuleCreateBulkRequest).Execute()

Create Firewall Rule Bulk



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
	firewallRuleCreateBulkRequest := *openapiclient.NewFirewallRuleCreateBulkRequest("b156740b6335468d8354eb9ef8eddf5a", []openapiclient.FirewallRuleCreateRequest{*openapiclient.NewFirewallRuleCreateRequest(openapiclient.FirewallRuleAction("ALLOW"), []string{"DestinationAddress_example"}, openapiclient.FirewallRuleDirection("INBOUND"), []openapiclient.FirewallPort{*openapiclient.NewFirewallPort(openapiclient.FirewallServiceType("TCP"))}, []string{"SourceAddress_example"}, openapiclient.FirewallStatusType("ENABLE"))}) // FirewallRuleCreateBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FirewallV1FirewallRulesApiAPI.CreateFirewallRuleBulk(context.Background()).FirewallRuleCreateBulkRequest(firewallRuleCreateBulkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallRulesApiAPI.CreateFirewallRuleBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallRuleBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallRuleCreateBulkRequest** | [**FirewallRuleCreateBulkRequest**](FirewallRuleCreateBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewallRule

> DeleteFirewallRule(ctx, firewallRuleId).Execute()

Delete Firewall Rule



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
	firewallRuleId := "b156740b6335468d8354eb9ef8eddf5a" // string | Firewall Rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FirewallV1FirewallRulesApiAPI.DeleteFirewallRule(context.Background(), firewallRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallRulesApiAPI.DeleteFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallRuleId** | **string** | Firewall Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRuleRequest struct via the builder pattern


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


## DeleteFirewallRuleBulk

> DeleteFirewallRuleBulk(ctx).FirewallRuleDeleteBulkRequest(firewallRuleDeleteBulkRequest).Execute()

Delete Firewall Rule Bulk



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
	firewallRuleDeleteBulkRequest := *openapiclient.NewFirewallRuleDeleteBulkRequest([]string{"FirewallRuleId_example"}) // FirewallRuleDeleteBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FirewallV1FirewallRulesApiAPI.DeleteFirewallRuleBulk(context.Background()).FirewallRuleDeleteBulkRequest(firewallRuleDeleteBulkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallRulesApiAPI.DeleteFirewallRuleBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRuleBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallRuleDeleteBulkRequest** | [**FirewallRuleDeleteBulkRequest**](FirewallRuleDeleteBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFirewallRules

> DownloadFirewallRules(ctx).FirewallId(firewallId).Execute()

Download Firewall Rules



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
	firewallId := "b156740b6335468d8354eb9ef8eddf5a" // string | Firewall ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FirewallV1FirewallRulesApiAPI.DownloadFirewallRules(context.Background()).FirewallId(firewallId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallRulesApiAPI.DownloadFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallId** | **string** | Firewall ID | 

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


## ListFirewallRules

> FirewallRuleListResponse ListFirewallRules(ctx).FirewallId(firewallId).Size(size).Page(page).Sort(sort).SrcIp(srcIp).DstIp(dstIp).Description(description).State(state).Status(status).FetchAll(fetchAll).Execute()

List Firewall Rules



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
	firewallId := "b156740b6335468d8354eb9ef8eddf5a" // string | Firewall ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	srcIp := "192.168.1.1" // string | Source IP (optional)
	dstIp := "192.168.1.1" // string | Destination IP (optional)
	description := "Firewall rule description example" // string | Firewall Rule Description (optional)
	state := []openapiclient.FirewallRuleState{openapiclient.FirewallRuleState("CREATING")} // []FirewallRuleState | Firewall Rule State (optional) (default to [])
	status := openapiclient.FirewallStatusType("ENABLE") // FirewallStatusType | Firewall Rule Status (optional)
	fetchAll := true // bool | Firewall Rule Fetch ALL (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallV1FirewallRulesApiAPI.ListFirewallRules(context.Background()).FirewallId(firewallId).Size(size).Page(page).Sort(sort).SrcIp(srcIp).DstIp(dstIp).Description(description).State(state).Status(status).FetchAll(fetchAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallRulesApiAPI.ListFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFirewallRules`: FirewallRuleListResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallV1FirewallRulesApiAPI.ListFirewallRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallId** | **string** | Firewall ID | 
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **srcIp** | **string** | Source IP | 
 **dstIp** | **string** | Destination IP | 
 **description** | **string** | Firewall Rule Description | 
 **state** | [**[]FirewallRuleState**](FirewallRuleState.md) | Firewall Rule State | [default to []]
 **status** | [**FirewallStatusType**](FirewallStatusType.md) | Firewall Rule Status | 
 **fetchAll** | **bool** | Firewall Rule Fetch ALL | [default to false]

### Return type

[**FirewallRuleListResponse**](FirewallRuleListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFirewallRule

> FirewallRuleShowResponse SetFirewallRule(ctx, firewallRuleId).FirewallRuleSetRequest(firewallRuleSetRequest).Execute()

Set Firewall Rule



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
	firewallRuleId := "b156740b6335468d8354eb9ef8eddf5a" // string | Firewall Rule ID
	firewallRuleSetRequest := *openapiclient.NewFirewallRuleSetRequest(openapiclient.FirewallRuleAction("ALLOW"), []string{"DestinationAddress_example"}, openapiclient.FirewallRuleDirection("INBOUND"), []openapiclient.FirewallPort{*openapiclient.NewFirewallPort(openapiclient.FirewallServiceType("TCP"))}, []string{"SourceAddress_example"}) // FirewallRuleSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallV1FirewallRulesApiAPI.SetFirewallRule(context.Background(), firewallRuleId).FirewallRuleSetRequest(firewallRuleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallRulesApiAPI.SetFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFirewallRule`: FirewallRuleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallV1FirewallRulesApiAPI.SetFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallRuleId** | **string** | Firewall Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallRuleSetRequest** | [**FirewallRuleSetRequest**](FirewallRuleSetRequest.md) |  | 

### Return type

[**FirewallRuleShowResponse**](FirewallRuleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFirewallRuleBulk

> SetFirewallRuleBulk(ctx).FirewallRuleUpdateBulkRequest(firewallRuleUpdateBulkRequest).Execute()

Set Firewall Rule Bulk



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
	firewallRuleUpdateBulkRequest := *openapiclient.NewFirewallRuleUpdateBulkRequest([]string{"FirewallRuleId_example"}, openapiclient.FirewallStatusType("ENABLE")) // FirewallRuleUpdateBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FirewallV1FirewallRulesApiAPI.SetFirewallRuleBulk(context.Background()).FirewallRuleUpdateBulkRequest(firewallRuleUpdateBulkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallRulesApiAPI.SetFirewallRuleBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetFirewallRuleBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallRuleUpdateBulkRequest** | [**FirewallRuleUpdateBulkRequest**](FirewallRuleUpdateBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFirewallRuleOrder

> FirewallRuleShowResponse SetFirewallRuleOrder(ctx, firewallRuleId).FirewallRuleSetOrderRequest(firewallRuleSetOrderRequest).Execute()

Set Firewall Rule Order



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
	firewallRuleId := "b156740b6335468d8354eb9ef8eddf5a" // string | Firewall Rule ID
	firewallRuleSetOrderRequest := *openapiclient.NewFirewallRuleSetOrderRequest(openapiclient.FirewallRuleOrderDirection("BEFORE")) // FirewallRuleSetOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallV1FirewallRulesApiAPI.SetFirewallRuleOrder(context.Background(), firewallRuleId).FirewallRuleSetOrderRequest(firewallRuleSetOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallRulesApiAPI.SetFirewallRuleOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFirewallRuleOrder`: FirewallRuleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallV1FirewallRulesApiAPI.SetFirewallRuleOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallRuleId** | **string** | Firewall Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFirewallRuleOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallRuleSetOrderRequest** | [**FirewallRuleSetOrderRequest**](FirewallRuleSetOrderRequest.md) |  | 

### Return type

[**FirewallRuleShowResponse**](FirewallRuleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFirewallRule

> FirewallRuleShowResponse ShowFirewallRule(ctx, firewallRuleId).Execute()

Show Firewall Rule



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
	firewallRuleId := "b156740b6335468d8354eb9ef8eddf5a" // string | Firewall Rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallV1FirewallRulesApiAPI.ShowFirewallRule(context.Background(), firewallRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallRulesApiAPI.ShowFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowFirewallRule`: FirewallRuleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallV1FirewallRulesApiAPI.ShowFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallRuleId** | **string** | Firewall Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirewallRuleShowResponse**](FirewallRuleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

