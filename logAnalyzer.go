package chap4

type LogAnalyzer struct {
	service *WebService
}

func (la *LogAnalyzer) Analyze(filename string){
	if(len(filename) < 8){
		la.service.LogError("File name too short: " + filename)
	}
}

func NewLogAnalyzer() *LogAnalyzer {
	ws := &WebService{}
	
    return &LogAnalyzer{
        service: ws,
    }
}


