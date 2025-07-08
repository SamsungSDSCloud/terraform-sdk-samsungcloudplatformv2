# VolumeGroupMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Volume group id | [optional] 
**MemberVolumes** | Pointer to [**[]MemberVolume**](MemberVolume.md) | List of member volumes | [optional] 

## Methods

### NewVolumeGroupMemberResponse

`func NewVolumeGroupMemberResponse() *VolumeGroupMemberResponse`

NewVolumeGroupMemberResponse instantiates a new VolumeGroupMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupMemberResponseWithDefaults

`func NewVolumeGroupMemberResponseWithDefaults() *VolumeGroupMemberResponse`

NewVolumeGroupMemberResponseWithDefaults instantiates a new VolumeGroupMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VolumeGroupMemberResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeGroupMemberResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeGroupMemberResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeGroupMemberResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemberVolumes

`func (o *VolumeGroupMemberResponse) GetMemberVolumes() []MemberVolume`

GetMemberVolumes returns the MemberVolumes field if non-nil, zero value otherwise.

### GetMemberVolumesOk

`func (o *VolumeGroupMemberResponse) GetMemberVolumesOk() (*[]MemberVolume, bool)`

GetMemberVolumesOk returns a tuple with the MemberVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberVolumes

`func (o *VolumeGroupMemberResponse) SetMemberVolumes(v []MemberVolume)`

SetMemberVolumes sets MemberVolumes field to given value.

### HasMemberVolumes

`func (o *VolumeGroupMemberResponse) HasMemberVolumes() bool`

HasMemberVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


