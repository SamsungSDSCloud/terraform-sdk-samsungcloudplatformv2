# ClusterCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudLoggingEnabled** | **bool** | Cloud Logging Enabled | 
**KubernetesVersion** | **string** | Cluster Version | 
**Name** | **string** | Cluster Name | 
**PrivateEndpointAccessControlResources** | Pointer to [**[]PrivateEndpointAccessControlResource**](PrivateEndpointAccessControlResource.md) |  | [optional] 
**PublicEndpointAccessControlIp** | Pointer to **NullableString** |  | [optional] 
**SecurityGroupIdList** | **[]string** | Security Group ID List | 
**SubnetId** | **string** | Subnet ID | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**VolumeId** | **string** | Volume ID | 
**VpcId** | **string** | VPC ID | 

## Methods

### NewClusterCreateRequest

`func NewClusterCreateRequest(cloudLoggingEnabled bool, kubernetesVersion string, name string, securityGroupIdList []string, subnetId string, volumeId string, vpcId string, ) *ClusterCreateRequest`

NewClusterCreateRequest instantiates a new ClusterCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateRequestWithDefaults

`func NewClusterCreateRequestWithDefaults() *ClusterCreateRequest`

NewClusterCreateRequestWithDefaults instantiates a new ClusterCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudLoggingEnabled

`func (o *ClusterCreateRequest) GetCloudLoggingEnabled() bool`

GetCloudLoggingEnabled returns the CloudLoggingEnabled field if non-nil, zero value otherwise.

### GetCloudLoggingEnabledOk

`func (o *ClusterCreateRequest) GetCloudLoggingEnabledOk() (*bool, bool)`

GetCloudLoggingEnabledOk returns a tuple with the CloudLoggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudLoggingEnabled

`func (o *ClusterCreateRequest) SetCloudLoggingEnabled(v bool)`

SetCloudLoggingEnabled sets CloudLoggingEnabled field to given value.


### GetKubernetesVersion

`func (o *ClusterCreateRequest) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *ClusterCreateRequest) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *ClusterCreateRequest) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### GetName

`func (o *ClusterCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPrivateEndpointAccessControlResources

`func (o *ClusterCreateRequest) GetPrivateEndpointAccessControlResources() []PrivateEndpointAccessControlResource`

GetPrivateEndpointAccessControlResources returns the PrivateEndpointAccessControlResources field if non-nil, zero value otherwise.

### GetPrivateEndpointAccessControlResourcesOk

`func (o *ClusterCreateRequest) GetPrivateEndpointAccessControlResourcesOk() (*[]PrivateEndpointAccessControlResource, bool)`

GetPrivateEndpointAccessControlResourcesOk returns a tuple with the PrivateEndpointAccessControlResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointAccessControlResources

`func (o *ClusterCreateRequest) SetPrivateEndpointAccessControlResources(v []PrivateEndpointAccessControlResource)`

SetPrivateEndpointAccessControlResources sets PrivateEndpointAccessControlResources field to given value.

### HasPrivateEndpointAccessControlResources

`func (o *ClusterCreateRequest) HasPrivateEndpointAccessControlResources() bool`

HasPrivateEndpointAccessControlResources returns a boolean if a field has been set.

### SetPrivateEndpointAccessControlResourcesNil

`func (o *ClusterCreateRequest) SetPrivateEndpointAccessControlResourcesNil(b bool)`

 SetPrivateEndpointAccessControlResourcesNil sets the value for PrivateEndpointAccessControlResources to be an explicit nil

### UnsetPrivateEndpointAccessControlResources
`func (o *ClusterCreateRequest) UnsetPrivateEndpointAccessControlResources()`

UnsetPrivateEndpointAccessControlResources ensures that no value is present for PrivateEndpointAccessControlResources, not even an explicit nil
### GetPublicEndpointAccessControlIp

`func (o *ClusterCreateRequest) GetPublicEndpointAccessControlIp() string`

GetPublicEndpointAccessControlIp returns the PublicEndpointAccessControlIp field if non-nil, zero value otherwise.

### GetPublicEndpointAccessControlIpOk

`func (o *ClusterCreateRequest) GetPublicEndpointAccessControlIpOk() (*string, bool)`

GetPublicEndpointAccessControlIpOk returns a tuple with the PublicEndpointAccessControlIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEndpointAccessControlIp

`func (o *ClusterCreateRequest) SetPublicEndpointAccessControlIp(v string)`

SetPublicEndpointAccessControlIp sets PublicEndpointAccessControlIp field to given value.

### HasPublicEndpointAccessControlIp

`func (o *ClusterCreateRequest) HasPublicEndpointAccessControlIp() bool`

HasPublicEndpointAccessControlIp returns a boolean if a field has been set.

### SetPublicEndpointAccessControlIpNil

`func (o *ClusterCreateRequest) SetPublicEndpointAccessControlIpNil(b bool)`

 SetPublicEndpointAccessControlIpNil sets the value for PublicEndpointAccessControlIp to be an explicit nil

### UnsetPublicEndpointAccessControlIp
`func (o *ClusterCreateRequest) UnsetPublicEndpointAccessControlIp()`

UnsetPublicEndpointAccessControlIp ensures that no value is present for PublicEndpointAccessControlIp, not even an explicit nil
### GetSecurityGroupIdList

`func (o *ClusterCreateRequest) GetSecurityGroupIdList() []string`

GetSecurityGroupIdList returns the SecurityGroupIdList field if non-nil, zero value otherwise.

### GetSecurityGroupIdListOk

`func (o *ClusterCreateRequest) GetSecurityGroupIdListOk() (*[]string, bool)`

GetSecurityGroupIdListOk returns a tuple with the SecurityGroupIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIdList

`func (o *ClusterCreateRequest) SetSecurityGroupIdList(v []string)`

SetSecurityGroupIdList sets SecurityGroupIdList field to given value.


### GetSubnetId

`func (o *ClusterCreateRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *ClusterCreateRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *ClusterCreateRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *ClusterCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClusterCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClusterCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClusterCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ClusterCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ClusterCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetVolumeId

`func (o *ClusterCreateRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *ClusterCreateRequest) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *ClusterCreateRequest) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.


### GetVpcId

`func (o *ClusterCreateRequest) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *ClusterCreateRequest) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *ClusterCreateRequest) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


