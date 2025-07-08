# ConfigMapListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMaps** | [**[]ConfigMapSummary**](ConfigMapSummary.md) |  | 
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewConfigMapListResponse

`func NewConfigMapListResponse(configMaps []ConfigMapSummary, ) *ConfigMapListResponse`

NewConfigMapListResponse instantiates a new ConfigMapListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigMapListResponseWithDefaults

`func NewConfigMapListResponseWithDefaults() *ConfigMapListResponse`

NewConfigMapListResponseWithDefaults instantiates a new ConfigMapListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMaps

`func (o *ConfigMapListResponse) GetConfigMaps() []ConfigMapSummary`

GetConfigMaps returns the ConfigMaps field if non-nil, zero value otherwise.

### GetConfigMapsOk

`func (o *ConfigMapListResponse) GetConfigMapsOk() (*[]ConfigMapSummary, bool)`

GetConfigMapsOk returns a tuple with the ConfigMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMaps

`func (o *ConfigMapListResponse) SetConfigMaps(v []ConfigMapSummary)`

SetConfigMaps sets ConfigMaps field to given value.


### GetCount

`func (o *ConfigMapListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ConfigMapListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ConfigMapListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ConfigMapListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ConfigMapListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ConfigMapListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *ConfigMapListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConfigMapListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConfigMapListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConfigMapListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ConfigMapListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ConfigMapListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


