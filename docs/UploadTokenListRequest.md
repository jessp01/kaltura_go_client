# UploadTokenListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaUploadTokenFilter**](KalturaUploadTokenFilter.md) |  | [optional] 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 

## Methods

### NewUploadTokenListRequest

`func NewUploadTokenListRequest() *UploadTokenListRequest`

NewUploadTokenListRequest instantiates a new UploadTokenListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadTokenListRequestWithDefaults

`func NewUploadTokenListRequestWithDefaults() *UploadTokenListRequest`

NewUploadTokenListRequestWithDefaults instantiates a new UploadTokenListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *UploadTokenListRequest) GetFilter() KalturaUploadTokenFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UploadTokenListRequest) GetFilterOk() (*KalturaUploadTokenFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UploadTokenListRequest) SetFilter(v KalturaUploadTokenFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UploadTokenListRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPager

`func (o *UploadTokenListRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *UploadTokenListRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *UploadTokenListRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *UploadTokenListRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


