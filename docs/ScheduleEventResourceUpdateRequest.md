# ScheduleEventResourceUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleEventId** | **int32** |  | 
**ScheduleEventResource** | [**KalturaScheduleEventResource**](KalturaScheduleEventResource.md) |  | 
**ScheduleResourceId** | **int32** |  | 

## Methods

### NewScheduleEventResourceUpdateRequest

`func NewScheduleEventResourceUpdateRequest(scheduleEventId int32, scheduleEventResource KalturaScheduleEventResource, scheduleResourceId int32, ) *ScheduleEventResourceUpdateRequest`

NewScheduleEventResourceUpdateRequest instantiates a new ScheduleEventResourceUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEventResourceUpdateRequestWithDefaults

`func NewScheduleEventResourceUpdateRequestWithDefaults() *ScheduleEventResourceUpdateRequest`

NewScheduleEventResourceUpdateRequestWithDefaults instantiates a new ScheduleEventResourceUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleEventId

`func (o *ScheduleEventResourceUpdateRequest) GetScheduleEventId() int32`

GetScheduleEventId returns the ScheduleEventId field if non-nil, zero value otherwise.

### GetScheduleEventIdOk

`func (o *ScheduleEventResourceUpdateRequest) GetScheduleEventIdOk() (*int32, bool)`

GetScheduleEventIdOk returns a tuple with the ScheduleEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEventId

`func (o *ScheduleEventResourceUpdateRequest) SetScheduleEventId(v int32)`

SetScheduleEventId sets ScheduleEventId field to given value.


### GetScheduleEventResource

`func (o *ScheduleEventResourceUpdateRequest) GetScheduleEventResource() KalturaScheduleEventResource`

GetScheduleEventResource returns the ScheduleEventResource field if non-nil, zero value otherwise.

### GetScheduleEventResourceOk

`func (o *ScheduleEventResourceUpdateRequest) GetScheduleEventResourceOk() (*KalturaScheduleEventResource, bool)`

GetScheduleEventResourceOk returns a tuple with the ScheduleEventResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEventResource

`func (o *ScheduleEventResourceUpdateRequest) SetScheduleEventResource(v KalturaScheduleEventResource)`

SetScheduleEventResource sets ScheduleEventResource field to given value.


### GetScheduleResourceId

`func (o *ScheduleEventResourceUpdateRequest) GetScheduleResourceId() int32`

GetScheduleResourceId returns the ScheduleResourceId field if non-nil, zero value otherwise.

### GetScheduleResourceIdOk

`func (o *ScheduleEventResourceUpdateRequest) GetScheduleResourceIdOk() (*int32, bool)`

GetScheduleResourceIdOk returns a tuple with the ScheduleResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleResourceId

`func (o *ScheduleEventResourceUpdateRequest) SetScheduleResourceId(v int32)`

SetScheduleResourceId sets ScheduleResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


