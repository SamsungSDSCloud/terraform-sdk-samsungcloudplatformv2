# Amount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Krw** | Pointer to **string** | KRW amount | [optional] [default to "0"]
**Usd** | Pointer to **string** | USD amount | [optional] [default to "0"]

## Methods

### NewAmount

`func NewAmount() *Amount`

NewAmount instantiates a new Amount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmountWithDefaults

`func NewAmountWithDefaults() *Amount`

NewAmountWithDefaults instantiates a new Amount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKrw

`func (o *Amount) GetKrw() string`

GetKrw returns the Krw field if non-nil, zero value otherwise.

### GetKrwOk

`func (o *Amount) GetKrwOk() (*string, bool)`

GetKrwOk returns a tuple with the Krw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrw

`func (o *Amount) SetKrw(v string)`

SetKrw sets Krw field to given value.

### HasKrw

`func (o *Amount) HasKrw() bool`

HasKrw returns a boolean if a field has been set.

### GetUsd

`func (o *Amount) GetUsd() string`

GetUsd returns the Usd field if non-nil, zero value otherwise.

### GetUsdOk

`func (o *Amount) GetUsdOk() (*string, bool)`

GetUsdOk returns a tuple with the Usd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsd

`func (o *Amount) SetUsd(v string)`

SetUsd sets Usd field to given value.

### HasUsd

`func (o *Amount) HasUsd() bool`

HasUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


