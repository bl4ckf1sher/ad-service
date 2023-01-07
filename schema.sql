CREATE TYPE user_role AS ENUM ('admin', 'user');
CREATE TABLE user (
    id UUID PRIMARY KEY,
    role user_role,
    email VARCHAR,
    password VARCHAR,
    name VARCHAR,
    surname VARCHAR,
    created_at TIMESTAMP WITHOUT TIME ZONE,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);