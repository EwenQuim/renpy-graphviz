package parser

type situation string

const (
	situationPending situation = ""
	situationJump    situation = "jump"
	situationCall    situation = "call"
	situationLabel   situation = "label"
)
