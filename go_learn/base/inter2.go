package bs

import (
	"github.com/pkg/errors"
)

var _ Study = (*study)(nil)

type study struct {
	Name     string
	UserName string
}

// Create implements Study
func (s *study) Create(str string) string {
	return s.Name + " c " + str + " " + s.UserName
}

// Delete implements Study
func (s *study) Delete(str string) string {
	return s.Name + " d " + str
}

// Read implements Study
func (s *study) Read(str string) string {
	return s.Name + " r " + str
}

// Update implements Study
func (s *study) Update(str string) string {
	return s.Name + " u " + str
}

type Study interface {
	Create(str string) string
	Update(str string) string
	Read(str string) string
	Delete(str string) string
}

func New(n string) (Study, error) {
	if n == "" {
		return nil, errors.New("name不能为空")
	}
	//实现所有接口
	return &study{
		Name:     n,
		UserName: "ck",
	}, nil
}
