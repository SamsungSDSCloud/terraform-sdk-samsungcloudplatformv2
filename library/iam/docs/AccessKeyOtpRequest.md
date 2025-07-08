# AccessKeyOtpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to [**LocaleEnum**](LocaleEnum.md) | Locale | [optional] 
**Method** | Pointer to [**OtpMethodEnum**](OtpMethodEnum.md) | OTP method (EMAIL or PHONE | [optional] 

## Methods

### NewAccessKeyOtpRequest

`func NewAccessKeyOtpRequest() *AccessKeyOtpRequest`

NewAccessKeyOtpRequest instantiates a new AccessKeyOtpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyOtpRequestWithDefaults

`func NewAccessKeyOtpRequestWithDefaults() *AccessKeyOtpRequest`

NewAccessKeyOtpRequestWithDefaults instantiates a new AccessKeyOtpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *AccessKeyOtpRequest) GetLocale() LocaleEnum`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *AccessKeyOtpRequest) GetLocaleOk() (*LocaleEnum, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *AccessKeyOtpRequest) SetLocale(v LocaleEnum)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *AccessKeyOtpRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetMethod

`func (o *AccessKeyOtpRequest) GetMethod() OtpMethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AccessKeyOtpRequest) GetMethodOk() (*OtpMethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AccessKeyOtpRequest) SetMethod(v OtpMethodEnum)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AccessKeyOtpRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


