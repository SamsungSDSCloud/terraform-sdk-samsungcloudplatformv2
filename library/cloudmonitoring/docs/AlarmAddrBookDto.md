# AlarmAddrBookDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddrBookName** | **string** | 주소록 이름 | 
**AddrbookId** | **string** | 주소록 아이디 | 
**CreatedBy** | **string** | 생성자 아이디 | 
**CreatedByName** | **string** | 생성자 이름 | 
**CreatedDt** | **time.Time** | 생성 일시 | 
**MemberCount** | **int32** | 멤버 수 | 

## Methods

### NewAlarmAddrBookDto

`func NewAlarmAddrBookDto(addrBookName string, addrbookId string, createdBy string, createdByName string, createdDt time.Time, memberCount int32, ) *AlarmAddrBookDto`

NewAlarmAddrBookDto instantiates a new AlarmAddrBookDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmAddrBookDtoWithDefaults

`func NewAlarmAddrBookDtoWithDefaults() *AlarmAddrBookDto`

NewAlarmAddrBookDtoWithDefaults instantiates a new AlarmAddrBookDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddrBookName

`func (o *AlarmAddrBookDto) GetAddrBookName() string`

GetAddrBookName returns the AddrBookName field if non-nil, zero value otherwise.

### GetAddrBookNameOk

`func (o *AlarmAddrBookDto) GetAddrBookNameOk() (*string, bool)`

GetAddrBookNameOk returns a tuple with the AddrBookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrBookName

`func (o *AlarmAddrBookDto) SetAddrBookName(v string)`

SetAddrBookName sets AddrBookName field to given value.


### GetAddrbookId

`func (o *AlarmAddrBookDto) GetAddrbookId() string`

GetAddrbookId returns the AddrbookId field if non-nil, zero value otherwise.

### GetAddrbookIdOk

`func (o *AlarmAddrBookDto) GetAddrbookIdOk() (*string, bool)`

GetAddrbookIdOk returns a tuple with the AddrbookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrbookId

`func (o *AlarmAddrBookDto) SetAddrbookId(v string)`

SetAddrbookId sets AddrbookId field to given value.


### GetCreatedBy

`func (o *AlarmAddrBookDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AlarmAddrBookDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AlarmAddrBookDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedByName

`func (o *AlarmAddrBookDto) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *AlarmAddrBookDto) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *AlarmAddrBookDto) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.


### GetCreatedDt

`func (o *AlarmAddrBookDto) GetCreatedDt() time.Time`

GetCreatedDt returns the CreatedDt field if non-nil, zero value otherwise.

### GetCreatedDtOk

`func (o *AlarmAddrBookDto) GetCreatedDtOk() (*time.Time, bool)`

GetCreatedDtOk returns a tuple with the CreatedDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDt

`func (o *AlarmAddrBookDto) SetCreatedDt(v time.Time)`

SetCreatedDt sets CreatedDt field to given value.


### GetMemberCount

`func (o *AlarmAddrBookDto) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *AlarmAddrBookDto) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *AlarmAddrBookDto) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


