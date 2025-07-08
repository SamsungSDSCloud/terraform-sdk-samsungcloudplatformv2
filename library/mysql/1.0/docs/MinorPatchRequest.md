# MinorPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupBeforeUpgrade** | **bool** | backup_before_upgrade | 
**SoftwareVersion** | **string** | software_version | 

## Methods

### NewMinorPatchRequest

`func NewMinorPatchRequest(backupBeforeUpgrade bool, softwareVersion string, ) *MinorPatchRequest`

NewMinorPatchRequest instantiates a new MinorPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinorPatchRequestWithDefaults

`func NewMinorPatchRequestWithDefaults() *MinorPatchRequest`

NewMinorPatchRequestWithDefaults instantiates a new MinorPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupBeforeUpgrade

`func (o *MinorPatchRequest) GetBackupBeforeUpgrade() bool`

GetBackupBeforeUpgrade returns the BackupBeforeUpgrade field if non-nil, zero value otherwise.

### GetBackupBeforeUpgradeOk

`func (o *MinorPatchRequest) GetBackupBeforeUpgradeOk() (*bool, bool)`

GetBackupBeforeUpgradeOk returns a tuple with the BackupBeforeUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupBeforeUpgrade

`func (o *MinorPatchRequest) SetBackupBeforeUpgrade(v bool)`

SetBackupBeforeUpgrade sets BackupBeforeUpgrade field to given value.


### GetSoftwareVersion

`func (o *MinorPatchRequest) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *MinorPatchRequest) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *MinorPatchRequest) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


