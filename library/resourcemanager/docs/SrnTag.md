# SrnTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Srn** | **string** | SRN | 
**Tag** | [**Tag**](Tag.md) |  | 

## Methods

### NewSrnTag

`func NewSrnTag(srn string, tag Tag, ) *SrnTag`

NewSrnTag instantiates a new SrnTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrnTagWithDefaults

`func NewSrnTagWithDefaults() *SrnTag`

NewSrnTagWithDefaults instantiates a new SrnTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrn

`func (o *SrnTag) GetSrn() string`

GetSrn returns the Srn field if non-nil, zero value otherwise.

### GetSrnOk

`func (o *SrnTag) GetSrnOk() (*string, bool)`

GetSrnOk returns a tuple with the Srn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrn

`func (o *SrnTag) SetSrn(v string)`

SetSrn sets Srn field to given value.


### GetTag

`func (o *SrnTag) GetTag() Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *SrnTag) GetTagOk() (*Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *SrnTag) SetTag(v Tag)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


