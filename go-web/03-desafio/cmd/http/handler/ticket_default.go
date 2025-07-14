package handler

import (
	"app/internal"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type TicketHandlerDefault struct {
	service internal.ServiceTicket
}

func NewTicketsHandlerDefault(sv internal.ServiceTicket) *TicketHandlerDefault {
	return &TicketHandlerDefault{
		service: sv,
	}
}

func (hd *TicketHandlerDefault) GetAllTickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tickets, err := hd.service.GetAllTickets()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Error geting all tickets")
			return
		}

		body := struct {
			Data map[int]internal.TicketAttributes `json:"data"`
		}{
			Data: tickets,
		}

		response.JSON(w, http.StatusOK, body)
	}
}

func (hd *TicketHandlerDefault) GetTicketsByCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := chi.URLParam(r, "dest")
		// CR1 - Aqui poderiamos adionar uma validacao para verificar se o parâmetro 'dest' está vazio ou inválido?
		tickets, err := hd.service.GetTicketsByDestinationCountry(country)
		if err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		body := struct {
			Data map[int]internal.TicketAttributes `json:"data"`
		}{
			Data: tickets,
		}

		response.JSON(w, http.StatusOK, body)
	}
}

func (hd *TicketHandlerDefault) GetAverageTicketsByCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := chi.URLParam(r, "dest")
		// CR2 - Aqui poderiamos adionar uma validacao para verificar se o parâmetro 'dest' está vazio ou inválido?
		avg, err := hd.service.GetPercentageTicketsByDestinationCountry(country)
		if err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		body := struct {
			Data float64 `json:"data"`
		}{
			Data: avg,
		}

		response.JSON(w, http.StatusOK, body)
	}
}
