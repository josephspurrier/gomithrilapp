/// <reference types="cypress" />
// ***********************************************************
// This example plugins/index.js can be used to load plugins
//
// You can change the location of this file or turn off loading
// the plugins file with the 'pluginsFile' configuration option.
//
// You can read more here:
// https://on.cypress.io/plugins-guide
// ***********************************************************

// This function is called when a project is opened or re-opened (e.g. due to
// the project's config changing)

// webpack import
// const webpack = require('@cypress/webpack-preprocessor')
const webpackPreProcessor = require('./ts-preprocessor')

/**
 * @type {Cypress.PluginConfig}
 */
module.exports = (on, config) => {
  // `on` is used to hook into various events Cypress emits
  // `config` is the resolved Cypress config
  // file:processor event
  /* on(
    'file:preprocessor',
    webpack({
      // webpackOptions: require('@vue/cli-service/webpack.config'),
      // webpackOptions: require('../../../src/index.js'),
      watchOptions: {}
    })
  ) */

  on('file:preprocessor', webpackPreProcessor)

  // on('file:preprocessor', webpackPreProcessor)

  return config

  // // return Object.assign({})
  // return Object.assign({}, config, {
  //   baseUrl: 'http://localhost:8080',
  //   testFiles: '**/*.spec.js',
  //   fixturesFolder: 'test/e2e/fixtures',
  //   integrationFolder: 'test/e2e/integration',
  //   pluginsFile: 'test/e2e/plugins/index.js',
  //   screenshotsFolder: 'test/e2e/screenshots',
  //   supportFile: 'test/e2e/support/index.js',
  //   videosFolder: 'test/e2e/videos'
  // })
}
