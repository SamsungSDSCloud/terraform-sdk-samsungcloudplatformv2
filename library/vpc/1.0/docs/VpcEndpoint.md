# VpcEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**EndpointIpAddress** | **string** | VPC Endpoint IP Address | 
**Id** | **string** | VPC Endpoint ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | VPC Endpoint Name | 
**ResourceInfo** | **string** | VPC Endpoint Resource Key Info | 
**ResourceKey** | **string** | VPC Endpoint Resource Key | 
**ResourceType** | [**VpcEndpointResourceType**](VpcEndpointResourceType.md) | VPC Endpoint Resource Type | 
**State** | [**VpcEndpointState**](VpcEndpointState.md) | VPC Endpoint State | 
**SubnetId** | **string** | Subnet Id | 
**SubnetName** | **string** | Subnet Name | 
**VpcId** | **string** | VPC Id | 
**VpcName** | **string** | VPC Name | 

## Methods

### NewVpcEndpoint

`func NewVpcEndpoint(accountId string, createdAt time.Time, createdBy string, endpointIpAddress string, id string, modifiedAt time.Time, modifiedBy string, name string, resourceInfo string, resourceKey string, resourceType VpcEndpointResourceType, state VpcEndpointState, subnetId string, subnetName string, vpcId string, vpcName string, ) *VpcEndpoint`

NewVpcEndpoint instantiates a new VpcEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcEndpointWithDefaults

`func NewVpcEndpointWithDefaults() *VpcEndpoint`

NewVpcEndpointWithDefaults instantiates a new VpcEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *VpcEndpoint) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VpcEndpoint) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VpcEndpoint) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *VpcEndpoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VpcEndpoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VpcEndpoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *VpcEndpoint) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VpcEndpoint) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VpcEndpoint) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *VpcEndpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpcEndpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpcEndpoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpcEndpoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VpcEndpoint) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VpcEndpoint) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndpointIpAddress

`func (o *VpcEndpoint) GetEndpointIpAddress() string`

GetEndpointIpAddress returns the EndpointIpAddress field if non-nil, zero value otherwise.

### GetEndpointIpAddressOk

`func (o *VpcEndpoint) GetEndpointIpAddressOk() (*string, bool)`

GetEndpointIpAddressOk returns a tuple with the EndpointIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointIpAddress

`func (o *VpcEndpoint) SetEndpointIpAddress(v string)`

SetEndpointIpAddress sets EndpointIpAddress field to given value.


### GetId

`func (o *VpcEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpcEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpcEndpoint) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *VpcEndpoint) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VpcEndpoint) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VpcEndpoint) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *VpcEndpoint) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VpcEndpoint) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VpcEndpoint) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *VpcEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpcEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpcEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetResourceInfo

`func (o *VpcEndpoint) GetResourceInfo() string`

GetResourceInfo returns the ResourceInfo field if non-nil, zero value otherwise.

### GetResourceInfoOk

`func (o *VpcEndpoint) GetResourceInfoOk() (*string, bool)`

GetResourceInfoOk returns a tuple with the ResourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInfo

`func (o *VpcEndpoint) SetResourceInfo(v string)`

SetResourceInfo sets ResourceInfo field to given value.


### GetResourceKey

`func (o *VpcEndpoint) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *VpcEndpoint) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *VpcEndpoint) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetResourceType

`func (o *VpcEndpoint) GetResourceType() VpcEndpointResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *VpcEndpoint) GetResourceTypeOk() (*VpcEndpointResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *VpcEndpoint) SetResourceType(v VpcEndpointResourceType)`

SetResourceType sets ResourceType field to given value.


### GetState

`func (o *VpcEndpoint) GetState() VpcEndpointState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpcEndpoint) GetStateOk() (*VpcEndpointState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpcEndpoint) SetState(v VpcEndpointState)`

SetState sets State field to given value.


### GetSubnetId

`func (o *VpcEndpoint) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *VpcEndpoint) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *VpcEndpoint) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetName

`func (o *VpcEndpoint) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *VpcEndpoint) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *VpcEndpoint) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.


### GetVpcId

`func (o *VpcEndpoint) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *VpcEndpoint) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *VpcEndpoint) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *VpcEndpoint) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *VpcEndpoint) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *VpcEndpoint) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


