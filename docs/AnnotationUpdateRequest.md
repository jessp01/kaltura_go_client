# AnnotationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | [**KalturaCuePoint**](KalturaCuePoint.md) |  | 
**Id** | **string** |  | 

## Methods

### NewAnnotationUpdateRequest

`func NewAnnotationUpdateRequest(annotation KalturaCuePoint, id string, ) *AnnotationUpdateRequest`

NewAnnotationUpdateRequest instantiates a new AnnotationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationUpdateRequestWithDefaults

`func NewAnnotationUpdateRequestWithDefaults() *AnnotationUpdateRequest`

NewAnnotationUpdateRequestWithDefaults instantiates a new AnnotationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *AnnotationUpdateRequest) GetAnnotation() KalturaCuePoint`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *AnnotationUpdateRequest) GetAnnotationOk() (*KalturaCuePoint, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *AnnotationUpdateRequest) SetAnnotation(v KalturaCuePoint)`

SetAnnotation sets Annotation field to given value.


### GetId

`func (o *AnnotationUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnnotationUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnnotationUpdateRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


