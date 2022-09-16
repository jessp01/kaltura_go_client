# LiveStreamCreateRecordedEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**LiveEntryStatus** | **int32** |  | 
**MediaServerIndex** | **string** |  | 

## Methods

### NewLiveStreamCreateRecordedEntryRequest

`func NewLiveStreamCreateRecordedEntryRequest(entryId string, liveEntryStatus int32, mediaServerIndex string, ) *LiveStreamCreateRecordedEntryRequest`

NewLiveStreamCreateRecordedEntryRequest instantiates a new LiveStreamCreateRecordedEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamCreateRecordedEntryRequestWithDefaults

`func NewLiveStreamCreateRecordedEntryRequestWithDefaults() *LiveStreamCreateRecordedEntryRequest`

NewLiveStreamCreateRecordedEntryRequestWithDefaults instantiates a new LiveStreamCreateRecordedEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *LiveStreamCreateRecordedEntryRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *LiveStreamCreateRecordedEntryRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *LiveStreamCreateRecordedEntryRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetLiveEntryStatus

`func (o *LiveStreamCreateRecordedEntryRequest) GetLiveEntryStatus() int32`

GetLiveEntryStatus returns the LiveEntryStatus field if non-nil, zero value otherwise.

### GetLiveEntryStatusOk

`func (o *LiveStreamCreateRecordedEntryRequest) GetLiveEntryStatusOk() (*int32, bool)`

GetLiveEntryStatusOk returns a tuple with the LiveEntryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveEntryStatus

`func (o *LiveStreamCreateRecordedEntryRequest) SetLiveEntryStatus(v int32)`

SetLiveEntryStatus sets LiveEntryStatus field to given value.


### GetMediaServerIndex

`func (o *LiveStreamCreateRecordedEntryRequest) GetMediaServerIndex() string`

GetMediaServerIndex returns the MediaServerIndex field if non-nil, zero value otherwise.

### GetMediaServerIndexOk

`func (o *LiveStreamCreateRecordedEntryRequest) GetMediaServerIndexOk() (*string, bool)`

GetMediaServerIndexOk returns a tuple with the MediaServerIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaServerIndex

`func (o *LiveStreamCreateRecordedEntryRequest) SetMediaServerIndex(v string)`

SetMediaServerIndex sets MediaServerIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


