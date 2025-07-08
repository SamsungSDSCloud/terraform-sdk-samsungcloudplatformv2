# NodepoolListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**Nodepools** | Pointer to [**[]NodepoolSummary**](NodepoolSummary.md) | Nodepools List | [optional] 

## Methods

### NewNodepoolListResponse

`func NewNodepoolListResponse() *NodepoolListResponse`

NewNodepoolListResponse instantiates a new NodepoolListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodepoolListResponseWithDefaults

`func NewNodepoolListResponseWithDefaults() *NodepoolListResponse`

NewNodepoolListResponseWithDefaults instantiates a new NodepoolListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *NodepoolListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NodepoolListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NodepoolListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NodepoolListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *NodepoolListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *NodepoolListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *NodepoolListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NodepoolListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NodepoolListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NodepoolListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *NodepoolListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *NodepoolListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetNodepools

`func (o *NodepoolListResponse) GetNodepools() []NodepoolSummary`

GetNodepools returns the Nodepools field if non-nil, zero value otherwise.

### GetNodepoolsOk

`func (o *NodepoolListResponse) GetNodepoolsOk() (*[]NodepoolSummary, bool)`

GetNodepoolsOk returns a tuple with the Nodepools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodepools

`func (o *NodepoolListResponse) SetNodepools(v []NodepoolSummary)`

SetNodepools sets Nodepools field to given value.

### HasNodepools

`func (o *NodepoolListResponse) HasNodepools() bool`

HasNodepools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


