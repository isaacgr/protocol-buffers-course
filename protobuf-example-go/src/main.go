package main

import (
	complexpb "complex"
	enumpb "enum"
	"fmt"
	"io/ioutil"
	"log"
	simplepb "simple"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	// sm := doSimple()
	// writeToFile("simple.bin", sm)

	// sm2 := &simplepb.SimpleMessage{}

	// readFromFile("simple.bin", sm2)
	// fmt.Println(sm2)

	// smAsString := toJSON(sm)
	// fmt.Println(smAsString)

	// sm3 := &simplepb.SimpleMessage{}
	// fromJSON(smAsString, sm3)
	// fmt.Println(sm3)
	// doEnum()
	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Firsts message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Firstss message",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_FRIDAY,
	}
	fmt.Println(em)
}

func toJSON(pb proto.Message) string {
	marshaler := protojson.MarshalOptions{
		Indent:          "  ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}
	out, err := marshaler.Marshal(pb)
	if err != nil {
		return ""
	}
	return string(out)
}

func fromJSON(in string, pb proto.Message) {
	err := protojson.Unmarshal([]byte(in), pb)
	if err != nil {
		log.Fatal("error")
	}
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err != nil {
		return err2
	}
	return nil
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		return err
	}
	fmt.Println("Data has been written")
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	sm.Name = "New message"
	return &sm
}
