package middlewares


import (
	"net/http"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Middleware ...
type Middleware func(http.HandlerFunc) http.HandlerFunc

// MultipleMiddleware ...
func MultipleMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {

   if len(m) < 1 {
      return h
   }

   wrapped := h

   // loop in reverse to preserve middleware order
   for i := len(m) - 1; i >= 0; i-- {
      wrapped = m[i](wrapped)
   }

   return wrapped
}

// SecreteKey ...
var SecreteKey string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	SecreteKey, _ = os.LookupEnv("SECRETE_KEY")
}