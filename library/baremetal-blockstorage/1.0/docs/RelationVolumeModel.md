# RelationVolumeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**VolumePermission**](VolumePermission.md) | Auth | [optional] 
**Id** | Pointer to **string** | Id | [optional] 
**Name** | Pointer to **string** | Volume name | [optional] 
**Purpose** | Pointer to [**BlockStoragePurpose**](BlockStoragePurpose.md) | Volume purpose | [optional] 
**Region** | Pointer to **string** | Region | [optional] 

## Methods

### NewRelationVolumeModel

`func NewRelationVolumeModel() *RelationVolumeModel`

NewRelationVolumeModel instantiates a new RelationVolumeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationVolumeModelWithDefaults

`func NewRelationVolumeModelWithDefaults() *RelationVolumeModel`

NewRelationVolumeModelWithDefaults instantiates a new RelationVolumeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *RelationVolumeModel) GetAuth() VolumePermission`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *RelationVolumeModel) GetAuthOk() (*VolumePermission, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *RelationVolumeModel) SetAuth(v VolumePermission)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *RelationVolumeModel) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetId

`func (o *RelationVolumeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationVolumeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationVolumeModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationVolumeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RelationVolumeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelationVolumeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelationVolumeModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RelationVolumeModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPurpose

`func (o *RelationVolumeModel) GetPurpose() BlockStoragePurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *RelationVolumeModel) GetPurposeOk() (*BlockStoragePurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *RelationVolumeModel) SetPurpose(v BlockStoragePurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *RelationVolumeModel) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetRegion

`func (o *RelationVolumeModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RelationVolumeModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RelationVolumeModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RelationVolumeModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


