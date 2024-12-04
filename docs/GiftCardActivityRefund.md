# GiftCardActivityRefund

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedeemActivityId** | **string** | The ID of the refunded &#x60;REDEEM&#x60; gift card activity. Square populates this field if the  &#x60;payment_id&#x60; in the corresponding [RefundPayment](api-endpoint:Refunds-RefundPayment) request  represents a gift card redemption.  For applications that use a custom payment processing system, this field is required when creating a &#x60;REFUND&#x60; activity. The provided &#x60;REDEEM&#x60; activity ID must be linked to the same gift card. | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**ReferenceId** | **string** | A client-specified ID that associates the gift card activity with an entity in another system. | [optional] [default to null]
**PaymentId** | **string** | The ID of the refunded payment. Square populates this field if the refund is for a  payment processed by Square. This field matches the &#x60;payment_id&#x60; in the corresponding [RefundPayment](api-endpoint:Refunds-RefundPayment) request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

