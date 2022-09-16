# KalturaEntryReplacementOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeepManualThumbnails** | Pointer to **int32** | If true manually created thumbnails will not be deleted on entry replacement | [optional] 
**PluginOptionItems** | Pointer to [**[]KalturaPluginReplacementOptionsItem**](KalturaPluginReplacementOptionsItem.md) |  | [optional] 

## Methods

### NewKalturaEntryReplacementOptions

`func NewKalturaEntryReplacementOptions() *KalturaEntryReplacementOptions`

NewKalturaEntryReplacementOptions instantiates a new KalturaEntryReplacementOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaEntryReplacementOptionsWithDefaults

`func NewKalturaEntryReplacementOptionsWithDefaults() *KalturaEntryReplacementOptions`

NewKalturaEntryReplacementOptionsWithDefaults instantiates a new KalturaEntryReplacementOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeepManualThumbnails

`func (o *KalturaEntryReplacementOptions) GetKeepManualThumbnails() int32`

GetKeepManualThumbnails returns the KeepManualThumbnails field if non-nil, zero value otherwise.

### GetKeepManualThumbnailsOk

`func (o *KalturaEntryReplacementOptions) GetKeepManualThumbnailsOk() (*int32, bool)`

GetKeepManualThumbnailsOk returns a tuple with the KeepManualThumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepManualThumbnails

`func (o *KalturaEntryReplacementOptions) SetKeepManualThumbnails(v int32)`

SetKeepManualThumbnails sets KeepManualThumbnails field to given value.

### HasKeepManualThumbnails

`func (o *KalturaEntryReplacementOptions) HasKeepManualThumbnails() bool`

HasKeepManualThumbnails returns a boolean if a field has been set.

### GetPluginOptionItems

`func (o *KalturaEntryReplacementOptions) GetPluginOptionItems() []KalturaPluginReplacementOptionsItem`

GetPluginOptionItems returns the PluginOptionItems field if non-nil, zero value otherwise.

### GetPluginOptionItemsOk

`func (o *KalturaEntryReplacementOptions) GetPluginOptionItemsOk() (*[]KalturaPluginReplacementOptionsItem, bool)`

GetPluginOptionItemsOk returns a tuple with the PluginOptionItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginOptionItems

`func (o *KalturaEntryReplacementOptions) SetPluginOptionItems(v []KalturaPluginReplacementOptionsItem)`

SetPluginOptionItems sets PluginOptionItems field to given value.

### HasPluginOptionItems

`func (o *KalturaEntryReplacementOptions) HasPluginOptionItems() bool`

HasPluginOptionItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


