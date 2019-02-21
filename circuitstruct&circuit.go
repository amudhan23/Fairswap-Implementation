//these struct and SetCircuitTuples is common to both sender and receiver

type circuit struct{
	Index int
	Operation int
	InputToTheGate []int
}

//function to model the circuit given....
func setCircuitTuples(circuitobjects *[TotalNumberOfgates]circuit){

	//Giving tuple value to input gates
	for i:=0;i<NoOfInputGates;i++{
			circuitobjects[i].Index = i
			circuitobjects[i].Operation =1
			circuitobjects[i].InputToTheGate = nil
	}

	//Giving tuple values to !inputgates && !lastgate
	k:=0
	for i:=NoOfInputGates;i<TotalNumberOfgates-1;i++ {
			circuitobjects[i].Index=i
			circuitobjects[i].Operation = 2
			for j:=0;j<MaxLinesToGate;j++{
				circuitobjects[i].InputToTheGate=append(circuitobjects[i].InputToTheGate,k)
				k++
			}
	}

	//Last gate
	circuitobjects[TotalNumberOfgates-1].Index = TotalNumberOfgates-1
	circuitobjects[TotalNumberOfgates-1].Operation = 3
	circuitobjects[TotalNumberOfgates-1].InputToTheGate=append(circuitobjects[TotalNumberOfgates-1].InputToTheGate,TotalNumberOfgates-2)

}
