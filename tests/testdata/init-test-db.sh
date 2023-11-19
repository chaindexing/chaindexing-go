#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER docker;
    GRANT ALL PRIVILEGES ON DATABASE chaindexing_test TO docker;
    CREATE TABLE IF NOT EXISTS chaindexing_contract_address (
        id SERIAL PRIMARY KEY,
        address VARCHAR(255) NOT NULL,
        contract_name VARCHAR(255) NOT NULL,
        chain_id INT NOT NULL,
        start_block_number BIGINT NOT NULL,
        next_block_to_ingest_from BIGINT NOT NULL,
        next_block_to_handle_from BIGINT NOT NULL
    );
EOSQL