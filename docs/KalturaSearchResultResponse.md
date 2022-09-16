# KalturaSearchResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeedMediaInfo** | Pointer to **bool** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Objects** | Pointer to [**[]KalturaSearchResult**](KalturaSearchResult.md) |  | [optional] 

## Methods

### NewKalturaSearchResultResponse

`func NewKalturaSearchResultResponse() *KalturaSearchResultResponse`

NewKalturaSearchResultResponse instantiates a new KalturaSearchResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaSearchResultResponseWithDefaults

`func NewKalturaSearchResultResponseWithDefaults() *KalturaSearchResultResponse`

NewKalturaSearchResultResponseWithDefaults instantiates a new KalturaSearchResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeedMediaInfo

`func (o *KalturaSearchResultResponse) GetNeedMediaInfo() bool`

GetNeedMediaInfo returns the NeedMediaInfo field if non-nil, zero value otherwise.

### GetNeedMediaInfoOk

`func (o *KalturaSearchResultResponse) GetNeedMediaInfoOk() (*bool, bool)`

GetNeedMediaInfoOk returns a tuple with the NeedMediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedMediaInfo

`func (o *KalturaSearchResultResponse) SetNeedMediaInfo(v bool)`

SetNeedMediaInfo sets NeedMediaInfo field to given value.

### HasNeedMediaInfo

`func (o *KalturaSearchResultResponse) HasNeedMediaInfo() bool`

HasNeedMediaInfo returns a boolean if a field has been set.

### GetObjects

`func (o *KalturaSearchResultResponse) GetObjects() []KalturaSearchResult`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *KalturaSearchResultResponse) GetObjectsOk() (*[]KalturaSearchResult, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *KalturaSearchResultResponse) SetObjects(v []KalturaSearchResult)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *KalturaSearchResultResponse) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


