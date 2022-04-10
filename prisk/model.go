package prisk

type Area struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Key        string `json:"key"`
	ChildAreas Areas  `json:"items"`
}

func (a Areas) ToMap() map[string]AreaMap {
	areaMap := make(map[string]AreaMap)
	for _, v := range a {
		areaMap[v.Name] = v.ToMap()
	}
	return areaMap
}

func (a Area) ToMap() AreaMap {
	areaMap := AreaMap{
		Code:       a.Code,
		Name:       a.Name,
		Key:        a.Key,
		ChildAreas: make(map[string]AreaMap),
	}
	for _, v := range a.ChildAreas {
		areaMap.ChildAreas[v.Name] = AreaMap{
			Code:       v.Code,
			Name:       v.Name,
			Key:        v.Key,
			ChildAreas: v.ChildAreas.ToMap(),
		}
	}
	return areaMap
}

type AreaMap struct {
	Code       string             `json:"code"`
	Name       string             `json:"name"`
	Key        string             `json:"key"`
	ChildAreas map[string]AreaMap `json:"items"`
}

type DangerAreaListResponse struct {
	Data struct {
		EndUpdateTime string        `json:"end_update_time"`
		HighCount     int           `json:"hcount"`
		MiddleCount   int           `json:"mcount"`
		HighList      []*DangerArea `json:"highlist"`
		MiddleList    []*DangerArea `json:"middlelist"`
	}
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type DangerArea struct {
	Type        string   `json:"type"`
	Province    string   `json:"province"`
	City        string   `json:"city"`
	Country     string   `json:"county"`
	AreaName    string   `json:"area_name"`
	Communities []string `json:"communitys"`
}

type DangerAreaArr struct {
	HighRisk   map[string]string
	MiddleRisk map[string]string
}
