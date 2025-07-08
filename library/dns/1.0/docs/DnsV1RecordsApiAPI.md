# \DnsV1RecordsApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRecord**](DnsV1RecordsApiAPI.md#CreateRecord) | **Post** /v1/hosted-zones/{hosted_zone_id}/records | Create Record
[**DeleteRecord**](DnsV1RecordsApiAPI.md#DeleteRecord) | **Delete** /v1/hosted-zones/{hosted_zone_id}/records/{record_id} | Delete Record
[**ListRecords**](DnsV1RecordsApiAPI.md#ListRecords) | **Get** /v1/hosted-zones/{hosted_zone_id}/records | List Records
[**SetRecord**](DnsV1RecordsApiAPI.md#SetRecord) | **Put** /v1/hosted-zones/{hosted_zone_id}/records/{record_id} | Update Record
[**ShowRecord**](DnsV1RecordsApiAPI.md#ShowRecord) | **Get** /v1/hosted-zones/{hosted_zone_id}/records/{record_id} | Show Record



## CreateRecord

> RecordCreateResponse CreateRecord(ctx, hostedZoneId).RecordCreateRequest(recordCreateRequest).Execute()

Create Record



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
	recordCreateRequest := *openapiclient.NewRecordCreateRequest("Name_example", []interface{}{nil}, "Type_example") // RecordCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1RecordsApiAPI.CreateRecord(context.Background(), hostedZoneId).RecordCreateRequest(recordCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1RecordsApiAPI.CreateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRecord`: RecordCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1RecordsApiAPI.CreateRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostedZoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordCreateRequest** | [**RecordCreateRequest**](RecordCreateRequest.md) |  | 

### Return type

[**RecordCreateResponse**](RecordCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecord

> RecordCreateResponse DeleteRecord(ctx, hostedZoneId, recordId).Execute()

Delete Record



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
	hostedZoneId := "hostedZoneId_example" // string | Zone ID
	recordId := "recordId_example" // string | Record ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1RecordsApiAPI.DeleteRecord(context.Background(), hostedZoneId, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1RecordsApiAPI.DeleteRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRecord`: RecordCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1RecordsApiAPI.DeleteRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostedZoneId** | **string** | Zone ID | 
**recordId** | **string** | Record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RecordCreateResponse**](RecordCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecords

> RecordListResponse ListRecords(ctx, hostedZoneId).Limit(limit).Marker(marker).SortDir(sortDir).SortKey(sortKey).Name(name).ExactName(exactName).Type_(type_).Data(data).Status(status).Description(description).Ttl(ttl).Execute()

List Records



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
	limit := int32(56) // int32 | limit (optional)
	marker := "marker_example" // string | marker (optional)
	sortDir := "sortDir_example" // string | sort direction (optional)
	sortKey := "sortKey_example" // string | sort key (optional)
	name := "name_example" // string | name (optional)
	exactName := "exactName_example" // string | name (optional)
	type_ := "type__example" // string | Type of the record (optional)
	data := "data_example" // string | Record data (optional)
	status := "status_example" // string | The status (optional)
	description := "description_example" // string | description (optional)
	ttl := int32(56) // int32 | TTL for the zone. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1RecordsApiAPI.ListRecords(context.Background(), hostedZoneId).Limit(limit).Marker(marker).SortDir(sortDir).SortKey(sortKey).Name(name).ExactName(exactName).Type_(type_).Data(data).Status(status).Description(description).Ttl(ttl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1RecordsApiAPI.ListRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRecords`: RecordListResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1RecordsApiAPI.ListRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostedZoneId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sortDir** | **string** | sort direction | 
 **sortKey** | **string** | sort key | 
 **name** | **string** | name | 
 **exactName** | **string** | name | 
 **type_** | **string** | Type of the record | 
 **data** | **string** | Record data | 
 **status** | **string** | The status | 
 **description** | **string** | description | 
 **ttl** | **int32** | TTL for the zone. | 

### Return type

[**RecordListResponse**](RecordListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRecord

> RecordSetResponse SetRecord(ctx, hostedZoneId, recordId).RecordSetRequest(recordSetRequest).Execute()

Update Record



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
	hostedZoneId := "hostedZoneId_example" // string | Zone ID
	recordId := "recordId_example" // string | Record ID
	recordSetRequest := *openapiclient.NewRecordSetRequest() // RecordSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1RecordsApiAPI.SetRecord(context.Background(), hostedZoneId, recordId).RecordSetRequest(recordSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1RecordsApiAPI.SetRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRecord`: RecordSetResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1RecordsApiAPI.SetRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostedZoneId** | **string** | Zone ID | 
**recordId** | **string** | Record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recordSetRequest** | [**RecordSetRequest**](RecordSetRequest.md) |  | 

### Return type

[**RecordSetResponse**](RecordSetResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowRecord

> RecordShowResponse ShowRecord(ctx, hostedZoneId, recordId).Execute()

Show Record



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
	hostedZoneId := "hostedZoneId_example" // string | Zone ID
	recordId := "recordId_example" // string | Record ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1RecordsApiAPI.ShowRecord(context.Background(), hostedZoneId, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1RecordsApiAPI.ShowRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowRecord`: RecordShowResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1RecordsApiAPI.ShowRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostedZoneId** | **string** | Zone ID | 
**recordId** | **string** | Record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RecordShowResponse**](RecordShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

