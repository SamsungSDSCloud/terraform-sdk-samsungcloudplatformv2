# Namespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**CreatedAt** | **time.Time** | Created At | 
**Labels** | **[]string** | Labels | 
**Name** | **string** | Namespace Name | 
**NamespaceStatus** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 
**Yaml** | **string** | YAML | 

## Methods

### NewNamespace

`func NewNamespace(age string, annotations []string, clusterId string, clusterName string, createdAt time.Time, labels []string, name string, namespaceStatus NullableString, uid string, yaml string, ) *Namespace`

NewNamespace instantiates a new Namespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceWithDefaults

`func NewNamespaceWithDefaults() *Namespace`

NewNamespaceWithDefaults instantiates a new Namespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *Namespace) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Namespace) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Namespace) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *Namespace) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Namespace) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Namespace) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetClusterId

`func (o *Namespace) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Namespace) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Namespace) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *Namespace) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Namespace) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Namespace) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCreatedAt

`func (o *Namespace) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Namespace) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Namespace) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLabels

`func (o *Namespace) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Namespace) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Namespace) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *Namespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Namespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Namespace) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceStatus

`func (o *Namespace) GetNamespaceStatus() string`

GetNamespaceStatus returns the NamespaceStatus field if non-nil, zero value otherwise.

### GetNamespaceStatusOk

`func (o *Namespace) GetNamespaceStatusOk() (*string, bool)`

GetNamespaceStatusOk returns a tuple with the NamespaceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceStatus

`func (o *Namespace) SetNamespaceStatus(v string)`

SetNamespaceStatus sets NamespaceStatus field to given value.


### SetNamespaceStatusNil

`func (o *Namespace) SetNamespaceStatusNil(b bool)`

 SetNamespaceStatusNil sets the value for NamespaceStatus to be an explicit nil

### UnsetNamespaceStatus
`func (o *Namespace) UnsetNamespaceStatus()`

UnsetNamespaceStatus ensures that no value is present for NamespaceStatus, not even an explicit nil
### GetUid

`func (o *Namespace) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Namespace) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Namespace) SetUid(v string)`

SetUid sets Uid field to given value.


### GetYaml

`func (o *Namespace) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *Namespace) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *Namespace) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


