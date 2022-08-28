package historico

import (
	"net/http"

	"github.com/go-pdf/fpdf"
	"github.com/rs/zerolog/log"
)

const (
	REPORT_NAME    = "AUTO MECÂNICA BAPTISTA"
	REPORT_ADDRESS = "R: São José, 32 - Vila Irerê - Lençóis Paulista - SP - 18682-100"
	REPORT_PHONE   = "(14)3264-4598"
	REPORT_MOBILE  = "(14) 9 9126-2313"
	REPORT_EMAIL   = "mecanicacarrit@gmail.com"
)

func GenerateReport(historico *HistoricoReport, w http.ResponseWriter) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.Output(w)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
