package practice

type MyArry struct {
	Elemet []MyType
}

type MyType int

func (this *MyArry)Add(e MyType)  {
	this.Elemet = myAppend(this.Elemet, e)
}

func (this *MyArry)Size() int  {

	return len(this.Elemet)
}

func myAppend(slice []MyType, elems MyType)[]MyType  {

	var result []MyType
	//获取当前切片长度
	slen := len(slice)
	scap := cap(slice)
	//拿到增加后的长度
	i := slen + 1

	//判断当前slice 容量 能否满足增加后的长度
	if i<= scap{
		//容量满足可以直接赋值
		result=slice[:i]
	}else{
		rcap := i
		if rcap<len(slice)*2{
			rcap = len(slice)*2
		}
		result = make([]MyType,i,rcap)
	}

	result[len(slice)] = elems
	copy(result, slice)

	return result
}