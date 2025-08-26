package dockerfile

import (
	"maps"
	"os"
	"slices"
	"testing"
)

func testDockerfileExists(t *testing.T, scoreDec func()) {
	// Check if Dockerfiles exist in the directories students create
	requiredFiles := map[string]string{
		"Task1_BasicWeb":   "task1/my-web-app/Dockerfile",
		"Task2_Advanced":   "task2/advanced-app/Dockerfile",
		"Task3_Optimized":  "task3/optimized-app/Dockerfile",
		"Task4_Go":         "task4/go-app/Dockerfile",
		"Task5_Production": "task5/production-app/Dockerfile",
	}

	for _, task := range slices.Sorted(maps.Keys(requiredFiles)) {
		t.Run(task, func(t *testing.T) {
			file := requiredFiles[task]
			if _, err := os.Stat(file); os.IsNotExist(err) {
				t.Errorf("Missing required file: %s", file)
				scoreDec()
			}
		})
	}
}
