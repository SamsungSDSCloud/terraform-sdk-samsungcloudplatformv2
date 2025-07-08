# SqlserverAddSecondaryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**License** | **string** | license | 
**Name** | **string** | Secondary Name | 
**ServiceIpAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSqlserverAddSecondaryRequest

`func NewSqlserverAddSecondaryRequest(license string, name string, ) *SqlserverAddSecondaryRequest`

NewSqlserverAddSecondaryRequest instantiates a new SqlserverAddSecondaryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlserverAddSecondaryRequestWithDefaults

`func NewSqlserverAddSecondaryRequestWithDefaults() *SqlserverAddSecondaryRequest`

NewSqlserverAddSecondaryRequestWithDefaults instantiates a new SqlserverAddSecondaryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicense

`func (o *SqlserverAddSecondaryRequest) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *SqlserverAddSecondaryRequest) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *SqlserverAddSecondaryRequest) SetLicense(v string)`

SetLicense sets License field to given value.


### GetName

`func (o *SqlserverAddSecondaryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlserverAddSecondaryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlserverAddSecondaryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetServiceIpAddress

`func (o *SqlserverAddSecondaryRequest) GetServiceIpAddress() string`

GetServiceIpAddress returns the ServiceIpAddress field if non-nil, zero value otherwise.

### GetServiceIpAddressOk

`func (o *SqlserverAddSecondaryRequest) GetServiceIpAddressOk() (*string, bool)`

GetServiceIpAddressOk returns a tuple with the ServiceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIpAddress

`func (o *SqlserverAddSecondaryRequest) SetServiceIpAddress(v string)`

SetServiceIpAddress sets ServiceIpAddress field to given value.

### HasServiceIpAddress

`func (o *SqlserverAddSecondaryRequest) HasServiceIpAddress() bool`

HasServiceIpAddress returns a boolean if a field has been set.

### SetServiceIpAddressNil

`func (o *SqlserverAddSecondaryRequest) SetServiceIpAddressNil(b bool)`

 SetServiceIpAddressNil sets the value for ServiceIpAddress to be an explicit nil

### UnsetServiceIpAddress
`func (o *SqlserverAddSecondaryRequest) UnsetServiceIpAddress()`

UnsetServiceIpAddress ensures that no value is present for ServiceIpAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


