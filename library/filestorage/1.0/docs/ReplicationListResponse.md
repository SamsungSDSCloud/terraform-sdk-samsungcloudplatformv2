# ReplicationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Replications** | [**[]Replication**](Replication.md) |  | 

## Methods

### NewReplicationListResponse

`func NewReplicationListResponse(count int32, replications []Replication, ) *ReplicationListResponse`

NewReplicationListResponse instantiates a new ReplicationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationListResponseWithDefaults

`func NewReplicationListResponseWithDefaults() *ReplicationListResponse`

NewReplicationListResponseWithDefaults instantiates a new ReplicationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ReplicationListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ReplicationListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ReplicationListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetReplications

`func (o *ReplicationListResponse) GetReplications() []Replication`

GetReplications returns the Replications field if non-nil, zero value otherwise.

### GetReplicationsOk

`func (o *ReplicationListResponse) GetReplicationsOk() (*[]Replication, bool)`

GetReplicationsOk returns a tuple with the Replications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplications

`func (o *ReplicationListResponse) SetReplications(v []Replication)`

SetReplications sets Replications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


