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

// Represents a group of customer profiles.   Customer groups can be created, be modified, and have their membership defined using  the Customers API or within the Customer Directory in the Square Seller Dashboard or Point of Sale.
type CustomerGroup struct {
	// A unique Square-generated ID for the customer group.
	Id string `json:"id,omitempty"`
	// The name of the customer group.
	Name string `json:"name"`
	// The timestamp when the customer group was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp when the customer group was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
}
