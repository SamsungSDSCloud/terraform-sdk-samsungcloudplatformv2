# NodepoolLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Node Pool Label Key | 
**Value** | Pointer to **string** | Node Pool Label Value | [optional] [default to ""]

## Methods

### NewNodepoolLabel

`func NewNodepoolLabel(key string, ) *NodepoolLabel`

NewNodepoolLabel instantiates a new NodepoolLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodepoolLabelWithDefaults

`func NewNodepoolLabelWithDefaults() *NodepoolLabel`

NewNodepoolLabelWithDefaults instantiates a new NodepoolLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *NodepoolLabel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NodepoolLabel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NodepoolLabel) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *NodepoolLabel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NodepoolLabel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NodepoolLabel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NodepoolLabel) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


