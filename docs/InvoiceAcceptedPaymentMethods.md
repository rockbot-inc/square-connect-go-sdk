# InvoiceAcceptedPaymentMethods

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | **bool** | Indicates whether credit card or debit card payments are accepted. The default value is &#x60;false&#x60;. | [optional] [default to null]
**SquareGiftCard** | **bool** | Indicates whether Square gift card payments are accepted. The default value is &#x60;false&#x60;. | [optional] [default to null]
**BankAccount** | **bool** | Indicates whether ACH bank transfer payments are accepted. The default value is &#x60;false&#x60;. | [optional] [default to null]
**BuyNowPayLater** | **bool** | Indicates whether Afterpay (also known as Clearpay) payments are accepted. The default value is &#x60;false&#x60;.  This option is allowed only for invoices that have a single payment request of the &#x60;BALANCE&#x60; type. This payment method is supported if the seller account accepts Afterpay payments and the seller location is in a country where Afterpay invoice payments are supported. As a best practice, consider enabling an additional payment method when allowing &#x60;buy_now_pay_later&#x60; payments. For more information, including detailed requirements and processing limits, see [Buy Now Pay Later payments with Afterpay](https://developer.squareup.com/docs/invoices-api/overview#buy-now-pay-later). | [optional] [default to null]
**CashAppPay** | **bool** | Indicates whether Cash App payments are accepted. The default value is &#x60;false&#x60;.  This payment method is supported only for seller [locations](entity:Location) in the United States. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

