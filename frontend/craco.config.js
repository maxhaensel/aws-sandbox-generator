const webpack = require('webpack')

module.exports = function (env) {
  const myPlugins = []

  if (process.env.REACT_APP_MOCK_API) {
    myPlugins.push(
      new webpack.NormalModuleReplacementPlugin(
        /(.*)production(\.*)/,
        function (resource) {
          resource.request = resource.request.replace(
            /production/,
            `development`,
          )
        },
      ),
    )
  }

  return {
    webpack: {
      plugins: myPlugins,
    },
  }
}
