package inner2

func ExternalMultiplication(values ... int) (returnValue int){
	for _, value := range values  {
		returnValue *= value
	}
	return
}