# PolicyDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | [**[]Statement**](Statement.md) | Statement | 
**Version** | **string** | Policy Version | 

## Methods

### NewPolicyDocument

`func NewPolicyDocument(statement []Statement, version string, ) *PolicyDocument`

NewPolicyDocument instantiates a new PolicyDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDocumentWithDefaults

`func NewPolicyDocumentWithDefaults() *PolicyDocument`

NewPolicyDocumentWithDefaults instantiates a new PolicyDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *PolicyDocument) GetStatement() []Statement`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *PolicyDocument) GetStatementOk() (*[]Statement, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *PolicyDocument) SetStatement(v []Statement)`

SetStatement sets Statement field to given value.


### GetVersion

`func (o *PolicyDocument) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PolicyDocument) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PolicyDocument) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


