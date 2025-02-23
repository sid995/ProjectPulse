-- Create extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create enum types
DO $$ 
BEGIN 
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_status') THEN
        CREATE TYPE task_status AS ENUM ('TODO', 'IN_PROGRESS', 'REVIEW', 'DONE');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_priority') THEN
        CREATE TYPE task_priority AS ENUM ('LOW', 'MEDIUM', 'HIGH', 'URGENT');
    END IF;
END $$;

-- Drop existing constraints if they exist
DO $$ 
BEGIN
    -- Drop team constraints
    IF EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'fk_product_owner') THEN
        ALTER TABLE teams DROP CONSTRAINT fk_product_owner;
    END IF;
    IF EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'fk_project_manager') THEN
        ALTER TABLE teams DROP CONSTRAINT fk_project_manager;
    END IF;
    
    -- Drop user constraints
    IF EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'fk_team') THEN
        ALTER TABLE users DROP CONSTRAINT fk_team;
    END IF;
END $$;

-- Create base tables first (no foreign keys)
CREATE TABLE IF NOT EXISTS teams (
    id SERIAL PRIMARY KEY,
    team_name VARCHAR(255) NOT NULL,
    product_owner_user_id INTEGER,
    project_manager_user_id INTEGER
);

CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    profile_picture_url TEXT,
    team_id INTEGER NULL
);

-- Add foreign key constraints in correct order
ALTER TABLE users
    ADD CONSTRAINT fk_team FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE SET NULL;

ALTER TABLE teams
    ADD CONSTRAINT fk_product_owner FOREIGN KEY (product_owner_user_id) REFERENCES users(user_id) ON DELETE SET NULL,
    ADD CONSTRAINT fk_project_manager FOREIGN KEY (project_manager_user_id) REFERENCES users(user_id) ON DELETE SET NULL;

CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    start_date TIMESTAMP WITH TIME ZONE,
    end_date TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS project_teams (
    id SERIAL PRIMARY KEY,
    team_id INTEGER NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    UNIQUE(team_id, project_id)
);

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status task_status,
    priority task_priority,
    tags TEXT,
    start_date TIMESTAMP WITH TIME ZONE,
    due_date TIMESTAMP WITH TIME ZONE,
    points INTEGER,
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    author_user_id INTEGER NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    assigned_user_id INTEGER REFERENCES users(user_id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS task_assignments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    task_id INTEGER NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    UNIQUE(user_id, task_id)
);

CREATE TABLE IF NOT EXISTS attachments (
    id SERIAL PRIMARY KEY,
    file_url TEXT NOT NULL,
    file_name TEXT,
    task_id INTEGER NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    uploaded_by_id INTEGER NOT NULL REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    task_id INTEGER NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(user_id) ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_tasks_project_id ON tasks(project_id);
CREATE INDEX IF NOT EXISTS idx_tasks_author_user_id ON tasks(author_user_id);
CREATE INDEX IF NOT EXISTS idx_tasks_assigned_user_id ON tasks(assigned_user_id);
CREATE INDEX IF NOT EXISTS idx_comments_task_id ON comments(task_id);
CREATE INDEX IF NOT EXISTS idx_comments_user_id ON comments(user_id);
CREATE INDEX IF NOT EXISTS idx_attachments_task_id ON attachments(task_id);
CREATE INDEX IF NOT EXISTS idx_attachments_uploaded_by_id ON attachments(uploaded_by_id);
CREATE INDEX IF NOT EXISTS idx_project_teams_team_id ON project_teams(team_id);
CREATE INDEX IF NOT EXISTS idx_project_teams_project_id ON project_teams(project_id);
CREATE INDEX IF NOT EXISTS idx_users_team_id ON users(team_id);
CREATE INDEX IF NOT EXISTS idx_teams_product_owner ON teams(product_owner_user_id);
CREATE INDEX IF NOT EXISTS idx_teams_project_manager ON teams(project_manager_user_id);
