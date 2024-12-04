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
// BookingStatus : Supported booking statuses.
type BookingStatus string

// List of BookingStatus
const (
	PENDING_BookingStatus BookingStatus = "PENDING"
	CANCELLED_BY_CUSTOMER_BookingStatus BookingStatus = "CANCELLED_BY_CUSTOMER"
	CANCELLED_BY_SELLER_BookingStatus BookingStatus = "CANCELLED_BY_SELLER"
	DECLINED_BookingStatus BookingStatus = "DECLINED"
	ACCEPTED_BookingStatus BookingStatus = "ACCEPTED"
	NO_SHOW_BookingStatus BookingStatus = "NO_SHOW"
)
