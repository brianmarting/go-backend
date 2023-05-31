DROP TABLE IF EXISTS crypto;
DROP TABLE IF EXISTS wallet;
DROP TABLE IF EXISTS wallet_crypto;

CREATE TABLE crypto (
    id BIGSERIAL PRIMARY KEY,
    uuid VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL
);

CREATE TABLE wallet (
    id BIGSERIAL PRIMARY KEY,
    uuid VARCHAR NOT NULL,
    crypto_id SERIAL NOT NULL,
    address VARCHAR NOT NULL,
    amount BIGINT NOT NULL,

    CONSTRAINT wallet_crypto_id_fk FOREIGN KEY (crypto_id) REFERENCES crypto (id)
);

CREATE TABLE exchange_user (
    id BIGSERIAL PRIMARY KEY,
    uuid VARCHAR NOT NULL,
    name VARCHAR NOT NULL
);

CREATE TABLE user_wallets (
    id BIGSERIAL PRIMARY KEY,
    wallet_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,

    CONSTRAINT user_wallets_wallet_id_fk FOREIGN KEY (wallet_id) REFERENCES wallet (id),
    CONSTRAINT user_wallets_user_id_fk FOREIGN KEY (user_id) REFERENCES user (id)
)

INSERT INTO crypto VALUES (1, '484cb884-fc16-42dc-b562-15720d7eb29f', 'btc', 'Bitcoin');
INSERT INTO crypto VALUES (2, '666cb884-fc16-42dc-b562-15720d7eb29f', 'eth', 'Ethereum');

INSERT INTO wallet VALUES (1, '781273f8-ea1c-4286-93c2-d52dc120a9d7', 1, 'bc1d42UNb54eBiGm0qEM0h6r2h8n532to9jtp186ns', 10);
INSERT INTO wallet VALUES (2, 'f494b4a6-ce50-49b8-95b9-2bc722f11421', 2, 'ff2d42UNb54eBiGm0qEM0h6r2h8n532to9jtp186ns', 25);

INSERT INTO user VALUES (1, 'a8cbb0b4-e44d-4d8d-a424-75f4b3657d23', 'John');

INSERT INTO user_wallets VALUES (1, 1, 1);
INSERT INTO user_wallets VALUES (1, 2, 1)
