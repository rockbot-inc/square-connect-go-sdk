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
import (
	"os"
)

type ImagesImageIdBody struct {
	Request *UpdateCatalogImageRequest `json:"request,omitempty"`
	ImageFile **os.File `json:"image_file,omitempty"`
}
