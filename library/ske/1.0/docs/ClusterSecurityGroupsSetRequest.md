# ClusterSecurityGroupsSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupIdList** | **[]string** | Security Group ID List | 

## Methods

### NewClusterSecurityGroupsSetRequest

`func NewClusterSecurityGroupsSetRequest(securityGroupIdList []string, ) *ClusterSecurityGroupsSetRequest`

NewClusterSecurityGroupsSetRequest instantiates a new ClusterSecurityGroupsSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSecurityGroupsSetRequestWithDefaults

`func NewClusterSecurityGroupsSetRequestWithDefaults() *ClusterSecurityGroupsSetRequest`

NewClusterSecurityGroupsSetRequestWithDefaults instantiates a new ClusterSecurityGroupsSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupIdList

`func (o *ClusterSecurityGroupsSetRequest) GetSecurityGroupIdList() []string`

GetSecurityGroupIdList returns the SecurityGroupIdList field if non-nil, zero value otherwise.

### GetSecurityGroupIdListOk

`func (o *ClusterSecurityGroupsSetRequest) GetSecurityGroupIdListOk() (*[]string, bool)`

GetSecurityGroupIdListOk returns a tuple with the SecurityGroupIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIdList

`func (o *ClusterSecurityGroupsSetRequest) SetSecurityGroupIdList(v []string)`

SetSecurityGroupIdList sets SecurityGroupIdList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


