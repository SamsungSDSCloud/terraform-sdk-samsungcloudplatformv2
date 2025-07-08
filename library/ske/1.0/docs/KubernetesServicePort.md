# KubernetesServicePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodePort** | **NullableString** |  | 
**Port** | **NullableString** |  | 
**PortName** | **NullableString** |  | 
**Protocol** | **NullableString** |  | 
**TargetPort** | **NullableString** |  | 

## Methods

### NewKubernetesServicePort

`func NewKubernetesServicePort(nodePort NullableString, port NullableString, portName NullableString, protocol NullableString, targetPort NullableString, ) *KubernetesServicePort`

NewKubernetesServicePort instantiates a new KubernetesServicePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesServicePortWithDefaults

`func NewKubernetesServicePortWithDefaults() *KubernetesServicePort`

NewKubernetesServicePortWithDefaults instantiates a new KubernetesServicePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodePort

`func (o *KubernetesServicePort) GetNodePort() string`

GetNodePort returns the NodePort field if non-nil, zero value otherwise.

### GetNodePortOk

`func (o *KubernetesServicePort) GetNodePortOk() (*string, bool)`

GetNodePortOk returns a tuple with the NodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePort

`func (o *KubernetesServicePort) SetNodePort(v string)`

SetNodePort sets NodePort field to given value.


### SetNodePortNil

`func (o *KubernetesServicePort) SetNodePortNil(b bool)`

 SetNodePortNil sets the value for NodePort to be an explicit nil

### UnsetNodePort
`func (o *KubernetesServicePort) UnsetNodePort()`

UnsetNodePort ensures that no value is present for NodePort, not even an explicit nil
### GetPort

`func (o *KubernetesServicePort) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KubernetesServicePort) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KubernetesServicePort) SetPort(v string)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *KubernetesServicePort) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *KubernetesServicePort) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPortName

`func (o *KubernetesServicePort) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *KubernetesServicePort) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *KubernetesServicePort) SetPortName(v string)`

SetPortName sets PortName field to given value.


### SetPortNameNil

`func (o *KubernetesServicePort) SetPortNameNil(b bool)`

 SetPortNameNil sets the value for PortName to be an explicit nil

### UnsetPortName
`func (o *KubernetesServicePort) UnsetPortName()`

UnsetPortName ensures that no value is present for PortName, not even an explicit nil
### GetProtocol

`func (o *KubernetesServicePort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *KubernetesServicePort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *KubernetesServicePort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### SetProtocolNil

`func (o *KubernetesServicePort) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *KubernetesServicePort) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetTargetPort

`func (o *KubernetesServicePort) GetTargetPort() string`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *KubernetesServicePort) GetTargetPortOk() (*string, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *KubernetesServicePort) SetTargetPort(v string)`

SetTargetPort sets TargetPort field to given value.


### SetTargetPortNil

`func (o *KubernetesServicePort) SetTargetPortNil(b bool)`

 SetTargetPortNil sets the value for TargetPort to be an explicit nil

### UnsetTargetPort
`func (o *KubernetesServicePort) UnsetTargetPort()`

UnsetTargetPort ensures that no value is present for TargetPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


