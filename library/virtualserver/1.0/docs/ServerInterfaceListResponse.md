# ServerInterfaceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interfaces** | [**[]InterfaceResponse**](InterfaceResponse.md) | Interface List | 

## Methods

### NewServerInterfaceListResponse

`func NewServerInterfaceListResponse(interfaces []InterfaceResponse, ) *ServerInterfaceListResponse`

NewServerInterfaceListResponse instantiates a new ServerInterfaceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInterfaceListResponseWithDefaults

`func NewServerInterfaceListResponseWithDefaults() *ServerInterfaceListResponse`

NewServerInterfaceListResponseWithDefaults instantiates a new ServerInterfaceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaces

`func (o *ServerInterfaceListResponse) GetInterfaces() []InterfaceResponse`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ServerInterfaceListResponse) GetInterfacesOk() (*[]InterfaceResponse, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ServerInterfaceListResponse) SetInterfaces(v []InterfaceResponse)`

SetInterfaces sets Interfaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


