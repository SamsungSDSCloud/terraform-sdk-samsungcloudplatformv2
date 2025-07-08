# SecurityGroupRuleCreateRequestParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Direction** | **string** | Security Group Rule Direction (Ingress, Egress) | 
**Ethertype** | Pointer to **NullableString** |  | [optional] 
**PortRangeMax** | Pointer to **NullableInt32** |  | [optional] 
**PortRangeMin** | Pointer to **NullableInt32** |  | [optional] 
**Protocol** | **NullableString** |  | 
**RemoteGroupId** | Pointer to **NullableString** |  | [optional] 
**RemoteIpPrefix** | Pointer to **NullableString** |  | [optional] 
**SecurityGroupId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSecurityGroupRuleCreateRequestParam

`func NewSecurityGroupRuleCreateRequestParam(direction string, protocol NullableString, ) *SecurityGroupRuleCreateRequestParam`

NewSecurityGroupRuleCreateRequestParam instantiates a new SecurityGroupRuleCreateRequestParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRuleCreateRequestParamWithDefaults

`func NewSecurityGroupRuleCreateRequestParamWithDefaults() *SecurityGroupRuleCreateRequestParam`

NewSecurityGroupRuleCreateRequestParamWithDefaults instantiates a new SecurityGroupRuleCreateRequestParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SecurityGroupRuleCreateRequestParam) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroupRuleCreateRequestParam) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroupRuleCreateRequestParam) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityGroupRuleCreateRequestParam) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityGroupRuleCreateRequestParam) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityGroupRuleCreateRequestParam) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDirection

`func (o *SecurityGroupRuleCreateRequestParam) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SecurityGroupRuleCreateRequestParam) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SecurityGroupRuleCreateRequestParam) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetEthertype

`func (o *SecurityGroupRuleCreateRequestParam) GetEthertype() string`

GetEthertype returns the Ethertype field if non-nil, zero value otherwise.

### GetEthertypeOk

`func (o *SecurityGroupRuleCreateRequestParam) GetEthertypeOk() (*string, bool)`

GetEthertypeOk returns a tuple with the Ethertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthertype

`func (o *SecurityGroupRuleCreateRequestParam) SetEthertype(v string)`

SetEthertype sets Ethertype field to given value.

### HasEthertype

`func (o *SecurityGroupRuleCreateRequestParam) HasEthertype() bool`

HasEthertype returns a boolean if a field has been set.

### SetEthertypeNil

`func (o *SecurityGroupRuleCreateRequestParam) SetEthertypeNil(b bool)`

 SetEthertypeNil sets the value for Ethertype to be an explicit nil

### UnsetEthertype
`func (o *SecurityGroupRuleCreateRequestParam) UnsetEthertype()`

UnsetEthertype ensures that no value is present for Ethertype, not even an explicit nil
### GetPortRangeMax

`func (o *SecurityGroupRuleCreateRequestParam) GetPortRangeMax() int32`

GetPortRangeMax returns the PortRangeMax field if non-nil, zero value otherwise.

### GetPortRangeMaxOk

`func (o *SecurityGroupRuleCreateRequestParam) GetPortRangeMaxOk() (*int32, bool)`

GetPortRangeMaxOk returns a tuple with the PortRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMax

`func (o *SecurityGroupRuleCreateRequestParam) SetPortRangeMax(v int32)`

SetPortRangeMax sets PortRangeMax field to given value.

### HasPortRangeMax

`func (o *SecurityGroupRuleCreateRequestParam) HasPortRangeMax() bool`

HasPortRangeMax returns a boolean if a field has been set.

### SetPortRangeMaxNil

`func (o *SecurityGroupRuleCreateRequestParam) SetPortRangeMaxNil(b bool)`

 SetPortRangeMaxNil sets the value for PortRangeMax to be an explicit nil

### UnsetPortRangeMax
`func (o *SecurityGroupRuleCreateRequestParam) UnsetPortRangeMax()`

UnsetPortRangeMax ensures that no value is present for PortRangeMax, not even an explicit nil
### GetPortRangeMin

`func (o *SecurityGroupRuleCreateRequestParam) GetPortRangeMin() int32`

GetPortRangeMin returns the PortRangeMin field if non-nil, zero value otherwise.

### GetPortRangeMinOk

`func (o *SecurityGroupRuleCreateRequestParam) GetPortRangeMinOk() (*int32, bool)`

GetPortRangeMinOk returns a tuple with the PortRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMin

`func (o *SecurityGroupRuleCreateRequestParam) SetPortRangeMin(v int32)`

SetPortRangeMin sets PortRangeMin field to given value.

### HasPortRangeMin

`func (o *SecurityGroupRuleCreateRequestParam) HasPortRangeMin() bool`

HasPortRangeMin returns a boolean if a field has been set.

### SetPortRangeMinNil

`func (o *SecurityGroupRuleCreateRequestParam) SetPortRangeMinNil(b bool)`

 SetPortRangeMinNil sets the value for PortRangeMin to be an explicit nil

### UnsetPortRangeMin
`func (o *SecurityGroupRuleCreateRequestParam) UnsetPortRangeMin()`

UnsetPortRangeMin ensures that no value is present for PortRangeMin, not even an explicit nil
### GetProtocol

`func (o *SecurityGroupRuleCreateRequestParam) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityGroupRuleCreateRequestParam) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityGroupRuleCreateRequestParam) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### SetProtocolNil

`func (o *SecurityGroupRuleCreateRequestParam) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *SecurityGroupRuleCreateRequestParam) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetRemoteGroupId

`func (o *SecurityGroupRuleCreateRequestParam) GetRemoteGroupId() string`

GetRemoteGroupId returns the RemoteGroupId field if non-nil, zero value otherwise.

### GetRemoteGroupIdOk

`func (o *SecurityGroupRuleCreateRequestParam) GetRemoteGroupIdOk() (*string, bool)`

GetRemoteGroupIdOk returns a tuple with the RemoteGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGroupId

`func (o *SecurityGroupRuleCreateRequestParam) SetRemoteGroupId(v string)`

SetRemoteGroupId sets RemoteGroupId field to given value.

### HasRemoteGroupId

`func (o *SecurityGroupRuleCreateRequestParam) HasRemoteGroupId() bool`

HasRemoteGroupId returns a boolean if a field has been set.

### SetRemoteGroupIdNil

`func (o *SecurityGroupRuleCreateRequestParam) SetRemoteGroupIdNil(b bool)`

 SetRemoteGroupIdNil sets the value for RemoteGroupId to be an explicit nil

### UnsetRemoteGroupId
`func (o *SecurityGroupRuleCreateRequestParam) UnsetRemoteGroupId()`

UnsetRemoteGroupId ensures that no value is present for RemoteGroupId, not even an explicit nil
### GetRemoteIpPrefix

`func (o *SecurityGroupRuleCreateRequestParam) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *SecurityGroupRuleCreateRequestParam) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *SecurityGroupRuleCreateRequestParam) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.

### HasRemoteIpPrefix

`func (o *SecurityGroupRuleCreateRequestParam) HasRemoteIpPrefix() bool`

HasRemoteIpPrefix returns a boolean if a field has been set.

### SetRemoteIpPrefixNil

`func (o *SecurityGroupRuleCreateRequestParam) SetRemoteIpPrefixNil(b bool)`

 SetRemoteIpPrefixNil sets the value for RemoteIpPrefix to be an explicit nil

### UnsetRemoteIpPrefix
`func (o *SecurityGroupRuleCreateRequestParam) UnsetRemoteIpPrefix()`

UnsetRemoteIpPrefix ensures that no value is present for RemoteIpPrefix, not even an explicit nil
### GetSecurityGroupId

`func (o *SecurityGroupRuleCreateRequestParam) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *SecurityGroupRuleCreateRequestParam) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *SecurityGroupRuleCreateRequestParam) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *SecurityGroupRuleCreateRequestParam) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupIdNil

`func (o *SecurityGroupRuleCreateRequestParam) SetSecurityGroupIdNil(b bool)`

 SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil

### UnsetSecurityGroupId
`func (o *SecurityGroupRuleCreateRequestParam) UnsetSecurityGroupId()`

UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


