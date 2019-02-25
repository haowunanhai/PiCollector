package log

type logLevel int

const (
	debugLevel logLevel = iota
	traceLevel
	infoLevel
	warningLevel
	errorLevel
	fatalLevel
)

func (l logLevel) String() string {
	switch l {
	case debugLevel:
		return "D"
	case traceLevel:
		return "T"
	case infoLevel:
		return "I"
	case warningLevel:
		return "W"
	case errorLevel:
		return "E"
	case fatalLevel:
		return "F"
	}
	return "N"
}
