# FirewallRuleSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**FirewallRuleAction**](FirewallRuleAction.md) | Firewall Rule Action | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DestinationAddress** | **[]string** | Destination Address | 
**Direction** | [**FirewallRuleDirection**](FirewallRuleDirection.md) | Firewall Rule Direction | 
**Service** | [**[]FirewallPort**](FirewallPort.md) | Service Port | 
**SourceAddress** | **[]string** | Source Address | 

## Methods

### NewFirewallRuleSetRequest

`func NewFirewallRuleSetRequest(action FirewallRuleAction, destinationAddress []string, direction FirewallRuleDirection, service []FirewallPort, sourceAddress []string, ) *FirewallRuleSetRequest`

NewFirewallRuleSetRequest instantiates a new FirewallRuleSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleSetRequestWithDefaults

`func NewFirewallRuleSetRequestWithDefaults() *FirewallRuleSetRequest`

NewFirewallRuleSetRequestWithDefaults instantiates a new FirewallRuleSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FirewallRuleSetRequest) GetAction() FirewallRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FirewallRuleSetRequest) GetActionOk() (*FirewallRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FirewallRuleSetRequest) SetAction(v FirewallRuleAction)`

SetAction sets Action field to given value.


### GetDescription

`func (o *FirewallRuleSetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallRuleSetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallRuleSetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallRuleSetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FirewallRuleSetRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FirewallRuleSetRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDestinationAddress

`func (o *FirewallRuleSetRequest) GetDestinationAddress() []string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *FirewallRuleSetRequest) GetDestinationAddressOk() (*[]string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *FirewallRuleSetRequest) SetDestinationAddress(v []string)`

SetDestinationAddress sets DestinationAddress field to given value.


### GetDirection

`func (o *FirewallRuleSetRequest) GetDirection() FirewallRuleDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FirewallRuleSetRequest) GetDirectionOk() (*FirewallRuleDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FirewallRuleSetRequest) SetDirection(v FirewallRuleDirection)`

SetDirection sets Direction field to given value.


### GetService

`func (o *FirewallRuleSetRequest) GetService() []FirewallPort`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *FirewallRuleSetRequest) GetServiceOk() (*[]FirewallPort, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *FirewallRuleSetRequest) SetService(v []FirewallPort)`

SetService sets Service field to given value.


### GetSourceAddress

`func (o *FirewallRuleSetRequest) GetSourceAddress() []string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *FirewallRuleSetRequest) GetSourceAddressOk() (*[]string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *FirewallRuleSetRequest) SetSourceAddress(v []string)`

SetSourceAddress sets SourceAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


