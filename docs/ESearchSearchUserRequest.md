# ESearchSearchUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pager** | Pointer to [**KalturaPager**](KalturaPager.md) |  | [optional] 
**SearchParams** | [**KalturaESearchUserParams**](KalturaESearchUserParams.md) |  | 

## Methods

### NewESearchSearchUserRequest

`func NewESearchSearchUserRequest(searchParams KalturaESearchUserParams, ) *ESearchSearchUserRequest`

NewESearchSearchUserRequest instantiates a new ESearchSearchUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESearchSearchUserRequestWithDefaults

`func NewESearchSearchUserRequestWithDefaults() *ESearchSearchUserRequest`

NewESearchSearchUserRequestWithDefaults instantiates a new ESearchSearchUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPager

`func (o *ESearchSearchUserRequest) GetPager() KalturaPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *ESearchSearchUserRequest) GetPagerOk() (*KalturaPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *ESearchSearchUserRequest) SetPager(v KalturaPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *ESearchSearchUserRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetSearchParams

`func (o *ESearchSearchUserRequest) GetSearchParams() KalturaESearchUserParams`

GetSearchParams returns the SearchParams field if non-nil, zero value otherwise.

### GetSearchParamsOk

`func (o *ESearchSearchUserRequest) GetSearchParamsOk() (*KalturaESearchUserParams, bool)`

GetSearchParamsOk returns a tuple with the SearchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchParams

`func (o *ESearchSearchUserRequest) SetSearchParams(v KalturaESearchUserParams)`

SetSearchParams sets SearchParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


