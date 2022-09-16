# EmailIngestionProfileAddMediaEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailMsgId** | **string** |  | 
**EmailProfId** | **int32** |  | 
**FromAddress** | **string** |  | 
**MediaEntry** | [**KalturaMediaEntry**](KalturaMediaEntry.md) |  | 
**UploadTokenId** | **string** |  | 

## Methods

### NewEmailIngestionProfileAddMediaEntryRequest

`func NewEmailIngestionProfileAddMediaEntryRequest(emailMsgId string, emailProfId int32, fromAddress string, mediaEntry KalturaMediaEntry, uploadTokenId string, ) *EmailIngestionProfileAddMediaEntryRequest`

NewEmailIngestionProfileAddMediaEntryRequest instantiates a new EmailIngestionProfileAddMediaEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailIngestionProfileAddMediaEntryRequestWithDefaults

`func NewEmailIngestionProfileAddMediaEntryRequestWithDefaults() *EmailIngestionProfileAddMediaEntryRequest`

NewEmailIngestionProfileAddMediaEntryRequestWithDefaults instantiates a new EmailIngestionProfileAddMediaEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailMsgId

`func (o *EmailIngestionProfileAddMediaEntryRequest) GetEmailMsgId() string`

GetEmailMsgId returns the EmailMsgId field if non-nil, zero value otherwise.

### GetEmailMsgIdOk

`func (o *EmailIngestionProfileAddMediaEntryRequest) GetEmailMsgIdOk() (*string, bool)`

GetEmailMsgIdOk returns a tuple with the EmailMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMsgId

`func (o *EmailIngestionProfileAddMediaEntryRequest) SetEmailMsgId(v string)`

SetEmailMsgId sets EmailMsgId field to given value.


### GetEmailProfId

`func (o *EmailIngestionProfileAddMediaEntryRequest) GetEmailProfId() int32`

GetEmailProfId returns the EmailProfId field if non-nil, zero value otherwise.

### GetEmailProfIdOk

`func (o *EmailIngestionProfileAddMediaEntryRequest) GetEmailProfIdOk() (*int32, bool)`

GetEmailProfIdOk returns a tuple with the EmailProfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailProfId

`func (o *EmailIngestionProfileAddMediaEntryRequest) SetEmailProfId(v int32)`

SetEmailProfId sets EmailProfId field to given value.


### GetFromAddress

`func (o *EmailIngestionProfileAddMediaEntryRequest) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *EmailIngestionProfileAddMediaEntryRequest) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *EmailIngestionProfileAddMediaEntryRequest) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetMediaEntry

`func (o *EmailIngestionProfileAddMediaEntryRequest) GetMediaEntry() KalturaMediaEntry`

GetMediaEntry returns the MediaEntry field if non-nil, zero value otherwise.

### GetMediaEntryOk

`func (o *EmailIngestionProfileAddMediaEntryRequest) GetMediaEntryOk() (*KalturaMediaEntry, bool)`

GetMediaEntryOk returns a tuple with the MediaEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEntry

`func (o *EmailIngestionProfileAddMediaEntryRequest) SetMediaEntry(v KalturaMediaEntry)`

SetMediaEntry sets MediaEntry field to given value.


### GetUploadTokenId

`func (o *EmailIngestionProfileAddMediaEntryRequest) GetUploadTokenId() string`

GetUploadTokenId returns the UploadTokenId field if non-nil, zero value otherwise.

### GetUploadTokenIdOk

`func (o *EmailIngestionProfileAddMediaEntryRequest) GetUploadTokenIdOk() (*string, bool)`

GetUploadTokenIdOk returns a tuple with the UploadTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTokenId

`func (o *EmailIngestionProfileAddMediaEntryRequest) SetUploadTokenId(v string)`

SetUploadTokenId sets UploadTokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


