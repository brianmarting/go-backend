DROP TABLE IF EXISTS crypto;
DROP TABLE IF EXISTS wallet;
DROP TABLE IF EXISTS wallet_crypto;

CREATE TABLE crypto (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL
);

CREATE TABLE wallet (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR NOT NULL,
    address VARCHAR NOT NULL
);

CREATE TABLE wallet_crypto (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR NOT NULL,
    wallet_id SERIAL NOT NULL,
    crypto_id SERIAL NOT NULL,
    amount SERIAL NOT NULL
);

INSERT INTO crypto VALUES (1, '484cb884-fc16-42dc-b562-15720d7eb29f', 'btc', 'Bitcoin');
INSERT INTO crypto VALUES (2, '666cb884-fc16-42dc-b562-15720d7eb29f', 'eth', 'Ethereum');

INSERT INTO wallet VALUES (1, '781273f8-ea1c-4286-93c2-d52dc120a9d7', 'bc1d42UNb54eBiGm0qEM0h6r2h8n532to9jtp186ns');
INSERT INTO wallet VALUES (2, 'f494b4a6-ce50-49b8-95b9-2bc722f11421', 'ff2d42UNb54eBiGm0qEM0h6r2h8n532to9jtp186ns');

INSERT INTO wallet_crypto VALUES (1, 'd72280e4-1be9-4ba9-9f28-ceadbce5b805', 1, 1, 10);
INSERT INTO wallet_crypto VALUES (2, '06dcb1ec-44a2-4813-ab8d-3fe5f99619e3', 2, 2, 15);
