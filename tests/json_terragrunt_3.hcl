terraform {
  source = "relatvie/path/to/terraform"
}

dependency {
  config_path = "relative/path/to/terragrunt/dependency"
}

include {
  path = find_in_parent_folders()
}

locals {
  local_var = [
    1, 2, 3
  ]
}

inputs = {
  var_1 = "brandfolder-storage-dev"
  var_2 : dependency.some_terragrunt_file.outputs.some_var
  var_3 = [
    for num in local.local_var :
    {
      var_4 : dependency.some_terragrunt_file.outputs.some_var[num],
      var_5 : ependency.some_terragrunt_file.outputs.some_var[num],
    }
  ]
  var_6 = [
    "2",
  ]
}
