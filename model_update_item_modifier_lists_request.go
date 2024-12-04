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

type UpdateItemModifierListsRequest struct {
	// The IDs of the catalog items associated with the CatalogModifierList objects being updated.
	ItemIds []string `json:"item_ids"`
	// The IDs of the CatalogModifierList objects to enable for the CatalogItem. At least one of `modifier_lists_to_enable` or `modifier_lists_to_disable` must be specified.
	ModifierListsToEnable []string `json:"modifier_lists_to_enable,omitempty"`
	// The IDs of the CatalogModifierList objects to disable for the CatalogItem. At least one of `modifier_lists_to_enable` or `modifier_lists_to_disable` must be specified.
	ModifierListsToDisable []string `json:"modifier_lists_to_disable,omitempty"`
}
