# FileAssetListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**KalturaFileAssetFilter**](KalturaFileAssetFilter.md) |  | 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 

## Methods

### NewFileAssetListRequest

`func NewFileAssetListRequest(filter KalturaFileAssetFilter, ) *FileAssetListRequest`

NewFileAssetListRequest instantiates a new FileAssetListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileAssetListRequestWithDefaults

`func NewFileAssetListRequestWithDefaults() *FileAssetListRequest`

NewFileAssetListRequestWithDefaults instantiates a new FileAssetListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *FileAssetListRequest) GetFilter() KalturaFileAssetFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *FileAssetListRequest) GetFilterOk() (*KalturaFileAssetFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *FileAssetListRequest) SetFilter(v KalturaFileAssetFilter)`

SetFilter sets Filter field to given value.


### GetPager

`func (o *FileAssetListRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *FileAssetListRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *FileAssetListRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *FileAssetListRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


