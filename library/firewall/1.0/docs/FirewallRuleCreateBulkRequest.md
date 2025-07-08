# FirewallRuleCreateBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirewallId** | **string** | Firewall ID | 
**FirewallRuleList** | [**[]FirewallRuleCreateRequest**](FirewallRuleCreateRequest.md) | Firewall Rule List | 

## Methods

### NewFirewallRuleCreateBulkRequest

`func NewFirewallRuleCreateBulkRequest(firewallId string, firewallRuleList []FirewallRuleCreateRequest, ) *FirewallRuleCreateBulkRequest`

NewFirewallRuleCreateBulkRequest instantiates a new FirewallRuleCreateBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleCreateBulkRequestWithDefaults

`func NewFirewallRuleCreateBulkRequestWithDefaults() *FirewallRuleCreateBulkRequest`

NewFirewallRuleCreateBulkRequestWithDefaults instantiates a new FirewallRuleCreateBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewallId

`func (o *FirewallRuleCreateBulkRequest) GetFirewallId() string`

GetFirewallId returns the FirewallId field if non-nil, zero value otherwise.

### GetFirewallIdOk

`func (o *FirewallRuleCreateBulkRequest) GetFirewallIdOk() (*string, bool)`

GetFirewallIdOk returns a tuple with the FirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallId

`func (o *FirewallRuleCreateBulkRequest) SetFirewallId(v string)`

SetFirewallId sets FirewallId field to given value.


### GetFirewallRuleList

`func (o *FirewallRuleCreateBulkRequest) GetFirewallRuleList() []FirewallRuleCreateRequest`

GetFirewallRuleList returns the FirewallRuleList field if non-nil, zero value otherwise.

### GetFirewallRuleListOk

`func (o *FirewallRuleCreateBulkRequest) GetFirewallRuleListOk() (*[]FirewallRuleCreateRequest, bool)`

GetFirewallRuleListOk returns a tuple with the FirewallRuleList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRuleList

`func (o *FirewallRuleCreateBulkRequest) SetFirewallRuleList(v []FirewallRuleCreateRequest)`

SetFirewallRuleList sets FirewallRuleList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


