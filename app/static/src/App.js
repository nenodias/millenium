const template = `<section class="todoapp">
    <header class="header">
      <h1>Tarefas {{ $route.params.id }}</h1>
      {{ msg }}
    </header>
  </section>
`

async function getContent () {
    try {
      let config = {
        method: 'GET'
      }
      let response = await fetch('https://api.icndb.com/jokes/random', config)
      console.log(response)
      return await response.json()
    } catch (err) {
      console.log(err)
    }
}

export default {
  name: 'app',
  data () {
    return {
        msg:''
    }
  },
  beforeMount () {
    getContent().then((data) => {
      console.log(data)
      this.msg = data.value.joke
    })
  },
  template: template
}
