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

// Defines the fields that are included in requests to the `BatchRetrieveOrders` endpoint.
type BatchRetrieveOrdersRequest struct {
	// The ID of the location for these orders. This field is optional: omit it to retrieve orders within the scope of the current authorization's merchant ID.
	LocationId string `json:"location_id,omitempty"`
	// The IDs of the orders to retrieve. A maximum of 100 orders can be retrieved per request.
	OrderIds []string `json:"order_ids"`
}
