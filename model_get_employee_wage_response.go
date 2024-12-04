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

// A response to a request to get an `EmployeeWage`. The response contains the requested `EmployeeWage` objects and might contain a set of `Error` objects if the request resulted in errors.
type GetEmployeeWageResponse struct {
	EmployeeWage *EmployeeWage `json:"employee_wage,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}