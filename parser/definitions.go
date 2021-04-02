package parser

type situation string

const (
	situationJump    situation = "jump"
	situationCall    situation = "call"
	situationLabel   situation = "label"
	situationPending situation = ""
)
