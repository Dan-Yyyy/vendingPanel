-- Роли пользователей
CREATE TABLE roles
(
    id   serial       not null primary key,
    name varchar(255) not null
);

-- Пользователи
CREATE TABLE users
(
    id            serial                                         not null primary key,
    name          varchar(255)                                   not null,
    email         varchar(255)                                   not null,
    password_hash varchar(255)                                   not null,
    role_id       int constraint fk_roles references roles (id) on delete cascade not null
);

-- Типы доп трат
CREATE TABLE additional_expense_types
(
    id   serial       not null primary key,
    name varchar(255) not null
);

-- Доп траты
CREATE TABLE additional_expenses
(
    id                         serial                                                         not null primary key,
    additional_expense_type_id int constraint fk_additional_expense_types references additional_expense_types (id) on delete cascade not null,
    name                       varchar(255)                                                   not null,
    amount                     int                                                            not null
);

CREATE TABLE consumables
(
    id   serial       not null primary key,
    name varchar(255) not null
);

CREATE TABLE coffee_machine_types
(
    id   serial       not null primary key,
    name varchar(255) not null
);

CREATE TABLE report_types
(
    id   serial       not null primary key,
    name varchar(255) not null
);

-- Торговые точки
CREATE TABLE points
(
    id      serial       not null primary key,
    name    varchar(255) not null,
    address varchar(255) not null
);

-- Связь торговых точек и кофейника
CREATE TABLE coffee_machine_points
(
    coffee_machine_type_id int constraint fk_coffee_machine_types references coffee_machine_types (id) on delete cascade not null,
    point_id               int constraint fk_points references points (id) on delete cascade               not null
);

-- Остатки
CREATE TABLE stocks
(
    id            serial                                            not null primary key,
    consumable_id int constraint fk_consumables references consumables (id) on delete cascade not null,
    amount        int                                               not null,
    created_at    timestamp                                         not null default CURRENT_TIMESTAMP,
    updated_at    timestamp                                         not null default CURRENT_TIMESTAMP
);

-- План объезда
CREATE TABLE visitation_plans
(
    id              serial                                      not null primary key,
    responsible_id  int constraint fk_users references users (id) on delete cascade not null,
    visitation_date date                                        not null
);

-- Пункт плана объезда
CREATE TABLE visitations
(
    id                 serial                                                 not null primary key,
    point_id           int constraint fk_points references points (id) on delete cascade           not null,
    visitation_plan_id int constraint fk_visitation_plans references visitation_plans (id) on delete cascade not null,
    is_visited         boolean                                                not null default false
);

-- Статистика по визиту пункта плана объезда
CREATE TABLE visitation_results
(
    id           serial    not null primary key,
    created_at   timestamp not null default CURRENT_TIMESTAMP,
    updated_at   timestamp not null default CURRENT_TIMESTAMP,
    cash_amount  int       not null,
    coins_amount int       not null,
    cups_amount  int       not null,
    comment      text
);

-- Отчет обслуживания
CREATE TABLE reports
(
    id                   serial                                                   not null primary key,
    user_id              int constraint fk_users references users (id) on delete cascade              not null,
    report_type_id       int constraint fk_report_types references report_types (id) on delete cascade      not null,
    visitation_id        int constraint fk_visitations references visitations (id) on delete cascade        not null,
    created_at           timestamp                                                not null default CURRENT_TIMESTAMP,
    updated_at           timestamp                                                not null default CURRENT_TIMESTAMP,
    is_product_added     boolean                                                  not null default false,
    is_encasement        boolean                                                  not null default false,
    visitation_result_id int constraint fk_visitation_results references visitation_results (id) on delete cascade not null
);

-- Закупки
CREATE TABLE purchases
(
    id             serial                                             not null primary key,
    user_id        int constraint fk_users references users (id) on delete cascade              not null,
    consumable_id  int constraint fk_consumables references consumables (id) on delete cascade  not null,
    amount         int                                                not null,
    price          int                                                not null,
    created_at     timestamp                                          not null default CURRENT_TIMESTAMP,
    updated_at     timestamp                                          not null default CURRENT_TIMESTAMP
);

-- Что было израсходовано со склада при обслуживании
CREATE TABLE report_stocks
(
    report_id int constraint fk_reports references reports (id) on delete cascade  not null,
    stock_id  int constraint fk_stocks references stocks (id)   on delete cascade  not null,
    amount    int                                                                  not null
);