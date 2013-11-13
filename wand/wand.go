package magickwand

import . "github.com/tHinqa/outside"

type (
	fix int

	AffineMatrix                  fix
	AlignType                     fix
	AlphaChannelType              fix
	ChannelFeatures               fix
	ChannelStatistics             fix
	ChannelType                   fix
	ClipPathUnits                 fix
	ColorspaceType                fix
	CompositeOperator             fix
	CompressionType               fix
	DecorationType                fix
	DisposeType                   fix
	DistortImageMethod            fix
	DitherMethod                  fix
	DrawBooleanType               fix
	DrawInfo                      fix
	DrawingWand                   struct{}
	DrawWand                      fix
	DuplexTransferPixelViewMethod fix
	DuplexTransferWandViewMethod  fix
	EndianType                    fix
	ExceptionInfo                 fix
	ExceptionType                 fix
	FILE                          fix
	FillRule                      fix
	FilterTypes                   fix
	GetPixelViewMethod            fix
	GetWandViewMethod             fix
	GravityType                   fix
	Image                         struct{}
	ImageInfo                     fix
	ImageLayerMethod              fix
	ImageType                     fix
	ImageView                     fix
	IndexPacket                   fix
	InterlaceType                 fix
	InterpolateMethodPixel        fix
	InterpolatePixelMethod        fix
	KernelInfo                    struct{}
	LineCap                       fix
	LineJoin                      fix
	MagickCommand                 fix
	MagickEvaluateOperator        fix
	MagickFunction                fix
	MagickOffsetType              fix
	MagickPixelPacket             fix
	MagickProgressMonitorType     fix
	MagickSizeType                fix
	MagickWand                    struct{}
	MetricType                    fix
	MontageMode                   fix
	MorphologyMethod              fix
	NoiseType                     fix
	OrientationType               fix
	PaintMethod                   fix
	PixelIterator                 fix
	Pixeliterator                 fix
	PixelPacket                   fix
	PixelView                     fix
	PixelWand                     struct{}
	PointInfo                     fix
	PreviewType                   fix
	Quantum                       fix
	RectangeInfo                  fix
	RectangleInfo                 fix
	RenderingIntent               fix
	ResolutionType                fix
	ResourceType                  fix
	SetPixelViewMethod            fix
	SetWandViewMethod             fix
	size_t                        fix
	SparseColorMethod             fix
	ssize_t                       fix
	StatisticType                 fix
	StorageType                   fix
	StretchType                   fix
	StyleType                     fix
	TransferPixelViewMethod       fix
	TransferWandViewMethod        fix
	UpdatePixelViewMethod         fix
	UpdateWandViewMethod          fix
	VirtualPixelMethod            fix
	Void                          struct{}
	WandView                      fix
)

//TODO(t): Is... -> IsValid?

var (
	NewMagickWand func() *MagickWand

	NewPixelIterator       func(m *MagickWand) *PixelIterator
	NewPixelRegionIterator func(m *MagickWand, x, y ssize_t, width, height size_t) *PixelIterator
	NewPixelView           func(m *MagickWand) *PixelView
	NewPixelViewRegion     func(m *MagickWand, x, y ssize_t, width, height size_t) *PixelView
	NewWandView            func(m *MagickWand) *WandView
	NewWandViewExtent      func(m *MagickWand, x, y ssize_t, width, height size_t) *WandView

	ClearMagickWand        func(m *MagickWand)
	CloneMagickWand        func(m *MagickWand) *MagickWand
	DestroyMagickWand      func(m *MagickWand) *MagickWand
	GetImageFromMagickWand func(m *MagickWand) *Image
	IsMagickWand           func(m *MagickWand) bool

	MagickAdaptiveBlurImage              func(m *MagickWand, radius, sigma float64) bool
	MagickAdaptiveBlurImageChannel       func(m *MagickWand, channel ChannelType, radius, sigma float64) bool
	MagickAdaptiveResizeImage            func(m *MagickWand, columns, rows size_t) bool
	MagickAdaptiveSharpenImage           func(m *MagickWand, radius, sigma float64) bool
	MagickAdaptiveSharpenImageChannel    func(m *MagickWand, channel ChannelType, radius, sigma float64) bool
	MagickAdaptiveThresholdImage         func(m *MagickWand, width, height size_t, offset ssize_t) bool
	MagickAddImage                       func(m *MagickWand, addWand *MagickWand) bool
	MagickAddNoiseImage                  func(m *MagickWand, noise_type NoiseType) bool
	MagickAddNoiseImageChannel           func(m *MagickWand, channel ChannelType, noise_type NoiseType) bool
	MagickAffineTransformImage           func(m *MagickWand, d *DrawingWand) bool
	MagickAnimateImages                  func(m *MagickWand, server_name string) bool
	MagickAnnotateImage                  func(m *MagickWand, d *DrawingWand, x, y, angle float64, text string) bool
	MagickAppendImages                   func(m *MagickWand, stack bool) *MagickWand
	MagickAutoGammaImage                 func(m *MagickWand) bool
	MagickAutoGammaImageChannel          func(m *MagickWand, channel ChannelType) bool
	MagickAutoLevelImage                 func(m *MagickWand) bool
	MagickAutoLevelImageChannel          func(m *MagickWand, channel ChannelType) bool
	MagickAverageImages                  func(m *MagickWand) *MagickWand
	MagickBlackThresholdImage            func(m *MagickWand, threshold *PixelWand) bool
	MagickBlueShiftImage                 func(m *MagickWand, factor float64) bool
	MagickBlurImage                      func(m *MagickWand, radius, sigma float64) bool
	MagickBlurImageChannel               func(m *MagickWand, channel ChannelType, radius, sigma float64) bool
	MagickBorderImage                    func(m *MagickWand, bordercolor *PixelWand, width, height size_t) bool
	MagickBrightnessContrastImage        func(m *MagickWand, brightness, contrast float64) bool
	MagickBrightnessContrastImageChannel func(m *MagickWand, channel ChannelType, brightness, contrast float64) bool
	MagickCharcoalImage                  func(m *MagickWand, radius, sigma float64) bool
	MagickChopImage                      func(m *MagickWand, width, height size_t, x, y ssize_t) bool
	MagickClampImage                     func(m *MagickWand) bool
	MagickClampImageChannel              func(m *MagickWand, channel ChannelType) bool
	MagickClearException                 func(m *MagickWand) bool
	MagickClipImage                      func(m *MagickWand) bool
	MagickClipImagePath                  func(m *MagickWand, pathname string, inside bool) bool
	MagickClipPathImage                  func(m *MagickWand, pathname string, inside bool) bool
	MagickClutImage                      func(m *MagickWand, clutWand *MagickWand) bool
	MagickClutImageChannel               func(m *MagickWand, channel ChannelType, clutWand *MagickWand) bool
	MagickCoalesceImages                 func(m *MagickWand) *MagickWand
	MagickColorDecisionListImage         func(m *MagickWand, color_correction_collection string) bool
	MagickColorFloodfillImage            func(m *MagickWand, fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y ssize_t) bool
	MagickColorizeImage                  func(m *MagickWand, colorize, opacity *PixelWand) bool
	MagickColorMatrixImage               func(m *MagickWand, color_matrix *KernelInfo) bool
	MagickCombineImages                  func(m *MagickWand, channel ChannelType) *MagickWand
	MagickCommentImage                   func(m *MagickWand, comment string) bool
	MagickCompareImageChannels           func(m *MagickWand, reference *MagickWand, channel ChannelType, metric MetricType, distortion *float64) *MagickWand
	MagickCompareImageLayers             func(m *MagickWand, method ImageLayerMethod) *MagickWand
	MagickCompareImages                  func(m *MagickWand, reference *MagickWand, metric MetricType, distortion *float64) *MagickWand
	MagickCompositeImage                 func(m *MagickWand, sourceWand *MagickWand, compose CompositeOperator, x, y ssize_t) bool
	MagickCompositeImageChannel          func(m *MagickWand, channel ChannelType, compositeWand *MagickWand, compose CompositeOperator, x, y ssize_t) bool
	MagickCompositeLayers                func(m *MagickWand, sourceWand *MagickWand, compose CompositeOperator, x, y ssize_t) bool
	MagickConstituteImage                func(m *MagickWand, columns, rows size_t, map_ string, storage StorageType, pixels *Void) bool
	MagickContrastImage                  func(m *MagickWand, sharpen bool) bool
	MagickContrastStretchImage           func(m *MagickWand, black_point, white_point float64) bool
	MagickContrastStretchImageChannel    func(m *MagickWand, channel ChannelType, black_point, white_point float64) bool
	MagickConvolveImage                  func(m *MagickWand, order size_t, kernel *float64) bool
	MagickConvolveImageChannel           func(m *MagickWand, channel ChannelType, order size_t, kernel *float64) bool
	MagickCropImage                      func(m *MagickWand, width, height size_t, x, y ssize_t) bool
	MagickCycleColormapImage             func(m *MagickWand, displace ssize_t) bool
	MagickDecipherImage                  func(m *MagickWand, passphrase string) bool
	MagickDeconstructImages              func(m *MagickWand) *MagickWand
	MagickDeleteImageArtifact            func(m *MagickWand, artifact string) bool
	MagickDeleteImageProperty            func(m *MagickWand, property string) bool
	MagickDeleteOption                   func(m *MagickWand, option string) bool
	MagickDescribeImage                  func(m *MagickWand) string
	MagickDeskewImage                    func(m *MagickWand, threshold float64) bool
	MagickDespeckleImage                 func(m *MagickWand) bool
	MagickDisplayImage                   func(m *MagickWand, server_name string) bool
	MagickDisplayImages                  func(m *MagickWand, server_name string) bool
	MagickDistortImage                   func(m *MagickWand, method DistortImageMethod, number_arguments size_t, arguments *float64, bestfit bool) bool
	MagickDrawImage                      func(m *MagickWand, d *DrawingWand) bool
	MagickEdgeImage                      func(m *MagickWand, radius float64) bool
	MagickEmbossImage                    func(m *MagickWand, radius, sigma float64) bool
	MagickEncipherImage                  func(m *MagickWand, passphrase string) bool
	MagickEnhanceImage                   func(m *MagickWand) bool
	MagickEqualizeImage                  func(m *MagickWand) bool
	MagickEqualizeImageChannel           func(m *MagickWand, channel ChannelType) bool
	MagickEvaluateImage                  func(m *MagickWand, operator MagickEvaluateOperator, value float64) bool
	MagickEvaluateImageChannel           func(m *MagickWand, channel ChannelType, op MagickEvaluateOperator, value float64) bool
	MagickEvaluateImages                 func(m *MagickWand, operator MagickEvaluateOperator) bool
	MagickExportImagePixels              func(m *MagickWand, x, y ssize_t, columns, rows size_t, map_ string, storage StorageType, pixels *Void) bool
	MagickExtentImage                    func(m *MagickWand, width, height size_t, x, y ssize_t) bool
	MagickFilterImage                    func(m *MagickWand, kernel *KernelInfo) bool
	MagickFilterImageChannel             func(m *MagickWand, channel ChannelType, kernel *KernelInfo) bool
	MagickFlattenImages                  func(m *MagickWand) *MagickWand
	MagickFlipImage                      func(m *MagickWand) bool
	MagickFloodfillPaintImage            func(m *MagickWand, channel ChannelType, fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y ssize_t, invert bool) bool
	MagickFlopImage                      func(m *MagickWand) bool
	MagickForwardFourierTransformImage   func(m *MagickWand, magnitude bool) bool
	MagickFrameImage                     func(m *MagickWand, matte_color *PixelWand, width, height size_t, inner_bevel, outer_bevel ssize_t) bool
	MagickFunctionImage                  func(m *MagickWand, function MagickFunction, number_arguments size_t, arguments *float64) bool
	MagickFunctionImageChannel           func(m *MagickWand, channel ChannelType, function MagickFunction, number_arguments size_t, arguments *float64) bool
	MagickFxImage                        func(m *MagickWand, expression string) *MagickWand
	MagickFxImageChannel                 func(m *MagickWand, channel ChannelType, expression string) *MagickWand
	MagickGammaImage                     func(m *MagickWand, gamma float64) bool
	MagickGammaImageChannel              func(m *MagickWand, channel ChannelType, gamma float64) bool
	MagickGaussianBlurImage              func(m *MagickWand, radius, sigma float64) bool
	MagickGaussianBlurImageChannel       func(m *MagickWand, channel ChannelType, radius, sigma float64) bool
	MagickGetAntialias                   func(m *MagickWand) bool
	MagickGetBackgroundColor             func(m *MagickWand) *PixelWand
	MagickGetColorspace                  func(m *MagickWand) ColorspaceType
	MagickGetCompression                 func(m *MagickWand) CompressionType
	MagickGetCompressionQuality          func(m *MagickWand) size_t
	MagickGetException                   func(m *MagickWand, severity *ExceptionType) string
	MagickGetExceptionType               func(m *MagickWand) ExceptionType
	MagickGetFilename                    func(m *MagickWand) string
	MagickGetFont                        func(m *MagickWand) string
	MagickGetFormat                      func(m *MagickWand) string
	MagickGetGravity                     func(m *MagickWand) GravityType
	MagickGetImage                       func(m *MagickWand) *MagickWand
	MagickGetImageAlphaChannel           func(m *MagickWand) size_t
	MagickGetImageArtifact               func(m *MagickWand, artifact string) string
	MagickGetImageArtifacts              func(m *MagickWand, pattern, number_artifacts *size_t) string
	MagickGetImageAttribute              func(m *MagickWand, property string) string
	MagickGetImageBackgroundColor        func(m *MagickWand, background_color *PixelWand) bool
	MagickGetImageBlob                   func(m *MagickWand, length *size_t) *uint8
	MagickGetImageBluePrimary            func(m *MagickWand, x, y *float64) bool
	MagickGetImageBorderColor            func(m *MagickWand, border_color *PixelWand) bool
	MagickGetImageChannelDepth           func(m *MagickWand, channel ChannelType) size_t
	MagickGetImageChannelDistortion      func(m *MagickWand, reference *MagickWand, channel ChannelType, metric MetricType, distortion *float64) bool
	MagickGetImageChannelDistortionFIX   func(m *MagickWand, reference *MagickWand, metric MetricType) *float64
	MagickGetImageChannelFeatures        func(m *MagickWand, distance size_t) *ChannelFeatures
	MagickGetImageChannelKurtosis        func(m *MagickWand, channel ChannelType, kurtosis, skewness *float64) bool
	MagickGetImageChannelMean            func(m *MagickWand, channel ChannelType, mean, standard_deviation *float64) bool
	MagickGetImageChannelRange           func(m *MagickWand, channel ChannelType, minima, maxima *float64) bool
	MagickGetImageChannelStatistics      func(m *MagickWand) *ChannelStatistics
	MagickGetImageClipMask               func(m *MagickWand) *MagickWand
	MagickGetImageColormapColor          func(m *MagickWand, index size_t, color *PixelWand) bool
	MagickGetImageColors                 func(m *MagickWand) size_t
	MagickGetImageColorspace             func(m *MagickWand) ColorspaceType
	MagickGetImageCompose                func(m *MagickWand) CompositeOperator
	MagickGetImageCompression            func(m *MagickWand) CompressionType
	MagickGetImageCompressionQuality     func(m *MagickWand) size_t
	MagickGetImageDelay                  func(m *MagickWand) size_t
	MagickGetImageDepth                  func(m *MagickWand) size_t
	MagickGetImageDispose                func(m *MagickWand) DisposeType
	MagickGetImageDistortion             func(m *MagickWand, reference *MagickWand, metric MetricType, distortion *float64) bool
	MagickGetImageFilename               func(m *MagickWand) string
	MagickGetImageFormat                 func(m *MagickWand) string
	MagickGetImageFuzz                   func(m *MagickWand) float64
	MagickGetImageGamma                  func(m *MagickWand) float64
	MagickGetImageGravity                func(m *MagickWand) GravityType
	MagickGetImageGreenPrimary           func(m *MagickWand, x, y *float64) bool
	MagickGetImageHeight                 func(m *MagickWand) size_t
	MagickGetImageHistogram              func(m *MagickWand, number_colors *size_t) **PixelWand
	MagickGetImageInterlaceScheme        func(m *MagickWand) InterlaceType
	MagickGetImageInterpolateMethod      func(m *MagickWand) InterpolatePixelMethod
	MagickGetImageIterations             func(m *MagickWand) size_t
	MagickGetImageLength                 func(m *MagickWand, length *MagickSizeType) bool
	MagickGetImageMatte                  func(m *MagickWand) size_t
	MagickGetImageMatteColor             func(m *MagickWand, matte_color *PixelWand) bool
	MagickGetImageOrientation            func(m *MagickWand) OrientationType
	MagickGetImagePage                   func(m *MagickWand, width, height *size_t, x, y *ssize_t) bool
	MagickGetImagePixelColor             func(m *MagickWand, x, y ssize_t, color *PixelWand) bool
	MagickGetImagePixels                 func(m *MagickWand, x, y ssize_t, columns, rows size_t, map_ string, storage StorageType, pixels *Void) bool
	MagickGetImageProfile                func(m *MagickWand, name string, length *size_t) *uint8
	MagickGetImageProfiles               func(m *MagickWand, pattern, number_profiles *size_t) string
	MagickGetImageProperties             func(m *MagickWand, pattern, number_properties *size_t) string
	MagickGetImageProperty               func(m *MagickWand, property string) string
	MagickGetImageRedPrimary             func(m *MagickWand, x, y *float64) bool
	MagickGetImageRegion                 func(m *MagickWand, width, height size_t, x, y ssize_t) *MagickWand
	MagickGetImageRenderingIntent        func(m *MagickWand) RenderingIntent
	MagickGetImageResolution             func(m *MagickWand, x, y *float64) bool
	MagickGetImagesBlob                  func(m *MagickWand, length *size_t) *uint8
	MagickGetImageScene                  func(m *MagickWand) size_t
	MagickGetImageSignature              func(m *MagickWand) string
	MagickGetImageSize                   func(m *MagickWand, length *MagickSizeType) bool
	MagickGetImageTicksPerSecond         func(m *MagickWand) size_t
	MagickGetImageTotalInkDensity        func(m *MagickWand) float64
	MagickGetImageType                   func(m *MagickWand) ImageType
	MagickGetImageUnits                  func(m *MagickWand) ResolutionType
	MagickGetImageVirtualPixelMethod     func(m *MagickWand) VirtualPixelMethod
	MagickGetImageWhitePoint             func(m *MagickWand, x, y *float64) bool
	MagickGetImageWidth                  func(m *MagickWand) size_t
	MagickGetInterlaceScheme             func(m *MagickWand) InterlaceType
	MagickGetInterpolateMethod           func(m *MagickWand) InterpolatePixelMethod
	MagickGetIteratorIndex               func(m *MagickWand) ssize_t
	MagickGetNumberImages                func(m *MagickWand) size_t
	MagickGetOption                      func(m *MagickWand, key string) string
	MagickGetOptions                     func(m *MagickWand, pattern, number_options *size_t) string
	MagickGetOrientation                 func(m *MagickWand) OrientationType
	MagickGetPage                        func(m *MagickWand, width, height *size_t, x, y *ssize_t) bool
	MagickGetPointsize                   func(m *MagickWand) float64
	MagickGetResolution                  func(m *MagickWand, x, y *float64) bool
	MagickGetSamplingFactor              func(m *MagickWand, number_factors *size_t) *float64
	MagickGetSize                        func(m *MagickWand, columns, rows *size_t) bool
	MagickGetSizeOffset                  func(m *MagickWand, offset *ssize_t) bool
	MagickGetType                        func(m *MagickWand) ImageType
	MagickHaldClutImage                  func(m *MagickWand, haldWand *MagickWand) bool
	MagickHaldClutImageChannel           func(m *MagickWand, channel ChannelType, haldWand *MagickWand) bool
	MagickHasNextImage                   func(m *MagickWand) bool
	MagickHasPreviousImage               func(m *MagickWand) bool
	MagickIdentifyImage                  func(m *MagickWand) string
	MagickImplodeImage                   func(m *MagickWand, radius float64) bool
	MagickImportImagePixels              func(m *MagickWand, x, y ssize_t, columns, rows size_t, map_ string, storage StorageType, pixels *Void) bool
	MagickInverseFourierTransformImage   func(m *MagickWand, phaseWand *MagickWand, magnitude bool) bool
	MagickLabelImage                     func(m *MagickWand, label string) bool
	MagickLevelImage                     func(m *MagickWand, black_point, gamma, white_point float64) bool
	MagickLevelImageChannel              func(m *MagickWand, channel ChannelType, black_point, gamma, white_point float64) bool
	MagickLinearStretchImage             func(m *MagickWand, black_point, white_point float64) bool
	MagickMagnifyImage                   func(m *MagickWand) bool
	MagickMapImage                       func(m *MagickWand, mapWand *MagickWand, dither bool) bool
	MagickMatteFloodfillImage            func(m *MagickWand, alpha, fuzz float64, bordercolor *PixelWand, x, y ssize_t) bool
	MagickMaximumImages                  func(m *MagickWand) *MagickWand
	MagickMedianFilterImage              func(m *MagickWand, radius float64) bool
	MagickMergeImageLayers               func(m *MagickWand, method ImageLayerMethod) *MagickWand
	MagickMinifyImage                    func(m *MagickWand) bool
	MagickMinimumImages                  func(m *MagickWand) *MagickWand
	MagickModeImage                      func(m *MagickWand, radius float64) bool
	MagickModulateImage                  func(m *MagickWand, brightness, saturation, hue float64) bool
	MagickMontageImage                   func(m *MagickWand, drawingWand DrawingWand, tile_geometry, thumbnail_geometry string, mode MontageMode, frame string) *MagickWand
	MagickMorphImages                    func(m *MagickWand, number_frames size_t) *MagickWand
	MagickMorphologyImage                func(m *MagickWand, method MorphologyMethod, iterations ssize_t, kernel *KernelInfo) bool
	MagickMorphologyImageChannel         func(m *MagickWand, channel ChannelType, method MorphologyMethod, iterations ssize_t, kernel *KernelInfo) bool
	MagickMosaicImages                   func(m *MagickWand) *MagickWand
	MagickMotionBlurImage                func(m *MagickWand, radius, sigma, angle float64) bool
	MagickMotionBlurImageChannel         func(m *MagickWand, channel ChannelType, radius, sigma, angle float64) bool
	MagickNegateImage                    func(m *MagickWand, gray bool) bool
	MagickNegateImageChannel             func(m *MagickWand, channel ChannelType, gray bool) bool
	MagickNewImage                       func(m *MagickWand, columns, rows size_t, background *PixelWand) bool
	MagickNextImage                      func(m *MagickWand) bool
	MagickNormalizeImage                 func(m *MagickWand) bool
	MagickNormalizeImageChannel          func(m *MagickWand, channel ChannelType) bool
	MagickOilPaintImage                  func(m *MagickWand, radius float64) bool
	MagickOpaqueImage                    func(m *MagickWand, target, fill *PixelWand, fuzz float64) bool
	MagickOpaquePaintImage               func(m *MagickWand, target, fill *PixelWand, fuzz float64, invert bool) bool
	MagickOpaquePaintImageChannel        func(m *MagickWand, channel ChannelType, target, fill *PixelWand, fuzz float64, invert bool) bool
	MagickOptimizeImageLayers            func(m *MagickWand) *MagickWand
	MagickOrderedPosterizeImage          func(m *MagickWand, threshold_map string) bool
	MagickOrderedPosterizeImageChannel   func(m *MagickWand, channel ChannelType, threshold_map string) bool
	MagickPaintFloodfillImage            func(m *MagickWand, channel ChannelType, fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y ssize_t) bool
	MagickPaintOpaqueImage               func(m *MagickWand, target, fill *PixelWand, fuzz float64) bool
	MagickPaintOpaqueImageChannel        func(m *MagickWand, channel ChannelType, target, fill *PixelWand, fuzz float64) bool
	MagickPaintTransparentImage          func(m *MagickWand, target *PixelWand, alpha, fuzz float64) bool
	MagickPingImage                      func(m *MagickWand, filename string) bool
	MagickPingImageBlob                  func(m *MagickWand, blob *Void, length size_t) bool
	MagickPingImageFile                  func(m *MagickWand, file *FILE) bool
	MagickPolaroidImage                  func(m *MagickWand, d *DrawingWand, angle float64) bool
	MagickPosterizeImage                 func(m *MagickWand, levels size_t, dither bool) bool
	MagickPreviewImages                  func(m *MagickWand, preview PreviewType) *MagickWand
	MagickPreviousImage                  func(m *MagickWand) bool
	MagickProfileImage                   func(m *MagickWand, name string, profile *Void, length size_t) bool
	MagickQuantizeImage                  func(m *MagickWand, number_colors size_t, colorspace ColorspaceType, treedepth size_t, dither, measure_error bool) bool
	MagickQuantizeImages                 func(m *MagickWand, number_colors size_t, colorspace ColorspaceType, treedepth size_t, dither, measure_error bool) bool
	MagickQueryFontMetrics               func(m *MagickWand, drawingWand *DrawingWand, text string) []float64
	MagickQueryMultilineFontMetrics      func(m *MagickWand, drawingWand *DrawingWand, text string) []float64
	MagickRadialBlurImage                func(m *MagickWand, angle float64) bool
	MagickRadialBlurImageChannel         func(m *MagickWand, channel ChannelType, angle float64) bool
	MagickRaiseImage                     func(m *MagickWand, width, height size_t, x, y ssize_t, raise bool) bool
	MagickRandomThresholdImage           func(m *MagickWand, low, high float64) bool
	MagickRandomThresholdImageChannel    func(m *MagickWand, channel ChannelType, low, high float64) bool
	MagickReadImage                      func(m *MagickWand, filename string) bool
	MagickReadImageBlob                  func(m *MagickWand, blob *Void, length size_t) bool
	MagickReadImageFile                  func(m *MagickWand, file *FILE) bool
	MagickRecolorImage                   func(m *MagickWand, order size_t, color_matrix *float64) bool
	MagickReduceNoiseImage               func(m *MagickWand, radius float64) bool
	MagickRegionOfInterestImage          func(m *MagickWand, width, height size_t, x, y ssize_t) *MagickWand
	MagickRemapImage                     func(m *MagickWand, remapWand *MagickWand, method DitherMethod) bool
	MagickRemoveImage                    func(m *MagickWand) bool
	MagickRemoveImageProfile             func(m *MagickWand, name string, length *size_t) *uint8
	MagickResampleImage                  func(m *MagickWand, x_resolution, y_resolution float64, filter FilterTypes, blur float64) bool
	MagickResetImagePage                 func(m *MagickWand, page string) bool
	MagickResetIterator                  func(m *MagickWand)
	MagickResizeImage                    func(m *MagickWand, columns, rows size_t, filter FilterTypes, blur float64) bool
	MagickRollImage                      func(m *MagickWand, x ssize_t, y size_t) bool
	MagickRotateImage                    func(m *MagickWand, background *PixelWand, degrees float64) bool
	MagickSampleImage                    func(m *MagickWand, columns, rows size_t) bool
	MagickScaleImage                     func(m *MagickWand, columns, rows size_t) bool
	MagickSegmentImage                   func(m *MagickWand, colorspace ColorspaceType, verbose bool, cluster_threshold, smooth_threshold float64) bool
	MagickSelectiveBlurImage             func(m *MagickWand, radius, sigma, threshold float64) bool
	MagickSelectiveBlurImageChannel      func(m *MagickWand, channel ChannelType, radius, sigma, threshold float64) bool
	MagickSeparateImageChannel           func(m *MagickWand, channel ChannelType) bool
	MagickSepiaToneImage                 func(m *MagickWand, threshold float64) bool
	MagickSetAntialias                   func(m *MagickWand, antialias bool) bool
	MagickSetBackgroundColor             func(m *MagickWand, background *PixelWand) bool
	MagickSetColorspace                  func(m *MagickWand, colorspace ColorspaceType) bool
	MagickSetCompression                 func(m *MagickWand, compression CompressionType) bool
	MagickSetCompressionQuality          func(m *MagickWand, quality size_t) bool
	MagickSetDepth                       func(m *MagickWand, depth size_t) bool
	MagickSetExtract                     func(m *MagickWand, geometry string) bool
	MagickSetFilename                    func(m *MagickWand, filename string) bool
	MagickSetFirstIterator               func(m *MagickWand)
	MagickSetFont                        func(m *MagickWand, font string) bool
	MagickSetFormat                      func(m *MagickWand, format string) bool
	MagickSetGravity                     func(m *MagickWand, type_ GravityType) bool
	MagickSetImage                       func(m *MagickWand, setWand *MagickWand) bool
	MagickSetImageAlphaChannel           func(m *MagickWand, alpha_type AlphaChannelType) bool
	MagickSetImageArtifact               func(m *MagickWand, artifact, value string) bool
	MagickSetImageAttribute              func(m *MagickWand, property, value string) bool
	MagickSetImageBackgroundColor        func(m *MagickWand, background *PixelWand) bool
	MagickSetImageBias                   func(m *MagickWand, bias float64) bool
	MagickSetImageBluePrimary            func(m *MagickWand, x, y float64) bool
	MagickSetImageBorderColor            func(m *MagickWand, border *PixelWand) bool
	MagickSetImageChannelDepth           func(m *MagickWand, channel ChannelType, depth size_t) bool
	MagickSetImageClipMask               func(m *MagickWand, clip_mask *MagickWand) bool
	MagickSetImageColor                  func(m *MagickWand, color *PixelWand) bool
	MagickSetImageColormapColor          func(m *MagickWand, index size_t, color *PixelWand) bool
	MagickSetImageColorspace             func(m *MagickWand, colorspace ColorspaceType) bool
	MagickSetImageCompose                func(m *MagickWand, compose CompositeOperator) bool
	MagickSetImageCompression            func(m *MagickWand, compression CompressionType) bool
	MagickSetImageCompressionQuality     func(m *MagickWand, quality size_t) bool
	MagickSetImageDelay                  func(m *MagickWand, delay size_t) bool
	MagickSetImageDepth                  func(m *MagickWand, depth size_t) bool
	MagickSetImageDispose                func(m *MagickWand, dispose DisposeType) bool
	MagickSetImageEndian                 func(m *MagickWand, endian EndianType) bool
	MagickSetImageExtent                 func(m *MagickWand, columns size_t, rows uint) bool
	MagickSetImageFilename               func(m *MagickWand, filename string) bool
	MagickSetImageFormat                 func(m *MagickWand, format string) bool
	MagickSetImageFuzz                   func(m *MagickWand, fuzz float64) bool
	MagickSetImageGamma                  func(m *MagickWand, gamma float64) bool
	MagickSetImageGravity                func(m *MagickWand, gravity GravityType) bool
	MagickSetImageGreenPrimary           func(m *MagickWand, x, y float64) bool
	MagickSetImageIndex                  func(m *MagickWand, index ssize_t) bool
	MagickSetImageInterlaceScheme        func(m *MagickWand, interlace InterlaceType) bool
	MagickSetImageInterpolateMethod      func(m *MagickWand, method InterpolatePixelMethod) bool
	MagickSetImageIterations             func(m *MagickWand, iterations size_t) bool
	MagickSetImageMatteColor             func(m *MagickWand, matte *bool) bool
	MagickSetImageMatteColorFIX          func(m *MagickWand, matte *PixelWand) bool
	MagickSetImageOpacity                func(m *MagickWand, alpha float64) bool
	MagickSetImageOrientation            func(m *MagickWand, orientation OrientationType) bool
	MagickSetImagePage                   func(m *MagickWand, width, height size_t, x, y ssize_t) bool
	MagickSetImagePixels                 func(m *MagickWand, x, y ssize_t, columns, rows size_t, map_ string, storage StorageType, pixels *Void) bool
	MagickSetImageProfile                func(m *MagickWand, name string, profile *Void, length size_t) bool
	MagickSetImageProgressMonitor        func(m *MagickWand, progress_monitor, client_data *Void) MagickProgressMonitorType
	MagickSetImageProperty               func(m *MagickWand, property, value string) bool
	MagickSetImageRedPrimary             func(m *MagickWand, x, y float64) bool
	MagickSetImageRenderingIntent        func(m *MagickWand, rendering_intent RenderingIntent) bool
	MagickSetImageResolution             func(m *MagickWand, x_resolution, y_resolution float64) bool
	MagickSetImageScene                  func(m *MagickWand, scene size_t) bool
	MagickSetImageTicksPerSecond         func(m *MagickWand, ticks_per_second ssize_t) bool
	MagickSetImageType                   func(m *MagickWand, image_type ImageType) bool
	MagickSetImageUnits                  func(m *MagickWand, units ResolutionType) bool
	MagickSetImageVirtualPixelMethod     func(m *MagickWand, method VirtualPixelMethod) VirtualPixelMethod
	MagickSetImageWhitePoint             func(m *MagickWand, x, y float64) bool
	MagickSetInterlaceScheme             func(m *MagickWand, interlace_scheme InterlaceType) bool
	MagickSetInterpolateMethod           func(m *MagickWand, method InterpolateMethodPixel) bool
	MagickSetIteratorIndex               func(m *MagickWand, index ssize_t) bool
	MagickSetLastIterator                func(m *MagickWand)
	MagickSetOption                      func(m *MagickWand, key, value string) bool
	MagickSetOrientation                 func(m *MagickWand, orientation OrientationType) bool
	MagickSetPage                        func(m *MagickWand, width, height size_t, x, y ssize_t) bool
	MagickSetPassphrase                  func(m *MagickWand, passphrase string) bool
	MagickSetPointsize                   func(m *MagickWand, pointsize float64) bool
	MagickSetProgressMonitor             func(m *MagickWand, progress_monitor, client_data *Void) MagickProgressMonitorType
	MagickSetResolution                  func(m *MagickWand, x_resolution, y_resolution float64) bool
	MagickSetSamplingFactors             func(m *MagickWand, number_factors size_t, sampling_factors *float64) bool
	MagickSetSize                        func(m *MagickWand, columns, rows size_t) bool
	MagickSetSizeOffset                  func(m *MagickWand, columns, rows size_t, offset ssize_t) bool
	MagickSetType                        func(m *MagickWand, image_type ImageType) bool
	MagickShadeImage                     func(m *MagickWand, gray bool, azimuth, elevation float64) bool
	MagickShadowImage                    func(m *MagickWand, opacity, sigma float64, x, y ssize_t) bool
	MagickSharpenImage                   func(m *MagickWand, radius, sigma float64) bool
	MagickSharpenImageChannel            func(m *MagickWand, channel ChannelType, radius, sigma float64) bool
	MagickShaveImage                     func(m *MagickWand, columns, rows size_t) bool
	MagickShearImage                     func(m *MagickWand, background *PixelWand, x_shear, y_shear float64) bool
	MagickSigmoidalContrastImage         func(m *MagickWand, sharpen bool, alpha, beta float64) bool
	MagickSigmoidalContrastImageChannel  func(m *MagickWand, channel ChannelType, sharpen bool, alpha, beta float64) bool
	MagickSimilarityImage                func(m *MagickWand, reference *MagickWand, offset *RectangeInfo, similarity *float64) *MagickWand
	MagickSketchImage                    func(m *MagickWand, radius, sigma, angle float64) bool
	MagickSmushImages                    func(m *MagickWand, stack bool, offset ssize_t) *MagickWand
	MagickSolarizeImage                  func(m *MagickWand, threshold float64) bool
	MagickSparseColorImage               func(m *MagickWand, channel ChannelType, method SparseColorMethod, number_arguments size_t, arguments *float64) bool
	MagickSpliceImage                    func(m *MagickWand, width, height size_t, x, y ssize_t) bool
	MagickSpreadImage                    func(m *MagickWand, radius float64) bool
	MagickStatisticImage                 func(m *MagickWand, type_ StatisticType, width float64, height size_t) bool
	MagickStatisticImageChannel          func(m *MagickWand, channel ChannelType, type_ StatisticType, width float64, height size_t) bool
	MagickSteganoImage                   func(m *MagickWand, watermarkWand, offset ssize_t) *MagickWand
	MagickStereoImage                    func(m *MagickWand, offsetWand *MagickWand) *MagickWand
	MagickStripImage                     func(m *MagickWand) bool
	MagickSwirlImage                     func(m *MagickWand, degrees float64) bool
	MagickTextureImage                   func(m *MagickWand, texture_wand *MagickWand) *MagickWand
	MagickThresholdImage                 func(m *MagickWand, threshold float64) bool
	MagickThresholdImageChannel          func(m *MagickWand, channel ChannelType, threshold float64) bool
	MagickThumbnailImage                 func(m *MagickWand, columns, rows size_t) bool
	MagickTintImage                      func(m *MagickWand, tint, opacity *PixelWand) bool
	MagickTransformImage                 func(m *MagickWand, crop, geometry string) *MagickWand
	MagickTransformImageColorspace       func(m *MagickWand, colorspace ColorspaceType) bool
	MagickTransparentImage               func(m *MagickWand, target *PixelWand, alpha, fuzz float64) bool
	MagickTransparentPaintImage          func(m *MagickWand, target *PixelWand, alpha, fuzz float64, invert bool) bool
	MagickTransposeImage                 func(m *MagickWand) bool
	MagickTransverseImage                func(m *MagickWand) bool
	MagickTrimImage                      func(m *MagickWand, fuzz float64) bool
	MagickUniqueImageColors              func(m *MagickWand) bool
	MagickUnsharpMaskImage               func(m *MagickWand, radius, sigma, amount, threshold float64) bool
	MagickUnsharpMaskImageChannel        func(m *MagickWand, channel ChannelType, radius, sigma, amount, threshold float64) bool
	MagickVignetteImage                  func(m *MagickWand, black_point, white_point float64, x, y ssize_t) bool
	MagickWaveImage                      func(m *MagickWand, amplitude, wave_length float64) bool
	MagickWhiteThresholdImage            func(m *MagickWand, threshold *PixelWand) bool
	MagickWriteImage                     func(m *MagickWand, filename string) bool
	MagickWriteImageBlob                 func(m *MagickWand, length *size_t) *uint8
	MagickWriteImageFile                 func(m *MagickWand, file *FILE) bool
	MagickWriteImages                    func(m *MagickWand, filename string, adjoin bool) bool
	MagickWriteImagesFile                func(m *MagickWand, file *FILE) bool
)

func (m *MagickWand) Clear()               { ClearMagickWand(m) }
func (m *MagickWand) Clone() *MagickWand   { return CloneMagickWand(m) }
func (m *MagickWand) Destroy() *MagickWand { return DestroyMagickWand(m) }
func (m *MagickWand) ImageFrom() *Image    { return GetImageFromMagickWand(m) }
func (m *MagickWand) IsMagickWand() bool   { return IsMagickWand(m) }

func (m *MagickWand) AdaptiveBlurImage(radius, sigma float64) bool {
	return MagickAdaptiveBlurImage(m, radius, sigma)
}
func (m *MagickWand) AdaptiveBlurImageChannel(channel ChannelType, radius, sigma float64) bool {
	return MagickAdaptiveBlurImageChannel(m, channel, radius, sigma)
}
func (m *MagickWand) AdaptiveResizeImage(columns, rows size_t) bool {
	return MagickAdaptiveResizeImage(m, columns, rows)
}
func (m *MagickWand) AdaptiveSharpenImage(radius, sigma float64) bool {
	return MagickAdaptiveSharpenImage(m, radius, sigma)
}
func (m *MagickWand) AdaptiveSharpenImageChannel(channel ChannelType, radius, sigma float64) bool {
	return MagickAdaptiveSharpenImageChannel(m, channel, radius, sigma)
}
func (m *MagickWand) AdaptiveThresholdImage(width, height size_t, offset ssize_t) bool {
	return MagickAdaptiveThresholdImage(m, width, height, offset)
}
func (m *MagickWand) AddImage(add_wand *MagickWand) bool { return MagickAddImage(m, add_wand) }
func (m *MagickWand) AddNoiseImage(noise_type NoiseType) bool {
	return MagickAddNoiseImage(m, noise_type)
}
func (m *MagickWand) AddNoiseImageChannel(channel ChannelType, noise_type NoiseType) bool {
	return MagickAddNoiseImageChannel(m, channel, noise_type)
}
func (m *MagickWand) AffineTransformImage(d *DrawingWand) bool {
	return MagickAffineTransformImage(m, d)
}
func (m *MagickWand) AnimateImages(server_name string) bool {
	return MagickAnimateImages(m, server_name)
}
func (m *MagickWand) AnnotateImage(d *DrawingWand, x, y, angle float64, text string) bool {
	return MagickAnnotateImage(m, d, x, y, angle, text)
}
func (m *MagickWand) AppendImages(stack bool) *MagickWand { return MagickAppendImages(m, stack) }
func (m *MagickWand) AutoGammaImage() bool                { return MagickAutoGammaImage(m) }
func (m *MagickWand) AutoGammaImageChannel(channel ChannelType) bool {
	return MagickAutoGammaImageChannel(m, channel)
}
func (m *MagickWand) AutoLevelImage() bool { return MagickAutoLevelImage(m) }
func (m *MagickWand) AutoLevelImageChannel(channel ChannelType) bool {
	return MagickAutoLevelImageChannel(m, channel)
}
func (m *MagickWand) AverageImages() *MagickWand { return MagickAverageImages(m) }
func (m *MagickWand) BlackThresholdImage(threshold *PixelWand) bool {
	return MagickBlackThresholdImage(m, threshold)
}
func (m *MagickWand) BlueShiftImage(factor float64) bool   { return MagickBlueShiftImage(m, factor) }
func (m *MagickWand) BlurImage(radius, sigma float64) bool { return MagickBlurImage(m, radius, sigma) }
func (m *MagickWand) BlurImageChannel(channel ChannelType, radius, sigma float64) bool {
	return MagickBlurImageChannel(m, channel, radius, sigma)
}
func (m *MagickWand) BorderImage(bordercolor *PixelWand, width, height size_t) bool {
	return MagickBorderImage(m, bordercolor, width, height)
}
func (m *MagickWand) BrightnessContrastImage(brightness, contrast float64) bool {
	return MagickBrightnessContrastImage(m, brightness, contrast)
}
func (m *MagickWand) BrightnessContrastImageChannel(channel ChannelType, brightness, contrast float64) bool {
	return MagickBrightnessContrastImageChannel(m, channel, brightness, contrast)
}
func (m *MagickWand) CharcoalImage(radius, sigma float64) bool {
	return MagickCharcoalImage(m, radius, sigma)
}
func (m *MagickWand) ChopImage(width, height size_t, x, y ssize_t) bool {
	return MagickChopImage(m, width, height, x, y)
}
func (m *MagickWand) ClampImage() bool { return MagickClampImage(m) }
func (m *MagickWand) ClampImageChannel(channel ChannelType) bool {
	return MagickClampImageChannel(m, channel)
}
func (m *MagickWand) ClearException() bool { return MagickClearException(m) }
func (m *MagickWand) ClipImage() bool      { return MagickClipImage(m) }
func (m *MagickWand) ClipImagePath(pathname string, inside bool) bool {
	return MagickClipImagePath(m, pathname, inside)
}
func (m *MagickWand) ClipPathImage(pathname string, inside bool) bool {
	return MagickClipPathImage(m, pathname, inside)
}
func (m *MagickWand) ClutImage(clut_wand *MagickWand) bool { return MagickClutImage(m, clut_wand) }
func (m *MagickWand) ClutImageChannel(channel ChannelType, clut_wand *MagickWand) bool {
	return MagickClutImageChannel(m, channel, clut_wand)
}
func (m *MagickWand) CoalesceImages() *MagickWand { return MagickCoalesceImages(m) }
func (m *MagickWand) ColorDecisionListImage(color_correction_collection string) bool {
	return MagickColorDecisionListImage(m, color_correction_collection)
}
func (m *MagickWand) ColorFloodfillImage(fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y ssize_t) bool {
	return MagickColorFloodfillImage(m, fill, fuzz, bordercolor, x, y)
}
func (m *MagickWand) ColorizeImage(colorize, opacity *PixelWand) bool {
	return MagickColorizeImage(m, colorize, opacity)
}
func (m *MagickWand) ColorMatrixImage(color_matrix *KernelInfo) bool {
	return MagickColorMatrixImage(m, color_matrix)
}
func (m *MagickWand) CombineImages(channel ChannelType) *MagickWand {
	return MagickCombineImages(m, channel)
}
func (m *MagickWand) CommentImage(comment string) bool { return MagickCommentImage(m, comment) }
func (m *MagickWand) CompareImageChannels(reference *MagickWand, channel ChannelType, metric MetricType, distortion *float64) *MagickWand {
	return MagickCompareImageChannels(m, reference, channel, metric, distortion)
}
func (m *MagickWand) CompareImageLayers(method ImageLayerMethod) *MagickWand {
	return MagickCompareImageLayers(m, method)
}
func (m *MagickWand) CompareImages(reference *MagickWand, metric MetricType, distortion *float64) *MagickWand {
	return MagickCompareImages(m, reference, metric, distortion)
}
func (m *MagickWand) CompositeImage(source_wand *MagickWand, compose CompositeOperator, x, y ssize_t) bool {
	return MagickCompositeImage(m, source_wand, compose, x, y)
}
func (m *MagickWand) CompositeImageChannel(channel ChannelType, composite_wand *MagickWand, compose CompositeOperator, x, y ssize_t) bool {
	return MagickCompositeImageChannel(m, channel, composite_wand, compose, x, y)
}
func (m *MagickWand) CompositeLayers(source_wand *MagickWand, compose CompositeOperator, x, y ssize_t) bool {
	return MagickCompositeLayers(m, source_wand, compose, x, y)
}
func (m *MagickWand) ConstituteImage(columns, rows size_t, map_ string, storage StorageType, pixels *Void) bool {
	return MagickConstituteImage(m, columns, rows, map_, storage, pixels)
}
func (m *MagickWand) ContrastImage(sharpen bool) bool { return MagickContrastImage(m, sharpen) }
func (m *MagickWand) ContrastStretchImage(black_point, white_point float64) bool {
	return MagickContrastStretchImage(m, black_point, white_point)
}
func (m *MagickWand) ContrastStretchImageChannel(channel ChannelType, black_point, white_point float64) bool {
	return MagickContrastStretchImageChannel(m, channel, black_point, white_point)
}
func (m *MagickWand) ConvolveImage(order size_t, kernel *float64) bool {
	return MagickConvolveImage(m, order, kernel)
}
func (m *MagickWand) ConvolveImageChannel(channel ChannelType, order size_t, kernel *float64) bool {
	return MagickConvolveImageChannel(m, channel, order, kernel)
}
func (m *MagickWand) CropImage(width, height size_t, x, y ssize_t) bool {
	return MagickCropImage(m, width, height, x, y)
}
func (m *MagickWand) CycleColormapImage(displace ssize_t) bool {
	return MagickCycleColormapImage(m, displace)
}
func (m *MagickWand) DecipherImage(passphrase string) bool { return MagickDecipherImage(m, passphrase) }
func (m *MagickWand) DeconstructImages() *MagickWand       { return MagickDeconstructImages(m) }
func (m *MagickWand) DeleteImageArtifact(artifact string) bool {
	return MagickDeleteImageArtifact(m, artifact)
}
func (m *MagickWand) DeleteImageProperty(property string) bool {
	return MagickDeleteImageProperty(m, property)
}
func (m *MagickWand) DeleteOption(option string) bool      { return MagickDeleteOption(m, option) }
func (m *MagickWand) DescribeImage() string                { return MagickDescribeImage(m) }
func (m *MagickWand) DeskewImage(threshold float64) bool   { return MagickDeskewImage(m, threshold) }
func (m *MagickWand) DespeckleImage() bool                 { return MagickDespeckleImage(m) }
func (m *MagickWand) DisplayImage(server_name string) bool { return MagickDisplayImage(m, server_name) }
func (m *MagickWand) DisplayImages(server_name string) bool {
	return MagickDisplayImages(m, server_name)
}
func (m *MagickWand) DistortImage(method DistortImageMethod, number_arguments size_t, arguments *float64, bestfit bool) bool {
	return MagickDistortImage(m, method, number_arguments, arguments, bestfit)
}
func (m *MagickWand) DrawImage(d *DrawingWand) bool { return MagickDrawImage(m, d) }
func (m *MagickWand) EdgeImage(radius float64) bool { return MagickEdgeImage(m, radius) }
func (m *MagickWand) EmbossImage(radius, sigma float64) bool {
	return MagickEmbossImage(m, radius, sigma)
}
func (m *MagickWand) EncipherImage(passphrase string) bool { return MagickEncipherImage(m, passphrase) }
func (m *MagickWand) EnhanceImage() bool                   { return MagickEnhanceImage(m) }
func (m *MagickWand) EqualizeImage() bool                  { return MagickEqualizeImage(m) }
func (m *MagickWand) EqualizeImageChannel(channel ChannelType) bool {
	return MagickEqualizeImageChannel(m, channel)
}
func (m *MagickWand) EvaluateImage(operator MagickEvaluateOperator, value float64) bool {
	return MagickEvaluateImage(m, operator, value)
}
func (m *MagickWand) EvaluateImageChannel(channel ChannelType, op MagickEvaluateOperator, value float64) bool {
	return MagickEvaluateImageChannel(m, channel, op, value)
}
func (m *MagickWand) EvaluateImages(operator MagickEvaluateOperator) bool {
	return MagickEvaluateImages(m, operator)
}
func (m *MagickWand) ExportImagePixels(x, y ssize_t, columns, rows size_t, map_ string, storage StorageType, pixels *Void) bool {
	return MagickExportImagePixels(m, x, y, columns, rows, map_, storage, pixels)
}
func (m *MagickWand) ExtentImage(width, height size_t, x, y ssize_t) bool {
	return MagickExtentImage(m, width, height, x, y)
}
func (m *MagickWand) FilterImage(kernel *KernelInfo) bool { return MagickFilterImage(m, kernel) }
func (m *MagickWand) FilterImageChannel(channel ChannelType, kernel *KernelInfo) bool {
	return MagickFilterImageChannel(m, channel, kernel)
}
func (m *MagickWand) FlattenImages() *MagickWand { return MagickFlattenImages(m) }
func (m *MagickWand) FlipImage() bool            { return MagickFlipImage(m) }
func (m *MagickWand) FloodfillPaintImage(channel ChannelType, fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y ssize_t, invert bool) bool {
	return MagickFloodfillPaintImage(m, channel, fill, fuzz, bordercolor, x, y, invert)
}
func (m *MagickWand) FlopImage() bool { return MagickFlopImage(m) }
func (m *MagickWand) ForwardFourierTransformImage(magnitude bool) bool {
	return MagickForwardFourierTransformImage(m, magnitude)
}
func (m *MagickWand) FrameImage(matte_color *PixelWand, width, height size_t, inner_bevel, outer_bevel ssize_t) bool {
	return MagickFrameImage(m, matte_color, width, height, inner_bevel, outer_bevel)
}
func (m *MagickWand) FunctionImage(function MagickFunction, number_arguments size_t, arguments *float64) bool {
	return MagickFunctionImage(m, function, number_arguments, arguments)
}
func (m *MagickWand) FunctionImageChannel(channel ChannelType, function MagickFunction, number_arguments size_t, arguments *float64) bool {
	return MagickFunctionImageChannel(m, channel, function, number_arguments, arguments)
}
func (m *MagickWand) FxImage(expression string) *MagickWand { return MagickFxImage(m, expression) }
func (m *MagickWand) FxImageChannel(channel ChannelType, expression string) *MagickWand {
	return MagickFxImageChannel(m, channel, expression)
}
func (m *MagickWand) GammaImage(gamma float64) bool { return MagickGammaImage(m, gamma) }
func (m *MagickWand) GammaImageChannel(channel ChannelType, gamma float64) bool {
	return MagickGammaImageChannel(m, channel, gamma)
}
func (m *MagickWand) GaussianBlurImage(radius, sigma float64) bool {
	return MagickGaussianBlurImage(m, radius, sigma)
}
func (m *MagickWand) GaussianBlurImageChannel(channel ChannelType, radius, sigma float64) bool {
	return MagickGaussianBlurImageChannel(m, channel, radius, sigma)
}
func (m *MagickWand) Antialias() bool                          { return MagickGetAntialias(m) }
func (m *MagickWand) BackgroundColor() *PixelWand              { return MagickGetBackgroundColor(m) }
func (m *MagickWand) Colorspace() ColorspaceType               { return MagickGetColorspace(m) }
func (m *MagickWand) Compression() CompressionType             { return MagickGetCompression(m) }
func (m *MagickWand) CompressionQuality() size_t               { return MagickGetCompressionQuality(m) }
func (m *MagickWand) Exception(severity *ExceptionType) string { return MagickGetException(m, severity) }
func (m *MagickWand) ExceptionType() ExceptionType             { return MagickGetExceptionType(m) }
func (m *MagickWand) Filename() string                         { return MagickGetFilename(m) }
func (m *MagickWand) Font() string                             { return MagickGetFont(m) }
func (m *MagickWand) Format() string                           { return MagickGetFormat(m) }
func (m *MagickWand) Gravity() GravityType                     { return MagickGetGravity(m) }
func (m *MagickWand) Image() *MagickWand                       { return MagickGetImage(m) }
func (m *MagickWand) ImageAlphaChannel() size_t                { return MagickGetImageAlphaChannel(m) }
func (m *MagickWand) ImageArtifact(artifact string) string     { return MagickGetImageArtifact(m, artifact) }
func (m *MagickWand) ImageArtifacts(pattern, number_artifacts *size_t) string {
	return MagickGetImageArtifacts(m, pattern, number_artifacts)
}
func (m *MagickWand) ImageAttribute(property string) string {
	return MagickGetImageAttribute(m, property)
}
func (m *MagickWand) ImageBackgroundColor(background_color *PixelWand) bool {
	return MagickGetImageBackgroundColor(m, background_color)
}
func (m *MagickWand) ImageBlob(length *size_t) *uint8     { return MagickGetImageBlob(m, length) }
func (m *MagickWand) ImageBluePrimary(x, y *float64) bool { return MagickGetImageBluePrimary(m, x, y) }
func (m *MagickWand) ImageBorderColor(border_color *PixelWand) bool {
	return MagickGetImageBorderColor(m, border_color)
}
func (m *MagickWand) ImageChannelDepth(channel ChannelType) size_t {
	return MagickGetImageChannelDepth(m, channel)
}
func (m *MagickWand) ImageChannelDistortion(reference *MagickWand, channel ChannelType, metric MetricType, distortion *float64) bool {
	return MagickGetImageChannelDistortion(m, reference, channel, metric, distortion)
}
func (m *MagickWand) ImageChannelDistortionFIX(reference *MagickWand, metric MetricType) *float64 {
	return MagickGetImageChannelDistortionFIX(m, reference, metric)
}
func (m *MagickWand) ImageChannelFeatures(distance size_t) *ChannelFeatures {
	return MagickGetImageChannelFeatures(m, distance)
}
func (m *MagickWand) ImageChannelKurtosis(channel ChannelType, kurtosis, skewness *float64) bool {
	return MagickGetImageChannelKurtosis(m, channel, kurtosis, skewness)
}
func (m *MagickWand) ImageChannelMean(channel ChannelType, mean, standard_deviation *float64) bool {
	return MagickGetImageChannelMean(m, channel, mean, standard_deviation)
}
func (m *MagickWand) ImageChannelRange(channel ChannelType, minima, maxima *float64) bool {
	return MagickGetImageChannelRange(m, channel, minima, maxima)
}
func (m *MagickWand) ImageChannelStatistics() *ChannelStatistics {
	return MagickGetImageChannelStatistics(m)
}
func (m *MagickWand) ImageClipMask() *MagickWand { return MagickGetImageClipMask(m) }
func (m *MagickWand) ImageColormapColor(index size_t, color *PixelWand) bool {
	return MagickGetImageColormapColor(m, index, color)
}
func (m *MagickWand) ImageColors() size_t               { return MagickGetImageColors(m) }
func (m *MagickWand) ImageColorspace() ColorspaceType   { return MagickGetImageColorspace(m) }
func (m *MagickWand) ImageCompose() CompositeOperator   { return MagickGetImageCompose(m) }
func (m *MagickWand) ImageCompression() CompressionType { return MagickGetImageCompression(m) }
func (m *MagickWand) ImageCompressionQuality() size_t   { return MagickGetImageCompressionQuality(m) }
func (m *MagickWand) ImageDelay() size_t                { return MagickGetImageDelay(m) }
func (m *MagickWand) ImageDepth() size_t                { return MagickGetImageDepth(m) }
func (m *MagickWand) ImageDispose() DisposeType         { return MagickGetImageDispose(m) }
func (m *MagickWand) ImageDistortion(reference *MagickWand, metric MetricType, distortion *float64) bool {
	return MagickGetImageDistortion(m, reference, metric, distortion)
}
func (m *MagickWand) ImageFilename() string                { return MagickGetImageFilename(m) }
func (m *MagickWand) ImageFormat() string                  { return MagickGetImageFormat(m) }
func (m *MagickWand) ImageFuzz() float64                   { return MagickGetImageFuzz(m) }
func (m *MagickWand) ImageGamma() float64                  { return MagickGetImageGamma(m) }
func (m *MagickWand) ImageGravity() GravityType            { return MagickGetImageGravity(m) }
func (m *MagickWand) ImageGreenPrimary(x, y *float64) bool { return MagickGetImageGreenPrimary(m, x, y) }
func (m *MagickWand) ImageHeight() size_t                  { return MagickGetImageHeight(m) }
func (m *MagickWand) ImageHistogram(number_colors *size_t) **PixelWand {
	return MagickGetImageHistogram(m, number_colors)
}
func (m *MagickWand) ImageInterlaceScheme() InterlaceType { return MagickGetImageInterlaceScheme(m) }
func (m *MagickWand) ImageInterpolateMethod() InterpolatePixelMethod {
	return MagickGetImageInterpolateMethod(m)
}
func (m *MagickWand) ImageIterations() size_t                 { return MagickGetImageIterations(m) }
func (m *MagickWand) ImageLength(length *MagickSizeType) bool { return MagickGetImageLength(m, length) }
func (m *MagickWand) ImageMatte() size_t                      { return MagickGetImageMatte(m) }
func (m *MagickWand) ImageMatteColor(matte_color *PixelWand) bool {
	return MagickGetImageMatteColor(m, matte_color)
}
func (m *MagickWand) ImageOrientation() OrientationType { return MagickGetImageOrientation(m) }
func (m *MagickWand) ImagePage(width, height *size_t, x, y *ssize_t) bool {
	return MagickGetImagePage(m, width, height, x, y)
}
func (m *MagickWand) ImagePixelColor(x, y ssize_t, color *PixelWand) bool {
	return MagickGetImagePixelColor(m, x, y, color)
}
func (m *MagickWand) ImagePixels(x, y ssize_t, columns, rows size_t, map_ string, storage StorageType, pixels *Void) bool {
	return MagickGetImagePixels(m, x, y, columns, rows, map_, storage, pixels)
}
func (m *MagickWand) ImageProfile(name string, length *size_t) *uint8 {
	return MagickGetImageProfile(m, name, length)
}
func (m *MagickWand) ImageProfiles(pattern, number_profiles *size_t) string {
	return MagickGetImageProfiles(m, pattern, number_profiles)
}
func (m *MagickWand) ImageProperties(pattern, number_properties *size_t) string {
	return MagickGetImageProperties(m, pattern, number_properties)
}
func (m *MagickWand) ImageProperty(property string) string { return MagickGetImageProperty(m, property) }
func (m *MagickWand) ImageRedPrimary(x, y *float64) bool   { return MagickGetImageRedPrimary(m, x, y) }
func (m *MagickWand) ImageRegion(width, height size_t, x, y ssize_t) *MagickWand {
	return MagickGetImageRegion(m, width, height, x, y)
}
func (m *MagickWand) ImageRenderingIntent() RenderingIntent { return MagickGetImageRenderingIntent(m) }
func (m *MagickWand) ImageResolution(x, y *float64) bool    { return MagickGetImageResolution(m, x, y) }
func (m *MagickWand) ImagesBlob(length *size_t) *uint8      { return MagickGetImagesBlob(m, length) }
func (m *MagickWand) ImageScene() size_t                    { return MagickGetImageScene(m) }
func (m *MagickWand) ImageSignature() string                { return MagickGetImageSignature(m) }
func (m *MagickWand) ImageSize(length *MagickSizeType) bool { return MagickGetImageSize(m, length) }
func (m *MagickWand) ImageTicksPerSecond() size_t           { return MagickGetImageTicksPerSecond(m) }
func (m *MagickWand) ImageTotalInkDensity() float64         { return MagickGetImageTotalInkDensity(m) }
func (m *MagickWand) ImageType() ImageType                  { return MagickGetImageType(m) }
func (m *MagickWand) ImageUnits() ResolutionType            { return MagickGetImageUnits(m) }
func (m *MagickWand) ImageVirtualPixelMethod() VirtualPixelMethod {
	return MagickGetImageVirtualPixelMethod(m)
}
func (m *MagickWand) ImageWhitePoint(x, y *float64) bool        { return MagickGetImageWhitePoint(m, x, y) }
func (m *MagickWand) ImageWidth() size_t                        { return MagickGetImageWidth(m) }
func (m *MagickWand) InterlaceScheme() InterlaceType            { return MagickGetInterlaceScheme(m) }
func (m *MagickWand) InterpolateMethod() InterpolatePixelMethod { return MagickGetInterpolateMethod(m) }
func (m *MagickWand) IteratorIndex() ssize_t                    { return MagickGetIteratorIndex(m) }
func (m *MagickWand) NumberImages() size_t                      { return MagickGetNumberImages(m) }
func (m *MagickWand) Option(key string) string                  { return MagickGetOption(m, key) }
func (m *MagickWand) Options(pattern, number_options *size_t) string {
	return MagickGetOptions(m, pattern, number_options)
}
func (m *MagickWand) Orientation() OrientationType { return MagickGetOrientation(m) }
func (m *MagickWand) Page(width, height *size_t, x, y *ssize_t) bool {
	return MagickGetPage(m, width, height, x, y)
}
func (m *MagickWand) Pointsize() float64            { return MagickGetPointsize(m) }
func (m *MagickWand) Resolution(x, y *float64) bool { return MagickGetResolution(m, x, y) }
func (m *MagickWand) SamplingFactor(number_factors *size_t) *float64 {
	return MagickGetSamplingFactor(m, number_factors)
}
func (m *MagickWand) Size(columns, rows *size_t) bool { return MagickGetSize(m, columns, rows) }
func (m *MagickWand) SizeOffset(offset *ssize_t) bool { return MagickGetSizeOffset(m, offset) }
func (m *MagickWand) Type() ImageType                 { return MagickGetType(m) }
func (m *MagickWand) HaldClutImage(hald_wand *MagickWand) bool {
	return MagickHaldClutImage(m, hald_wand)
}
func (m *MagickWand) HaldClutImageChannel(channel ChannelType, hald_wand *MagickWand) bool {
	return MagickHaldClutImageChannel(m, channel, hald_wand)
}
func (m *MagickWand) HasNextImage() bool               { return MagickHasNextImage(m) }
func (m *MagickWand) HasPreviousImage() bool           { return MagickHasPreviousImage(m) }
func (m *MagickWand) IdentifyImage() string            { return MagickIdentifyImage(m) }
func (m *MagickWand) ImplodeImage(radius float64) bool { return MagickImplodeImage(m, radius) }
func (m *MagickWand) ImportImagePixels(x, y ssize_t, columns, rows size_t, map_ string, storage StorageType, pixels *Void) bool {
	return MagickImportImagePixels(m, x, y, columns, rows, map_, storage, pixels)
}
func (m *MagickWand) InverseFourierTransformImage(phase_wand *MagickWand, magnitude bool) bool {
	return MagickInverseFourierTransformImage(m, phase_wand, magnitude)
}
func (m *MagickWand) LabelImage(label string) bool { return MagickLabelImage(m, label) }
func (m *MagickWand) LevelImage(black_point, gamma, white_point float64) bool {
	return MagickLevelImage(m, black_point, gamma, white_point)
}
func (m *MagickWand) LevelImageChannel(channel ChannelType, black_point, gamma, white_point float64) bool {
	return MagickLevelImageChannel(m, channel, black_point, gamma, white_point)
}
func (m *MagickWand) LinearStretchImage(black_point, white_point float64) bool {
	return MagickLinearStretchImage(m, black_point, white_point)
}
func (m *MagickWand) MagnifyImage() bool { return MagickMagnifyImage(m) }
func (m *MagickWand) MapImage(map_wand *MagickWand, dither bool) bool {
	return MagickMapImage(m, map_wand, dither)
}
func (m *MagickWand) MatteFloodfillImage(alpha, fuzz float64, bordercolor *PixelWand, x, y ssize_t) bool {
	return MagickMatteFloodfillImage(m, alpha, fuzz, bordercolor, x, y)
}
func (m *MagickWand) MaximumImages() *MagickWand { return MagickMaximumImages(m) }
func (m *MagickWand) MedianFilterImage(radius float64) bool {
	return MagickMedianFilterImage(m, radius)
}
func (m *MagickWand) MergeImageLayers(method ImageLayerMethod) *MagickWand {
	return MagickMergeImageLayers(m, method)
}
func (m *MagickWand) MinifyImage() bool             { return MagickMinifyImage(m) }
func (m *MagickWand) MinimumImages() *MagickWand    { return MagickMinimumImages(m) }
func (m *MagickWand) ModeImage(radius float64) bool { return MagickModeImage(m, radius) }
func (m *MagickWand) ModulateImage(brightness, saturation, hue float64) bool {
	return MagickModulateImage(m, brightness, saturation, hue)
}
func (m *MagickWand) MontageImage(drawing_wand DrawingWand, tile_geometry, thumbnail_geometry string, mode MontageMode, frame string) *MagickWand {
	return MagickMontageImage(m, drawing_wand, tile_geometry, thumbnail_geometry, mode, frame)
}
func (m *MagickWand) MorphImages(number_frames size_t) *MagickWand {
	return MagickMorphImages(m, number_frames)
}
func (m *MagickWand) MorphologyImage(method MorphologyMethod, iterations ssize_t, kernel *KernelInfo) bool {
	return MagickMorphologyImage(m, method, iterations, kernel)
}
func (m *MagickWand) MorphologyImageChannel(channel ChannelType, method MorphologyMethod, iterations ssize_t, kernel *KernelInfo) bool {
	return MagickMorphologyImageChannel(m, channel, method, iterations, kernel)
}
func (m *MagickWand) MosaicImages() *MagickWand { return MagickMosaicImages(m) }
func (m *MagickWand) MotionBlurImage(radius, sigma, angle float64) bool {
	return MagickMotionBlurImage(m, radius, sigma, angle)
}
func (m *MagickWand) MotionBlurImageChannel(channel ChannelType, radius, sigma, angle float64) bool {
	return MagickMotionBlurImageChannel(m, channel, radius, sigma, angle)
}
func (m *MagickWand) NegateImage(gray bool) bool { return MagickNegateImage(m, gray) }
func (m *MagickWand) NegateImageChannel(channel ChannelType, gray bool) bool {
	return MagickNegateImageChannel(m, channel, gray)
}
func (m *MagickWand) NewImage(columns, rows size_t, background *PixelWand) bool {
	return MagickNewImage(m, columns, rows, background)
}
func (m *MagickWand) NextImage() bool      { return MagickNextImage(m) }
func (m *MagickWand) NormalizeImage() bool { return MagickNormalizeImage(m) }
func (m *MagickWand) NormalizeImageChannel(channel ChannelType) bool {
	return MagickNormalizeImageChannel(m, channel)
}
func (m *MagickWand) OilPaintImage(radius float64) bool { return MagickOilPaintImage(m, radius) }
func (m *MagickWand) OpaqueImage(target, fill *PixelWand, fuzz float64) bool {
	return MagickOpaqueImage(m, target, fill, fuzz)
}
func (m *MagickWand) OpaquePaintImage(target, fill *PixelWand, fuzz float64, invert bool) bool {
	return MagickOpaquePaintImage(m, target, fill, fuzz, invert)
}
func (m *MagickWand) OpaquePaintImageChannel(channel ChannelType, target, fill *PixelWand, fuzz float64, invert bool) bool {
	return MagickOpaquePaintImageChannel(m, channel, target, fill, fuzz, invert)
}
func (m *MagickWand) OptimizeImageLayers() *MagickWand { return MagickOptimizeImageLayers(m) }
func (m *MagickWand) OrderedPosterizeImage(threshold_map string) bool {
	return MagickOrderedPosterizeImage(m, threshold_map)
}
func (m *MagickWand) OrderedPosterizeImageChannel(channel ChannelType, threshold_map string) bool {
	return MagickOrderedPosterizeImageChannel(m, channel, threshold_map)
}
func (m *MagickWand) PaintFloodfillImage(channel ChannelType, fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y ssize_t) bool {
	return MagickPaintFloodfillImage(m, channel, fill, fuzz, bordercolor, x, y)
}
func (m *MagickWand) PaintOpaqueImage(target, fill *PixelWand, fuzz float64) bool {
	return MagickPaintOpaqueImage(m, target, fill, fuzz)
}
func (m *MagickWand) PaintOpaqueImageChannel(channel ChannelType, target, fill *PixelWand, fuzz float64) bool {
	return MagickPaintOpaqueImageChannel(m, channel, target, fill, fuzz)
}
func (m *MagickWand) PaintTransparentImage(target *PixelWand, alpha, fuzz float64) bool {
	return MagickPaintTransparentImage(m, target, alpha, fuzz)
}
func (m *MagickWand) PingImage(filename string) bool { return MagickPingImage(m, filename) }
func (m *MagickWand) PingImageBlob(blob *Void, length size_t) bool {
	return MagickPingImageBlob(m, blob, length)
}
func (m *MagickWand) PingImageFile(file *FILE) bool { return MagickPingImageFile(m, file) }
func (m *MagickWand) PolaroidImage(d *DrawingWand, angle float64) bool {
	return MagickPolaroidImage(m, d, angle)
}
func (m *MagickWand) PosterizeImage(levels size_t, dither bool) bool {
	return MagickPosterizeImage(m, levels, dither)
}
func (m *MagickWand) PreviewImages(preview PreviewType) *MagickWand {
	return MagickPreviewImages(m, preview)
}
func (m *MagickWand) PreviousImage() bool { return MagickPreviousImage(m) }
func (m *MagickWand) ProfileImage(name string, profile *Void, length size_t) bool {
	return MagickProfileImage(m, name, profile, length)
}
func (m *MagickWand) QuantizeImage(number_colors size_t, colorspace ColorspaceType, treedepth size_t, dither, measure_error bool) bool {
	return MagickQuantizeImage(m, number_colors, colorspace, treedepth, dither, measure_error)
}
func (m *MagickWand) QuantizeImages(number_colors size_t, colorspace ColorspaceType, treedepth size_t, dither, measure_error bool) bool {
	return MagickQuantizeImages(m, number_colors, colorspace, treedepth, dither, measure_error)
}
func (m *MagickWand) QueryFontMetrics(drawingWand *DrawingWand, text string) []float64 {
	return MagickQueryFontMetrics(m, drawingWand, text)
}
func (m *MagickWand) QueryMultilineFontMetrics(drawingWand *DrawingWand, text string) []float64 {
	return MagickQueryMultilineFontMetrics(m, drawingWand, text)
}
func (m *MagickWand) RadialBlurImage(angle float64) bool { return MagickRadialBlurImage(m, angle) }
func (m *MagickWand) RadialBlurImageChannel(channel ChannelType, angle float64) bool {
	return MagickRadialBlurImageChannel(m, channel, angle)
}
func (m *MagickWand) RaiseImage(width, height size_t, x, y ssize_t, raise bool) bool {
	return MagickRaiseImage(m, width, height, x, y, raise)
}
func (m *MagickWand) RandomThresholdImage(low, high float64) bool {
	return MagickRandomThresholdImage(m, low, high)
}
func (m *MagickWand) RandomThresholdImageChannel(channel ChannelType, low, high float64) bool {
	return MagickRandomThresholdImageChannel(m, channel, low, high)
}
func (m *MagickWand) ReadImage(filename string) bool { return MagickReadImage(m, filename) }
func (m *MagickWand) ReadImageBlob(blob *Void, length size_t) bool {
	return MagickReadImageBlob(m, blob, length)
}
func (m *MagickWand) ReadImageFile(file *FILE) bool { return MagickReadImageFile(m, file) }
func (m *MagickWand) RecolorImage(order size_t, color_matrix *float64) bool {
	return MagickRecolorImage(m, order, color_matrix)
}
func (m *MagickWand) ReduceNoiseImage(radius float64) bool { return MagickReduceNoiseImage(m, radius) }
func (m *MagickWand) RegionOfInterestImage(width, height size_t, x, y ssize_t) *MagickWand {
	return MagickRegionOfInterestImage(m, width, height, x, y)
}
func (m *MagickWand) RemapImage(remap_wand *MagickWand, method DitherMethod) bool {
	return MagickRemapImage(m, remap_wand, method)
}
func (m *MagickWand) RemoveImage() bool { return MagickRemoveImage(m) }
func (m *MagickWand) RemoveImageProfile(name string, length *size_t) *uint8 {
	return MagickRemoveImageProfile(m, name, length)
}
func (m *MagickWand) ResampleImage(x_resolution, y_resolution float64, filter FilterTypes, blur float64) bool {
	return MagickResampleImage(m, x_resolution, y_resolution, filter, blur)
}
func (m *MagickWand) ResetImagePage(page string) bool { return MagickResetImagePage(m, page) }
func (m *MagickWand) ResetIterator()                  { MagickResetIterator(m) }
func (m *MagickWand) ResizeImage(columns, rows size_t, filter FilterTypes, blur float64) bool {
	return MagickResizeImage(m, columns, rows, filter, blur)
}
func (m *MagickWand) RollImage(x ssize_t, y size_t) bool { return MagickRollImage(m, x, y) }
func (m *MagickWand) RotateImage(background *PixelWand, degrees float64) bool {
	return MagickRotateImage(m, background, degrees)
}
func (m *MagickWand) SampleImage(columns, rows size_t) bool {
	return MagickSampleImage(m, columns, rows)
}
func (m *MagickWand) ScaleImage(columns, rows size_t) bool { return MagickScaleImage(m, columns, rows) }
func (m *MagickWand) SegmentImage(colorspace ColorspaceType, verbose bool, cluster_threshold, smooth_threshold float64) bool {
	return MagickSegmentImage(m, colorspace, verbose, cluster_threshold, smooth_threshold)
}
func (m *MagickWand) SelectiveBlurImage(radius, sigma, threshold float64) bool {
	return MagickSelectiveBlurImage(m, radius, sigma, threshold)
}
func (m *MagickWand) SelectiveBlurImageChannel(channel ChannelType, radius, sigma, threshold float64) bool {
	return MagickSelectiveBlurImageChannel(m, channel, radius, sigma, threshold)
}
func (m *MagickWand) SeparateImageChannel(channel ChannelType) bool {
	return MagickSeparateImageChannel(m, channel)
}
func (m *MagickWand) SepiaToneImage(threshold float64) bool {
	return MagickSepiaToneImage(m, threshold)
}
func (m *MagickWand) SetAntialias(antialias bool) bool { return MagickSetAntialias(m, antialias) }
func (m *MagickWand) SetBackgroundColor(background *PixelWand) bool {
	return MagickSetBackgroundColor(m, background)
}
func (m *MagickWand) SetColorspace(colorspace ColorspaceType) bool {
	return MagickSetColorspace(m, colorspace)
}
func (m *MagickWand) SetCompression(compression CompressionType) bool {
	return MagickSetCompression(m, compression)
}
func (m *MagickWand) SetCompressionQuality(quality size_t) bool {
	return MagickSetCompressionQuality(m, quality)
}
func (m *MagickWand) SetDepth(depth size_t) bool         { return MagickSetDepth(m, depth) }
func (m *MagickWand) SetExtract(geometry string) bool    { return MagickSetExtract(m, geometry) }
func (m *MagickWand) SetFilename(filename string) bool   { return MagickSetFilename(m, filename) }
func (m *MagickWand) SetFirstIterator()                  { MagickSetFirstIterator(m) }
func (m *MagickWand) SetFont(font string) bool           { return MagickSetFont(m, font) }
func (m *MagickWand) SetFormat(format string) bool       { return MagickSetFormat(m, format) }
func (m *MagickWand) SetGravity(type_ GravityType) bool  { return MagickSetGravity(m, type_) }
func (m *MagickWand) SetImage(set_wand *MagickWand) bool { return MagickSetImage(m, set_wand) }
func (m *MagickWand) SetImageAlphaChannel(alpha_type AlphaChannelType) bool {
	return MagickSetImageAlphaChannel(m, alpha_type)
}
func (m *MagickWand) SetImageArtifact(artifact, value string) bool {
	return MagickSetImageArtifact(m, artifact, value)
}
func (m *MagickWand) SetImageAttribute(property, value string) bool {
	return MagickSetImageAttribute(m, property, value)
}
func (m *MagickWand) SetImageBackgroundColor(background *PixelWand) bool {
	return MagickSetImageBackgroundColor(m, background)
}
func (m *MagickWand) SetImageBias(bias float64) bool { return MagickSetImageBias(m, bias) }
func (m *MagickWand) SetImageBluePrimary(x, y float64) bool {
	return MagickSetImageBluePrimary(m, x, y)
}
func (m *MagickWand) SetImageBorderColor(border *PixelWand) bool {
	return MagickSetImageBorderColor(m, border)
}
func (m *MagickWand) SetImageChannelDepth(channel ChannelType, depth size_t) bool {
	return MagickSetImageChannelDepth(m, channel, depth)
}
func (m *MagickWand) SetImageClipMask(clip_mask *MagickWand) bool {
	return MagickSetImageClipMask(m, clip_mask)
}
func (m *MagickWand) SetImageColor(color *PixelWand) bool { return MagickSetImageColor(m, color) }
func (m *MagickWand) SetImageColormapColor(index size_t, color *PixelWand) bool {
	return MagickSetImageColormapColor(m, index, color)
}
func (m *MagickWand) SetImageColorspace(colorspace ColorspaceType) bool {
	return MagickSetImageColorspace(m, colorspace)
}
func (m *MagickWand) SetImageCompose(compose CompositeOperator) bool {
	return MagickSetImageCompose(m, compose)
}
func (m *MagickWand) SetImageCompression(compression CompressionType) bool {
	return MagickSetImageCompression(m, compression)
}
func (m *MagickWand) SetImageCompressionQuality(quality size_t) bool {
	return MagickSetImageCompressionQuality(m, quality)
}
func (m *MagickWand) SetImageDelay(delay size_t) bool { return MagickSetImageDelay(m, delay) }
func (m *MagickWand) SetImageDepth(depth size_t) bool { return MagickSetImageDepth(m, depth) }
func (m *MagickWand) SetImageDispose(dispose DisposeType) bool {
	return MagickSetImageDispose(m, dispose)
}
func (m *MagickWand) SetImageEndian(endian EndianType) bool { return MagickSetImageEndian(m, endian) }
func (m *MagickWand) SetImageExtent(columns size_t, rows uint) bool {
	return MagickSetImageExtent(m, columns, rows)
}
func (m *MagickWand) SetImageFilename(filename string) bool {
	return MagickSetImageFilename(m, filename)
}
func (m *MagickWand) SetImageFormat(format string) bool { return MagickSetImageFormat(m, format) }
func (m *MagickWand) SetImageFuzz(fuzz float64) bool    { return MagickSetImageFuzz(m, fuzz) }
func (m *MagickWand) SetImageGamma(gamma float64) bool  { return MagickSetImageGamma(m, gamma) }
func (m *MagickWand) SetImageGravity(gravity GravityType) bool {
	return MagickSetImageGravity(m, gravity)
}
func (m *MagickWand) SetImageGreenPrimary(x, y float64) bool {
	return MagickSetImageGreenPrimary(m, x, y)
}
func (m *MagickWand) SetImageIndex(index ssize_t) bool { return MagickSetImageIndex(m, index) }
func (m *MagickWand) SetImageInterlaceScheme(interlace InterlaceType) bool {
	return MagickSetImageInterlaceScheme(m, interlace)
}
func (m *MagickWand) SetImageInterpolateMethod(method InterpolatePixelMethod) bool {
	return MagickSetImageInterpolateMethod(m, method)
}
func (m *MagickWand) SetImageIterations(iterations size_t) bool {
	return MagickSetImageIterations(m, iterations)
}
func (m *MagickWand) SetImageMatteColor(matte *bool) bool { return MagickSetImageMatteColor(m, matte) }
func (m *MagickWand) SetImageMatteColorFIX(matte *PixelWand) bool {
	return MagickSetImageMatteColorFIX(m, matte)
}
func (m *MagickWand) SetImageOpacity(alpha float64) bool { return MagickSetImageOpacity(m, alpha) }
func (m *MagickWand) SetImageOrientation(orientation OrientationType) bool {
	return MagickSetImageOrientation(m, orientation)
}
func (m *MagickWand) SetImagePage(width, height size_t, x, y ssize_t) bool {
	return MagickSetImagePage(m, width, height, x, y)
}
func (m *MagickWand) SetImagePixels(x, y ssize_t, columns, rows size_t, map_ string, storage StorageType, pixels *Void) bool {
	return MagickSetImagePixels(m, x, y, columns, rows, map_, storage, pixels)
}
func (m *MagickWand) SetImageProfile(name string, profile *Void, length size_t) bool {
	return MagickSetImageProfile(m, name, profile, length)
}
func (m *MagickWand) SetImageProgressMonitor(progress_monitor, client_data *Void) MagickProgressMonitorType {
	return MagickSetImageProgressMonitor(m, progress_monitor, client_data)
}
func (m *MagickWand) SetImageProperty(property, value string) bool {
	return MagickSetImageProperty(m, property, value)
}
func (m *MagickWand) SetImageRedPrimary(x, y float64) bool { return MagickSetImageRedPrimary(m, x, y) }
func (m *MagickWand) SetImageRenderingIntent(rendering_intent RenderingIntent) bool {
	return MagickSetImageRenderingIntent(m, rendering_intent)
}
func (m *MagickWand) SetImageResolution(x_resolution, y_resolution float64) bool {
	return MagickSetImageResolution(m, x_resolution, y_resolution)
}
func (m *MagickWand) SetImageScene(scene size_t) bool { return MagickSetImageScene(m, scene) }
func (m *MagickWand) SetImageTicksPerSecond(ticks_per_second ssize_t) bool {
	return MagickSetImageTicksPerSecond(m, ticks_per_second)
}
func (m *MagickWand) SetImageType(image_type ImageType) bool {
	return MagickSetImageType(m, image_type)
}
func (m *MagickWand) SetImageUnits(units ResolutionType) bool { return MagickSetImageUnits(m, units) }
func (m *MagickWand) SetImageVirtualPixelMethod(method VirtualPixelMethod) VirtualPixelMethod {
	return MagickSetImageVirtualPixelMethod(m, method)
}
func (m *MagickWand) SetImageWhitePoint(x, y float64) bool { return MagickSetImageWhitePoint(m, x, y) }
func (m *MagickWand) SetInterlaceScheme(interlace_scheme InterlaceType) bool {
	return MagickSetInterlaceScheme(m, interlace_scheme)
}
func (m *MagickWand) SetInterpolateMethod(method InterpolateMethodPixel) bool {
	return MagickSetInterpolateMethod(m, method)
}
func (m *MagickWand) SetIteratorIndex(index ssize_t) bool { return MagickSetIteratorIndex(m, index) }
func (m *MagickWand) SetLastIterator()                    { MagickSetLastIterator(m) }
func (m *MagickWand) SetOption(key, value string) bool    { return MagickSetOption(m, key, value) }
func (m *MagickWand) SetOrientation(orientation OrientationType) bool {
	return MagickSetOrientation(m, orientation)
}
func (m *MagickWand) SetPage(width, height size_t, x, y ssize_t) bool {
	return MagickSetPage(m, width, height, x, y)
}
func (m *MagickWand) SetPassphrase(passphrase string) bool { return MagickSetPassphrase(m, passphrase) }
func (m *MagickWand) SetPointsize(pointsize float64) bool  { return MagickSetPointsize(m, pointsize) }
func (m *MagickWand) SetProgressMonitor(progress_monitor, client_data *Void) MagickProgressMonitorType {
	return MagickSetProgressMonitor(m, progress_monitor, client_data)
}
func (m *MagickWand) SetResolution(x_resolution, y_resolution float64) bool {
	return MagickSetResolution(m, x_resolution, y_resolution)
}
func (m *MagickWand) SetSamplingFactors(number_factors size_t, sampling_factors *float64) bool {
	return MagickSetSamplingFactors(m, number_factors, sampling_factors)
}
func (m *MagickWand) SetSize(columns, rows size_t) bool { return MagickSetSize(m, columns, rows) }
func (m *MagickWand) SetSizeOffset(columns, rows size_t, offset ssize_t) bool {
	return MagickSetSizeOffset(m, columns, rows, offset)
}
func (m *MagickWand) SetType(image_type ImageType) bool { return MagickSetType(m, image_type) }
func (m *MagickWand) ShadeImage(gray bool, azimuth, elevation float64) bool {
	return MagickShadeImage(m, gray, azimuth, elevation)
}
func (m *MagickWand) ShadowImage(opacity, sigma float64, x, y ssize_t) bool {
	return MagickShadowImage(m, opacity, sigma, x, y)
}
func (m *MagickWand) SharpenImage(radius, sigma float64) bool {
	return MagickSharpenImage(m, radius, sigma)
}
func (m *MagickWand) SharpenImageChannel(channel ChannelType, radius, sigma float64) bool {
	return MagickSharpenImageChannel(m, channel, radius, sigma)
}
func (m *MagickWand) ShaveImage(columns, rows size_t) bool { return MagickShaveImage(m, columns, rows) }
func (m *MagickWand) ShearImage(background *PixelWand, x_shear, y_shear float64) bool {
	return MagickShearImage(m, background, x_shear, y_shear)
}
func (m *MagickWand) SigmoidalContrastImage(sharpen bool, alpha, beta float64) bool {
	return MagickSigmoidalContrastImage(m, sharpen, alpha, beta)
}
func (m *MagickWand) SigmoidalContrastImageChannel(channel ChannelType, sharpen bool, alpha, beta float64) bool {
	return MagickSigmoidalContrastImageChannel(m, channel, sharpen, alpha, beta)
}
func (m *MagickWand) SimilarityImage(reference *MagickWand, offset *RectangeInfo, similarity *float64) *MagickWand {
	return MagickSimilarityImage(m, reference, offset, similarity)
}
func (m *MagickWand) SketchImage(radius, sigma, angle float64) bool {
	return MagickSketchImage(m, radius, sigma, angle)
}
func (m *MagickWand) SmushImages(stack bool, offset ssize_t) *MagickWand {
	return MagickSmushImages(m, stack, offset)
}
func (m *MagickWand) SolarizeImage(threshold float64) bool { return MagickSolarizeImage(m, threshold) }
func (m *MagickWand) SparseColorImage(channel ChannelType, method SparseColorMethod, number_arguments size_t, arguments *float64) bool {
	return MagickSparseColorImage(m, channel, method, number_arguments, arguments)
}
func (m *MagickWand) SpliceImage(width, height size_t, x, y ssize_t) bool {
	return MagickSpliceImage(m, width, height, x, y)
}
func (m *MagickWand) SpreadImage(radius float64) bool { return MagickSpreadImage(m, radius) }
func (m *MagickWand) StatisticImage(type_ StatisticType, width float64, height size_t) bool {
	return MagickStatisticImage(m, type_, width, height)
}
func (m *MagickWand) StatisticImageChannel(channel ChannelType, type_ StatisticType, width float64, height size_t) bool {
	return MagickStatisticImageChannel(m, channel, type_, width, height)
}
func (m *MagickWand) SteganoImage(watermark_wand, offset ssize_t) *MagickWand {
	return MagickSteganoImage(m, watermark_wand, offset)
}
func (m *MagickWand) StereoImage(offset_wand *MagickWand) *MagickWand {
	return MagickStereoImage(m, offset_wand)
}
func (m *MagickWand) StripImage() bool                { return MagickStripImage(m) }
func (m *MagickWand) SwirlImage(degrees float64) bool { return MagickSwirlImage(m, degrees) }
func (m *MagickWand) TextureImage(texture_wand *MagickWand) *MagickWand {
	return MagickTextureImage(m, texture_wand)
}
func (m *MagickWand) ThresholdImage(threshold float64) bool {
	return MagickThresholdImage(m, threshold)
}
func (m *MagickWand) ThresholdImageChannel(channel ChannelType, threshold float64) bool {
	return MagickThresholdImageChannel(m, channel, threshold)
}
func (m *MagickWand) ThumbnailImage(columns, rows size_t) bool {
	return MagickThumbnailImage(m, columns, rows)
}
func (m *MagickWand) TintImage(tint, opacity *PixelWand) bool {
	return MagickTintImage(m, tint, opacity)
}
func (m *MagickWand) TransformImage(crop, geometry string) *MagickWand {
	return MagickTransformImage(m, crop, geometry)
}
func (m *MagickWand) TransformImageColorspace(colorspace ColorspaceType) bool {
	return MagickTransformImageColorspace(m, colorspace)
}
func (m *MagickWand) TransparentImage(target *PixelWand, alpha, fuzz float64) bool {
	return MagickTransparentImage(m, target, alpha, fuzz)
}
func (m *MagickWand) TransparentPaintImage(target *PixelWand, alpha, fuzz float64, invert bool) bool {
	return MagickTransparentPaintImage(m, target, alpha, fuzz, invert)
}
func (m *MagickWand) TransposeImage() bool        { return MagickTransposeImage(m) }
func (m *MagickWand) TransverseImage() bool       { return MagickTransverseImage(m) }
func (m *MagickWand) TrimImage(fuzz float64) bool { return MagickTrimImage(m, fuzz) }
func (m *MagickWand) UniqueImageColors() bool     { return MagickUniqueImageColors(m) }
func (m *MagickWand) UnsharpMaskImage(radius, sigma, amount, threshold float64) bool {
	return MagickUnsharpMaskImage(m, radius, sigma, amount, threshold)
}
func (m *MagickWand) UnsharpMaskImageChannel(channel ChannelType, radius, sigma, amount, threshold float64) bool {
	return MagickUnsharpMaskImageChannel(m, channel, radius, sigma, amount, threshold)
}
func (m *MagickWand) VignetteImage(black_point, white_point float64, x, y ssize_t) bool {
	return MagickVignetteImage(m, black_point, white_point, x, y)
}
func (m *MagickWand) WaveImage(amplitude, wave_length float64) bool {
	return MagickWaveImage(m, amplitude, wave_length)
}
func (m *MagickWand) WhiteThresholdImage(threshold *PixelWand) bool {
	return MagickWhiteThresholdImage(m, threshold)
}
func (m *MagickWand) WriteImage(filename string) bool      { return MagickWriteImage(m, filename) }
func (m *MagickWand) WriteImageBlob(length *size_t) *uint8 { return MagickWriteImageBlob(m, length) }
func (m *MagickWand) WriteImageFile(file *FILE) bool       { return MagickWriteImageFile(m, file) }
func (m *MagickWand) WriteImages(filename string, adjoin bool) bool {
	return MagickWriteImages(m, filename, adjoin)
}
func (m *MagickWand) WriteImagesFile(file *FILE) bool { return MagickWriteImagesFile(m, file) }

var (
	NewDrawingWand func() *DrawingWand

	ClearDrawingWand   func(d *DrawingWand)
	CloneDrawingWand   func(d *DrawingWand) *DrawingWand
	DestroyDrawingWand func(d *DrawingWand) *DrawingWand
	IsDrawingWand      func(d *DrawingWand) bool
	PeekDrawingWand    func(d *DrawingWand) *DrawInfo
	PopDrawingWand     func(d *DrawingWand) bool
	PushDrawingWand    func(d *DrawingWand) bool

	DrawAffine                                   func(d *DrawingWand, affine *AffineMatrix)
	DrawAnnotation                               func(d *DrawingWand, x, y float64, text *uint8)
	DrawArc                                      func(d *DrawingWand, sx, sy, ex, ey, sd, ed float64)
	DrawBezier                                   func(d *DrawingWand, number_coordinates size_t, coordinates *PointInfo)
	DrawCircle                                   func(d *DrawingWand, ox, oy, px, py float64)
	DrawClearException                           func(d *DrawingWand) bool
	DrawColor                                    func(d *DrawingWand, x, y float64, paint_method PaintMethod)
	DrawComment                                  func(d *DrawingWand, comment string)
	DrawComposite                                func(d *DrawingWand, compose CompositeOperator, x, y, width, height float64, magick_wand *MagickWand) bool
	DrawEllipse                                  func(d *DrawingWand, ox, oy, rx, ry, start, end float64)
	DrawGetBorderColor                           func(d *DrawingWand, border_color *PixelWand)
	DrawGetClipPath                              func(d *DrawingWand) string
	DrawGetClipUnits                             func(d *DrawingWand) ClipPathUnits
	DrawGetException                             func(d *DrawingWand, severity *ExceptionType) string
	DrawGetExceptionType                         func(d *DrawingWand) ExceptionType
	DrawGetFillAlpha                             func(d *DrawingWand) float64
	DrawGetFillColor                             func(d *DrawingWand, fill_color *PixelWand)
	DrawGetFillOpacity                           func(d *DrawingWand) float64
	DrawGetFillRule                              func(d *DrawingWand) FillRule
	DrawGetFont                                  func(d *DrawingWand) string
	DrawGetFontFamily                            func(d *DrawingWand) string
	DrawGetFontResolution                        func(d *DrawingWand, x, y *float64) DrawBooleanType
	DrawGetFontSize                              func(d *DrawingWand) float64
	DrawGetFontStretch                           func(d *DrawingWand) StretchType
	DrawGetFontStyle                             func(d *DrawingWand) StyleType
	DrawGetFontWeight                            func(d *DrawingWand) size_t
	DrawGetGravity                               func(d *DrawingWand) GravityType
	DrawGetOpacity                               func(d *DrawingWand) float64
	DrawGetStrokeAlpha                           func(d *DrawingWand) float64
	DrawGetStrokeAntialias                       func(d *DrawingWand) bool
	DrawGetStrokeColor                           func(d *DrawingWand, stroke_color *PixelWand)
	DrawGetStrokeDashArray                       func(d *DrawingWand, number_elements *size_t) *float64
	DrawGetStrokeDashOffset                      func(d *DrawingWand) float64
	DrawGetStrokeLineCap                         func(d *DrawingWand) LineCap
	DrawGetStrokeLineJoin                        func(d *DrawingWand) LineJoin
	DrawGetStrokeMiterLimit                      func(d *DrawingWand) size_t
	DrawGetStrokeOpacity                         func(d *DrawingWand) float64
	DrawGetStrokeWidth                           func(d *DrawingWand) float64
	DrawGetTextAlignment                         func(d *DrawingWand) AlignType
	DrawGetTextAntialias                         func(d *DrawingWand) bool
	DrawGetTextDecoration                        func(d *DrawingWand) DecorationType
	DrawGetTextEncoding                          func(d *DrawingWand) string
	DrawGetTextInterwordSpacing                  func(d *DrawingWand) float64
	DrawGetTextKerning                           func(d *DrawingWand) float64
	DrawGetTextUnderColor                        func(d *DrawingWand, under_color *PixelWand)
	DrawGetVectorGraphics                        func(d *DrawingWand) string
	DrawLine                                     func(d *DrawingWand, sx, sy, ex, ey float64)
	DrawMatte                                    func(d *DrawingWand, x, y float64, paint_method PaintMethod)
	DrawPathClose                                func(d *DrawingWand)
	DrawPathCurveToAbsolute                      func(d *DrawingWand, x1, y1, x2, y2, x, y float64)
	DrawPathCurveToQuadraticBezierAbsolute       func(d *DrawingWand, x1, y1, x, y float64)
	DrawPathCurveToQuadraticBezierRelative       func(d *DrawingWand, x1, y1, x, y float64)
	DrawPathCurveToQuadraticBezierSmoothAbsolute func(d *DrawingWand, x, y float64)
	DrawPathCurveToQuadraticBezierSmoothRelative func(d *DrawingWand, x, y float64)
	DrawPathCurveToRelative                      func(d *DrawingWand, x1, y1, x2, y2, x, y float64)
	DrawPathCurveToSmoothAbsolute                func(d *DrawingWand, x2, y2, x, y float64)
	DrawPathCurveToSmoothRelative                func(d *DrawingWand, x2, y2, x, y float64)
	DrawPathEllipticArcAbsolute                  func(d *DrawingWand, rx, ry, x_axis_rotation float64, large_arc_flag, sweep_flag bool, x, y float64)
	DrawPathEllipticArcRelative                  func(d *DrawingWand, rx, ry, x_axis_rotation float64, large_arc_flag, sweep_flag bool, x, y float64)
	DrawPathFinish                               func(d *DrawingWand)
	DrawPathLineToAbsolute                       func(d *DrawingWand, x, y float64)
	DrawPathLineToHorizontalAbsolute             func(d *DrawingWand, x float64)
	DrawPathLineToHorizontalRelative             func(d *DrawingWand, x float64)
	DrawPathLineToRelative                       func(d *DrawingWand, x, y float64)
	DrawPathLineToVerticalAbsolute               func(d *DrawingWand, y float64)
	DrawPathLineToVerticalRelative               func(d *DrawingWand, y float64)
	DrawPathMoveToAbsolute                       func(d *DrawingWand, x, y float64)
	DrawPathMoveToRelative                       func(d *DrawingWand, x, y float64)
	DrawPathStart                                func(d *DrawingWand)
	DrawPeekGraphicWand                          func(d *DrawingWand) *DrawInfo
	DrawPoint                                    func(d *DrawingWand, x, y float64)
	DrawPolygon                                  func(d *DrawingWand, number_coordinates size_t, coordinates *PointInfo)
	DrawPolyline                                 func(d *DrawingWand, number_coordinates size_t, coordinates *PointInfo)
	DrawPopClipPath                              func(d *DrawingWand)
	DrawPopDefs                                  func(d *DrawingWand)
	DrawPopGraphicContext                        func(d *DrawingWand) bool
	DrawPopPattern                               func(d *DrawingWand) bool
	DrawPushClipPath                             func(d *DrawingWand, clip_mask_id string)
	DrawPushDefs                                 func(d *DrawingWand)
	DrawPushGraphicContext                       func(d *DrawingWand) bool
	DrawPushPattern                              func(d *DrawingWand, pattern_id string, x, y, width, height float64) bool
	DrawRectangle                                func(d *DrawingWand, x1, y1, x2, y2 float64)
	DrawResetVectorGraphics                      func(d *DrawingWand)
	DrawRotate                                   func(d *DrawingWand, degrees float64)
	DrawRoundRectangle                           func(d *DrawingWand, x1, y1, x2, y2, rx, ry float64)
	DrawScale                                    func(d *DrawingWand, x, y float64)
	DrawSetBorderColor                           func(d *DrawingWand, border_wand *PixelWand)
	DrawSetClipPath                              func(d *DrawingWand, clip_mask string) bool
	DrawSetClipRule                              func(d *DrawingWand, fill_rule FillRule)
	DrawSetClipUnits                             func(d *DrawingWand, clip_units ClipPathUnits)
	DrawSetFillAlpha                             func(d *DrawingWand, fill_alpha float64)
	DrawSetFillColor                             func(d *DrawingWand, fill_wand *PixelWand)
	DrawSetFillOpacity                           func(d *DrawingWand, fill_opacity float64)
	DrawSetFillPatternURL                        func(d *DrawingWand, fill_url string) bool
	DrawSetFillRule                              func(d *DrawingWand, fill_rule FillRule)
	DrawSetFont                                  func(d *DrawingWand, font_name string) bool
	DrawSetFontFamily                            func(d *DrawingWand, font_family string) bool
	DrawSetFontResolution                        func(d *DrawingWand, x_resolution, y_resolution float64) bool
	DrawSetFontSize                              func(d *DrawingWand, pointsize float64)
	DrawSetFontStretch                           func(d *DrawingWand, font_stretch StretchType)
	DrawSetFontStyle                             func(d *DrawingWand, style StyleType)
	DrawSetFontWeight                            func(d *DrawingWand, font_weight size_t)
	DrawSetGravity                               func(d *DrawingWand, gravity GravityType)
	DrawSetOpacity                               func(d *DrawingWand, opacity float64)
	DrawSetStrokeAlpha                           func(d *DrawingWand, stroke_alpha float64)
	DrawSetStrokeAntialias                       func(d *DrawingWand, stroke_antialias bool)
	DrawSetStrokeColor                           func(d *DrawingWand, stroke_wand *PixelWand)
	DrawSetStrokeDashArray                       func(d *DrawingWand, number_elements size_t, dash_array *float64) bool
	DrawSetStrokeDashOffset                      func(d *DrawingWand, dash_offset float64)
	DrawSetStrokeLineCap                         func(d *DrawingWand, linecap LineCap)
	DrawSetStrokeLineJoin                        func(d *DrawingWand, linejoin LineJoin)
	DrawSetStrokeMiterLimit                      func(d *DrawingWand, miterlimit size_t)
	DrawSetStrokeOpacity                         func(d *DrawingWand, stroke_opacity float64)
	DrawSetStrokePatternURL                      func(d *DrawingWand, stroke_url string) bool
	DrawSetStrokeWidth                           func(d *DrawingWand, stroke_width float64)
	DrawSetTextAlignment                         func(d *DrawingWand, alignment AlignType)
	DrawSetTextAntialias                         func(d *DrawingWand, text_antialias bool)
	DrawSetTextDecoration                        func(d *DrawingWand, decoration DecorationType)
	DrawSetTextEncoding                          func(d *DrawingWand, encoding string)
	DrawSetTextInterlineSpacing                  func(d *DrawingWand, interline_spacing float64)
	DrawSetTextInterwordSpacing                  func(d *DrawingWand, interword_spacing float64)
	DrawSetTextKerning                           func(d *DrawingWand, kerning float64)
	DrawSetTextUnderColor                        func(d *DrawingWand, under_wand *PixelWand)
	DrawSetVectorGraphics                        func(d *DrawingWand, xml string) bool
	DrawSetViewbox                               func(d *DrawingWand, x1, y1, x2, y2 ssize_t)
	DrawSkewX                                    func(d *DrawingWand, degrees float64)
	DrawSkewY                                    func(d *DrawingWand, degrees float64)
	DrawTranslate                                func(d *DrawingWand, x, y float64)
)

func (d *DrawingWand) Clear()                { ClearDrawingWand(d) }
func (d *DrawingWand) Clone() *DrawingWand   { return CloneDrawingWand(d) }
func (d *DrawingWand) Destroy() *DrawingWand { return DestroyDrawingWand(d) }
func (d *DrawingWand) IsDrawingWand() bool   { return IsDrawingWand(d) }
func (d *DrawingWand) Peek() *DrawInfo       { return PeekDrawingWand(d) }
func (d *DrawingWand) Pop() bool             { return PopDrawingWand(d) }
func (d *DrawingWand) Push() bool            { return PushDrawingWand(d) }

func (d *DrawingWand) Affine(affine *AffineMatrix)          { DrawAffine(d, affine) }
func (d *DrawingWand) Annotation(x, y float64, text *uint8) { DrawAnnotation(d, x, y, text) }
func (d *DrawingWand) Arc(sx, sy, ex, ey, sd, ed float64)   { DrawArc(d, sx, sy, ex, ey, sd, ed) }
func (d *DrawingWand) Bezier(number_coordinates size_t, coordinates *PointInfo) {
	DrawBezier(d, number_coordinates, coordinates)
}
func (d *DrawingWand) Circle(ox, oy, px, py float64) { DrawCircle(d, ox, oy, px, py) }
func (d *DrawingWand) ClearException() bool          { return DrawClearException(d) }
func (d *DrawingWand) Color(x, y float64, paint_method PaintMethod) {
	DrawColor(d, x, y, paint_method)
}
func (d *DrawingWand) Comment(comment string) { DrawComment(d, comment) }
func (d *DrawingWand) Composite(compose CompositeOperator, x, y, width, height float64, magick_wand *MagickWand) bool {
	return DrawComposite(d, compose, x, y, width, height, magick_wand)
}
func (d *DrawingWand) Ellipse(ox, oy, rx, ry, start, end float64) {
	DrawEllipse(d, ox, oy, rx, ry, start, end)
}
func (d *DrawingWand) BorderColor(border_color *PixelWand) { DrawGetBorderColor(d, border_color) }
func (d *DrawingWand) ClipPath() string                    { return DrawGetClipPath(d) }
func (d *DrawingWand) ClipUnits() ClipPathUnits            { return DrawGetClipUnits(d) }
func (d *DrawingWand) Exception(severity *ExceptionType) string {
	return DrawGetException(d, severity)
}
func (d *DrawingWand) ExceptionType() ExceptionType    { return DrawGetExceptionType(d) }
func (d *DrawingWand) FillAlpha() float64              { return DrawGetFillAlpha(d) }
func (d *DrawingWand) FillColor(fill_color *PixelWand) { DrawGetFillColor(d, fill_color) }
func (d *DrawingWand) FillOpacity() float64            { return DrawGetFillOpacity(d) }
func (d *DrawingWand) FillRule() FillRule              { return DrawGetFillRule(d) }
func (d *DrawingWand) Font() string                    { return DrawGetFont(d) }
func (d *DrawingWand) FontFamily() string              { return DrawGetFontFamily(d) }
func (d *DrawingWand) FontResolution(x, y *float64) DrawBooleanType {
	return DrawGetFontResolution(d, x, y)
}
func (d *DrawingWand) FontSize() float64                   { return DrawGetFontSize(d) }
func (d *DrawingWand) FontStretch() StretchType            { return DrawGetFontStretch(d) }
func (d *DrawingWand) FontStyle() StyleType                { return DrawGetFontStyle(d) }
func (d *DrawingWand) FontWeight() size_t                  { return DrawGetFontWeight(d) }
func (d *DrawingWand) Gravity() GravityType                { return DrawGetGravity(d) }
func (d *DrawingWand) Opacity() float64                    { return DrawGetOpacity(d) }
func (d *DrawingWand) StrokeAlpha() float64                { return DrawGetStrokeAlpha(d) }
func (d *DrawingWand) StrokeAntialias() bool               { return DrawGetStrokeAntialias(d) }
func (d *DrawingWand) StrokeColor(stroke_color *PixelWand) { DrawGetStrokeColor(d, stroke_color) }
func (d *DrawingWand) StrokeDashArray(number_elements *size_t) *float64 {
	return DrawGetStrokeDashArray(d, number_elements)
}
func (d *DrawingWand) StrokeDashOffset() float64             { return DrawGetStrokeDashOffset(d) }
func (d *DrawingWand) StrokeLineCap() LineCap                { return DrawGetStrokeLineCap(d) }
func (d *DrawingWand) StrokeLineJoin() LineJoin              { return DrawGetStrokeLineJoin(d) }
func (d *DrawingWand) StrokeMiterLimit() size_t              { return DrawGetStrokeMiterLimit(d) }
func (d *DrawingWand) StrokeOpacity() float64                { return DrawGetStrokeOpacity(d) }
func (d *DrawingWand) StrokeWidth() float64                  { return DrawGetStrokeWidth(d) }
func (d *DrawingWand) TextAlignment() AlignType              { return DrawGetTextAlignment(d) }
func (d *DrawingWand) TextAntialias() bool                   { return DrawGetTextAntialias(d) }
func (d *DrawingWand) TextDecoration() DecorationType        { return DrawGetTextDecoration(d) }
func (d *DrawingWand) TextEncoding() string                  { return DrawGetTextEncoding(d) }
func (d *DrawingWand) TextInterwordSpacing() float64         { return DrawGetTextInterwordSpacing(d) }
func (d *DrawingWand) TextKerning() float64                  { return DrawGetTextKerning(d) }
func (d *DrawingWand) TextUnderColor(under_color *PixelWand) { DrawGetTextUnderColor(d, under_color) }
func (d *DrawingWand) VectorGraphics() string                { return DrawGetVectorGraphics(d) }
func (d *DrawingWand) Line(sx, sy, ex, ey float64)           { DrawLine(d, sx, sy, ex, ey) }
func (d *DrawingWand) Matte(x, y float64, paint_method PaintMethod) {
	DrawMatte(d, x, y, paint_method)
}
func (d *DrawingWand) PathClose() { DrawPathClose(d) }
func (d *DrawingWand) PathCurveToAbsolute(x1, y1, x2, y2, x, y float64) {
	DrawPathCurveToAbsolute(d, x1, y1, x2, y2, x, y)
}
func (d *DrawingWand) PathCurveToQuadraticBezierAbsolute(x1, y1, x, y float64) {
	DrawPathCurveToQuadraticBezierAbsolute(d, x1, y1, x, y)
}
func (d *DrawingWand) PathCurveToQuadraticBezierRelative(x1, y1, x, y float64) {
	DrawPathCurveToQuadraticBezierRelative(d, x1, y1, x, y)
}
func (d *DrawingWand) PathCurveToQuadraticBezierSmoothAbsolute(x, y float64) {
	DrawPathCurveToQuadraticBezierSmoothAbsolute(d, x, y)
}
func (d *DrawingWand) PathCurveToQuadraticBezierSmoothRelative(x, y float64) {
	DrawPathCurveToQuadraticBezierSmoothRelative(d, x, y)
}
func (d *DrawingWand) PathCurveToRelative(x1, y1, x2, y2, x, y float64) {
	DrawPathCurveToRelative(d, x1, y1, x2, y2, x, y)
}
func (d *DrawingWand) PathCurveToSmoothAbsolute(x2, y2, x, y float64) {
	DrawPathCurveToSmoothAbsolute(d, x2, y2, x, y)
}
func (d *DrawingWand) PathCurveToSmoothRelative(x2, y2, x, y float64) {
	DrawPathCurveToSmoothRelative(d, x2, y2, x, y)
}
func (d *DrawingWand) PathEllipticArcAbsolute(rx, ry, x_axis_rotation float64, large_arc_flag, sweep_flag bool, x, y float64) {
	DrawPathEllipticArcAbsolute(d, rx, ry, x_axis_rotation, large_arc_flag, sweep_flag, x, y)
}
func (d *DrawingWand) PathEllipticArcRelative(rx, ry, x_axis_rotation float64, large_arc_flag, sweep_flag bool, x, y float64) {
	DrawPathEllipticArcRelative(d, rx, ry, x_axis_rotation, large_arc_flag, sweep_flag, x, y)
}
func (d *DrawingWand) PathFinish()                     { DrawPathFinish(d) }
func (d *DrawingWand) PathLineToAbsolute(x, y float64) { DrawPathLineToAbsolute(d, x, y) }
func (d *DrawingWand) PathLineToHorizontalAbsolute(x float64) {
	DrawPathLineToHorizontalAbsolute(d, x)
}
func (d *DrawingWand) PathLineToHorizontalRelative(x float64) {
	DrawPathLineToHorizontalRelative(d, x)
}
func (d *DrawingWand) PathLineToRelative(x, y float64)      { DrawPathLineToRelative(d, x, y) }
func (d *DrawingWand) PathLineToVerticalAbsolute(y float64) { DrawPathLineToVerticalAbsolute(d, y) }
func (d *DrawingWand) PathLineToVerticalRelative(y float64) { DrawPathLineToVerticalRelative(d, y) }
func (d *DrawingWand) PathMoveToAbsolute(x, y float64)      { DrawPathMoveToAbsolute(d, x, y) }
func (d *DrawingWand) PathMoveToRelative(x, y float64)      { DrawPathMoveToRelative(d, x, y) }
func (d *DrawingWand) PathStart()                           { DrawPathStart(d) }
func (d *DrawingWand) PeekGraphicWand() *DrawInfo           { return DrawPeekGraphicWand(d) }
func (d *DrawingWand) Point(x, y float64)                   { DrawPoint(d, x, y) }
func (d *DrawingWand) Polygon(number_coordinates size_t, coordinates *PointInfo) {
	DrawPolygon(d, number_coordinates, coordinates)
}
func (d *DrawingWand) Polyline(number_coordinates size_t, coordinates *PointInfo) {
	DrawPolyline(d, number_coordinates, coordinates)
}
func (d *DrawingWand) PopClipPath()                     { DrawPopClipPath(d) }
func (d *DrawingWand) PopDefs()                         { DrawPopDefs(d) }
func (d *DrawingWand) PopGraphicContext() bool          { return DrawPopGraphicContext(d) }
func (d *DrawingWand) PopPattern() bool                 { return DrawPopPattern(d) }
func (d *DrawingWand) PushClipPath(clip_mask_id string) { DrawPushClipPath(d, clip_mask_id) }
func (d *DrawingWand) PushDefs()                        { DrawPushDefs(d) }
func (d *DrawingWand) PushGraphicContext() bool         { return DrawPushGraphicContext(d) }
func (d *DrawingWand) PushPattern(pattern_id string, x, y, width, height float64) bool {
	return DrawPushPattern(d, pattern_id, x, y, width, height)
}
func (d *DrawingWand) Rectangle(x1, y1, x2, y2 float64) { DrawRectangle(d, x1, y1, x2, y2) }
func (d *DrawingWand) ResetVectorGraphics()             { DrawResetVectorGraphics(d) }
func (d *DrawingWand) Rotate(degrees float64)           { DrawRotate(d, degrees) }
func (d *DrawingWand) RoundRectangle(x1, y1, x2, y2, rx, ry float64) {
	DrawRoundRectangle(d, x1, y1, x2, y2, rx, ry)
}
func (d *DrawingWand) Scale(x, y float64)                    { DrawScale(d, x, y) }
func (d *DrawingWand) SetBorderColor(border_wand *PixelWand) { DrawSetBorderColor(d, border_wand) }
func (d *DrawingWand) SetClipPath(clip_mask string) bool     { return DrawSetClipPath(d, clip_mask) }
func (d *DrawingWand) SetClipRule(fill_rule FillRule)        { DrawSetClipRule(d, fill_rule) }
func (d *DrawingWand) SetClipUnits(clip_units ClipPathUnits) { DrawSetClipUnits(d, clip_units) }
func (d *DrawingWand) SetFillAlpha(fill_alpha float64)       { DrawSetFillAlpha(d, fill_alpha) }
func (d *DrawingWand) SetFillColor(fill_wand *PixelWand)     { DrawSetFillColor(d, fill_wand) }
func (d *DrawingWand) SetFillOpacity(fill_opacity float64)   { DrawSetFillOpacity(d, fill_opacity) }
func (d *DrawingWand) SetFillPatternURL(fill_url string) bool {
	return DrawSetFillPatternURL(d, fill_url)
}
func (d *DrawingWand) SetFillRule(fill_rule FillRule) { DrawSetFillRule(d, fill_rule) }
func (d *DrawingWand) SetFont(font_name string) bool  { return DrawSetFont(d, font_name) }
func (d *DrawingWand) SetFontFamily(font_family string) bool {
	return DrawSetFontFamily(d, font_family)
}
func (d *DrawingWand) SetFontResolution(x_resolution, y_resolution float64) bool {
	return DrawSetFontResolution(d, x_resolution, y_resolution)
}
func (d *DrawingWand) SetFontSize(pointsize float64)           { DrawSetFontSize(d, pointsize) }
func (d *DrawingWand) SetFontStretch(font_stretch StretchType) { DrawSetFontStretch(d, font_stretch) }
func (d *DrawingWand) SetFontStyle(style StyleType)            { DrawSetFontStyle(d, style) }
func (d *DrawingWand) SetFontWeight(font_weight size_t)        { DrawSetFontWeight(d, font_weight) }
func (d *DrawingWand) SetGravity(gravity GravityType)          { DrawSetGravity(d, gravity) }
func (d *DrawingWand) SetOpacity(opacity float64)              { DrawSetOpacity(d, opacity) }
func (d *DrawingWand) SetStrokeAlpha(stroke_alpha float64)     { DrawSetStrokeAlpha(d, stroke_alpha) }
func (d *DrawingWand) SetStrokeAntialias(stroke_antialias bool) {
	DrawSetStrokeAntialias(d, stroke_antialias)
}
func (d *DrawingWand) SetStrokeColor(stroke_wand *PixelWand) { DrawSetStrokeColor(d, stroke_wand) }
func (d *DrawingWand) SetStrokeDashArray(number_elements size_t, dash_array *float64) bool {
	return DrawSetStrokeDashArray(d, number_elements, dash_array)
}
func (d *DrawingWand) SetStrokeDashOffset(dash_offset float64) {
	DrawSetStrokeDashOffset(d, dash_offset)
}
func (d *DrawingWand) SetStrokeLineCap(linecap LineCap)      { DrawSetStrokeLineCap(d, linecap) }
func (d *DrawingWand) SetStrokeLineJoin(linejoin LineJoin)   { DrawSetStrokeLineJoin(d, linejoin) }
func (d *DrawingWand) SetStrokeMiterLimit(miterlimit size_t) { DrawSetStrokeMiterLimit(d, miterlimit) }
func (d *DrawingWand) SetStrokeOpacity(stroke_opacity float64) {
	DrawSetStrokeOpacity(d, stroke_opacity)
}
func (d *DrawingWand) SetStrokePatternURL(stroke_url string) bool {
	return DrawSetStrokePatternURL(d, stroke_url)
}
func (d *DrawingWand) SetStrokeWidth(stroke_width float64)  { DrawSetStrokeWidth(d, stroke_width) }
func (d *DrawingWand) SetTextAlignment(alignment AlignType) { DrawSetTextAlignment(d, alignment) }
func (d *DrawingWand) SetTextAntialias(text_antialias bool) { DrawSetTextAntialias(d, text_antialias) }
func (d *DrawingWand) SetTextDecoration(decoration DecorationType) {
	DrawSetTextDecoration(d, decoration)
}
func (d *DrawingWand) SetTextEncoding(encoding string) { DrawSetTextEncoding(d, encoding) }
func (d *DrawingWand) SetTextInterlineSpacing(interline_spacing float64) {
	DrawSetTextInterlineSpacing(d, interline_spacing)
}
func (d *DrawingWand) SetTextInterwordSpacing(interword_spacing float64) {
	DrawSetTextInterwordSpacing(d, interword_spacing)
}
func (d *DrawingWand) SetTextKerning(kerning float64)          { DrawSetTextKerning(d, kerning) }
func (d *DrawingWand) SetTextUnderColor(under_wand *PixelWand) { DrawSetTextUnderColor(d, under_wand) }
func (d *DrawingWand) SetVectorGraphics(xml string) bool       { return DrawSetVectorGraphics(d, xml) }
func (d *DrawingWand) SetViewbox(x1, y1, x2, y2 ssize_t)       { DrawSetViewbox(d, x1, y1, x2, y2) }
func (d *DrawingWand) SkewX(degrees float64)                   { DrawSkewX(d, degrees) }
func (d *DrawingWand) SkewY(degrees float64)                   { DrawSkewY(d, degrees) }
func (d *DrawingWand) Translate(x, y float64)                  { DrawTranslate(d, x, y) }

var (
	NewPixelWand  func() *PixelWand
	NewPixelWands func(number_wands size_t) **PixelWand

	ClonePixelWands   func(wands **PixelWand, number_wands size_t) **PixelWand
	DestroyPixelWands func(wand, number_wands size_t) **PixelWand

	ClearPixelWand        func(p *PixelWand)
	ClonePixelWand        func(p *PixelWand) *PixelWand
	DestroyPixelWand      func(p *PixelWand) *PixelWand
	GetPixelViewException func(p *PixelWand, severity *ExceptionType) string
	IsPixelWand           func(p *PixelWand) bool
	IsPixelWandSimilar    func(p *PixelWand, q *PixelWand, fuzz float64) bool

	PixelClearException             func(p *PixelWand) bool
	PixelGetAlpha                   func(p *PixelWand) float64
	PixelGetAlphaQuantum            func(p *PixelWand) Quantum
	PixelGetBlack                   func(p *PixelWand) float64
	PixelGetBlackQuantum            func(p *PixelWand) Quantum
	PixelGetBlue                    func(p *PixelWand) float64
	PixelGetBlueQuantum             func(p *PixelWand) Quantum
	PixelGetColorAsNormalizedString func(p *PixelWand) string
	PixelGetColorAsString           func(p *PixelWand) string
	PixelGetColorCount              func(p *PixelWand) size_t
	PixelGetCyan                    func(p *PixelWand) float64
	PixelGetCyanQuantum             func(p *PixelWand) Quantum
	PixelGetException               func(p *PixelWand, severity *ExceptionType) string
	PixelGetExceptionType           func(p *PixelWand) ExceptionType
	PixelGetFuzz                    func(p *PixelWand) float64
	PixelGetGreen                   func(p *PixelWand) float64
	PixelGetGreenQuantum            func(p *PixelWand) Quantum
	PixelGetHSL                     func(p *PixelWand, hue, saturation, lightness *float64)
	PixelGetIndex                   func(p *PixelWand) IndexPacket
	PixelGetMagenta                 func(p *PixelWand) float64
	PixelGetMagentaQuantum          func(p *PixelWand) Quantum
	PixelGetMagickColor             func(p *PixelWand, color *MagickPixelPacket)
	PixelGetOpacity                 func(p *PixelWand) float64
	PixelGetOpacityQuantum          func(p *PixelWand) Quantum
	PixelGetQuantumColor            func(p *PixelWand, color *PixelPacket)
	PixelGetRed                     func(p *PixelWand) float64
	PixelGetRedQuantum              func(p *PixelWand) Quantum
	PixelGetYellow                  func(p *PixelWand) float64
	PixelGetYellowQuantum           func(p *PixelWand) Quantum
	PixelSetAlpha                   func(p *PixelWand, alpha float64)
	PixelSetAlphaQuantum            func(p *PixelWand, opacity Quantum)
	PixelSetBlack                   func(p *PixelWand, black float64)
	PixelSetBlackQuantum            func(p *PixelWand, black Quantum)
	PixelSetBlue                    func(p *PixelWand, blue float64)
	PixelSetBlueQuantum             func(p *PixelWand, blue Quantum)
	PixelSetColor                   func(p *PixelWand, color string) bool
	PixelSetColorCount              func(p *PixelWand, count size_t)
	PixelSetColorFromWand           func(p *PixelWand, color *PixelWand)
	PixelSetCyan                    func(p *PixelWand, cyan float64)
	PixelSetCyanQuantum             func(p *PixelWand, cyan Quantum)
	PixelSetFuzz                    func(p *PixelWand, fuzz float64)
	PixelSetGreen                   func(p *PixelWand, green float64)
	PixelSetGreenQuantum            func(p *PixelWand, green Quantum)
	PixelSetHSL                     func(p *PixelWand, hue, saturation, lightness float64)
	PixelSetIndex                   func(p *PixelWand, index IndexPacket)
	PixelSetMagenta                 func(p *PixelWand, magenta float64)
	PixelSetMagentaQuantum          func(p *PixelWand, magenta Quantum)
	PixelSetMagickColor             func(p *PixelWand, color *MagickPixelPacket)
	PixelSetOpacity                 func(p *PixelWand, opacity float64)
	PixelSetOpacityQuantum          func(p *PixelWand, opacity Quantum)
	PixelSetQuantumColor            func(p *PixelWand, color *PixelPacket)
	PixelSetRed                     func(p *PixelWand, red float64)
	PixelSetRedQuantum              func(p *PixelWand, red Quantum)
	PixelSetYellow                  func(p *PixelWand, yellow float64)
	PixelSetYellowQuantum           func(p *PixelWand, yellow Quantum)
)

func (p  *PixelWand) Clear() {	return ClearPixelWand(p)}
func (p  *PixelWand) Clone() *PixelWand {	return ClonePixelWand(p)}
func (p  *PixelWand) Destroy() *PixelWand {	return DestroyPixelWand(p)}
func (p  *PixelWand) PixelViewException( severity *ExceptionType) string {	return GetPixelViewException(p, severity)}
func (p  *PixelWand) IsPixelWand() bool {	return IsPixelWand(p)}
func (p  *PixelWand) IsSimilar( q *PixelWand, fuzz float64) bool {	return IsPixelWandSimilar(p, q, fuzz)}

func (p *PixelWand) ClearException() bool                     { return PixelClearException(p) }
func (p *PixelWand) Alpha() float64                           { return PixelGetAlpha(p) }
func (p *PixelWand) AlphaQuantum() Quantum                    { return PixelGetAlphaQuantum(p) }
func (p *PixelWand) Black() float64                           { return PixelGetBlack(p) }
func (p *PixelWand) BlackQuantum() Quantum                    { return PixelGetBlackQuantum(p) }
func (p *PixelWand) Blue() float64                            { return PixelGetBlue(p) }
func (p *PixelWand) BlueQuantum() Quantum                     { return PixelGetBlueQuantum(p) }
func (p *PixelWand) ColorAsNormalizedString() string          { return PixelGetColorAsNormalizedString(p) }
func (p *PixelWand) ColorAsString() string                    { return PixelGetColorAsString(p) }
func (p *PixelWand) ColorCount() size_t                       { return PixelGetColorCount(p) }
func (p *PixelWand) Cyan() float64                            { return PixelGetCyan(p) }
func (p *PixelWand) CyanQuantum() Quantum                     { return PixelGetCyanQuantum(p) }
func (p *PixelWand) Exception(severity *ExceptionType) string { return PixelGetException(p, severity) }
func (p *PixelWand) ExceptionType() ExceptionType             { return PixelGetExceptionType(p) }
func (p *PixelWand) Fuzz() float64                            { return PixelGetFuzz(p) }
func (p *PixelWand) Green() float64                           { return PixelGetGreen(p) }
func (p *PixelWand) GreenQuantum() Quantum                    { return PixelGetGreenQuantum(p) }
func (p *PixelWand) HSL(hue, saturation, lightness *float64) {
	return PixelGetHSL(p, hue, saturation, lightness)
}
func (p *PixelWand) Index() IndexPacket                   { return PixelGetIndex(p) }
func (p *PixelWand) Magenta() float64                     { return PixelGetMagenta(p) }
func (p *PixelWand) MagentaQuantum() Quantum              { return PixelGetMagentaQuantum(p) }
func (p *PixelWand) MagickColor(color *MagickPixelPacket) { return PixelGetMagickColor(p, color) }
func (p *PixelWand) Opacity() float64                     { return PixelGetOpacity(p) }
func (p *PixelWand) OpacityQuantum() Quantum              { return PixelGetOpacityQuantum(p) }
func (p *PixelWand) QuantumColor(color *PixelPacket)      { return PixelGetQuantumColor(p, color) }
func (p *PixelWand) Red() float64                         { return PixelGetRed(p) }
func (p *PixelWand) RedQuantum() Quantum                  { return PixelGetRedQuantum(p) }
func (p *PixelWand) Yellow() float64                      { return PixelGetYellow(p) }
func (p *PixelWand) YellowQuantum() Quantum               { return PixelGetYellowQuantum(p) }
func (p *PixelWand) SetAlpha(alpha float64)               { return PixelSetAlpha(p, alpha) }
func (p *PixelWand) SetAlphaQuantum(opacity Quantum)      { return PixelSetAlphaQuantum(p, opacity) }
func (p *PixelWand) SetBlack(black float64)               { return PixelSetBlack(p, black) }
func (p *PixelWand) SetBlackQuantum(black Quantum)        { return PixelSetBlackQuantum(p, black) }
func (p *PixelWand) SetBlue(blue float64)                 { return PixelSetBlue(p, blue) }
func (p *PixelWand) SetBlueQuantum(blue Quantum)          { return PixelSetBlueQuantum(p, blue) }
func (p *PixelWand) SetColor(color string) bool           { return PixelSetColor(p, color) }
func (p *PixelWand) SetColorCount(count size_t)           { return PixelSetColorCount(p, count) }
func (p *PixelWand) SetColorFromWand(color *PixelWand)    { return PixelSetColorFromWand(p, color) }
func (p *PixelWand) SetCyan(cyan float64)                 { return PixelSetCyan(p, cyan) }
func (p *PixelWand) SetCyanQuantum(cyan Quantum)          { return PixelSetCyanQuantum(p, cyan) }
func (p *PixelWand) SetFuzz(fuzz float64)                 { return PixelSetFuzz(p, fuzz) }
func (p *PixelWand) SetGreen(green float64)               { return PixelSetGreen(p, green) }
func (p *PixelWand) SetGreenQuantum(green Quantum)        { return PixelSetGreenQuantum(p, green) }
func (p *PixelWand) SetHSL(hue, saturation, lightness float64) {
	return PixelSetHSL(p, hue, saturation, lightness)
}
func (p *PixelWand) SetIndex(index IndexPacket)              { return PixelSetIndex(p, index) }
func (p *PixelWand) SetMagenta(magenta float64)              { return PixelSetMagenta(p, magenta) }
func (p *PixelWand) SetMagentaQuantum(magenta Quantum)       { return PixelSetMagentaQuantum(p, magenta) }
func (p *PixelWand) SetMagickColor(color *MagickPixelPacket) { return PixelSetMagickColor(p, color) }
func (p *PixelWand) SetOpacity(opacity float64)              { return PixelSetOpacity(p, opacity) }
func (p *PixelWand) SetOpacityQuantum(opacity Quantum)       { return PixelSetOpacityQuantum(p, opacity) }
func (p *PixelWand) SetQuantumColor(color *PixelPacket)      { return PixelSetQuantumColor(p, color) }
func (p *PixelWand) SetRed(red float64)                      { return PixelSetRed(p, red) }
func (p *PixelWand) SetRedQuantum(red Quantum)               { return PixelSetRedQuantum(p, red) }
func (p *PixelWand) SetYellow(yellow float64)                { return PixelSetYellow(p, yellow) }
func (p *PixelWand) SetYellowQuantum(yellow Quantum)         { return PixelSetYellowQuantum(p, yellow) }
var (
	ClonePixelView                  func(pixel_view *PixelView) *PixelView
	DestroyPixelView                func(pixel_view *PixelView, number_wands, number_threads size_t) *PixelView

	DuplexTransferPixelViewIterator func(source, duplex, destination *PixelView, transfer DuplexTransferPixelViewMethod, context *Void) bool
	GetPixelViewHeight              func(pixel_view *PixelView) size_t
	GetPixelViewIterator            func(source *PixelView, get GetPixelViewMethod, context *Void) bool
	GetPixelViewPixels              func(pixel_view *PixelView) *PixelWand
	GetPixelViewWand                func(pixel_view *PixelView) *MagickWand
	GetPixelViewWidth               func(pixel_view *PixelView) size_t
	GetPixelViewX                   func(pixel_view *PixelView) ssize_t
	GetPixelViewY                   func(pixel_view *PixelView) ssize_t
	IsPixelView                     func(pixel_view *PixelView) bool
	SetPixelViewIterator            func(destination *PixelView, set SetPixelViewMethod, context *Void) bool
	TransferPixelViewIterator       func(source, destination *PixelView, transfer TransferPixelViewMethod, context *Void) bool
	UpdatePixelViewIterator         func(source *PixelView, update UpdatePixelViewMethod, context *Void) bool
)

var (
	ClearPixelIterator              func(p *PixelIterator)
	ClonePixelIterator              func(p *PixelIterator) *PixelIterator
	DestroyPixelIterator            func(p *PixelIterator) *PixelIterator
	IsPixelIterator                 func(p *PixelIterator) bool

	PixelClearIteratorException     func(p *PixelIterator) bool
	PixelGetCurrentIteratorRow      func(p *PixelIterator, number_wands *size_t) **PixelWand
	PixelGetIteratorException       func(p *PixelIterator, severity *ExceptionType) string
	PixelGetIteratorExceptionType   func(p *PixelIterator) ExceptionType
	PixelGetIteratorRow             func(p *PixelIterator) bool
	PixelGetNextIteratorRow         func(p *PixelIterator, number_wands *size_t) **PixelWand
	PixelGetNextRow                 func(p *PixelIterator, number_wands *size_t) **PixelWand
	PixelGetPreviousIteratorRow     func(p *PixelIterator, number_wands *size_t) **PixelWand
	PixelIteratorGetException       func(p *Pixeliterator, severity *ExceptionType) string
	PixelResetIterator              func(p *PixelIterator)
	PixelSetFirstIteratorRow        func(p *PixelIterator)
	PixelSetIteratorRow             func(p *PixelIterator, row ssize_t) bool
	PixelSetLastIteratorRow         func(p *PixelIterator)
	PixelSyncIterator               func(p *PixelIterator) bool
)
func (p  *PixelIterator) Clear() {	return ClearPixelIterator(p)}
func (p  *PixelIterator) Clone() *PixelIterator {	return ClonePixelIterator(p)}
func (p  *PixelIterator) Destroy() *PixelIterator {	return DestroyPixelIterator(p)}
func (p  *PixelIterator) IsPixelIterator() bool {	return IsPixelIterator(p)}

//TODO(t): w w/o Iterator in name
func (p  *PixelIterator) CurrentIteratorRow( number_wands *size_t) **PixelWand {	return PixelGetCurrentIteratorRow(p, number_wands *size_t)}
func (p  *PixelIterator) IteratorException( severity *ExceptionType) string {	return PixelGetIteratorException(p, severity *ExceptionType)}
func (p  *PixelIterator) IteratorExceptionType() ExceptionType {	return PixelGetIteratorExceptionType(p)}
func (p  *PixelIterator) IteratorRow() bool {	return PixelGetIteratorRow(p)}
func (p  *PixelIterator) NextIteratorRow( number_wands *size_t) **PixelWand {	return PixelGetNextIteratorRow(p, number_wands *size_t)}
func (p  *PixelIterator) NextRow( number_wands *size_t) **PixelWand {	return PixelGetNextRow(p, number_wands *size_t)}
func (p  *PixelIterator) PreviousIteratorRow( number_wands *size_t) **PixelWand {	return PixelGetPreviousIteratorRow(p, number_wands *size_t)}
func (p  *PixelIterator) Exception( severity *ExceptionType) string {	return PixelIteratorGetException(p, severity *ExceptionType)}

func (p  *PixelIterator) ClearException() bool {	return PixelClearIteratorException(p)}
func (p  *PixelIterator) Reset() {	return PixelResetIterator(p)}
func (p  *PixelIterator) SetFirstIteratorRow() {	return PixelSetFirstIteratorRow(p)}
func (p  *PixelIterator) SetIteratorRow( row ssize_t) bool {	return PixelSetIteratorRow(p, row ssize_t)}
func (p  *PixelIterator) SetLastIteratorRow() {	return PixelSetLastIteratorRow(p)}
func (p  *PixelIterator) Sync() bool {	return PixelSyncIterator(p)}

var(
	CloneWandView                   func(wand_view *WandView) *WandView
	DestroyWandView                 func(wand_view *WandView) *WandView
	DuplexTransferImageViewMethod   func(source, duplex, destination *WandView, y ssize_t, thread_id int, context *Void) bool
	DuplexTransferWandViewIterator  func(source, duplex, destination *WandView, transfer DuplexTransferWandViewMethod, context *Void) bool
	GetImageViewMethod              func(source *WandView, y ssize_t, thread_id int, context *Void) bool
	GetWandViewException            func(wand_view *WandView, severity *ExceptionType) string
	GetWandViewExtent               func(wand_view *WandView) RectangleInfo
	GetWandViewIterator             func(source *WandView, get GetWandViewMethod, context *Void) bool
	GetWandViewPixels               func(wand_view *WandView) *PixelWand
	GetWandViewWand                 func(wand_view *WandView) *MagickWand
	IsWandView                      func(wand_view *WandView) bool
	SetWandViewDescription          func(image_view *WandView, description string)
	SetWandViewIterator             func(destination *WandView, set SetWandViewMethod, context *Void) bool
	SetWandViewThreads              func(image_view *WandView, number_threads size_t)
	TransferImageViewMethod         func(source, destination *WandView, y ssize_t, thread_id int, context *Void) bool
	TransferWandViewIterator        func(source, destination *WandView, transfer TransferWandViewMethod, context *Void) bool
	UpdateImageViewMethod           func(source *WandView, y ssize_t, thread_id int, context *Void) bool
	UpdateWandViewIterator          func(source *WandView, update UpdateWandViewMethod, context *Void) bool
)

var (
	MagickCommandGenesis            func(image_info *ImageInfo, command MagickCommand, argc int, argv, metadata []string, exception *ExceptionInfo) bool
	MagickDestroyImage              func(image *Image) *Image
	MagickGetCopyright              func() string
	MagickGetHomeURL                func() string
	MagickGetPackageName            func() string
	MagickGetQuantumDepth           func(depth *size_t) string
	MagickGetQuantumRange           func(range_ *size_t) string
	MagickGetReleaseDate            func() string
	MagickGetResource               func(type_ ResourceType) MagickSizeType
	MagickGetResourceLimit          func(type_ ResourceType) MagickSizeType
	MagickGetVersion                func(version *size_t) string
	MagickProgressMonitor           func(text string, offset MagickOffsetType, span MagickSizeType, client_data *Void) bool
	MagickProgressMonitorFIX        func(text string, offset MagickOffsetType, span MagickSizeType, client_data *Void) bool
	MagickQueryConfigureOption      func(option string) string
	MagickQueryConfigureOptions     func(pattern string, numberOptions *size_t) []string
	MagickQueryFonts                func(pattern string, numberFonts *size_t) []string
	MagickQueryFormats              func(pattern string, numberFormats *size_t) []string
	MagickRelinquishMemory          func(resource *Void) *Void
	MagickSetResourceLimit          func(type_ ResourceType, limit MagickSizeType) bool
	MagickWandGenesis               func()
	MagickWandTerminus              func()
	NewMagickWandFromImage          func(image *Image) *MagickWand
	SetImageViewMethod              func(destination *ImageView, y ssize_t, thread_id int, context *Void) bool
)

func init() {
	AddDllApis("CORE_RL_wand_.dll", false, Apis{
		// {"AcquireWandId", &AcquireWandId},
		// {"AnimateImageCommand", &AnimateImageCommand},
		{"ClearDrawingWand", &ClearDrawingWand},
		{"ClearMagickWand", &ClearMagickWand},
		{"ClearPixelIterator", &ClearPixelIterator},
		{"ClearPixelWand", &ClearPixelWand},
		{"CloneDrawingWand", &CloneDrawingWand},
		{"CloneMagickWand", &CloneMagickWand},
		{"ClonePixelIterator", &ClonePixelIterator},
		{"ClonePixelView", &ClonePixelView},
		{"ClonePixelWand", &ClonePixelWand},
		{"ClonePixelWands", &ClonePixelWands},
		{"CloneWandView", &CloneWandView},
		// {"CompareImageCommand", &CompareImageCommand},
		// {"CompositeImageCommand", &CompositeImageCommand},
		// {"ConjureImageCommand", &ConjureImageCommand},
		// {"ConvertImageCommand", &ConvertImageCommand},
		{"DestroyDrawingWand", &DestroyDrawingWand},
		{"DestroyMagickWand", &DestroyMagickWand},
		{"DestroyPixelIterator", &DestroyPixelIterator},
		{"DestroyPixelView", &DestroyPixelView},
		{"DestroyPixelWand", &DestroyPixelWand},
		{"DestroyPixelWands", &DestroyPixelWands},
		// {"DestroyWandIds", &DestroyWandIds},
		{"DestroyWandView", &DestroyWandView},
		// {"DisplayImageCommand", &DisplayImageCommand},
		{"DrawAffine", &DrawAffine},
		// {"DrawAllocateWand", &DrawAllocateWand},
		{"DrawAnnotation", &DrawAnnotation},
		{"DrawArc", &DrawArc},
		{"DrawBezier", &DrawBezier},
		{"DrawCircle", &DrawCircle},
		{"DrawClearException", &DrawClearException},
		{"DrawColor", &DrawColor},
		{"DrawComment", &DrawComment},
		{"DrawComposite", &DrawComposite},
		{"DrawEllipse", &DrawEllipse},
		{"DrawGetBorderColor", &DrawGetBorderColor},
		{"DrawGetClipPath", &DrawGetClipPath},
		// {"DrawGetClipRule", &DrawGetClipRule},
		{"DrawGetClipUnits", &DrawGetClipUnits},
		{"DrawGetException", &DrawGetException},
		{"DrawGetExceptionType", &DrawGetExceptionType},
		{"DrawGetFillAlpha", &DrawGetFillAlpha},
		{"DrawGetFillColor", &DrawGetFillColor},
		{"DrawGetFillOpacity", &DrawGetFillOpacity},
		{"DrawGetFillRule", &DrawGetFillRule},
		{"DrawGetFont", &DrawGetFont},
		{"DrawGetFontFamily", &DrawGetFontFamily},
		{"DrawGetFontResolution", &DrawGetFontResolution},
		{"DrawGetFontSize", &DrawGetFontSize},
		{"DrawGetFontStretch", &DrawGetFontStretch},
		{"DrawGetFontStyle", &DrawGetFontStyle},
		{"DrawGetFontWeight", &DrawGetFontWeight},
		{"DrawGetGravity", &DrawGetGravity},
		{"DrawGetOpacity", &DrawGetOpacity},
		{"DrawGetStrokeAlpha", &DrawGetStrokeAlpha},
		{"DrawGetStrokeAntialias", &DrawGetStrokeAntialias},
		{"DrawGetStrokeColor", &DrawGetStrokeColor},
		{"DrawGetStrokeDashArray", &DrawGetStrokeDashArray},
		{"DrawGetStrokeDashOffset", &DrawGetStrokeDashOffset},
		{"DrawGetStrokeLineCap", &DrawGetStrokeLineCap},
		{"DrawGetStrokeLineJoin", &DrawGetStrokeLineJoin},
		{"DrawGetStrokeMiterLimit", &DrawGetStrokeMiterLimit},
		{"DrawGetStrokeOpacity", &DrawGetStrokeOpacity},
		{"DrawGetStrokeWidth", &DrawGetStrokeWidth},
		{"DrawGetTextAlignment", &DrawGetTextAlignment},
		{"DrawGetTextAntialias", &DrawGetTextAntialias},
		{"DrawGetTextDecoration", &DrawGetTextDecoration},
		{"DrawGetTextEncoding", &DrawGetTextEncoding},
		// {"DrawGetTextInterlineSpacing", &DrawGetTextInterlineSpacing},
		{"DrawGetTextInterwordSpacing", &DrawGetTextInterwordSpacing},
		{"DrawGetTextKerning", &DrawGetTextKerning},
		{"DrawGetTextUnderColor", &DrawGetTextUnderColor},
		{"DrawGetVectorGraphics", &DrawGetVectorGraphics},
		{"DrawLine", &DrawLine},
		{"DrawMatte", &DrawMatte},
		{"DrawPathClose", &DrawPathClose},
		{"DrawPathCurveToAbsolute", &DrawPathCurveToAbsolute},
		{"DrawPathCurveToQuadraticBezierAbsolute", &DrawPathCurveToQuadraticBezierAbsolute},
		{"DrawPathCurveToQuadraticBezierRelative", &DrawPathCurveToQuadraticBezierRelative},
		{"DrawPathCurveToQuadraticBezierSmoothAbsolute", &DrawPathCurveToQuadraticBezierSmoothAbsolute},
		{"DrawPathCurveToQuadraticBezierSmoothRelative", &DrawPathCurveToQuadraticBezierSmoothRelative},
		{"DrawPathCurveToRelative", &DrawPathCurveToRelative},
		{"DrawPathCurveToSmoothAbsolute", &DrawPathCurveToSmoothAbsolute},
		{"DrawPathCurveToSmoothRelative", &DrawPathCurveToSmoothRelative},
		{"DrawPathEllipticArcAbsolute", &DrawPathEllipticArcAbsolute},
		{"DrawPathEllipticArcRelative", &DrawPathEllipticArcRelative},
		{"DrawPathFinish", &DrawPathFinish},
		{"DrawPathLineToAbsolute", &DrawPathLineToAbsolute},
		{"DrawPathLineToHorizontalAbsolute", &DrawPathLineToHorizontalAbsolute},
		{"DrawPathLineToHorizontalRelative", &DrawPathLineToHorizontalRelative},
		{"DrawPathLineToRelative", &DrawPathLineToRelative},
		{"DrawPathLineToVerticalAbsolute", &DrawPathLineToVerticalAbsolute},
		{"DrawPathLineToVerticalRelative", &DrawPathLineToVerticalRelative},
		{"DrawPathMoveToAbsolute", &DrawPathMoveToAbsolute},
		{"DrawPathMoveToRelative", &DrawPathMoveToRelative},
		{"DrawPathStart", &DrawPathStart},
		{"DrawPeekGraphicWand", &DrawPeekGraphicWand},
		{"DrawPoint", &DrawPoint},
		{"DrawPolygon", &DrawPolygon},
		{"DrawPolyline", &DrawPolyline},
		{"DrawPopClipPath", &DrawPopClipPath},
		{"DrawPopDefs", &DrawPopDefs},
		{"DrawPopGraphicContext", &DrawPopGraphicContext},
		{"DrawPopPattern", &DrawPopPattern},
		{"DrawPushClipPath", &DrawPushClipPath},
		{"DrawPushDefs", &DrawPushDefs},
		{"DrawPushGraphicContext", &DrawPushGraphicContext},
		{"DrawPushPattern", &DrawPushPattern},
		{"DrawRectangle", &DrawRectangle},
		// {"DrawRender", &DrawRender},
		{"DrawResetVectorGraphics", &DrawResetVectorGraphics},
		{"DrawRotate", &DrawRotate},
		{"DrawRoundRectangle", &DrawRoundRectangle},
		{"DrawScale", &DrawScale},
		{"DrawSetBorderColor", &DrawSetBorderColor},
		{"DrawSetClipPath", &DrawSetClipPath},
		{"DrawSetClipRule", &DrawSetClipRule},
		{"DrawSetClipUnits", &DrawSetClipUnits},
		{"DrawSetFillAlpha", &DrawSetFillAlpha},
		{"DrawSetFillColor", &DrawSetFillColor},
		{"DrawSetFillOpacity", &DrawSetFillOpacity},
		{"DrawSetFillPatternURL", &DrawSetFillPatternURL},
		{"DrawSetFillRule", &DrawSetFillRule},
		{"DrawSetFont", &DrawSetFont},
		{"DrawSetFontFamily", &DrawSetFontFamily},
		{"DrawSetFontResolution", &DrawSetFontResolution},
		{"DrawSetFontSize", &DrawSetFontSize},
		{"DrawSetFontStretch", &DrawSetFontStretch},
		{"DrawSetFontStyle", &DrawSetFontStyle},
		{"DrawSetFontWeight", &DrawSetFontWeight},
		{"DrawSetGravity", &DrawSetGravity},
		{"DrawSetOpacity", &DrawSetOpacity},
		{"DrawSetStrokeAlpha", &DrawSetStrokeAlpha},
		{"DrawSetStrokeAntialias", &DrawSetStrokeAntialias},
		{"DrawSetStrokeColor", &DrawSetStrokeColor},
		{"DrawSetStrokeDashArray", &DrawSetStrokeDashArray},
		{"DrawSetStrokeDashOffset", &DrawSetStrokeDashOffset},
		{"DrawSetStrokeLineCap", &DrawSetStrokeLineCap},
		{"DrawSetStrokeLineJoin", &DrawSetStrokeLineJoin},
		{"DrawSetStrokeMiterLimit", &DrawSetStrokeMiterLimit},
		{"DrawSetStrokeOpacity", &DrawSetStrokeOpacity},
		{"DrawSetStrokePatternURL", &DrawSetStrokePatternURL},
		{"DrawSetStrokeWidth", &DrawSetStrokeWidth},
		{"DrawSetTextAlignment", &DrawSetTextAlignment},
		{"DrawSetTextAntialias", &DrawSetTextAntialias},
		{"DrawSetTextDecoration", &DrawSetTextDecoration},
		{"DrawSetTextEncoding", &DrawSetTextEncoding},
		{"DrawSetTextInterlineSpacing", &DrawSetTextInterlineSpacing},
		{"DrawSetTextInterwordSpacing", &DrawSetTextInterwordSpacing},
		{"DrawSetTextKerning", &DrawSetTextKerning},
		{"DrawSetTextUnderColor", &DrawSetTextUnderColor},
		{"DrawSetVectorGraphics", &DrawSetVectorGraphics},
		{"DrawSetViewbox", &DrawSetViewbox},
		{"DrawSkewX", &DrawSkewX},
		{"DrawSkewY", &DrawSkewY},
		{"DrawTranslate", &DrawTranslate},
		{"DuplexTransferPixelViewIterator", &DuplexTransferPixelViewIterator},
		{"DuplexTransferWandViewIterator", &DuplexTransferWandViewIterator},
		{"GetImageFromMagickWand", &GetImageFromMagickWand},
		{"GetPixelViewException", &GetPixelViewException},
		{"GetPixelViewHeight", &GetPixelViewHeight},
		{"GetPixelViewIterator", &GetPixelViewIterator},
		{"GetPixelViewPixels", &GetPixelViewPixels},
		{"GetPixelViewWand", &GetPixelViewWand},
		{"GetPixelViewWidth", &GetPixelViewWidth},
		{"GetPixelViewX", &GetPixelViewX},
		{"GetPixelViewY", &GetPixelViewY},
		{"GetWandViewException", &GetWandViewException},
		{"GetWandViewExtent", &GetWandViewExtent},
		{"GetWandViewIterator", &GetWandViewIterator},
		{"GetWandViewPixels", &GetWandViewPixels},
		{"GetWandViewWand", &GetWandViewWand},
		// {"IdentifyImageCommand", &IdentifyImageCommand},
		// {"ImportImageCommand", &ImportImageCommand},
		{"IsDrawingWand", &IsDrawingWand},
		{"IsMagickWand", &IsMagickWand},
		{"IsPixelIterator", &IsPixelIterator},
		{"IsPixelView", &IsPixelView},
		{"IsPixelWand", &IsPixelWand},
		{"IsPixelWandSimilar", &IsPixelWandSimilar},
		{"IsWandView", &IsWandView},
		{"MagickAdaptiveBlurImage", &MagickAdaptiveBlurImage},
		{"MagickAdaptiveBlurImageChannel", &MagickAdaptiveBlurImageChannel},
		{"MagickAdaptiveResizeImage", &MagickAdaptiveResizeImage},
		{"MagickAdaptiveSharpenImage", &MagickAdaptiveSharpenImage},
		{"MagickAdaptiveSharpenImageChannel", &MagickAdaptiveSharpenImageChannel},
		{"MagickAdaptiveThresholdImage", &MagickAdaptiveThresholdImage},
		{"MagickAddImage", &MagickAddImage},
		{"MagickAddNoiseImage", &MagickAddNoiseImage},
		{"MagickAddNoiseImageChannel", &MagickAddNoiseImageChannel},
		{"MagickAffineTransformImage", &MagickAffineTransformImage},
		{"MagickAnimateImages", &MagickAnimateImages},
		{"MagickAnnotateImage", &MagickAnnotateImage},
		{"MagickAppendImages", &MagickAppendImages},
		{"MagickAutoGammaImage", &MagickAutoGammaImage},
		{"MagickAutoGammaImageChannel", &MagickAutoGammaImageChannel},
		{"MagickAutoLevelImage", &MagickAutoLevelImage},
		{"MagickAutoLevelImageChannel", &MagickAutoLevelImageChannel},
		{"MagickAverageImages", &MagickAverageImages},
		{"MagickBlackThresholdImage", &MagickBlackThresholdImage},
		{"MagickBlueShiftImage", &MagickBlueShiftImage},
		{"MagickBlurImage", &MagickBlurImage},
		{"MagickBlurImageChannel", &MagickBlurImageChannel},
		{"MagickBorderImage", &MagickBorderImage},
		{"MagickBrightnessContrastImage", &MagickBrightnessContrastImage},
		{"MagickBrightnessContrastImageChannel", &MagickBrightnessContrastImageChannel},
		{"MagickCharcoalImage", &MagickCharcoalImage},
		{"MagickChopImage", &MagickChopImage},
		{"MagickClampImage", &MagickClampImage},
		{"MagickClampImageChannel", &MagickClampImageChannel},
		{"MagickClearException", &MagickClearException},
		{"MagickClipImage", &MagickClipImage},
		{"MagickClipImagePath", &MagickClipImagePath},
		{"MagickClipPathImage", &MagickClipPathImage},
		{"MagickClutImage", &MagickClutImage},
		{"MagickClutImageChannel", &MagickClutImageChannel},
		{"MagickCoalesceImages", &MagickCoalesceImages},
		{"MagickColorDecisionListImage", &MagickColorDecisionListImage},
		{"MagickColorFloodfillImage", &MagickColorFloodfillImage},
		{"MagickColorMatrixImage", &MagickColorMatrixImage},
		{"MagickColorizeImage", &MagickColorizeImage},
		{"MagickCombineImages", &MagickCombineImages},
		{"MagickCommandGenesis", &MagickCommandGenesis},
		{"MagickCommentImage", &MagickCommentImage},
		{"MagickCompareImageChannels", &MagickCompareImageChannels},
		{"MagickCompareImageLayers", &MagickCompareImageLayers},
		{"MagickCompareImages", &MagickCompareImages},
		{"MagickCompositeImage", &MagickCompositeImage},
		{"MagickCompositeImageChannel", &MagickCompositeImageChannel},
		{"MagickCompositeLayers", &MagickCompositeLayers},
		{"MagickConstituteImage", &MagickConstituteImage},
		{"MagickContrastImage", &MagickContrastImage},
		{"MagickContrastStretchImage", &MagickContrastStretchImage},
		{"MagickContrastStretchImageChannel", &MagickContrastStretchImageChannel},
		{"MagickConvolveImage", &MagickConvolveImage},
		{"MagickConvolveImageChannel", &MagickConvolveImageChannel},
		{"MagickCropImage", &MagickCropImage},
		{"MagickCycleColormapImage", &MagickCycleColormapImage},
		{"MagickDecipherImage", &MagickDecipherImage},
		{"MagickDeconstructImages", &MagickDeconstructImages},
		{"MagickDeleteImageArtifact", &MagickDeleteImageArtifact},
		{"MagickDeleteImageProperty", &MagickDeleteImageProperty},
		{"MagickDeleteOption", &MagickDeleteOption},
		{"MagickDescribeImage", &MagickDescribeImage},
		{"MagickDeskewImage", &MagickDeskewImage},
		{"MagickDespeckleImage", &MagickDespeckleImage},
		{"MagickDestroyImage", &MagickDestroyImage},
		{"MagickDisplayImage", &MagickDisplayImage},
		{"MagickDisplayImages", &MagickDisplayImages},
		{"MagickDistortImage", &MagickDistortImage},
		{"MagickDrawImage", &MagickDrawImage},
		{"MagickEdgeImage", &MagickEdgeImage},
		{"MagickEmbossImage", &MagickEmbossImage},
		{"MagickEncipherImage", &MagickEncipherImage},
		{"MagickEnhanceImage", &MagickEnhanceImage},
		{"MagickEqualizeImage", &MagickEqualizeImage},
		{"MagickEqualizeImageChannel", &MagickEqualizeImageChannel},
		{"MagickEvaluateImage", &MagickEvaluateImage},
		{"MagickEvaluateImageChannel", &MagickEvaluateImageChannel},
		{"MagickEvaluateImages", &MagickEvaluateImages},
		{"MagickExportImagePixels", &MagickExportImagePixels},
		{"MagickExtentImage", &MagickExtentImage},
		{"MagickFilterImage", &MagickFilterImage},
		{"MagickFilterImageChannel", &MagickFilterImageChannel},
		{"MagickFlattenImages", &MagickFlattenImages},
		{"MagickFlipImage", &MagickFlipImage},
		{"MagickFloodfillPaintImage", &MagickFloodfillPaintImage},
		{"MagickFlopImage", &MagickFlopImage},
		{"MagickForwardFourierTransformImage", &MagickForwardFourierTransformImage},
		{"MagickFrameImage", &MagickFrameImage},
		{"MagickFunctionImage", &MagickFunctionImage},
		{"MagickFunctionImageChannel", &MagickFunctionImageChannel},
		{"MagickFxImage", &MagickFxImage},
		{"MagickFxImageChannel", &MagickFxImageChannel},
		{"MagickGammaImage", &MagickGammaImage},
		{"MagickGammaImageChannel", &MagickGammaImageChannel},
		{"MagickGaussianBlurImage", &MagickGaussianBlurImage},
		{"MagickGaussianBlurImageChannel", &MagickGaussianBlurImageChannel},
		{"MagickGetAntialias", &MagickGetAntialias},
		{"MagickGetBackgroundColor", &MagickGetBackgroundColor},
		{"MagickGetColorspace", &MagickGetColorspace},
		{"MagickGetCompression", &MagickGetCompression},
		{"MagickGetCompressionQuality", &MagickGetCompressionQuality},
		{"MagickGetCopyright", &MagickGetCopyright},
		{"MagickGetException", &MagickGetException},
		{"MagickGetExceptionType", &MagickGetExceptionType},
		{"MagickGetFilename", &MagickGetFilename},
		{"MagickGetFont", &MagickGetFont},
		{"MagickGetFormat", &MagickGetFormat},
		{"MagickGetGravity", &MagickGetGravity},
		{"MagickGetHomeURL", &MagickGetHomeURL},
		{"MagickGetImage", &MagickGetImage},
		{"MagickGetImageAlphaChannel", &MagickGetImageAlphaChannel},
		{"MagickGetImageArtifact", &MagickGetImageArtifact},
		{"MagickGetImageArtifacts", &MagickGetImageArtifacts},
		{"MagickGetImageAttribute", &MagickGetImageAttribute},
		{"MagickGetImageBackgroundColor", &MagickGetImageBackgroundColor},
		{"MagickGetImageBlob", &MagickGetImageBlob},
		{"MagickGetImageBluePrimary", &MagickGetImageBluePrimary},
		{"MagickGetImageBorderColor", &MagickGetImageBorderColor},
		{"MagickGetImageChannelDepth", &MagickGetImageChannelDepth},
		{"MagickGetImageChannelDistortion", &MagickGetImageChannelDistortion},
		// {"MagickGetImageChannelDistortions", &MagickGetImageChannelDistortions},
		// {"MagickGetImageChannelExtrema", &MagickGetImageChannelExtrema},
		{"MagickGetImageChannelFeatures", &MagickGetImageChannelFeatures},
		{"MagickGetImageChannelKurtosis", &MagickGetImageChannelKurtosis},
		{"MagickGetImageChannelMean", &MagickGetImageChannelMean},
		{"MagickGetImageChannelRange", &MagickGetImageChannelRange},
		{"MagickGetImageChannelStatistics", &MagickGetImageChannelStatistics},
		{"MagickGetImageClipMask", &MagickGetImageClipMask},
		{"MagickGetImageColormapColor", &MagickGetImageColormapColor},
		{"MagickGetImageColors", &MagickGetImageColors},
		{"MagickGetImageColorspace", &MagickGetImageColorspace},
		{"MagickGetImageCompose", &MagickGetImageCompose},
		{"MagickGetImageCompression", &MagickGetImageCompression},
		{"MagickGetImageCompressionQuality", &MagickGetImageCompressionQuality},
		{"MagickGetImageDelay", &MagickGetImageDelay},
		{"MagickGetImageDepth", &MagickGetImageDepth},
		{"MagickGetImageDispose", &MagickGetImageDispose},
		{"MagickGetImageDistortion", &MagickGetImageDistortion},
		// {"MagickGetImageEndian", &MagickGetImageEndian},
		// {"MagickGetImageExtrema", &MagickGetImageExtrema},
		{"MagickGetImageFilename", &MagickGetImageFilename},
		{"MagickGetImageFormat", &MagickGetImageFormat},
		{"MagickGetImageFuzz", &MagickGetImageFuzz},
		{"MagickGetImageGamma", &MagickGetImageGamma},
		{"MagickGetImageGravity", &MagickGetImageGravity},
		{"MagickGetImageGreenPrimary", &MagickGetImageGreenPrimary},
		{"MagickGetImageHeight", &MagickGetImageHeight},
		{"MagickGetImageHistogram", &MagickGetImageHistogram},
		// {"MagickGetImageIndex", &MagickGetImageIndex},
		{"MagickGetImageInterlaceScheme", &MagickGetImageInterlaceScheme},
		{"MagickGetImageInterpolateMethod", &MagickGetImageInterpolateMethod},
		{"MagickGetImageIterations", &MagickGetImageIterations},
		{"MagickGetImageLength", &MagickGetImageLength},
		{"MagickGetImageMatte", &MagickGetImageMatte},
		{"MagickGetImageMatteColor", &MagickGetImageMatteColor},
		{"MagickGetImageOrientation", &MagickGetImageOrientation},
		{"MagickGetImagePage", &MagickGetImagePage},
		{"MagickGetImagePixelColor", &MagickGetImagePixelColor},
		{"MagickGetImagePixels", &MagickGetImagePixels},
		{"MagickGetImageProfile", &MagickGetImageProfile},
		{"MagickGetImageProfiles", &MagickGetImageProfiles},
		{"MagickGetImageProperties", &MagickGetImageProperties},
		{"MagickGetImageProperty", &MagickGetImageProperty},
		// {"MagickGetImageRange", &MagickGetImageRange},
		{"MagickGetImageRedPrimary", &MagickGetImageRedPrimary},
		{"MagickGetImageRegion", &MagickGetImageRegion},
		{"MagickGetImageRenderingIntent", &MagickGetImageRenderingIntent},
		{"MagickGetImageResolution", &MagickGetImageResolution},
		{"MagickGetImageScene", &MagickGetImageScene},
		{"MagickGetImageSignature", &MagickGetImageSignature},
		{"MagickGetImageSize", &MagickGetImageSize},
		{"MagickGetImageTicksPerSecond", &MagickGetImageTicksPerSecond},
		{"MagickGetImageTotalInkDensity", &MagickGetImageTotalInkDensity},
		{"MagickGetImageType", &MagickGetImageType},
		{"MagickGetImageUnits", &MagickGetImageUnits},
		{"MagickGetImageVirtualPixelMethod", &MagickGetImageVirtualPixelMethod},
		{"MagickGetImageWhitePoint", &MagickGetImageWhitePoint},
		{"MagickGetImageWidth", &MagickGetImageWidth},
		{"MagickGetImagesBlob", &MagickGetImagesBlob},
		{"MagickGetInterlaceScheme", &MagickGetInterlaceScheme},
		{"MagickGetInterpolateMethod", &MagickGetInterpolateMethod},
		{"MagickGetIteratorIndex", &MagickGetIteratorIndex},
		{"MagickGetNumberImages", &MagickGetNumberImages},
		{"MagickGetOption", &MagickGetOption},
		{"MagickGetOptions", &MagickGetOptions},
		{"MagickGetOrientation", &MagickGetOrientation},
		{"MagickGetPackageName", &MagickGetPackageName},
		{"MagickGetPage", &MagickGetPage},
		{"MagickGetPointsize", &MagickGetPointsize},
		{"MagickGetQuantumDepth", &MagickGetQuantumDepth},
		{"MagickGetQuantumRange", &MagickGetQuantumRange},
		{"MagickGetReleaseDate", &MagickGetReleaseDate},
		{"MagickGetResolution", &MagickGetResolution},
		{"MagickGetResource", &MagickGetResource},
		{"MagickGetResourceLimit", &MagickGetResourceLimit},
		// {"MagickGetSamplingFactors", &MagickGetSamplingFactors},
		{"MagickGetSize", &MagickGetSize},
		{"MagickGetSizeOffset", &MagickGetSizeOffset},
		{"MagickGetType", &MagickGetType},
		{"MagickGetVersion", &MagickGetVersion},
		{"MagickHaldClutImage", &MagickHaldClutImage},
		{"MagickHaldClutImageChannel", &MagickHaldClutImageChannel},
		{"MagickHasNextImage", &MagickHasNextImage},
		{"MagickHasPreviousImage", &MagickHasPreviousImage},
		{"MagickIdentifyImage", &MagickIdentifyImage},
		{"MagickImplodeImage", &MagickImplodeImage},
		{"MagickImportImagePixels", &MagickImportImagePixels},
		{"MagickInverseFourierTransformImage", &MagickInverseFourierTransformImage},
		{"MagickLabelImage", &MagickLabelImage},
		{"MagickLevelImage", &MagickLevelImage},
		{"MagickLevelImageChannel", &MagickLevelImageChannel},
		{"MagickLinearStretchImage", &MagickLinearStretchImage},
		// {"MagickLiquidRescaleImage", &MagickLiquidRescaleImage},
		{"MagickMagnifyImage", &MagickMagnifyImage},
		{"MagickMapImage", &MagickMapImage},
		{"MagickMatteFloodfillImage", &MagickMatteFloodfillImage},
		{"MagickMaximumImages", &MagickMaximumImages},
		{"MagickMedianFilterImage", &MagickMedianFilterImage},
		{"MagickMergeImageLayers", &MagickMergeImageLayers},
		{"MagickMinifyImage", &MagickMinifyImage},
		{"MagickMinimumImages", &MagickMinimumImages},
		{"MagickModeImage", &MagickModeImage},
		{"MagickModulateImage", &MagickModulateImage},
		{"MagickMontageImage", &MagickMontageImage},
		{"MagickMorphImages", &MagickMorphImages},
		{"MagickMorphologyImage", &MagickMorphologyImage},
		{"MagickMorphologyImageChannel", &MagickMorphologyImageChannel},
		{"MagickMosaicImages", &MagickMosaicImages},
		{"MagickMotionBlurImage", &MagickMotionBlurImage},
		{"MagickMotionBlurImageChannel", &MagickMotionBlurImageChannel},
		{"MagickNegateImage", &MagickNegateImage},
		{"MagickNegateImageChannel", &MagickNegateImageChannel},
		{"MagickNewImage", &MagickNewImage},
		{"MagickNextImage", &MagickNextImage},
		{"MagickNormalizeImage", &MagickNormalizeImage},
		{"MagickNormalizeImageChannel", &MagickNormalizeImageChannel},
		{"MagickOilPaintImage", &MagickOilPaintImage},
		{"MagickOpaqueImage", &MagickOpaqueImage},
		{"MagickOpaquePaintImage", &MagickOpaquePaintImage},
		{"MagickOpaquePaintImageChannel", &MagickOpaquePaintImageChannel},
		{"MagickOptimizeImageLayers", &MagickOptimizeImageLayers},
		// {"MagickOptimizeImageTransparency", &MagickOptimizeImageTransparency},
		{"MagickOrderedPosterizeImage", &MagickOrderedPosterizeImage},
		{"MagickOrderedPosterizeImageChannel", &MagickOrderedPosterizeImageChannel},
		{"MagickPaintFloodfillImage", &MagickPaintFloodfillImage},
		{"MagickPaintOpaqueImage", &MagickPaintOpaqueImage},
		{"MagickPaintOpaqueImageChannel", &MagickPaintOpaqueImageChannel},
		{"MagickPaintTransparentImage", &MagickPaintTransparentImage},
		{"MagickPingImage", &MagickPingImage},
		{"MagickPingImageBlob", &MagickPingImageBlob},
		{"MagickPingImageFile", &MagickPingImageFile},
		{"MagickPolaroidImage", &MagickPolaroidImage},
		{"MagickPosterizeImage", &MagickPosterizeImage},
		{"MagickPreviewImages", &MagickPreviewImages},
		{"MagickPreviousImage", &MagickPreviousImage},
		{"MagickProfileImage", &MagickProfileImage},
		{"MagickQuantizeImage", &MagickQuantizeImage},
		{"MagickQuantizeImages", &MagickQuantizeImages},
		{"MagickQueryConfigureOption", &MagickQueryConfigureOption},
		{"MagickQueryConfigureOptions", &MagickQueryConfigureOptions},
		{"MagickQueryFontMetrics", &MagickQueryFontMetrics},
		{"MagickQueryFonts", &MagickQueryFonts},
		{"MagickQueryFormats", &MagickQueryFormats},
		{"MagickQueryMultilineFontMetrics", &MagickQueryMultilineFontMetrics},
		{"MagickRadialBlurImage", &MagickRadialBlurImage},
		{"MagickRadialBlurImageChannel", &MagickRadialBlurImageChannel},
		{"MagickRaiseImage", &MagickRaiseImage},
		{"MagickRandomThresholdImage", &MagickRandomThresholdImage},
		{"MagickRandomThresholdImageChannel", &MagickRandomThresholdImageChannel},
		{"MagickReadImage", &MagickReadImage},
		{"MagickReadImageBlob", &MagickReadImageBlob},
		{"MagickReadImageFile", &MagickReadImageFile},
		{"MagickRecolorImage", &MagickRecolorImage},
		{"MagickReduceNoiseImage", &MagickReduceNoiseImage},
		{"MagickRegionOfInterestImage", &MagickRegionOfInterestImage},
		{"MagickRelinquishMemory", &MagickRelinquishMemory},
		{"MagickRemapImage", &MagickRemapImage},
		{"MagickRemoveImage", &MagickRemoveImage},
		{"MagickRemoveImageProfile", &MagickRemoveImageProfile},
		{"MagickResampleImage", &MagickResampleImage},
		{"MagickResetImagePage", &MagickResetImagePage},
		{"MagickResetIterator", &MagickResetIterator},
		{"MagickResizeImage", &MagickResizeImage},
		{"MagickRollImage", &MagickRollImage},
		{"MagickRotateImage", &MagickRotateImage},
		{"MagickSampleImage", &MagickSampleImage},
		{"MagickScaleImage", &MagickScaleImage},
		{"MagickSegmentImage", &MagickSegmentImage},
		{"MagickSelectiveBlurImage", &MagickSelectiveBlurImage},
		{"MagickSelectiveBlurImageChannel", &MagickSelectiveBlurImageChannel},
		{"MagickSeparateImageChannel", &MagickSeparateImageChannel},
		{"MagickSepiaToneImage", &MagickSepiaToneImage},
		{"MagickSetAntialias", &MagickSetAntialias},
		{"MagickSetBackgroundColor", &MagickSetBackgroundColor},
		{"MagickSetColorspace", &MagickSetColorspace},
		{"MagickSetCompression", &MagickSetCompression},
		{"MagickSetCompressionQuality", &MagickSetCompressionQuality},
		{"MagickSetDepth", &MagickSetDepth},
		{"MagickSetExtract", &MagickSetExtract},
		{"MagickSetFilename", &MagickSetFilename},
		{"MagickSetFirstIterator", &MagickSetFirstIterator},
		{"MagickSetFont", &MagickSetFont},
		{"MagickSetFormat", &MagickSetFormat},
		{"MagickSetGravity", &MagickSetGravity},
		{"MagickSetImage", &MagickSetImage},
		{"MagickSetImageAlphaChannel", &MagickSetImageAlphaChannel},
		{"MagickSetImageArtifact", &MagickSetImageArtifact},
		{"MagickSetImageAttribute", &MagickSetImageAttribute},
		{"MagickSetImageBackgroundColor", &MagickSetImageBackgroundColor},
		{"MagickSetImageBias", &MagickSetImageBias},
		{"MagickSetImageBluePrimary", &MagickSetImageBluePrimary},
		{"MagickSetImageBorderColor", &MagickSetImageBorderColor},
		{"MagickSetImageChannelDepth", &MagickSetImageChannelDepth},
		{"MagickSetImageClipMask", &MagickSetImageClipMask},
		{"MagickSetImageColor", &MagickSetImageColor},
		{"MagickSetImageColormapColor", &MagickSetImageColormapColor},
		{"MagickSetImageColorspace", &MagickSetImageColorspace},
		{"MagickSetImageCompose", &MagickSetImageCompose},
		{"MagickSetImageCompression", &MagickSetImageCompression},
		{"MagickSetImageCompressionQuality", &MagickSetImageCompressionQuality},
		{"MagickSetImageDelay", &MagickSetImageDelay},
		{"MagickSetImageDepth", &MagickSetImageDepth},
		{"MagickSetImageDispose", &MagickSetImageDispose},
		{"MagickSetImageEndian", &MagickSetImageEndian},
		{"MagickSetImageExtent", &MagickSetImageExtent},
		{"MagickSetImageFilename", &MagickSetImageFilename},
		{"MagickSetImageFormat", &MagickSetImageFormat},
		{"MagickSetImageFuzz", &MagickSetImageFuzz},
		{"MagickSetImageGamma", &MagickSetImageGamma},
		{"MagickSetImageGravity", &MagickSetImageGravity},
		{"MagickSetImageGreenPrimary", &MagickSetImageGreenPrimary},
		{"MagickSetImageIndex", &MagickSetImageIndex},
		{"MagickSetImageInterlaceScheme", &MagickSetImageInterlaceScheme},
		{"MagickSetImageInterpolateMethod", &MagickSetImageInterpolateMethod},
		{"MagickSetImageIterations", &MagickSetImageIterations},
		// {"MagickSetImageMatte", &MagickSetImageMatte},
		{"MagickSetImageMatteColor", &MagickSetImageMatteColor},
		{"MagickSetImageOpacity", &MagickSetImageOpacity},
		// {"MagickSetImageOption", &MagickSetImageOption},
		{"MagickSetImageOrientation", &MagickSetImageOrientation},
		{"MagickSetImagePage", &MagickSetImagePage},
		{"MagickSetImagePixels", &MagickSetImagePixels},
		{"MagickSetImageProfile", &MagickSetImageProfile},
		{"MagickSetImageProgressMonitor", &MagickSetImageProgressMonitor},
		{"MagickSetImageProperty", &MagickSetImageProperty},
		{"MagickSetImageRedPrimary", &MagickSetImageRedPrimary},
		{"MagickSetImageRenderingIntent", &MagickSetImageRenderingIntent},
		{"MagickSetImageResolution", &MagickSetImageResolution},
		{"MagickSetImageScene", &MagickSetImageScene},
		{"MagickSetImageTicksPerSecond", &MagickSetImageTicksPerSecond},
		{"MagickSetImageType", &MagickSetImageType},
		{"MagickSetImageUnits", &MagickSetImageUnits},
		{"MagickSetImageVirtualPixelMethod", &MagickSetImageVirtualPixelMethod},
		{"MagickSetImageWhitePoint", &MagickSetImageWhitePoint},
		{"MagickSetInterlaceScheme", &MagickSetInterlaceScheme},
		{"MagickSetInterpolateMethod", &MagickSetInterpolateMethod},
		{"MagickSetIteratorIndex", &MagickSetIteratorIndex},
		{"MagickSetLastIterator", &MagickSetLastIterator},
		{"MagickSetOption", &MagickSetOption},
		{"MagickSetOrientation", &MagickSetOrientation},
		{"MagickSetPage", &MagickSetPage},
		{"MagickSetPassphrase", &MagickSetPassphrase},
		{"MagickSetPointsize", &MagickSetPointsize},
		{"MagickSetProgressMonitor", &MagickSetProgressMonitor},
		{"MagickSetResolution", &MagickSetResolution},
		{"MagickSetResourceLimit", &MagickSetResourceLimit},
		{"MagickSetSamplingFactors", &MagickSetSamplingFactors},
		{"MagickSetSize", &MagickSetSize},
		{"MagickSetSizeOffset", &MagickSetSizeOffset},
		{"MagickSetType", &MagickSetType},
		{"MagickShadeImage", &MagickShadeImage},
		{"MagickShadowImage", &MagickShadowImage},
		{"MagickSharpenImage", &MagickSharpenImage},
		{"MagickSharpenImageChannel", &MagickSharpenImageChannel},
		{"MagickShaveImage", &MagickShaveImage},
		{"MagickShearImage", &MagickShearImage},
		{"MagickSigmoidalContrastImage", &MagickSigmoidalContrastImage},
		{"MagickSigmoidalContrastImageChannel", &MagickSigmoidalContrastImageChannel},
		{"MagickSimilarityImage", &MagickSimilarityImage},
		{"MagickSketchImage", &MagickSketchImage},
		{"MagickSmushImages", &MagickSmushImages},
		{"MagickSolarizeImage", &MagickSolarizeImage},
		// {"MagickSolarizeImageChannel", &MagickSolarizeImageChannel},
		{"MagickSparseColorImage", &MagickSparseColorImage},
		{"MagickSpliceImage", &MagickSpliceImage},
		{"MagickSpreadImage", &MagickSpreadImage},
		{"MagickStatisticImage", &MagickStatisticImage},
		{"MagickSteganoImage", &MagickSteganoImage},
		{"MagickStereoImage", &MagickStereoImage},
		{"MagickStripImage", &MagickStripImage},
		{"MagickSwirlImage", &MagickSwirlImage},
		{"MagickTextureImage", &MagickTextureImage},
		{"MagickThresholdImage", &MagickThresholdImage},
		{"MagickThresholdImageChannel", &MagickThresholdImageChannel},
		{"MagickThumbnailImage", &MagickThumbnailImage},
		{"MagickTintImage", &MagickTintImage},
		{"MagickTransformImage", &MagickTransformImage},
		{"MagickTransformImageColorspace", &MagickTransformImageColorspace},
		{"MagickTransparentImage", &MagickTransparentImage},
		{"MagickTransparentPaintImage", &MagickTransparentPaintImage},
		{"MagickTransposeImage", &MagickTransposeImage},
		{"MagickTransverseImage", &MagickTransverseImage},
		{"MagickTrimImage", &MagickTrimImage},
		{"MagickUniqueImageColors", &MagickUniqueImageColors},
		{"MagickUnsharpMaskImage", &MagickUnsharpMaskImage},
		{"MagickUnsharpMaskImageChannel", &MagickUnsharpMaskImageChannel},
		{"MagickVignetteImage", &MagickVignetteImage},
		{"MagickWandGenesis", &MagickWandGenesis},
		{"MagickWandTerminus", &MagickWandTerminus},
		{"MagickWaveImage", &MagickWaveImage},
		{"MagickWhiteThresholdImage", &MagickWhiteThresholdImage},
		{"MagickWriteImage", &MagickWriteImage},
		{"MagickWriteImageBlob", &MagickWriteImageBlob},
		{"MagickWriteImageFile", &MagickWriteImageFile},
		{"MagickWriteImages", &MagickWriteImages},
		{"MagickWriteImagesFile", &MagickWriteImagesFile},
		// {"MogrifyImage", &MogrifyImage},
		// {"MogrifyImageCommand", &MogrifyImageCommand},
		// {"MogrifyImageInfo", &MogrifyImageInfo},
		// {"MogrifyImageList", &MogrifyImageList},
		// {"MogrifyImages", &MogrifyImages},
		// {"MontageImageCommand", &MontageImageCommand},
		{"NewDrawingWand", &NewDrawingWand},
		{"NewMagickWand", &NewMagickWand},
		{"NewMagickWandFromImage", &NewMagickWandFromImage},
		{"NewPixelIterator", &NewPixelIterator},
		{"NewPixelRegionIterator", &NewPixelRegionIterator},
		{"NewPixelView", &NewPixelView},
		{"NewPixelViewRegion", &NewPixelViewRegion},
		{"NewPixelWand", &NewPixelWand},
		{"NewPixelWands", &NewPixelWands},
		{"NewWandView", &NewWandView},
		{"NewWandViewExtent", &NewWandViewExtent},
		{"PeekDrawingWand", &PeekDrawingWand},
		{"PixelClearException", &PixelClearException},
		{"PixelClearIteratorException", &PixelClearIteratorException},
		{"PixelGetAlpha", &PixelGetAlpha},
		{"PixelGetAlphaQuantum", &PixelGetAlphaQuantum},
		{"PixelGetBlack", &PixelGetBlack},
		{"PixelGetBlackQuantum", &PixelGetBlackQuantum},
		{"PixelGetBlue", &PixelGetBlue},
		{"PixelGetBlueQuantum", &PixelGetBlueQuantum},
		{"PixelGetColorAsNormalizedString", &PixelGetColorAsNormalizedString},
		{"PixelGetColorAsString", &PixelGetColorAsString},
		{"PixelGetColorCount", &PixelGetColorCount},
		{"PixelGetCurrentIteratorRow", &PixelGetCurrentIteratorRow},
		{"PixelGetCyan", &PixelGetCyan},
		{"PixelGetCyanQuantum", &PixelGetCyanQuantum},
		{"PixelGetException", &PixelGetException},
		{"PixelGetExceptionType", &PixelGetExceptionType},
		{"PixelGetFuzz", &PixelGetFuzz},
		{"PixelGetGreen", &PixelGetGreen},
		{"PixelGetGreenQuantum", &PixelGetGreenQuantum},
		{"PixelGetHSL", &PixelGetHSL},
		{"PixelGetIndex", &PixelGetIndex},
		{"PixelGetIteratorException", &PixelGetIteratorException},
		{"PixelGetIteratorExceptionType", &PixelGetIteratorExceptionType},
		{"PixelGetIteratorRow", &PixelGetIteratorRow},
		{"PixelGetMagenta", &PixelGetMagenta},
		{"PixelGetMagentaQuantum", &PixelGetMagentaQuantum},
		{"PixelGetMagickColor", &PixelGetMagickColor},
		{"PixelGetNextIteratorRow", &PixelGetNextIteratorRow},
		{"PixelGetNextRow", &PixelGetNextRow},
		{"PixelGetOpacity", &PixelGetOpacity},
		{"PixelGetOpacityQuantum", &PixelGetOpacityQuantum},
		{"PixelGetPreviousIteratorRow", &PixelGetPreviousIteratorRow},
		// {"PixelGetPreviousRow", &PixelGetPreviousRow},
		{"PixelGetQuantumColor", &PixelGetQuantumColor},
		{"PixelGetRed", &PixelGetRed},
		{"PixelGetRedQuantum", &PixelGetRedQuantum},
		{"PixelGetYellow", &PixelGetYellow},
		{"PixelGetYellowQuantum", &PixelGetYellowQuantum},
		{"PixelIteratorGetException", &PixelIteratorGetException},
		{"PixelResetIterator", &PixelResetIterator},
		{"PixelSetAlpha", &PixelSetAlpha},
		{"PixelSetAlphaQuantum", &PixelSetAlphaQuantum},
		{"PixelSetBlack", &PixelSetBlack},
		{"PixelSetBlackQuantum", &PixelSetBlackQuantum},
		{"PixelSetBlue", &PixelSetBlue},
		{"PixelSetBlueQuantum", &PixelSetBlueQuantum},
		{"PixelSetColor", &PixelSetColor},
		{"PixelSetColorCount", &PixelSetColorCount},
		{"PixelSetColorFromWand", &PixelSetColorFromWand},
		{"PixelSetCyan", &PixelSetCyan},
		{"PixelSetCyanQuantum", &PixelSetCyanQuantum},
		{"PixelSetFirstIteratorRow", &PixelSetFirstIteratorRow},
		{"PixelSetFuzz", &PixelSetFuzz},
		{"PixelSetGreen", &PixelSetGreen},
		{"PixelSetGreenQuantum", &PixelSetGreenQuantum},
		{"PixelSetHSL", &PixelSetHSL},
		{"PixelSetIndex", &PixelSetIndex},
		{"PixelSetIteratorRow", &PixelSetIteratorRow},
		{"PixelSetLastIteratorRow", &PixelSetLastIteratorRow},
		{"PixelSetMagenta", &PixelSetMagenta},
		{"PixelSetMagentaQuantum", &PixelSetMagentaQuantum},
		{"PixelSetMagickColor", &PixelSetMagickColor},
		{"PixelSetOpacity", &PixelSetOpacity},
		{"PixelSetOpacityQuantum", &PixelSetOpacityQuantum},
		{"PixelSetQuantumColor", &PixelSetQuantumColor},
		{"PixelSetRed", &PixelSetRed},
		{"PixelSetRedQuantum", &PixelSetRedQuantum},
		{"PixelSetYellow", &PixelSetYellow},
		{"PixelSetYellowQuantum", &PixelSetYellowQuantum},
		{"PixelSyncIterator", &PixelSyncIterator},
		{"PopDrawingWand", &PopDrawingWand},
		{"PushDrawingWand", &PushDrawingWand},
		// {"RelinquishWandId", &RelinquishWandId},
		{"SetPixelViewIterator", &SetPixelViewIterator},
		{"SetWandViewDescription", &SetWandViewDescription},
		{"SetWandViewIterator", &SetWandViewIterator},
		{"SetWandViewThreads", &SetWandViewThreads},
		// {"StreamImageCommand", &StreamImageCommand},
		{"TransferPixelViewIterator", &TransferPixelViewIterator},
		{"TransferWandViewIterator", &TransferWandViewIterator},
		{"UpdatePixelViewIterator", &UpdatePixelViewIterator},
		{"UpdateWandViewIterator", &UpdateWandViewIterator},
	})
}
