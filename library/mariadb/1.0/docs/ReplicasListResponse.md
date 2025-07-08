# ReplicasListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]ReplicaItem**](ReplicaItem.md) | Replicas list | 
**Count** | **int32** | Replica count | 

## Methods

### NewReplicasListResponse

`func NewReplicasListResponse(contents []ReplicaItem, count int32, ) *ReplicasListResponse`

NewReplicasListResponse instantiates a new ReplicasListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicasListResponseWithDefaults

`func NewReplicasListResponseWithDefaults() *ReplicasListResponse`

NewReplicasListResponseWithDefaults instantiates a new ReplicasListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *ReplicasListResponse) GetContents() []ReplicaItem`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ReplicasListResponse) GetContentsOk() (*[]ReplicaItem, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ReplicasListResponse) SetContents(v []ReplicaItem)`

SetContents sets Contents field to given value.


### GetCount

`func (o *ReplicasListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ReplicasListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ReplicasListResponse) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


