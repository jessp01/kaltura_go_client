# EmailIngestionProfileUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailIP** | [**KalturaEmailIngestionProfile**](KalturaEmailIngestionProfile.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewEmailIngestionProfileUpdateRequest

`func NewEmailIngestionProfileUpdateRequest(emailIP KalturaEmailIngestionProfile, id int32, ) *EmailIngestionProfileUpdateRequest`

NewEmailIngestionProfileUpdateRequest instantiates a new EmailIngestionProfileUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailIngestionProfileUpdateRequestWithDefaults

`func NewEmailIngestionProfileUpdateRequestWithDefaults() *EmailIngestionProfileUpdateRequest`

NewEmailIngestionProfileUpdateRequestWithDefaults instantiates a new EmailIngestionProfileUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailIP

`func (o *EmailIngestionProfileUpdateRequest) GetEmailIP() KalturaEmailIngestionProfile`

GetEmailIP returns the EmailIP field if non-nil, zero value otherwise.

### GetEmailIPOk

`func (o *EmailIngestionProfileUpdateRequest) GetEmailIPOk() (*KalturaEmailIngestionProfile, bool)`

GetEmailIPOk returns a tuple with the EmailIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIP

`func (o *EmailIngestionProfileUpdateRequest) SetEmailIP(v KalturaEmailIngestionProfile)`

SetEmailIP sets EmailIP field to given value.


### GetId

`func (o *EmailIngestionProfileUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailIngestionProfileUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailIngestionProfileUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


