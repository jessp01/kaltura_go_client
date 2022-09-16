# MetadataProfileUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**MetadataProfile** | [**KalturaMetadataProfile**](KalturaMetadataProfile.md) |  | 
**ViewsData** | Pointer to **string** |  | [optional] 
**XsdData** | Pointer to **string** |  | [optional] 

## Methods

### NewMetadataProfileUpdateRequest

`func NewMetadataProfileUpdateRequest(id int32, metadataProfile KalturaMetadataProfile, ) *MetadataProfileUpdateRequest`

NewMetadataProfileUpdateRequest instantiates a new MetadataProfileUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataProfileUpdateRequestWithDefaults

`func NewMetadataProfileUpdateRequestWithDefaults() *MetadataProfileUpdateRequest`

NewMetadataProfileUpdateRequestWithDefaults instantiates a new MetadataProfileUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataProfileUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataProfileUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataProfileUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetMetadataProfile

`func (o *MetadataProfileUpdateRequest) GetMetadataProfile() KalturaMetadataProfile`

GetMetadataProfile returns the MetadataProfile field if non-nil, zero value otherwise.

### GetMetadataProfileOk

`func (o *MetadataProfileUpdateRequest) GetMetadataProfileOk() (*KalturaMetadataProfile, bool)`

GetMetadataProfileOk returns a tuple with the MetadataProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfile

`func (o *MetadataProfileUpdateRequest) SetMetadataProfile(v KalturaMetadataProfile)`

SetMetadataProfile sets MetadataProfile field to given value.


### GetViewsData

`func (o *MetadataProfileUpdateRequest) GetViewsData() string`

GetViewsData returns the ViewsData field if non-nil, zero value otherwise.

### GetViewsDataOk

`func (o *MetadataProfileUpdateRequest) GetViewsDataOk() (*string, bool)`

GetViewsDataOk returns a tuple with the ViewsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewsData

`func (o *MetadataProfileUpdateRequest) SetViewsData(v string)`

SetViewsData sets ViewsData field to given value.

### HasViewsData

`func (o *MetadataProfileUpdateRequest) HasViewsData() bool`

HasViewsData returns a boolean if a field has been set.

### GetXsdData

`func (o *MetadataProfileUpdateRequest) GetXsdData() string`

GetXsdData returns the XsdData field if non-nil, zero value otherwise.

### GetXsdDataOk

`func (o *MetadataProfileUpdateRequest) GetXsdDataOk() (*string, bool)`

GetXsdDataOk returns a tuple with the XsdData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsdData

`func (o *MetadataProfileUpdateRequest) SetXsdData(v string)`

SetXsdData sets XsdData field to given value.

### HasXsdData

`func (o *MetadataProfileUpdateRequest) HasXsdData() bool`

HasXsdData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


