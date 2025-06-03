package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Ticket struct {
	ID          int
	Name        string
	Email       string
	Destination string
	Time        string
	Price       float64
}

const (
	Dawn      = "Dawn"      // 00:00 to 06:59
	Morning   = "Morning"   // 07:00 to 12:59
	Afternoon = "Afternoon" // 13:00 to 19:59
	Evening   = "Evening"   // 20:00 to 23:59
)

// Exercício 1:
// Uma função que calcula quantas pessoas viajam para um determinado país:
func GetTicketsByDestination(destination string) (int, error) {
	if destination == "" {
		return 0, fmt.Errorf("destination cannot be empty")
	}

	// ler o arquivo
	tickets := getTicketsFromCsv()
	destinationTicketsCount := 0

	// filtrar os tickets pelo destino
	for _, ticket := range tickets {
		if ticket.Destination == destination {
			destinationTicketsCount++
		}
	}

	if destinationTicketsCount == 0 {
		return 0, fmt.Errorf("no tickets found for destination: %s", destination)
	}

	// retornar a quantidade
	return destinationTicketsCount, nil
}

// Exercício 2:
// Uma ou mais funções que calculam quantas pessoas viajam no início da manhã (0 → 6), manhã (7 → 12), tarde (13 → 19) e noite (20 → 23):
func GetTicketsByPeriod(period string) (int, error) {
	if period == "" {
		return 0, fmt.Errorf("period cannot be empty")
	}

	start, end := "", ""

	switch period {
	case Dawn:
		start, end = "00:00", "06:59"
	case Morning:
		start, end = "07:00", "12:59"
	case Afternoon:
		start, end = "13:00", "19:59"
	case Evening:
		start, end = "20:00", "23:59"
	default:
		return 0, fmt.Errorf("invalid period: %s", period)
	}

	startTime, err := time.Parse("15:04", start)
	if err != nil {
		return 0, fmt.Errorf("failed to parse start time: %v", err)
	}

	endTime, err := time.Parse("15:04", end)
	if err != nil {
		return 0, fmt.Errorf("failed to parse end time: %v", err)
	}

	tickets := getTicketsFromCsv()
	count := 0

	for _, ticket := range tickets {
		ticketTime, err := time.Parse("15:04", ticket.Time)
		if err != nil {
			return 0, fmt.Errorf("failed to parse ticket time for ticket ID %d: %v", ticket.ID, err)
		}
		if ticketTime.After(startTime) && ticketTime.Before(endTime) || ticketTime.Equal(startTime) || ticketTime.Equal(endTime) {
			count++
		}
	}

	return count, nil
}

// Exercício 3:
// Calcule a porcentagem de pessoas que viajam para um determinado país em um determinado dia, em relação ao restante:
// Dica: A porcentagem de x é calculada da seguinte forma: porcentagem = (x / total)
func DestinationPercentage(destination string) (float64, error) {
	if destination == "" {
		return 0.0, fmt.Errorf("destination cannot be empty")
	}

	tickets := getTicketsFromCsv()
	totalTickets := len(tickets)
	if totalTickets == 0 {
		return 0.0, fmt.Errorf("no tickets found")
	}

	destinationCount, err := GetTicketsByDestination(destination)
	if err != nil {
		return 0.0, err
	}

	result := float64(destinationCount) / float64(totalTickets)
	return result, nil
}

// Função auxiliar para ler o arquivo CSV e retornar seu conteúdo como string
func readCsv() (result string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	content, err := os.ReadFile("go-bases/desafio-go-bases/tickets.csv")
	if err != nil {
		panic("Failed to read tickets.csv: " + err.Error())
	}
	result = string(content)
	//fmt.Println("File read successfully:\n", result)
	return
}

// Função auxiliar para retornar todos os tickets como uma slice de Ticket
func getTicketsFromCsv() []Ticket {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	reader := csv.NewReader(strings.NewReader(readCsv()))
	records, err := reader.ReadAll()
	if err != nil {
		panic("Failed to read CSV: " + err.Error())
	}
	var tickets []Ticket
	for i, record := range records {

		//fmt.Printf("Record %d: %v\n", i, record)

		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse price for record %d: %v", i, err))
		}

		tickets = append(tickets, Ticket{
			ID:          i,
			Name:        record[1],
			Email:       record[2],
			Destination: record[3],
			Time:        record[4],
			Price:       price,
		})
	}

	return tickets
}
