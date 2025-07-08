# \BillingplanV1PlannedComputeApiAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlannedComputes**](BillingplanV1PlannedComputeApiAPI.md#CreatePlannedComputes) | **Post** /v1/planned-computes | Create Request Planned Computes
[**ListContractTypes**](BillingplanV1PlannedComputeApiAPI.md#ListContractTypes) | **Get** /v1/planned-computes/contract-types | List Contract Type
[**ListOsTypes**](BillingplanV1PlannedComputeApiAPI.md#ListOsTypes) | **Get** /v1/planned-computes/os-types | List OS System
[**ListPlannedComputeInstances**](BillingplanV1PlannedComputeApiAPI.md#ListPlannedComputeInstances) | **Get** /v1/planned-computes/instances | List Planned Compute Covered Resource
[**ListPlannedComputeServerTypes**](BillingplanV1PlannedComputeApiAPI.md#ListPlannedComputeServerTypes) | **Get** /v1/planned-computes/server-types | ListServerType
[**ListPlannedComputeServiceTypes**](BillingplanV1PlannedComputeApiAPI.md#ListPlannedComputeServiceTypes) | **Get** /v1/planned-computes/service-types | ListServiceType
[**ListPlannedComputes**](BillingplanV1PlannedComputeApiAPI.md#ListPlannedComputes) | **Get** /v1/planned-computes | ListPlannedComputes
[**ShowCancellationFee**](BillingplanV1PlannedComputeApiAPI.md#ShowCancellationFee) | **Post** /v1/planned-computes/cancellation-fee | GetTerminationInfo
[**ShowPlannedCompute**](BillingplanV1PlannedComputeApiAPI.md#ShowPlannedCompute) | **Get** /v1/planned-computes/{planned_compute_id} | Get Planned Compute
[**UpdatePlannedCompute**](BillingplanV1PlannedComputeApiAPI.md#UpdatePlannedCompute) | **Put** /v1/planned-computes/{planned_compute_id} | Change Request Planned Compute



## CreatePlannedComputes

> PlannedComputeResponse CreatePlannedComputes(ctx).PlannedComputeCreateRequest(plannedComputeCreateRequest).Execute()

Create Request Planned Computes



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
	plannedComputeCreateRequest := *openapiclient.NewPlannedComputeCreateRequest("228cb9e4a7934f84853594c7f26f7a21", "s1v1m2", "VIRTUAL_SERVER") // PlannedComputeCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingplanV1PlannedComputeApiAPI.CreatePlannedComputes(context.Background()).PlannedComputeCreateRequest(plannedComputeCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingplanV1PlannedComputeApiAPI.CreatePlannedComputes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlannedComputes`: PlannedComputeResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingplanV1PlannedComputeApiAPI.CreatePlannedComputes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlannedComputesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plannedComputeCreateRequest** | [**PlannedComputeCreateRequest**](PlannedComputeCreateRequest.md) |  | 

### Return type

[**PlannedComputeResponse**](PlannedComputeResponse.md)

### Authorization

[keystone-auth-token](../README.md#keystone-auth-token), [user-id](../README.md#user-id), [project-id](../README.md#project-id)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContractTypes

> ContractTypeListResponse ListContractTypes(ctx).ServiceId(serviceId).ServerType(serverType).Execute()

List Contract Type



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
	serviceId := "VIRTUAL_SERVER" // string | Service Id (optional)
	serverType := "s1v1m2" // string | Server Type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingplanV1PlannedComputeApiAPI.ListContractTypes(context.Background()).ServiceId(serviceId).ServerType(serverType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingplanV1PlannedComputeApiAPI.ListContractTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContractTypes`: ContractTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingplanV1PlannedComputeApiAPI.ListContractTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContractTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId** | **string** | Service Id | 
 **serverType** | **string** | Server Type | 

### Return type

[**ContractTypeListResponse**](ContractTypeListResponse.md)

### Authorization

[keystone-auth-token](../README.md#keystone-auth-token), [user-id](../README.md#user-id), [project-id](../README.md#project-id)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOsTypes

> OSTypeListResponse ListOsTypes(ctx).ServiceId(serviceId).Execute()

List OS System



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
	serviceId := "VIRTUAL_SERVER" // string | Service Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingplanV1PlannedComputeApiAPI.ListOsTypes(context.Background()).ServiceId(serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingplanV1PlannedComputeApiAPI.ListOsTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOsTypes`: OSTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingplanV1PlannedComputeApiAPI.ListOsTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOsTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId** | **string** | Service Id | 

### Return type

[**OSTypeListResponse**](OSTypeListResponse.md)

### Authorization

[keystone-auth-token](../README.md#keystone-auth-token), [user-id](../README.md#user-id), [project-id](../README.md#project-id)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlannedComputeInstances

> PlannedComputeResourceListResponse ListPlannedComputeInstances(ctx).StartDate(startDate).EndDate(endDate).ServiceId(serviceId).OsType(osType).ServerType(serverType).Sort(sort).Execute()

List Planned Compute Covered Resource



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
	startDate := "2024-08-01" // string | Start date (YYYY-MM-dd) (optional)
	endDate := "2024-08-01" // string | End date (YYYY-MM-dd) (optional)
	serviceId := "VIRTUAL_SERVER" // string | Service Id (optional)
	osType := "OPEN_SOURCE" // string | OS Type (optional)
	serverType := "s1v1m2" // string | Server Type (optional)
	sort := "asc" // string | Sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingplanV1PlannedComputeApiAPI.ListPlannedComputeInstances(context.Background()).StartDate(startDate).EndDate(endDate).ServiceId(serviceId).OsType(osType).ServerType(serverType).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingplanV1PlannedComputeApiAPI.ListPlannedComputeInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlannedComputeInstances`: PlannedComputeResourceListResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingplanV1PlannedComputeApiAPI.ListPlannedComputeInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlannedComputeInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Start date (YYYY-MM-dd) | 
 **endDate** | **string** | End date (YYYY-MM-dd) | 
 **serviceId** | **string** | Service Id | 
 **osType** | **string** | OS Type | 
 **serverType** | **string** | Server Type | 
 **sort** | **string** | Sort | 

### Return type

[**PlannedComputeResourceListResponse**](PlannedComputeResourceListResponse.md)

### Authorization

[keystone-auth-token](../README.md#keystone-auth-token), [user-id](../README.md#user-id), [project-id](../README.md#project-id)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlannedComputeServerTypes

> ServerTypeListResponse ListPlannedComputeServerTypes(ctx).ServiceId(serviceId).CurrentServerType(currentServerType).Execute()

ListServerType



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
	serviceId := "VIRTUAL_SERVER" // string | Service Id (optional)
	currentServerType := "s1v1m2" // string | Server Type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingplanV1PlannedComputeApiAPI.ListPlannedComputeServerTypes(context.Background()).ServiceId(serviceId).CurrentServerType(currentServerType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingplanV1PlannedComputeApiAPI.ListPlannedComputeServerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlannedComputeServerTypes`: ServerTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingplanV1PlannedComputeApiAPI.ListPlannedComputeServerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlannedComputeServerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId** | **string** | Service Id | 
 **currentServerType** | **string** | Server Type | 

### Return type

[**ServerTypeListResponse**](ServerTypeListResponse.md)

### Authorization

[keystone-auth-token](../README.md#keystone-auth-token), [user-id](../README.md#user-id), [project-id](../README.md#project-id)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlannedComputeServiceTypes

> ProductListResponse ListPlannedComputeServiceTypes(ctx).Execute()

ListServiceType



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
	resp, r, err := apiClient.BillingplanV1PlannedComputeApiAPI.ListPlannedComputeServiceTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingplanV1PlannedComputeApiAPI.ListPlannedComputeServiceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlannedComputeServiceTypes`: ProductListResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingplanV1PlannedComputeApiAPI.ListPlannedComputeServiceTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPlannedComputeServiceTypesRequest struct via the builder pattern


### Return type

[**ProductListResponse**](ProductListResponse.md)

### Authorization

[keystone-auth-token](../README.md#keystone-auth-token), [user-id](../README.md#user-id), [project-id](../README.md#project-id)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlannedComputes

> PlannedComputeListResponse ListPlannedComputes(ctx).Limit(limit).Page(page).StartDate(startDate).EndDate(endDate).ServerType(serverType).ContractId(contractId).ContractType(contractType).NextContractType(nextContractType).ServiceId(serviceId).OsType(osType).State(state).CreatedBy(createdBy).ModifiedBy(modifiedBy).Sort(sort).Execute()

ListPlannedComputes



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
	limit := int32(10) // int32 | LIMIT (optional)
	page := int32(1) // int32 | Page (optional)
	startDate := "2024-08-01" // string | Start date (YYYY-MM-dd) (optional)
	endDate := "2024-08-01" // string | End date (YYYY-MM-dd) (optional)
	serverType := "s1v1m2" // string | Server Type (optional)
	contractId := "C002612775" // string | Contract Id (optional)
	contractType := []*string{"Inner_example"} // []*string | Contract Type (optional)
	nextContractType := []*string{"Inner_example"} // []*string | Next Contract Type (optional)
	serviceId := []*string{"Inner_example"} // []*string | Service Id (optional)
	osType := []*string{"Inner_example"} // []*string | OS Type (optional)
	state := []*string{"Inner_example"} // []*string | Planned Compute State (optional)
	createdBy := "kim" // string | Created by (optional)
	modifiedBy := "kim" // string | Modified by (optional)
	sort := "asc" // string | Sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingplanV1PlannedComputeApiAPI.ListPlannedComputes(context.Background()).Limit(limit).Page(page).StartDate(startDate).EndDate(endDate).ServerType(serverType).ContractId(contractId).ContractType(contractType).NextContractType(nextContractType).ServiceId(serviceId).OsType(osType).State(state).CreatedBy(createdBy).ModifiedBy(modifiedBy).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingplanV1PlannedComputeApiAPI.ListPlannedComputes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlannedComputes`: PlannedComputeListResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingplanV1PlannedComputeApiAPI.ListPlannedComputes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlannedComputesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | LIMIT | 
 **page** | **int32** | Page | 
 **startDate** | **string** | Start date (YYYY-MM-dd) | 
 **endDate** | **string** | End date (YYYY-MM-dd) | 
 **serverType** | **string** | Server Type | 
 **contractId** | **string** | Contract Id | 
 **contractType** | **[]string** | Contract Type | 
 **nextContractType** | **[]string** | Next Contract Type | 
 **serviceId** | **[]string** | Service Id | 
 **osType** | **[]string** | OS Type | 
 **state** | **[]string** | Planned Compute State | 
 **createdBy** | **string** | Created by | 
 **modifiedBy** | **string** | Modified by | 
 **sort** | **string** | Sort | 

### Return type

[**PlannedComputeListResponse**](PlannedComputeListResponse.md)

### Authorization

[keystone-auth-token](../README.md#keystone-auth-token), [user-id](../README.md#user-id), [project-id](../README.md#project-id)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowCancellationFee

> CancellationFeeResponse ShowCancellationFee(ctx).CancellationFeeRequest(cancellationFeeRequest).Execute()

GetTerminationInfo



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
	cancellationFeeRequest := *openapiclient.NewCancellationFeeRequest() // CancellationFeeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingplanV1PlannedComputeApiAPI.ShowCancellationFee(context.Background()).CancellationFeeRequest(cancellationFeeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingplanV1PlannedComputeApiAPI.ShowCancellationFee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowCancellationFee`: CancellationFeeResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingplanV1PlannedComputeApiAPI.ShowCancellationFee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowCancellationFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancellationFeeRequest** | [**CancellationFeeRequest**](CancellationFeeRequest.md) |  | 

### Return type

[**CancellationFeeResponse**](CancellationFeeResponse.md)

### Authorization

[keystone-auth-token](../README.md#keystone-auth-token), [user-id](../README.md#user-id), [project-id](../README.md#project-id)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPlannedCompute

> PlannedComputeResponse ShowPlannedCompute(ctx, plannedComputeId).Execute()

Get Planned Compute



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
	plannedComputeId := "5f3688a192094ddeb0c94166a3e88935" // string | Planned Compute Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingplanV1PlannedComputeApiAPI.ShowPlannedCompute(context.Background(), plannedComputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingplanV1PlannedComputeApiAPI.ShowPlannedCompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowPlannedCompute`: PlannedComputeResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingplanV1PlannedComputeApiAPI.ShowPlannedCompute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannedComputeId** | **string** | Planned Compute Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowPlannedComputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlannedComputeResponse**](PlannedComputeResponse.md)

### Authorization

[keystone-auth-token](../README.md#keystone-auth-token), [user-id](../README.md#user-id), [project-id](../README.md#project-id)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlannedCompute

> PlannedComputeResponse UpdatePlannedCompute(ctx, plannedComputeId).PlannedComputeChangeRequest(plannedComputeChangeRequest).Execute()

Change Request Planned Compute



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
	plannedComputeId := "5f3688a192094ddeb0c94166a3e88935" // string | Planned Compute Id
	plannedComputeChangeRequest := *openapiclient.NewPlannedComputeChangeRequest() // PlannedComputeChangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingplanV1PlannedComputeApiAPI.UpdatePlannedCompute(context.Background(), plannedComputeId).PlannedComputeChangeRequest(plannedComputeChangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingplanV1PlannedComputeApiAPI.UpdatePlannedCompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePlannedCompute`: PlannedComputeResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingplanV1PlannedComputeApiAPI.UpdatePlannedCompute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannedComputeId** | **string** | Planned Compute Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlannedComputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **plannedComputeChangeRequest** | [**PlannedComputeChangeRequest**](PlannedComputeChangeRequest.md) |  | 

### Return type

[**PlannedComputeResponse**](PlannedComputeResponse.md)

### Authorization

[keystone-auth-token](../README.md#keystone-auth-token), [user-id](../README.md#user-id), [project-id](../README.md#project-id)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

