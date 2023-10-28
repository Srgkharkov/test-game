package apiserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Srgkharkov/test-game/internal/game"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/golang/gddo/httputil/header"
)

func (h *APIHandler) AddConfig(w http.ResponseWriter, r *http.Request) {
	(*h.metrics.RequestsTotal).Inc()
	// If the Content-Type header is present, check that it has the value
	// multipart/form-data. Note that we are using the gddo/httputil/header
	// package to parse and extract the value here, so the check works
	// even if the client includes additional charset or boundary
	// information in the header.
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "multipart/form-data" {
			msg := "Content-Type header is not multipart/form-data"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			(*h.metrics.ErrorResponseTotal).Inc()
			return
		}
	}

	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
	// response body. A request body larger than that will now result in
	// Decode() returning a "http: request body too large" error.
	r.ParseMultipartForm(128 << 10)

	configType := r.FormValue("configType")
	if configType == "" {
		http.Error(w, "Not found type of config", http.StatusBadRequest)
		(*h.metrics.ErrorResponseTotal).Inc()
		return
	}

	configName := r.FormValue("configName")
	if configName == "" {
		http.Error(w, "Not found name config", http.StatusBadRequest)
		(*h.metrics.ErrorResponseTotal).Inc()
		return
	}

	//conftext := r.FormValue("conftext")
	//if conftext == "" {
	//	http.Error(w, "Not found config text", http.StatusBadRequest)
	//	return
	//}

	conffile, _, err := r.FormFile("config")
	if err != nil {
		http.Error(w, "Bad or not found file config", http.StatusBadRequest)
		(*h.metrics.ErrorResponseTotal).Inc()
		return
	}
	defer conffile.Close()

	switch configType {
	case "reels":
		var Config_reels game.Config_reels
		Config_reels.Name = configName

		var field [3][5]string

		decoder := json.NewDecoder(conffile)
		err = decoder.Decode(&field)
		if err != nil {
			http.Error(w, "Bad file config, I can`t decode JSON", http.StatusBadRequest)
			(*h.metrics.ErrorResponseTotal).Inc()
			return
		}

		// Convert [3][5]string to [3][5]rune
		for i := 0; i < 3; i++ {
			for j := 0; j < 5; j++ {
				for _, c := range field[i][j] {
					Config_reels.Reels[i][j] = c
					break
				}
			}
		}

		if err := h.game.Configs_reels.AddConfig(&Config_reels); err != nil {
			http.Error(w, "Internal error, Can`t add Config", http.StatusInternalServerError)
			(*h.metrics.ErrorResponseTotal).Inc()
			return
		}

	case "lines":
		var Config_lines game.Config_lines
		Config_lines.Name = configName
		decoder := json.NewDecoder(conffile)
		err = decoder.Decode(&Config_lines.Lines)
		if err != nil {
			http.Error(w, "Bad file config, I can`t decode JSON", http.StatusBadRequest)
			(*h.metrics.ErrorResponseTotal).Inc()
			return
		}
		if err := h.game.Configs_lines.AddConfig(&Config_lines); err != nil {
			http.Error(w, "Internal error, Can`t add Config", http.StatusInternalServerError)
			(*h.metrics.ErrorResponseTotal).Inc()
			return
		}

	case "payouts":
		type tPayouts struct {
			Symbol string `json:"symbol"`
			Payout []int  `json:"payout"`
		}

		var JSONPayouts []*tPayouts

		//type tConfig_payouts struct {
		//	Name     string
		//	Payouts  []*tPayouts
		//	mPayouts map[rune]*tPayouts
		//}

		var Config_payouts game.Config_payouts
		Config_payouts.Name = configName

		decoder := json.NewDecoder(conffile)
		err = decoder.Decode(&JSONPayouts)
		if err != nil {
			http.Error(w, "Bad file config, I can`t decode JSON", http.StatusBadRequest)
			(*h.metrics.ErrorResponseTotal).Inc()
			return
		}

		Payouts := make([]game.Payouts, len(JSONPayouts))

		for i, v := range JSONPayouts {
			//Payouts[i].Payout = make([]int, len(v.Payout))
			//copy(Config_payouts.Payouts[i].Payout, v.Payout)
			Payouts[i].Payout = v.Payout
			for _, c := range v.Symbol {
				Payouts[i].Symbol = c
				break
			}
		}

		Config_payouts.Payouts = Payouts

		if err := h.game.Configs_payouts.AddConfig(&Config_payouts); err != nil {
			http.Error(w, "Internal error, Can`t add Config", http.StatusInternalServerError)
			(*h.metrics.ErrorResponseTotal).Inc()
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, "OK")
	w.WriteHeader(http.StatusOK)
}

// GetResult godoc
// @Summary      Get result
// @Description  get result by Config reels name, Config lines name and Config payouts name
// @Tags         result
// @Accept       json
// @Produce      json
// @Param        conf_reels_name   conf_lines_name      conf_payouts_name
// @Success      200  {string}  OK
// @Failure      400  {object}  httputil.HTTPError
// @Failure      413  {object}  httputil.HTTPError
// @Failure      415  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /getresult [get]
func (h *APIHandler) GetResult(w http.ResponseWriter, r *http.Request) {
	(*h.metrics.RequestsTotal).Inc()
	// If the Content-Type header is present, check that it has the value
	// application/json. Note that we are using the gddo/httputil/header
	// package to parse and extract the value here, so the check works
	// even if the client includes additional charset or boundary
	// information in the header.
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			(*h.metrics.ErrorResponseTotal).Inc()
			return
		}
	}

	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
	// response body. A request body larger than that will now result in
	// Decode() returning a "http: request body too large" error.
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	// Setup the decoder and call the DisallowUnknownFields() method on it.
	// This will cause Decode() to return a "json: unknown field ..." error
	// if it encounters any extra unexpected fields in the JSON. Strictly
	// speaking, it returns an error for "keys which do not match any
	// non-ignored, exported fields in the destination".
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var ReqResult game.ReqResult
	err := dec.Decode(&ReqResult)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
		// is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			http.Error(w, msg, http.StatusBadRequest)

		// Catch any type errors, like trying to assign a string in the
		// JSON request body to a int field in our Person struct. We can
		// interpolate the relevant field name and position into the error
		// message to make it easier for the client to fix.
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		// Catch the error caused by extra unexpected fields in the request
		// body. We extract the field name from the error message and
		// interpolate it in our custom error message. There is an open
		// issue at https://github.com/golang/go/issues/29035 regarding
		// turning this into a sentinel error.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			http.Error(w, msg, http.StatusBadRequest)

		// An io.EOF error is returned by Decode() if the request body is
		// empty.
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			http.Error(w, msg, http.StatusBadRequest)

		// Catch the error caused by the request body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		// error at https://github.com/golang/go/issues/30715.
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			http.Error(w, msg, http.StatusRequestEntityTooLarge)

		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		(*h.metrics.ErrorResponseTotal).Inc()
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, err := h.game.GetResult(&ReqResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		(*h.metrics.ErrorResponseTotal).Inc()
		return
	}
	bresult, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		(*h.metrics.ErrorResponseTotal).Inc()
		return
	}
	w.Write(bresult)
	w.WriteHeader(http.StatusOK)
}
