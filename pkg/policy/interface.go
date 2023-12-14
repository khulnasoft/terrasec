package policy

// Engine Policy Engine interface
type Engine interface {
	//Init method to initialize engine with policy path, and a pre load filter
	Init(string, PreLoadFilter) error
	Configure() error
	Evaluate(EngineInput, PreScanFilter) (EngineOutput, error)
	GetResults() EngineOutput
	Release() error
}

// FilterSpecification defines a function that
// RegoMetadata filter specifications should implement
type FilterSpecification interface {
	IsSatisfied(r *RegoMetadata) bool
}

// PreLoadFilter defines functions, that a pre load filter should implement
type PreLoadFilter interface {
	IsAllowed(r *RegoMetadata) bool
	IsFiltered(r *RegoMetadata) bool
}

// PreScanFilter defines function, that a pre scan filter should implement
type PreScanFilter interface {
	Filter(rmap map[string]*RegoData, input EngineInput) map[string]*RegoData
}
