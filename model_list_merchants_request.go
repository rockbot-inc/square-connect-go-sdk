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

// Request object for the [ListMerchant]($e/Merchants/ListMerchants) endpoint.
type ListMerchantsRequest struct {
	// The cursor generated by the previous response.
	Cursor int32 `json:"cursor,omitempty"`
}
