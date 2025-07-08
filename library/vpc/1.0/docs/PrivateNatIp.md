# PrivateNatIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachedResourceId** | Pointer to **NullableString** |  | [optional] 
**AttachedResourceName** | Pointer to **NullableString** |  | [optional] 
**AttachedResourceType** | Pointer to [**NullablePrivateNatIpAttachedResourceType**](PrivateNatIpAttachedResourceType.md) |  | [optional] 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | **string** | Private NAT IP ID | 
**IpAddress** | **string** | Private NAT IP Address | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**PrivateNatId** | **string** | Private NAT ID | 
**State** | [**PrivateNatIpState**](PrivateNatIpState.md) | Private NAT IP State | 

## Methods

### NewPrivateNatIp

`func NewPrivateNatIp(createdAt time.Time, createdBy string, id string, ipAddress string, modifiedAt time.Time, modifiedBy string, privateNatId string, state PrivateNatIpState, ) *PrivateNatIp`

NewPrivateNatIp instantiates a new PrivateNatIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNatIpWithDefaults

`func NewPrivateNatIpWithDefaults() *PrivateNatIp`

NewPrivateNatIpWithDefaults instantiates a new PrivateNatIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachedResourceId

`func (o *PrivateNatIp) GetAttachedResourceId() string`

GetAttachedResourceId returns the AttachedResourceId field if non-nil, zero value otherwise.

### GetAttachedResourceIdOk

`func (o *PrivateNatIp) GetAttachedResourceIdOk() (*string, bool)`

GetAttachedResourceIdOk returns a tuple with the AttachedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedResourceId

`func (o *PrivateNatIp) SetAttachedResourceId(v string)`

SetAttachedResourceId sets AttachedResourceId field to given value.

### HasAttachedResourceId

`func (o *PrivateNatIp) HasAttachedResourceId() bool`

HasAttachedResourceId returns a boolean if a field has been set.

### SetAttachedResourceIdNil

`func (o *PrivateNatIp) SetAttachedResourceIdNil(b bool)`

 SetAttachedResourceIdNil sets the value for AttachedResourceId to be an explicit nil

### UnsetAttachedResourceId
`func (o *PrivateNatIp) UnsetAttachedResourceId()`

UnsetAttachedResourceId ensures that no value is present for AttachedResourceId, not even an explicit nil
### GetAttachedResourceName

`func (o *PrivateNatIp) GetAttachedResourceName() string`

GetAttachedResourceName returns the AttachedResourceName field if non-nil, zero value otherwise.

### GetAttachedResourceNameOk

`func (o *PrivateNatIp) GetAttachedResourceNameOk() (*string, bool)`

GetAttachedResourceNameOk returns a tuple with the AttachedResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedResourceName

`func (o *PrivateNatIp) SetAttachedResourceName(v string)`

SetAttachedResourceName sets AttachedResourceName field to given value.

### HasAttachedResourceName

`func (o *PrivateNatIp) HasAttachedResourceName() bool`

HasAttachedResourceName returns a boolean if a field has been set.

### SetAttachedResourceNameNil

`func (o *PrivateNatIp) SetAttachedResourceNameNil(b bool)`

 SetAttachedResourceNameNil sets the value for AttachedResourceName to be an explicit nil

### UnsetAttachedResourceName
`func (o *PrivateNatIp) UnsetAttachedResourceName()`

UnsetAttachedResourceName ensures that no value is present for AttachedResourceName, not even an explicit nil
### GetAttachedResourceType

`func (o *PrivateNatIp) GetAttachedResourceType() PrivateNatIpAttachedResourceType`

GetAttachedResourceType returns the AttachedResourceType field if non-nil, zero value otherwise.

### GetAttachedResourceTypeOk

`func (o *PrivateNatIp) GetAttachedResourceTypeOk() (*PrivateNatIpAttachedResourceType, bool)`

GetAttachedResourceTypeOk returns a tuple with the AttachedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedResourceType

`func (o *PrivateNatIp) SetAttachedResourceType(v PrivateNatIpAttachedResourceType)`

SetAttachedResourceType sets AttachedResourceType field to given value.

### HasAttachedResourceType

`func (o *PrivateNatIp) HasAttachedResourceType() bool`

HasAttachedResourceType returns a boolean if a field has been set.

### SetAttachedResourceTypeNil

`func (o *PrivateNatIp) SetAttachedResourceTypeNil(b bool)`

 SetAttachedResourceTypeNil sets the value for AttachedResourceType to be an explicit nil

### UnsetAttachedResourceType
`func (o *PrivateNatIp) UnsetAttachedResourceType()`

UnsetAttachedResourceType ensures that no value is present for AttachedResourceType, not even an explicit nil
### GetCreatedAt

`func (o *PrivateNatIp) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrivateNatIp) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrivateNatIp) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *PrivateNatIp) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PrivateNatIp) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PrivateNatIp) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *PrivateNatIp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateNatIp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateNatIp) SetId(v string)`

SetId sets Id field to given value.


### GetIpAddress

`func (o *PrivateNatIp) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PrivateNatIp) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PrivateNatIp) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetModifiedAt

`func (o *PrivateNatIp) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PrivateNatIp) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PrivateNatIp) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *PrivateNatIp) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PrivateNatIp) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PrivateNatIp) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetPrivateNatId

`func (o *PrivateNatIp) GetPrivateNatId() string`

GetPrivateNatId returns the PrivateNatId field if non-nil, zero value otherwise.

### GetPrivateNatIdOk

`func (o *PrivateNatIp) GetPrivateNatIdOk() (*string, bool)`

GetPrivateNatIdOk returns a tuple with the PrivateNatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNatId

`func (o *PrivateNatIp) SetPrivateNatId(v string)`

SetPrivateNatId sets PrivateNatId field to given value.


### GetState

`func (o *PrivateNatIp) GetState() PrivateNatIpState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PrivateNatIp) GetStateOk() (*PrivateNatIpState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PrivateNatIp) SetState(v PrivateNatIpState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


