package services

import "fmt"

type ReportType struct {
	TypeId   string `json:"type_id"`
	Name     string `json:Name`
	Category string `json:Category`
}

type ReportTypes []ReportType

var AllReportTypes = ReportTypes{
	ReportType{TypeId: "ROBBERY", Name: "Assalto", Category: "CRIME"},
	ReportType{TypeId: "MASSIVE_ROBBERY", Name: "Arrastão", Category: "CRIME"},
	ReportType{TypeId: "SUSPECT_ACTIVITY", Name: "Atividade Suspeita", Category: "CRIME"},
	ReportType{TypeId: "DRUG_TRAFFICKING", Name: "Tráfico de Drogas", Category: "CRIME"},
	ReportType{TypeId: "VANDALISM", Name: "Vandalismo", Category: "CRIME"},
	ReportType{TypeId: "KIDNAPPING", Name: "Sequestro", Category: "CRIME"},
	ReportType{TypeId: "THEFT", Name: "Furto", Category: "CRIME"},
	ReportType{TypeId: "ASSAULT", Name: "Agressão", Category: "CRIME"},
	ReportType{TypeId: "SEXUAL_ASSAULT", Name: "Assédio Sexual", Category: "CRIME"},
	ReportType{TypeId: "DOMESTIC_VIOLENCE", Name: "Violência Doméstica", Category: "CRIME"},
	ReportType{TypeId: "STALKING", Name: "Perseguição", Category: "CRIME"},
	ReportType{TypeId: "DEATH", Name: "Morte", Category: "CRIME"},
	ReportType{TypeId: "MISSING_PERSON", Name: "Pessoa Desaparecida", Category: "CRIME"},
	ReportType{TypeId: "HARASSMENT", Name: "Assédio", Category: "CRIME"},
	ReportType{TypeId: "FIGHT", Name: "Briga", Category: "CRIME"},
	ReportType{TypeId: "DRUNKENNESS", Name: "Embriaguez", Category: "CRIME"},
	ReportType{TypeId: "DRUG_USE", Name: "Uso de Drogas", Category: "CRIME"},
	ReportType{TypeId: "PROSTITUTION", Name: "Prostituição", Category: "CRIME"},
	ReportType{TypeId: "OTHER_CRIME", Name: "Outro Crime", Category: "CRIME"},
	ReportType{TypeId: "ACCIDENT_SERIOUS", Name: "Acidente Grave", Category: "ACCIDENT"},
	ReportType{TypeId: "ACCIDENT_LIGHT", Name: "Acidente Leve", Category: "ACCIDENT"},
	ReportType{TypeId: "ACCIDENT_WITH_INJURIES", Name: "Acidente com Feridos", Category: "ACCIDENT"},
	ReportType{TypeId: "ACCIDENT_WITH_FATALITY", Name: "Acidente com Fatalidade", Category: "ACCIDENT"},
	ReportType{TypeId: "ACCIDENT_WITH_TRAPPED_PEOPLE", Name: "Acidente com Pessoas Presas", Category: "ACCIDENT"},
	ReportType{TypeId: "ACCIDENT_WITH_FIRE", Name: "Acidente com Fogo", Category: "ACCIDENT"},
	ReportType{TypeId: "ACCIDENT_WITH_SPILL", Name: "Acidente com Derramamento", Category: "ACCIDENT"},
	ReportType{TypeId: "ACCIDENT_WITH_ELECTRICITY", Name: "Acidente com Eletricidade", Category: "ACCIDENT"},
	ReportType{TypeId: "ACCIDENT_WITH_WATER", Name: "Acidente com Água", Category: "ACCIDENT"},
	ReportType{TypeId: "ACCIDENT_WITH_ANIMALS", Name: "Acidente com Animais", Category: "ACCIDENT"},
	ReportType{TypeId: "FLOOD", Name: "Enchente", Category: "NATURE"},
	ReportType{TypeId: "LANDSLIDE", Name: "Deslizamento", Category: "NATURE"},
	ReportType{TypeId: "FOG", Name: "Neblina", Category: "NATURE"},
	ReportType{TypeId: "FREEZE_RAIN", Name: "Chuva de Granizo", Category: "NATURE"},
	ReportType{TypeId: "SEVERAL_RAIN", Name: "Chuva Forte", Category: "NATURE"},
	ReportType{TypeId: "OTHER_NATURE", Name: "Outra Natureza", Category: "NATURE"},
	ReportType{TypeId: "FIRE", Name: "Incêndio", Category: "FIRE"},
	ReportType{TypeId: "EXPLOSION", Name: "Explosão", Category: "FIRE"},
	ReportType{TypeId: "GAS_LEAK", Name: "Vazamento de Gás", Category: "FIRE"},
	ReportType{TypeId: "OTHER_FIRE", Name: "Outro Fogo", Category: "FIRE"},
	ReportType{TypeId: "ELECTRIC_PROBLEM", Name: "Problema Elétrico", Category: "INFRASTRUCTURE"},
	ReportType{TypeId: "WATER_PROBLEM", Name: "Problema de Água", Category: "INFRASTRUCTURE"},
	ReportType{TypeId: "ROAD_PROBLEM", Name: "Problema de Estrada", Category: "INFRASTRUCTURE"},
	ReportType{TypeId: "HOLE_ROAD", Name: "Buraco na Estrada", Category: "INFRASTRUCTURE"},
	ReportType{TypeId: "OIL_SPILL", Name: "Derramamento de Óleo", Category: "INFRASTRUCTURE"},
	ReportType{TypeId: "CLOSED_ROAD", Name: "Estrada Fechada", Category: "INFRASTRUCTURE"},
	ReportType{TypeId: "CONSTRUCTION", Name: "Construção", Category: "INFRASTRUCTURE"},
	ReportType{TypeId: "BAD_TRAFFIC_LIGHT", Name: "Semáforo Ruim", Category: "INFRASTRUCTURE"},
	ReportType{TypeId: "BAD_ROAD_SIGNAL", Name: "Sinalização Ruim", Category: "INFRASTRUCTURE"},
	ReportType{TypeId: "ANIMALS_ON_THE_ROAD", Name: "Animais na Estrada", Category: "INFRASTRUCTURE"},
	ReportType{TypeId: "OTHER_INFRASTRUCTURE", Name: "Outra Infraestrutura", Category: "INFRASTRUCTURE"},
	ReportType{TypeId: "GENERIC", Name: "Não Listado", Category: "GENERIC"},
}

func GetReportType(typeId string) (*ReportType, error) {
	if typeId == "" {
		return nil, fmt.Errorf("report is null")
	}
	for _, reportType := range AllReportTypes {
		if reportType.TypeId == typeId {
			return &reportType, nil
		}
	}
	return nil, fmt.Errorf("report type not found")
}

func GetReportTypeName(typeId string) *string {
	if typeId == "" {
		genericReturn := &ReportType{TypeId: "GENERIC", Name: "Não Listado", Category: "GENERIC"}
		return &genericReturn.Name
	}
	reportType, err := GetReportType(typeId)
	if err != nil {
		return nil
	}
	return &reportType.Name
}

func GetReportTypesByCategory(category string) (*ReportTypes, error) {
	var reportTypes ReportTypes
	for _, reportType := range AllReportTypes {
		if reportType.Category == category {
			reportTypes = append(reportTypes, reportType)
		}
	}
	return &reportTypes, nil
}

func GetReportTypes() ReportTypes {
	return AllReportTypes
}
