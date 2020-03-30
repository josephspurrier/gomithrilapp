import { resolve } from "path";

var Webpack = {
  entry: "./src/index.js",
  output: {
    path: resolve(__dirname, "./bin"),
    filename: "app.js"
  },
  module: {
    rules: [
      {
        enforce: "pre",
        test: /\.js$/,
        exclude: /node_modules/,
        loader: "eslint-loader"
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: "babel-loader"
      }
    ]
  }
};

export default Webpack;
