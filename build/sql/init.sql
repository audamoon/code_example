create schema if not exists billing;


comment on schema billing is 'Схема для системы взаимодействия с кошельками';


create schema if not exists core_pkg;


comment on schema core_pkg is 'Схема для функций и утилит';


create or replace function core_pkg.update_updated_at() returns trigger
    language plpgsql
as $$
begin
    new.updated_at = now();
    return new;
end;
$$;


create table if not exists billing.users
(
    id serial primary key,
    account_name varchar,
    first_name varchar,
    last_name varchar,
    created_at timestamp not null default now(),
    updated_at timestamp
);


create unique index uq_idx1_users on billing.users(account_name);


create trigger tbu_billing_users before update
    on billing.users
for each row execute procedure core_pkg.update_updated_at();


create table if not exists billing.wallets
(
    id serial primary key,
    balance integer not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);


create trigger tbu_billing_wallets before update
    on billing.wallets
for each row execute procedure core_pkg.update_updated_at();


create table if not exists billing.users_wallets
(
    wallet_id integer not null,
    user_id integer not null,
    created_at timestamp not null default now(),
    updated_at timestamp,

    constraint fk1_users_wallets foreign key (wallet_id)
        references billing.wallets(id),
    constraint fk2_users_wallets foreign key (user_id)
        references billing.users(id)
);


create trigger tbu_billing_users_wallets
    before update
    on billing.users_wallets
    for each row
execute procedure core_pkg.update_updated_at();


create index idx1_users_wallets on billing.users_wallets(wallet_id);

create index idx2_users_wallets on billing.users_wallets(user_id);


insert into billing.wallets(balance)
values (200), (300), (400);

