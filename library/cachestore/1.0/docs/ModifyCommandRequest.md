# ModifyCommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | [**[]ModifyCommandItem**](ModifyCommandItem.md) | Apply command information | 

## Methods

### NewModifyCommandRequest

`func NewModifyCommandRequest(commands []ModifyCommandItem, ) *ModifyCommandRequest`

NewModifyCommandRequest instantiates a new ModifyCommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyCommandRequestWithDefaults

`func NewModifyCommandRequestWithDefaults() *ModifyCommandRequest`

NewModifyCommandRequestWithDefaults instantiates a new ModifyCommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *ModifyCommandRequest) GetCommands() []ModifyCommandItem`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *ModifyCommandRequest) GetCommandsOk() (*[]ModifyCommandItem, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *ModifyCommandRequest) SetCommands(v []ModifyCommandItem)`

SetCommands sets Commands field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


