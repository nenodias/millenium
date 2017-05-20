
export async function login ({usuario, senha}){
  let resp = await fetch('/login',{
    method:"POST",
    headers: {
      "Content-Type": "application/json",
    }, body:{
      usuario,
      senha
    }
  })
  let dados = await resp.json()
  return dados
}
