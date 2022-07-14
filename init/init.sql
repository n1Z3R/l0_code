CREATE TABLE orders (
                      id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                      content jsonb NOT NULL
);