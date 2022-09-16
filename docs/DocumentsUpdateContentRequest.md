# DocumentsUpdateContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionProfileId** | Pointer to **int32** |  | [optional] 
**EntryId** | **string** |  | 
**Resource** | [**KalturaResource**](KalturaResource.md) |  | 

## Methods

### NewDocumentsUpdateContentRequest

`func NewDocumentsUpdateContentRequest(entryId string, resource KalturaResource, ) *DocumentsUpdateContentRequest`

NewDocumentsUpdateContentRequest instantiates a new DocumentsUpdateContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentsUpdateContentRequestWithDefaults

`func NewDocumentsUpdateContentRequestWithDefaults() *DocumentsUpdateContentRequest`

NewDocumentsUpdateContentRequestWithDefaults instantiates a new DocumentsUpdateContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversionProfileId

`func (o *DocumentsUpdateContentRequest) GetConversionProfileId() int32`

GetConversionProfileId returns the ConversionProfileId field if non-nil, zero value otherwise.

### GetConversionProfileIdOk

`func (o *DocumentsUpdateContentRequest) GetConversionProfileIdOk() (*int32, bool)`

GetConversionProfileIdOk returns a tuple with the ConversionProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionProfileId

`func (o *DocumentsUpdateContentRequest) SetConversionProfileId(v int32)`

SetConversionProfileId sets ConversionProfileId field to given value.

### HasConversionProfileId

`func (o *DocumentsUpdateContentRequest) HasConversionProfileId() bool`

HasConversionProfileId returns a boolean if a field has been set.

### GetEntryId

`func (o *DocumentsUpdateContentRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *DocumentsUpdateContentRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *DocumentsUpdateContentRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetResource

`func (o *DocumentsUpdateContentRequest) GetResource() KalturaResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DocumentsUpdateContentRequest) GetResourceOk() (*KalturaResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DocumentsUpdateContentRequest) SetResource(v KalturaResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


