# ServerSecurityGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroups** | [**[]SecurityGroupResponse**](SecurityGroupResponse.md) | Security Groups | 

## Methods

### NewServerSecurityGroupListResponse

`func NewServerSecurityGroupListResponse(securityGroups []SecurityGroupResponse, ) *ServerSecurityGroupListResponse`

NewServerSecurityGroupListResponse instantiates a new ServerSecurityGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerSecurityGroupListResponseWithDefaults

`func NewServerSecurityGroupListResponseWithDefaults() *ServerSecurityGroupListResponse`

NewServerSecurityGroupListResponseWithDefaults instantiates a new ServerSecurityGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroups

`func (o *ServerSecurityGroupListResponse) GetSecurityGroups() []SecurityGroupResponse`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ServerSecurityGroupListResponse) GetSecurityGroupsOk() (*[]SecurityGroupResponse, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ServerSecurityGroupListResponse) SetSecurityGroups(v []SecurityGroupResponse)`

SetSecurityGroups sets SecurityGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


