# KalturaDistributionProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityUpdateEnabled** | Pointer to **bool** |  | [optional] 
**DeleteInsteadUpdate** | Pointer to **bool** |  | [optional] 
**IntervalBeforeSunrise** | Pointer to **int32** |  | [optional] 
**IntervalBeforeSunset** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**ScheduleUpdateEnabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaDistributionProviderType&#x60; | [optional] [readonly] 
**UpdateRequiredEntryFields** | Pointer to **string** |  | [optional] 
**UpdateRequiredMetadataXPaths** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaDistributionProvider

`func NewKalturaDistributionProvider() *KalturaDistributionProvider`

NewKalturaDistributionProvider instantiates a new KalturaDistributionProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaDistributionProviderWithDefaults

`func NewKalturaDistributionProviderWithDefaults() *KalturaDistributionProvider`

NewKalturaDistributionProviderWithDefaults instantiates a new KalturaDistributionProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityUpdateEnabled

`func (o *KalturaDistributionProvider) GetAvailabilityUpdateEnabled() bool`

GetAvailabilityUpdateEnabled returns the AvailabilityUpdateEnabled field if non-nil, zero value otherwise.

### GetAvailabilityUpdateEnabledOk

`func (o *KalturaDistributionProvider) GetAvailabilityUpdateEnabledOk() (*bool, bool)`

GetAvailabilityUpdateEnabledOk returns a tuple with the AvailabilityUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityUpdateEnabled

`func (o *KalturaDistributionProvider) SetAvailabilityUpdateEnabled(v bool)`

SetAvailabilityUpdateEnabled sets AvailabilityUpdateEnabled field to given value.

### HasAvailabilityUpdateEnabled

`func (o *KalturaDistributionProvider) HasAvailabilityUpdateEnabled() bool`

HasAvailabilityUpdateEnabled returns a boolean if a field has been set.

### GetDeleteInsteadUpdate

`func (o *KalturaDistributionProvider) GetDeleteInsteadUpdate() bool`

GetDeleteInsteadUpdate returns the DeleteInsteadUpdate field if non-nil, zero value otherwise.

### GetDeleteInsteadUpdateOk

`func (o *KalturaDistributionProvider) GetDeleteInsteadUpdateOk() (*bool, bool)`

GetDeleteInsteadUpdateOk returns a tuple with the DeleteInsteadUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInsteadUpdate

`func (o *KalturaDistributionProvider) SetDeleteInsteadUpdate(v bool)`

SetDeleteInsteadUpdate sets DeleteInsteadUpdate field to given value.

### HasDeleteInsteadUpdate

`func (o *KalturaDistributionProvider) HasDeleteInsteadUpdate() bool`

HasDeleteInsteadUpdate returns a boolean if a field has been set.

### GetIntervalBeforeSunrise

`func (o *KalturaDistributionProvider) GetIntervalBeforeSunrise() int32`

GetIntervalBeforeSunrise returns the IntervalBeforeSunrise field if non-nil, zero value otherwise.

### GetIntervalBeforeSunriseOk

`func (o *KalturaDistributionProvider) GetIntervalBeforeSunriseOk() (*int32, bool)`

GetIntervalBeforeSunriseOk returns a tuple with the IntervalBeforeSunrise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalBeforeSunrise

`func (o *KalturaDistributionProvider) SetIntervalBeforeSunrise(v int32)`

SetIntervalBeforeSunrise sets IntervalBeforeSunrise field to given value.

### HasIntervalBeforeSunrise

`func (o *KalturaDistributionProvider) HasIntervalBeforeSunrise() bool`

HasIntervalBeforeSunrise returns a boolean if a field has been set.

### GetIntervalBeforeSunset

`func (o *KalturaDistributionProvider) GetIntervalBeforeSunset() int32`

GetIntervalBeforeSunset returns the IntervalBeforeSunset field if non-nil, zero value otherwise.

### GetIntervalBeforeSunsetOk

`func (o *KalturaDistributionProvider) GetIntervalBeforeSunsetOk() (*int32, bool)`

GetIntervalBeforeSunsetOk returns a tuple with the IntervalBeforeSunset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalBeforeSunset

`func (o *KalturaDistributionProvider) SetIntervalBeforeSunset(v int32)`

SetIntervalBeforeSunset sets IntervalBeforeSunset field to given value.

### HasIntervalBeforeSunset

`func (o *KalturaDistributionProvider) HasIntervalBeforeSunset() bool`

HasIntervalBeforeSunset returns a boolean if a field has been set.

### GetName

`func (o *KalturaDistributionProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaDistributionProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaDistributionProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaDistributionProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaDistributionProvider) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaDistributionProvider) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaDistributionProvider) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaDistributionProvider) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetScheduleUpdateEnabled

`func (o *KalturaDistributionProvider) GetScheduleUpdateEnabled() bool`

GetScheduleUpdateEnabled returns the ScheduleUpdateEnabled field if non-nil, zero value otherwise.

### GetScheduleUpdateEnabledOk

`func (o *KalturaDistributionProvider) GetScheduleUpdateEnabledOk() (*bool, bool)`

GetScheduleUpdateEnabledOk returns a tuple with the ScheduleUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUpdateEnabled

`func (o *KalturaDistributionProvider) SetScheduleUpdateEnabled(v bool)`

SetScheduleUpdateEnabled sets ScheduleUpdateEnabled field to given value.

### HasScheduleUpdateEnabled

`func (o *KalturaDistributionProvider) HasScheduleUpdateEnabled() bool`

HasScheduleUpdateEnabled returns a boolean if a field has been set.

### GetType

`func (o *KalturaDistributionProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaDistributionProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaDistributionProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaDistributionProvider) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateRequiredEntryFields

`func (o *KalturaDistributionProvider) GetUpdateRequiredEntryFields() string`

GetUpdateRequiredEntryFields returns the UpdateRequiredEntryFields field if non-nil, zero value otherwise.

### GetUpdateRequiredEntryFieldsOk

`func (o *KalturaDistributionProvider) GetUpdateRequiredEntryFieldsOk() (*string, bool)`

GetUpdateRequiredEntryFieldsOk returns a tuple with the UpdateRequiredEntryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRequiredEntryFields

`func (o *KalturaDistributionProvider) SetUpdateRequiredEntryFields(v string)`

SetUpdateRequiredEntryFields sets UpdateRequiredEntryFields field to given value.

### HasUpdateRequiredEntryFields

`func (o *KalturaDistributionProvider) HasUpdateRequiredEntryFields() bool`

HasUpdateRequiredEntryFields returns a boolean if a field has been set.

### GetUpdateRequiredMetadataXPaths

`func (o *KalturaDistributionProvider) GetUpdateRequiredMetadataXPaths() string`

GetUpdateRequiredMetadataXPaths returns the UpdateRequiredMetadataXPaths field if non-nil, zero value otherwise.

### GetUpdateRequiredMetadataXPathsOk

`func (o *KalturaDistributionProvider) GetUpdateRequiredMetadataXPathsOk() (*string, bool)`

GetUpdateRequiredMetadataXPathsOk returns a tuple with the UpdateRequiredMetadataXPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRequiredMetadataXPaths

`func (o *KalturaDistributionProvider) SetUpdateRequiredMetadataXPaths(v string)`

SetUpdateRequiredMetadataXPaths sets UpdateRequiredMetadataXPaths field to given value.

### HasUpdateRequiredMetadataXPaths

`func (o *KalturaDistributionProvider) HasUpdateRequiredMetadataXPaths() bool`

HasUpdateRequiredMetadataXPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


