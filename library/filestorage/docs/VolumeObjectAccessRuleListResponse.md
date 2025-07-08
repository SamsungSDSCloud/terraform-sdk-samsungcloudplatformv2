# VolumeObjectAccessRuleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRules** | [**[]VolumeObjectAccessRuleListBase**](VolumeObjectAccessRuleListBase.md) |  | 
**Count** | **int32** | count | 
**VolumeId** | **string** | Volume ID | 

## Methods

### NewVolumeObjectAccessRuleListResponse

`func NewVolumeObjectAccessRuleListResponse(accessRules []VolumeObjectAccessRuleListBase, count int32, volumeId string, ) *VolumeObjectAccessRuleListResponse`

NewVolumeObjectAccessRuleListResponse instantiates a new VolumeObjectAccessRuleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeObjectAccessRuleListResponseWithDefaults

`func NewVolumeObjectAccessRuleListResponseWithDefaults() *VolumeObjectAccessRuleListResponse`

NewVolumeObjectAccessRuleListResponseWithDefaults instantiates a new VolumeObjectAccessRuleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRules

`func (o *VolumeObjectAccessRuleListResponse) GetAccessRules() []VolumeObjectAccessRuleListBase`

GetAccessRules returns the AccessRules field if non-nil, zero value otherwise.

### GetAccessRulesOk

`func (o *VolumeObjectAccessRuleListResponse) GetAccessRulesOk() (*[]VolumeObjectAccessRuleListBase, bool)`

GetAccessRulesOk returns a tuple with the AccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRules

`func (o *VolumeObjectAccessRuleListResponse) SetAccessRules(v []VolumeObjectAccessRuleListBase)`

SetAccessRules sets AccessRules field to given value.


### GetCount

`func (o *VolumeObjectAccessRuleListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VolumeObjectAccessRuleListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VolumeObjectAccessRuleListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetVolumeId

`func (o *VolumeObjectAccessRuleListResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeObjectAccessRuleListResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeObjectAccessRuleListResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


