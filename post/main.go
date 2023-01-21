package post

import "app/core"

func New(args Args) core.Element {

	return core.Box(
		core.Args{
			Class: "post-parent",
		},
		core.Box(
			core.Args{
				Class: "post-avatar",
			},
			core.Img(
				core.Args{
					Alt: "avatar",
					Src: args.Image,
				}),
		),
		core.Box(
			core.Args{
				Class: "post-body",
			},
			core.Label(
				core.Args{
					Type:  "h3",
					Class: "post-title",
					Value: args.Title,
				}),
			core.Label(
				core.Args{
					Type:  "span",
					Class: "post-sms",
					Value: args.Content,
				}),
		),
		core.Styles(
			gcssPost(
				args.ColorTitle,
				args.ColorContent,
				args.BgBody,
			)),
	)
}
