# SubnetDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AllocationPools** | Pointer to **[]interface{}** | Allocation Pools | [optional] [default to []]
**Cidr** | **string** | Subnet Cidr | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DnsNameservers** | Pointer to **[]string** | DNS Name Servers | [optional] [default to []]
**GatewayIpAddress** | **NullableString** |  | 
**HostRoutes** | Pointer to **[]interface{}** | Host Routes | [optional] [default to []]
**Id** | **string** | Subnet Id | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Subnet Name | 
**State** | [**SubnetState**](SubnetState.md) | State | 
**Type** | [**SubnetType**](SubnetType.md) | Subnet Type | 
**VpcId** | **string** | VPC Id | 
**VpcName** | **string** | VPC Name | 

## Methods

### NewSubnetDetail

`func NewSubnetDetail(accountId string, cidr string, createdAt time.Time, createdBy string, gatewayIpAddress NullableString, id string, modifiedAt time.Time, modifiedBy string, name string, state SubnetState, type_ SubnetType, vpcId string, vpcName string, ) *SubnetDetail`

NewSubnetDetail instantiates a new SubnetDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetDetailWithDefaults

`func NewSubnetDetailWithDefaults() *SubnetDetail`

NewSubnetDetailWithDefaults instantiates a new SubnetDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SubnetDetail) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SubnetDetail) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SubnetDetail) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAllocationPools

`func (o *SubnetDetail) GetAllocationPools() []interface{}`

GetAllocationPools returns the AllocationPools field if non-nil, zero value otherwise.

### GetAllocationPoolsOk

`func (o *SubnetDetail) GetAllocationPoolsOk() (*[]interface{}, bool)`

GetAllocationPoolsOk returns a tuple with the AllocationPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationPools

`func (o *SubnetDetail) SetAllocationPools(v []interface{})`

SetAllocationPools sets AllocationPools field to given value.

### HasAllocationPools

`func (o *SubnetDetail) HasAllocationPools() bool`

HasAllocationPools returns a boolean if a field has been set.

### GetCidr

`func (o *SubnetDetail) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *SubnetDetail) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *SubnetDetail) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetCreatedAt

`func (o *SubnetDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubnetDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubnetDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SubnetDetail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SubnetDetail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SubnetDetail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *SubnetDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubnetDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubnetDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubnetDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SubnetDetail) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SubnetDetail) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDnsNameservers

`func (o *SubnetDetail) GetDnsNameservers() []string`

GetDnsNameservers returns the DnsNameservers field if non-nil, zero value otherwise.

### GetDnsNameserversOk

`func (o *SubnetDetail) GetDnsNameserversOk() (*[]string, bool)`

GetDnsNameserversOk returns a tuple with the DnsNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameservers

`func (o *SubnetDetail) SetDnsNameservers(v []string)`

SetDnsNameservers sets DnsNameservers field to given value.

### HasDnsNameservers

`func (o *SubnetDetail) HasDnsNameservers() bool`

HasDnsNameservers returns a boolean if a field has been set.

### GetGatewayIpAddress

`func (o *SubnetDetail) GetGatewayIpAddress() string`

GetGatewayIpAddress returns the GatewayIpAddress field if non-nil, zero value otherwise.

### GetGatewayIpAddressOk

`func (o *SubnetDetail) GetGatewayIpAddressOk() (*string, bool)`

GetGatewayIpAddressOk returns a tuple with the GatewayIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIpAddress

`func (o *SubnetDetail) SetGatewayIpAddress(v string)`

SetGatewayIpAddress sets GatewayIpAddress field to given value.


### SetGatewayIpAddressNil

`func (o *SubnetDetail) SetGatewayIpAddressNil(b bool)`

 SetGatewayIpAddressNil sets the value for GatewayIpAddress to be an explicit nil

### UnsetGatewayIpAddress
`func (o *SubnetDetail) UnsetGatewayIpAddress()`

UnsetGatewayIpAddress ensures that no value is present for GatewayIpAddress, not even an explicit nil
### GetHostRoutes

`func (o *SubnetDetail) GetHostRoutes() []interface{}`

GetHostRoutes returns the HostRoutes field if non-nil, zero value otherwise.

### GetHostRoutesOk

`func (o *SubnetDetail) GetHostRoutesOk() (*[]interface{}, bool)`

GetHostRoutesOk returns a tuple with the HostRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostRoutes

`func (o *SubnetDetail) SetHostRoutes(v []interface{})`

SetHostRoutes sets HostRoutes field to given value.

### HasHostRoutes

`func (o *SubnetDetail) HasHostRoutes() bool`

HasHostRoutes returns a boolean if a field has been set.

### GetId

`func (o *SubnetDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubnetDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubnetDetail) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *SubnetDetail) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SubnetDetail) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SubnetDetail) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *SubnetDetail) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SubnetDetail) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *SubnetDetail) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *SubnetDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubnetDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubnetDetail) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *SubnetDetail) GetState() SubnetState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SubnetDetail) GetStateOk() (*SubnetState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SubnetDetail) SetState(v SubnetState)`

SetState sets State field to given value.


### GetType

`func (o *SubnetDetail) GetType() SubnetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubnetDetail) GetTypeOk() (*SubnetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubnetDetail) SetType(v SubnetType)`

SetType sets Type field to given value.


### GetVpcId

`func (o *SubnetDetail) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *SubnetDetail) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *SubnetDetail) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *SubnetDetail) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *SubnetDetail) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *SubnetDetail) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


