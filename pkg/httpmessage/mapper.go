package httpmessage

import (
	"net/http"

	"github.com/SoroushBeigi/knowledge-game/pkg/errmessage"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
)

func CodeAndMessage(err error) (message string, code int) {
	switch err := err.(type) {

	case richerror.RichError:
		re := err
		msg := re.Message()
		status := errorCodeToHTTPStatus(re.Code())

		//Security practice to prevent expoing error details
		if status >= 500 {
			msg = errmessage.SomethingWentWrong
		}

		return msg, status

	default:
		return err.Error(), http.StatusInternalServerError
	}
}

func errorCodeToHTTPStatus(code richerror.ErrorCode) int {
	switch code {
	case richerror.InvalidCode:
		return http.StatusUnprocessableEntity
	case richerror.ForbiddenCode:
		return http.StatusForbidden
	case richerror.UnexpectedCode:
		return http.StatusInternalServerError
	case richerror.NotFoundCode:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}

}
