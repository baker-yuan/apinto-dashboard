CREATE TABLE `api`
(
    `id`                 int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `namespace`          int(11) NOT NULL COMMENT '工作空间',
    `uuid`               varchar(36)  NOT NULL COMMENT 'api的UUID',
    `group_uuid`         varchar(36)  NOT NULL COMMENT 'api所在分组的UUID',
    `name`               varchar(255) NOT NULL COMMENT 'api名称',
    `request_path`       varchar(255) NOT NULL COMMENT 'api请求路径',
    `request_path_label` varchar(255) NOT NULL COMMENT 'api请求路径Label',
    `source_type`        varchar(36)  NOT NULL COMMENT '来源类型',
    `source_id`          int(11) NOT NULL COMMENT '来源id,用于关联外部应用',
    `source_label`       varchar(36)  NOT NULL COMMENT '来源标签',
    `desc`               varchar(255)          DEFAULT NULL COMMENT '描述',
    `operator`           int(11) DEFAULT NULL COMMENT '更新人/操作人',
    `create_time`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `namespace_name` (`namespace`,`uuid`),
    KEY                  `group_uuid` (`group_uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='API表';

CREATE TABLE `audit_log`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `namespace`  int(11) NOT NULL COMMENT '工作空间',
    `user_id`    int(11) NOT NULL COMMENT '用户id',
    `username`   varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
    `ip`         varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'ip地址',
    `operate`    tinyint(2) NOT NULL COMMENT '操作类型 1.创建 2.编辑 3.删除 4.发布',
    `kind`       varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '操作对象',
    `object`     text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '对象信息',
    `url`        text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求url,包括了query参数',
    `start_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '请求开始时间',
    `end_time`   timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '请求结束时间',
    `user_agent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'user_agent',
    `body`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求内容',
    `error`      text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '错误信息',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '操作日志' ROW_FORMAT = Dynamic;

CREATE TABLE `cluster`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '集群ID',
    `namespace`   int(11) NOT NULL COMMENT '工作空间',
    `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '集群名称',
    `name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '集群名称',
    `desc`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
    `addr`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '集群地址',
    `env`         varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL COMMENT '环境',
    `create_time` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `update_time` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP (0) COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `namespace`(`namespace`, `name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '集群信息表' ROW_FORMAT = Dynamic;


CREATE TABLE `cluster_certificate`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '证书ID',
    `cluster`     int(11) NOT NULL COMMENT '集群ID',
    `namespace`   int(11) NOT NULL COMMENT '工作空间',
    `operator`    int(11) NULL DEFAULT NULL COMMENT '更新人/操作人',
    `uuid`        varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `key`         text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'key',
    `pem`         text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'pem',
    `create_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `update_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP (0) COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX         `cluster`(`cluster`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '集群绑定的证书信息表' ROW_FORMAT = Dynamic;


CREATE TABLE `cluster_config`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `namespace`   int(11) NOT NULL COMMENT '工作空间',
    `cluster`     int(11) NOT NULL COMMENT '集群ID',
    `type`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型',
    `is_enable`   tinyint(1) NOT NULL COMMENT '是否启用',
    `data`        text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置数据',
    `operator`    int(11) NULL DEFAULT NULL COMMENT '更新人/操作人',
    `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP (0) COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX         `cluster_type`(`type`, `cluster`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '集群配置表' ROW_FORMAT = DYNAMIC;


CREATE TABLE `cluster_node`
(
    `id`           int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `name`         varchar(255) NOT NULL COMMENT '节点名称/ID',
    `namespace`    int(11) NOT NULL COMMENT '工作空间',
    `cluster`      int(11) NOT NULL COMMENT '集群ID',
    `admin_addr`   text COMMENT '管理地址',
    `service_addr` text         NOT NULL COMMENT '服务地址',
    `create_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `cluster` (`cluster`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=158 DEFAULT CHARSET=utf8mb4 COMMENT='集群绑定的节点信息表';

CREATE TABLE `common_group`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '目录ID',
    `uuid`        varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'uuid',
    `namespace`   int(11) NOT NULL COMMENT '工作空间',
    `tag`         int(11) NOT NULL COMMENT '根据type区分是哪个表的ID',
    `type`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型',
    `name`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
    `parent_id`   int(11) NOT NULL DEFAULT 0 COMMENT '父级目录ID',
    `operator`    int(11) NULL DEFAULT NULL COMMENT '更新人/操作人',
    `sort`        int(11) NOT NULL,
    `create_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `update_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP (0) COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `uuid`(`uuid`) USING BTREE,
    UNIQUE INDEX `name_parentID`(`name`, `parent_id`) USING BTREE,
    INDEX         `type`(`namespace`, `type`, `tag`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '通用目录表' ROW_FORMAT = Dynamic;

CREATE TABLE `history`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `namespace` int(11) NOT NULL COMMENT '工作空间',
    `target`    int(11) NOT NULL COMMENT '根据kind区分',
    `kind`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'kind',
    `old_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'old_value',
    `new_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'new_value',
    `opt_type`  tinyint(4) NULL DEFAULT 0 COMMENT '1新增 2修改 3删除',
    `operator`  int(11) NOT NULL COMMENT '更新人/操作人',
    `opt_time`  timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '操作时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX       `kind_target`(`kind`, `target`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1476 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '变更记录表' ROW_FORMAT = Dynamic;


CREATE TABLE `namespace`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '工作空间name',
    `create_time` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '注册时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '工作空间' ROW_FORMAT = Dynamic;


CREATE TABLE `publish_history`
(
    `id`           int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `version_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '版本号',
    `version`      int(11) NOT NULL COMMENT '版本ID',
    `cluster`      int(11) NOT NULL COMMENT '集群ID',
    `namespace`    int(11) NOT NULL COMMENT '工作空间',
    `target`       int(11) NOT NULL COMMENT '根据kind区分',
    `opt_type`     int(11) NOT NULL COMMENT '1发布 2回滚',
    `desc`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
    `kind`         varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'kind',
    `data`         text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'data',
    `operator`     int(11) NOT NULL COMMENT '操作人',
    `opt_time`     timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '操作时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX          `cluster`(`cluster`) USING BTREE,
    INDEX          `kind_target`(`kind`, `target`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '发布记录表' ROW_FORMAT = DYNAMIC;


CREATE TABLE `quote`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `source`      int(11) NOT NULL COMMENT '根据kind区分是哪个表的ID,被引用的id',
    `kind`        varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型',
    `target`      int(11) NOT NULL COMMENT '引用的id',
    `target_kind` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '引用的类型',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX         `kind`(`kind`) USING BTREE,
    INDEX         `target_kind`(`target_kind`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '引用表' ROW_FORMAT = Dynamic;


CREATE TABLE `runtime`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `cluster`     int(11) NOT NULL COMMENT '集群ID',
    `target`      int(11) NULL DEFAULT 0 COMMENT '根据type区分是哪个表的ID',
    `namespace`   int(11) NOT NULL COMMENT '工作空间',
    `kind`        varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型',
    `disable`     tinyint(1) NOT NULL COMMENT '禁用状态',
    `is_online`   tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否上线',
    `version`     int(11) NOT NULL COMMENT '版本ID',
    `operator`    int(11) NULL DEFAULT NULL COMMENT '更新人/操作人',
    `create_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `update_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP (0) COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `cluster`(`cluster`, `kind`, `target`) USING BTREE,
    INDEX         `kind`(`kind`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '正在运行的版本表' ROW_FORMAT = Dynamic;


CREATE TABLE `stat`
(
    `id`      int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `target`  int(11) NOT NULL COMMENT '根据type区分是哪个表的ID',
    `kind`    varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型',
    `version` int(11) NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `kind`(`kind`, `target`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '最新版本关联表' ROW_FORMAT = Dynamic;


CREATE TABLE `strategy`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `uuid`        varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'uuid',
    `namespace`   int(11) NOT NULL COMMENT '工作空间',
    `cluster`     int(11) NOT NULL COMMENT '集群ID',
    `name`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '策略名称',
    `desc`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
    `type`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型',
    `priority`    int(11) NOT NULL COMMENT '优先级',
    `is_stop`     tinyint(1) NOT NULL COMMENT '是否停用',
    `is_delete`   tinyint(1) NOT NULL COMMENT '是否已删除（软删除）',
    `operator`    int(11) NULL DEFAULT NULL COMMENT '更新人/操作人',
    `create_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `update_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP (0) COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `uuid`(`uuid`) USING BTREE,
    INDEX         `priority_type`(`cluster`, `type`, `priority`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '流量策略表' ROW_FORMAT = DYNAMIC;


CREATE TABLE `system_info`
(
    `id`    int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `key`   varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '健',
    `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '值',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `unique_key`(`key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;



CREATE TABLE `variable_cluster`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `namespace`   int(11) NOT NULL COMMENT '工作空间',
    `variable`    int(11) NOT NULL COMMENT '全局环境变量ID',
    `cluster`     int(11) NOT NULL COMMENT '集群ID',
    `status`      int(11) NOT NULL DEFAULT 0 COMMENT '状态 1 删除',
    `key`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'key',
    `value`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'value',
    `operator`    int(11) NOT NULL COMMENT '更新人/操作人',
    `create_time` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `update_time` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP (0) COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `cluster`(`cluster`, `variable`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '集群绑定的环境变量信息表' ROW_FORMAT = Dynamic;


CREATE TABLE `variables`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `namespace`   int(11) NOT NULL COMMENT '工作空间',
    `key`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'key',
    `desc`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
    `operator`    int(11) NOT NULL COMMENT '更新人/操作人',
    `create_time` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `update_time` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP (0) COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `namespace`(`namespace`, `key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '全局环境变量信息表' ROW_FORMAT = Dynamic;


CREATE TABLE `version`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `namespace`   int(11) NOT NULL COMMENT '工作空间',
    `target`      int(11) NOT NULL COMMENT '根据type区分是哪个表的ID',
    `kind`        varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型',
    `data`        text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置信息JSON',
    `operator`    int(11) NULL DEFAULT NULL COMMENT '更新人/操作人',
    `create_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `update_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP (0) COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX         `kind`(`kind`, `target`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '版本表' ROW_FORMAT = Dynamic;



CREATE TABLE `external_app`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `uuid`        varchar(36)                                                  NOT NULL COMMENT '应用id，随机生成的长度32的字符串',
    `namespace`   int(11) NOT NULL COMMENT '工作空间',
    `name`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '应用名',
    `token`       varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '鉴权token',
    `desc`        text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
    `tags`        text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联标签',
    `is_disable`  tinyint(1) NULL DEFAULT 0 COMMENT '禁用状态，0启用，1禁用',
    `is_delete`   tinyint(1) NULL DEFAULT 0 COMMENT '是否删除',
    `operator`    int(11) NULL DEFAULT NULL COMMENT '更新人/操作人',
    `create_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `update_time` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP (0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `unique_uuid` (`uuid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '外部应用表' ROW_FORMAT = Dynamic;