package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configPathes ...string) (*Setting, error) {
	v := viper.New()

	//越前面加進去的越先採用
	for _, path := range configPathes {
		if path != "" {
			v.AddConfigPath(path)
		}
	}
	// v.AddConfigPath("configs/")
	// v.AddConfigPath("../configs/")
	// v.AddConfigPath("../../configs/")
	// v.AddConfigPath(".")

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp: v}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v) // unmarshaling to a struct
	return err
}
