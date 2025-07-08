# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CloudLoggingEnabled** | **bool** | Cloud Logging Enabled | 
**ClusterNamespace** | **string** | Cluster Namespace | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | **string** | ID | 
**KubernetesVersion** | **string** | Cluster Version | 
**ManagedSecurityGroup** | [**NullableExternalResource**](ExternalResource.md) |  | 
**MaxNodeCount** | **NullableInt32** |  | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Cluster Name | 
**NodeCount** | **NullableInt32** |  | 
**PrivateEndpointAccessControlResources** | [**[]PrivateEndpointAccessControlResource**](PrivateEndpointAccessControlResource.md) |  | 
**PrivateEndpointUrl** | **string** | Private Endpoint URL | 
**PublicEndpointAccessControlIp** | **NullableString** |  | 
**PublicEndpointUrl** | **NullableString** |  | 
**SecurityGroupList** | [**[]ExternalResource**](ExternalResource.md) | Connected Security Group List | 
**Status** | **string** | Cluster Status | 
**Subnet** | [**NullableExternalResource**](ExternalResource.md) |  | 
**Volume** | [**NullableExternalResource**](ExternalResource.md) |  | 
**Vpc** | [**NullableExternalResource**](ExternalResource.md) |  | 

## Methods

### NewCluster

`func NewCluster(accountId string, cloudLoggingEnabled bool, clusterNamespace string, createdAt time.Time, createdBy string, id string, kubernetesVersion string, managedSecurityGroup NullableExternalResource, maxNodeCount NullableInt32, modifiedAt time.Time, modifiedBy string, name string, nodeCount NullableInt32, privateEndpointAccessControlResources []PrivateEndpointAccessControlResource, privateEndpointUrl string, publicEndpointAccessControlIp NullableString, publicEndpointUrl NullableString, securityGroupList []ExternalResource, status string, subnet NullableExternalResource, volume NullableExternalResource, vpc NullableExternalResource, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Cluster) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Cluster) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Cluster) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCloudLoggingEnabled

`func (o *Cluster) GetCloudLoggingEnabled() bool`

GetCloudLoggingEnabled returns the CloudLoggingEnabled field if non-nil, zero value otherwise.

### GetCloudLoggingEnabledOk

`func (o *Cluster) GetCloudLoggingEnabledOk() (*bool, bool)`

GetCloudLoggingEnabledOk returns a tuple with the CloudLoggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudLoggingEnabled

`func (o *Cluster) SetCloudLoggingEnabled(v bool)`

SetCloudLoggingEnabled sets CloudLoggingEnabled field to given value.


### GetClusterNamespace

`func (o *Cluster) GetClusterNamespace() string`

GetClusterNamespace returns the ClusterNamespace field if non-nil, zero value otherwise.

### GetClusterNamespaceOk

`func (o *Cluster) GetClusterNamespaceOk() (*string, bool)`

GetClusterNamespaceOk returns a tuple with the ClusterNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNamespace

`func (o *Cluster) SetClusterNamespace(v string)`

SetClusterNamespace sets ClusterNamespace field to given value.


### GetCreatedAt

`func (o *Cluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Cluster) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Cluster) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Cluster) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.


### GetKubernetesVersion

`func (o *Cluster) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *Cluster) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *Cluster) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### GetManagedSecurityGroup

`func (o *Cluster) GetManagedSecurityGroup() ExternalResource`

GetManagedSecurityGroup returns the ManagedSecurityGroup field if non-nil, zero value otherwise.

### GetManagedSecurityGroupOk

`func (o *Cluster) GetManagedSecurityGroupOk() (*ExternalResource, bool)`

GetManagedSecurityGroupOk returns a tuple with the ManagedSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedSecurityGroup

`func (o *Cluster) SetManagedSecurityGroup(v ExternalResource)`

SetManagedSecurityGroup sets ManagedSecurityGroup field to given value.


### SetManagedSecurityGroupNil

`func (o *Cluster) SetManagedSecurityGroupNil(b bool)`

 SetManagedSecurityGroupNil sets the value for ManagedSecurityGroup to be an explicit nil

### UnsetManagedSecurityGroup
`func (o *Cluster) UnsetManagedSecurityGroup()`

UnsetManagedSecurityGroup ensures that no value is present for ManagedSecurityGroup, not even an explicit nil
### GetMaxNodeCount

`func (o *Cluster) GetMaxNodeCount() int32`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *Cluster) GetMaxNodeCountOk() (*int32, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *Cluster) SetMaxNodeCount(v int32)`

SetMaxNodeCount sets MaxNodeCount field to given value.


### SetMaxNodeCountNil

`func (o *Cluster) SetMaxNodeCountNil(b bool)`

 SetMaxNodeCountNil sets the value for MaxNodeCount to be an explicit nil

### UnsetMaxNodeCount
`func (o *Cluster) UnsetMaxNodeCount()`

UnsetMaxNodeCount ensures that no value is present for MaxNodeCount, not even an explicit nil
### GetModifiedAt

`func (o *Cluster) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Cluster) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Cluster) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *Cluster) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Cluster) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Cluster) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetNodeCount

`func (o *Cluster) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *Cluster) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *Cluster) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### SetNodeCountNil

`func (o *Cluster) SetNodeCountNil(b bool)`

 SetNodeCountNil sets the value for NodeCount to be an explicit nil

### UnsetNodeCount
`func (o *Cluster) UnsetNodeCount()`

UnsetNodeCount ensures that no value is present for NodeCount, not even an explicit nil
### GetPrivateEndpointAccessControlResources

`func (o *Cluster) GetPrivateEndpointAccessControlResources() []PrivateEndpointAccessControlResource`

GetPrivateEndpointAccessControlResources returns the PrivateEndpointAccessControlResources field if non-nil, zero value otherwise.

### GetPrivateEndpointAccessControlResourcesOk

`func (o *Cluster) GetPrivateEndpointAccessControlResourcesOk() (*[]PrivateEndpointAccessControlResource, bool)`

GetPrivateEndpointAccessControlResourcesOk returns a tuple with the PrivateEndpointAccessControlResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointAccessControlResources

`func (o *Cluster) SetPrivateEndpointAccessControlResources(v []PrivateEndpointAccessControlResource)`

SetPrivateEndpointAccessControlResources sets PrivateEndpointAccessControlResources field to given value.


### SetPrivateEndpointAccessControlResourcesNil

`func (o *Cluster) SetPrivateEndpointAccessControlResourcesNil(b bool)`

 SetPrivateEndpointAccessControlResourcesNil sets the value for PrivateEndpointAccessControlResources to be an explicit nil

### UnsetPrivateEndpointAccessControlResources
`func (o *Cluster) UnsetPrivateEndpointAccessControlResources()`

UnsetPrivateEndpointAccessControlResources ensures that no value is present for PrivateEndpointAccessControlResources, not even an explicit nil
### GetPrivateEndpointUrl

`func (o *Cluster) GetPrivateEndpointUrl() string`

GetPrivateEndpointUrl returns the PrivateEndpointUrl field if non-nil, zero value otherwise.

### GetPrivateEndpointUrlOk

`func (o *Cluster) GetPrivateEndpointUrlOk() (*string, bool)`

GetPrivateEndpointUrlOk returns a tuple with the PrivateEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointUrl

`func (o *Cluster) SetPrivateEndpointUrl(v string)`

SetPrivateEndpointUrl sets PrivateEndpointUrl field to given value.


### GetPublicEndpointAccessControlIp

`func (o *Cluster) GetPublicEndpointAccessControlIp() string`

GetPublicEndpointAccessControlIp returns the PublicEndpointAccessControlIp field if non-nil, zero value otherwise.

### GetPublicEndpointAccessControlIpOk

`func (o *Cluster) GetPublicEndpointAccessControlIpOk() (*string, bool)`

GetPublicEndpointAccessControlIpOk returns a tuple with the PublicEndpointAccessControlIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEndpointAccessControlIp

`func (o *Cluster) SetPublicEndpointAccessControlIp(v string)`

SetPublicEndpointAccessControlIp sets PublicEndpointAccessControlIp field to given value.


### SetPublicEndpointAccessControlIpNil

`func (o *Cluster) SetPublicEndpointAccessControlIpNil(b bool)`

 SetPublicEndpointAccessControlIpNil sets the value for PublicEndpointAccessControlIp to be an explicit nil

### UnsetPublicEndpointAccessControlIp
`func (o *Cluster) UnsetPublicEndpointAccessControlIp()`

UnsetPublicEndpointAccessControlIp ensures that no value is present for PublicEndpointAccessControlIp, not even an explicit nil
### GetPublicEndpointUrl

`func (o *Cluster) GetPublicEndpointUrl() string`

GetPublicEndpointUrl returns the PublicEndpointUrl field if non-nil, zero value otherwise.

### GetPublicEndpointUrlOk

`func (o *Cluster) GetPublicEndpointUrlOk() (*string, bool)`

GetPublicEndpointUrlOk returns a tuple with the PublicEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEndpointUrl

`func (o *Cluster) SetPublicEndpointUrl(v string)`

SetPublicEndpointUrl sets PublicEndpointUrl field to given value.


### SetPublicEndpointUrlNil

`func (o *Cluster) SetPublicEndpointUrlNil(b bool)`

 SetPublicEndpointUrlNil sets the value for PublicEndpointUrl to be an explicit nil

### UnsetPublicEndpointUrl
`func (o *Cluster) UnsetPublicEndpointUrl()`

UnsetPublicEndpointUrl ensures that no value is present for PublicEndpointUrl, not even an explicit nil
### GetSecurityGroupList

`func (o *Cluster) GetSecurityGroupList() []ExternalResource`

GetSecurityGroupList returns the SecurityGroupList field if non-nil, zero value otherwise.

### GetSecurityGroupListOk

`func (o *Cluster) GetSecurityGroupListOk() (*[]ExternalResource, bool)`

GetSecurityGroupListOk returns a tuple with the SecurityGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupList

`func (o *Cluster) SetSecurityGroupList(v []ExternalResource)`

SetSecurityGroupList sets SecurityGroupList field to given value.


### GetStatus

`func (o *Cluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Cluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Cluster) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubnet

`func (o *Cluster) GetSubnet() ExternalResource`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *Cluster) GetSubnetOk() (*ExternalResource, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *Cluster) SetSubnet(v ExternalResource)`

SetSubnet sets Subnet field to given value.


### SetSubnetNil

`func (o *Cluster) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *Cluster) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetVolume

`func (o *Cluster) GetVolume() ExternalResource`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Cluster) GetVolumeOk() (*ExternalResource, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Cluster) SetVolume(v ExternalResource)`

SetVolume sets Volume field to given value.


### SetVolumeNil

`func (o *Cluster) SetVolumeNil(b bool)`

 SetVolumeNil sets the value for Volume to be an explicit nil

### UnsetVolume
`func (o *Cluster) UnsetVolume()`

UnsetVolume ensures that no value is present for Volume, not even an explicit nil
### GetVpc

`func (o *Cluster) GetVpc() ExternalResource`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *Cluster) GetVpcOk() (*ExternalResource, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *Cluster) SetVpc(v ExternalResource)`

SetVpc sets Vpc field to given value.


### SetVpcNil

`func (o *Cluster) SetVpcNil(b bool)`

 SetVpcNil sets the value for Vpc to be an explicit nil

### UnsetVpc
`func (o *Cluster) UnsetVpc()`

UnsetVpc ensures that no value is present for Vpc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


