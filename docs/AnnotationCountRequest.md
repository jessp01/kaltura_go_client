# AnnotationCountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaCuePointFilter**](KalturaCuePointFilter.md) |  | [optional] 

## Methods

### NewAnnotationCountRequest

`func NewAnnotationCountRequest() *AnnotationCountRequest`

NewAnnotationCountRequest instantiates a new AnnotationCountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationCountRequestWithDefaults

`func NewAnnotationCountRequestWithDefaults() *AnnotationCountRequest`

NewAnnotationCountRequestWithDefaults instantiates a new AnnotationCountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *AnnotationCountRequest) GetFilter() KalturaCuePointFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AnnotationCountRequest) GetFilterOk() (*KalturaCuePointFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AnnotationCountRequest) SetFilter(v KalturaCuePointFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AnnotationCountRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


