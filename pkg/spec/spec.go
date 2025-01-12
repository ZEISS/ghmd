package spec

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/zeiss/pkg/filex"
	"gopkg.in/yaml.v3"
)

// Example returns an example of the Spec struct
func Example() *Spec {
	return &Spec{
		Version:     DefaultVersion,
		Name:        "example",
		Description: "This is an example of the specification",
	}
}

const (
	// DefaultVersion is the default version of the specification
	DefaultVersion = 1
	// DefaultFilename is the default filename of the specification
	DefaultFilename = ".ghmd.yaml"
)

var validate = validator.New()

// Spec is the specification of the repository
type Spec struct {
	// Version is the version of the specification
	Version int `yaml:"version" validate:"required"`
	// Name is a given name of the project or repository (optional)
	Name string `yaml:"name,omitempty"`
	// Description is a short description of the project or repository (optional)
	Description string `yaml:"description,omitempty"`
	// Header is the header of the markdown file
	Header string `yaml:"header"`
	// Footer is the footer of the markdown file
	Footer      string         `yaml:"footer"`
	Groups      []Group        `yaml:"groups"`
	Contents    ContentsConfig `yaml:"contents"`
	EntryConfig EntryConfig    `yaml:"entry"`
}

type ContentsConfig struct {
	Title string `yaml:"title"`
}

type EntryRequirementsConfig struct {
	Title string `yaml:"title"`
}

type EntryExampleConfig struct {
	Title string `yaml:"title"`
}

type EntryConfig struct {
	TitlePrefix  string                  `yaml:"title_prefix"`
	Back         string                  `yaml:"back"`
	Requirements EntryRequirementsConfig `yaml:"requirements"`
	Example      EntryExampleConfig      `yaml:"example"`
}

// Entry in page containing single post.
type Entry struct {
	Title             string   `yaml:"title"`
	Name              string   `yaml:"name"`
	URL               string   `yaml:"url"`
	Description       string   `yaml:"description"`
	Author            string   `yaml:"author"`
	ExampleImageURL   string   `yaml:"example_image_url"`
	ExampleContent    string   `yaml:"example_content"`
	ExampleContentExt string   `yaml:"example_content_ext"`
	ExampleOutput     string   `yaml:"example_output"`
	Requirements      []string `yaml:"requirements"`
	Commands          []string `yaml:"commands"`
}

// Group is named ordered collection of Entries.
type Group struct {
	Title   string    `yaml:"title"`
	Type    GroupType `yaml:"type"`
	Entries []Entry   `yaml:"entries"`
}

type GroupType string

// UnmarshalYAML unmarshals the YAML data into the Spec struct
func (s *Spec) UnmarshalYAML(data []byte) error {
	ss := struct {
		Version     int    `yaml:"version"`
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
	}{}

	if err := yaml.Unmarshal(data, &ss); err != nil {
		return errors.WithStack(err)
	}

	s.Version = ss.Version
	s.Name = ss.Name
	s.Description = ss.Description

	err := validate.Struct(s)
	if err != nil {
		return err
	}

	return err
}

// Write writes the specification to the given file.
func Write(s *Spec, file string, force bool) error {
	b, err := yaml.Marshal(s)
	if err != nil {
		return err
	}

	ok, _ := filex.FileExists(filepath.Clean(file))
	if ok && !force {
		return fmt.Errorf("%s already exists, use --force to overwrite", file)
	}

	f, err := os.Create(filepath.Clean(file))
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}
