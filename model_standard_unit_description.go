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

// Contains the name and abbreviation for standard measurement unit.
type StandardUnitDescription struct {
	Unit *MeasurementUnit `json:"unit,omitempty"`
	// UI display name of the measurement unit. For example, 'Pound'.
	Name string `json:"name,omitempty"`
	// UI display abbreviation for the measurement unit. For example, 'lb'.
	Abbreviation string `json:"abbreviation,omitempty"`
}
