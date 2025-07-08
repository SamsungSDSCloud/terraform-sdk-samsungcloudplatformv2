# BackupListRegionRelationshipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]BackupRegionRelationshipResponse**](BackupRegionRelationshipResponse.md) |  | 
**Count** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBackupListRegionRelationshipResponse

`func NewBackupListRegionRelationshipResponse(contents []BackupRegionRelationshipResponse, ) *BackupListRegionRelationshipResponse`

NewBackupListRegionRelationshipResponse instantiates a new BackupListRegionRelationshipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupListRegionRelationshipResponseWithDefaults

`func NewBackupListRegionRelationshipResponseWithDefaults() *BackupListRegionRelationshipResponse`

NewBackupListRegionRelationshipResponseWithDefaults instantiates a new BackupListRegionRelationshipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *BackupListRegionRelationshipResponse) GetContents() []BackupRegionRelationshipResponse`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *BackupListRegionRelationshipResponse) GetContentsOk() (*[]BackupRegionRelationshipResponse, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *BackupListRegionRelationshipResponse) SetContents(v []BackupRegionRelationshipResponse)`

SetContents sets Contents field to given value.


### SetContentsNil

`func (o *BackupListRegionRelationshipResponse) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *BackupListRegionRelationshipResponse) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetCount

`func (o *BackupListRegionRelationshipResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BackupListRegionRelationshipResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BackupListRegionRelationshipResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BackupListRegionRelationshipResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *BackupListRegionRelationshipResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *BackupListRegionRelationshipResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


