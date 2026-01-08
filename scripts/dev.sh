#!/bin/bash

# DadMail Development Environment Starter

set -e

echo "========================================="
echo "DadMail - Development Environment"
echo "========================================="
echo ""

# Check if setup has been run
if [ ! -f .env ]; then
    echo "Error: .env file not found. Please run ./scripts/setup.sh first"
    exit 1
fi

cd docker

echo "Starting all services..."
docker-compose up -d

echo ""
echo "========================================="
echo "Development Environment Started!"
echo "========================================="
echo ""
echo "Services:"
echo "  Frontend:       http://localhost:5173"
echo "  Backend API:    http://localhost:8080"
echo "  API Health:     http://localhost:8080/health"
echo "  MinIO Console:  http://localhost:9001 (minioadmin/minioadmin)"
echo ""
echo "Database:"
echo "  PostgreSQL:     localhost:5432 (dadmail/dadmail_dev_password)"
echo "  Redis:          localhost:6379"
echo ""
echo "Useful commands:"
echo "  - View logs:           docker-compose logs -f"
echo "  - View backend logs:   docker-compose logs -f backend"
echo "  - View frontend logs:  docker-compose logs -f frontend"
echo "  - Stop services:       docker-compose down"
echo "  - Restart backend:     docker-compose restart backend"
echo "  - Restart frontend:    docker-compose restart frontend"
echo ""
echo "Hot reload is enabled for both backend and frontend."
echo "Changes to source files will automatically trigger rebuilds."
echo ""
