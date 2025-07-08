# VerticaClusterDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AllowableIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**ConsoleIncluded** | Pointer to **bool** | Management Console included state | [optional] [default to false]
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**DbaasEngine** | [**DbaasEngine**](DbaasEngine.md) | DBaaS engine | 
**DbaasEngineVersionName** | **string** | DBaaS engine version name | 
**Id** | **string** | ID | 
**InitConfigOption** | [**VerticaInitConfigOptionResponse**](VerticaInitConfigOptionResponse.md) | DB initial configuration option | 
**InstanceCount** | Pointer to **int32** | Instance Count | [optional] [default to 0]
**InstanceGroups** | [**[]InstanceGroupResponse**](InstanceGroupResponse.md) | Instance groups list | 
**MaintenanceOption** | Pointer to [**NullableMaintenanceResponseOption**](MaintenanceResponseOption.md) |  | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Cluster name | 
**NatEnabled** | Pointer to **NullableBool** |  | [optional] 
**ProductType** | [**ProductType**](ProductType.md) | Product type | 
**RoleType** | [**NullableClusterRoleType**](ClusterRoleType.md) |  | 
**ServiceState** | [**ServiceState**](ServiceState.md) | Service state | 
**SoftwareVersion** | **string** | Software version | 
**SubnetId** | **string** | Subnet ID | 
**Timezone** | **string** | Timezone | 

## Methods

### NewVerticaClusterDetailResponse

`func NewVerticaClusterDetailResponse(accountId string, createdAt time.Time, createdBy string, dbaasEngine DbaasEngine, dbaasEngineVersionName string, id string, initConfigOption VerticaInitConfigOptionResponse, instanceGroups []InstanceGroupResponse, modifiedAt time.Time, modifiedBy string, name string, productType ProductType, roleType NullableClusterRoleType, serviceState ServiceState, softwareVersion string, subnetId string, timezone string, ) *VerticaClusterDetailResponse`

NewVerticaClusterDetailResponse instantiates a new VerticaClusterDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerticaClusterDetailResponseWithDefaults

`func NewVerticaClusterDetailResponseWithDefaults() *VerticaClusterDetailResponse`

NewVerticaClusterDetailResponseWithDefaults instantiates a new VerticaClusterDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *VerticaClusterDetailResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VerticaClusterDetailResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VerticaClusterDetailResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAllowableIpAddresses

`func (o *VerticaClusterDetailResponse) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *VerticaClusterDetailResponse) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *VerticaClusterDetailResponse) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *VerticaClusterDetailResponse) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetConsoleIncluded

`func (o *VerticaClusterDetailResponse) GetConsoleIncluded() bool`

GetConsoleIncluded returns the ConsoleIncluded field if non-nil, zero value otherwise.

### GetConsoleIncludedOk

`func (o *VerticaClusterDetailResponse) GetConsoleIncludedOk() (*bool, bool)`

GetConsoleIncludedOk returns a tuple with the ConsoleIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleIncluded

`func (o *VerticaClusterDetailResponse) SetConsoleIncluded(v bool)`

SetConsoleIncluded sets ConsoleIncluded field to given value.

### HasConsoleIncluded

`func (o *VerticaClusterDetailResponse) HasConsoleIncluded() bool`

HasConsoleIncluded returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VerticaClusterDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VerticaClusterDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VerticaClusterDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *VerticaClusterDetailResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VerticaClusterDetailResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VerticaClusterDetailResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDbaasEngine

`func (o *VerticaClusterDetailResponse) GetDbaasEngine() DbaasEngine`

GetDbaasEngine returns the DbaasEngine field if non-nil, zero value otherwise.

### GetDbaasEngineOk

`func (o *VerticaClusterDetailResponse) GetDbaasEngineOk() (*DbaasEngine, bool)`

GetDbaasEngineOk returns a tuple with the DbaasEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngine

`func (o *VerticaClusterDetailResponse) SetDbaasEngine(v DbaasEngine)`

SetDbaasEngine sets DbaasEngine field to given value.


### GetDbaasEngineVersionName

`func (o *VerticaClusterDetailResponse) GetDbaasEngineVersionName() string`

GetDbaasEngineVersionName returns the DbaasEngineVersionName field if non-nil, zero value otherwise.

### GetDbaasEngineVersionNameOk

`func (o *VerticaClusterDetailResponse) GetDbaasEngineVersionNameOk() (*string, bool)`

GetDbaasEngineVersionNameOk returns a tuple with the DbaasEngineVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngineVersionName

`func (o *VerticaClusterDetailResponse) SetDbaasEngineVersionName(v string)`

SetDbaasEngineVersionName sets DbaasEngineVersionName field to given value.


### GetId

`func (o *VerticaClusterDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerticaClusterDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerticaClusterDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInitConfigOption

`func (o *VerticaClusterDetailResponse) GetInitConfigOption() VerticaInitConfigOptionResponse`

GetInitConfigOption returns the InitConfigOption field if non-nil, zero value otherwise.

### GetInitConfigOptionOk

`func (o *VerticaClusterDetailResponse) GetInitConfigOptionOk() (*VerticaInitConfigOptionResponse, bool)`

GetInitConfigOptionOk returns a tuple with the InitConfigOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitConfigOption

`func (o *VerticaClusterDetailResponse) SetInitConfigOption(v VerticaInitConfigOptionResponse)`

SetInitConfigOption sets InitConfigOption field to given value.


### GetInstanceCount

`func (o *VerticaClusterDetailResponse) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *VerticaClusterDetailResponse) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *VerticaClusterDetailResponse) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *VerticaClusterDetailResponse) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetInstanceGroups

`func (o *VerticaClusterDetailResponse) GetInstanceGroups() []InstanceGroupResponse`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *VerticaClusterDetailResponse) GetInstanceGroupsOk() (*[]InstanceGroupResponse, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *VerticaClusterDetailResponse) SetInstanceGroups(v []InstanceGroupResponse)`

SetInstanceGroups sets InstanceGroups field to given value.


### GetMaintenanceOption

`func (o *VerticaClusterDetailResponse) GetMaintenanceOption() MaintenanceResponseOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *VerticaClusterDetailResponse) GetMaintenanceOptionOk() (*MaintenanceResponseOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *VerticaClusterDetailResponse) SetMaintenanceOption(v MaintenanceResponseOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *VerticaClusterDetailResponse) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *VerticaClusterDetailResponse) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *VerticaClusterDetailResponse) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetModifiedAt

`func (o *VerticaClusterDetailResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VerticaClusterDetailResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VerticaClusterDetailResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *VerticaClusterDetailResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VerticaClusterDetailResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VerticaClusterDetailResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *VerticaClusterDetailResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VerticaClusterDetailResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VerticaClusterDetailResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNatEnabled

`func (o *VerticaClusterDetailResponse) GetNatEnabled() bool`

GetNatEnabled returns the NatEnabled field if non-nil, zero value otherwise.

### GetNatEnabledOk

`func (o *VerticaClusterDetailResponse) GetNatEnabledOk() (*bool, bool)`

GetNatEnabledOk returns a tuple with the NatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatEnabled

`func (o *VerticaClusterDetailResponse) SetNatEnabled(v bool)`

SetNatEnabled sets NatEnabled field to given value.

### HasNatEnabled

`func (o *VerticaClusterDetailResponse) HasNatEnabled() bool`

HasNatEnabled returns a boolean if a field has been set.

### SetNatEnabledNil

`func (o *VerticaClusterDetailResponse) SetNatEnabledNil(b bool)`

 SetNatEnabledNil sets the value for NatEnabled to be an explicit nil

### UnsetNatEnabled
`func (o *VerticaClusterDetailResponse) UnsetNatEnabled()`

UnsetNatEnabled ensures that no value is present for NatEnabled, not even an explicit nil
### GetProductType

`func (o *VerticaClusterDetailResponse) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *VerticaClusterDetailResponse) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *VerticaClusterDetailResponse) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetRoleType

`func (o *VerticaClusterDetailResponse) GetRoleType() ClusterRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *VerticaClusterDetailResponse) GetRoleTypeOk() (*ClusterRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *VerticaClusterDetailResponse) SetRoleType(v ClusterRoleType)`

SetRoleType sets RoleType field to given value.


### SetRoleTypeNil

`func (o *VerticaClusterDetailResponse) SetRoleTypeNil(b bool)`

 SetRoleTypeNil sets the value for RoleType to be an explicit nil

### UnsetRoleType
`func (o *VerticaClusterDetailResponse) UnsetRoleType()`

UnsetRoleType ensures that no value is present for RoleType, not even an explicit nil
### GetServiceState

`func (o *VerticaClusterDetailResponse) GetServiceState() ServiceState`

GetServiceState returns the ServiceState field if non-nil, zero value otherwise.

### GetServiceStateOk

`func (o *VerticaClusterDetailResponse) GetServiceStateOk() (*ServiceState, bool)`

GetServiceStateOk returns a tuple with the ServiceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceState

`func (o *VerticaClusterDetailResponse) SetServiceState(v ServiceState)`

SetServiceState sets ServiceState field to given value.


### GetSoftwareVersion

`func (o *VerticaClusterDetailResponse) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *VerticaClusterDetailResponse) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *VerticaClusterDetailResponse) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.


### GetSubnetId

`func (o *VerticaClusterDetailResponse) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *VerticaClusterDetailResponse) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *VerticaClusterDetailResponse) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTimezone

`func (o *VerticaClusterDetailResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *VerticaClusterDetailResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *VerticaClusterDetailResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


