# InstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIpId** | Pointer to **NullableString** |  | [optional] 
**RoleType** | [**InstanceRoleType**](InstanceRoleType.md) | Role type | 
**ServiceIpAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInstanceRequest

`func NewInstanceRequest(roleType InstanceRoleType, ) *InstanceRequest`

NewInstanceRequest instantiates a new InstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceRequestWithDefaults

`func NewInstanceRequestWithDefaults() *InstanceRequest`

NewInstanceRequestWithDefaults instantiates a new InstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIpId

`func (o *InstanceRequest) GetPublicIpId() string`

GetPublicIpId returns the PublicIpId field if non-nil, zero value otherwise.

### GetPublicIpIdOk

`func (o *InstanceRequest) GetPublicIpIdOk() (*string, bool)`

GetPublicIpIdOk returns a tuple with the PublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpId

`func (o *InstanceRequest) SetPublicIpId(v string)`

SetPublicIpId sets PublicIpId field to given value.

### HasPublicIpId

`func (o *InstanceRequest) HasPublicIpId() bool`

HasPublicIpId returns a boolean if a field has been set.

### SetPublicIpIdNil

`func (o *InstanceRequest) SetPublicIpIdNil(b bool)`

 SetPublicIpIdNil sets the value for PublicIpId to be an explicit nil

### UnsetPublicIpId
`func (o *InstanceRequest) UnsetPublicIpId()`

UnsetPublicIpId ensures that no value is present for PublicIpId, not even an explicit nil
### GetRoleType

`func (o *InstanceRequest) GetRoleType() InstanceRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *InstanceRequest) GetRoleTypeOk() (*InstanceRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *InstanceRequest) SetRoleType(v InstanceRoleType)`

SetRoleType sets RoleType field to given value.


### GetServiceIpAddress

`func (o *InstanceRequest) GetServiceIpAddress() string`

GetServiceIpAddress returns the ServiceIpAddress field if non-nil, zero value otherwise.

### GetServiceIpAddressOk

`func (o *InstanceRequest) GetServiceIpAddressOk() (*string, bool)`

GetServiceIpAddressOk returns a tuple with the ServiceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIpAddress

`func (o *InstanceRequest) SetServiceIpAddress(v string)`

SetServiceIpAddress sets ServiceIpAddress field to given value.

### HasServiceIpAddress

`func (o *InstanceRequest) HasServiceIpAddress() bool`

HasServiceIpAddress returns a boolean if a field has been set.

### SetServiceIpAddressNil

`func (o *InstanceRequest) SetServiceIpAddressNil(b bool)`

 SetServiceIpAddressNil sets the value for ServiceIpAddress to be an explicit nil

### UnsetServiceIpAddress
`func (o *InstanceRequest) UnsetServiceIpAddress()`

UnsetServiceIpAddress ensures that no value is present for ServiceIpAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


