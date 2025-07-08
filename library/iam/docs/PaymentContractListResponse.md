# PaymentContractListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**PaymentContracts** | [**[]PaymentContract**](PaymentContract.md) |  | 

## Methods

### NewPaymentContractListResponse

`func NewPaymentContractListResponse(paymentContracts []PaymentContract, ) *PaymentContractListResponse`

NewPaymentContractListResponse instantiates a new PaymentContractListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentContractListResponseWithDefaults

`func NewPaymentContractListResponseWithDefaults() *PaymentContractListResponse`

NewPaymentContractListResponseWithDefaults instantiates a new PaymentContractListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaymentContractListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaymentContractListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaymentContractListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaymentContractListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *PaymentContractListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PaymentContractListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *PaymentContractListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaymentContractListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaymentContractListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaymentContractListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *PaymentContractListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *PaymentContractListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetPaymentContracts

`func (o *PaymentContractListResponse) GetPaymentContracts() []PaymentContract`

GetPaymentContracts returns the PaymentContracts field if non-nil, zero value otherwise.

### GetPaymentContractsOk

`func (o *PaymentContractListResponse) GetPaymentContractsOk() (*[]PaymentContract, bool)`

GetPaymentContractsOk returns a tuple with the PaymentContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentContracts

`func (o *PaymentContractListResponse) SetPaymentContracts(v []PaymentContract)`

SetPaymentContracts sets PaymentContracts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


