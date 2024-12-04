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

// A request for a set of `EmployeeWage` objects.
type ListEmployeeWagesRequest struct {
	// Filter the returned wages to only those that are associated with the specified employee.
	EmployeeId string `json:"employee_id,omitempty"`
	// The maximum number of `EmployeeWage` results to return per page. The number can range between 1 and 200. The default is 200.
	Limit int32 `json:"limit,omitempty"`
	// A pointer to the next page of `EmployeeWage` results to fetch.
	Cursor string `json:"cursor,omitempty"`
}