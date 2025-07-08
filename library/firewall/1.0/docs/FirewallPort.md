# FirewallPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceType** | [**FirewallServiceType**](FirewallServiceType.md) | Service Protocol Type | 
**ServiceValue** | Pointer to **string** | Service Port Value | [optional] 

## Methods

### NewFirewallPort

`func NewFirewallPort(serviceType FirewallServiceType, ) *FirewallPort`

NewFirewallPort instantiates a new FirewallPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallPortWithDefaults

`func NewFirewallPortWithDefaults() *FirewallPort`

NewFirewallPortWithDefaults instantiates a new FirewallPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceType

`func (o *FirewallPort) GetServiceType() FirewallServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *FirewallPort) GetServiceTypeOk() (*FirewallServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *FirewallPort) SetServiceType(v FirewallServiceType)`

SetServiceType sets ServiceType field to given value.


### GetServiceValue

`func (o *FirewallPort) GetServiceValue() string`

GetServiceValue returns the ServiceValue field if non-nil, zero value otherwise.

### GetServiceValueOk

`func (o *FirewallPort) GetServiceValueOk() (*string, bool)`

GetServiceValueOk returns a tuple with the ServiceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceValue

`func (o *FirewallPort) SetServiceValue(v string)`

SetServiceValue sets ServiceValue field to given value.

### HasServiceValue

`func (o *FirewallPort) HasServiceValue() bool`

HasServiceValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


