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
// EmployeeStatus : The status of the Employee being retrieved.  DEPRECATED at version 2020-08-26. Replaced by [TeamMemberStatus](entity:TeamMemberStatus).
type EmployeeStatus string

// List of EmployeeStatus
const (
	ACTIVE_EmployeeStatus EmployeeStatus = "ACTIVE"
	INACTIVE_EmployeeStatus EmployeeStatus = "INACTIVE"
)
