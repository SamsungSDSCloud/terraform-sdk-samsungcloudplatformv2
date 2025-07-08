# ClusterPrivateAccessControlSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateEndpointAccessControlResources** | [**[]PrivateEndpointAccessControlResource**](PrivateEndpointAccessControlResource.md) |  | 

## Methods

### NewClusterPrivateAccessControlSetRequest

`func NewClusterPrivateAccessControlSetRequest(privateEndpointAccessControlResources []PrivateEndpointAccessControlResource, ) *ClusterPrivateAccessControlSetRequest`

NewClusterPrivateAccessControlSetRequest instantiates a new ClusterPrivateAccessControlSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPrivateAccessControlSetRequestWithDefaults

`func NewClusterPrivateAccessControlSetRequestWithDefaults() *ClusterPrivateAccessControlSetRequest`

NewClusterPrivateAccessControlSetRequestWithDefaults instantiates a new ClusterPrivateAccessControlSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateEndpointAccessControlResources

`func (o *ClusterPrivateAccessControlSetRequest) GetPrivateEndpointAccessControlResources() []PrivateEndpointAccessControlResource`

GetPrivateEndpointAccessControlResources returns the PrivateEndpointAccessControlResources field if non-nil, zero value otherwise.

### GetPrivateEndpointAccessControlResourcesOk

`func (o *ClusterPrivateAccessControlSetRequest) GetPrivateEndpointAccessControlResourcesOk() (*[]PrivateEndpointAccessControlResource, bool)`

GetPrivateEndpointAccessControlResourcesOk returns a tuple with the PrivateEndpointAccessControlResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointAccessControlResources

`func (o *ClusterPrivateAccessControlSetRequest) SetPrivateEndpointAccessControlResources(v []PrivateEndpointAccessControlResource)`

SetPrivateEndpointAccessControlResources sets PrivateEndpointAccessControlResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


