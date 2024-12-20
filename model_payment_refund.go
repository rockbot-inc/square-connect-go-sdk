/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package square

// Represents a refund of a payment made using Square. Contains information about the original payment and the amount of money refunded.
type PaymentRefund struct {
	// The unique ID for this refund, generated by Square.
	Id string `json:"id"`
	// The refund's status: - `PENDING` - Awaiting approval. - `COMPLETED` - Successfully completed. - `REJECTED` - The refund was rejected. - `FAILED` - An error occurred.
	Status string `json:"status,omitempty"`
	// The location ID associated with the payment this refund is attached to.
	LocationId string `json:"location_id,omitempty"`
	// Flag indicating whether or not the refund is linked to an existing payment in Square.
	Unlinked bool `json:"unlinked,omitempty"`
	// The destination type for this refund.  Current values include `CARD`, `BANK_ACCOUNT`, `WALLET`, `BUY_NOW_PAY_LATER`, `CASH`, `EXTERNAL`, and `SQUARE_ACCOUNT`.
	DestinationType string `json:"destination_type,omitempty"`
	DestinationDetails *DestinationDetails `json:"destination_details,omitempty"`
	AmountMoney *Money `json:"amount_money"`
	AppFeeMoney *Money `json:"app_fee_money,omitempty"`
	// Processing fees and fee adjustments assessed by Square for this refund.
	ProcessingFee []ProcessingFee `json:"processing_fee,omitempty"`
	// The ID of the payment associated with this refund.
	PaymentId string `json:"payment_id,omitempty"`
	// The ID of the order associated with the refund.
	OrderId string `json:"order_id,omitempty"`
	// The reason for the refund.
	Reason string `json:"reason,omitempty"`
	// The timestamp of when the refund was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp of when the refund was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
	// An optional ID of the team member associated with taking the payment.
	TeamMemberId string `json:"team_member_id,omitempty"`
	// An optional ID for a Terminal refund.
	TerminalRefundId string `json:"terminal_refund_id,omitempty"`
}
