# ServerRebootRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RebootType** | Pointer to [**NullableServerRebootType**](ServerRebootType.md) |  | [optional] 

## Methods

### NewServerRebootRequestBody

`func NewServerRebootRequestBody() *ServerRebootRequestBody`

NewServerRebootRequestBody instantiates a new ServerRebootRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerRebootRequestBodyWithDefaults

`func NewServerRebootRequestBodyWithDefaults() *ServerRebootRequestBody`

NewServerRebootRequestBodyWithDefaults instantiates a new ServerRebootRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRebootType

`func (o *ServerRebootRequestBody) GetRebootType() ServerRebootType`

GetRebootType returns the RebootType field if non-nil, zero value otherwise.

### GetRebootTypeOk

`func (o *ServerRebootRequestBody) GetRebootTypeOk() (*ServerRebootType, bool)`

GetRebootTypeOk returns a tuple with the RebootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootType

`func (o *ServerRebootRequestBody) SetRebootType(v ServerRebootType)`

SetRebootType sets RebootType field to given value.

### HasRebootType

`func (o *ServerRebootRequestBody) HasRebootType() bool`

HasRebootType returns a boolean if a field has been set.

### SetRebootTypeNil

`func (o *ServerRebootRequestBody) SetRebootTypeNil(b bool)`

 SetRebootTypeNil sets the value for RebootType to be an explicit nil

### UnsetRebootType
`func (o *ServerRebootRequestBody) UnsetRebootType()`

UnsetRebootType ensures that no value is present for RebootType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


