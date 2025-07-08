# RelationVolumeGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**VolumePermission**](VolumePermission.md) | Auth | [optional] 
**Id** | Pointer to **string** | Id | [optional] 
**Name** | Pointer to **string** | Volume group name | [optional] 
**Purpose** | Pointer to [**VolumeGroupPurpose**](VolumeGroupPurpose.md) | Volume group purpose | [optional] 
**Region** | Pointer to **string** | Region | [optional] 

## Methods

### NewRelationVolumeGroupModel

`func NewRelationVolumeGroupModel() *RelationVolumeGroupModel`

NewRelationVolumeGroupModel instantiates a new RelationVolumeGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationVolumeGroupModelWithDefaults

`func NewRelationVolumeGroupModelWithDefaults() *RelationVolumeGroupModel`

NewRelationVolumeGroupModelWithDefaults instantiates a new RelationVolumeGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *RelationVolumeGroupModel) GetAuth() VolumePermission`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *RelationVolumeGroupModel) GetAuthOk() (*VolumePermission, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *RelationVolumeGroupModel) SetAuth(v VolumePermission)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *RelationVolumeGroupModel) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetId

`func (o *RelationVolumeGroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationVolumeGroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationVolumeGroupModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationVolumeGroupModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RelationVolumeGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelationVolumeGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelationVolumeGroupModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RelationVolumeGroupModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPurpose

`func (o *RelationVolumeGroupModel) GetPurpose() VolumeGroupPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *RelationVolumeGroupModel) GetPurposeOk() (*VolumeGroupPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *RelationVolumeGroupModel) SetPurpose(v VolumeGroupPurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *RelationVolumeGroupModel) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetRegion

`func (o *RelationVolumeGroupModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RelationVolumeGroupModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RelationVolumeGroupModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RelationVolumeGroupModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


