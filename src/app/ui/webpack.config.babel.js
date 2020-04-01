import { resolve } from "path";
import { CleanWebpackPlugin } from "clean-webpack-plugin";
import HtmlWebpackPlugin from "html-webpack-plugin";
import MiniCssExtractPlugin from "mini-css-extract-plugin";
import CopyPlugin from "copy-webpack-plugin";

var Webpack = {
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
    new CopyPlugin([
      { from: "./static/swagger", to: "static/swagger" },
      { from: "./static/healthcheck.html", to: "static/healthcheck.html" },
    ]),
  ],
  output: {
    path: resolve(__dirname, "./dist"),
    filename: "static/[name].[contenthash].js",
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
