# FirewallSetStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**FirewallStatusType**](FirewallStatusType.md) | Firewall Status | 

## Methods

### NewFirewallSetStatusRequest

`func NewFirewallSetStatusRequest(status FirewallStatusType, ) *FirewallSetStatusRequest`

NewFirewallSetStatusRequest instantiates a new FirewallSetStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSetStatusRequestWithDefaults

`func NewFirewallSetStatusRequestWithDefaults() *FirewallSetStatusRequest`

NewFirewallSetStatusRequestWithDefaults instantiates a new FirewallSetStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FirewallSetStatusRequest) GetStatus() FirewallStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallSetStatusRequest) GetStatusOk() (*FirewallStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallSetStatusRequest) SetStatus(v FirewallStatusType)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


