System.config({
  transpiler: 'plugin-babel',
  packages: {
      '.': {
          defaultJSExtensions: 'js'
      }
  },
  map: {
    'plugin-babel': './static/node_modules/systemjs-plugin-babel/plugin-babel.js',
    'systemjs-babel-build': './static/node_modules/systemjs-plugin-babel/systemjs-babel-browser.js',
    'vue': './static/node_modules/vue/dist/vue.js',
    'vue-router': './static/node_modules/vue-router/dist/vue-router.js',
    'vuex': './static/node_modules/vuex/dist/vuex.js'
  }
});
