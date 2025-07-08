# ServerCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** | Image ID | 
**KeypairName** | **string** | Keypair name | 
**Lock** | Pointer to **NullableBool** |  | [optional] 
**MaxCount** | Pointer to **NullableInt32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** | Server name | 
**Networks** | [**[]Network**](Network.md) | List of Network | 
**ProductCategory** | Pointer to [**NullableServerProductCategory**](ServerProductCategory.md) |  | [optional] 
**ProductOffering** | Pointer to [**NullableServerProductOffering**](ServerProductOffering.md) |  | [optional] 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 
**ServerGroupId** | Pointer to **NullableString** |  | [optional] 
**ServerTypeId** | **string** | Server type ID | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**UserData** | Pointer to **NullableString** |  | [optional] 
**Volumes** | Pointer to [**[]Volume**](Volume.md) |  | [optional] 

## Methods

### NewServerCreateRequest

`func NewServerCreateRequest(imageId string, keypairName string, name string, networks []Network, serverTypeId string, ) *ServerCreateRequest`

NewServerCreateRequest instantiates a new ServerCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerCreateRequestWithDefaults

`func NewServerCreateRequestWithDefaults() *ServerCreateRequest`

NewServerCreateRequestWithDefaults instantiates a new ServerCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *ServerCreateRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ServerCreateRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ServerCreateRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetKeypairName

`func (o *ServerCreateRequest) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *ServerCreateRequest) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *ServerCreateRequest) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.


### GetLock

`func (o *ServerCreateRequest) GetLock() bool`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *ServerCreateRequest) GetLockOk() (*bool, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *ServerCreateRequest) SetLock(v bool)`

SetLock sets Lock field to given value.

### HasLock

`func (o *ServerCreateRequest) HasLock() bool`

HasLock returns a boolean if a field has been set.

### SetLockNil

`func (o *ServerCreateRequest) SetLockNil(b bool)`

 SetLockNil sets the value for Lock to be an explicit nil

### UnsetLock
`func (o *ServerCreateRequest) UnsetLock()`

UnsetLock ensures that no value is present for Lock, not even an explicit nil
### GetMaxCount

`func (o *ServerCreateRequest) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *ServerCreateRequest) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *ServerCreateRequest) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *ServerCreateRequest) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### SetMaxCountNil

`func (o *ServerCreateRequest) SetMaxCountNil(b bool)`

 SetMaxCountNil sets the value for MaxCount to be an explicit nil

### UnsetMaxCount
`func (o *ServerCreateRequest) UnsetMaxCount()`

UnsetMaxCount ensures that no value is present for MaxCount, not even an explicit nil
### GetMetadata

`func (o *ServerCreateRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServerCreateRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServerCreateRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServerCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ServerCreateRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ServerCreateRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *ServerCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNetworks

`func (o *ServerCreateRequest) GetNetworks() []Network`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ServerCreateRequest) GetNetworksOk() (*[]Network, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ServerCreateRequest) SetNetworks(v []Network)`

SetNetworks sets Networks field to given value.


### GetProductCategory

`func (o *ServerCreateRequest) GetProductCategory() ServerProductCategory`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *ServerCreateRequest) GetProductCategoryOk() (*ServerProductCategory, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *ServerCreateRequest) SetProductCategory(v ServerProductCategory)`

SetProductCategory sets ProductCategory field to given value.

### HasProductCategory

`func (o *ServerCreateRequest) HasProductCategory() bool`

HasProductCategory returns a boolean if a field has been set.

### SetProductCategoryNil

`func (o *ServerCreateRequest) SetProductCategoryNil(b bool)`

 SetProductCategoryNil sets the value for ProductCategory to be an explicit nil

### UnsetProductCategory
`func (o *ServerCreateRequest) UnsetProductCategory()`

UnsetProductCategory ensures that no value is present for ProductCategory, not even an explicit nil
### GetProductOffering

`func (o *ServerCreateRequest) GetProductOffering() ServerProductOffering`

GetProductOffering returns the ProductOffering field if non-nil, zero value otherwise.

### GetProductOfferingOk

`func (o *ServerCreateRequest) GetProductOfferingOk() (*ServerProductOffering, bool)`

GetProductOfferingOk returns a tuple with the ProductOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOffering

`func (o *ServerCreateRequest) SetProductOffering(v ServerProductOffering)`

SetProductOffering sets ProductOffering field to given value.

### HasProductOffering

`func (o *ServerCreateRequest) HasProductOffering() bool`

HasProductOffering returns a boolean if a field has been set.

### SetProductOfferingNil

`func (o *ServerCreateRequest) SetProductOfferingNil(b bool)`

 SetProductOfferingNil sets the value for ProductOffering to be an explicit nil

### UnsetProductOffering
`func (o *ServerCreateRequest) UnsetProductOffering()`

UnsetProductOffering ensures that no value is present for ProductOffering, not even an explicit nil
### GetSecurityGroups

`func (o *ServerCreateRequest) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ServerCreateRequest) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ServerCreateRequest) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ServerCreateRequest) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *ServerCreateRequest) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *ServerCreateRequest) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetServerGroupId

`func (o *ServerCreateRequest) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *ServerCreateRequest) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *ServerCreateRequest) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.

### HasServerGroupId

`func (o *ServerCreateRequest) HasServerGroupId() bool`

HasServerGroupId returns a boolean if a field has been set.

### SetServerGroupIdNil

`func (o *ServerCreateRequest) SetServerGroupIdNil(b bool)`

 SetServerGroupIdNil sets the value for ServerGroupId to be an explicit nil

### UnsetServerGroupId
`func (o *ServerCreateRequest) UnsetServerGroupId()`

UnsetServerGroupId ensures that no value is present for ServerGroupId, not even an explicit nil
### GetServerTypeId

`func (o *ServerCreateRequest) GetServerTypeId() string`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *ServerCreateRequest) GetServerTypeIdOk() (*string, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *ServerCreateRequest) SetServerTypeId(v string)`

SetServerTypeId sets ServerTypeId field to given value.


### GetTags

`func (o *ServerCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ServerCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ServerCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUserData

`func (o *ServerCreateRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *ServerCreateRequest) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *ServerCreateRequest) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *ServerCreateRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *ServerCreateRequest) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *ServerCreateRequest) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetVolumes

`func (o *ServerCreateRequest) GetVolumes() []Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ServerCreateRequest) GetVolumesOk() (*[]Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ServerCreateRequest) SetVolumes(v []Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ServerCreateRequest) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *ServerCreateRequest) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *ServerCreateRequest) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


