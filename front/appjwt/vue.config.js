// vue.config.js
module.exports = {
  devServer: {
    port: 8010,
    proxy: 'http://localhost:8011',
}
}