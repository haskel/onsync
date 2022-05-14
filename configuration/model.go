package configuration

type Target struct {
	Host        string            `yaml:"host"`
	Path        string            `yaml:"path"`
	Credentials TargetCredentials `yaml:"credentials"`
}

type TargetCredentials struct {
	User    string `yaml:"user"`
	UseKey  bool   `yaml:"use_key"`
	KeyPath string `yaml:"key_path"`
	Port    uint   `yaml:"port"`
}

type DirectoriesSync struct {
	Excluded []string `yaml:"excluded"`
	Only     []string `yaml:"only"`
}

type Sync struct {
	Source      string          `yaml:"source"`
	Target      string          `yaml:"target"`
	Directories DirectoriesSync `yaml:"directories"`
}

type Config struct {
	Targets map[string]Target `yaml:"targets"`
	Syncs   map[string]Sync   `yaml:"syncs"`
}

const (
	Localhost = "local"
)
