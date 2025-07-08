# ResourceUsedAmountDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NonAppliedAmount** | Pointer to [**NullableAmount**](Amount.md) |  | [optional] 
**RequestAt** | Pointer to **NullableString** |  | [optional] 
**ResourceName** | Pointer to **NullableString** |  | [optional] 
**UnitPrice** | Pointer to [**NullableAmount**](Amount.md) |  | [optional] 
**UsedAmount** | Pointer to [**NullableAmount**](Amount.md) |  | [optional] 
**UsedTime** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewResourceUsedAmountDTO

`func NewResourceUsedAmountDTO() *ResourceUsedAmountDTO`

NewResourceUsedAmountDTO instantiates a new ResourceUsedAmountDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUsedAmountDTOWithDefaults

`func NewResourceUsedAmountDTOWithDefaults() *ResourceUsedAmountDTO`

NewResourceUsedAmountDTOWithDefaults instantiates a new ResourceUsedAmountDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonAppliedAmount

`func (o *ResourceUsedAmountDTO) GetNonAppliedAmount() Amount`

GetNonAppliedAmount returns the NonAppliedAmount field if non-nil, zero value otherwise.

### GetNonAppliedAmountOk

`func (o *ResourceUsedAmountDTO) GetNonAppliedAmountOk() (*Amount, bool)`

GetNonAppliedAmountOk returns a tuple with the NonAppliedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonAppliedAmount

`func (o *ResourceUsedAmountDTO) SetNonAppliedAmount(v Amount)`

SetNonAppliedAmount sets NonAppliedAmount field to given value.

### HasNonAppliedAmount

`func (o *ResourceUsedAmountDTO) HasNonAppliedAmount() bool`

HasNonAppliedAmount returns a boolean if a field has been set.

### SetNonAppliedAmountNil

`func (o *ResourceUsedAmountDTO) SetNonAppliedAmountNil(b bool)`

 SetNonAppliedAmountNil sets the value for NonAppliedAmount to be an explicit nil

### UnsetNonAppliedAmount
`func (o *ResourceUsedAmountDTO) UnsetNonAppliedAmount()`

UnsetNonAppliedAmount ensures that no value is present for NonAppliedAmount, not even an explicit nil
### GetRequestAt

`func (o *ResourceUsedAmountDTO) GetRequestAt() string`

GetRequestAt returns the RequestAt field if non-nil, zero value otherwise.

### GetRequestAtOk

`func (o *ResourceUsedAmountDTO) GetRequestAtOk() (*string, bool)`

GetRequestAtOk returns a tuple with the RequestAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAt

`func (o *ResourceUsedAmountDTO) SetRequestAt(v string)`

SetRequestAt sets RequestAt field to given value.

### HasRequestAt

`func (o *ResourceUsedAmountDTO) HasRequestAt() bool`

HasRequestAt returns a boolean if a field has been set.

### SetRequestAtNil

`func (o *ResourceUsedAmountDTO) SetRequestAtNil(b bool)`

 SetRequestAtNil sets the value for RequestAt to be an explicit nil

### UnsetRequestAt
`func (o *ResourceUsedAmountDTO) UnsetRequestAt()`

UnsetRequestAt ensures that no value is present for RequestAt, not even an explicit nil
### GetResourceName

`func (o *ResourceUsedAmountDTO) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceUsedAmountDTO) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceUsedAmountDTO) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ResourceUsedAmountDTO) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *ResourceUsedAmountDTO) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *ResourceUsedAmountDTO) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetUnitPrice

`func (o *ResourceUsedAmountDTO) GetUnitPrice() Amount`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *ResourceUsedAmountDTO) GetUnitPriceOk() (*Amount, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *ResourceUsedAmountDTO) SetUnitPrice(v Amount)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *ResourceUsedAmountDTO) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *ResourceUsedAmountDTO) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *ResourceUsedAmountDTO) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetUsedAmount

`func (o *ResourceUsedAmountDTO) GetUsedAmount() Amount`

GetUsedAmount returns the UsedAmount field if non-nil, zero value otherwise.

### GetUsedAmountOk

`func (o *ResourceUsedAmountDTO) GetUsedAmountOk() (*Amount, bool)`

GetUsedAmountOk returns a tuple with the UsedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAmount

`func (o *ResourceUsedAmountDTO) SetUsedAmount(v Amount)`

SetUsedAmount sets UsedAmount field to given value.

### HasUsedAmount

`func (o *ResourceUsedAmountDTO) HasUsedAmount() bool`

HasUsedAmount returns a boolean if a field has been set.

### SetUsedAmountNil

`func (o *ResourceUsedAmountDTO) SetUsedAmountNil(b bool)`

 SetUsedAmountNil sets the value for UsedAmount to be an explicit nil

### UnsetUsedAmount
`func (o *ResourceUsedAmountDTO) UnsetUsedAmount()`

UnsetUsedAmount ensures that no value is present for UsedAmount, not even an explicit nil
### GetUsedTime

`func (o *ResourceUsedAmountDTO) GetUsedTime() int32`

GetUsedTime returns the UsedTime field if non-nil, zero value otherwise.

### GetUsedTimeOk

`func (o *ResourceUsedAmountDTO) GetUsedTimeOk() (*int32, bool)`

GetUsedTimeOk returns a tuple with the UsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTime

`func (o *ResourceUsedAmountDTO) SetUsedTime(v int32)`

SetUsedTime sets UsedTime field to given value.

### HasUsedTime

`func (o *ResourceUsedAmountDTO) HasUsedTime() bool`

HasUsedTime returns a boolean if a field has been set.

### SetUsedTimeNil

`func (o *ResourceUsedAmountDTO) SetUsedTimeNil(b bool)`

 SetUsedTimeNil sets the value for UsedTime to be an explicit nil

### UnsetUsedTime
`func (o *ResourceUsedAmountDTO) UnsetUsedTime()`

UnsetUsedTime ensures that no value is present for UsedTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


