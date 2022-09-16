# BaseEntryUpdateContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedOptions** | Pointer to [**KalturaEntryReplacementOptions**](KalturaEntryReplacementOptions.md) |  | [optional] 
**ConversionProfileId** | Pointer to **int32** |  | [optional] 
**EntryId** | **string** |  | 
**Resource** | [**KalturaResource**](KalturaResource.md) |  | 

## Methods

### NewBaseEntryUpdateContentRequest

`func NewBaseEntryUpdateContentRequest(entryId string, resource KalturaResource, ) *BaseEntryUpdateContentRequest`

NewBaseEntryUpdateContentRequest instantiates a new BaseEntryUpdateContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntryUpdateContentRequestWithDefaults

`func NewBaseEntryUpdateContentRequestWithDefaults() *BaseEntryUpdateContentRequest`

NewBaseEntryUpdateContentRequestWithDefaults instantiates a new BaseEntryUpdateContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedOptions

`func (o *BaseEntryUpdateContentRequest) GetAdvancedOptions() KalturaEntryReplacementOptions`

GetAdvancedOptions returns the AdvancedOptions field if non-nil, zero value otherwise.

### GetAdvancedOptionsOk

`func (o *BaseEntryUpdateContentRequest) GetAdvancedOptionsOk() (*KalturaEntryReplacementOptions, bool)`

GetAdvancedOptionsOk returns a tuple with the AdvancedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedOptions

`func (o *BaseEntryUpdateContentRequest) SetAdvancedOptions(v KalturaEntryReplacementOptions)`

SetAdvancedOptions sets AdvancedOptions field to given value.

### HasAdvancedOptions

`func (o *BaseEntryUpdateContentRequest) HasAdvancedOptions() bool`

HasAdvancedOptions returns a boolean if a field has been set.

### GetConversionProfileId

`func (o *BaseEntryUpdateContentRequest) GetConversionProfileId() int32`

GetConversionProfileId returns the ConversionProfileId field if non-nil, zero value otherwise.

### GetConversionProfileIdOk

`func (o *BaseEntryUpdateContentRequest) GetConversionProfileIdOk() (*int32, bool)`

GetConversionProfileIdOk returns a tuple with the ConversionProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionProfileId

`func (o *BaseEntryUpdateContentRequest) SetConversionProfileId(v int32)`

SetConversionProfileId sets ConversionProfileId field to given value.

### HasConversionProfileId

`func (o *BaseEntryUpdateContentRequest) HasConversionProfileId() bool`

HasConversionProfileId returns a boolean if a field has been set.

### GetEntryId

`func (o *BaseEntryUpdateContentRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *BaseEntryUpdateContentRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *BaseEntryUpdateContentRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetResource

`func (o *BaseEntryUpdateContentRequest) GetResource() KalturaResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BaseEntryUpdateContentRequest) GetResourceOk() (*KalturaResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BaseEntryUpdateContentRequest) SetResource(v KalturaResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


