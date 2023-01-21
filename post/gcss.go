package post

import "app/core/gcss"
import "app/core/color"

func style(colorTitle, colorSms, bgColorPost string) string {

	return gcss.Sheet{
		".post-parent": &gcss.Rule{
			BgColor:      bgColorPost,
			Display:      "flex",
			MinHeight:    "150px",
			Border:       "none",
			BorderRadius: "5px",
			Margin:       "5px",
			Hover: &gcss.Rule{
				BoxShadow: "0px 0px 5px 2px" + color.RGBA(0, 0, 0, 0.25),
			},
		},
		".post-avatar": &gcss.Rule{
			Display:        "flex",
			Width:          "20%",
			JustifyContent: "center",
			AlignItems:     "flex-start",
		},
		".post-avatar *": &gcss.Rule{
			Margin:       "10px",
			BorderRadius: "50%",
			Width:        "60px",
			Height:       "60px",
		},
		".post-body": &gcss.Rule{
			Display:       "flex",
			FlexDirection: "column",
			Width:         "80%",
		},
		".post-title": &gcss.Rule{
			Color:   colorTitle,
			Margin:  "0",
			Padding: "5px",
			Height:  "20%",
		},
		".post-sms": &gcss.Rule{
			Color:   colorSms,
			Margin:  "0",
			Padding: "5px 10px",
			Height:  "80%",
		},
	}.Parse()
}
