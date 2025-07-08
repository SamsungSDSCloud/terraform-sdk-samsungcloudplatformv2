# PlannedComputeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current Page | [optional] [default to 0]
**PlannedComputes** | [**[]PlannedComputeDTO**](PlannedComputeDTO.md) |  | 
**TotalCount** | Pointer to **int32** | Total Count | [optional] [default to 0]
**TotalPages** | Pointer to **int32** | Total Pages | [optional] [default to 0]

## Methods

### NewPlannedComputeListResponse

`func NewPlannedComputeListResponse(plannedComputes []PlannedComputeDTO, ) *PlannedComputeListResponse`

NewPlannedComputeListResponse instantiates a new PlannedComputeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannedComputeListResponseWithDefaults

`func NewPlannedComputeListResponseWithDefaults() *PlannedComputeListResponse`

NewPlannedComputeListResponseWithDefaults instantiates a new PlannedComputeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *PlannedComputeListResponse) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *PlannedComputeListResponse) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *PlannedComputeListResponse) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *PlannedComputeListResponse) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetPlannedComputes

`func (o *PlannedComputeListResponse) GetPlannedComputes() []PlannedComputeDTO`

GetPlannedComputes returns the PlannedComputes field if non-nil, zero value otherwise.

### GetPlannedComputesOk

`func (o *PlannedComputeListResponse) GetPlannedComputesOk() (*[]PlannedComputeDTO, bool)`

GetPlannedComputesOk returns a tuple with the PlannedComputes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedComputes

`func (o *PlannedComputeListResponse) SetPlannedComputes(v []PlannedComputeDTO)`

SetPlannedComputes sets PlannedComputes field to given value.


### GetTotalCount

`func (o *PlannedComputeListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PlannedComputeListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PlannedComputeListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PlannedComputeListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *PlannedComputeListResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PlannedComputeListResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PlannedComputeListResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PlannedComputeListResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


