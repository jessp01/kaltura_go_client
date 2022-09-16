# KalturaVendorTaskData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryDuration** | Pointer to **int32** | &#x60;readOnly&#x60;  The duration of the entry for which the task was created for in milliseconds | [optional] [readonly] 
**ObjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaVendorTaskData

`func NewKalturaVendorTaskData() *KalturaVendorTaskData`

NewKalturaVendorTaskData instantiates a new KalturaVendorTaskData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaVendorTaskDataWithDefaults

`func NewKalturaVendorTaskDataWithDefaults() *KalturaVendorTaskData`

NewKalturaVendorTaskDataWithDefaults instantiates a new KalturaVendorTaskData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryDuration

`func (o *KalturaVendorTaskData) GetEntryDuration() int32`

GetEntryDuration returns the EntryDuration field if non-nil, zero value otherwise.

### GetEntryDurationOk

`func (o *KalturaVendorTaskData) GetEntryDurationOk() (*int32, bool)`

GetEntryDurationOk returns a tuple with the EntryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryDuration

`func (o *KalturaVendorTaskData) SetEntryDuration(v int32)`

SetEntryDuration sets EntryDuration field to given value.

### HasEntryDuration

`func (o *KalturaVendorTaskData) HasEntryDuration() bool`

HasEntryDuration returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaVendorTaskData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaVendorTaskData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaVendorTaskData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaVendorTaskData) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


