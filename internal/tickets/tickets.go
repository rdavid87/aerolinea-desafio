package tickets

import (
	"errors"
	"strconv"
	"strings"
)

const (
	MADRUGADA string = "Madrugada"
	MANANA    string = "Mañana"
	TARDE     string = "Tarde"
	NOCHE     string = "Noche"
)

type Ticket struct {
	Id             int64
	PaisDestino    string
	PasajeroNombre string
	PasajeroEmail  string
	HoraVuelo      string
	Precio         float64
}

type TicketArr []Ticket

func (t *TicketArr) AddTicket(data Ticket) TicketArr {
	return append(*t, data)
}

// ejemplo 1
func GetTotalTickets(t []Ticket, destination string) (int, error) {
	total := 0
	for _, v := range t {
		if v.PaisDestino == destination {
			total += 1
		}
	}
	return total, nil
}

// ejemplo 2
func GetMornings(t []Ticket, periodoTiempo string) (int, error) {
	total := 0
	switch periodoTiempo {
	case MADRUGADA:
		total = searchRangeTime(t, Madrugada())
	case MANANA:
		total = searchRangeTime(t, Manana())
	case TARDE:
		total = searchRangeTime(t, Tarde())
	case NOCHE:
		total = searchRangeTime(t, Noche())
	}
	return total, nil
}

// ejemplo 3
func AverageDestination(t []Ticket, destination string) (float64, error) {

	total, err := GetTotalTickets(t, destination)
	porcentaje := 0.0
	if len(t) == 0 {
		return 0, errors.New("No hay información para hacer los calculos")
	}
	if err == nil {
		porcentaje = (float64(total) / float64(len(t))) * 100.0
	}
	return porcentaje, err
}

func Madrugada() []int {
	return []int{0, 6}
}

func Manana() []int {
	return []int{7, 12}
}

func Tarde() []int {
	return []int{13, 19}
}

func Noche() []int {
	return []int{20, 23}
}

func searchRangeTime(t []Ticket, rango []int) int {
	total := 0
	for _, v := range t {
		hora := strings.Split(v.HoraVuelo, ":")
		newHora, _ := strconv.ParseInt(hora[0], 8, 0)

		if int(newHora) >= rango[0] && int(newHora) <= rango[1] {
			total += 1
		}
	}
	return total
}
