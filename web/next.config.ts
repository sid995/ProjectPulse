import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  output: 'standalone',
  // Enable hot reloading in Docker
  webpack: (config, { dev, isServer }) => {
    // Only apply in development mode
    if (dev && !isServer) {
      // Add support for hot module replacement in Docker
      config.watchOptions = {
        ...config.watchOptions,
        poll: 1000, // Check for changes every second
        aggregateTimeout: 300, // Delay before rebuilding
      };
    }
    return config;
  },
};

export default nextConfig;
