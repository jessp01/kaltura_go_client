# KalturaLinkedScheduleEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **int32** | The id of the event influencing the start of the event holding this object | [optional] 
**Offset** | Pointer to **int32** | The time between the end of the event which it&#39;s id is in $eventId and the start of the event holding this object | [optional] 

## Methods

### NewKalturaLinkedScheduleEvent

`func NewKalturaLinkedScheduleEvent() *KalturaLinkedScheduleEvent`

NewKalturaLinkedScheduleEvent instantiates a new KalturaLinkedScheduleEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaLinkedScheduleEventWithDefaults

`func NewKalturaLinkedScheduleEventWithDefaults() *KalturaLinkedScheduleEvent`

NewKalturaLinkedScheduleEventWithDefaults instantiates a new KalturaLinkedScheduleEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *KalturaLinkedScheduleEvent) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *KalturaLinkedScheduleEvent) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *KalturaLinkedScheduleEvent) SetEventId(v int32)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *KalturaLinkedScheduleEvent) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetOffset

`func (o *KalturaLinkedScheduleEvent) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *KalturaLinkedScheduleEvent) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *KalturaLinkedScheduleEvent) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *KalturaLinkedScheduleEvent) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


