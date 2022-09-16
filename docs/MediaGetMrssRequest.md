# MediaGetMrssRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**ExtendingItemsArray** | Pointer to [**[]KalturaExtendingItemMrssParameter**](KalturaExtendingItemMrssParameter.md) |  | [optional] 
**Features** | Pointer to **string** |  | [optional] 

## Methods

### NewMediaGetMrssRequest

`func NewMediaGetMrssRequest(entryId string, ) *MediaGetMrssRequest`

NewMediaGetMrssRequest instantiates a new MediaGetMrssRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaGetMrssRequestWithDefaults

`func NewMediaGetMrssRequestWithDefaults() *MediaGetMrssRequest`

NewMediaGetMrssRequestWithDefaults instantiates a new MediaGetMrssRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *MediaGetMrssRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *MediaGetMrssRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *MediaGetMrssRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetExtendingItemsArray

`func (o *MediaGetMrssRequest) GetExtendingItemsArray() []KalturaExtendingItemMrssParameter`

GetExtendingItemsArray returns the ExtendingItemsArray field if non-nil, zero value otherwise.

### GetExtendingItemsArrayOk

`func (o *MediaGetMrssRequest) GetExtendingItemsArrayOk() (*[]KalturaExtendingItemMrssParameter, bool)`

GetExtendingItemsArrayOk returns a tuple with the ExtendingItemsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendingItemsArray

`func (o *MediaGetMrssRequest) SetExtendingItemsArray(v []KalturaExtendingItemMrssParameter)`

SetExtendingItemsArray sets ExtendingItemsArray field to given value.

### HasExtendingItemsArray

`func (o *MediaGetMrssRequest) HasExtendingItemsArray() bool`

HasExtendingItemsArray returns a boolean if a field has been set.

### GetFeatures

`func (o *MediaGetMrssRequest) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *MediaGetMrssRequest) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *MediaGetMrssRequest) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *MediaGetMrssRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


