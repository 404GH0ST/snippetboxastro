# Notes

## Creating a User

```bash
# as postgres user
initdb -D /var/lib/postgres/data
createuser --interactive
createdb snippetbox -O myUserName
ALTER USER myUserName WITH ENCRYPTED PASSWORD 's3cur3P4$$w0rd';
```

## Setup PostgreSQL database:
If haven't already done so
```psql
CREATE DATABASE snippetbox;
```

```psql
CREATE TABLE snippets (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created timestamptz NOT NULL,
    expires timestamptz NOT NULL
);
```

```psql
INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n- Frodo Baggins',
    now(),
    now() + interval '365 days' 
);
```
