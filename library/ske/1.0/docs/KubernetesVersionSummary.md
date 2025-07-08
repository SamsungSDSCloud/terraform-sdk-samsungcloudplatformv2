# KubernetesVersionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**KubernetesVersion** | **string** | Kubernetes version | 

## Methods

### NewKubernetesVersionSummary

`func NewKubernetesVersionSummary(kubernetesVersion string, ) *KubernetesVersionSummary`

NewKubernetesVersionSummary instantiates a new KubernetesVersionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVersionSummaryWithDefaults

`func NewKubernetesVersionSummaryWithDefaults() *KubernetesVersionSummary`

NewKubernetesVersionSummaryWithDefaults instantiates a new KubernetesVersionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *KubernetesVersionSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesVersionSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesVersionSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KubernetesVersionSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *KubernetesVersionSummary) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KubernetesVersionSummary) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKubernetesVersion

`func (o *KubernetesVersionSummary) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *KubernetesVersionSummary) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *KubernetesVersionSummary) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


