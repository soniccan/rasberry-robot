

# CircleCI のワークロード ID プールを作成
locals {
  circleci_org = "rasberry-robot"
}

resource "google_iam_workload_identity_pool" "circleci" {
  provider                  = google-beta
  workload_identity_pool_id = "circleci"
  display_name              = "circleci"
}

resource "google_iam_workload_identity_pool_provider" "circleci" {
  provider                           = google-beta
  workload_identity_pool_id          = google_iam_workload_identity_pool.circleci.workload_identity_pool_id
  workload_identity_pool_provider_id = "circleci"
  display_name                       = "circleci"
  attribute_mapping = {
    "attribute.project_id" = "assertion['oidc.circleci.com/project-id']" # CircleCI の Project を制限するためにマッピングを行います。
    "google.subject"       = "assertion.sub"
  }
  oidc {
    issuer_uri = "https://oidc.circleci.com/org/${local.circleci_org}" # Issuer URL は CircleCI Organization ごとに発行されます。
    allowed_audiences = [local.circleci_org]
  }
}


# デフォルトの GSA を作成
locals {
  gsa_name = "circleci-sa"
}

resource "google_service_account" "default" {
  account_id   = local.gsa_name
  display_name = local.gsa_name
}

resource "google_service_account_iam_member" "default" {
  for_each = toset([
    "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.circleci.name}/attribute.project_id/<CircleCI Project ID>",
  ])
  service_account_id = google_service_account.default.id
  role               = "roles/iam.workloadIdentityUser"
  member             = each.value
}

resource "google_project_iam_member" "default" {
  for_each = toset([
    "roles/editor", # 操作したい権限を付与します
  ])
  project = var.project
  role    = each.value
  member  = "serviceAccount:${google_service_account.default.email}"
}