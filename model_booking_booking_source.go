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
// BookingBookingSource : Supported sources a booking was created from.
type BookingBookingSource string

// List of BookingBookingSource
const (
	FIRST_PARTY_MERCHANT_BookingBookingSource BookingBookingSource = "FIRST_PARTY_MERCHANT"
	FIRST_PARTY_BUYER_BookingBookingSource BookingBookingSource = "FIRST_PARTY_BUYER"
	THIRD_PARTY_BUYER_BookingBookingSource BookingBookingSource = "THIRD_PARTY_BUYER"
	API_BookingBookingSource BookingBookingSource = "API"
)
