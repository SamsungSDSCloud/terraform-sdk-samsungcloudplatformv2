# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** |  | 
**AvailableReplicas** | **int32** | Available Replicas | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**CreatedAt** | **time.Time** | Created At | 
**Labels** | **[]string** |  | 
**MaxSurge** | **NullableString** |  | 
**MaxUnavailable** | **NullableString** |  | 
**MinReadySeconds** | **NullableString** |  | 
**Name** | **string** | Deployment Name | 
**NamespaceName** | **string** | Namespace Name | 
**ReadyReplicas** | **int32** | Ready Replicas | 
**Selector** | **[]string** |  | 
**SelectorDetails** | **[]string** |  | 
**SpecReplicas** | **int32** | Spec Replicas | 
**StatusReplicas** | **int32** | Status Replicas | 
**StrategyType** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 
**UnavailableReplicas** | **int32** | Unavailable Replicas | 
**UpdatedReplicas** | **int32** | Updated Replicas | 
**Yaml** | **string** | YAML | 

## Methods

### NewDeployment

`func NewDeployment(age string, annotations []string, availableReplicas int32, clusterId string, clusterName string, createdAt time.Time, labels []string, maxSurge NullableString, maxUnavailable NullableString, minReadySeconds NullableString, name string, namespaceName string, readyReplicas int32, selector []string, selectorDetails []string, specReplicas int32, statusReplicas int32, strategyType NullableString, uid string, unavailableReplicas int32, updatedReplicas int32, yaml string, ) *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *Deployment) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Deployment) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Deployment) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *Deployment) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Deployment) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Deployment) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### SetAnnotationsNil

`func (o *Deployment) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *Deployment) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetAvailableReplicas

`func (o *Deployment) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *Deployment) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *Deployment) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.


### GetClusterId

`func (o *Deployment) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Deployment) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Deployment) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *Deployment) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Deployment) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Deployment) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCreatedAt

`func (o *Deployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Deployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Deployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLabels

`func (o *Deployment) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Deployment) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Deployment) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### SetLabelsNil

`func (o *Deployment) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *Deployment) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMaxSurge

`func (o *Deployment) GetMaxSurge() string`

GetMaxSurge returns the MaxSurge field if non-nil, zero value otherwise.

### GetMaxSurgeOk

`func (o *Deployment) GetMaxSurgeOk() (*string, bool)`

GetMaxSurgeOk returns a tuple with the MaxSurge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSurge

`func (o *Deployment) SetMaxSurge(v string)`

SetMaxSurge sets MaxSurge field to given value.


### SetMaxSurgeNil

`func (o *Deployment) SetMaxSurgeNil(b bool)`

 SetMaxSurgeNil sets the value for MaxSurge to be an explicit nil

### UnsetMaxSurge
`func (o *Deployment) UnsetMaxSurge()`

UnsetMaxSurge ensures that no value is present for MaxSurge, not even an explicit nil
### GetMaxUnavailable

`func (o *Deployment) GetMaxUnavailable() string`

GetMaxUnavailable returns the MaxUnavailable field if non-nil, zero value otherwise.

### GetMaxUnavailableOk

`func (o *Deployment) GetMaxUnavailableOk() (*string, bool)`

GetMaxUnavailableOk returns a tuple with the MaxUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnavailable

`func (o *Deployment) SetMaxUnavailable(v string)`

SetMaxUnavailable sets MaxUnavailable field to given value.


### SetMaxUnavailableNil

`func (o *Deployment) SetMaxUnavailableNil(b bool)`

 SetMaxUnavailableNil sets the value for MaxUnavailable to be an explicit nil

### UnsetMaxUnavailable
`func (o *Deployment) UnsetMaxUnavailable()`

UnsetMaxUnavailable ensures that no value is present for MaxUnavailable, not even an explicit nil
### GetMinReadySeconds

`func (o *Deployment) GetMinReadySeconds() string`

GetMinReadySeconds returns the MinReadySeconds field if non-nil, zero value otherwise.

### GetMinReadySecondsOk

`func (o *Deployment) GetMinReadySecondsOk() (*string, bool)`

GetMinReadySecondsOk returns a tuple with the MinReadySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadySeconds

`func (o *Deployment) SetMinReadySeconds(v string)`

SetMinReadySeconds sets MinReadySeconds field to given value.


### SetMinReadySecondsNil

`func (o *Deployment) SetMinReadySecondsNil(b bool)`

 SetMinReadySecondsNil sets the value for MinReadySeconds to be an explicit nil

### UnsetMinReadySeconds
`func (o *Deployment) UnsetMinReadySeconds()`

UnsetMinReadySeconds ensures that no value is present for MinReadySeconds, not even an explicit nil
### GetName

`func (o *Deployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Deployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Deployment) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *Deployment) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *Deployment) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *Deployment) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetReadyReplicas

`func (o *Deployment) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *Deployment) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *Deployment) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.


### GetSelector

`func (o *Deployment) GetSelector() []string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *Deployment) GetSelectorOk() (*[]string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *Deployment) SetSelector(v []string)`

SetSelector sets Selector field to given value.


### SetSelectorNil

`func (o *Deployment) SetSelectorNil(b bool)`

 SetSelectorNil sets the value for Selector to be an explicit nil

### UnsetSelector
`func (o *Deployment) UnsetSelector()`

UnsetSelector ensures that no value is present for Selector, not even an explicit nil
### GetSelectorDetails

`func (o *Deployment) GetSelectorDetails() []string`

GetSelectorDetails returns the SelectorDetails field if non-nil, zero value otherwise.

### GetSelectorDetailsOk

`func (o *Deployment) GetSelectorDetailsOk() (*[]string, bool)`

GetSelectorDetailsOk returns a tuple with the SelectorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorDetails

`func (o *Deployment) SetSelectorDetails(v []string)`

SetSelectorDetails sets SelectorDetails field to given value.


### SetSelectorDetailsNil

`func (o *Deployment) SetSelectorDetailsNil(b bool)`

 SetSelectorDetailsNil sets the value for SelectorDetails to be an explicit nil

### UnsetSelectorDetails
`func (o *Deployment) UnsetSelectorDetails()`

UnsetSelectorDetails ensures that no value is present for SelectorDetails, not even an explicit nil
### GetSpecReplicas

`func (o *Deployment) GetSpecReplicas() int32`

GetSpecReplicas returns the SpecReplicas field if non-nil, zero value otherwise.

### GetSpecReplicasOk

`func (o *Deployment) GetSpecReplicasOk() (*int32, bool)`

GetSpecReplicasOk returns a tuple with the SpecReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecReplicas

`func (o *Deployment) SetSpecReplicas(v int32)`

SetSpecReplicas sets SpecReplicas field to given value.


### GetStatusReplicas

`func (o *Deployment) GetStatusReplicas() int32`

GetStatusReplicas returns the StatusReplicas field if non-nil, zero value otherwise.

### GetStatusReplicasOk

`func (o *Deployment) GetStatusReplicasOk() (*int32, bool)`

GetStatusReplicasOk returns a tuple with the StatusReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReplicas

`func (o *Deployment) SetStatusReplicas(v int32)`

SetStatusReplicas sets StatusReplicas field to given value.


### GetStrategyType

`func (o *Deployment) GetStrategyType() string`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *Deployment) GetStrategyTypeOk() (*string, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *Deployment) SetStrategyType(v string)`

SetStrategyType sets StrategyType field to given value.


### SetStrategyTypeNil

`func (o *Deployment) SetStrategyTypeNil(b bool)`

 SetStrategyTypeNil sets the value for StrategyType to be an explicit nil

### UnsetStrategyType
`func (o *Deployment) UnsetStrategyType()`

UnsetStrategyType ensures that no value is present for StrategyType, not even an explicit nil
### GetUid

`func (o *Deployment) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Deployment) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Deployment) SetUid(v string)`

SetUid sets Uid field to given value.


### GetUnavailableReplicas

`func (o *Deployment) GetUnavailableReplicas() int32`

GetUnavailableReplicas returns the UnavailableReplicas field if non-nil, zero value otherwise.

### GetUnavailableReplicasOk

`func (o *Deployment) GetUnavailableReplicasOk() (*int32, bool)`

GetUnavailableReplicasOk returns a tuple with the UnavailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableReplicas

`func (o *Deployment) SetUnavailableReplicas(v int32)`

SetUnavailableReplicas sets UnavailableReplicas field to given value.


### GetUpdatedReplicas

`func (o *Deployment) GetUpdatedReplicas() int32`

GetUpdatedReplicas returns the UpdatedReplicas field if non-nil, zero value otherwise.

### GetUpdatedReplicasOk

`func (o *Deployment) GetUpdatedReplicasOk() (*int32, bool)`

GetUpdatedReplicasOk returns a tuple with the UpdatedReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedReplicas

`func (o *Deployment) SetUpdatedReplicas(v int32)`

SetUpdatedReplicas sets UpdatedReplicas field to given value.


### GetYaml

`func (o *Deployment) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *Deployment) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *Deployment) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


