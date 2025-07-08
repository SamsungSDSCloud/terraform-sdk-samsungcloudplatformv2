# InstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**PublicIpAddress** | Pointer to **NullableString** |  | [optional] 
**PublicIpId** | Pointer to **NullableString** |  | [optional] 
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


### GetPublicIpAddress

`func (o *InstanceResponse) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *InstanceResponse) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *InstanceResponse) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *InstanceResponse) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### SetPublicIpAddressNil

`func (o *InstanceResponse) SetPublicIpAddressNil(b bool)`

 SetPublicIpAddressNil sets the value for PublicIpAddress to be an explicit nil

### UnsetPublicIpAddress
`func (o *InstanceResponse) UnsetPublicIpAddress()`

UnsetPublicIpAddress ensures that no value is present for PublicIpAddress, not even an explicit nil
### GetPublicIpId

`func (o *InstanceResponse) GetPublicIpId() string`

GetPublicIpId returns the PublicIpId field if non-nil, zero value otherwise.

### GetPublicIpIdOk

`func (o *InstanceResponse) GetPublicIpIdOk() (*string, bool)`

GetPublicIpIdOk returns a tuple with the PublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpId

`func (o *InstanceResponse) SetPublicIpId(v string)`

SetPublicIpId sets PublicIpId field to given value.

### HasPublicIpId

`func (o *InstanceResponse) HasPublicIpId() bool`

HasPublicIpId returns a boolean if a field has been set.

### SetPublicIpIdNil

`func (o *InstanceResponse) SetPublicIpIdNil(b bool)`

 SetPublicIpIdNil sets the value for PublicIpId to be an explicit nil

### UnsetPublicIpId
`func (o *InstanceResponse) UnsetPublicIpId()`

UnsetPublicIpId ensures that no value is present for PublicIpId, not even an explicit nil
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


