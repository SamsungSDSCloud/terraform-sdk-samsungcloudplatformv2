# JobSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**ClusterId** | **string** | Cluster ID | 
**Completion** | **string** | Completion | 
**CreatedAt** | **time.Time** | Created At | 
**EndAt** | **NullableTime** |  | 
**Labels** | **[]string** | Labels | 
**Name** | **string** | Job Name | 
**NamespaceName** | **string** | Namespace Name | 
**RunningAge** | **NullableString** |  | 
**SelectorDetails** | **[]string** | Selector Details | 
**Selectors** | **[]string** | Selectors | 
**Uid** | **string** | Resource ID | 

## Methods

### NewJobSummary

`func NewJobSummary(age string, annotations []string, clusterId string, completion string, createdAt time.Time, endAt NullableTime, labels []string, name string, namespaceName string, runningAge NullableString, selectorDetails []string, selectors []string, uid string, ) *JobSummary`

NewJobSummary instantiates a new JobSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobSummaryWithDefaults

`func NewJobSummaryWithDefaults() *JobSummary`

NewJobSummaryWithDefaults instantiates a new JobSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *JobSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *JobSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *JobSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *JobSummary) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *JobSummary) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *JobSummary) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetClusterId

`func (o *JobSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *JobSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *JobSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCompletion

`func (o *JobSummary) GetCompletion() string`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *JobSummary) GetCompletionOk() (*string, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *JobSummary) SetCompletion(v string)`

SetCompletion sets Completion field to given value.


### GetCreatedAt

`func (o *JobSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JobSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JobSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEndAt

`func (o *JobSummary) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *JobSummary) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *JobSummary) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.


### SetEndAtNil

`func (o *JobSummary) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *JobSummary) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetLabels

`func (o *JobSummary) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *JobSummary) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *JobSummary) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *JobSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *JobSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *JobSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *JobSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetRunningAge

`func (o *JobSummary) GetRunningAge() string`

GetRunningAge returns the RunningAge field if non-nil, zero value otherwise.

### GetRunningAgeOk

`func (o *JobSummary) GetRunningAgeOk() (*string, bool)`

GetRunningAgeOk returns a tuple with the RunningAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningAge

`func (o *JobSummary) SetRunningAge(v string)`

SetRunningAge sets RunningAge field to given value.


### SetRunningAgeNil

`func (o *JobSummary) SetRunningAgeNil(b bool)`

 SetRunningAgeNil sets the value for RunningAge to be an explicit nil

### UnsetRunningAge
`func (o *JobSummary) UnsetRunningAge()`

UnsetRunningAge ensures that no value is present for RunningAge, not even an explicit nil
### GetSelectorDetails

`func (o *JobSummary) GetSelectorDetails() []string`

GetSelectorDetails returns the SelectorDetails field if non-nil, zero value otherwise.

### GetSelectorDetailsOk

`func (o *JobSummary) GetSelectorDetailsOk() (*[]string, bool)`

GetSelectorDetailsOk returns a tuple with the SelectorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorDetails

`func (o *JobSummary) SetSelectorDetails(v []string)`

SetSelectorDetails sets SelectorDetails field to given value.


### GetSelectors

`func (o *JobSummary) GetSelectors() []string`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *JobSummary) GetSelectorsOk() (*[]string, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *JobSummary) SetSelectors(v []string)`

SetSelectors sets Selectors field to given value.


### GetUid

`func (o *JobSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *JobSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *JobSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


