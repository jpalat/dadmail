# DadMail

> Email designed for seniors

## The Problem

Seniors can feel overwhelmed by email. From spammers and scammers to managing the scheduling of medical appointments, it's hard to know:
- What is important?
- How do I keep up with the deluge of email?
- Which messages require immediate attention?

## The Solution

DadMail is an email client specifically designed for seniors that:

**For Seniors:**
- **Lower cognitive overhead** - Simplified interface with large, clear typography
- **Clear labeling** - Goes beyond spam/not-spam with meaningful categories (Medical, Financial, Family, etc.)
- **Conversation profiles** - Builds profiles beyond simple threading to understand relationships and context
- **Works with existing email** - Compatible with Gmail, IMAP/SMTP providers

**For Caregivers:**
- **Co-pilot features** - Visibility and access to help manage senior's email
- **Rule creation** - Create categorization rules on behalf of seniors
- **Appointment tracking** - Automatically detect and track appointments
- **Privacy-conscious** - Summarize conversations without full access when appropriate

## Tech Stack

- **Backend**: Go 1.22+ with Fiber framework
- **Frontend**: React 18+ with TypeScript, Vite, and TailwindCSS
- **Database**: PostgreSQL 15+
- **Cache**: Redis 7+
- **Storage**: MinIO (S3-compatible)
- **Development**: Docker Compose with hot reload

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Git
- (Optional) Go 1.22+ for local backend development
- (Optional) Node.js 20+ for local frontend development

### Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/dadmail.git
   cd dadmail
   ```

2. **Run initial setup**
   ```bash
   ./scripts/setup.sh
   ```

3. **Start development environment**
   ```bash
   ./scripts/dev.sh
   ```

4. **Access the application**
   - Frontend: http://localhost:5173
   - Backend API: http://localhost:8080
   - API Health: http://localhost:8080/health
   - MinIO Console: http://localhost:9001 (minioadmin/minioadmin)

## Project Structure

```
dadmail/
├── backend/          # Go backend (Fiber + PostgreSQL)
│   ├── cmd/          # Entry points (api, worker)
│   ├── internal/     # Internal packages
│   └── migrations/   # Database migrations
├── frontend/         # React frontend (TypeScript + TailwindCSS)
│   └── src/          # Source code
├── docker/           # Docker configuration
│   ├── docker-compose.yml
│   ├── backend.Dockerfile
│   └── frontend.Dockerfile
├── scripts/          # Development scripts
└── docs/             # Documentation
```

## Core Features

### Email Categorization

Intelligently categorize emails beyond simple spam filtering:
- **Medical** - Appointments, prescriptions, test results
- **Financial** - Bills, statements, important notices
- **Family & Friends** - Personal correspondence
- **Commercial** - Purchases, shipping, newsletters
- **Administrative** - Government, utilities, services
- **Spam** - Unwanted or promotional emails

### Conversation Profiling

Build richer context beyond traditional threading:
- Track participants and relationships
- Extract topics and keywords
- Detect action items (appointments, deadlines)
- Assign importance scores
- Quick summaries

### Senior-Friendly Design

- Large, clear typography (18px base minimum)
- High contrast colors (WCAG AAA compliant)
- Simple navigation (maximum 3 levels)
- Large touch targets (44px minimum, 48px on mobile)
- Clear focus indicators
- Obvious action buttons

### Caregiver Co-Pilot

- View-only or managed access levels
- Create and manage categorization rules
- Activity logging and monitoring
- Calendar integration
- Privacy controls

## Development

See [CLAUDE.md](CLAUDE.md) for detailed development documentation including:
- Build commands
- Testing
- Database migrations
- Architecture details

For AI agents: See [AGENTS.md](AGENTS.md) for session completion protocol and beads workflow.

## Issue Tracking

This project uses [Beads](https://github.com/steveyegge/beads) for issue tracking.

```bash
# Get started with beads
bd onboard

# View open issues
bd list --status=open

# View ready-to-work issues
bd ready

# Create new issue
bd create --title="..." --type=feature --priority=2

# Track progress
bd stats
```

## Roadmap

- [x] **Phase 1: Foundation** - Project setup, backend/frontend scaffolding, Docker environment
- [ ] **Phase 2: Email Integration** - IMAP/SMTP and Gmail API clients
- [ ] **Phase 3: Categorization Engine** - Rule-based and ML categorization
- [ ] **Phase 4: Core UI Features** - Inbox, reading, and compose interfaces
- [ ] **Phase 5: Caregiver Features** - Access control and dashboard

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

[MIT License](LICENSE)

## Acknowledgments

Built with care for seniors and their caregivers.

