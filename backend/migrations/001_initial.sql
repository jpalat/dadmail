-- Initial database schema for DadMail

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'senior', -- senior, caregiver, admin
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_login_at TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);

-- Email accounts table (stores user's external email accounts)
CREATE TABLE email_accounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    provider VARCHAR(50) NOT NULL, -- gmail, imap, outlook
    email_address VARCHAR(255) NOT NULL,
    display_name VARCHAR(255),
    credentials_encrypted TEXT NOT NULL, -- AES-256 encrypted credentials
    is_primary BOOLEAN DEFAULT false,
    sync_enabled BOOLEAN DEFAULT true,
    last_synced_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, email_address)
);

CREATE INDEX idx_email_accounts_user ON email_accounts(user_id);
CREATE INDEX idx_email_accounts_sync ON email_accounts(sync_enabled, last_synced_at);

-- Categories table
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    color VARCHAR(7), -- hex color code
    icon VARCHAR(50),
    is_system BOOLEAN DEFAULT false, -- system categories can't be deleted
    display_order INT DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Insert default system categories
INSERT INTO categories (name, description, color, icon, is_system, display_order) VALUES
    ('medical', 'Medical appointments, prescriptions, test results', '#dc2626', 'heart-pulse', true, 1),
    ('financial', 'Bills, statements, important notices', '#16a34a', 'dollar-sign', true, 2),
    ('family', 'Family and friends correspondence', '#2563eb', 'users', true, 3),
    ('commercial', 'Purchases, shipping, newsletters', '#9333ea', 'shopping-cart', true, 4),
    ('administrative', 'Government, utilities, services', '#ea580c', 'file-text', true, 5),
    ('spam', 'Unwanted or promotional emails', '#6b7280', 'trash-2', true, 6);

-- User preferences table
CREATE TABLE user_preferences (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    theme VARCHAR(20) DEFAULT 'light', -- light, dark, high-contrast
    font_size VARCHAR(20) DEFAULT 'large', -- small, medium, large, extra-large
    auto_categorize BOOLEAN DEFAULT true,
    show_images BOOLEAN DEFAULT false, -- for security, default off
    notifications_enabled BOOLEAN DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Categorization rules table
CREATE TABLE categorization_rules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE, -- null means system rule
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    priority INT DEFAULT 0, -- higher number = higher priority
    enabled BOOLEAN DEFAULT true,

    -- Rule conditions (JSON format for flexibility)
    conditions JSONB NOT NULL, -- e.g., {"from": ["*@hospital.com"], "subject": ["appointment"]}

    created_by UUID REFERENCES users(id), -- for caregiver-created rules
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_rules_user ON categorization_rules(user_id);
CREATE INDEX idx_rules_priority ON categorization_rules(priority DESC);
CREATE INDEX idx_rules_enabled ON categorization_rules(enabled);

-- Emails metadata table (we don't store full email body, just metadata)
CREATE TABLE emails (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_id UUID NOT NULL REFERENCES email_accounts(id) ON DELETE CASCADE,
    external_id VARCHAR(255) NOT NULL, -- provider's message ID
    thread_id VARCHAR(255), -- for conversation threading

    from_address VARCHAR(255) NOT NULL,
    from_name VARCHAR(255),
    to_addresses TEXT[], -- array of email addresses
    cc_addresses TEXT[],
    subject TEXT,
    snippet TEXT, -- short preview

    category_id UUID REFERENCES categories(id),
    is_read BOOLEAN DEFAULT false,
    is_starred BOOLEAN DEFAULT false,
    has_attachments BOOLEAN DEFAULT false,

    received_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE(account_id, external_id)
);

CREATE INDEX idx_emails_account ON emails(account_id);
CREATE INDEX idx_emails_category ON emails(category_id);
CREATE INDEX idx_emails_thread ON emails(thread_id);
CREATE INDEX idx_emails_received ON emails(received_at DESC);
CREATE INDEX idx_emails_unread ON emails(account_id, is_read) WHERE is_read = false;

-- Conversation profiles table
CREATE TABLE conversations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_id UUID NOT NULL REFERENCES email_accounts(id) ON DELETE CASCADE,
    thread_id VARCHAR(255) NOT NULL,

    participants TEXT[], -- array of email addresses involved
    subject TEXT,
    message_count INT DEFAULT 0,
    last_message_at TIMESTAMP,

    -- Extracted metadata
    topics TEXT[], -- extracted topics/keywords
    importance_score INT DEFAULT 0, -- 0-100
    has_action_items BOOLEAN DEFAULT false,
    action_items JSONB, -- detected appointments, deadlines

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE(account_id, thread_id)
);

CREATE INDEX idx_conversations_account ON conversations(account_id);
CREATE INDEX idx_conversations_last_message ON conversations(last_message_at DESC);
CREATE INDEX idx_conversations_importance ON conversations(importance_score DESC);

-- Caregiver access table
CREATE TABLE caregiver_access (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    senior_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    caregiver_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    access_level VARCHAR(50) NOT NULL DEFAULT 'view', -- view, manage, full
    status VARCHAR(50) NOT NULL DEFAULT 'pending', -- pending, active, revoked

    can_view_emails BOOLEAN DEFAULT true,
    can_create_rules BOOLEAN DEFAULT false,
    can_manage_categories BOOLEAN DEFAULT false,

    invited_at TIMESTAMP NOT NULL DEFAULT NOW(),
    accepted_at TIMESTAMP,
    revoked_at TIMESTAMP,

    UNIQUE(senior_id, caregiver_id)
);

CREATE INDEX idx_caregiver_access_senior ON caregiver_access(senior_id);
CREATE INDEX idx_caregiver_access_caregiver ON caregiver_access(caregiver_id);
CREATE INDEX idx_caregiver_access_status ON caregiver_access(status);

-- Activity log table (for caregiver monitoring)
CREATE TABLE activity_log (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    actor_id UUID REFERENCES users(id), -- who performed the action (caregiver)
    action_type VARCHAR(100) NOT NULL, -- email_read, email_sent, rule_created, etc.
    resource_type VARCHAR(100), -- email, rule, category, etc.
    resource_id UUID,
    details JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_activity_user ON activity_log(user_id, created_at DESC);
CREATE INDEX idx_activity_actor ON activity_log(actor_id, created_at DESC);
CREATE INDEX idx_activity_type ON activity_log(action_type);

-- Sessions table (for JWT refresh tokens)
CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    refresh_token VARCHAR(255) NOT NULL UNIQUE,
    user_agent TEXT,
    ip_address VARCHAR(45),
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_sessions_user ON sessions(user_id);
CREATE INDEX idx_sessions_token ON sessions(refresh_token);
CREATE INDEX idx_sessions_expires ON sessions(expires_at);

-- Function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Apply update trigger to relevant tables
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_email_accounts_updated_at BEFORE UPDATE ON email_accounts
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_user_preferences_updated_at BEFORE UPDATE ON user_preferences
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_categorization_rules_updated_at BEFORE UPDATE ON categorization_rules
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_emails_updated_at BEFORE UPDATE ON emails
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_conversations_updated_at BEFORE UPDATE ON conversations
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
