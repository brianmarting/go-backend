DROP TABLE IF EXISTS crypto;
DROP TABLE IF EXISTS wallet;
DROP TABLE IF EXISTS wallet_crypto;

CREATE TABLE crypto (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL
);

CREATE TABLE wallet (
    id SERIAL PRIMARY KEY,
    address VARCHAR NOT NULL
)

CREATE TABLE wallet_crypto (
    id SERIAL PRIMARY KEY,
    wallet_id SERIAL NOT NULL,
    crypto_id SERIAL NOT NULL,
    amount SERIAL NOT NULL
)