# Publicip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AttachedResourceId** | Pointer to **NullableString** |  | [optional] 
**AttachedResourceName** | Pointer to **NullableString** |  | [optional] 
**AttachedResourceType** | Pointer to [**NullablePublicipAttachedResourceType**](PublicipAttachedResourceType.md) |  | [optional] 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | PublicIP ID | 
**IpAddress** | **string** | IP Address | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**State** | [**PublicipState**](PublicipState.md) | PublicIP State | 
**Type** | [**PublicipType**](PublicipType.md) | PublicIP Type | 

## Methods

### NewPublicip

`func NewPublicip(accountId string, createdAt time.Time, createdBy string, id string, ipAddress string, modifiedAt time.Time, modifiedBy string, state PublicipState, type_ PublicipType, ) *Publicip`

NewPublicip instantiates a new Publicip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicipWithDefaults

`func NewPublicipWithDefaults() *Publicip`

NewPublicipWithDefaults instantiates a new Publicip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Publicip) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Publicip) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Publicip) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAttachedResourceId

`func (o *Publicip) GetAttachedResourceId() string`

GetAttachedResourceId returns the AttachedResourceId field if non-nil, zero value otherwise.

### GetAttachedResourceIdOk

`func (o *Publicip) GetAttachedResourceIdOk() (*string, bool)`

GetAttachedResourceIdOk returns a tuple with the AttachedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedResourceId

`func (o *Publicip) SetAttachedResourceId(v string)`

SetAttachedResourceId sets AttachedResourceId field to given value.

### HasAttachedResourceId

`func (o *Publicip) HasAttachedResourceId() bool`

HasAttachedResourceId returns a boolean if a field has been set.

### SetAttachedResourceIdNil

`func (o *Publicip) SetAttachedResourceIdNil(b bool)`

 SetAttachedResourceIdNil sets the value for AttachedResourceId to be an explicit nil

### UnsetAttachedResourceId
`func (o *Publicip) UnsetAttachedResourceId()`

UnsetAttachedResourceId ensures that no value is present for AttachedResourceId, not even an explicit nil
### GetAttachedResourceName

`func (o *Publicip) GetAttachedResourceName() string`

GetAttachedResourceName returns the AttachedResourceName field if non-nil, zero value otherwise.

### GetAttachedResourceNameOk

`func (o *Publicip) GetAttachedResourceNameOk() (*string, bool)`

GetAttachedResourceNameOk returns a tuple with the AttachedResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedResourceName

`func (o *Publicip) SetAttachedResourceName(v string)`

SetAttachedResourceName sets AttachedResourceName field to given value.

### HasAttachedResourceName

`func (o *Publicip) HasAttachedResourceName() bool`

HasAttachedResourceName returns a boolean if a field has been set.

### SetAttachedResourceNameNil

`func (o *Publicip) SetAttachedResourceNameNil(b bool)`

 SetAttachedResourceNameNil sets the value for AttachedResourceName to be an explicit nil

### UnsetAttachedResourceName
`func (o *Publicip) UnsetAttachedResourceName()`

UnsetAttachedResourceName ensures that no value is present for AttachedResourceName, not even an explicit nil
### GetAttachedResourceType

`func (o *Publicip) GetAttachedResourceType() PublicipAttachedResourceType`

GetAttachedResourceType returns the AttachedResourceType field if non-nil, zero value otherwise.

### GetAttachedResourceTypeOk

`func (o *Publicip) GetAttachedResourceTypeOk() (*PublicipAttachedResourceType, bool)`

GetAttachedResourceTypeOk returns a tuple with the AttachedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedResourceType

`func (o *Publicip) SetAttachedResourceType(v PublicipAttachedResourceType)`

SetAttachedResourceType sets AttachedResourceType field to given value.

### HasAttachedResourceType

`func (o *Publicip) HasAttachedResourceType() bool`

HasAttachedResourceType returns a boolean if a field has been set.

### SetAttachedResourceTypeNil

`func (o *Publicip) SetAttachedResourceTypeNil(b bool)`

 SetAttachedResourceTypeNil sets the value for AttachedResourceType to be an explicit nil

### UnsetAttachedResourceType
`func (o *Publicip) UnsetAttachedResourceType()`

UnsetAttachedResourceType ensures that no value is present for AttachedResourceType, not even an explicit nil
### GetCreatedAt

`func (o *Publicip) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Publicip) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Publicip) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Publicip) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Publicip) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Publicip) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *Publicip) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Publicip) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Publicip) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Publicip) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Publicip) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Publicip) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *Publicip) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Publicip) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Publicip) SetId(v string)`

SetId sets Id field to given value.


### GetIpAddress

`func (o *Publicip) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Publicip) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Publicip) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetModifiedAt

`func (o *Publicip) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Publicip) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Publicip) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *Publicip) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Publicip) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Publicip) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetState

`func (o *Publicip) GetState() PublicipState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Publicip) GetStateOk() (*PublicipState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Publicip) SetState(v PublicipState)`

SetState sets State field to given value.


### GetType

`func (o *Publicip) GetType() PublicipType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Publicip) GetTypeOk() (*PublicipType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Publicip) SetType(v PublicipType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


