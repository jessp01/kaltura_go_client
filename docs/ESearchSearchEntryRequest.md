# ESearchSearchEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pager** | Pointer to [**KalturaPager**](KalturaPager.md) |  | [optional] 
**SearchParams** | [**KalturaESearchEntryParams**](KalturaESearchEntryParams.md) |  | 

## Methods

### NewESearchSearchEntryRequest

`func NewESearchSearchEntryRequest(searchParams KalturaESearchEntryParams, ) *ESearchSearchEntryRequest`

NewESearchSearchEntryRequest instantiates a new ESearchSearchEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESearchSearchEntryRequestWithDefaults

`func NewESearchSearchEntryRequestWithDefaults() *ESearchSearchEntryRequest`

NewESearchSearchEntryRequestWithDefaults instantiates a new ESearchSearchEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPager

`func (o *ESearchSearchEntryRequest) GetPager() KalturaPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *ESearchSearchEntryRequest) GetPagerOk() (*KalturaPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *ESearchSearchEntryRequest) SetPager(v KalturaPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *ESearchSearchEntryRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetSearchParams

`func (o *ESearchSearchEntryRequest) GetSearchParams() KalturaESearchEntryParams`

GetSearchParams returns the SearchParams field if non-nil, zero value otherwise.

### GetSearchParamsOk

`func (o *ESearchSearchEntryRequest) GetSearchParamsOk() (*KalturaESearchEntryParams, bool)`

GetSearchParamsOk returns a tuple with the SearchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchParams

`func (o *ESearchSearchEntryRequest) SetSearchParams(v KalturaESearchEntryParams)`

SetSearchParams sets SearchParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


