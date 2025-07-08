# HostedZoneListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**HostedZones** | [**[]HostedZone**](HostedZone.md) | hosted zones list | 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewHostedZoneListResponse

`func NewHostedZoneListResponse(hostedZones []HostedZone, ) *HostedZoneListResponse`

NewHostedZoneListResponse instantiates a new HostedZoneListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedZoneListResponseWithDefaults

`func NewHostedZoneListResponseWithDefaults() *HostedZoneListResponse`

NewHostedZoneListResponseWithDefaults instantiates a new HostedZoneListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *HostedZoneListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HostedZoneListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HostedZoneListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HostedZoneListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *HostedZoneListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *HostedZoneListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetHostedZones

`func (o *HostedZoneListResponse) GetHostedZones() []HostedZone`

GetHostedZones returns the HostedZones field if non-nil, zero value otherwise.

### GetHostedZonesOk

`func (o *HostedZoneListResponse) GetHostedZonesOk() (*[]HostedZone, bool)`

GetHostedZonesOk returns a tuple with the HostedZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZones

`func (o *HostedZoneListResponse) SetHostedZones(v []HostedZone)`

SetHostedZones sets HostedZones field to given value.


### GetLinks

`func (o *HostedZoneListResponse) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *HostedZoneListResponse) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *HostedZoneListResponse) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *HostedZoneListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *HostedZoneListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *HostedZoneListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *HostedZoneListResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HostedZoneListResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HostedZoneListResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HostedZoneListResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *HostedZoneListResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *HostedZoneListResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


