# VpcPeeringCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproverVpcAccountId** | Pointer to **string** | Approver VPC Account ID | [optional] [default to ""]
**ApproverVpcId** | **string** | Approver VPC ID | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | VPC Peering Name | 
**RequesterVpcId** | **string** | Requester VPC ID | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]

## Methods

### NewVpcPeeringCreateRequest

`func NewVpcPeeringCreateRequest(approverVpcId string, name string, requesterVpcId string, ) *VpcPeeringCreateRequest`

NewVpcPeeringCreateRequest instantiates a new VpcPeeringCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcPeeringCreateRequestWithDefaults

`func NewVpcPeeringCreateRequestWithDefaults() *VpcPeeringCreateRequest`

NewVpcPeeringCreateRequestWithDefaults instantiates a new VpcPeeringCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproverVpcAccountId

`func (o *VpcPeeringCreateRequest) GetApproverVpcAccountId() string`

GetApproverVpcAccountId returns the ApproverVpcAccountId field if non-nil, zero value otherwise.

### GetApproverVpcAccountIdOk

`func (o *VpcPeeringCreateRequest) GetApproverVpcAccountIdOk() (*string, bool)`

GetApproverVpcAccountIdOk returns a tuple with the ApproverVpcAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverVpcAccountId

`func (o *VpcPeeringCreateRequest) SetApproverVpcAccountId(v string)`

SetApproverVpcAccountId sets ApproverVpcAccountId field to given value.

### HasApproverVpcAccountId

`func (o *VpcPeeringCreateRequest) HasApproverVpcAccountId() bool`

HasApproverVpcAccountId returns a boolean if a field has been set.

### GetApproverVpcId

`func (o *VpcPeeringCreateRequest) GetApproverVpcId() string`

GetApproverVpcId returns the ApproverVpcId field if non-nil, zero value otherwise.

### GetApproverVpcIdOk

`func (o *VpcPeeringCreateRequest) GetApproverVpcIdOk() (*string, bool)`

GetApproverVpcIdOk returns a tuple with the ApproverVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverVpcId

`func (o *VpcPeeringCreateRequest) SetApproverVpcId(v string)`

SetApproverVpcId sets ApproverVpcId field to given value.


### GetDescription

`func (o *VpcPeeringCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpcPeeringCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpcPeeringCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpcPeeringCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VpcPeeringCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VpcPeeringCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *VpcPeeringCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpcPeeringCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpcPeeringCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRequesterVpcId

`func (o *VpcPeeringCreateRequest) GetRequesterVpcId() string`

GetRequesterVpcId returns the RequesterVpcId field if non-nil, zero value otherwise.

### GetRequesterVpcIdOk

`func (o *VpcPeeringCreateRequest) GetRequesterVpcIdOk() (*string, bool)`

GetRequesterVpcIdOk returns a tuple with the RequesterVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterVpcId

`func (o *VpcPeeringCreateRequest) SetRequesterVpcId(v string)`

SetRequesterVpcId sets RequesterVpcId field to given value.


### GetTags

`func (o *VpcPeeringCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpcPeeringCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpcPeeringCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpcPeeringCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


