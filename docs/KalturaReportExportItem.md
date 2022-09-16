# KalturaReportExportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **int32** | Enum Type: &#x60;KalturaReportExportItemType&#x60; | [optional] 
**Filter** | Pointer to [**KalturaReportInputFilter**](KalturaReportInputFilter.md) |  | [optional] 
**ObjectIds** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**ReportTitle** | Pointer to **string** |  | [optional] 
**ReportType** | Pointer to **string** | Enum Type: &#x60;KalturaReportType&#x60; | [optional] 
**ResponseOptions** | Pointer to [**KalturaReportResponseOptions**](KalturaReportResponseOptions.md) |  | [optional] 

## Methods

### NewKalturaReportExportItem

`func NewKalturaReportExportItem() *KalturaReportExportItem`

NewKalturaReportExportItem instantiates a new KalturaReportExportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaReportExportItemWithDefaults

`func NewKalturaReportExportItemWithDefaults() *KalturaReportExportItem`

NewKalturaReportExportItemWithDefaults instantiates a new KalturaReportExportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *KalturaReportExportItem) GetAction() int32`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *KalturaReportExportItem) GetActionOk() (*int32, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *KalturaReportExportItem) SetAction(v int32)`

SetAction sets Action field to given value.

### HasAction

`func (o *KalturaReportExportItem) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetFilter

`func (o *KalturaReportExportItem) GetFilter() KalturaReportInputFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *KalturaReportExportItem) GetFilterOk() (*KalturaReportInputFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *KalturaReportExportItem) SetFilter(v KalturaReportInputFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *KalturaReportExportItem) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetObjectIds

`func (o *KalturaReportExportItem) GetObjectIds() string`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *KalturaReportExportItem) GetObjectIdsOk() (*string, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *KalturaReportExportItem) SetObjectIds(v string)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *KalturaReportExportItem) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.

### GetOrder

`func (o *KalturaReportExportItem) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *KalturaReportExportItem) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *KalturaReportExportItem) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *KalturaReportExportItem) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReportTitle

`func (o *KalturaReportExportItem) GetReportTitle() string`

GetReportTitle returns the ReportTitle field if non-nil, zero value otherwise.

### GetReportTitleOk

`func (o *KalturaReportExportItem) GetReportTitleOk() (*string, bool)`

GetReportTitleOk returns a tuple with the ReportTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTitle

`func (o *KalturaReportExportItem) SetReportTitle(v string)`

SetReportTitle sets ReportTitle field to given value.

### HasReportTitle

`func (o *KalturaReportExportItem) HasReportTitle() bool`

HasReportTitle returns a boolean if a field has been set.

### GetReportType

`func (o *KalturaReportExportItem) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *KalturaReportExportItem) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *KalturaReportExportItem) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *KalturaReportExportItem) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetResponseOptions

`func (o *KalturaReportExportItem) GetResponseOptions() KalturaReportResponseOptions`

GetResponseOptions returns the ResponseOptions field if non-nil, zero value otherwise.

### GetResponseOptionsOk

`func (o *KalturaReportExportItem) GetResponseOptionsOk() (*KalturaReportResponseOptions, bool)`

GetResponseOptionsOk returns a tuple with the ResponseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseOptions

`func (o *KalturaReportExportItem) SetResponseOptions(v KalturaReportResponseOptions)`

SetResponseOptions sets ResponseOptions field to given value.

### HasResponseOptions

`func (o *KalturaReportExportItem) HasResponseOptions() bool`

HasResponseOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


