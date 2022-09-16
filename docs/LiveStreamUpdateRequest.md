# LiveStreamUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**LiveStreamEntry** | [**KalturaLiveStreamEntry**](KalturaLiveStreamEntry.md) |  | 

## Methods

### NewLiveStreamUpdateRequest

`func NewLiveStreamUpdateRequest(entryId string, liveStreamEntry KalturaLiveStreamEntry, ) *LiveStreamUpdateRequest`

NewLiveStreamUpdateRequest instantiates a new LiveStreamUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamUpdateRequestWithDefaults

`func NewLiveStreamUpdateRequestWithDefaults() *LiveStreamUpdateRequest`

NewLiveStreamUpdateRequestWithDefaults instantiates a new LiveStreamUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *LiveStreamUpdateRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *LiveStreamUpdateRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *LiveStreamUpdateRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetLiveStreamEntry

`func (o *LiveStreamUpdateRequest) GetLiveStreamEntry() KalturaLiveStreamEntry`

GetLiveStreamEntry returns the LiveStreamEntry field if non-nil, zero value otherwise.

### GetLiveStreamEntryOk

`func (o *LiveStreamUpdateRequest) GetLiveStreamEntryOk() (*KalturaLiveStreamEntry, bool)`

GetLiveStreamEntryOk returns a tuple with the LiveStreamEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamEntry

`func (o *LiveStreamUpdateRequest) SetLiveStreamEntry(v KalturaLiveStreamEntry)`

SetLiveStreamEntry sets LiveStreamEntry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


