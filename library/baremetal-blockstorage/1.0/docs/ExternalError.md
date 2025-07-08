# ExternalError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code | [optional] 
**ObjectType** | Pointer to [**BlockStorageAttachmentObjectType**](BlockStorageAttachmentObjectType.md) | Server type | [optional] 
**Status** | Pointer to **int32** | Status | [optional] 

## Methods

### NewExternalError

`func NewExternalError() *ExternalError`

NewExternalError instantiates a new ExternalError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalErrorWithDefaults

`func NewExternalErrorWithDefaults() *ExternalError`

NewExternalErrorWithDefaults instantiates a new ExternalError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ExternalError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExternalError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExternalError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ExternalError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetObjectType

`func (o *ExternalError) GetObjectType() BlockStorageAttachmentObjectType`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ExternalError) GetObjectTypeOk() (*BlockStorageAttachmentObjectType, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ExternalError) SetObjectType(v BlockStorageAttachmentObjectType)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ExternalError) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetStatus

`func (o *ExternalError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExternalError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


