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

// Defines supported query expressions to search for vendors by.
type SearchVendorsRequestFilter struct {
	// The names of the [Vendor](entity:Vendor) objects to retrieve.
	Name []string `json:"name,omitempty"`
	// The statuses of the [Vendor](entity:Vendor) objects to retrieve. See [VendorStatus](#type-vendorstatus) for possible values
	Status []VendorStatus `json:"status,omitempty"`
}