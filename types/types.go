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
	FILE                 fix
	PixelIntensityMethod fix
	XrmDatabase          fix
)
type (
	fix uintptr

	Char             int8
	Enum             int
	IndexPacket      Quantum
	Long             int    // TODO(t):size on gcc vs msc 32 vs 64
	MagickOffsetType int64  // TODO(t):size on gcc vs msc 32 vs 64
	MagickSizeType   uint64 // TODO(t):size on gcc vs msc 32 vs 64
	MagickStatusType uint
	Size             uint // TODO(t):size_t on gcc vs msc 32 vs 64
	SSize            int  // TODO(t):ssize_t on gcc vs msc 32 vs 64
	Time             fix
	UnsignedShort    uint16 // TODO(t):ssize_t on gcc vs msc 32 vs 64
	Void             struct{}
)
type (
	Ascii85Info   struct{}
	BlobInfo      struct{}
	CacheView     struct{}
	FxInfo        struct{}
	PixelView     struct{}
	SemaphoreInfo struct{}
	WandView      struct{}
)

// 	ErrorHandler          func(ExceptionType, string, string)
// 	FatalErrorHandler     func(ExceptionType, string, string)
// 	GeometryInfo          struct{ Rho, Sigma, Xi, Psi, Chi float64 }
// 	MonitorHandler        func(string, MagickOffsetType, MagickSizeType, *ExceptionInfo) bool

// 	RandomInfo      struct{}
// 	RealPixelPacket struct{ Red, Green, Blue, Opacity MagickRealType }

// 	WarningHandler func(ExceptionType, string, string)
// 	XImportInfo    struct{ Frame, Borders, Screen, Descend, Silent MagickBooleanType }

// 	AcquireMemoryHandler func(size_t) *Void

// 	Atom            fix
// 	Colormap        fix
// 	CompositeMethod fix

// 	Cursor                            fix
// 	DestroyMemoryHandler              fix
// 	DIR                               fix
// 	DuplexTransferImageViewMethodFunc fix
// 	GC                            fix
// 	GetImageViewMethodFunc        fix
// 	Long                          fix
// 	MagickCore                    fix
// 	MagickThreadType              pthread_t // ok // or DWORD
// 	PixelImage                    fix
// 	ResizeMemoryHandler           fix
// 	SetImageViewMethodFunc        fix
// 	Status                        fix
// 	Time                          fix
// 	TransferImageViewMethodFunc   fix
// 	UnsignedLong                  fix
// 	UpdateImageViewMethodFunc     fix
// 	Window                        fix
// 	XAnnotateInfo                 fix
// 	XClassHint                    fix
// 	XColor                        fix
// 	XDrawInfo                     fix
// 	XErrorEvent                   fix
// 	XEvent                        fix
// 	XFontStruct                   fix
// 	XImage                        fix
// 	XPixelInfo                    fix
// 	XSegment                      fix
// 	XStandardColormap             fix
// 	XVisualInfo                   fix
// 	XWindowInfo                   fix
// 	XWindows                      fix
// 	XWMHints                      fix
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

// type BlobMode Enum

// const (
// 	UndefinedBlobMode BlobMode = iota
// 	ReadBlobMode
// 	ReadBinaryBlobMode
// 	WriteBlobMode
// 	WriteBinaryBlobMode
// )

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

// type CacheMethods struct {
// 	// acquireIndexesFromHandler   AcquireIndexesFromHandler
// 	// acquireOnePixelFromHandler AcquireOnePixelFromHandler
// 	// acquirePixelHandler          AcquirePixelHandler
// 	// destroyPixelHandler          DestroyPixelHandler
// 	// getIndexesFromHandler       GetIndexesFromHandler
// 	// getOnePixelFromHandler     GetOnePixelFromHandler
// 	// getPixelHandler              GetPixelHandler
// 	// getPixelsFromHandler        GetPixelsFromHandler
// 	// setPixelHandler              SetPixelHandler
// 	// syncPixelHandler             SyncPixelHandler
// }
// type CacheType Enum

// const (
// 	UndefinedCache CacheType = iota
// 	MemoryCache
// 	MapCache
// 	DiskCache
// )

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

// type CoderInfo struct {
// 	Path      *uint8
// 	Magick    *uint8
// 	Name      *uint8
// 	Stealth   MagickBooleanType
// 	Previous  *CoderInfo
// 	Next      *CoderInfo
// 	Signature UnsignedLong
// }
// type ColorInfo struct {
// 	Path       *uint8
// 	Name       *uint8
// 	Compliance ComplianceType
// 	Color      MagickPixelPacket
// 	Stealth    MagickBooleanType
// 	Previous   *ColorInfo
// 	Next       *ColorInfo
// 	Signature  UnsignedLong
// }
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

// type ConfigureInfo struct {
// 	Path      *char
// 	Name      *char
// 	Value     *char
// 	Stealth   MagickBooleanType
// 	Previous  *ConfigureInfo
// 	Next      *ConfigureInfo
// 	Signature unsignedlong
// }
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
	Primitive        *Char
	Geometry         *Char
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
	Text             *Char
	Face             Size
	Font             *Char
	Metrics          *Char
	Family           *Char
	Style            StyleType
	Stretch          StretchType
	Weight           Size
	Encoding         *Char
	Pointsize        float64
	Density          *Char
	Align            AlignType
	Undercolor       PixelPacket
	BorderColor      PixelPacket
	ServerName       *Char
	DashPattern      *float64
	ClipMask         *Char
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
	Id        *Char
	Type      ReferenceType
	Gradient  GradientInfo
	Signature Size
	Previous  *ElementReference
	Next      *ElementReference
}
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
	Reason      *Char
	Description *Char
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

type GetPixelViewMethod func(*PixelView, *Void) bool

type GetWandViewMethod func(*WandView, SSize, int, *Void) bool

// type GhostscriptVectors struct {
// 	// doc garbage
// 	// int (MagickDLLCall *exit)(gsMainInstance *)
// 	// int (MagickDLLCall *initWithArgs)(gsMainInstance *
// 	// 	char ** int (MagickDLLCall *newInstance)(gsMainInstance **
// 	// 	char **void * int (MagickDLLCall *runString)(gsMainInstance *
// 	// 	char **void *const char int * void (MagickDLLCall *deleteInstance)(gsMainInstance *)
// }
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
	Montage                *Char
	Directory              *Char
	Geometry_              *Char
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
	Key         *Char
	Value       *Char
	Compression MagickBooleanType
	Previous    *ImageAttribute
	Next        *ImageAttribute // Deprecated
}
type ImageInfo struct {
	Compression        CompressionType
	Orientation        OrientationType
	Temporary          MagickBooleanType
	Adjoin             MagickBooleanType
	Affirm             MagickBooleanType
	Antialias          MagickBooleanType
	Size               *Char
	Extract            *Char
	Page               *Char
	Scenes             *Char
	Scene              Size
	NumberScenes       Size
	Depth              Size
	Interlace          InterlaceType
	Endian             EndianType
	Units              ResolutionType
	Quality            Size
	SamplingFactor     *Char
	ServerName         *Char
	Font               *Char
	Texture            *Char
	Density            *Char
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
	View               *Char
	Authenticate       *Char
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
	Tile               *Char
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

// type LocaleInfo struct {
// 	Path      *char
// 	Tag       *char
// 	Message   *char
// 	Stealth   MagickBooleanType
// 	Previous  *LocaleInfo
// 	Next      *LocaleInfo
// 	Signature unsignedlong
// }
// type LogEventType Enum

// const (
// 	UndefinedEvents LogEventType = iota
// 	NoEvents
// 	TraceEvent
// 	AnnotateEvent
// 	BlobEvent
// 	CacheEvent
// 	CoderEvent
// 	ConfigureEvent
// 	DeprecateEvent
// 	DrawEvent
// 	ExceptionEvent
// 	LocaleEvent
// 	ModuleEvent
// 	ResourceEvent
// 	TransformEvent
// 	UserEvent
// 	WandEvent
// 	X11Event
// 	AllEvents
// )

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

// type LogInfo struct {
// 	EventMask   LogEventType
// 	HandlerMask LogHandlerType
// 	Path        *char
// 	Name        *char
// 	Filename    *char
// 	Format      *char
// 	Generations unsignedlong
// 	Limit       unsignedlong
// 	File        *FILE
// 	Generation  unsignedlong
// 	Append      MagickBooleanType
// 	Stealth     MagickBooleanType
// 	Timer       TimerInfo
// 	Signature   unsignedlong
// }
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

// type MagicInfo struct {
// 	Path      *char
// 	Name      *char
// 	Target    *char
// 	Magic     *byte
// 	Length    size_t
// 	Offset    MagickOffsetType
// 	Stealth   MagickBooleanType
// 	Previous  *MagicInfo
// 	Next      *MagicInfo
// 	Signature size_t
// }

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

// type MapMode Enum

// const (
// 	ReadMode MapMode = iota
// 	WriteMode
// 	IOMode
// )

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

// type MimeInfo struct {
// 	Path        *char
// 	Type        *char
// 	Description *char
// 	Pattern     *char
// 	Priority    ssize_t
// 	Offset      MagickOffsetType
// 	Extent      size_t
// 	DataType    DataType
// 	Mask        ssize_t
// 	Value       ssize_t
// 	Endian      EndianType
// 	Length      size_t
// 	Magic       *byte
// 	Stealth     MagickBooleanType
// 	Signature   size_t
// }
// type ModuleInfo struct {
// 	Path             *char
// 	Tag              *char
// 	Handle           *void
// 	UnregisterModule func()
// 	RegisterModule   func() unsignedlong
// 	LoadTime         time_t
// 	Stealth          MagickBooleanType
// 	Previous         *ModuleInfo
// 	Next             *ModuleInfo
// 	Signature        unsignedlong
// }
// type MontageInfo struct {
// 	Geometry        *char
// 	Tile            *char
// 	Title           *char
// 	Frame           *char
// 	Texture         *char
// 	Font            *char
// 	Pointsize       float64
// 	BorderWidth     unsignedlong
// 	Shadow          MagickBooleanType
// 	Fill            PixelPacket
// 	Stroke          PixelPacket
// 	BackgroundColor PixelPacket
// 	BorderColor     PixelPacket
// 	MatteColor      PixelPacket
// 	Gravity         GravityType
// 	Filename        [MaxTextExtent]char
// 	Debug           MagickBooleanType
// 	Signature       unsignedlong
// }
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

// type NexusInfo struct {
// 	Available MagickBooleanType
// 	Mapped    MagickBooleanType
// 	Columns   unsignedlong
// 	Rows      unsignedlong
// 	X         long
// 	Y         long
// 	Length    MagickSizeType
// 	Cache     *PixelPacket
// 	Pixels    *PixelPacket
// 	Indexes   *IndexPacket
// }
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

// type PathType Enum

// const (
// 	UndefinedPath PathType = iota
// 	MagickPath
// 	RootPath
// 	HeadPath
// 	TailPath
// 	BasePath
// 	ExtensionPath
// 	SubimagePath
// 	CanonicalPath
// )

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

// type PolicyDomain Enum

// const (
// 	UndefinedPolicyDomain PolicyDomain = iota
// 	CoderPolicyDomain
// 	DelegatePolicyDomain
// 	FilterPolicyDomain
// 	PathPolicyDomain
// 	ResourcePolicyDomain
// 	SystemPolicyDomain
// )

// type PolicyRights Enum

// const (
// 	ReadPolicyRights PolicyRights = 1 << iota
// 	WritePolicyRights
// 	ExecutePolicyRights
// 	UndefinedPolicyRights PolicyRights = 0
// 	NoPolicyRights                     = UndefinedPolicyRights
// )

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
	Text        *Char
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
	Name      *Char
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

// type RegistryType Enum

// const (
// 	UndefinedRegistryType RegistryType = iota
// 	ImageRegistryType
// 	ImageInfoRegistryType
// 	StringRegistryType
// )

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
type SetPixelViewMethod func(*PixelView, *Void) bool
type SetWandViewMethod func(*WandView, SSize, int, *Void) bool

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

// type ThresholdMap struct {
// 	MapId       *char
// 	Description *char
// 	Width       int
// 	Height      int
// 	Divisor     int
// 	Levels      *int
// }
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

type TransferPixelViewMethod func(*PixelView, *PixelView, *Void) bool
type TransferWandViewMethod func(*WandView, *WandView, SSize, int, *Void) bool

// type TypeInfo struct {
// 	Face        unsignedlong
// 	Path        *char
// 	Name        *char
// 	Description *char
// 	Family      *char
// 	Style       StyleType
// 	Stretch     StretchType
// 	Weight      unsignedlong
// 	Encoding    *char
// 	Foundry     *char
// 	Format      *char
// 	Metrics     *char
// 	Glyphs      *char
// 	Stealth     MagickBooleanType
// 	Previous    *TypeInfo
// 	Next        *TypeInfo
// 	Signature   unsignedlong
// }
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

type XColormapType Enum

const (
	UndefinedColormap XColormapType = iota
	PrivateColormap
	SharedColormap
)

type XResourceInfo struct {
	ResourceDatabase XrmDatabase
	ImageInfo        *ImageInfo
	QuantizeInfo     *QuantizeInfo
	Colors           Size
	CloseServer      MagickBooleanType
	Backdrop         MagickBooleanType
	BackgroundColor  *Char
	BorderColor      *Char
	ClientName       *Char
	Colormap         XColormapType
	BorderWidth      uint
	Delay            Size
	ColorRecovery    MagickBooleanType
	ConfirmExit      MagickBooleanType
	ConfirmEdit      MagickBooleanType
	DisplayGamma     *Char
	Font             *Char
	FontName         [MaxNumberFonts]*Char
	ForegroundColor  *Char
	DisplayWarnings  MagickBooleanType
	GammaCorrect     MagickBooleanType
	IconGeometry     *Char
	Iconic           MagickBooleanType
	Immutable        MagickBooleanType
	ImageGeometry    *Char
	MapType          *Char
	MatteColor       *Char
	Name             *Char
	Magnify          uint
	Pause            uint
	PenColors        [MaxNumberPens]*Char
	TextFont         *Char
	Title            *Char
	Quantum          int
	Update           uint
	UsePixmap        MagickBooleanType
	UseSharedMemory  MagickBooleanType
	UndoCache        Size
	VisualType       *Char
	WindowGroup      *Char
	WindowId         *Char
	WriteFilename    *Char
	CopyImage        *Image
	Gravity          int
	HomeDirectory    [MaxTextExtent]Char
}
