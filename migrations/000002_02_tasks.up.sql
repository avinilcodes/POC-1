CREATE TYPE task_status AS ENUM ('not_scoped','scoped','in_progress', 'code_review', 'mr_approved');
CREATE TABLE "tasks" (
  "id" varchar(64) PRIMARY KEY,
  "description" varchar(200) NOT NULL DEFAULT '',
  "task_status_code" task_status,
  "started_at" timestamp,
  "ended_at" timestamp
);