# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | 생성 일시 | 
**CreatedBy** | **string** | 생성자 | 
**CreatorEmail** | Pointer to **string** | 생성자 Email | [optional] [default to "-"]
**CreatorName** | Pointer to **string** | 생성자 성, 이름 | [optional] [default to "-"]
**Description** | **NullableString** |  | 
**DomainName** | **string** | 도메인 이름 | 
**Id** | **string** | ID | 
**Members** | Pointer to [**[]GroupMember**](GroupMember.md) |  | [optional] 
**ModifiedAt** | **time.Time** | 수정 일시 | 
**ModifiedBy** | **string** | 수정자 | 
**ModifierEmail** | Pointer to **string** | 수정자 Email | [optional] [default to "-"]
**ModifierName** | Pointer to **string** | 수정자 성, 이름 | [optional] [default to "-"]
**Name** | **string** | Group 이름 | 
**Policies** | [**[]Policy**](Policy.md) |  | 
**ResourceType** | Pointer to **NullableString** |  | [optional] 
**Roles** | [**[]Role**](Role.md) |  | 
**ServiceName** | Pointer to **NullableString** |  | [optional] 
**ServiceType** | Pointer to **NullableString** |  | [optional] 
**Srn** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** | Group Type | 

## Methods

### NewGroup

`func NewGroup(createdAt time.Time, createdBy string, description NullableString, domainName string, id string, modifiedAt time.Time, modifiedBy string, name string, policies []Policy, roles []Role, type_ string, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Group) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Group) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Group) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Group) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Group) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Group) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatorEmail

`func (o *Group) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *Group) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *Group) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *Group) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetCreatorName

`func (o *Group) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *Group) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *Group) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *Group) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Group) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Group) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomainName

`func (o *Group) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Group) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Group) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.


### GetMembers

`func (o *Group) GetMembers() []GroupMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Group) GetMembersOk() (*[]GroupMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Group) SetMembers(v []GroupMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Group) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *Group) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *Group) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetModifiedAt

`func (o *Group) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Group) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Group) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *Group) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Group) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Group) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetModifierEmail

`func (o *Group) GetModifierEmail() string`

GetModifierEmail returns the ModifierEmail field if non-nil, zero value otherwise.

### GetModifierEmailOk

`func (o *Group) GetModifierEmailOk() (*string, bool)`

GetModifierEmailOk returns a tuple with the ModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierEmail

`func (o *Group) SetModifierEmail(v string)`

SetModifierEmail sets ModifierEmail field to given value.

### HasModifierEmail

`func (o *Group) HasModifierEmail() bool`

HasModifierEmail returns a boolean if a field has been set.

### GetModifierName

`func (o *Group) GetModifierName() string`

GetModifierName returns the ModifierName field if non-nil, zero value otherwise.

### GetModifierNameOk

`func (o *Group) GetModifierNameOk() (*string, bool)`

GetModifierNameOk returns a tuple with the ModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierName

`func (o *Group) SetModifierName(v string)`

SetModifierName sets ModifierName field to given value.

### HasModifierName

`func (o *Group) HasModifierName() bool`

HasModifierName returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### GetPolicies

`func (o *Group) GetPolicies() []Policy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *Group) GetPoliciesOk() (*[]Policy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *Group) SetPolicies(v []Policy)`

SetPolicies sets Policies field to given value.


### SetPoliciesNil

`func (o *Group) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *Group) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil
### GetResourceType

`func (o *Group) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Group) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Group) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Group) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *Group) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *Group) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetRoles

`func (o *Group) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Group) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Group) SetRoles(v []Role)`

SetRoles sets Roles field to given value.


### SetRolesNil

`func (o *Group) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *Group) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetServiceName

`func (o *Group) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Group) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Group) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *Group) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *Group) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *Group) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetServiceType

`func (o *Group) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Group) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Group) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *Group) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### SetServiceTypeNil

`func (o *Group) SetServiceTypeNil(b bool)`

 SetServiceTypeNil sets the value for ServiceType to be an explicit nil

### UnsetServiceType
`func (o *Group) UnsetServiceType()`

UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
### GetSrn

`func (o *Group) GetSrn() string`

GetSrn returns the Srn field if non-nil, zero value otherwise.

### GetSrnOk

`func (o *Group) GetSrnOk() (*string, bool)`

GetSrnOk returns a tuple with the Srn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrn

`func (o *Group) SetSrn(v string)`

SetSrn sets Srn field to given value.

### HasSrn

`func (o *Group) HasSrn() bool`

HasSrn returns a boolean if a field has been set.

### SetSrnNil

`func (o *Group) SetSrnNil(b bool)`

 SetSrnNil sets the value for Srn to be an explicit nil

### UnsetSrn
`func (o *Group) UnsetSrn()`

UnsetSrn ensures that no value is present for Srn, not even an explicit nil
### GetType

`func (o *Group) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Group) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Group) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


