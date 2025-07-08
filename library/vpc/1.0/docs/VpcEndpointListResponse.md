# VpcEndpointListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**VpcEndpoints** | [**[]VpcEndpoint**](VpcEndpoint.md) |  | 

## Methods

### NewVpcEndpointListResponse

`func NewVpcEndpointListResponse(vpcEndpoints []VpcEndpoint, ) *VpcEndpointListResponse`

NewVpcEndpointListResponse instantiates a new VpcEndpointListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcEndpointListResponseWithDefaults

`func NewVpcEndpointListResponseWithDefaults() *VpcEndpointListResponse`

NewVpcEndpointListResponseWithDefaults instantiates a new VpcEndpointListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VpcEndpointListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VpcEndpointListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VpcEndpointListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *VpcEndpointListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *VpcEndpointListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *VpcEndpointListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *VpcEndpointListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VpcEndpointListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VpcEndpointListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VpcEndpointListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *VpcEndpointListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VpcEndpointListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetVpcEndpoints

`func (o *VpcEndpointListResponse) GetVpcEndpoints() []VpcEndpoint`

GetVpcEndpoints returns the VpcEndpoints field if non-nil, zero value otherwise.

### GetVpcEndpointsOk

`func (o *VpcEndpointListResponse) GetVpcEndpointsOk() (*[]VpcEndpoint, bool)`

GetVpcEndpointsOk returns a tuple with the VpcEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpoints

`func (o *VpcEndpointListResponse) SetVpcEndpoints(v []VpcEndpoint)`

SetVpcEndpoints sets VpcEndpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


