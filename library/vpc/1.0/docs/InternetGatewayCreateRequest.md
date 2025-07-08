# InternetGatewayCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**FirewallEnabled** | Pointer to **bool** | Firewall Enable | [optional] [default to false]
**FirewallLoggable** | Pointer to **bool** | Firewall Loggable | [optional] [default to false]
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]
**Type** | [**InternetGatewayType**](InternetGatewayType.md) | Internet Gateway Type | 
**VpcId** | **string** | VPC Id | 

## Methods

### NewInternetGatewayCreateRequest

`func NewInternetGatewayCreateRequest(type_ InternetGatewayType, vpcId string, ) *InternetGatewayCreateRequest`

NewInternetGatewayCreateRequest instantiates a new InternetGatewayCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetGatewayCreateRequestWithDefaults

`func NewInternetGatewayCreateRequestWithDefaults() *InternetGatewayCreateRequest`

NewInternetGatewayCreateRequestWithDefaults instantiates a new InternetGatewayCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *InternetGatewayCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternetGatewayCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternetGatewayCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternetGatewayCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InternetGatewayCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InternetGatewayCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFirewallEnabled

`func (o *InternetGatewayCreateRequest) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *InternetGatewayCreateRequest) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *InternetGatewayCreateRequest) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *InternetGatewayCreateRequest) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetFirewallLoggable

`func (o *InternetGatewayCreateRequest) GetFirewallLoggable() bool`

GetFirewallLoggable returns the FirewallLoggable field if non-nil, zero value otherwise.

### GetFirewallLoggableOk

`func (o *InternetGatewayCreateRequest) GetFirewallLoggableOk() (*bool, bool)`

GetFirewallLoggableOk returns a tuple with the FirewallLoggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallLoggable

`func (o *InternetGatewayCreateRequest) SetFirewallLoggable(v bool)`

SetFirewallLoggable sets FirewallLoggable field to given value.

### HasFirewallLoggable

`func (o *InternetGatewayCreateRequest) HasFirewallLoggable() bool`

HasFirewallLoggable returns a boolean if a field has been set.

### GetTags

`func (o *InternetGatewayCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InternetGatewayCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InternetGatewayCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InternetGatewayCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *InternetGatewayCreateRequest) GetType() InternetGatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternetGatewayCreateRequest) GetTypeOk() (*InternetGatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternetGatewayCreateRequest) SetType(v InternetGatewayType)`

SetType sets Type field to given value.


### GetVpcId

`func (o *InternetGatewayCreateRequest) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *InternetGatewayCreateRequest) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *InternetGatewayCreateRequest) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


