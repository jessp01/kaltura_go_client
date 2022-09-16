# ESearchSearchGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pager** | Pointer to [**KalturaPager**](KalturaPager.md) |  | [optional] 
**SearchParams** | [**KalturaESearchGroupParams**](KalturaESearchGroupParams.md) |  | 

## Methods

### NewESearchSearchGroupRequest

`func NewESearchSearchGroupRequest(searchParams KalturaESearchGroupParams, ) *ESearchSearchGroupRequest`

NewESearchSearchGroupRequest instantiates a new ESearchSearchGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESearchSearchGroupRequestWithDefaults

`func NewESearchSearchGroupRequestWithDefaults() *ESearchSearchGroupRequest`

NewESearchSearchGroupRequestWithDefaults instantiates a new ESearchSearchGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPager

`func (o *ESearchSearchGroupRequest) GetPager() KalturaPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *ESearchSearchGroupRequest) GetPagerOk() (*KalturaPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *ESearchSearchGroupRequest) SetPager(v KalturaPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *ESearchSearchGroupRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetSearchParams

`func (o *ESearchSearchGroupRequest) GetSearchParams() KalturaESearchGroupParams`

GetSearchParams returns the SearchParams field if non-nil, zero value otherwise.

### GetSearchParamsOk

`func (o *ESearchSearchGroupRequest) GetSearchParamsOk() (*KalturaESearchGroupParams, bool)`

GetSearchParamsOk returns a tuple with the SearchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchParams

`func (o *ESearchSearchGroupRequest) SetSearchParams(v KalturaESearchGroupParams)`

SetSearchParams sets SearchParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


