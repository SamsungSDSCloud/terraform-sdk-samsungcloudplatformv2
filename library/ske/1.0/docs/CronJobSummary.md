# CronJobSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveJobCount** | **int32** | Active Job Count | 
**ActiveJobs** | **[]string** | Active Jobs | 
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**ClusterId** | **string** | Cluster ID | 
**Completions** | **int32** | Completions | 
**ConcurrencyPolicy** | **NullableString** |  | 
**CreatedAt** | **time.Time** | Created At | 
**FailedJobsHistoryLimit** | **int32** | Failed Jobs History Limit | 
**Labels** | **[]string** | Labels | 
**LastScheduleTime** | **NullableTime** |  | 
**Name** | **string** | Cron Job Name | 
**NamespaceName** | **string** | Namespace Name | 
**Parallelism** | **int32** | Parallelism | 
**Schedule** | **NullableString** |  | 
**SelectorDetails** | **[]string** | Selector Details | 
**Selectors** | **[]string** | Selectors | 
**StartingDeadlineSeconds** | **int32** | Starting Deadline Seconds | 
**SuccessfulJobsHistoryLimit** | **int32** | Successful Jobs History Limit | 
**Suspended** | **bool** | Suspended | 
**Uid** | **string** | Resource ID | 

## Methods

### NewCronJobSummary

`func NewCronJobSummary(activeJobCount int32, activeJobs []string, age string, annotations []string, clusterId string, completions int32, concurrencyPolicy NullableString, createdAt time.Time, failedJobsHistoryLimit int32, labels []string, lastScheduleTime NullableTime, name string, namespaceName string, parallelism int32, schedule NullableString, selectorDetails []string, selectors []string, startingDeadlineSeconds int32, successfulJobsHistoryLimit int32, suspended bool, uid string, ) *CronJobSummary`

NewCronJobSummary instantiates a new CronJobSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronJobSummaryWithDefaults

`func NewCronJobSummaryWithDefaults() *CronJobSummary`

NewCronJobSummaryWithDefaults instantiates a new CronJobSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveJobCount

`func (o *CronJobSummary) GetActiveJobCount() int32`

GetActiveJobCount returns the ActiveJobCount field if non-nil, zero value otherwise.

### GetActiveJobCountOk

`func (o *CronJobSummary) GetActiveJobCountOk() (*int32, bool)`

GetActiveJobCountOk returns a tuple with the ActiveJobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveJobCount

`func (o *CronJobSummary) SetActiveJobCount(v int32)`

SetActiveJobCount sets ActiveJobCount field to given value.


### GetActiveJobs

`func (o *CronJobSummary) GetActiveJobs() []string`

GetActiveJobs returns the ActiveJobs field if non-nil, zero value otherwise.

### GetActiveJobsOk

`func (o *CronJobSummary) GetActiveJobsOk() (*[]string, bool)`

GetActiveJobsOk returns a tuple with the ActiveJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveJobs

`func (o *CronJobSummary) SetActiveJobs(v []string)`

SetActiveJobs sets ActiveJobs field to given value.


### GetAge

`func (o *CronJobSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *CronJobSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *CronJobSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *CronJobSummary) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CronJobSummary) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CronJobSummary) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetClusterId

`func (o *CronJobSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CronJobSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CronJobSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCompletions

`func (o *CronJobSummary) GetCompletions() int32`

GetCompletions returns the Completions field if non-nil, zero value otherwise.

### GetCompletionsOk

`func (o *CronJobSummary) GetCompletionsOk() (*int32, bool)`

GetCompletionsOk returns a tuple with the Completions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletions

`func (o *CronJobSummary) SetCompletions(v int32)`

SetCompletions sets Completions field to given value.


### GetConcurrencyPolicy

`func (o *CronJobSummary) GetConcurrencyPolicy() string`

GetConcurrencyPolicy returns the ConcurrencyPolicy field if non-nil, zero value otherwise.

### GetConcurrencyPolicyOk

`func (o *CronJobSummary) GetConcurrencyPolicyOk() (*string, bool)`

GetConcurrencyPolicyOk returns a tuple with the ConcurrencyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyPolicy

`func (o *CronJobSummary) SetConcurrencyPolicy(v string)`

SetConcurrencyPolicy sets ConcurrencyPolicy field to given value.


### SetConcurrencyPolicyNil

`func (o *CronJobSummary) SetConcurrencyPolicyNil(b bool)`

 SetConcurrencyPolicyNil sets the value for ConcurrencyPolicy to be an explicit nil

### UnsetConcurrencyPolicy
`func (o *CronJobSummary) UnsetConcurrencyPolicy()`

UnsetConcurrencyPolicy ensures that no value is present for ConcurrencyPolicy, not even an explicit nil
### GetCreatedAt

`func (o *CronJobSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CronJobSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CronJobSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFailedJobsHistoryLimit

`func (o *CronJobSummary) GetFailedJobsHistoryLimit() int32`

GetFailedJobsHistoryLimit returns the FailedJobsHistoryLimit field if non-nil, zero value otherwise.

### GetFailedJobsHistoryLimitOk

`func (o *CronJobSummary) GetFailedJobsHistoryLimitOk() (*int32, bool)`

GetFailedJobsHistoryLimitOk returns a tuple with the FailedJobsHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedJobsHistoryLimit

`func (o *CronJobSummary) SetFailedJobsHistoryLimit(v int32)`

SetFailedJobsHistoryLimit sets FailedJobsHistoryLimit field to given value.


### GetLabels

`func (o *CronJobSummary) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CronJobSummary) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CronJobSummary) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetLastScheduleTime

`func (o *CronJobSummary) GetLastScheduleTime() time.Time`

GetLastScheduleTime returns the LastScheduleTime field if non-nil, zero value otherwise.

### GetLastScheduleTimeOk

`func (o *CronJobSummary) GetLastScheduleTimeOk() (*time.Time, bool)`

GetLastScheduleTimeOk returns a tuple with the LastScheduleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScheduleTime

`func (o *CronJobSummary) SetLastScheduleTime(v time.Time)`

SetLastScheduleTime sets LastScheduleTime field to given value.


### SetLastScheduleTimeNil

`func (o *CronJobSummary) SetLastScheduleTimeNil(b bool)`

 SetLastScheduleTimeNil sets the value for LastScheduleTime to be an explicit nil

### UnsetLastScheduleTime
`func (o *CronJobSummary) UnsetLastScheduleTime()`

UnsetLastScheduleTime ensures that no value is present for LastScheduleTime, not even an explicit nil
### GetName

`func (o *CronJobSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CronJobSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CronJobSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *CronJobSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *CronJobSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *CronJobSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetParallelism

`func (o *CronJobSummary) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *CronJobSummary) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *CronJobSummary) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.


### GetSchedule

`func (o *CronJobSummary) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CronJobSummary) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CronJobSummary) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### SetScheduleNil

`func (o *CronJobSummary) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *CronJobSummary) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetSelectorDetails

`func (o *CronJobSummary) GetSelectorDetails() []string`

GetSelectorDetails returns the SelectorDetails field if non-nil, zero value otherwise.

### GetSelectorDetailsOk

`func (o *CronJobSummary) GetSelectorDetailsOk() (*[]string, bool)`

GetSelectorDetailsOk returns a tuple with the SelectorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorDetails

`func (o *CronJobSummary) SetSelectorDetails(v []string)`

SetSelectorDetails sets SelectorDetails field to given value.


### GetSelectors

`func (o *CronJobSummary) GetSelectors() []string`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *CronJobSummary) GetSelectorsOk() (*[]string, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *CronJobSummary) SetSelectors(v []string)`

SetSelectors sets Selectors field to given value.


### GetStartingDeadlineSeconds

`func (o *CronJobSummary) GetStartingDeadlineSeconds() int32`

GetStartingDeadlineSeconds returns the StartingDeadlineSeconds field if non-nil, zero value otherwise.

### GetStartingDeadlineSecondsOk

`func (o *CronJobSummary) GetStartingDeadlineSecondsOk() (*int32, bool)`

GetStartingDeadlineSecondsOk returns a tuple with the StartingDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingDeadlineSeconds

`func (o *CronJobSummary) SetStartingDeadlineSeconds(v int32)`

SetStartingDeadlineSeconds sets StartingDeadlineSeconds field to given value.


### GetSuccessfulJobsHistoryLimit

`func (o *CronJobSummary) GetSuccessfulJobsHistoryLimit() int32`

GetSuccessfulJobsHistoryLimit returns the SuccessfulJobsHistoryLimit field if non-nil, zero value otherwise.

### GetSuccessfulJobsHistoryLimitOk

`func (o *CronJobSummary) GetSuccessfulJobsHistoryLimitOk() (*int32, bool)`

GetSuccessfulJobsHistoryLimitOk returns a tuple with the SuccessfulJobsHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulJobsHistoryLimit

`func (o *CronJobSummary) SetSuccessfulJobsHistoryLimit(v int32)`

SetSuccessfulJobsHistoryLimit sets SuccessfulJobsHistoryLimit field to given value.


### GetSuspended

`func (o *CronJobSummary) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *CronJobSummary) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *CronJobSummary) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.


### GetUid

`func (o *CronJobSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CronJobSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CronJobSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


