# CronJobListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**CronJobs** | [**[]CronJobSummary**](CronJobSummary.md) |  | 
**Links** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewCronJobListResponse

`func NewCronJobListResponse(cronJobs []CronJobSummary, ) *CronJobListResponse`

NewCronJobListResponse instantiates a new CronJobListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronJobListResponseWithDefaults

`func NewCronJobListResponseWithDefaults() *CronJobListResponse`

NewCronJobListResponseWithDefaults instantiates a new CronJobListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CronJobListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CronJobListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CronJobListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CronJobListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *CronJobListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *CronJobListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetCronJobs

`func (o *CronJobListResponse) GetCronJobs() []CronJobSummary`

GetCronJobs returns the CronJobs field if non-nil, zero value otherwise.

### GetCronJobsOk

`func (o *CronJobListResponse) GetCronJobsOk() (*[]CronJobSummary, bool)`

GetCronJobsOk returns a tuple with the CronJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronJobs

`func (o *CronJobListResponse) SetCronJobs(v []CronJobSummary)`

SetCronJobs sets CronJobs field to given value.


### GetLinks

`func (o *CronJobListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CronJobListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CronJobListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CronJobListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *CronJobListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *CronJobListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


