# ResourceUsageLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limitation** | **int32** | resource limitation count | 
**Remain** | **int32** | resource remain count | 
**Usage** | **int32** | resource usage count | 

## Methods

### NewResourceUsageLimitResponse

`func NewResourceUsageLimitResponse(limitation int32, remain int32, usage int32, ) *ResourceUsageLimitResponse`

NewResourceUsageLimitResponse instantiates a new ResourceUsageLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUsageLimitResponseWithDefaults

`func NewResourceUsageLimitResponseWithDefaults() *ResourceUsageLimitResponse`

NewResourceUsageLimitResponseWithDefaults instantiates a new ResourceUsageLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitation

`func (o *ResourceUsageLimitResponse) GetLimitation() int32`

GetLimitation returns the Limitation field if non-nil, zero value otherwise.

### GetLimitationOk

`func (o *ResourceUsageLimitResponse) GetLimitationOk() (*int32, bool)`

GetLimitationOk returns a tuple with the Limitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitation

`func (o *ResourceUsageLimitResponse) SetLimitation(v int32)`

SetLimitation sets Limitation field to given value.


### GetRemain

`func (o *ResourceUsageLimitResponse) GetRemain() int32`

GetRemain returns the Remain field if non-nil, zero value otherwise.

### GetRemainOk

`func (o *ResourceUsageLimitResponse) GetRemainOk() (*int32, bool)`

GetRemainOk returns a tuple with the Remain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemain

`func (o *ResourceUsageLimitResponse) SetRemain(v int32)`

SetRemain sets Remain field to given value.


### GetUsage

`func (o *ResourceUsageLimitResponse) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ResourceUsageLimitResponse) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ResourceUsageLimitResponse) SetUsage(v int32)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


