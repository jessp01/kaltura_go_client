# KalturaExportToCsvOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultHeader** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**Format** | Pointer to **string** | The format of the outputted date string. There are also several predefined date constants that may be used instead, so for example DATE_RSS contains the format string &#39;D, d M Y H:i:s&#39;.  https://www.php.net/manual/en/function.date.php | [optional] 
**TypeEqual** | Pointer to **string** | Enum Type: &#x60;KalturaEntryType&#x60;  Setting this property will cause additional columns to be added to the final report. The columns will be related to the specific object type passed (currently only MEDIA_CLIP is supported).  Please note that this property will NOT change the result filter in any way (i.e passing MEDIA_CLIP here will not force the report to return only media items). | [optional] 

## Methods

### NewKalturaExportToCsvOptions

`func NewKalturaExportToCsvOptions() *KalturaExportToCsvOptions`

NewKalturaExportToCsvOptions instantiates a new KalturaExportToCsvOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaExportToCsvOptionsWithDefaults

`func NewKalturaExportToCsvOptionsWithDefaults() *KalturaExportToCsvOptions`

NewKalturaExportToCsvOptionsWithDefaults instantiates a new KalturaExportToCsvOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultHeader

`func (o *KalturaExportToCsvOptions) GetDefaultHeader() int32`

GetDefaultHeader returns the DefaultHeader field if non-nil, zero value otherwise.

### GetDefaultHeaderOk

`func (o *KalturaExportToCsvOptions) GetDefaultHeaderOk() (*int32, bool)`

GetDefaultHeaderOk returns a tuple with the DefaultHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultHeader

`func (o *KalturaExportToCsvOptions) SetDefaultHeader(v int32)`

SetDefaultHeader sets DefaultHeader field to given value.

### HasDefaultHeader

`func (o *KalturaExportToCsvOptions) HasDefaultHeader() bool`

HasDefaultHeader returns a boolean if a field has been set.

### GetFormat

`func (o *KalturaExportToCsvOptions) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *KalturaExportToCsvOptions) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *KalturaExportToCsvOptions) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *KalturaExportToCsvOptions) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetTypeEqual

`func (o *KalturaExportToCsvOptions) GetTypeEqual() string`

GetTypeEqual returns the TypeEqual field if non-nil, zero value otherwise.

### GetTypeEqualOk

`func (o *KalturaExportToCsvOptions) GetTypeEqualOk() (*string, bool)`

GetTypeEqualOk returns a tuple with the TypeEqual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeEqual

`func (o *KalturaExportToCsvOptions) SetTypeEqual(v string)`

SetTypeEqual sets TypeEqual field to given value.

### HasTypeEqual

`func (o *KalturaExportToCsvOptions) HasTypeEqual() bool`

HasTypeEqual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


