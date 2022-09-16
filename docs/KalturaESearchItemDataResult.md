# KalturaESearchItemDataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]KalturaESearchItemData**](KalturaESearchItemData.md) |  | [optional] 
**ItemsType** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewKalturaESearchItemDataResult

`func NewKalturaESearchItemDataResult() *KalturaESearchItemDataResult`

NewKalturaESearchItemDataResult instantiates a new KalturaESearchItemDataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaESearchItemDataResultWithDefaults

`func NewKalturaESearchItemDataResultWithDefaults() *KalturaESearchItemDataResult`

NewKalturaESearchItemDataResultWithDefaults instantiates a new KalturaESearchItemDataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *KalturaESearchItemDataResult) GetItems() []KalturaESearchItemData`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KalturaESearchItemDataResult) GetItemsOk() (*[]KalturaESearchItemData, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KalturaESearchItemDataResult) SetItems(v []KalturaESearchItemData)`

SetItems sets Items field to given value.

### HasItems

`func (o *KalturaESearchItemDataResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetItemsType

`func (o *KalturaESearchItemDataResult) GetItemsType() string`

GetItemsType returns the ItemsType field if non-nil, zero value otherwise.

### GetItemsTypeOk

`func (o *KalturaESearchItemDataResult) GetItemsTypeOk() (*string, bool)`

GetItemsTypeOk returns a tuple with the ItemsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsType

`func (o *KalturaESearchItemDataResult) SetItemsType(v string)`

SetItemsType sets ItemsType field to given value.

### HasItemsType

`func (o *KalturaESearchItemDataResult) HasItemsType() bool`

HasItemsType returns a boolean if a field has been set.

### GetTotalCount

`func (o *KalturaESearchItemDataResult) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *KalturaESearchItemDataResult) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *KalturaESearchItemDataResult) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *KalturaESearchItemDataResult) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


