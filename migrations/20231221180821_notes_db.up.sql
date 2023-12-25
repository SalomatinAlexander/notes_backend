CREATE TABLE notes_table(
    id bigserial not null primary key,
    user_id bigserial not null,
    list_id bigserial not null,
    title varchar not null,
    description_text varchar not null,
    create_at varchar not null,
    update_at varchar not null
);