# StatefulSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**CreatedAt** | **time.Time** | Created At | 
**Labels** | **[]string** | Labels | 
**Name** | **string** | Stateful Set Name | 
**NamespaceName** | **string** | Namespace Name | 
**ReadyReplicas** | **int32** | Ready Replicas | 
**SelectorDetails** | **[]string** | Selector Details | 
**Selectors** | **[]string** | Selectors | 
**SpecReplicas** | **int32** | Spec Replicas | 
**StatusReplicas** | **int32** | Status Replicas | 
**Uid** | **string** | Resource ID | 
**UpdateStrategy** | **NullableString** |  | 
**Yaml** | **string** | YAML | 

## Methods

### NewStatefulSet

`func NewStatefulSet(age string, annotations []string, clusterId string, clusterName string, createdAt time.Time, labels []string, name string, namespaceName string, readyReplicas int32, selectorDetails []string, selectors []string, specReplicas int32, statusReplicas int32, uid string, updateStrategy NullableString, yaml string, ) *StatefulSet`

NewStatefulSet instantiates a new StatefulSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatefulSetWithDefaults

`func NewStatefulSetWithDefaults() *StatefulSet`

NewStatefulSetWithDefaults instantiates a new StatefulSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *StatefulSet) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *StatefulSet) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *StatefulSet) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *StatefulSet) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *StatefulSet) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *StatefulSet) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetClusterId

`func (o *StatefulSet) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StatefulSet) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StatefulSet) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *StatefulSet) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *StatefulSet) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *StatefulSet) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCreatedAt

`func (o *StatefulSet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StatefulSet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StatefulSet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLabels

`func (o *StatefulSet) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StatefulSet) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StatefulSet) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *StatefulSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatefulSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatefulSet) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *StatefulSet) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *StatefulSet) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *StatefulSet) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetReadyReplicas

`func (o *StatefulSet) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *StatefulSet) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *StatefulSet) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.


### GetSelectorDetails

`func (o *StatefulSet) GetSelectorDetails() []string`

GetSelectorDetails returns the SelectorDetails field if non-nil, zero value otherwise.

### GetSelectorDetailsOk

`func (o *StatefulSet) GetSelectorDetailsOk() (*[]string, bool)`

GetSelectorDetailsOk returns a tuple with the SelectorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorDetails

`func (o *StatefulSet) SetSelectorDetails(v []string)`

SetSelectorDetails sets SelectorDetails field to given value.


### GetSelectors

`func (o *StatefulSet) GetSelectors() []string`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *StatefulSet) GetSelectorsOk() (*[]string, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *StatefulSet) SetSelectors(v []string)`

SetSelectors sets Selectors field to given value.


### GetSpecReplicas

`func (o *StatefulSet) GetSpecReplicas() int32`

GetSpecReplicas returns the SpecReplicas field if non-nil, zero value otherwise.

### GetSpecReplicasOk

`func (o *StatefulSet) GetSpecReplicasOk() (*int32, bool)`

GetSpecReplicasOk returns a tuple with the SpecReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecReplicas

`func (o *StatefulSet) SetSpecReplicas(v int32)`

SetSpecReplicas sets SpecReplicas field to given value.


### GetStatusReplicas

`func (o *StatefulSet) GetStatusReplicas() int32`

GetStatusReplicas returns the StatusReplicas field if non-nil, zero value otherwise.

### GetStatusReplicasOk

`func (o *StatefulSet) GetStatusReplicasOk() (*int32, bool)`

GetStatusReplicasOk returns a tuple with the StatusReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReplicas

`func (o *StatefulSet) SetStatusReplicas(v int32)`

SetStatusReplicas sets StatusReplicas field to given value.


### GetUid

`func (o *StatefulSet) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *StatefulSet) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *StatefulSet) SetUid(v string)`

SetUid sets Uid field to given value.


### GetUpdateStrategy

`func (o *StatefulSet) GetUpdateStrategy() string`

GetUpdateStrategy returns the UpdateStrategy field if non-nil, zero value otherwise.

### GetUpdateStrategyOk

`func (o *StatefulSet) GetUpdateStrategyOk() (*string, bool)`

GetUpdateStrategyOk returns a tuple with the UpdateStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStrategy

`func (o *StatefulSet) SetUpdateStrategy(v string)`

SetUpdateStrategy sets UpdateStrategy field to given value.


### SetUpdateStrategyNil

`func (o *StatefulSet) SetUpdateStrategyNil(b bool)`

 SetUpdateStrategyNil sets the value for UpdateStrategy to be an explicit nil

### UnsetUpdateStrategy
`func (o *StatefulSet) UnsetUpdateStrategy()`

UnsetUpdateStrategy ensures that no value is present for UpdateStrategy, not even an explicit nil
### GetYaml

`func (o *StatefulSet) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *StatefulSet) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *StatefulSet) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


