# CancellationFeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlannedComputeId** | Pointer to **[]string** | Planned Compute Id | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCancellationFeeRequest

`func NewCancellationFeeRequest() *CancellationFeeRequest`

NewCancellationFeeRequest instantiates a new CancellationFeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancellationFeeRequestWithDefaults

`func NewCancellationFeeRequestWithDefaults() *CancellationFeeRequest`

NewCancellationFeeRequestWithDefaults instantiates a new CancellationFeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlannedComputeId

`func (o *CancellationFeeRequest) GetPlannedComputeId() []string`

GetPlannedComputeId returns the PlannedComputeId field if non-nil, zero value otherwise.

### GetPlannedComputeIdOk

`func (o *CancellationFeeRequest) GetPlannedComputeIdOk() (*[]string, bool)`

GetPlannedComputeIdOk returns a tuple with the PlannedComputeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedComputeId

`func (o *CancellationFeeRequest) SetPlannedComputeId(v []string)`

SetPlannedComputeId sets PlannedComputeId field to given value.

### HasPlannedComputeId

`func (o *CancellationFeeRequest) HasPlannedComputeId() bool`

HasPlannedComputeId returns a boolean if a field has been set.

### GetRegion

`func (o *CancellationFeeRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CancellationFeeRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CancellationFeeRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CancellationFeeRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *CancellationFeeRequest) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CancellationFeeRequest) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


