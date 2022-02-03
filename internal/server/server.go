package server

import (
	"SimpleInferencer/pkg/api"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/logger"
	"github.com/gorilla/mux"
)

type Server struct {
	ctx context.Context
}

func NewServer(parentContext context.Context) (*Server, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parentContext)
	return &Server{ctx: ctx}, cancel
}

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

func (s *Server) Routes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range s.routes() {
		router.Methods(route.Method).Path(route.Path).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}

func (s *Server) routes() []Route {
	return []Route{
		{"Health", http.MethodGet, "/healthcheck", s.HealthcheckHandler},
		{"Inference", http.MethodPost, "/inference", s.InferencingHandler},
	}
}

func (s *Server) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	logger.Infof("Received HealthCheck request")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) InferencingHandler(w http.ResponseWriter, r *http.Request) {
	var req api.InferenceRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Deserialization error %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Errorf("Deserialization error %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Infof("Received Inferencing request %+v", req)

	resp := api.InferenceResponse{
		Id:     req.Id,
		Type:   req.Type,
		Output: fmt.Sprintf("echoing %s", req.Input),
	}

	respBody, err := json.Marshal(&resp)
	if err != nil {
		logger.Errorf("Response serialization error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("input", req.Input)
	w.Write(respBody)
}
