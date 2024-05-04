import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Action func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error

type Controller struct{}