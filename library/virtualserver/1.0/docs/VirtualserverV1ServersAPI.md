# \VirtualserverV1ServersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerInterface**](VirtualserverV1ServersAPI.md#CreateServerInterface) | **Post** /v1/servers/{server_id}/interfaces | Create Server Interface
[**CreateServerInterfaceNat**](VirtualserverV1ServersAPI.md#CreateServerInterfaceNat) | **Post** /v1/servers/{server_id}/interfaces/{port_id}/static-nats | Create Server Interface Nat
[**CreateServerInterfacePrivateNat**](VirtualserverV1ServersAPI.md#CreateServerInterfacePrivateNat) | **Post** /v1/servers/{server_id}/interfaces/{port_id}/private-static-nats | Create Server Interface Private Nat
[**CreateServerVolume**](VirtualserverV1ServersAPI.md#CreateServerVolume) | **Post** /v1/servers/{server_id}/volumes | Create Server Volume
[**CreateVirtualServer**](VirtualserverV1ServersAPI.md#CreateVirtualServer) | **Post** /v1/servers | Create Virtual Server
[**DeleteServerInterface**](VirtualserverV1ServersAPI.md#DeleteServerInterface) | **Delete** /v1/servers/{server_id}/interfaces/{port_id} | Delete Server Interface
[**DeleteServerInterfaceNat**](VirtualserverV1ServersAPI.md#DeleteServerInterfaceNat) | **Delete** /v1/servers/{server_id}/interfaces/{port_id}/static-nats/{nat_id} | Delete Server Interface Nat
[**DeleteServerInterfacePrivateNat**](VirtualserverV1ServersAPI.md#DeleteServerInterfacePrivateNat) | **Delete** /v1/servers/{server_id}/interfaces/{port_id}/private-static-nats/{nat_id} | Delete Server Interface Private Nat
[**DeleteServerVolume**](VirtualserverV1ServersAPI.md#DeleteServerVolume) | **Delete** /v1/servers/{server_id}/volumes/{volume_id} | Delete Server Volume
[**DeleteVirtualServer**](VirtualserverV1ServersAPI.md#DeleteVirtualServer) | **Delete** /v1/servers/{server_id} | Delete Virtual Server
[**ListServerInterfaces**](VirtualserverV1ServersAPI.md#ListServerInterfaces) | **Get** /v1/servers/{server_id}/interfaces | List Server Interfaces
[**ListServerIps**](VirtualserverV1ServersAPI.md#ListServerIps) | **Get** /v1/servers/{server_id}/ips | List Server Ips
[**ListServerVolumes**](VirtualserverV1ServersAPI.md#ListServerVolumes) | **Get** /v1/servers/{server_id}/volumes | List Server Volumes
[**ListVirtualServers**](VirtualserverV1ServersAPI.md#ListVirtualServers) | **Get** /v1/servers | List Virtual Servers
[**ShowServerInterface**](VirtualserverV1ServersAPI.md#ShowServerInterface) | **Get** /v1/servers/{server_id}/interfaces/{port_id} | Show Server Interface
[**ShowServerIp**](VirtualserverV1ServersAPI.md#ShowServerIp) | **Get** /v1/servers/{server_id}/ips/{subnet_id} | Show Server Ip
[**ShowServerVolume**](VirtualserverV1ServersAPI.md#ShowServerVolume) | **Get** /v1/servers/{server_id}/volumes/{volume_id} | Show Server Volume
[**ShowVirtualServer**](VirtualserverV1ServersAPI.md#ShowVirtualServer) | **Get** /v1/servers/{server_id} | Show Virtual Server
[**UpdateServerInterface**](VirtualserverV1ServersAPI.md#UpdateServerInterface) | **Put** /v1/servers/{server_id}/interfaces/{port_id} | Update Server Interface
[**UpdateServerVolume**](VirtualserverV1ServersAPI.md#UpdateServerVolume) | **Put** /v1/servers/{server_id}/volumes/{volume_id} | Update Server Volume
[**UpdateVirtualServer**](VirtualserverV1ServersAPI.md#UpdateVirtualServer) | **Put** /v1/servers/{server_id} | Update Virtual Server



## CreateServerInterface

> InterfaceResponse CreateServerInterface(ctx, serverId).ServerInterfaceCreateRequest(serverInterfaceCreateRequest).Execute()

Create Server Interface



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
	serverInterfaceCreateRequest := *openapiclient.NewServerInterfaceCreateRequest() // ServerInterfaceCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersAPI.CreateServerInterface(context.Background(), serverId).ServerInterfaceCreateRequest(serverInterfaceCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.CreateServerInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerInterface`: InterfaceResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersAPI.CreateServerInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverInterfaceCreateRequest** | [**ServerInterfaceCreateRequest**](ServerInterfaceCreateRequest.md) |  | 

### Return type

[**InterfaceResponse**](InterfaceResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServerInterfaceNat

> CreateServerInterfaceNat(ctx, serverId, portId).ServerStaticNatCreateRequest(serverStaticNatCreateRequest).Execute()

Create Server Interface Nat



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
	portId := "f2b02fc5-a5fd-4b5f-b1a9-987f4c0d66cc" // string | Port ID
	serverStaticNatCreateRequest := *openapiclient.NewServerStaticNatCreateRequest("1f0cb5390c40483592ddc5a282f53496") // ServerStaticNatCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServersAPI.CreateServerInterfaceNat(context.Background(), serverId, portId).ServerStaticNatCreateRequest(serverStaticNatCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.CreateServerInterfaceNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerInterfaceNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serverStaticNatCreateRequest** | [**ServerStaticNatCreateRequest**](ServerStaticNatCreateRequest.md) |  | 

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


## CreateServerInterfacePrivateNat

> CreateServerInterfacePrivateNat(ctx, serverId, portId).ServerPrivateStaticNatCreateRequest(serverPrivateStaticNatCreateRequest).Execute()

Create Server Interface Private Nat



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
	portId := "f2b02fc5-a5fd-4b5f-b1a9-987f4c0d66cc" // string | Port ID
	serverPrivateStaticNatCreateRequest := *openapiclient.NewServerPrivateStaticNatCreateRequest("3a6c1dc6b1b24e02a444d672c2798a1c", "1f0cb5390c40483592ddc5a282f53496") // ServerPrivateStaticNatCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServersAPI.CreateServerInterfacePrivateNat(context.Background(), serverId, portId).ServerPrivateStaticNatCreateRequest(serverPrivateStaticNatCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.CreateServerInterfacePrivateNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerInterfacePrivateNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serverPrivateStaticNatCreateRequest** | [**ServerPrivateStaticNatCreateRequest**](ServerPrivateStaticNatCreateRequest.md) |  | 

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


## CreateServerVolume

> ServersVolumeResponse CreateServerVolume(ctx, serverId).ServerVolumesCreateRequest(serverVolumesCreateRequest).Execute()

Create Server Volume



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
	serverVolumesCreateRequest := *openapiclient.NewServerVolumesCreateRequest("VolumeId_example") // ServerVolumesCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersAPI.CreateServerVolume(context.Background(), serverId).ServerVolumesCreateRequest(serverVolumesCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.CreateServerVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServerVolume`: ServersVolumeResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersAPI.CreateServerVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverVolumesCreateRequest** | [**ServerVolumesCreateRequest**](ServerVolumesCreateRequest.md) |  | 

### Return type

[**ServersVolumeResponse**](ServersVolumeResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVirtualServer

> ServerCreateResponse CreateVirtualServer(ctx).ServerCreateRequest(serverCreateRequest).Execute()

Create Virtual Server



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
	serverCreateRequest := *openapiclient.NewServerCreateRequest("70a599e0-31e7-49b7-b260-868f441e862b", "keypairname", "servername", []openapiclient.Network{*openapiclient.NewNetwork()}, "s1v1m2") // ServerCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersAPI.CreateVirtualServer(context.Background()).ServerCreateRequest(serverCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.CreateVirtualServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVirtualServer`: ServerCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersAPI.CreateVirtualServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverCreateRequest** | [**ServerCreateRequest**](ServerCreateRequest.md) |  | 

### Return type

[**ServerCreateResponse**](ServerCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerInterface

> DeleteServerInterface(ctx, serverId, portId).Execute()

Delete Server Interface



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
	portId := "f2b02fc5-a5fd-4b5f-b1a9-987f4c0d66cc" // string | Port ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServersAPI.DeleteServerInterface(context.Background(), serverId, portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.DeleteServerInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerInterfaceNat

> DeleteServerInterfaceNat(ctx, serverId, portId, natId).Execute()

Delete Server Interface Nat



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
	portId := "f2b02fc5-a5fd-4b5f-b1a9-987f4c0d66cc" // string | Port ID
	natId := "224b80106e6f41b38efe98ac9ddbf280" // string | NAT ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServersAPI.DeleteServerInterfaceNat(context.Background(), serverId, portId, natId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.DeleteServerInterfaceNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 
**portId** | **string** | Port ID | 
**natId** | **string** | NAT ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerInterfaceNatRequest struct via the builder pattern


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


## DeleteServerInterfacePrivateNat

> DeleteServerInterfacePrivateNat(ctx, serverId, portId, natId).Execute()

Delete Server Interface Private Nat



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
	portId := "f2b02fc5-a5fd-4b5f-b1a9-987f4c0d66cc" // string | Port ID
	natId := "224b80106e6f41b38efe98ac9ddbf280" // string | NAT ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServersAPI.DeleteServerInterfacePrivateNat(context.Background(), serverId, portId, natId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.DeleteServerInterfacePrivateNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 
**portId** | **string** | Port ID | 
**natId** | **string** | NAT ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerInterfacePrivateNatRequest struct via the builder pattern


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


## DeleteServerVolume

> DeleteServerVolume(ctx, serverId, volumeId).Execute()

Delete Server Volume



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
	volumeId := "3f500f00-ed80-4566-a057-e31760226f9a" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServersAPI.DeleteServerVolume(context.Background(), serverId, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.DeleteServerVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualServer

> DeleteVirtualServer(ctx, serverId).Execute()

Delete Virtual Server



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServersAPI.DeleteVirtualServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.DeleteVirtualServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServerInterfaces

> ServerInterfaceListResponse ListServerInterfaces(ctx, serverId).Execute()

List Server Interfaces



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersAPI.ListServerInterfaces(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.ListServerInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServerInterfaces`: ServerInterfaceListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersAPI.ListServerInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServerInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerInterfaceListResponse**](ServerInterfaceListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServerIps

> ServerIpsResponse ListServerIps(ctx, serverId).Execute()

List Server Ips



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersAPI.ListServerIps(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.ListServerIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServerIps`: ServerIpsResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersAPI.ListServerIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServerIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerIpsResponse**](ServerIpsResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServerVolumes

> ServerVolumesResponse ListServerVolumes(ctx, serverId).Execute()

List Server Volumes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersAPI.ListServerVolumes(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.ListServerVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServerVolumes`: ServerVolumesResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersAPI.ListServerVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServerVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerVolumesResponse**](ServerVolumesResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualServers

> ServerListResponse ListVirtualServers(ctx).Name(name).Ip(ip).State(state).ProductCategory(productCategory).ProductOffering(productOffering).VpcId(vpcId).ServerTypeId(serverTypeId).Limit(limit).Marker(marker).Sort(sort).AutoScalingGroupId(autoScalingGroupId).Execute()

List Virtual Servers



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
	name := "servername" // string | Server name (optional)
	ip := "192.169.3.2" // string | IP address (optional)
	state := "ACTIVE" // string | Server state (optional)
	productCategory := openapiclient.ServerProductCategory("compute") // ServerProductCategory | Product category (optional)
	productOffering := *openapiclient.NewProductOffering() // ProductOffering | Product offering (optional)
	vpcId := "cc976b621087484ea5fd527f4b78708b" // string | VPC ID (optional)
	serverTypeId := "s1v1m2" // string | Server type ID (optional)
	limit := int32(20) // int32 | Limit (optional)
	marker := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Marker (optional)
	sort := "name:asc" // string | Sort (optional)
	autoScalingGroupId := "52613bd852b04b39adcb15a8364d856d" // string | Auto-Scaling Group ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersAPI.ListVirtualServers(context.Background()).Name(name).Ip(ip).State(state).ProductCategory(productCategory).ProductOffering(productOffering).VpcId(vpcId).ServerTypeId(serverTypeId).Limit(limit).Marker(marker).Sort(sort).AutoScalingGroupId(autoScalingGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.ListVirtualServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualServers`: ServerListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersAPI.ListVirtualServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Server name | 
 **ip** | **string** | IP address | 
 **state** | **string** | Server state | 
 **productCategory** | [**ServerProductCategory**](ServerProductCategory.md) | Product category | 
 **productOffering** | [**ProductOffering**](ProductOffering.md) | Product offering | 
 **vpcId** | **string** | VPC ID | 
 **serverTypeId** | **string** | Server type ID | 
 **limit** | **int32** | Limit | 
 **marker** | **string** | Marker | 
 **sort** | **string** | Sort | 
 **autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Return type

[**ServerListResponse**](ServerListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowServerInterface

> ServerInterfaceListResponse ShowServerInterface(ctx, serverId, portId).Execute()

Show Server Interface



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
	portId := "f2b02fc5-a5fd-4b5f-b1a9-987f4c0d66cc" // string | Port ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersAPI.ShowServerInterface(context.Background(), serverId, portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.ShowServerInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowServerInterface`: ServerInterfaceListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersAPI.ShowServerInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowServerInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServerInterfaceListResponse**](ServerInterfaceListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowServerIp

> ServerShowResponseAddress ShowServerIp(ctx, serverId, subnetId).Execute()

Show Server Ip



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
	subnetId := "183c6c30df0c4fcbb16422aa2d64aa21" // string | Subnet ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersAPI.ShowServerIp(context.Background(), serverId, subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.ShowServerIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowServerIp`: ServerShowResponseAddress
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersAPI.ShowServerIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 
**subnetId** | **string** | Subnet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowServerIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServerShowResponseAddress**](ServerShowResponseAddress.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowServerVolume

> ServersVolumeResponse ShowServerVolume(ctx, serverId, volumeId).Execute()

Show Server Volume



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
	volumeId := "3f500f00-ed80-4566-a057-e31760226f9a" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersAPI.ShowServerVolume(context.Background(), serverId, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.ShowServerVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowServerVolume`: ServersVolumeResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersAPI.ShowServerVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowServerVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServersVolumeResponse**](ServersVolumeResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVirtualServer

> ServerShowResponse ShowVirtualServer(ctx, serverId).Execute()

Show Virtual Server



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersAPI.ShowVirtualServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.ShowVirtualServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVirtualServer`: ServerShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersAPI.ShowVirtualServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerShowResponse**](ServerShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerInterface

> UpdateServerInterface(ctx, serverId, portId).ServerInterfaceUpdateRequest(serverInterfaceUpdateRequest).Execute()

Update Server Interface



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
	portId := "f2b02fc5-a5fd-4b5f-b1a9-987f4c0d66cc" // string | Port ID
	serverInterfaceUpdateRequest := *openapiclient.NewServerInterfaceUpdateRequest() // ServerInterfaceUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServersAPI.UpdateServerInterface(context.Background(), serverId, portId).ServerInterfaceUpdateRequest(serverInterfaceUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.UpdateServerInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serverInterfaceUpdateRequest** | [**ServerInterfaceUpdateRequest**](ServerInterfaceUpdateRequest.md) |  | 

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


## UpdateServerVolume

> UpdateServerVolume(ctx, serverId, volumeId).ServerVolumesUpdateRequest(serverVolumesUpdateRequest).Execute()

Update Server Volume



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
	volumeId := "3f500f00-ed80-4566-a057-e31760226f9a" // string | Volume ID
	serverVolumesUpdateRequest := *openapiclient.NewServerVolumesUpdateRequest("3f500f00-ed80-4566-a057-e31760226f9a") // ServerVolumesUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ServersAPI.UpdateServerVolume(context.Background(), serverId, volumeId).ServerVolumesUpdateRequest(serverVolumesUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.UpdateServerVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serverVolumesUpdateRequest** | [**ServerVolumesUpdateRequest**](ServerVolumesUpdateRequest.md) |  | 

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


## UpdateVirtualServer

> ServerShowResponse UpdateVirtualServer(ctx, serverId).ServerUpdateRequest(serverUpdateRequest).Execute()

Update Virtual Server



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
	serverUpdateRequest := *openapiclient.NewServerUpdateRequest("servername") // ServerUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServersAPI.UpdateVirtualServer(context.Background(), serverId).ServerUpdateRequest(serverUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServersAPI.UpdateVirtualServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVirtualServer`: ServerShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServersAPI.UpdateVirtualServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverUpdateRequest** | [**ServerUpdateRequest**](ServerUpdateRequest.md) |  | 

### Return type

[**ServerShowResponse**](ServerShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

