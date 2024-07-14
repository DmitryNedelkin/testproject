package database

type Message struct {
	Index    int32  `json:"index"`
	Id       string `json:"id"`
	Guid     string `json:"guid"`
	IsActive bool   `json:"isActive"`
}

type Messages []Message

var (
	AllMessages = []Messages{}
)

func GetMessageFromDatabase(index int32) *Message {
	for i, message := range AllMessages {
		if message[i].Index == index {
			return &message[i]
		}
	}

	return nil
}

func PutMessagesToDatabase(messages []Message) {
	AllMessages = append(AllMessages, messages)
}
