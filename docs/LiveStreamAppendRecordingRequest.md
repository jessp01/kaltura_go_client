# LiveStreamAppendRecordingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** |  | 
**Duration** | **float32** |  | 
**EntryId** | **string** |  | 
**IsLastChunk** | Pointer to **bool** |  | [optional] [default to false]
**MediaServerIndex** | **string** |  | 
**Resource** | [**KalturaDataCenterContentResource**](KalturaDataCenterContentResource.md) |  | 

## Methods

### NewLiveStreamAppendRecordingRequest

`func NewLiveStreamAppendRecordingRequest(assetId string, duration float32, entryId string, mediaServerIndex string, resource KalturaDataCenterContentResource, ) *LiveStreamAppendRecordingRequest`

NewLiveStreamAppendRecordingRequest instantiates a new LiveStreamAppendRecordingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamAppendRecordingRequestWithDefaults

`func NewLiveStreamAppendRecordingRequestWithDefaults() *LiveStreamAppendRecordingRequest`

NewLiveStreamAppendRecordingRequestWithDefaults instantiates a new LiveStreamAppendRecordingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *LiveStreamAppendRecordingRequest) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *LiveStreamAppendRecordingRequest) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *LiveStreamAppendRecordingRequest) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetDuration

`func (o *LiveStreamAppendRecordingRequest) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *LiveStreamAppendRecordingRequest) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *LiveStreamAppendRecordingRequest) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetEntryId

`func (o *LiveStreamAppendRecordingRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *LiveStreamAppendRecordingRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *LiveStreamAppendRecordingRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetIsLastChunk

`func (o *LiveStreamAppendRecordingRequest) GetIsLastChunk() bool`

GetIsLastChunk returns the IsLastChunk field if non-nil, zero value otherwise.

### GetIsLastChunkOk

`func (o *LiveStreamAppendRecordingRequest) GetIsLastChunkOk() (*bool, bool)`

GetIsLastChunkOk returns a tuple with the IsLastChunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastChunk

`func (o *LiveStreamAppendRecordingRequest) SetIsLastChunk(v bool)`

SetIsLastChunk sets IsLastChunk field to given value.

### HasIsLastChunk

`func (o *LiveStreamAppendRecordingRequest) HasIsLastChunk() bool`

HasIsLastChunk returns a boolean if a field has been set.

### GetMediaServerIndex

`func (o *LiveStreamAppendRecordingRequest) GetMediaServerIndex() string`

GetMediaServerIndex returns the MediaServerIndex field if non-nil, zero value otherwise.

### GetMediaServerIndexOk

`func (o *LiveStreamAppendRecordingRequest) GetMediaServerIndexOk() (*string, bool)`

GetMediaServerIndexOk returns a tuple with the MediaServerIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaServerIndex

`func (o *LiveStreamAppendRecordingRequest) SetMediaServerIndex(v string)`

SetMediaServerIndex sets MediaServerIndex field to given value.


### GetResource

`func (o *LiveStreamAppendRecordingRequest) GetResource() KalturaDataCenterContentResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *LiveStreamAppendRecordingRequest) GetResourceOk() (*KalturaDataCenterContentResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *LiveStreamAppendRecordingRequest) SetResource(v KalturaDataCenterContentResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


