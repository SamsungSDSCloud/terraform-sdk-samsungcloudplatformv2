# RoutingRule

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
**DestinationType** | [**RoutingRuleDestinationType**](RoutingRuleDestinationType.md) | Destination Type | 
**Id** | **string** | Routing Rule ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**OwnerId** | **string** | Routing Rule Owner ID | 
**OwnerType** | [**RoutingRuleOwnerType**](RoutingRuleOwnerType.md) | Routing Rule Owner Type | 
**State** | [**RoutingRuleState**](RoutingRuleState.md) | State | 

## Methods

### NewRoutingRule

`func NewRoutingRule(accountId string, createdAt time.Time, createdBy string, description string, destinationCidr string, destinationResourceId NullableString, destinationType RoutingRuleDestinationType, id string, modifiedAt time.Time, modifiedBy string, ownerId string, ownerType RoutingRuleOwnerType, state RoutingRuleState, ) *RoutingRule`

NewRoutingRule instantiates a new RoutingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingRuleWithDefaults

`func NewRoutingRuleWithDefaults() *RoutingRule`

NewRoutingRuleWithDefaults instantiates a new RoutingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RoutingRule) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RoutingRule) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RoutingRule) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *RoutingRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoutingRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoutingRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *RoutingRule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RoutingRule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RoutingRule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *RoutingRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingRule) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationCidr

`func (o *RoutingRule) GetDestinationCidr() string`

GetDestinationCidr returns the DestinationCidr field if non-nil, zero value otherwise.

### GetDestinationCidrOk

`func (o *RoutingRule) GetDestinationCidrOk() (*string, bool)`

GetDestinationCidrOk returns a tuple with the DestinationCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidr

`func (o *RoutingRule) SetDestinationCidr(v string)`

SetDestinationCidr sets DestinationCidr field to given value.


### GetDestinationResourceId

`func (o *RoutingRule) GetDestinationResourceId() string`

GetDestinationResourceId returns the DestinationResourceId field if non-nil, zero value otherwise.

### GetDestinationResourceIdOk

`func (o *RoutingRule) GetDestinationResourceIdOk() (*string, bool)`

GetDestinationResourceIdOk returns a tuple with the DestinationResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationResourceId

`func (o *RoutingRule) SetDestinationResourceId(v string)`

SetDestinationResourceId sets DestinationResourceId field to given value.


### SetDestinationResourceIdNil

`func (o *RoutingRule) SetDestinationResourceIdNil(b bool)`

 SetDestinationResourceIdNil sets the value for DestinationResourceId to be an explicit nil

### UnsetDestinationResourceId
`func (o *RoutingRule) UnsetDestinationResourceId()`

UnsetDestinationResourceId ensures that no value is present for DestinationResourceId, not even an explicit nil
### GetDestinationResourceName

`func (o *RoutingRule) GetDestinationResourceName() string`

GetDestinationResourceName returns the DestinationResourceName field if non-nil, zero value otherwise.

### GetDestinationResourceNameOk

`func (o *RoutingRule) GetDestinationResourceNameOk() (*string, bool)`

GetDestinationResourceNameOk returns a tuple with the DestinationResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationResourceName

`func (o *RoutingRule) SetDestinationResourceName(v string)`

SetDestinationResourceName sets DestinationResourceName field to given value.

### HasDestinationResourceName

`func (o *RoutingRule) HasDestinationResourceName() bool`

HasDestinationResourceName returns a boolean if a field has been set.

### SetDestinationResourceNameNil

`func (o *RoutingRule) SetDestinationResourceNameNil(b bool)`

 SetDestinationResourceNameNil sets the value for DestinationResourceName to be an explicit nil

### UnsetDestinationResourceName
`func (o *RoutingRule) UnsetDestinationResourceName()`

UnsetDestinationResourceName ensures that no value is present for DestinationResourceName, not even an explicit nil
### GetDestinationType

`func (o *RoutingRule) GetDestinationType() RoutingRuleDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *RoutingRule) GetDestinationTypeOk() (*RoutingRuleDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *RoutingRule) SetDestinationType(v RoutingRuleDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetId

`func (o *RoutingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoutingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoutingRule) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *RoutingRule) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RoutingRule) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RoutingRule) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *RoutingRule) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *RoutingRule) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *RoutingRule) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetOwnerId

`func (o *RoutingRule) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *RoutingRule) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *RoutingRule) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerType

`func (o *RoutingRule) GetOwnerType() RoutingRuleOwnerType`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *RoutingRule) GetOwnerTypeOk() (*RoutingRuleOwnerType, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *RoutingRule) SetOwnerType(v RoutingRuleOwnerType)`

SetOwnerType sets OwnerType field to given value.


### GetState

`func (o *RoutingRule) GetState() RoutingRuleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RoutingRule) GetStateOk() (*RoutingRuleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RoutingRule) SetState(v RoutingRuleState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


