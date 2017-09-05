// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import VueRouter from 'vue-router'
import VueUtils from './plugins/utils'
import routes from './routes'
import store from './store'

//Vue.config.productionTip = false //uncomment to make a prod
Vue.use(VueRouter)
Vue.use(VueUtils)

const router = new VueRouter({
  mode: 'hash',//mode: 'history',
  routes
})

/* eslint-disable no-new */
new Vue({
  router,
  store,
  el: '#app',
  template: `
  <div id="app">
      <transition name="fade" mode="out-in">
        <!-- componente será mostrado aqui -->
        <router-view class="view"></router-view>
      </transition>
    </div>
  `,
  components: { App }
})
