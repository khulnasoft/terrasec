package dockerv1

import (
	"bytes"
	"os"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"go.uber.org/zap"
)

// DockerConfig holds information about individual docker instructions
type DockerConfig struct {
	Cmd   string `json:"cmd"`
	Value string `json:"value"`
	Line  int    `json:"line"`
}

const (
	stringJoinCharacter = " "
	commentPrefix       = "#"
	newLine             = "\n"
)

// ValidateInstruction validates the dockerfile instructions
func (dc *DockerV1) ValidateInstruction(node *parser.Node) error {
	_, err := instructions.ParseInstruction(node)
	return err
}

// Parse parses the given dockerfile and gives docker config and string of comments present in dockerfile.
func (dc *DockerV1) Parse(filepath string) ([]DockerConfig, string, error) {
	config := []DockerConfig{}
	comments := ""

	data, err := os.ReadFile(filepath)
	if err != nil {
		zap.S().Error("error loading docker file", filepath, zap.Error(err))
		return config, "", err
	}
	r := bytes.NewReader(data)
	res, err := parser.Parse(r)
	if err != nil {
		zap.S().Errorf("error while parsing iac file", filepath, zap.Error(err))
		return config, "", err
	}

	for _, child := range res.AST.Children {
		values := []string{}
		err = dc.ValidateInstruction(child)
		if err != nil {
			return config, "", err
		}

		// loop over all the comments before the instruction is found to create one single string of comments
		// appending # prefix and new line since it is removed by the parser while creating the AST
		// Purpose of adding them back is to use the command function to find skiprules and min max severity.
		for _, comment := range child.PrevComment {
			comments = comments + commentPrefix + comment + newLine
		}

		values = append(values, child.Flags...)

		for i := child.Next; i != nil; i = i.Next {
			values = append(values, i.Value)
		}

		value := strings.Join(values, stringJoinCharacter)

		tempConfig := DockerConfig{
			Cmd:   child.Value,
			Value: value,
			Line:  child.StartLine,
		}
		config = append(config, tempConfig)
	}
	return config, comments, nil
}
