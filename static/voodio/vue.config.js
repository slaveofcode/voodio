module.exports = {
  runtimeCompiler: true,
  chainWebpack: (config) => {
    config.plugins.delete('prefetch')
  }
}