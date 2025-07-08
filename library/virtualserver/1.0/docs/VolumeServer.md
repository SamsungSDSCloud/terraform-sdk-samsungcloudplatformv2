# VolumeServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachedAt** | **time.Time** | Attached at | 
**Device** | **string** | Device | 
**Id** | **string** | Server ID | 

## Methods

### NewVolumeServer

`func NewVolumeServer(attachedAt time.Time, device string, id string, ) *VolumeServer`

NewVolumeServer instantiates a new VolumeServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeServerWithDefaults

`func NewVolumeServerWithDefaults() *VolumeServer`

NewVolumeServerWithDefaults instantiates a new VolumeServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachedAt

`func (o *VolumeServer) GetAttachedAt() time.Time`

GetAttachedAt returns the AttachedAt field if non-nil, zero value otherwise.

### GetAttachedAtOk

`func (o *VolumeServer) GetAttachedAtOk() (*time.Time, bool)`

GetAttachedAtOk returns a tuple with the AttachedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedAt

`func (o *VolumeServer) SetAttachedAt(v time.Time)`

SetAttachedAt sets AttachedAt field to given value.


### GetDevice

`func (o *VolumeServer) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VolumeServer) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VolumeServer) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetId

`func (o *VolumeServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeServer) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


