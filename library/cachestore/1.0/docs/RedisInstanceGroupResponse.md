# RedisInstanceGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockStorageGroups** | [**[]BlockStorageGroupResponse**](BlockStorageGroupResponse.md) | Block storage groups list | 
**Id** | **string** | ID | 
**Instances** | [**[]RedisInstanceResponse**](RedisInstanceResponse.md) |  | 
**RoleType** | [**InstanceGroupRoleType**](InstanceGroupRoleType.md) | Role type | 
**ServerTypeName** | **string** | Server type name | 

## Methods

### NewRedisInstanceGroupResponse

`func NewRedisInstanceGroupResponse(blockStorageGroups []BlockStorageGroupResponse, id string, instances []RedisInstanceResponse, roleType InstanceGroupRoleType, serverTypeName string, ) *RedisInstanceGroupResponse`

NewRedisInstanceGroupResponse instantiates a new RedisInstanceGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedisInstanceGroupResponseWithDefaults

`func NewRedisInstanceGroupResponseWithDefaults() *RedisInstanceGroupResponse`

NewRedisInstanceGroupResponseWithDefaults instantiates a new RedisInstanceGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockStorageGroups

`func (o *RedisInstanceGroupResponse) GetBlockStorageGroups() []BlockStorageGroupResponse`

GetBlockStorageGroups returns the BlockStorageGroups field if non-nil, zero value otherwise.

### GetBlockStorageGroupsOk

`func (o *RedisInstanceGroupResponse) GetBlockStorageGroupsOk() (*[]BlockStorageGroupResponse, bool)`

GetBlockStorageGroupsOk returns a tuple with the BlockStorageGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockStorageGroups

`func (o *RedisInstanceGroupResponse) SetBlockStorageGroups(v []BlockStorageGroupResponse)`

SetBlockStorageGroups sets BlockStorageGroups field to given value.


### GetId

`func (o *RedisInstanceGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RedisInstanceGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RedisInstanceGroupResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInstances

`func (o *RedisInstanceGroupResponse) GetInstances() []RedisInstanceResponse`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *RedisInstanceGroupResponse) GetInstancesOk() (*[]RedisInstanceResponse, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *RedisInstanceGroupResponse) SetInstances(v []RedisInstanceResponse)`

SetInstances sets Instances field to given value.


### SetInstancesNil

`func (o *RedisInstanceGroupResponse) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *RedisInstanceGroupResponse) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil
### GetRoleType

`func (o *RedisInstanceGroupResponse) GetRoleType() InstanceGroupRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *RedisInstanceGroupResponse) GetRoleTypeOk() (*InstanceGroupRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *RedisInstanceGroupResponse) SetRoleType(v InstanceGroupRoleType)`

SetRoleType sets RoleType field to given value.


### GetServerTypeName

`func (o *RedisInstanceGroupResponse) GetServerTypeName() string`

GetServerTypeName returns the ServerTypeName field if non-nil, zero value otherwise.

### GetServerTypeNameOk

`func (o *RedisInstanceGroupResponse) GetServerTypeNameOk() (*string, bool)`

GetServerTypeNameOk returns a tuple with the ServerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeName

`func (o *RedisInstanceGroupResponse) SetServerTypeName(v string)`

SetServerTypeName sets ServerTypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


