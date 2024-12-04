# ListPaymentsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeginTime** | **string** | Indicates the start of the time range to retrieve payments for, in RFC 3339 format.   The range is determined using the &#x60;created_at&#x60; field for each Payment. Inclusive. Default: The current time minus one year. | [optional] [default to null]
**EndTime** | **string** | Indicates the end of the time range to retrieve payments for, in RFC 3339 format.  The  range is determined using the &#x60;created_at&#x60; field for each Payment.  Default: The current time. | [optional] [default to null]
**SortOrder** | **string** | The order in which results are listed by &#x60;Payment.created_at&#x60;: - &#x60;ASC&#x60; - Oldest to newest. - &#x60;DESC&#x60; - Newest to oldest (default). | [optional] [default to null]
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | [optional] [default to null]
**LocationId** | **string** | Limit results to the location supplied. By default, results are returned for the default (main) location associated with the seller. | [optional] [default to null]
**Total** | **int64** | The exact amount in the &#x60;total_money&#x60; for a payment. | [optional] [default to null]
**Last4** | **string** | The last four digits of a payment card. | [optional] [default to null]
**CardBrand** | **string** | The brand of the payment card (for example, VISA). | [optional] [default to null]
**Limit** | **int32** | The maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page.  The default value of 100 is also the maximum allowed value. If the provided value is  greater than 100, it is ignored and the default value is used instead.  Default: &#x60;100&#x60; | [optional] [default to null]
**IsOfflinePayment** | **bool** | Whether the payment was taken offline or not. | [optional] [default to null]
**OfflineBeginTime** | **string** | Indicates the start of the time range for which to retrieve offline payments, in RFC 3339 format for timestamps. The range is determined using the &#x60;offline_payment_details.client_created_at&#x60; field for each Payment. If set, payments without a value set in &#x60;offline_payment_details.client_created_at&#x60; will not be returned.  Default: The current time. | [optional] [default to null]
**OfflineEndTime** | **string** | Indicates the end of the time range for which to retrieve offline payments, in RFC 3339 format for timestamps. The range is determined using the &#x60;offline_payment_details.client_created_at&#x60; field for each Payment. If set, payments without a value set in &#x60;offline_payment_details.client_created_at&#x60; will not be returned.  Default: The current time. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

