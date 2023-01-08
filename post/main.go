package post

import "app/core"
import "fmt"

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
			core.Args{
				Src: "./components/post/go.css",
				Key: "post-default",
			}),
		core.Styles(
			core.Args{
				Value: fmt.Sprintf(
					"$.post-title{color:%s;} $.post-sms{color:%s;} ${background-color:%s;}",
					core.Default(args.TitleFg, "green"),
					core.Default(args.ContenFg, "white"),
					core.Default(args.BodyBg, "#484343"),
				),
			}),
	)
}
