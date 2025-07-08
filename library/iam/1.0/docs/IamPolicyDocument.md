# IamPolicyDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | [**[]Statement**](Statement.md) | Statement | 
**Version** | **string** | Policy Version | 

## Methods

### NewIamPolicyDocument

`func NewIamPolicyDocument(statement []Statement, version string, ) *IamPolicyDocument`

NewIamPolicyDocument instantiates a new IamPolicyDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPolicyDocumentWithDefaults

`func NewIamPolicyDocumentWithDefaults() *IamPolicyDocument`

NewIamPolicyDocumentWithDefaults instantiates a new IamPolicyDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *IamPolicyDocument) GetStatement() []Statement`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *IamPolicyDocument) GetStatementOk() (*[]Statement, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *IamPolicyDocument) SetStatement(v []Statement)`

SetStatement sets Statement field to given value.


### GetVersion

`func (o *IamPolicyDocument) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IamPolicyDocument) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IamPolicyDocument) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


