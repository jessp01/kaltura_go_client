# ScheduleEventResourceListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaScheduleEventResourceFilter**](KalturaScheduleEventResourceFilter.md) |  | [optional] 
**FilterBlackoutConflicts** | Pointer to **bool** |  | [optional] [default to true]
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 

## Methods

### NewScheduleEventResourceListRequest

`func NewScheduleEventResourceListRequest() *ScheduleEventResourceListRequest`

NewScheduleEventResourceListRequest instantiates a new ScheduleEventResourceListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEventResourceListRequestWithDefaults

`func NewScheduleEventResourceListRequestWithDefaults() *ScheduleEventResourceListRequest`

NewScheduleEventResourceListRequestWithDefaults instantiates a new ScheduleEventResourceListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ScheduleEventResourceListRequest) GetFilter() KalturaScheduleEventResourceFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ScheduleEventResourceListRequest) GetFilterOk() (*KalturaScheduleEventResourceFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ScheduleEventResourceListRequest) SetFilter(v KalturaScheduleEventResourceFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ScheduleEventResourceListRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFilterBlackoutConflicts

`func (o *ScheduleEventResourceListRequest) GetFilterBlackoutConflicts() bool`

GetFilterBlackoutConflicts returns the FilterBlackoutConflicts field if non-nil, zero value otherwise.

### GetFilterBlackoutConflictsOk

`func (o *ScheduleEventResourceListRequest) GetFilterBlackoutConflictsOk() (*bool, bool)`

GetFilterBlackoutConflictsOk returns a tuple with the FilterBlackoutConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBlackoutConflicts

`func (o *ScheduleEventResourceListRequest) SetFilterBlackoutConflicts(v bool)`

SetFilterBlackoutConflicts sets FilterBlackoutConflicts field to given value.

### HasFilterBlackoutConflicts

`func (o *ScheduleEventResourceListRequest) HasFilterBlackoutConflicts() bool`

HasFilterBlackoutConflicts returns a boolean if a field has been set.

### GetPager

`func (o *ScheduleEventResourceListRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *ScheduleEventResourceListRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *ScheduleEventResourceListRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *ScheduleEventResourceListRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


