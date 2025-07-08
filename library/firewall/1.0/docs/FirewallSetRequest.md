# FirewallSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlavorName** | Pointer to [**FirewallFlavorType**](FirewallFlavorType.md) | Firewall Size Name | [optional] 
**Loggable** | Pointer to **bool** | Logging Use | [optional] 

## Methods

### NewFirewallSetRequest

`func NewFirewallSetRequest() *FirewallSetRequest`

NewFirewallSetRequest instantiates a new FirewallSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSetRequestWithDefaults

`func NewFirewallSetRequestWithDefaults() *FirewallSetRequest`

NewFirewallSetRequestWithDefaults instantiates a new FirewallSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavorName

`func (o *FirewallSetRequest) GetFlavorName() FirewallFlavorType`

GetFlavorName returns the FlavorName field if non-nil, zero value otherwise.

### GetFlavorNameOk

`func (o *FirewallSetRequest) GetFlavorNameOk() (*FirewallFlavorType, bool)`

GetFlavorNameOk returns a tuple with the FlavorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorName

`func (o *FirewallSetRequest) SetFlavorName(v FirewallFlavorType)`

SetFlavorName sets FlavorName field to given value.

### HasFlavorName

`func (o *FirewallSetRequest) HasFlavorName() bool`

HasFlavorName returns a boolean if a field has been set.

### GetLoggable

`func (o *FirewallSetRequest) GetLoggable() bool`

GetLoggable returns the Loggable field if non-nil, zero value otherwise.

### GetLoggableOk

`func (o *FirewallSetRequest) GetLoggableOk() (*bool, bool)`

GetLoggableOk returns a tuple with the Loggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggable

`func (o *FirewallSetRequest) SetLoggable(v bool)`

SetLoggable sets Loggable field to given value.

### HasLoggable

`func (o *FirewallSetRequest) HasLoggable() bool`

HasLoggable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


