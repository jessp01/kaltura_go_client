# ESearchSearchCategoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pager** | Pointer to [**KalturaPager**](KalturaPager.md) |  | [optional] 
**SearchParams** | [**KalturaESearchCategoryParams**](KalturaESearchCategoryParams.md) |  | 

## Methods

### NewESearchSearchCategoryRequest

`func NewESearchSearchCategoryRequest(searchParams KalturaESearchCategoryParams, ) *ESearchSearchCategoryRequest`

NewESearchSearchCategoryRequest instantiates a new ESearchSearchCategoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESearchSearchCategoryRequestWithDefaults

`func NewESearchSearchCategoryRequestWithDefaults() *ESearchSearchCategoryRequest`

NewESearchSearchCategoryRequestWithDefaults instantiates a new ESearchSearchCategoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPager

`func (o *ESearchSearchCategoryRequest) GetPager() KalturaPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *ESearchSearchCategoryRequest) GetPagerOk() (*KalturaPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *ESearchSearchCategoryRequest) SetPager(v KalturaPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *ESearchSearchCategoryRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetSearchParams

`func (o *ESearchSearchCategoryRequest) GetSearchParams() KalturaESearchCategoryParams`

GetSearchParams returns the SearchParams field if non-nil, zero value otherwise.

### GetSearchParamsOk

`func (o *ESearchSearchCategoryRequest) GetSearchParamsOk() (*KalturaESearchCategoryParams, bool)`

GetSearchParamsOk returns a tuple with the SearchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchParams

`func (o *ESearchSearchCategoryRequest) SetSearchParams(v KalturaESearchCategoryParams)`

SetSearchParams sets SearchParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


