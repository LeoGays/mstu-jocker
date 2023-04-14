package server

import (
	"jocer/internal/server/generated"
	"jocer/internal/server/mapper"
	"jocer/pkg/response"
	"jocer/pkg/serial"
	"net/http"
)

func (s Server) GetJocks(w http.ResponseWriter, r *http.Request) {
	jocks, err := s.useCase.GetJocks(r.Context())
	if err != nil {
		response.JSON(w, http.StatusBadRequest, generated.Error{
			Message: err.Error(),
		},
		)
	}

	response.JSON(w, http.StatusOK, mapper.CreateJockListResponse(jocks))
}

func (s Server) CreateJock(w http.ResponseWriter, r *http.Request) {
	newJock, err := serial.JSONDecode[generated.JockRequestBody](r.Body)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, generated.Error{Message: err.Error()})
	}
	jock, err := s.useCase.CreateJock(r.Context(), mapper.CreateJock(newJock))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, generated.Error{Message: err.Error()})
	}

	response.JSON(w, http.StatusCreated, mapper.CreateJockResponse(jock))
}
