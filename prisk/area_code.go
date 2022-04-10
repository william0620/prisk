package prisk

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

func (pr PRisk) GetAreaList() Areas {
	var request = pr.GetRequest()
	var url = "http://bmfw.www.gov.cn/myqfxdjcx/source/WAP/js/area.js"
	resp, body, errs := request.Get(url).End()
	fmt.Println(resp.StatusCode)
	if errs != nil {
		panic(errs[0])
	}
	rawJson := mustParseArea(body)
	var areas Areas
	err := json.Unmarshal([]byte(rawJson), &areas)
	if err != nil {
		panic(err)
	}
	areaMap := areas.ToMap()
	areaMapJson, _ := json.Marshal(areaMap)
	err = ioutil.WriteFile("areaMap.json", areaMapJson, 0644)
	if err != nil {
		panic(err)
	}
	return areas
}

func GetDangerAreaMapFromFile() map[string]AreaMap {
	areaMap := make(map[string]AreaMap)
	bytes, err := ioutil.ReadFile("areaMap.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &areaMap)
	if err != nil {
		panic(err)
	}
	return areaMap
}

type Areas []Area

func mustParseArea(areaJS string) string {
	vm := otto.New()

	_, err := vm.Run(string(areaJS))
	if err != nil {
		panic(err)
	}

	value, err := vm.Eval("JSON.stringify(area)")
	if err != nil {
		panic(err)
	}
	var result = value.String()
	//fmt.Println(result)
	return result
}
