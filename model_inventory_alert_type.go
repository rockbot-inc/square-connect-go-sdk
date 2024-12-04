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
// InventoryAlertType : Indicates whether Square should alert the merchant when the inventory quantity of a CatalogItemVariation is low.
type InventoryAlertType string

// List of InventoryAlertType
const (
	NONE_InventoryAlertType InventoryAlertType = "NONE"
	LOW_QUANTITY_InventoryAlertType InventoryAlertType = "LOW_QUANTITY"
)