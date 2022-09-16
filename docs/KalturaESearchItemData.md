# KalturaESearchItemData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Highlight** | Pointer to [**[]KalturaESearchHighlight**](KalturaESearchHighlight.md) |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaESearchItemData

`func NewKalturaESearchItemData() *KalturaESearchItemData`

NewKalturaESearchItemData instantiates a new KalturaESearchItemData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaESearchItemDataWithDefaults

`func NewKalturaESearchItemDataWithDefaults() *KalturaESearchItemData`

NewKalturaESearchItemDataWithDefaults instantiates a new KalturaESearchItemData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighlight

`func (o *KalturaESearchItemData) GetHighlight() []KalturaESearchHighlight`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *KalturaESearchItemData) GetHighlightOk() (*[]KalturaESearchHighlight, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *KalturaESearchItemData) SetHighlight(v []KalturaESearchHighlight)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *KalturaESearchItemData) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaESearchItemData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaESearchItemData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaESearchItemData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaESearchItemData) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


