# VolumeServerAttachResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnTermination** | **bool** | When delete server, delete volumes together | 
**Device** | **string** | Device | 
**ServerId** | **string** | Server ID | 
**VolumeId** | **string** | Volume ID | 

## Methods

### NewVolumeServerAttachResponse

`func NewVolumeServerAttachResponse(deleteOnTermination bool, device string, serverId string, volumeId string, ) *VolumeServerAttachResponse`

NewVolumeServerAttachResponse instantiates a new VolumeServerAttachResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeServerAttachResponseWithDefaults

`func NewVolumeServerAttachResponseWithDefaults() *VolumeServerAttachResponse`

NewVolumeServerAttachResponseWithDefaults instantiates a new VolumeServerAttachResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnTermination

`func (o *VolumeServerAttachResponse) GetDeleteOnTermination() bool`

GetDeleteOnTermination returns the DeleteOnTermination field if non-nil, zero value otherwise.

### GetDeleteOnTerminationOk

`func (o *VolumeServerAttachResponse) GetDeleteOnTerminationOk() (*bool, bool)`

GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnTermination

`func (o *VolumeServerAttachResponse) SetDeleteOnTermination(v bool)`

SetDeleteOnTermination sets DeleteOnTermination field to given value.


### GetDevice

`func (o *VolumeServerAttachResponse) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VolumeServerAttachResponse) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VolumeServerAttachResponse) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetServerId

`func (o *VolumeServerAttachResponse) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *VolumeServerAttachResponse) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *VolumeServerAttachResponse) SetServerId(v string)`

SetServerId sets ServerId field to given value.


### GetVolumeId

`func (o *VolumeServerAttachResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeServerAttachResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeServerAttachResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


