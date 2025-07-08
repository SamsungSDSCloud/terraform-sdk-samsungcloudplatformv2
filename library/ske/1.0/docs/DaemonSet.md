# DaemonSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**CreatedAt** | **time.Time** | Created At | 
**CurrentNumberScheduled** | **int32** | Current Number Scheduled | 
**DesiredNumberScheduled** | **int32** | Desired Number Scheduled | 
**Labels** | **[]string** | Labels | 
**Name** | **string** | Daemon Set Name | 
**NamespaceName** | **string** | Namespace Name | 
**NodeSelector** | **[]string** | Node Selector | 
**NumberAvailable** | **int32** | Number Available | 
**NumberMisscheduled** | **int32** | Number Misscheduled | 
**NumberReady** | **int32** | Number Ready | 
**SelectorDetails** | **[]string** | Selector Details | 
**Selectors** | **[]string** | Selectors | 
**Uid** | **string** | Resource ID | 
**UpdatedNumberScheduled** | **int32** | Updated Number Scheduled | 
**Yaml** | **string** | YAML | 

## Methods

### NewDaemonSet

`func NewDaemonSet(age string, annotations []string, clusterId string, clusterName string, createdAt time.Time, currentNumberScheduled int32, desiredNumberScheduled int32, labels []string, name string, namespaceName string, nodeSelector []string, numberAvailable int32, numberMisscheduled int32, numberReady int32, selectorDetails []string, selectors []string, uid string, updatedNumberScheduled int32, yaml string, ) *DaemonSet`

NewDaemonSet instantiates a new DaemonSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaemonSetWithDefaults

`func NewDaemonSetWithDefaults() *DaemonSet`

NewDaemonSetWithDefaults instantiates a new DaemonSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *DaemonSet) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *DaemonSet) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *DaemonSet) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *DaemonSet) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *DaemonSet) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *DaemonSet) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetClusterId

`func (o *DaemonSet) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *DaemonSet) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *DaemonSet) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *DaemonSet) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *DaemonSet) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *DaemonSet) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCreatedAt

`func (o *DaemonSet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaemonSet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaemonSet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrentNumberScheduled

`func (o *DaemonSet) GetCurrentNumberScheduled() int32`

GetCurrentNumberScheduled returns the CurrentNumberScheduled field if non-nil, zero value otherwise.

### GetCurrentNumberScheduledOk

`func (o *DaemonSet) GetCurrentNumberScheduledOk() (*int32, bool)`

GetCurrentNumberScheduledOk returns a tuple with the CurrentNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNumberScheduled

`func (o *DaemonSet) SetCurrentNumberScheduled(v int32)`

SetCurrentNumberScheduled sets CurrentNumberScheduled field to given value.


### GetDesiredNumberScheduled

`func (o *DaemonSet) GetDesiredNumberScheduled() int32`

GetDesiredNumberScheduled returns the DesiredNumberScheduled field if non-nil, zero value otherwise.

### GetDesiredNumberScheduledOk

`func (o *DaemonSet) GetDesiredNumberScheduledOk() (*int32, bool)`

GetDesiredNumberScheduledOk returns a tuple with the DesiredNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNumberScheduled

`func (o *DaemonSet) SetDesiredNumberScheduled(v int32)`

SetDesiredNumberScheduled sets DesiredNumberScheduled field to given value.


### GetLabels

`func (o *DaemonSet) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *DaemonSet) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *DaemonSet) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *DaemonSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaemonSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaemonSet) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *DaemonSet) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *DaemonSet) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *DaemonSet) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetNodeSelector

`func (o *DaemonSet) GetNodeSelector() []string`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *DaemonSet) GetNodeSelectorOk() (*[]string, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *DaemonSet) SetNodeSelector(v []string)`

SetNodeSelector sets NodeSelector field to given value.


### GetNumberAvailable

`func (o *DaemonSet) GetNumberAvailable() int32`

GetNumberAvailable returns the NumberAvailable field if non-nil, zero value otherwise.

### GetNumberAvailableOk

`func (o *DaemonSet) GetNumberAvailableOk() (*int32, bool)`

GetNumberAvailableOk returns a tuple with the NumberAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberAvailable

`func (o *DaemonSet) SetNumberAvailable(v int32)`

SetNumberAvailable sets NumberAvailable field to given value.


### GetNumberMisscheduled

`func (o *DaemonSet) GetNumberMisscheduled() int32`

GetNumberMisscheduled returns the NumberMisscheduled field if non-nil, zero value otherwise.

### GetNumberMisscheduledOk

`func (o *DaemonSet) GetNumberMisscheduledOk() (*int32, bool)`

GetNumberMisscheduledOk returns a tuple with the NumberMisscheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberMisscheduled

`func (o *DaemonSet) SetNumberMisscheduled(v int32)`

SetNumberMisscheduled sets NumberMisscheduled field to given value.


### GetNumberReady

`func (o *DaemonSet) GetNumberReady() int32`

GetNumberReady returns the NumberReady field if non-nil, zero value otherwise.

### GetNumberReadyOk

`func (o *DaemonSet) GetNumberReadyOk() (*int32, bool)`

GetNumberReadyOk returns a tuple with the NumberReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberReady

`func (o *DaemonSet) SetNumberReady(v int32)`

SetNumberReady sets NumberReady field to given value.


### GetSelectorDetails

`func (o *DaemonSet) GetSelectorDetails() []string`

GetSelectorDetails returns the SelectorDetails field if non-nil, zero value otherwise.

### GetSelectorDetailsOk

`func (o *DaemonSet) GetSelectorDetailsOk() (*[]string, bool)`

GetSelectorDetailsOk returns a tuple with the SelectorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorDetails

`func (o *DaemonSet) SetSelectorDetails(v []string)`

SetSelectorDetails sets SelectorDetails field to given value.


### GetSelectors

`func (o *DaemonSet) GetSelectors() []string`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *DaemonSet) GetSelectorsOk() (*[]string, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *DaemonSet) SetSelectors(v []string)`

SetSelectors sets Selectors field to given value.


### GetUid

`func (o *DaemonSet) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DaemonSet) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DaemonSet) SetUid(v string)`

SetUid sets Uid field to given value.


### GetUpdatedNumberScheduled

`func (o *DaemonSet) GetUpdatedNumberScheduled() int32`

GetUpdatedNumberScheduled returns the UpdatedNumberScheduled field if non-nil, zero value otherwise.

### GetUpdatedNumberScheduledOk

`func (o *DaemonSet) GetUpdatedNumberScheduledOk() (*int32, bool)`

GetUpdatedNumberScheduledOk returns a tuple with the UpdatedNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedNumberScheduled

`func (o *DaemonSet) SetUpdatedNumberScheduled(v int32)`

SetUpdatedNumberScheduled sets UpdatedNumberScheduled field to given value.


### GetYaml

`func (o *DaemonSet) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *DaemonSet) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *DaemonSet) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


