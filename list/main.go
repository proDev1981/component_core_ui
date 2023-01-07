package list

import . "app/core"

func ListDataBase(data *State) Element {

	return List(Args{State: data},
		Box(Args{Class: "item-list"},
			Label(Args{
				Type:  "span",
				Value: "{{.Age}}",
			}),
			Label(Args{
				Type:  "span",
				Value: "{{.Name}}",
			}),
		),
	)
}
