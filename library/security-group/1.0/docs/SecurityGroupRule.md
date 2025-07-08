# SecurityGroupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | **NullableString** |  | 
**Direction** | [**Direction**](Direction.md) | Security Group Rule Direction (Ingress, Egress) | 
**Ethertype** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Security Group Rule ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**PortRangeMax** | **NullableInt32** |  | 
**PortRangeMin** | **NullableInt32** |  | 
**Protocol** | **NullableString** |  | 
**RemoteGroupId** | **NullableString** |  | 
**RemoteGroupName** | Pointer to **NullableString** |  | [optional] 
**RemoteIpPrefix** | **NullableString** |  | 
**SecurityGroupId** | **string** | Security Group ID | 

## Methods

### NewSecurityGroupRule

`func NewSecurityGroupRule(createdAt time.Time, createdBy string, description NullableString, direction Direction, id string, modifiedAt time.Time, modifiedBy string, portRangeMax NullableInt32, portRangeMin NullableInt32, protocol NullableString, remoteGroupId NullableString, remoteIpPrefix NullableString, securityGroupId string, ) *SecurityGroupRule`

NewSecurityGroupRule instantiates a new SecurityGroupRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRuleWithDefaults

`func NewSecurityGroupRuleWithDefaults() *SecurityGroupRule`

NewSecurityGroupRuleWithDefaults instantiates a new SecurityGroupRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SecurityGroupRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityGroupRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityGroupRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SecurityGroupRule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SecurityGroupRule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SecurityGroupRule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *SecurityGroupRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroupRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroupRule) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *SecurityGroupRule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityGroupRule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDirection

`func (o *SecurityGroupRule) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SecurityGroupRule) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SecurityGroupRule) SetDirection(v Direction)`

SetDirection sets Direction field to given value.


### GetEthertype

`func (o *SecurityGroupRule) GetEthertype() string`

GetEthertype returns the Ethertype field if non-nil, zero value otherwise.

### GetEthertypeOk

`func (o *SecurityGroupRule) GetEthertypeOk() (*string, bool)`

GetEthertypeOk returns a tuple with the Ethertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthertype

`func (o *SecurityGroupRule) SetEthertype(v string)`

SetEthertype sets Ethertype field to given value.

### HasEthertype

`func (o *SecurityGroupRule) HasEthertype() bool`

HasEthertype returns a boolean if a field has been set.

### SetEthertypeNil

`func (o *SecurityGroupRule) SetEthertypeNil(b bool)`

 SetEthertypeNil sets the value for Ethertype to be an explicit nil

### UnsetEthertype
`func (o *SecurityGroupRule) UnsetEthertype()`

UnsetEthertype ensures that no value is present for Ethertype, not even an explicit nil
### GetId

`func (o *SecurityGroupRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupRule) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *SecurityGroupRule) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SecurityGroupRule) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SecurityGroupRule) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *SecurityGroupRule) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SecurityGroupRule) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *SecurityGroupRule) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetPortRangeMax

`func (o *SecurityGroupRule) GetPortRangeMax() int32`

GetPortRangeMax returns the PortRangeMax field if non-nil, zero value otherwise.

### GetPortRangeMaxOk

`func (o *SecurityGroupRule) GetPortRangeMaxOk() (*int32, bool)`

GetPortRangeMaxOk returns a tuple with the PortRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMax

`func (o *SecurityGroupRule) SetPortRangeMax(v int32)`

SetPortRangeMax sets PortRangeMax field to given value.


### SetPortRangeMaxNil

`func (o *SecurityGroupRule) SetPortRangeMaxNil(b bool)`

 SetPortRangeMaxNil sets the value for PortRangeMax to be an explicit nil

### UnsetPortRangeMax
`func (o *SecurityGroupRule) UnsetPortRangeMax()`

UnsetPortRangeMax ensures that no value is present for PortRangeMax, not even an explicit nil
### GetPortRangeMin

`func (o *SecurityGroupRule) GetPortRangeMin() int32`

GetPortRangeMin returns the PortRangeMin field if non-nil, zero value otherwise.

### GetPortRangeMinOk

`func (o *SecurityGroupRule) GetPortRangeMinOk() (*int32, bool)`

GetPortRangeMinOk returns a tuple with the PortRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMin

`func (o *SecurityGroupRule) SetPortRangeMin(v int32)`

SetPortRangeMin sets PortRangeMin field to given value.


### SetPortRangeMinNil

`func (o *SecurityGroupRule) SetPortRangeMinNil(b bool)`

 SetPortRangeMinNil sets the value for PortRangeMin to be an explicit nil

### UnsetPortRangeMin
`func (o *SecurityGroupRule) UnsetPortRangeMin()`

UnsetPortRangeMin ensures that no value is present for PortRangeMin, not even an explicit nil
### GetProtocol

`func (o *SecurityGroupRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityGroupRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityGroupRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### SetProtocolNil

`func (o *SecurityGroupRule) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *SecurityGroupRule) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetRemoteGroupId

`func (o *SecurityGroupRule) GetRemoteGroupId() string`

GetRemoteGroupId returns the RemoteGroupId field if non-nil, zero value otherwise.

### GetRemoteGroupIdOk

`func (o *SecurityGroupRule) GetRemoteGroupIdOk() (*string, bool)`

GetRemoteGroupIdOk returns a tuple with the RemoteGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGroupId

`func (o *SecurityGroupRule) SetRemoteGroupId(v string)`

SetRemoteGroupId sets RemoteGroupId field to given value.


### SetRemoteGroupIdNil

`func (o *SecurityGroupRule) SetRemoteGroupIdNil(b bool)`

 SetRemoteGroupIdNil sets the value for RemoteGroupId to be an explicit nil

### UnsetRemoteGroupId
`func (o *SecurityGroupRule) UnsetRemoteGroupId()`

UnsetRemoteGroupId ensures that no value is present for RemoteGroupId, not even an explicit nil
### GetRemoteGroupName

`func (o *SecurityGroupRule) GetRemoteGroupName() string`

GetRemoteGroupName returns the RemoteGroupName field if non-nil, zero value otherwise.

### GetRemoteGroupNameOk

`func (o *SecurityGroupRule) GetRemoteGroupNameOk() (*string, bool)`

GetRemoteGroupNameOk returns a tuple with the RemoteGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGroupName

`func (o *SecurityGroupRule) SetRemoteGroupName(v string)`

SetRemoteGroupName sets RemoteGroupName field to given value.

### HasRemoteGroupName

`func (o *SecurityGroupRule) HasRemoteGroupName() bool`

HasRemoteGroupName returns a boolean if a field has been set.

### SetRemoteGroupNameNil

`func (o *SecurityGroupRule) SetRemoteGroupNameNil(b bool)`

 SetRemoteGroupNameNil sets the value for RemoteGroupName to be an explicit nil

### UnsetRemoteGroupName
`func (o *SecurityGroupRule) UnsetRemoteGroupName()`

UnsetRemoteGroupName ensures that no value is present for RemoteGroupName, not even an explicit nil
### GetRemoteIpPrefix

`func (o *SecurityGroupRule) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *SecurityGroupRule) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *SecurityGroupRule) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.


### SetRemoteIpPrefixNil

`func (o *SecurityGroupRule) SetRemoteIpPrefixNil(b bool)`

 SetRemoteIpPrefixNil sets the value for RemoteIpPrefix to be an explicit nil

### UnsetRemoteIpPrefix
`func (o *SecurityGroupRule) UnsetRemoteIpPrefix()`

UnsetRemoteIpPrefix ensures that no value is present for RemoteIpPrefix, not even an explicit nil
### GetSecurityGroupId

`func (o *SecurityGroupRule) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *SecurityGroupRule) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *SecurityGroupRule) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


