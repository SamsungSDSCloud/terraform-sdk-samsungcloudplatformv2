# \VirtualserverV1ServersPasswordAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowVirtualServerPassword**](VirtualserverV1ServersPasswordAPI.md#ShowVirtualServerPassword) | **Post** /v1/servers/{server_id}/password | Show Virtual Server Password



## ShowVirtualServerPassword

> ServerPasswordResponse ShowVirtualServerPassword(ctx, serverId).ServerPasswordRequest(serverPasswordRequest).Execute()

Show Virtual Server Password



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
	serverId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Server ID
	serverPasswordRequest := *openapiclient.NewServerPasswordRequest() // ServerPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersPasswordAPI.ShowVirtualServerPassword(context.Background(), serverId).ServerPasswordRequest(serverPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersPasswordAPI.ShowVirtualServerPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVirtualServerPassword`: ServerPasswordResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersPasswordAPI.ShowVirtualServerPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVirtualServerPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverPasswordRequest** | [**ServerPasswordRequest**](ServerPasswordRequest.md) |  | 

### Return type

[**ServerPasswordResponse**](ServerPasswordResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

