# KalturaReportExportParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** |  | [optional] 
**RecipientEmail** | Pointer to **string** |  | [optional] 
**ReportItems** | Pointer to [**[]KalturaReportExportItem**](KalturaReportExportItem.md) |  | [optional] 
**ReportsItemsGroup** | Pointer to **string** |  | [optional] 
**TimeZoneOffset** | Pointer to **int32** | Time zone offset in minutes (between client to UTC) | [optional] 

## Methods

### NewKalturaReportExportParams

`func NewKalturaReportExportParams() *KalturaReportExportParams`

NewKalturaReportExportParams instantiates a new KalturaReportExportParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaReportExportParamsWithDefaults

`func NewKalturaReportExportParamsWithDefaults() *KalturaReportExportParams`

NewKalturaReportExportParamsWithDefaults instantiates a new KalturaReportExportParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *KalturaReportExportParams) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *KalturaReportExportParams) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *KalturaReportExportParams) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *KalturaReportExportParams) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetRecipientEmail

`func (o *KalturaReportExportParams) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *KalturaReportExportParams) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *KalturaReportExportParams) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *KalturaReportExportParams) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### GetReportItems

`func (o *KalturaReportExportParams) GetReportItems() []KalturaReportExportItem`

GetReportItems returns the ReportItems field if non-nil, zero value otherwise.

### GetReportItemsOk

`func (o *KalturaReportExportParams) GetReportItemsOk() (*[]KalturaReportExportItem, bool)`

GetReportItemsOk returns a tuple with the ReportItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportItems

`func (o *KalturaReportExportParams) SetReportItems(v []KalturaReportExportItem)`

SetReportItems sets ReportItems field to given value.

### HasReportItems

`func (o *KalturaReportExportParams) HasReportItems() bool`

HasReportItems returns a boolean if a field has been set.

### GetReportsItemsGroup

`func (o *KalturaReportExportParams) GetReportsItemsGroup() string`

GetReportsItemsGroup returns the ReportsItemsGroup field if non-nil, zero value otherwise.

### GetReportsItemsGroupOk

`func (o *KalturaReportExportParams) GetReportsItemsGroupOk() (*string, bool)`

GetReportsItemsGroupOk returns a tuple with the ReportsItemsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsItemsGroup

`func (o *KalturaReportExportParams) SetReportsItemsGroup(v string)`

SetReportsItemsGroup sets ReportsItemsGroup field to given value.

### HasReportsItemsGroup

`func (o *KalturaReportExportParams) HasReportsItemsGroup() bool`

HasReportsItemsGroup returns a boolean if a field has been set.

### GetTimeZoneOffset

`func (o *KalturaReportExportParams) GetTimeZoneOffset() int32`

GetTimeZoneOffset returns the TimeZoneOffset field if non-nil, zero value otherwise.

### GetTimeZoneOffsetOk

`func (o *KalturaReportExportParams) GetTimeZoneOffsetOk() (*int32, bool)`

GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneOffset

`func (o *KalturaReportExportParams) SetTimeZoneOffset(v int32)`

SetTimeZoneOffset sets TimeZoneOffset field to given value.

### HasTimeZoneOffset

`func (o *KalturaReportExportParams) HasTimeZoneOffset() bool`

HasTimeZoneOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


