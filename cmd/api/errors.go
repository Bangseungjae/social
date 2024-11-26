package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw(
		"internal error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw(
		"bad request error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusNotFound, "not found")
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw(
		"conflict error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusConflict, err.Error())
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw(
		"unauthorized basic error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	/*
		"WWW-Authenticate" 헤더:
				•	HTTP 프로토콜에서 인증이 필요한 리소스에 접근할 때 서버가 클라이언트에게 인증 방식을 알리기 위해 사용하는 헤더입니다.
				•	주로 401 Unauthorized 응답과 함께 사용됩니다.
		Basic realm="restricted", charset="UTF-8" 값:
			•	Basic: HTTP Basic Authentication 방식을 사용함을 나타냅니다. 이 방식은 사용자 이름과 비밀번호를 Base64로 인코딩하여 전송합니다.
			•	realm="restricted": 인증이 필요한 영역(Realm)의 이름을 지정합니다. 클라이언트는 이 값을 참고하여 사용자에게 인증을 요청할 때 표시할 수 있습니다.
			•	charset="UTF-8": 클라이언트와 서버 간의 문자 인코딩 방식을 지정합니다. 여기서는 UTF-8을 사용함을 명시하고 있습니다.

	*/
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	writeJSONError(w, http.StatusNotFound, "unauthorized")
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw(
		"unauthorized error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusNotFound, "unauthorized")
}
