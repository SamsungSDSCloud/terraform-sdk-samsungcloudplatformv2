# ServerVolumesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumes** | [**[]ServersVolumeResponse**](ServersVolumeResponse.md) | Volumes | 

## Methods

### NewServerVolumesResponse

`func NewServerVolumesResponse(volumes []ServersVolumeResponse, ) *ServerVolumesResponse`

NewServerVolumesResponse instantiates a new ServerVolumesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerVolumesResponseWithDefaults

`func NewServerVolumesResponseWithDefaults() *ServerVolumesResponse`

NewServerVolumesResponseWithDefaults instantiates a new ServerVolumesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumes

`func (o *ServerVolumesResponse) GetVolumes() []ServersVolumeResponse`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ServerVolumesResponse) GetVolumesOk() (*[]ServersVolumeResponse, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ServerVolumesResponse) SetVolumes(v []ServersVolumeResponse)`

SetVolumes sets Volumes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


