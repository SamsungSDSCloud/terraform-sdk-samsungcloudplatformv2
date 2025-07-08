# PlannedComputeCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account Id | 
**ContractType** | Pointer to [**PlannedComputeContractEnum**](PlannedComputeContractEnum.md) | Contract Type | [optional] 
**OsType** | Pointer to [**PlannedComputeOSTypeEnum**](PlannedComputeOSTypeEnum.md) | OS Type | [optional] 
**ServerType** | **string** | Server Type | 
**ServiceId** | **string** | Service Id | 
**ServiceName** | Pointer to **string** | Service Name | [optional] 
**Tag** | Pointer to [**[]TagDTO**](TagDTO.md) | Tag Form | [optional] 

## Methods

### NewPlannedComputeCreateRequest

`func NewPlannedComputeCreateRequest(accountId string, serverType string, serviceId string, ) *PlannedComputeCreateRequest`

NewPlannedComputeCreateRequest instantiates a new PlannedComputeCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannedComputeCreateRequestWithDefaults

`func NewPlannedComputeCreateRequestWithDefaults() *PlannedComputeCreateRequest`

NewPlannedComputeCreateRequestWithDefaults instantiates a new PlannedComputeCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PlannedComputeCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PlannedComputeCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PlannedComputeCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetContractType

`func (o *PlannedComputeCreateRequest) GetContractType() PlannedComputeContractEnum`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *PlannedComputeCreateRequest) GetContractTypeOk() (*PlannedComputeContractEnum, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *PlannedComputeCreateRequest) SetContractType(v PlannedComputeContractEnum)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *PlannedComputeCreateRequest) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetOsType

`func (o *PlannedComputeCreateRequest) GetOsType() PlannedComputeOSTypeEnum`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *PlannedComputeCreateRequest) GetOsTypeOk() (*PlannedComputeOSTypeEnum, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *PlannedComputeCreateRequest) SetOsType(v PlannedComputeOSTypeEnum)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *PlannedComputeCreateRequest) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetServerType

`func (o *PlannedComputeCreateRequest) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *PlannedComputeCreateRequest) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *PlannedComputeCreateRequest) SetServerType(v string)`

SetServerType sets ServerType field to given value.


### GetServiceId

`func (o *PlannedComputeCreateRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PlannedComputeCreateRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PlannedComputeCreateRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceName

`func (o *PlannedComputeCreateRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *PlannedComputeCreateRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *PlannedComputeCreateRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *PlannedComputeCreateRequest) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetTag

`func (o *PlannedComputeCreateRequest) GetTag() []TagDTO`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PlannedComputeCreateRequest) GetTagOk() (*[]TagDTO, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PlannedComputeCreateRequest) SetTag(v []TagDTO)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PlannedComputeCreateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


