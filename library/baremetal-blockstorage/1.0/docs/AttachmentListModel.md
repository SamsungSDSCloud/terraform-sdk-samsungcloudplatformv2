# AttachmentListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | Pointer to **string** | Object id | [optional] 
**ObjectType** | Pointer to [**BlockStorageAttachmentObjectType**](BlockStorageAttachmentObjectType.md) | Object type | [optional] 

## Methods

### NewAttachmentListModel

`func NewAttachmentListModel() *AttachmentListModel`

NewAttachmentListModel instantiates a new AttachmentListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentListModelWithDefaults

`func NewAttachmentListModelWithDefaults() *AttachmentListModel`

NewAttachmentListModelWithDefaults instantiates a new AttachmentListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *AttachmentListModel) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *AttachmentListModel) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *AttachmentListModel) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *AttachmentListModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *AttachmentListModel) GetObjectType() BlockStorageAttachmentObjectType`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AttachmentListModel) GetObjectTypeOk() (*BlockStorageAttachmentObjectType, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AttachmentListModel) SetObjectType(v BlockStorageAttachmentObjectType)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *AttachmentListModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


