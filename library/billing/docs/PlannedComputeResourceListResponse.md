# PlannedComputeResourceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**ResourceGroupTotalAmountDTO**](ResourceGroupTotalAmountDTO.md) | Resource Group Amount | [optional] 
**Coverages** | [**[]ResourceUsedAmountDTO**](ResourceUsedAmountDTO.md) |  | 
**OrderCount** | Pointer to **int32** | Order Count | [optional] 
**Os** | Pointer to [**OSTypeDTO**](OSTypeDTO.md) | OS Type | [optional] 
**ServerType** | Pointer to **string** | Server Type | [optional] 
**ServerTypeDescription** | Pointer to **map[string]interface{}** | Server Type Description | [optional] 
**Service** | Pointer to [**ProductResponse**](ProductResponse.md) | Product Type Name | [optional] 

## Methods

### NewPlannedComputeResourceListResponse

`func NewPlannedComputeResourceListResponse(coverages []ResourceUsedAmountDTO, ) *PlannedComputeResourceListResponse`

NewPlannedComputeResourceListResponse instantiates a new PlannedComputeResourceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannedComputeResourceListResponseWithDefaults

`func NewPlannedComputeResourceListResponseWithDefaults() *PlannedComputeResourceListResponse`

NewPlannedComputeResourceListResponseWithDefaults instantiates a new PlannedComputeResourceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PlannedComputeResourceListResponse) GetAmount() ResourceGroupTotalAmountDTO`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PlannedComputeResourceListResponse) GetAmountOk() (*ResourceGroupTotalAmountDTO, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PlannedComputeResourceListResponse) SetAmount(v ResourceGroupTotalAmountDTO)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PlannedComputeResourceListResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCoverages

`func (o *PlannedComputeResourceListResponse) GetCoverages() []ResourceUsedAmountDTO`

GetCoverages returns the Coverages field if non-nil, zero value otherwise.

### GetCoveragesOk

`func (o *PlannedComputeResourceListResponse) GetCoveragesOk() (*[]ResourceUsedAmountDTO, bool)`

GetCoveragesOk returns a tuple with the Coverages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverages

`func (o *PlannedComputeResourceListResponse) SetCoverages(v []ResourceUsedAmountDTO)`

SetCoverages sets Coverages field to given value.


### GetOrderCount

`func (o *PlannedComputeResourceListResponse) GetOrderCount() int32`

GetOrderCount returns the OrderCount field if non-nil, zero value otherwise.

### GetOrderCountOk

`func (o *PlannedComputeResourceListResponse) GetOrderCountOk() (*int32, bool)`

GetOrderCountOk returns a tuple with the OrderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCount

`func (o *PlannedComputeResourceListResponse) SetOrderCount(v int32)`

SetOrderCount sets OrderCount field to given value.

### HasOrderCount

`func (o *PlannedComputeResourceListResponse) HasOrderCount() bool`

HasOrderCount returns a boolean if a field has been set.

### GetOs

`func (o *PlannedComputeResourceListResponse) GetOs() OSTypeDTO`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *PlannedComputeResourceListResponse) GetOsOk() (*OSTypeDTO, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *PlannedComputeResourceListResponse) SetOs(v OSTypeDTO)`

SetOs sets Os field to given value.

### HasOs

`func (o *PlannedComputeResourceListResponse) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetServerType

`func (o *PlannedComputeResourceListResponse) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *PlannedComputeResourceListResponse) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *PlannedComputeResourceListResponse) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *PlannedComputeResourceListResponse) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerTypeDescription

`func (o *PlannedComputeResourceListResponse) GetServerTypeDescription() map[string]interface{}`

GetServerTypeDescription returns the ServerTypeDescription field if non-nil, zero value otherwise.

### GetServerTypeDescriptionOk

`func (o *PlannedComputeResourceListResponse) GetServerTypeDescriptionOk() (*map[string]interface{}, bool)`

GetServerTypeDescriptionOk returns a tuple with the ServerTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeDescription

`func (o *PlannedComputeResourceListResponse) SetServerTypeDescription(v map[string]interface{})`

SetServerTypeDescription sets ServerTypeDescription field to given value.

### HasServerTypeDescription

`func (o *PlannedComputeResourceListResponse) HasServerTypeDescription() bool`

HasServerTypeDescription returns a boolean if a field has been set.

### GetService

`func (o *PlannedComputeResourceListResponse) GetService() ProductResponse`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PlannedComputeResourceListResponse) GetServiceOk() (*ProductResponse, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PlannedComputeResourceListResponse) SetService(v ProductResponse)`

SetService sets Service field to given value.

### HasService

`func (o *PlannedComputeResourceListResponse) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


