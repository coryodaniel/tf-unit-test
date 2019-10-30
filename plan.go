package main

import "encoding/json"
import "fmt"
import "io/ioutil"
import "os"

type ChangedResource struct {
	Address       string `json:"address"`
	ModuleAddress string `json:"module_address"`
	Mode          string `json:"mode"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	ProviderName  string `json:"provider_name"`
	Change        struct {
		Actions []string    `json:"actions"`
		Before  interface{} `json:"before"`
		After   struct {
			Content             string      `json:"content"`
			ContentBase64       interface{} `json:"content_base64"`
			DirectoryPermission string      `json:"directory_permission"`
			FilePermission      string      `json:"file_permission"`
			Filename            string      `json:"filename"`
			SensitiveContent    interface{} `json:"sensitive_content"`
		} `json:"after"`
		AfterUnknown struct {
			ID bool `json:"id"`
		} `json:"after_unknown"`
	} `json:"change"`
}

type Plan struct {
	FormatVersion    string `json:"format_version"`
	TerraformVersion string `json:"terraform_version"`
	Variables        struct {
		Name struct {
			Value string `json:"value"`
		} `json:"name"`
	} `json:"variables"`
	PlannedValues struct {
		Outputs struct {
			File struct {
				Sensitive bool   `json:"sensitive"`
				Value     string `json:"value"`
			} `json:"file"`
		} `json:"outputs"`
		RootModule struct {
			ChildModules []struct {
				Resources []struct {
					Address       string `json:"address"`
					Mode          string `json:"mode"`
					Type          string `json:"type"`
					Name          string `json:"name"`
					ProviderName  string `json:"provider_name"`
					SchemaVersion int    `json:"schema_version"`
					Values        struct {
						Content             string      `json:"content"`
						ContentBase64       interface{} `json:"content_base64"`
						DirectoryPermission string      `json:"directory_permission"`
						FilePermission      string      `json:"file_permission"`
						Filename            string      `json:"filename"`
						SensitiveContent    interface{} `json:"sensitive_content"`
					} `json:"values"`
				} `json:"resources"`
				Address string `json:"address"`
			} `json:"child_modules"`
		} `json:"root_module"`
	} `json:"planned_values"`
	ResourceChanges []ChangedResource `json:"resource_changes"`
	OutputChanges   struct {
		File struct {
			Actions      []string    `json:"actions"`
			Before       interface{} `json:"before"`
			After        string      `json:"after"`
			AfterUnknown bool        `json:"after_unknown"`
		} `json:"file"`
	} `json:"output_changes"`
	Configuration struct {
		RootModule struct {
			Outputs struct {
				File struct {
					Expression struct {
						References []string `json:"references"`
					} `json:"expression"`
				} `json:"file"`
			} `json:"outputs"`
			ModuleCalls struct {
				Mymod struct {
					Source      string `json:"source"`
					Expressions struct {
						Ext struct {
							References []string `json:"references"`
						} `json:"ext"`
						Name struct {
							References []string `json:"references"`
						} `json:"name"`
					} `json:"expressions"`
					Module struct {
						Outputs struct {
							Filename struct {
								Expression struct {
									References []string `json:"references"`
								} `json:"expression"`
							} `json:"filename"`
						} `json:"outputs"`
						Resources []struct {
							Address           string `json:"address"`
							Mode              string `json:"mode"`
							Type              string `json:"type"`
							Name              string `json:"name"`
							ProviderConfigKey string `json:"provider_config_key"`
							Expressions       struct {
								Content struct {
									ConstantValue string `json:"constant_value"`
								} `json:"content"`
								Filename struct {
									References []string `json:"references"`
								} `json:"filename"`
							} `json:"expressions"`
							SchemaVersion int `json:"schema_version"`
						} `json:"resources"`
						Variables struct {
							Ext struct {
							} `json:"ext"`
							Name struct {
							} `json:"name"`
						} `json:"variables"`
					} `json:"module"`
				} `json:"mymod"`
			} `json:"module_calls"`
			Variables struct {
				Name struct {
				} `json:"name"`
			} `json:"variables"`
		} `json:"root_module"`
	} `json:"configuration"`
}

// GetPlan returns Plan
func GetPlan() Plan {
	const plan = "./tf.json"
	jsonFile, err := os.Open(plan)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Plan
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
