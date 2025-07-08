# SqlserverInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIpId** | Pointer to **NullableString** |  | [optional] 
**RoleType** | [**InstanceRoleType**](InstanceRoleType.md) | Role type | 
**ServiceIpAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSqlserverInstanceRequest

`func NewSqlserverInstanceRequest(roleType InstanceRoleType, ) *SqlserverInstanceRequest`

NewSqlserverInstanceRequest instantiates a new SqlserverInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlserverInstanceRequestWithDefaults

`func NewSqlserverInstanceRequestWithDefaults() *SqlserverInstanceRequest`

NewSqlserverInstanceRequestWithDefaults instantiates a new SqlserverInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIpId

`func (o *SqlserverInstanceRequest) GetPublicIpId() string`

GetPublicIpId returns the PublicIpId field if non-nil, zero value otherwise.

### GetPublicIpIdOk

`func (o *SqlserverInstanceRequest) GetPublicIpIdOk() (*string, bool)`

GetPublicIpIdOk returns a tuple with the PublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpId

`func (o *SqlserverInstanceRequest) SetPublicIpId(v string)`

SetPublicIpId sets PublicIpId field to given value.

### HasPublicIpId

`func (o *SqlserverInstanceRequest) HasPublicIpId() bool`

HasPublicIpId returns a boolean if a field has been set.

### SetPublicIpIdNil

`func (o *SqlserverInstanceRequest) SetPublicIpIdNil(b bool)`

 SetPublicIpIdNil sets the value for PublicIpId to be an explicit nil

### UnsetPublicIpId
`func (o *SqlserverInstanceRequest) UnsetPublicIpId()`

UnsetPublicIpId ensures that no value is present for PublicIpId, not even an explicit nil
### GetRoleType

`func (o *SqlserverInstanceRequest) GetRoleType() InstanceRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *SqlserverInstanceRequest) GetRoleTypeOk() (*InstanceRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *SqlserverInstanceRequest) SetRoleType(v InstanceRoleType)`

SetRoleType sets RoleType field to given value.


### GetServiceIpAddress

`func (o *SqlserverInstanceRequest) GetServiceIpAddress() string`

GetServiceIpAddress returns the ServiceIpAddress field if non-nil, zero value otherwise.

### GetServiceIpAddressOk

`func (o *SqlserverInstanceRequest) GetServiceIpAddressOk() (*string, bool)`

GetServiceIpAddressOk returns a tuple with the ServiceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIpAddress

`func (o *SqlserverInstanceRequest) SetServiceIpAddress(v string)`

SetServiceIpAddress sets ServiceIpAddress field to given value.

### HasServiceIpAddress

`func (o *SqlserverInstanceRequest) HasServiceIpAddress() bool`

HasServiceIpAddress returns a boolean if a field has been set.

### SetServiceIpAddressNil

`func (o *SqlserverInstanceRequest) SetServiceIpAddressNil(b bool)`

 SetServiceIpAddressNil sets the value for ServiceIpAddress to be an explicit nil

### UnsetServiceIpAddress
`func (o *SqlserverInstanceRequest) UnsetServiceIpAddress()`

UnsetServiceIpAddress ensures that no value is present for ServiceIpAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


