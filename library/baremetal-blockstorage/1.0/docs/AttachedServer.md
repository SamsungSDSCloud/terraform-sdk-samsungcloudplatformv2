# AttachedServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Server id | [optional] 
**ImageName** | Pointer to **string** | Image name | [optional] 
**Name** | Pointer to **string** | Server name | [optional] 
**Srn** | Pointer to **string** | SRN | [optional] 
**State** | Pointer to **string** | Server state | [optional] 
**Type** | Pointer to [**BlockStorageAttachmentObjectType**](BlockStorageAttachmentObjectType.md) | Server type | [optional] 

## Methods

### NewAttachedServer

`func NewAttachedServer() *AttachedServer`

NewAttachedServer instantiates a new AttachedServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachedServerWithDefaults

`func NewAttachedServerWithDefaults() *AttachedServer`

NewAttachedServerWithDefaults instantiates a new AttachedServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttachedServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachedServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachedServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttachedServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageName

`func (o *AttachedServer) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *AttachedServer) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *AttachedServer) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *AttachedServer) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetName

`func (o *AttachedServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttachedServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttachedServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttachedServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSrn

`func (o *AttachedServer) GetSrn() string`

GetSrn returns the Srn field if non-nil, zero value otherwise.

### GetSrnOk

`func (o *AttachedServer) GetSrnOk() (*string, bool)`

GetSrnOk returns a tuple with the Srn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrn

`func (o *AttachedServer) SetSrn(v string)`

SetSrn sets Srn field to given value.

### HasSrn

`func (o *AttachedServer) HasSrn() bool`

HasSrn returns a boolean if a field has been set.

### GetState

`func (o *AttachedServer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AttachedServer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AttachedServer) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AttachedServer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *AttachedServer) GetType() BlockStorageAttachmentObjectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachedServer) GetTypeOk() (*BlockStorageAttachmentObjectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachedServer) SetType(v BlockStorageAttachmentObjectType)`

SetType sets Type field to given value.

### HasType

`func (o *AttachedServer) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


