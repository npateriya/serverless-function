package connectors

// Part of this code is copied from github.com/docker-exec/dexec for
// customization

import (
	"fmt"
	"log"
	"os"

	"github.com/docker-exec/dexec/dexec"
	"github.com/docker-exec/dexec/util"
	"github.com/fsouza/go-dockerclient"
	"github.com/npateriya/serverless-agent/models"
	"github.com/npateriya/serverless-agent/utils"
)

func RunContainer(funcData *models.Function) int {
	if len(funcData.Type) > 0 && funcData.Type == models.FUNCTION_TYPE_URL {
		filepath, err := utils.DownloadFile(funcData.CacheDir, funcData.SourceURL, true)
		if err != nil {
			log.Fatal(err)
		}
		funcData.SourceFile = filepath
	}
	//fmt.Printf("%+v", funcData)
	return RunDexecContainer(funcData)
}

// RunDexecContainer runs an anonymous Docker container with a Docker Exec
// image, mounting the specified sources and includes and passing the
// list of sources and arguments to the entrypoint.
func RunDexecContainer(funcData *models.Function) int {

	// Removing clean image and update image option for now. Add back if needed
	// Ideally these need to be seperate functions
	updateImage := false

	client, err := docker.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	dexecImage, err := ImageFromOptions(funcData)
	if err != nil {
		log.Fatal(err)
	}

	dockerImage := fmt.Sprintf("%s:%s", dexecImage.Image, dexecImage.Version)

	if err = dexec.FetchImage(
		dexecImage.Image,
		dexecImage.Version,
		updateImage,
		client); err != nil {
		log.Fatal(err)
	}

	// TODO : Add check if SourceFile
	var sourceBasenames []string
	if len(funcData.SourceFile) > 0 {
		basename, _ := dexec.ExtractBasenameAndPermission(funcData.SourceFile)
		sourceBasenames = append(sourceBasenames, []string{basename}...)

	}

	entrypointArgs := util.JoinStringSlices(
		sourceBasenames,
		util.AddPrefix(funcData.BuildArgs, "-b"),
		util.AddPrefix(funcData.RunParams, "-a"),
	)

	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Config: &docker.Config{
			Image:     dockerImage,
			Cmd:       entrypointArgs,
			StdinOnce: true,
			OpenStdin: true,
		},

		HostConfig: &docker.HostConfig{
			Binds: dexec.BuildVolumeArgs(
				util.RetrievePath([]string{funcData.TargetDir}),
				append([]string{funcData.SourceFile}, funcData.IncludeDir...)),
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.RemoveContainer(docker.RemoveContainerOptions{
			ID: container.ID,
		}); err != nil {
			log.Fatal(err)
		}
	}()

	if err = client.StartContainer(container.ID, &docker.HostConfig{}); err != nil {
		log.Fatal(err)
	}

	go func() {
		if err = client.AttachToContainer(docker.AttachToContainerOptions{
			Container:   container.ID,
			InputStream: os.Stdin,
			Stream:      true,
			Stdin:       true,
		}); err != nil {
			log.Fatal(err)
		}
	}()

	code, err := client.WaitContainer(container.ID)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Logs(docker.LogsOptions{
		Container:    container.ID,
		Stdout:       true,
		Stderr:       true,
		OutputStream: os.Stdout,
		ErrorStream:  os.Stderr,
	})

	if err != nil {
		log.Fatal(err)
	}

	return code
}

// ImageFromOptions returns an image from a set of functionData.
func ImageFromOptions(funcData *models.Function) (image *dexec.Image, err error) {
	useExtension := len(funcData.SourceLang) == 1
	useImage := len(funcData.BaseImage) == 1

	if useStdin := len(funcData.SourceFile) == 0; useStdin {
		if useExtension {
			image, err = dexec.LookupImageByExtension(funcData.SourceLang)
		} else if useImage {
			overrideImage, err := dexec.LookupImageByOverride(funcData.BaseImage, "unknown")
			if err != nil {
				return nil, err
			}
			image, err = dexec.LookupImageByName(overrideImage.Image)
			image.Version = overrideImage.Version
		} else {
			err = fmt.Errorf("STDIN requested but no extension or image supplied")
		}
	} else {
		if extension := util.ExtractFileExtension(funcData.SourceFile); useExtension {
			image, err = dexec.LookupImageByExtension(funcData.SourceLang)
		} else if useImage {
			image, err = dexec.LookupImageByOverride(funcData.BaseImage, extension)
		} else {
			image, err = dexec.LookupImageByExtension(extension)
		}
	}
	return image, err
}
