package gen

type GenConfig struct {
	Key              []string
	Configs          []DataConfig
	PermutationRange [2]int
	URL              string
}

type DataConfig interface {
	Config() string
}

func GenData(config GenConfig) (interface{}, error) {

}
