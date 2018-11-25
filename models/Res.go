package models

type Res struct {
	Code         int  // -1   0   99
	Info         string
	Data         map[string] string
	SecretData   map[string] []Secret
}
