# AutoScalingGroupLbServerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | LB Server Group ID | 
**Port** | **int32** | LB Server Group port | 

## Methods

### NewAutoScalingGroupLbServerGroup

`func NewAutoScalingGroupLbServerGroup(id string, port int32, ) *AutoScalingGroupLbServerGroup`

NewAutoScalingGroupLbServerGroup instantiates a new AutoScalingGroupLbServerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupLbServerGroupWithDefaults

`func NewAutoScalingGroupLbServerGroupWithDefaults() *AutoScalingGroupLbServerGroup`

NewAutoScalingGroupLbServerGroupWithDefaults instantiates a new AutoScalingGroupLbServerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoScalingGroupLbServerGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoScalingGroupLbServerGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoScalingGroupLbServerGroup) SetId(v string)`

SetId sets Id field to given value.


### GetPort

`func (o *AutoScalingGroupLbServerGroup) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AutoScalingGroupLbServerGroup) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AutoScalingGroupLbServerGroup) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


