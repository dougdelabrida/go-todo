/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  // allow hot reload via docker
  webpackDevMiddleware: (config) => {
    config.watchOptions = {
      poll: 1000,
      aggregateTimeout: 300,
    }
    return config
  },
  rewrites: () => {
    return [
      {
        source: '/server/:path*',
        destination: 'http://todo-server:4000/:path*',
      },
    ]
  },
}

module.exports = nextConfig
