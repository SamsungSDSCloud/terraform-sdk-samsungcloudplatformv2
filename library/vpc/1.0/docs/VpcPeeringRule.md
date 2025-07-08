# VpcPeeringRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**DestinationCidr** | **string** | Destination CIDR | 
**DestinationVpcId** | **string** | Destination VPC ID | 
**DestinationVpcName** | **string** | Destination VPC Name | 
**DestinationVpcType** | [**VpcPeeringRuleDestinationVpcType**](VpcPeeringRuleDestinationVpcType.md) | Destination VPC Type | 
**Id** | **string** | VPC Peering Rule ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**SourceVpcId** | **string** | Source VPC ID | 
**SourceVpcName** | **string** | Source VPC Name | 
**SourceVpcType** | [**VpcPeeringRuleDestinationVpcType**](VpcPeeringRuleDestinationVpcType.md) | Source VPC Type | 
**State** | [**RoutingRuleState**](RoutingRuleState.md) | State | 
**VpcPeeringId** | **string** | VPC Peering ID | 

## Methods

### NewVpcPeeringRule

`func NewVpcPeeringRule(createdAt time.Time, createdBy string, destinationCidr string, destinationVpcId string, destinationVpcName string, destinationVpcType VpcPeeringRuleDestinationVpcType, id string, modifiedAt time.Time, modifiedBy string, sourceVpcId string, sourceVpcName string, sourceVpcType VpcPeeringRuleDestinationVpcType, state RoutingRuleState, vpcPeeringId string, ) *VpcPeeringRule`

NewVpcPeeringRule instantiates a new VpcPeeringRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcPeeringRuleWithDefaults

`func NewVpcPeeringRuleWithDefaults() *VpcPeeringRule`

NewVpcPeeringRuleWithDefaults instantiates a new VpcPeeringRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *VpcPeeringRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VpcPeeringRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VpcPeeringRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *VpcPeeringRule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VpcPeeringRule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VpcPeeringRule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDestinationCidr

`func (o *VpcPeeringRule) GetDestinationCidr() string`

GetDestinationCidr returns the DestinationCidr field if non-nil, zero value otherwise.

### GetDestinationCidrOk

`func (o *VpcPeeringRule) GetDestinationCidrOk() (*string, bool)`

GetDestinationCidrOk returns a tuple with the DestinationCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidr

`func (o *VpcPeeringRule) SetDestinationCidr(v string)`

SetDestinationCidr sets DestinationCidr field to given value.


### GetDestinationVpcId

`func (o *VpcPeeringRule) GetDestinationVpcId() string`

GetDestinationVpcId returns the DestinationVpcId field if non-nil, zero value otherwise.

### GetDestinationVpcIdOk

`func (o *VpcPeeringRule) GetDestinationVpcIdOk() (*string, bool)`

GetDestinationVpcIdOk returns a tuple with the DestinationVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationVpcId

`func (o *VpcPeeringRule) SetDestinationVpcId(v string)`

SetDestinationVpcId sets DestinationVpcId field to given value.


### GetDestinationVpcName

`func (o *VpcPeeringRule) GetDestinationVpcName() string`

GetDestinationVpcName returns the DestinationVpcName field if non-nil, zero value otherwise.

### GetDestinationVpcNameOk

`func (o *VpcPeeringRule) GetDestinationVpcNameOk() (*string, bool)`

GetDestinationVpcNameOk returns a tuple with the DestinationVpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationVpcName

`func (o *VpcPeeringRule) SetDestinationVpcName(v string)`

SetDestinationVpcName sets DestinationVpcName field to given value.


### GetDestinationVpcType

`func (o *VpcPeeringRule) GetDestinationVpcType() VpcPeeringRuleDestinationVpcType`

GetDestinationVpcType returns the DestinationVpcType field if non-nil, zero value otherwise.

### GetDestinationVpcTypeOk

`func (o *VpcPeeringRule) GetDestinationVpcTypeOk() (*VpcPeeringRuleDestinationVpcType, bool)`

GetDestinationVpcTypeOk returns a tuple with the DestinationVpcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationVpcType

`func (o *VpcPeeringRule) SetDestinationVpcType(v VpcPeeringRuleDestinationVpcType)`

SetDestinationVpcType sets DestinationVpcType field to given value.


### GetId

`func (o *VpcPeeringRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpcPeeringRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpcPeeringRule) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *VpcPeeringRule) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VpcPeeringRule) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VpcPeeringRule) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *VpcPeeringRule) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VpcPeeringRule) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VpcPeeringRule) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetSourceVpcId

`func (o *VpcPeeringRule) GetSourceVpcId() string`

GetSourceVpcId returns the SourceVpcId field if non-nil, zero value otherwise.

### GetSourceVpcIdOk

`func (o *VpcPeeringRule) GetSourceVpcIdOk() (*string, bool)`

GetSourceVpcIdOk returns a tuple with the SourceVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVpcId

`func (o *VpcPeeringRule) SetSourceVpcId(v string)`

SetSourceVpcId sets SourceVpcId field to given value.


### GetSourceVpcName

`func (o *VpcPeeringRule) GetSourceVpcName() string`

GetSourceVpcName returns the SourceVpcName field if non-nil, zero value otherwise.

### GetSourceVpcNameOk

`func (o *VpcPeeringRule) GetSourceVpcNameOk() (*string, bool)`

GetSourceVpcNameOk returns a tuple with the SourceVpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVpcName

`func (o *VpcPeeringRule) SetSourceVpcName(v string)`

SetSourceVpcName sets SourceVpcName field to given value.


### GetSourceVpcType

`func (o *VpcPeeringRule) GetSourceVpcType() VpcPeeringRuleDestinationVpcType`

GetSourceVpcType returns the SourceVpcType field if non-nil, zero value otherwise.

### GetSourceVpcTypeOk

`func (o *VpcPeeringRule) GetSourceVpcTypeOk() (*VpcPeeringRuleDestinationVpcType, bool)`

GetSourceVpcTypeOk returns a tuple with the SourceVpcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVpcType

`func (o *VpcPeeringRule) SetSourceVpcType(v VpcPeeringRuleDestinationVpcType)`

SetSourceVpcType sets SourceVpcType field to given value.


### GetState

`func (o *VpcPeeringRule) GetState() RoutingRuleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpcPeeringRule) GetStateOk() (*RoutingRuleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpcPeeringRule) SetState(v RoutingRuleState)`

SetState sets State field to given value.


### GetVpcPeeringId

`func (o *VpcPeeringRule) GetVpcPeeringId() string`

GetVpcPeeringId returns the VpcPeeringId field if non-nil, zero value otherwise.

### GetVpcPeeringIdOk

`func (o *VpcPeeringRule) GetVpcPeeringIdOk() (*string, bool)`

GetVpcPeeringIdOk returns a tuple with the VpcPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcPeeringId

`func (o *VpcPeeringRule) SetVpcPeeringId(v string)`

SetVpcPeeringId sets VpcPeeringId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


