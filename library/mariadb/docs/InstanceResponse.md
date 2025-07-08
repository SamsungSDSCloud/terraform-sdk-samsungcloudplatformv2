# InstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**RoleType** | [**InstanceRoleType**](InstanceRoleType.md) | Role type | 
**ServiceIpAddress** | Pointer to **NullableString** |  | [optional] 
**ServiceState** | [**ServiceState**](ServiceState.md) | Service state | 

## Methods

### NewInstanceResponse

`func NewInstanceResponse(name string, roleType InstanceRoleType, serviceState ServiceState, ) *InstanceResponse`

NewInstanceResponse instantiates a new InstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceResponseWithDefaults

`func NewInstanceResponseWithDefaults() *InstanceResponse`

NewInstanceResponseWithDefaults instantiates a new InstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRoleType

`func (o *InstanceResponse) GetRoleType() InstanceRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *InstanceResponse) GetRoleTypeOk() (*InstanceRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *InstanceResponse) SetRoleType(v InstanceRoleType)`

SetRoleType sets RoleType field to given value.


### GetServiceIpAddress

`func (o *InstanceResponse) GetServiceIpAddress() string`

GetServiceIpAddress returns the ServiceIpAddress field if non-nil, zero value otherwise.

### GetServiceIpAddressOk

`func (o *InstanceResponse) GetServiceIpAddressOk() (*string, bool)`

GetServiceIpAddressOk returns a tuple with the ServiceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIpAddress

`func (o *InstanceResponse) SetServiceIpAddress(v string)`

SetServiceIpAddress sets ServiceIpAddress field to given value.

### HasServiceIpAddress

`func (o *InstanceResponse) HasServiceIpAddress() bool`

HasServiceIpAddress returns a boolean if a field has been set.

### SetServiceIpAddressNil

`func (o *InstanceResponse) SetServiceIpAddressNil(b bool)`

 SetServiceIpAddressNil sets the value for ServiceIpAddress to be an explicit nil

### UnsetServiceIpAddress
`func (o *InstanceResponse) UnsetServiceIpAddress()`

UnsetServiceIpAddress ensures that no value is present for ServiceIpAddress, not even an explicit nil
### GetServiceState

`func (o *InstanceResponse) GetServiceState() ServiceState`

GetServiceState returns the ServiceState field if non-nil, zero value otherwise.

### GetServiceStateOk

`func (o *InstanceResponse) GetServiceStateOk() (*ServiceState, bool)`

GetServiceStateOk returns a tuple with the ServiceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceState

`func (o *InstanceResponse) SetServiceState(v ServiceState)`

SetServiceState sets ServiceState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


