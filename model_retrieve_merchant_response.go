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

// The response object returned by the [RetrieveMerchant]($e/Merchants/RetrieveMerchant) endpoint.
type RetrieveMerchantResponse struct {
	// Information on errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Merchant *Merchant `json:"merchant,omitempty"`
}
