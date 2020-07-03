const { resolve } = require("path");
const { CleanWebpackPlugin } = require("clean-webpack-plugin");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const CopyWebpackPlugin = require("copy-webpack-plugin");

// Try the environment variable, otherwise use root.
const ASSET_PATH = process.env.ASSET_PATH || "/";

module.exports = {
  entry: "./src/index.js",
  plugins: [
    new CleanWebpackPlugin(),
    new HtmlWebpackPlugin({
      filename: "./index.html",
      title: "gomithrilapp",
    }),
    new MiniCssExtractPlugin({
      filename: "static/[name].[contenthash].css",
    }),
    new CopyWebpackPlugin(
      [
        { from: "./static/healthcheck.html", to: "static/" },
        { from: "./static/swagger.json", to: "static/" },
      ],
      { copyUnmodified: true }
    ),
  ],
  resolve: {
    alias: {
      "@": resolve(__dirname, "./src"),
      "~": resolve(__dirname),
    },
  },
  output: {
    path: resolve(__dirname, "./dist"),
    filename: "static/[name].[contenthash].js",
    publicPath: ASSET_PATH,
  },
  optimization: {
    splitChunks: {
      chunks: "all",
    },
  },
  performance: {
    hints: false,
  },
  module: {
    rules: [
      {
        enforce: "pre",
        test: /\.js$/,
        exclude: /node_modules/,
        loader: "eslint-loader",
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: "babel-loader",
      },
      {
        test: /\.scss$/,
        use: [
          MiniCssExtractPlugin.loader,
          {
            loader: "css-loader",
          },
          {
            loader: "sass-loader",
            options: {
              sourceMap: true,
            },
          },
        ],
      },
    ],
  },
};
