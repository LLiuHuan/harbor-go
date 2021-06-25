package schema

type GetSystemInfoOptions struct {
}

type GetSystemInfo struct {
	ReadOnly          bool `json:"read_only"`
	AuthproxySettings struct {
		ServerCertificate   string `json:"server_certificate"`
		TokenreivewEndpoint string `json:"tokenreivew_endpoint"`
		Endpoint            string `json:"endpoint"`
		VerifyCert          bool   `json:"verify_cert"`
		SkipSearch          bool   `json:"skip_search"`
	} `json:"authproxy_settings"`
	HarborVersion               string `json:"harbor_version"`
	NotificationEnable          bool   `json:"notification_enable"`
	AuthMode                    string `json:"auth_mode"`
	SelfRegistration            bool   `json:"self_registration"`
	ExternalUrl                 string `json:"external_url"`
	ProjectCreationRestriction  string `json:"project_creation_restriction"`
	HasCaRoot                   bool   `json:"has_ca_root"`
	WithNotary                  bool   `json:"with_notary"`
	RegistryStorageProviderName string `json:"registry_storage_provider_name"`
	WithChartmuseum             bool   `json:"with_chartmuseum"`
	RegistryUrl                 string `json:"registry_url"`
}
