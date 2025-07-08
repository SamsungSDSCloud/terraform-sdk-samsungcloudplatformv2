# SearchEngineClusterRestoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowableIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**BackupHistoryNumber** | **string** | Backup id | 
**InstanceGroups** | [**[]SearchEngineRestoreInstanceGroup**](SearchEngineRestoreInstanceGroup.md) | Instance groups list | 
**InstanceNamePrefix** | **string** | Instance name prefix | 
**License** | Pointer to **NullableString** |  | [optional] 
**MaintenanceOption** | Pointer to [**NullableMaintenanceOption**](MaintenanceOption.md) |  | [optional] 
**Name** | **string** | Cluster name | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewSearchEngineClusterRestoreRequest

`func NewSearchEngineClusterRestoreRequest(backupHistoryNumber string, instanceGroups []SearchEngineRestoreInstanceGroup, instanceNamePrefix string, name string, ) *SearchEngineClusterRestoreRequest`

NewSearchEngineClusterRestoreRequest instantiates a new SearchEngineClusterRestoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEngineClusterRestoreRequestWithDefaults

`func NewSearchEngineClusterRestoreRequestWithDefaults() *SearchEngineClusterRestoreRequest`

NewSearchEngineClusterRestoreRequestWithDefaults instantiates a new SearchEngineClusterRestoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowableIpAddresses

`func (o *SearchEngineClusterRestoreRequest) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *SearchEngineClusterRestoreRequest) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *SearchEngineClusterRestoreRequest) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *SearchEngineClusterRestoreRequest) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetBackupHistoryNumber

`func (o *SearchEngineClusterRestoreRequest) GetBackupHistoryNumber() string`

GetBackupHistoryNumber returns the BackupHistoryNumber field if non-nil, zero value otherwise.

### GetBackupHistoryNumberOk

`func (o *SearchEngineClusterRestoreRequest) GetBackupHistoryNumberOk() (*string, bool)`

GetBackupHistoryNumberOk returns a tuple with the BackupHistoryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupHistoryNumber

`func (o *SearchEngineClusterRestoreRequest) SetBackupHistoryNumber(v string)`

SetBackupHistoryNumber sets BackupHistoryNumber field to given value.


### GetInstanceGroups

`func (o *SearchEngineClusterRestoreRequest) GetInstanceGroups() []SearchEngineRestoreInstanceGroup`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *SearchEngineClusterRestoreRequest) GetInstanceGroupsOk() (*[]SearchEngineRestoreInstanceGroup, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *SearchEngineClusterRestoreRequest) SetInstanceGroups(v []SearchEngineRestoreInstanceGroup)`

SetInstanceGroups sets InstanceGroups field to given value.


### GetInstanceNamePrefix

`func (o *SearchEngineClusterRestoreRequest) GetInstanceNamePrefix() string`

GetInstanceNamePrefix returns the InstanceNamePrefix field if non-nil, zero value otherwise.

### GetInstanceNamePrefixOk

`func (o *SearchEngineClusterRestoreRequest) GetInstanceNamePrefixOk() (*string, bool)`

GetInstanceNamePrefixOk returns a tuple with the InstanceNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceNamePrefix

`func (o *SearchEngineClusterRestoreRequest) SetInstanceNamePrefix(v string)`

SetInstanceNamePrefix sets InstanceNamePrefix field to given value.


### GetLicense

`func (o *SearchEngineClusterRestoreRequest) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *SearchEngineClusterRestoreRequest) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *SearchEngineClusterRestoreRequest) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *SearchEngineClusterRestoreRequest) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### SetLicenseNil

`func (o *SearchEngineClusterRestoreRequest) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *SearchEngineClusterRestoreRequest) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetMaintenanceOption

`func (o *SearchEngineClusterRestoreRequest) GetMaintenanceOption() MaintenanceOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *SearchEngineClusterRestoreRequest) GetMaintenanceOptionOk() (*MaintenanceOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *SearchEngineClusterRestoreRequest) SetMaintenanceOption(v MaintenanceOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *SearchEngineClusterRestoreRequest) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *SearchEngineClusterRestoreRequest) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *SearchEngineClusterRestoreRequest) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetName

`func (o *SearchEngineClusterRestoreRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchEngineClusterRestoreRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchEngineClusterRestoreRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *SearchEngineClusterRestoreRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SearchEngineClusterRestoreRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SearchEngineClusterRestoreRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SearchEngineClusterRestoreRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SearchEngineClusterRestoreRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SearchEngineClusterRestoreRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


