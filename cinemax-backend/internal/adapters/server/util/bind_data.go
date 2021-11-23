package util

import (
	"log"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func init() {
	// Se implementa para que los errores devuelvan el tag de la etiqueta
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

// invalidArgument define el formato en que se regresan los errores
type invalidArgument struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
	Tag   string      `json:"tag"`
	Param string      `json:"param"`
}

type ErrorArgs struct {
	Error       string        `json:"error"`
	InvalidArgs []invalidArgs `json:"invalidArgs"`
}

type invalidArgs struct {
	Param string `json:"param"`
	Tag   string `json:"tag"`
}

// BindData devuelve falso si los datos no coinciden con los esperados
func BindData(c *gin.Context, req interface{}) bool {
	// Enlaza el JSON a la estructura y verifica errores de validación
	if err := c.ShouldBind(req); err != nil {
		invalidArgs := GetErrorsOfBinding(err)
		if invalidArgs != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{
					"error":       "Invalid request parameters",
					"invalidArgs": invalidArgs})
			return false
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return false
	}
	return true
}

// GetErrorsOfBinding devuleve una lista con los errores encontrados al realizar
// la validación.
func GetErrorsOfBinding(err error) []invalidArgument {
	var invalidArgs []invalidArgument
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			invalidArgs = append(invalidArgs, invalidArgument{
				Field: err.Field(),
				Value: err.Value(),
				Tag:   err.Tag(),
				Param: err.Param(),
			})
		}
		return invalidArgs
	}
	return nil
}

func BindFile(c *gin.Context, file string, req **multipart.FileHeader) bool {
	var err error
	*req, err = c.FormFile(file)
	if err != nil {
		log.Printf("Unable parse multipart/form-data: %+v", err)
		if err.Error() == "http: request body too large" {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"error": "Max request body size exceded\n",
			})
			return false
		}
		e := domain.NewBadRequest("Unable to parse multipart/form-data")
		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return false
	}
	return true
}

func BindQuery(c *gin.Context, query string, req *string) bool {
	*req = c.Query(query)
	var errorArgs ErrorArgs
	if *req == "" {
		errorArgs.Error = "Invalid request parameters"
		errorArgs.InvalidArgs = append(errorArgs.InvalidArgs, invalidArgs{
			Param: query,
			Tag:   "required",
		})
		c.JSON(http.StatusBadRequest, errorArgs)
		return false
	}
	return true
}
