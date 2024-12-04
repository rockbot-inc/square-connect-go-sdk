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

// Represents an output from a call to [BulkUpdateVendors]($e/Vendors/BulkUpdateVendors).
type BulkUpdateVendorsResponse struct {
	// Errors encountered when the request fails.
	Errors []ModelError `json:"errors,omitempty"`
	// A set of [UpdateVendorResponse](entity:UpdateVendorResponse) objects encapsulating successfully created [Vendor](entity:Vendor) objects or error responses for failed attempts. The set is represented by a collection of `Vendor`-ID/`UpdateVendorResponse`-object or  `Vendor`-ID/error-object pairs.
	Responses map[string]UpdateVendorResponse `json:"responses,omitempty"`
}
