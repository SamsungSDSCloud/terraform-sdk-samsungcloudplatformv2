# PortList

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
**State** | **string** | State | 
**SubnetId** | **string** | Subnet Id | 
**SubnetName** | **string** | Subnet Name | 
**VpcId** | **string** | VPC Id | 
**VpcName** | **string** | VPC Name | 

## Methods

### NewPortList

`func NewPortList(accountId string, attachedResourceId string, attachedResourceType string, createdAt time.Time, description string, fixedIpAddress string, id string, macAddress string, modifiedAt time.Time, name string, state string, subnetId string, subnetName string, vpcId string, vpcName string, ) *PortList`

NewPortList instantiates a new PortList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortListWithDefaults

`func NewPortListWithDefaults() *PortList`

NewPortListWithDefaults instantiates a new PortList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PortList) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PortList) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PortList) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAttachedResourceId

`func (o *PortList) GetAttachedResourceId() string`

GetAttachedResourceId returns the AttachedResourceId field if non-nil, zero value otherwise.

### GetAttachedResourceIdOk

`func (o *PortList) GetAttachedResourceIdOk() (*string, bool)`

GetAttachedResourceIdOk returns a tuple with the AttachedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedResourceId

`func (o *PortList) SetAttachedResourceId(v string)`

SetAttachedResourceId sets AttachedResourceId field to given value.


### GetAttachedResourceType

`func (o *PortList) GetAttachedResourceType() string`

GetAttachedResourceType returns the AttachedResourceType field if non-nil, zero value otherwise.

### GetAttachedResourceTypeOk

`func (o *PortList) GetAttachedResourceTypeOk() (*string, bool)`

GetAttachedResourceTypeOk returns a tuple with the AttachedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedResourceType

`func (o *PortList) SetAttachedResourceType(v string)`

SetAttachedResourceType sets AttachedResourceType field to given value.


### GetCreatedAt

`func (o *PortList) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortList) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortList) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *PortList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PortList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PortList) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFixedIpAddress

`func (o *PortList) GetFixedIpAddress() string`

GetFixedIpAddress returns the FixedIpAddress field if non-nil, zero value otherwise.

### GetFixedIpAddressOk

`func (o *PortList) GetFixedIpAddressOk() (*string, bool)`

GetFixedIpAddressOk returns a tuple with the FixedIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAddress

`func (o *PortList) SetFixedIpAddress(v string)`

SetFixedIpAddress sets FixedIpAddress field to given value.


### GetId

`func (o *PortList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortList) SetId(v string)`

SetId sets Id field to given value.


### GetMacAddress

`func (o *PortList) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *PortList) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *PortList) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetModifiedAt

`func (o *PortList) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PortList) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PortList) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetName

`func (o *PortList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortList) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *PortList) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PortList) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PortList) SetState(v string)`

SetState sets State field to given value.


### GetSubnetId

`func (o *PortList) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *PortList) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *PortList) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetName

`func (o *PortList) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *PortList) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *PortList) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.


### GetVpcId

`func (o *PortList) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *PortList) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *PortList) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *PortList) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *PortList) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *PortList) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


