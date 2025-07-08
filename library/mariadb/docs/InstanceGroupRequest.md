# InstanceGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockStorageGroups** | [**[]BlockStorageGroupRequest**](BlockStorageGroupRequest.md) | Block storage groups list | 
**Instances** | Pointer to [**[]InstanceRequest**](InstanceRequest.md) |  | [optional] 
**RoleType** | [**InstanceGroupRoleType**](InstanceGroupRoleType.md) | Role type | 
**ServerTypeName** | **string** | Server type name | 

## Methods

### NewInstanceGroupRequest

`func NewInstanceGroupRequest(blockStorageGroups []BlockStorageGroupRequest, roleType InstanceGroupRoleType, serverTypeName string, ) *InstanceGroupRequest`

NewInstanceGroupRequest instantiates a new InstanceGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceGroupRequestWithDefaults

`func NewInstanceGroupRequestWithDefaults() *InstanceGroupRequest`

NewInstanceGroupRequestWithDefaults instantiates a new InstanceGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockStorageGroups

`func (o *InstanceGroupRequest) GetBlockStorageGroups() []BlockStorageGroupRequest`

GetBlockStorageGroups returns the BlockStorageGroups field if non-nil, zero value otherwise.

### GetBlockStorageGroupsOk

`func (o *InstanceGroupRequest) GetBlockStorageGroupsOk() (*[]BlockStorageGroupRequest, bool)`

GetBlockStorageGroupsOk returns a tuple with the BlockStorageGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockStorageGroups

`func (o *InstanceGroupRequest) SetBlockStorageGroups(v []BlockStorageGroupRequest)`

SetBlockStorageGroups sets BlockStorageGroups field to given value.


### GetInstances

`func (o *InstanceGroupRequest) GetInstances() []InstanceRequest`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *InstanceGroupRequest) GetInstancesOk() (*[]InstanceRequest, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *InstanceGroupRequest) SetInstances(v []InstanceRequest)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *InstanceGroupRequest) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### SetInstancesNil

`func (o *InstanceGroupRequest) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *InstanceGroupRequest) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil
### GetRoleType

`func (o *InstanceGroupRequest) GetRoleType() InstanceGroupRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *InstanceGroupRequest) GetRoleTypeOk() (*InstanceGroupRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *InstanceGroupRequest) SetRoleType(v InstanceGroupRoleType)`

SetRoleType sets RoleType field to given value.


### GetServerTypeName

`func (o *InstanceGroupRequest) GetServerTypeName() string`

GetServerTypeName returns the ServerTypeName field if non-nil, zero value otherwise.

### GetServerTypeNameOk

`func (o *InstanceGroupRequest) GetServerTypeNameOk() (*string, bool)`

GetServerTypeNameOk returns a tuple with the ServerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeName

`func (o *InstanceGroupRequest) SetServerTypeName(v string)`

SetServerTypeName sets ServerTypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


