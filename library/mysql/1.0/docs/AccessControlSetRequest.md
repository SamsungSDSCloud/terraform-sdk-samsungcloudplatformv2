# AccessControlSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddAccessControl** | Pointer to [**[]AccessControlRequest**](AccessControlRequest.md) |  | [optional] 
**DelAccessControl** | Pointer to [**[]AccessControlRequest**](AccessControlRequest.md) |  | [optional] 

## Methods

### NewAccessControlSetRequest

`func NewAccessControlSetRequest() *AccessControlSetRequest`

NewAccessControlSetRequest instantiates a new AccessControlSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlSetRequestWithDefaults

`func NewAccessControlSetRequestWithDefaults() *AccessControlSetRequest`

NewAccessControlSetRequestWithDefaults instantiates a new AccessControlSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddAccessControl

`func (o *AccessControlSetRequest) GetAddAccessControl() []AccessControlRequest`

GetAddAccessControl returns the AddAccessControl field if non-nil, zero value otherwise.

### GetAddAccessControlOk

`func (o *AccessControlSetRequest) GetAddAccessControlOk() (*[]AccessControlRequest, bool)`

GetAddAccessControlOk returns a tuple with the AddAccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessControl

`func (o *AccessControlSetRequest) SetAddAccessControl(v []AccessControlRequest)`

SetAddAccessControl sets AddAccessControl field to given value.

### HasAddAccessControl

`func (o *AccessControlSetRequest) HasAddAccessControl() bool`

HasAddAccessControl returns a boolean if a field has been set.

### SetAddAccessControlNil

`func (o *AccessControlSetRequest) SetAddAccessControlNil(b bool)`

 SetAddAccessControlNil sets the value for AddAccessControl to be an explicit nil

### UnsetAddAccessControl
`func (o *AccessControlSetRequest) UnsetAddAccessControl()`

UnsetAddAccessControl ensures that no value is present for AddAccessControl, not even an explicit nil
### GetDelAccessControl

`func (o *AccessControlSetRequest) GetDelAccessControl() []AccessControlRequest`

GetDelAccessControl returns the DelAccessControl field if non-nil, zero value otherwise.

### GetDelAccessControlOk

`func (o *AccessControlSetRequest) GetDelAccessControlOk() (*[]AccessControlRequest, bool)`

GetDelAccessControlOk returns a tuple with the DelAccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelAccessControl

`func (o *AccessControlSetRequest) SetDelAccessControl(v []AccessControlRequest)`

SetDelAccessControl sets DelAccessControl field to given value.

### HasDelAccessControl

`func (o *AccessControlSetRequest) HasDelAccessControl() bool`

HasDelAccessControl returns a boolean if a field has been set.

### SetDelAccessControlNil

`func (o *AccessControlSetRequest) SetDelAccessControlNil(b bool)`

 SetDelAccessControlNil sets the value for DelAccessControl to be an explicit nil

### UnsetDelAccessControl
`func (o *AccessControlSetRequest) UnsetDelAccessControl()`

UnsetDelAccessControl ensures that no value is present for DelAccessControl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


