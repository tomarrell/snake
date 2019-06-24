package main

type mAck struct {
	MType string `json:"type"`
	Data  string `json:"data"`
	Err   string `json:"err,omitempty"`
}

func newAckError(err string) mAck {
	return mAck{
		MType: "ack",
		Data:  "error",
		Err:   err,
	}
}

func newAckOk() mAck {
	return mAck{
		MType: "ack",
		Data:  "ok",
	}
}
