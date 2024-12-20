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

// Represents an output from a call to [UpdateVendor]($e/Vendors/UpdateVendor).
type UpdateVendorResponse struct {
	// Errors occurred when the request fails.
	Errors []ModelError `json:"errors,omitempty"`
	Vendor *Vendor `json:"vendor,omitempty"`
}
