# NodepoolTaint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | Pointer to [**TaintEffectEnum**](TaintEffectEnum.md) | Node Pool Taint Effect | [optional] 
**Key** | **string** | Node Pool Taint Key | 
**Value** | Pointer to **string** | Node Pool Taint Value | [optional] [default to ""]

## Methods

### NewNodepoolTaint

`func NewNodepoolTaint(key string, ) *NodepoolTaint`

NewNodepoolTaint instantiates a new NodepoolTaint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodepoolTaintWithDefaults

`func NewNodepoolTaintWithDefaults() *NodepoolTaint`

NewNodepoolTaintWithDefaults instantiates a new NodepoolTaint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *NodepoolTaint) GetEffect() TaintEffectEnum`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *NodepoolTaint) GetEffectOk() (*TaintEffectEnum, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *NodepoolTaint) SetEffect(v TaintEffectEnum)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *NodepoolTaint) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetKey

`func (o *NodepoolTaint) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NodepoolTaint) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NodepoolTaint) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *NodepoolTaint) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NodepoolTaint) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NodepoolTaint) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NodepoolTaint) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


