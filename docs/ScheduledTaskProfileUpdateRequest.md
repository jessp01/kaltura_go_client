# ScheduledTaskProfileUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ScheduledTaskProfile** | [**KalturaScheduledTaskProfile**](KalturaScheduledTaskProfile.md) |  | 

## Methods

### NewScheduledTaskProfileUpdateRequest

`func NewScheduledTaskProfileUpdateRequest(id int32, scheduledTaskProfile KalturaScheduledTaskProfile, ) *ScheduledTaskProfileUpdateRequest`

NewScheduledTaskProfileUpdateRequest instantiates a new ScheduledTaskProfileUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledTaskProfileUpdateRequestWithDefaults

`func NewScheduledTaskProfileUpdateRequestWithDefaults() *ScheduledTaskProfileUpdateRequest`

NewScheduledTaskProfileUpdateRequestWithDefaults instantiates a new ScheduledTaskProfileUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledTaskProfileUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledTaskProfileUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledTaskProfileUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetScheduledTaskProfile

`func (o *ScheduledTaskProfileUpdateRequest) GetScheduledTaskProfile() KalturaScheduledTaskProfile`

GetScheduledTaskProfile returns the ScheduledTaskProfile field if non-nil, zero value otherwise.

### GetScheduledTaskProfileOk

`func (o *ScheduledTaskProfileUpdateRequest) GetScheduledTaskProfileOk() (*KalturaScheduledTaskProfile, bool)`

GetScheduledTaskProfileOk returns a tuple with the ScheduledTaskProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTaskProfile

`func (o *ScheduledTaskProfileUpdateRequest) SetScheduledTaskProfile(v KalturaScheduledTaskProfile)`

SetScheduledTaskProfile sets ScheduledTaskProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


