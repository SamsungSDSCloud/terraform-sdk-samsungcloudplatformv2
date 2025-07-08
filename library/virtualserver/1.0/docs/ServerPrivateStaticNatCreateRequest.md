# ServerPrivateStaticNatCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateNatId** | **string** | Private NAT ID | 
**PrivateNatIpId** | **string** | Private NAT IP ID | 

## Methods

### NewServerPrivateStaticNatCreateRequest

`func NewServerPrivateStaticNatCreateRequest(privateNatId string, privateNatIpId string, ) *ServerPrivateStaticNatCreateRequest`

NewServerPrivateStaticNatCreateRequest instantiates a new ServerPrivateStaticNatCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPrivateStaticNatCreateRequestWithDefaults

`func NewServerPrivateStaticNatCreateRequestWithDefaults() *ServerPrivateStaticNatCreateRequest`

NewServerPrivateStaticNatCreateRequestWithDefaults instantiates a new ServerPrivateStaticNatCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateNatId

`func (o *ServerPrivateStaticNatCreateRequest) GetPrivateNatId() string`

GetPrivateNatId returns the PrivateNatId field if non-nil, zero value otherwise.

### GetPrivateNatIdOk

`func (o *ServerPrivateStaticNatCreateRequest) GetPrivateNatIdOk() (*string, bool)`

GetPrivateNatIdOk returns a tuple with the PrivateNatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNatId

`func (o *ServerPrivateStaticNatCreateRequest) SetPrivateNatId(v string)`

SetPrivateNatId sets PrivateNatId field to given value.


### GetPrivateNatIpId

`func (o *ServerPrivateStaticNatCreateRequest) GetPrivateNatIpId() string`

GetPrivateNatIpId returns the PrivateNatIpId field if non-nil, zero value otherwise.

### GetPrivateNatIpIdOk

`func (o *ServerPrivateStaticNatCreateRequest) GetPrivateNatIpIdOk() (*string, bool)`

GetPrivateNatIpIdOk returns a tuple with the PrivateNatIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNatIpId

`func (o *ServerPrivateStaticNatCreateRequest) SetPrivateNatIpId(v string)`

SetPrivateNatIpId sets PrivateNatIpId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


