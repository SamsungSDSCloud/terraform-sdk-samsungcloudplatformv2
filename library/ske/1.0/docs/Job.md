# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**BackoffLimit** | **NullableInt32** |  | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**Completion** | **string** | Completion | 
**CompletionMode** | **NullableString** |  | 
**Completions** | **NullableInt32** |  | 
**CreatedAt** | **time.Time** | Created At | 
**EndAt** | **NullableTime** |  | 
**Labels** | **[]string** | Labels | 
**Name** | **string** | Job Name | 
**NamespaceName** | **string** | Namespace Name | 
**Parallelism** | **NullableInt32** |  | 
**RunningAge** | **NullableString** |  | 
**SelectorDetails** | **[]string** | Selector Details | 
**Selectors** | **[]string** | Selectors | 
**StartTime** | **NullableTime** |  | 
**Uid** | **string** | Resource ID | 
**Yaml** | **string** | YAML | 

## Methods

### NewJob

`func NewJob(age string, annotations []string, backoffLimit NullableInt32, clusterId string, clusterName string, completion string, completionMode NullableString, completions NullableInt32, createdAt time.Time, endAt NullableTime, labels []string, name string, namespaceName string, parallelism NullableInt32, runningAge NullableString, selectorDetails []string, selectors []string, startTime NullableTime, uid string, yaml string, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *Job) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Job) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Job) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *Job) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Job) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Job) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetBackoffLimit

`func (o *Job) GetBackoffLimit() int32`

GetBackoffLimit returns the BackoffLimit field if non-nil, zero value otherwise.

### GetBackoffLimitOk

`func (o *Job) GetBackoffLimitOk() (*int32, bool)`

GetBackoffLimitOk returns a tuple with the BackoffLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoffLimit

`func (o *Job) SetBackoffLimit(v int32)`

SetBackoffLimit sets BackoffLimit field to given value.


### SetBackoffLimitNil

`func (o *Job) SetBackoffLimitNil(b bool)`

 SetBackoffLimitNil sets the value for BackoffLimit to be an explicit nil

### UnsetBackoffLimit
`func (o *Job) UnsetBackoffLimit()`

UnsetBackoffLimit ensures that no value is present for BackoffLimit, not even an explicit nil
### GetClusterId

`func (o *Job) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Job) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Job) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *Job) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Job) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Job) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCompletion

`func (o *Job) GetCompletion() string`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *Job) GetCompletionOk() (*string, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *Job) SetCompletion(v string)`

SetCompletion sets Completion field to given value.


### GetCompletionMode

`func (o *Job) GetCompletionMode() string`

GetCompletionMode returns the CompletionMode field if non-nil, zero value otherwise.

### GetCompletionModeOk

`func (o *Job) GetCompletionModeOk() (*string, bool)`

GetCompletionModeOk returns a tuple with the CompletionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionMode

`func (o *Job) SetCompletionMode(v string)`

SetCompletionMode sets CompletionMode field to given value.


### SetCompletionModeNil

`func (o *Job) SetCompletionModeNil(b bool)`

 SetCompletionModeNil sets the value for CompletionMode to be an explicit nil

### UnsetCompletionMode
`func (o *Job) UnsetCompletionMode()`

UnsetCompletionMode ensures that no value is present for CompletionMode, not even an explicit nil
### GetCompletions

`func (o *Job) GetCompletions() int32`

GetCompletions returns the Completions field if non-nil, zero value otherwise.

### GetCompletionsOk

`func (o *Job) GetCompletionsOk() (*int32, bool)`

GetCompletionsOk returns a tuple with the Completions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletions

`func (o *Job) SetCompletions(v int32)`

SetCompletions sets Completions field to given value.


### SetCompletionsNil

`func (o *Job) SetCompletionsNil(b bool)`

 SetCompletionsNil sets the value for Completions to be an explicit nil

### UnsetCompletions
`func (o *Job) UnsetCompletions()`

UnsetCompletions ensures that no value is present for Completions, not even an explicit nil
### GetCreatedAt

`func (o *Job) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Job) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Job) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEndAt

`func (o *Job) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *Job) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *Job) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.


### SetEndAtNil

`func (o *Job) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *Job) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetLabels

`func (o *Job) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Job) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Job) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *Job) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Job) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Job) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *Job) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *Job) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *Job) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetParallelism

`func (o *Job) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *Job) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *Job) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.


### SetParallelismNil

`func (o *Job) SetParallelismNil(b bool)`

 SetParallelismNil sets the value for Parallelism to be an explicit nil

### UnsetParallelism
`func (o *Job) UnsetParallelism()`

UnsetParallelism ensures that no value is present for Parallelism, not even an explicit nil
### GetRunningAge

`func (o *Job) GetRunningAge() string`

GetRunningAge returns the RunningAge field if non-nil, zero value otherwise.

### GetRunningAgeOk

`func (o *Job) GetRunningAgeOk() (*string, bool)`

GetRunningAgeOk returns a tuple with the RunningAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningAge

`func (o *Job) SetRunningAge(v string)`

SetRunningAge sets RunningAge field to given value.


### SetRunningAgeNil

`func (o *Job) SetRunningAgeNil(b bool)`

 SetRunningAgeNil sets the value for RunningAge to be an explicit nil

### UnsetRunningAge
`func (o *Job) UnsetRunningAge()`

UnsetRunningAge ensures that no value is present for RunningAge, not even an explicit nil
### GetSelectorDetails

`func (o *Job) GetSelectorDetails() []string`

GetSelectorDetails returns the SelectorDetails field if non-nil, zero value otherwise.

### GetSelectorDetailsOk

`func (o *Job) GetSelectorDetailsOk() (*[]string, bool)`

GetSelectorDetailsOk returns a tuple with the SelectorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorDetails

`func (o *Job) SetSelectorDetails(v []string)`

SetSelectorDetails sets SelectorDetails field to given value.


### GetSelectors

`func (o *Job) GetSelectors() []string`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *Job) GetSelectorsOk() (*[]string, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *Job) SetSelectors(v []string)`

SetSelectors sets Selectors field to given value.


### GetStartTime

`func (o *Job) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Job) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Job) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### SetStartTimeNil

`func (o *Job) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *Job) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetUid

`func (o *Job) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Job) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Job) SetUid(v string)`

SetUid sets Uid field to given value.


### GetYaml

`func (o *Job) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *Job) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *Job) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


