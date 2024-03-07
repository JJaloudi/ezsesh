package ezsesh

type EzStoreMethods interface {
	Generate(assoc string)
	/*
		Delete()

		OnGenerate()
		OnDelete()
		OnExpire()
	*/
}
