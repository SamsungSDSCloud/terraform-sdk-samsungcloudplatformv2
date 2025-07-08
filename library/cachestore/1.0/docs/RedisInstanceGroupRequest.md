# RedisInstanceGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockStorageGroups** | [**[]RedisBlockStorageGroupRequest**](RedisBlockStorageGroupRequest.md) | Block storage groups list | 
**Instances** | [**[]InstanceRequest**](InstanceRequest.md) |  | 
**RoleType** | [**InstanceGroupRoleType**](InstanceGroupRoleType.md) | Role type | 
**ServerTypeName** | **string** | Server type name | 

## Methods

### NewRedisInstanceGroupRequest

`func NewRedisInstanceGroupRequest(blockStorageGroups []RedisBlockStorageGroupRequest, instances []InstanceRequest, roleType InstanceGroupRoleType, serverTypeName string, ) *RedisInstanceGroupRequest`

NewRedisInstanceGroupRequest instantiates a new RedisInstanceGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedisInstanceGroupRequestWithDefaults

`func NewRedisInstanceGroupRequestWithDefaults() *RedisInstanceGroupRequest`

NewRedisInstanceGroupRequestWithDefaults instantiates a new RedisInstanceGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockStorageGroups

`func (o *RedisInstanceGroupRequest) GetBlockStorageGroups() []RedisBlockStorageGroupRequest`

GetBlockStorageGroups returns the BlockStorageGroups field if non-nil, zero value otherwise.

### GetBlockStorageGroupsOk

`func (o *RedisInstanceGroupRequest) GetBlockStorageGroupsOk() (*[]RedisBlockStorageGroupRequest, bool)`

GetBlockStorageGroupsOk returns a tuple with the BlockStorageGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockStorageGroups

`func (o *RedisInstanceGroupRequest) SetBlockStorageGroups(v []RedisBlockStorageGroupRequest)`

SetBlockStorageGroups sets BlockStorageGroups field to given value.


### GetInstances

`func (o *RedisInstanceGroupRequest) GetInstances() []InstanceRequest`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *RedisInstanceGroupRequest) GetInstancesOk() (*[]InstanceRequest, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *RedisInstanceGroupRequest) SetInstances(v []InstanceRequest)`

SetInstances sets Instances field to given value.


### SetInstancesNil

`func (o *RedisInstanceGroupRequest) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *RedisInstanceGroupRequest) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil
### GetRoleType

`func (o *RedisInstanceGroupRequest) GetRoleType() InstanceGroupRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *RedisInstanceGroupRequest) GetRoleTypeOk() (*InstanceGroupRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *RedisInstanceGroupRequest) SetRoleType(v InstanceGroupRoleType)`

SetRoleType sets RoleType field to given value.


### GetServerTypeName

`func (o *RedisInstanceGroupRequest) GetServerTypeName() string`

GetServerTypeName returns the ServerTypeName field if non-nil, zero value otherwise.

### GetServerTypeNameOk

`func (o *RedisInstanceGroupRequest) GetServerTypeNameOk() (*string, bool)`

GetServerTypeNameOk returns a tuple with the ServerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeName

`func (o *RedisInstanceGroupRequest) SetServerTypeName(v string)`

SetServerTypeName sets ServerTypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


