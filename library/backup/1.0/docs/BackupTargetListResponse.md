# BackupTargetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | Pointer to [**[]BackupTargetResponse**](BackupTargetResponse.md) |  | [optional] 
**Count** | **int32** | Count | 

## Methods

### NewBackupTargetListResponse

`func NewBackupTargetListResponse(count int32, ) *BackupTargetListResponse`

NewBackupTargetListResponse instantiates a new BackupTargetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupTargetListResponseWithDefaults

`func NewBackupTargetListResponseWithDefaults() *BackupTargetListResponse`

NewBackupTargetListResponseWithDefaults instantiates a new BackupTargetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *BackupTargetListResponse) GetContents() []BackupTargetResponse`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *BackupTargetListResponse) GetContentsOk() (*[]BackupTargetResponse, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *BackupTargetListResponse) SetContents(v []BackupTargetResponse)`

SetContents sets Contents field to given value.

### HasContents

`func (o *BackupTargetListResponse) HasContents() bool`

HasContents returns a boolean if a field has been set.

### SetContentsNil

`func (o *BackupTargetListResponse) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *BackupTargetListResponse) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetCount

`func (o *BackupTargetListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BackupTargetListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BackupTargetListResponse) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


