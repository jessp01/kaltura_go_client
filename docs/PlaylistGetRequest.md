# PlaylistGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewPlaylistGetRequest

`func NewPlaylistGetRequest(id string, ) *PlaylistGetRequest`

NewPlaylistGetRequest instantiates a new PlaylistGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistGetRequestWithDefaults

`func NewPlaylistGetRequestWithDefaults() *PlaylistGetRequest`

NewPlaylistGetRequestWithDefaults instantiates a new PlaylistGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlaylistGetRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaylistGetRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaylistGetRequest) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *PlaylistGetRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PlaylistGetRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PlaylistGetRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PlaylistGetRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


