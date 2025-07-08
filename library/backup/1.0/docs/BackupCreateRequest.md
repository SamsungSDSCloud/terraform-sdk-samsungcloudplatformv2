# BackupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptEnabled** | Pointer to **bool** | Whether to use Encryption | [optional] [default to false]
**FilesystemPaths** | Pointer to **[]string** |  | [optional] 
**Name** | **string** | Backup name | 
**PolicyCategory** | [**BackupPolicyCategory**](BackupPolicyCategory.md) | Backup policy category | 
**PolicyType** | [**BackupPolicyType**](BackupPolicyType.md) | Backup policy type | 
**Region** | Pointer to **NullableString** |  | [optional] 
**RetentionPeriod** | Pointer to [**BackupRetentionPeriod**](BackupRetentionPeriod.md) | Backup retention period | [optional] [default to BACKUPRETENTIONPERIOD_WEEK_2]
**Schedules** | [**[]BackupScheduleCreateRequest**](BackupScheduleCreateRequest.md) | Schedules | 
**ServerCategory** | [**ServerCategory**](ServerCategory.md) | Backup server category | 
**ServerGuid** | Pointer to **NullableString** |  | [optional] 
**ServerUuid** | **string** | Backup server UUID | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewBackupCreateRequest

`func NewBackupCreateRequest(name string, policyCategory BackupPolicyCategory, policyType BackupPolicyType, schedules []BackupScheduleCreateRequest, serverCategory ServerCategory, serverUuid string, ) *BackupCreateRequest`

NewBackupCreateRequest instantiates a new BackupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupCreateRequestWithDefaults

`func NewBackupCreateRequestWithDefaults() *BackupCreateRequest`

NewBackupCreateRequestWithDefaults instantiates a new BackupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptEnabled

`func (o *BackupCreateRequest) GetEncryptEnabled() bool`

GetEncryptEnabled returns the EncryptEnabled field if non-nil, zero value otherwise.

### GetEncryptEnabledOk

`func (o *BackupCreateRequest) GetEncryptEnabledOk() (*bool, bool)`

GetEncryptEnabledOk returns a tuple with the EncryptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptEnabled

`func (o *BackupCreateRequest) SetEncryptEnabled(v bool)`

SetEncryptEnabled sets EncryptEnabled field to given value.

### HasEncryptEnabled

`func (o *BackupCreateRequest) HasEncryptEnabled() bool`

HasEncryptEnabled returns a boolean if a field has been set.

### GetFilesystemPaths

`func (o *BackupCreateRequest) GetFilesystemPaths() []string`

GetFilesystemPaths returns the FilesystemPaths field if non-nil, zero value otherwise.

### GetFilesystemPathsOk

`func (o *BackupCreateRequest) GetFilesystemPathsOk() (*[]string, bool)`

GetFilesystemPathsOk returns a tuple with the FilesystemPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemPaths

`func (o *BackupCreateRequest) SetFilesystemPaths(v []string)`

SetFilesystemPaths sets FilesystemPaths field to given value.

### HasFilesystemPaths

`func (o *BackupCreateRequest) HasFilesystemPaths() bool`

HasFilesystemPaths returns a boolean if a field has been set.

### SetFilesystemPathsNil

`func (o *BackupCreateRequest) SetFilesystemPathsNil(b bool)`

 SetFilesystemPathsNil sets the value for FilesystemPaths to be an explicit nil

### UnsetFilesystemPaths
`func (o *BackupCreateRequest) UnsetFilesystemPaths()`

UnsetFilesystemPaths ensures that no value is present for FilesystemPaths, not even an explicit nil
### GetName

`func (o *BackupCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPolicyCategory

`func (o *BackupCreateRequest) GetPolicyCategory() BackupPolicyCategory`

GetPolicyCategory returns the PolicyCategory field if non-nil, zero value otherwise.

### GetPolicyCategoryOk

`func (o *BackupCreateRequest) GetPolicyCategoryOk() (*BackupPolicyCategory, bool)`

GetPolicyCategoryOk returns a tuple with the PolicyCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCategory

`func (o *BackupCreateRequest) SetPolicyCategory(v BackupPolicyCategory)`

SetPolicyCategory sets PolicyCategory field to given value.


### GetPolicyType

`func (o *BackupCreateRequest) GetPolicyType() BackupPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *BackupCreateRequest) GetPolicyTypeOk() (*BackupPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *BackupCreateRequest) SetPolicyType(v BackupPolicyType)`

SetPolicyType sets PolicyType field to given value.


### GetRegion

`func (o *BackupCreateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BackupCreateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BackupCreateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BackupCreateRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *BackupCreateRequest) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *BackupCreateRequest) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetRetentionPeriod

`func (o *BackupCreateRequest) GetRetentionPeriod() BackupRetentionPeriod`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *BackupCreateRequest) GetRetentionPeriodOk() (*BackupRetentionPeriod, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *BackupCreateRequest) SetRetentionPeriod(v BackupRetentionPeriod)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *BackupCreateRequest) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetSchedules

`func (o *BackupCreateRequest) GetSchedules() []BackupScheduleCreateRequest`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *BackupCreateRequest) GetSchedulesOk() (*[]BackupScheduleCreateRequest, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *BackupCreateRequest) SetSchedules(v []BackupScheduleCreateRequest)`

SetSchedules sets Schedules field to given value.


### GetServerCategory

`func (o *BackupCreateRequest) GetServerCategory() ServerCategory`

GetServerCategory returns the ServerCategory field if non-nil, zero value otherwise.

### GetServerCategoryOk

`func (o *BackupCreateRequest) GetServerCategoryOk() (*ServerCategory, bool)`

GetServerCategoryOk returns a tuple with the ServerCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCategory

`func (o *BackupCreateRequest) SetServerCategory(v ServerCategory)`

SetServerCategory sets ServerCategory field to given value.


### GetServerGuid

`func (o *BackupCreateRequest) GetServerGuid() string`

GetServerGuid returns the ServerGuid field if non-nil, zero value otherwise.

### GetServerGuidOk

`func (o *BackupCreateRequest) GetServerGuidOk() (*string, bool)`

GetServerGuidOk returns a tuple with the ServerGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGuid

`func (o *BackupCreateRequest) SetServerGuid(v string)`

SetServerGuid sets ServerGuid field to given value.

### HasServerGuid

`func (o *BackupCreateRequest) HasServerGuid() bool`

HasServerGuid returns a boolean if a field has been set.

### SetServerGuidNil

`func (o *BackupCreateRequest) SetServerGuidNil(b bool)`

 SetServerGuidNil sets the value for ServerGuid to be an explicit nil

### UnsetServerGuid
`func (o *BackupCreateRequest) UnsetServerGuid()`

UnsetServerGuid ensures that no value is present for ServerGuid, not even an explicit nil
### GetServerUuid

`func (o *BackupCreateRequest) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *BackupCreateRequest) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *BackupCreateRequest) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.


### GetTags

`func (o *BackupCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BackupCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BackupCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BackupCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *BackupCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *BackupCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


