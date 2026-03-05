package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"u-short/internal/service"
	"u-short/internal/utils"

	"github.com/go-chi/chi/v5"
)

type UrlHandler struct {
	svc  *service.UrlService
	tmpl *template.Template
}

func NewUrlHandler(svc *service.UrlService) *UrlHandler {
	t := template.Must(template.ParseGlob("web/templates/*.html"))
	template.Must(t.ParseGlob("web/templates/partials/*.html"))

	return &UrlHandler{
		svc:  svc,
		tmpl: t,
	}
}

// Render Layout HTML
func (h *UrlHandler) Index(w http.ResponseWriter, r *http.Request) {
	links, clicks := h.svc.GetStats(r.Context())

	data := map[string]interface{}{
		"Links":  links,
		"Clicks": clicks,
	}

	err := h.tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ShortLink form
func (h *UrlHandler) ShortLink(w http.ResponseWriter, r *http.Request) {
	h.tmpl.ExecuteTemplate(w, "form-shorten", nil)
}

// QRCode form
func (h *UrlHandler) QRCode(w http.ResponseWriter, r *http.Request) {
	h.tmpl.ExecuteTemplate(w, "form-qr", nil)
}

// Create ShortLink
func (h *UrlHandler) Create(baseUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		originalUrl := r.FormValue("url")
		customAlias := r.FormValue("code")
		FormType := r.FormValue("type")
		cleanedUrl := strings.TrimSuffix(originalUrl, "/")

		if cleanedUrl == "" {
			http.Error(w, "Url is required", http.StatusBadRequest)
			return
		}

		result, err := h.svc.Shorten(r.Context(), cleanedUrl, customAlias)
		if err != nil {
			data := map[string]interface{}{
				"Message": err.Error(),
			}

			h.tmpl.ExecuteTemplate(w, "error", data)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		shortLink := baseUrl + result.ShortCode

		triggerData := map[string]string{
			"id":       fmt.Sprint(result.ID),
			"original": cleanedUrl,
			"short":    shortLink,
		}

		jsonTrigger, _ := json.Marshal(map[string]interface{}{
			"linkCreated": triggerData,
		})

		w.Header().Set("HX-Trigger", string(jsonTrigger))

		data := map[string]interface{}{
			"OriginalUrl": cleanedUrl,
			"ShortLink":   shortLink,
		}

		if FormType == "qr" {

			qrBase64, err := utils.GetQrCode(shortLink)
			if err != nil {
				http.Error(w, "Failed to generate QR code", http.StatusBadRequest)
				return
			}
			data["QRCode"] = qrBase64

			err = h.tmpl.ExecuteTemplate(w, "result-qr", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			err = h.tmpl.ExecuteTemplate(w, "result-shorten", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

// QR Scan
func (h *UrlHandler) Scan(w http.ResponseWriter, r *http.Request) {
	originalUrl := r.FormValue("url")

	qrBase64, err := utils.GetQrCode(originalUrl)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	h.tmpl.ExecuteTemplate(w, "scan-qr", qrBase64)
}

// Redirect link
func (h *UrlHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")

	if shortCode == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	originalUrl, err := h.svc.GetOriginalUrl(r.Context(), shortCode)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, originalUrl, http.StatusMovedPermanently)
}
