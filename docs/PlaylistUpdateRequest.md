# PlaylistUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Playlist** | [**KalturaPlaylist**](KalturaPlaylist.md) |  | 
**UpdateStats** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPlaylistUpdateRequest

`func NewPlaylistUpdateRequest(id string, playlist KalturaPlaylist, ) *PlaylistUpdateRequest`

NewPlaylistUpdateRequest instantiates a new PlaylistUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistUpdateRequestWithDefaults

`func NewPlaylistUpdateRequestWithDefaults() *PlaylistUpdateRequest`

NewPlaylistUpdateRequestWithDefaults instantiates a new PlaylistUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlaylistUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaylistUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaylistUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPlaylist

`func (o *PlaylistUpdateRequest) GetPlaylist() KalturaPlaylist`

GetPlaylist returns the Playlist field if non-nil, zero value otherwise.

### GetPlaylistOk

`func (o *PlaylistUpdateRequest) GetPlaylistOk() (*KalturaPlaylist, bool)`

GetPlaylistOk returns a tuple with the Playlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylist

`func (o *PlaylistUpdateRequest) SetPlaylist(v KalturaPlaylist)`

SetPlaylist sets Playlist field to given value.


### GetUpdateStats

`func (o *PlaylistUpdateRequest) GetUpdateStats() bool`

GetUpdateStats returns the UpdateStats field if non-nil, zero value otherwise.

### GetUpdateStatsOk

`func (o *PlaylistUpdateRequest) GetUpdateStatsOk() (*bool, bool)`

GetUpdateStatsOk returns a tuple with the UpdateStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStats

`func (o *PlaylistUpdateRequest) SetUpdateStats(v bool)`

SetUpdateStats sets UpdateStats field to given value.

### HasUpdateStats

`func (o *PlaylistUpdateRequest) HasUpdateStats() bool`

HasUpdateStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


