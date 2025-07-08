# SecretSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**ClusterId** | **string** | Cluster ID | 
**CreatedAt** | **time.Time** | Created At | 
**Name** | **string** | Secret Name | 
**NamespaceName** | **string** | Namespace Name | 
**SecretData** | **[]string** | Secret Data | 
**SecretType** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 

## Methods

### NewSecretSummary

`func NewSecretSummary(age string, clusterId string, createdAt time.Time, name string, namespaceName string, secretData []string, secretType NullableString, uid string, ) *SecretSummary`

NewSecretSummary instantiates a new SecretSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretSummaryWithDefaults

`func NewSecretSummaryWithDefaults() *SecretSummary`

NewSecretSummaryWithDefaults instantiates a new SecretSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *SecretSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *SecretSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *SecretSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetClusterId

`func (o *SecretSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *SecretSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *SecretSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCreatedAt

`func (o *SecretSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecretSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecretSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *SecretSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *SecretSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *SecretSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *SecretSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetSecretData

`func (o *SecretSummary) GetSecretData() []string`

GetSecretData returns the SecretData field if non-nil, zero value otherwise.

### GetSecretDataOk

`func (o *SecretSummary) GetSecretDataOk() (*[]string, bool)`

GetSecretDataOk returns a tuple with the SecretData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretData

`func (o *SecretSummary) SetSecretData(v []string)`

SetSecretData sets SecretData field to given value.


### GetSecretType

`func (o *SecretSummary) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *SecretSummary) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *SecretSummary) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.


### SetSecretTypeNil

`func (o *SecretSummary) SetSecretTypeNil(b bool)`

 SetSecretTypeNil sets the value for SecretType to be an explicit nil

### UnsetSecretType
`func (o *SecretSummary) UnsetSecretType()`

UnsetSecretType ensures that no value is present for SecretType, not even an explicit nil
### GetUid

`func (o *SecretSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SecretSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SecretSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


