# ValidationErrorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ctx** | Pointer to **map[string]interface{}** |  | [optional] 
**Loc** | Pointer to **[]string** |  | [optional] 
**Msg** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewValidationErrorModel

`func NewValidationErrorModel() *ValidationErrorModel`

NewValidationErrorModel instantiates a new ValidationErrorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorModelWithDefaults

`func NewValidationErrorModelWithDefaults() *ValidationErrorModel`

NewValidationErrorModelWithDefaults instantiates a new ValidationErrorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtx

`func (o *ValidationErrorModel) GetCtx() map[string]interface{}`

GetCtx returns the Ctx field if non-nil, zero value otherwise.

### GetCtxOk

`func (o *ValidationErrorModel) GetCtxOk() (*map[string]interface{}, bool)`

GetCtxOk returns a tuple with the Ctx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtx

`func (o *ValidationErrorModel) SetCtx(v map[string]interface{})`

SetCtx sets Ctx field to given value.

### HasCtx

`func (o *ValidationErrorModel) HasCtx() bool`

HasCtx returns a boolean if a field has been set.

### SetCtxNil

`func (o *ValidationErrorModel) SetCtxNil(b bool)`

 SetCtxNil sets the value for Ctx to be an explicit nil

### UnsetCtx
`func (o *ValidationErrorModel) UnsetCtx()`

UnsetCtx ensures that no value is present for Ctx, not even an explicit nil
### GetLoc

`func (o *ValidationErrorModel) GetLoc() []string`

GetLoc returns the Loc field if non-nil, zero value otherwise.

### GetLocOk

`func (o *ValidationErrorModel) GetLocOk() (*[]string, bool)`

GetLocOk returns a tuple with the Loc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoc

`func (o *ValidationErrorModel) SetLoc(v []string)`

SetLoc sets Loc field to given value.

### HasLoc

`func (o *ValidationErrorModel) HasLoc() bool`

HasLoc returns a boolean if a field has been set.

### SetLocNil

`func (o *ValidationErrorModel) SetLocNil(b bool)`

 SetLocNil sets the value for Loc to be an explicit nil

### UnsetLoc
`func (o *ValidationErrorModel) UnsetLoc()`

UnsetLoc ensures that no value is present for Loc, not even an explicit nil
### GetMsg

`func (o *ValidationErrorModel) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ValidationErrorModel) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ValidationErrorModel) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ValidationErrorModel) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *ValidationErrorModel) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *ValidationErrorModel) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetType

`func (o *ValidationErrorModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidationErrorModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidationErrorModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ValidationErrorModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ValidationErrorModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ValidationErrorModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


