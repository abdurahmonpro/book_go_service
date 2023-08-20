CREATE TABLE IF NOT EXISTS "book" (
    "id" SERIAL PRIMARY KEY,
    "isbn" VARCHAR(30) NOT NULL,
    "title" VARCHAR(100) NOT NULL,
    "cover" TEXT NOT NULL UNIQUE,
    "author" VARCHAR(50) NOT NULL,
    "published" VARCHAR(30) NOT NULL,
    "pages" INTEGER NOT NULL,
    -- 0 new, 1 reading, 2 finished
    "status" SMALLINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
