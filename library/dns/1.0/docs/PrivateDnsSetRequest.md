# PrivateDnsSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedVpcIds** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPrivateDnsSetRequest

`func NewPrivateDnsSetRequest() *PrivateDnsSetRequest`

NewPrivateDnsSetRequest instantiates a new PrivateDnsSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateDnsSetRequestWithDefaults

`func NewPrivateDnsSetRequestWithDefaults() *PrivateDnsSetRequest`

NewPrivateDnsSetRequestWithDefaults instantiates a new PrivateDnsSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedVpcIds

`func (o *PrivateDnsSetRequest) GetConnectedVpcIds() []string`

GetConnectedVpcIds returns the ConnectedVpcIds field if non-nil, zero value otherwise.

### GetConnectedVpcIdsOk

`func (o *PrivateDnsSetRequest) GetConnectedVpcIdsOk() (*[]string, bool)`

GetConnectedVpcIdsOk returns a tuple with the ConnectedVpcIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedVpcIds

`func (o *PrivateDnsSetRequest) SetConnectedVpcIds(v []string)`

SetConnectedVpcIds sets ConnectedVpcIds field to given value.

### HasConnectedVpcIds

`func (o *PrivateDnsSetRequest) HasConnectedVpcIds() bool`

HasConnectedVpcIds returns a boolean if a field has been set.

### SetConnectedVpcIdsNil

`func (o *PrivateDnsSetRequest) SetConnectedVpcIdsNil(b bool)`

 SetConnectedVpcIdsNil sets the value for ConnectedVpcIds to be an explicit nil

### UnsetConnectedVpcIds
`func (o *PrivateDnsSetRequest) UnsetConnectedVpcIds()`

UnsetConnectedVpcIds ensures that no value is present for ConnectedVpcIds, not even an explicit nil
### GetDescription

`func (o *PrivateDnsSetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateDnsSetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateDnsSetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivateDnsSetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PrivateDnsSetRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PrivateDnsSetRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


