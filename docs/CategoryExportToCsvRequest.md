# CategoryExportToCsvRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalFields** | Pointer to [**[]KalturaCsvAdditionalFieldInfo**](KalturaCsvAdditionalFieldInfo.md) |  | [optional] 
**Filter** | Pointer to [**KalturaCategoryFilter**](KalturaCategoryFilter.md) |  | [optional] 
**MappedFields** | Pointer to [**[]KalturaKeyValue**](KalturaKeyValue.md) |  | [optional] 
**MetadataProfileId** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**KalturaExportToCsvOptions**](KalturaExportToCsvOptions.md) |  | [optional] 

## Methods

### NewCategoryExportToCsvRequest

`func NewCategoryExportToCsvRequest() *CategoryExportToCsvRequest`

NewCategoryExportToCsvRequest instantiates a new CategoryExportToCsvRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryExportToCsvRequestWithDefaults

`func NewCategoryExportToCsvRequestWithDefaults() *CategoryExportToCsvRequest`

NewCategoryExportToCsvRequestWithDefaults instantiates a new CategoryExportToCsvRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalFields

`func (o *CategoryExportToCsvRequest) GetAdditionalFields() []KalturaCsvAdditionalFieldInfo`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CategoryExportToCsvRequest) GetAdditionalFieldsOk() (*[]KalturaCsvAdditionalFieldInfo, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CategoryExportToCsvRequest) SetAdditionalFields(v []KalturaCsvAdditionalFieldInfo)`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CategoryExportToCsvRequest) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetFilter

`func (o *CategoryExportToCsvRequest) GetFilter() KalturaCategoryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CategoryExportToCsvRequest) GetFilterOk() (*KalturaCategoryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CategoryExportToCsvRequest) SetFilter(v KalturaCategoryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CategoryExportToCsvRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMappedFields

`func (o *CategoryExportToCsvRequest) GetMappedFields() []KalturaKeyValue`

GetMappedFields returns the MappedFields field if non-nil, zero value otherwise.

### GetMappedFieldsOk

`func (o *CategoryExportToCsvRequest) GetMappedFieldsOk() (*[]KalturaKeyValue, bool)`

GetMappedFieldsOk returns a tuple with the MappedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedFields

`func (o *CategoryExportToCsvRequest) SetMappedFields(v []KalturaKeyValue)`

SetMappedFields sets MappedFields field to given value.

### HasMappedFields

`func (o *CategoryExportToCsvRequest) HasMappedFields() bool`

HasMappedFields returns a boolean if a field has been set.

### GetMetadataProfileId

`func (o *CategoryExportToCsvRequest) GetMetadataProfileId() int32`

GetMetadataProfileId returns the MetadataProfileId field if non-nil, zero value otherwise.

### GetMetadataProfileIdOk

`func (o *CategoryExportToCsvRequest) GetMetadataProfileIdOk() (*int32, bool)`

GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileId

`func (o *CategoryExportToCsvRequest) SetMetadataProfileId(v int32)`

SetMetadataProfileId sets MetadataProfileId field to given value.

### HasMetadataProfileId

`func (o *CategoryExportToCsvRequest) HasMetadataProfileId() bool`

HasMetadataProfileId returns a boolean if a field has been set.

### GetOptions

`func (o *CategoryExportToCsvRequest) GetOptions() KalturaExportToCsvOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CategoryExportToCsvRequest) GetOptionsOk() (*KalturaExportToCsvOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CategoryExportToCsvRequest) SetOptions(v KalturaExportToCsvOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CategoryExportToCsvRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


