package loader

import (
	"app/internal"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (l *LoaderTicketCSV) Load() (t map[int]internal.TicketAttributes, err error) {
	// open the file
	f, err := os.Open(l.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	t = make(map[int]internal.TicketAttributes)
	for {
		record, errr := r.Read()
		if errr != nil {
			if errr == io.EOF {
				break
			}

			err = fmt.Errorf("error reading record: %v", err)
			return
		}

		// serialize the record
		id, err := strconv.Atoi(record[0])
		if err != nil {

		}
		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {

		}
		ticket := internal.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   price,
		}

		// add the ticket to the map
		t[id] = ticket
	}

	return
}
