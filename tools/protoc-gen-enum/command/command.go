package command

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gogo/protobuf/plugin/testgen"
	"github.com/gogo/protobuf/proto"
	plugin_go "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	"github.com/gogo/protobuf/vanity/command"
	"github.com/liov/tiga/tools/protoc-gen-enum/generator"
)

func Read() *plugin_go.CodeGeneratorRequest {
	g := generator.New()
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}
	return g.Request
}

func Generate(req *plugin_go.CodeGeneratorRequest) *plugin_go.CodeGeneratorResponse {
	// Begin by allocating a generator. The request and response structures are stored there
	// so we can do error handling easily - the response structure contains the field to
	// report failure.
	g := generator.New()
	g.Request = req

	g.CommandLineParameters(g.Request.GetParameter())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	g.WrapTypes()

	g.SetPackageNames()
	g.BuildTypeNameMap()

	g.GenerateAllFiles()

	if err := goformat(g.Response); err != nil {
		g.Error(err)
	}

	testReq := proto.Clone(req).(*plugin_go.CodeGeneratorRequest)

	testResp := command.GeneratePlugin(testReq, testgen.NewPlugin(), "pb_test.go")

	for i := 0; i < len(testResp.File); i++ {
		if strings.Contains(*testResp.File[i].Content, `//These tests are generated by github.com/gogo/protobuf/plugin/testgen`) {
			g.Response.File = append(g.Response.File, testResp.File[i])
		}
	}

	return g.Response
}

func Write(resp *plugin_go.CodeGeneratorResponse) {
	g := generator.New()
	// Send back the results.
	data, err := proto.Marshal(resp)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}

func goformat(resp *plugin_go.CodeGeneratorResponse) error {
	for i := 0; i < len(resp.File); i++ {
		formatted, err := format.Source([]byte(resp.File[i].GetContent()))
		if err != nil {
			return fmt.Errorf("go format error: %v", err)
		}
		fmts := string(formatted)
		resp.File[i].Content = &fmts
	}
	return nil
}