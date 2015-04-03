// Copyright 2015 Matthew Collins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	idSearchString = "Currently the packet id is: 0x"
)

var (
	protocol, dir string

	structs = map[string]*ast.TypeSpec{}
	packets []packet
	imports = map[string]struct{}{}
)

type packet struct {
	id   int
	name string
}

func main() {
	if len(os.Args) != 4 {
		log.Println("Missing target, protocol or dir")
		os.Exit(4)
	}

	input := os.Args[1]
	protocol = os.Args[2]
	dir = os.Args[3]

	fs := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fs, input, nil, parser.ParseComments)
	if err != nil {
		log.Fatalln(err)
	}

	for _, decl := range parsedFile.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			if decl.Tok != token.TYPE {
				continue
			}

			if len(decl.Specs) != 1 {
				return
			}

			tSpec, ok := decl.Specs[0].(*ast.TypeSpec)
			if !ok {
				continue
			}
			_, ok = tSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			structs[tSpec.Name.Name] = tSpec

			if decl.Doc == nil {
				continue
			}
			doc := decl.Doc.Text()
			pos := strings.Index(doc, idSearchString)
			if pos == -1 {
				continue
			}

			packetID, err := strconv.ParseInt(strings.TrimSpace(doc[pos+len(idSearchString):]), 16, 32)
			if err != nil {
				panic(err)
			}
			packets = append(packets, packet{
				id:   int(packetID),
				name: tSpec.Name.Name,
			})
		}
	}

	var buf bytes.Buffer

	// Packets
	for _, p := range packets {
		imports["io"] = struct{}{}
		short := string(strings.ToLower(p.name)[0])

		fmt.Fprintf(&buf, "func (%s *%s) id() int { return %d; }\n", short, p.name, p.id)

		fmt.Fprintf(&buf, "func (%s *%s) write(ww io.Writer) (err error) { \n", short, p.name)
		w := &writing{
			base: short,
			out:  &buf,
		}
		w.writeStruct(structs[p.name].Type.(*ast.StructType), short)
		w.flush()
		buf.WriteString("return; }\n")

		fmt.Fprintf(&buf, "func (%s *%s) read(rr io.Reader) (err error) { \n", short, p.name)
		r := &reading{
			base: short,
			out:  &buf,
		}
		r.readStruct(structs[p.name].Type.(*ast.StructType), short)
		r.flush()
		buf.WriteString("return; }\n")

		buf.WriteString("\n\n")
	}

	// Packet constructors
	buf.WriteString("func init() {\n")
	for _, p := range packets {
		fmt.Fprintf(&buf, "packetCreator[%s][%s][%d] = func () Packet { return &%s{} }\n", protocol, dir, p.id, p.name)
	}
	buf.WriteString("}\n")

	// Write the header last because of imports

	var header bytes.Buffer
	header.WriteString("// Generated by protocol_builder\n")
	header.WriteString("// Do not edit\n\n")
	fmt.Fprintf(&header, "package %s\n", parsedFile.Name)

	header.WriteString("import (")
	for impt := range imports {
		fmt.Fprintf(&header, "\"%s\"\n", impt)
	}
	header.WriteString(")\n")

	buf.WriteTo(&header)

	b, err := format.Source(header.Bytes())
	if err != nil {
		log.Println(header.String())
		log.Fatalf("format error: %s", err)
	}

	o, err := os.Create(input[:len(input)-len(filepath.Ext(input))] + "_proto.go")
	if err != nil {
		log.Fatalln(err)
	}
	defer o.Close()
	o.Write(b)
}
