// ---------------for has-one usage----------------------------------------
CREATE TABLE employees
(
	`id`              INT NOT NULL AUTO_INCREMENT,
	`name`            VARCHAR(125),
	`email`           VARCHAR(125),
	`age`             tinyint,
	`birthday`        TIMESTAMP NULL,
	`member_number`   VARCHAR(125),
	`activated_at`    TIMESTAMP NULL,
	`created_at`      TIMESTAMP NULL,
	`updated_at`      TIMESTAMP NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
DEFAULT CHARSET = UTF8MB4;

CREATE TABLE credit_cards 
(
`id`              INT NOT NULL AUTO_INCREMENT,
`number`   		  VARCHAR(125),
`empoyee_id`      INT NOT NULL,
`created_at`      TIMESTAMP NULL,
`updated_at`      TIMESTAMP NULL,
PRIMARY KEY (`id`),
FOREIGN KEY (empoyee_id) REFERENCES employees(`id`)
)
ENGINE = InnoDB
DEFAULT CHARSET = UTF8MB4;

// ---------------for polymorphic usage ----------------------------------------
CREATE TABLE shapes
(
`id`              INT NOT NULL AUTO_INCREMENT,
`name`            VARCHAR(125),
`child_id`        INT,
`child_type`      VARCHAR(125),
`created_at`      TIMESTAMP NULL,
`updated_at`      TIMESTAMP NULL,
`deleted_at`	  TIMESTAMP NULL,
PRIMARY KEY (`id`)
)
ENGINE = InnoDB
DEFAULT CHARSET = UTF8MB4;

CREATE TABLE circles
(
`id`              INT NOT NULL AUTO_INCREMENT,
`name`            VARCHAR(125),
`created_at`      TIMESTAMP NULL,
`updated_at`      TIMESTAMP NULL,
`deleted_at`	  TIMESTAMP NULL,
PRIMARY KEY (`id`)
)
ENGINE = InnoDB
DEFAULT CHARSET = UTF8MB4;

CREATE TABLE squares
(
`id`              INT NOT NULL AUTO_INCREMENT,
`name`            VARCHAR(125),
`created_at`      TIMESTAMP NULL,
`updated_at`      TIMESTAMP NULL,
`deleted_at`	  TIMESTAMP NULL,
PRIMARY KEY (`id`)
)
ENGINE = InnoDB
DEFAULT CHARSET = UTF8MB4;

// --------------for has many usage---------------------------------
CREATE TABLE teachers
(
`id`              INT NOT NULL AUTO_INCREMENT,
`name`            VARCHAR(125),
`created_at`      TIMESTAMP NULL,
`updated_at`      TIMESTAMP NULL,
`deleted_at`	  TIMESTAMP NULL,
PRIMARY KEY (`id`)
)
ENGINE = InnoDB
DEFAULT CHARSET = UTF8MB4;

CREATE TABLE students
(
`id`              INT NOT NULL AUTO_INCREMENT,
`name`            VARCHAR(125),
`age`      		  INT,
`teacher_id`      INT,
`created_at`      TIMESTAMP NULL,
`updated_at`      TIMESTAMP NULL,
`deleted_at`	  TIMESTAMP NULL,
PRIMARY KEY (`id`),
FOREIGN KEY (teacher_id) REFERENCES teachers(`id`)
)
ENGINE = InnoDB
DEFAULT CHARSET = UTF8MB4;

insert into teachers(`name`, `created_at`, `updated_at`) values('teacher1', NOW(), NOW());
insert into students(`name`, `age`, `teacher_id`, `created_at`, `updated_at`) values('student1', 10,  1, NOW(), NOW());
insert into students(`name`, `age`, `teacher_id`, `created_at`, `updated_at`) values('student2', 11,  1, NOW(), NOW());

// --------------for many to many usage---------------------------------
CREATE TABLE humen
(
`id`              BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
`name`            VARCHAR(125),
`created_at`      TIMESTAMP NULL,
`updated_at`      TIMESTAMP NULL,
`deleted_at`	  TIMESTAMP NULL,
PRIMARY KEY (`id`)
)
ENGINE = InnoDB
DEFAULT CHARSET = UTF8MB4;

CREATE TABLE `languages`
(
`id`              BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
`name`            VARCHAR(125),
`created_at`      TIMESTAMP NULL,
`updated_at`      TIMESTAMP NULL,
`deleted_at`	  TIMESTAMP NULL,
PRIMARY KEY (`id`)
)
ENGINE = InnoDB
DEFAULT CHARSET = UTF8MB4;

CREATE TABLE human_languages
(
	`language_id` BIGINT(20) UNSIGNED NOT NULL,
	`human_id` BIGINT(20) UNSIGNED NOT NULL,
	PRIMARY KEY (`language_id`,`human_id`),
	KEY `fk_human_languages_human` (`human_id`),
	CONSTRAINT `fk_human_languages_human` FOREIGN KEY (`human_id`) REFERENCES `humen` (`id`),
    CONSTRAINT `fk_human_languages_language` FOREIGN KEY (`language_id`) REFERENCES `languages` (`id`)
)
ENGINE = InnoDB
DEFAULT CHARSET = UTF8MB4;

insert into humen(`name`, `created_at`, `updated_at` ) values('Janessa', NOW(), NOW());
insert into humen(`name`, `created_at`, `updated_at` ) values('yarenty', NOW(), NOW());
insert into `languages`(`name`, `created_at`, `updated_at` ) values('english', NOW(), NOW());
insert into `languages`(`name`, `created_at`, `updated_at` ) values('chinese', NOW(), NOW());
insert into `human_languages`(`language_id`, `human_id`) values(1, 1);
insert into `human_languages`(`language_id`, `human_id`) values(1, 2);
insert into `human_languages`(`language_id`, `human_id`) values(2, 1);