/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Represents a payment processed by the Square API.
type Payment struct {
	// A unique ID for the payment.
	Id string `json:"id,omitempty"`
	// The timestamp of when the payment was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp of when the payment was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
	AmountMoney *Money `json:"amount_money,omitempty"`
	TipMoney *Money `json:"tip_money,omitempty"`
	TotalMoney *Money `json:"total_money,omitempty"`
	AppFeeMoney *Money `json:"app_fee_money,omitempty"`
	ApprovedMoney *Money `json:"approved_money,omitempty"`
	// The processing fees and fee adjustments assessed by Square for this payment.
	ProcessingFee []ProcessingFee `json:"processing_fee,omitempty"`
	RefundedMoney *Money `json:"refunded_money,omitempty"`
	// Indicates whether the payment is APPROVED, PENDING, COMPLETED, CANCELED, or FAILED.
	Status string `json:"status,omitempty"`
	// The duration of time after the payment's creation when Square automatically applies the `delay_action` to the payment. This automatic `delay_action` applies only to payments that do not reach a terminal state (COMPLETED, CANCELED, or FAILED) before the `delay_duration` time period.  This field is specified as a time duration, in RFC 3339 format.  Notes: This feature is only supported for card payments.  Default:  - Card-present payments: \"PT36H\" (36 hours) from the creation time. - Card-not-present payments: \"P7D\" (7 days) from the creation time.
	DelayDuration string `json:"delay_duration,omitempty"`
	// The action to be applied to the payment when the `delay_duration` has elapsed.  Current values include `CANCEL` and `COMPLETE`.
	DelayAction string `json:"delay_action,omitempty"`
	// The read-only timestamp of when the `delay_action` is automatically applied, in RFC 3339 format.  Note that this field is calculated by summing the payment's `delay_duration` and `created_at` fields. The `created_at` field is generated by Square and might not exactly match the time on your local machine.
	DelayedUntil string `json:"delayed_until,omitempty"`
	// The source type for this payment.  Current values include `CARD`, `BANK_ACCOUNT`, `WALLET`, `BUY_NOW_PAY_LATER`, `SQUARE_ACCOUNT`, `CASH` and `EXTERNAL`. For information about these payment source types, see [Take Payments](https://developer.squareup.com/docs/payments-api/take-payments).
	SourceType string `json:"source_type,omitempty"`
	CardDetails *CardPaymentDetails `json:"card_details,omitempty"`
	CashDetails *CashPaymentDetails `json:"cash_details,omitempty"`
	BankAccountDetails *BankAccountPaymentDetails `json:"bank_account_details,omitempty"`
	ExternalDetails *ExternalPaymentDetails `json:"external_details,omitempty"`
	WalletDetails *DigitalWalletDetails `json:"wallet_details,omitempty"`
	BuyNowPayLaterDetails *BuyNowPayLaterDetails `json:"buy_now_pay_later_details,omitempty"`
	SquareAccountDetails *SquareAccountDetails `json:"square_account_details,omitempty"`
	// The ID of the location associated with the payment.
	LocationId string `json:"location_id,omitempty"`
	// The ID of the order associated with the payment.
	OrderId string `json:"order_id,omitempty"`
	// An optional ID that associates the payment with an entity in another system.
	ReferenceId string `json:"reference_id,omitempty"`
	// The ID of the customer associated with the payment. If the ID is  not provided in the `CreatePayment` request that was used to create the `Payment`,  Square may use information in the request  (such as the billing and shipping address, email address, and payment source)  to identify a matching customer profile in the Customer Directory.  If found, the profile ID is used. If a profile is not found, the  API attempts to create an  [instant profile](https://developer.squareup.com/docs/customers-api/what-it-does#instant-profiles).  If the API cannot create an  instant profile (either because the seller has disabled it or the  seller's region prevents creating it), this field remains unset. Note that  this process is asynchronous and it may take some time before a  customer ID is added to the payment.
	CustomerId string `json:"customer_id,omitempty"`
	// __Deprecated__: Use `Payment.team_member_id` instead.  An optional ID of the employee associated with taking the payment.
	EmployeeId string `json:"employee_id,omitempty"`
	// An optional ID of the [TeamMember](entity:TeamMember) associated with taking the payment.
	TeamMemberId string `json:"team_member_id,omitempty"`
	// A list of `refund_id`s identifying refunds for the payment.
	RefundIds []string `json:"refund_ids,omitempty"`
	RiskEvaluation *RiskEvaluation `json:"risk_evaluation,omitempty"`
	// An optional ID for a Terminal checkout that is associated with the payment.
	TerminalCheckoutId string `json:"terminal_checkout_id,omitempty"`
	// The buyer's email address.
	BuyerEmailAddress string `json:"buyer_email_address,omitempty"`
	BillingAddress *Address `json:"billing_address,omitempty"`
	ShippingAddress *Address `json:"shipping_address,omitempty"`
	// An optional note to include when creating a payment.
	Note string `json:"note,omitempty"`
	// Additional payment information that gets added to the customer's card statement as part of the statement description.  Note that the `statement_description_identifier` might get truncated on the statement description to fit the required information including the Square identifier (SQ *) and the name of the seller taking the payment.
	StatementDescriptionIdentifier string `json:"statement_description_identifier,omitempty"`
	// Actions that can be performed on this payment: - `EDIT_AMOUNT_UP` - The payment amount can be edited up. - `EDIT_AMOUNT_DOWN` - The payment amount can be edited down. - `EDIT_TIP_AMOUNT_UP` - The tip amount can be edited up. - `EDIT_TIP_AMOUNT_DOWN` - The tip amount can be edited down. - `EDIT_DELAY_ACTION` - The delay_action can be edited.
	Capabilities []string `json:"capabilities,omitempty"`
	// The payment's receipt number. The field is missing if a payment is canceled.
	ReceiptNumber string `json:"receipt_number,omitempty"`
	// The URL for the payment's receipt. The field is only populated for COMPLETED payments.
	ReceiptUrl string `json:"receipt_url,omitempty"`
	DeviceDetails *DeviceDetails `json:"device_details,omitempty"`
	ApplicationDetails *ApplicationDetails `json:"application_details,omitempty"`
	// Whether or not this payment was taken offline.
	IsOfflinePayment bool `json:"is_offline_payment,omitempty"`
	OfflinePaymentDetails *OfflinePaymentDetails `json:"offline_payment_details,omitempty"`
	// Used for optimistic concurrency. This opaque token identifies a specific version of the `Payment` object.
	VersionToken string `json:"version_token,omitempty"`
}