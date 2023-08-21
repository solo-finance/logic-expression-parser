package lep

type Between struct {
	statement
	Start Value
	End   Value
}

var _ Expression = (*Between)(nil)
var _ Statement = (*Between)(nil)

func BetweenXandY(param *ParamX, start Value, end Value) *Between {
	return &Between{
		statement: statement{
			Param: param,
		},
		Start: start,
		End:   end,
	}
}

func (s Between) Equals(other Expression) bool {
	if expr, ok := other.(*Between); ok {
		return s.Param.Equals(expr.Param) && s.Start.Equals(expr.Start) && s.End.Equals(expr.End)
	}
	return false
}

func (s Between) String() string {
	return s.Param.String() + " between " + s.Start.String() + " and " + s.End.String()
}

func parseBetween(left, start, end interface{}) (*Between, error) {
	param, valuestart, err := parseStatement(left, start)
	if err != nil {
		return nil, err
	}
	_, valueend, err := parseStatement(left, end)
	if err != nil {
		return nil, err
	}

	return BetweenXandY(param, valuestart, valueend), nil
}
