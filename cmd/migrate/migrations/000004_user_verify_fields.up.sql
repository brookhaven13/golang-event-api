ALTER TABLE users
ADD COLUMN verified boolean NOT NULL DEFAULT false,
ADD COLUMN verify_token text,
ADD COLUMN verify_token_expires timestamp with time zone;