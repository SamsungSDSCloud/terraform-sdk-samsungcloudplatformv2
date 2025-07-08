# AutoScalingGroupPolicyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Count | 
**Policies** | [**[]AutoScalingGroupPolicyShowResponse**](AutoScalingGroupPolicyShowResponse.md) |  | 

## Methods

### NewAutoScalingGroupPolicyListResponse

`func NewAutoScalingGroupPolicyListResponse(count int32, policies []AutoScalingGroupPolicyShowResponse, ) *AutoScalingGroupPolicyListResponse`

NewAutoScalingGroupPolicyListResponse instantiates a new AutoScalingGroupPolicyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupPolicyListResponseWithDefaults

`func NewAutoScalingGroupPolicyListResponseWithDefaults() *AutoScalingGroupPolicyListResponse`

NewAutoScalingGroupPolicyListResponseWithDefaults instantiates a new AutoScalingGroupPolicyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AutoScalingGroupPolicyListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AutoScalingGroupPolicyListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AutoScalingGroupPolicyListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPolicies

`func (o *AutoScalingGroupPolicyListResponse) GetPolicies() []AutoScalingGroupPolicyShowResponse`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AutoScalingGroupPolicyListResponse) GetPoliciesOk() (*[]AutoScalingGroupPolicyShowResponse, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AutoScalingGroupPolicyListResponse) SetPolicies(v []AutoScalingGroupPolicyShowResponse)`

SetPolicies sets Policies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


