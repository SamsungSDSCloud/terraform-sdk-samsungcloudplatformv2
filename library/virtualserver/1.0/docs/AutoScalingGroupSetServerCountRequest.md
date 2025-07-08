# AutoScalingGroupSetServerCountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredServerCount** | Pointer to **NullableInt32** |  | [optional] 
**MaxServerCount** | Pointer to **NullableInt32** |  | [optional] 
**MinServerCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAutoScalingGroupSetServerCountRequest

`func NewAutoScalingGroupSetServerCountRequest() *AutoScalingGroupSetServerCountRequest`

NewAutoScalingGroupSetServerCountRequest instantiates a new AutoScalingGroupSetServerCountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupSetServerCountRequestWithDefaults

`func NewAutoScalingGroupSetServerCountRequestWithDefaults() *AutoScalingGroupSetServerCountRequest`

NewAutoScalingGroupSetServerCountRequestWithDefaults instantiates a new AutoScalingGroupSetServerCountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredServerCount

`func (o *AutoScalingGroupSetServerCountRequest) GetDesiredServerCount() int32`

GetDesiredServerCount returns the DesiredServerCount field if non-nil, zero value otherwise.

### GetDesiredServerCountOk

`func (o *AutoScalingGroupSetServerCountRequest) GetDesiredServerCountOk() (*int32, bool)`

GetDesiredServerCountOk returns a tuple with the DesiredServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredServerCount

`func (o *AutoScalingGroupSetServerCountRequest) SetDesiredServerCount(v int32)`

SetDesiredServerCount sets DesiredServerCount field to given value.

### HasDesiredServerCount

`func (o *AutoScalingGroupSetServerCountRequest) HasDesiredServerCount() bool`

HasDesiredServerCount returns a boolean if a field has been set.

### SetDesiredServerCountNil

`func (o *AutoScalingGroupSetServerCountRequest) SetDesiredServerCountNil(b bool)`

 SetDesiredServerCountNil sets the value for DesiredServerCount to be an explicit nil

### UnsetDesiredServerCount
`func (o *AutoScalingGroupSetServerCountRequest) UnsetDesiredServerCount()`

UnsetDesiredServerCount ensures that no value is present for DesiredServerCount, not even an explicit nil
### GetMaxServerCount

`func (o *AutoScalingGroupSetServerCountRequest) GetMaxServerCount() int32`

GetMaxServerCount returns the MaxServerCount field if non-nil, zero value otherwise.

### GetMaxServerCountOk

`func (o *AutoScalingGroupSetServerCountRequest) GetMaxServerCountOk() (*int32, bool)`

GetMaxServerCountOk returns a tuple with the MaxServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxServerCount

`func (o *AutoScalingGroupSetServerCountRequest) SetMaxServerCount(v int32)`

SetMaxServerCount sets MaxServerCount field to given value.

### HasMaxServerCount

`func (o *AutoScalingGroupSetServerCountRequest) HasMaxServerCount() bool`

HasMaxServerCount returns a boolean if a field has been set.

### SetMaxServerCountNil

`func (o *AutoScalingGroupSetServerCountRequest) SetMaxServerCountNil(b bool)`

 SetMaxServerCountNil sets the value for MaxServerCount to be an explicit nil

### UnsetMaxServerCount
`func (o *AutoScalingGroupSetServerCountRequest) UnsetMaxServerCount()`

UnsetMaxServerCount ensures that no value is present for MaxServerCount, not even an explicit nil
### GetMinServerCount

`func (o *AutoScalingGroupSetServerCountRequest) GetMinServerCount() int32`

GetMinServerCount returns the MinServerCount field if non-nil, zero value otherwise.

### GetMinServerCountOk

`func (o *AutoScalingGroupSetServerCountRequest) GetMinServerCountOk() (*int32, bool)`

GetMinServerCountOk returns a tuple with the MinServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinServerCount

`func (o *AutoScalingGroupSetServerCountRequest) SetMinServerCount(v int32)`

SetMinServerCount sets MinServerCount field to given value.

### HasMinServerCount

`func (o *AutoScalingGroupSetServerCountRequest) HasMinServerCount() bool`

HasMinServerCount returns a boolean if a field has been set.

### SetMinServerCountNil

`func (o *AutoScalingGroupSetServerCountRequest) SetMinServerCountNil(b bool)`

 SetMinServerCountNil sets the value for MinServerCount to be an explicit nil

### UnsetMinServerCount
`func (o *AutoScalingGroupSetServerCountRequest) UnsetMinServerCount()`

UnsetMinServerCount ensures that no value is present for MinServerCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


