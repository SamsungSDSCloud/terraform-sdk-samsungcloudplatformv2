# LogDownloadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndAt** | **string** | Query finish date | 
**FileType** | **string** | Download file type | 
**ProductName** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**ResourceName** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to **NullableString** |  | [optional] 
**RootResourceId** | Pointer to **NullableString** |  | [optional] 
**StartAt** | **string** | Query start date | 
**Status** | Pointer to **NullableString** |  | [optional] 
**TimeZoneInfo** | **string** | Timezone info | 
**UserName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLogDownloadRequest

`func NewLogDownloadRequest(endAt string, fileType string, startAt string, timeZoneInfo string, ) *LogDownloadRequest`

NewLogDownloadRequest instantiates a new LogDownloadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogDownloadRequestWithDefaults

`func NewLogDownloadRequestWithDefaults() *LogDownloadRequest`

NewLogDownloadRequestWithDefaults instantiates a new LogDownloadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndAt

`func (o *LogDownloadRequest) GetEndAt() string`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *LogDownloadRequest) GetEndAtOk() (*string, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *LogDownloadRequest) SetEndAt(v string)`

SetEndAt sets EndAt field to given value.


### GetFileType

`func (o *LogDownloadRequest) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *LogDownloadRequest) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *LogDownloadRequest) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetProductName

`func (o *LogDownloadRequest) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *LogDownloadRequest) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *LogDownloadRequest) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *LogDownloadRequest) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *LogDownloadRequest) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *LogDownloadRequest) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetRegion

`func (o *LogDownloadRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LogDownloadRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LogDownloadRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LogDownloadRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *LogDownloadRequest) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *LogDownloadRequest) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetResourceName

`func (o *LogDownloadRequest) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *LogDownloadRequest) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *LogDownloadRequest) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *LogDownloadRequest) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *LogDownloadRequest) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *LogDownloadRequest) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceType

`func (o *LogDownloadRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *LogDownloadRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *LogDownloadRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *LogDownloadRequest) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *LogDownloadRequest) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *LogDownloadRequest) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetRootResourceId

`func (o *LogDownloadRequest) GetRootResourceId() string`

GetRootResourceId returns the RootResourceId field if non-nil, zero value otherwise.

### GetRootResourceIdOk

`func (o *LogDownloadRequest) GetRootResourceIdOk() (*string, bool)`

GetRootResourceIdOk returns a tuple with the RootResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootResourceId

`func (o *LogDownloadRequest) SetRootResourceId(v string)`

SetRootResourceId sets RootResourceId field to given value.

### HasRootResourceId

`func (o *LogDownloadRequest) HasRootResourceId() bool`

HasRootResourceId returns a boolean if a field has been set.

### SetRootResourceIdNil

`func (o *LogDownloadRequest) SetRootResourceIdNil(b bool)`

 SetRootResourceIdNil sets the value for RootResourceId to be an explicit nil

### UnsetRootResourceId
`func (o *LogDownloadRequest) UnsetRootResourceId()`

UnsetRootResourceId ensures that no value is present for RootResourceId, not even an explicit nil
### GetStartAt

`func (o *LogDownloadRequest) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *LogDownloadRequest) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *LogDownloadRequest) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.


### GetStatus

`func (o *LogDownloadRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogDownloadRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogDownloadRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LogDownloadRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *LogDownloadRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *LogDownloadRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTimeZoneInfo

`func (o *LogDownloadRequest) GetTimeZoneInfo() string`

GetTimeZoneInfo returns the TimeZoneInfo field if non-nil, zero value otherwise.

### GetTimeZoneInfoOk

`func (o *LogDownloadRequest) GetTimeZoneInfoOk() (*string, bool)`

GetTimeZoneInfoOk returns a tuple with the TimeZoneInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneInfo

`func (o *LogDownloadRequest) SetTimeZoneInfo(v string)`

SetTimeZoneInfo sets TimeZoneInfo field to given value.


### GetUserName

`func (o *LogDownloadRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *LogDownloadRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *LogDownloadRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *LogDownloadRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *LogDownloadRequest) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *LogDownloadRequest) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


