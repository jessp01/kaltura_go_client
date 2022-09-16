# MediaUpdateThumbnailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**FlavorParamsId** | Pointer to **int32** |  | [optional] 
**TimeOffset** | **int32** |  | 

## Methods

### NewMediaUpdateThumbnailRequest

`func NewMediaUpdateThumbnailRequest(entryId string, timeOffset int32, ) *MediaUpdateThumbnailRequest`

NewMediaUpdateThumbnailRequest instantiates a new MediaUpdateThumbnailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaUpdateThumbnailRequestWithDefaults

`func NewMediaUpdateThumbnailRequestWithDefaults() *MediaUpdateThumbnailRequest`

NewMediaUpdateThumbnailRequestWithDefaults instantiates a new MediaUpdateThumbnailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *MediaUpdateThumbnailRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *MediaUpdateThumbnailRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *MediaUpdateThumbnailRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetFlavorParamsId

`func (o *MediaUpdateThumbnailRequest) GetFlavorParamsId() int32`

GetFlavorParamsId returns the FlavorParamsId field if non-nil, zero value otherwise.

### GetFlavorParamsIdOk

`func (o *MediaUpdateThumbnailRequest) GetFlavorParamsIdOk() (*int32, bool)`

GetFlavorParamsIdOk returns a tuple with the FlavorParamsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorParamsId

`func (o *MediaUpdateThumbnailRequest) SetFlavorParamsId(v int32)`

SetFlavorParamsId sets FlavorParamsId field to given value.

### HasFlavorParamsId

`func (o *MediaUpdateThumbnailRequest) HasFlavorParamsId() bool`

HasFlavorParamsId returns a boolean if a field has been set.

### GetTimeOffset

`func (o *MediaUpdateThumbnailRequest) GetTimeOffset() int32`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *MediaUpdateThumbnailRequest) GetTimeOffsetOk() (*int32, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *MediaUpdateThumbnailRequest) SetTimeOffset(v int32)`

SetTimeOffset sets TimeOffset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


