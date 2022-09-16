# PlaylistExecuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detailed** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**KalturaMediaEntryFilterForPlaylist**](KalturaMediaEntryFilterForPlaylist.md) |  | [optional] 
**Id** | **string** |  | 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 
**PlaylistContext** | Pointer to [**KalturaContext**](KalturaContext.md) |  | [optional] 

## Methods

### NewPlaylistExecuteRequest

`func NewPlaylistExecuteRequest(id string, ) *PlaylistExecuteRequest`

NewPlaylistExecuteRequest instantiates a new PlaylistExecuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistExecuteRequestWithDefaults

`func NewPlaylistExecuteRequestWithDefaults() *PlaylistExecuteRequest`

NewPlaylistExecuteRequestWithDefaults instantiates a new PlaylistExecuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetailed

`func (o *PlaylistExecuteRequest) GetDetailed() string`

GetDetailed returns the Detailed field if non-nil, zero value otherwise.

### GetDetailedOk

`func (o *PlaylistExecuteRequest) GetDetailedOk() (*string, bool)`

GetDetailedOk returns a tuple with the Detailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailed

`func (o *PlaylistExecuteRequest) SetDetailed(v string)`

SetDetailed sets Detailed field to given value.

### HasDetailed

`func (o *PlaylistExecuteRequest) HasDetailed() bool`

HasDetailed returns a boolean if a field has been set.

### GetFilter

`func (o *PlaylistExecuteRequest) GetFilter() KalturaMediaEntryFilterForPlaylist`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PlaylistExecuteRequest) GetFilterOk() (*KalturaMediaEntryFilterForPlaylist, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PlaylistExecuteRequest) SetFilter(v KalturaMediaEntryFilterForPlaylist)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PlaylistExecuteRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *PlaylistExecuteRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaylistExecuteRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaylistExecuteRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPager

`func (o *PlaylistExecuteRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *PlaylistExecuteRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *PlaylistExecuteRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *PlaylistExecuteRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetPlaylistContext

`func (o *PlaylistExecuteRequest) GetPlaylistContext() KalturaContext`

GetPlaylistContext returns the PlaylistContext field if non-nil, zero value otherwise.

### GetPlaylistContextOk

`func (o *PlaylistExecuteRequest) GetPlaylistContextOk() (*KalturaContext, bool)`

GetPlaylistContextOk returns a tuple with the PlaylistContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistContext

`func (o *PlaylistExecuteRequest) SetPlaylistContext(v KalturaContext)`

SetPlaylistContext sets PlaylistContext field to given value.

### HasPlaylistContext

`func (o *PlaylistExecuteRequest) HasPlaylistContext() bool`

HasPlaylistContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


