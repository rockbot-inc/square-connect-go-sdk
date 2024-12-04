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
// CustomAttributeDefinitionVisibility : The level of permission that a seller or other applications requires to view this custom attribute definition. The `Visibility` field controls who can read and write the custom attribute values and custom attribute definition.
type CustomAttributeDefinitionVisibility string

// List of CustomAttributeDefinitionVisibility
const (
	HIDDEN_CustomAttributeDefinitionVisibility CustomAttributeDefinitionVisibility = "VISIBILITY_HIDDEN"
	READ_ONLY_CustomAttributeDefinitionVisibility CustomAttributeDefinitionVisibility = "VISIBILITY_READ_ONLY"
	READ_WRITE_VALUES_CustomAttributeDefinitionVisibility CustomAttributeDefinitionVisibility = "VISIBILITY_READ_WRITE_VALUES"
)
