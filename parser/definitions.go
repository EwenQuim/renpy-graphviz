package parser

type situation string

const (
	situationPending      situation = ""
	situationJump         situation = "jump"
	situationCall         situation = "call"
	situationLabel        situation = "label"
	situationFakeLabel    situation = "fakelabel"
	situationFakeJump     situation = "fakejump"
	situationScreen       situation = "screen"
	situationScreenSwitch situation = "screenswitch"
)
