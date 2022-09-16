# MediaUpdateThumbnailFromSourceEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**FlavorParamsId** | Pointer to **int32** |  | [optional] 
**SourceEntryId** | **string** |  | 
**TimeOffset** | **int32** |  | 

## Methods

### NewMediaUpdateThumbnailFromSourceEntryRequest

`func NewMediaUpdateThumbnailFromSourceEntryRequest(entryId string, sourceEntryId string, timeOffset int32, ) *MediaUpdateThumbnailFromSourceEntryRequest`

NewMediaUpdateThumbnailFromSourceEntryRequest instantiates a new MediaUpdateThumbnailFromSourceEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaUpdateThumbnailFromSourceEntryRequestWithDefaults

`func NewMediaUpdateThumbnailFromSourceEntryRequestWithDefaults() *MediaUpdateThumbnailFromSourceEntryRequest`

NewMediaUpdateThumbnailFromSourceEntryRequestWithDefaults instantiates a new MediaUpdateThumbnailFromSourceEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetFlavorParamsId

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) GetFlavorParamsId() int32`

GetFlavorParamsId returns the FlavorParamsId field if non-nil, zero value otherwise.

### GetFlavorParamsIdOk

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) GetFlavorParamsIdOk() (*int32, bool)`

GetFlavorParamsIdOk returns a tuple with the FlavorParamsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorParamsId

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) SetFlavorParamsId(v int32)`

SetFlavorParamsId sets FlavorParamsId field to given value.

### HasFlavorParamsId

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) HasFlavorParamsId() bool`

HasFlavorParamsId returns a boolean if a field has been set.

### GetSourceEntryId

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) GetSourceEntryId() string`

GetSourceEntryId returns the SourceEntryId field if non-nil, zero value otherwise.

### GetSourceEntryIdOk

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) GetSourceEntryIdOk() (*string, bool)`

GetSourceEntryIdOk returns a tuple with the SourceEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntryId

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) SetSourceEntryId(v string)`

SetSourceEntryId sets SourceEntryId field to given value.


### GetTimeOffset

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) GetTimeOffset() int32`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) GetTimeOffsetOk() (*int32, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *MediaUpdateThumbnailFromSourceEntryRequest) SetTimeOffset(v int32)`

SetTimeOffset sets TimeOffset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


