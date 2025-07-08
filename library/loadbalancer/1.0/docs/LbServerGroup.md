# LbServerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | LB Server Group ID | 
**LbHealthCheckId** | **NullableString** |  | 
**LbMethod** | [**LbServerGroupLbMethod**](LbServerGroupLbMethod.md) | LB Method | 
**LbName** | Pointer to **NullableString** |  | [optional] 
**LoadbalancerId** | Pointer to **NullableString** |  | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | LB Server Group name | 
**Protocol** | [**LbServerGroupProtocol**](LbServerGroupProtocol.md) | Protocol | 
**State** | **string** | LB Server Group state | 
**SubnetId** | **string** | Service Subnet ID | 
**VpcId** | **string** | VPC ID | 

## Methods

### NewLbServerGroup

`func NewLbServerGroup(accountId string, createdAt time.Time, createdBy string, id string, lbHealthCheckId NullableString, lbMethod LbServerGroupLbMethod, modifiedAt time.Time, modifiedBy string, name string, protocol LbServerGroupProtocol, state string, subnetId string, vpcId string, ) *LbServerGroup`

NewLbServerGroup instantiates a new LbServerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbServerGroupWithDefaults

`func NewLbServerGroupWithDefaults() *LbServerGroup`

NewLbServerGroupWithDefaults instantiates a new LbServerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LbServerGroup) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LbServerGroup) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LbServerGroup) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *LbServerGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LbServerGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LbServerGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *LbServerGroup) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LbServerGroup) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LbServerGroup) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *LbServerGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LbServerGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LbServerGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LbServerGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LbServerGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LbServerGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *LbServerGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LbServerGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LbServerGroup) SetId(v string)`

SetId sets Id field to given value.


### GetLbHealthCheckId

`func (o *LbServerGroup) GetLbHealthCheckId() string`

GetLbHealthCheckId returns the LbHealthCheckId field if non-nil, zero value otherwise.

### GetLbHealthCheckIdOk

`func (o *LbServerGroup) GetLbHealthCheckIdOk() (*string, bool)`

GetLbHealthCheckIdOk returns a tuple with the LbHealthCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbHealthCheckId

`func (o *LbServerGroup) SetLbHealthCheckId(v string)`

SetLbHealthCheckId sets LbHealthCheckId field to given value.


### SetLbHealthCheckIdNil

`func (o *LbServerGroup) SetLbHealthCheckIdNil(b bool)`

 SetLbHealthCheckIdNil sets the value for LbHealthCheckId to be an explicit nil

### UnsetLbHealthCheckId
`func (o *LbServerGroup) UnsetLbHealthCheckId()`

UnsetLbHealthCheckId ensures that no value is present for LbHealthCheckId, not even an explicit nil
### GetLbMethod

`func (o *LbServerGroup) GetLbMethod() LbServerGroupLbMethod`

GetLbMethod returns the LbMethod field if non-nil, zero value otherwise.

### GetLbMethodOk

`func (o *LbServerGroup) GetLbMethodOk() (*LbServerGroupLbMethod, bool)`

GetLbMethodOk returns a tuple with the LbMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbMethod

`func (o *LbServerGroup) SetLbMethod(v LbServerGroupLbMethod)`

SetLbMethod sets LbMethod field to given value.


### GetLbName

`func (o *LbServerGroup) GetLbName() string`

GetLbName returns the LbName field if non-nil, zero value otherwise.

### GetLbNameOk

`func (o *LbServerGroup) GetLbNameOk() (*string, bool)`

GetLbNameOk returns a tuple with the LbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbName

`func (o *LbServerGroup) SetLbName(v string)`

SetLbName sets LbName field to given value.

### HasLbName

`func (o *LbServerGroup) HasLbName() bool`

HasLbName returns a boolean if a field has been set.

### SetLbNameNil

`func (o *LbServerGroup) SetLbNameNil(b bool)`

 SetLbNameNil sets the value for LbName to be an explicit nil

### UnsetLbName
`func (o *LbServerGroup) UnsetLbName()`

UnsetLbName ensures that no value is present for LbName, not even an explicit nil
### GetLoadbalancerId

`func (o *LbServerGroup) GetLoadbalancerId() string`

GetLoadbalancerId returns the LoadbalancerId field if non-nil, zero value otherwise.

### GetLoadbalancerIdOk

`func (o *LbServerGroup) GetLoadbalancerIdOk() (*string, bool)`

GetLoadbalancerIdOk returns a tuple with the LoadbalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadbalancerId

`func (o *LbServerGroup) SetLoadbalancerId(v string)`

SetLoadbalancerId sets LoadbalancerId field to given value.

### HasLoadbalancerId

`func (o *LbServerGroup) HasLoadbalancerId() bool`

HasLoadbalancerId returns a boolean if a field has been set.

### SetLoadbalancerIdNil

`func (o *LbServerGroup) SetLoadbalancerIdNil(b bool)`

 SetLoadbalancerIdNil sets the value for LoadbalancerId to be an explicit nil

### UnsetLoadbalancerId
`func (o *LbServerGroup) UnsetLoadbalancerId()`

UnsetLoadbalancerId ensures that no value is present for LoadbalancerId, not even an explicit nil
### GetModifiedAt

`func (o *LbServerGroup) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LbServerGroup) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LbServerGroup) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *LbServerGroup) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *LbServerGroup) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *LbServerGroup) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *LbServerGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LbServerGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LbServerGroup) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *LbServerGroup) GetProtocol() LbServerGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LbServerGroup) GetProtocolOk() (*LbServerGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LbServerGroup) SetProtocol(v LbServerGroupProtocol)`

SetProtocol sets Protocol field to given value.


### GetState

`func (o *LbServerGroup) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LbServerGroup) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LbServerGroup) SetState(v string)`

SetState sets State field to given value.


### GetSubnetId

`func (o *LbServerGroup) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *LbServerGroup) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *LbServerGroup) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetVpcId

`func (o *LbServerGroup) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *LbServerGroup) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *LbServerGroup) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


