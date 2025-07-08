# VolumeTransferListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transfers** | [**[]VolumeTransferAcceptResponse**](VolumeTransferAcceptResponse.md) |  | 

## Methods

### NewVolumeTransferListResponse

`func NewVolumeTransferListResponse(transfers []VolumeTransferAcceptResponse, ) *VolumeTransferListResponse`

NewVolumeTransferListResponse instantiates a new VolumeTransferListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeTransferListResponseWithDefaults

`func NewVolumeTransferListResponseWithDefaults() *VolumeTransferListResponse`

NewVolumeTransferListResponseWithDefaults instantiates a new VolumeTransferListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransfers

`func (o *VolumeTransferListResponse) GetTransfers() []VolumeTransferAcceptResponse`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *VolumeTransferListResponse) GetTransfersOk() (*[]VolumeTransferAcceptResponse, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *VolumeTransferListResponse) SetTransfers(v []VolumeTransferAcceptResponse)`

SetTransfers sets Transfers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


