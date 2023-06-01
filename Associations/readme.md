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
