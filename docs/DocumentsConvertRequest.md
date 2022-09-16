# DocumentsConvertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionProfileId** | Pointer to **int32** |  | [optional] 
**DynamicConversionAttributes** | Pointer to [**[]KalturaConversionAttribute**](KalturaConversionAttribute.md) |  | [optional] 
**EntryId** | **string** |  | 

## Methods

### NewDocumentsConvertRequest

`func NewDocumentsConvertRequest(entryId string, ) *DocumentsConvertRequest`

NewDocumentsConvertRequest instantiates a new DocumentsConvertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentsConvertRequestWithDefaults

`func NewDocumentsConvertRequestWithDefaults() *DocumentsConvertRequest`

NewDocumentsConvertRequestWithDefaults instantiates a new DocumentsConvertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversionProfileId

`func (o *DocumentsConvertRequest) GetConversionProfileId() int32`

GetConversionProfileId returns the ConversionProfileId field if non-nil, zero value otherwise.

### GetConversionProfileIdOk

`func (o *DocumentsConvertRequest) GetConversionProfileIdOk() (*int32, bool)`

GetConversionProfileIdOk returns a tuple with the ConversionProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionProfileId

`func (o *DocumentsConvertRequest) SetConversionProfileId(v int32)`

SetConversionProfileId sets ConversionProfileId field to given value.

### HasConversionProfileId

`func (o *DocumentsConvertRequest) HasConversionProfileId() bool`

HasConversionProfileId returns a boolean if a field has been set.

### GetDynamicConversionAttributes

`func (o *DocumentsConvertRequest) GetDynamicConversionAttributes() []KalturaConversionAttribute`

GetDynamicConversionAttributes returns the DynamicConversionAttributes field if non-nil, zero value otherwise.

### GetDynamicConversionAttributesOk

`func (o *DocumentsConvertRequest) GetDynamicConversionAttributesOk() (*[]KalturaConversionAttribute, bool)`

GetDynamicConversionAttributesOk returns a tuple with the DynamicConversionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicConversionAttributes

`func (o *DocumentsConvertRequest) SetDynamicConversionAttributes(v []KalturaConversionAttribute)`

SetDynamicConversionAttributes sets DynamicConversionAttributes field to given value.

### HasDynamicConversionAttributes

`func (o *DocumentsConvertRequest) HasDynamicConversionAttributes() bool`

HasDynamicConversionAttributes returns a boolean if a field has been set.

### GetEntryId

`func (o *DocumentsConvertRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *DocumentsConvertRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *DocumentsConvertRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


