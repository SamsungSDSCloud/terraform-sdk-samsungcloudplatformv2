# AccessControlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseUserName** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAccessControlRequest

`func NewAccessControlRequest() *AccessControlRequest`

NewAccessControlRequest instantiates a new AccessControlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlRequestWithDefaults

`func NewAccessControlRequestWithDefaults() *AccessControlRequest`

NewAccessControlRequestWithDefaults instantiates a new AccessControlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseUserName

`func (o *AccessControlRequest) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *AccessControlRequest) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *AccessControlRequest) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.

### HasDatabaseUserName

`func (o *AccessControlRequest) HasDatabaseUserName() bool`

HasDatabaseUserName returns a boolean if a field has been set.

### SetDatabaseUserNameNil

`func (o *AccessControlRequest) SetDatabaseUserNameNil(b bool)`

 SetDatabaseUserNameNil sets the value for DatabaseUserName to be an explicit nil

### UnsetDatabaseUserName
`func (o *AccessControlRequest) UnsetDatabaseUserName()`

UnsetDatabaseUserName ensures that no value is present for DatabaseUserName, not even an explicit nil
### GetIpAddress

`func (o *AccessControlRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AccessControlRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AccessControlRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *AccessControlRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *AccessControlRequest) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *AccessControlRequest) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


