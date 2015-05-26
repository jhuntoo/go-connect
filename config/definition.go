package config

type AppConfigDefinition struct {
	 SoftwareStatement SoftwareStatementSectionDefinition
}

func NewAppConfigDefinition() AppConfigDefinition {
	return AppConfigDefinition {
		SoftwareStatement : newSoftwareStatementSectionDefinition(),
	}
}

type SectionDefinition struct{
	SectionName string
	SectionPath string
}


type SoftwareStatementSectionDefinition struct {
	SectionDefinition
	SigningMethod stringItemDefinition
	
		
}
func newSoftwareStatementSectionDefinition() SoftwareStatementSectionDefinition {
	return SoftwareStatementSectionDefinition { 
			SectionDefinition:SectionDefinition {SectionName: "softwarestatement", SectionPath: "auth2"},
		 	SigningMethod: stringItemDefinition { Key:"signing_method",
												DefaultValue: "RS256", 
												AllowedValues : []string{"RS256"},
											},
		}
}

type itemDefinition interface {
	GetKey() string
}
    
type stringItemDefinition struct {
	Key string
	itemDefinition
	DefaultValue string
	AllowedValues []string
}

func(i *stringItemDefinition) GetKey() string {
	return i.Key
}

