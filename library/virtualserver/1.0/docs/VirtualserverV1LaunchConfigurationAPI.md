# \VirtualserverV1LaunchConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLaunchConfiguration**](VirtualserverV1LaunchConfigurationAPI.md#CreateLaunchConfiguration) | **Post** /v1/launch-configurations | Create Launch Configuration
[**DeleteLaunchConfiguration**](VirtualserverV1LaunchConfigurationAPI.md#DeleteLaunchConfiguration) | **Delete** /v1/launch-configurations/{launch_configuration_id} | Delete Launch Configuration
[**ListLaunchConfigurations**](VirtualserverV1LaunchConfigurationAPI.md#ListLaunchConfigurations) | **Get** /v1/launch-configurations | List Launch Configurations
[**ShowLaunchConfiguration**](VirtualserverV1LaunchConfigurationAPI.md#ShowLaunchConfiguration) | **Get** /v1/launch-configurations/{launch_configuration_id} | Show Launch Configuration



## CreateLaunchConfiguration

> LaunchConfigurationDetailShowResponse CreateLaunchConfiguration(ctx).LaunchConfigurationCreateRequest(launchConfigurationCreateRequest).Execute()

Create Launch Configuration



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
	launchConfigurationCreateRequest := *openapiclient.NewLaunchConfigurationCreateRequest("70a599e0-31e7-49b7-b260-868f441e862b", "keypairname", "launch-configuration-name", "s1v1m2") // LaunchConfigurationCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1LaunchConfigurationAPI.CreateLaunchConfiguration(context.Background()).LaunchConfigurationCreateRequest(launchConfigurationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1LaunchConfigurationAPI.CreateLaunchConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLaunchConfiguration`: LaunchConfigurationDetailShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1LaunchConfigurationAPI.CreateLaunchConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLaunchConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **launchConfigurationCreateRequest** | [**LaunchConfigurationCreateRequest**](LaunchConfigurationCreateRequest.md) |  | 

### Return type

[**LaunchConfigurationDetailShowResponse**](LaunchConfigurationDetailShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLaunchConfiguration

> DeleteLaunchConfiguration(ctx, launchConfigurationId).Execute()

Delete Launch Configuration



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
	launchConfigurationId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Launch Configuration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1LaunchConfigurationAPI.DeleteLaunchConfiguration(context.Background(), launchConfigurationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1LaunchConfigurationAPI.DeleteLaunchConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**launchConfigurationId** | **string** | Launch Configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLaunchConfigurationRequest struct via the builder pattern


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


## ListLaunchConfigurations

> LaunchConfigurationShowResponse ListLaunchConfigurations(ctx).Name(name).ImageId(imageId).Offset(offset).Limit(limit).Sort(sort).Execute()

List Launch Configurations



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
	name := "launch-configuration-name" // string | Launch Configuration name (optional)
	imageId := "70a599e0-31e7-49b7-b260-868f441e862b" // string | Image ID (optional)
	offset := int32(56) // int32 | Offset (optional)
	limit := int32(56) // int32 | Limit (optional)
	sort := "sort_example" // string | Sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1LaunchConfigurationAPI.ListLaunchConfigurations(context.Background()).Name(name).ImageId(imageId).Offset(offset).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1LaunchConfigurationAPI.ListLaunchConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLaunchConfigurations`: LaunchConfigurationShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1LaunchConfigurationAPI.ListLaunchConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLaunchConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Launch Configuration name | 
 **imageId** | **string** | Image ID | 
 **offset** | **int32** | Offset | 
 **limit** | **int32** | Limit | 
 **sort** | **string** | Sort | 

### Return type

[**LaunchConfigurationShowResponse**](LaunchConfigurationShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowLaunchConfiguration

> LaunchConfigurationDetailShowResponse ShowLaunchConfiguration(ctx, launchConfigurationId).Execute()

Show Launch Configuration



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
	launchConfigurationId := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Launch Configuration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1LaunchConfigurationAPI.ShowLaunchConfiguration(context.Background(), launchConfigurationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1LaunchConfigurationAPI.ShowLaunchConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowLaunchConfiguration`: LaunchConfigurationDetailShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1LaunchConfigurationAPI.ShowLaunchConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**launchConfigurationId** | **string** | Launch Configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowLaunchConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LaunchConfigurationDetailShowResponse**](LaunchConfigurationDetailShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

