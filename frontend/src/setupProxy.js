const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function (app) {
  const target =
    process.env.REACT_APP_API_PROXY_TARGET ||
    process.env.API_PROXY_TARGET ||
    'http://localhost:8080';

  app.use(
    '/api',
    createProxyMiddleware({
      target,
      changeOrigin: true,
      logLevel: 'warn',
      proxyTimeout: 10000,
    })
  );
};
