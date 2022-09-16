# BaseEntryListByReferenceIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 
**RefId** | **string** |  | 

## Methods

### NewBaseEntryListByReferenceIdRequest

`func NewBaseEntryListByReferenceIdRequest(refId string, ) *BaseEntryListByReferenceIdRequest`

NewBaseEntryListByReferenceIdRequest instantiates a new BaseEntryListByReferenceIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntryListByReferenceIdRequestWithDefaults

`func NewBaseEntryListByReferenceIdRequestWithDefaults() *BaseEntryListByReferenceIdRequest`

NewBaseEntryListByReferenceIdRequestWithDefaults instantiates a new BaseEntryListByReferenceIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPager

`func (o *BaseEntryListByReferenceIdRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *BaseEntryListByReferenceIdRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *BaseEntryListByReferenceIdRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *BaseEntryListByReferenceIdRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetRefId

`func (o *BaseEntryListByReferenceIdRequest) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *BaseEntryListByReferenceIdRequest) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *BaseEntryListByReferenceIdRequest) SetRefId(v string)`

SetRefId sets RefId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


