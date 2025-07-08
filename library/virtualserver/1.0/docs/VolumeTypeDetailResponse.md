# VolumeTypeDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**ExtraSpecs** | [**NullableVolumeTypeExtraSpecs**](VolumeTypeExtraSpecs.md) |  | 
**Id** | **string** | Volume type ID | 
**Name** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVolumeTypeDetailResponse

`func NewVolumeTypeDetailResponse(extraSpecs NullableVolumeTypeExtraSpecs, id string, ) *VolumeTypeDetailResponse`

NewVolumeTypeDetailResponse instantiates a new VolumeTypeDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeTypeDetailResponseWithDefaults

`func NewVolumeTypeDetailResponseWithDefaults() *VolumeTypeDetailResponse`

NewVolumeTypeDetailResponseWithDefaults instantiates a new VolumeTypeDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VolumeTypeDetailResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeTypeDetailResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeTypeDetailResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeTypeDetailResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VolumeTypeDetailResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VolumeTypeDetailResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExtraSpecs

`func (o *VolumeTypeDetailResponse) GetExtraSpecs() VolumeTypeExtraSpecs`

GetExtraSpecs returns the ExtraSpecs field if non-nil, zero value otherwise.

### GetExtraSpecsOk

`func (o *VolumeTypeDetailResponse) GetExtraSpecsOk() (*VolumeTypeExtraSpecs, bool)`

GetExtraSpecsOk returns a tuple with the ExtraSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraSpecs

`func (o *VolumeTypeDetailResponse) SetExtraSpecs(v VolumeTypeExtraSpecs)`

SetExtraSpecs sets ExtraSpecs field to given value.


### SetExtraSpecsNil

`func (o *VolumeTypeDetailResponse) SetExtraSpecsNil(b bool)`

 SetExtraSpecsNil sets the value for ExtraSpecs to be an explicit nil

### UnsetExtraSpecs
`func (o *VolumeTypeDetailResponse) UnsetExtraSpecs()`

UnsetExtraSpecs ensures that no value is present for ExtraSpecs, not even an explicit nil
### GetId

`func (o *VolumeTypeDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeTypeDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeTypeDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VolumeTypeDetailResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeTypeDetailResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeTypeDetailResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeTypeDetailResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VolumeTypeDetailResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VolumeTypeDetailResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


