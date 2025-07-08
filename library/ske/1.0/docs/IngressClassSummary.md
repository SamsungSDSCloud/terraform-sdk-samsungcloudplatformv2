# IngressClassSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**ClusterId** | **string** | Cluster ID | 
**Controller** | **NullableString** |  | 
**CreatedAt** | **time.Time** | Created At | 
**IsDefaultClass** | **string** | Is Default Class | 
**Name** | **string** | Ingress Class Name | 
**Parameter** | [**NullableParameter**](Parameter.md) |  | 
**Uid** | **string** | Resource ID | 

## Methods

### NewIngressClassSummary

`func NewIngressClassSummary(age string, clusterId string, controller NullableString, createdAt time.Time, isDefaultClass string, name string, parameter NullableParameter, uid string, ) *IngressClassSummary`

NewIngressClassSummary instantiates a new IngressClassSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressClassSummaryWithDefaults

`func NewIngressClassSummaryWithDefaults() *IngressClassSummary`

NewIngressClassSummaryWithDefaults instantiates a new IngressClassSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *IngressClassSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *IngressClassSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *IngressClassSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetClusterId

`func (o *IngressClassSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *IngressClassSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *IngressClassSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetController

`func (o *IngressClassSummary) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *IngressClassSummary) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *IngressClassSummary) SetController(v string)`

SetController sets Controller field to given value.


### SetControllerNil

`func (o *IngressClassSummary) SetControllerNil(b bool)`

 SetControllerNil sets the value for Controller to be an explicit nil

### UnsetController
`func (o *IngressClassSummary) UnsetController()`

UnsetController ensures that no value is present for Controller, not even an explicit nil
### GetCreatedAt

`func (o *IngressClassSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IngressClassSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IngressClassSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIsDefaultClass

`func (o *IngressClassSummary) GetIsDefaultClass() string`

GetIsDefaultClass returns the IsDefaultClass field if non-nil, zero value otherwise.

### GetIsDefaultClassOk

`func (o *IngressClassSummary) GetIsDefaultClassOk() (*string, bool)`

GetIsDefaultClassOk returns a tuple with the IsDefaultClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultClass

`func (o *IngressClassSummary) SetIsDefaultClass(v string)`

SetIsDefaultClass sets IsDefaultClass field to given value.


### GetName

`func (o *IngressClassSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngressClassSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngressClassSummary) SetName(v string)`

SetName sets Name field to given value.


### GetParameter

`func (o *IngressClassSummary) GetParameter() Parameter`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *IngressClassSummary) GetParameterOk() (*Parameter, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *IngressClassSummary) SetParameter(v Parameter)`

SetParameter sets Parameter field to given value.


### SetParameterNil

`func (o *IngressClassSummary) SetParameterNil(b bool)`

 SetParameterNil sets the value for Parameter to be an explicit nil

### UnsetParameter
`func (o *IngressClassSummary) UnsetParameter()`

UnsetParameter ensures that no value is present for Parameter, not even an explicit nil
### GetUid

`func (o *IngressClassSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IngressClassSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IngressClassSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


