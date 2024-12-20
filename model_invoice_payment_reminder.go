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

// Describes a payment request reminder (automatic notification) that Square sends to the customer. You configure a reminder relative to the payment request `due_date`.
type InvoicePaymentReminder struct {
	// A Square-assigned ID that uniquely identifies the reminder within the `InvoicePaymentRequest`.
	Uid string `json:"uid,omitempty"`
	// The number of days before (a negative number) or after (a positive number) the payment request `due_date` when the reminder is sent. For example, -3 indicates that the reminder should be sent 3 days before the payment request `due_date`.
	RelativeScheduledDays int32 `json:"relative_scheduled_days,omitempty"`
	// The reminder message.
	Message string `json:"message,omitempty"`
	Status *InvoicePaymentReminderStatus `json:"status,omitempty"`
	// If sent, the timestamp when the reminder was sent, in RFC 3339 format.
	SentAt string `json:"sent_at,omitempty"`
}
