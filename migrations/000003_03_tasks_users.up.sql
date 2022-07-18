CREATE TABLE "users_tasks" (
user_id varchar(64) NOT NULL references users(id),
task_id varchar(64) NOT NULL references tasks(id)
);