# AutoScalingGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScalingGroups** | [**[]AutoScalingGroup**](AutoScalingGroup.md) | Auto-Scaling Group list | 
**Count** | **int32** | Count | 

## Methods

### NewAutoScalingGroupListResponse

`func NewAutoScalingGroupListResponse(autoScalingGroups []AutoScalingGroup, count int32, ) *AutoScalingGroupListResponse`

NewAutoScalingGroupListResponse instantiates a new AutoScalingGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupListResponseWithDefaults

`func NewAutoScalingGroupListResponseWithDefaults() *AutoScalingGroupListResponse`

NewAutoScalingGroupListResponseWithDefaults instantiates a new AutoScalingGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScalingGroups

`func (o *AutoScalingGroupListResponse) GetAutoScalingGroups() []AutoScalingGroup`

GetAutoScalingGroups returns the AutoScalingGroups field if non-nil, zero value otherwise.

### GetAutoScalingGroupsOk

`func (o *AutoScalingGroupListResponse) GetAutoScalingGroupsOk() (*[]AutoScalingGroup, bool)`

GetAutoScalingGroupsOk returns a tuple with the AutoScalingGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroups

`func (o *AutoScalingGroupListResponse) SetAutoScalingGroups(v []AutoScalingGroup)`

SetAutoScalingGroups sets AutoScalingGroups field to given value.


### GetCount

`func (o *AutoScalingGroupListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AutoScalingGroupListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AutoScalingGroupListResponse) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


