# NodepoolSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AutoRecoveryEnabled** | **bool** | Is Auto Recovery | 
**AutoScaleEnabled** | **bool** | Is Auto Scale | 
**CurrentNodeCount** | **int32** | Current Node Count | 
**DesiredNodeCount** | **int32** | Desired Node Count | 
**Id** | **string** | Nodepool ID | 
**Image** | [**Image**](Image.md) | Image | 
**KubernetesVersion** | **string** | Kubernetes Version | 
**Name** | **string** | Nodepool Name | 
**ServerType** | [**ServerType**](ServerType.md) | Server Type | 
**Status** | **string** | Node Pool Status | 
**VolumeType** | [**VolumeTypeSummary**](VolumeTypeSummary.md) | Volume Type | 

## Methods

### NewNodepoolSummary

`func NewNodepoolSummary(accountId string, autoRecoveryEnabled bool, autoScaleEnabled bool, currentNodeCount int32, desiredNodeCount int32, id string, image Image, kubernetesVersion string, name string, serverType ServerType, status string, volumeType VolumeTypeSummary, ) *NodepoolSummary`

NewNodepoolSummary instantiates a new NodepoolSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodepoolSummaryWithDefaults

`func NewNodepoolSummaryWithDefaults() *NodepoolSummary`

NewNodepoolSummaryWithDefaults instantiates a new NodepoolSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NodepoolSummary) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NodepoolSummary) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NodepoolSummary) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAutoRecoveryEnabled

`func (o *NodepoolSummary) GetAutoRecoveryEnabled() bool`

GetAutoRecoveryEnabled returns the AutoRecoveryEnabled field if non-nil, zero value otherwise.

### GetAutoRecoveryEnabledOk

`func (o *NodepoolSummary) GetAutoRecoveryEnabledOk() (*bool, bool)`

GetAutoRecoveryEnabledOk returns a tuple with the AutoRecoveryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoveryEnabled

`func (o *NodepoolSummary) SetAutoRecoveryEnabled(v bool)`

SetAutoRecoveryEnabled sets AutoRecoveryEnabled field to given value.


### GetAutoScaleEnabled

`func (o *NodepoolSummary) GetAutoScaleEnabled() bool`

GetAutoScaleEnabled returns the AutoScaleEnabled field if non-nil, zero value otherwise.

### GetAutoScaleEnabledOk

`func (o *NodepoolSummary) GetAutoScaleEnabledOk() (*bool, bool)`

GetAutoScaleEnabledOk returns a tuple with the AutoScaleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaleEnabled

`func (o *NodepoolSummary) SetAutoScaleEnabled(v bool)`

SetAutoScaleEnabled sets AutoScaleEnabled field to given value.


### GetCurrentNodeCount

`func (o *NodepoolSummary) GetCurrentNodeCount() int32`

GetCurrentNodeCount returns the CurrentNodeCount field if non-nil, zero value otherwise.

### GetCurrentNodeCountOk

`func (o *NodepoolSummary) GetCurrentNodeCountOk() (*int32, bool)`

GetCurrentNodeCountOk returns a tuple with the CurrentNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNodeCount

`func (o *NodepoolSummary) SetCurrentNodeCount(v int32)`

SetCurrentNodeCount sets CurrentNodeCount field to given value.


### GetDesiredNodeCount

`func (o *NodepoolSummary) GetDesiredNodeCount() int32`

GetDesiredNodeCount returns the DesiredNodeCount field if non-nil, zero value otherwise.

### GetDesiredNodeCountOk

`func (o *NodepoolSummary) GetDesiredNodeCountOk() (*int32, bool)`

GetDesiredNodeCountOk returns a tuple with the DesiredNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNodeCount

`func (o *NodepoolSummary) SetDesiredNodeCount(v int32)`

SetDesiredNodeCount sets DesiredNodeCount field to given value.


### GetId

`func (o *NodepoolSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodepoolSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodepoolSummary) SetId(v string)`

SetId sets Id field to given value.


### GetImage

`func (o *NodepoolSummary) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *NodepoolSummary) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *NodepoolSummary) SetImage(v Image)`

SetImage sets Image field to given value.


### GetKubernetesVersion

`func (o *NodepoolSummary) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *NodepoolSummary) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *NodepoolSummary) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### GetName

`func (o *NodepoolSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodepoolSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodepoolSummary) SetName(v string)`

SetName sets Name field to given value.


### GetServerType

`func (o *NodepoolSummary) GetServerType() ServerType`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *NodepoolSummary) GetServerTypeOk() (*ServerType, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *NodepoolSummary) SetServerType(v ServerType)`

SetServerType sets ServerType field to given value.


### GetStatus

`func (o *NodepoolSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodepoolSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodepoolSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVolumeType

`func (o *NodepoolSummary) GetVolumeType() VolumeTypeSummary`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *NodepoolSummary) GetVolumeTypeOk() (*VolumeTypeSummary, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *NodepoolSummary) SetVolumeType(v VolumeTypeSummary)`

SetVolumeType sets VolumeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


