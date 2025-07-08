# FirewallRuleSetOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderDirection** | [**FirewallRuleOrderDirection**](FirewallRuleOrderDirection.md) | Order Direction | 
**OrderRuleId** | Pointer to **string** | Order Rule ID | [optional] 

## Methods

### NewFirewallRuleSetOrderRequest

`func NewFirewallRuleSetOrderRequest(orderDirection FirewallRuleOrderDirection, ) *FirewallRuleSetOrderRequest`

NewFirewallRuleSetOrderRequest instantiates a new FirewallRuleSetOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleSetOrderRequestWithDefaults

`func NewFirewallRuleSetOrderRequestWithDefaults() *FirewallRuleSetOrderRequest`

NewFirewallRuleSetOrderRequestWithDefaults instantiates a new FirewallRuleSetOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderDirection

`func (o *FirewallRuleSetOrderRequest) GetOrderDirection() FirewallRuleOrderDirection`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *FirewallRuleSetOrderRequest) GetOrderDirectionOk() (*FirewallRuleOrderDirection, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *FirewallRuleSetOrderRequest) SetOrderDirection(v FirewallRuleOrderDirection)`

SetOrderDirection sets OrderDirection field to given value.


### GetOrderRuleId

`func (o *FirewallRuleSetOrderRequest) GetOrderRuleId() string`

GetOrderRuleId returns the OrderRuleId field if non-nil, zero value otherwise.

### GetOrderRuleIdOk

`func (o *FirewallRuleSetOrderRequest) GetOrderRuleIdOk() (*string, bool)`

GetOrderRuleIdOk returns a tuple with the OrderRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRuleId

`func (o *FirewallRuleSetOrderRequest) SetOrderRuleId(v string)`

SetOrderRuleId sets OrderRuleId field to given value.

### HasOrderRuleId

`func (o *FirewallRuleSetOrderRequest) HasOrderRuleId() bool`

HasOrderRuleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


