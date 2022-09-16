# KalturaLiveReportExportParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationUrlTemplate** | Pointer to **string** | Optional argument that allows controlling the prefix of the exported csv url | [optional] 
**EntryIds** | Pointer to **string** |  | [optional] 
**RecpientEmail** | Pointer to **string** |  | [optional] 
**TimeZoneOffset** | Pointer to **int32** | Time zone offset in minutes (between client to UTC) | [optional] 

## Methods

### NewKalturaLiveReportExportParams

`func NewKalturaLiveReportExportParams() *KalturaLiveReportExportParams`

NewKalturaLiveReportExportParams instantiates a new KalturaLiveReportExportParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaLiveReportExportParamsWithDefaults

`func NewKalturaLiveReportExportParamsWithDefaults() *KalturaLiveReportExportParams`

NewKalturaLiveReportExportParamsWithDefaults instantiates a new KalturaLiveReportExportParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationUrlTemplate

`func (o *KalturaLiveReportExportParams) GetApplicationUrlTemplate() string`

GetApplicationUrlTemplate returns the ApplicationUrlTemplate field if non-nil, zero value otherwise.

### GetApplicationUrlTemplateOk

`func (o *KalturaLiveReportExportParams) GetApplicationUrlTemplateOk() (*string, bool)`

GetApplicationUrlTemplateOk returns a tuple with the ApplicationUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUrlTemplate

`func (o *KalturaLiveReportExportParams) SetApplicationUrlTemplate(v string)`

SetApplicationUrlTemplate sets ApplicationUrlTemplate field to given value.

### HasApplicationUrlTemplate

`func (o *KalturaLiveReportExportParams) HasApplicationUrlTemplate() bool`

HasApplicationUrlTemplate returns a boolean if a field has been set.

### GetEntryIds

`func (o *KalturaLiveReportExportParams) GetEntryIds() string`

GetEntryIds returns the EntryIds field if non-nil, zero value otherwise.

### GetEntryIdsOk

`func (o *KalturaLiveReportExportParams) GetEntryIdsOk() (*string, bool)`

GetEntryIdsOk returns a tuple with the EntryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryIds

`func (o *KalturaLiveReportExportParams) SetEntryIds(v string)`

SetEntryIds sets EntryIds field to given value.

### HasEntryIds

`func (o *KalturaLiveReportExportParams) HasEntryIds() bool`

HasEntryIds returns a boolean if a field has been set.

### GetRecpientEmail

`func (o *KalturaLiveReportExportParams) GetRecpientEmail() string`

GetRecpientEmail returns the RecpientEmail field if non-nil, zero value otherwise.

### GetRecpientEmailOk

`func (o *KalturaLiveReportExportParams) GetRecpientEmailOk() (*string, bool)`

GetRecpientEmailOk returns a tuple with the RecpientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecpientEmail

`func (o *KalturaLiveReportExportParams) SetRecpientEmail(v string)`

SetRecpientEmail sets RecpientEmail field to given value.

### HasRecpientEmail

`func (o *KalturaLiveReportExportParams) HasRecpientEmail() bool`

HasRecpientEmail returns a boolean if a field has been set.

### GetTimeZoneOffset

`func (o *KalturaLiveReportExportParams) GetTimeZoneOffset() int32`

GetTimeZoneOffset returns the TimeZoneOffset field if non-nil, zero value otherwise.

### GetTimeZoneOffsetOk

`func (o *KalturaLiveReportExportParams) GetTimeZoneOffsetOk() (*int32, bool)`

GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneOffset

`func (o *KalturaLiveReportExportParams) SetTimeZoneOffset(v int32)`

SetTimeZoneOffset sets TimeZoneOffset field to given value.

### HasTimeZoneOffset

`func (o *KalturaLiveReportExportParams) HasTimeZoneOffset() bool`

HasTimeZoneOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


