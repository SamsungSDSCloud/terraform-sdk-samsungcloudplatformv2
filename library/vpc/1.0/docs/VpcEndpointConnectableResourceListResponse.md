# VpcEndpointConnectableResourceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**VpcEndpointConnectableResources** | [**[]VpcEndpointConnectableResourceDetail**](VpcEndpointConnectableResourceDetail.md) |  | 

## Methods

### NewVpcEndpointConnectableResourceListResponse

`func NewVpcEndpointConnectableResourceListResponse(vpcEndpointConnectableResources []VpcEndpointConnectableResourceDetail, ) *VpcEndpointConnectableResourceListResponse`

NewVpcEndpointConnectableResourceListResponse instantiates a new VpcEndpointConnectableResourceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcEndpointConnectableResourceListResponseWithDefaults

`func NewVpcEndpointConnectableResourceListResponseWithDefaults() *VpcEndpointConnectableResourceListResponse`

NewVpcEndpointConnectableResourceListResponseWithDefaults instantiates a new VpcEndpointConnectableResourceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VpcEndpointConnectableResourceListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VpcEndpointConnectableResourceListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VpcEndpointConnectableResourceListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *VpcEndpointConnectableResourceListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *VpcEndpointConnectableResourceListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *VpcEndpointConnectableResourceListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *VpcEndpointConnectableResourceListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VpcEndpointConnectableResourceListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VpcEndpointConnectableResourceListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VpcEndpointConnectableResourceListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *VpcEndpointConnectableResourceListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VpcEndpointConnectableResourceListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetVpcEndpointConnectableResources

`func (o *VpcEndpointConnectableResourceListResponse) GetVpcEndpointConnectableResources() []VpcEndpointConnectableResourceDetail`

GetVpcEndpointConnectableResources returns the VpcEndpointConnectableResources field if non-nil, zero value otherwise.

### GetVpcEndpointConnectableResourcesOk

`func (o *VpcEndpointConnectableResourceListResponse) GetVpcEndpointConnectableResourcesOk() (*[]VpcEndpointConnectableResourceDetail, bool)`

GetVpcEndpointConnectableResourcesOk returns a tuple with the VpcEndpointConnectableResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcEndpointConnectableResources

`func (o *VpcEndpointConnectableResourceListResponse) SetVpcEndpointConnectableResources(v []VpcEndpointConnectableResourceDetail)`

SetVpcEndpointConnectableResources sets VpcEndpointConnectableResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


