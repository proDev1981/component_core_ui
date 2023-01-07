package post

import . "app/core"

func Post() Element {
	return Box(Args{
		Class: "parent",
	},
		Box(Args{
			Class: "avatar",
		},
			Img(Args{
				Alt: "avatar",
			}),
		),
		Box(Args{
			Class: "body",
		},
			Label(Args{
				Type:  "h3",
				Class: "title",
				Value: "Title",
			}),
			Label(Args{
				Type:  "span",
				Class: "sms",
				Value: "Contenido",
			}),
		),
		Styles("./components/post/go.css"),
	)
}
