# ScheduleResourceListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaScheduleResourceFilter**](KalturaScheduleResourceFilter.md) |  | [optional] 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 

## Methods

### NewScheduleResourceListRequest

`func NewScheduleResourceListRequest() *ScheduleResourceListRequest`

NewScheduleResourceListRequest instantiates a new ScheduleResourceListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleResourceListRequestWithDefaults

`func NewScheduleResourceListRequestWithDefaults() *ScheduleResourceListRequest`

NewScheduleResourceListRequestWithDefaults instantiates a new ScheduleResourceListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ScheduleResourceListRequest) GetFilter() KalturaScheduleResourceFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ScheduleResourceListRequest) GetFilterOk() (*KalturaScheduleResourceFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ScheduleResourceListRequest) SetFilter(v KalturaScheduleResourceFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ScheduleResourceListRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPager

`func (o *ScheduleResourceListRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *ScheduleResourceListRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *ScheduleResourceListRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *ScheduleResourceListRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


