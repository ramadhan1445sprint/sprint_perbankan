variable "sg_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "docker_image_url" {
  type = string
}

variable "db_host" {
  type = string
}

variable "db_user" {
  type = string
}

variable "db_pass" {
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

variable "s3_bucket" {
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