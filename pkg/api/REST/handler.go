package REST

import (
	"encoding/json"
	"fmt"
	jss "golang-interview-project-masaru-ohashi/pkg/serializer/json"
	"golang-interview-project-masaru-ohashi/pkg/serializer/msgpack"
	"golang-interview-project-masaru-ohashi/pkg/serializer/xml"
	"golang-interview-project-masaru-ohashi/pkg/team"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

type MemberHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type handler struct {
	memberService team.MemberService
}

func NewHandler(memberService team.MemberService) MemberHandler {
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

func (h *handler) serializer(contentType string) team.MemberSerializer {
	if contentType == "application/x-msgpack" {
		return &msgpack.Member{}
	} else if contentType == "application/xml" {
		return &xml.Member{}
	}
	return &jss.Member{}
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	members, err := h.memberService.GetAll()
	if err != nil {
		if errors.Cause(err) == team.ErrMemberNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	sMembers, err := json.Marshal(members)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(sMembers))
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	fmt.Println(name)
	member, err := h.memberService.Get(name)
	if err != nil {
		if errors.Cause(err) == team.ErrMemberNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	sMember, err := json.Marshal(member)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(sMember))
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	requestMember, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = h.memberService.Create(requestMember.Member)
	if err != nil {
		if errors.Cause(err) == team.ErrMemberNotFound {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		} else if errors.Cause(err) == team.ErrMemberInvalid {
			println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).Encode(requestMember.Member)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}

func (h *handler) Put(w http.ResponseWriter, r *http.Request) {
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
}
