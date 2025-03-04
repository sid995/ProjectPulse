#!/bin/bash

# Exit on error
set -e

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Starting ProjectPulse Development Environment${NC}"
echo -e "${YELLOW}This will start all services with hot-reloading enabled${NC}"

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
  echo -e "${YELLOW}Docker is not running. Please start Docker first.${NC}"
  exit 1
fi

# Start development environment with docker-compose
echo -e "${GREEN}Starting Docker containers...${NC}"
docker-compose up --build 