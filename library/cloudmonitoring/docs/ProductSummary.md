# ProductSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentState** | **string** | 상품 상태 | 
**LbName** | Pointer to **string** | swagger.event.productSummary.lbName.value | [optional] 
**LbSize** | Pointer to **string** | swagger.event.productSummary.lbSize.value | [optional] 
**ProductIpAddress** | Pointer to **string** | 상품 IP주소 | [optional] 
**ProductName** | **string** | 상품 이름 | 
**ProductResourceId** | **string** | 상품 리소스 아이디 | 
**ProductSq** | Pointer to **int64** | 상품 시퀀스 | [optional] 
**ProductState** | **string** | 상품 상태 | 
**ProductSubName** | Pointer to **string** | swagger.event.productSummary.productSubName.value | [optional] 
**ProductSubType** | Pointer to **string** | swagger.event.productSummary.productSubType.value | [optional] 
**ProductTypeCode** | **string** | 상품 유형 코드 | 
**ProductTypeName** | **string** | 상품 유형 이름 | 
**VpcName** | Pointer to **string** | swagger.event.productSummary.vpcName.value | [optional] 

## Methods

### NewProductSummary

`func NewProductSummary(agentState string, productName string, productResourceId string, productState string, productTypeCode string, productTypeName string, ) *ProductSummary`

NewProductSummary instantiates a new ProductSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSummaryWithDefaults

`func NewProductSummaryWithDefaults() *ProductSummary`

NewProductSummaryWithDefaults instantiates a new ProductSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentState

`func (o *ProductSummary) GetAgentState() string`

GetAgentState returns the AgentState field if non-nil, zero value otherwise.

### GetAgentStateOk

`func (o *ProductSummary) GetAgentStateOk() (*string, bool)`

GetAgentStateOk returns a tuple with the AgentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentState

`func (o *ProductSummary) SetAgentState(v string)`

SetAgentState sets AgentState field to given value.


### GetLbName

`func (o *ProductSummary) GetLbName() string`

GetLbName returns the LbName field if non-nil, zero value otherwise.

### GetLbNameOk

`func (o *ProductSummary) GetLbNameOk() (*string, bool)`

GetLbNameOk returns a tuple with the LbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbName

`func (o *ProductSummary) SetLbName(v string)`

SetLbName sets LbName field to given value.

### HasLbName

`func (o *ProductSummary) HasLbName() bool`

HasLbName returns a boolean if a field has been set.

### GetLbSize

`func (o *ProductSummary) GetLbSize() string`

GetLbSize returns the LbSize field if non-nil, zero value otherwise.

### GetLbSizeOk

`func (o *ProductSummary) GetLbSizeOk() (*string, bool)`

GetLbSizeOk returns a tuple with the LbSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbSize

`func (o *ProductSummary) SetLbSize(v string)`

SetLbSize sets LbSize field to given value.

### HasLbSize

`func (o *ProductSummary) HasLbSize() bool`

HasLbSize returns a boolean if a field has been set.

### GetProductIpAddress

`func (o *ProductSummary) GetProductIpAddress() string`

GetProductIpAddress returns the ProductIpAddress field if non-nil, zero value otherwise.

### GetProductIpAddressOk

`func (o *ProductSummary) GetProductIpAddressOk() (*string, bool)`

GetProductIpAddressOk returns a tuple with the ProductIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIpAddress

`func (o *ProductSummary) SetProductIpAddress(v string)`

SetProductIpAddress sets ProductIpAddress field to given value.

### HasProductIpAddress

`func (o *ProductSummary) HasProductIpAddress() bool`

HasProductIpAddress returns a boolean if a field has been set.

### GetProductName

`func (o *ProductSummary) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ProductSummary) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ProductSummary) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetProductResourceId

`func (o *ProductSummary) GetProductResourceId() string`

GetProductResourceId returns the ProductResourceId field if non-nil, zero value otherwise.

### GetProductResourceIdOk

`func (o *ProductSummary) GetProductResourceIdOk() (*string, bool)`

GetProductResourceIdOk returns a tuple with the ProductResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductResourceId

`func (o *ProductSummary) SetProductResourceId(v string)`

SetProductResourceId sets ProductResourceId field to given value.


### GetProductSq

`func (o *ProductSummary) GetProductSq() int64`

GetProductSq returns the ProductSq field if non-nil, zero value otherwise.

### GetProductSqOk

`func (o *ProductSummary) GetProductSqOk() (*int64, bool)`

GetProductSqOk returns a tuple with the ProductSq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSq

`func (o *ProductSummary) SetProductSq(v int64)`

SetProductSq sets ProductSq field to given value.

### HasProductSq

`func (o *ProductSummary) HasProductSq() bool`

HasProductSq returns a boolean if a field has been set.

### GetProductState

`func (o *ProductSummary) GetProductState() string`

GetProductState returns the ProductState field if non-nil, zero value otherwise.

### GetProductStateOk

`func (o *ProductSummary) GetProductStateOk() (*string, bool)`

GetProductStateOk returns a tuple with the ProductState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductState

`func (o *ProductSummary) SetProductState(v string)`

SetProductState sets ProductState field to given value.


### GetProductSubName

`func (o *ProductSummary) GetProductSubName() string`

GetProductSubName returns the ProductSubName field if non-nil, zero value otherwise.

### GetProductSubNameOk

`func (o *ProductSummary) GetProductSubNameOk() (*string, bool)`

GetProductSubNameOk returns a tuple with the ProductSubName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubName

`func (o *ProductSummary) SetProductSubName(v string)`

SetProductSubName sets ProductSubName field to given value.

### HasProductSubName

`func (o *ProductSummary) HasProductSubName() bool`

HasProductSubName returns a boolean if a field has been set.

### GetProductSubType

`func (o *ProductSummary) GetProductSubType() string`

GetProductSubType returns the ProductSubType field if non-nil, zero value otherwise.

### GetProductSubTypeOk

`func (o *ProductSummary) GetProductSubTypeOk() (*string, bool)`

GetProductSubTypeOk returns a tuple with the ProductSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubType

`func (o *ProductSummary) SetProductSubType(v string)`

SetProductSubType sets ProductSubType field to given value.

### HasProductSubType

`func (o *ProductSummary) HasProductSubType() bool`

HasProductSubType returns a boolean if a field has been set.

### GetProductTypeCode

`func (o *ProductSummary) GetProductTypeCode() string`

GetProductTypeCode returns the ProductTypeCode field if non-nil, zero value otherwise.

### GetProductTypeCodeOk

`func (o *ProductSummary) GetProductTypeCodeOk() (*string, bool)`

GetProductTypeCodeOk returns a tuple with the ProductTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeCode

`func (o *ProductSummary) SetProductTypeCode(v string)`

SetProductTypeCode sets ProductTypeCode field to given value.


### GetProductTypeName

`func (o *ProductSummary) GetProductTypeName() string`

GetProductTypeName returns the ProductTypeName field if non-nil, zero value otherwise.

### GetProductTypeNameOk

`func (o *ProductSummary) GetProductTypeNameOk() (*string, bool)`

GetProductTypeNameOk returns a tuple with the ProductTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeName

`func (o *ProductSummary) SetProductTypeName(v string)`

SetProductTypeName sets ProductTypeName field to given value.


### GetVpcName

`func (o *ProductSummary) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *ProductSummary) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *ProductSummary) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *ProductSummary) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


