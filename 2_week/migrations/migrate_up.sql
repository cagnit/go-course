CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE terminals (
    id BIGINT PRIMARY KEY,
    client_id TEXT NOT NULL,
    client_secret TEXT NOT NULL,
    uuid UUID NOT NULL UNIQUE
);

CREATE TABLE transactions (
    id BIGINT PRIMARY KEY,
    terminal_id UUID REFERENCES terminals(uuid) ON DELETE CASCADE,
    order_id TEXT NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status_changed TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    code TEXT,
    message TEXT
);

-- Тестовый терминал с заданными ID и UUID
INSERT INTO terminals (id, client_id, client_secret, uuid)
VALUES (
    1001,
    'demo_client',
    'super_secret',
    '11111111-1111-1111-1111-111111111111'
);

-- Тестовая транзакция с ручным ID
INSERT INTO transactions (id, terminal_id, order_id, amount, status, code, message)
VALUES (
    5001,
    '11111111-1111-1111-1111-111111111111',
    'order-001',
    150.00,
    'NEW',
    '00',
    'Создана вручную'
);
