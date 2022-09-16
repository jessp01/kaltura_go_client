# ScheduleEventGetConflictsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceIds** | **string** |  | 
**ScheduleEvent** | [**KalturaScheduleEvent**](KalturaScheduleEvent.md) |  | 
**ScheduleEventConflictType** | Pointer to **int32** |  | [optional] 
**ScheduleEventIdToIgnore** | Pointer to **string** |  | [optional] 

## Methods

### NewScheduleEventGetConflictsRequest

`func NewScheduleEventGetConflictsRequest(resourceIds string, scheduleEvent KalturaScheduleEvent, ) *ScheduleEventGetConflictsRequest`

NewScheduleEventGetConflictsRequest instantiates a new ScheduleEventGetConflictsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEventGetConflictsRequestWithDefaults

`func NewScheduleEventGetConflictsRequestWithDefaults() *ScheduleEventGetConflictsRequest`

NewScheduleEventGetConflictsRequestWithDefaults instantiates a new ScheduleEventGetConflictsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceIds

`func (o *ScheduleEventGetConflictsRequest) GetResourceIds() string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *ScheduleEventGetConflictsRequest) GetResourceIdsOk() (*string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *ScheduleEventGetConflictsRequest) SetResourceIds(v string)`

SetResourceIds sets ResourceIds field to given value.


### GetScheduleEvent

`func (o *ScheduleEventGetConflictsRequest) GetScheduleEvent() KalturaScheduleEvent`

GetScheduleEvent returns the ScheduleEvent field if non-nil, zero value otherwise.

### GetScheduleEventOk

`func (o *ScheduleEventGetConflictsRequest) GetScheduleEventOk() (*KalturaScheduleEvent, bool)`

GetScheduleEventOk returns a tuple with the ScheduleEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEvent

`func (o *ScheduleEventGetConflictsRequest) SetScheduleEvent(v KalturaScheduleEvent)`

SetScheduleEvent sets ScheduleEvent field to given value.


### GetScheduleEventConflictType

`func (o *ScheduleEventGetConflictsRequest) GetScheduleEventConflictType() int32`

GetScheduleEventConflictType returns the ScheduleEventConflictType field if non-nil, zero value otherwise.

### GetScheduleEventConflictTypeOk

`func (o *ScheduleEventGetConflictsRequest) GetScheduleEventConflictTypeOk() (*int32, bool)`

GetScheduleEventConflictTypeOk returns a tuple with the ScheduleEventConflictType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEventConflictType

`func (o *ScheduleEventGetConflictsRequest) SetScheduleEventConflictType(v int32)`

SetScheduleEventConflictType sets ScheduleEventConflictType field to given value.

### HasScheduleEventConflictType

`func (o *ScheduleEventGetConflictsRequest) HasScheduleEventConflictType() bool`

HasScheduleEventConflictType returns a boolean if a field has been set.

### GetScheduleEventIdToIgnore

`func (o *ScheduleEventGetConflictsRequest) GetScheduleEventIdToIgnore() string`

GetScheduleEventIdToIgnore returns the ScheduleEventIdToIgnore field if non-nil, zero value otherwise.

### GetScheduleEventIdToIgnoreOk

`func (o *ScheduleEventGetConflictsRequest) GetScheduleEventIdToIgnoreOk() (*string, bool)`

GetScheduleEventIdToIgnoreOk returns a tuple with the ScheduleEventIdToIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEventIdToIgnore

`func (o *ScheduleEventGetConflictsRequest) SetScheduleEventIdToIgnore(v string)`

SetScheduleEventIdToIgnore sets ScheduleEventIdToIgnore field to given value.

### HasScheduleEventIdToIgnore

`func (o *ScheduleEventGetConflictsRequest) HasScheduleEventIdToIgnore() bool`

HasScheduleEventIdToIgnore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


