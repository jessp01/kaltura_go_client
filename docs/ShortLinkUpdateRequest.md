# ShortLinkUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ShortLink** | [**KalturaShortLink**](KalturaShortLink.md) |  | 

## Methods

### NewShortLinkUpdateRequest

`func NewShortLinkUpdateRequest(id string, shortLink KalturaShortLink, ) *ShortLinkUpdateRequest`

NewShortLinkUpdateRequest instantiates a new ShortLinkUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortLinkUpdateRequestWithDefaults

`func NewShortLinkUpdateRequestWithDefaults() *ShortLinkUpdateRequest`

NewShortLinkUpdateRequestWithDefaults instantiates a new ShortLinkUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShortLinkUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShortLinkUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShortLinkUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetShortLink

`func (o *ShortLinkUpdateRequest) GetShortLink() KalturaShortLink`

GetShortLink returns the ShortLink field if non-nil, zero value otherwise.

### GetShortLinkOk

`func (o *ShortLinkUpdateRequest) GetShortLinkOk() (*KalturaShortLink, bool)`

GetShortLinkOk returns a tuple with the ShortLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortLink

`func (o *ShortLinkUpdateRequest) SetShortLink(v KalturaShortLink)`

SetShortLink sets ShortLink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


