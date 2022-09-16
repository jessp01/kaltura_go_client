# ResponseProfileCloneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Profile** | [**KalturaResponseProfile**](KalturaResponseProfile.md) |  | 

## Methods

### NewResponseProfileCloneRequest

`func NewResponseProfileCloneRequest(id int32, profile KalturaResponseProfile, ) *ResponseProfileCloneRequest`

NewResponseProfileCloneRequest instantiates a new ResponseProfileCloneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseProfileCloneRequestWithDefaults

`func NewResponseProfileCloneRequestWithDefaults() *ResponseProfileCloneRequest`

NewResponseProfileCloneRequestWithDefaults instantiates a new ResponseProfileCloneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseProfileCloneRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseProfileCloneRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseProfileCloneRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetProfile

`func (o *ResponseProfileCloneRequest) GetProfile() KalturaResponseProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ResponseProfileCloneRequest) GetProfileOk() (*KalturaResponseProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ResponseProfileCloneRequest) SetProfile(v KalturaResponseProfile)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


