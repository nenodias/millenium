import login from './login'
import index from './components/index'

const template = `<section class="app">
  <login v-if="!logged"></login>
  <index v-if="logged"></index>
</section>
`

export default {
  name: 'App',
  components:	{
    login,
    index
  },
  computed: {
    logged: function () {
      return this.$store.getters.login_done
    }
  },
  template: template
}
