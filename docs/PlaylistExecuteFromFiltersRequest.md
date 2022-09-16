# PlaylistExecuteFromFiltersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detailed** | Pointer to **string** |  | [optional] [default to "1"]
**Filters** | [**[]KalturaMediaEntryFilterForPlaylist**](KalturaMediaEntryFilterForPlaylist.md) |  | 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 
**TotalResults** | **int32** |  | 

## Methods

### NewPlaylistExecuteFromFiltersRequest

`func NewPlaylistExecuteFromFiltersRequest(filters []KalturaMediaEntryFilterForPlaylist, totalResults int32, ) *PlaylistExecuteFromFiltersRequest`

NewPlaylistExecuteFromFiltersRequest instantiates a new PlaylistExecuteFromFiltersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistExecuteFromFiltersRequestWithDefaults

`func NewPlaylistExecuteFromFiltersRequestWithDefaults() *PlaylistExecuteFromFiltersRequest`

NewPlaylistExecuteFromFiltersRequestWithDefaults instantiates a new PlaylistExecuteFromFiltersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetailed

`func (o *PlaylistExecuteFromFiltersRequest) GetDetailed() string`

GetDetailed returns the Detailed field if non-nil, zero value otherwise.

### GetDetailedOk

`func (o *PlaylistExecuteFromFiltersRequest) GetDetailedOk() (*string, bool)`

GetDetailedOk returns a tuple with the Detailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailed

`func (o *PlaylistExecuteFromFiltersRequest) SetDetailed(v string)`

SetDetailed sets Detailed field to given value.

### HasDetailed

`func (o *PlaylistExecuteFromFiltersRequest) HasDetailed() bool`

HasDetailed returns a boolean if a field has been set.

### GetFilters

`func (o *PlaylistExecuteFromFiltersRequest) GetFilters() []KalturaMediaEntryFilterForPlaylist`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PlaylistExecuteFromFiltersRequest) GetFiltersOk() (*[]KalturaMediaEntryFilterForPlaylist, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PlaylistExecuteFromFiltersRequest) SetFilters(v []KalturaMediaEntryFilterForPlaylist)`

SetFilters sets Filters field to given value.


### GetPager

`func (o *PlaylistExecuteFromFiltersRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *PlaylistExecuteFromFiltersRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *PlaylistExecuteFromFiltersRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *PlaylistExecuteFromFiltersRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetTotalResults

`func (o *PlaylistExecuteFromFiltersRequest) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *PlaylistExecuteFromFiltersRequest) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *PlaylistExecuteFromFiltersRequest) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


