resource "aws_ecs_cluster" "ecs_cluster" {
  name = "tf_ilham_paimonbank_go_pg"
}

resource "aws_ecs_task_definition" "web_backend_task" {
  cpu                      = 1024
  memory                   = 4096
  family                   = "web_backend_task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = var.ecs_exec_role_arn

  container_definitions = jsonencode([
    {
      name      = "paimon_bank_ilhamnyto"
      cpu       = 1024
      memory    = 4096
      image     = var.docker_image_url
      essential = true
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
      environment = [
        { "name" : "DB_HOST", "value" : "${var.db_host}" },
        { "name" : "DB_USERNAME", "value" : "${var.db_username}" },
        { "name" : "DB_NAME", "value" : "${var.db_name}" },
        { "name" : "DB_PASSWORD", "value" : "${var.db_password}" },
        { "name" : "DB_PORT", "value" : "${var.db_port}" },
        { "name" : "DB_PARAMS", "value" : "${var.db_params}" },
        { "name" : "BCRYPT_SALT", "value" : "${var.bcrypt_salt}" },
        { "name" : "JWT_SECRET", "value" : "${var.jwt_secret}" },
        { "name" : "S3_ID", "value" : "${var.s3_id}" },
        { "name" : "S3_SECRET", "value" : "${var.s3_secret}" },
        { "name" : "S3_BUCKET_NAME", "value" : "${var.s3_bucket_name}" },
        { "name" : "S3_REGION", "value" : "ap-southeast-1" },
      ]
    }
  ])
}

resource "aws_ecs_service" "web_backend_service" {
  name            = "web_backend_service"
  cluster         = aws_ecs_cluster.ecs_cluster.id
  task_definition = aws_ecs_task_definition.web_backend_task.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = var.subnet_ids
    security_groups  = [var.sg_id]
    assign_public_ip = true
  }
}