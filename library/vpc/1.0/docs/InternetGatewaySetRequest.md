# InternetGatewaySetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Loggable** | Pointer to **bool** | NAT Loggable | [optional] 

## Methods

### NewInternetGatewaySetRequest

`func NewInternetGatewaySetRequest() *InternetGatewaySetRequest`

NewInternetGatewaySetRequest instantiates a new InternetGatewaySetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetGatewaySetRequestWithDefaults

`func NewInternetGatewaySetRequestWithDefaults() *InternetGatewaySetRequest`

NewInternetGatewaySetRequestWithDefaults instantiates a new InternetGatewaySetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *InternetGatewaySetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternetGatewaySetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternetGatewaySetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternetGatewaySetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InternetGatewaySetRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InternetGatewaySetRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLoggable

`func (o *InternetGatewaySetRequest) GetLoggable() bool`

GetLoggable returns the Loggable field if non-nil, zero value otherwise.

### GetLoggableOk

`func (o *InternetGatewaySetRequest) GetLoggableOk() (*bool, bool)`

GetLoggableOk returns a tuple with the Loggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggable

`func (o *InternetGatewaySetRequest) SetLoggable(v bool)`

SetLoggable sets Loggable field to given value.

### HasLoggable

`func (o *InternetGatewaySetRequest) HasLoggable() bool`

HasLoggable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


