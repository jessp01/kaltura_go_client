# TagSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 
**TagFilter** | [**KalturaTagFilter**](KalturaTagFilter.md) |  | 

## Methods

### NewTagSearchRequest

`func NewTagSearchRequest(tagFilter KalturaTagFilter, ) *TagSearchRequest`

NewTagSearchRequest instantiates a new TagSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagSearchRequestWithDefaults

`func NewTagSearchRequestWithDefaults() *TagSearchRequest`

NewTagSearchRequestWithDefaults instantiates a new TagSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPager

`func (o *TagSearchRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *TagSearchRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *TagSearchRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *TagSearchRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetTagFilter

`func (o *TagSearchRequest) GetTagFilter() KalturaTagFilter`

GetTagFilter returns the TagFilter field if non-nil, zero value otherwise.

### GetTagFilterOk

`func (o *TagSearchRequest) GetTagFilterOk() (*KalturaTagFilter, bool)`

GetTagFilterOk returns a tuple with the TagFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilter

`func (o *TagSearchRequest) SetTagFilter(v KalturaTagFilter)`

SetTagFilter sets TagFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


