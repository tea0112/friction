CREATE SCHEMA IF NOT EXISTS friction;

CREATE TABLE IF NOT EXISTS friction.users (
    id BIGSERIAL,
    email VARCHAR(100) UNIQUE NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS friction.roles (
    id BIGSERIAL,
    name VARCHAR(100) UNIQUE NOT NULL,
    CONSTRAINT pk_roles PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS friction.permissions (
    id BIGSERIAL,
    name VARCHAR(100) UNIQUE NOT NULL,
    CONSTRAINT pk_permissions PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS friction.users_roles (
    user_id BIGSERIAL,
    role_id BIGSERIAL,
    CONSTRAINT pk_users_roles PRIMARY KEY (user_id, role_id)
);

CREATE TABLE IF NOT EXISTS friction.roles_permissions (
    role_id BIGSERIAL,
    permission_id BIGSERIAL,
    CONSTRAINT pk_roles_permissions PRIMARY KEY (role_id, permission_id)
);

