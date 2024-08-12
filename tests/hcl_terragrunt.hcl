terraform {
  source = "relative/path/to/terraform"
}

include {
  path = find_in_parent_folders()
}
inputs = {
  var_1 = "Some:pure:hcl:values:may:have:colons."
  var_2 = "For instance, urls may have colons."
  var_3 = "In those instances, terragrunt files shouldn't get flag the linter."

  # Comments : may : also : have : colons.
  var_4 = "Those instances shouldn't get flagged either."
  var_5 = "Some Cases may not be tested fully."
  var_6 = "But we can always add more!"
}
