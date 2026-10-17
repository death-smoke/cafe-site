CREATE TABLE IF NOT EXISTS categories (
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE IF NOT EXISTS images (
    id         SERIAL PRIMARY KEY,
    path       TEXT NOT NULL,
    -- Пример attributes:
    -- {"filename":"ab12.png","original_name":"salad.png","mime_type":"image/png",
    --  "size_bytes":123456,"width":800,"height":600,"alt":"Цезарь","uploaded_at":"2026-10-17T12:00:00Z"}
    attributes JSONB NOT NULL DEFAULT '{}'::jsonb
);

CREATE TABLE IF NOT EXISTS menu (
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    category_id INT  NOT NULL REFERENCES categories(id),
    description TEXT,
    ingredients TEXT,
    image_id    INT REFERENCES images(id) ON DELETE SET NULL,
    price       INT  NOT NULL CHECK (price >= 0),
    UNIQUE (name)
);

CREATE INDEX IF NOT EXISTS idx_menu_category ON menu(category_id);
