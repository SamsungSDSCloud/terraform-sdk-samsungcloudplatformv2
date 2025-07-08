# BackupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]BackupResponse**](BackupResponse.md) |  | 
**Count** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBackupListResponse

`func NewBackupListResponse(contents []BackupResponse, ) *BackupListResponse`

NewBackupListResponse instantiates a new BackupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupListResponseWithDefaults

`func NewBackupListResponseWithDefaults() *BackupListResponse`

NewBackupListResponseWithDefaults instantiates a new BackupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *BackupListResponse) GetContents() []BackupResponse`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *BackupListResponse) GetContentsOk() (*[]BackupResponse, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *BackupListResponse) SetContents(v []BackupResponse)`

SetContents sets Contents field to given value.


### SetContentsNil

`func (o *BackupListResponse) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *BackupListResponse) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetCount

`func (o *BackupListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BackupListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BackupListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BackupListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *BackupListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *BackupListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


