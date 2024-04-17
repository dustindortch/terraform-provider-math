terraform {
  required_providers {
    math = {
      source = "hashicorp.com/edu/math"
    }
  }
}

variable "function" {
  type    = number
  default = 10
}

variable "resource" {
  type    = number
  default = 12
}

resource "math_fib" "sequence" {
  number = var.resource
}

output "func_sequence" {
  value = provider::math::fib(var.function)
}


output "res_sequence" {
  value = math_fib.sequence.result
}