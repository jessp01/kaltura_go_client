# PlaylistExecuteFromContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detailed** | Pointer to **string** |  | [optional] 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 
**PlaylistContent** | **string** |  | 
**PlaylistType** | **int32** |  | 

## Methods

### NewPlaylistExecuteFromContentRequest

`func NewPlaylistExecuteFromContentRequest(playlistContent string, playlistType int32, ) *PlaylistExecuteFromContentRequest`

NewPlaylistExecuteFromContentRequest instantiates a new PlaylistExecuteFromContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistExecuteFromContentRequestWithDefaults

`func NewPlaylistExecuteFromContentRequestWithDefaults() *PlaylistExecuteFromContentRequest`

NewPlaylistExecuteFromContentRequestWithDefaults instantiates a new PlaylistExecuteFromContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetailed

`func (o *PlaylistExecuteFromContentRequest) GetDetailed() string`

GetDetailed returns the Detailed field if non-nil, zero value otherwise.

### GetDetailedOk

`func (o *PlaylistExecuteFromContentRequest) GetDetailedOk() (*string, bool)`

GetDetailedOk returns a tuple with the Detailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailed

`func (o *PlaylistExecuteFromContentRequest) SetDetailed(v string)`

SetDetailed sets Detailed field to given value.

### HasDetailed

`func (o *PlaylistExecuteFromContentRequest) HasDetailed() bool`

HasDetailed returns a boolean if a field has been set.

### GetPager

`func (o *PlaylistExecuteFromContentRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *PlaylistExecuteFromContentRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *PlaylistExecuteFromContentRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *PlaylistExecuteFromContentRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetPlaylistContent

`func (o *PlaylistExecuteFromContentRequest) GetPlaylistContent() string`

GetPlaylistContent returns the PlaylistContent field if non-nil, zero value otherwise.

### GetPlaylistContentOk

`func (o *PlaylistExecuteFromContentRequest) GetPlaylistContentOk() (*string, bool)`

GetPlaylistContentOk returns a tuple with the PlaylistContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistContent

`func (o *PlaylistExecuteFromContentRequest) SetPlaylistContent(v string)`

SetPlaylistContent sets PlaylistContent field to given value.


### GetPlaylistType

`func (o *PlaylistExecuteFromContentRequest) GetPlaylistType() int32`

GetPlaylistType returns the PlaylistType field if non-nil, zero value otherwise.

### GetPlaylistTypeOk

`func (o *PlaylistExecuteFromContentRequest) GetPlaylistTypeOk() (*int32, bool)`

GetPlaylistTypeOk returns a tuple with the PlaylistType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistType

`func (o *PlaylistExecuteFromContentRequest) SetPlaylistType(v int32)`

SetPlaylistType sets PlaylistType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


