# TransitGatewayRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | **string** | Description | 
**DestinationCidr** | **string** | Destination CIDR | 
**DestinationResourceId** | **NullableString** |  | 
**DestinationResourceName** | Pointer to **NullableString** |  | [optional] 
**DestinationType** | [**TransitGatewayRuleDestinationType**](TransitGatewayRuleDestinationType.md) | Destination Type | 
**Id** | **string** | Routing Rule ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**SourceResourceId** | **NullableString** |  | 
**SourceResourceName** | Pointer to **NullableString** |  | [optional] 
**SourceType** | [**TransitGatewayRuleDestinationType**](TransitGatewayRuleDestinationType.md) | Source Type | 
**State** | [**RoutingRuleState**](RoutingRuleState.md) | State | 
**TgwConnectionVpcId** | **NullableString** |  | 
**TgwConnectionVpcName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTransitGatewayRule

`func NewTransitGatewayRule(accountId string, createdAt time.Time, createdBy string, description string, destinationCidr string, destinationResourceId NullableString, destinationType TransitGatewayRuleDestinationType, id string, modifiedAt time.Time, modifiedBy string, sourceResourceId NullableString, sourceType TransitGatewayRuleDestinationType, state RoutingRuleState, tgwConnectionVpcId NullableString, ) *TransitGatewayRule`

NewTransitGatewayRule instantiates a new TransitGatewayRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayRuleWithDefaults

`func NewTransitGatewayRuleWithDefaults() *TransitGatewayRule`

NewTransitGatewayRuleWithDefaults instantiates a new TransitGatewayRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TransitGatewayRule) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransitGatewayRule) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransitGatewayRule) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *TransitGatewayRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransitGatewayRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransitGatewayRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *TransitGatewayRule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TransitGatewayRule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TransitGatewayRule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *TransitGatewayRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransitGatewayRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransitGatewayRule) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationCidr

`func (o *TransitGatewayRule) GetDestinationCidr() string`

GetDestinationCidr returns the DestinationCidr field if non-nil, zero value otherwise.

### GetDestinationCidrOk

`func (o *TransitGatewayRule) GetDestinationCidrOk() (*string, bool)`

GetDestinationCidrOk returns a tuple with the DestinationCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidr

`func (o *TransitGatewayRule) SetDestinationCidr(v string)`

SetDestinationCidr sets DestinationCidr field to given value.


### GetDestinationResourceId

`func (o *TransitGatewayRule) GetDestinationResourceId() string`

GetDestinationResourceId returns the DestinationResourceId field if non-nil, zero value otherwise.

### GetDestinationResourceIdOk

`func (o *TransitGatewayRule) GetDestinationResourceIdOk() (*string, bool)`

GetDestinationResourceIdOk returns a tuple with the DestinationResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationResourceId

`func (o *TransitGatewayRule) SetDestinationResourceId(v string)`

SetDestinationResourceId sets DestinationResourceId field to given value.


### SetDestinationResourceIdNil

`func (o *TransitGatewayRule) SetDestinationResourceIdNil(b bool)`

 SetDestinationResourceIdNil sets the value for DestinationResourceId to be an explicit nil

### UnsetDestinationResourceId
`func (o *TransitGatewayRule) UnsetDestinationResourceId()`

UnsetDestinationResourceId ensures that no value is present for DestinationResourceId, not even an explicit nil
### GetDestinationResourceName

`func (o *TransitGatewayRule) GetDestinationResourceName() string`

GetDestinationResourceName returns the DestinationResourceName field if non-nil, zero value otherwise.

### GetDestinationResourceNameOk

`func (o *TransitGatewayRule) GetDestinationResourceNameOk() (*string, bool)`

GetDestinationResourceNameOk returns a tuple with the DestinationResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationResourceName

`func (o *TransitGatewayRule) SetDestinationResourceName(v string)`

SetDestinationResourceName sets DestinationResourceName field to given value.

### HasDestinationResourceName

`func (o *TransitGatewayRule) HasDestinationResourceName() bool`

HasDestinationResourceName returns a boolean if a field has been set.

### SetDestinationResourceNameNil

`func (o *TransitGatewayRule) SetDestinationResourceNameNil(b bool)`

 SetDestinationResourceNameNil sets the value for DestinationResourceName to be an explicit nil

### UnsetDestinationResourceName
`func (o *TransitGatewayRule) UnsetDestinationResourceName()`

UnsetDestinationResourceName ensures that no value is present for DestinationResourceName, not even an explicit nil
### GetDestinationType

`func (o *TransitGatewayRule) GetDestinationType() TransitGatewayRuleDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransitGatewayRule) GetDestinationTypeOk() (*TransitGatewayRuleDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransitGatewayRule) SetDestinationType(v TransitGatewayRuleDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetId

`func (o *TransitGatewayRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransitGatewayRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransitGatewayRule) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *TransitGatewayRule) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *TransitGatewayRule) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *TransitGatewayRule) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *TransitGatewayRule) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *TransitGatewayRule) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *TransitGatewayRule) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetSourceResourceId

`func (o *TransitGatewayRule) GetSourceResourceId() string`

GetSourceResourceId returns the SourceResourceId field if non-nil, zero value otherwise.

### GetSourceResourceIdOk

`func (o *TransitGatewayRule) GetSourceResourceIdOk() (*string, bool)`

GetSourceResourceIdOk returns a tuple with the SourceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceResourceId

`func (o *TransitGatewayRule) SetSourceResourceId(v string)`

SetSourceResourceId sets SourceResourceId field to given value.


### SetSourceResourceIdNil

`func (o *TransitGatewayRule) SetSourceResourceIdNil(b bool)`

 SetSourceResourceIdNil sets the value for SourceResourceId to be an explicit nil

### UnsetSourceResourceId
`func (o *TransitGatewayRule) UnsetSourceResourceId()`

UnsetSourceResourceId ensures that no value is present for SourceResourceId, not even an explicit nil
### GetSourceResourceName

`func (o *TransitGatewayRule) GetSourceResourceName() string`

GetSourceResourceName returns the SourceResourceName field if non-nil, zero value otherwise.

### GetSourceResourceNameOk

`func (o *TransitGatewayRule) GetSourceResourceNameOk() (*string, bool)`

GetSourceResourceNameOk returns a tuple with the SourceResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceResourceName

`func (o *TransitGatewayRule) SetSourceResourceName(v string)`

SetSourceResourceName sets SourceResourceName field to given value.

### HasSourceResourceName

`func (o *TransitGatewayRule) HasSourceResourceName() bool`

HasSourceResourceName returns a boolean if a field has been set.

### SetSourceResourceNameNil

`func (o *TransitGatewayRule) SetSourceResourceNameNil(b bool)`

 SetSourceResourceNameNil sets the value for SourceResourceName to be an explicit nil

### UnsetSourceResourceName
`func (o *TransitGatewayRule) UnsetSourceResourceName()`

UnsetSourceResourceName ensures that no value is present for SourceResourceName, not even an explicit nil
### GetSourceType

`func (o *TransitGatewayRule) GetSourceType() TransitGatewayRuleDestinationType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransitGatewayRule) GetSourceTypeOk() (*TransitGatewayRuleDestinationType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransitGatewayRule) SetSourceType(v TransitGatewayRuleDestinationType)`

SetSourceType sets SourceType field to given value.


### GetState

`func (o *TransitGatewayRule) GetState() RoutingRuleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TransitGatewayRule) GetStateOk() (*RoutingRuleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TransitGatewayRule) SetState(v RoutingRuleState)`

SetState sets State field to given value.


### GetTgwConnectionVpcId

`func (o *TransitGatewayRule) GetTgwConnectionVpcId() string`

GetTgwConnectionVpcId returns the TgwConnectionVpcId field if non-nil, zero value otherwise.

### GetTgwConnectionVpcIdOk

`func (o *TransitGatewayRule) GetTgwConnectionVpcIdOk() (*string, bool)`

GetTgwConnectionVpcIdOk returns a tuple with the TgwConnectionVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwConnectionVpcId

`func (o *TransitGatewayRule) SetTgwConnectionVpcId(v string)`

SetTgwConnectionVpcId sets TgwConnectionVpcId field to given value.


### SetTgwConnectionVpcIdNil

`func (o *TransitGatewayRule) SetTgwConnectionVpcIdNil(b bool)`

 SetTgwConnectionVpcIdNil sets the value for TgwConnectionVpcId to be an explicit nil

### UnsetTgwConnectionVpcId
`func (o *TransitGatewayRule) UnsetTgwConnectionVpcId()`

UnsetTgwConnectionVpcId ensures that no value is present for TgwConnectionVpcId, not even an explicit nil
### GetTgwConnectionVpcName

`func (o *TransitGatewayRule) GetTgwConnectionVpcName() string`

GetTgwConnectionVpcName returns the TgwConnectionVpcName field if non-nil, zero value otherwise.

### GetTgwConnectionVpcNameOk

`func (o *TransitGatewayRule) GetTgwConnectionVpcNameOk() (*string, bool)`

GetTgwConnectionVpcNameOk returns a tuple with the TgwConnectionVpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwConnectionVpcName

`func (o *TransitGatewayRule) SetTgwConnectionVpcName(v string)`

SetTgwConnectionVpcName sets TgwConnectionVpcName field to given value.

### HasTgwConnectionVpcName

`func (o *TransitGatewayRule) HasTgwConnectionVpcName() bool`

HasTgwConnectionVpcName returns a boolean if a field has been set.

### SetTgwConnectionVpcNameNil

`func (o *TransitGatewayRule) SetTgwConnectionVpcNameNil(b bool)`

 SetTgwConnectionVpcNameNil sets the value for TgwConnectionVpcName to be an explicit nil

### UnsetTgwConnectionVpcName
`func (o *TransitGatewayRule) UnsetTgwConnectionVpcName()`

UnsetTgwConnectionVpcName ensures that no value is present for TgwConnectionVpcName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


