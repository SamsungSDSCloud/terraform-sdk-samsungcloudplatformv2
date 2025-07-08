# ParameterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Parameter group ID | 
**MajorVersion** | **string** | Software major version | 
**Name** | **string** | Parameter group name | 
**ProductType** | **string** | Product type | 

## Methods

### NewParameterGroup

`func NewParameterGroup(id string, majorVersion string, name string, productType string, ) *ParameterGroup`

NewParameterGroup instantiates a new ParameterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterGroupWithDefaults

`func NewParameterGroupWithDefaults() *ParameterGroup`

NewParameterGroupWithDefaults instantiates a new ParameterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParameterGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterGroup) SetId(v string)`

SetId sets Id field to given value.


### GetMajorVersion

`func (o *ParameterGroup) GetMajorVersion() string`

GetMajorVersion returns the MajorVersion field if non-nil, zero value otherwise.

### GetMajorVersionOk

`func (o *ParameterGroup) GetMajorVersionOk() (*string, bool)`

GetMajorVersionOk returns a tuple with the MajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorVersion

`func (o *ParameterGroup) SetMajorVersion(v string)`

SetMajorVersion sets MajorVersion field to given value.


### GetName

`func (o *ParameterGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterGroup) SetName(v string)`

SetName sets Name field to given value.


### GetProductType

`func (o *ParameterGroup) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ParameterGroup) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ParameterGroup) SetProductType(v string)`

SetProductType sets ProductType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


