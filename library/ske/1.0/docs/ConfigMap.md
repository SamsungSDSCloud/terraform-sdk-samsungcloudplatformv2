# ConfigMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**ConfigMapBinaryData** | **[]string** | Config Map Binary Data | 
**ConfigMapData** | **[]string** | Config Map Data | 
**CreatedAt** | **time.Time** | Created At | 
**Labels** | **[]string** | Labels | 
**Name** | **string** | Config Map Name | 
**NamespaceName** | **string** | Namespace Name | 
**Uid** | **string** | Resource ID | 
**Yaml** | **string** | YAML | 

## Methods

### NewConfigMap

`func NewConfigMap(age string, annotations []string, clusterId string, clusterName string, configMapBinaryData []string, configMapData []string, createdAt time.Time, labels []string, name string, namespaceName string, uid string, yaml string, ) *ConfigMap`

NewConfigMap instantiates a new ConfigMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigMapWithDefaults

`func NewConfigMapWithDefaults() *ConfigMap`

NewConfigMapWithDefaults instantiates a new ConfigMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *ConfigMap) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ConfigMap) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ConfigMap) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *ConfigMap) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ConfigMap) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ConfigMap) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetClusterId

`func (o *ConfigMap) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConfigMap) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConfigMap) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *ConfigMap) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ConfigMap) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ConfigMap) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetConfigMapBinaryData

`func (o *ConfigMap) GetConfigMapBinaryData() []string`

GetConfigMapBinaryData returns the ConfigMapBinaryData field if non-nil, zero value otherwise.

### GetConfigMapBinaryDataOk

`func (o *ConfigMap) GetConfigMapBinaryDataOk() (*[]string, bool)`

GetConfigMapBinaryDataOk returns a tuple with the ConfigMapBinaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapBinaryData

`func (o *ConfigMap) SetConfigMapBinaryData(v []string)`

SetConfigMapBinaryData sets ConfigMapBinaryData field to given value.


### GetConfigMapData

`func (o *ConfigMap) GetConfigMapData() []string`

GetConfigMapData returns the ConfigMapData field if non-nil, zero value otherwise.

### GetConfigMapDataOk

`func (o *ConfigMap) GetConfigMapDataOk() (*[]string, bool)`

GetConfigMapDataOk returns a tuple with the ConfigMapData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapData

`func (o *ConfigMap) SetConfigMapData(v []string)`

SetConfigMapData sets ConfigMapData field to given value.


### GetCreatedAt

`func (o *ConfigMap) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConfigMap) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConfigMap) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLabels

`func (o *ConfigMap) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ConfigMap) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ConfigMap) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *ConfigMap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigMap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigMap) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *ConfigMap) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *ConfigMap) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *ConfigMap) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetUid

`func (o *ConfigMap) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ConfigMap) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ConfigMap) SetUid(v string)`

SetUid sets Uid field to given value.


### GetYaml

`func (o *ConfigMap) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *ConfigMap) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *ConfigMap) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


