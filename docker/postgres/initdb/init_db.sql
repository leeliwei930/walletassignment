-- Create development user and database
CREATE USER wallet_app_db_user WITH PASSWORD 'wallet_app_db_user';
CREATE DATABASE wallet_app_db_dev WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';
ALTER DATABASE wallet_app_db_dev OWNER TO wallet_app_db_user;
ALTER DATABASE wallet_app_db_dev SET default_text_search_config = 'pg_catalog.english';

-- Connect to the development database to set permissions
\c wallet_app_db_dev
GRANT ALL PRIVILEGES ON SCHEMA public TO wallet_app_db_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO wallet_app_db_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON SEQUENCES TO wallet_app_db_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON FUNCTIONS TO wallet_app_db_user;

-- Create test user and database
CREATE USER wallet_app_db_test WITH PASSWORD 'wallet_app_db_test';
CREATE DATABASE wallet_app_db_test WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';
ALTER DATABASE wallet_app_db_test OWNER TO wallet_app_db_test;
ALTER DATABASE wallet_app_db_test SET default_text_search_config = 'pg_catalog.english';

-- Connect to the test database to set permissions
\c wallet_app_db_test
GRANT ALL PRIVILEGES ON SCHEMA public TO wallet_app_db_test;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO wallet_app_db_test;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON SEQUENCES TO wallet_app_db_test;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON FUNCTIONS TO wallet_app_db_test;
