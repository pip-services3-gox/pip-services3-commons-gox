package run

// IExecutable interface for components that can be called to execute work.
//	Example:
//		type EchoComponent {}
//		...
//		func  (ec* EchoComponent) Execute(correlationId: string, args: Parameters) (result any, err error) {
//			return nil, result = args.getAsObject("message")
//		}
//		echo := EchoComponent{};
//		message = "Test";
//		res, err = echo.Execute("123", NewParametersFromTuples("message", message));
//		fmt.Println(res);
type IExecutable interface {
	// Execute component with arguments and receives execution result.
	//	Parameters:
	//		- correlationId string transaction id to trace execution through call chain.
	//		- args *Parameters execution arguments.
	//	Returns: any, error result or execution and error
	Execute(correlationId string, args *Parameters) (result any, err error)
}
