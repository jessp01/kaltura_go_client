# BulkListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkUploadFilter** | Pointer to [**KalturaBulkUploadFilter**](KalturaBulkUploadFilter.md) |  | [optional] 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 

## Methods

### NewBulkListRequest

`func NewBulkListRequest() *BulkListRequest`

NewBulkListRequest instantiates a new BulkListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkListRequestWithDefaults

`func NewBulkListRequestWithDefaults() *BulkListRequest`

NewBulkListRequestWithDefaults instantiates a new BulkListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkUploadFilter

`func (o *BulkListRequest) GetBulkUploadFilter() KalturaBulkUploadFilter`

GetBulkUploadFilter returns the BulkUploadFilter field if non-nil, zero value otherwise.

### GetBulkUploadFilterOk

`func (o *BulkListRequest) GetBulkUploadFilterOk() (*KalturaBulkUploadFilter, bool)`

GetBulkUploadFilterOk returns a tuple with the BulkUploadFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkUploadFilter

`func (o *BulkListRequest) SetBulkUploadFilter(v KalturaBulkUploadFilter)`

SetBulkUploadFilter sets BulkUploadFilter field to given value.

### HasBulkUploadFilter

`func (o *BulkListRequest) HasBulkUploadFilter() bool`

HasBulkUploadFilter returns a boolean if a field has been set.

### GetPager

`func (o *BulkListRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *BulkListRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *BulkListRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *BulkListRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


