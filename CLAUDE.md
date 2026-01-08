# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**DadMail** is an email client designed specifically for seniors to reduce cognitive overhead when managing email. The project aims to:
- Provide clear labeling beyond simple spam filtering
- Build conversation profiles beyond traditional threading
- Offer co-pilot features for caregivers (visibility, rule creation, appointment understanding)
- Work alongside existing email tools

**Current Status**: Phase 1 (Foundation) complete. Backend (Go/Fiber), Frontend (React/Vite/TailwindCSS), and Docker development environment are set up.

## Issue Tracking with Beads

This project uses **Beads** (bd) for AI-native issue tracking. All issues live in `.beads/issues.jsonl` and sync with git.

**Quick Start:** Run `bd onboard` to get started with beads in this repository.

### Essential Commands

```bash
# Finding work
bd ready                  # Show issues ready to work (no blockers)
bd list --status=open     # All open issues
bd show <id>              # View issue details with dependencies

# Creating issues
bd create --title="..." --type=task|bug|feature --priority=2
# Priority: 0-4 or P0-P4 (0=critical, 2=medium, 4=backlog)

# Working on issues
bd update <id> --status=in_progress    # Claim work
bd close <id>                          # Mark complete
bd close <id1> <id2> ...              # Close multiple (more efficient)

# Dependencies
bd dep add <issue> <depends-on>   # Add dependency (issue depends on depends-on)
bd blocked                        # Show all blocked issues

# Sync
bd sync                   # Sync with git remote
bd sync --status         # Check sync status without syncing
```

### Issue Management Guidelines

- **Use bd for strategic work**: Multi-session tasks, dependencies, discovered work during implementation
- **Use TodoWrite for execution**: Simple single-session task tracking
- **When in doubt, prefer bd**: Persistence beats lost context
- **Creating many issues**: Use parallel subagents for efficiency

## Session Completion Protocol

**CRITICAL**: When ending a work session, you MUST complete ALL steps below. Work is NOT complete until `git push` succeeds.

**MANDATORY WORKFLOW:**

1. **File issues for remaining work** - Create issues for anything that needs follow-up
2. **Run quality gates** (if code changed) - Tests, linters, builds
3. **Update issue status** - Close finished work, update in-progress items
4. **PUSH TO REMOTE** - This is MANDATORY:
   ```bash
   git pull --rebase
   bd sync
   git push
   git status  # MUST show "up to date with origin"
   ```
5. **Clean up** - Clear stashes, prune remote branches
6. **Verify** - All changes committed AND pushed
7. **Hand off** - Provide context for next session

**CRITICAL RULES:**
- Work is NOT complete until `git push` succeeds
- NEVER stop before pushing - that leaves work stranded locally
- NEVER say "ready to push when you are" - YOU must push
- If push fails, resolve and retry until it succeeds

## Development Setup

### Tech Stack

**Backend:**
- Go 1.22+ with Fiber web framework
- PostgreSQL 15+ (database)
- Redis 7+ (cache & sessions)
- MinIO (S3-compatible object storage)

**Frontend:**
- React 18+ with TypeScript
- Vite (build tool)
- TailwindCSS (senior-friendly styling)
- React Query (data fetching)
- Zustand (state management)
- Radix UI (accessible components)

### Initial Setup

```bash
# Run once to set up the project
./scripts/setup.sh
```

This will:
- Create .env files from examples
- Build Docker images
- Start database and cache services
- Run initial database migrations

### Development

```bash
# Start all services (with hot reload)
./scripts/dev.sh

# Or manually with docker-compose
cd docker
docker-compose up
```

**Access URLs:**
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080
- API Health: http://localhost:8080/health
- MinIO Console: http://localhost:9001 (minioadmin/minioadmin)

**Database:**
- PostgreSQL: localhost:5432 (dadmail/dadmail_dev_password)
- Redis: localhost:6379

### Building & Testing

**Backend:**
```bash
cd backend

# Build
go build -o bin/api ./cmd/api

# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Format code
go fmt ./...

# Lint (requires golangci-lint)
golangci-lint run
```

**Frontend:**
```bash
cd frontend

# Install dependencies
npm install

# Development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview

# Lint
npm run lint

# Format (requires prettier)
npm run format
```

### Docker Commands

```bash
cd docker

# View logs
docker-compose logs -f

# View specific service logs
docker-compose logs -f backend
docker-compose logs -f frontend

# Restart a service
docker-compose restart backend

# Stop all services
docker-compose down

# Stop and remove volumes (reset database)
docker-compose down -v

# Rebuild specific service
docker-compose build backend
```

### Database Migrations

Migrations are in `backend/migrations/` and run automatically when PostgreSQL container starts.

To run migrations manually:
```bash
cd docker
docker-compose exec postgres psql -U dadmail -d dadmail -f /docker-entrypoint-initdb.d/001_initial.sql
```

## Architecture

### System Overview

DadMail uses a web-based architecture with Go backend and React frontend:

```
Frontend (React) ←→ REST API + WebSocket ←→ Backend (Go)
                                                    ↓
                                    ┌───────────────┴────────────┐
                                    ↓               ↓            ↓
                              PostgreSQL       Redis        MinIO
```

### Backend Structure

```
backend/
├── cmd/
│   ├── api/          # Main API server entry point
│   └── worker/       # Background email sync worker (future)
├── internal/
│   ├── api/          # HTTP handlers and routes
│   ├── auth/         # Authentication & authorization
│   ├── email/        # Email sync & protocol handling
│   │   ├── gmail/    # Gmail API client
│   │   ├── imap/     # IMAP client
│   │   └── smtp/     # SMTP client
│   ├── categorizer/  # Email categorization engine
│   ├── profile/      # Conversation profiling
│   ├── caregiver/    # Caregiver features
│   ├── models/       # Database models
│   ├── repository/   # Data access layer
│   └── config/       # Configuration management
├── migrations/       # Database migrations
└── pkg/              # Public packages
```

**Key Files:**
- `backend/cmd/api/main.go` - API server entry point
- `backend/internal/config/config.go` - Configuration loader
- `backend/internal/api/router.go` - Route definitions
- `backend/migrations/001_initial.sql` - Initial database schema

### Frontend Structure

```
frontend/src/
├── components/
│   ├── senior/       # Senior-optimized UI components
│   ├── caregiver/    # Caregiver dashboard components
│   └── shared/       # Shared components
├── features/
│   ├── inbox/        # Inbox feature
│   ├── compose/      # Email composition
│   ├── categories/   # Category management
│   └── settings/     # User settings
├── hooks/            # Custom React hooks
├── services/         # API clients
├── stores/           # Zustand state stores
├── styles/           # Global styles
├── types/            # TypeScript type definitions
└── utils/            # Utility functions
```

**Key Files:**
- `frontend/src/App.tsx` - Root component
- `frontend/src/services/api.ts` - API client with auth interceptors
- `frontend/src/config/env.ts` - Environment configuration
- `frontend/tailwind.config.js` - Senior-friendly theme configuration
- `frontend/src/index.css` - Senior-friendly base styles

### Database Schema

Main tables:
- `users` - User accounts (seniors, caregivers, admins)
- `email_accounts` - External email accounts (Gmail, IMAP)
- `categories` - Email categories (medical, financial, family, etc.)
- `categorization_rules` - User/caregiver defined rules
- `emails` - Email metadata (not full content)
- `conversations` - Conversation profiling data
- `caregiver_access` - Caregiver permissions
- `activity_log` - Action audit trail
- `sessions` - JWT refresh tokens

See `backend/migrations/001_initial.sql` for complete schema.

### API Endpoints

**Auth:**
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Refresh access token
- `POST /api/v1/auth/logout` - Logout

**Users:**
- `GET /api/v1/users/me` - Get current user
- `PATCH /api/v1/users/me` - Update user profile

**Emails:**
- `GET /api/v1/emails` - List emails
- `GET /api/v1/emails/:id` - Get email details
- `POST /api/v1/emails` - Send email
- `GET /api/v1/emails/categories/:category` - Filter by category

**Caregivers:**
- `GET /api/v1/caregivers/dashboard` - Caregiver overview
- `POST /api/v1/caregivers/rules` - Create categorization rule
- `GET /api/v1/caregivers/activity` - View activity log

### Senior-Friendly UI Design

**Key Principles:**
- Large, clear typography (18px base font, up to 48px for headings)
- High contrast colors (WCAG AAA compliant)
- Simple navigation (max 3 levels deep)
- Large touch targets (minimum 44px, 48px on mobile)
- Clear focus indicators (4px ring)
- Obvious action buttons with descriptive labels

**TailwindCSS Senior Theme:**
- Custom color palette for categories (medical, financial, family, etc.)
- Larger spacing scale
- Reduced border radius for clarity
- Senior-friendly scrollbars (16px wide)
- Custom utility classes (.btn-primary, .card, .email-item, etc.)

See `frontend/src/index.css` and `frontend/tailwind.config.js` for implementation.

### Next Implementation Phases

1. **Phase 2: Email Integration** - IMAP/SMTP and Gmail API clients
2. **Phase 3: Categorization Engine** - Rule-based and ML categorization
3. **Phase 4: Core UI Features** - Inbox, reading, and compose interfaces
4. **Phase 5: Caregiver Features** - Access control and dashboard

Track progress: `bd list --status=open`
