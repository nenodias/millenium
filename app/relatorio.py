from pdb import set_trace
import os, sys
from io import BytesIO, StringIO
from fpdf import FPDF
from app.models import tupla_tipo_historico

class MyPDF(FPDF):
    
    def __init__(self,dados):
        super(MyPDF, self).__init__()
        self.dados = dados
        self.tipo = tupla_tipo_historico[0][1] if (self.dados['tipo'] == tupla_tipo_historico[0][0]) else tupla_tipo_historico[1][1]
        self.numero_ordem = str(self.dados['numero_ordem'])
        self.nome = self.dados['cliente']['nome']
     
    #----------------------------------------------------------------------
    def header(self):
        base = os.path.dirname(os.path.abspath(__file__))
        logo = os.path.join(base,"static/img/carrit.png")
        self.image(logo, x=10, y=10, w=23) # logo
        self.cell(w=25, h = 0, txt = '', border = 0, ln = 0, align = '', fill = False, link = '')# empty space for logo

        self.set_font("Arial", style="B", size=15)
        self.cell(w=100, h = 10, txt = self.tipo, border = 1, ln = 0, align = 'C')
        self.cell(w=10, h = 10, txt = 'N°', border = 1, ln = 0, align = 'C')
        self.set_font("Arial", style="B", size=20)
        self.cell(0,10, self.numero_ordem, border=1, ln=0, align="C")
        self.ln(10) # Line Break

        self.set_font("Arial", style="B", size=11)
        self.cell(w=25, h = 10, txt = 'Nome:', border=1, ln=0, align='R')
        self.set_font("Arial", size=11)
        self.cell(0,10, self.nome, border=1, ln=0, align="C")
        self.set_font("Arial", style="B", size=15)
        self.ln(10)
 
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

def gerar_pdf(dados):
    pdf = MyPDF(dados)
    pdf.alias_nb_pages()
    pdf.set_font('Arial', 'B', 16)
    pdf.add_page()
    pdf.cell(40, 10, 'Hello World!')
    pdf.output(dest='S')
    byte_string = pdf.output(dest="S")  # Probably what you want
    stream = BytesIO(byte_string) 
    return stream