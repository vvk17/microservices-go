package utilities

include(
	"github.com/gofiber/fiber/v2"
)

func GetBaseResponseObject() map[string]interface{} {
	response := make(map[string]interface{})
	response["status"] = "fail"
	response["message"] = "General error"
	return response
}

func PostBodyValidation (c *fiber.Ctx, postBody interface{}) interface{} {
	response := GetBaseResponseObject()

	if err := c.BodyParser(postBody); err != nil {
		response["error"] = err.Error()
		return response
	} else {
		if errors := validators.ValidateStruct(postBody); errors != nil {
			return errors
		} else {
			return nil
		}
	}
}