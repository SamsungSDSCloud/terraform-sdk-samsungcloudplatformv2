# SqlserverInstanceGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockStorageGroups** | [**[]BlockStorageGroupRequest**](BlockStorageGroupRequest.md) | Block storage groups list | 
**Instances** | [**[]SqlserverInstanceRequest**](SqlserverInstanceRequest.md) | Instances list | 
**RoleType** | Pointer to **interface{}** |  | [optional] 
**ServerTypeName** | **string** | Server type name | 

## Methods

### NewSqlserverInstanceGroupRequest

`func NewSqlserverInstanceGroupRequest(blockStorageGroups []BlockStorageGroupRequest, instances []SqlserverInstanceRequest, serverTypeName string, ) *SqlserverInstanceGroupRequest`

NewSqlserverInstanceGroupRequest instantiates a new SqlserverInstanceGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlserverInstanceGroupRequestWithDefaults

`func NewSqlserverInstanceGroupRequestWithDefaults() *SqlserverInstanceGroupRequest`

NewSqlserverInstanceGroupRequestWithDefaults instantiates a new SqlserverInstanceGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockStorageGroups

`func (o *SqlserverInstanceGroupRequest) GetBlockStorageGroups() []BlockStorageGroupRequest`

GetBlockStorageGroups returns the BlockStorageGroups field if non-nil, zero value otherwise.

### GetBlockStorageGroupsOk

`func (o *SqlserverInstanceGroupRequest) GetBlockStorageGroupsOk() (*[]BlockStorageGroupRequest, bool)`

GetBlockStorageGroupsOk returns a tuple with the BlockStorageGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockStorageGroups

`func (o *SqlserverInstanceGroupRequest) SetBlockStorageGroups(v []BlockStorageGroupRequest)`

SetBlockStorageGroups sets BlockStorageGroups field to given value.


### GetInstances

`func (o *SqlserverInstanceGroupRequest) GetInstances() []SqlserverInstanceRequest`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *SqlserverInstanceGroupRequest) GetInstancesOk() (*[]SqlserverInstanceRequest, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *SqlserverInstanceGroupRequest) SetInstances(v []SqlserverInstanceRequest)`

SetInstances sets Instances field to given value.


### GetRoleType

`func (o *SqlserverInstanceGroupRequest) GetRoleType() interface{}`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *SqlserverInstanceGroupRequest) GetRoleTypeOk() (*interface{}, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *SqlserverInstanceGroupRequest) SetRoleType(v interface{})`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *SqlserverInstanceGroupRequest) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### SetRoleTypeNil

`func (o *SqlserverInstanceGroupRequest) SetRoleTypeNil(b bool)`

 SetRoleTypeNil sets the value for RoleType to be an explicit nil

### UnsetRoleType
`func (o *SqlserverInstanceGroupRequest) UnsetRoleType()`

UnsetRoleType ensures that no value is present for RoleType, not even an explicit nil
### GetServerTypeName

`func (o *SqlserverInstanceGroupRequest) GetServerTypeName() string`

GetServerTypeName returns the ServerTypeName field if non-nil, zero value otherwise.

### GetServerTypeNameOk

`func (o *SqlserverInstanceGroupRequest) GetServerTypeNameOk() (*string, bool)`

GetServerTypeNameOk returns a tuple with the ServerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeName

`func (o *SqlserverInstanceGroupRequest) SetServerTypeName(v string)`

SetServerTypeName sets ServerTypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


