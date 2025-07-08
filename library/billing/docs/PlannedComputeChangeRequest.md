# PlannedComputeChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**PlannedComputeChangeActionEnum**](PlannedComputeChangeActionEnum.md) | Planned Compute change request action type | [optional] 
**ContractType** | Pointer to [**NullablePlannedComputeContractEnum**](PlannedComputeContractEnum.md) |  | [optional] 
**ServerType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPlannedComputeChangeRequest

`func NewPlannedComputeChangeRequest() *PlannedComputeChangeRequest`

NewPlannedComputeChangeRequest instantiates a new PlannedComputeChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannedComputeChangeRequestWithDefaults

`func NewPlannedComputeChangeRequestWithDefaults() *PlannedComputeChangeRequest`

NewPlannedComputeChangeRequestWithDefaults instantiates a new PlannedComputeChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PlannedComputeChangeRequest) GetAction() PlannedComputeChangeActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PlannedComputeChangeRequest) GetActionOk() (*PlannedComputeChangeActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PlannedComputeChangeRequest) SetAction(v PlannedComputeChangeActionEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *PlannedComputeChangeRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetContractType

`func (o *PlannedComputeChangeRequest) GetContractType() PlannedComputeContractEnum`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *PlannedComputeChangeRequest) GetContractTypeOk() (*PlannedComputeContractEnum, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *PlannedComputeChangeRequest) SetContractType(v PlannedComputeContractEnum)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *PlannedComputeChangeRequest) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### SetContractTypeNil

`func (o *PlannedComputeChangeRequest) SetContractTypeNil(b bool)`

 SetContractTypeNil sets the value for ContractType to be an explicit nil

### UnsetContractType
`func (o *PlannedComputeChangeRequest) UnsetContractType()`

UnsetContractType ensures that no value is present for ContractType, not even an explicit nil
### GetServerType

`func (o *PlannedComputeChangeRequest) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *PlannedComputeChangeRequest) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *PlannedComputeChangeRequest) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *PlannedComputeChangeRequest) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### SetServerTypeNil

`func (o *PlannedComputeChangeRequest) SetServerTypeNil(b bool)`

 SetServerTypeNil sets the value for ServerType to be an explicit nil

### UnsetServerType
`func (o *PlannedComputeChangeRequest) UnsetServerType()`

UnsetServerType ensures that no value is present for ServerType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


