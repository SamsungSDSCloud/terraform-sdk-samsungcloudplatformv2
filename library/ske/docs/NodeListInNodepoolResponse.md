# NodeListInNodepoolResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**Nodes** | Pointer to [**[]NodeInNodepool**](NodeInNodepool.md) |  | [optional] 

## Methods

### NewNodeListInNodepoolResponse

`func NewNodeListInNodepoolResponse() *NodeListInNodepoolResponse`

NewNodeListInNodepoolResponse instantiates a new NodeListInNodepoolResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeListInNodepoolResponseWithDefaults

`func NewNodeListInNodepoolResponseWithDefaults() *NodeListInNodepoolResponse`

NewNodeListInNodepoolResponseWithDefaults instantiates a new NodeListInNodepoolResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *NodeListInNodepoolResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NodeListInNodepoolResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NodeListInNodepoolResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NodeListInNodepoolResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *NodeListInNodepoolResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *NodeListInNodepoolResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *NodeListInNodepoolResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NodeListInNodepoolResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NodeListInNodepoolResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NodeListInNodepoolResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *NodeListInNodepoolResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *NodeListInNodepoolResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetNodes

`func (o *NodeListInNodepoolResponse) GetNodes() []NodeInNodepool`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *NodeListInNodepoolResponse) GetNodesOk() (*[]NodeInNodepool, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *NodeListInNodepoolResponse) SetNodes(v []NodeInNodepool)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *NodeListInNodepoolResponse) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *NodeListInNodepoolResponse) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *NodeListInNodepoolResponse) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


