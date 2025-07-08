# DeploymentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Deployments** | [**[]DeploymentSummary**](DeploymentSummary.md) |  | 
**Links** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewDeploymentListResponse

`func NewDeploymentListResponse(deployments []DeploymentSummary, ) *DeploymentListResponse`

NewDeploymentListResponse instantiates a new DeploymentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentListResponseWithDefaults

`func NewDeploymentListResponseWithDefaults() *DeploymentListResponse`

NewDeploymentListResponseWithDefaults instantiates a new DeploymentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DeploymentListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DeploymentListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DeploymentListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DeploymentListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *DeploymentListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *DeploymentListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetDeployments

`func (o *DeploymentListResponse) GetDeployments() []DeploymentSummary`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *DeploymentListResponse) GetDeploymentsOk() (*[]DeploymentSummary, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *DeploymentListResponse) SetDeployments(v []DeploymentSummary)`

SetDeployments sets Deployments field to given value.


### GetLinks

`func (o *DeploymentListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeploymentListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeploymentListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeploymentListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *DeploymentListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *DeploymentListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


