/* eslint-disable @typescript-eslint/no-var-requires */

const wp = require('@cypress/webpack-preprocessor')

// https://nuxtjs.org/api/nuxt/
// https://github.com/bahmutov/cypress-vue-unit-test/issues/200
// https://github.com/hex-digital/nuxt-cypress-example/blob/master/plugins/cypress.js
// https://github.com/bahmutov/cypress-vue-unit-test#bundling
// https://gist.github.com/danielroe/72618a5617c637e15613a51a911221ba
// https://stackoverflow.com/questions/52168030/how-to-do-cypress-unit-tests-with-vue-cli

const { Nuxt, Builder } = require('nuxt')
const config = require('../../../nuxt.config.js')
const nuxt = new Nuxt({
  ...config,
  dev: false
})

// process.env.DEBUG = 'nuxt:*'

/* eslint-enable */
config.modules.forEach(module => nuxt.moduleContainer.addModule(module))
config.buildModules.forEach(module => nuxt.moduleContainer.addModule(module))

const builder = new Builder(nuxt)
const nuxtWebpack = builder.bundleBuilder.getWebpackConfig('Client')

const webpackOptions = {
  resolve: nuxtWebpack.resolve,
  module: {
    rules: nuxtWebpack.module.rules
  },
  plugins: nuxtWebpack.plugins
}

module.exports = wp({
  webpackOptions
})
