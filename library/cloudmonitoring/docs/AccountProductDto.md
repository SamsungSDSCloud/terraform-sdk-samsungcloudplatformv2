# AccountProductDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | 프로젝트 아이디 | 
**EndDt** | Pointer to **time.Time** | 상품 삭제 일시 | [optional] 
**LastEventLevel** | Pointer to **string** | 최종 이벤트 레벨 | [optional] 
**PoolName** | Pointer to **string** | 풀 이름 | [optional] 
**ProductIpAddress** | Pointer to **string** | 상품 IP주소 | [optional] 
**ProductName** | **string** | 상품 이름 | 
**ProductResourceId** | **string** | 상품 리소스 아이디 | 
**ProductState** | **string** | 상품 상태 - 유효한 값 : Running, Down | 
**ProductTypeCode** | **string** | 상품 유형코드 | 
**ProductTypeName** | **string** | 상품 유형이름 | 
**StartDt** | **time.Time** | 상품 생성 일시 | 

## Methods

### NewAccountProductDto

`func NewAccountProductDto(accountId string, productName string, productResourceId string, productState string, productTypeCode string, productTypeName string, startDt time.Time, ) *AccountProductDto`

NewAccountProductDto instantiates a new AccountProductDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountProductDtoWithDefaults

`func NewAccountProductDtoWithDefaults() *AccountProductDto`

NewAccountProductDtoWithDefaults instantiates a new AccountProductDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountProductDto) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountProductDto) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountProductDto) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetEndDt

`func (o *AccountProductDto) GetEndDt() time.Time`

GetEndDt returns the EndDt field if non-nil, zero value otherwise.

### GetEndDtOk

`func (o *AccountProductDto) GetEndDtOk() (*time.Time, bool)`

GetEndDtOk returns a tuple with the EndDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDt

`func (o *AccountProductDto) SetEndDt(v time.Time)`

SetEndDt sets EndDt field to given value.

### HasEndDt

`func (o *AccountProductDto) HasEndDt() bool`

HasEndDt returns a boolean if a field has been set.

### GetLastEventLevel

`func (o *AccountProductDto) GetLastEventLevel() string`

GetLastEventLevel returns the LastEventLevel field if non-nil, zero value otherwise.

### GetLastEventLevelOk

`func (o *AccountProductDto) GetLastEventLevelOk() (*string, bool)`

GetLastEventLevelOk returns a tuple with the LastEventLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventLevel

`func (o *AccountProductDto) SetLastEventLevel(v string)`

SetLastEventLevel sets LastEventLevel field to given value.

### HasLastEventLevel

`func (o *AccountProductDto) HasLastEventLevel() bool`

HasLastEventLevel returns a boolean if a field has been set.

### GetPoolName

`func (o *AccountProductDto) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AccountProductDto) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AccountProductDto) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *AccountProductDto) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetProductIpAddress

`func (o *AccountProductDto) GetProductIpAddress() string`

GetProductIpAddress returns the ProductIpAddress field if non-nil, zero value otherwise.

### GetProductIpAddressOk

`func (o *AccountProductDto) GetProductIpAddressOk() (*string, bool)`

GetProductIpAddressOk returns a tuple with the ProductIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIpAddress

`func (o *AccountProductDto) SetProductIpAddress(v string)`

SetProductIpAddress sets ProductIpAddress field to given value.

### HasProductIpAddress

`func (o *AccountProductDto) HasProductIpAddress() bool`

HasProductIpAddress returns a boolean if a field has been set.

### GetProductName

`func (o *AccountProductDto) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *AccountProductDto) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *AccountProductDto) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetProductResourceId

`func (o *AccountProductDto) GetProductResourceId() string`

GetProductResourceId returns the ProductResourceId field if non-nil, zero value otherwise.

### GetProductResourceIdOk

`func (o *AccountProductDto) GetProductResourceIdOk() (*string, bool)`

GetProductResourceIdOk returns a tuple with the ProductResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductResourceId

`func (o *AccountProductDto) SetProductResourceId(v string)`

SetProductResourceId sets ProductResourceId field to given value.


### GetProductState

`func (o *AccountProductDto) GetProductState() string`

GetProductState returns the ProductState field if non-nil, zero value otherwise.

### GetProductStateOk

`func (o *AccountProductDto) GetProductStateOk() (*string, bool)`

GetProductStateOk returns a tuple with the ProductState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductState

`func (o *AccountProductDto) SetProductState(v string)`

SetProductState sets ProductState field to given value.


### GetProductTypeCode

`func (o *AccountProductDto) GetProductTypeCode() string`

GetProductTypeCode returns the ProductTypeCode field if non-nil, zero value otherwise.

### GetProductTypeCodeOk

`func (o *AccountProductDto) GetProductTypeCodeOk() (*string, bool)`

GetProductTypeCodeOk returns a tuple with the ProductTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeCode

`func (o *AccountProductDto) SetProductTypeCode(v string)`

SetProductTypeCode sets ProductTypeCode field to given value.


### GetProductTypeName

`func (o *AccountProductDto) GetProductTypeName() string`

GetProductTypeName returns the ProductTypeName field if non-nil, zero value otherwise.

### GetProductTypeNameOk

`func (o *AccountProductDto) GetProductTypeNameOk() (*string, bool)`

GetProductTypeNameOk returns a tuple with the ProductTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeName

`func (o *AccountProductDto) SetProductTypeName(v string)`

SetProductTypeName sets ProductTypeName field to given value.


### GetStartDt

`func (o *AccountProductDto) GetStartDt() time.Time`

GetStartDt returns the StartDt field if non-nil, zero value otherwise.

### GetStartDtOk

`func (o *AccountProductDto) GetStartDtOk() (*time.Time, bool)`

GetStartDtOk returns a tuple with the StartDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDt

`func (o *AccountProductDto) SetStartDt(v time.Time)`

SetStartDt sets StartDt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


