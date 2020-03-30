module.exports = {
    root: true,
    env: {
        browser: true,
        node: true,
        es6: true
    },
    parserOptions: {
        parser: 'babel-eslint',
        sourceType: "module",
        allowImportExportEverywhere: true
    },
    extends: [
        'eslint:recommended',
        'plugin:prettier/recommended',
        'plugin:mithril/recommended'
    ],
    plugins: ['prettier'],
    rules: {}
}
