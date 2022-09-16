# ScheduleEventUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleEvent** | [**KalturaScheduleEvent**](KalturaScheduleEvent.md) |  | 
**ScheduleEventId** | **int32** |  | 

## Methods

### NewScheduleEventUpdateRequest

`func NewScheduleEventUpdateRequest(scheduleEvent KalturaScheduleEvent, scheduleEventId int32, ) *ScheduleEventUpdateRequest`

NewScheduleEventUpdateRequest instantiates a new ScheduleEventUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEventUpdateRequestWithDefaults

`func NewScheduleEventUpdateRequestWithDefaults() *ScheduleEventUpdateRequest`

NewScheduleEventUpdateRequestWithDefaults instantiates a new ScheduleEventUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleEvent

`func (o *ScheduleEventUpdateRequest) GetScheduleEvent() KalturaScheduleEvent`

GetScheduleEvent returns the ScheduleEvent field if non-nil, zero value otherwise.

### GetScheduleEventOk

`func (o *ScheduleEventUpdateRequest) GetScheduleEventOk() (*KalturaScheduleEvent, bool)`

GetScheduleEventOk returns a tuple with the ScheduleEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEvent

`func (o *ScheduleEventUpdateRequest) SetScheduleEvent(v KalturaScheduleEvent)`

SetScheduleEvent sets ScheduleEvent field to given value.


### GetScheduleEventId

`func (o *ScheduleEventUpdateRequest) GetScheduleEventId() int32`

GetScheduleEventId returns the ScheduleEventId field if non-nil, zero value otherwise.

### GetScheduleEventIdOk

`func (o *ScheduleEventUpdateRequest) GetScheduleEventIdOk() (*int32, bool)`

GetScheduleEventIdOk returns a tuple with the ScheduleEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEventId

`func (o *ScheduleEventUpdateRequest) SetScheduleEventId(v int32)`

SetScheduleEventId sets ScheduleEventId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


