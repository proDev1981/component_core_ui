package post

import "app/core"

func New(args Args) core.Element {

	return core.Box(
		core.Args{
			Class: "parent",
		},
		core.Box(
			core.Args{
				Class: "avatar",
			},
			core.Img(
				core.Args{
					Alt: "avatar",
					Src: args.Image,
				}),
		),
		core.Box(
			core.Args{
				Class: "body",
			},
			core.Label(
				core.Args{
					Type:  "h3",
					Class: "title",
					Value: args.Title,
				}),
			core.Label(
				core.Args{
					Type:  "span",
					Class: "sms",
					Value: args.Content,
				}),
		),
		core.Styles("./components/post/go.css"),
	)
}
