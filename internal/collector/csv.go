package collector

import (
	"encoding/csv"
	"log"
	"os"
)

type Row struct {
	Address string
	Pk      string
}

func WriterCsv(row *Row) {
	records := [][]string{
		{"address", "pk"},
		{row.Address, row.Pk},
	}

	f, err := os.Create("wallets.csv ")
	defer f.Close()

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("Error writing record to file", err)
		}
	}
}
