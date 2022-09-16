# MetadataProfileAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataProfile** | [**KalturaMetadataProfile**](KalturaMetadataProfile.md) |  | 
**ViewsData** | Pointer to **string** |  | [optional] 
**XsdData** | **string** |  | 

## Methods

### NewMetadataProfileAddRequest

`func NewMetadataProfileAddRequest(metadataProfile KalturaMetadataProfile, xsdData string, ) *MetadataProfileAddRequest`

NewMetadataProfileAddRequest instantiates a new MetadataProfileAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataProfileAddRequestWithDefaults

`func NewMetadataProfileAddRequestWithDefaults() *MetadataProfileAddRequest`

NewMetadataProfileAddRequestWithDefaults instantiates a new MetadataProfileAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataProfile

`func (o *MetadataProfileAddRequest) GetMetadataProfile() KalturaMetadataProfile`

GetMetadataProfile returns the MetadataProfile field if non-nil, zero value otherwise.

### GetMetadataProfileOk

`func (o *MetadataProfileAddRequest) GetMetadataProfileOk() (*KalturaMetadataProfile, bool)`

GetMetadataProfileOk returns a tuple with the MetadataProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfile

`func (o *MetadataProfileAddRequest) SetMetadataProfile(v KalturaMetadataProfile)`

SetMetadataProfile sets MetadataProfile field to given value.


### GetViewsData

`func (o *MetadataProfileAddRequest) GetViewsData() string`

GetViewsData returns the ViewsData field if non-nil, zero value otherwise.

### GetViewsDataOk

`func (o *MetadataProfileAddRequest) GetViewsDataOk() (*string, bool)`

GetViewsDataOk returns a tuple with the ViewsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewsData

`func (o *MetadataProfileAddRequest) SetViewsData(v string)`

SetViewsData sets ViewsData field to given value.

### HasViewsData

`func (o *MetadataProfileAddRequest) HasViewsData() bool`

HasViewsData returns a boolean if a field has been set.

### GetXsdData

`func (o *MetadataProfileAddRequest) GetXsdData() string`

GetXsdData returns the XsdData field if non-nil, zero value otherwise.

### GetXsdDataOk

`func (o *MetadataProfileAddRequest) GetXsdDataOk() (*string, bool)`

GetXsdDataOk returns a tuple with the XsdData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsdData

`func (o *MetadataProfileAddRequest) SetXsdData(v string)`

SetXsdData sets XsdData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


