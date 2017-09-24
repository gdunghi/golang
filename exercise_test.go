package golang

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

func TestJsonStructure(t *testing.T) {
	jsonStr := `{
	"id": 1,
	"lang": "go"		
}`

	//type testStruct = langStruct
	var myStruct langStruct

	err := json.Unmarshal([]byte(jsonStr), &myStruct)

	if err != nil {
		t.Error("it should not error when you make the righ struct that represent json string")
	}
}

func TestJsonStructures(t *testing.T) {
	jsonStr := `[
	{
		"id": 1,
		"lang": "go"
	},
	{
		"id": 2,
		"lang": "java"
	},
	{
		"id": 3,
		"lang": "python"
	}
]`

	//type testStruct = langStruct
	var myStruct []langStruct

	err := json.Unmarshal([]byte(jsonStr), &myStruct)

	if err != nil {
		t.Error("it should not error when you make the righ struct that represent json string")
	}
}

func TestJsonAndXMLStructure(t *testing.T) {
	xmlStr := `<xml>
	<id>1</id>
	<lang>go</lang>
	<Book>
		<auther>tom</auther>
		<name>naja</name>
	</Book>
	</xml>
}`

	//type testStruct = langStruct
	var myStruct langStruct

	err := xml.Unmarshal([]byte(xmlStr), &myStruct)

	if err != nil {
		t.Error("it should not error when you make the righ struct that represent json string")
	}

}
