# VpcPeering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**VpcPeeringAccountType**](VpcPeeringAccountType.md) | Account Type | 
**ApproverVpcAccountId** | **string** | Approver VPC Account ID | 
**ApproverVpcId** | **string** | Approver VPC ID | 
**ApproverVpcName** | **string** | Approver VPC Name | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | VPC Peering ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | VPC Peering Name | 
**RequesterVpcAccountId** | **string** | Requester VPC Account ID | 
**RequesterVpcId** | **string** | Requester VPC ID | 
**RequesterVpcName** | **string** | Requester VPC Name | 
**State** | [**VpcPeeringState**](VpcPeeringState.md) | State | 

## Methods

### NewVpcPeering

`func NewVpcPeering(accountType VpcPeeringAccountType, approverVpcAccountId string, approverVpcId string, approverVpcName string, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, name string, requesterVpcAccountId string, requesterVpcId string, requesterVpcName string, state VpcPeeringState, ) *VpcPeering`

NewVpcPeering instantiates a new VpcPeering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcPeeringWithDefaults

`func NewVpcPeeringWithDefaults() *VpcPeering`

NewVpcPeeringWithDefaults instantiates a new VpcPeering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *VpcPeering) GetAccountType() VpcPeeringAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *VpcPeering) GetAccountTypeOk() (*VpcPeeringAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *VpcPeering) SetAccountType(v VpcPeeringAccountType)`

SetAccountType sets AccountType field to given value.


### GetApproverVpcAccountId

`func (o *VpcPeering) GetApproverVpcAccountId() string`

GetApproverVpcAccountId returns the ApproverVpcAccountId field if non-nil, zero value otherwise.

### GetApproverVpcAccountIdOk

`func (o *VpcPeering) GetApproverVpcAccountIdOk() (*string, bool)`

GetApproverVpcAccountIdOk returns a tuple with the ApproverVpcAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverVpcAccountId

`func (o *VpcPeering) SetApproverVpcAccountId(v string)`

SetApproverVpcAccountId sets ApproverVpcAccountId field to given value.


### GetApproverVpcId

`func (o *VpcPeering) GetApproverVpcId() string`

GetApproverVpcId returns the ApproverVpcId field if non-nil, zero value otherwise.

### GetApproverVpcIdOk

`func (o *VpcPeering) GetApproverVpcIdOk() (*string, bool)`

GetApproverVpcIdOk returns a tuple with the ApproverVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverVpcId

`func (o *VpcPeering) SetApproverVpcId(v string)`

SetApproverVpcId sets ApproverVpcId field to given value.


### GetApproverVpcName

`func (o *VpcPeering) GetApproverVpcName() string`

GetApproverVpcName returns the ApproverVpcName field if non-nil, zero value otherwise.

### GetApproverVpcNameOk

`func (o *VpcPeering) GetApproverVpcNameOk() (*string, bool)`

GetApproverVpcNameOk returns a tuple with the ApproverVpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverVpcName

`func (o *VpcPeering) SetApproverVpcName(v string)`

SetApproverVpcName sets ApproverVpcName field to given value.


### GetCreatedAt

`func (o *VpcPeering) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VpcPeering) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VpcPeering) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *VpcPeering) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VpcPeering) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VpcPeering) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *VpcPeering) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpcPeering) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpcPeering) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpcPeering) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VpcPeering) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VpcPeering) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *VpcPeering) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpcPeering) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpcPeering) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *VpcPeering) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VpcPeering) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VpcPeering) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *VpcPeering) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VpcPeering) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VpcPeering) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *VpcPeering) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpcPeering) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpcPeering) SetName(v string)`

SetName sets Name field to given value.


### GetRequesterVpcAccountId

`func (o *VpcPeering) GetRequesterVpcAccountId() string`

GetRequesterVpcAccountId returns the RequesterVpcAccountId field if non-nil, zero value otherwise.

### GetRequesterVpcAccountIdOk

`func (o *VpcPeering) GetRequesterVpcAccountIdOk() (*string, bool)`

GetRequesterVpcAccountIdOk returns a tuple with the RequesterVpcAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterVpcAccountId

`func (o *VpcPeering) SetRequesterVpcAccountId(v string)`

SetRequesterVpcAccountId sets RequesterVpcAccountId field to given value.


### GetRequesterVpcId

`func (o *VpcPeering) GetRequesterVpcId() string`

GetRequesterVpcId returns the RequesterVpcId field if non-nil, zero value otherwise.

### GetRequesterVpcIdOk

`func (o *VpcPeering) GetRequesterVpcIdOk() (*string, bool)`

GetRequesterVpcIdOk returns a tuple with the RequesterVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterVpcId

`func (o *VpcPeering) SetRequesterVpcId(v string)`

SetRequesterVpcId sets RequesterVpcId field to given value.


### GetRequesterVpcName

`func (o *VpcPeering) GetRequesterVpcName() string`

GetRequesterVpcName returns the RequesterVpcName field if non-nil, zero value otherwise.

### GetRequesterVpcNameOk

`func (o *VpcPeering) GetRequesterVpcNameOk() (*string, bool)`

GetRequesterVpcNameOk returns a tuple with the RequesterVpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterVpcName

`func (o *VpcPeering) SetRequesterVpcName(v string)`

SetRequesterVpcName sets RequesterVpcName field to given value.


### GetState

`func (o *VpcPeering) GetState() VpcPeeringState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpcPeering) GetStateOk() (*VpcPeeringState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpcPeering) SetState(v VpcPeeringState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


