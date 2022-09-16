# KalturaESearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Highlight** | Pointer to [**[]KalturaESearchHighlight**](KalturaESearchHighlight.md) |  | [optional] 
**ItemsData** | Pointer to [**[]KalturaESearchItemDataResult**](KalturaESearchItemDataResult.md) |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaESearchResult

`func NewKalturaESearchResult() *KalturaESearchResult`

NewKalturaESearchResult instantiates a new KalturaESearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaESearchResultWithDefaults

`func NewKalturaESearchResultWithDefaults() *KalturaESearchResult`

NewKalturaESearchResultWithDefaults instantiates a new KalturaESearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighlight

`func (o *KalturaESearchResult) GetHighlight() []KalturaESearchHighlight`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *KalturaESearchResult) GetHighlightOk() (*[]KalturaESearchHighlight, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *KalturaESearchResult) SetHighlight(v []KalturaESearchHighlight)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *KalturaESearchResult) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetItemsData

`func (o *KalturaESearchResult) GetItemsData() []KalturaESearchItemDataResult`

GetItemsData returns the ItemsData field if non-nil, zero value otherwise.

### GetItemsDataOk

`func (o *KalturaESearchResult) GetItemsDataOk() (*[]KalturaESearchItemDataResult, bool)`

GetItemsDataOk returns a tuple with the ItemsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsData

`func (o *KalturaESearchResult) SetItemsData(v []KalturaESearchItemDataResult)`

SetItemsData sets ItemsData field to given value.

### HasItemsData

`func (o *KalturaESearchResult) HasItemsData() bool`

HasItemsData returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaESearchResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaESearchResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaESearchResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaESearchResult) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


