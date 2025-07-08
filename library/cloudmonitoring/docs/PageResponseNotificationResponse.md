# PageResponseNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | Pointer to [**[]NotificationResponse**](NotificationResponse.md) |  | [optional] 
**IsChanged** | Pointer to **bool** |  | [optional] 
**ReturnCount** | Pointer to **int64** |  | [optional] 
**ReturnPage** | Pointer to **int64** |  | [optional] 
**Sort** | Pointer to **[]string** |  | [optional] 
**SortableFields** | Pointer to **[]string** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewPageResponseNotificationResponse

`func NewPageResponseNotificationResponse() *PageResponseNotificationResponse`

NewPageResponseNotificationResponse instantiates a new PageResponseNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageResponseNotificationResponseWithDefaults

`func NewPageResponseNotificationResponseWithDefaults() *PageResponseNotificationResponse`

NewPageResponseNotificationResponseWithDefaults instantiates a new PageResponseNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *PageResponseNotificationResponse) GetContents() []NotificationResponse`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *PageResponseNotificationResponse) GetContentsOk() (*[]NotificationResponse, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *PageResponseNotificationResponse) SetContents(v []NotificationResponse)`

SetContents sets Contents field to given value.

### HasContents

`func (o *PageResponseNotificationResponse) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetIsChanged

`func (o *PageResponseNotificationResponse) GetIsChanged() bool`

GetIsChanged returns the IsChanged field if non-nil, zero value otherwise.

### GetIsChangedOk

`func (o *PageResponseNotificationResponse) GetIsChangedOk() (*bool, bool)`

GetIsChangedOk returns a tuple with the IsChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChanged

`func (o *PageResponseNotificationResponse) SetIsChanged(v bool)`

SetIsChanged sets IsChanged field to given value.

### HasIsChanged

`func (o *PageResponseNotificationResponse) HasIsChanged() bool`

HasIsChanged returns a boolean if a field has been set.

### GetReturnCount

`func (o *PageResponseNotificationResponse) GetReturnCount() int64`

GetReturnCount returns the ReturnCount field if non-nil, zero value otherwise.

### GetReturnCountOk

`func (o *PageResponseNotificationResponse) GetReturnCountOk() (*int64, bool)`

GetReturnCountOk returns a tuple with the ReturnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCount

`func (o *PageResponseNotificationResponse) SetReturnCount(v int64)`

SetReturnCount sets ReturnCount field to given value.

### HasReturnCount

`func (o *PageResponseNotificationResponse) HasReturnCount() bool`

HasReturnCount returns a boolean if a field has been set.

### GetReturnPage

`func (o *PageResponseNotificationResponse) GetReturnPage() int64`

GetReturnPage returns the ReturnPage field if non-nil, zero value otherwise.

### GetReturnPageOk

`func (o *PageResponseNotificationResponse) GetReturnPageOk() (*int64, bool)`

GetReturnPageOk returns a tuple with the ReturnPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPage

`func (o *PageResponseNotificationResponse) SetReturnPage(v int64)`

SetReturnPage sets ReturnPage field to given value.

### HasReturnPage

`func (o *PageResponseNotificationResponse) HasReturnPage() bool`

HasReturnPage returns a boolean if a field has been set.

### GetSort

`func (o *PageResponseNotificationResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PageResponseNotificationResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PageResponseNotificationResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PageResponseNotificationResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSortableFields

`func (o *PageResponseNotificationResponse) GetSortableFields() []string`

GetSortableFields returns the SortableFields field if non-nil, zero value otherwise.

### GetSortableFieldsOk

`func (o *PageResponseNotificationResponse) GetSortableFieldsOk() (*[]string, bool)`

GetSortableFieldsOk returns a tuple with the SortableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortableFields

`func (o *PageResponseNotificationResponse) SetSortableFields(v []string)`

SetSortableFields sets SortableFields field to given value.

### HasSortableFields

`func (o *PageResponseNotificationResponse) HasSortableFields() bool`

HasSortableFields returns a boolean if a field has been set.

### GetTotalCount

`func (o *PageResponseNotificationResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PageResponseNotificationResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PageResponseNotificationResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PageResponseNotificationResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


