# FirewallRuleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**FirewallRuleAction**](FirewallRuleAction.md) | Firewall Rule Action | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DestinationAddress** | **[]string** | Destination Address | 
**Direction** | [**FirewallRuleDirection**](FirewallRuleDirection.md) | Firewall Rule Direction | 
**OrderDirection** | Pointer to [**FirewallRuleOrderDirection**](FirewallRuleOrderDirection.md) | Order Direction | [optional] 
**OrderRuleId** | Pointer to **string** | Order Rule ID | [optional] 
**Service** | [**[]FirewallPort**](FirewallPort.md) | Service Port | 
**SourceAddress** | **[]string** | Source Address | 
**Status** | [**FirewallStatusType**](FirewallStatusType.md) | Firewall Rule Status | 

## Methods

### NewFirewallRuleCreateRequest

`func NewFirewallRuleCreateRequest(action FirewallRuleAction, destinationAddress []string, direction FirewallRuleDirection, service []FirewallPort, sourceAddress []string, status FirewallStatusType, ) *FirewallRuleCreateRequest`

NewFirewallRuleCreateRequest instantiates a new FirewallRuleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleCreateRequestWithDefaults

`func NewFirewallRuleCreateRequestWithDefaults() *FirewallRuleCreateRequest`

NewFirewallRuleCreateRequestWithDefaults instantiates a new FirewallRuleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FirewallRuleCreateRequest) GetAction() FirewallRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FirewallRuleCreateRequest) GetActionOk() (*FirewallRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FirewallRuleCreateRequest) SetAction(v FirewallRuleAction)`

SetAction sets Action field to given value.


### GetDescription

`func (o *FirewallRuleCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallRuleCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallRuleCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallRuleCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FirewallRuleCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FirewallRuleCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDestinationAddress

`func (o *FirewallRuleCreateRequest) GetDestinationAddress() []string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *FirewallRuleCreateRequest) GetDestinationAddressOk() (*[]string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *FirewallRuleCreateRequest) SetDestinationAddress(v []string)`

SetDestinationAddress sets DestinationAddress field to given value.


### GetDirection

`func (o *FirewallRuleCreateRequest) GetDirection() FirewallRuleDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FirewallRuleCreateRequest) GetDirectionOk() (*FirewallRuleDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FirewallRuleCreateRequest) SetDirection(v FirewallRuleDirection)`

SetDirection sets Direction field to given value.


### GetOrderDirection

`func (o *FirewallRuleCreateRequest) GetOrderDirection() FirewallRuleOrderDirection`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *FirewallRuleCreateRequest) GetOrderDirectionOk() (*FirewallRuleOrderDirection, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *FirewallRuleCreateRequest) SetOrderDirection(v FirewallRuleOrderDirection)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *FirewallRuleCreateRequest) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetOrderRuleId

`func (o *FirewallRuleCreateRequest) GetOrderRuleId() string`

GetOrderRuleId returns the OrderRuleId field if non-nil, zero value otherwise.

### GetOrderRuleIdOk

`func (o *FirewallRuleCreateRequest) GetOrderRuleIdOk() (*string, bool)`

GetOrderRuleIdOk returns a tuple with the OrderRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRuleId

`func (o *FirewallRuleCreateRequest) SetOrderRuleId(v string)`

SetOrderRuleId sets OrderRuleId field to given value.

### HasOrderRuleId

`func (o *FirewallRuleCreateRequest) HasOrderRuleId() bool`

HasOrderRuleId returns a boolean if a field has been set.

### GetService

`func (o *FirewallRuleCreateRequest) GetService() []FirewallPort`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *FirewallRuleCreateRequest) GetServiceOk() (*[]FirewallPort, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *FirewallRuleCreateRequest) SetService(v []FirewallPort)`

SetService sets Service field to given value.


### GetSourceAddress

`func (o *FirewallRuleCreateRequest) GetSourceAddress() []string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *FirewallRuleCreateRequest) GetSourceAddressOk() (*[]string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *FirewallRuleCreateRequest) SetSourceAddress(v []string)`

SetSourceAddress sets SourceAddress field to given value.


### GetStatus

`func (o *FirewallRuleCreateRequest) GetStatus() FirewallStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallRuleCreateRequest) GetStatusOk() (*FirewallStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallRuleCreateRequest) SetStatus(v FirewallStatusType)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


