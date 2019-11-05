package terraform

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/trawler/cna-installer/pkg/terraform/plugins"
)

// var plugins = map[string]*plugin.ServeOpts{
// 	"local":    {ProviderFunc: local.Provider},
// 	"template": {ProviderFunc: template.Provider},
// 	//"azurerm":  {ProviderFunc: azurerm.Provider},
// }
//
// var pluginsConfigTemplate = gtemplate.Must(gtemplate.New("").Parse(`
// providers {
// 	{{- range $name, $plugin := .Plugins }}
// 	{{ $name }} = "{{ $.BinaryPath }}-TFSPACE-{{ $name }}"
// 	{{- end }}
// }`))
//
// // ServePlugin serves every vendored TerraForm providers/provisioners. This
// // function never returns and should be the final function called in the main
// // function of the plugin. Additionally, there should be no stdout/stderr
// // outputs, which may interfere with the handshake and further communications.
// func ServePlugin(name string) {
// 	p := plugins[name]
// 	if p == nil {
// 		log.Fatalf("could not find plugin %q", name)
// 	}
//
// 	plugin.Serve(p)
// }
//
// // BuildPluginsConfig comment
// func BuildPluginsConfig() (string, error) {
// 	execPath, err := osext.Executable()
// 	if err != nil {
// 		return "", err
// 	}
//
// 	var buffer bytes.Buffer
// 	var data = struct {
// 		Plugins    map[string]*plugin.ServeOpts
// 		BinaryPath string
// 	}{plugins, execPath}
//
// 	err = pluginsConfigTemplate.Execute(&buffer, data)
// 	return buffer.String(), err
// }

// unpackAndInit unpacks the platform-specific Terraform modules into
// the given directory and then runs 'terraform init'.
func unpackAndInit(dir string, platform string) (err error) {
	err = unpack(dir, platform)
	if err != nil {
		return errors.Wrap(err, "failed to unpack Terraform modules")
	}

	if err := setupEmbeddedPlugins(dir); err != nil {
		return errors.Wrap(err, "failed to setup embedded Terraform plugins")
	}

	tDebug := &lineprinter.Trimmer{WrappedPrint: logrus.Debug}
	tError := &lineprinter.Trimmer{WrappedPrint: logrus.Error}
	lpDebug := &lineprinter.LinePrinter{Print: tDebug.Print}
	lpError := &lineprinter.LinePrinter{Print: tError.Print}
	defer lpDebug.Close()
	defer lpError.Close()

	args := []string{
		"-get-plugins=false",
	}
	args = append(args, dir)
	if exitCode := texec.Init(dir, args, lpDebug, lpError); exitCode != 0 {
		return errors.New("failed to initialize Terraform")
	}
	return nil
}

func setupEmbeddedPlugins(dir string) error {
	execPath, err := os.Executable()
	if err != nil {
		return errors.Wrap(err, "failed to find path for the executable")
	}

	pdir := filepath.Join(dir, "plugins")
	if err := os.MkdirAll(pdir, 0777); err != nil {
		return err
	}
	for name := range plugins.KnownPlugins {
		dst := filepath.Join(pdir, name)
		if runtime.GOOS == "windows" {
			dst = fmt.Sprintf("%s.exe", dst)
		}
		if _, err := os.Stat(dst); err == nil {
			// stat succeeded, the plugin already exists.
			continue
		}
		logrus.Debugf("Symlinking plugin %s src: %q dst: %q", name, execPath, dst)
		if err := os.Symlink(execPath, dst); err != nil {
			return err
		}
	}
	return nil
}
