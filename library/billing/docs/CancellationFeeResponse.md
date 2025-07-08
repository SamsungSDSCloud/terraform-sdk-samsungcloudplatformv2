# CancellationFeeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillYearMonth** | Pointer to **string** | Bill year Month | [optional] 
**CancellationFee** | Pointer to [**Amount**](Amount.md) | 해지 수수료 | [optional] 
**SameDayCancel** | Pointer to **bool** | Same day Cancel or not | [optional] [default to false]

## Methods

### NewCancellationFeeResponse

`func NewCancellationFeeResponse() *CancellationFeeResponse`

NewCancellationFeeResponse instantiates a new CancellationFeeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancellationFeeResponseWithDefaults

`func NewCancellationFeeResponseWithDefaults() *CancellationFeeResponse`

NewCancellationFeeResponseWithDefaults instantiates a new CancellationFeeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillYearMonth

`func (o *CancellationFeeResponse) GetBillYearMonth() string`

GetBillYearMonth returns the BillYearMonth field if non-nil, zero value otherwise.

### GetBillYearMonthOk

`func (o *CancellationFeeResponse) GetBillYearMonthOk() (*string, bool)`

GetBillYearMonthOk returns a tuple with the BillYearMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillYearMonth

`func (o *CancellationFeeResponse) SetBillYearMonth(v string)`

SetBillYearMonth sets BillYearMonth field to given value.

### HasBillYearMonth

`func (o *CancellationFeeResponse) HasBillYearMonth() bool`

HasBillYearMonth returns a boolean if a field has been set.

### GetCancellationFee

`func (o *CancellationFeeResponse) GetCancellationFee() Amount`

GetCancellationFee returns the CancellationFee field if non-nil, zero value otherwise.

### GetCancellationFeeOk

`func (o *CancellationFeeResponse) GetCancellationFeeOk() (*Amount, bool)`

GetCancellationFeeOk returns a tuple with the CancellationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationFee

`func (o *CancellationFeeResponse) SetCancellationFee(v Amount)`

SetCancellationFee sets CancellationFee field to given value.

### HasCancellationFee

`func (o *CancellationFeeResponse) HasCancellationFee() bool`

HasCancellationFee returns a boolean if a field has been set.

### GetSameDayCancel

`func (o *CancellationFeeResponse) GetSameDayCancel() bool`

GetSameDayCancel returns the SameDayCancel field if non-nil, zero value otherwise.

### GetSameDayCancelOk

`func (o *CancellationFeeResponse) GetSameDayCancelOk() (*bool, bool)`

GetSameDayCancelOk returns a tuple with the SameDayCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameDayCancel

`func (o *CancellationFeeResponse) SetSameDayCancel(v bool)`

SetSameDayCancel sets SameDayCancel field to given value.

### HasSameDayCancel

`func (o *CancellationFeeResponse) HasSameDayCancel() bool`

HasSameDayCancel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


