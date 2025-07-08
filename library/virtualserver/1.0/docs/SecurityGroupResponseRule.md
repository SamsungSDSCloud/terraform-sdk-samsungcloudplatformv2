# SecurityGroupResponseRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPort** | **int32** | From port rule | 
**Id** | **string** | Security group rule ID | 
**IpProtocol** | **string** | Ip protocol | 
**IpRange** | [**SecurityGroupResponseRuleIpRange**](SecurityGroupResponseRuleIpRange.md) | Ip range | 
**ToPort** | **int32** | To port rule | 

## Methods

### NewSecurityGroupResponseRule

`func NewSecurityGroupResponseRule(fromPort int32, id string, ipProtocol string, ipRange SecurityGroupResponseRuleIpRange, toPort int32, ) *SecurityGroupResponseRule`

NewSecurityGroupResponseRule instantiates a new SecurityGroupResponseRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupResponseRuleWithDefaults

`func NewSecurityGroupResponseRuleWithDefaults() *SecurityGroupResponseRule`

NewSecurityGroupResponseRuleWithDefaults instantiates a new SecurityGroupResponseRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPort

`func (o *SecurityGroupResponseRule) GetFromPort() int32`

GetFromPort returns the FromPort field if non-nil, zero value otherwise.

### GetFromPortOk

`func (o *SecurityGroupResponseRule) GetFromPortOk() (*int32, bool)`

GetFromPortOk returns a tuple with the FromPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPort

`func (o *SecurityGroupResponseRule) SetFromPort(v int32)`

SetFromPort sets FromPort field to given value.


### GetId

`func (o *SecurityGroupResponseRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupResponseRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupResponseRule) SetId(v string)`

SetId sets Id field to given value.


### GetIpProtocol

`func (o *SecurityGroupResponseRule) GetIpProtocol() string`

GetIpProtocol returns the IpProtocol field if non-nil, zero value otherwise.

### GetIpProtocolOk

`func (o *SecurityGroupResponseRule) GetIpProtocolOk() (*string, bool)`

GetIpProtocolOk returns a tuple with the IpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpProtocol

`func (o *SecurityGroupResponseRule) SetIpProtocol(v string)`

SetIpProtocol sets IpProtocol field to given value.


### GetIpRange

`func (o *SecurityGroupResponseRule) GetIpRange() SecurityGroupResponseRuleIpRange`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *SecurityGroupResponseRule) GetIpRangeOk() (*SecurityGroupResponseRuleIpRange, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *SecurityGroupResponseRule) SetIpRange(v SecurityGroupResponseRuleIpRange)`

SetIpRange sets IpRange field to given value.


### GetToPort

`func (o *SecurityGroupResponseRule) GetToPort() int32`

GetToPort returns the ToPort field if non-nil, zero value otherwise.

### GetToPortOk

`func (o *SecurityGroupResponseRule) GetToPortOk() (*int32, bool)`

GetToPortOk returns a tuple with the ToPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPort

`func (o *SecurityGroupResponseRule) SetToPort(v int32)`

SetToPort sets ToPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


