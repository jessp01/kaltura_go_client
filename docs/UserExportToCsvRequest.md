# UserExportToCsvRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalFields** | Pointer to [**[]KalturaCsvAdditionalFieldInfo**](KalturaCsvAdditionalFieldInfo.md) |  | [optional] 
**Filter** | Pointer to [**KalturaUserFilter**](KalturaUserFilter.md) |  | [optional] 
**MappedFields** | Pointer to [**[]KalturaKeyValue**](KalturaKeyValue.md) |  | [optional] 
**MetadataProfileId** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**KalturaExportToCsvOptions**](KalturaExportToCsvOptions.md) |  | [optional] 

## Methods

### NewUserExportToCsvRequest

`func NewUserExportToCsvRequest() *UserExportToCsvRequest`

NewUserExportToCsvRequest instantiates a new UserExportToCsvRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExportToCsvRequestWithDefaults

`func NewUserExportToCsvRequestWithDefaults() *UserExportToCsvRequest`

NewUserExportToCsvRequestWithDefaults instantiates a new UserExportToCsvRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalFields

`func (o *UserExportToCsvRequest) GetAdditionalFields() []KalturaCsvAdditionalFieldInfo`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *UserExportToCsvRequest) GetAdditionalFieldsOk() (*[]KalturaCsvAdditionalFieldInfo, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *UserExportToCsvRequest) SetAdditionalFields(v []KalturaCsvAdditionalFieldInfo)`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *UserExportToCsvRequest) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetFilter

`func (o *UserExportToCsvRequest) GetFilter() KalturaUserFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UserExportToCsvRequest) GetFilterOk() (*KalturaUserFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UserExportToCsvRequest) SetFilter(v KalturaUserFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UserExportToCsvRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMappedFields

`func (o *UserExportToCsvRequest) GetMappedFields() []KalturaKeyValue`

GetMappedFields returns the MappedFields field if non-nil, zero value otherwise.

### GetMappedFieldsOk

`func (o *UserExportToCsvRequest) GetMappedFieldsOk() (*[]KalturaKeyValue, bool)`

GetMappedFieldsOk returns a tuple with the MappedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedFields

`func (o *UserExportToCsvRequest) SetMappedFields(v []KalturaKeyValue)`

SetMappedFields sets MappedFields field to given value.

### HasMappedFields

`func (o *UserExportToCsvRequest) HasMappedFields() bool`

HasMappedFields returns a boolean if a field has been set.

### GetMetadataProfileId

`func (o *UserExportToCsvRequest) GetMetadataProfileId() int32`

GetMetadataProfileId returns the MetadataProfileId field if non-nil, zero value otherwise.

### GetMetadataProfileIdOk

`func (o *UserExportToCsvRequest) GetMetadataProfileIdOk() (*int32, bool)`

GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileId

`func (o *UserExportToCsvRequest) SetMetadataProfileId(v int32)`

SetMetadataProfileId sets MetadataProfileId field to given value.

### HasMetadataProfileId

`func (o *UserExportToCsvRequest) HasMetadataProfileId() bool`

HasMetadataProfileId returns a boolean if a field has been set.

### GetOptions

`func (o *UserExportToCsvRequest) GetOptions() KalturaExportToCsvOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UserExportToCsvRequest) GetOptionsOk() (*KalturaExportToCsvOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UserExportToCsvRequest) SetOptions(v KalturaExportToCsvOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UserExportToCsvRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


