# AsyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Request ID | 
**Resource** | [**ResourceInfo**](ResourceInfo.md) | Resource | 

## Methods

### NewAsyncResponse

`func NewAsyncResponse(requestId string, resource ResourceInfo, ) *AsyncResponse`

NewAsyncResponse instantiates a new AsyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncResponseWithDefaults

`func NewAsyncResponseWithDefaults() *AsyncResponse`

NewAsyncResponseWithDefaults instantiates a new AsyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *AsyncResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AsyncResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AsyncResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetResource

`func (o *AsyncResponse) GetResource() ResourceInfo`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AsyncResponse) GetResourceOk() (*ResourceInfo, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AsyncResponse) SetResource(v ResourceInfo)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


