# ShortLinkGotoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Proxy** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewShortLinkGotoRequest

`func NewShortLinkGotoRequest(id string, ) *ShortLinkGotoRequest`

NewShortLinkGotoRequest instantiates a new ShortLinkGotoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortLinkGotoRequestWithDefaults

`func NewShortLinkGotoRequestWithDefaults() *ShortLinkGotoRequest`

NewShortLinkGotoRequestWithDefaults instantiates a new ShortLinkGotoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShortLinkGotoRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShortLinkGotoRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShortLinkGotoRequest) SetId(v string)`

SetId sets Id field to given value.


### GetProxy

`func (o *ShortLinkGotoRequest) GetProxy() bool`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ShortLinkGotoRequest) GetProxyOk() (*bool, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ShortLinkGotoRequest) SetProxy(v bool)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ShortLinkGotoRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


