DROP TABLE IF EXISTS asset_genesis;
DROP TABLE IF EXISTS assets;
DROP TABLE IF EXISTS asset_proofs;
DROP TABLE IF EXISTS asset_minting_state;
DROP TABLE IF EXISTS chain_txns;
DROP TABLE IF EXISTS genesis_points;
DROP TABLE IF EXISTS genesis_assets;
DROP INDEX IF EXISTS asset_ids;
DROP TABLE IF EXISTS internal_keys;
DROP TABLE IF EXISTS asset_groups;
DROP TABLE IF EXISTS asset_group_witnesses;
DROP TABLE IF EXISTS managed_utxos;
DROP TABLE IF EXISTS script_keys;
DROP TABLE IF EXISTS asset_witnesses;
DROP TABLE IF EXISTS asset_seedlings;
DROP TABLE IF EXISTS asset_minting_batches;
DROP INDEX IF EXISTS batch_state_lookup;
DROP VIEW IF EXISTS genesis_info_view;
DROP VIEW IF EXISTS key_group_info_view;
DROP TABLE IF EXISTS assets_meta;
DROP INDEX IF EXISTS meta_hash_index;