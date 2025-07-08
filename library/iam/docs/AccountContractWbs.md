# AccountContractWbs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountContractId** | **string** | Account 계약 ID | 
**CreatedAt** | **time.Time** | 생성 일시 | 
**CreatedBy** | **string** | 생성자 | 
**CreatorEmail** | **string** | 생성자 Email | 
**CreatorName** | **string** | 생성자 성, 이름 | 
**WbsCode** | **string** | WBS Code | 

## Methods

### NewAccountContractWbs

`func NewAccountContractWbs(accountContractId string, createdAt time.Time, createdBy string, creatorEmail string, creatorName string, wbsCode string, ) *AccountContractWbs`

NewAccountContractWbs instantiates a new AccountContractWbs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountContractWbsWithDefaults

`func NewAccountContractWbsWithDefaults() *AccountContractWbs`

NewAccountContractWbsWithDefaults instantiates a new AccountContractWbs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountContractId

`func (o *AccountContractWbs) GetAccountContractId() string`

GetAccountContractId returns the AccountContractId field if non-nil, zero value otherwise.

### GetAccountContractIdOk

`func (o *AccountContractWbs) GetAccountContractIdOk() (*string, bool)`

GetAccountContractIdOk returns a tuple with the AccountContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountContractId

`func (o *AccountContractWbs) SetAccountContractId(v string)`

SetAccountContractId sets AccountContractId field to given value.


### GetCreatedAt

`func (o *AccountContractWbs) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountContractWbs) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountContractWbs) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AccountContractWbs) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AccountContractWbs) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AccountContractWbs) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatorEmail

`func (o *AccountContractWbs) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *AccountContractWbs) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *AccountContractWbs) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.


### GetCreatorName

`func (o *AccountContractWbs) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *AccountContractWbs) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *AccountContractWbs) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.


### GetWbsCode

`func (o *AccountContractWbs) GetWbsCode() string`

GetWbsCode returns the WbsCode field if non-nil, zero value otherwise.

### GetWbsCodeOk

`func (o *AccountContractWbs) GetWbsCodeOk() (*string, bool)`

GetWbsCodeOk returns a tuple with the WbsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbsCode

`func (o *AccountContractWbs) SetWbsCode(v string)`

SetWbsCode sets WbsCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


