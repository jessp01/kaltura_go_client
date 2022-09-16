# PlaylistAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Playlist** | [**KalturaPlaylist**](KalturaPlaylist.md) |  | 
**UpdateStats** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPlaylistAddRequest

`func NewPlaylistAddRequest(playlist KalturaPlaylist, ) *PlaylistAddRequest`

NewPlaylistAddRequest instantiates a new PlaylistAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistAddRequestWithDefaults

`func NewPlaylistAddRequestWithDefaults() *PlaylistAddRequest`

NewPlaylistAddRequestWithDefaults instantiates a new PlaylistAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaylist

`func (o *PlaylistAddRequest) GetPlaylist() KalturaPlaylist`

GetPlaylist returns the Playlist field if non-nil, zero value otherwise.

### GetPlaylistOk

`func (o *PlaylistAddRequest) GetPlaylistOk() (*KalturaPlaylist, bool)`

GetPlaylistOk returns a tuple with the Playlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylist

`func (o *PlaylistAddRequest) SetPlaylist(v KalturaPlaylist)`

SetPlaylist sets Playlist field to given value.


### GetUpdateStats

`func (o *PlaylistAddRequest) GetUpdateStats() bool`

GetUpdateStats returns the UpdateStats field if non-nil, zero value otherwise.

### GetUpdateStatsOk

`func (o *PlaylistAddRequest) GetUpdateStatsOk() (*bool, bool)`

GetUpdateStatsOk returns a tuple with the UpdateStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStats

`func (o *PlaylistAddRequest) SetUpdateStats(v bool)`

SetUpdateStats sets UpdateStats field to given value.

### HasUpdateStats

`func (o *PlaylistAddRequest) HasUpdateStats() bool`

HasUpdateStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


