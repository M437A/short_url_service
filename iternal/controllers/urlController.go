package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"net/http"
	"short_url/iternal/entity"
	"short_url/iternal/exceptions"
	"short_url/iternal/services"
	"short_url/iternal/validators"
)

var Router *chi.Mux

func CreateShortUrl(responseWriter http.ResponseWriter, request *http.Request) {
	var urlLink entity.Link
	getDataFromRequest(responseWriter, request, &urlLink)

	validateErr := validators.ValidateCreateShortUrl(urlLink)
	if exceptions.JSONResponseException(validateErr, responseWriter) {
		return
	}

	errorGenerate := services.GetNewShortUrl(&urlLink)
	if exceptions.JSONResponseException(errorGenerate, responseWriter) {
		return
	}

	Router.Get(urlLink.ShortURL, OpenShorUrl)
	JSONResponse(responseWriter, viper.GetString("url.prefix")+request.Host+urlLink.ShortURL)
}

func OpenShorUrl(responseWriter http.ResponseWriter, request *http.Request) {
	var urlLink entity.Link
	urlLink.ShortURL = request.URL.Path

	newUrlLink, err := services.GetLongUrl(urlLink)

	if exceptions.JSONResponseException(err, responseWriter) {
		return
	}

	http.Redirect(responseWriter, request, newUrlLink.LongURL, http.StatusSeeOther)
}

func JSONResponse(responseWriter http.ResponseWriter, data interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	err := json.NewEncoder(responseWriter).Encode(data)
	exceptions.JSONResponseException(err, responseWriter)
}

func getDataFromRequest[T any](responseWriter http.ResponseWriter, request *http.Request, object *T) {
	errorMapping := json.NewDecoder(request.Body).Decode(&object)
	exceptions.JSONResponseException(errorMapping, responseWriter)
}
