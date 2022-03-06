package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

/*
let newAdder = fn(x) { fn(y) { x + y } };
let addTwo = newAdder(2);

newAdder는 고차 함수. 클로저는 그 함수가 정의된 환경을 담아내는 함수. 자신의 환경을 담고 움직이고 호출됐을 때 그 환경에 접근.
addTwo가 호출됐을 때 갖고 있는 바인딩이 addTwo를 클로저로 만든다. 클로저 addTwo는 정의될 당시의 환경에 접근할 수 있다.
"당시"는 newAdder함수 몸체의 마지막 행이 평가됐을 때. 마지막 행은 함수 리터럴. 함수 리터럴이 평가될 떄 objectFunction을 생성하고 현재 환경에 대한 참조를 .Env 필드에 저장.
나중에 addTwo의 몸체를 평가할 따 현재 환경에서 평가하지 않고, addTwo 함수가 가진 환경에서 평가하고 현재 환경을 Eval에 넘기는 대신, 함수가 가진 환경을 확장해서 Eval 함수에 넘김.
정의할 때 사용한 환경을 넘겨야 그 환경으로 정의한 이름에 접근할 수 있다.
*/
