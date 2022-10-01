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
	"github.com/nenodias/millenium/config"
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
	Falhas          []HistoricoItem
	Pecas           []HistoricoItem
	Servicos        []HistoricoItem
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
	pdf.SetHeaderFunc(func() { MakeHeader(historico, pdf, tr) })
	pdf.SetFooterFunc(func() { MakeFooter(historico, pdf, tr) })
	pdf.AddPage()
	MakeFalhas(historico, pdf, tr)
	MakePecas(historico, pdf, tr)
	MakeServicos(historico, pdf, tr)
	MakeTotal(historico, pdf, tr)
	MakeObservacoes(historico, pdf, tr)
	err := pdf.Output(w)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}

func MakeFooter(historico *HistoricoReport, pdf *fpdf.Fpdf, tr func(string) string) {
	pdf.SetY(275)
	pdf.SetFont("Arial", "I", 8)
	pageNum := fmt.Sprintf("%d/{nb}", pdf.PageNo())
	pdf.CellFormat(0, 0, pageNum, "", 0, "C", false, 0, "")
}

func MakeHeader(historico *HistoricoReport, pdf *fpdf.Fpdf, tr func(string) string) {
	logo := config.GetEnv("LOGO_REPORT","d://workspace//logo_oficina.png")
	pdf.Image(logo, 8, 10, 25, 0, false, "", 0, "")
	pdf.CellFormat(25, 0, "", "", 0, "", false, 0, "")
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(100, 4, tr(REPORT_NAME), "", 0, "L", false, 0, "")
	pdf.Ln(4)
	pdf.SetFont("Arial", "", 9)
	pdf.CellFormat(25, 0, "", "", 0, "", false, 0, "")
	pdf.CellFormat(100, 4, tr(REPORT_ADDRESS), "", 0, "L", false, 0, "")
	pdf.Ln(4)
	pdf.CellFormat(25, 0, "", "", 0, "", false, 0, "")
	pdf.CellFormat(80, 4, "Fone: "+tr(REPORT_PHONE), "", 0, "L", false, 0, "")
	pdf.CellFormat(80, 4, "Celular: "+tr(REPORT_MOBILE), "", 0, "L", false, 0, "")
	pdf.Ln(4)
	pdf.CellFormat(25, 0, "", "", 0, "", false, 0, "")
	pdf.CellFormat(100, 4, "Email: "+tr(REPORT_EMAIL), "", 0, "L", false, 0, "")
	pdf.Ln(4)
	pdf.Line(5, 27, 205, 27)

	MakeCustomer(historico, pdf, tr)
}

func MakeCustomer(historico *HistoricoReport, pdf *fpdf.Fpdf, tr func(string) string) {
	pdf.Ln(1)
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

	pdf.Ln(8)
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(50, 8, tr(historico.Historico.Tipo.String()), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(110, 8, fmt.Sprintf("%d", historico.Historico.NumeroOrdem), "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 8, utils.DateToString(historico.Historico.Data), "", 0, "L", false, 0, "")
	pdf.Line(5, 64, 205, 64)

	pdf.Ln(10)
	pdf.Rect(5, 5, 200, 287, "D")
}

func MakeFalhas(historico *HistoricoReport, pdf *fpdf.Fpdf, tr func(string) string) {
	desenharTitulo := func() {
		pdf.Ln(1)
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(120, 8, "FALHAS", "", 0, "L", false, 0, "")
		pdf.Ln(8)
	}
	if len(historico.Falhas) > 0 {
		desenharTitulo()
		for _, falha := range historico.Falhas {
			pdf.SetFont("Arial", "", 10)
			pdf.CellFormat(120, 8, tr(falha.Descricao), "", 0, "L", false, 0, "")
			pdf.Ln(7)
			if pdf.GetY() > 275.0 {
				pdf.AddPage()
				desenharTitulo()
			}
		}
	}
}

func MakePecas(historico *HistoricoReport, pdf *fpdf.Fpdf, tr func(string) string) {
	desenharTitulo := func() {
		pdf.Ln(1)
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(120, 8, tr("PEÇAS"), "", 0, "L", false, 0, "")
		pdf.SetFont("Arial", "B", 10)
		pdf.CellFormat(25, 8, "Qtd.", "", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, "Valor", "", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, "Total", "", 0, "C", false, 0, "")
		pdf.Ln(8)
	}
	if len(historico.Pecas) > 0 {
		desenharTitulo()
		for _, peca := range historico.Pecas {
			pdf.SetFont("Arial", "", 9)
			pdf.CellFormat(120, 6, tr(peca.Descricao), "1", 0, "L", false, 0, "")
			pdf.CellFormat(25, 6, utils.IntToString(peca.Quantidade), "1", 0, "R", false, 0, "")
			pdf.CellFormat(25, 6, utils.FloatMoney(peca.Valor), "1", 0, "R", false, 0, "")
			pdf.CellFormat(25, 6, utils.FloatMoney(float64(peca.Quantidade)*peca.Valor), "1", 0, "R", false, 0, "")
			pdf.Ln(6)
			if pdf.GetY() > 275.0 {
				pdf.AddPage()
				desenharTitulo()
			}
		}
		pdf.CellFormat(165, 6, "Sub Total:", "", 0, "R", false, 0, "")
		pdf.CellFormat(30, 6, utils.FloatMoney(historico.SubTotalPeca), "", 6, "R", false, 0, "")
	}
}

func MakeServicos(historico *HistoricoReport, pdf *fpdf.Fpdf, tr func(string) string) {
	desenharTitulo := func() {
		pdf.Ln(1)
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(120, 8, tr("SERVIÇOS"), "", 0, "L", false, 0, "")
		pdf.SetFont("Arial", "B", 10)
		pdf.CellFormat(25, 8, "Qtd.", "", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, "Valor", "", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, "Total", "", 0, "C", false, 0, "")
		pdf.Ln(8)
	}
	if len(historico.Servicos) > 0 {
		desenharTitulo()
		for _, servico := range historico.Servicos {
			pdf.SetFont("Arial", "", 9)
			pdf.CellFormat(120, 6, tr(servico.Descricao), "1", 0, "L", false, 0, "")
			pdf.CellFormat(25, 6, utils.IntToString(servico.Quantidade), "1", 0, "R", false, 0, "")
			pdf.CellFormat(25, 6, utils.FloatMoney(servico.Valor), "1", 0, "R", false, 0, "")
			pdf.CellFormat(25, 6, utils.FloatMoney(float64(servico.Quantidade)*servico.Valor), "1", 0, "R", false, 0, "")
			pdf.Ln(6)
			if pdf.GetY() > 275.0 {
				pdf.AddPage()
				desenharTitulo()
			}
		}
		pdf.CellFormat(165, 6, "Sub Total:", "", 0, "R", false, 0, "")
		pdf.CellFormat(30, 6, utils.FloatMoney(historico.SubTotalServico), "", 6, "R", false, 0, "")
	}
}

func MakeTotal(historico *HistoricoReport, pdf *fpdf.Fpdf, tr func(string) string) {
	pdf.Ln(2)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(160, 6, "Total:", "", 0, "R", false, 0, "")
	pdf.CellFormat(35, 6, utils.FloatMoney(historico.Total), "", 6, "R", false, 0, "")
	pdf.Ln(1)
}

func MakeObservacoes(historico *HistoricoReport, pdf *fpdf.Fpdf, tr func(string) string) {
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(160, 6, tr("Observação:"), "", 6, "L", false, 0, "")
	pdf.SetFont("Arial", "", 9)
	pdf.Write(5, historico.Historico.Observacao)
}

func GetVeiculoDescription(historico *HistoricoReport) string {
	veiculo := ""
	veiculo += historico.Veiculo.Placa
	veiculo += " - "
	veiculo += historico.Modelo.Nome
	veiculo += " - "
	veiculo += historico.Montadora.Nome
	veiculo += " - Ano: "
	veiculo += historico.Veiculo.Ano
	veiculo += " - Km: "
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
		if v.Tipo == FALHA {
			historico.Falhas = append(historico.Falhas, v)
		} else if v.Tipo == PECA {
			historico.Pecas = append(historico.Pecas, v)
			historico.SubTotalPeca += (float64(v.Quantidade) * v.Valor)
		} else if v.Tipo == SERVICO {
			historico.Servicos = append(historico.Servicos, v)
			historico.SubTotalServico += (float64(v.Quantidade) * v.Valor)
		}
	}
	historico.Total = historico.SubTotalPeca + historico.SubTotalServico
}
