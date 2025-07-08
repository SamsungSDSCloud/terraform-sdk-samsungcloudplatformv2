# InternetGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**FirewallId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Internet Gateway ID | 
**Loggable** | Pointer to **bool** | NAT Loggable | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Internet Gateway Name | 
**State** | [**InternetGatewayState**](InternetGatewayState.md) | State | 
**Type** | [**InternetGatewayType**](InternetGatewayType.md) | Internet Gateway Type | 
**VpcId** | **string** | VPC Id | 
**VpcName** | **string** | VPC Name | 

## Methods

### NewInternetGateway

`func NewInternetGateway(accountId string, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, name string, state InternetGatewayState, type_ InternetGatewayType, vpcId string, vpcName string, ) *InternetGateway`

NewInternetGateway instantiates a new InternetGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetGatewayWithDefaults

`func NewInternetGatewayWithDefaults() *InternetGateway`

NewInternetGatewayWithDefaults instantiates a new InternetGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *InternetGateway) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InternetGateway) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InternetGateway) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *InternetGateway) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InternetGateway) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InternetGateway) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *InternetGateway) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InternetGateway) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InternetGateway) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *InternetGateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternetGateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternetGateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternetGateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InternetGateway) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InternetGateway) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFirewallId

`func (o *InternetGateway) GetFirewallId() string`

GetFirewallId returns the FirewallId field if non-nil, zero value otherwise.

### GetFirewallIdOk

`func (o *InternetGateway) GetFirewallIdOk() (*string, bool)`

GetFirewallIdOk returns a tuple with the FirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallId

`func (o *InternetGateway) SetFirewallId(v string)`

SetFirewallId sets FirewallId field to given value.

### HasFirewallId

`func (o *InternetGateway) HasFirewallId() bool`

HasFirewallId returns a boolean if a field has been set.

### SetFirewallIdNil

`func (o *InternetGateway) SetFirewallIdNil(b bool)`

 SetFirewallIdNil sets the value for FirewallId to be an explicit nil

### UnsetFirewallId
`func (o *InternetGateway) UnsetFirewallId()`

UnsetFirewallId ensures that no value is present for FirewallId, not even an explicit nil
### GetId

`func (o *InternetGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternetGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternetGateway) SetId(v string)`

SetId sets Id field to given value.


### GetLoggable

`func (o *InternetGateway) GetLoggable() bool`

GetLoggable returns the Loggable field if non-nil, zero value otherwise.

### GetLoggableOk

`func (o *InternetGateway) GetLoggableOk() (*bool, bool)`

GetLoggableOk returns a tuple with the Loggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggable

`func (o *InternetGateway) SetLoggable(v bool)`

SetLoggable sets Loggable field to given value.

### HasLoggable

`func (o *InternetGateway) HasLoggable() bool`

HasLoggable returns a boolean if a field has been set.

### GetModifiedAt

`func (o *InternetGateway) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *InternetGateway) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *InternetGateway) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *InternetGateway) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *InternetGateway) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *InternetGateway) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *InternetGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternetGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternetGateway) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *InternetGateway) GetState() InternetGatewayState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InternetGateway) GetStateOk() (*InternetGatewayState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InternetGateway) SetState(v InternetGatewayState)`

SetState sets State field to given value.


### GetType

`func (o *InternetGateway) GetType() InternetGatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternetGateway) GetTypeOk() (*InternetGatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternetGateway) SetType(v InternetGatewayType)`

SetType sets Type field to given value.


### GetVpcId

`func (o *InternetGateway) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *InternetGateway) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *InternetGateway) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *InternetGateway) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *InternetGateway) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *InternetGateway) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


