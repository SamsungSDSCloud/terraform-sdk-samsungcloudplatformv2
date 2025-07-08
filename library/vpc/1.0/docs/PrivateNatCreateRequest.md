# PrivateNatCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **string** | Private NAT IP range | 
**Description** | Pointer to **string** | Description | [optional] [default to ""]
**DirectConnectId** | **string** | Direct Connect ID | 
**Name** | **string** | Private NAT Name | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]

## Methods

### NewPrivateNatCreateRequest

`func NewPrivateNatCreateRequest(cidr string, directConnectId string, name string, ) *PrivateNatCreateRequest`

NewPrivateNatCreateRequest instantiates a new PrivateNatCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNatCreateRequestWithDefaults

`func NewPrivateNatCreateRequestWithDefaults() *PrivateNatCreateRequest`

NewPrivateNatCreateRequestWithDefaults instantiates a new PrivateNatCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *PrivateNatCreateRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *PrivateNatCreateRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *PrivateNatCreateRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetDescription

`func (o *PrivateNatCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateNatCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateNatCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivateNatCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirectConnectId

`func (o *PrivateNatCreateRequest) GetDirectConnectId() string`

GetDirectConnectId returns the DirectConnectId field if non-nil, zero value otherwise.

### GetDirectConnectIdOk

`func (o *PrivateNatCreateRequest) GetDirectConnectIdOk() (*string, bool)`

GetDirectConnectIdOk returns a tuple with the DirectConnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectConnectId

`func (o *PrivateNatCreateRequest) SetDirectConnectId(v string)`

SetDirectConnectId sets DirectConnectId field to given value.


### GetName

`func (o *PrivateNatCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateNatCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateNatCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *PrivateNatCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PrivateNatCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PrivateNatCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PrivateNatCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


