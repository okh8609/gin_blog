package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	v := viper.New()
	v.AddConfigPath("configs/")
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

