#!/bin/bash

# DadMail Initial Setup Script

set -e

echo "========================================="
echo "DadMail - Initial Setup"
echo "========================================="
echo ""

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "Error: Docker is not installed. Please install Docker first."
    echo "Visit: https://docs.docker.com/get-docker/"
    exit 1
fi

if ! command -v docker-compose &> /dev/null && ! docker compose version &> /dev/null; then
    echo "Error: Docker Compose is not installed. Please install Docker Compose first."
    exit 1
fi

# Check if Go is installed (for local development)
if ! command -v go &> /dev/null; then
    echo "Warning: Go is not installed. You'll need it for local backend development."
    echo "Visit: https://golang.org/doc/install"
fi

# Check if Node.js is installed (for local development)
if ! command -v node &> /dev/null; then
    echo "Warning: Node.js is not installed. You'll need it for local frontend development."
    echo "Visit: https://nodejs.org/"
fi

echo "1. Creating .env files from examples..."

# Backend .env
if [ ! -f .env ]; then
    if [ -f .env.example ]; then
        cp .env.example .env
        echo "   Created .env from .env.example"
        echo "   ⚠️  IMPORTANT: Update .env with secure values for production!"
    else
        echo "   Warning: .env.example not found"
    fi
else
    echo "   .env already exists, skipping..."
fi

# Frontend .env
if [ ! -f frontend/.env ]; then
    if [ -f frontend/.env.example ]; then
        cp frontend/.env.example frontend/.env
        echo "   Created frontend/.env from frontend/.env.example"
    else
        echo "   Warning: frontend/.env.example not found"
    fi
else
    echo "   frontend/.env already exists, skipping..."
fi

echo ""
echo "2. Building Docker images..."
cd docker
docker-compose build

echo ""
echo "3. Starting services..."
docker-compose up -d postgres redis minio

echo ""
echo "4. Waiting for services to be healthy..."
sleep 10

echo ""
echo "5. Running database migrations..."
docker-compose exec -T postgres psql -U dadmail -d dadmail -f /docker-entrypoint-initdb.d/001_initial.sql || true

echo ""
echo "========================================="
echo "Setup Complete!"
echo "========================================="
echo ""
echo "Next steps:"
echo "  1. Review and update .env files with secure values"
echo "  2. Run './scripts/dev.sh' to start the development environment"
echo "  3. Access the application:"
echo "     - Frontend: http://localhost:5173"
echo "     - Backend API: http://localhost:8080"
echo "     - MinIO Console: http://localhost:9001 (admin/minioadmin)"
echo ""
echo "Useful commands:"
echo "  - Stop all services: cd docker && docker-compose down"
echo "  - View logs: cd docker && docker-compose logs -f"
echo "  - Reset database: cd docker && docker-compose down -v"
echo ""
