package Handler

import (
	"net/http"
	"simple-rest-go-echo/Config"
	"simple-rest-go-echo/Models"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

// Get All Courses Data
func GetCourse(c echo.Context) error {
	// Get the database instance from Config package
	db := Config.GetDB()

	var course []*Models.Course

	if err := db.Find(&course); err.Error != nil {
		data := map[string]interface{}{
			"message": err.Error.Error(),
		}
		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"message": "Fetch Successfully",
		"data":    course,
	}

	return c.JSON(http.StatusOK, response)
}

// Create Course Data
func CreateCourse(c echo.Context) error {
	db := Config.GetDB()

	course := new(Models.Course)

	// Bind data
	if err := c.Bind(course); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	// Validate the course struct
	validate := validator.New()
	if err := validate.Struct(course); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	if err := db.Create(&course).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Create Successfully",
		"data":    course,
	}
	return c.JSON(http.StatusOK, response)
}

// Update Courses Data
func UpdateCourse(c echo.Context) error {
	db := Config.GetDB()

	course := new(Models.Course)

	// Bind data
	if err := c.Bind(course); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	// Validate the course struct
	validate := validator.New()
	if err := validate.Struct(course); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	// Retrieve ID from the URL parameter
	id := c.Param("id")
	existing_course := new(Models.Course)

	if err := db.First(&existing_course, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existing_course.Name = course.Name
	existing_course.Description = course.Description
	existing_course.Price = course.Price

	if err := db.Save(&existing_course).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Update Successfully",
		"data":    existing_course,
	}
	return c.JSON(http.StatusOK, response)
}

// Delete Course Data
func DeleteCourse(c echo.Context) error {
	db := Config.GetDB()

	course := new(Models.Course)

	id := c.Param("id")

	err := db.Delete(&course, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Delete Successfully!",
	}
	return c.JSON(http.StatusOK, response)
}
