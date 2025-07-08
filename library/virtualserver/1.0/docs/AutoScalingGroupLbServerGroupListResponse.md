# AutoScalingGroupLbServerGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Count | 
**LbServerGroups** | [**[]AutoScalingGroupLbServerGroupShowResponse**](AutoScalingGroupLbServerGroupShowResponse.md) | LB Server Group list | 

## Methods

### NewAutoScalingGroupLbServerGroupListResponse

`func NewAutoScalingGroupLbServerGroupListResponse(count int32, lbServerGroups []AutoScalingGroupLbServerGroupShowResponse, ) *AutoScalingGroupLbServerGroupListResponse`

NewAutoScalingGroupLbServerGroupListResponse instantiates a new AutoScalingGroupLbServerGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupLbServerGroupListResponseWithDefaults

`func NewAutoScalingGroupLbServerGroupListResponseWithDefaults() *AutoScalingGroupLbServerGroupListResponse`

NewAutoScalingGroupLbServerGroupListResponseWithDefaults instantiates a new AutoScalingGroupLbServerGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AutoScalingGroupLbServerGroupListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AutoScalingGroupLbServerGroupListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AutoScalingGroupLbServerGroupListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLbServerGroups

`func (o *AutoScalingGroupLbServerGroupListResponse) GetLbServerGroups() []AutoScalingGroupLbServerGroupShowResponse`

GetLbServerGroups returns the LbServerGroups field if non-nil, zero value otherwise.

### GetLbServerGroupsOk

`func (o *AutoScalingGroupLbServerGroupListResponse) GetLbServerGroupsOk() (*[]AutoScalingGroupLbServerGroupShowResponse, bool)`

GetLbServerGroupsOk returns a tuple with the LbServerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbServerGroups

`func (o *AutoScalingGroupLbServerGroupListResponse) SetLbServerGroups(v []AutoScalingGroupLbServerGroupShowResponse)`

SetLbServerGroups sets LbServerGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


