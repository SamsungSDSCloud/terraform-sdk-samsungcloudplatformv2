# ClusterRestoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowableIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**BackupHistoryNumber** | Pointer to **NullableString** |  | [optional] 
**BackupRecoveryTime** | Pointer to **NullableTime** |  | [optional] 
**BlockStorageGroups** | Pointer to [**[]BlockStorageGroupRequest**](BlockStorageGroupRequest.md) |  | [optional] 
**InstanceNamePrefix** | **string** | Instance name prefix | 
**MaintenanceOption** | Pointer to [**NullableMaintenanceOption**](MaintenanceOption.md) |  | [optional] 
**Name** | **string** | Cluster name | 
**ServerTypeName** | **string** | Server type name | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewClusterRestoreRequest

`func NewClusterRestoreRequest(instanceNamePrefix string, name string, serverTypeName string, ) *ClusterRestoreRequest`

NewClusterRestoreRequest instantiates a new ClusterRestoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRestoreRequestWithDefaults

`func NewClusterRestoreRequestWithDefaults() *ClusterRestoreRequest`

NewClusterRestoreRequestWithDefaults instantiates a new ClusterRestoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowableIpAddresses

`func (o *ClusterRestoreRequest) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *ClusterRestoreRequest) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *ClusterRestoreRequest) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *ClusterRestoreRequest) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetBackupHistoryNumber

`func (o *ClusterRestoreRequest) GetBackupHistoryNumber() string`

GetBackupHistoryNumber returns the BackupHistoryNumber field if non-nil, zero value otherwise.

### GetBackupHistoryNumberOk

`func (o *ClusterRestoreRequest) GetBackupHistoryNumberOk() (*string, bool)`

GetBackupHistoryNumberOk returns a tuple with the BackupHistoryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupHistoryNumber

`func (o *ClusterRestoreRequest) SetBackupHistoryNumber(v string)`

SetBackupHistoryNumber sets BackupHistoryNumber field to given value.

### HasBackupHistoryNumber

`func (o *ClusterRestoreRequest) HasBackupHistoryNumber() bool`

HasBackupHistoryNumber returns a boolean if a field has been set.

### SetBackupHistoryNumberNil

`func (o *ClusterRestoreRequest) SetBackupHistoryNumberNil(b bool)`

 SetBackupHistoryNumberNil sets the value for BackupHistoryNumber to be an explicit nil

### UnsetBackupHistoryNumber
`func (o *ClusterRestoreRequest) UnsetBackupHistoryNumber()`

UnsetBackupHistoryNumber ensures that no value is present for BackupHistoryNumber, not even an explicit nil
### GetBackupRecoveryTime

`func (o *ClusterRestoreRequest) GetBackupRecoveryTime() time.Time`

GetBackupRecoveryTime returns the BackupRecoveryTime field if non-nil, zero value otherwise.

### GetBackupRecoveryTimeOk

`func (o *ClusterRestoreRequest) GetBackupRecoveryTimeOk() (*time.Time, bool)`

GetBackupRecoveryTimeOk returns a tuple with the BackupRecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRecoveryTime

`func (o *ClusterRestoreRequest) SetBackupRecoveryTime(v time.Time)`

SetBackupRecoveryTime sets BackupRecoveryTime field to given value.

### HasBackupRecoveryTime

`func (o *ClusterRestoreRequest) HasBackupRecoveryTime() bool`

HasBackupRecoveryTime returns a boolean if a field has been set.

### SetBackupRecoveryTimeNil

`func (o *ClusterRestoreRequest) SetBackupRecoveryTimeNil(b bool)`

 SetBackupRecoveryTimeNil sets the value for BackupRecoveryTime to be an explicit nil

### UnsetBackupRecoveryTime
`func (o *ClusterRestoreRequest) UnsetBackupRecoveryTime()`

UnsetBackupRecoveryTime ensures that no value is present for BackupRecoveryTime, not even an explicit nil
### GetBlockStorageGroups

`func (o *ClusterRestoreRequest) GetBlockStorageGroups() []BlockStorageGroupRequest`

GetBlockStorageGroups returns the BlockStorageGroups field if non-nil, zero value otherwise.

### GetBlockStorageGroupsOk

`func (o *ClusterRestoreRequest) GetBlockStorageGroupsOk() (*[]BlockStorageGroupRequest, bool)`

GetBlockStorageGroupsOk returns a tuple with the BlockStorageGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockStorageGroups

`func (o *ClusterRestoreRequest) SetBlockStorageGroups(v []BlockStorageGroupRequest)`

SetBlockStorageGroups sets BlockStorageGroups field to given value.

### HasBlockStorageGroups

`func (o *ClusterRestoreRequest) HasBlockStorageGroups() bool`

HasBlockStorageGroups returns a boolean if a field has been set.

### SetBlockStorageGroupsNil

`func (o *ClusterRestoreRequest) SetBlockStorageGroupsNil(b bool)`

 SetBlockStorageGroupsNil sets the value for BlockStorageGroups to be an explicit nil

### UnsetBlockStorageGroups
`func (o *ClusterRestoreRequest) UnsetBlockStorageGroups()`

UnsetBlockStorageGroups ensures that no value is present for BlockStorageGroups, not even an explicit nil
### GetInstanceNamePrefix

`func (o *ClusterRestoreRequest) GetInstanceNamePrefix() string`

GetInstanceNamePrefix returns the InstanceNamePrefix field if non-nil, zero value otherwise.

### GetInstanceNamePrefixOk

`func (o *ClusterRestoreRequest) GetInstanceNamePrefixOk() (*string, bool)`

GetInstanceNamePrefixOk returns a tuple with the InstanceNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceNamePrefix

`func (o *ClusterRestoreRequest) SetInstanceNamePrefix(v string)`

SetInstanceNamePrefix sets InstanceNamePrefix field to given value.


### GetMaintenanceOption

`func (o *ClusterRestoreRequest) GetMaintenanceOption() MaintenanceOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *ClusterRestoreRequest) GetMaintenanceOptionOk() (*MaintenanceOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *ClusterRestoreRequest) SetMaintenanceOption(v MaintenanceOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *ClusterRestoreRequest) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *ClusterRestoreRequest) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *ClusterRestoreRequest) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetName

`func (o *ClusterRestoreRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterRestoreRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterRestoreRequest) SetName(v string)`

SetName sets Name field to given value.


### GetServerTypeName

`func (o *ClusterRestoreRequest) GetServerTypeName() string`

GetServerTypeName returns the ServerTypeName field if non-nil, zero value otherwise.

### GetServerTypeNameOk

`func (o *ClusterRestoreRequest) GetServerTypeNameOk() (*string, bool)`

GetServerTypeNameOk returns a tuple with the ServerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeName

`func (o *ClusterRestoreRequest) SetServerTypeName(v string)`

SetServerTypeName sets ServerTypeName field to given value.


### GetTags

`func (o *ClusterRestoreRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClusterRestoreRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClusterRestoreRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClusterRestoreRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ClusterRestoreRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ClusterRestoreRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


