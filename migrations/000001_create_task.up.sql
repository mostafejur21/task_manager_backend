CREATE TABLE IF NOT EXISTS tasks (
    id bigserial PRIMARY KEY,
    title varchar(255) NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
