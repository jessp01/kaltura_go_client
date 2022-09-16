# EntryDistributionListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaEntryDistributionFilter**](KalturaEntryDistributionFilter.md) |  | [optional] 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 

## Methods

### NewEntryDistributionListRequest

`func NewEntryDistributionListRequest() *EntryDistributionListRequest`

NewEntryDistributionListRequest instantiates a new EntryDistributionListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryDistributionListRequestWithDefaults

`func NewEntryDistributionListRequestWithDefaults() *EntryDistributionListRequest`

NewEntryDistributionListRequestWithDefaults instantiates a new EntryDistributionListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *EntryDistributionListRequest) GetFilter() KalturaEntryDistributionFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *EntryDistributionListRequest) GetFilterOk() (*KalturaEntryDistributionFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *EntryDistributionListRequest) SetFilter(v KalturaEntryDistributionFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *EntryDistributionListRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPager

`func (o *EntryDistributionListRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *EntryDistributionListRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *EntryDistributionListRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *EntryDistributionListRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


