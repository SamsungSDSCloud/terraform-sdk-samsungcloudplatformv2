# SubnetCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocationPools** | Pointer to **[]interface{}** | Allocation Pools | [optional] [default to []]
**Cidr** | **string** | Subnet Cidr | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DnsNameservers** | Pointer to **[]string** | DNS Name Servers | [optional] [default to []]
**HostRoutes** | Pointer to **[]interface{}** | Host Routes | [optional] [default to []]
**Name** | **string** | Subnet Name | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]
**Type** | [**SubnetType**](SubnetType.md) | Subnet Type | 
**VpcId** | **string** | VPC Id | 

## Methods

### NewSubnetCreateRequest

`func NewSubnetCreateRequest(cidr string, name string, type_ SubnetType, vpcId string, ) *SubnetCreateRequest`

NewSubnetCreateRequest instantiates a new SubnetCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetCreateRequestWithDefaults

`func NewSubnetCreateRequestWithDefaults() *SubnetCreateRequest`

NewSubnetCreateRequestWithDefaults instantiates a new SubnetCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocationPools

`func (o *SubnetCreateRequest) GetAllocationPools() []interface{}`

GetAllocationPools returns the AllocationPools field if non-nil, zero value otherwise.

### GetAllocationPoolsOk

`func (o *SubnetCreateRequest) GetAllocationPoolsOk() (*[]interface{}, bool)`

GetAllocationPoolsOk returns a tuple with the AllocationPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationPools

`func (o *SubnetCreateRequest) SetAllocationPools(v []interface{})`

SetAllocationPools sets AllocationPools field to given value.

### HasAllocationPools

`func (o *SubnetCreateRequest) HasAllocationPools() bool`

HasAllocationPools returns a boolean if a field has been set.

### GetCidr

`func (o *SubnetCreateRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *SubnetCreateRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *SubnetCreateRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetDescription

`func (o *SubnetCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubnetCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubnetCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubnetCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SubnetCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SubnetCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDnsNameservers

`func (o *SubnetCreateRequest) GetDnsNameservers() []string`

GetDnsNameservers returns the DnsNameservers field if non-nil, zero value otherwise.

### GetDnsNameserversOk

`func (o *SubnetCreateRequest) GetDnsNameserversOk() (*[]string, bool)`

GetDnsNameserversOk returns a tuple with the DnsNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameservers

`func (o *SubnetCreateRequest) SetDnsNameservers(v []string)`

SetDnsNameservers sets DnsNameservers field to given value.

### HasDnsNameservers

`func (o *SubnetCreateRequest) HasDnsNameservers() bool`

HasDnsNameservers returns a boolean if a field has been set.

### GetHostRoutes

`func (o *SubnetCreateRequest) GetHostRoutes() []interface{}`

GetHostRoutes returns the HostRoutes field if non-nil, zero value otherwise.

### GetHostRoutesOk

`func (o *SubnetCreateRequest) GetHostRoutesOk() (*[]interface{}, bool)`

GetHostRoutesOk returns a tuple with the HostRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostRoutes

`func (o *SubnetCreateRequest) SetHostRoutes(v []interface{})`

SetHostRoutes sets HostRoutes field to given value.

### HasHostRoutes

`func (o *SubnetCreateRequest) HasHostRoutes() bool`

HasHostRoutes returns a boolean if a field has been set.

### GetName

`func (o *SubnetCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubnetCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubnetCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *SubnetCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SubnetCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SubnetCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SubnetCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *SubnetCreateRequest) GetType() SubnetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubnetCreateRequest) GetTypeOk() (*SubnetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubnetCreateRequest) SetType(v SubnetType)`

SetType sets Type field to given value.


### GetVpcId

`func (o *SubnetCreateRequest) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *SubnetCreateRequest) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *SubnetCreateRequest) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


