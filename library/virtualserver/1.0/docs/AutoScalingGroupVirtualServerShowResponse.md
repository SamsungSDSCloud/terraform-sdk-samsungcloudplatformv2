# AutoScalingGroupVirtualServerShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**Addresses** | Pointer to [**[]ServerShowResponseAddress**](ServerShowResponseAddress.md) |  | [optional] 
**AutoScalingGroupId** | **string** | Auto-Scaling Group ID | 
**CreatedAt** | **time.Time** | Created at | 
**CreatedBy** | **string** | Created by | 
**DiskConfig** | **string** | Disk config | 
**Id** | **string** | Server ID | 
**ImageId** | Pointer to **NullableString** |  | [optional] 
**IsMarketplace** | Pointer to **NullableBool** |  | [optional] 
**KeypairName** | Pointer to **NullableString** |  | [optional] 
**LaunchConfigurationId** | Pointer to **NullableString** |  | [optional] 
**LbLinkedState** | Pointer to **NullableString** |  | [optional] 
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

### NewAutoScalingGroupVirtualServerShowResponse

`func NewAutoScalingGroupVirtualServerShowResponse(accountId string, autoScalingGroupId string, createdAt time.Time, createdBy string, diskConfig string, id string, locked bool, metadata map[string]interface{}, modifiedAt time.Time, name string, serverType ServerShowResponseServerType, state string, volumes []ServerShowResponseVolume, ) *AutoScalingGroupVirtualServerShowResponse`

NewAutoScalingGroupVirtualServerShowResponse instantiates a new AutoScalingGroupVirtualServerShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupVirtualServerShowResponseWithDefaults

`func NewAutoScalingGroupVirtualServerShowResponseWithDefaults() *AutoScalingGroupVirtualServerShowResponse`

NewAutoScalingGroupVirtualServerShowResponseWithDefaults instantiates a new AutoScalingGroupVirtualServerShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AutoScalingGroupVirtualServerShowResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AutoScalingGroupVirtualServerShowResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAddresses

`func (o *AutoScalingGroupVirtualServerShowResponse) GetAddresses() []ServerShowResponseAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetAddressesOk() (*[]ServerShowResponseAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *AutoScalingGroupVirtualServerShowResponse) SetAddresses(v []ServerShowResponseAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *AutoScalingGroupVirtualServerShowResponse) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *AutoScalingGroupVirtualServerShowResponse) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *AutoScalingGroupVirtualServerShowResponse) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetAutoScalingGroupId

`func (o *AutoScalingGroupVirtualServerShowResponse) GetAutoScalingGroupId() string`

GetAutoScalingGroupId returns the AutoScalingGroupId field if non-nil, zero value otherwise.

### GetAutoScalingGroupIdOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetAutoScalingGroupIdOk() (*string, bool)`

GetAutoScalingGroupIdOk returns a tuple with the AutoScalingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroupId

`func (o *AutoScalingGroupVirtualServerShowResponse) SetAutoScalingGroupId(v string)`

SetAutoScalingGroupId sets AutoScalingGroupId field to given value.


### GetCreatedAt

`func (o *AutoScalingGroupVirtualServerShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutoScalingGroupVirtualServerShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AutoScalingGroupVirtualServerShowResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AutoScalingGroupVirtualServerShowResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDiskConfig

`func (o *AutoScalingGroupVirtualServerShowResponse) GetDiskConfig() string`

GetDiskConfig returns the DiskConfig field if non-nil, zero value otherwise.

### GetDiskConfigOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetDiskConfigOk() (*string, bool)`

GetDiskConfigOk returns a tuple with the DiskConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskConfig

`func (o *AutoScalingGroupVirtualServerShowResponse) SetDiskConfig(v string)`

SetDiskConfig sets DiskConfig field to given value.


### GetId

`func (o *AutoScalingGroupVirtualServerShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoScalingGroupVirtualServerShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetImageId

`func (o *AutoScalingGroupVirtualServerShowResponse) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AutoScalingGroupVirtualServerShowResponse) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *AutoScalingGroupVirtualServerShowResponse) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *AutoScalingGroupVirtualServerShowResponse) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *AutoScalingGroupVirtualServerShowResponse) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetIsMarketplace

`func (o *AutoScalingGroupVirtualServerShowResponse) GetIsMarketplace() bool`

GetIsMarketplace returns the IsMarketplace field if non-nil, zero value otherwise.

### GetIsMarketplaceOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetIsMarketplaceOk() (*bool, bool)`

GetIsMarketplaceOk returns a tuple with the IsMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketplace

`func (o *AutoScalingGroupVirtualServerShowResponse) SetIsMarketplace(v bool)`

SetIsMarketplace sets IsMarketplace field to given value.

### HasIsMarketplace

`func (o *AutoScalingGroupVirtualServerShowResponse) HasIsMarketplace() bool`

HasIsMarketplace returns a boolean if a field has been set.

### SetIsMarketplaceNil

`func (o *AutoScalingGroupVirtualServerShowResponse) SetIsMarketplaceNil(b bool)`

 SetIsMarketplaceNil sets the value for IsMarketplace to be an explicit nil

### UnsetIsMarketplace
`func (o *AutoScalingGroupVirtualServerShowResponse) UnsetIsMarketplace()`

UnsetIsMarketplace ensures that no value is present for IsMarketplace, not even an explicit nil
### GetKeypairName

`func (o *AutoScalingGroupVirtualServerShowResponse) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *AutoScalingGroupVirtualServerShowResponse) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.

### HasKeypairName

`func (o *AutoScalingGroupVirtualServerShowResponse) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### SetKeypairNameNil

`func (o *AutoScalingGroupVirtualServerShowResponse) SetKeypairNameNil(b bool)`

 SetKeypairNameNil sets the value for KeypairName to be an explicit nil

### UnsetKeypairName
`func (o *AutoScalingGroupVirtualServerShowResponse) UnsetKeypairName()`

UnsetKeypairName ensures that no value is present for KeypairName, not even an explicit nil
### GetLaunchConfigurationId

`func (o *AutoScalingGroupVirtualServerShowResponse) GetLaunchConfigurationId() string`

GetLaunchConfigurationId returns the LaunchConfigurationId field if non-nil, zero value otherwise.

### GetLaunchConfigurationIdOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetLaunchConfigurationIdOk() (*string, bool)`

GetLaunchConfigurationIdOk returns a tuple with the LaunchConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchConfigurationId

`func (o *AutoScalingGroupVirtualServerShowResponse) SetLaunchConfigurationId(v string)`

SetLaunchConfigurationId sets LaunchConfigurationId field to given value.

### HasLaunchConfigurationId

`func (o *AutoScalingGroupVirtualServerShowResponse) HasLaunchConfigurationId() bool`

HasLaunchConfigurationId returns a boolean if a field has been set.

### SetLaunchConfigurationIdNil

`func (o *AutoScalingGroupVirtualServerShowResponse) SetLaunchConfigurationIdNil(b bool)`

 SetLaunchConfigurationIdNil sets the value for LaunchConfigurationId to be an explicit nil

### UnsetLaunchConfigurationId
`func (o *AutoScalingGroupVirtualServerShowResponse) UnsetLaunchConfigurationId()`

UnsetLaunchConfigurationId ensures that no value is present for LaunchConfigurationId, not even an explicit nil
### GetLbLinkedState

`func (o *AutoScalingGroupVirtualServerShowResponse) GetLbLinkedState() string`

GetLbLinkedState returns the LbLinkedState field if non-nil, zero value otherwise.

### GetLbLinkedStateOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetLbLinkedStateOk() (*string, bool)`

GetLbLinkedStateOk returns a tuple with the LbLinkedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbLinkedState

`func (o *AutoScalingGroupVirtualServerShowResponse) SetLbLinkedState(v string)`

SetLbLinkedState sets LbLinkedState field to given value.

### HasLbLinkedState

`func (o *AutoScalingGroupVirtualServerShowResponse) HasLbLinkedState() bool`

HasLbLinkedState returns a boolean if a field has been set.

### SetLbLinkedStateNil

`func (o *AutoScalingGroupVirtualServerShowResponse) SetLbLinkedStateNil(b bool)`

 SetLbLinkedStateNil sets the value for LbLinkedState to be an explicit nil

### UnsetLbLinkedState
`func (o *AutoScalingGroupVirtualServerShowResponse) UnsetLbLinkedState()`

UnsetLbLinkedState ensures that no value is present for LbLinkedState, not even an explicit nil
### GetLocked

`func (o *AutoScalingGroupVirtualServerShowResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *AutoScalingGroupVirtualServerShowResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetMetadata

`func (o *AutoScalingGroupVirtualServerShowResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AutoScalingGroupVirtualServerShowResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetModifiedAt

`func (o *AutoScalingGroupVirtualServerShowResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AutoScalingGroupVirtualServerShowResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetName

`func (o *AutoScalingGroupVirtualServerShowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoScalingGroupVirtualServerShowResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPlannedComputeOsType

`func (o *AutoScalingGroupVirtualServerShowResponse) GetPlannedComputeOsType() PlannedComputeOsType`

GetPlannedComputeOsType returns the PlannedComputeOsType field if non-nil, zero value otherwise.

### GetPlannedComputeOsTypeOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetPlannedComputeOsTypeOk() (*PlannedComputeOsType, bool)`

GetPlannedComputeOsTypeOk returns a tuple with the PlannedComputeOsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedComputeOsType

`func (o *AutoScalingGroupVirtualServerShowResponse) SetPlannedComputeOsType(v PlannedComputeOsType)`

SetPlannedComputeOsType sets PlannedComputeOsType field to given value.

### HasPlannedComputeOsType

`func (o *AutoScalingGroupVirtualServerShowResponse) HasPlannedComputeOsType() bool`

HasPlannedComputeOsType returns a boolean if a field has been set.

### SetPlannedComputeOsTypeNil

`func (o *AutoScalingGroupVirtualServerShowResponse) SetPlannedComputeOsTypeNil(b bool)`

 SetPlannedComputeOsTypeNil sets the value for PlannedComputeOsType to be an explicit nil

### UnsetPlannedComputeOsType
`func (o *AutoScalingGroupVirtualServerShowResponse) UnsetPlannedComputeOsType()`

UnsetPlannedComputeOsType ensures that no value is present for PlannedComputeOsType, not even an explicit nil
### GetProductCategory

`func (o *AutoScalingGroupVirtualServerShowResponse) GetProductCategory() ServerProductCategory`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetProductCategoryOk() (*ServerProductCategory, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *AutoScalingGroupVirtualServerShowResponse) SetProductCategory(v ServerProductCategory)`

SetProductCategory sets ProductCategory field to given value.

### HasProductCategory

`func (o *AutoScalingGroupVirtualServerShowResponse) HasProductCategory() bool`

HasProductCategory returns a boolean if a field has been set.

### SetProductCategoryNil

`func (o *AutoScalingGroupVirtualServerShowResponse) SetProductCategoryNil(b bool)`

 SetProductCategoryNil sets the value for ProductCategory to be an explicit nil

### UnsetProductCategory
`func (o *AutoScalingGroupVirtualServerShowResponse) UnsetProductCategory()`

UnsetProductCategory ensures that no value is present for ProductCategory, not even an explicit nil
### GetProductOffering

`func (o *AutoScalingGroupVirtualServerShowResponse) GetProductOffering() ServerProductOffering`

GetProductOffering returns the ProductOffering field if non-nil, zero value otherwise.

### GetProductOfferingOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetProductOfferingOk() (*ServerProductOffering, bool)`

GetProductOfferingOk returns a tuple with the ProductOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOffering

`func (o *AutoScalingGroupVirtualServerShowResponse) SetProductOffering(v ServerProductOffering)`

SetProductOffering sets ProductOffering field to given value.

### HasProductOffering

`func (o *AutoScalingGroupVirtualServerShowResponse) HasProductOffering() bool`

HasProductOffering returns a boolean if a field has been set.

### SetProductOfferingNil

`func (o *AutoScalingGroupVirtualServerShowResponse) SetProductOfferingNil(b bool)`

 SetProductOfferingNil sets the value for ProductOffering to be an explicit nil

### UnsetProductOffering
`func (o *AutoScalingGroupVirtualServerShowResponse) UnsetProductOffering()`

UnsetProductOffering ensures that no value is present for ProductOffering, not even an explicit nil
### GetSecurityGroups

`func (o *AutoScalingGroupVirtualServerShowResponse) GetSecurityGroups() []ServerShowResponseSecurityGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetSecurityGroupsOk() (*[]ServerShowResponseSecurityGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *AutoScalingGroupVirtualServerShowResponse) SetSecurityGroups(v []ServerShowResponseSecurityGroup)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *AutoScalingGroupVirtualServerShowResponse) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *AutoScalingGroupVirtualServerShowResponse) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *AutoScalingGroupVirtualServerShowResponse) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetServerGroupId

`func (o *AutoScalingGroupVirtualServerShowResponse) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *AutoScalingGroupVirtualServerShowResponse) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.

### HasServerGroupId

`func (o *AutoScalingGroupVirtualServerShowResponse) HasServerGroupId() bool`

HasServerGroupId returns a boolean if a field has been set.

### SetServerGroupIdNil

`func (o *AutoScalingGroupVirtualServerShowResponse) SetServerGroupIdNil(b bool)`

 SetServerGroupIdNil sets the value for ServerGroupId to be an explicit nil

### UnsetServerGroupId
`func (o *AutoScalingGroupVirtualServerShowResponse) UnsetServerGroupId()`

UnsetServerGroupId ensures that no value is present for ServerGroupId, not even an explicit nil
### GetServerType

`func (o *AutoScalingGroupVirtualServerShowResponse) GetServerType() ServerShowResponseServerType`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetServerTypeOk() (*ServerShowResponseServerType, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *AutoScalingGroupVirtualServerShowResponse) SetServerType(v ServerShowResponseServerType)`

SetServerType sets ServerType field to given value.


### GetState

`func (o *AutoScalingGroupVirtualServerShowResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingGroupVirtualServerShowResponse) SetState(v string)`

SetState sets State field to given value.


### GetVolumes

`func (o *AutoScalingGroupVirtualServerShowResponse) GetVolumes() []ServerShowResponseVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetVolumesOk() (*[]ServerShowResponseVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AutoScalingGroupVirtualServerShowResponse) SetVolumes(v []ServerShowResponseVolume)`

SetVolumes sets Volumes field to given value.


### GetVpcId

`func (o *AutoScalingGroupVirtualServerShowResponse) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AutoScalingGroupVirtualServerShowResponse) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AutoScalingGroupVirtualServerShowResponse) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AutoScalingGroupVirtualServerShowResponse) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *AutoScalingGroupVirtualServerShowResponse) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *AutoScalingGroupVirtualServerShowResponse) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


