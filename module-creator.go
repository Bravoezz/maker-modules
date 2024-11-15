package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"unicode/utf8"
)

type ModuleCreator struct {
	nameModule string
	rootPath   string
}

func NewModuleCreator(name, rootPath string) (*ModuleCreator, error) {
	instance := &ModuleCreator{name, rootPath}
	return instance, instance.createModule()
}

func (md ModuleCreator) createModule() error {
	err := os.Mkdir(path.Join(md.rootPath, md.nameModule), 0777)
	if err != nil {
		return err
	}
	return nil
}

func (md ModuleCreator) ExecAsync() []error {
	wg := new(sync.WaitGroup)
	makers := [](func() error){
		md.createRepository,
		md.createService,
		md.createController,
		md.createContract,
		md.createModel,
	}

	var errs []error
	for _, maker := range makers {
		wg.Add(1)
		go func(w *sync.WaitGroup, listErr *[]error) {
			defer w.Done()
			*listErr = append(*listErr, maker())
		}(wg, &errs)
	}

	wg.Wait()
	return errs
}

func (md ModuleCreator) createRepository() error {
	//dir: repositories fileName: test.repository.ts
	var (
		fileName    = "repositories"
		repoName    = fmt.Sprintf("%s.repository.ts", md.nameModule)
		nameMdUpper = md.capitalizeFirstLetter(md.nameModule)
	)
	contentRepo := fmt.Sprintf(`
import {I%sRepository} from '../contracts/%s.contract'

export class %sRepository implements I%sRepository {}
	`, nameMdUpper, md.nameModule, nameMdUpper, nameMdUpper)

	if err := os.Mkdir(path.Join(md.rootPath, md.nameModule, fileName), 0777); err != nil {
		return err
	}
	file, err := os.Create(path.Join(md.rootPath, md.nameModule, fileName, repoName))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(contentRepo)
	if err != nil {
		return err
	}
	return nil
}

func (md ModuleCreator) createService() error {
	//dir: services fileName: test.service.ts

	var (
		fileName    = "services"
		repoName    = fmt.Sprintf("%s.service.ts", md.nameModule)
		nameMdUpper = md.capitalizeFirstLetter(md.nameModule)
	)
	contentRepo := fmt.Sprintf(`
import {I%sService} from '../contracts/%s.contract'

export class %sService implements I%sService {}
	`, nameMdUpper, md.nameModule, nameMdUpper, nameMdUpper)

	if err := os.Mkdir(path.Join(md.rootPath, md.nameModule, fileName), 0777); err != nil {
		return err
	}
	file, err := os.Create(path.Join(md.rootPath, md.nameModule, fileName, repoName))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(contentRepo)
	if err != nil {
		return err
	}
	return nil
}

func (md ModuleCreator) createController() error {
	//dir: controllers fileName: test.controller.ts

	var (
		fileName    = "controllers"
		repoName    = fmt.Sprintf("%s.controller.ts", md.nameModule)
		nameMdUpper = md.capitalizeFirstLetter(md.nameModule)
	)
	contentRepo := fmt.Sprintf(`
export class %sController {}
	`, nameMdUpper)

	if err := os.Mkdir(path.Join(md.rootPath, md.nameModule, fileName), 0777); err != nil {
		return err
	}
	file, err := os.Create(path.Join(md.rootPath, md.nameModule, fileName, repoName))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(contentRepo)
	if err != nil {
		return err
	}
	return nil
}

func (md ModuleCreator) createContract() error {
	//dir: contracts fileName: test.contract.ts

	var (
		fileName    = "contracts"
		repoName    = fmt.Sprintf("%s.contract.ts", md.nameModule)
		nameMdUpper = md.capitalizeFirstLetter(md.nameModule)
	)
	contentRepo := fmt.Sprintf(`
export interface I%sRepository {}
export interface I%sService {}
	`, nameMdUpper, nameMdUpper)

	if err := os.Mkdir(path.Join(md.rootPath, md.nameModule, fileName), 0777); err != nil {
		return err
	}
	file, err := os.Create(path.Join(md.rootPath, md.nameModule, fileName, repoName))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(contentRepo)
	if err != nil {
		return err
	}
	return nil
}

func (md ModuleCreator) createModel() error {
	//dir: models fileName: test.model.ts

	var (
		fileName    = "models"
		repoName    = fmt.Sprintf("%s.model.ts", md.nameModule)
		nameMdUpper = md.capitalizeFirstLetter(md.nameModule)
	)
	contentRepo := fmt.Sprintf(`

export interface %s {}
	`, nameMdUpper)

	if err := os.Mkdir(path.Join(md.rootPath, md.nameModule, fileName), 0777); err != nil {
		return err
	}
	file, err := os.Create(path.Join(md.rootPath, md.nameModule, fileName, repoName))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(contentRepo)
	if err != nil {
		return err
	}
	return nil
}

func (md ModuleCreator) capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s // Cadena vacía, no hay nada que cambiar
	}
	r, n := utf8.DecodeRuneInString(s)
	return strings.ToUpper(string(r)) + strings.ToLower(s[n:])
}
