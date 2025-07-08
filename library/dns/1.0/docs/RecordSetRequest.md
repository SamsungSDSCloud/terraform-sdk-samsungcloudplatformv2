# RecordSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to **[]interface{}** |  | [optional] 
**Ttl** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewRecordSetRequest

`func NewRecordSetRequest() *RecordSetRequest`

NewRecordSetRequest instantiates a new RecordSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordSetRequestWithDefaults

`func NewRecordSetRequestWithDefaults() *RecordSetRequest`

NewRecordSetRequestWithDefaults instantiates a new RecordSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *RecordSetRequest) GetRecords() []interface{}`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *RecordSetRequest) GetRecordsOk() (*[]interface{}, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *RecordSetRequest) SetRecords(v []interface{})`

SetRecords sets Records field to given value.

### HasRecords

`func (o *RecordSetRequest) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### SetRecordsNil

`func (o *RecordSetRequest) SetRecordsNil(b bool)`

 SetRecordsNil sets the value for Records to be an explicit nil

### UnsetRecords
`func (o *RecordSetRequest) UnsetRecords()`

UnsetRecords ensures that no value is present for Records, not even an explicit nil
### GetTtl

`func (o *RecordSetRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordSetRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordSetRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordSetRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *RecordSetRequest) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *RecordSetRequest) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


