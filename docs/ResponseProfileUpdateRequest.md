# ResponseProfileUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**UpdateResponseProfile** | [**KalturaResponseProfile**](KalturaResponseProfile.md) |  | 

## Methods

### NewResponseProfileUpdateRequest

`func NewResponseProfileUpdateRequest(id int32, updateResponseProfile KalturaResponseProfile, ) *ResponseProfileUpdateRequest`

NewResponseProfileUpdateRequest instantiates a new ResponseProfileUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseProfileUpdateRequestWithDefaults

`func NewResponseProfileUpdateRequestWithDefaults() *ResponseProfileUpdateRequest`

NewResponseProfileUpdateRequestWithDefaults instantiates a new ResponseProfileUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseProfileUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseProfileUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseProfileUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetUpdateResponseProfile

`func (o *ResponseProfileUpdateRequest) GetUpdateResponseProfile() KalturaResponseProfile`

GetUpdateResponseProfile returns the UpdateResponseProfile field if non-nil, zero value otherwise.

### GetUpdateResponseProfileOk

`func (o *ResponseProfileUpdateRequest) GetUpdateResponseProfileOk() (*KalturaResponseProfile, bool)`

GetUpdateResponseProfileOk returns a tuple with the UpdateResponseProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateResponseProfile

`func (o *ResponseProfileUpdateRequest) SetUpdateResponseProfile(v KalturaResponseProfile)`

SetUpdateResponseProfile sets UpdateResponseProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


