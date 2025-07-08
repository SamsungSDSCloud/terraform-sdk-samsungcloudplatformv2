# FirewallRuleCreateSingleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirewallId** | **string** | Firewall ID | 
**FirewallRule** | [**FirewallRuleCreateRequest**](FirewallRuleCreateRequest.md) | Firewall Rule List | 

## Methods

### NewFirewallRuleCreateSingleRequest

`func NewFirewallRuleCreateSingleRequest(firewallId string, firewallRule FirewallRuleCreateRequest, ) *FirewallRuleCreateSingleRequest`

NewFirewallRuleCreateSingleRequest instantiates a new FirewallRuleCreateSingleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleCreateSingleRequestWithDefaults

`func NewFirewallRuleCreateSingleRequestWithDefaults() *FirewallRuleCreateSingleRequest`

NewFirewallRuleCreateSingleRequestWithDefaults instantiates a new FirewallRuleCreateSingleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewallId

`func (o *FirewallRuleCreateSingleRequest) GetFirewallId() string`

GetFirewallId returns the FirewallId field if non-nil, zero value otherwise.

### GetFirewallIdOk

`func (o *FirewallRuleCreateSingleRequest) GetFirewallIdOk() (*string, bool)`

GetFirewallIdOk returns a tuple with the FirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallId

`func (o *FirewallRuleCreateSingleRequest) SetFirewallId(v string)`

SetFirewallId sets FirewallId field to given value.


### GetFirewallRule

`func (o *FirewallRuleCreateSingleRequest) GetFirewallRule() FirewallRuleCreateRequest`

GetFirewallRule returns the FirewallRule field if non-nil, zero value otherwise.

### GetFirewallRuleOk

`func (o *FirewallRuleCreateSingleRequest) GetFirewallRuleOk() (*FirewallRuleCreateRequest, bool)`

GetFirewallRuleOk returns a tuple with the FirewallRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRule

`func (o *FirewallRuleCreateSingleRequest) SetFirewallRule(v FirewallRuleCreateRequest)`

SetFirewallRule sets FirewallRule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


