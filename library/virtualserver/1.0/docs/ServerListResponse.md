# ServerListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | [**[]ServerShowResponse**](ServerShowResponse.md) | Server list | 

## Methods

### NewServerListResponse

`func NewServerListResponse(servers []ServerShowResponse, ) *ServerListResponse`

NewServerListResponse instantiates a new ServerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerListResponseWithDefaults

`func NewServerListResponseWithDefaults() *ServerListResponse`

NewServerListResponseWithDefaults instantiates a new ServerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *ServerListResponse) GetServers() []ServerShowResponse`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ServerListResponse) GetServersOk() (*[]ServerShowResponse, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ServerListResponse) SetServers(v []ServerShowResponse)`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


