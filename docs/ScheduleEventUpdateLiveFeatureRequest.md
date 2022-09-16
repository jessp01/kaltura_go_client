# ScheduleEventUpdateLiveFeatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureName** | **string** |  | 
**LiveFeature** | [**KalturaLiveFeature**](KalturaLiveFeature.md) |  | 
**ScheduledEventId** | **int32** |  | 

## Methods

### NewScheduleEventUpdateLiveFeatureRequest

`func NewScheduleEventUpdateLiveFeatureRequest(featureName string, liveFeature KalturaLiveFeature, scheduledEventId int32, ) *ScheduleEventUpdateLiveFeatureRequest`

NewScheduleEventUpdateLiveFeatureRequest instantiates a new ScheduleEventUpdateLiveFeatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEventUpdateLiveFeatureRequestWithDefaults

`func NewScheduleEventUpdateLiveFeatureRequestWithDefaults() *ScheduleEventUpdateLiveFeatureRequest`

NewScheduleEventUpdateLiveFeatureRequestWithDefaults instantiates a new ScheduleEventUpdateLiveFeatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureName

`func (o *ScheduleEventUpdateLiveFeatureRequest) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *ScheduleEventUpdateLiveFeatureRequest) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *ScheduleEventUpdateLiveFeatureRequest) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.


### GetLiveFeature

`func (o *ScheduleEventUpdateLiveFeatureRequest) GetLiveFeature() KalturaLiveFeature`

GetLiveFeature returns the LiveFeature field if non-nil, zero value otherwise.

### GetLiveFeatureOk

`func (o *ScheduleEventUpdateLiveFeatureRequest) GetLiveFeatureOk() (*KalturaLiveFeature, bool)`

GetLiveFeatureOk returns a tuple with the LiveFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveFeature

`func (o *ScheduleEventUpdateLiveFeatureRequest) SetLiveFeature(v KalturaLiveFeature)`

SetLiveFeature sets LiveFeature field to given value.


### GetScheduledEventId

`func (o *ScheduleEventUpdateLiveFeatureRequest) GetScheduledEventId() int32`

GetScheduledEventId returns the ScheduledEventId field if non-nil, zero value otherwise.

### GetScheduledEventIdOk

`func (o *ScheduleEventUpdateLiveFeatureRequest) GetScheduledEventIdOk() (*int32, bool)`

GetScheduledEventIdOk returns a tuple with the ScheduledEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEventId

`func (o *ScheduleEventUpdateLiveFeatureRequest) SetScheduledEventId(v int32)`

SetScheduledEventId sets ScheduledEventId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


