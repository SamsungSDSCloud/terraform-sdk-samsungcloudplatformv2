# ParameterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Parameter id | 
**NewValue** | **string** | Parameter new value | 
**OldValue** | **string** | Parameter old value | 

## Methods

### NewParameterRequest

`func NewParameterRequest(id string, newValue string, oldValue string, ) *ParameterRequest`

NewParameterRequest instantiates a new ParameterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterRequestWithDefaults

`func NewParameterRequestWithDefaults() *ParameterRequest`

NewParameterRequestWithDefaults instantiates a new ParameterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParameterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterRequest) SetId(v string)`

SetId sets Id field to given value.


### GetNewValue

`func (o *ParameterRequest) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *ParameterRequest) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *ParameterRequest) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.


### GetOldValue

`func (o *ParameterRequest) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *ParameterRequest) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *ParameterRequest) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


