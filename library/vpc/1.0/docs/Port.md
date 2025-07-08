# Port

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AttachedResourceId** | **string** | Connected resource ID | 
**AttachedResourceType** | **string** | Connected resource Type | 
**CreatedAt** | **time.Time** | Created date | 
**Description** | **string** | Description | 
**FixedIpAddress** | **string** | Fixed IP | 
**Id** | **string** | Port ID | 
**MacAddress** | **string** | MAC Address | 
**ModifiedAt** | **time.Time** | Modified date | 
**Name** | **string** | Port Name | 
**SecurityGroups** | [**[]PortSecurityGroup**](PortSecurityGroup.md) | Security Group List | 
**State** | **string** | State | 
**SubnetId** | **string** | Subnet Id | 
**SubnetName** | **string** | Subnet Name | 
**VpcId** | **string** | VPC Id | 
**VpcName** | **string** | VPC Name | 

## Methods

### NewPort

`func NewPort(accountId string, attachedResourceId string, attachedResourceType string, createdAt time.Time, description string, fixedIpAddress string, id string, macAddress string, modifiedAt time.Time, name string, securityGroups []PortSecurityGroup, state string, subnetId string, subnetName string, vpcId string, vpcName string, ) *Port`

NewPort instantiates a new Port object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortWithDefaults

`func NewPortWithDefaults() *Port`

NewPortWithDefaults instantiates a new Port object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Port) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Port) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Port) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAttachedResourceId

`func (o *Port) GetAttachedResourceId() string`

GetAttachedResourceId returns the AttachedResourceId field if non-nil, zero value otherwise.

### GetAttachedResourceIdOk

`func (o *Port) GetAttachedResourceIdOk() (*string, bool)`

GetAttachedResourceIdOk returns a tuple with the AttachedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedResourceId

`func (o *Port) SetAttachedResourceId(v string)`

SetAttachedResourceId sets AttachedResourceId field to given value.


### GetAttachedResourceType

`func (o *Port) GetAttachedResourceType() string`

GetAttachedResourceType returns the AttachedResourceType field if non-nil, zero value otherwise.

### GetAttachedResourceTypeOk

`func (o *Port) GetAttachedResourceTypeOk() (*string, bool)`

GetAttachedResourceTypeOk returns a tuple with the AttachedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedResourceType

`func (o *Port) SetAttachedResourceType(v string)`

SetAttachedResourceType sets AttachedResourceType field to given value.


### GetCreatedAt

`func (o *Port) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Port) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Port) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *Port) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Port) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Port) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFixedIpAddress

`func (o *Port) GetFixedIpAddress() string`

GetFixedIpAddress returns the FixedIpAddress field if non-nil, zero value otherwise.

### GetFixedIpAddressOk

`func (o *Port) GetFixedIpAddressOk() (*string, bool)`

GetFixedIpAddressOk returns a tuple with the FixedIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAddress

`func (o *Port) SetFixedIpAddress(v string)`

SetFixedIpAddress sets FixedIpAddress field to given value.


### GetId

`func (o *Port) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Port) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Port) SetId(v string)`

SetId sets Id field to given value.


### GetMacAddress

`func (o *Port) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Port) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Port) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetModifiedAt

`func (o *Port) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Port) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Port) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetName

`func (o *Port) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Port) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Port) SetName(v string)`

SetName sets Name field to given value.


### GetSecurityGroups

`func (o *Port) GetSecurityGroups() []PortSecurityGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *Port) GetSecurityGroupsOk() (*[]PortSecurityGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *Port) SetSecurityGroups(v []PortSecurityGroup)`

SetSecurityGroups sets SecurityGroups field to given value.


### GetState

`func (o *Port) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Port) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Port) SetState(v string)`

SetState sets State field to given value.


### GetSubnetId

`func (o *Port) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *Port) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *Port) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetName

`func (o *Port) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *Port) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *Port) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.


### GetVpcId

`func (o *Port) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *Port) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *Port) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *Port) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *Port) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *Port) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


