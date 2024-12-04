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

// A tax to block from applying to a line item. The tax must be identified by either `tax_uid` or `tax_catalog_object_id`, but not both.
type OrderLineItemPricingBlocklistsBlockedTax struct {
	// A unique ID of the `BlockedTax` within the order.
	Uid string `json:"uid,omitempty"`
	// The `uid` of the tax that should be blocked. Use this field to block ad hoc taxes. For catalog, taxes use the `tax_catalog_object_id` field.
	TaxUid string `json:"tax_uid,omitempty"`
	// The `catalog_object_id` of the tax that should be blocked. Use this field to block catalog taxes. For ad hoc taxes, use the `tax_uid` field.
	TaxCatalogObjectId string `json:"tax_catalog_object_id,omitempty"`
}