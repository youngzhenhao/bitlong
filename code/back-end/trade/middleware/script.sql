create table bill_balance
(
    id           bigint unsigned auto_increment
        primary key,
    account_id   bigint unsigned null,
    bill_type    smallint        null,
    away         bigint unsigned null,
    amount       decimal(10, 2)  null,
    unit         smallint        not null comment '0:satoshi',
    invoice      varchar(512)    null comment '关联发票',
    payment_hash varchar(100)    null comment '关联支付哈希',
    created_at   datetime(3)     null,
    updated_at   datetime(3)     null,
    deleted_at   datetime(3)     null,
    state        smallint        null comment '0.unkown1.confirmed2.failure'
)
    comment 'User balance change' collate = utf32_unicode_520_ci;

create index uni_bill_balance_account_id
    on bill_balance (account_id);

create table bill_balance_ext
(
    id            bigint unsigned auto_increment comment 'id'
        primary key,
    balance_id    bigint unsigned null,
    bill_ext_desc longtext        null,
    created_at    datetime(3)     null,
    updated_at    datetime(3)     null,
    deleted_at    datetime(3)     null
)
    comment '单据信息扩充表';

create index idx_bill_balance_ext_deleted_at
    on bill_balance_ext (deleted_at);

create table scheduled_tasks
(
    id              int auto_increment
        primary key,
    name            varchar(255)       not null,
    cron_expression varchar(100)       not null,
    function_name   varchar(100)       not null,
    package         varchar(255)       not null,
    status          smallint default 1 null,
    created_at      datetime           null,
    updated_at      datetime           null,
    deleted_at      datetime           null
);

create table user
(
    id         bigint unsigned auto_increment comment 'id'
        primary key,
    user_name  varchar(191) null,
    password   longtext     null,
    status     smallint     null,
    created_at datetime(3)  null,
    update_at  datetime     null comment '更新时间',
    deleted_at datetime(3)  null,
    updated_at datetime(3)  null,
    constraint uni_user_user_name
        unique (user_name)
)
    comment '登陆用户信息';

create index idx_deleted_at
    on user (deleted_at);

create index idx_status
    on user (status);

create index idx_user_deleted_at
    on user (deleted_at);

create index idx_user_name
    on user (user_name);

create table user_account
(
    id                bigint unsigned auto_increment comment 'id'
        primary key,
    user_id           bigint unsigned not null comment '登陆用户id',
    user_name         varchar(100)    not null comment '登陆账号',
    user_account_code varchar(100)    not null comment '账户账户编码',
    created_at        datetime        null comment '创建时间',
    updated_at        datetime        null comment '更新时间',
    deleted_at        datetime        null,
    status            smallint        not null comment '0.无效1.有效',
    label             varchar(100)    null
)
    comment '用户账户信息表';

create table user_invoice
(
    id          bigint unsigned auto_increment comment 'id'
        primary key,
    user_id     bigint unsigned not null,
    account_id  bigint unsigned null,
    invoice     longtext        null,
    amount      decimal(10, 2)  null,
    create_date datetime(3)     null,
    expiry      bigint          null,
    status      smallint        null,
    created_at  datetime(3)     null,
    updated_at  datetime(3)     null,
    deleted_at  datetime(3)     null
);

create index idx_user_invoice_deleted_at
    on user_invoice (deleted_at);


