# PrivateDns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthDnsName** | **NullableString** |  | 
**ConnectedVpcIds** | Pointer to **[]string** | Connected VPC ID | [optional] 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | **NullableString** |  | 
**Id** | **string** | Private DNS ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Private DNS Name | 
**PoolId** | **NullableString** |  | 
**PoolName** | **NullableString** |  | 
**RegisteredRegion** | **NullableString** |  | 
**ResolverIp** | **NullableString** |  | 
**ResolverName** | **NullableString** |  | 
**State** | [**PrivateDnsState**](PrivateDnsState.md) | State | 

## Methods

### NewPrivateDns

`func NewPrivateDns(authDnsName NullableString, createdAt time.Time, createdBy string, description NullableString, id string, modifiedAt time.Time, modifiedBy string, name string, poolId NullableString, poolName NullableString, registeredRegion NullableString, resolverIp NullableString, resolverName NullableString, state PrivateDnsState, ) *PrivateDns`

NewPrivateDns instantiates a new PrivateDns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateDnsWithDefaults

`func NewPrivateDnsWithDefaults() *PrivateDns`

NewPrivateDnsWithDefaults instantiates a new PrivateDns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthDnsName

`func (o *PrivateDns) GetAuthDnsName() string`

GetAuthDnsName returns the AuthDnsName field if non-nil, zero value otherwise.

### GetAuthDnsNameOk

`func (o *PrivateDns) GetAuthDnsNameOk() (*string, bool)`

GetAuthDnsNameOk returns a tuple with the AuthDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDnsName

`func (o *PrivateDns) SetAuthDnsName(v string)`

SetAuthDnsName sets AuthDnsName field to given value.


### SetAuthDnsNameNil

`func (o *PrivateDns) SetAuthDnsNameNil(b bool)`

 SetAuthDnsNameNil sets the value for AuthDnsName to be an explicit nil

### UnsetAuthDnsName
`func (o *PrivateDns) UnsetAuthDnsName()`

UnsetAuthDnsName ensures that no value is present for AuthDnsName, not even an explicit nil
### GetConnectedVpcIds

`func (o *PrivateDns) GetConnectedVpcIds() []string`

GetConnectedVpcIds returns the ConnectedVpcIds field if non-nil, zero value otherwise.

### GetConnectedVpcIdsOk

`func (o *PrivateDns) GetConnectedVpcIdsOk() (*[]string, bool)`

GetConnectedVpcIdsOk returns a tuple with the ConnectedVpcIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedVpcIds

`func (o *PrivateDns) SetConnectedVpcIds(v []string)`

SetConnectedVpcIds sets ConnectedVpcIds field to given value.

### HasConnectedVpcIds

`func (o *PrivateDns) HasConnectedVpcIds() bool`

HasConnectedVpcIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PrivateDns) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrivateDns) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrivateDns) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *PrivateDns) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PrivateDns) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PrivateDns) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *PrivateDns) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateDns) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateDns) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *PrivateDns) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PrivateDns) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *PrivateDns) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateDns) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateDns) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *PrivateDns) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PrivateDns) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PrivateDns) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *PrivateDns) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PrivateDns) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PrivateDns) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *PrivateDns) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateDns) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateDns) SetName(v string)`

SetName sets Name field to given value.


### GetPoolId

`func (o *PrivateDns) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PrivateDns) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PrivateDns) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### SetPoolIdNil

`func (o *PrivateDns) SetPoolIdNil(b bool)`

 SetPoolIdNil sets the value for PoolId to be an explicit nil

### UnsetPoolId
`func (o *PrivateDns) UnsetPoolId()`

UnsetPoolId ensures that no value is present for PoolId, not even an explicit nil
### GetPoolName

`func (o *PrivateDns) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *PrivateDns) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *PrivateDns) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.


### SetPoolNameNil

`func (o *PrivateDns) SetPoolNameNil(b bool)`

 SetPoolNameNil sets the value for PoolName to be an explicit nil

### UnsetPoolName
`func (o *PrivateDns) UnsetPoolName()`

UnsetPoolName ensures that no value is present for PoolName, not even an explicit nil
### GetRegisteredRegion

`func (o *PrivateDns) GetRegisteredRegion() string`

GetRegisteredRegion returns the RegisteredRegion field if non-nil, zero value otherwise.

### GetRegisteredRegionOk

`func (o *PrivateDns) GetRegisteredRegionOk() (*string, bool)`

GetRegisteredRegionOk returns a tuple with the RegisteredRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredRegion

`func (o *PrivateDns) SetRegisteredRegion(v string)`

SetRegisteredRegion sets RegisteredRegion field to given value.


### SetRegisteredRegionNil

`func (o *PrivateDns) SetRegisteredRegionNil(b bool)`

 SetRegisteredRegionNil sets the value for RegisteredRegion to be an explicit nil

### UnsetRegisteredRegion
`func (o *PrivateDns) UnsetRegisteredRegion()`

UnsetRegisteredRegion ensures that no value is present for RegisteredRegion, not even an explicit nil
### GetResolverIp

`func (o *PrivateDns) GetResolverIp() string`

GetResolverIp returns the ResolverIp field if non-nil, zero value otherwise.

### GetResolverIpOk

`func (o *PrivateDns) GetResolverIpOk() (*string, bool)`

GetResolverIpOk returns a tuple with the ResolverIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverIp

`func (o *PrivateDns) SetResolverIp(v string)`

SetResolverIp sets ResolverIp field to given value.


### SetResolverIpNil

`func (o *PrivateDns) SetResolverIpNil(b bool)`

 SetResolverIpNil sets the value for ResolverIp to be an explicit nil

### UnsetResolverIp
`func (o *PrivateDns) UnsetResolverIp()`

UnsetResolverIp ensures that no value is present for ResolverIp, not even an explicit nil
### GetResolverName

`func (o *PrivateDns) GetResolverName() string`

GetResolverName returns the ResolverName field if non-nil, zero value otherwise.

### GetResolverNameOk

`func (o *PrivateDns) GetResolverNameOk() (*string, bool)`

GetResolverNameOk returns a tuple with the ResolverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverName

`func (o *PrivateDns) SetResolverName(v string)`

SetResolverName sets ResolverName field to given value.


### SetResolverNameNil

`func (o *PrivateDns) SetResolverNameNil(b bool)`

 SetResolverNameNil sets the value for ResolverName to be an explicit nil

### UnsetResolverName
`func (o *PrivateDns) UnsetResolverName()`

UnsetResolverName ensures that no value is present for ResolverName, not even an explicit nil
### GetState

`func (o *PrivateDns) GetState() PrivateDnsState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PrivateDns) GetStateOk() (*PrivateDnsState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PrivateDns) SetState(v PrivateDnsState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


