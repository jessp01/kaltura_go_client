# MediaAddFromSearchResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaEntry** | Pointer to [**KalturaMediaEntry**](KalturaMediaEntry.md) |  | [optional] 
**SearchResult** | Pointer to [**KalturaSearchResult**](KalturaSearchResult.md) |  | [optional] 

## Methods

### NewMediaAddFromSearchResultRequest

`func NewMediaAddFromSearchResultRequest() *MediaAddFromSearchResultRequest`

NewMediaAddFromSearchResultRequest instantiates a new MediaAddFromSearchResultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaAddFromSearchResultRequestWithDefaults

`func NewMediaAddFromSearchResultRequestWithDefaults() *MediaAddFromSearchResultRequest`

NewMediaAddFromSearchResultRequestWithDefaults instantiates a new MediaAddFromSearchResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaEntry

`func (o *MediaAddFromSearchResultRequest) GetMediaEntry() KalturaMediaEntry`

GetMediaEntry returns the MediaEntry field if non-nil, zero value otherwise.

### GetMediaEntryOk

`func (o *MediaAddFromSearchResultRequest) GetMediaEntryOk() (*KalturaMediaEntry, bool)`

GetMediaEntryOk returns a tuple with the MediaEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEntry

`func (o *MediaAddFromSearchResultRequest) SetMediaEntry(v KalturaMediaEntry)`

SetMediaEntry sets MediaEntry field to given value.

### HasMediaEntry

`func (o *MediaAddFromSearchResultRequest) HasMediaEntry() bool`

HasMediaEntry returns a boolean if a field has been set.

### GetSearchResult

`func (o *MediaAddFromSearchResultRequest) GetSearchResult() KalturaSearchResult`

GetSearchResult returns the SearchResult field if non-nil, zero value otherwise.

### GetSearchResultOk

`func (o *MediaAddFromSearchResultRequest) GetSearchResultOk() (*KalturaSearchResult, bool)`

GetSearchResultOk returns a tuple with the SearchResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResult

`func (o *MediaAddFromSearchResultRequest) SetSearchResult(v KalturaSearchResult)`

SetSearchResult sets SearchResult field to given value.

### HasSearchResult

`func (o *MediaAddFromSearchResultRequest) HasSearchResult() bool`

HasSearchResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


