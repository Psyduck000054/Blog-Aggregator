-- +goose Up
create table feed_follows (
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id uuid references users(id) on delete cascade not null,
    feed_id uuid references feeds(id) on delete cascade not null,
    constraint unique_usr_fd unique (user_id, feed_id)
);
    
-- +goose Down
drop table feed_follows;
