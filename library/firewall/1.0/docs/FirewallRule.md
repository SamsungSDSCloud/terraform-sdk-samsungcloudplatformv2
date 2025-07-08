# FirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**FirewallRuleAction**](FirewallRuleAction.md) | Firewall Rule Action | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DestinationAddress** | **[]string** | Destination Address | 
**DestinationInterface** | **string** | Destination Interface Name | 
**Direction** | [**FirewallRuleDirection**](FirewallRuleDirection.md) | Firewall Rule Direction | 
**FirewallId** | **string** | Firewall ID | 
**Id** | **string** | Firewall Rule ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Sequence** | **int32** | Firewall Rule Sequence | 
**Service** | [**[]FirewallPort**](FirewallPort.md) | Service Port | 
**SourceAddress** | **[]string** | Source Address | 
**SourceInterface** | **string** | Source Interface Name | 
**State** | [**FirewallRuleState**](FirewallRuleState.md) | Firewall Rule State | 
**Status** | [**FirewallStatusType**](FirewallStatusType.md) | Firewall Rule Status | 
**VendorRuleId** | **string** | Vendor Rule ID | 

## Methods

### NewFirewallRule

`func NewFirewallRule(action FirewallRuleAction, createdAt time.Time, createdBy string, destinationAddress []string, destinationInterface string, direction FirewallRuleDirection, firewallId string, id string, modifiedAt time.Time, modifiedBy string, sequence int32, service []FirewallPort, sourceAddress []string, sourceInterface string, state FirewallRuleState, status FirewallStatusType, vendorRuleId string, ) *FirewallRule`

NewFirewallRule instantiates a new FirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleWithDefaults

`func NewFirewallRuleWithDefaults() *FirewallRule`

NewFirewallRuleWithDefaults instantiates a new FirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FirewallRule) GetAction() FirewallRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FirewallRule) GetActionOk() (*FirewallRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FirewallRule) SetAction(v FirewallRuleAction)`

SetAction sets Action field to given value.


### GetCreatedAt

`func (o *FirewallRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirewallRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirewallRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *FirewallRule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FirewallRule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FirewallRule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *FirewallRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FirewallRule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FirewallRule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDestinationAddress

`func (o *FirewallRule) GetDestinationAddress() []string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *FirewallRule) GetDestinationAddressOk() (*[]string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *FirewallRule) SetDestinationAddress(v []string)`

SetDestinationAddress sets DestinationAddress field to given value.


### GetDestinationInterface

`func (o *FirewallRule) GetDestinationInterface() string`

GetDestinationInterface returns the DestinationInterface field if non-nil, zero value otherwise.

### GetDestinationInterfaceOk

`func (o *FirewallRule) GetDestinationInterfaceOk() (*string, bool)`

GetDestinationInterfaceOk returns a tuple with the DestinationInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationInterface

`func (o *FirewallRule) SetDestinationInterface(v string)`

SetDestinationInterface sets DestinationInterface field to given value.


### GetDirection

`func (o *FirewallRule) GetDirection() FirewallRuleDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FirewallRule) GetDirectionOk() (*FirewallRuleDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FirewallRule) SetDirection(v FirewallRuleDirection)`

SetDirection sets Direction field to given value.


### GetFirewallId

`func (o *FirewallRule) GetFirewallId() string`

GetFirewallId returns the FirewallId field if non-nil, zero value otherwise.

### GetFirewallIdOk

`func (o *FirewallRule) GetFirewallIdOk() (*string, bool)`

GetFirewallIdOk returns a tuple with the FirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallId

`func (o *FirewallRule) SetFirewallId(v string)`

SetFirewallId sets FirewallId field to given value.


### GetId

`func (o *FirewallRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallRule) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *FirewallRule) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *FirewallRule) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *FirewallRule) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *FirewallRule) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *FirewallRule) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *FirewallRule) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *FirewallRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallRule) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FirewallRule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FirewallRule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSequence

`func (o *FirewallRule) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *FirewallRule) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *FirewallRule) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetService

`func (o *FirewallRule) GetService() []FirewallPort`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *FirewallRule) GetServiceOk() (*[]FirewallPort, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *FirewallRule) SetService(v []FirewallPort)`

SetService sets Service field to given value.


### GetSourceAddress

`func (o *FirewallRule) GetSourceAddress() []string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *FirewallRule) GetSourceAddressOk() (*[]string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *FirewallRule) SetSourceAddress(v []string)`

SetSourceAddress sets SourceAddress field to given value.


### GetSourceInterface

`func (o *FirewallRule) GetSourceInterface() string`

GetSourceInterface returns the SourceInterface field if non-nil, zero value otherwise.

### GetSourceInterfaceOk

`func (o *FirewallRule) GetSourceInterfaceOk() (*string, bool)`

GetSourceInterfaceOk returns a tuple with the SourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterface

`func (o *FirewallRule) SetSourceInterface(v string)`

SetSourceInterface sets SourceInterface field to given value.


### GetState

`func (o *FirewallRule) GetState() FirewallRuleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FirewallRule) GetStateOk() (*FirewallRuleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FirewallRule) SetState(v FirewallRuleState)`

SetState sets State field to given value.


### GetStatus

`func (o *FirewallRule) GetStatus() FirewallStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallRule) GetStatusOk() (*FirewallStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallRule) SetStatus(v FirewallStatusType)`

SetStatus sets Status field to given value.


### GetVendorRuleId

`func (o *FirewallRule) GetVendorRuleId() string`

GetVendorRuleId returns the VendorRuleId field if non-nil, zero value otherwise.

### GetVendorRuleIdOk

`func (o *FirewallRule) GetVendorRuleIdOk() (*string, bool)`

GetVendorRuleIdOk returns a tuple with the VendorRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorRuleId

`func (o *FirewallRule) SetVendorRuleId(v string)`

SetVendorRuleId sets VendorRuleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


