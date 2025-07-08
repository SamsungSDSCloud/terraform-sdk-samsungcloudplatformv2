# VolumeGroupListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Account id | [optional] 
**CreatedAt** | Pointer to **time.Time** | Created at | [optional] 
**CreatedBy** | Pointer to **string** | Created by | [optional] 
**Id** | Pointer to **string** | Id | [optional] 
**ModifiedAt** | Pointer to **time.Time** | Modified at | [optional] 
**ModifiedBy** | Pointer to **string** | Modified by | [optional] 
**Name** | Pointer to **string** | Volume group name | [optional] 
**NumOfBlockStorages** | Pointer to **int32** | Number of volumes | [optional] [default to 1]
**Purpose** | Pointer to **string** | Purpose | [optional] 
**State** | Pointer to [**VolumeGroupState**](VolumeGroupState.md) | Current state | [optional] 

## Methods

### NewVolumeGroupListModel

`func NewVolumeGroupListModel() *VolumeGroupListModel`

NewVolumeGroupListModel instantiates a new VolumeGroupListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupListModelWithDefaults

`func NewVolumeGroupListModelWithDefaults() *VolumeGroupListModel`

NewVolumeGroupListModelWithDefaults instantiates a new VolumeGroupListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *VolumeGroupListModel) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VolumeGroupListModel) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VolumeGroupListModel) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *VolumeGroupListModel) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VolumeGroupListModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeGroupListModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeGroupListModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VolumeGroupListModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VolumeGroupListModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VolumeGroupListModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VolumeGroupListModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VolumeGroupListModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetId

`func (o *VolumeGroupListModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeGroupListModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeGroupListModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeGroupListModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *VolumeGroupListModel) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VolumeGroupListModel) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VolumeGroupListModel) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *VolumeGroupListModel) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *VolumeGroupListModel) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VolumeGroupListModel) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VolumeGroupListModel) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *VolumeGroupListModel) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *VolumeGroupListModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeGroupListModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeGroupListModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeGroupListModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumOfBlockStorages

`func (o *VolumeGroupListModel) GetNumOfBlockStorages() int32`

GetNumOfBlockStorages returns the NumOfBlockStorages field if non-nil, zero value otherwise.

### GetNumOfBlockStoragesOk

`func (o *VolumeGroupListModel) GetNumOfBlockStoragesOk() (*int32, bool)`

GetNumOfBlockStoragesOk returns a tuple with the NumOfBlockStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfBlockStorages

`func (o *VolumeGroupListModel) SetNumOfBlockStorages(v int32)`

SetNumOfBlockStorages sets NumOfBlockStorages field to given value.

### HasNumOfBlockStorages

`func (o *VolumeGroupListModel) HasNumOfBlockStorages() bool`

HasNumOfBlockStorages returns a boolean if a field has been set.

### GetPurpose

`func (o *VolumeGroupListModel) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *VolumeGroupListModel) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *VolumeGroupListModel) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *VolumeGroupListModel) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetState

`func (o *VolumeGroupListModel) GetState() VolumeGroupState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VolumeGroupListModel) GetStateOk() (*VolumeGroupState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VolumeGroupListModel) SetState(v VolumeGroupState)`

SetState sets State field to given value.

### HasState

`func (o *VolumeGroupListModel) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


