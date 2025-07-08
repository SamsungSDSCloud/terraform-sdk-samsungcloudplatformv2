# GslbResponseCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | The GSLB Algorithm. | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**EnvUsage** | **string** | The GSLB Environment Usage. | 
**Id** | **string** | ID | 
**LinkedResourceCount** | **int32** | The GSLB Linked Resource Count. | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | The Name of the gslb. | 
**State** | **string** | The GSLB State. | 

## Methods

### NewGslbResponseCommon

`func NewGslbResponseCommon(algorithm string, createdAt time.Time, createdBy string, envUsage string, id string, linkedResourceCount int32, modifiedAt time.Time, modifiedBy string, name string, state string, ) *GslbResponseCommon`

NewGslbResponseCommon instantiates a new GslbResponseCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGslbResponseCommonWithDefaults

`func NewGslbResponseCommonWithDefaults() *GslbResponseCommon`

NewGslbResponseCommonWithDefaults instantiates a new GslbResponseCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *GslbResponseCommon) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *GslbResponseCommon) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *GslbResponseCommon) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetCreatedAt

`func (o *GslbResponseCommon) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GslbResponseCommon) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GslbResponseCommon) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *GslbResponseCommon) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GslbResponseCommon) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GslbResponseCommon) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *GslbResponseCommon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GslbResponseCommon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GslbResponseCommon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GslbResponseCommon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GslbResponseCommon) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GslbResponseCommon) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnvUsage

`func (o *GslbResponseCommon) GetEnvUsage() string`

GetEnvUsage returns the EnvUsage field if non-nil, zero value otherwise.

### GetEnvUsageOk

`func (o *GslbResponseCommon) GetEnvUsageOk() (*string, bool)`

GetEnvUsageOk returns a tuple with the EnvUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvUsage

`func (o *GslbResponseCommon) SetEnvUsage(v string)`

SetEnvUsage sets EnvUsage field to given value.


### GetId

`func (o *GslbResponseCommon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GslbResponseCommon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GslbResponseCommon) SetId(v string)`

SetId sets Id field to given value.


### GetLinkedResourceCount

`func (o *GslbResponseCommon) GetLinkedResourceCount() int32`

GetLinkedResourceCount returns the LinkedResourceCount field if non-nil, zero value otherwise.

### GetLinkedResourceCountOk

`func (o *GslbResponseCommon) GetLinkedResourceCountOk() (*int32, bool)`

GetLinkedResourceCountOk returns a tuple with the LinkedResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedResourceCount

`func (o *GslbResponseCommon) SetLinkedResourceCount(v int32)`

SetLinkedResourceCount sets LinkedResourceCount field to given value.


### GetModifiedAt

`func (o *GslbResponseCommon) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *GslbResponseCommon) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *GslbResponseCommon) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *GslbResponseCommon) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *GslbResponseCommon) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *GslbResponseCommon) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *GslbResponseCommon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GslbResponseCommon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GslbResponseCommon) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *GslbResponseCommon) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GslbResponseCommon) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GslbResponseCommon) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


