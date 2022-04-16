package handle

import "basic_gin/model"

type MemberData struct {
	data []model.Member
}

func NewMember() *MemberData {
	return &MemberData{}
}

func (m *MemberData) AllData() []model.Member {
	return m.data
}

func (m *MemberData) AddData(d model.Member) model.Member {
	m.data = append(m.data, d)
	return d
}
