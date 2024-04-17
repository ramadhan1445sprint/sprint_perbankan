variable "sg_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "aws_vpc_id" {
  type = string
}

variable "docker_image_url" {
  type = string
}

variable "docker_image_url_prometheus" {
  type = string
}

variable "docker_image_url_grafana" {
  type = string
}

variable "db_host" {
  type = string
}

variable "db_username" {
  type = string
}

variable "db_password" {
  type = string
  sensitive = true
}

variable "db_name" {
  type = string
}

variable "db_port" {
  type = string
}

variable "db_params" {
  type = string
}

variable "ecs_ak_id" {
  type = string
  sensitive = true
}

variable "ecs_ak_secret" {
  type = string
  sensitive = true
}

variable "ecs_exec_role_arn" {
  type = string
}

variable "s3_bucket_name" {
  type = string
}

variable "s3_id" {
  type = string
  sensitive = true
}

variable "s3_secret" {
  type = string
  sensitive = true
}

variable "bcrypt_salt" {
  type = string
  sensitive = true
}

variable "jwt_secret" {
  type = string
  sensitive = true
}