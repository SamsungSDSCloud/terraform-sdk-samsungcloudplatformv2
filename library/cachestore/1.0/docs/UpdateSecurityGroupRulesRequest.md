# UpdateSecurityGroupRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**DelIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]

## Methods

### NewUpdateSecurityGroupRulesRequest

`func NewUpdateSecurityGroupRulesRequest() *UpdateSecurityGroupRulesRequest`

NewUpdateSecurityGroupRulesRequest instantiates a new UpdateSecurityGroupRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecurityGroupRulesRequestWithDefaults

`func NewUpdateSecurityGroupRulesRequestWithDefaults() *UpdateSecurityGroupRulesRequest`

NewUpdateSecurityGroupRulesRequestWithDefaults instantiates a new UpdateSecurityGroupRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddIpAddresses

`func (o *UpdateSecurityGroupRulesRequest) GetAddIpAddresses() []string`

GetAddIpAddresses returns the AddIpAddresses field if non-nil, zero value otherwise.

### GetAddIpAddressesOk

`func (o *UpdateSecurityGroupRulesRequest) GetAddIpAddressesOk() (*[]string, bool)`

GetAddIpAddressesOk returns a tuple with the AddIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddIpAddresses

`func (o *UpdateSecurityGroupRulesRequest) SetAddIpAddresses(v []string)`

SetAddIpAddresses sets AddIpAddresses field to given value.

### HasAddIpAddresses

`func (o *UpdateSecurityGroupRulesRequest) HasAddIpAddresses() bool`

HasAddIpAddresses returns a boolean if a field has been set.

### GetDelIpAddresses

`func (o *UpdateSecurityGroupRulesRequest) GetDelIpAddresses() []string`

GetDelIpAddresses returns the DelIpAddresses field if non-nil, zero value otherwise.

### GetDelIpAddressesOk

`func (o *UpdateSecurityGroupRulesRequest) GetDelIpAddressesOk() (*[]string, bool)`

GetDelIpAddressesOk returns a tuple with the DelIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelIpAddresses

`func (o *UpdateSecurityGroupRulesRequest) SetDelIpAddresses(v []string)`

SetDelIpAddresses sets DelIpAddresses field to given value.

### HasDelIpAddresses

`func (o *UpdateSecurityGroupRulesRequest) HasDelIpAddresses() bool`

HasDelIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


