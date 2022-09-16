# BaseEntryExportToCsvRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalFields** | Pointer to [**[]KalturaCsvAdditionalFieldInfo**](KalturaCsvAdditionalFieldInfo.md) |  | [optional] 
**Filter** | Pointer to [**KalturaBaseEntryFilter**](KalturaBaseEntryFilter.md) |  | [optional] 
**MappedFields** | Pointer to [**[]KalturaKeyValue**](KalturaKeyValue.md) |  | [optional] 
**MetadataProfileId** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**KalturaExportToCsvOptions**](KalturaExportToCsvOptions.md) |  | [optional] 

## Methods

### NewBaseEntryExportToCsvRequest

`func NewBaseEntryExportToCsvRequest() *BaseEntryExportToCsvRequest`

NewBaseEntryExportToCsvRequest instantiates a new BaseEntryExportToCsvRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntryExportToCsvRequestWithDefaults

`func NewBaseEntryExportToCsvRequestWithDefaults() *BaseEntryExportToCsvRequest`

NewBaseEntryExportToCsvRequestWithDefaults instantiates a new BaseEntryExportToCsvRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalFields

`func (o *BaseEntryExportToCsvRequest) GetAdditionalFields() []KalturaCsvAdditionalFieldInfo`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *BaseEntryExportToCsvRequest) GetAdditionalFieldsOk() (*[]KalturaCsvAdditionalFieldInfo, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *BaseEntryExportToCsvRequest) SetAdditionalFields(v []KalturaCsvAdditionalFieldInfo)`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *BaseEntryExportToCsvRequest) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetFilter

`func (o *BaseEntryExportToCsvRequest) GetFilter() KalturaBaseEntryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *BaseEntryExportToCsvRequest) GetFilterOk() (*KalturaBaseEntryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *BaseEntryExportToCsvRequest) SetFilter(v KalturaBaseEntryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *BaseEntryExportToCsvRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMappedFields

`func (o *BaseEntryExportToCsvRequest) GetMappedFields() []KalturaKeyValue`

GetMappedFields returns the MappedFields field if non-nil, zero value otherwise.

### GetMappedFieldsOk

`func (o *BaseEntryExportToCsvRequest) GetMappedFieldsOk() (*[]KalturaKeyValue, bool)`

GetMappedFieldsOk returns a tuple with the MappedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedFields

`func (o *BaseEntryExportToCsvRequest) SetMappedFields(v []KalturaKeyValue)`

SetMappedFields sets MappedFields field to given value.

### HasMappedFields

`func (o *BaseEntryExportToCsvRequest) HasMappedFields() bool`

HasMappedFields returns a boolean if a field has been set.

### GetMetadataProfileId

`func (o *BaseEntryExportToCsvRequest) GetMetadataProfileId() int32`

GetMetadataProfileId returns the MetadataProfileId field if non-nil, zero value otherwise.

### GetMetadataProfileIdOk

`func (o *BaseEntryExportToCsvRequest) GetMetadataProfileIdOk() (*int32, bool)`

GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileId

`func (o *BaseEntryExportToCsvRequest) SetMetadataProfileId(v int32)`

SetMetadataProfileId sets MetadataProfileId field to given value.

### HasMetadataProfileId

`func (o *BaseEntryExportToCsvRequest) HasMetadataProfileId() bool`

HasMetadataProfileId returns a boolean if a field has been set.

### GetOptions

`func (o *BaseEntryExportToCsvRequest) GetOptions() KalturaExportToCsvOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BaseEntryExportToCsvRequest) GetOptionsOk() (*KalturaExportToCsvOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BaseEntryExportToCsvRequest) SetOptions(v KalturaExportToCsvOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BaseEntryExportToCsvRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


