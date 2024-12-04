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

// Represents an output from a call to [BulkRetrieveVendors]($e/Vendors/BulkRetrieveVendors).
type BulkRetrieveVendorsResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The set of [RetrieveVendorResponse](entity:RetrieveVendorResponse) objects encapsulating successfully retrieved [Vendor](entity:Vendor) objects or error responses for failed attempts. The set is represented by  a collection of `Vendor`-ID/`Vendor`-object or `Vendor`-ID/error-object pairs.
	Responses map[string]RetrieveVendorResponse `json:"responses,omitempty"`
}
