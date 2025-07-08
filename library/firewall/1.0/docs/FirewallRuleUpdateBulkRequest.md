# FirewallRuleUpdateBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirewallRuleId** | **[]string** | Firewall Rule ID | 
**Status** | [**FirewallStatusType**](FirewallStatusType.md) | Firewall Rule Status | 

## Methods

### NewFirewallRuleUpdateBulkRequest

`func NewFirewallRuleUpdateBulkRequest(firewallRuleId []string, status FirewallStatusType, ) *FirewallRuleUpdateBulkRequest`

NewFirewallRuleUpdateBulkRequest instantiates a new FirewallRuleUpdateBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleUpdateBulkRequestWithDefaults

`func NewFirewallRuleUpdateBulkRequestWithDefaults() *FirewallRuleUpdateBulkRequest`

NewFirewallRuleUpdateBulkRequestWithDefaults instantiates a new FirewallRuleUpdateBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewallRuleId

`func (o *FirewallRuleUpdateBulkRequest) GetFirewallRuleId() []string`

GetFirewallRuleId returns the FirewallRuleId field if non-nil, zero value otherwise.

### GetFirewallRuleIdOk

`func (o *FirewallRuleUpdateBulkRequest) GetFirewallRuleIdOk() (*[]string, bool)`

GetFirewallRuleIdOk returns a tuple with the FirewallRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRuleId

`func (o *FirewallRuleUpdateBulkRequest) SetFirewallRuleId(v []string)`

SetFirewallRuleId sets FirewallRuleId field to given value.


### GetStatus

`func (o *FirewallRuleUpdateBulkRequest) GetStatus() FirewallStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallRuleUpdateBulkRequest) GetStatusOk() (*FirewallStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallRuleUpdateBulkRequest) SetStatus(v FirewallStatusType)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


