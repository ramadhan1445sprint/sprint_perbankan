resource "aws_ecs_cluster" "ecs_cluster" {
  name = "ilhamnyto_paimon_bank"
}

data "aws_ecs_cluster" "ecs_cluster" {
  cluster_name = "ilhamnyto_paimon_bank"
}

resource "aws_ecs_task_definition" "web_backend_task" {
  cpu                      = 2048
  memory                   = 6144
  family                   = "ilhamnyto_task"
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
        { "name" : "S3_SECRET_KEY", "value" : "${var.s3_secret}" },
        { "name" : "S3_BUCKET_NAME", "value" : "${var.s3_bucket_name}" },
        { "name" : "S3_REGION", "value" : "ap-southeast-1" },
      ]
    },
    {
      name      = "paimon_bank_ilhamnyto_prometheus"
      image     = var.docker_image_url_prometheus
      cpu       = 512
      memory    = 1024
      essential = false
      portMappings = [
        {
          containerPort = 9090
          hostPort      = 9090
          protocol      = "tcp"
        }
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-create-group"  = "true"
          "awslogs-group"         = "/ecs/web_backend_service"
          "awslogs-region"        = "ap-southeast-1"
          "awslogs-stream-prefix" = "ecs"
        }
      }
    },
    {
      name      = "paimon_bank_ilhamnyto_grafana"
      image     = var.docker_image_url_grafana
      cpu       = 512
      memory    = 1024
      essential = false
      portMappings = [
        {
          containerPort = 3000
          hostPort      = 3000
        }
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-create-group"  = "true"
          "awslogs-group"         = "/ecs/web_backend_service"
          "awslogs-region"        = "ap-southeast-1"
          "awslogs-stream-prefix" = "ecs"
        }
      }
    }
  ])
}

resource "aws_ecs_service" "web_backend_service" {
  name            = "ilhamnyto_service"
  cluster         = aws_ecs_cluster.ecs_cluster.id
  task_definition = aws_ecs_task_definition.web_backend_task.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = var.subnet_ids
    security_groups  = [var.sg_id]
    assign_public_ip = true
  }

  # load_balancer {
  # 	target_group_arn = aws_lb_target_group.backend_tg.arn
  # 	container_name   = var.docker_image_url
  # 	container_port   = 8080
  # }
}

# resource "aws_lb" "backend_lb" {
# 	name               = "backend-lb"
# 	internal           = false
# 	subnets            = var.subnet_ids
# 	security_groups    = [var.sg_id]
# 	load_balancer_type = "application"
# }

# resource "aws_lb_target_group" "backend_tg" {
# 	name        = "backend-tg"
# 	port        = 8080
# 	protocol    = "HTTP"
# 	vpc_id      = var.aws_vpc_id
# 	target_type = "ip"
# }

# resource "aws_lb_listener" "backend_listener" {
# 	load_balancer_arn = aws_lb.backend_lb.arn
# 	port              = "8080"
# 	default_action {
# 		type             = "forward"
# 		target_group_arn = aws_lb_target_group.backend_tg.arn
# 	}
# }
