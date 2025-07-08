# ProductResourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectList** | Pointer to **[]string** | 개별항목 목록 | [optional] 
**ProductResourceId** | **string** | 상품 리소스 아이디 | 

## Methods

### NewProductResourceInfo

`func NewProductResourceInfo(productResourceId string, ) *ProductResourceInfo`

NewProductResourceInfo instantiates a new ProductResourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductResourceInfoWithDefaults

`func NewProductResourceInfoWithDefaults() *ProductResourceInfo`

NewProductResourceInfoWithDefaults instantiates a new ProductResourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectList

`func (o *ProductResourceInfo) GetObjectList() []string`

GetObjectList returns the ObjectList field if non-nil, zero value otherwise.

### GetObjectListOk

`func (o *ProductResourceInfo) GetObjectListOk() (*[]string, bool)`

GetObjectListOk returns a tuple with the ObjectList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectList

`func (o *ProductResourceInfo) SetObjectList(v []string)`

SetObjectList sets ObjectList field to given value.

### HasObjectList

`func (o *ProductResourceInfo) HasObjectList() bool`

HasObjectList returns a boolean if a field has been set.

### GetProductResourceId

`func (o *ProductResourceInfo) GetProductResourceId() string`

GetProductResourceId returns the ProductResourceId field if non-nil, zero value otherwise.

### GetProductResourceIdOk

`func (o *ProductResourceInfo) GetProductResourceIdOk() (*string, bool)`

GetProductResourceIdOk returns a tuple with the ProductResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductResourceId

`func (o *ProductResourceInfo) SetProductResourceId(v string)`

SetProductResourceId sets ProductResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


