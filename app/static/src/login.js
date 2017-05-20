const template = `<section class="hero is-fullheight is-dark is-bold">
    <div class="hero-body">
        <div class="container">
            <div class="columns is-vcentered">
                <div class="column is-4 is-offset-4">
                    <h1 class="title">
                        Login
                    </h1>
                    <div class="box">
                        <label class="label">Login</label>
                        <p class="control">
                            <input class="input upper" v-model="usuario" type="text" name="usuario">
                        </p>
                        <label class="label">Password</label>
                        <p class="control">
                            <input class="input upper" type="password" v-model="senha" name="senha">
                        </p>
                        <hr>
                        <p class="control">
                            <button @click="login" :class="{ 'button is-primary': true, 'is-loading': loading }" type="button">Login</button>
                            <button @click="redirect" class="button is-default">Cancel</button>
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</section>
`

export default {
  name: 'login',
  data () {
    return {
        usuario:'',
        senha:''
    }
  },
  computed: {
    loading: function () {
      return this.$store.getters.processing
    }
  },
  methods: {
    login () {
      const usuario = this.usuario;
      const senha = this.senha;
      const dados = {
        usuario,
        senha
      }
      console.log(this.$store)
      this.$store.dispatch('doLogin', dados)
    },
    redirect () {
      window.location.href = "http://google.com"
    }
  },
  template: template
}
