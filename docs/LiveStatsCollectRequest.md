# LiveStatsCollectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**KalturaLiveStatsEvent**](KalturaLiveStatsEvent.md) |  | 

## Methods

### NewLiveStatsCollectRequest

`func NewLiveStatsCollectRequest(event KalturaLiveStatsEvent, ) *LiveStatsCollectRequest`

NewLiveStatsCollectRequest instantiates a new LiveStatsCollectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStatsCollectRequestWithDefaults

`func NewLiveStatsCollectRequestWithDefaults() *LiveStatsCollectRequest`

NewLiveStatsCollectRequestWithDefaults instantiates a new LiveStatsCollectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *LiveStatsCollectRequest) GetEvent() KalturaLiveStatsEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *LiveStatsCollectRequest) GetEventOk() (*KalturaLiveStatsEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *LiveStatsCollectRequest) SetEvent(v KalturaLiveStatsEvent)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


