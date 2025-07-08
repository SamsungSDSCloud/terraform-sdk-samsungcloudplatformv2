# LbHealthCheckCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**HealthCheckCount** | Pointer to **int32** | LB Health Check count | [optional] [default to 3]
**HealthCheckInterval** | Pointer to **int32** | LB Health Check interval | [optional] [default to 5]
**HealthCheckPort** | Pointer to **int32** | LB Health Check port | [optional] 
**HealthCheckTimeout** | Pointer to **int32** | LB Health Check timeout | [optional] [default to 5]
**HealthCheckUrl** | Pointer to **NullableString** |  | [optional] 
**HttpMethod** | Pointer to [**NullableLbMonitorHttpMethod**](LbMonitorHttpMethod.md) |  | [optional] 
**Name** | **string** | LB Health Check name | 
**Protocol** | [**LbMonitorProtocol**](LbMonitorProtocol.md) | LB Health Check Protocol | 
**RequestData** | Pointer to **NullableString** |  | [optional] 
**ResponseCode** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | **string** | Service Subnet ID | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]
**VpcId** | **string** | VPC ID | 

## Methods

### NewLbHealthCheckCreate

`func NewLbHealthCheckCreate(name string, protocol LbMonitorProtocol, subnetId string, vpcId string, ) *LbHealthCheckCreate`

NewLbHealthCheckCreate instantiates a new LbHealthCheckCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbHealthCheckCreateWithDefaults

`func NewLbHealthCheckCreateWithDefaults() *LbHealthCheckCreate`

NewLbHealthCheckCreateWithDefaults instantiates a new LbHealthCheckCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LbHealthCheckCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LbHealthCheckCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LbHealthCheckCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LbHealthCheckCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LbHealthCheckCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LbHealthCheckCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHealthCheckCount

`func (o *LbHealthCheckCreate) GetHealthCheckCount() int32`

GetHealthCheckCount returns the HealthCheckCount field if non-nil, zero value otherwise.

### GetHealthCheckCountOk

`func (o *LbHealthCheckCreate) GetHealthCheckCountOk() (*int32, bool)`

GetHealthCheckCountOk returns a tuple with the HealthCheckCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckCount

`func (o *LbHealthCheckCreate) SetHealthCheckCount(v int32)`

SetHealthCheckCount sets HealthCheckCount field to given value.

### HasHealthCheckCount

`func (o *LbHealthCheckCreate) HasHealthCheckCount() bool`

HasHealthCheckCount returns a boolean if a field has been set.

### GetHealthCheckInterval

`func (o *LbHealthCheckCreate) GetHealthCheckInterval() int32`

GetHealthCheckInterval returns the HealthCheckInterval field if non-nil, zero value otherwise.

### GetHealthCheckIntervalOk

`func (o *LbHealthCheckCreate) GetHealthCheckIntervalOk() (*int32, bool)`

GetHealthCheckIntervalOk returns a tuple with the HealthCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckInterval

`func (o *LbHealthCheckCreate) SetHealthCheckInterval(v int32)`

SetHealthCheckInterval sets HealthCheckInterval field to given value.

### HasHealthCheckInterval

`func (o *LbHealthCheckCreate) HasHealthCheckInterval() bool`

HasHealthCheckInterval returns a boolean if a field has been set.

### GetHealthCheckPort

`func (o *LbHealthCheckCreate) GetHealthCheckPort() int32`

GetHealthCheckPort returns the HealthCheckPort field if non-nil, zero value otherwise.

### GetHealthCheckPortOk

`func (o *LbHealthCheckCreate) GetHealthCheckPortOk() (*int32, bool)`

GetHealthCheckPortOk returns a tuple with the HealthCheckPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckPort

`func (o *LbHealthCheckCreate) SetHealthCheckPort(v int32)`

SetHealthCheckPort sets HealthCheckPort field to given value.

### HasHealthCheckPort

`func (o *LbHealthCheckCreate) HasHealthCheckPort() bool`

HasHealthCheckPort returns a boolean if a field has been set.

### GetHealthCheckTimeout

`func (o *LbHealthCheckCreate) GetHealthCheckTimeout() int32`

GetHealthCheckTimeout returns the HealthCheckTimeout field if non-nil, zero value otherwise.

### GetHealthCheckTimeoutOk

`func (o *LbHealthCheckCreate) GetHealthCheckTimeoutOk() (*int32, bool)`

GetHealthCheckTimeoutOk returns a tuple with the HealthCheckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckTimeout

`func (o *LbHealthCheckCreate) SetHealthCheckTimeout(v int32)`

SetHealthCheckTimeout sets HealthCheckTimeout field to given value.

### HasHealthCheckTimeout

`func (o *LbHealthCheckCreate) HasHealthCheckTimeout() bool`

HasHealthCheckTimeout returns a boolean if a field has been set.

### GetHealthCheckUrl

`func (o *LbHealthCheckCreate) GetHealthCheckUrl() string`

GetHealthCheckUrl returns the HealthCheckUrl field if non-nil, zero value otherwise.

### GetHealthCheckUrlOk

`func (o *LbHealthCheckCreate) GetHealthCheckUrlOk() (*string, bool)`

GetHealthCheckUrlOk returns a tuple with the HealthCheckUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckUrl

`func (o *LbHealthCheckCreate) SetHealthCheckUrl(v string)`

SetHealthCheckUrl sets HealthCheckUrl field to given value.

### HasHealthCheckUrl

`func (o *LbHealthCheckCreate) HasHealthCheckUrl() bool`

HasHealthCheckUrl returns a boolean if a field has been set.

### SetHealthCheckUrlNil

`func (o *LbHealthCheckCreate) SetHealthCheckUrlNil(b bool)`

 SetHealthCheckUrlNil sets the value for HealthCheckUrl to be an explicit nil

### UnsetHealthCheckUrl
`func (o *LbHealthCheckCreate) UnsetHealthCheckUrl()`

UnsetHealthCheckUrl ensures that no value is present for HealthCheckUrl, not even an explicit nil
### GetHttpMethod

`func (o *LbHealthCheckCreate) GetHttpMethod() LbMonitorHttpMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *LbHealthCheckCreate) GetHttpMethodOk() (*LbMonitorHttpMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *LbHealthCheckCreate) SetHttpMethod(v LbMonitorHttpMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *LbHealthCheckCreate) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### SetHttpMethodNil

`func (o *LbHealthCheckCreate) SetHttpMethodNil(b bool)`

 SetHttpMethodNil sets the value for HttpMethod to be an explicit nil

### UnsetHttpMethod
`func (o *LbHealthCheckCreate) UnsetHttpMethod()`

UnsetHttpMethod ensures that no value is present for HttpMethod, not even an explicit nil
### GetName

`func (o *LbHealthCheckCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LbHealthCheckCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LbHealthCheckCreate) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *LbHealthCheckCreate) GetProtocol() LbMonitorProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LbHealthCheckCreate) GetProtocolOk() (*LbMonitorProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LbHealthCheckCreate) SetProtocol(v LbMonitorProtocol)`

SetProtocol sets Protocol field to given value.


### GetRequestData

`func (o *LbHealthCheckCreate) GetRequestData() string`

GetRequestData returns the RequestData field if non-nil, zero value otherwise.

### GetRequestDataOk

`func (o *LbHealthCheckCreate) GetRequestDataOk() (*string, bool)`

GetRequestDataOk returns a tuple with the RequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestData

`func (o *LbHealthCheckCreate) SetRequestData(v string)`

SetRequestData sets RequestData field to given value.

### HasRequestData

`func (o *LbHealthCheckCreate) HasRequestData() bool`

HasRequestData returns a boolean if a field has been set.

### SetRequestDataNil

`func (o *LbHealthCheckCreate) SetRequestDataNil(b bool)`

 SetRequestDataNil sets the value for RequestData to be an explicit nil

### UnsetRequestData
`func (o *LbHealthCheckCreate) UnsetRequestData()`

UnsetRequestData ensures that no value is present for RequestData, not even an explicit nil
### GetResponseCode

`func (o *LbHealthCheckCreate) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *LbHealthCheckCreate) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *LbHealthCheckCreate) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *LbHealthCheckCreate) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### SetResponseCodeNil

`func (o *LbHealthCheckCreate) SetResponseCodeNil(b bool)`

 SetResponseCodeNil sets the value for ResponseCode to be an explicit nil

### UnsetResponseCode
`func (o *LbHealthCheckCreate) UnsetResponseCode()`

UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
### GetSubnetId

`func (o *LbHealthCheckCreate) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *LbHealthCheckCreate) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *LbHealthCheckCreate) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *LbHealthCheckCreate) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LbHealthCheckCreate) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LbHealthCheckCreate) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LbHealthCheckCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVpcId

`func (o *LbHealthCheckCreate) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *LbHealthCheckCreate) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *LbHealthCheckCreate) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


