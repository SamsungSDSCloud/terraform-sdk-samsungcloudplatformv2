# TransitGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**Bandwidth** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**FirewallIds** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Transit Gateway ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Transit Gateway Name | 
**State** | [**TransitGatewayState**](TransitGatewayState.md) | State | 
**UplinkEnabled** | Pointer to **bool** | Uplink Enabled | [optional] [default to false]

## Methods

### NewTransitGateway

`func NewTransitGateway(accountId string, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, name string, state TransitGatewayState, ) *TransitGateway`

NewTransitGateway instantiates a new TransitGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayWithDefaults

`func NewTransitGatewayWithDefaults() *TransitGateway`

NewTransitGatewayWithDefaults instantiates a new TransitGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TransitGateway) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransitGateway) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransitGateway) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBandwidth

`func (o *TransitGateway) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *TransitGateway) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *TransitGateway) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *TransitGateway) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### SetBandwidthNil

`func (o *TransitGateway) SetBandwidthNil(b bool)`

 SetBandwidthNil sets the value for Bandwidth to be an explicit nil

### UnsetBandwidth
`func (o *TransitGateway) UnsetBandwidth()`

UnsetBandwidth ensures that no value is present for Bandwidth, not even an explicit nil
### GetCreatedAt

`func (o *TransitGateway) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransitGateway) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransitGateway) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *TransitGateway) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TransitGateway) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TransitGateway) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *TransitGateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransitGateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransitGateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransitGateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TransitGateway) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TransitGateway) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFirewallIds

`func (o *TransitGateway) GetFirewallIds() string`

GetFirewallIds returns the FirewallIds field if non-nil, zero value otherwise.

### GetFirewallIdsOk

`func (o *TransitGateway) GetFirewallIdsOk() (*string, bool)`

GetFirewallIdsOk returns a tuple with the FirewallIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallIds

`func (o *TransitGateway) SetFirewallIds(v string)`

SetFirewallIds sets FirewallIds field to given value.

### HasFirewallIds

`func (o *TransitGateway) HasFirewallIds() bool`

HasFirewallIds returns a boolean if a field has been set.

### SetFirewallIdsNil

`func (o *TransitGateway) SetFirewallIdsNil(b bool)`

 SetFirewallIdsNil sets the value for FirewallIds to be an explicit nil

### UnsetFirewallIds
`func (o *TransitGateway) UnsetFirewallIds()`

UnsetFirewallIds ensures that no value is present for FirewallIds, not even an explicit nil
### GetId

`func (o *TransitGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransitGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransitGateway) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *TransitGateway) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *TransitGateway) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *TransitGateway) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *TransitGateway) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *TransitGateway) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *TransitGateway) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *TransitGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransitGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransitGateway) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *TransitGateway) GetState() TransitGatewayState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TransitGateway) GetStateOk() (*TransitGatewayState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TransitGateway) SetState(v TransitGatewayState)`

SetState sets State field to given value.


### GetUplinkEnabled

`func (o *TransitGateway) GetUplinkEnabled() bool`

GetUplinkEnabled returns the UplinkEnabled field if non-nil, zero value otherwise.

### GetUplinkEnabledOk

`func (o *TransitGateway) GetUplinkEnabledOk() (*bool, bool)`

GetUplinkEnabledOk returns a tuple with the UplinkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkEnabled

`func (o *TransitGateway) SetUplinkEnabled(v bool)`

SetUplinkEnabled sets UplinkEnabled field to given value.

### HasUplinkEnabled

`func (o *TransitGateway) HasUplinkEnabled() bool`

HasUplinkEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


