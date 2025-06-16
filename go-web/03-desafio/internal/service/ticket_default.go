package service

import (
	"app/internal"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp internal.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetAllTickets returns all Tickets
func (s *ServiceTicketDefault) GetAllTickets() (tickets map[int]internal.TicketAttributes, err error) {
	return s.rp.Get(nil)
}

// GetTicketsByDestinationCountry returns Tickets by destination country
func (s *ServiceTicketDefault) GetTicketsByDestinationCountry(country string) (tickets map[int]internal.TicketAttributes, err error) {
	return s.rp.GetTicketsByDestinationCountry(nil, country)
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	tickets, err := s.rp.Get(nil)
	total = len(tickets)
	return
}

// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(country string) (total int, err error) {
	tickets, err := s.rp.GetTicketsByDestinationCountry(nil, country)
	total = len(tickets)
	return
}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(country string) (avg float64, err error) {
	total, err := s.GetTotalAmountTickets()
	if err != nil {
		return
	}
	countryAmount, err := s.GetTicketsAmountByDestinationCountry(country)
	if err != nil {
		return
	}
	avg = float64(countryAmount) / float64(total)
	return
}
