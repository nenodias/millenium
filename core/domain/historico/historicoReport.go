package historico

import (
	"fmt"
	"net/http"

	"github.com/go-pdf/fpdf"
	clienteDomain "github.com/nenodias/millenium/core/domain/cliente"
	modeloDomain "github.com/nenodias/millenium/core/domain/modelo"
	montadoraDomain "github.com/nenodias/millenium/core/domain/montadora"
	tecnicoDomain "github.com/nenodias/millenium/core/domain/tecnico"
	"github.com/nenodias/millenium/core/domain/utils"
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
	Historico       Historico
	Cliente         clienteDomain.Cliente
	Veiculo         veiculoDomain.Veiculo
	Modelo          modeloDomain.Modelo
	Montadora       montadoraDomain.Montadora
	Tecnico         tecnicoDomain.Tecnico
	SubTotalServico float64
	SubTotalPeca    float64
	Total           float64
}

func GenerateReport(historico *HistoricoReport, w http.ResponseWriter) {
	Calculate(historico)
	pdf := fpdf.New("P", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Arial", "B", 16)
	pdf.AliasNbPages("{nb}")
	pdf.AddPage()
	MakeHeader(historico, pdf, tr)
	MakeFooter(historico, pdf, tr)
	err := pdf.Output(w)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}

func MakeFooter(historico *HistoricoReport, pdf *fpdf.Fpdf, tr func(string) string) {
	pdf.SetY(-15)
	pdf.SetFont("Arial", "I", 8)
	pageNum := fmt.Sprintf("%d/{nb}", pdf.PageNo())
	pdf.CellFormat(0, 10, pageNum, "", 0, "C", false, 0, "")
}

func MakeHeader(historico *HistoricoReport, pdf *fpdf.Fpdf, tr func(string) string) {
	logo := "d://workspace//logo_oficina.png"
	pdf.Image(logo, 8, 10, 25, 0, false, "", 0, "")
	pdf.CellFormat(25, 0, "", "", 0, "", false, 0, "")
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(100, 4, tr(REPORT_NAME), "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 9)
	pdf.CellFormat(25, 0, "", "", 0, "", false, 0, "")
	pdf.CellFormat(100, 4, tr(REPORT_ADDRESS), "", 1, "L", false, 0, "")
	pdf.CellFormat(25, 0, "", "", 0, "", false, 0, "")
	pdf.CellFormat(80, 4, "Fone: "+tr(REPORT_PHONE), "", 0, "L", false, 0, "")
	pdf.CellFormat(80, 4, "Celular: "+tr(REPORT_MOBILE), "", 1, "L", false, 0, "")
	pdf.CellFormat(25, 0, "", "", 0, "", false, 0, "")
	pdf.CellFormat(100, 4, "Email: "+tr(REPORT_EMAIL), "", 0, "L", false, 0, "")
	pdf.Line(5, 27, 205, 27)

	MakeCustomer(historico, pdf, tr)
}

func MakeCustomer(historico *HistoricoReport, pdf *fpdf.Fpdf, tr func(string) string) {
	pdf.Ln(4)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(20, 7, "Cliente       :", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(100, 7, tr(historico.Cliente.Nome), "", 0, "L", false, 0, "")

	pdf.Ln(4)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(20, 7, tr("Endereço  :"), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(100, 7, tr(GetEnderecoCliente(historico)), "", 0, "L", false, 0, "")

	pdf.Ln(4)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(21, 7, "Telefone    :", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(50, 7, tr(historico.Cliente.Telefone), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(18, 7, " - Celular:", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(50, 7, tr(historico.Cliente.Celular), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(23, 7, " - Fone Coml:", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(50, 7, tr(historico.Cliente.TelefoneComercial), "", 0, "L", false, 0, "")

	pdf.Ln(4)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(20, 7, "CPF/CNPJ :", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(100, 7, tr(historico.Cliente.CPF), "", 0, "L", false, 0, "")

	pdf.Ln(4)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(20, 7, tr("Técnico     :"), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(100, 7, tr(historico.Tecnico.Nome), "", 0, "L", false, 0, "")

	pdf.Ln(4)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(20, 7, tr("Veículo      :"), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(100, 7, tr(GetVeiculoDescription(historico)), "", 0, "L", false, 0, "")

	pdf.Line(5, 54, 205, 54)

	pdf.Ln(9)
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(50, 8, tr(historico.Historico.Tipo.String()), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(110, 8, fmt.Sprintf("%d", historico.Historico.NumeroOrdem), "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 8, utils.DateToString(historico.Historico.Data), "", 0, "L", false, 0, "")
	pdf.Line(5, 64, 205, 64)

	pdf.Ln(10)
	pdf.Rect(5, 5, 200, 287, "D")
}

func GetVeiculoDescription(historico *HistoricoReport) string {
	veiculo := ""
	veiculo += historico.Veiculo.Placa
	veiculo += "-"
	veiculo += historico.Modelo.Nome
	veiculo += "-"
	veiculo += historico.Montadora.Nome
	veiculo += "- Ano: "
	veiculo += historico.Veiculo.Ano
	veiculo += "- Km: "
	veiculo += fmt.Sprintf("%.0f", historico.Historico.Kilometragem)
	return veiculo
}

func GetEnderecoCliente(historico *HistoricoReport) string {
	endereco := ""
	endereco += historico.Cliente.Complemento
	endereco += " - "
	endereco += historico.Cliente.Bairro
	endereco += " - "
	endereco += historico.Cliente.Cidade
	endereco += " - "
	endereco += historico.Cliente.Estado
	return endereco
}

func Calculate(historico *HistoricoReport) {
	for _, v := range historico.Historico.Items {
		if v.Tipo == PECA {
			historico.SubTotalPeca += (float64(v.Quantidade) * v.Valor)
		} else if v.Tipo == SERVICO {
			historico.SubTotalServico += (float64(v.Quantidade) * v.Valor)
		}
	}
	historico.Total = historico.SubTotalPeca + historico.SubTotalServico
}
