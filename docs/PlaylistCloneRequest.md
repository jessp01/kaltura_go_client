# PlaylistCloneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**NewPlaylist** | Pointer to [**KalturaPlaylist**](KalturaPlaylist.md) |  | [optional] 

## Methods

### NewPlaylistCloneRequest

`func NewPlaylistCloneRequest(id string, ) *PlaylistCloneRequest`

NewPlaylistCloneRequest instantiates a new PlaylistCloneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistCloneRequestWithDefaults

`func NewPlaylistCloneRequestWithDefaults() *PlaylistCloneRequest`

NewPlaylistCloneRequestWithDefaults instantiates a new PlaylistCloneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlaylistCloneRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaylistCloneRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaylistCloneRequest) SetId(v string)`

SetId sets Id field to given value.


### GetNewPlaylist

`func (o *PlaylistCloneRequest) GetNewPlaylist() KalturaPlaylist`

GetNewPlaylist returns the NewPlaylist field if non-nil, zero value otherwise.

### GetNewPlaylistOk

`func (o *PlaylistCloneRequest) GetNewPlaylistOk() (*KalturaPlaylist, bool)`

GetNewPlaylistOk returns a tuple with the NewPlaylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlaylist

`func (o *PlaylistCloneRequest) SetNewPlaylist(v KalturaPlaylist)`

SetNewPlaylist sets NewPlaylist field to given value.

### HasNewPlaylist

`func (o *PlaylistCloneRequest) HasNewPlaylist() bool`

HasNewPlaylist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


