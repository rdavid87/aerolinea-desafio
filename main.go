package main

import (
	"fmt"
	"os"
	"primer/entregable/internal/tickets"
	"strconv"
	"strings"
)

func main() {

	rawData := loadFileCSV("./tickets.csv")

	lista := readFile(rawData)

	//fmt.Println("Tickets:", lista)

	destino := "China"
	jornada := tickets.MADRUGADA

	total, _ := tickets.GetTotalTickets(lista, destino)
	fmt.Printf("Total de tickets con destino %s: %d", destino, total)
	fmt.Println()

	total, _ = tickets.GetMornings(lista, jornada)
	fmt.Printf("Total de personas que viajan en jornada %s: %d", jornada, total)
	fmt.Println()

	porcentaje, _ := tickets.AverageDestination(lista, destino)
	fmt.Printf("Porcentaje de personas con destino %s: %.2f %s", destino, porcentaje, "%")
	fmt.Println()
}

func loadFileCSV(nombreFile string) string {
	rawData, err := os.ReadFile(nombreFile)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return string(rawData)

}

func readFile(rawData string) tickets.TicketArr {
	lista := tickets.TicketArr{}

	data := strings.Split(rawData, "\n")

	for _, v := range data {
		linea := strings.TrimSpace(v)

		if len(linea) > 0 {
			column := strings.Split(linea, ",")
			if len(column) == 6 {
				precio, _ := strconv.ParseFloat(column[5], 64)
				id, _ := strconv.ParseInt(column[0], 64, 0)
				ticket := tickets.Ticket{
					Id:             id,
					PaisDestino:    column[3],
					PasajeroNombre: column[1],
					PasajeroEmail:  column[2],
					HoraVuelo:      column[4],
					Precio:         precio,
				}
				lista = lista.AddTicket(ticket)
			} else {
				printerErrorLine(column)
			}
		}
	}

	return lista
}

func printerErrorLine(data []string) {
	fmt.Println("Error de estructura en la Linea=> ", data)
}
