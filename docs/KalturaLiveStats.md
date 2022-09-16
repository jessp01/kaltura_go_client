# KalturaLiveStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **int32** |  | [optional] 
**AvgBitrate** | Pointer to **float32** |  | [optional] 
**BufferTime** | Pointer to **int32** |  | [optional] 
**DvrAudience** | Pointer to **int32** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Plays** | Pointer to **int32** |  | [optional] 
**SecondsViewed** | Pointer to **int32** |  | [optional] 
**StartEvent** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 

## Methods

### NewKalturaLiveStats

`func NewKalturaLiveStats() *KalturaLiveStats`

NewKalturaLiveStats instantiates a new KalturaLiveStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaLiveStatsWithDefaults

`func NewKalturaLiveStatsWithDefaults() *KalturaLiveStats`

NewKalturaLiveStatsWithDefaults instantiates a new KalturaLiveStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *KalturaLiveStats) GetAudience() int32`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *KalturaLiveStats) GetAudienceOk() (*int32, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *KalturaLiveStats) SetAudience(v int32)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *KalturaLiveStats) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetAvgBitrate

`func (o *KalturaLiveStats) GetAvgBitrate() float32`

GetAvgBitrate returns the AvgBitrate field if non-nil, zero value otherwise.

### GetAvgBitrateOk

`func (o *KalturaLiveStats) GetAvgBitrateOk() (*float32, bool)`

GetAvgBitrateOk returns a tuple with the AvgBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgBitrate

`func (o *KalturaLiveStats) SetAvgBitrate(v float32)`

SetAvgBitrate sets AvgBitrate field to given value.

### HasAvgBitrate

`func (o *KalturaLiveStats) HasAvgBitrate() bool`

HasAvgBitrate returns a boolean if a field has been set.

### GetBufferTime

`func (o *KalturaLiveStats) GetBufferTime() int32`

GetBufferTime returns the BufferTime field if non-nil, zero value otherwise.

### GetBufferTimeOk

`func (o *KalturaLiveStats) GetBufferTimeOk() (*int32, bool)`

GetBufferTimeOk returns a tuple with the BufferTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferTime

`func (o *KalturaLiveStats) SetBufferTime(v int32)`

SetBufferTime sets BufferTime field to given value.

### HasBufferTime

`func (o *KalturaLiveStats) HasBufferTime() bool`

HasBufferTime returns a boolean if a field has been set.

### GetDvrAudience

`func (o *KalturaLiveStats) GetDvrAudience() int32`

GetDvrAudience returns the DvrAudience field if non-nil, zero value otherwise.

### GetDvrAudienceOk

`func (o *KalturaLiveStats) GetDvrAudienceOk() (*int32, bool)`

GetDvrAudienceOk returns a tuple with the DvrAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvrAudience

`func (o *KalturaLiveStats) SetDvrAudience(v int32)`

SetDvrAudience sets DvrAudience field to given value.

### HasDvrAudience

`func (o *KalturaLiveStats) HasDvrAudience() bool`

HasDvrAudience returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaLiveStats) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaLiveStats) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaLiveStats) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaLiveStats) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPlays

`func (o *KalturaLiveStats) GetPlays() int32`

GetPlays returns the Plays field if non-nil, zero value otherwise.

### GetPlaysOk

`func (o *KalturaLiveStats) GetPlaysOk() (*int32, bool)`

GetPlaysOk returns a tuple with the Plays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlays

`func (o *KalturaLiveStats) SetPlays(v int32)`

SetPlays sets Plays field to given value.

### HasPlays

`func (o *KalturaLiveStats) HasPlays() bool`

HasPlays returns a boolean if a field has been set.

### GetSecondsViewed

`func (o *KalturaLiveStats) GetSecondsViewed() int32`

GetSecondsViewed returns the SecondsViewed field if non-nil, zero value otherwise.

### GetSecondsViewedOk

`func (o *KalturaLiveStats) GetSecondsViewedOk() (*int32, bool)`

GetSecondsViewedOk returns a tuple with the SecondsViewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsViewed

`func (o *KalturaLiveStats) SetSecondsViewed(v int32)`

SetSecondsViewed sets SecondsViewed field to given value.

### HasSecondsViewed

`func (o *KalturaLiveStats) HasSecondsViewed() bool`

HasSecondsViewed returns a boolean if a field has been set.

### GetStartEvent

`func (o *KalturaLiveStats) GetStartEvent() int32`

GetStartEvent returns the StartEvent field if non-nil, zero value otherwise.

### GetStartEventOk

`func (o *KalturaLiveStats) GetStartEventOk() (*int32, bool)`

GetStartEventOk returns a tuple with the StartEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartEvent

`func (o *KalturaLiveStats) SetStartEvent(v int32)`

SetStartEvent sets StartEvent field to given value.

### HasStartEvent

`func (o *KalturaLiveStats) HasStartEvent() bool`

HasStartEvent returns a boolean if a field has been set.

### GetTimestamp

`func (o *KalturaLiveStats) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *KalturaLiveStats) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *KalturaLiveStats) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *KalturaLiveStats) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


