provider "google" {
  project = var.project
  region  = var.region
}

provider "google-beta" {
  project = var.project
  region  = var.region
}

variable "region" {
  description = "GCPリージョン"
  type        = string
  default     = "asia-northeast1"
}