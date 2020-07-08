const path = require('path');
const rootWebpack = require('../webpack.config.js');

module.exports = {
    stories: ['../src/**/*.stories.js'],
    addons: [
        '@storybook/addon-storysource',
        '@storybook/addon-actions/register',
        '@storybook/addon-knobs/register',
        '@storybook/addon-a11y/register',
    ],
    webpackFinal: (config) => {
        return {
            ...config,
            plugins: [
                ...config.plugins,
                ...rootWebpack.plugins,
            ],
            resolve: {
                alias: {
                    ...config.resolve.alias,
                    ...rootWebpack.resolve.alias,
                },
            },
            module: {
                ...config.module,
                rules: [
                    ...config.module.rules,
                    ...rootWebpack.module.rules,
                ]
            }
        };
    },
};