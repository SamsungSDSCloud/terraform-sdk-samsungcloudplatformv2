# ContractTypeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractTypes** | [**[]ContractTypeDTO**](ContractTypeDTO.md) |  | 
**ExtensionTypes** | [**[]ContractTypeDTO**](ContractTypeDTO.md) |  | 

## Methods

### NewContractTypeListResponse

`func NewContractTypeListResponse(contractTypes []ContractTypeDTO, extensionTypes []ContractTypeDTO, ) *ContractTypeListResponse`

NewContractTypeListResponse instantiates a new ContractTypeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractTypeListResponseWithDefaults

`func NewContractTypeListResponseWithDefaults() *ContractTypeListResponse`

NewContractTypeListResponseWithDefaults instantiates a new ContractTypeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractTypes

`func (o *ContractTypeListResponse) GetContractTypes() []ContractTypeDTO`

GetContractTypes returns the ContractTypes field if non-nil, zero value otherwise.

### GetContractTypesOk

`func (o *ContractTypeListResponse) GetContractTypesOk() (*[]ContractTypeDTO, bool)`

GetContractTypesOk returns a tuple with the ContractTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTypes

`func (o *ContractTypeListResponse) SetContractTypes(v []ContractTypeDTO)`

SetContractTypes sets ContractTypes field to given value.


### GetExtensionTypes

`func (o *ContractTypeListResponse) GetExtensionTypes() []ContractTypeDTO`

GetExtensionTypes returns the ExtensionTypes field if non-nil, zero value otherwise.

### GetExtensionTypesOk

`func (o *ContractTypeListResponse) GetExtensionTypesOk() (*[]ContractTypeDTO, bool)`

GetExtensionTypesOk returns a tuple with the ExtensionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionTypes

`func (o *ContractTypeListResponse) SetExtensionTypes(v []ContractTypeDTO)`

SetExtensionTypes sets ExtensionTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


