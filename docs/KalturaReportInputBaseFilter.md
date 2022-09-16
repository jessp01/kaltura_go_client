# KalturaReportInputBaseFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromDate** | Pointer to **int32** | Start date as Unix timestamp (In seconds) | [optional] 
**FromDay** | Pointer to **string** | Start day as string (YYYYMMDD) | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**ToDate** | Pointer to **int32** | End date as Unix timestamp (In seconds) | [optional] 
**ToDay** | Pointer to **string** | End date as string (YYYYMMDD) | [optional] 

## Methods

### NewKalturaReportInputBaseFilter

`func NewKalturaReportInputBaseFilter() *KalturaReportInputBaseFilter`

NewKalturaReportInputBaseFilter instantiates a new KalturaReportInputBaseFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaReportInputBaseFilterWithDefaults

`func NewKalturaReportInputBaseFilterWithDefaults() *KalturaReportInputBaseFilter`

NewKalturaReportInputBaseFilterWithDefaults instantiates a new KalturaReportInputBaseFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromDate

`func (o *KalturaReportInputBaseFilter) GetFromDate() int32`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *KalturaReportInputBaseFilter) GetFromDateOk() (*int32, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *KalturaReportInputBaseFilter) SetFromDate(v int32)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *KalturaReportInputBaseFilter) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetFromDay

`func (o *KalturaReportInputBaseFilter) GetFromDay() string`

GetFromDay returns the FromDay field if non-nil, zero value otherwise.

### GetFromDayOk

`func (o *KalturaReportInputBaseFilter) GetFromDayOk() (*string, bool)`

GetFromDayOk returns a tuple with the FromDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDay

`func (o *KalturaReportInputBaseFilter) SetFromDay(v string)`

SetFromDay sets FromDay field to given value.

### HasFromDay

`func (o *KalturaReportInputBaseFilter) HasFromDay() bool`

HasFromDay returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaReportInputBaseFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaReportInputBaseFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaReportInputBaseFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaReportInputBaseFilter) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetToDate

`func (o *KalturaReportInputBaseFilter) GetToDate() int32`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *KalturaReportInputBaseFilter) GetToDateOk() (*int32, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *KalturaReportInputBaseFilter) SetToDate(v int32)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *KalturaReportInputBaseFilter) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetToDay

`func (o *KalturaReportInputBaseFilter) GetToDay() string`

GetToDay returns the ToDay field if non-nil, zero value otherwise.

### GetToDayOk

`func (o *KalturaReportInputBaseFilter) GetToDayOk() (*string, bool)`

GetToDayOk returns a tuple with the ToDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDay

`func (o *KalturaReportInputBaseFilter) SetToDay(v string)`

SetToDay sets ToDay field to given value.

### HasToDay

`func (o *KalturaReportInputBaseFilter) HasToDay() bool`

HasToDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


