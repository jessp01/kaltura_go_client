# CaptionAssetItemListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptionAssetId** | **string** |  | 
**CaptionAssetItemFilter** | Pointer to [**KalturaCaptionAssetItemFilter**](KalturaCaptionAssetItemFilter.md) |  | [optional] 
**CaptionAssetItemPager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 

## Methods

### NewCaptionAssetItemListRequest

`func NewCaptionAssetItemListRequest(captionAssetId string, ) *CaptionAssetItemListRequest`

NewCaptionAssetItemListRequest instantiates a new CaptionAssetItemListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptionAssetItemListRequestWithDefaults

`func NewCaptionAssetItemListRequestWithDefaults() *CaptionAssetItemListRequest`

NewCaptionAssetItemListRequestWithDefaults instantiates a new CaptionAssetItemListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptionAssetId

`func (o *CaptionAssetItemListRequest) GetCaptionAssetId() string`

GetCaptionAssetId returns the CaptionAssetId field if non-nil, zero value otherwise.

### GetCaptionAssetIdOk

`func (o *CaptionAssetItemListRequest) GetCaptionAssetIdOk() (*string, bool)`

GetCaptionAssetIdOk returns a tuple with the CaptionAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionAssetId

`func (o *CaptionAssetItemListRequest) SetCaptionAssetId(v string)`

SetCaptionAssetId sets CaptionAssetId field to given value.


### GetCaptionAssetItemFilter

`func (o *CaptionAssetItemListRequest) GetCaptionAssetItemFilter() KalturaCaptionAssetItemFilter`

GetCaptionAssetItemFilter returns the CaptionAssetItemFilter field if non-nil, zero value otherwise.

### GetCaptionAssetItemFilterOk

`func (o *CaptionAssetItemListRequest) GetCaptionAssetItemFilterOk() (*KalturaCaptionAssetItemFilter, bool)`

GetCaptionAssetItemFilterOk returns a tuple with the CaptionAssetItemFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionAssetItemFilter

`func (o *CaptionAssetItemListRequest) SetCaptionAssetItemFilter(v KalturaCaptionAssetItemFilter)`

SetCaptionAssetItemFilter sets CaptionAssetItemFilter field to given value.

### HasCaptionAssetItemFilter

`func (o *CaptionAssetItemListRequest) HasCaptionAssetItemFilter() bool`

HasCaptionAssetItemFilter returns a boolean if a field has been set.

### GetCaptionAssetItemPager

`func (o *CaptionAssetItemListRequest) GetCaptionAssetItemPager() KalturaFilterPager`

GetCaptionAssetItemPager returns the CaptionAssetItemPager field if non-nil, zero value otherwise.

### GetCaptionAssetItemPagerOk

`func (o *CaptionAssetItemListRequest) GetCaptionAssetItemPagerOk() (*KalturaFilterPager, bool)`

GetCaptionAssetItemPagerOk returns a tuple with the CaptionAssetItemPager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionAssetItemPager

`func (o *CaptionAssetItemListRequest) SetCaptionAssetItemPager(v KalturaFilterPager)`

SetCaptionAssetItemPager sets CaptionAssetItemPager field to given value.

### HasCaptionAssetItemPager

`func (o *CaptionAssetItemListRequest) HasCaptionAssetItemPager() bool`

HasCaptionAssetItemPager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


