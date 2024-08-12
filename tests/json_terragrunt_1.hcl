terraform {
  source = "relative/path/to/terraform"
}

include {
  path = find_in_parent_folders()
}


inputs = {
  variable = {
    "The following colon should be an equal sign" : {
      var_1 = "Colons indicate JSON syntax."
      var_2 = "'Pure' HCL is a superset of json."
      var_3 = "mixing the two is stylistically inconsistent."
    }
  }
}
