from pdb import set_trace
import os, sys
from io import BytesIO, StringIO
from fpdf import FPDF
from app.models import tupla_tipo_historico
from app.utils import from_str_to_datetime_or_none, format_date

class DadosOficina():

    def __init__(self):
       self.nome = 'AUTO MECÂNICA CARRIT' 
       self.endereco = 'R: Rio Grande do Sul, 460 - J.D.CRUZIERO - Lençóis Paulista - SP - 186806-380'
       self.telefone = '(14)3264-4598'
       self.celular = '(14)991262313'
       self.email = 'mecanicacarrit@gmail.com'



class MyPDF(FPDF):
    
    def __init__(self,dados):
        super(MyPDF, self).__init__()
        self.oficina = DadosOficina()
        self.dados = dados
        self.tipo = tupla_tipo_historico[0][1] if (dados['tipo'] == tupla_tipo_historico[0][0]) else tupla_tipo_historico[1][1]
        self.numero_ordem = str(dados['numero_ordem'])
        self.cliente = dados['cliente']['nome']
        self.endereco = dados['cliente']['endereco']
        self.endereco += ' - '+ dados['cliente']['complemento']
        self.endereco += ' - '+ dados['cliente']['bairro']
        self.endereco += ' - '+ dados['cliente']['cidade']
        self.endereco += ' - '+ dados['cliente']['estado']
        self.telefone = dados['cliente']['telefone']
        self.celular = dados['cliente']['celular']
        self.telefone_comercial = dados['cliente']['telefone_comercial']
        self.cpf = dados['cliente']['cpf']
        self.tecnico = ''
        if self.dados.get('tecnico'):
            self.tecnico = dados['tecnico']['nome']
        veiculo = ''
        ano = ''
        if self.dados.get('veiculo'):
            veiculo += dados['veiculo']['placa']
            ano = str(dados['veiculo']['ano'])
        if self.dados.get('modelo'):
            veiculo += ' - '+ dados['modelo']['nome']
        if self.dados.get('montadora'):
            veiculo += ' - '+ dados['montadora']['nome']
        veiculo += ' - Ano: '+ ano
        veiculo += ' - Km: '+ '%.0f'%(dados['vistoria']['kilometragem'])
        self.veiculo = veiculo
        self.data = format_date( from_str_to_datetime_or_none(dados['data']) )
        

        self.falhas = []
        self.servicos = []
        self.pecas = []

        self.sub_total_servicos = 0
        self.sub_total_pecas = 0

        for item in dados['items']:
            if item['tipo'] == 'F':
                self.falhas.append( item )
            elif item['tipo'] == 'P':
                self.pecas.append( item )
                self.sub_total_pecas += item['valor']
            elif item['tipo'] == 'S':
                self.servicos.append( item )
                self.sub_total_servicos += item['valor']

        self.total = self.sub_total_servicos + self.sub_total_pecas
        self.observacoes = dados['observacao']
     
    #----------------------------------------------------------------------
    def header(self):
        base = os.path.dirname(os.path.abspath(__file__))
        logo = os.path.join(base,"static/img/carrit.png")
        self.image(logo, x=10, y=10, w=23) # logo
        self.cell(w=25, h = 0, txt = '', border = 0, ln = 0, align = '', fill = False, link = '')# empty space for logo

        # Dados Oficina

        self.set_font("Arial", style="B", size=12)
        self.cell(w=100, h = 4, txt = self.oficina.nome, border = 0, ln = 0, align = 'L')
        self.ln(4)
        
        self.set_font("Arial", style="", size=9)
        self.cell(w=25, h = 0, txt = '', border = 0, ln = 0, align = '', fill = False, link = '')
        self.cell(w=100, h = 4, txt = self.oficina.endereco, border = 0, ln = 0, align = 'L')
        self.ln(4)
        
        self.cell(w=25, h = 0, txt = '', border = 0, ln = 0, align = '', fill = False, link = '')
        self.cell(w=80, h = 4, txt = 'Fone: %s'%(self.oficina.telefone), border = 0, ln = 0, align = 'L')
        self.cell(w=80, h = 4, txt = 'Celular: %s'%(self.oficina.celular), border = 0, ln = 0, align = 'L')
        self.ln(4)
        
        self.cell(w=25, h = 0, txt = '', border = 0, ln = 0, align = '', fill = False, link = '')
        self.cell(w=100, h = 4, txt = 'E-mail: %s'%(self.oficina.email), border = 0, ln = 0, align = 'L')
        self.ln(4)

        self.line(5, 27, 205, 27)
        
        # Dados Cliente
        
        self.ln(1)
        self.set_font("Arial", style="B", size=10)
        self.cell(w=20, h = 7, txt ='Cliente      :', border = 0, ln = 0, align = 'L')
        self.set_font("Arial", style="", size=10)
        self.cell(w=100, h = 7, txt =self.cliente, border = 0, ln = 0, align = 'L')

        self.ln(4)
        self.set_font("Arial", style="B", size=10)
        self.cell(w=20, h = 7, txt ='Endereço  :', border = 0, ln = 0, align = 'L')
        self.set_font("Arial", style="", size=10)
        self.cell(w=100, h = 7, txt =self.endereco, border = 0, ln = 0, align = 'L')

        self.ln(4)
        self.set_font("Arial", style="B", size=10)
        self.cell(w=5, h = 7, txt ='Telefone    :', border = 0, ln = 0, align = 'L')
        self.set_font("Arial", style="", size=10)
        self.cell(w=50, h = 7, txt =self.telefone, border = 0, ln = 0, align = 'L')
        self.set_font("Arial", style="B", size=10)
        self.cell(w=5, h = 7, txt =' - Celular:', border = 0, ln = 0, align = 'L')
        self.set_font("Arial", style="", size=10)
        self.cell(w=50, h = 7, txt =self.celular, border = 0, ln = 0, align = 'L')
        self.set_font("Arial", style="B", size=10)
        self.cell(w=5, h = 7, txt =' - Fone Coml:', border = 0, ln = 0, align = 'L')
        self.set_font("Arial", style="", size=10)
        self.cell(w=50, h = 7, txt =self.telefone_comercial, border = 0, ln = 0, align = 'L')

        self.ln(4)
        self.set_font("Arial", style="B", size=10)
        self.cell(w=20, h = 7, txt ='CPF/CNPJ :', border = 0, ln = 0, align = 'L')
        self.set_font("Arial", style="", size=10)
        self.cell(w=100, h = 7, txt =self.cpf, border = 0, ln = 0, align = 'L')

        self.ln(4)
        self.set_font("Arial", style="B", size=10)
        self.cell(w=20, h = 7, txt ='Técnico     :', border = 0, ln = 0, align = 'L')
        self.set_font("Arial", style="", size=10)
        self.cell(w=100, h = 7, txt =self.tecnico, border = 0, ln = 0, align = 'L')

        self.ln(4)
        self.set_font("Arial", style="B", size=10)
        self.cell(w=20, h = 7, txt ='Veículo     :', border = 0, ln = 0, align = 'L')
        self.set_font("Arial", style="", size=10)
        self.cell(w=100, h = 7, txt =self.veiculo, border = 0, ln = 0, align = 'L')

        self.line(5, 54, 205, 54)

        self.ln(8)
        self.set_font("Arial", style="B", size=14)
        self.cell(w=50, h = 8, txt =self.tipo, border = 0, ln = 0, align = 'L')
        self.set_font("Arial", style="B", size=16)
        self.cell(w=110, h = 8, txt =self.numero_ordem, border = 0, ln = 0, align = 'L')
        self.cell(w=20, h = 8, txt =self.data, border = 0, ln = 0, align = 'L')
        self.line(5, 64, 205, 64)
        
        self.ln(10)

        # Margem
        self.rect(5, 5, 200, 287, 'D')
 
    #----------------------------------------------------------------------
    def footer(self):
        """
        Footer on each page
        """
        # position footer at 15mm from the bottom
        self.set_y(-15)
 
        # set the font, I=italic
        self.set_font("Arial", style="I", size=8)
 
        # display the page number and center it
        pageNum = "Page %s/{nb}" % self.page_no()
        self.cell(0, 10, pageNum, align="C")

    def draw_falhas(self):
        def desenhar_titulo():
            self.set_font("Arial", style="B", size=12)
            self.cell(w=120, h = 8, txt ='FALHAS', border = 0, ln = 0, align = 'L')
            self.ln(8)
        desenhar_titulo()
        for falha in self.falhas:
            self.set_font("Arial", style="", size=10)
            self.cell(w=120, h = 8, txt =falha['descricao'], border = 0, ln = 0, align = 'L')
            self.ln(8)
            if self.y >= 275.0:
                self.add_page()
                desenhar_titulo()




    def draw_pecas(self):
        def desenhar_titulo():
            self.set_font("Arial", style="B", size=12)
            self.cell(w=120, h = 8, txt ='PEÇAS', border = 0, ln = 0, align = 'L')
            self.set_font("Arial", style="B", size=10)
            self.cell(w=25, h = 8, txt ='Qtd.', border = 0, ln = 0, align = 'C')
            self.cell(w=25, h = 8, txt ='Valor', border = 0, ln = 0, align = 'C')
            self.cell(w=25, h = 8, txt ='Total', border = 0, ln = 0, align = 'C')
            self.ln(8)
        desenhar_titulo()
        for peca in self.pecas:
            self.set_font("Arial", style="", size=9)
            self.cell(w=120, h = 6, txt =peca['descricao'], border = 1, ln = 0, align = 'L')
            self.cell(w=25, h = 6, txt ='%i'%(peca['quantidade']), border = 1, ln = 0, align = 'R')
            self.cell(w=25, h = 6, txt ='%.2f'%(peca['valor']), border = 1, ln = 0, align = 'R')
            self.cell(w=25, h = 6, txt ='%.2f'%(peca['quantidade']*peca['valor']), border = 1, ln = 0, align = 'R')
            self.ln(6)
            if self.y >= 275.0:
                self.add_page()
                desenhar_titulo()
        self.cell(w=165, h = 6, txt ='Sub Total:', border = 0, ln = 0, align = 'R')
        self.cell(w=30, h = 6, txt ='%.2f'%(self.sub_total_pecas), border = 0, ln = 6, align = 'R')

    def draw_servicos(self):
        def desenhar_titulo():
            self.set_font("Arial", style="B", size=12)
            self.cell(w=120, h = 8, txt ='SERVIÇOS', border = 0, ln = 0, align = 'L')
            self.set_font("Arial", style="B", size=10)
            self.cell(w=25, h = 8, txt ='Qtd.', border = 0, ln = 0, align = 'C')
            self.cell(w=25, h = 8, txt ='Valor', border = 0, ln = 0, align = 'C')
            self.cell(w=25, h = 8, txt ='Total', border = 0, ln = 0, align = 'C')
            self.ln(8)
        desenhar_titulo()
        for servico in self.servicos:
            self.set_font("Arial", style="", size=9)
            self.cell(w=120, h = 6, txt =servico['descricao'], border = 1, ln = 0, align = 'L')
            self.cell(w=25, h = 6, txt ='%i'%(servico['quantidade']), border = 1, ln = 0, align = 'R')
            self.cell(w=25, h = 6, txt ='%.2f'%(servico['valor']), border = 1, ln = 0, align = 'R')
            self.cell(w=25, h = 6, txt ='%.2f'%(servico['quantidade']*servico['valor']), border = 1, ln = 0, align = 'R')
            self.ln(6)
            if self.y >= 275.0:
                self.add_page()
                desenhar_titulo()
        self.cell(w=165, h = 6, txt ='Sub Total:', border = 0, ln = 0, align = 'R')
        self.cell(w=30, h = 6, txt ='%.2f'%(self.sub_total_servicos), border = 0, ln = 6, align = 'R')

        self.ln(2)
        self.set_font("Arial", style="B", size=10)
        self.cell(w=160, h = 6, txt ='Total:', border = 0, ln = 0, align = 'R')
        self.cell(w=35, h = 6, txt ='%.2f'%(self.total), border = 0, ln = 6, align = 'R')
        self.ln(1)

    def draw_observacoes(self):
        self.set_font("Arial", style="B", size=10)
        self.cell(w=160, h = 6, txt ='Observação:', border = 0, ln = 6, align = 'L')
        self.set_font("Arial", style="", size=9)
        self.write(5,self.observacoes)

def gerar_pdf(dados):
    pdf = MyPDF(dados)
    pdf.alias_nb_pages()
    pdf.set_font('Arial', 'B', 16)
    pdf.add_page()
    
    pdf.draw_falhas()
    pdf.draw_pecas()
    pdf.draw_servicos()
    pdf.draw_observacoes()

    byte_string = pdf.output(dest="S")  # Probably what you want
    stream = BytesIO(byte_string) 
    return stream