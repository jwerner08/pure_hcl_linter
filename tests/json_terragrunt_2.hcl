terraform {
  source = "relatvie/path/to/terraform"
}

include {
  path = find_in_parent_folders()
}

inputs = {
  var_1 = "'Pure' HCL doesn't quote keys, and doesn't have colons."
  var_2 = "It is generally easier for humans to read vs. JSON."
  var_3 = [
    { nested_var_1 : "This is JSON-like.", },
    { nested_var_2 : "and so is this", },
  ]
}
