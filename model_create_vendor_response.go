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

// Represents an output from a call to [CreateVendor]($e/Vendors/CreateVendor).
type CreateVendorResponse struct {
	// Errors encountered when the request fails.
	Errors []ModelError `json:"errors,omitempty"`
	Vendor *Vendor `json:"vendor,omitempty"`
}