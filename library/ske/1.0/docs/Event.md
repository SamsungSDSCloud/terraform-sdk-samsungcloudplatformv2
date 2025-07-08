# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**EventComponent** | **NullableString** |  | 
**EventCount** | **NullableString** |  | 
**EventFirstSeenDate** | **NullableTime** |  | 
**EventHost** | **NullableString** |  | 
**EventKind** | **string** | Event Kind | 
**EventLastSeenDate** | **NullableTime** |  | 
**EventMessage** | **NullableString** |  | 
**EventReason** | **NullableString** |  | 
**EventType** | **NullableString** |  | 
**Name** | **string** | Name | 
**NamespaceName** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 

## Methods

### NewEvent

`func NewEvent(age string, eventComponent NullableString, eventCount NullableString, eventFirstSeenDate NullableTime, eventHost NullableString, eventKind string, eventLastSeenDate NullableTime, eventMessage NullableString, eventReason NullableString, eventType NullableString, name string, namespaceName NullableString, uid string, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *Event) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Event) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Event) SetAge(v string)`

SetAge sets Age field to given value.


### GetEventComponent

`func (o *Event) GetEventComponent() string`

GetEventComponent returns the EventComponent field if non-nil, zero value otherwise.

### GetEventComponentOk

`func (o *Event) GetEventComponentOk() (*string, bool)`

GetEventComponentOk returns a tuple with the EventComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventComponent

`func (o *Event) SetEventComponent(v string)`

SetEventComponent sets EventComponent field to given value.


### SetEventComponentNil

`func (o *Event) SetEventComponentNil(b bool)`

 SetEventComponentNil sets the value for EventComponent to be an explicit nil

### UnsetEventComponent
`func (o *Event) UnsetEventComponent()`

UnsetEventComponent ensures that no value is present for EventComponent, not even an explicit nil
### GetEventCount

`func (o *Event) GetEventCount() string`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *Event) GetEventCountOk() (*string, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *Event) SetEventCount(v string)`

SetEventCount sets EventCount field to given value.


### SetEventCountNil

`func (o *Event) SetEventCountNil(b bool)`

 SetEventCountNil sets the value for EventCount to be an explicit nil

### UnsetEventCount
`func (o *Event) UnsetEventCount()`

UnsetEventCount ensures that no value is present for EventCount, not even an explicit nil
### GetEventFirstSeenDate

`func (o *Event) GetEventFirstSeenDate() time.Time`

GetEventFirstSeenDate returns the EventFirstSeenDate field if non-nil, zero value otherwise.

### GetEventFirstSeenDateOk

`func (o *Event) GetEventFirstSeenDateOk() (*time.Time, bool)`

GetEventFirstSeenDateOk returns a tuple with the EventFirstSeenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFirstSeenDate

`func (o *Event) SetEventFirstSeenDate(v time.Time)`

SetEventFirstSeenDate sets EventFirstSeenDate field to given value.


### SetEventFirstSeenDateNil

`func (o *Event) SetEventFirstSeenDateNil(b bool)`

 SetEventFirstSeenDateNil sets the value for EventFirstSeenDate to be an explicit nil

### UnsetEventFirstSeenDate
`func (o *Event) UnsetEventFirstSeenDate()`

UnsetEventFirstSeenDate ensures that no value is present for EventFirstSeenDate, not even an explicit nil
### GetEventHost

`func (o *Event) GetEventHost() string`

GetEventHost returns the EventHost field if non-nil, zero value otherwise.

### GetEventHostOk

`func (o *Event) GetEventHostOk() (*string, bool)`

GetEventHostOk returns a tuple with the EventHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHost

`func (o *Event) SetEventHost(v string)`

SetEventHost sets EventHost field to given value.


### SetEventHostNil

`func (o *Event) SetEventHostNil(b bool)`

 SetEventHostNil sets the value for EventHost to be an explicit nil

### UnsetEventHost
`func (o *Event) UnsetEventHost()`

UnsetEventHost ensures that no value is present for EventHost, not even an explicit nil
### GetEventKind

`func (o *Event) GetEventKind() string`

GetEventKind returns the EventKind field if non-nil, zero value otherwise.

### GetEventKindOk

`func (o *Event) GetEventKindOk() (*string, bool)`

GetEventKindOk returns a tuple with the EventKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKind

`func (o *Event) SetEventKind(v string)`

SetEventKind sets EventKind field to given value.


### GetEventLastSeenDate

`func (o *Event) GetEventLastSeenDate() time.Time`

GetEventLastSeenDate returns the EventLastSeenDate field if non-nil, zero value otherwise.

### GetEventLastSeenDateOk

`func (o *Event) GetEventLastSeenDateOk() (*time.Time, bool)`

GetEventLastSeenDateOk returns a tuple with the EventLastSeenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLastSeenDate

`func (o *Event) SetEventLastSeenDate(v time.Time)`

SetEventLastSeenDate sets EventLastSeenDate field to given value.


### SetEventLastSeenDateNil

`func (o *Event) SetEventLastSeenDateNil(b bool)`

 SetEventLastSeenDateNil sets the value for EventLastSeenDate to be an explicit nil

### UnsetEventLastSeenDate
`func (o *Event) UnsetEventLastSeenDate()`

UnsetEventLastSeenDate ensures that no value is present for EventLastSeenDate, not even an explicit nil
### GetEventMessage

`func (o *Event) GetEventMessage() string`

GetEventMessage returns the EventMessage field if non-nil, zero value otherwise.

### GetEventMessageOk

`func (o *Event) GetEventMessageOk() (*string, bool)`

GetEventMessageOk returns a tuple with the EventMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMessage

`func (o *Event) SetEventMessage(v string)`

SetEventMessage sets EventMessage field to given value.


### SetEventMessageNil

`func (o *Event) SetEventMessageNil(b bool)`

 SetEventMessageNil sets the value for EventMessage to be an explicit nil

### UnsetEventMessage
`func (o *Event) UnsetEventMessage()`

UnsetEventMessage ensures that no value is present for EventMessage, not even an explicit nil
### GetEventReason

`func (o *Event) GetEventReason() string`

GetEventReason returns the EventReason field if non-nil, zero value otherwise.

### GetEventReasonOk

`func (o *Event) GetEventReasonOk() (*string, bool)`

GetEventReasonOk returns a tuple with the EventReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReason

`func (o *Event) SetEventReason(v string)`

SetEventReason sets EventReason field to given value.


### SetEventReasonNil

`func (o *Event) SetEventReasonNil(b bool)`

 SetEventReasonNil sets the value for EventReason to be an explicit nil

### UnsetEventReason
`func (o *Event) UnsetEventReason()`

UnsetEventReason ensures that no value is present for EventReason, not even an explicit nil
### GetEventType

`func (o *Event) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Event) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Event) SetEventType(v string)`

SetEventType sets EventType field to given value.


### SetEventTypeNil

`func (o *Event) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *Event) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetName

`func (o *Event) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Event) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Event) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *Event) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *Event) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *Event) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### SetNamespaceNameNil

`func (o *Event) SetNamespaceNameNil(b bool)`

 SetNamespaceNameNil sets the value for NamespaceName to be an explicit nil

### UnsetNamespaceName
`func (o *Event) UnsetNamespaceName()`

UnsetNamespaceName ensures that no value is present for NamespaceName, not even an explicit nil
### GetUid

`func (o *Event) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Event) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Event) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


