# MediaAddFromEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaEntry** | Pointer to [**KalturaMediaEntry**](KalturaMediaEntry.md) |  | [optional] 
**SourceEntryId** | **string** |  | 
**SourceFlavorParamsId** | Pointer to **int32** |  | [optional] 

## Methods

### NewMediaAddFromEntryRequest

`func NewMediaAddFromEntryRequest(sourceEntryId string, ) *MediaAddFromEntryRequest`

NewMediaAddFromEntryRequest instantiates a new MediaAddFromEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaAddFromEntryRequestWithDefaults

`func NewMediaAddFromEntryRequestWithDefaults() *MediaAddFromEntryRequest`

NewMediaAddFromEntryRequestWithDefaults instantiates a new MediaAddFromEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaEntry

`func (o *MediaAddFromEntryRequest) GetMediaEntry() KalturaMediaEntry`

GetMediaEntry returns the MediaEntry field if non-nil, zero value otherwise.

### GetMediaEntryOk

`func (o *MediaAddFromEntryRequest) GetMediaEntryOk() (*KalturaMediaEntry, bool)`

GetMediaEntryOk returns a tuple with the MediaEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEntry

`func (o *MediaAddFromEntryRequest) SetMediaEntry(v KalturaMediaEntry)`

SetMediaEntry sets MediaEntry field to given value.

### HasMediaEntry

`func (o *MediaAddFromEntryRequest) HasMediaEntry() bool`

HasMediaEntry returns a boolean if a field has been set.

### GetSourceEntryId

`func (o *MediaAddFromEntryRequest) GetSourceEntryId() string`

GetSourceEntryId returns the SourceEntryId field if non-nil, zero value otherwise.

### GetSourceEntryIdOk

`func (o *MediaAddFromEntryRequest) GetSourceEntryIdOk() (*string, bool)`

GetSourceEntryIdOk returns a tuple with the SourceEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntryId

`func (o *MediaAddFromEntryRequest) SetSourceEntryId(v string)`

SetSourceEntryId sets SourceEntryId field to given value.


### GetSourceFlavorParamsId

`func (o *MediaAddFromEntryRequest) GetSourceFlavorParamsId() int32`

GetSourceFlavorParamsId returns the SourceFlavorParamsId field if non-nil, zero value otherwise.

### GetSourceFlavorParamsIdOk

`func (o *MediaAddFromEntryRequest) GetSourceFlavorParamsIdOk() (*int32, bool)`

GetSourceFlavorParamsIdOk returns a tuple with the SourceFlavorParamsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFlavorParamsId

`func (o *MediaAddFromEntryRequest) SetSourceFlavorParamsId(v int32)`

SetSourceFlavorParamsId sets SourceFlavorParamsId field to given value.

### HasSourceFlavorParamsId

`func (o *MediaAddFromEntryRequest) HasSourceFlavorParamsId() bool`

HasSourceFlavorParamsId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


