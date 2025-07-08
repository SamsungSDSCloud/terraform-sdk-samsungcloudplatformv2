# VerticaClusterRestoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowableIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**BackupHistoryNumber** | **string** | Backup id | 
**InstanceGroups** | [**[]VerticaRestoreInstanceGroup**](VerticaRestoreInstanceGroup.md) | Instance groups list | 
**InstanceNamePrefix** | **string** | Instance name prefix | 
**License** | Pointer to **string** | license | [optional] 
**MaintenanceOption** | Pointer to [**NullableMaintenanceOption**](MaintenanceOption.md) |  | [optional] 
**Name** | **string** | Cluster name | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewVerticaClusterRestoreRequest

`func NewVerticaClusterRestoreRequest(backupHistoryNumber string, instanceGroups []VerticaRestoreInstanceGroup, instanceNamePrefix string, name string, ) *VerticaClusterRestoreRequest`

NewVerticaClusterRestoreRequest instantiates a new VerticaClusterRestoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerticaClusterRestoreRequestWithDefaults

`func NewVerticaClusterRestoreRequestWithDefaults() *VerticaClusterRestoreRequest`

NewVerticaClusterRestoreRequestWithDefaults instantiates a new VerticaClusterRestoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowableIpAddresses

`func (o *VerticaClusterRestoreRequest) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *VerticaClusterRestoreRequest) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *VerticaClusterRestoreRequest) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *VerticaClusterRestoreRequest) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetBackupHistoryNumber

`func (o *VerticaClusterRestoreRequest) GetBackupHistoryNumber() string`

GetBackupHistoryNumber returns the BackupHistoryNumber field if non-nil, zero value otherwise.

### GetBackupHistoryNumberOk

`func (o *VerticaClusterRestoreRequest) GetBackupHistoryNumberOk() (*string, bool)`

GetBackupHistoryNumberOk returns a tuple with the BackupHistoryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupHistoryNumber

`func (o *VerticaClusterRestoreRequest) SetBackupHistoryNumber(v string)`

SetBackupHistoryNumber sets BackupHistoryNumber field to given value.


### GetInstanceGroups

`func (o *VerticaClusterRestoreRequest) GetInstanceGroups() []VerticaRestoreInstanceGroup`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *VerticaClusterRestoreRequest) GetInstanceGroupsOk() (*[]VerticaRestoreInstanceGroup, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *VerticaClusterRestoreRequest) SetInstanceGroups(v []VerticaRestoreInstanceGroup)`

SetInstanceGroups sets InstanceGroups field to given value.


### GetInstanceNamePrefix

`func (o *VerticaClusterRestoreRequest) GetInstanceNamePrefix() string`

GetInstanceNamePrefix returns the InstanceNamePrefix field if non-nil, zero value otherwise.

### GetInstanceNamePrefixOk

`func (o *VerticaClusterRestoreRequest) GetInstanceNamePrefixOk() (*string, bool)`

GetInstanceNamePrefixOk returns a tuple with the InstanceNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceNamePrefix

`func (o *VerticaClusterRestoreRequest) SetInstanceNamePrefix(v string)`

SetInstanceNamePrefix sets InstanceNamePrefix field to given value.


### GetLicense

`func (o *VerticaClusterRestoreRequest) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *VerticaClusterRestoreRequest) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *VerticaClusterRestoreRequest) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *VerticaClusterRestoreRequest) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMaintenanceOption

`func (o *VerticaClusterRestoreRequest) GetMaintenanceOption() MaintenanceOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *VerticaClusterRestoreRequest) GetMaintenanceOptionOk() (*MaintenanceOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *VerticaClusterRestoreRequest) SetMaintenanceOption(v MaintenanceOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *VerticaClusterRestoreRequest) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *VerticaClusterRestoreRequest) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *VerticaClusterRestoreRequest) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetName

`func (o *VerticaClusterRestoreRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VerticaClusterRestoreRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VerticaClusterRestoreRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *VerticaClusterRestoreRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VerticaClusterRestoreRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VerticaClusterRestoreRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VerticaClusterRestoreRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *VerticaClusterRestoreRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *VerticaClusterRestoreRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


