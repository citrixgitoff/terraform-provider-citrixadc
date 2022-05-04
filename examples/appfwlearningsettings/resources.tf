resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name                     = "tf_appfwprofile"
  bufferoverflowaction     = ["none"]
  contenttypeaction        = ["none"]
  cookieconsistencyaction  = ["none"]
  creditcard               = ["none"]
  creditcardaction         = ["none"]
  crosssitescriptingaction = ["none"]
  csrftagaction            = ["none"]
  denyurlaction            = ["none"]
  dynamiclearning          = ["none"]
  fieldconsistencyaction   = ["none"]
  fieldformataction        = ["none"]
  fileuploadtypesaction    = ["none"]
  inspectcontenttypes      = ["none"]
  jsondosaction            = ["none"]
  jsonsqlinjectionaction   = ["none"]
  jsonxssaction            = ["none"]
  multipleheaderaction     = ["none"]
  sqlinjectionaction       = ["none"]
  starturlaction           = ["none"]
  type                     = ["HTML"]
  xmlattachmentaction      = ["none"]
  xmldosaction             = ["none"]
  xmlformataction          = ["none"]
  xmlsoapfaultaction       = ["none"]
  xmlsqlinjectionaction    = ["none"]
  xmlvalidationaction      = ["none"]
  xmlwsiaction             = ["none"]
  xmlxssaction             = ["none"]
}
resource "citrixadc_appfwlearningsettings" "tf_learningsetting" {
  profilename                        = citrixadc_appfwprofile.tf_appfwprofile.name
  starturlminthreshold               = 9
  starturlpercentthreshold           = 10
  cookieconsistencyminthreshold      = 2
  cookieconsistencypercentthreshold  = 1
  csrftagminthreshold                = 2
  csrftagpercentthreshold            = 10
  fieldconsistencyminthreshold       = 20
  fieldconsistencypercentthreshold   = 8
  crosssitescriptingminthreshold     = 10
  crosssitescriptingpercentthreshold = 1
  sqlinjectionminthreshold           = 10
  sqlinjectionpercentthreshold       = 1
  fieldformatminthreshold            = 10
  fieldformatpercentthreshold        = 1
  creditcardnumberminthreshold       = 1
  creditcardnumberpercentthreshold   = 0
  contenttypeminthreshold            = 1
  contenttypepercentthreshold        = 0
}