package module

type SendMethod interface {
	Send(text []string) int
}

type SendModule struct {
	sendMethod SendMethod
	textsInfo  map[int]TextsInfo
}

func NewSendModule(sendMethod SendMethod) *SendModule {
	return &SendModule{
		sendMethod: sendMethod,
		textsInfo:  make(map[int]TextsInfo),
	}
}

func (s SendModule) Send(text []string) int {
	id := s.sendMethod.Send(text)

	info := TextsInfo{
		text: text,
	}

	s.textsInfo[id] = info
	return id

}

func (s SendModule) ExactMessageById(id int) TextsInfo {
	info, ok := s.textsInfo[id]
	if !ok {
		return TextsInfo{}
	}

	return info
}

func (s SendModule) AllInfo() map[int]TextsInfo {
	tempMap := make(map[int]TextsInfo, len(s.textsInfo))

	for k, v := range s.textsInfo {
		tempMap[k] = v
	}
	return tempMap
}
