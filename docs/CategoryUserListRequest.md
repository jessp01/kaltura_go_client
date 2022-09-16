# CategoryUserListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaCategoryUserFilter**](KalturaCategoryUserFilter.md) |  | [optional] 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 

## Methods

### NewCategoryUserListRequest

`func NewCategoryUserListRequest() *CategoryUserListRequest`

NewCategoryUserListRequest instantiates a new CategoryUserListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryUserListRequestWithDefaults

`func NewCategoryUserListRequestWithDefaults() *CategoryUserListRequest`

NewCategoryUserListRequestWithDefaults instantiates a new CategoryUserListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *CategoryUserListRequest) GetFilter() KalturaCategoryUserFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CategoryUserListRequest) GetFilterOk() (*KalturaCategoryUserFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CategoryUserListRequest) SetFilter(v KalturaCategoryUserFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CategoryUserListRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPager

`func (o *CategoryUserListRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *CategoryUserListRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *CategoryUserListRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *CategoryUserListRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


