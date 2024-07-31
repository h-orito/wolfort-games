/** @type {import('next').NextConfig} */
const path = require('path')

const nextConfig = {
    basePath: '/games',
    images: {
        domains: [
            "wolfort.net",
            "image.wolfort.dev"
        ]
    },
    sassOptions: {
        includePaths: [path.join(__dirname, 'styles')],
    },
}

module.exports = nextConfig
