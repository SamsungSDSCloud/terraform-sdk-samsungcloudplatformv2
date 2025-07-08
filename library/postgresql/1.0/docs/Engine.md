# Engine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndOfService** | Pointer to **bool** | End of Service | [optional] [default to false]
**Id** | **string** | ID | 
**MajorVersion** | **string** | Software major version | 
**Name** | **string** | DBaaS engine version name | 
**OsType** | **string** | OS type | 
**OsVersion** | **string** | OS version | 
**ProductImageType** | Pointer to **NullableString** |  | [optional] 
**SoftwareVersion** | **string** | Software version | 

## Methods

### NewEngine

`func NewEngine(id string, majorVersion string, name string, osType string, osVersion string, softwareVersion string, ) *Engine`

NewEngine instantiates a new Engine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineWithDefaults

`func NewEngineWithDefaults() *Engine`

NewEngineWithDefaults instantiates a new Engine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndOfService

`func (o *Engine) GetEndOfService() bool`

GetEndOfService returns the EndOfService field if non-nil, zero value otherwise.

### GetEndOfServiceOk

`func (o *Engine) GetEndOfServiceOk() (*bool, bool)`

GetEndOfServiceOk returns a tuple with the EndOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfService

`func (o *Engine) SetEndOfService(v bool)`

SetEndOfService sets EndOfService field to given value.

### HasEndOfService

`func (o *Engine) HasEndOfService() bool`

HasEndOfService returns a boolean if a field has been set.

### GetId

`func (o *Engine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Engine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Engine) SetId(v string)`

SetId sets Id field to given value.


### GetMajorVersion

`func (o *Engine) GetMajorVersion() string`

GetMajorVersion returns the MajorVersion field if non-nil, zero value otherwise.

### GetMajorVersionOk

`func (o *Engine) GetMajorVersionOk() (*string, bool)`

GetMajorVersionOk returns a tuple with the MajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorVersion

`func (o *Engine) SetMajorVersion(v string)`

SetMajorVersion sets MajorVersion field to given value.


### GetName

`func (o *Engine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Engine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Engine) SetName(v string)`

SetName sets Name field to given value.


### GetOsType

`func (o *Engine) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *Engine) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *Engine) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetOsVersion

`func (o *Engine) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *Engine) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *Engine) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetProductImageType

`func (o *Engine) GetProductImageType() string`

GetProductImageType returns the ProductImageType field if non-nil, zero value otherwise.

### GetProductImageTypeOk

`func (o *Engine) GetProductImageTypeOk() (*string, bool)`

GetProductImageTypeOk returns a tuple with the ProductImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImageType

`func (o *Engine) SetProductImageType(v string)`

SetProductImageType sets ProductImageType field to given value.

### HasProductImageType

`func (o *Engine) HasProductImageType() bool`

HasProductImageType returns a boolean if a field has been set.

### SetProductImageTypeNil

`func (o *Engine) SetProductImageTypeNil(b bool)`

 SetProductImageTypeNil sets the value for ProductImageType to be an explicit nil

### UnsetProductImageType
`func (o *Engine) UnsetProductImageType()`

UnsetProductImageType ensures that no value is present for ProductImageType, not even an explicit nil
### GetSoftwareVersion

`func (o *Engine) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *Engine) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *Engine) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


