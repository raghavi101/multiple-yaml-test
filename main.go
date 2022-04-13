package amin

import (
	"github.com/bitfield/script"
)

func main() {

	for i := 1; i < 3; i++ {
		input := script.File("kustomization.yaml")
		input.Replace("kustomize{i}-", "kustomize{i+1}-")
		input.Replace("argo${i}.yaml", "argo${i+1}.yaml")
		script.Exec("kustomize build > ./argo${i+1}.yaml").Stdout()
	}

}
