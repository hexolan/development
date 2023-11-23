CREATE TABLE passwd_auth_methods (
    user_id varchar(64) PRIMARY KEY,
    hashed_passwd varchar(128) NOT NULL
);