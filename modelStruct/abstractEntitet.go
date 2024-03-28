package modelStruct

type NoviEntitet interface{
	GetId() int
}

type entitetBase struct{
	Id int
}

func (a entitetBase) GetId() int{
	return a.Id
}

func SetEntitet (id int) NoviEntitet{
	return entitetBase{Id: id}
}