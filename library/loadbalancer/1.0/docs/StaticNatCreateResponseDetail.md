# StaticNatCreateResponseDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | account_id | 
**ActionType** | **string** | action_type | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | **string** | description | 
**ExternalIpAddress** | **string** | external_ip_address | 
**Id** | **string** | ID | 
**InternalIpAddress** | **string** | internal_ip_address | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | name | 
**OwnerId** | **string** | owner_id | 
**OwnerName** | **string** | owner_name | 
**OwnerType** | **string** | owner_type | 
**PublicipId** | **string** | publicip_id | 
**ServiceIpPortId** | **string** | service_ip_port_id | 
**State** | **string** | state | 
**SubnetId** | **string** | subnet_id | 
**Type** | **string** | type | 
**VpcId** | **string** | vpc_id | 

## Methods

### NewStaticNatCreateResponseDetail

`func NewStaticNatCreateResponseDetail(accountId string, actionType string, createdAt time.Time, createdBy string, description string, externalIpAddress string, id string, internalIpAddress string, modifiedAt time.Time, modifiedBy string, name string, ownerId string, ownerName string, ownerType string, publicipId string, serviceIpPortId string, state string, subnetId string, type_ string, vpcId string, ) *StaticNatCreateResponseDetail`

NewStaticNatCreateResponseDetail instantiates a new StaticNatCreateResponseDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticNatCreateResponseDetailWithDefaults

`func NewStaticNatCreateResponseDetailWithDefaults() *StaticNatCreateResponseDetail`

NewStaticNatCreateResponseDetailWithDefaults instantiates a new StaticNatCreateResponseDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *StaticNatCreateResponseDetail) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StaticNatCreateResponseDetail) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StaticNatCreateResponseDetail) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetActionType

`func (o *StaticNatCreateResponseDetail) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *StaticNatCreateResponseDetail) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *StaticNatCreateResponseDetail) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetCreatedAt

`func (o *StaticNatCreateResponseDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StaticNatCreateResponseDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StaticNatCreateResponseDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *StaticNatCreateResponseDetail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *StaticNatCreateResponseDetail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *StaticNatCreateResponseDetail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *StaticNatCreateResponseDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StaticNatCreateResponseDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StaticNatCreateResponseDetail) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExternalIpAddress

`func (o *StaticNatCreateResponseDetail) GetExternalIpAddress() string`

GetExternalIpAddress returns the ExternalIpAddress field if non-nil, zero value otherwise.

### GetExternalIpAddressOk

`func (o *StaticNatCreateResponseDetail) GetExternalIpAddressOk() (*string, bool)`

GetExternalIpAddressOk returns a tuple with the ExternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpAddress

`func (o *StaticNatCreateResponseDetail) SetExternalIpAddress(v string)`

SetExternalIpAddress sets ExternalIpAddress field to given value.


### GetId

`func (o *StaticNatCreateResponseDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StaticNatCreateResponseDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StaticNatCreateResponseDetail) SetId(v string)`

SetId sets Id field to given value.


### GetInternalIpAddress

`func (o *StaticNatCreateResponseDetail) GetInternalIpAddress() string`

GetInternalIpAddress returns the InternalIpAddress field if non-nil, zero value otherwise.

### GetInternalIpAddressOk

`func (o *StaticNatCreateResponseDetail) GetInternalIpAddressOk() (*string, bool)`

GetInternalIpAddressOk returns a tuple with the InternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIpAddress

`func (o *StaticNatCreateResponseDetail) SetInternalIpAddress(v string)`

SetInternalIpAddress sets InternalIpAddress field to given value.


### GetModifiedAt

`func (o *StaticNatCreateResponseDetail) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *StaticNatCreateResponseDetail) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *StaticNatCreateResponseDetail) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *StaticNatCreateResponseDetail) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *StaticNatCreateResponseDetail) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *StaticNatCreateResponseDetail) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *StaticNatCreateResponseDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StaticNatCreateResponseDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StaticNatCreateResponseDetail) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *StaticNatCreateResponseDetail) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *StaticNatCreateResponseDetail) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *StaticNatCreateResponseDetail) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetOwnerName

`func (o *StaticNatCreateResponseDetail) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *StaticNatCreateResponseDetail) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *StaticNatCreateResponseDetail) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.


### GetOwnerType

`func (o *StaticNatCreateResponseDetail) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *StaticNatCreateResponseDetail) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *StaticNatCreateResponseDetail) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.


### GetPublicipId

`func (o *StaticNatCreateResponseDetail) GetPublicipId() string`

GetPublicipId returns the PublicipId field if non-nil, zero value otherwise.

### GetPublicipIdOk

`func (o *StaticNatCreateResponseDetail) GetPublicipIdOk() (*string, bool)`

GetPublicipIdOk returns a tuple with the PublicipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicipId

`func (o *StaticNatCreateResponseDetail) SetPublicipId(v string)`

SetPublicipId sets PublicipId field to given value.


### GetServiceIpPortId

`func (o *StaticNatCreateResponseDetail) GetServiceIpPortId() string`

GetServiceIpPortId returns the ServiceIpPortId field if non-nil, zero value otherwise.

### GetServiceIpPortIdOk

`func (o *StaticNatCreateResponseDetail) GetServiceIpPortIdOk() (*string, bool)`

GetServiceIpPortIdOk returns a tuple with the ServiceIpPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIpPortId

`func (o *StaticNatCreateResponseDetail) SetServiceIpPortId(v string)`

SetServiceIpPortId sets ServiceIpPortId field to given value.


### GetState

`func (o *StaticNatCreateResponseDetail) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StaticNatCreateResponseDetail) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StaticNatCreateResponseDetail) SetState(v string)`

SetState sets State field to given value.


### GetSubnetId

`func (o *StaticNatCreateResponseDetail) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *StaticNatCreateResponseDetail) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *StaticNatCreateResponseDetail) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetType

`func (o *StaticNatCreateResponseDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StaticNatCreateResponseDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StaticNatCreateResponseDetail) SetType(v string)`

SetType sets Type field to given value.


### GetVpcId

`func (o *StaticNatCreateResponseDetail) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *StaticNatCreateResponseDetail) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *StaticNatCreateResponseDetail) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


