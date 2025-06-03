package main

import (
	"bgw7-aulas/go-bases/desafio-go-bases/internal/tickets"
	"fmt"
)

func main() {
	// Exercício 1:
	ticketsCount, err := tickets.GetTicketsByDestination("Brazil")
	if err != nil {
		panic(err)
	}
	fmt.Println("Total tickets for Brazil: ", ticketsCount)

	// Exercício 2:
	countByPeriod, countAll := 0, 0
	// Dawn
	countByPeriod, _ = tickets.GetTicketsByPeriod(tickets.Dawn)
	countAll += countByPeriod
	fmt.Printf("Total tickets in the %s: %d\n", tickets.Dawn, countByPeriod)
	// Morning
	countByPeriod, _ = tickets.GetTicketsByPeriod(tickets.Morning)
	countAll += countByPeriod
	fmt.Printf("Total tickets in the %s: %d\n", tickets.Morning, countByPeriod)
	// Afternoon
	countByPeriod, _ = tickets.GetTicketsByPeriod(tickets.Afternoon)
	countAll += countByPeriod
	fmt.Printf("Total tickets in the %s: %d\n", tickets.Afternoon, countByPeriod)
	// Evening
	countByPeriod, _ = tickets.GetTicketsByPeriod(tickets.Evening)
	countAll += countByPeriod
	fmt.Printf("Total tickets in the %s: %d\n", tickets.Evening, countByPeriod)
	// Total tickets count to assure all tickets are counted
	fmt.Printf("Total tickets count: %d\n", countAll)

	// Exercício 3:
	destination, err := tickets.DestinationPercentage("Brazil")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Percentage of tickets for Brazil: %.2f%%\n", destination*100)

}
