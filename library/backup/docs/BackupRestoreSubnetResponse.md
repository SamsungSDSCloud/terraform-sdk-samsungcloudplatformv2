# BackupRestoreSubnetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account  ID | 
**Cidr** | **string** | Subnet Cidr | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**GatewayIpAddress** | **NullableString** |  | 
**Id** | **string** | Subnet ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Subnet Name | 
**State** | **string** | Subnet state | 
**Type** | **string** | Subnet Type | 
**VpcId** | **string** | VPC ID | 
**VpcName** | **string** | VPC name | 

## Methods

### NewBackupRestoreSubnetResponse

`func NewBackupRestoreSubnetResponse(accountId string, cidr string, createdAt time.Time, createdBy string, gatewayIpAddress NullableString, id string, modifiedAt time.Time, modifiedBy string, name string, state string, type_ string, vpcId string, vpcName string, ) *BackupRestoreSubnetResponse`

NewBackupRestoreSubnetResponse instantiates a new BackupRestoreSubnetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreSubnetResponseWithDefaults

`func NewBackupRestoreSubnetResponseWithDefaults() *BackupRestoreSubnetResponse`

NewBackupRestoreSubnetResponseWithDefaults instantiates a new BackupRestoreSubnetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BackupRestoreSubnetResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BackupRestoreSubnetResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BackupRestoreSubnetResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCidr

`func (o *BackupRestoreSubnetResponse) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *BackupRestoreSubnetResponse) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *BackupRestoreSubnetResponse) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetCreatedAt

`func (o *BackupRestoreSubnetResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackupRestoreSubnetResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackupRestoreSubnetResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *BackupRestoreSubnetResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BackupRestoreSubnetResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BackupRestoreSubnetResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetGatewayIpAddress

`func (o *BackupRestoreSubnetResponse) GetGatewayIpAddress() string`

GetGatewayIpAddress returns the GatewayIpAddress field if non-nil, zero value otherwise.

### GetGatewayIpAddressOk

`func (o *BackupRestoreSubnetResponse) GetGatewayIpAddressOk() (*string, bool)`

GetGatewayIpAddressOk returns a tuple with the GatewayIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIpAddress

`func (o *BackupRestoreSubnetResponse) SetGatewayIpAddress(v string)`

SetGatewayIpAddress sets GatewayIpAddress field to given value.


### SetGatewayIpAddressNil

`func (o *BackupRestoreSubnetResponse) SetGatewayIpAddressNil(b bool)`

 SetGatewayIpAddressNil sets the value for GatewayIpAddress to be an explicit nil

### UnsetGatewayIpAddress
`func (o *BackupRestoreSubnetResponse) UnsetGatewayIpAddress()`

UnsetGatewayIpAddress ensures that no value is present for GatewayIpAddress, not even an explicit nil
### GetId

`func (o *BackupRestoreSubnetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupRestoreSubnetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupRestoreSubnetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *BackupRestoreSubnetResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BackupRestoreSubnetResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BackupRestoreSubnetResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *BackupRestoreSubnetResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BackupRestoreSubnetResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BackupRestoreSubnetResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *BackupRestoreSubnetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupRestoreSubnetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupRestoreSubnetResponse) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *BackupRestoreSubnetResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackupRestoreSubnetResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackupRestoreSubnetResponse) SetState(v string)`

SetState sets State field to given value.


### GetType

`func (o *BackupRestoreSubnetResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupRestoreSubnetResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupRestoreSubnetResponse) SetType(v string)`

SetType sets Type field to given value.


### GetVpcId

`func (o *BackupRestoreSubnetResponse) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BackupRestoreSubnetResponse) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BackupRestoreSubnetResponse) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *BackupRestoreSubnetResponse) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BackupRestoreSubnetResponse) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BackupRestoreSubnetResponse) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


