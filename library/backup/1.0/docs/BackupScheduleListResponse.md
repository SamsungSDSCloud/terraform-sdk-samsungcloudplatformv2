# BackupScheduleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]BackupScheduleResponse**](BackupScheduleResponse.md) |  | 
**Count** | **int32** | Count | 

## Methods

### NewBackupScheduleListResponse

`func NewBackupScheduleListResponse(contents []BackupScheduleResponse, count int32, ) *BackupScheduleListResponse`

NewBackupScheduleListResponse instantiates a new BackupScheduleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleListResponseWithDefaults

`func NewBackupScheduleListResponseWithDefaults() *BackupScheduleListResponse`

NewBackupScheduleListResponseWithDefaults instantiates a new BackupScheduleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *BackupScheduleListResponse) GetContents() []BackupScheduleResponse`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *BackupScheduleListResponse) GetContentsOk() (*[]BackupScheduleResponse, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *BackupScheduleListResponse) SetContents(v []BackupScheduleResponse)`

SetContents sets Contents field to given value.


### SetContentsNil

`func (o *BackupScheduleListResponse) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *BackupScheduleListResponse) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetCount

`func (o *BackupScheduleListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BackupScheduleListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BackupScheduleListResponse) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


