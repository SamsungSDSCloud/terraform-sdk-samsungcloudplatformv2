# Statement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **[]string** |  | [optional] 
**Condition** | Pointer to [**map[string]map[string][]interface{}**](map.md) |  | [optional] 
**Effect** | **string** | Effect | 
**NotAction** | Pointer to **[]string** |  | [optional] 
**Principal** | Pointer to [**NullablePrincipal**](Principal.md) |  | [optional] 
**Resource** | Pointer to **[]string** | Resource | [optional] 
**Sid** | Pointer to **string** | Statement Id | [optional] [default to "statement1"]

## Methods

### NewStatement

`func NewStatement(effect string, ) *Statement`

NewStatement instantiates a new Statement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementWithDefaults

`func NewStatementWithDefaults() *Statement`

NewStatementWithDefaults instantiates a new Statement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Statement) GetAction() []string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Statement) GetActionOk() (*[]string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Statement) SetAction(v []string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Statement) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *Statement) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *Statement) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetCondition

`func (o *Statement) GetCondition() map[string]map[string][]interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *Statement) GetConditionOk() (*map[string]map[string][]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *Statement) SetCondition(v map[string]map[string][]interface{})`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *Statement) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *Statement) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *Statement) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetEffect

`func (o *Statement) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *Statement) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *Statement) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetNotAction

`func (o *Statement) GetNotAction() []string`

GetNotAction returns the NotAction field if non-nil, zero value otherwise.

### GetNotActionOk

`func (o *Statement) GetNotActionOk() (*[]string, bool)`

GetNotActionOk returns a tuple with the NotAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAction

`func (o *Statement) SetNotAction(v []string)`

SetNotAction sets NotAction field to given value.

### HasNotAction

`func (o *Statement) HasNotAction() bool`

HasNotAction returns a boolean if a field has been set.

### SetNotActionNil

`func (o *Statement) SetNotActionNil(b bool)`

 SetNotActionNil sets the value for NotAction to be an explicit nil

### UnsetNotAction
`func (o *Statement) UnsetNotAction()`

UnsetNotAction ensures that no value is present for NotAction, not even an explicit nil
### GetPrincipal

`func (o *Statement) GetPrincipal() Principal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *Statement) GetPrincipalOk() (*Principal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *Statement) SetPrincipal(v Principal)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *Statement) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### SetPrincipalNil

`func (o *Statement) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *Statement) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil
### GetResource

`func (o *Statement) GetResource() []string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Statement) GetResourceOk() (*[]string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Statement) SetResource(v []string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Statement) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSid

`func (o *Statement) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *Statement) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *Statement) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *Statement) HasSid() bool`

HasSid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


