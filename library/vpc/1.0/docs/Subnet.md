# Subnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**Cidr** | **string** | Subnet Cidr | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**GatewayIpAddress** | **NullableString** |  | 
**Id** | **string** | Subnet Id | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Subnet Name | 
**State** | [**SubnetState**](SubnetState.md) | State | 
**Type** | [**SubnetType**](SubnetType.md) | Subnet Type | 
**VpcId** | **string** | VPC Id | 
**VpcName** | **string** | VPC Name | 

## Methods

### NewSubnet

`func NewSubnet(accountId string, cidr string, createdAt time.Time, createdBy string, gatewayIpAddress NullableString, id string, modifiedAt time.Time, modifiedBy string, name string, state SubnetState, type_ SubnetType, vpcId string, vpcName string, ) *Subnet`

NewSubnet instantiates a new Subnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetWithDefaults

`func NewSubnetWithDefaults() *Subnet`

NewSubnetWithDefaults instantiates a new Subnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Subnet) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Subnet) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Subnet) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCidr

`func (o *Subnet) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Subnet) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Subnet) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetCreatedAt

`func (o *Subnet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subnet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subnet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Subnet) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Subnet) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Subnet) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetGatewayIpAddress

`func (o *Subnet) GetGatewayIpAddress() string`

GetGatewayIpAddress returns the GatewayIpAddress field if non-nil, zero value otherwise.

### GetGatewayIpAddressOk

`func (o *Subnet) GetGatewayIpAddressOk() (*string, bool)`

GetGatewayIpAddressOk returns a tuple with the GatewayIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIpAddress

`func (o *Subnet) SetGatewayIpAddress(v string)`

SetGatewayIpAddress sets GatewayIpAddress field to given value.


### SetGatewayIpAddressNil

`func (o *Subnet) SetGatewayIpAddressNil(b bool)`

 SetGatewayIpAddressNil sets the value for GatewayIpAddress to be an explicit nil

### UnsetGatewayIpAddress
`func (o *Subnet) UnsetGatewayIpAddress()`

UnsetGatewayIpAddress ensures that no value is present for GatewayIpAddress, not even an explicit nil
### GetId

`func (o *Subnet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subnet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subnet) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *Subnet) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Subnet) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Subnet) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *Subnet) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Subnet) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Subnet) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *Subnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subnet) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *Subnet) GetState() SubnetState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Subnet) GetStateOk() (*SubnetState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Subnet) SetState(v SubnetState)`

SetState sets State field to given value.


### GetType

`func (o *Subnet) GetType() SubnetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Subnet) GetTypeOk() (*SubnetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Subnet) SetType(v SubnetType)`

SetType sets Type field to given value.


### GetVpcId

`func (o *Subnet) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *Subnet) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *Subnet) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *Subnet) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *Subnet) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *Subnet) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


