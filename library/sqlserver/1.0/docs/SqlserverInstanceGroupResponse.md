# SqlserverInstanceGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockStorageGroups** | [**[]SqlserverBlockStorageGroupResponse**](SqlserverBlockStorageGroupResponse.md) |  | 
**Id** | **string** | ID | 
**Instances** | [**[]InstanceResponse**](InstanceResponse.md) | Instances list | 
**RoleType** | [**InstanceGroupRoleType**](InstanceGroupRoleType.md) | Role type | 
**ServerTypeName** | **string** | Server type name | 

## Methods

### NewSqlserverInstanceGroupResponse

`func NewSqlserverInstanceGroupResponse(blockStorageGroups []SqlserverBlockStorageGroupResponse, id string, instances []InstanceResponse, roleType InstanceGroupRoleType, serverTypeName string, ) *SqlserverInstanceGroupResponse`

NewSqlserverInstanceGroupResponse instantiates a new SqlserverInstanceGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlserverInstanceGroupResponseWithDefaults

`func NewSqlserverInstanceGroupResponseWithDefaults() *SqlserverInstanceGroupResponse`

NewSqlserverInstanceGroupResponseWithDefaults instantiates a new SqlserverInstanceGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockStorageGroups

`func (o *SqlserverInstanceGroupResponse) GetBlockStorageGroups() []SqlserverBlockStorageGroupResponse`

GetBlockStorageGroups returns the BlockStorageGroups field if non-nil, zero value otherwise.

### GetBlockStorageGroupsOk

`func (o *SqlserverInstanceGroupResponse) GetBlockStorageGroupsOk() (*[]SqlserverBlockStorageGroupResponse, bool)`

GetBlockStorageGroupsOk returns a tuple with the BlockStorageGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockStorageGroups

`func (o *SqlserverInstanceGroupResponse) SetBlockStorageGroups(v []SqlserverBlockStorageGroupResponse)`

SetBlockStorageGroups sets BlockStorageGroups field to given value.


### GetId

`func (o *SqlserverInstanceGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SqlserverInstanceGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SqlserverInstanceGroupResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInstances

`func (o *SqlserverInstanceGroupResponse) GetInstances() []InstanceResponse`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *SqlserverInstanceGroupResponse) GetInstancesOk() (*[]InstanceResponse, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *SqlserverInstanceGroupResponse) SetInstances(v []InstanceResponse)`

SetInstances sets Instances field to given value.


### GetRoleType

`func (o *SqlserverInstanceGroupResponse) GetRoleType() InstanceGroupRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *SqlserverInstanceGroupResponse) GetRoleTypeOk() (*InstanceGroupRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *SqlserverInstanceGroupResponse) SetRoleType(v InstanceGroupRoleType)`

SetRoleType sets RoleType field to given value.


### GetServerTypeName

`func (o *SqlserverInstanceGroupResponse) GetServerTypeName() string`

GetServerTypeName returns the ServerTypeName field if non-nil, zero value otherwise.

### GetServerTypeNameOk

`func (o *SqlserverInstanceGroupResponse) GetServerTypeNameOk() (*string, bool)`

GetServerTypeNameOk returns a tuple with the ServerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeName

`func (o *SqlserverInstanceGroupResponse) SetServerTypeName(v string)`

SetServerTypeName sets ServerTypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


