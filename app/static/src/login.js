import * as actions from './store/modules/login/actions'

const template = `<section class="hero is-fullheight is-dark is-bold">
    <div class="hero-body">
        <div class="container">
            <div class="columns is-vcentered">
                <div class="column is-4 is-offset-4">
                    <div v-if="hasMessage" :class="{'notification': true, [message.type]: true }">
                        <button @click="clearMessage" class="delete"></button>
                        {{ message.text }}
                    </div>
                    <h1 class="title">
                        Login
                    </h1>
                    <div class="box">
                        <label class="label">Login</label>
                        <p class="control">
                            <input class="input upper" @keyup.enter="login" v-model="usuario" type="text" name="usuario">
                        </p>
                        <label class="label">Password</label>
                        <p class="control">
                            <input class="input upper" @keyup.enter="login" type="password" v-model="senha" name="senha">
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
      return this.$store.getters.login_processing
    },
    hasMessage: function () {
      return this.$store.getters.login_message != null
    },
    message: function () {
      if (this.$store.getters.token) {
        let tomorrow = new Date.today()
        tomorrow.addDays(1)
        this.$cookies.set('token', this.$store.getters.token, { expires: tomorrow })
      }
      return this.$store.getters.login_message
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
      this.$store.dispatch(actions.LOGIN_DO_LOGIN, dados)
    },
    clearMessage (){
      this.$store.dispatch(actions.LOGIN_CLEAR_MESSAGE)
    },
    redirect () {
      window.location.href = "http://google.com"
    }
  },
  beforeMount () {
    // Verificar cookie pegar token
    if(this.$cookies.get('token')){
      this.$store.dispatch(actions.LOGIN_VERIFY_COOKIE, this.$cookies.get('token'))
    }
  },
  template: template
}
