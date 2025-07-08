# DeploymentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**AvailableReplicas** | **int32** | Available Replicas | 
**ClusterId** | **string** | Cluster ID | 
**CreatedAt** | **time.Time** | Created At | 
**Name** | **string** | Deployment Name | 
**NamespaceName** | **string** | Namespace Name | 
**ReadyReplicas** | **int32** | Ready Replicas | 
**SpecReplicas** | **int32** | Spec Replicas | 
**Uid** | **string** | Resource ID | 
**UpdatedReplicas** | **int32** | Updated Replicas | 

## Methods

### NewDeploymentSummary

`func NewDeploymentSummary(age string, availableReplicas int32, clusterId string, createdAt time.Time, name string, namespaceName string, readyReplicas int32, specReplicas int32, uid string, updatedReplicas int32, ) *DeploymentSummary`

NewDeploymentSummary instantiates a new DeploymentSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentSummaryWithDefaults

`func NewDeploymentSummaryWithDefaults() *DeploymentSummary`

NewDeploymentSummaryWithDefaults instantiates a new DeploymentSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *DeploymentSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *DeploymentSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *DeploymentSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetAvailableReplicas

`func (o *DeploymentSummary) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *DeploymentSummary) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *DeploymentSummary) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.


### GetClusterId

`func (o *DeploymentSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *DeploymentSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *DeploymentSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCreatedAt

`func (o *DeploymentSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *DeploymentSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *DeploymentSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *DeploymentSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *DeploymentSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetReadyReplicas

`func (o *DeploymentSummary) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *DeploymentSummary) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *DeploymentSummary) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.


### GetSpecReplicas

`func (o *DeploymentSummary) GetSpecReplicas() int32`

GetSpecReplicas returns the SpecReplicas field if non-nil, zero value otherwise.

### GetSpecReplicasOk

`func (o *DeploymentSummary) GetSpecReplicasOk() (*int32, bool)`

GetSpecReplicasOk returns a tuple with the SpecReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecReplicas

`func (o *DeploymentSummary) SetSpecReplicas(v int32)`

SetSpecReplicas sets SpecReplicas field to given value.


### GetUid

`func (o *DeploymentSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DeploymentSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DeploymentSummary) SetUid(v string)`

SetUid sets Uid field to given value.


### GetUpdatedReplicas

`func (o *DeploymentSummary) GetUpdatedReplicas() int32`

GetUpdatedReplicas returns the UpdatedReplicas field if non-nil, zero value otherwise.

### GetUpdatedReplicasOk

`func (o *DeploymentSummary) GetUpdatedReplicasOk() (*int32, bool)`

GetUpdatedReplicasOk returns a tuple with the UpdatedReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedReplicas

`func (o *DeploymentSummary) SetUpdatedReplicas(v int32)`

SetUpdatedReplicas sets UpdatedReplicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


