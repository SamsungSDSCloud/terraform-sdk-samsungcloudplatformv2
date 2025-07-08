# SqlserverBlockStorageGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveLetter** | **string** | Drive letter | 
**Id** | **string** | ID | 
**Name** | **string** | Name | 
**RoleType** | [**BlockStorageGroupRoleType**](BlockStorageGroupRoleType.md) | Role type | 
**SizeGb** | **int32** | Size in GB | 
**VolumeType** | [**VolumeType**](VolumeType.md) | Volume type | 

## Methods

### NewSqlserverBlockStorageGroupResponse

`func NewSqlserverBlockStorageGroupResponse(driveLetter string, id string, name string, roleType BlockStorageGroupRoleType, sizeGb int32, volumeType VolumeType, ) *SqlserverBlockStorageGroupResponse`

NewSqlserverBlockStorageGroupResponse instantiates a new SqlserverBlockStorageGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlserverBlockStorageGroupResponseWithDefaults

`func NewSqlserverBlockStorageGroupResponseWithDefaults() *SqlserverBlockStorageGroupResponse`

NewSqlserverBlockStorageGroupResponseWithDefaults instantiates a new SqlserverBlockStorageGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveLetter

`func (o *SqlserverBlockStorageGroupResponse) GetDriveLetter() string`

GetDriveLetter returns the DriveLetter field if non-nil, zero value otherwise.

### GetDriveLetterOk

`func (o *SqlserverBlockStorageGroupResponse) GetDriveLetterOk() (*string, bool)`

GetDriveLetterOk returns a tuple with the DriveLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveLetter

`func (o *SqlserverBlockStorageGroupResponse) SetDriveLetter(v string)`

SetDriveLetter sets DriveLetter field to given value.


### GetId

`func (o *SqlserverBlockStorageGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SqlserverBlockStorageGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SqlserverBlockStorageGroupResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SqlserverBlockStorageGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlserverBlockStorageGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlserverBlockStorageGroupResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRoleType

`func (o *SqlserverBlockStorageGroupResponse) GetRoleType() BlockStorageGroupRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *SqlserverBlockStorageGroupResponse) GetRoleTypeOk() (*BlockStorageGroupRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *SqlserverBlockStorageGroupResponse) SetRoleType(v BlockStorageGroupRoleType)`

SetRoleType sets RoleType field to given value.


### GetSizeGb

`func (o *SqlserverBlockStorageGroupResponse) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *SqlserverBlockStorageGroupResponse) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *SqlserverBlockStorageGroupResponse) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.


### GetVolumeType

`func (o *SqlserverBlockStorageGroupResponse) GetVolumeType() VolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *SqlserverBlockStorageGroupResponse) GetVolumeTypeOk() (*VolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *SqlserverBlockStorageGroupResponse) SetVolumeType(v VolumeType)`

SetVolumeType sets VolumeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


