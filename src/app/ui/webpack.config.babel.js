import { resolve } from "path";
import { CleanWebpackPlugin } from "clean-webpack-plugin";
import HtmlWebpackPlugin from "html-webpack-plugin";
import MiniCssExtractPlugin from "mini-css-extract-plugin";
import CopyWebpackPlugin from "copy-webpack-plugin";

// Try the environment variable, otherwise use root.
const ASSET_PATH = process.env.ASSET_PATH || "/";

var Webpack = {
  entry: "./script/index.js",
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

export default Webpack;
