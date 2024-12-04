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
// InvoicePaymentReminderStatus : The status of a payment request reminder.
type InvoicePaymentReminderStatus string

// List of InvoicePaymentReminderStatus
const (
	PENDING_InvoicePaymentReminderStatus InvoicePaymentReminderStatus = "PENDING"
	NOT_APPLICABLE_InvoicePaymentReminderStatus InvoicePaymentReminderStatus = "NOT_APPLICABLE"
	SENT_InvoicePaymentReminderStatus InvoicePaymentReminderStatus = "SENT"
)
