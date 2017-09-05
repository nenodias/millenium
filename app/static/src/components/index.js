const template = `<section class="todoapp">
    <header class="header">
      <h1>Tarefas {{ $route.params.id }}</h1>
      <p>{{ msg }}</p>
    </header>
  </section>
`

async function getContent () {
    try {
      let config = {
        method: 'GET'
      }
      let response = await fetch('https://api.icndb.com/jokes/random', config)
      return await response.json()
    } catch (err) {
      console.error(err)
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
      this.msg = data.value.joke
    })
  },
  template: template
}
