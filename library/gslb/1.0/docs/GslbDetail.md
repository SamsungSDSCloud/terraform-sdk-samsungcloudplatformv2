# GslbDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | The GSLB Algorithm. | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**EnvUsage** | **string** | The GSLB Environment Usage. | 
**HealthCheck** | [**NullableGslbHeathCheckShowResponse**](GslbHeathCheckShowResponse.md) |  | 
**Id** | **string** | ID | 
**LinkedResourceCount** | **int32** | The GSLB Linked Resource Count. | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | The Name of the gslb. | 
**State** | **string** | The GSLB State. | 

## Methods

### NewGslbDetail

`func NewGslbDetail(algorithm string, createdAt time.Time, createdBy string, envUsage string, healthCheck NullableGslbHeathCheckShowResponse, id string, linkedResourceCount int32, modifiedAt time.Time, modifiedBy string, name string, state string, ) *GslbDetail`

NewGslbDetail instantiates a new GslbDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGslbDetailWithDefaults

`func NewGslbDetailWithDefaults() *GslbDetail`

NewGslbDetailWithDefaults instantiates a new GslbDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *GslbDetail) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *GslbDetail) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *GslbDetail) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetCreatedAt

`func (o *GslbDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GslbDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GslbDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *GslbDetail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GslbDetail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GslbDetail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *GslbDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GslbDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GslbDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GslbDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GslbDetail) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GslbDetail) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnvUsage

`func (o *GslbDetail) GetEnvUsage() string`

GetEnvUsage returns the EnvUsage field if non-nil, zero value otherwise.

### GetEnvUsageOk

`func (o *GslbDetail) GetEnvUsageOk() (*string, bool)`

GetEnvUsageOk returns a tuple with the EnvUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvUsage

`func (o *GslbDetail) SetEnvUsage(v string)`

SetEnvUsage sets EnvUsage field to given value.


### GetHealthCheck

`func (o *GslbDetail) GetHealthCheck() GslbHeathCheckShowResponse`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *GslbDetail) GetHealthCheckOk() (*GslbHeathCheckShowResponse, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *GslbDetail) SetHealthCheck(v GslbHeathCheckShowResponse)`

SetHealthCheck sets HealthCheck field to given value.


### SetHealthCheckNil

`func (o *GslbDetail) SetHealthCheckNil(b bool)`

 SetHealthCheckNil sets the value for HealthCheck to be an explicit nil

### UnsetHealthCheck
`func (o *GslbDetail) UnsetHealthCheck()`

UnsetHealthCheck ensures that no value is present for HealthCheck, not even an explicit nil
### GetId

`func (o *GslbDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GslbDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GslbDetail) SetId(v string)`

SetId sets Id field to given value.


### GetLinkedResourceCount

`func (o *GslbDetail) GetLinkedResourceCount() int32`

GetLinkedResourceCount returns the LinkedResourceCount field if non-nil, zero value otherwise.

### GetLinkedResourceCountOk

`func (o *GslbDetail) GetLinkedResourceCountOk() (*int32, bool)`

GetLinkedResourceCountOk returns a tuple with the LinkedResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedResourceCount

`func (o *GslbDetail) SetLinkedResourceCount(v int32)`

SetLinkedResourceCount sets LinkedResourceCount field to given value.


### GetModifiedAt

`func (o *GslbDetail) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *GslbDetail) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *GslbDetail) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *GslbDetail) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *GslbDetail) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *GslbDetail) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *GslbDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GslbDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GslbDetail) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *GslbDetail) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GslbDetail) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GslbDetail) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


