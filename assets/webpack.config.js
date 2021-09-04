module.exports = {
  mode: "development",
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: "ts-loader",
      },
    ],
  },
  resolve: {
    extensions: [".ts", ".tsx", ".js", ".json"],
  },
  target: ["web", "es5"],
  name: "private",
  entry: `./src/main.tsx`,
  output: {
    path: `${__dirname}/dist`,
    filename: "bundle.js",
  },
  devServer: {
    static: "dist",
    proxy: {
      "/api": "http://localhost:9080",
      "/auth": "http://localhost:9080",
    },
    hot: true,
    historyApiFallback: {
      rewrites: [{ from: /^\/*/, to: "/index.html" }],
    },
  },
};
