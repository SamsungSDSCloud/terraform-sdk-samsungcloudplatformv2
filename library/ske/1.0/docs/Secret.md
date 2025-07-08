# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**CreatedAt** | **time.Time** | Created At | 
**Labels** | **[]string** | Labels | 
**Name** | **string** | Secret Name | 
**NamespaceName** | **string** | Namespace Name | 
**SecretData** | **[]string** | Secret Data | 
**SecretType** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 
**Yaml** | **string** | YAML | 

## Methods

### NewSecret

`func NewSecret(age string, annotations []string, clusterId string, clusterName string, createdAt time.Time, labels []string, name string, namespaceName string, secretData []string, secretType NullableString, uid string, yaml string, ) *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *Secret) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Secret) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Secret) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *Secret) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Secret) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Secret) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetClusterId

`func (o *Secret) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Secret) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Secret) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *Secret) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Secret) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Secret) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCreatedAt

`func (o *Secret) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Secret) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Secret) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLabels

`func (o *Secret) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Secret) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Secret) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *Secret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Secret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Secret) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *Secret) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *Secret) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *Secret) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetSecretData

`func (o *Secret) GetSecretData() []string`

GetSecretData returns the SecretData field if non-nil, zero value otherwise.

### GetSecretDataOk

`func (o *Secret) GetSecretDataOk() (*[]string, bool)`

GetSecretDataOk returns a tuple with the SecretData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretData

`func (o *Secret) SetSecretData(v []string)`

SetSecretData sets SecretData field to given value.


### GetSecretType

`func (o *Secret) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *Secret) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *Secret) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.


### SetSecretTypeNil

`func (o *Secret) SetSecretTypeNil(b bool)`

 SetSecretTypeNil sets the value for SecretType to be an explicit nil

### UnsetSecretType
`func (o *Secret) UnsetSecretType()`

UnsetSecretType ensures that no value is present for SecretType, not even an explicit nil
### GetUid

`func (o *Secret) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Secret) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Secret) SetUid(v string)`

SetUid sets Uid field to given value.


### GetYaml

`func (o *Secret) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *Secret) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *Secret) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


