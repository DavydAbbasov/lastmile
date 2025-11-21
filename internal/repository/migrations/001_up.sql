BEGIN;

-- === ENUM TYPES
CREATE TYPE user_role AS ENUM ('client', 'moderator');

CREATE TYPE city AS ENUM ('new_york', 'los_angeles', 'chicago');

CREATE TYPE intake_status AS ENUM ('in_progress', 'closed');

CREATE TYPE parcel_type AS ENUM ('electronics', 'clothes', 'shoes');

-- === TABLES
CREATE TABLE users (
    id           BIGSERIAL PRIMARY KEY,
    email        TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role         user_role NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE TABLE pickup_points (
    id          BIGSERIAL PRIMARY KEY,
    city        city NOT NULL,
    address     TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE TABLE intakes (
    id               BIGSERIAL PRIMARY KEY,
    pickup_point_id  BIGINT NOT NULL REFERENCES pickup_points(id),
    status           intake_status NOT NULL,
    started_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    closed_at        TIMESTAMPTZ,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE TABLE parcels (
    id           BIGSERIAL PRIMARY KEY,
    intake_id    BIGINT NOT NULL REFERENCES intakes(id) ON DELETE CASCADE,
    type         parcel_type NOT NULL,
    received_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- === INDEXES
CREATE INDEX idx_intakes_pvz_status
    ON intakes (pickup_point_id, status);
CREATE INDEX idx_parcels_intake_received_at
    ON parcels (intake_id, received_at DESC);

COMMIT;