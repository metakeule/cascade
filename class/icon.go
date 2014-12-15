package class

import (
	"github.com/metakeule/goh4"
)

func icon(name string) goh4.Class {
	return class("icon-" + name)
}

var (
	Icon                  = class("icon")
	IconCollapse          = icon("collapse")
	Icon16                = icon("16")
	Icon32                = icon("32")
	Icon64                = icon("64")
	Icon128               = icon("128")
	IconButton            = icon("button")
	IconBorder            = icon("border")
	Icon_Glass            = icon("glass")
	Icon_Music            = icon("music")
	Icon_Search           = icon("search")
	Icon_Envelope         = icon("envelope")
	Icon_Heart            = icon("heart")
	Icon_Star             = icon("star")
	Icon_StarEmpty        = icon("star-empty")
	Icon_User             = icon("user")
	Icon_Film             = icon("film")
	Icon_ThLarge          = icon("th-large")
	Icon_Th               = icon("th")
	Icon_ThList           = icon("th-list")
	Icon_Ok               = icon("ok")
	Icon_Remove           = icon("remove")
	Icon_ZoomIn           = icon("zoom-in")
	Icon_ZoomOut          = icon("zoom-out")
	Icon_Off              = icon("off")
	Icon_Signal           = icon("signal")
	Icon_Cog              = icon("cog")
	Icon_Trash            = icon("trash")
	Icon_Home             = icon("home")
	Icon_File             = icon("file")
	Icon_Time             = icon("time")
	Icon_Road             = icon("road")
	Icon_DownloadAlt      = icon("download-alt")
	Icon_Download         = icon("download")
	Icon_Upload           = icon("upload")
	Icon_Inbox            = icon("inbox")
	Icon_PlayCircle       = icon("play-circle")
	Icon_Repeat           = icon("repeat")
	Icon_Refresh          = icon("refresh")
	Icon_ListArt          = icon("list-alt")
	Icon_Lock             = icon("lock")
	Icon_Flag             = icon("flag")
	Icon_Headphones       = icon("headphones")
	Icon_VolumeOff        = icon("volume-off")
	Icon_VolumeDown       = icon("volume-down")
	Icon_VolumeUp         = icon("volume-up")
	Icon_QRcode           = icon("qrcode")
	Icon_Barcode          = icon("barcode")
	Icon_Tag              = icon("tag")
	Icon_Tags             = icon("tags")
	Icon_Book             = icon("book")
	Icon_Bookmark         = icon("bookmark")
	Icon_Print            = icon("print")
	Icon_Camera           = icon("camera")
	Icon_Font             = icon("font")
	Icon_Bold             = icon("bold")
	Icon_Italic           = icon("italic")
	Icon_TextHeight       = icon("text-height")
	Icon_TextWidth        = icon("text-width")
	Icon_AlignLeft        = icon("align-left")
	Icon_AlignCenter      = icon("align-center")
	Icon_AlignRight       = icon("align-right")
	Icon_AlignJustify     = icon("align-justify")
	Icon_List             = icon("list")
	Icon_IndentLeft       = icon("indent-left")
	Icon_IndentRight      = icon("indent-right")
	Icon_FacetimeVideo    = icon("facetime-video")
	Icon_Picture          = icon("picture")
	Icon_Pencil           = icon("pencil")
	Icon_MapMarker        = icon("map-marker")
	Icon_Adjust           = icon("adjust")
	Icon_Tint             = icon("tint")
	Icon_Edit             = icon("edit")
	Icon_Share            = icon("share")
	Icon_Check            = icon("check")
	Icon_Move             = icon("move")
	Icon_StepBackward     = icon("step-backward")
	Icon_FastBackward     = icon("fast-backward")
	Icon_Backward         = icon("backward")
	Icon_Play             = icon("play")
	Icon_Pause            = icon("pause")
	Icon_Stop             = icon("stop")
	Icon_Forward          = icon("forward")
	Icon_FastForward      = icon("fast-forward")
	Icon_StepForward      = icon("step-forward")
	Icon_Eject            = icon("eject")
	Icon_ChevronLeft      = icon("chevron-left")
	Icon_ChevronRight     = icon("chevron-right")
	Icon_PlusSign         = icon("plus-sign")
	Icon_MinusSign        = icon("minus-sign")
	Icon_RemoveSign       = icon("remove-sign")
	Icon_OkSign           = icon("ok-sign")
	Icon_QuestionSign     = icon("question-sign")
	Icon_InfoSign         = icon("info-sign")
	Icon_Screenshot       = icon("screenshot")
	Icon_RemoveCircle     = icon("remove-circle")
	Icon_OkCircle         = icon("ok-circle")
	Icon_BanCircle        = icon("ban-circle")
	Icon_ArrowLeft        = icon("arrow-left")
	Icon_ArrowRight       = icon("arrow-right")
	Icon_ArrowUp          = icon("arrow-up")
	Icon_ArrowDown        = icon("arrow-down")
	Icon_ShareAlt         = icon("share-alt")
	Icon_ResizeFull       = icon("resize-full")
	Icon_ResizeSmall      = icon("resize-small")
	Icon_Plus             = icon("plus")
	Icon_Minus            = icon("minus")
	Icon_Asterisk         = icon("asterisk")
	Icon_ExclamationSign  = icon("exclamation-sign")
	Icon_Gift             = icon("gift")
	Icon_Leaf             = icon("leaf")
	Icon_Fire             = icon("fire")
	Icon_EyeOpen          = icon("eye-open")
	Icon_EyeClose         = icon("eye-close")
	Icon_WarningSign      = icon("warning-sign")
	Icon_Plane            = icon("plane")
	Icon_Calendar         = icon("calendar")
	Icon_Random           = icon("random")
	Icon_Comment          = icon("comment")
	Icon_Magnet           = icon("magnet")
	Icon_ChevronUp        = icon("chevron-up")
	Icon_ChevronDown      = icon("chevron-down")
	Icon_Retweet          = icon("retweet")
	Icon_ShoppingCart     = icon("shopping-cart")
	Icon_FolderClose      = icon("folder-close")
	Icon_FolderOpen       = icon("folder-open")
	Icon_ResizeVertical   = icon("resize-vertical")
	Icon_ResizeHorizontal = icon("resize-horizontal")
	Icon_BarChart         = icon("bar-chart")
	Icon_TwitterSign      = icon("twitter-sign")
	Icon_FacebookSign     = icon("facebook-sign")
	Icon_CameraRetro      = icon("camera-retro")
	Icon_Key              = icon("key")
	Icon_Cogs             = icon("cogs")
	Icon_Comments         = icon("comments")
	Icon_ThumbsUp         = icon("thumbs-up")
	Icon_ThumbsDown       = icon("thumbs-down")
	Icon_StarHalf         = icon("star-half")
	Icon_HeartEmpty       = icon("heart-empty")
	Icon_Signout          = icon("signout")
	Icon_LinkedinSign     = icon("linkedin-sign")
	Icon_PushPin          = icon("pushpin")
	Icon_ExternalLink     = icon("external-link")
	Icon_SignIn           = icon("signin")
	Icon_Trophy           = icon("trophy")
	Icon_GithubSign       = icon("github-sign")
	Icon_UploadAlt        = icon("upload-alt")
	Icon_Lemon            = icon("lemon")
	Icon_Phone            = icon("phone")
	Icon_CheckEmpty       = icon("check-empty")
	Icon_BookmarkEmpty    = icon("bookmark-empty")
	Icon_PhoneSign        = icon("phone-sign")
	Icon_Twitter          = icon("twitter")
	Icon_Facebook         = icon("facebook")
	Icon_Github           = icon("github")
	Icon_Unlock           = icon("unlock")
	Icon_CreditCard       = icon("credit-card")
	Icon_Rss              = icon("rss")
	Icon_Hdd              = icon("hdd")
	Icon_Bullhorn         = icon("bullhorn")
	Icon_Bell             = icon("bell")
	Icon_Certificate      = icon("certificate")
	Icon_HandRight        = icon("hand-right")
	Icon_HandLeft         = icon("hand-left")
	Icon_HandUp           = icon("hand-up")
	Icon_HandDown         = icon("hand-down")
	Icon_CircleArrowLeft  = icon("circle-arrow-left")
	Icon_CircleArrowRight = icon("circle-arrow-right")
	Icon_CircleArrowUp    = icon("circle-arrow-up")
	Icon_CircleArrowDown  = icon("circle-arrow-down")
	Icon_Globe            = icon("globe")
	Icon_Wrench           = icon("wrench")
	Icon_Tasks            = icon("tasks")
	Icon_Filter           = icon("filter")
	Icon_Briefcase        = icon("briefcase")
	Icon_Fullscreen       = icon("fullscreen")
	Icon_Group            = icon("group")
	Icon_Link             = icon("link")
	Icon_Cloud            = icon("cloud")
	Icon_Beaker           = icon("beaker")
	Icon_Cut              = icon("cut")
	Icon_Copy             = icon("copy")
	Icon_PaperClip        = icon("paper-clip")
	Icon_Save             = icon("save")
	Icon_SignBlank        = icon("sign-blank")
	Icon_Reorder          = icon("reorder")
	Icon_ListUl           = icon("list-ul")
	Icon_ListOl           = icon("list-ol")
	Icon_StrikeThrough    = icon("strikethrough")
	Icon_Underline        = icon("underline")
	Icon_Table            = icon("table")
	Icon_Magic            = icon("magic")
	Icon_Truck            = icon("truck")
	Icon_Pinterest        = icon("pinterest")
	Icon_PinterestSign    = icon("pinterest-sign")
	Icon_GooglePlusSign   = icon("google-plus-sign")
	Icon_GooglePlus       = icon("google-plus")
	Icon_Money            = icon("money")
	Icon_CaretDown        = icon("caret-down")
	Icon_CaretUp          = icon("caret-up")
	Icon_CaretLeft        = icon("caret-left")
	Icon_CaretRight       = icon("caret-right")
	Icon_Columns          = icon("columns")
	Icon_Sort             = icon("sort")
	Icon_SortDown         = icon("sort-down")
	Icon_SortUp           = icon("sort-up")
	Icon_EnvelopeAlt      = icon("envelope-alt")
	Icon_Linkedin         = icon("linkedin")
	Icon_Undo             = icon("undo")
	Icon_Legal            = icon("legal")
	Icon_Dashboard        = icon("dashboard")
	Icon_CommentAlt       = icon("comment-alt")
	Icon_CommentsAlt      = icon("comments-alt")
	Icon_Bolt             = icon("bolt")
	Icon_Sitemap          = icon("sitemap")
	Icon_Umbrella         = icon("umbrella")
	Icon_Paste            = icon("paste")
	Icon_Lightbulb        = icon("lightbulb")
	Icon_Exchange         = icon("exchange")
	Icon_CloudDownload    = icon("cloud-download")
	Icon_CloudUpload      = icon("cloud-upload")
	Icon_UserMd           = icon("user-md")
	Icon_Stethoscope      = icon("stethoscope")
	Icon_Suitcase         = icon("suitcase")
	Icon_BellAlt          = icon("bell-alt")
	Icon_Coffee           = icon("coffee")
	Icon_Food             = icon("food")
	Icon_FileAlt          = icon("file-alt")
	Icon_Building         = icon("building")
	Icon_Hospital         = icon("hospital")
	Icon_Ambulance        = icon("ambulance")
	Icon_Medkit           = icon("medkit")
	Icon_FighterJet       = icon("fighter-jet")
	Icon_Beer             = icon("beer")
	Icon_HSign            = icon("h-sign")
	Icon_PlusSignAlt      = icon("plus-sign-alt")
	Icon_DoubleAngleLeft  = icon("double-angle-left")
	Icon_DoubleAngleRight = icon("double-angle-right")
	Icon_DoubleAngleUp    = icon("double-angle-up")
	Icon_DoubleAngleDown  = icon("double-angle-down")
	Icon_AngleLeft        = icon("angle-left")
	Icon_AngleRight       = icon("angle-right")
	Icon_AngleUp          = icon("angle-up")
	Icon_AngleDown        = icon("angle-down")
	Icon_Desktop          = icon("desktop")
	Icon_Laptop           = icon("laptop")
	Icon_Tablet           = icon("tablet")
	Icon_MobilePhone      = icon("mobile-phone")
	Icon_CircleBlank      = icon("circle-blank")
	Icon_QuoteLeft        = icon("quote-left")
	Icon_QuoteRight       = icon("quote-right")
	Icon_Spinner          = icon("spinner")
	Icon_Circle           = icon("circle")
	Icon_Reply            = icon("reply")
	Icon_GithubAlt        = icon("github-alt")
	Icon_FolderCloseAlt   = icon("folder-close-alt")
	Icon_FolderOpenAlt    = icon("folder-open-alt")
)