# CronJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveJobCount** | **int32** | Active Job Count | 
**ActiveJobs** | **[]string** | Active Jobs | 
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
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
**Yaml** | **string** | YAML | 

## Methods

### NewCronJob

`func NewCronJob(activeJobCount int32, activeJobs []string, age string, annotations []string, clusterId string, clusterName string, completions int32, concurrencyPolicy NullableString, createdAt time.Time, failedJobsHistoryLimit int32, labels []string, lastScheduleTime NullableTime, name string, namespaceName string, parallelism int32, schedule NullableString, selectorDetails []string, selectors []string, startingDeadlineSeconds int32, successfulJobsHistoryLimit int32, suspended bool, uid string, yaml string, ) *CronJob`

NewCronJob instantiates a new CronJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronJobWithDefaults

`func NewCronJobWithDefaults() *CronJob`

NewCronJobWithDefaults instantiates a new CronJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveJobCount

`func (o *CronJob) GetActiveJobCount() int32`

GetActiveJobCount returns the ActiveJobCount field if non-nil, zero value otherwise.

### GetActiveJobCountOk

`func (o *CronJob) GetActiveJobCountOk() (*int32, bool)`

GetActiveJobCountOk returns a tuple with the ActiveJobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveJobCount

`func (o *CronJob) SetActiveJobCount(v int32)`

SetActiveJobCount sets ActiveJobCount field to given value.


### GetActiveJobs

`func (o *CronJob) GetActiveJobs() []string`

GetActiveJobs returns the ActiveJobs field if non-nil, zero value otherwise.

### GetActiveJobsOk

`func (o *CronJob) GetActiveJobsOk() (*[]string, bool)`

GetActiveJobsOk returns a tuple with the ActiveJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveJobs

`func (o *CronJob) SetActiveJobs(v []string)`

SetActiveJobs sets ActiveJobs field to given value.


### GetAge

`func (o *CronJob) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *CronJob) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *CronJob) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *CronJob) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CronJob) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CronJob) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetClusterId

`func (o *CronJob) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CronJob) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CronJob) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *CronJob) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CronJob) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CronJob) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCompletions

`func (o *CronJob) GetCompletions() int32`

GetCompletions returns the Completions field if non-nil, zero value otherwise.

### GetCompletionsOk

`func (o *CronJob) GetCompletionsOk() (*int32, bool)`

GetCompletionsOk returns a tuple with the Completions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletions

`func (o *CronJob) SetCompletions(v int32)`

SetCompletions sets Completions field to given value.


### GetConcurrencyPolicy

`func (o *CronJob) GetConcurrencyPolicy() string`

GetConcurrencyPolicy returns the ConcurrencyPolicy field if non-nil, zero value otherwise.

### GetConcurrencyPolicyOk

`func (o *CronJob) GetConcurrencyPolicyOk() (*string, bool)`

GetConcurrencyPolicyOk returns a tuple with the ConcurrencyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyPolicy

`func (o *CronJob) SetConcurrencyPolicy(v string)`

SetConcurrencyPolicy sets ConcurrencyPolicy field to given value.


### SetConcurrencyPolicyNil

`func (o *CronJob) SetConcurrencyPolicyNil(b bool)`

 SetConcurrencyPolicyNil sets the value for ConcurrencyPolicy to be an explicit nil

### UnsetConcurrencyPolicy
`func (o *CronJob) UnsetConcurrencyPolicy()`

UnsetConcurrencyPolicy ensures that no value is present for ConcurrencyPolicy, not even an explicit nil
### GetCreatedAt

`func (o *CronJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CronJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CronJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFailedJobsHistoryLimit

`func (o *CronJob) GetFailedJobsHistoryLimit() int32`

GetFailedJobsHistoryLimit returns the FailedJobsHistoryLimit field if non-nil, zero value otherwise.

### GetFailedJobsHistoryLimitOk

`func (o *CronJob) GetFailedJobsHistoryLimitOk() (*int32, bool)`

GetFailedJobsHistoryLimitOk returns a tuple with the FailedJobsHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedJobsHistoryLimit

`func (o *CronJob) SetFailedJobsHistoryLimit(v int32)`

SetFailedJobsHistoryLimit sets FailedJobsHistoryLimit field to given value.


### GetLabels

`func (o *CronJob) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CronJob) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CronJob) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetLastScheduleTime

`func (o *CronJob) GetLastScheduleTime() time.Time`

GetLastScheduleTime returns the LastScheduleTime field if non-nil, zero value otherwise.

### GetLastScheduleTimeOk

`func (o *CronJob) GetLastScheduleTimeOk() (*time.Time, bool)`

GetLastScheduleTimeOk returns a tuple with the LastScheduleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScheduleTime

`func (o *CronJob) SetLastScheduleTime(v time.Time)`

SetLastScheduleTime sets LastScheduleTime field to given value.


### SetLastScheduleTimeNil

`func (o *CronJob) SetLastScheduleTimeNil(b bool)`

 SetLastScheduleTimeNil sets the value for LastScheduleTime to be an explicit nil

### UnsetLastScheduleTime
`func (o *CronJob) UnsetLastScheduleTime()`

UnsetLastScheduleTime ensures that no value is present for LastScheduleTime, not even an explicit nil
### GetName

`func (o *CronJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CronJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CronJob) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *CronJob) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *CronJob) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *CronJob) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetParallelism

`func (o *CronJob) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *CronJob) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *CronJob) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.


### GetSchedule

`func (o *CronJob) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CronJob) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CronJob) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### SetScheduleNil

`func (o *CronJob) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *CronJob) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetSelectorDetails

`func (o *CronJob) GetSelectorDetails() []string`

GetSelectorDetails returns the SelectorDetails field if non-nil, zero value otherwise.

### GetSelectorDetailsOk

`func (o *CronJob) GetSelectorDetailsOk() (*[]string, bool)`

GetSelectorDetailsOk returns a tuple with the SelectorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorDetails

`func (o *CronJob) SetSelectorDetails(v []string)`

SetSelectorDetails sets SelectorDetails field to given value.


### GetSelectors

`func (o *CronJob) GetSelectors() []string`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *CronJob) GetSelectorsOk() (*[]string, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *CronJob) SetSelectors(v []string)`

SetSelectors sets Selectors field to given value.


### GetStartingDeadlineSeconds

`func (o *CronJob) GetStartingDeadlineSeconds() int32`

GetStartingDeadlineSeconds returns the StartingDeadlineSeconds field if non-nil, zero value otherwise.

### GetStartingDeadlineSecondsOk

`func (o *CronJob) GetStartingDeadlineSecondsOk() (*int32, bool)`

GetStartingDeadlineSecondsOk returns a tuple with the StartingDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingDeadlineSeconds

`func (o *CronJob) SetStartingDeadlineSeconds(v int32)`

SetStartingDeadlineSeconds sets StartingDeadlineSeconds field to given value.


### GetSuccessfulJobsHistoryLimit

`func (o *CronJob) GetSuccessfulJobsHistoryLimit() int32`

GetSuccessfulJobsHistoryLimit returns the SuccessfulJobsHistoryLimit field if non-nil, zero value otherwise.

### GetSuccessfulJobsHistoryLimitOk

`func (o *CronJob) GetSuccessfulJobsHistoryLimitOk() (*int32, bool)`

GetSuccessfulJobsHistoryLimitOk returns a tuple with the SuccessfulJobsHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulJobsHistoryLimit

`func (o *CronJob) SetSuccessfulJobsHistoryLimit(v int32)`

SetSuccessfulJobsHistoryLimit sets SuccessfulJobsHistoryLimit field to given value.


### GetSuspended

`func (o *CronJob) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *CronJob) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *CronJob) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.


### GetUid

`func (o *CronJob) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CronJob) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CronJob) SetUid(v string)`

SetUid sets Uid field to given value.


### GetYaml

`func (o *CronJob) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *CronJob) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *CronJob) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


