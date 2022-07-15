DROP TABLE IF EXISTS users cascade;
DROP TABLE IF EXISTS tasks cascade;
DROP TABLE IF EXISTS users_tasks;

CREATE TYPE user_role AS ENUM ('super_admin', 'admin','end_user');

CREATE TABLE "users" (
  "id" varchar(64) PRIMARY KEY,
  "name" varchar(200) NOT NULL DEFAULT '',    
  "email" varchar(512) UNIQUE NOT NULL DEFAULT '',
  "password" varchar(200),
  "role_type" user_role, 
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TYPE task_status AS ENUM ('not_scoped','scoped','in_progress', 'code_review', 'mr_approved');
CREATE TABLE "tasks" (
  "id" varchar(64) PRIMARY KEY,
  "descreption" varchar(200) NOT NULL DEFAULT '',
  "task_status_code" task_status,
  "started_at" timestamp,
  "ended_at" timestamp
);

CREATE TABLE "users_tasks" (
"id" varchar(64) PRIMARY KEY,
user_id varchar(64) NOT NULL references users(id),
task_id varchar(64) NOT NULL references tasks(id)
);