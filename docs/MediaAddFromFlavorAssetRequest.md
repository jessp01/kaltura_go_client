# MediaAddFromFlavorAssetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaEntry** | Pointer to [**KalturaMediaEntry**](KalturaMediaEntry.md) |  | [optional] 
**SourceFlavorAssetId** | **string** |  | 

## Methods

### NewMediaAddFromFlavorAssetRequest

`func NewMediaAddFromFlavorAssetRequest(sourceFlavorAssetId string, ) *MediaAddFromFlavorAssetRequest`

NewMediaAddFromFlavorAssetRequest instantiates a new MediaAddFromFlavorAssetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaAddFromFlavorAssetRequestWithDefaults

`func NewMediaAddFromFlavorAssetRequestWithDefaults() *MediaAddFromFlavorAssetRequest`

NewMediaAddFromFlavorAssetRequestWithDefaults instantiates a new MediaAddFromFlavorAssetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaEntry

`func (o *MediaAddFromFlavorAssetRequest) GetMediaEntry() KalturaMediaEntry`

GetMediaEntry returns the MediaEntry field if non-nil, zero value otherwise.

### GetMediaEntryOk

`func (o *MediaAddFromFlavorAssetRequest) GetMediaEntryOk() (*KalturaMediaEntry, bool)`

GetMediaEntryOk returns a tuple with the MediaEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEntry

`func (o *MediaAddFromFlavorAssetRequest) SetMediaEntry(v KalturaMediaEntry)`

SetMediaEntry sets MediaEntry field to given value.

### HasMediaEntry

`func (o *MediaAddFromFlavorAssetRequest) HasMediaEntry() bool`

HasMediaEntry returns a boolean if a field has been set.

### GetSourceFlavorAssetId

`func (o *MediaAddFromFlavorAssetRequest) GetSourceFlavorAssetId() string`

GetSourceFlavorAssetId returns the SourceFlavorAssetId field if non-nil, zero value otherwise.

### GetSourceFlavorAssetIdOk

`func (o *MediaAddFromFlavorAssetRequest) GetSourceFlavorAssetIdOk() (*string, bool)`

GetSourceFlavorAssetIdOk returns a tuple with the SourceFlavorAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFlavorAssetId

`func (o *MediaAddFromFlavorAssetRequest) SetSourceFlavorAssetId(v string)`

SetSourceFlavorAssetId sets SourceFlavorAssetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


