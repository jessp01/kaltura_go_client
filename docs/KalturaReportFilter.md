# KalturaReportFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | Pointer to **string** | The dimension whose values should be filtered | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **string** | The (comma separated) values to include in the filter | [optional] 

## Methods

### NewKalturaReportFilter

`func NewKalturaReportFilter() *KalturaReportFilter`

NewKalturaReportFilter instantiates a new KalturaReportFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaReportFilterWithDefaults

`func NewKalturaReportFilterWithDefaults() *KalturaReportFilter`

NewKalturaReportFilterWithDefaults instantiates a new KalturaReportFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *KalturaReportFilter) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *KalturaReportFilter) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *KalturaReportFilter) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *KalturaReportFilter) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetOrderBy

`func (o *KalturaReportFilter) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *KalturaReportFilter) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *KalturaReportFilter) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *KalturaReportFilter) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetValues

`func (o *KalturaReportFilter) GetValues() string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *KalturaReportFilter) GetValuesOk() (*string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *KalturaReportFilter) SetValues(v string)`

SetValues sets Values field to given value.

### HasValues

`func (o *KalturaReportFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


