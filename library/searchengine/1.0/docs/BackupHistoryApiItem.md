# BackupHistoryApiItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupEndAt** | **string** | Backup end time | 
**BackupHistoryNumber** | **string** | Backup id | 
**BackupSizeGb** | **float32** | Backup amount(GB) | 
**BackupStartAt** | **string** | Backup start time | 
**BackupState** | **string** | Backup state | 
**BackupStateDetail** | Pointer to **NullableString** |  | [optional] 
**SoftwareVersion** | **string** | Software version | 

## Methods

### NewBackupHistoryApiItem

`func NewBackupHistoryApiItem(backupEndAt string, backupHistoryNumber string, backupSizeGb float32, backupStartAt string, backupState string, softwareVersion string, ) *BackupHistoryApiItem`

NewBackupHistoryApiItem instantiates a new BackupHistoryApiItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupHistoryApiItemWithDefaults

`func NewBackupHistoryApiItemWithDefaults() *BackupHistoryApiItem`

NewBackupHistoryApiItemWithDefaults instantiates a new BackupHistoryApiItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupEndAt

`func (o *BackupHistoryApiItem) GetBackupEndAt() string`

GetBackupEndAt returns the BackupEndAt field if non-nil, zero value otherwise.

### GetBackupEndAtOk

`func (o *BackupHistoryApiItem) GetBackupEndAtOk() (*string, bool)`

GetBackupEndAtOk returns a tuple with the BackupEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEndAt

`func (o *BackupHistoryApiItem) SetBackupEndAt(v string)`

SetBackupEndAt sets BackupEndAt field to given value.


### GetBackupHistoryNumber

`func (o *BackupHistoryApiItem) GetBackupHistoryNumber() string`

GetBackupHistoryNumber returns the BackupHistoryNumber field if non-nil, zero value otherwise.

### GetBackupHistoryNumberOk

`func (o *BackupHistoryApiItem) GetBackupHistoryNumberOk() (*string, bool)`

GetBackupHistoryNumberOk returns a tuple with the BackupHistoryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupHistoryNumber

`func (o *BackupHistoryApiItem) SetBackupHistoryNumber(v string)`

SetBackupHistoryNumber sets BackupHistoryNumber field to given value.


### GetBackupSizeGb

`func (o *BackupHistoryApiItem) GetBackupSizeGb() float32`

GetBackupSizeGb returns the BackupSizeGb field if non-nil, zero value otherwise.

### GetBackupSizeGbOk

`func (o *BackupHistoryApiItem) GetBackupSizeGbOk() (*float32, bool)`

GetBackupSizeGbOk returns a tuple with the BackupSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSizeGb

`func (o *BackupHistoryApiItem) SetBackupSizeGb(v float32)`

SetBackupSizeGb sets BackupSizeGb field to given value.


### GetBackupStartAt

`func (o *BackupHistoryApiItem) GetBackupStartAt() string`

GetBackupStartAt returns the BackupStartAt field if non-nil, zero value otherwise.

### GetBackupStartAtOk

`func (o *BackupHistoryApiItem) GetBackupStartAtOk() (*string, bool)`

GetBackupStartAtOk returns a tuple with the BackupStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStartAt

`func (o *BackupHistoryApiItem) SetBackupStartAt(v string)`

SetBackupStartAt sets BackupStartAt field to given value.


### GetBackupState

`func (o *BackupHistoryApiItem) GetBackupState() string`

GetBackupState returns the BackupState field if non-nil, zero value otherwise.

### GetBackupStateOk

`func (o *BackupHistoryApiItem) GetBackupStateOk() (*string, bool)`

GetBackupStateOk returns a tuple with the BackupState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupState

`func (o *BackupHistoryApiItem) SetBackupState(v string)`

SetBackupState sets BackupState field to given value.


### GetBackupStateDetail

`func (o *BackupHistoryApiItem) GetBackupStateDetail() string`

GetBackupStateDetail returns the BackupStateDetail field if non-nil, zero value otherwise.

### GetBackupStateDetailOk

`func (o *BackupHistoryApiItem) GetBackupStateDetailOk() (*string, bool)`

GetBackupStateDetailOk returns a tuple with the BackupStateDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStateDetail

`func (o *BackupHistoryApiItem) SetBackupStateDetail(v string)`

SetBackupStateDetail sets BackupStateDetail field to given value.

### HasBackupStateDetail

`func (o *BackupHistoryApiItem) HasBackupStateDetail() bool`

HasBackupStateDetail returns a boolean if a field has been set.

### SetBackupStateDetailNil

`func (o *BackupHistoryApiItem) SetBackupStateDetailNil(b bool)`

 SetBackupStateDetailNil sets the value for BackupStateDetail to be an explicit nil

### UnsetBackupStateDetail
`func (o *BackupHistoryApiItem) UnsetBackupStateDetail()`

UnsetBackupStateDetail ensures that no value is present for BackupStateDetail, not even an explicit nil
### GetSoftwareVersion

`func (o *BackupHistoryApiItem) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *BackupHistoryApiItem) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *BackupHistoryApiItem) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


