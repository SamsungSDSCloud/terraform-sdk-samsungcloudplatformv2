# LoadbalancerCreateValidateResponseDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **bool** | Available IP Quantity Check Result | 
**Result** | **bool** | Load balancer creation availability | 
**Subnet** | **bool** | Subnet quantity check result | 

## Methods

### NewLoadbalancerCreateValidateResponseDetail

`func NewLoadbalancerCreateValidateResponseDetail(ip bool, result bool, subnet bool, ) *LoadbalancerCreateValidateResponseDetail`

NewLoadbalancerCreateValidateResponseDetail instantiates a new LoadbalancerCreateValidateResponseDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadbalancerCreateValidateResponseDetailWithDefaults

`func NewLoadbalancerCreateValidateResponseDetailWithDefaults() *LoadbalancerCreateValidateResponseDetail`

NewLoadbalancerCreateValidateResponseDetailWithDefaults instantiates a new LoadbalancerCreateValidateResponseDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *LoadbalancerCreateValidateResponseDetail) GetIp() bool`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LoadbalancerCreateValidateResponseDetail) GetIpOk() (*bool, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LoadbalancerCreateValidateResponseDetail) SetIp(v bool)`

SetIp sets Ip field to given value.


### GetResult

`func (o *LoadbalancerCreateValidateResponseDetail) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *LoadbalancerCreateValidateResponseDetail) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *LoadbalancerCreateValidateResponseDetail) SetResult(v bool)`

SetResult sets Result field to given value.


### GetSubnet

`func (o *LoadbalancerCreateValidateResponseDetail) GetSubnet() bool`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *LoadbalancerCreateValidateResponseDetail) GetSubnetOk() (*bool, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *LoadbalancerCreateValidateResponseDetail) SetSubnet(v bool)`

SetSubnet sets Subnet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


