# RedisInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FloatingIpId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Name | 
**PublicIpAddress** | Pointer to **NullableString** |  | [optional] 
**PublicIpId** | Pointer to **NullableString** |  | [optional] 
**RoleType** | [**InstanceRoleType**](InstanceRoleType.md) | Role type | 
**ServiceIpAddress** | Pointer to **NullableString** |  | [optional] 
**ServiceState** | [**ServiceState**](ServiceState.md) | Service state | 

## Methods

### NewRedisInstanceResponse

`func NewRedisInstanceResponse(name string, roleType InstanceRoleType, serviceState ServiceState, ) *RedisInstanceResponse`

NewRedisInstanceResponse instantiates a new RedisInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedisInstanceResponseWithDefaults

`func NewRedisInstanceResponseWithDefaults() *RedisInstanceResponse`

NewRedisInstanceResponseWithDefaults instantiates a new RedisInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloatingIpId

`func (o *RedisInstanceResponse) GetFloatingIpId() string`

GetFloatingIpId returns the FloatingIpId field if non-nil, zero value otherwise.

### GetFloatingIpIdOk

`func (o *RedisInstanceResponse) GetFloatingIpIdOk() (*string, bool)`

GetFloatingIpIdOk returns a tuple with the FloatingIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpId

`func (o *RedisInstanceResponse) SetFloatingIpId(v string)`

SetFloatingIpId sets FloatingIpId field to given value.

### HasFloatingIpId

`func (o *RedisInstanceResponse) HasFloatingIpId() bool`

HasFloatingIpId returns a boolean if a field has been set.

### SetFloatingIpIdNil

`func (o *RedisInstanceResponse) SetFloatingIpIdNil(b bool)`

 SetFloatingIpIdNil sets the value for FloatingIpId to be an explicit nil

### UnsetFloatingIpId
`func (o *RedisInstanceResponse) UnsetFloatingIpId()`

UnsetFloatingIpId ensures that no value is present for FloatingIpId, not even an explicit nil
### GetName

`func (o *RedisInstanceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RedisInstanceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RedisInstanceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPublicIpAddress

`func (o *RedisInstanceResponse) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *RedisInstanceResponse) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *RedisInstanceResponse) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *RedisInstanceResponse) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### SetPublicIpAddressNil

`func (o *RedisInstanceResponse) SetPublicIpAddressNil(b bool)`

 SetPublicIpAddressNil sets the value for PublicIpAddress to be an explicit nil

### UnsetPublicIpAddress
`func (o *RedisInstanceResponse) UnsetPublicIpAddress()`

UnsetPublicIpAddress ensures that no value is present for PublicIpAddress, not even an explicit nil
### GetPublicIpId

`func (o *RedisInstanceResponse) GetPublicIpId() string`

GetPublicIpId returns the PublicIpId field if non-nil, zero value otherwise.

### GetPublicIpIdOk

`func (o *RedisInstanceResponse) GetPublicIpIdOk() (*string, bool)`

GetPublicIpIdOk returns a tuple with the PublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpId

`func (o *RedisInstanceResponse) SetPublicIpId(v string)`

SetPublicIpId sets PublicIpId field to given value.

### HasPublicIpId

`func (o *RedisInstanceResponse) HasPublicIpId() bool`

HasPublicIpId returns a boolean if a field has been set.

### SetPublicIpIdNil

`func (o *RedisInstanceResponse) SetPublicIpIdNil(b bool)`

 SetPublicIpIdNil sets the value for PublicIpId to be an explicit nil

### UnsetPublicIpId
`func (o *RedisInstanceResponse) UnsetPublicIpId()`

UnsetPublicIpId ensures that no value is present for PublicIpId, not even an explicit nil
### GetRoleType

`func (o *RedisInstanceResponse) GetRoleType() InstanceRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *RedisInstanceResponse) GetRoleTypeOk() (*InstanceRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *RedisInstanceResponse) SetRoleType(v InstanceRoleType)`

SetRoleType sets RoleType field to given value.


### GetServiceIpAddress

`func (o *RedisInstanceResponse) GetServiceIpAddress() string`

GetServiceIpAddress returns the ServiceIpAddress field if non-nil, zero value otherwise.

### GetServiceIpAddressOk

`func (o *RedisInstanceResponse) GetServiceIpAddressOk() (*string, bool)`

GetServiceIpAddressOk returns a tuple with the ServiceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIpAddress

`func (o *RedisInstanceResponse) SetServiceIpAddress(v string)`

SetServiceIpAddress sets ServiceIpAddress field to given value.

### HasServiceIpAddress

`func (o *RedisInstanceResponse) HasServiceIpAddress() bool`

HasServiceIpAddress returns a boolean if a field has been set.

### SetServiceIpAddressNil

`func (o *RedisInstanceResponse) SetServiceIpAddressNil(b bool)`

 SetServiceIpAddressNil sets the value for ServiceIpAddress to be an explicit nil

### UnsetServiceIpAddress
`func (o *RedisInstanceResponse) UnsetServiceIpAddress()`

UnsetServiceIpAddress ensures that no value is present for ServiceIpAddress, not even an explicit nil
### GetServiceState

`func (o *RedisInstanceResponse) GetServiceState() ServiceState`

GetServiceState returns the ServiceState field if non-nil, zero value otherwise.

### GetServiceStateOk

`func (o *RedisInstanceResponse) GetServiceStateOk() (*ServiceState, bool)`

GetServiceStateOk returns a tuple with the ServiceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceState

`func (o *RedisInstanceResponse) SetServiceState(v ServiceState)`

SetServiceState sets ServiceState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


