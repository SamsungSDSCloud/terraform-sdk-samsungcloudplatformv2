# BackupDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupAgentId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**EncryptEnabled** | Pointer to **NullableBool** |  | [optional] 
**FilesystemPaths** | Pointer to **[]string** |  | [optional] 
**Id** | **string** | ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Backup name | 
**PolicyCategory** | **string** | Backup policy category | 
**PolicyType** | **string** | Backup policy type | 
**Region** | Pointer to **NullableString** |  | [optional] 
**RetentionPeriod** | **string** | Backup retention period | 
**RoleType** | **string** | Backup role type | 
**ServerCategory** | **string** | Backup server category | 
**ServerName** | **string** | Backup server name | 
**ServerOsType** | Pointer to **NullableString** |  | [optional] 
**ServerRegion** | **NullableString** |  | 
**ServerUuid** | **string** | Backup server UUID | 
**State** | **string** | Backup state | 

## Methods

### NewBackupDetailResponse

`func NewBackupDetailResponse(createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, name string, policyCategory string, policyType string, retentionPeriod string, roleType string, serverCategory string, serverName string, serverRegion NullableString, serverUuid string, state string, ) *BackupDetailResponse`

NewBackupDetailResponse instantiates a new BackupDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupDetailResponseWithDefaults

`func NewBackupDetailResponseWithDefaults() *BackupDetailResponse`

NewBackupDetailResponseWithDefaults instantiates a new BackupDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupAgentId

`func (o *BackupDetailResponse) GetBackupAgentId() string`

GetBackupAgentId returns the BackupAgentId field if non-nil, zero value otherwise.

### GetBackupAgentIdOk

`func (o *BackupDetailResponse) GetBackupAgentIdOk() (*string, bool)`

GetBackupAgentIdOk returns a tuple with the BackupAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAgentId

`func (o *BackupDetailResponse) SetBackupAgentId(v string)`

SetBackupAgentId sets BackupAgentId field to given value.

### HasBackupAgentId

`func (o *BackupDetailResponse) HasBackupAgentId() bool`

HasBackupAgentId returns a boolean if a field has been set.

### SetBackupAgentIdNil

`func (o *BackupDetailResponse) SetBackupAgentIdNil(b bool)`

 SetBackupAgentIdNil sets the value for BackupAgentId to be an explicit nil

### UnsetBackupAgentId
`func (o *BackupDetailResponse) UnsetBackupAgentId()`

UnsetBackupAgentId ensures that no value is present for BackupAgentId, not even an explicit nil
### GetCreatedAt

`func (o *BackupDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackupDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackupDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *BackupDetailResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BackupDetailResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BackupDetailResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetEncryptEnabled

`func (o *BackupDetailResponse) GetEncryptEnabled() bool`

GetEncryptEnabled returns the EncryptEnabled field if non-nil, zero value otherwise.

### GetEncryptEnabledOk

`func (o *BackupDetailResponse) GetEncryptEnabledOk() (*bool, bool)`

GetEncryptEnabledOk returns a tuple with the EncryptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptEnabled

`func (o *BackupDetailResponse) SetEncryptEnabled(v bool)`

SetEncryptEnabled sets EncryptEnabled field to given value.

### HasEncryptEnabled

`func (o *BackupDetailResponse) HasEncryptEnabled() bool`

HasEncryptEnabled returns a boolean if a field has been set.

### SetEncryptEnabledNil

`func (o *BackupDetailResponse) SetEncryptEnabledNil(b bool)`

 SetEncryptEnabledNil sets the value for EncryptEnabled to be an explicit nil

### UnsetEncryptEnabled
`func (o *BackupDetailResponse) UnsetEncryptEnabled()`

UnsetEncryptEnabled ensures that no value is present for EncryptEnabled, not even an explicit nil
### GetFilesystemPaths

`func (o *BackupDetailResponse) GetFilesystemPaths() []string`

GetFilesystemPaths returns the FilesystemPaths field if non-nil, zero value otherwise.

### GetFilesystemPathsOk

`func (o *BackupDetailResponse) GetFilesystemPathsOk() (*[]string, bool)`

GetFilesystemPathsOk returns a tuple with the FilesystemPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemPaths

`func (o *BackupDetailResponse) SetFilesystemPaths(v []string)`

SetFilesystemPaths sets FilesystemPaths field to given value.

### HasFilesystemPaths

`func (o *BackupDetailResponse) HasFilesystemPaths() bool`

HasFilesystemPaths returns a boolean if a field has been set.

### SetFilesystemPathsNil

`func (o *BackupDetailResponse) SetFilesystemPathsNil(b bool)`

 SetFilesystemPathsNil sets the value for FilesystemPaths to be an explicit nil

### UnsetFilesystemPaths
`func (o *BackupDetailResponse) UnsetFilesystemPaths()`

UnsetFilesystemPaths ensures that no value is present for FilesystemPaths, not even an explicit nil
### GetId

`func (o *BackupDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *BackupDetailResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BackupDetailResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BackupDetailResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *BackupDetailResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BackupDetailResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BackupDetailResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *BackupDetailResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupDetailResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupDetailResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPolicyCategory

`func (o *BackupDetailResponse) GetPolicyCategory() string`

GetPolicyCategory returns the PolicyCategory field if non-nil, zero value otherwise.

### GetPolicyCategoryOk

`func (o *BackupDetailResponse) GetPolicyCategoryOk() (*string, bool)`

GetPolicyCategoryOk returns a tuple with the PolicyCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCategory

`func (o *BackupDetailResponse) SetPolicyCategory(v string)`

SetPolicyCategory sets PolicyCategory field to given value.


### GetPolicyType

`func (o *BackupDetailResponse) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *BackupDetailResponse) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *BackupDetailResponse) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.


### GetRegion

`func (o *BackupDetailResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BackupDetailResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BackupDetailResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BackupDetailResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *BackupDetailResponse) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *BackupDetailResponse) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetRetentionPeriod

`func (o *BackupDetailResponse) GetRetentionPeriod() string`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *BackupDetailResponse) GetRetentionPeriodOk() (*string, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *BackupDetailResponse) SetRetentionPeriod(v string)`

SetRetentionPeriod sets RetentionPeriod field to given value.


### GetRoleType

`func (o *BackupDetailResponse) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *BackupDetailResponse) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *BackupDetailResponse) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.


### GetServerCategory

`func (o *BackupDetailResponse) GetServerCategory() string`

GetServerCategory returns the ServerCategory field if non-nil, zero value otherwise.

### GetServerCategoryOk

`func (o *BackupDetailResponse) GetServerCategoryOk() (*string, bool)`

GetServerCategoryOk returns a tuple with the ServerCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCategory

`func (o *BackupDetailResponse) SetServerCategory(v string)`

SetServerCategory sets ServerCategory field to given value.


### GetServerName

`func (o *BackupDetailResponse) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *BackupDetailResponse) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *BackupDetailResponse) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetServerOsType

`func (o *BackupDetailResponse) GetServerOsType() string`

GetServerOsType returns the ServerOsType field if non-nil, zero value otherwise.

### GetServerOsTypeOk

`func (o *BackupDetailResponse) GetServerOsTypeOk() (*string, bool)`

GetServerOsTypeOk returns a tuple with the ServerOsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOsType

`func (o *BackupDetailResponse) SetServerOsType(v string)`

SetServerOsType sets ServerOsType field to given value.

### HasServerOsType

`func (o *BackupDetailResponse) HasServerOsType() bool`

HasServerOsType returns a boolean if a field has been set.

### SetServerOsTypeNil

`func (o *BackupDetailResponse) SetServerOsTypeNil(b bool)`

 SetServerOsTypeNil sets the value for ServerOsType to be an explicit nil

### UnsetServerOsType
`func (o *BackupDetailResponse) UnsetServerOsType()`

UnsetServerOsType ensures that no value is present for ServerOsType, not even an explicit nil
### GetServerRegion

`func (o *BackupDetailResponse) GetServerRegion() string`

GetServerRegion returns the ServerRegion field if non-nil, zero value otherwise.

### GetServerRegionOk

`func (o *BackupDetailResponse) GetServerRegionOk() (*string, bool)`

GetServerRegionOk returns a tuple with the ServerRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRegion

`func (o *BackupDetailResponse) SetServerRegion(v string)`

SetServerRegion sets ServerRegion field to given value.


### SetServerRegionNil

`func (o *BackupDetailResponse) SetServerRegionNil(b bool)`

 SetServerRegionNil sets the value for ServerRegion to be an explicit nil

### UnsetServerRegion
`func (o *BackupDetailResponse) UnsetServerRegion()`

UnsetServerRegion ensures that no value is present for ServerRegion, not even an explicit nil
### GetServerUuid

`func (o *BackupDetailResponse) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *BackupDetailResponse) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *BackupDetailResponse) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.


### GetState

`func (o *BackupDetailResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackupDetailResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackupDetailResponse) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


