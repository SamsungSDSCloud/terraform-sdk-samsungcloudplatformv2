# PlannedComputeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Account Id | [optional] 
**ContractId** | Pointer to **NullableString** |  | [optional] 
**ContractType** | Pointer to **string** | Contract Type | [optional] [default to "01"]
**CreatedAt** | Pointer to **time.Time** | Created date | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**DeleteYn** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **string** | End date (YYYY-MM-dd) | [optional] 
**FirstContractStartAt** | Pointer to **string** | Start date (YYYY-MM-dd) | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | Modified date | [optional] 
**ModifiedBy** | Pointer to **NullableString** |  | [optional] 
**NextContractType** | Pointer to **NullableString** |  | [optional] 
**NextEndDate** | Pointer to **NullableString** |  | [optional] 
**NextStartDate** | Pointer to **NullableString** |  | [optional] 
**OsName** | Pointer to **string** | OS Name | [optional] 
**OsType** | Pointer to **string** | OS Type | [optional] 
**Region** | Pointer to **string** | Region | [optional] 
**ResourceName** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to **string** | Resource Type | [optional] 
**ServerType** | Pointer to **string** | Server Type | [optional] 
**ServerTypeDescription** | Pointer to **map[string]interface{}** | Server Type Description | [optional] 
**ServiceId** | Pointer to **string** | Service Id | [optional] 
**ServiceName** | Pointer to **string** | Service Name | [optional] 
**Srn** | Pointer to **string** | SRN | [optional] 
**StartDate** | Pointer to **string** | Start date (YYYY-MM-dd) | [optional] 
**State** | Pointer to **string** | Planned Compute State | [optional] [default to "ACTIVE"]

## Methods

### NewPlannedComputeDTO

`func NewPlannedComputeDTO() *PlannedComputeDTO`

NewPlannedComputeDTO instantiates a new PlannedComputeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannedComputeDTOWithDefaults

`func NewPlannedComputeDTOWithDefaults() *PlannedComputeDTO`

NewPlannedComputeDTOWithDefaults instantiates a new PlannedComputeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PlannedComputeDTO) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PlannedComputeDTO) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PlannedComputeDTO) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PlannedComputeDTO) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetContractId

`func (o *PlannedComputeDTO) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *PlannedComputeDTO) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *PlannedComputeDTO) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *PlannedComputeDTO) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### SetContractIdNil

`func (o *PlannedComputeDTO) SetContractIdNil(b bool)`

 SetContractIdNil sets the value for ContractId to be an explicit nil

### UnsetContractId
`func (o *PlannedComputeDTO) UnsetContractId()`

UnsetContractId ensures that no value is present for ContractId, not even an explicit nil
### GetContractType

`func (o *PlannedComputeDTO) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *PlannedComputeDTO) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *PlannedComputeDTO) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *PlannedComputeDTO) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlannedComputeDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlannedComputeDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlannedComputeDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlannedComputeDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PlannedComputeDTO) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PlannedComputeDTO) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PlannedComputeDTO) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PlannedComputeDTO) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *PlannedComputeDTO) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *PlannedComputeDTO) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDeleteYn

`func (o *PlannedComputeDTO) GetDeleteYn() string`

GetDeleteYn returns the DeleteYn field if non-nil, zero value otherwise.

### GetDeleteYnOk

`func (o *PlannedComputeDTO) GetDeleteYnOk() (*string, bool)`

GetDeleteYnOk returns a tuple with the DeleteYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteYn

`func (o *PlannedComputeDTO) SetDeleteYn(v string)`

SetDeleteYn sets DeleteYn field to given value.

### HasDeleteYn

`func (o *PlannedComputeDTO) HasDeleteYn() bool`

HasDeleteYn returns a boolean if a field has been set.

### SetDeleteYnNil

`func (o *PlannedComputeDTO) SetDeleteYnNil(b bool)`

 SetDeleteYnNil sets the value for DeleteYn to be an explicit nil

### UnsetDeleteYn
`func (o *PlannedComputeDTO) UnsetDeleteYn()`

UnsetDeleteYn ensures that no value is present for DeleteYn, not even an explicit nil
### GetEndDate

`func (o *PlannedComputeDTO) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PlannedComputeDTO) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PlannedComputeDTO) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PlannedComputeDTO) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFirstContractStartAt

`func (o *PlannedComputeDTO) GetFirstContractStartAt() string`

GetFirstContractStartAt returns the FirstContractStartAt field if non-nil, zero value otherwise.

### GetFirstContractStartAtOk

`func (o *PlannedComputeDTO) GetFirstContractStartAtOk() (*string, bool)`

GetFirstContractStartAtOk returns a tuple with the FirstContractStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstContractStartAt

`func (o *PlannedComputeDTO) SetFirstContractStartAt(v string)`

SetFirstContractStartAt sets FirstContractStartAt field to given value.

### HasFirstContractStartAt

`func (o *PlannedComputeDTO) HasFirstContractStartAt() bool`

HasFirstContractStartAt returns a boolean if a field has been set.

### GetId

`func (o *PlannedComputeDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlannedComputeDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlannedComputeDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlannedComputeDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PlannedComputeDTO) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PlannedComputeDTO) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetModifiedAt

`func (o *PlannedComputeDTO) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PlannedComputeDTO) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PlannedComputeDTO) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PlannedComputeDTO) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *PlannedComputeDTO) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PlannedComputeDTO) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PlannedComputeDTO) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *PlannedComputeDTO) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedByNil

`func (o *PlannedComputeDTO) SetModifiedByNil(b bool)`

 SetModifiedByNil sets the value for ModifiedBy to be an explicit nil

### UnsetModifiedBy
`func (o *PlannedComputeDTO) UnsetModifiedBy()`

UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
### GetNextContractType

`func (o *PlannedComputeDTO) GetNextContractType() string`

GetNextContractType returns the NextContractType field if non-nil, zero value otherwise.

### GetNextContractTypeOk

`func (o *PlannedComputeDTO) GetNextContractTypeOk() (*string, bool)`

GetNextContractTypeOk returns a tuple with the NextContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextContractType

`func (o *PlannedComputeDTO) SetNextContractType(v string)`

SetNextContractType sets NextContractType field to given value.

### HasNextContractType

`func (o *PlannedComputeDTO) HasNextContractType() bool`

HasNextContractType returns a boolean if a field has been set.

### SetNextContractTypeNil

`func (o *PlannedComputeDTO) SetNextContractTypeNil(b bool)`

 SetNextContractTypeNil sets the value for NextContractType to be an explicit nil

### UnsetNextContractType
`func (o *PlannedComputeDTO) UnsetNextContractType()`

UnsetNextContractType ensures that no value is present for NextContractType, not even an explicit nil
### GetNextEndDate

`func (o *PlannedComputeDTO) GetNextEndDate() string`

GetNextEndDate returns the NextEndDate field if non-nil, zero value otherwise.

### GetNextEndDateOk

`func (o *PlannedComputeDTO) GetNextEndDateOk() (*string, bool)`

GetNextEndDateOk returns a tuple with the NextEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEndDate

`func (o *PlannedComputeDTO) SetNextEndDate(v string)`

SetNextEndDate sets NextEndDate field to given value.

### HasNextEndDate

`func (o *PlannedComputeDTO) HasNextEndDate() bool`

HasNextEndDate returns a boolean if a field has been set.

### SetNextEndDateNil

`func (o *PlannedComputeDTO) SetNextEndDateNil(b bool)`

 SetNextEndDateNil sets the value for NextEndDate to be an explicit nil

### UnsetNextEndDate
`func (o *PlannedComputeDTO) UnsetNextEndDate()`

UnsetNextEndDate ensures that no value is present for NextEndDate, not even an explicit nil
### GetNextStartDate

`func (o *PlannedComputeDTO) GetNextStartDate() string`

GetNextStartDate returns the NextStartDate field if non-nil, zero value otherwise.

### GetNextStartDateOk

`func (o *PlannedComputeDTO) GetNextStartDateOk() (*string, bool)`

GetNextStartDateOk returns a tuple with the NextStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStartDate

`func (o *PlannedComputeDTO) SetNextStartDate(v string)`

SetNextStartDate sets NextStartDate field to given value.

### HasNextStartDate

`func (o *PlannedComputeDTO) HasNextStartDate() bool`

HasNextStartDate returns a boolean if a field has been set.

### SetNextStartDateNil

`func (o *PlannedComputeDTO) SetNextStartDateNil(b bool)`

 SetNextStartDateNil sets the value for NextStartDate to be an explicit nil

### UnsetNextStartDate
`func (o *PlannedComputeDTO) UnsetNextStartDate()`

UnsetNextStartDate ensures that no value is present for NextStartDate, not even an explicit nil
### GetOsName

`func (o *PlannedComputeDTO) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *PlannedComputeDTO) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *PlannedComputeDTO) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *PlannedComputeDTO) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsType

`func (o *PlannedComputeDTO) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *PlannedComputeDTO) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *PlannedComputeDTO) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *PlannedComputeDTO) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetRegion

`func (o *PlannedComputeDTO) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PlannedComputeDTO) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PlannedComputeDTO) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PlannedComputeDTO) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResourceName

`func (o *PlannedComputeDTO) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PlannedComputeDTO) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PlannedComputeDTO) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *PlannedComputeDTO) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *PlannedComputeDTO) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *PlannedComputeDTO) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceType

`func (o *PlannedComputeDTO) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *PlannedComputeDTO) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *PlannedComputeDTO) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *PlannedComputeDTO) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetServerType

`func (o *PlannedComputeDTO) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *PlannedComputeDTO) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *PlannedComputeDTO) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *PlannedComputeDTO) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerTypeDescription

`func (o *PlannedComputeDTO) GetServerTypeDescription() map[string]interface{}`

GetServerTypeDescription returns the ServerTypeDescription field if non-nil, zero value otherwise.

### GetServerTypeDescriptionOk

`func (o *PlannedComputeDTO) GetServerTypeDescriptionOk() (*map[string]interface{}, bool)`

GetServerTypeDescriptionOk returns a tuple with the ServerTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeDescription

`func (o *PlannedComputeDTO) SetServerTypeDescription(v map[string]interface{})`

SetServerTypeDescription sets ServerTypeDescription field to given value.

### HasServerTypeDescription

`func (o *PlannedComputeDTO) HasServerTypeDescription() bool`

HasServerTypeDescription returns a boolean if a field has been set.

### GetServiceId

`func (o *PlannedComputeDTO) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PlannedComputeDTO) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PlannedComputeDTO) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *PlannedComputeDTO) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *PlannedComputeDTO) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *PlannedComputeDTO) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *PlannedComputeDTO) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *PlannedComputeDTO) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSrn

`func (o *PlannedComputeDTO) GetSrn() string`

GetSrn returns the Srn field if non-nil, zero value otherwise.

### GetSrnOk

`func (o *PlannedComputeDTO) GetSrnOk() (*string, bool)`

GetSrnOk returns a tuple with the Srn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrn

`func (o *PlannedComputeDTO) SetSrn(v string)`

SetSrn sets Srn field to given value.

### HasSrn

`func (o *PlannedComputeDTO) HasSrn() bool`

HasSrn returns a boolean if a field has been set.

### GetStartDate

`func (o *PlannedComputeDTO) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PlannedComputeDTO) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PlannedComputeDTO) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PlannedComputeDTO) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetState

`func (o *PlannedComputeDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PlannedComputeDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PlannedComputeDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PlannedComputeDTO) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


