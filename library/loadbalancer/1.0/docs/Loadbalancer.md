# Loadbalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account ID of the load balancer. | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**FirewallId** | Pointer to **NullableString** |  | [optional] 
**FirewallName** | Pointer to **NullableString** |  | [optional] 
**HealthCheckIp** | **[]string** | Health check IP | 
**Id** | **string** | ID | 
**LayerType** | **string** | Layer type | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | name | 
**PublicNatIp** | Pointer to **NullableString** |  | [optional] 
**PublicNatState** | Pointer to **NullableString** |  | [optional] 
**ServiceIp** | **string** | Service IP | 
**SourceNatIp** | **string** | Source NAT IP | 
**State** | **string** | The state of the load balancer. | 
**SubnetId** | **string** | Subnet ID | 
**VpcId** | **string** | VPC ID | 

## Methods

### NewLoadbalancer

`func NewLoadbalancer(accountId string, createdAt time.Time, createdBy string, healthCheckIp []string, id string, layerType string, modifiedAt time.Time, modifiedBy string, name string, serviceIp string, sourceNatIp string, state string, subnetId string, vpcId string, ) *Loadbalancer`

NewLoadbalancer instantiates a new Loadbalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadbalancerWithDefaults

`func NewLoadbalancerWithDefaults() *Loadbalancer`

NewLoadbalancerWithDefaults instantiates a new Loadbalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Loadbalancer) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Loadbalancer) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Loadbalancer) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *Loadbalancer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Loadbalancer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Loadbalancer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Loadbalancer) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Loadbalancer) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Loadbalancer) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *Loadbalancer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Loadbalancer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Loadbalancer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Loadbalancer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Loadbalancer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Loadbalancer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFirewallId

`func (o *Loadbalancer) GetFirewallId() string`

GetFirewallId returns the FirewallId field if non-nil, zero value otherwise.

### GetFirewallIdOk

`func (o *Loadbalancer) GetFirewallIdOk() (*string, bool)`

GetFirewallIdOk returns a tuple with the FirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallId

`func (o *Loadbalancer) SetFirewallId(v string)`

SetFirewallId sets FirewallId field to given value.

### HasFirewallId

`func (o *Loadbalancer) HasFirewallId() bool`

HasFirewallId returns a boolean if a field has been set.

### SetFirewallIdNil

`func (o *Loadbalancer) SetFirewallIdNil(b bool)`

 SetFirewallIdNil sets the value for FirewallId to be an explicit nil

### UnsetFirewallId
`func (o *Loadbalancer) UnsetFirewallId()`

UnsetFirewallId ensures that no value is present for FirewallId, not even an explicit nil
### GetFirewallName

`func (o *Loadbalancer) GetFirewallName() string`

GetFirewallName returns the FirewallName field if non-nil, zero value otherwise.

### GetFirewallNameOk

`func (o *Loadbalancer) GetFirewallNameOk() (*string, bool)`

GetFirewallNameOk returns a tuple with the FirewallName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallName

`func (o *Loadbalancer) SetFirewallName(v string)`

SetFirewallName sets FirewallName field to given value.

### HasFirewallName

`func (o *Loadbalancer) HasFirewallName() bool`

HasFirewallName returns a boolean if a field has been set.

### SetFirewallNameNil

`func (o *Loadbalancer) SetFirewallNameNil(b bool)`

 SetFirewallNameNil sets the value for FirewallName to be an explicit nil

### UnsetFirewallName
`func (o *Loadbalancer) UnsetFirewallName()`

UnsetFirewallName ensures that no value is present for FirewallName, not even an explicit nil
### GetHealthCheckIp

`func (o *Loadbalancer) GetHealthCheckIp() []string`

GetHealthCheckIp returns the HealthCheckIp field if non-nil, zero value otherwise.

### GetHealthCheckIpOk

`func (o *Loadbalancer) GetHealthCheckIpOk() (*[]string, bool)`

GetHealthCheckIpOk returns a tuple with the HealthCheckIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckIp

`func (o *Loadbalancer) SetHealthCheckIp(v []string)`

SetHealthCheckIp sets HealthCheckIp field to given value.


### GetId

`func (o *Loadbalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Loadbalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Loadbalancer) SetId(v string)`

SetId sets Id field to given value.


### GetLayerType

`func (o *Loadbalancer) GetLayerType() string`

GetLayerType returns the LayerType field if non-nil, zero value otherwise.

### GetLayerTypeOk

`func (o *Loadbalancer) GetLayerTypeOk() (*string, bool)`

GetLayerTypeOk returns a tuple with the LayerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerType

`func (o *Loadbalancer) SetLayerType(v string)`

SetLayerType sets LayerType field to given value.


### GetModifiedAt

`func (o *Loadbalancer) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Loadbalancer) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Loadbalancer) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *Loadbalancer) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Loadbalancer) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Loadbalancer) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *Loadbalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Loadbalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Loadbalancer) SetName(v string)`

SetName sets Name field to given value.


### GetPublicNatIp

`func (o *Loadbalancer) GetPublicNatIp() string`

GetPublicNatIp returns the PublicNatIp field if non-nil, zero value otherwise.

### GetPublicNatIpOk

`func (o *Loadbalancer) GetPublicNatIpOk() (*string, bool)`

GetPublicNatIpOk returns a tuple with the PublicNatIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNatIp

`func (o *Loadbalancer) SetPublicNatIp(v string)`

SetPublicNatIp sets PublicNatIp field to given value.

### HasPublicNatIp

`func (o *Loadbalancer) HasPublicNatIp() bool`

HasPublicNatIp returns a boolean if a field has been set.

### SetPublicNatIpNil

`func (o *Loadbalancer) SetPublicNatIpNil(b bool)`

 SetPublicNatIpNil sets the value for PublicNatIp to be an explicit nil

### UnsetPublicNatIp
`func (o *Loadbalancer) UnsetPublicNatIp()`

UnsetPublicNatIp ensures that no value is present for PublicNatIp, not even an explicit nil
### GetPublicNatState

`func (o *Loadbalancer) GetPublicNatState() string`

GetPublicNatState returns the PublicNatState field if non-nil, zero value otherwise.

### GetPublicNatStateOk

`func (o *Loadbalancer) GetPublicNatStateOk() (*string, bool)`

GetPublicNatStateOk returns a tuple with the PublicNatState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNatState

`func (o *Loadbalancer) SetPublicNatState(v string)`

SetPublicNatState sets PublicNatState field to given value.

### HasPublicNatState

`func (o *Loadbalancer) HasPublicNatState() bool`

HasPublicNatState returns a boolean if a field has been set.

### SetPublicNatStateNil

`func (o *Loadbalancer) SetPublicNatStateNil(b bool)`

 SetPublicNatStateNil sets the value for PublicNatState to be an explicit nil

### UnsetPublicNatState
`func (o *Loadbalancer) UnsetPublicNatState()`

UnsetPublicNatState ensures that no value is present for PublicNatState, not even an explicit nil
### GetServiceIp

`func (o *Loadbalancer) GetServiceIp() string`

GetServiceIp returns the ServiceIp field if non-nil, zero value otherwise.

### GetServiceIpOk

`func (o *Loadbalancer) GetServiceIpOk() (*string, bool)`

GetServiceIpOk returns a tuple with the ServiceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIp

`func (o *Loadbalancer) SetServiceIp(v string)`

SetServiceIp sets ServiceIp field to given value.


### GetSourceNatIp

`func (o *Loadbalancer) GetSourceNatIp() string`

GetSourceNatIp returns the SourceNatIp field if non-nil, zero value otherwise.

### GetSourceNatIpOk

`func (o *Loadbalancer) GetSourceNatIpOk() (*string, bool)`

GetSourceNatIpOk returns a tuple with the SourceNatIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNatIp

`func (o *Loadbalancer) SetSourceNatIp(v string)`

SetSourceNatIp sets SourceNatIp field to given value.


### GetState

`func (o *Loadbalancer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Loadbalancer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Loadbalancer) SetState(v string)`

SetState sets State field to given value.


### GetSubnetId

`func (o *Loadbalancer) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *Loadbalancer) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *Loadbalancer) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetVpcId

`func (o *Loadbalancer) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *Loadbalancer) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *Loadbalancer) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


