# DefaultConfigListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configs** | **map[string]interface{}** | Configs | 
**Count** | **int32** | Count | 

## Methods

### NewDefaultConfigListResponse

`func NewDefaultConfigListResponse(configs map[string]interface{}, count int32, ) *DefaultConfigListResponse`

NewDefaultConfigListResponse instantiates a new DefaultConfigListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultConfigListResponseWithDefaults

`func NewDefaultConfigListResponseWithDefaults() *DefaultConfigListResponse`

NewDefaultConfigListResponseWithDefaults instantiates a new DefaultConfigListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigs

`func (o *DefaultConfigListResponse) GetConfigs() map[string]interface{}`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *DefaultConfigListResponse) GetConfigsOk() (*map[string]interface{}, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *DefaultConfigListResponse) SetConfigs(v map[string]interface{})`

SetConfigs sets Configs field to given value.


### GetCount

`func (o *DefaultConfigListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DefaultConfigListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DefaultConfigListResponse) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


