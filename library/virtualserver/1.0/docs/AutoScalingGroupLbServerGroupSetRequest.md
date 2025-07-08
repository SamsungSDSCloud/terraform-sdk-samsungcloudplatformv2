# AutoScalingGroupLbServerGroupSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LbServerGroups** | [**[]AutoScalingGroupLbServerGroup**](AutoScalingGroupLbServerGroup.md) | LB Server Group list | 

## Methods

### NewAutoScalingGroupLbServerGroupSetRequest

`func NewAutoScalingGroupLbServerGroupSetRequest(lbServerGroups []AutoScalingGroupLbServerGroup, ) *AutoScalingGroupLbServerGroupSetRequest`

NewAutoScalingGroupLbServerGroupSetRequest instantiates a new AutoScalingGroupLbServerGroupSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupLbServerGroupSetRequestWithDefaults

`func NewAutoScalingGroupLbServerGroupSetRequestWithDefaults() *AutoScalingGroupLbServerGroupSetRequest`

NewAutoScalingGroupLbServerGroupSetRequestWithDefaults instantiates a new AutoScalingGroupLbServerGroupSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLbServerGroups

`func (o *AutoScalingGroupLbServerGroupSetRequest) GetLbServerGroups() []AutoScalingGroupLbServerGroup`

GetLbServerGroups returns the LbServerGroups field if non-nil, zero value otherwise.

### GetLbServerGroupsOk

`func (o *AutoScalingGroupLbServerGroupSetRequest) GetLbServerGroupsOk() (*[]AutoScalingGroupLbServerGroup, bool)`

GetLbServerGroupsOk returns a tuple with the LbServerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbServerGroups

`func (o *AutoScalingGroupLbServerGroupSetRequest) SetLbServerGroups(v []AutoScalingGroupLbServerGroup)`

SetLbServerGroups sets LbServerGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


