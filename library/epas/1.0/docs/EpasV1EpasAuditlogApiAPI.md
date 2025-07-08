# \EpasV1EpasAuditlogApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EpasSetAuditLog**](EpasV1EpasAuditlogApiAPI.md#EpasSetAuditLog) | **Put** /v1/clusters/{cluster_id}/audit-log | Set Audit Log



## EpasSetAuditLog

> AsyncResponse EpasSetAuditLog(ctx, clusterId).AuditLogRequest(auditLogRequest).Execute()

Set Audit Log



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
	clusterId := "clusterId_example" // string | Cluster ID
	auditLogRequest := *openapiclient.NewAuditLogRequest(false) // AuditLogRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasAuditlogApiAPI.EpasSetAuditLog(context.Background(), clusterId).AuditLogRequest(auditLogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasAuditlogApiAPI.EpasSetAuditLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSetAuditLog`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasAuditlogApiAPI.EpasSetAuditLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSetAuditLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auditLogRequest** | [**AuditLogRequest**](AuditLogRequest.md) |  | 

### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

