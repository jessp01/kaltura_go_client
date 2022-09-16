# LiveStreamSetRecordedContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **float32** |  | 
**EntryId** | **string** |  | 
**FlavorParamsId** | Pointer to **int32** |  | [optional] 
**MediaServerIndex** | **string** |  | 
**RecordedEntryId** | Pointer to **string** |  | [optional] 
**Resource** | [**KalturaDataCenterContentResource**](KalturaDataCenterContentResource.md) |  | 

## Methods

### NewLiveStreamSetRecordedContentRequest

`func NewLiveStreamSetRecordedContentRequest(duration float32, entryId string, mediaServerIndex string, resource KalturaDataCenterContentResource, ) *LiveStreamSetRecordedContentRequest`

NewLiveStreamSetRecordedContentRequest instantiates a new LiveStreamSetRecordedContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamSetRecordedContentRequestWithDefaults

`func NewLiveStreamSetRecordedContentRequestWithDefaults() *LiveStreamSetRecordedContentRequest`

NewLiveStreamSetRecordedContentRequestWithDefaults instantiates a new LiveStreamSetRecordedContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *LiveStreamSetRecordedContentRequest) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *LiveStreamSetRecordedContentRequest) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *LiveStreamSetRecordedContentRequest) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetEntryId

`func (o *LiveStreamSetRecordedContentRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *LiveStreamSetRecordedContentRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *LiveStreamSetRecordedContentRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetFlavorParamsId

`func (o *LiveStreamSetRecordedContentRequest) GetFlavorParamsId() int32`

GetFlavorParamsId returns the FlavorParamsId field if non-nil, zero value otherwise.

### GetFlavorParamsIdOk

`func (o *LiveStreamSetRecordedContentRequest) GetFlavorParamsIdOk() (*int32, bool)`

GetFlavorParamsIdOk returns a tuple with the FlavorParamsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorParamsId

`func (o *LiveStreamSetRecordedContentRequest) SetFlavorParamsId(v int32)`

SetFlavorParamsId sets FlavorParamsId field to given value.

### HasFlavorParamsId

`func (o *LiveStreamSetRecordedContentRequest) HasFlavorParamsId() bool`

HasFlavorParamsId returns a boolean if a field has been set.

### GetMediaServerIndex

`func (o *LiveStreamSetRecordedContentRequest) GetMediaServerIndex() string`

GetMediaServerIndex returns the MediaServerIndex field if non-nil, zero value otherwise.

### GetMediaServerIndexOk

`func (o *LiveStreamSetRecordedContentRequest) GetMediaServerIndexOk() (*string, bool)`

GetMediaServerIndexOk returns a tuple with the MediaServerIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaServerIndex

`func (o *LiveStreamSetRecordedContentRequest) SetMediaServerIndex(v string)`

SetMediaServerIndex sets MediaServerIndex field to given value.


### GetRecordedEntryId

`func (o *LiveStreamSetRecordedContentRequest) GetRecordedEntryId() string`

GetRecordedEntryId returns the RecordedEntryId field if non-nil, zero value otherwise.

### GetRecordedEntryIdOk

`func (o *LiveStreamSetRecordedContentRequest) GetRecordedEntryIdOk() (*string, bool)`

GetRecordedEntryIdOk returns a tuple with the RecordedEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedEntryId

`func (o *LiveStreamSetRecordedContentRequest) SetRecordedEntryId(v string)`

SetRecordedEntryId sets RecordedEntryId field to given value.

### HasRecordedEntryId

`func (o *LiveStreamSetRecordedContentRequest) HasRecordedEntryId() bool`

HasRecordedEntryId returns a boolean if a field has been set.

### GetResource

`func (o *LiveStreamSetRecordedContentRequest) GetResource() KalturaDataCenterContentResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *LiveStreamSetRecordedContentRequest) GetResourceOk() (*KalturaDataCenterContentResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *LiveStreamSetRecordedContentRequest) SetResource(v KalturaDataCenterContentResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


