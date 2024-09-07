package reports

import "fmt"

type ReportType struct {
	TypeId       string `json:"id"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	CategoryName string `json:"category_name"`
}

type ReportTypes []ReportType

var AllReportTypes = ReportTypes{
	ReportType{TypeId: "ROBBERY", Name: "Assalto", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "MASSIVE_ROBBERY", Name: "Arrastão", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "SUSPECT_ACTIVITY", Name: "Atividade Suspeita", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "DRUG_TRAFFICKING", Name: "Tráfico de Drogas", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "VANDALISM", Name: "Vandalismo", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "KIDNAPPING", Name: "Sequestro", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "THEFT", Name: "Furto", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "ASSAULT", Name: "Agressão", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "SEXUAL_ASSAULT", Name: "Assédio Sexual", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "DOMESTIC_VIOLENCE", Name: "Violência Doméstica", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "STALKING", Name: "Perseguição", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "DEATH", Name: "Morte", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "MISSING_PERSON", Name: "Pessoa Desaparecida", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "HARASSMENT", Name: "Assédio", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "FIGHT", Name: "Briga", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "DRUNKENNESS", Name: "Embriaguez", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "DRUG_USE", Name: "Uso de Drogas", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "PROSTITUTION", Name: "Prostituição", Category: "CRIME", CategoryName: "Crime"},
	ReportType{TypeId: "ACCIDENT_SERIOUS", Name: "Acidente Grave", Category: "ACCIDENT", CategoryName: "Acidente"},
	ReportType{TypeId: "ACCIDENT_LIGHT", Name: "Acidente Leve", Category: "ACCIDENT", CategoryName: "Acidente"},
	ReportType{TypeId: "ACCIDENT_WITH_INJURIES", Name: "Acidente com Feridos", Category: "ACCIDENT", CategoryName: "Acidente"},
	ReportType{TypeId: "ACCIDENT_WITH_FATALITY", Name: "Acidente com Fatalidade", Category: "ACCIDENT", CategoryName: "Acidente"},
	ReportType{TypeId: "ACCIDENT_WITH_TRAPPED_PEOPLE", Name: "Acidente com Pessoas Presas", Category: "ACCIDENT", CategoryName: "Acidente"},
	ReportType{TypeId: "ACCIDENT_WITH_FIRE", Name: "Acidente com Fogo", Category: "ACCIDENT", CategoryName: "Acidente"},
	ReportType{TypeId: "ACCIDENT_WITH_SPILL", Name: "Acidente com Derramamento", Category: "ACCIDENT", CategoryName: "Acidente"},
	ReportType{TypeId: "ACCIDENT_WITH_ELECTRICITY", Name: "Acidente com Eletricidade", Category: "ACCIDENT", CategoryName: "Acidente"},
	ReportType{TypeId: "ACCIDENT_WITH_WATER", Name: "Acidente com Água", Category: "ACCIDENT", CategoryName: "Acidente"},
	ReportType{TypeId: "ACCIDENT_WITH_ANIMALS", Name: "Acidente com Animais", Category: "ACCIDENT", CategoryName: "Acidente"},
	ReportType{TypeId: "FLOOD", Name: "Enchente", Category: "NATURE", CategoryName: "Causas Naturais"},
	ReportType{TypeId: "LANDSLIDE", Name: "Deslizamento", Category: "NATURE", CategoryName: "Causas Naturais"},
	ReportType{TypeId: "FOG", Name: "Neblina", Category: "NATURE", CategoryName: "Causas Naturais"},
	ReportType{TypeId: "FREEZE_RAIN", Name: "Chuva de Granizo", Category: "NATURE", CategoryName: "Causas Naturais"},
	ReportType{TypeId: "SEVERAL_RAIN", Name: "Chuva Forte", Category: "NATURE", CategoryName: "Causas Naturais"},
	ReportType{TypeId: "FIRE", Name: "Incêndio", Category: "FIRE", CategoryName: "Incêndio"},
	ReportType{TypeId: "EXPLOSION", Name: "Explosão", Category: "FIRE", CategoryName: "Incêndio"},
	ReportType{TypeId: "GAS_LEAK", Name: "Vazamento de Gás", Category: "FIRE", CategoryName: "Incêndio"},
	ReportType{TypeId: "ELECTRIC_PROBLEM", Name: "Problema Elétrico", Category: "INFRASTRUCTURE", CategoryName: "Infraestrutura"},
	ReportType{TypeId: "WATER_PROBLEM", Name: "Problema de Água", Category: "INFRASTRUCTURE", CategoryName: "Infraestrutura"},
	ReportType{TypeId: "ROAD_PROBLEM", Name: "Problema de Estrada", Category: "INFRASTRUCTURE", CategoryName: "Infraestrutura"},
	ReportType{TypeId: "HOLE_ROAD", Name: "Buraco na Estrada", Category: "INFRASTRUCTURE", CategoryName: "Infraestrutura"},
	ReportType{TypeId: "OIL_SPILL", Name: "Derramamento de Óleo", Category: "INFRASTRUCTURE", CategoryName: "Infraestrutura"},
	ReportType{TypeId: "CLOSED_ROAD", Name: "Estrada Fechada", Category: "INFRASTRUCTURE", CategoryName: "Infraestrutura"},
	ReportType{TypeId: "CONSTRUCTION", Name: "Construção", Category: "INFRASTRUCTURE", CategoryName: "Infraestrutura"},
	ReportType{TypeId: "BAD_TRAFFIC_LIGHT", Name: "Semáforo Ruim", Category: "INFRASTRUCTURE", CategoryName: "Infraestrutura"},
	ReportType{TypeId: "BAD_ROAD_SIGNAL", Name: "Sinalização Ruim", Category: "INFRASTRUCTURE", CategoryName: "Infraestrutura"},
	ReportType{TypeId: "ANIMALS_ON_THE_ROAD", Name: "Animais na Estrada", Category: "INFRASTRUCTURE", CategoryName: "Infraestrutura"},
	ReportType{TypeId: "GENERIC", Name: "Não Listado", Category: "GENERIC", CategoryName: "Outros"},
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

//func GetTypesByCategory(category string) (*ReportTypes, error) {
//	var reportTypes ReportTypes
//	for _, reportType := range AllReportTypes {
//		if reportType.Category == category {
//			reportTypes = append(reportTypes, reportType)
//		}
//	}
//	return &reportTypes, nil
//}

func GetReportTypes() ReportTypes {
	return AllReportTypes
}
