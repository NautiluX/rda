package project

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"gopkg.in/yaml.v2"
)

const (
	RegistryDir string = "registry"
	RenderDir   string = "rendered-registry"
)

var parsedProjects []Project

func ReadProjects() ([]Project, error) {
	if parsedProjects != nil {
		return parsedProjects, nil
	}
	parsedProjects = []Project{}
	err := filepath.Walk(RegistryDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			yamlContent, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read project file %s: %w", path, err)
			}
			project := Project{}
			err = yaml.Unmarshal(yamlContent, &project)
			if err != nil {
				return fmt.Errorf("failed to parse project yaml from file %s: %w", path, err)
			}
			parsedProjects = append(parsedProjects, project)
		}
		return nil
	})
	if err != nil {
		return []Project{}, fmt.Errorf("error processing registry: %w", err)
	}

	return parsedProjects, nil
}
func AddToRegistry(p Project) error {
	nextId, err := GetNextProjectID()
	if err != nil {
		return fmt.Errorf("couldn't get next project ID: %w", err)
	}
	fmt.Printf("Assigning project ID %s", nextId)
	p.ID = nextId
	yamlOutput, err := yaml.Marshal(p)
	if err != nil {
		return fmt.Errorf("couldn't marshal project: %w", err)
	}
	err = os.WriteFile(filepath.Join(RegistryDir, p.ID+".yaml"), yamlOutput, 0644)
	if err != nil {
		return fmt.Errorf("unable to write data into file: %w", err)
	}
	return nil
}

func RenderRegistry() error {
	index := []string{}
	projects, err := ReadProjects()
	if err != nil {
		return fmt.Errorf("failed to read projects from registry: %w", err)
	}

	for _, p := range projects {
		filename := fmt.Sprintf("%s %s.md", p.ID, p.Name)

		err = os.WriteFile(filepath.Join(RenderDir, filename), []byte(p.RenderMarkdown()), 0644)
		if err != nil {
			return fmt.Errorf("unable to write markdown to file %s: %w", filename, err)
		}

		index = append(index, fmt.Sprintf("[%s %s](/%s/%s)\n\n", p.ID, p.Name, RenderDir, url.QueryEscape(filename)))
	}

	err = RenderReadme(index)
	if err != nil {
		return fmt.Errorf("unable to render readme: %w", err)
	}
	return nil
}

func RenderReadme(index []string) error {
	readme, err := os.ReadFile("README_INPUT.md")
	if err != nil {
		return fmt.Errorf("unable to read file README_INPUT.md: %w", err)
	}

	readmeString := string(readme)
	for _, rdaLink := range index {
		readmeString += fmt.Sprintf("\n* %s\n", rdaLink)
	}

	err = os.WriteFile("README.md", []byte(readmeString), 0644)
	if err != nil {
		return fmt.Errorf("unable to write README.md: %w", err)
	}
	return nil
}

func GetNextProjectID() (string, error) {
	projects, err := ReadProjects()
	if err != nil {
		return "", fmt.Errorf("failed to read projects from registry: %w", err)
	}
	projectIdregex := "RDA([0-9]+)$"
	biggestNumericId := 0
	re, err := regexp.Compile(projectIdregex)
	if err != nil {
		return "", fmt.Errorf("failed to parse regex %s: %w", projectIdregex, err)
	}
	for _, project := range projects {
		projectIdResult := re.FindStringSubmatch(project.ID)
		if len(projectIdResult) < 2 {
			return "", fmt.Errorf("unexpected project ID, couldn't find numeric ID: %s", project.ID)
		}
		numericId, err := strconv.Atoi(projectIdResult[1])
		if err != nil {
			return "", fmt.Errorf("failed to parse numeric project ID in %s: %w", project.ID, err)
		}
		if biggestNumericId < numericId {
			biggestNumericId = numericId
		}
	}

	nextId := fmt.Sprintf("RDA%04d", biggestNumericId+1)
	return nextId, nil
}
