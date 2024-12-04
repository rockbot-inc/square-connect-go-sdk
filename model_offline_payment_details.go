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

// Details specific to offline payments.
type OfflinePaymentDetails struct {
	// The client-side timestamp of when the offline payment was created, in RFC 3339 format.
	ClientCreatedAt string `json:"client_created_at,omitempty"`
}
