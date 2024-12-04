# RefundPaymentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | **string** |  A unique string that identifies this &#x60;RefundPayment&#x60; request. The key can be any valid string but must be unique for every &#x60;RefundPayment&#x60; request.  Keys are limited to a max of 45 characters - however, the number of allowed characters might be less than 45, if multi-byte characters are used.  For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency). | [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**AppFeeMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**PaymentId** | **string** | The unique ID of the payment being refunded. Required when unlinked&#x3D;false, otherwise must not be set. | [optional] [default to null]
**DestinationId** | **string** | The ID indicating where funds will be refunded to. Required for unlinked refunds. For more information, see [Process an Unlinked Refund](https://developer.squareup.com/docs/refunds-api/unlinked-refunds).  For refunds linked to Square payments, &#x60;destination_id&#x60; is usually omitted; in this case, funds will be returned to the original payment source. The field may be specified in order to request a cross-method refund to a gift card. For more information, see [Cross-method refunds to gift cards](https://developer.squareup.com/docs/payments-api/refund-payments#cross-method-refunds-to-gift-cards). | [optional] [default to null]
**Unlinked** | **bool** | Indicates that the refund is not linked to a Square payment. If set to true, &#x60;destination_id&#x60; and &#x60;location_id&#x60; must be supplied while &#x60;payment_id&#x60; must not be provided. | [optional] [default to null]
**LocationId** | **string** | The location ID associated with the unlinked refund. Required for requests specifying &#x60;unlinked&#x3D;true&#x60;. Otherwise, if included when &#x60;unlinked&#x3D;false&#x60;, will throw an error. | [optional] [default to null]
**CustomerId** | **string** | The [Customer](entity:Customer) ID of the customer associated with the refund. This is required if the &#x60;destination_id&#x60; refers to a card on file created using the Cards API. Only allowed when &#x60;unlinked&#x3D;true&#x60;. | [optional] [default to null]
**Reason** | **string** | A description of the reason for the refund. | [optional] [default to null]
**PaymentVersionToken** | **string** |  Used for optimistic concurrency. This opaque token identifies the current &#x60;Payment&#x60; version that the caller expects. If the server has a different version of the Payment, the update fails and a response with a VERSION_MISMATCH error is returned. If the versions match, or the field is not provided, the refund proceeds as normal. | [optional] [default to null]
**TeamMemberId** | **string** | An optional [TeamMember](entity:TeamMember) ID to associate with this refund. | [optional] [default to null]
**CashDetails** | [***DestinationDetailsCashRefundDetails**](DestinationDetailsCashRefundDetails.md) |  | [optional] [default to null]
**ExternalDetails** | [***DestinationDetailsExternalRefundDetails**](DestinationDetailsExternalRefundDetails.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

