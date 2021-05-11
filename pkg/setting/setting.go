package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

var allSection = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v) // unmarshaling to a struct
	if err != nil {
		return err
	}

	if _, ok := allSection[k]; !ok { // 如果沒紀錄過的話，就把指標紀錄起來
		allSection[k] = v
	}

	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range allSection {
		err := s.vp.UnmarshalKey(k, v) // unmarshaling to a struct
		// fmt.Print(err)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Setting) WatchConfigChange() error {
	go func() {

		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			s.ReloadAllSection()
		})

	}()

	return nil
}
