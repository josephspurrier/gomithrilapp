// Source: https://github.com/bahmutov/cypress-vue-unit-test/issues/200
// Source: https://gist.github.com/danielroe/72618a5617c637e15613a51a911221ba

// This file will extrack the webpack config from the nuxt.config.js so it can
// be used with Cypress in the index.js file.

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

module.exports = webpackOptions
