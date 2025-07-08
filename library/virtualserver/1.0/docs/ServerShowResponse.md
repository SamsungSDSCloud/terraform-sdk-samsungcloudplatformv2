# ServerShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**Addresses** | Pointer to [**[]ServerShowResponseAddress**](ServerShowResponseAddress.md) |  | [optional] 
**AutoScalingGroupId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | Created at | 
**CreatedBy** | **string** | Created by | 
**DiskConfig** | **string** | Disk config | 
**Id** | **string** | Server ID | 
**ImageId** | Pointer to **NullableString** |  | [optional] 
**IsMarketplace** | Pointer to **NullableBool** |  | [optional] 
**KeypairName** | Pointer to **NullableString** |  | [optional] 
**LaunchConfigurationId** | Pointer to **NullableString** |  | [optional] 
**Locked** | **bool** | Locked | 
**Metadata** | **map[string]interface{}** | Metadata | 
**ModifiedAt** | **time.Time** | Modified at | 
**Name** | **string** | Server name | 
**PlannedComputeOsType** | Pointer to [**NullablePlannedComputeOsType**](PlannedComputeOsType.md) |  | [optional] 
**ProductCategory** | Pointer to [**NullableServerProductCategory**](ServerProductCategory.md) |  | [optional] 
**ProductOffering** | Pointer to [**NullableServerProductOffering**](ServerProductOffering.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]ServerShowResponseSecurityGroup**](ServerShowResponseSecurityGroup.md) |  | [optional] 
**ServerGroupId** | Pointer to **NullableString** |  | [optional] 
**ServerType** | [**ServerShowResponseServerType**](ServerShowResponseServerType.md) | Server type | 
**State** | **string** | Server state | 
**Volumes** | [**[]ServerShowResponseVolume**](ServerShowResponseVolume.md) | Volumes | 
**VpcId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServerShowResponse

`func NewServerShowResponse(accountId string, createdAt time.Time, createdBy string, diskConfig string, id string, locked bool, metadata map[string]interface{}, modifiedAt time.Time, name string, serverType ServerShowResponseServerType, state string, volumes []ServerShowResponseVolume, ) *ServerShowResponse`

NewServerShowResponse instantiates a new ServerShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerShowResponseWithDefaults

`func NewServerShowResponseWithDefaults() *ServerShowResponse`

NewServerShowResponseWithDefaults instantiates a new ServerShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ServerShowResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ServerShowResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ServerShowResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAddresses

`func (o *ServerShowResponse) GetAddresses() []ServerShowResponseAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ServerShowResponse) GetAddressesOk() (*[]ServerShowResponseAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ServerShowResponse) SetAddresses(v []ServerShowResponseAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *ServerShowResponse) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *ServerShowResponse) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *ServerShowResponse) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetAutoScalingGroupId

`func (o *ServerShowResponse) GetAutoScalingGroupId() string`

GetAutoScalingGroupId returns the AutoScalingGroupId field if non-nil, zero value otherwise.

### GetAutoScalingGroupIdOk

`func (o *ServerShowResponse) GetAutoScalingGroupIdOk() (*string, bool)`

GetAutoScalingGroupIdOk returns a tuple with the AutoScalingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroupId

`func (o *ServerShowResponse) SetAutoScalingGroupId(v string)`

SetAutoScalingGroupId sets AutoScalingGroupId field to given value.

### HasAutoScalingGroupId

`func (o *ServerShowResponse) HasAutoScalingGroupId() bool`

HasAutoScalingGroupId returns a boolean if a field has been set.

### SetAutoScalingGroupIdNil

`func (o *ServerShowResponse) SetAutoScalingGroupIdNil(b bool)`

 SetAutoScalingGroupIdNil sets the value for AutoScalingGroupId to be an explicit nil

### UnsetAutoScalingGroupId
`func (o *ServerShowResponse) UnsetAutoScalingGroupId()`

UnsetAutoScalingGroupId ensures that no value is present for AutoScalingGroupId, not even an explicit nil
### GetCreatedAt

`func (o *ServerShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ServerShowResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServerShowResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServerShowResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDiskConfig

`func (o *ServerShowResponse) GetDiskConfig() string`

GetDiskConfig returns the DiskConfig field if non-nil, zero value otherwise.

### GetDiskConfigOk

`func (o *ServerShowResponse) GetDiskConfigOk() (*string, bool)`

GetDiskConfigOk returns a tuple with the DiskConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskConfig

`func (o *ServerShowResponse) SetDiskConfig(v string)`

SetDiskConfig sets DiskConfig field to given value.


### GetId

`func (o *ServerShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetImageId

`func (o *ServerShowResponse) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ServerShowResponse) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ServerShowResponse) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ServerShowResponse) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *ServerShowResponse) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *ServerShowResponse) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetIsMarketplace

`func (o *ServerShowResponse) GetIsMarketplace() bool`

GetIsMarketplace returns the IsMarketplace field if non-nil, zero value otherwise.

### GetIsMarketplaceOk

`func (o *ServerShowResponse) GetIsMarketplaceOk() (*bool, bool)`

GetIsMarketplaceOk returns a tuple with the IsMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketplace

`func (o *ServerShowResponse) SetIsMarketplace(v bool)`

SetIsMarketplace sets IsMarketplace field to given value.

### HasIsMarketplace

`func (o *ServerShowResponse) HasIsMarketplace() bool`

HasIsMarketplace returns a boolean if a field has been set.

### SetIsMarketplaceNil

`func (o *ServerShowResponse) SetIsMarketplaceNil(b bool)`

 SetIsMarketplaceNil sets the value for IsMarketplace to be an explicit nil

### UnsetIsMarketplace
`func (o *ServerShowResponse) UnsetIsMarketplace()`

UnsetIsMarketplace ensures that no value is present for IsMarketplace, not even an explicit nil
### GetKeypairName

`func (o *ServerShowResponse) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *ServerShowResponse) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *ServerShowResponse) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.

### HasKeypairName

`func (o *ServerShowResponse) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### SetKeypairNameNil

`func (o *ServerShowResponse) SetKeypairNameNil(b bool)`

 SetKeypairNameNil sets the value for KeypairName to be an explicit nil

### UnsetKeypairName
`func (o *ServerShowResponse) UnsetKeypairName()`

UnsetKeypairName ensures that no value is present for KeypairName, not even an explicit nil
### GetLaunchConfigurationId

`func (o *ServerShowResponse) GetLaunchConfigurationId() string`

GetLaunchConfigurationId returns the LaunchConfigurationId field if non-nil, zero value otherwise.

### GetLaunchConfigurationIdOk

`func (o *ServerShowResponse) GetLaunchConfigurationIdOk() (*string, bool)`

GetLaunchConfigurationIdOk returns a tuple with the LaunchConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchConfigurationId

`func (o *ServerShowResponse) SetLaunchConfigurationId(v string)`

SetLaunchConfigurationId sets LaunchConfigurationId field to given value.

### HasLaunchConfigurationId

`func (o *ServerShowResponse) HasLaunchConfigurationId() bool`

HasLaunchConfigurationId returns a boolean if a field has been set.

### SetLaunchConfigurationIdNil

`func (o *ServerShowResponse) SetLaunchConfigurationIdNil(b bool)`

 SetLaunchConfigurationIdNil sets the value for LaunchConfigurationId to be an explicit nil

### UnsetLaunchConfigurationId
`func (o *ServerShowResponse) UnsetLaunchConfigurationId()`

UnsetLaunchConfigurationId ensures that no value is present for LaunchConfigurationId, not even an explicit nil
### GetLocked

`func (o *ServerShowResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ServerShowResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ServerShowResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetMetadata

`func (o *ServerShowResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServerShowResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServerShowResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetModifiedAt

`func (o *ServerShowResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ServerShowResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ServerShowResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetName

`func (o *ServerShowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerShowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerShowResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPlannedComputeOsType

`func (o *ServerShowResponse) GetPlannedComputeOsType() PlannedComputeOsType`

GetPlannedComputeOsType returns the PlannedComputeOsType field if non-nil, zero value otherwise.

### GetPlannedComputeOsTypeOk

`func (o *ServerShowResponse) GetPlannedComputeOsTypeOk() (*PlannedComputeOsType, bool)`

GetPlannedComputeOsTypeOk returns a tuple with the PlannedComputeOsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedComputeOsType

`func (o *ServerShowResponse) SetPlannedComputeOsType(v PlannedComputeOsType)`

SetPlannedComputeOsType sets PlannedComputeOsType field to given value.

### HasPlannedComputeOsType

`func (o *ServerShowResponse) HasPlannedComputeOsType() bool`

HasPlannedComputeOsType returns a boolean if a field has been set.

### SetPlannedComputeOsTypeNil

`func (o *ServerShowResponse) SetPlannedComputeOsTypeNil(b bool)`

 SetPlannedComputeOsTypeNil sets the value for PlannedComputeOsType to be an explicit nil

### UnsetPlannedComputeOsType
`func (o *ServerShowResponse) UnsetPlannedComputeOsType()`

UnsetPlannedComputeOsType ensures that no value is present for PlannedComputeOsType, not even an explicit nil
### GetProductCategory

`func (o *ServerShowResponse) GetProductCategory() ServerProductCategory`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *ServerShowResponse) GetProductCategoryOk() (*ServerProductCategory, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *ServerShowResponse) SetProductCategory(v ServerProductCategory)`

SetProductCategory sets ProductCategory field to given value.

### HasProductCategory

`func (o *ServerShowResponse) HasProductCategory() bool`

HasProductCategory returns a boolean if a field has been set.

### SetProductCategoryNil

`func (o *ServerShowResponse) SetProductCategoryNil(b bool)`

 SetProductCategoryNil sets the value for ProductCategory to be an explicit nil

### UnsetProductCategory
`func (o *ServerShowResponse) UnsetProductCategory()`

UnsetProductCategory ensures that no value is present for ProductCategory, not even an explicit nil
### GetProductOffering

`func (o *ServerShowResponse) GetProductOffering() ServerProductOffering`

GetProductOffering returns the ProductOffering field if non-nil, zero value otherwise.

### GetProductOfferingOk

`func (o *ServerShowResponse) GetProductOfferingOk() (*ServerProductOffering, bool)`

GetProductOfferingOk returns a tuple with the ProductOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOffering

`func (o *ServerShowResponse) SetProductOffering(v ServerProductOffering)`

SetProductOffering sets ProductOffering field to given value.

### HasProductOffering

`func (o *ServerShowResponse) HasProductOffering() bool`

HasProductOffering returns a boolean if a field has been set.

### SetProductOfferingNil

`func (o *ServerShowResponse) SetProductOfferingNil(b bool)`

 SetProductOfferingNil sets the value for ProductOffering to be an explicit nil

### UnsetProductOffering
`func (o *ServerShowResponse) UnsetProductOffering()`

UnsetProductOffering ensures that no value is present for ProductOffering, not even an explicit nil
### GetSecurityGroups

`func (o *ServerShowResponse) GetSecurityGroups() []ServerShowResponseSecurityGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ServerShowResponse) GetSecurityGroupsOk() (*[]ServerShowResponseSecurityGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ServerShowResponse) SetSecurityGroups(v []ServerShowResponseSecurityGroup)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ServerShowResponse) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *ServerShowResponse) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *ServerShowResponse) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetServerGroupId

`func (o *ServerShowResponse) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *ServerShowResponse) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *ServerShowResponse) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.

### HasServerGroupId

`func (o *ServerShowResponse) HasServerGroupId() bool`

HasServerGroupId returns a boolean if a field has been set.

### SetServerGroupIdNil

`func (o *ServerShowResponse) SetServerGroupIdNil(b bool)`

 SetServerGroupIdNil sets the value for ServerGroupId to be an explicit nil

### UnsetServerGroupId
`func (o *ServerShowResponse) UnsetServerGroupId()`

UnsetServerGroupId ensures that no value is present for ServerGroupId, not even an explicit nil
### GetServerType

`func (o *ServerShowResponse) GetServerType() ServerShowResponseServerType`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ServerShowResponse) GetServerTypeOk() (*ServerShowResponseServerType, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ServerShowResponse) SetServerType(v ServerShowResponseServerType)`

SetServerType sets ServerType field to given value.


### GetState

`func (o *ServerShowResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServerShowResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServerShowResponse) SetState(v string)`

SetState sets State field to given value.


### GetVolumes

`func (o *ServerShowResponse) GetVolumes() []ServerShowResponseVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ServerShowResponse) GetVolumesOk() (*[]ServerShowResponseVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ServerShowResponse) SetVolumes(v []ServerShowResponseVolume)`

SetVolumes sets Volumes field to given value.


### GetVpcId

`func (o *ServerShowResponse) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *ServerShowResponse) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *ServerShowResponse) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *ServerShowResponse) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *ServerShowResponse) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *ServerShowResponse) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


