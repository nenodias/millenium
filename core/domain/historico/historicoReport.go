package historico

import (
	"net/http"

	"github.com/go-pdf/fpdf"
	clienteDomain "github.com/nenodias/millenium/core/domain/cliente"
	modeloDomain "github.com/nenodias/millenium/core/domain/modelo"
	montadoraDomain "github.com/nenodias/millenium/core/domain/montadora"
	tecnicoDomain "github.com/nenodias/millenium/core/domain/tecnico"
	veiculoDomain "github.com/nenodias/millenium/core/domain/veiculo"
	"github.com/rs/zerolog/log"
)

const (
	REPORT_NAME    = "AUTO MECÂNICA BAPTISTA"
	REPORT_ADDRESS = "R: São José, 32 - Vila Irerê - Lençóis Paulista - SP - 18682-100"
	REPORT_PHONE   = "(14)3264-4598"
	REPORT_MOBILE  = "(14) 9 9126-2313"
	REPORT_EMAIL   = "mecanicacarrit@gmail.com"
)

type HistoricoReport struct {
	Historico Historico
	Cliente   clienteDomain.Cliente
	Veiculo   veiculoDomain.Veiculo
	Modelo    modeloDomain.Modelo
	Montadora montadoraDomain.Montadora
	Tecnico   tecnicoDomain.Tecnico
}

func GenerateReport(historico *HistoricoReport, w http.ResponseWriter) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "B", 16)
	pdf.AddPage()
	MakeHeader(historico, pdf)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.Output(w)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}

func MakeHeader(historico *HistoricoReport, pdf *fpdf.Fpdf) {
	logo := "d://workspace//logo_oficina.png"
	pdf.Image(logo, 8, 10, 25, 0, false, "", 0, "")
}
