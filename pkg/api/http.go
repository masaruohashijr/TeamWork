package api

import (
	"golang-interview-project-masaru-ohashi/pkg/member"
	m "golang-interview-project-masaru-ohashi/pkg/member"
	"golang-interview-project-masaru-ohashi/pkg/serializer/json"
	"golang-interview-project-masaru-ohashi/pkg/serializer/msgpack"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/go-chi/chi"
)

type MemberHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	memberService member.MemberService
}

func NewHandler(memberService member.MemberService) MemberHandler {
	return &handler{memberService: memberService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) member.MemberSerializer {
	if contentType == "application/x-msgpack" {
		return &msgpack.Member{}
	}
	return &json.Member{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	member, err := h.memberService.Find(name)
	if err != nil {
		if errors.Cause(err) == m.ErrMemberNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(member.Name))
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	member, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = h.memberService.Store(member)
	if err != nil {
		if errors.Cause(err) == m.ErrMemberNotFound {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).Encode(member)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}
