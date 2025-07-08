# IngressClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**Controller** | **NullableString** |  | 
**CreatedAt** | **time.Time** | Created At | 
**IsDefaultClass** | **string** | Is Default Class | 
**Labels** | **[]string** | Labels | 
**Name** | **string** | Ingress Class Name | 
**Parameter** | [**NullableParameter**](Parameter.md) |  | 
**Uid** | **string** | Resource ID | 
**Yaml** | **string** | YAML | 

## Methods

### NewIngressClass

`func NewIngressClass(age string, annotations []string, clusterId string, clusterName string, controller NullableString, createdAt time.Time, isDefaultClass string, labels []string, name string, parameter NullableParameter, uid string, yaml string, ) *IngressClass`

NewIngressClass instantiates a new IngressClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressClassWithDefaults

`func NewIngressClassWithDefaults() *IngressClass`

NewIngressClassWithDefaults instantiates a new IngressClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *IngressClass) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *IngressClass) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *IngressClass) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *IngressClass) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *IngressClass) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *IngressClass) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetClusterId

`func (o *IngressClass) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *IngressClass) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *IngressClass) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *IngressClass) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *IngressClass) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *IngressClass) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetController

`func (o *IngressClass) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *IngressClass) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *IngressClass) SetController(v string)`

SetController sets Controller field to given value.


### SetControllerNil

`func (o *IngressClass) SetControllerNil(b bool)`

 SetControllerNil sets the value for Controller to be an explicit nil

### UnsetController
`func (o *IngressClass) UnsetController()`

UnsetController ensures that no value is present for Controller, not even an explicit nil
### GetCreatedAt

`func (o *IngressClass) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IngressClass) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IngressClass) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIsDefaultClass

`func (o *IngressClass) GetIsDefaultClass() string`

GetIsDefaultClass returns the IsDefaultClass field if non-nil, zero value otherwise.

### GetIsDefaultClassOk

`func (o *IngressClass) GetIsDefaultClassOk() (*string, bool)`

GetIsDefaultClassOk returns a tuple with the IsDefaultClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultClass

`func (o *IngressClass) SetIsDefaultClass(v string)`

SetIsDefaultClass sets IsDefaultClass field to given value.


### GetLabels

`func (o *IngressClass) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IngressClass) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IngressClass) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *IngressClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngressClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngressClass) SetName(v string)`

SetName sets Name field to given value.


### GetParameter

`func (o *IngressClass) GetParameter() Parameter`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *IngressClass) GetParameterOk() (*Parameter, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *IngressClass) SetParameter(v Parameter)`

SetParameter sets Parameter field to given value.


### SetParameterNil

`func (o *IngressClass) SetParameterNil(b bool)`

 SetParameterNil sets the value for Parameter to be an explicit nil

### UnsetParameter
`func (o *IngressClass) UnsetParameter()`

UnsetParameter ensures that no value is present for Parameter, not even an explicit nil
### GetUid

`func (o *IngressClass) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IngressClass) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IngressClass) SetUid(v string)`

SetUid sets Uid field to given value.


### GetYaml

`func (o *IngressClass) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *IngressClass) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *IngressClass) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


