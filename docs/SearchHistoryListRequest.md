# SearchHistoryListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaESearchHistoryFilter**](KalturaESearchHistoryFilter.md) |  | [optional] 

## Methods

### NewSearchHistoryListRequest

`func NewSearchHistoryListRequest() *SearchHistoryListRequest`

NewSearchHistoryListRequest instantiates a new SearchHistoryListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchHistoryListRequestWithDefaults

`func NewSearchHistoryListRequestWithDefaults() *SearchHistoryListRequest`

NewSearchHistoryListRequestWithDefaults instantiates a new SearchHistoryListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *SearchHistoryListRequest) GetFilter() KalturaESearchHistoryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SearchHistoryListRequest) GetFilterOk() (*KalturaESearchHistoryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SearchHistoryListRequest) SetFilter(v KalturaESearchHistoryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SearchHistoryListRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


