# GslbCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | The GSLB Algorithm. | 
**Description** | Pointer to **NullableString** |  | [optional] 
**EnvUsage** | **string** | The GSLB Environment Usage. | 
**HealthCheck** | Pointer to [**NullableGslbHealthCheck**](GslbHealthCheck.md) |  | [optional] 
**Name** | **string** | The Name of the gslb. | 
**Resources** | [**[]GslbResource**](GslbResource.md) |  | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]

## Methods

### NewGslbCreateRequest

`func NewGslbCreateRequest(algorithm string, envUsage string, name string, resources []GslbResource, ) *GslbCreateRequest`

NewGslbCreateRequest instantiates a new GslbCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGslbCreateRequestWithDefaults

`func NewGslbCreateRequestWithDefaults() *GslbCreateRequest`

NewGslbCreateRequestWithDefaults instantiates a new GslbCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *GslbCreateRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *GslbCreateRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *GslbCreateRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetDescription

`func (o *GslbCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GslbCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GslbCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GslbCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GslbCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GslbCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnvUsage

`func (o *GslbCreateRequest) GetEnvUsage() string`

GetEnvUsage returns the EnvUsage field if non-nil, zero value otherwise.

### GetEnvUsageOk

`func (o *GslbCreateRequest) GetEnvUsageOk() (*string, bool)`

GetEnvUsageOk returns a tuple with the EnvUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvUsage

`func (o *GslbCreateRequest) SetEnvUsage(v string)`

SetEnvUsage sets EnvUsage field to given value.


### GetHealthCheck

`func (o *GslbCreateRequest) GetHealthCheck() GslbHealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *GslbCreateRequest) GetHealthCheckOk() (*GslbHealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *GslbCreateRequest) SetHealthCheck(v GslbHealthCheck)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *GslbCreateRequest) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### SetHealthCheckNil

`func (o *GslbCreateRequest) SetHealthCheckNil(b bool)`

 SetHealthCheckNil sets the value for HealthCheck to be an explicit nil

### UnsetHealthCheck
`func (o *GslbCreateRequest) UnsetHealthCheck()`

UnsetHealthCheck ensures that no value is present for HealthCheck, not even an explicit nil
### GetName

`func (o *GslbCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GslbCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GslbCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *GslbCreateRequest) GetResources() []GslbResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GslbCreateRequest) GetResourcesOk() (*[]GslbResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GslbCreateRequest) SetResources(v []GslbResource)`

SetResources sets Resources field to given value.


### GetTags

`func (o *GslbCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GslbCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GslbCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GslbCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


