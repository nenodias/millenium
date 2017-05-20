
export async function login ({usuario, senha}){
  const body = {
    usuario,
    senha
  }
  let resp = await fetch('/login',{
    method:"POST",
    headers: {
      "Content-Type": "application/json",
    }, body: JSON.stringify(body)
  })
  let dados = await resp.json()
  return dados
}
