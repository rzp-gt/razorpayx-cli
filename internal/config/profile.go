package config

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Profile handles all things related to managing the project specific configurations
type Profile struct {
	DeviceName     string
	ProfileName    string
	APIKey         string
	LiveModeAPIKey string
	TestModeAPIKey string
	DisplayName    string
}

// CreateProfile creates a profile when logging in
func (p *Profile) CreateProfile() error {
	writeErr := p.writeProfile(viper.GetViper())
	if writeErr != nil {
		return writeErr
	}

	return nil
}

func (p *Profile) writeProfile(runtimeViper *viper.Viper) error {
	profilesFile := viper.ConfigFileUsed()

	err := makePath(profilesFile)
	if err != nil {
		return err
	}

	if p.DeviceName != "" {
		runtimeViper.Set(p.GetConfigField("device_name"), strings.TrimSpace(p.DeviceName))
	}

	if p.LiveModeAPIKey != "" {
		runtimeViper.Set(p.GetConfigField("live_mode_api_key"), strings.TrimSpace(p.LiveModeAPIKey))
	}

	if p.TestModeAPIKey != "" {
		runtimeViper.Set(p.GetConfigField("test_mode_api_key"), strings.TrimSpace(p.TestModeAPIKey))
	}

	if p.DisplayName != "" {
		runtimeViper.Set(p.GetConfigField("display_name"), strings.TrimSpace(p.DisplayName))
	}

	runtimeViper.MergeInConfig()

	// Do this after we merge the old configs in
	if p.TestModeAPIKey != "" {
		runtimeViper = p.safeRemove(runtimeViper, "secret_key")
		runtimeViper = p.safeRemove(runtimeViper, "api_key")
	}

	runtimeViper.SetConfigFile(profilesFile)

	// Ensure we preserve the config file type
	runtimeViper.SetConfigType(filepath.Ext(profilesFile))

	err = runtimeViper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

func makePath(path string) error {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetConfigField returns the configuration field for the specific profile
func (p *Profile) GetConfigField(field string) string {
	return p.ProfileName + "." + field
}

func (p *Profile) safeRemove(v *viper.Viper, key string) *viper.Viper {
	if v.IsSet(p.GetConfigField(key)) {
		newViper, err := removeKey(v, p.GetConfigField(key))
		if err == nil {
			// I don't want to fail the entire login process on not being able to remove
			// the old secret_key field so if there's no error
			return newViper
		}
	}

	return v
}

// Temporary workaround until https://github.com/spf13/viper/pull/519 can remove a key from viper
func removeKey(v *viper.Viper, key string) (*viper.Viper, error) {
	configMap := v.AllSettings()
	path := strings.Split(key, ".")
	lastKey := strings.ToLower(path[len(path)-1])
	deepestMap := deepSearch(configMap, path[0:len(path)-1])
	delete(deepestMap, lastKey)

	buf := new(bytes.Buffer)

	encodeErr := toml.NewEncoder(buf).Encode(configMap)
	if encodeErr != nil {
		return nil, encodeErr
	}

	nv := viper.New()
	nv.SetConfigType("toml") // hint to viper that we've encoded the data as toml

	err := nv.ReadConfig(buf)
	if err != nil {
		return nil, err
	}

	return nv, nil
}

// taken from https://github.com/spf13/viper/blob/master/util.go#L199,
// we need this to delete configs, remove when viper supprts unset natively
func deepSearch(m map[string]interface{}, path []string) map[string]interface{} {
	for _, k := range path {
		m2, ok := m[k]
		if !ok {
			// intermediate key does not exist
			// => create it and continue from there
			m3 := make(map[string]interface{})
			m[k] = m3
			m = m3

			continue
		}

		m3, ok := m2.(map[string]interface{})
		if !ok {
			// intermediate key is a value
			// => replace with a new map
			m3 = make(map[string]interface{})
			m[k] = m3
		}

		// continue search from here
		m = m3
	}

	return m
}

// RegisterAlias registers an alias for a given key.
func (p *Profile) RegisterAlias(alias, key string) {
	viper.RegisterAlias(p.GetConfigField(alias), p.GetConfigField(key))
}

// WriteConfigField updates a configuration field and writes the updated
// configuration to disk.
func (p *Profile) WriteConfigField(field, value string) error {
	viper.Set(p.GetConfigField(field), value)
	return viper.WriteConfig()
}

// DeleteConfigField deletes a configuration field.
func (p *Profile) DeleteConfigField(field string) error {
	v, err := removeKey(viper.GetViper(), p.GetConfigField(field))
	if err != nil {
		return err
	}

	return p.writeProfile(v)
}

// EditConfig opens the configuration file in the default editor.
func (c *Config) EditConfig() error {
	var err error

	fmt.Println("Opening config file:", c.ProfilesFile)

	switch runtime.GOOS {
	case "darwin", "linux":
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "vi"
		}

		cmd := exec.Command(editor, c.ProfilesFile)
		// Some editors detect whether they have control of stdin/out and will
		// fail if they do not.
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		return cmd.Run()
	case "windows":
		// As far as I can tell, Windows doesn't have an easily accesible or
		// comparable option to $EDITOR, so default to notepad for now
		err = exec.Command("notepad", c.ProfilesFile).Run()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}

// PrintConfig outputs the contents of the configuration file.
func (c *Config) PrintConfig() error {
	if c.Profile.ProfileName == "default" {
		configFile, err := ioutil.ReadFile(c.ProfilesFile)
		if err != nil {
			return err
		}

		fmt.Print(string(configFile))
	} else {
		configs := viper.GetStringMapString(c.Profile.ProfileName)

		if len(configs) > 0 {
			fmt.Printf("[%s]\n", c.Profile.ProfileName)
			for field, value := range configs {
				fmt.Printf("  %s=%s\n", field, value)
			}
		}
	}

	return nil
}
