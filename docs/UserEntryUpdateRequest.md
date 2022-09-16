# UserEntryUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**UserEntry** | [**KalturaUserEntry**](KalturaUserEntry.md) |  | 

## Methods

### NewUserEntryUpdateRequest

`func NewUserEntryUpdateRequest(id int32, userEntry KalturaUserEntry, ) *UserEntryUpdateRequest`

NewUserEntryUpdateRequest instantiates a new UserEntryUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEntryUpdateRequestWithDefaults

`func NewUserEntryUpdateRequestWithDefaults() *UserEntryUpdateRequest`

NewUserEntryUpdateRequestWithDefaults instantiates a new UserEntryUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserEntryUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserEntryUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserEntryUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetUserEntry

`func (o *UserEntryUpdateRequest) GetUserEntry() KalturaUserEntry`

GetUserEntry returns the UserEntry field if non-nil, zero value otherwise.

### GetUserEntryOk

`func (o *UserEntryUpdateRequest) GetUserEntryOk() (*KalturaUserEntry, bool)`

GetUserEntryOk returns a tuple with the UserEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEntry

`func (o *UserEntryUpdateRequest) SetUserEntry(v KalturaUserEntry)`

SetUserEntry sets UserEntry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


