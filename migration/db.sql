create table if not exists conf_flow_atoms
(
    id            int unsigned auto_increment
    primary key,
    flow_iid      int unsigned                           not null comment 'flow的version id',
    name          varchar(255)                           not null,
    type          int unsigned                           null comment 'step类型，为了灵活封装预留。如远程原子，自定义原子等',
    resource_id   int unsigned                           not null comment '资源id',
    resource_conf json                                   not null comment '资源配置',
    before_run    int                                    null comment 'step运行前跑的内容',
    after_run     json                                   not null comment '运行的内容',
    image         varchar(255) default ''                not null comment '镜像地址',
    run           json                                   not null,
    modifier      varchar(255)                           not null comment '记录个性人',
    created_at    timestamp    default CURRENT_TIMESTAMP not null,
    updated_time  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
    );

create table if not exists conf_flow_dags
(
    id           int unsigned auto_increment
    primary key,
    flow_iid     int unsigned                           null comment 'flow的version id',
    stage_id     int unsigned                           not null,
    step_id      int unsigned                           not null,
    from_step_id int unsigned                           not null,
    next_step_id int unsigned default '0'               not null,
    created_at   timestamp    default CURRENT_TIMESTAMP not null,
    updated_at   timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

create table if not exists conf_flow_ids
(
    id         int unsigned auto_increment
    primary key,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

create table if not exists conf_flow_stages
(
    id         int unsigned auto_increment
    primary key,
    flow_iid   int unsigned                        not null comment 'flow的version id',
    name       varchar(255)                        not null,
    order_num  smallint unsigned                   not null comment '顺序',
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
    );

create table if not exists conf_flows
(
    id                int unsigned auto_increment
    primary key,
    flow_id           int unsigned                               not null comment '流水线id',
    name              varchar(255)                               not null comment '流水线名称',
    status            tinyint          default 0                 not null comment '流水线状态。0:默认状态，1:正常，2:废弃',
    version           int unsigned     default '0'               not null,
    latest_version    int unsigned     default '0'               not null,
    parallel_num      int unsigned     default '1'               not null comment '并行数',
    parallel_strategy tinyint unsigned default '1'               not null comment '并行策略。1:禁止新触发 2:取消最早的',
    resource_id       int unsigned     default '0'               not null comment '默认使用的资源，优先级低于step的资源配置',
    resource_conf     json                                       null comment '资源用量配置',
    before_run        json                                       null comment '在流水线调度前要执行的内容。比如当前这次触发是否要执行，是否要通知等',
    after_run         json                                       null comment '在流水线终态后要执行的内容。比如通知等',
    modifier          int unsigned                               not null,
    created_at        timestamp        default CURRENT_TIMESTAMP not null,
    updated_at        timestamp        default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
    );

create table if not exists conf_resource
(
    id         int unsigned auto_increment
    primary key,
    name       varchar(255)                               not null comment '资源名称',
    type       tinyint unsigned default '1'               not null comment '资源类型。1:k8s，2:普通机器',
    info       json                                       null comment '资源详情',
    created_at timestamp        default CURRENT_TIMESTAMP not null,
    updated_at timestamp        default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
    )
    comment '执行资源';

