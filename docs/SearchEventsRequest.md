# SearchEventsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of events for your original query.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | [optional] [default to null]
**Limit** | **int32** | The maximum number of events to return in a single page. The response might contain fewer events. The default value is 100, which is also the maximum allowed value.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).  Default: 100 | [optional] [default to null]
**Query** | [***SearchEventsQuery**](SearchEventsQuery.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

