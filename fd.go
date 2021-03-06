package gologic

func Difference(a,b,c interface{}) Goal {
        return AddC(Constraint{
                func (s S) (S, ConstraintResult) {
                        xo := Project(a,s)
                        x, xok := xo.(int)
                        yo := Project(b,s)
                        y, yok := yo.(int)
                        zo := Project(c,s)
                        z, zok := zo.(int)

                        if xok && yok && zok && x - y == z {
                                return s,Yes
                        } else if xok && yok && zok && x - y != z {
                                return s,No
                        } else if xok && yok  {
                                ns,success := Unifi(c,x-y,s)
                                if success {
                                        return ns, Yes
                                } else {
                                        return ns, No
                                }
                        } else if xok && zok  {
                                ns,success := Unifi(b,x-z,s)
                                if success {
                                        return ns, Yes
                                } else {
                                        return ns, No
                                }
                        } else if yok && zok {
                                ns,success := Unifi(a,y+z,s)
                                if success {
                                        return ns, Yes
                                } else {
                                        return ns, No
                                }
                        } else {
                                return s, Maybe
                        }
        }})
}

func Sum(a,b,c interface{}) Goal {
	return Difference(c,b,a)
}

func Increasing(a,b interface{}) Goal {
        return AddC(Constraint{
                func (s S) (S, ConstraintResult) {
                        o := Project(a,s)
                        x, xok := o.(int)
                        if xok {
                                o = Project(b,s)
                                y, yok := o.(int)
                                if yok {
                                        if y > x {
                                                return s,Yes
                                        } else {
                                                return s,No
                                        }
                                } else {
                                        return s,Maybe
                                }
                        } else {
                                return s, Maybe
                        }
        }})
}

func Divide(a,b,c interface{}) Goal {
        return AddC(Constraint{
                func (s S) (S, ConstraintResult) {
                        xo := Project(a,s)
                        x, xok := xo.(int)
                        yo := Project(b,s)
                        y, yok := yo.(int)
                        zo := Project(c,s)
                        z, zok := zo.(int)
			
                        if xok && yok && zok && x / y == z {
                                return s,Yes
                        } else if xok && yok && zok && x / y != z {
                                return s,No
                        } else if xok && yok  {
                                ns,success := Unifi(c,x/y,s)
                                if success {
                                        return ns, Yes
                                } else {
                                        return ns, No
                                }
                        } else if xok && zok  {
                                ns,success := Unifi(b,x/z,s)
                                if success {
                                        return ns, Yes
                                } else {
                                        return ns, No
                                }
                        } else if yok && zok {
                                ns,success := Unifi(a,z*y,s)
                                if success {
                                        return ns, Yes
                                } else {
                                        return ns, No
                                }
                        } else {
                                return s, Maybe
                        }
        }})
}


func Mult(a,b,c interface{}) Goal {
        return AddC(Constraint{
                func (s S) (S, ConstraintResult) {
                        xo := Project(a,s)
                        x, xok := xo.(int)
                        yo := Project(b,s)
                        y, yok := yo.(int)
                        zo := Project(c,s)
                        z, zok := zo.(int)
			
                        if xok && yok && zok && x * y == z {
                                return s,Yes
                        } else if xok && yok && zok && x * y != z {
                                return s,No
                        } else if xok && yok  {
                                ns,success := Unifi(c,x*y,s)
                                if success {
                                        return ns, Yes
                                } else {
                                        return ns, No
                                }
                        } else if xok && zok  {
                                ns,success := Unifi(b,z/x,s)
                                if success {
                                        return ns, Yes
                                } else {
                                        return ns, No
                                }
                        } else if yok && zok {
                                ns,success := Unifi(a,z/y,s)
                                if success {
                                        return ns, Yes
                                } else {
                                        return ns, No
                                }
                        } else {
                                return s, Maybe
                        }
        }})
}

func Neq(a,b interface{}) Goal {
        return AddC(Constraint{
                func (s S) (S, ConstraintResult) {
			if a == b {
				return s, No
			} else {
				x := Project(a,s)
				y := Project(b,s)
				if x == y {
					return s, No
				} else {
					_,xok := x.(V)
					_,yok := y.(V)
					if yok || xok {
						return s, Maybe
					} else {
						return s, Yes
					}
				}
			}
		}})
}
