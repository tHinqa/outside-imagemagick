// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package types provides API structs, enums and types for
//imagemagick and wand packages.
package types

type Quantum uint16 // byte uint16 uint32

type MagickRealType float32 // float64

const (
	MaxNumberFonts = 11
	MaxNumberPens  = 11
	MaxTextExtent  = 4096
)

type (
	CompositeMethod      Fixme
	DIR                  Fixme
	FILE                 Fixme
	PixelIntensityMethod Fixme
	Time                 Fixme // time_t
)

// X
type (
	Atom              UnsignedLong
	Colormap          UnsignedLong
	Cursor            UnsignedLong
	Pixmap            UnsignedLong
	Status            int
	Visual            struct{}
	Window            UnsignedLong
	XClassHint        struct{}
	XColor            struct{}
	XErrorEvent       struct{}
	XEvent            struct{} // union
	XFontStruct       struct{}
	XImage            struct{}
	XPoint            struct{ x, y Short }
	XrmDatabase       *struct{}
	XSegment          struct{ x1, y1, x2, y2 Short }
	XStandardColormap struct{}
	XVisualInfo       struct{}
	XWMHints          struct{}
)

// Placeholder type to be rectified.
type (
	Fixme uintptr

	Char             int8
	Enum             int
	IndexPacket      Quantum
	Long             int    // TODO(t):long on gcc vs msc 32 vs 64
	MagickOffsetType int64  // TODO(t):size on gcc vs msc 32 vs 64
	MagickSizeType   uint64 // TODO(t):size on gcc vs msc 32 vs 64
	MagickStatusType uint
	Short            uint16 // TODO(t):short on gcc vs msc 32 vs 64
	Size             uint   // TODO(t):size_t on gcc vs msc 32 vs 64
	SSize            int    // TODO(t):ssize_t on gcc vs msc 32 vs 64
	UnsignedShort    uint16 // TODO(t):unsigned short on gcc vs msc 32 vs 64
	UnsignedLong     uint   // TODO(t):unsigned long on gcc vs msc 32 vs 64
	Void             struct{}
)

// Opaque types
type (
	Ascii85Info   struct{}
	BlobInfo      struct{}
	CacheView     struct{}
	FxInfo        struct{}
	GC            *struct{}
	ImageView     struct{}
	PixelView     struct{}
	RandomInfo    struct{}
	SemaphoreInfo struct{}
	ThresholdMap  struct{}
	WandView      struct{}
)

// 	RealPixelPacket struct{ Red, Green, Blue, Opacity MagickRealType }
// 	XImportInfo    struct{ Frame, Borders, Screen, Descend, Silent MagickBooleanType }

// )

// type AESInfo struct {
// 	Key                      *StringInfo
// 	Blocksize                uint
// 	EncipherKey, DecipherKey *uint
// 	Rounds, Timestamp        ssize_t
// 	Signature                size_t
// }

type AffineMatrix struct {
	Sx, Rx, Ry, Sy, Tx, Ty float64
}
type AlignType Enum

const (
	UndefinedAlign AlignType = iota
	LeftAlign
	CenterAlign
	RightAlign
)

type AlphaChannelType Enum

const (
	UndefinedAlphaChannel AlphaChannelType = iota
	ActivateAlphaChannel
	BackgroundAlphaChannel
	CopyAlphaChannel
	DeactivateAlphaChannel
	ExtractAlphaChannel
	OpaqueAlphaChannel
	ResetAlphaChannel
	SetAlphaChannel
	ShapeAlphaChannel
	TransparentAlphaChannel
	FlattenAlphaChannel
	RemoveAlphaChannel
)

type AnnotationStencil Enum

const (
	ForegroundStencil AnnotationStencil = iota
	BackgroundStencil
	OpaqueStencil
	TransparentStencil
)

type BlobMode Enum

const (
	UndefinedBlobMode BlobMode = iota
	ReadBlobMode
	ReadBinaryBlobMode
	WriteBlobMode
	WriteBinaryBlobMode
	AppendBlobMode
	AppendBinaryBlobMode
)

// type CacheInfo struct { // ok
// 	StorageClass             ClassType
// 	Colorspace               ColorspaceType
// 	Channels                 size_t
// 	Type                     CacheType
// 	Mode                     MapMode
// 	Mapped                   MagickBooleanType
// 	Columns, Rows            size_t
// 	Offset                   MagickOffsetType
// 	Length                   MagickSizeType
// 	VirtualPixelMethod       VirtualPixelMethod
// 	VirtualPixelColor        MagickPixelPacket
// 	NumberThreads            size_t
// 	NexusInfo                **NexusInfo
// 	Pixels                   *PixelPacket
// 	Indexes                  *IndexPacket
// 	ActiveIndexChannel       MagickBooleanType
// 	File                     int
// 	Filename                 [MaxTextExtent]char
// 	CacheFilename            [MaxTextExtent]char
// 	Methods                  CacheMethods
// 	RandomInfo               *RandomInfo
// 	NumberConnections        size_t
// 	ServerInfo               *void
// 	Synchronize, Debug       MagickBooleanType
// 	Id                       MagickThreadType
// 	ReferenceCount           ssize_t
// 	Semaphore, FileSemaphore *SemaphoreInfo
// 	Timestamp                time_t
// 	Signature                size_t
// }

type CacheMethods struct {
	// GetVirtualPixelHandler          GetVirtualPixelHandler
	// GetVirtualPixelsHandler         GetVirtualPixelsHandler
	// GetVirtualIndexesFromHandler    GetVirtualIndexesFromHandler
	// GetOneVirtualPixelFromHandler   GetOneVirtualPixelFromHandler
	// GetAuthenticPixelsHandler       GetAuthenticPixelsHandler
	// GetAuthenticIndexesFromHandler  GetAuthenticIndexesFromHandler
	// GetOneAuthenticPixelFromHandler GetOneAuthenticPixelFromHandler
	// GetAuthenticPixelsFromHandler   GetAuthenticPixelsFromHandler
	// QueueAuthenticPixelsHandler     QueueAuthenticPixelsHandler
	// SyncAuthenticPixelsHandler      SyncAuthenticPixelsHandler
	// DestroyPixelHandler             DestroyPixelHandler
}

type CacheType Enum

const (
	UndefinedCache CacheType = iota
	MemoryCache
	MapCache
	DiskCache
	PingCache
	DistributedCache
)

type ChannelStatistics struct {
	Depth Size
	Minima,
	Maxima,
	Sum,
	SumSquared,
	SumCubed,
	SumFourthPower,
	Mean,
	Variance,
	StandardDeviation,
	Kurtosis,
	Skewness float64
}
type ChannelFeatures struct {
	AngularSecondMoment,
	Contrast,
	Correlation,
	VarianceSumOfSquares,
	InverseDifferenceMoment,
	SumAverage,
	SumVariance,
	SumEntropy,
	Entropy,
	DifferenceVariance,
	DifferenceEntropy,
	MeasureOfCorrelation1,
	MeasureOfCorrelation2,
	MaximumCorrelationCoefficient [4]float64
}

type ChannelType Enum

const (
	RedChannel ChannelType = 1 << iota
	GreenChannel
	BlueChannel
	AlphaChannel
	_
	BlackChannel
	TrueAlphaChannel
	RGBChannels
	SyncChannels

	DefaultChannels = ((AllChannels | SyncChannels) & ^OpacityChannel)

	UndefinedChannel ChannelType = 0
	GrayChannel                  = RedChannel
	CyanChannel                  = RedChannel
	MagentaChannel               = GreenChannel
	YellowChannel                = BlueChannel
	OpacityChannel               = AlphaChannel
	matteChannel                 = AlphaChannel // Deprecated
	IndexChannel                 = BlackChannel
	GrayChannels                 = RGBChannels

	CompositeChannels ChannelType = 0x2F
	AllChannels       ChannelType = 0x7ffffff
)

type ChromaticityInfo struct {
	RedPrimary,
	GreenPrimary,
	BluePrimary,
	WhitePoint PrimaryInfo
}

type ClassType Enum

const (
	UndefinedClass ClassType = iota
	DirectClass
	PseudoClass
)

type ClipPathUnits Enum

const (
	UndefinedPathUnits ClipPathUnits = iota
	UserSpace
	UserSpaceOnUse
	ObjectBoundingBox
)

type CoderInfo struct {
	Path      *string
	Magick    *string
	Name      *string
	Exempt    MagickBooleanType
	Stealth   MagickBooleanType
	Previous  *CoderInfo
	next      *CoderInfo // Deprecated
	Signature Size
}
type ColorInfo struct {
	Path       *string
	Name       *string
	Compliance ComplianceType
	Color      MagickPixelPacket
	Exempt     MagickBooleanType
	Stealth    MagickBooleanType
	Previous   *ColorInfo
	next       *ColorInfo // Deprecated
	Signature  UnsignedLong
}
type ColorPacket struct {
	Pixel PixelPacket
	Index IndexPacket
	Count MagickSizeType
}
type ColorspaceType Enum

const (
	UndefinedColorspace ColorspaceType = iota
	RGBColorspace
	GRAYColorspace
	TransparentColorspace
	OHTAColorspace
	LabColorspace
	XYZColorspace
	YCbCrColorspace
	YCCColorspace
	YIQColorspace
	YPbPrColorspace
	YUVColorspace
	CMYKColorspace
	SRGBColorspace
	HSBColorspace
	HSLColorspace
	HWBColorspace
	Rec601LumaColorspace
	Rec601YCbCrColorspace
	Rec709LumaColorspace
	Rec709YCbCrColorspace
	LogColorspace
	CMYColorspace
	LuvColorspace
	HCLColorspace
	LCHColorspace
	LMSColorspace
	LCHabColorspace
	LCHuvColorspace
	scRGBColorspace
	HSIColorspace
	HSVColorspace
	HCLpColorspace
	YDbDrColorspace
)

type CommandOption Enum

const (
	MagickUndefinedOptions CommandOption = iota - 1
	MagickAlignOptions
	MagickAlphaOptions
	MagickBooleanOptions
	MagickCacheOptions
	MagickChannelOptions
	MagickClassOptions
	MagickClipPathOptions
	MagickCoderOptions
	MagickColorOptions
	MagickColorspaceOptions
	MagickCommandOptions
	MagickComposeOptions
	MagickCompressOptions
	MagickConfigureOptions
	MagickDataTypeOptions
	MagickDebugOptions
	MagickDecorateOptions
	MagickDelegateOptions
	MagickDirectionOptions
	MagickDisposeOptions
	MagickDistortOptions
	MagickDitherOptions
	MagickEndianOptions
	MagickEvaluateOptions
	MagickFillRuleOptions
	MagickFilterOptions
	MagickFontOptions
	MagickFontsOptions
	MagickFormatOptions
	MagickFunctionOptions
	MagickGravityOptions
	MagickIntentOptions
	MagickInterlaceOptions
	MagickInterpolateOptions
	MagickKernelOptions
	MagickLayerOptions
	MagickLineCapOptions
	MagickLineJoinOptions
	MagickListOptions
	MagickLocaleOptions
	MagickLogEventOptions
	MagickLogOptions
	MagickMagicOptions
	MagickMethodOptions
	MagickMetricOptions
	MagickMimeOptions
	MagickModeOptions
	MagickModuleOptions
	MagickMorphologyOptions
	MagickNoiseOptions
	MagickOrientationOptions
	MagickPixelIntensityOptions
	MagickPolicyOptions
	MagickPolicyDomainOptions
	MagickPolicyRightsOptions
	MagickPreviewOptions
	MagickPrimitiveOptions
	MagickQuantumFormatOptions
	MagickResolutionOptions
	MagickResourceOptions
	MagickSparseColorOptions
	MagickStatisticOptions
	MagickStorageOptions
	MagickStretchOptions
	MagickStyleOptions
	MagickThresholdOptions
	MagickTypeOptions
	MagickValidateOptions
	MagickVirtualPixelOptionsOptions
	MagickComplexOptions
	MagickIntensityOptions
)

type ComplianceType Enum

const (
	SVGCompliance ComplianceType = 1 << iota
	X11Compliance
	XPMCompliance
	AllCompliance       ComplianceType = 0x7FFFFFFF
	UndefinedCompliance ComplianceType = 0
	NoCompliance                       = UndefinedCompliance
)

type CompositeOperator Enum

const (
	UndefinedCompositeOp CompositeOperator = iota
	NoCompositeOp
	ModulusAddCompositeOp
	AtopCompositeOp
	BlendCompositeOp
	BumpmapCompositeOp
	ChangeMaskCompositeOp
	ClearCompositeOp
	ColorBurnCompositeOp
	ColorDodgeCompositeOp
	ColorizeCompositeOp
	CopyBlackCompositeOp
	CopyBlueCompositeOp
	CopyCompositeOp
	CopyCyanCompositeOp
	CopyGreenCompositeOp
	CopyMagentaCompositeOp
	CopyOpacityCompositeOp
	CopyRedCompositeOp
	CopyYellowCompositeOp
	DarkenCompositeOp
	DstAtopCompositeOp
	DstCompositeOp
	DstInCompositeOp
	DstOutCompositeOp
	DstOverCompositeOp
	DifferenceCompositeOp
	DisplaceCompositeOp
	DissolveCompositeOp
	ExclusionCompositeOp
	HardLightCompositeOp
	HueCompositeOp
	InCompositeOp
	LightenCompositeOp
	LinearLightCompositeOp
	LuminizeCompositeOp
	MinusDstCompositeOp
	ModulateCompositeOp
	MultiplyCompositeOp
	OutCompositeOp
	OverCompositeOp
	OverlayCompositeOp
	PlusCompositeOp
	ReplaceCompositeOp
	SaturateCompositeOp
	ScreenCompositeOp
	SoftLightCompositeOp
	SrcAtopCompositeOp
	SrcCompositeOp
	SrcInCompositeOp
	SrcOutCompositeOp
	SrcOverCompositeOp
	ModulusSubtractCompositeOp
	ThresholdCompositeOp
	XorCompositeOp
	DivideDstCompositeOp
	DistortCompositeOp
	BlurCompositeOp
	PegtopLightCompositeOp
	VividLightCompositeOp
	PinLightCompositeOp
	LinearDodgeCompositeOp
	LinearBurnCompositeOp
	MathematicsCompositeOp
	DivideSrcCompositeOp
	MinusSrcCompositeOp
	DarkenIntensityCompositeOp
	LightenIntensityCompositeOp
)

type CompressionType Enum

const (
	UndefinedCompression CompressionType = iota
	NoCompression
	BZipCompression
	DXT1Compression
	DXT3Compression
	DXT5Compression
	FaxCompression
	Group4Compression
	JPEGCompression
	JPEG2000Compression
	LosslessJPEGCompression
	LZWCompression
	RLECompression
	ZipCompression
	ZipSCompression
	PizCompression
	Pxr24Compression
	B44Compression
	B44ACompression
	LZMACompression
	JBIG1Compression
	JBIG2Compression
)

type ConfigureInfo struct {
	Path      *string
	Name      *string
	Value     *string
	Exempt    MagickBooleanType
	Stealth   MagickBooleanType
	Previous  *ConfigureInfo
	next      *ConfigureInfo // Deprecated
	Signature Size
}

// type DataType Enum

// const (
// 	UndefinedData DataType = iota
// 	StringData
// 	ByteData
// 	ShortData
// 	LongData
// )

type DecodeImageHandler func(*ImageInfo, *ExceptionInfo) *Image
type DecorationType Enum

const (
	UndefinedDecoration DecorationType = iota
	NoDecoration
	UnderlineDecoration
	OverlineDecoration
	LineThroughDecoration
)

type DirectionType Enum

const (
	UndefinedDirection DirectionType = iota
	RightToLeftDirection
	LeftToRightDirection
)

type DisposeType Enum

const (
	UnrecognizedDispose DisposeType = iota
	NoneDispose
	BackgroundDispose
	PreviousDispose
	UndefinedDispose = UnrecognizedDispose
)

type DistortImageMethod Enum

const (
	UndefinedDistortion DistortImageMethod = iota
	AffineDistortion
	AffineProjectionDistortion
	ScaleRotateTranslateDistortion
	PerspectiveDistortion
	PerspectiveProjectionDistortion
	BilinearForwardDistortion
	BilinearReverseDistortion
	PolynomialDistortion
	ArcDistortion
	PolarDistortion
	DePolarDistortion
	Cylinder2PlaneDistortion
	Plane2CylinderDistortion
	BarrelDistortion
	BarrelInverseDistortion
	ShepardsDistortion
	ResizeDistortion
	SentinelDistortion
	BilinearDistortion = BilinearForwardDistortion
)

type DitherMethod Enum

const (
	UndefinedDitherMethod DitherMethod = iota
	NoDitherMethod
	RiemersmaDitherMethod
	FloydSteinbergDitherMethod
)

type DrawInfo struct {
	Primitive        *string
	Geometry         *string
	Viewbox          RectangleInfo
	Affine           AffineMatrix
	Gravity          GravityType
	Fill             PixelPacket
	Stroke           PixelPacket
	StrokeWidth      float64
	Gradient         GradientInfo
	FillPattern      *Image
	Tile             *Image
	StrokePattern    *Image
	StrokeAntialias  MagickBooleanType
	TextAntialias    MagickBooleanType
	FillRule         FillRule
	Linecap          LineCap
	Linejoin         LineJoin
	Miterlimit       Size
	DashOffset       float64
	Decorate         DecorationType
	Compose          CompositeOperator
	Text             *string
	Face             Size
	Font             *string
	Metrics          *string
	Family           *string
	Style            StyleType
	Stretch          StretchType
	Weight           Size
	Encoding         *string
	Pointsize        float64
	Density          *string
	Align            AlignType
	Undercolor       PixelPacket
	BorderColor      PixelPacket
	ServerName       *string
	DashPattern      *float64
	ClipMask         *string
	Bounds           SegmentInfo
	ClipUnits        ClipPathUnits
	Opacity          Quantum
	Render           MagickBooleanType
	ElementReference ElementReference
	Debug            MagickBooleanType
	Signature        Size
	Kerning          float64
	InterwordSpacing float64
	InterlineSpacing float64
	Direction        DirectionType
}
type DuplexTransferImageViewMethod func(
	*ImageView, *ImageView, *ImageView, SSize, int,
	*Void) MagickBooleanType
type DuplexTransferPixelViewMethod func(
	*PixelView, *PixelView, *PixelView, *Void) bool
type DuplexTransferWandViewMethod func(
	*WandView, *WandView, *WandView, SSize, int, *Void) bool

// type ElementInfo struct {
// 	Cx    MagickRealType
// 	Cy    MagickRealType
// 	Major MagickRealType
// 	Minor MagickRealType
// 	Angle MagickRealType
// 	Value *void
// 	Next  *ElementInfo
// }
type ElementReference struct {
	Id        *string
	Type      ReferenceType
	Gradient  GradientInfo
	Signature Size
	Previous  *ElementReference
	Next      *ElementReference
}
type ElementType Enum

const (
	UndefinedElement ElementType = iota
	PointElement
	LineElement
	RectangleElement
	FillRectangleElement
	CircleElement
	FillCircleElement
	EllipseElement
	FillEllipseElement
	PolygonElement
	FillPolygonElement
	ColorElement
	MatteElement
	TextElement
	ImageElement
)

type EncodeImageHandler func(*ImageInfo, *Image) bool
type EndianType Enum

const (
	UndefinedEndian EndianType = iota
	LSBEndian
	MSBEndian
)

type ErrorInfo struct {
	MeanErrorPerPixel,
	NormalizedMeanError,
	NormalizedMaximumError float64
}
type ExceptionInfo struct {
	Severity    ExceptionType
	ErrorNumber int
	Reason      *string
	Description *string
	Exceptions  *Void
	Relinquish  MagickBooleanType
	Semaphore   *SemaphoreInfo
	Signature   Size
}

type ExceptionType Enum

const (
	UndefinedException        ExceptionType = 0
	WarningException          ExceptionType = 300
	ResourceLimitWarning      ExceptionType = 300
	TypeWarning               ExceptionType = 305
	OptionWarning             ExceptionType = 310
	DelegateWarning           ExceptionType = 315
	MissingDelegateWarning    ExceptionType = 320
	CorruptImageWarning       ExceptionType = 325
	FileOpenWarning           ExceptionType = 330
	BlobWarning               ExceptionType = 335
	StreamWarning             ExceptionType = 340
	CacheWarning              ExceptionType = 345
	CoderWarning              ExceptionType = 350
	FilterWarning             ExceptionType = 352
	ModuleWarning             ExceptionType = 355
	DrawWarning               ExceptionType = 360
	ImageWarning              ExceptionType = 365
	WandWarning               ExceptionType = 370
	RandomWarning             ExceptionType = 375
	XServerWarning            ExceptionType = 380
	MonitorWarning            ExceptionType = 385
	RegistryWarning           ExceptionType = 390
	ConfigureWarning          ExceptionType = 395
	PolicyWarning             ExceptionType = 399
	ErrorException            ExceptionType = 400
	ResourceLimitError        ExceptionType = 400
	TypeError                 ExceptionType = 405
	OptionError               ExceptionType = 410
	DelegateError             ExceptionType = 415
	MissingDelegateError      ExceptionType = 420
	CorruptImageError         ExceptionType = 425
	FileOpenError             ExceptionType = 430
	BlobError                 ExceptionType = 435
	StreamError               ExceptionType = 440
	CacheError                ExceptionType = 445
	CoderError                ExceptionType = 450
	FilterError               ExceptionType = 452
	ModuleError               ExceptionType = 455
	DrawError                 ExceptionType = 460
	ImageError                ExceptionType = 465
	WandError                 ExceptionType = 470
	RandomError               ExceptionType = 475
	XServerError              ExceptionType = 480
	MonitorError              ExceptionType = 485
	RegistryError             ExceptionType = 490
	ConfigureError            ExceptionType = 495
	PolicyError               ExceptionType = 499
	FatalErrorException       ExceptionType = 700
	ResourceLimitFatalError   ExceptionType = 700
	TypeFatalError            ExceptionType = 705
	OptionFatalError          ExceptionType = 710
	DelegateFatalError        ExceptionType = 715
	MissingDelegateFatalError ExceptionType = 720
	CorruptImageFatalError    ExceptionType = 725
	FileOpenFatalError        ExceptionType = 730
	BlobFatalError            ExceptionType = 735
	StreamFatalError          ExceptionType = 740
	CacheFatalError           ExceptionType = 745
	CoderFatalError           ExceptionType = 750
	FilterFatalError          ExceptionType = 752
	ModuleFatalError          ExceptionType = 755
	DrawFatalError            ExceptionType = 760
	ImageFatalError           ExceptionType = 765
	WandFatalError            ExceptionType = 770
	RandomFatalError          ExceptionType = 775
	XServerFatalError         ExceptionType = 780
	MonitorFatalError         ExceptionType = 785
	RegistryFatalError        ExceptionType = 790
	ConfigureFatalError       ExceptionType = 795
	PolicyFatalError          ExceptionType = 799
)

type FillRule Enum

const (
	UndefinedRule FillRule = iota
	EvenOddRule
	NonZeroRule
)

type FilterTypes Enum

const (
	UndefinedFilter FilterTypes = iota
	PointFilter
	BoxFilter
	TriangleFilter
	HermiteFilter
	HanningFilter
	HammingFilter
	BlackmanFilter
	GaussianFilter
	QuadraticFilter
	CubicFilter
	CatromFilter
	MitchellFilter
	JincFilter
	SincFilter
	SincFastFilter
	KaiserFilter
	WelshFilter
	ParzenFilter
	BohmanFilter
	BartlettFilter
	LagrangeFilter
	LanczosFilter
	LanczosSharpFilter
	Lanczos2Filter
	Lanczos2SharpFilter
	RobidouxFilter
	RobidouxSharpFilter
	CosineFilter
	SplineFilter
	LanczosRadiusFilter
	SentinelFilter
)

type FrameInfo struct {
	Width, Height                Size
	X, Y, InnerBevel, OuterBevel SSize
}
type GeometryInfo struct{ Rho, Sigma, Xi, Psi, Chi float64 }
type GetImageViewMethod func(*ImageView, SSize, int, *Void) MagickBooleanType
type GetPixelViewMethod func(*PixelView, *Void) bool
type GetWandViewMethod func(*WandView, SSize, int, *Void) bool

type GhostInfo struct {
	// exit            func(*gs_main_instance) int
	// init_with_args  func(*gs_main_instance, int, []string) int
	// new_instance    func(**gs_main_instance, *Void) int
	// run_string      func(*gs_main_instance, string, int, *int) int
	// delete_instance func(*gs_main_instance)
}

type GradientInfo struct {
	Type           GradientType
	BoundingBox    RectangleInfo
	GradientVector SegmentInfo
	Stops          *StopInfo
	NumberStops    Size
	Spread         SpreadMethod
	Debug          MagickBooleanType
	Signature      Size
	Center         PointInfo
	Radius         MagickRealType
}

type GradientType Enum

const (
	UndefinedGradient GradientType = iota
	LinearGradient
	RadialGradient
)

type GravityType Enum

const (
	UndefinedGravity GravityType = iota
	NorthWestGravity
	NorthGravity
	NorthEastGravity
	WestGravity
	CenterGravity
	EastGravity
	SouthWestGravity
	SouthGravity
	SouthEastGravity
	StaticGravity
	ForgetGravity = UndefinedGravity
)

type Image struct {
	StorageClass           ClassType
	Colorspace             ColorspaceType
	Compression            CompressionType
	Quality                Size
	Orientation            OrientationType
	Taint_                 MagickBooleanType
	Matte                  MagickBooleanType
	Columns                Size
	Rows                   Size
	Depth_                 Size
	Colors                 Size
	Colormap               *PixelPacket
	BackgroundColor        PixelPacket
	BorderColor            PixelPacket
	MatteColor             PixelPacket
	Gamma_                 float64
	Chromaticity           ChromaticityInfo
	RenderingIntent        RenderingIntent
	Profiles               *Void
	Units                  ResolutionType
	Montage                *string
	Directory              *string
	Geometry_              *string
	Offset                 SSize
	XResolution            float64
	YResolution            float64
	Page                   RectangleInfo
	ExtractInfo            RectangleInfo
	TileInfo               RectangleInfo
	Bias                   float64
	Blur_                  float64
	Fuzz                   float64
	Filter_                FilterTypes
	Interlace              InterlaceType
	Endian                 EndianType
	Gravity                GravityType
	Compose                CompositeOperator
	Dispose_               DisposeType
	ClipMask_              *Image
	Scene                  Size
	Delay                  Size
	TicksPerSecond         SSize
	Iterations             Size
	TotalColors            Size
	StartLoop              SSize
	Error                  ErrorInfo
	Timer                  TimerInfo
	ProgressMonitor        MagickProgressMonitor
	ClientData             *Void
	Cache                  *Void
	Attributes             *Void
	Ascii85                *Ascii85Info
	Blob                   *BlobInfo
	Filename               [MaxTextExtent]Char
	MagickFilename         [MaxTextExtent]Char
	Magick                 [MaxTextExtent]Char
	MagickColumns          Size
	MagickRows             Size
	Exception_             ExceptionInfo
	Debug                  MagickBooleanType
	ReferenceCount         SSize
	Semaphore              *SemaphoreInfo
	ColorProfile           ProfileInfo
	IptcProfile            ProfileInfo
	GenericProfile         *ProfileInfo
	GenericProfiles        Size
	Signature_             Size
	Previous               *Image
	List                   *Image
	Next                   *Image
	Interpolate            InterpolatePixelMethod
	BlackPointCompensation MagickBooleanType
	TransparentColor       PixelPacket
	Mask_                  *Image
	TileOffset             RectangleInfo
	Properties             *Void
	Artifacts              *Void
	Type_                  ImageType
	Dither                 MagickBooleanType
	Extent_                MagickSizeType
	Ping                   MagickBooleanType
	Channels_              Size
	Timestamp              Time
	Intensity              PixelIntensityMethod
}
type ImageAttribute struct {
	Key         *string
	Value       *string
	Compression MagickBooleanType
	Previous    *ImageAttribute
	next        *ImageAttribute // Deprecated
}
type ImageInfo struct {
	Compression        CompressionType
	Orientation        OrientationType
	Temporary          MagickBooleanType
	Adjoin             MagickBooleanType
	Affirm             MagickBooleanType
	Antialias          MagickBooleanType
	Size               *string
	Extract            *string
	Page               *string
	Scenes             *string
	Scene              Size
	NumberScenes       Size
	Depth              Size
	Interlace          InterlaceType
	Endian             EndianType
	Units              ResolutionType
	Quality            Size
	SamplingFactor     *string
	ServerName         *string
	Font               *string
	Texture            *string
	Density            *string
	Pointsize          float64
	Fuzz               float64
	BackgroundColor    PixelPacket
	BorderColor        PixelPacket
	MatteColor         PixelPacket
	Dither             MagickBooleanType
	Monochrome         MagickBooleanType
	Colors             Size
	Colorspace         ColorspaceType
	Type               ImageType
	PreviewType        PreviewType
	Group              SSize
	Ping               MagickBooleanType
	Verbose            MagickBooleanType
	View               *string
	Authenticate       *string
	Channel            ChannelType
	Attributes         *Image
	Options            *Void
	ProgressMonitor    MagickProgressMonitor
	ClientData         *Void
	Cache              *Void
	Stream             StreamHandler
	File               *FILE
	Blob               *Void
	Length             Size
	Magick             [MaxTextExtent]Char
	Unique             [MaxTextExtent]Char
	Zero               [MaxTextExtent]Char
	Filename           [MaxTextExtent]Char
	Debug              MagickBooleanType
	Tile               *string
	Subimage           Size
	Subrange           Size
	Pen                PixelPacket
	Signature          Size
	VirtualPixelMethod VirtualPixelMethod
	TransparentColor   PixelPacket
	Profile            *Void
	Synchronize        MagickBooleanType
}
type ImageLayerMethod Enum

const (
	UndefinedLayer ImageLayerMethod = iota
	CoalesceLayer
	CompareAnyLayer
	CompareClearLayer
	CompareOverlayLayer
	DisposeLayer
	OptimizeLayer
	OptimizeImageLayer
	OptimizePlusLayer
	OptimizeTransLayer
	RemoveDupsLayer
	RemoveZeroLayer
	CompositeLayer
	MergeLayer
	FlattenLayer
	MosaicLayer
	TrimBoundsLayer
)

type ImageType Enum

const (
	UndefinedType ImageType = iota
	BilevelType
	GrayscaleType
	GrayscaleMatteType
	PaletteType
	PaletteMatteType
	TrueColorType
	TrueColorMatteType
	ColorSeparationType
	ColorSeparationMatteType
	OptimizeType
	PaletteBilevelMatteType
)

type InterlaceType Enum

const (
	UndefinedInterlace InterlaceType = iota
	NoInterlace
	LineInterlace
	PlaneInterlace
	PartitionInterlace
	GIFInterlace
	JPEGInterlace
	PNGInterlace
)

type InterpolatePixelMethod Enum

const (
	UndefinedInterpolatePixel InterpolatePixelMethod = iota
	AverageInterpolatePixel
	BicubicInterpolatePixel
	BilinearInterpolatePixel
	FilterInterpolatePixel
	IntegerInterpolatePixel
	MeshInterpolatePixel
	NearestNeighborInterpolatePixel
	SplineInterpolatePixel
	Average9InterpolatePixel
	Average16InterpolatePixel
	BlendInterpolatePixel
	BackgroundInterpolatePixel
	CatromInterpolatePixel
)

type KernelInfo struct {
	Type          KernelInfoType
	Width, Height Size
	X, Y          SSize
	Values        *float64
	Minimum,
	Maximum,
	NegativeRange,
	PositiveRange,
	Angle float64
	Next      *KernelInfo
	Signature Size
}
type KernelInfoType Enum

const (
	UndefinedKernel KernelInfoType = iota
	UnityKernel
	GaussianKernel
	DoGKernel
	LoGKernel
	BlurKernel
	CometKernel
	LaplacianKernel
	SobelKernel
	FreiChenKernel
	RobertsKernel
	PrewittKernel
	CompassKernel
	KirschKernel
	DiamondKernel
	SquareKernel
	RectangleKernel
	OctagonKernel
	DiskKernel
	PlusKernel
	CrossKernel
	RingKernel
	PeaksKernel
	EdgesKernel
	CornersKernel
	DiagonalsKernel
	LineEndsKernel
	LineJunctionsKernel
	RidgesKernel
	ConvexHullKernel
	ThinSEKernel
	SkeletonKernel
	ChebyshevKernel
	ManhattanKernel
	OctagonalKernel
	EuclideanKernel
	UserDefinedKernel
	BinomialKernel
)

type LineCap Enum

const (
	UndefinedCap LineCap = iota
	ButtCap
	RoundCap
	SquareCap
)

type LineJoin Enum

const (
	UndefinedJoin LineJoin = iota
	MiterJoin
	RoundJoin
	BevelJoin
)

type LocaleInfo struct {
	Path      *string
	Tag       *string
	Message   *string
	Stealth   MagickBooleanType
	Previous  *LocaleInfo
	next      *LocaleInfo // Deprecated
	Signature Size
}

type LogEventType Enum

const (
	TraceEvent LogEventType = 1 << iota
	AnnotateEvent
	BlobEvent
	CacheEvent
	CoderEvent
	ConfigureEvent
	DeprecateEvent
	DrawEvent
	ExceptionEvent
	ImageEvent
	LocaleEvent
	ModuleEvent
	PolicyEvent
	ResourceEvent
	TransformEvent
	_ // See UserEvent
	WandEvent
	X11Event
	AccelerateEvent
	UndefinedEvents LogEventType = 0
	NoEvents                     = UndefinedEvents
	UserEvent       LogEventType = 0x9000 // Weird
	AllEvents       LogEventType = 0x7FFFFFFF
)

// type LogHandlerType Enum

// const (
// 	UndefinedHandler LogHandlerType = iota
// 	NoHandler
// 	ConsoleHandler
// 	StdoutHandler
// 	StderrHandler
// 	FileHandler
// 	DebugHandler
// 	EventHandler
// )

// Changed to bool for uses other than in structs
type MagickBooleanType Enum

const (
	MagickFalse MagickBooleanType = iota
	MagickTrue
)

type MagickCommand func(*ImageInfo, int, []string, []string, *ExceptionInfo) bool

type MagickFunction Enum

const (
	UndefinedFunction MagickFunction = iota
	PolynomialFunction
	SinusoidFunction
	ArcsinFunction
	ArctanFunction
)

type MagicInfo struct {
	Path      *string
	Name      *string
	Target    *string
	Magic     *byte
	Length    Size
	Offset    MagickOffsetType
	Exempt    MagickBooleanType
	Stealth   MagickBooleanType
	Previous  *MagicInfo
	next      *MagicInfo // Deprecated
	Signature Size
}

type MagickEvaluateOperator Enum

const (
	UndefinedEvaluateOperator MagickEvaluateOperator = iota
	AddEvaluateOperator
	AndEvaluateOperator
	DivideEvaluateOperator
	LeftShiftEvaluateOperator
	MaxEvaluateOperator
	MinEvaluateOperator
	MultiplyEvaluateOperator
	OrEvaluateOperator
	RightShiftEvaluateOperator
	SetEvaluateOperator
	SubtractEvaluateOperator
	XorEvaluateOperator
	PowEvaluateOperator
	LogEvaluateOperator
	ThresholdEvaluateOperator
	ThresholdBlackEvaluateOperator
	ThresholdWhiteEvaluateOperator
	GaussianNoiseEvaluateOperator
	ImpulseNoiseEvaluateOperator
	LaplacianNoiseEvaluateOperator
	MultiplicativeNoiseEvaluateOperator
	PoissonNoiseEvaluateOperator
	UniformNoiseEvaluateOperator
	CosineEvaluateOperator
	SineEvaluateOperator
	AddModulusEvaluateOperator
	MeanEvaluateOperator
	AbsEvaluateOperator
	ExponentialEvaluateOperator
	MedianEvaluateOperator
	SumEvaluateOperator
)

// type MagickModuleType Enum

// const (
// 	MagickImageCoderModule MagickModuleType = iota
// 	MagickImageFilterModule
// )

type MagickPixelPacket struct {
	StorageClass ClassType
	Colorspace   ColorspaceType
	Matte        MagickBooleanType
	Fuzz         float64
	Depth        Size
	Red          MagickRealType
	Green        MagickRealType
	Blue         MagickRealType
	Opacity      MagickRealType
	Index        MagickRealType
}

type MagickProgressMonitor func(string, MagickOffsetType, MagickSizeType, *Void) bool

// const MagickSignatureSize = 64

type MapMode Enum

const (
	ReadMode MapMode = iota
	WriteMode
	IOMode
)

type MetricType Enum

const (
	UndefinedMetric MetricType = iota
	AbsoluteErrorMetric
	MeanAbsoluteErrorMetric
	MeanErrorPerPixelMetric
	MeanSquaredErrorMetric
	PeakAbsoluteErrorMetric
	PeakSignalToNoiseRatioMetric
	RootMeanSquaredErrorMetric
	NormalizedCrossCorrelationErrorMetric
	FuzzErrorMetric
	UndefinedErrorMetric = UndefinedMetric
)

type ModuleInfo struct {
	Path             *string
	Tag              *string
	Handle           *Void
	UnregisterModule func()
	RegisterModule   func() Size
	LoadTime         Time
	Stealth          MagickBooleanType
	Previous         *ModuleInfo
	next             *ModuleInfo // Deprecated
	Signature        Size
}

type MonitorHandler func(string, MagickOffsetType, MagickSizeType, *ExceptionInfo) bool
type MontageInfo struct {
	Geometry        *string
	Tile            *string
	Title           *string
	Frame           *string
	Texture         *string
	Font            *string
	Pointsize       float64
	BorderWidth     Size
	Shadow          MagickBooleanType
	Fill            PixelPacket
	Stroke          PixelPacket
	BackgroundColor PixelPacket
	BorderColor     PixelPacket
	MatteColor      PixelPacket
	Gravity         GravityType
	Filename        [MaxTextExtent]Char
	Debug           MagickBooleanType
	Signature       Size
}
type MontageMode Enum

const (
	UndefinedMode MontageMode = iota
	FrameMode
	UnframeMode
	ConcatenateMode
)

type MorphologyMethod Enum

const (
	UndefinedMorphology MorphologyMethod = iota
	ConvolveMorphology
	CorrelateMorphology
	ErodeMorphology
	DilateMorphology
	ErodeIntensityMorphology
	DilateIntensityMorphology
	DistanceMorphology
	OpenMorphology
	CloseMorphology
	OpenIntensityMorphology
	CloseIntensityMorphology
	SmoothMorphology
	EdgeInMorphology
	EdgeOutMorphology
	EdgeMorphology
	TopHatMorphology
	BottomHatMorphology
	HitAndMissMorphology
	ThinningMorphology
	ThickenMorphology
	VoronoiMorphology
	IterativeDistanceMorphology
)

type NexusInfo struct {
	Mapped              MagickBooleanType
	Region              RectangleInfo
	Length              MagickSizeType
	Cache               *PixelPacket
	Pixels              *PixelPacket
	AuthenticPixelCache MagickBooleanType
	Indexes             *IndexPacket
	Signature           Size
}

// type NodeInfo struct {
// 	Child         [16]*NodeInfo
// 	List          *ColorPacket
// 	NumberUnique  MagickSizeType
// 	Level         unsignedlong
// 	Parent        *NodeInfo
// 	TotalColor    RealPixelPacket
// 	QuantizeError MagickRealType
// 	ColorNumber   unsignedlong
// 	Id            unsignedlong
// 	Key           *void
// 	Value         *void
// 	Left          *NodeInfo
// 	Right         *NodeInfo
// }
type NoiseType Enum

const (
	UndefinedNoise NoiseType = iota
	UniformNoise
	GaussianNoise
	MultiplicativeGaussianNoise
	ImpulseNoise
	LaplacianNoise
	PoissonNoise
	RandomNoise
)

type OrientationType Enum

const (
	UndefinedOrientation OrientationType = iota
	TopLeftOrientation
	TopRightOrientation
	BottomRightOrientation
	BottomLeftOrientation
	LeftTopOrientation
	RightTopOrientation
	RightBottomOrientation
	LeftBottomOrientation
)

type PaintMethod Enum

const (
	UndefinedMethod PaintMethod = iota
	PointMethod
	ReplaceMethod
	FloodfillMethod
	FillToBorderMethod
	ResetMethod
)

type PathType Enum

const (
	UndefinedPath PathType = iota
	MagickPath
	RootPath
	HeadPath
	TailPath
	BasePath
	ExtensionPath
	SubimagePath
	CanonicalPath
)

// Little-endian
type PixelPacket struct {
	Blue, Green, Red, Opacity Quantum
}

type PixelPacketBE struct {
	Red, Green, Blue, Opacity Quantum
}

type PointInfo struct {
	X, Y float64
}

type PolicyDomain Enum

const (
	UndefinedPolicyDomain PolicyDomain = iota
	CoderPolicyDomain
	DelegatePolicyDomain
	FilterPolicyDomain
	PathPolicyDomain
	ResourcePolicyDomain
	SystemPolicyDomain
)

type PolicyRights Enum

const (
	ReadPolicyRights PolicyRights = 1 << iota
	WritePolicyRights
	ExecutePolicyRights
	UndefinedPolicyRights PolicyRights = 0
	NoPolicyRights                     = UndefinedPolicyRights
)

type PreviewType Enum

const (
	UndefinedPreview PreviewType = iota
	RotatePreview
	ShearPreview
	RollPreview
	HuePreview
	SaturationPreview
	BrightnessPreview
	GammaPreview
	SpiffPreview
	DullPreview
	GrayscalePreview
	QuantizePreview
	DespecklePreview
	ReduceNoisePreview
	AddNoisePreview
	SharpenPreview
	BlurPreview
	ThresholdPreview
	EdgeDetectPreview
	SpreadPreview
	SolarizePreview
	ShadePreview
	RaisePreview
	SegmentPreview
	SwirlPreview
	ImplodePreview
	WavePreview
	OilPaintPreview
	CharcoalDrawingPreview
	JPEGPreview
)

type PrimaryInfo struct {
	X, Y, Z float64
}

type PrimitiveInfo struct {
	Point       PointInfo
	Coordinates Size
	Primitive   PrimitiveType
	Method      PaintMethod
	Text        *string
}

type PrimitiveType Enum

const (
	UndefinedPrimitive PrimitiveType = iota
	PointPrimitive
	LinePrimitive
	RectanglePrimitive
	RoundRectanglePrimitive
	ArcPrimitive
	EllipsePrimitive
	CirclePrimitive
	PolylinePrimitive
	PolygonPrimitive
	BezierPrimitive
	ColorPrimitive
	MattePrimitive
	TextPrimitive
	ImagePrimitive
	PathPrimitive
)

type ProfileInfo struct {
	Name      *string
	Length    Size
	Info      *byte
	Signature Size
}

type QuantizeInfo struct {
	NumberColors Size
	TreeDepth    Size
	Dither       MagickBooleanType
	Colorspace   ColorspaceType
	MeasureError MagickBooleanType
	Signature    Size
	DitherMethod DitherMethod
}

type QuantumAlphaType Enum

const (
	UndefinedQuantumAlpha QuantumAlphaType = iota
	AssociatedQuantumAlpha
	DisassociatedQuantumAlpha
)

type QuantumFormatType Enum

const (
	UndefinedQuantumFormat QuantumFormatType = iota
	FloatingPointQuantumFormat
	SignedQuantumFormat
	UnsignedQuantumFormat
)

type QuantumInfo struct {
	Depth         Size
	Quantum       Size
	Format        QuantumFormatType
	Minimum       float64
	Maximum       float64
	Scale         float64
	Pad           Size
	MinIsWhite    MagickBooleanType
	Pack          MagickBooleanType
	AlphaType     QuantumAlphaType
	NumberThreads Size
	Pixels        **byte
	Extent        Size
	Endian        EndianType
	State         QuantumState
	Semaphore     *SemaphoreInfo
	Signature     Size
}

type QuantumState struct {
	InverseScale float64
	Pixel        uint
	Bits         Size
	Mask         *uint
}

type QuantumType Enum

const (
	UndefinedQuantum QuantumType = iota
	AlphaQuantum
	BlackQuantum
	BlueQuantum
	CMYKAQuantum
	CMYKQuantum
	CyanQuantum
	GrayAlphaQuantum
	GrayQuantum
	GreenQuantum
	IndexAlphaQuantum
	IndexQuantum
	MagentaQuantum
	OpacityQuantum
	RedQuantum
	RGBAQuantum
	BGRAQuantum
	RGBOQuantum
	RGBQuantum
	YellowQuantum
	grayPadQuantum // Deprecated
	RGBPadQuantum
	CbYCrYQuantum
	CbYCrQuantum
	CbYCrAQuantum
	CMYKOQuantum
	BGRQuantum
	BGROQuantum
)

type RectangleInfo struct {
	Width, Height Size
	X, Y          Long
}

type ReferenceType Enum

const (
	UndefinedReference ReferenceType = iota
	GradientReference
)

type RegistryType Enum

const (
	UndefinedRegistryType RegistryType = iota
	ImageRegistryType
	ImageInfoRegistryType
	StringRegistryType
)

type RenderingIntent Enum

const (
	UndefinedIntent RenderingIntent = iota
	SaturationIntent
	PerceptualIntent
	AbsoluteIntent
	RelativeIntent
)

type ResolutionType Enum

const (
	UndefinedResolution ResolutionType = iota
	PixelsPerInchResolution
	PixelsPerCentimeterResolution
)

type ResourceType Enum

const (
	UndefinedResource ResourceType = iota
	AreaResource
	DiskResource
	FileResource
	MapResource
	MemoryResource
	ThreadResource
	TimeResource
	ThrottleResource
)

type SegmentInfo struct {
	X1, Y1, X2, Y2 float64
}

type SparseColorMethod Enum

const (
	// Without the type, godoc doesn't recognise ownership
	UndefinedColorInterpolate   SparseColorMethod = SparseColorMethod(UndefinedDistortion)
	BarycentricColorInterpolate                   = SparseColorMethod(AffineDistortion)
	BilinearColorInterpolate                      = SparseColorMethod(BilinearReverseDistortion)
	PolynomialColorInterpolate                    = SparseColorMethod(PolynomialDistortion)
	ShepardsColorInterpolate                      = SparseColorMethod(ShepardsDistortion)
	VoronoiColorInterpolate                       = SparseColorMethod(SentinelDistortion)
	InverseColorInterpolate                       = VoronoiColorInterpolate + 1
)

type SpreadMethod Enum

const (
	UndefinedSpread SpreadMethod = iota
	PadSpread
	ReflectSpread
	RepeatSpread
)

type StatisticType Enum

const (
	UndefinedStatistic StatisticType = iota
	GradientStatistic
	MaximumStatistic
	MeanStatistic
	MedianStatistic
	MinimumStatistic
	ModeStatistic
	NonpeakStatistic
	StandardDeviationStatistic
)

type StopInfo struct {
	Color  MagickPixelPacket
	Offset MagickRealType
}

type StorageType Enum

const (
	UndefinedPixel StorageType = iota
	CharPixel
	float64Pixel
	FloatPixel
	IntegerPixel
	LongPixel
	QuantumPixel
	ShortPixel
)

type StreamHandler func(*Image, *Void, Size) Size

// type StreamType Enum

// const (
// 	UndefinedStream StreamType = iota
// 	FileStream
// 	StandardStream
// 	PipeStream
// 	ZipStream
// 	BZipStream
// 	FifoStream
// 	BlobStream
// )

type StretchType Enum

const (
	UndefinedStretch StretchType = iota
	NormalStretch
	UltraCondensedStretch
	ExtraCondensedStretch
	CondensedStretch
	SemiCondensedStretch
	SemiExpandedStretch
	ExpandedStretch
	ExtraExpandedStretch
	UltraExpandedStretch
	AnyStretch
)

type StringInfo struct {
	Path_     [MaxTextExtent]Char
	Datum_    *byte
	Length_   Size
	Signature Size
}
type StyleType Enum

const (
	UndefinedStyle StyleType = iota
	NormalStyle
	ItalicStyle
	ObliqueStyle
	AnyStyle
)

type Timer struct {
	Start,
	Stop,
	Total float64
}
type TimerInfo struct {
	User      Timer
	Elapsed   Timer
	State     TimerState
	Signature Size
}
type TimerState Enum

const (
	UndefinedTimerState TimerState = iota
	StoppedTimerState
	RunningTimerState
)

type TransferImageViewMethod func(*ImageView, *ImageView, SSize, int, *Void) MagickBooleanType
type TransferPixelViewMethod func(*PixelView, *PixelView, *Void) bool
type TransferWandViewMethod func(*WandView, *WandView, SSize, int, *Void) bool
type TypeInfo struct {
	Face        Size
	Path        *string
	Name        *string
	Description *string
	Family      *string
	Style       StyleType
	Stretch     StretchType
	Weight      Size
	Encoding    *string
	Foundry     *string
	Format      *string
	Metrics     *string
	Glyphs      *string
	Stealth     MagickBooleanType
	Previous    *TypeInfo
	next        *TypeInfo // Deprecated
	Signature   Size
}
type TypeMetric struct {
	PixelsPerEm        PointInfo
	Ascent             float64
	Descent            float64
	Width              float64
	Height             float64
	MaxAdvance         float64
	UnderlinePosition  float64
	UnderlineThickness float64
	Bounds             SegmentInfo
	Origin             PointInfo
}
type UpdateImageViewMethod func(*ImageView, SSize, int, *Void) MagickBooleanType
type UpdatePixelViewMethod func(*PixelView, *Void) bool
type UpdateWandViewMethod func(*WandView, SSize, int, *Void) bool
type VirtualPixelMethod Enum

const (
	UndefinedVirtualPixelMethod VirtualPixelMethod = iota
	BackgroundVirtualPixelMethod
	constantVirtualPixelMethod // Deprecated
	DitherVirtualPixelMethod
	EdgeVirtualPixelMethod
	MirrorVirtualPixelMethod
	RandomVirtualPixelMethod
	TileVirtualPixelMethod
	TransparentVirtualPixelMethod
	MaskVirtualPixelMethod
	BlackVirtualPixelMethod
	GrayVirtualPixelMethod
	WhiteVirtualPixelMethod
	HorizontalTileVirtualPixelMethod
	VerticalTileVirtualPixelMethod
	HorizontalTileEdgeVirtualPixelMethod
	VerticalTileEdgeVirtualPixelMethod
	CheckerTileVirtualPixelMethod
)

// const WLUT_WIDTH = 1024

type XAnnotateInfo struct {
	X        int
	Y        int
	Width    uint
	Height   uint
	Degrees  float64
	FontInfo *XFontStruct
	Text     *string
	Stencil  AnnotationStencil
	Geometry [MaxTextExtent]Char
	Next     *XAnnotateInfo
	Previous *XAnnotateInfo
}

type XColormapType Enum

const (
	UndefinedColormap XColormapType = iota
	PrivateColormap
	SharedColormap
)

type XDrawInfo struct {
	X                 int
	Y                 int
	Width             uint
	Height            uint
	Degrees           float64
	Stencil           AnnotationStencil
	Element           ElementType
	Stipple           Pixmap
	LineWidth         uint
	LineInfo          XSegment
	NumberCoordinates uint
	RectangleInfo     RectangleInfo
	CoordinateInfo    *XPoint
	Geometry          [MaxTextExtent]Char
}

type XImportInfo struct {
	Frame   MagickBooleanType
	Borders MagickBooleanType
	Screen  MagickBooleanType
	Descend MagickBooleanType
	Silent  MagickBooleanType
}

type XPixelInfo struct {
	Colors           SSize
	Pixels           *UnsignedLong
	ForegroundColor  XColor
	BackgroundColor  XColor
	BorderColor      XColor
	MatteColor       XColor
	HighlightColor   XColor
	ShadowColor      XColor
	DepthColor       XColor
	TroughColor      XColor
	BoxColor         XColor
	PenColor         XColor
	PenColors        [MaxNumberPens]XColor
	AnnotateContext  GC
	HighlightContext GC
	WidgetContext    GC
	BoxIndex         UnsignedShort
	PenIndex         UnsignedShort
}

type XResourceInfo struct {
	ResourceDatabase XrmDatabase
	ImageInfo        *ImageInfo
	QuantizeInfo     *QuantizeInfo
	Colors           Size
	CloseServer      MagickBooleanType
	Backdrop         MagickBooleanType
	BackgroundColor  *string
	BorderColor      *string
	ClientName       *string
	Colormap         XColormapType
	BorderWidth      uint
	Delay            Size
	ColorRecovery    MagickBooleanType
	ConfirmExit      MagickBooleanType
	ConfirmEdit      MagickBooleanType
	DisplayGamma     *string
	Font             *string
	FontName         [MaxNumberFonts]*string
	ForegroundColor  *string
	DisplayWarnings  MagickBooleanType
	GammaCorrect     MagickBooleanType
	IconGeometry     *string
	Iconic           MagickBooleanType
	Immutable        MagickBooleanType
	ImageGeometry    *string
	MapType          *string
	MatteColor       *string
	Name             *string
	Magnify          uint
	Pause            uint
	PenColors        [MaxNumberPens]*string
	TextFont         *string
	Title            *string
	Quantum          int
	Update           uint
	UsePixmap        MagickBooleanType
	UseSharedMemory  MagickBooleanType
	UndoCache        Size
	VisualType       *string
	WindowGroup      *string
	WindowId         *string
	WriteFilename    *string
	CopyImage        *Image
	Gravity          int
	HomeDirectory    [MaxTextExtent]Char
}

type XWindowInfo struct {
	// Id               Window
	// Root             Window
	// Visual           *Visual
	// StorageClass     uint
	// Depth            uint
	// VisualInfo       *XVisualInfo
	// MapInfo          *XStandardColormap
	// PixelInfo        *XPixelInfo
	// FontInfo         *XFontStruct
	// AnnotateContext  GC
	// HighlightContext GC
	// WidgetContext    GC
	// Cursor           Cursor
	// BusyCursor       Cursor
	// Name             *string
	// Geometry         *string
	// IconName         *string
	// IconGeometry     *string
	// CropGeometry     *string
	// Data             Size
	// Flags            Size
	// X                int
	// Y                int
	// Width            uint
	// Height           uint
	// MinWidth         uint
	// MinHeight        uint
	// WidthInc         uint
	// HeightInc        uint
	// BorderWidth      uint
	// UsePixmap        MagickBooleanType
	// Immutable        MagickBooleanType
	// Shape            MagickBooleanType
	// SharedMemory     MagickBooleanType
	// Screen           int
	// Ximage           *XImage
	// MatteImage       *XImage
	// HighlightStipple Pixmap
	// ShadowStipple    Pixmap
	// Pixmap           Pixmap
	// Pixmaps          *Pixmap
	// MattePixmap      Pixmap
	// MattePixmaps     *Pixmap
	// Attributes       XSetWindowAttributes // NOT POINTER
	// WindowChanges    XWindowChanges // NOT POINTER
	// SegmentInfo      *Void
	// Mask             Long
	// Orphan           MagickBooleanType
	// Mapped           MagickBooleanType
	// Stasis           MagickBooleanType
	// Image            *Image
	// Destroy          MagickBooleanType
}

type XWindows struct {
	// display            *Display
	// map_info           *XStandardColormap
	// icon_map           *XStandardColormap
	// visual_info        *XVisualInfo
	// icon_visual        *XVisualInfo
	// pixel_info         *XPixelInfo
	// icon_pixel         *XPixelInfo
	// font_info          *XFontStruct
	// icon_resources     *XResourceInfo
	// class_hints        *XClassHint
	// manager_hints      *XWMHints
	// context            XWindowInfo
	// group_leader       XWindowInfo
	// backdrop           XWindowInfo
	// icon               XWindowInfo
	// image              XWindowInfo
	// info               XWindowInfo
	// magnify            XWindowInfo
	// pan                XWindowInfo
	// command            XWindowInfo
	// widget             XWindowInfo
	// popup              XWindowInfo
	// wm_protocols       Atom
	// wm_delete_window   Atom
	// wm_take_focus      Atom
	// im_protocols       Atom
	// im_remote_command  Atom
	// im_update_widget   Atom
	// im_update_colormap Atom
	// im_former_image    Atom
	// im_retain_colors   Atom
	// im_next_image      Atom
	// im_exit            Atom
	// dnd_protocols      Atom
}
