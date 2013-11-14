// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package imagemagick provides API definitions for
//accessing CORE_RL_magick_.dll.
package core

import (
	. "github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-imagemagick/types"
	. "github.com/tHinqa/outside/types"
)

//Opaque types
type (
	LogInfo        struct{}
	MimeInfo       struct{}
	ResampleFilter struct{}
	SignatureInfo  struct{}
	SplayTreeInfo  struct{}
	StreamInfo     struct{}
)

// Func types
type (
	AcquireMemoryHandler func(T.Size) *T.Void
	DestroyMemoryHandler func(*T.Void)
	ErrorHandler         func(T.ExceptionType, string, string)
	FatalErrorHandler    func(T.ExceptionType, string, string)
	ResizeMemoryHandler  func(*T.Void, T.Size) *T.Void
	SetImageViewMethod   func(*ImageView, T.SSize, int, *T.Void) T.MagickBooleanType
	SetPixelViewMethod   func(*T.PixelView, *T.Void) bool
	SetWandViewMethod    func(*T.WandView, T.SSize, int, *T.Void) bool
	WarningHandler       func(T.ExceptionType, string, string)
)
type (
	fix int

	dirent   fix
	off_t    fix // system-dep 32/64
	stat     fix
	time_t   fix
	timeval  fix
	timezone fix
	wchar_t  fix
	char     int8

	BitStreamReadHandle               fix
	BitStreamWriteHandle              fix
	BlobInfo                          T.BlobInfo
	Cache                             fix
	CacheInfo                         fix // same as Cache?
	CacheView                         T.CacheView
	CommandOption                     T.CommandOption
	CompositeOptions_t                fix
	ConfirmAccessHandler              fix
	ConfirmAccessMode                 fix
	DelegateInfo                      fix
	DifferenceImageOptions            fix
	DifferenceStatistics              fix
	Display                           fix
	DrawContext                       fix
	DrawInfo                          T.DrawInfo
	DrawingWand                       fix
	ExceptionInfo                     T.ExceptionInfo
	ExportPixelAreaInfo               fix
	ExportPixelAreaOptions            fix
	FILE                              T.FILE
	FileIOMode                        fix
	fp_16bits                         [2]byte
	fp_24bits                         [3]byte
	FxInfo                            T.FxInfo
	HashmapInfo                       struct{}
	HighlightStyle                    fix
	HistogramColorPacket              fix
	Image                             T.Image
	ImageCharacteristics              fix
	ImageInfo                         T.ImageInfo
	ImageProfileIterator              fix
	ImageStatistics                   fix
	ImageView                         struct{}
	ImportPixelAreaInfo               fix
	ImportPixelAreaOptions            fix
	LinkedListInfo                    struct{}
	magick_off_t                      fix // system-dep 32/64?
	MagickFreeFunc                    fix
	MagickInfo                        fix
	MagickMallocFunc                  fix
	MagickMap                         fix
	MagickMapIterator                 fix
	MagickMapObjectClone              fix
	MagickMapObjectDeallocator        fix
	MagickPixelPacket                 T.MagickPixelPacket
	MagickRandomKernel                fix
	MagickReallocFunc                 fix
	MagickTextTranslate               fix
	MagickThreadKey                   fix
	MagickTsdKey_t                    fix
	PixelIteratorDualModifyCallback   fix
	PixelIteratorDualNewCallback      fix
	PixelIteratorDualReadCallback     fix
	PixelIteratorMonoModifyCallback   fix
	PixelIteratorMonoReadCallback     fix
	PixelIteratorOptions              fix
	PixelIteratorTripleModifyCallback fix
	PixelIteratorTripleNewCallback    fix
	PolicyInfo                        struct{}
	QuantizeInfo                      T.QuantizeInfo
	QuantumAlphaType                  fix
	QuantumOperator                   fix
	QuantumSampleType                 fix
	ResizeFilter                      struct{}
	ResourceType                      T.ResourceType
	SemaphoreInfo                     T.SemaphoreInfo
	StringInfo                        T.StringInfo
	ThreadViewDataSet                 fix
	ThreadViewSetPtr                  fix
	TimerInfo                         T.TimerInfo
	TokenInfo                         struct{}
	ViewInfo                          fix
	void                              T.Void
	WordStreamReadFunc                fix
	WordStreamReadHandle              fix
	WordStreamWriteFunc               fix
	WordStreamWriteHandle             fix
	WriteByteHook                     fix
	XMLTreeInfo                       fix
	XrmDatabase                       fix
)

// TODO(t): image *Image; images *Image; images **Image distinctions

var AccelerateConvolveImage func(*Image, *T.KernelInfo, *Image, *ExceptionInfo) bool

func (i *Image) AccelerateConvolve(k *T.KernelInfo, i2 *Image, e *ExceptionInfo) bool {
	return AccelerateConvolveImage(i, k, i2, e)
}

var AcquireFxInfo func(i *Image, expression string) *FxInfo

func (i *Image) AcquireFxInfo(expression string) *FxInfo { return AcquireFxInfo(i, expression) }

var AcquireImagePixels func(i *Image, x, y T.Long, columns, rows T.Size, exception *ExceptionInfo) *T.PixelPacket

func (i *Image) AcquireImagePixels(x, y T.Long, columns, rows T.Size, exception *ExceptionInfo) *T.PixelPacket {
	return AcquireImagePixels(i, x, y, columns, rows, exception)
}

var AcquireIndexes func(i *Image) *T.IndexPacket

func (i *Image) AcquireIndexes() *T.IndexPacket { return AcquireIndexes(i) }

var AcquireOneMagickPixel func(i *Image, x, y T.Long, exception *ExceptionInfo) MagickPixelPacket

func (i *Image) AcquireOneMagickPixel(x, y T.Long, exception *ExceptionInfo) MagickPixelPacket {
	return AcquireOneMagickPixel(i, x, y, exception)
}

var AcquireOnePixel func(i *Image, x, y T.Long, exception *ExceptionInfo) T.PixelPacket

func (i *Image) AcquireOnePixel(x, y T.Long, exception *ExceptionInfo) T.PixelPacket {
	return AcquireOnePixel(i, x, y, exception)
}

var AcquireOneVirtualPixel func(i *Image, virtualPixelMethod T.VirtualPixelMethod, x, y T.Long, exception *ExceptionInfo) T.PixelPacket

func (i *Image) AcquireOneVirtualPixel(virtualPixelMethod T.VirtualPixelMethod, x, y T.Long, exception *ExceptionInfo) T.PixelPacket {
	return AcquireOneVirtualPixel(i, virtualPixelMethod, x, y, exception)
}

var AcquireResampleFilter func(i *Image, exception *ExceptionInfo) *ResampleFilter

func (i *Image) AcquireResampleFilter(exception *ExceptionInfo) *ResampleFilter {
	return AcquireResampleFilter(i, exception)
}

var AcquireResizeFilter func(i *Image, filter T.FilterTypes, blur T.MagickRealType, cylindrical bool, exception *ExceptionInfo) *ResizeFilter

func (i *Image) AcquireResizeFilter(filter T.FilterTypes, blur T.MagickRealType, cylindrical bool, exception *ExceptionInfo) *ResizeFilter {
	return AcquireResizeFilter(i, filter, blur, cylindrical, exception)
}

var AdaptiveBlurImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveBlur(radius, sigma float64, exception *ExceptionInfo) *Image {
	return AdaptiveBlurImage(i, radius, sigma, exception)
}

var AdaptiveBlurImageChannel func(i *Image, channel T.ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveBlurChannel(channel T.ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return AdaptiveBlurImageChannel(i, channel, radius, sigma, exception)
}

var AdaptiveResizeImage func(i *Image, columns, rows T.Size, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveResize(columns, rows T.Size, exception *ExceptionInfo) *Image {
	return AdaptiveResizeImage(i, columns, rows, exception)
}

var AdaptiveSharpenImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveSharpen(radius, sigma float64, exception *ExceptionInfo) *Image {
	return AdaptiveSharpenImage(i, radius, sigma, exception)
}

var AdaptiveSharpenImageChannel func(i *Image, channel T.ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveSharpenChannel(channel T.ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return AdaptiveSharpenImageChannel(i, channel, radius, sigma, exception)
}

var AdaptiveThresholdImage func(i *Image, width, height T.Size, offset T.Long, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveThreshold(width, height T.Size, offset T.Long, exception *ExceptionInfo) *Image {
	return AdaptiveThresholdImage(i, width, height, offset, exception)
}

var AddNoiseImage func(i *Image, noiseType T.NoiseType, exception *ExceptionInfo) *Image

func (i *Image) AddNoise(noiseType T.NoiseType, exception *ExceptionInfo) *Image {
	return AddNoiseImage(i, noiseType, exception)
}

var AddNoiseImageChannel func(i *Image, channel T.ChannelType, noiseType T.NoiseType, exception *ExceptionInfo) *Image

func (i *Image) AddNoiseChannel(channel T.ChannelType, noiseType T.NoiseType, exception *ExceptionInfo) *Image {
	return AddNoiseImageChannel(i, channel, noiseType, exception)
}

var AffineTransformImage func(i *Image, affineMatrix *T.AffineMatrix, exception *ExceptionInfo) *Image

func (i *Image) AffineTransform(affineMatrix *T.AffineMatrix, exception *ExceptionInfo) *Image {
	return AffineTransformImage(i, affineMatrix, exception)
}

var AllocateImageColormap func(i *Image, colors T.Size) bool

func (i *Image) AllocateColormap(colors T.Size) bool {
	return AllocateImageColormap(i, colors)
}

var AnnotateImage func(i *Image, drawInfo *DrawInfo) bool

func (i *Image) Annotate(drawInfo *DrawInfo) bool { return AnnotateImage(i, drawInfo) }

var AppendImages func(i *Image, stack bool, exception *ExceptionInfo) *Image

func (i *Image) AppendImages(stack bool, exception *ExceptionInfo) *Image {
	return AppendImages(i, stack, exception)
}

var Ascii85Encode func(i *Image, code byte)

func (i *Image) Ascii85Encode(code byte) { Ascii85Encode(i, code) }

func (i *Image) Ascii85Flush() { Ascii85Flush(i) }

func (i *Image) Ascii85Initialize() { Ascii85Initialize(i) }

func (i *Image) AverageImages(exception *ExceptionInfo) *Image { return AverageImages(i, exception) }

func (i *Image) Bilevel(threshold float64) bool { return BilevelImage(i, threshold) }

func (i *Image) BilevelChannel(channel T.ChannelType, threshold float64) bool {
	return BilevelImageChannel(i, channel, threshold)
}

func (i *Image) BlackThreshold(threshold string) bool { return BlackThresholdImage(i, threshold) }

func (i *Image) Blur(radius, sigma float64, exception *ExceptionInfo) *Image {
	return BlurImage(i, radius, sigma, exception)
}

func (i *Image) BlurChannel(channel T.ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return BlurImageChannel(i, channel, radius, sigma, exception)
}

func (i *Image) Border(borderInfo *T.RectangleInfo, exception *ExceptionInfo) *Image {
	return BorderImage(i, borderInfo, exception)
}

var Ascii85Flush func(i *Image)

var Ascii85Initialize func(i *Image)

var AverageImages func(i *Image, exception *ExceptionInfo) *Image

var BilevelImage func(i *Image, threshold float64) bool

var BilevelImageChannel func(i *Image, channel T.ChannelType, threshold float64) bool

var BlackThresholdImage func(i *Image, threshold string) bool

var BlurImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

var BlurImageChannel func(i *Image, channel T.ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image

var BorderImage func(i *Image, borderInfo *T.RectangleInfo, exception *ExceptionInfo) *Image

var CatchImageException func(i *Image) T.ExceptionType

var ChannelImage func(i *Image, channel T.ChannelType) uint

var ChannelThresholdImage func(i *Image, level string) uint

var CharcoalImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

var ChopImage func(i *Image, chopInfo *T.RectangleInfo, exception *ExceptionInfo) *Image

var ClipImage func(i *Image) bool

var ClipImagePath func(i *Image, pathname string, inside bool) bool

var ClipPathImage func(i *Image, pathname string, inside bool) bool

var CloneImage func(i *Image, columns, rows T.Size, orphan bool, exception *ExceptionInfo) *Image

var CloneImageArtifacts func(i *Image, cloneImage *Image) bool

var CloneImageAttributes func(i *Image, cloneImage *Image) bool

var CloneImageProfiles func(i *Image, cloneImage *Image) bool

var CloneImageProperties func(i *Image, cloneImage *Image) bool

var CloseBlob func(i *Image) bool

var ClutImage func(i *Image, clutImage *Image) bool

var ClutImageChannel func(i *Image, channel T.ChannelType, clutImage *Image) bool

var CoalesceImages func(i *Image, exception *ExceptionInfo) *Image

var ColorFloodfillImage func(i *Image, drawInfo *DrawInfo, target T.PixelPacket, xOffset, yOffset T.Long, method T.PaintMethod) bool

var ColorizeImage func(i *Image, opacity string, colorize T.PixelPacket, exception *ExceptionInfo) *Image

var CombineImages func(i *Image, channel T.ChannelType, exception *ExceptionInfo) *Image

var CompareImageChannels func(i *Image, reconstructImage *Image, channel T.ChannelType, metric T.MetricType, distortion *float64, exception *ExceptionInfo) *Image

var CompareImageLayers func(i *Image, method T.ImageLayerMethod, exception *ExceptionInfo) *Image

var CompareImages func(i *Image, reconstructImage *Image, metric T.MetricType, distortion *float64, exception *ExceptionInfo) *Image

var CompositeImage func(i *Image, compose T.CompositeOperator, compositeImage *Image, xOffset, yOffset T.Long) bool

var CompositeImageChannel func(i *Image, channel T.ChannelType, compose T.CompositeOperator, compositeImage *Image, xOffset, yOffset T.Long) bool

var CompositeLayers func(i *Image, compose T.CompositeOperator, source *Image, xOffset, yOffset T.Long, exception *ExceptionInfo)

var CompressImageColormap func(i *Image)

var ContrastImage func(i *Image, sharpen bool) bool

var ContrastStretchImage func(i *Image, levels string) bool

var ContrastStretchImageChannel func(i *Image, channel T.ChannelType, blackPoint, whitePoint float64) bool

var ConvolveImage func(i *Image, order T.Size, kernel *float64, exception *ExceptionInfo) *Image

var ConvolveImageChannel func(i *Image, channel T.ChannelType, order T.Size, kernel *float64, exception *ExceptionInfo) *Image

var CropImage func(i *Image, geometry *T.RectangleInfo, exception *ExceptionInfo) *Image

var CycleColormapImage func(i *Image, displace T.Long) bool

var DefineImageArtifact func(i *Image, artifact string) bool

var DefineImageProperty func(i *Image, property string) bool

var DeleteImageArtifact func(i *Image, artifact string) bool

var DeleteImageAttribute func(i *Image, key string) bool

var DeleteImageProfile func(i *Image, name string) bool

var DeleteImageProperty func(i *Image, property string) bool

var DescribeImage func(i *Image, file *FILE, verbose bool) bool

var DespeckleImage func(i *Image, exception *ExceptionInfo) *Image

var DestroyBlob func(i *Image)

var DestroyImage func(i *Image) *Image

var DestroyImageArtifacts func(i *Image)

var DestroyImageAttributes func(i *Image)

var DestroyImagePixels func(i *Image)

var DestroyImageProfiles func(i *Image)

var DestroyImageProperties func(i *Image)

var DestroyImages func(i *Image)

var DisassociateImageStream func(i *Image)

var DispatchImage func(i *Image, xOffset, yOffset T.Long, columns, rows T.Size, map_ string, type_ T.StorageType, pixels *T.Void, exception *ExceptionInfo) uint

var DisposeImages func(i *Image, exception *ExceptionInfo) *Image

var DistortImage func(i *Image, method T.DistortImageMethod, numberArguments T.Size, arguments *float64, bestfit bool, exception *ExceptionInfo) *Image

var DrawAffineImage func(i *Image, source *Image, affine *T.AffineMatrix) bool

var DrawClipPath func(i *Image, drawInfo *DrawInfo, name string) bool

var DrawGradientImage func(i *Image, drawInfo *DrawInfo) bool

var DrawImage func(i *Image, drawInfo *DrawInfo) bool

var DrawPatternPath func(i *Image, drawInfo *DrawInfo, name string, pattern **Image) bool

var DrawPrimitive func(i *Image, drawInfo *DrawInfo, primitiveInfo *T.PrimitiveInfo) bool

var EdgeImage func(i *Image, radius float64, exception *ExceptionInfo) *Image

var EmbossImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

var EnhanceImage func(i *Image, exception *ExceptionInfo) *Image

var EOFBlob func(i *Image) int

var EqualizeImage func(i *Image) bool

var EqualizeImageChannel func(i *Image, channel T.ChannelType) bool

var EvaluateImage func(i *Image, op T.MagickEvaluateOperator, value float64, exception *ExceptionInfo) bool

var EvaluateImageChannel func(i *Image, channel T.ChannelType, op T.MagickEvaluateOperator, value float64, exception *ExceptionInfo) bool

var ExcerptImage func(i *Image, geometry *T.RectangleInfo, exception *ExceptionInfo) *Image

var ExportImagePixels func(i *Image, xOffset, yOffset T.Long, columns, rows T.Size, map_ string, type_ T.StorageType, pixels *T.Void, exception *ExceptionInfo) bool

var ExportQuantumPixels func(i *Image, quantumInfo *T.QuantumInfo, quantumType T.QuantumType, pixels *byte) bool

var ExtentImage func(i *Image, geometry *T.RectangleInfo, exception *ExceptionInfo) *Image

var FileToImage func(i *Image, filename string) bool

var FlattenImages func(i *Image, exception *ExceptionInfo) *Image

var FlipImage func(i *Image, exception *ExceptionInfo) *Image

var FloodfillPaintImage func(i *Image, channel T.ChannelType, drawInfo *DrawInfo, target *MagickPixelPacket, xOffset, yOffset T.Long, invert bool) bool

var FlopImage func(i *Image, exception *ExceptionInfo) *Image

var FormatImageAttribute func(i *Image, key, format string, va ...VArg) bool

var FormatImageAttributeList func(i *Image, key, format string, operands VAList) bool

var FormatImageProperty func(i *Image, property, format string, va ...VArg) bool

var FormatImagePropertyList func(i *Image, property, format string, operands VAList) bool

var FormatMagickCaption func(i *Image, drawInfo *DrawInfo, caption string, metrics *T.TypeMetric) T.Long

var FrameImage func(i *Image, frameInfo *T.FrameInfo, exception *ExceptionInfo) *Image

var FuzzyColorCompare func(i *Image, p, q *T.PixelPacket) bool

var FuzzyOpacityCompare func(i *Image, p, q *T.PixelPacket) bool

var FxImage func(i *Image, expression string, exception *ExceptionInfo) *Image

var FxImageChannel func(i *Image, channel T.ChannelType, expression string, exception *ExceptionInfo) *Image

var GammaImage func(i *Image, level string) bool

var GammaImageChannel func(i *Image, channel T.ChannelType, gamma float64) bool

var GaussianBlurImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

var GaussianBlurImageChannel func(i *Image, channel T.ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image

var GetBlobError func(i *Image) bool

var GetBlobFileHandle func(i *Image) *FILE

var GetBlobSize func(i *Image) T.MagickSizeType

var GetBlobStreamData func(i *Image) *byte

var GetBlobStreamHandler func(i *Image) T.StreamHandler

var GetImageArtifact func(i *Image, artifact string) string

var GetImageAttribute func(i *Image, key string) *T.ImageAttribute

var GetImageBoundingBox func(i *Image, exception *ExceptionInfo) T.RectangleInfo

var GetImageChannelDepth func(i *Image, channel T.ChannelType, exception *ExceptionInfo) T.Size

var GetImageChannelDistortion func(i *Image, reconstructImage *Image, channel T.ChannelType, metric T.MetricType, distortion *float64, exception *ExceptionInfo) bool

var GetImageChannelExtrema func(i *Image, channel T.ChannelType, minima, maxima *T.Size, exception *ExceptionInfo) bool

var GetImageChannelMean func(i *Image, channel T.ChannelType, mean, standardDeviation *float64, exception *ExceptionInfo) bool

var GetImageChannelRange func(i *Image, channel T.ChannelType, minima, maxima *float64, exception *ExceptionInfo) bool

var GetImageChannelStatistics func(i *Image, exception *ExceptionInfo) *T.ChannelStatistics

var GetImageClipMask func(i *Image, exception *ExceptionInfo) *Image

var GetImageClippingPathAttribute func(i *Image) *T.ImageAttribute

var GetImageDepth func(i *Image, exception *ExceptionInfo) T.Size

var GetImageDistortion func(i *Image, reconstructImage *Image, metric T.MetricType, distortion *float64, exception *ExceptionInfo) bool

var GetImageDynamicThreshold func(i *Image, clusterThreshold, smoothThreshold float64, exception *ExceptionInfo) MagickPixelPacket

var GetImageException func(i *Image, exception *ExceptionInfo)

var GetImageExtrema func(i *Image, minima, maxima *T.Size, exception *ExceptionInfo) bool

var GetImageGeometry func(i *Image, geometry string, sizeToFit uint, regionInfo *T.RectangleInfo) int

var GetImageHistogram func(i *Image, numberColors *T.Size, exception *ExceptionInfo) *T.ColorPacket

var GetImageMask func(i *Image, exception *ExceptionInfo) *Image

var GetImageMean func(i *Image, mean, standardDeviation *float64, exception *ExceptionInfo) bool

var GetImagePixels func(i *Image, x, y T.Long, columns, rows T.Size) *T.PixelPacket

var GetImageProfile func(i *Image, name string) *StringInfo

var GetImageProperty func(i *Image, property string) string

var GetImageQuantizeError func(i *Image) bool

var GetImageQuantumDepth func(i *Image, constrain bool) T.Size

var GetImageRange func(i *Image, minima, maxima *float64, exception *ExceptionInfo) bool

var GetImageTotalInkDensity func(i *Image) float64

var GetImageType func(i *Image, exception *ExceptionInfo) T.ImageType

var GetImageVirtualPixelMethod func(i *Image) T.VirtualPixelMethod

var GetIndexes func(i *Image) *T.IndexPacket

var GetMagickPixelPacket func(i *Image, pixel *MagickPixelPacket)

var GetMultilineTypeMetrics func(i *Image, drawInfo *DrawInfo, metrics *T.TypeMetric) bool

var GetNextImageArtifact func(i *Image) string

var GetNextImageAttribute func(i *Image) *T.ImageAttribute

var GetNextImageProfile func(i *Image) string

var GetNextImageProperty func(i *Image) string

var GetNumberColors func(i *Image, file *FILE, exception *ExceptionInfo) T.Size

var GetNumberScenes func(i *Image) uint

var GetOnePixel func(i *Image, x, y T.Long) T.PixelPacket

var GetPixelCacheArea func(i *Image) T.MagickSizeType

var GetPixels func(i *Image) *T.PixelPacket

var GetTypeMetrics func(i *Image, drawInfo *DrawInfo, metrics *T.TypeMetric) bool

var GradientImage func(i *Image, startColor, stopColor *T.PixelPacket) bool

var HuffmanDecodeImage func(i *Image) bool

var IdentifyImage func(i *Image, file *FILE, verbose bool) bool

var ImageToFile func(i *Image, filename string, exception *ExceptionInfo) bool

var ImplodeImage func(i *Image, amount float64, exception *ExceptionInfo) *Image

var ImportImagePixels func(i *Image, xOffset, yOffset T.Long, columns, rows T.Size, map_ string, type_ T.StorageType, pixels *T.Void) bool

var ImportQuantumPixels func(i *Image, quantumInfo *T.QuantumInfo, quantumType T.QuantumType, pixels *byte) bool

var InterpolatePixelColor func(i *Image, imageView *CacheView, method T.InterpolatePixelMethod, x, y float64, exception *ExceptionInfo) MagickPixelPacket

var IsBlobExempt func(i *Image) bool

var IsBlobSeekable func(i *Image) bool

var IsBlobTemporary func(i *Image) bool

var IsColorSimilar func(i *Image, p, q *T.PixelPacket) bool

var IsGrayImage func(i *Image, exception *ExceptionInfo) bool

var IsHighDynamicRangeImage func(i *Image, exception *ExceptionInfo) bool

var IsHistogramImage func(i *Image, exception *ExceptionInfo) bool

var IsImageObject func(i *Image) bool

var IsImagesEqual func(i *Image, reconstructImage *Image) bool

var IsImageSimilar func(i *Image, targetImage *Image, xOffset, yOffset *T.Long, exception *ExceptionInfo) bool

var IsMonochromeImage func(i *Image, exception *ExceptionInfo) bool

var IsOpacitySimilar func(i *Image, p, q *T.PixelPacket) bool

var IsOpaqueImage func(i *Image, exception *ExceptionInfo) bool

var IsPaletteImage func(i *Image, exception *ExceptionInfo) bool

var IsTaintImage func(i *Image) bool

var LevelImage func(i *Image, levels string) bool

var LevelImageChannel func(i *Image, channel T.ChannelType, blackPoint, whitePoint, gamma float64) bool

var LinearStretchImage func(i *Image, blackPoint, whitePoint float64) bool

var LiquidRescaleImage func(i *Image, columns, rows T.Size, deltaX, rigidity float64, exception *ExceptionInfo) *Image

var LZWEncodeImage func(i *Image, length uint32, pixels *byte) bool

var MagnifyImage func(i *Image, exception *ExceptionInfo) *Image

var MapImage func(i *Image, mapImage *Image, dither bool) bool

var MatteFloodfillImage func(i *Image, target T.PixelPacket, opacity T.Quantum, xOffset, yOffset T.Long, method T.PaintMethod) bool

var MedianFilterImage func(i *Image, radius float64, exception *ExceptionInfo) *Image

var MergeImageLayers func(i *Image, method T.ImageLayerMethod, exception *ExceptionInfo) *Image

var MinifyImage func(i *Image, exception *ExceptionInfo) *Image

var ModulateImage func(i *Image, modulate string) bool

var MorphImages func(i *Image, numberFrames T.Size, exception *ExceptionInfo) *Image

var MosaicImages func(i *Image, exception *ExceptionInfo) *Image

var MotionBlurImage func(i *Image, radius, sigma, angle float64, exception *ExceptionInfo) *Image

var NegateImage func(i *Image, grayscale bool) bool

var NegateImageChannel func(i *Image, channel T.ChannelType, grayscale bool) bool

var NormalizeImage func(i *Image) bool

var NormalizeImageChannel func(i *Image, channel T.ChannelType) bool

var OilPaintImage func(i *Image, radius float64, exception *ExceptionInfo) *Image

var OpaqueImage func(i *Image, target, fill T.PixelPacket) bool

var OpaquePaintImage func(i *Image, target, fill *MagickPixelPacket, invert bool) bool

var OpaquePaintImageChannel func(i *Image, channel T.ChannelType, target, fill *MagickPixelPacket, invert bool) bool

var OpenCacheView func(i *Image) *CacheView

var OptimizeImageLayers func(i *Image, exception *ExceptionInfo) *Image

var OptimizeImageTransparency func(i *Image, exception *ExceptionInfo)

var OptimizePlusImageLayers func(i *Image, exception *ExceptionInfo) *Image

var OrderedDitherImage func(i *Image) bool

var OrderedDitherImageChannel func(i *Image, channel T.ChannelType, exception *ExceptionInfo) bool

var OrderedPosterizeImage func(i *Image, thresholdMap string, exception *ExceptionInfo) bool

var OrderedPosterizeImageChannel func(i *Image, channel T.ChannelType, thresholdMap string, exception *ExceptionInfo) bool

var PackbitsEncodeImage func(i *Image, length uint32, pixels *byte) bool

var PaintFloodfillImage func(i *Image, channel T.ChannelType, target *MagickPixelPacket, x, y T.Long, drawInfo *DrawInfo, method T.PaintMethod) bool

var PaintOpaqueImage func(i *Image, target, fill *MagickPixelPacket) bool

var PaintOpaqueImageChannel func(i *Image, channel T.ChannelType, target, fill *MagickPixelPacket) bool

var PaintTransparentImage func(i *Image, target *MagickPixelPacket, opacity T.Quantum) bool

var ParseGravityGeometry func(i *Image, geometry string, regionInfo *T.RectangleInfo) T.MagickStatusType

var ParsePageGeometry func(i *Image, geometry string, regionInfo *T.RectangleInfo) T.MagickStatusType

var PersistCache func(i *Image, filename string, attach bool, offset *T.MagickOffsetType, exception *ExceptionInfo) bool

var PlasmaImage func(i *Image, segment *T.SegmentInfo, attenuate, depth T.Size) bool

var PolaroidImage func(i *Image, drawInfo *DrawInfo, angle float64, exception *ExceptionInfo) *Image

var PopImagePixels func(i *Image, quantum T.QuantumType, destination *byte) bool

var PosterizeImage func(i *Image, levels T.Size, dither bool) bool

var PreviewImage func(i *Image, preview T.PreviewType, exception *ExceptionInfo) *Image

var ProfileImage func(i *Image, name string, datum *T.Void, length uint32, clone bool) bool

var PushImagePixels func(i *Image, quantum T.QuantumType, source *byte) bool

var QuantizationError func(i *Image) uint

var QueryColorname func(i *Image, color *T.PixelPacket, compliance T.ComplianceType, name string, exception *ExceptionInfo) bool

var QueryMagickColorname func(i *Image, color *MagickPixelPacket, compliance T.ComplianceType, hex bool, name string, exception *ExceptionInfo) bool

var RadialBlurImage func(i *Image, angle float64, exception *ExceptionInfo) *Image

var RadialBlurImageChannel func(i *Image, channel T.ChannelType, angle float64, exception *ExceptionInfo) *Image

var RaiseImage func(i *Image, raiseInfo *T.RectangleInfo, raise bool) bool

var RandomChannelThresholdImage func(i *Image, channel, thresholds string, exception *ExceptionInfo) uint

var RandomThresholdImage func(i *Image, thresholds string, exception *ExceptionInfo) bool

var RandomThresholdImageChannel func(i *Image, channel T.ChannelType, thresholds string, exception *ExceptionInfo) bool

var ReadBlob func(i *Image, length uint32, data *byte) int32

var ReadBlobByte func(i *Image) int

var ReadBlobDouble func(i *Image) float64

var ReadBlobFloat func(i *Image) float32

var ReadBlobLong func(i *Image) T.Size

var ReadBlobLongLong func(i *Image) T.MagickSizeType

var ReadBlobLSBLong func(i *Image) T.Size

var ReadBlobLSBShort func(i *Image) T.UnsignedShort

var ReadBlobMSBLong func(i *Image) T.Size

var ReadBlobMSBShort func(i *Image) T.UnsignedShort

var ReadBlobShort func(i *Image) T.UnsignedShort

var ReadBlobString func(i *Image, str string) string

var RecolorImage func(i *Image, order T.Size, colorMatrix *float64, exception *ExceptionInfo) *Image

var ReduceNoiseImage func(i *Image, radius float64, exception *ExceptionInfo) *Image

var ReferenceImage func(i *Image) *Image

var RemoveImageArtifact func(i *Image, artifact string) string

var RemoveImageProfile func(i *Image, name string) *StringInfo

var RemoveImageProperty func(i *Image, property string) string

var ResampleImage func(i *Image, xResolution, yResolution float64, filter T.FilterTypes, blur float64, exception *ExceptionInfo) *Image

var ResetImageArtifactIterator func(i *Image)

var ResetImageAttributeIterator func(i *Image)

var ResetImagePage func(i *Image, page string) bool

var ResetImageProfileIterator func(i *Image)

var ResetImagePropertyIterator func(i *Image)

var ResizeImage func(i *Image, columns, rows T.Size, filter T.FilterTypes, blur float64, exception *ExceptionInfo) *Image

var RGBTransformImage func(i *Image, colorspace T.ColorspaceType) bool

var RollImage func(i *Image, xOffset, yOffset T.Long, exception *ExceptionInfo) *Image

var RotateImage func(i *Image, degrees float64, exception *ExceptionInfo) *Image

var SampleImage func(i *Image, columns, rows T.Size, exception *ExceptionInfo) *Image

var ScaleImage func(i *Image, columns, rows T.Size, exception *ExceptionInfo) *Image

var SeekBlob func(i *Image, offset T.MagickOffsetType, whence int) T.MagickOffsetType

var SegmentImage func(i *Image, colorspace T.ColorspaceType, verbose bool, clusterThreshold, smoothThreshold float64) bool

var SeparateImageChannel func(i *Image, channel T.ChannelType) bool

var SeparateImages func(i *Image, channel T.ChannelType, exception *ExceptionInfo) *Image

var SepiaToneImage func(i *Image, threshold float64, exception *ExceptionInfo) *Image

var SetBlobExempt func(i *Image, exempt bool)

var SetGeometry func(i *Image, geometry *T.RectangleInfo)

var SetImage func(i *Image, opacity T.Quantum)

var SetImageAlphaChannel func(i *Image, alphaType T.AlphaChannelType) bool

var SetImageArtifact func(i *Image, artifact, value string) bool

var SetImageAttribute func(i *Image, key, value string) bool

var SetImageBackgroundColor func(i *Image) bool

var SetImageChannelDepth func(i *Image, channel T.ChannelType, depth T.Size) bool

var SetImageClipMask func(i *Image, clipMask *Image) bool

var SetImageDepth func(i *Image, depth T.Size) bool

var SetImageExtent func(i *Image, columns, rows T.Size) bool

var SetImageMask func(i *Image, mask *Image) bool

var SetImageOpacity func(i *Image, opacity T.Quantum) bool

var SetImagePixels func(i *Image, x, y T.Long, columns, rows T.Size) *T.PixelPacket

var SetImageProfile func(i *Image, name string, profile *StringInfo) bool

var SetImageProgressMonitor func(i *Image, progressMonitor T.MagickProgressMonitor, clientData *T.Void) T.MagickProgressMonitor

var SetImageProperty func(i *Image, property, value string) bool

var SetImageStorageClass func(i *Image, storageClass T.ClassType) bool

var SetImageType func(i *Image, imageType T.ImageType) bool

var SetImageVirtualPixelMethod func(i *Image, virtualPixelMethod T.VirtualPixelMethod) T.VirtualPixelMethod

var ShadeImage func(i *Image, gray bool, azimuth, elevation float64, exception *ExceptionInfo) *Image

var ShadowImage func(i *Image, opacity, sigma float64, xOffset, yOffset T.Long, exception *ExceptionInfo) *Image

var SharpenImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

var SharpenImageChannel func(i *Image, channel T.ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image

var ShaveImage func(i *Image, shaveInfo *T.RectangleInfo, exception *ExceptionInfo) *Image

var ShearImage func(i *Image, xShear, yShear float64, exception *ExceptionInfo) *Image

var SigmoidalContrastImage func(i *Image, sharpen bool, levels string) bool

var SigmoidalContrastImageChannel func(i *Image, channel T.ChannelType, sharpen bool, contrast, midpoint float64) bool

var SignatureImage func(i *Image) bool

var SizeBlob func(i *Image) T.MagickOffsetType

var SketchImage func(i *Image, radius, sigma, angle float64, exception *ExceptionInfo) *Image

var SolarizeImage func(i *Image, threshold float64) bool

var SortColormapByIntensity func(i *Image) bool

var SpliceImage func(i *Image, geometry *T.RectangleInfo, exception *ExceptionInfo) *Image

var SpreadImage func(i *Image, radius float64, exception *ExceptionInfo) *Image

var SteganoImage func(i *Image, watermark *Image, exception *ExceptionInfo) *Image

var StereoImage func(i *Image, offsetImage *Image, exception *ExceptionInfo) *Image

var StripImage func(i *Image) bool

var SwirlImage func(i *Image, degrees float64, exception *ExceptionInfo) *Image

var SyncImage func(i *Image) bool

var SyncImagePixels func(i *Image) bool

var SyncImageProfiles func(i *Image) bool

var TellBlob func(i *Image) T.MagickOffsetType

var TextureImage func(i *Image, texture *Image) bool

var ThresholdImage func(i *Image, threshold float64) uint

var ThresholdImageChannel func(i *Image, threshold string) uint

var ThumbnailImage func(i *Image, columns, rows T.Size, exception *ExceptionInfo) *Image

var TintImage func(i *Image, opacity string, tint T.PixelPacket, exception *ExceptionInfo) *Image

var TransformColorspace func(i *Image, colorspace T.ColorspaceType) uint

var TransformImageColorspace func(i *Image, colorspace T.ColorspaceType) bool

var TransformRGBImage func(i *Image, colorspace T.ColorspaceType) bool

var TransparentImage func(i *Image, target T.PixelPacket, opacity T.Quantum) bool

var TransparentPaintImage func(i *Image, target *MagickPixelPacket, opacity T.Quantum, invert bool) bool

var TransposeImage func(i *Image, exception *ExceptionInfo) *Image

var TransverseImage func(i *Image, exception *ExceptionInfo) *Image

var TrimImage func(i *Image, exception *ExceptionInfo) *Image

var UniqueImageColors func(i *Image, exception *ExceptionInfo) *Image

var UnsharpMaskImage func(i *Image, radius, sigma, amount, threshold float64, exception *ExceptionInfo) *Image

var UnsharpMaskImageChannel func(i *Image, channel T.ChannelType, radius, sigma, amount, threshold float64, exception *ExceptionInfo) *Image

var ValidateColormapIndex func(i *Image, index T.Size) T.IndexPacket

var VignetteImage func(i *Image, radius, sigma float64, x, y T.Long, exception *ExceptionInfo) *Image

var WaveImage func(i *Image, amplitude, waveLength float64, exception *ExceptionInfo) *Image

var WhiteThresholdImage func(i *Image, threshold string) bool

var WriteBlob func(i *Image, length uint32, data *byte) int32

var WriteBlobByte func(i *Image, value byte) int32

var WriteBlobFloat func(i *Image, value float32) int32

var WriteBlobLong func(i *Image, value T.UnsignedLong) int32

var WriteBlobLSBLong func(i *Image, value T.UnsignedLong) int32

var WriteBlobLSBShort func(i *Image, value T.UnsignedShort) int32

var WriteBlobMSBLong func(i *Image, value T.UnsignedLong) int32

var WriteBlobMSBShort func(i *Image, value T.UnsignedShort) int32

var WriteBlobShort func(i *Image, value T.UnsignedShort) int32

var WriteBlobString func(i *Image, str string) int32

var ZLIBEncodeImage func(i *Image, length uint32, pixels *byte) bool

var ZoomImage func(i *Image, columns, rows T.Size, exception *ExceptionInfo) *Image

var AcquireAuthenticCacheView func(i *Image, exception *ExceptionInfo) *CacheView

var AcquireCacheView func(i *Image) *CacheView

var AcquireImageColormap func(i *Image, colors uint32) bool

var AcquirePixelCache func(*Image, T.VirtualPixelMethod, T.Long, T.Long, T.Size, T.Size, *ExceptionInfo) *T.PixelPacket // doc c-static
var AcquirePixels func(i *Image) *T.PixelPacket                                                                         // doc (image Image)
var AcquireVirtualCacheView func(i *Image, exception *ExceptionInfo) *CacheView

var AutoGammaImage func(i *Image) bool

var AutoGammaImageChannel func(i *Image, channel T.ChannelType) bool

var AutoLevelImage func(i *Image) bool

var AutoLevelImageChannel func(i *Image, channel T.ChannelType) bool

var BlueShiftImage func(i *Image, factor float64, exception *ExceptionInfo) *Image

var BrightnessContrastImage func(i *Image, brightness, contrast float64) bool

var BrightnessContrastImageChannel func(i *Image, channel T.ChannelType, brightness, contrast float64) bool

var ColorDecisionListImage func(i *Image, colorCorrectionCollection string) bool

var ColorMatrixImage func(i *Image, colorMatrix *T.KernelInfo, exception *ExceptionInfo) *Image

var CropImageToHBITMAP func(i *Image, r *T.RectangleInfo, e *ExceptionInfo) *T.Void

var CropImageToTiles func(i *Image, cropGeometry *T.RectangleInfo, exception *ExceptionInfo) *Image

var DecipherImage func(i *Image, passphrase string, exception *ExceptionInfo) bool

var DeskewImage func(i *Image, threshold float64, exception *ExceptionInfo) *Image

var EncipherImage func(i *Image, passphrase string, exception *ExceptionInfo) bool

var ExtractSubimageFromImage func(i *Image, reference *Image, exception *ExceptionInfo) *Image

var FilterImage func(i *Image, kernel *T.KernelInfo, exception *ExceptionInfo) *Image

var FilterImageChannel func(i *Image, channel T.ChannelType, kernel *T.KernelInfo, exception *ExceptionInfo) *Image

var ForwardFourierTransformImage func(i *Image, modulus bool, exception *ExceptionInfo) *Image

var FunctionImage func(i *Image, function T.MagickFunction, numberParameters int32, parameters *float64, exception *ExceptionInfo) bool

var FunctionImageChannel func(i *Image, channel T.ChannelType, function T.MagickFunction, numberParameters int32, argument *float64, exception *ExceptionInfo) bool

var GetAuthenticIndexQueue func(i *Image) *T.IndexPacket

var GetAuthenticPixelQueue func(i *Image) *T.PixelPacket // doc (image Image)
var GetAuthenticPixels func(i *Image, x, y int32, columns, rows uint32, exception *ExceptionInfo) *T.PixelPacket

var GetBlobProperties func(i *Image) *stat

var GetImageAlphaChannel func(i *Image) bool

var GetImageChannelDistortions func(i *Image, reconstructImage *Image, metric T.MetricType, exception *ExceptionInfo) *float64

var GetImageChannelFeatures func(i *Image, distance uint32, exception *ExceptionInfo) *T.ChannelFeatures

var GetImageChannelKurtosis func(i *Image, channel T.ChannelType, kurtosis, kewness *float64, exception *ExceptionInfo) bool

var GetImageChannels func(i *Image) uint32

var GetOneAuthenticPixel func(i *Image, x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) bool                                         // doc (image Image)
var GetOneVirtualMagickPixel func(i *Image, x, y int32, pixel *MagickPixelPacket, exception *ExceptionInfo) bool                                 // doc (image Image...exception ExceptionInfo)
var GetOneVirtualMethodPixel func(i *Image, virtualPixelMethod T.VirtualPixelMethod, x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) // doc (image Image...Pixelpacket...exception ExceptionInfo)
var GetOneVirtualPixel func(i *Image, x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) bool                                           // doc (image Image...exception ExceptionInfo)
var GetVirtualIndexQueue func(i *Image) *T.IndexPacket

var GetVirtualPixelQueue func(i *Image) *T.PixelPacket // doc (image Image)
var GetVirtualPixels func(i *Image, x, y int32, columns, rows uint32, exception *ExceptionInfo) *T.PixelPacket

var HaldClutImage func(i *Image, haldImage *Image) bool

var HaldClutImageChannel func(i *Image, channel T.ChannelType, haldImage *Image) bool

var ImageToHBITMAP func(i *Image) **T.Void

var IntegralRotateImage func(i *Image, rotations uint32, exception *ExceptionInfo) *Image

var InverseFourierTransformImage func(i *Image, phaseImage *Image, modulus bool, exception *ExceptionInfo) *Image

var LevelColorsImage func(i *Image, blackColor, whiteColor *MagickPixelPacket, invert bool) bool

var LevelColorsImageChannel func(i *Image, channel T.ChannelType, blackColor, whiteColor *MagickPixelPacket, invert bool) bool

var LevelImageColors func(i *Image, channel T.ChannelType, blackColor, whiteColor *MagickPixelPacket, invert bool) bool

var LevelizeImage func(i *Image, blackPoint, whitePoint, gamma float64) bool

var LevelizeImageChannel func(i *Image, channel T.ChannelType, blackPoint, whitePoint, gamma float64) bool

var MinMaxStretchImage func(i *Image, channel T.ChannelType, blackAdjust, whiteAdjust float64) bool

var ModeImage func(i *Image, radius float64, exception *ExceptionInfo) *Image

var MorphologyApply func(i *Image, method T.MorphologyMethod, channel T.ChannelType, iterations int32, kernel *T.KernelInfo, compose T.CompositeMethod, bias float64, exception *ExceptionInfo) *Image

var MorphologyImage func(i *Image, method T.MorphologyMethod, iterations int32, kernel *T.KernelInfo, exception *ExceptionInfo) *Image

var MorphologyImageChannel func(i *Image, channel T.ChannelType, method T.MorphologyMethod, iterations int32, kernel *T.KernelInfo, exception *ExceptionInfo) *Image

var MotionBlurImageChannel func(i *Image, channel T.ChannelType, radius, sigma, angle float64, exception *ExceptionInfo) *Image

var ParseSizeGeometry func(i *Image, geometry string, regionInfo *T.RectangleInfo) T.MagickStatusType // doc RectangeInfo
var PasskeyDecipherImage func(i *Image, passkey *StringInfo, exception *ExceptionInfo) bool

var PasskeyEncipherImage func(i *Image, passkey *StringInfo, exception *ExceptionInfo) bool

var PosterizeImageChannel func(i *Image, channel T.ChannelType, levels uint32, dither bool) bool

var QueueAuthenticPixels func(i *Image, x, y int32, columns, rows uint32, exception *ExceptionInfo) *T.PixelPacket

var SelectiveBlurImage func(i *Image, radius, sigma, threshold float64, exception *ExceptionInfo) *Image

var SelectiveBlurImageChannel func(i *Image, channel T.ChannelType, radius, sigma, threshold float64, exception *ExceptionInfo) *Image

var SetImageChannels func(i *Image, channels uint32) bool

var SetImageColor func(i *Image, color *MagickPixelPacket) bool

var SetImageColorspace func(i *Image, colorspace T.ColorspaceType) // return ???
var ShearRotateImage func(i *Image, degrees float64, exception *ExceptionInfo) *Image

var SimilarityImage func(i *Image, reference *Image, offset *T.RectangleInfo, similarity *float64, exception *ExceptionInfo) *Image

var SparseColorImage func(i *Image, channel T.ChannelType, method T.SparseColorMethod, numberArguments uint32, arguments *float64, exception *ExceptionInfo) *Image

var StatisticImage func(i *Image, type_ T.StatisticType, width, height uint32, exception *ExceptionInfo) *Image

var StatisticImageChannel func(i *Image, channel T.ChannelType, type_ T.StatisticType, width, height uint32, exception *ExceptionInfo) *Image

var StereoAnaglyphImage func(i *Image, rightImage *Image, xOffset, yOffset int32, exception *ExceptionInfo) *Image

var SyncAuthenticPixels func(i *Image, exception *ExceptionInfo) bool

func (i *Image) CatchException() T.ExceptionType { return CatchImageException(i) }

func (i *Image) Channel(channel T.ChannelType) uint { return ChannelImage(i, channel) }

func (i *Image) ChannelThreshold(level string) uint { return ChannelThresholdImage(i, level) }

func (i *Image) Charcoal(radius, sigma float64, exception *ExceptionInfo) *Image {
	return CharcoalImage(i, radius, sigma, exception)
}

func (i *Image) Chop(chopInfo *T.RectangleInfo, exception *ExceptionInfo) *Image {
	return ChopImage(i, chopInfo, exception)
}

func (i *Image) Clip() bool { return ClipImage(i) }

func (i *Image) ClipImagePath(pathname string, inside bool) bool {
	return ClipImagePath(i, pathname, inside)
}

func (i *Image) ClipPathImage(pathname string, inside bool) bool {
	return ClipPathImage(i, pathname, inside)
}

func (i *Image) Clone(columns, rows T.Size, orphan bool, exception *ExceptionInfo) *Image {
	return CloneImage(i, columns, rows, orphan, exception)
}

func (i *Image) CloneArtifacts(cloneImage *Image) bool {
	return CloneImageArtifacts(i, cloneImage)
}

func (i *Image) CloneAttributes(cloneImage *Image) bool {
	return CloneImageAttributes(i, cloneImage)
}

func (i *Image) CloneProfiles(cloneImage *Image) bool {
	return CloneImageProfiles(i, cloneImage)
}

func (i *Image) CloneImageProperties(cloneImage *Image) bool {
	return CloneImageProperties(i, cloneImage)
}

func (i *Image) CloseBlob() bool { return CloseBlob(i) }

func (i *Image) Clut(clutImage *Image) bool { return ClutImage(i, clutImage) }

func (i *Image) ClutChannel(channel T.ChannelType, clutImage *Image) bool {
	return ClutImageChannel(i, channel, clutImage)
}

func (i *Image) CoalesceImages(exception *ExceptionInfo) *Image { return CoalesceImages(i, exception) }

func (i *Image) ColorFloodfill(drawInfo *DrawInfo, target T.PixelPacket, xOffset, yOffset T.Long, method T.PaintMethod) bool {
	return ColorFloodfillImage(i, drawInfo, target, xOffset, yOffset, method)
}

func (i *Image) Colorize(opacity string, colorize T.PixelPacket, exception *ExceptionInfo) *Image {
	return ColorizeImage(i, opacity, colorize, exception)
}

func (i *Image) CombineImages(channel T.ChannelType, exception *ExceptionInfo) *Image {
	return CombineImages(i, channel, exception)
}

func (i *Image) CompareChannels(reconstructImage *Image, channel T.ChannelType, metric T.MetricType, distortion *float64, exception *ExceptionInfo) *Image {
	return CompareImageChannels(i, reconstructImage, channel, metric, distortion, exception)
}

func (i *Image) CompareImageLayers(method T.ImageLayerMethod, exception *ExceptionInfo) *Image {
	return CompareImageLayers(i, method, exception)
}

func (i *Image) CompareImages(reconstructImage *Image, metric T.MetricType, distortion *float64, exception *ExceptionInfo) *Image {
	return CompareImages(i, reconstructImage, metric, distortion, exception)
}

func (i *Image) CompositeImage(compose T.CompositeOperator, compositeImage *Image, xOffset, yOffset T.Long) bool {
	return CompositeImage(i, compose, compositeImage, xOffset, yOffset)
}

func (i *Image) CompositeChannel(channel T.ChannelType, compose T.CompositeOperator, compositeImage *Image, xOffset, yOffset T.Long) bool {
	return CompositeImageChannel(i, channel, compose, compositeImage, xOffset, yOffset)
}

func (i *Image) CompositeLayers(compose T.CompositeOperator, source *Image, xOffset, yOffset T.Long, exception *ExceptionInfo) {
	CompositeLayers(i, compose, source, xOffset, yOffset, exception)
}

func (i *Image) CompressColormap() { CompressImageColormap(i) }

func (i *Image) Contrast(sharpen bool) bool { return ContrastImage(i, sharpen) }

func (i *Image) ContrastStretch(levels string) bool { return ContrastStretchImage(i, levels) }

func (i *Image) ContrastStretchChannel(channel T.ChannelType, blackPoint, whitePoint float64) bool {
	return ContrastStretchImageChannel(i, channel, blackPoint, whitePoint)
}

func (i *Image) Convolve(order T.Size, kernel *float64, exception *ExceptionInfo) *Image {
	return ConvolveImage(i, order, kernel, exception)
}

func (i *Image) ConvolveChannel(channel T.ChannelType, order T.Size, kernel *float64, exception *ExceptionInfo) *Image {
	return ConvolveImageChannel(i, channel, order, kernel, exception)
}

func (i *Image) CropImage(geometry *T.RectangleInfo, exception *ExceptionInfo) *Image {
	return CropImage(i, geometry, exception)
}

func (i *Image) CycleColormap(displace T.Long) bool { return CycleColormapImage(i, displace) }

func (i *Image) DefineArtifact(artifact string) bool { return DefineImageArtifact(i, artifact) }

func (i *Image) DefineProperty(property string) bool { return DefineImageProperty(i, property) }

func (i *Image) DeleteArtifact(artifact string) bool { return DeleteImageArtifact(i, artifact) }

func (i *Image) DeleteAttribute(key string) bool { return DeleteImageAttribute(i, key) }

func (i *Image) DeleteProfile(name string) bool { return DeleteImageProfile(i, name) }

func (i *Image) DeleteProperty(property string) bool { return DeleteImageProperty(i, property) }

func (i *Image) Describe(file *FILE, verbose bool) bool {
	return DescribeImage(i, file, verbose)
}

func (i *Image) Despeckle(exception *ExceptionInfo) *Image { return DespeckleImage(i, exception) }

func (i *Image) DestroyBlob() { DestroyBlob(i) }

func (i *Image) Destroy() *Image { return DestroyImage(i) }

func (i *Image) DestroyArtifacts() { DestroyImageArtifacts(i) }

func (i *Image) DestroyAttributes() { DestroyImageAttributes(i) }

func (i *Image) DestroyPixels() { DestroyImagePixels(i) }

func (i *Image) DestroyProfiles() { DestroyImageProfiles(i) }

func (i *Image) DestroyProperties() { DestroyImageProperties(i) }

func (i *Image) DestroyImages() { DestroyImages(i) }

func (i *Image) DisassociateStream() { DisassociateImageStream(i) }

func (i *Image) Dispatch(xOffset, yOffset T.Long, columns, rows T.Size, map_ string, type_ T.StorageType, pixels *T.Void, exception *ExceptionInfo) uint {
	return DispatchImage(i, xOffset, yOffset, columns, rows, map_, type_, pixels, exception)
}

func (i *Image) Dispose(exception *ExceptionInfo) *Image { return DisposeImages(i, exception) }

func (i *Image) Distort(method T.DistortImageMethod, numberArguments T.Size, arguments *float64, bestfit bool, exception *ExceptionInfo) *Image {
	return DistortImage(i, method, numberArguments, arguments, bestfit, exception)
}

func (i *Image) DrawAffine(source *Image, affine *T.AffineMatrix) bool {
	return DrawAffineImage(i, source, affine)
}

func (i *Image) DrawClipPath(drawInfo *DrawInfo, name string) bool {
	return DrawClipPath(i, drawInfo, name)
}

func (i *Image) DrawGradient(drawInfo *DrawInfo) bool { return DrawGradientImage(i, drawInfo) }

func (i *Image) Draw(drawInfo *DrawInfo) bool { return DrawImage(i, drawInfo) }

func (i *Image) DrawPatternPath(drawInfo *DrawInfo, name string, pattern **Image) bool {
	return DrawPatternPath(i, drawInfo, name, pattern)
}

func (i *Image) DrawPrimitive(drawInfo *DrawInfo, primitiveInfo *T.PrimitiveInfo) bool {
	return DrawPrimitive(i, drawInfo, primitiveInfo)
}

func (i *Image) Edge(radius float64, exception *ExceptionInfo) *Image {
	return EdgeImage(i, radius, exception)
}

func (i *Image) Emboss(radius, sigma float64, exception *ExceptionInfo) *Image {
	return EmbossImage(i, radius, sigma, exception)
}

func (i *Image) Enhance(exception *ExceptionInfo) *Image { return EnhanceImage(i, exception) }

func (i *Image) EOFBlob() int { return EOFBlob(i) }

func (i *Image) Equalize() bool { return EqualizeImage(i) }

func (i *Image) EqualizeChannel(channel T.ChannelType) bool {
	return EqualizeImageChannel(i, channel)
}

func (i *Image) Evaluate(op T.MagickEvaluateOperator, value float64, exception *ExceptionInfo) bool {
	return EvaluateImage(i, op, value, exception)
}

func (i *Image) EvaluateChannel(channel T.ChannelType, op T.MagickEvaluateOperator, value float64, exception *ExceptionInfo) bool {
	return EvaluateImageChannel(i, channel, op, value, exception)
}

func (i *Image) Excerpt(geometry *T.RectangleInfo, exception *ExceptionInfo) *Image {
	return ExcerptImage(i, geometry, exception)
}

func (i *Image) ExportPixels(xOffset, yOffset T.Long, columns, rows T.Size, map_ string, type_ T.StorageType, pixels *T.Void, exception *ExceptionInfo) bool {
	return ExportImagePixels(i, xOffset, yOffset, columns, rows, map_, type_, pixels, exception)
}

func (i *Image) ExportQuantumPixels(quantumInfo *T.QuantumInfo, quantumType T.QuantumType, pixels *byte) bool {
	return ExportQuantumPixels(i, quantumInfo, quantumType, pixels)
}

func (i *Image) Extent(geometry *T.RectangleInfo, exception *ExceptionInfo) *Image {
	return ExtentImage(i, geometry, exception)
}

func (i *Image) FromFile(filename string) bool { return FileToImage(i, filename) }

func (i *Image) FlattenImages(exception *ExceptionInfo) *Image { return FlattenImages(i, exception) }

func (i *Image) Flip(exception *ExceptionInfo) *Image { return FlipImage(i, exception) }

func (i *Image) FloodfillPaint(channel T.ChannelType, drawInfo *DrawInfo, target *MagickPixelPacket, xOffset, yOffset T.Long, invert bool) bool {
	return FloodfillPaintImage(i, channel, drawInfo, target, xOffset, yOffset, invert)
}

func (i *Image) Flop(exception *ExceptionInfo) *Image { return FlopImage(i, exception) }

func (i *Image) FormatAttribute(key, format string, va ...VArg) bool {
	return FormatImageAttribute(i, key, format, va)
}

func (i *Image) FormatAttributeList(key, format string, operands VAList) bool {
	return FormatImageAttributeList(i, key, format, operands)
}

func (i *Image) FormatImageProperty(property, format string, va ...VArg) bool {
	return FormatImageProperty(i, property, format, va)
}

func (i *Image) FormatPropertyList(property, format string, operands VAList) bool {
	return FormatImagePropertyList(i, property, format, operands)
}

func (i *Image) FormatMagickCaption(drawInfo *DrawInfo, caption string, metrics *T.TypeMetric) T.Long {
	return FormatMagickCaption(i, drawInfo, caption, metrics)
}

func (i *Image) Frame(frameInfo *T.FrameInfo, exception *ExceptionInfo) *Image {
	return FrameImage(i, frameInfo, exception)
}

func (i *Image) FuzzyColorCompare(p, q *T.PixelPacket) bool { return FuzzyColorCompare(i, p, q) }

func (i *Image) FuzzyOpacityCompare(p, q *T.PixelPacket) bool { return FuzzyOpacityCompare(i, p, q) }

func (i *Image) Fx(expression string, exception *ExceptionInfo) *Image {
	return FxImage(i, expression, exception)
}

func (i *Image) FxChannel(channel T.ChannelType, expression string, exception *ExceptionInfo) *Image {
	return FxImageChannel(i, channel, expression, exception)
}

func (i *Image) Gamma(level string) bool { return GammaImage(i, level) }

func (i *Image) GammaChannel(channel T.ChannelType, gamma float64) bool {
	return GammaImageChannel(i, channel, gamma)
}

func (i *Image) GaussianBlur(radius, sigma float64, exception *ExceptionInfo) *Image {
	return GaussianBlurImage(i, radius, sigma, exception)
}

func (i *Image) GaussianBlurChannel(channel T.ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return GaussianBlurImageChannel(i, channel, radius, sigma, exception)
}

func (i *Image) BlobError() bool { return GetBlobError(i) }

func (i *Image) BlobFileHandle() *FILE { return GetBlobFileHandle(i) }

func (i *Image) BlobSize() T.MagickSizeType { return GetBlobSize(i) }

func (i *Image) BlobStreamData() *byte { return GetBlobStreamData(i) }

func (i *Image) BlobStreamHandler() T.StreamHandler { return GetBlobStreamHandler(i) }

func (i *Image) Artifact(artifact string) string { return GetImageArtifact(i, artifact) }

func (i *Image) Attribute(key string) *T.ImageAttribute { return GetImageAttribute(i, key) }

func (i *Image) BoundingBox(exception *ExceptionInfo) T.RectangleInfo {
	return GetImageBoundingBox(i, exception)
}

func (i *Image) ChannelDepth(channel T.ChannelType, exception *ExceptionInfo) T.Size {
	return GetImageChannelDepth(i, channel, exception)
}

func (i *Image) ChannelDistortion(reconstructImage *Image, channel T.ChannelType, metric T.MetricType, distortion *float64, exception *ExceptionInfo) bool {
	return GetImageChannelDistortion(i, reconstructImage, channel, metric, distortion, exception)
}

func (i *Image) ChannelExtrema(channel T.ChannelType, minima, maxima *T.Size, exception *ExceptionInfo) bool {
	return GetImageChannelExtrema(i, channel, minima, maxima, exception)
}

func (i *Image) ChannelMean(channel T.ChannelType, mean, standardDeviation *float64, exception *ExceptionInfo) bool {
	return GetImageChannelMean(i, channel, mean, standardDeviation, exception)
}

func (i *Image) ChannelRange(channel T.ChannelType, minima, maxima *float64, exception *ExceptionInfo) bool {
	return GetImageChannelRange(i, channel, minima, maxima, exception)
}

func (i *Image) ChannelStatistics(exception *ExceptionInfo) *T.ChannelStatistics {
	return GetImageChannelStatistics(i, exception)
}

func (i *Image) ClipMask(exception *ExceptionInfo) *Image {
	return GetImageClipMask(i, exception)
}

func (i *Image) ClippingPathAttribute() *T.ImageAttribute {
	return GetImageClippingPathAttribute(i)
}

func (i *Image) Depth(exception *ExceptionInfo) T.Size {
	return GetImageDepth(i, exception)
}

func (i *Image) Distortion(reconstructImage *Image, metric T.MetricType, distortion *float64, exception *ExceptionInfo) bool {
	return GetImageDistortion(i, reconstructImage, metric, distortion, exception)
}

func (i *Image) DynamicThreshold(clusterThreshold, smoothThreshold float64, exception *ExceptionInfo) MagickPixelPacket {
	return GetImageDynamicThreshold(i, clusterThreshold, smoothThreshold, exception)
}

func (i *Image) Exception(exception *ExceptionInfo) { GetImageException(i, exception) }

func (i *Image) Extrema(minima, maxima *T.Size, exception *ExceptionInfo) bool {
	return GetImageExtrema(i, minima, maxima, exception)
}

func (i *Image) Geometry(geometry string, sizeToFit uint, regionInfo *T.RectangleInfo) int {
	return GetImageGeometry(i, geometry, sizeToFit, regionInfo)
}

func (i *Image) Histogram(numberColors *T.Size, exception *ExceptionInfo) *T.ColorPacket {
	return GetImageHistogram(i, numberColors, exception)
}

func (i *Image) Mask(exception *ExceptionInfo) *Image { return GetImageMask(i, exception) }

func (i *Image) Mean(mean, standardDeviation *float64, exception *ExceptionInfo) bool {
	return GetImageMean(i, mean, standardDeviation, exception)
}

func (i *Image) ImagePixels(x, y T.Long, columns, rows T.Size) *T.PixelPacket {
	return GetImagePixels(i, x, y, columns, rows)
}

func (i *Image) Profile(name string) *StringInfo { return GetImageProfile(i, name) }

func (i *Image) Property(property string) string { return GetImageProperty(i, property) }

func (i *Image) QuantizeError() bool { return GetImageQuantizeError(i) }

func (i *Image) QuantumDepth(constrain bool) T.Size {
	return GetImageQuantumDepth(i, constrain)
}

func (i *Image) Range(minima, maxima *float64, exception *ExceptionInfo) bool {
	return GetImageRange(i, minima, maxima, exception)
}

func (i *Image) TotalInkDensity() float64 { return GetImageTotalInkDensity(i) }

func (i *Image) Type(exception *ExceptionInfo) T.ImageType { return GetImageType(i, exception) }

func (i *Image) VirtualPixelMethod() T.VirtualPixelMethod { return GetImageVirtualPixelMethod(i) }

func (i *Image) Indexes() *T.IndexPacket { return GetIndexes(i) }

func (i *Image) MagickPixelPacket(pixel *MagickPixelPacket) { GetMagickPixelPacket(i, pixel) }

func (i *Image) MultilineTypeMetrics(drawInfo *DrawInfo, metrics *T.TypeMetric) bool {
	return GetMultilineTypeMetrics(i, drawInfo, metrics)
}

func (i *Image) NextArtifact() string { return GetNextImageArtifact(i) }

func (i *Image) NextAttribute() *T.ImageAttribute { return GetNextImageAttribute(i) }

func (i *Image) NextProfile() string { return GetNextImageProfile(i) }

func (i *Image) NextProperty() string { return GetNextImageProperty(i) }

func (i *Image) NumberColors(file *FILE, exception *ExceptionInfo) T.Size {
	return GetNumberColors(i, file, exception)
}

func (i *Image) NumberScenes() uint { return GetNumberScenes(i) }

func (i *Image) OnePixel(x, y T.Long) T.PixelPacket { return GetOnePixel(i, x, y) }

func (i *Image) PixelCacheArea() T.MagickSizeType { return GetPixelCacheArea(i) }

func (i *Image) Pixels() *T.PixelPacket { return GetPixels(i) }

func (i *Image) TypeMetrics(drawInfo *DrawInfo, metrics *T.TypeMetric) bool {
	return GetTypeMetrics(i, drawInfo, metrics)
}

func (i *Image) Gradient(startColor, stopColor *T.PixelPacket) bool {
	return GradientImage(i, startColor, stopColor)
}

func (i *Image) HuffmanDecode() bool { return HuffmanDecodeImage(i) }

func (i *Image) Identify(file *FILE, verbose bool) bool { return IdentifyImage(i, file, verbose) }

func (i *Image) ToFile(filename string, exception *ExceptionInfo) bool {
	return ImageToFile(i, filename, exception)
}

func (i *Image) Implode(amount float64, exception *ExceptionInfo) *Image {
	return ImplodeImage(i, amount, exception)
}

func (i *Image) ImportPixels(xOffset, yOffset T.Long, columns, rows T.Size, map_ string, type_ T.StorageType, pixels *T.Void) bool {
	return ImportImagePixels(i, xOffset, yOffset, columns, rows, map_, type_, pixels)
}

func (i *Image) ImportQuantumPixels(quantumInfo *T.QuantumInfo, quantumType T.QuantumType, pixels *byte) bool {
	return ImportQuantumPixels(i, quantumInfo, quantumType, pixels)
}

func (i *Image) InterpolatePixelColor(imageView *CacheView, method T.InterpolatePixelMethod, x, y float64, exception *ExceptionInfo) MagickPixelPacket {
	return InterpolatePixelColor(i, imageView, method, x, y, exception)
}

func (i *Image) BlobExempt() bool { return IsBlobExempt(i) }

func (i *Image) BlobSeekable() bool { return IsBlobSeekable(i) }

func (i *Image) BlobTemporary() bool { return IsBlobTemporary(i) }

func (i *Image) ColorSimilar(p, q *T.PixelPacket) bool { return IsColorSimilar(i, p, q) }

func (i *Image) Gray(exception *ExceptionInfo) bool { return IsGrayImage(i, exception) }

func (i *Image) HighDynamicRange(exception *ExceptionInfo) bool {
	return IsHighDynamicRangeImage(i, exception)
}

func (i *Image) IsHistogram(exception *ExceptionInfo) bool {
	return IsHistogramImage(i, exception)
}

func (i *Image) Object() bool { return IsImageObject(i) }

func (i *Image) Equal(reconstructImage *Image) bool {
	return IsImagesEqual(i, reconstructImage)
}

func (i *Image) Similar(targetImage *Image, xOffset, yOffset *T.Long, exception *ExceptionInfo) bool {
	return IsImageSimilar(i, targetImage, xOffset, yOffset, exception)
}

func (i *Image) Monochrome(exception *ExceptionInfo) bool {
	return IsMonochromeImage(i, exception)
}

func (i *Image) OpacitySimilar(p, q *T.PixelPacket) bool { return IsOpacitySimilar(i, p, q) }

func (i *Image) IsOpaque(exception *ExceptionInfo) bool { return IsOpaqueImage(i, exception) }

func (i *Image) Palette(exception *ExceptionInfo) bool { return IsPaletteImage(i, exception) }

func (i *Image) Taint() bool { return IsTaintImage(i) }

func (i *Image) Level(levels string) bool { return LevelImage(i, levels) }

func (i *Image) LevelChannel(channel T.ChannelType, blackPoint, whitePoint, gamma float64) bool {
	return LevelImageChannel(i, channel, blackPoint, whitePoint, gamma)
}

func (i *Image) LinearStretch(blackPoint, whitePoint float64) bool {
	return LinearStretchImage(i, blackPoint, whitePoint)
}

func (i *Image) LiquidRescale(columns, rows T.Size, deltaX, rigidity float64, exception *ExceptionInfo) *Image {
	return LiquidRescaleImage(i, columns, rows, deltaX, rigidity, exception)
}

func (i *Image) LZWEncode(length uint32, pixels *byte) bool {
	return LZWEncodeImage(i, length, pixels)
}

func (i *Image) Magnify(exception *ExceptionInfo) *Image { return MagnifyImage(i, exception) }

func (i *Image) Map(mapImage *Image, dither bool) bool { return MapImage(i, mapImage, dither) }

func (i *Image) MatteFloodfill(target T.PixelPacket, opacity T.Quantum, xOffset, yOffset T.Long, method T.PaintMethod) bool {
	return MatteFloodfillImage(i, target, opacity, xOffset, yOffset, method)
}

func (i *Image) MedianFilter(radius float64, exception *ExceptionInfo) *Image {
	return MedianFilterImage(i, radius, exception)
}

func (i *Image) MergeImageLayers(method T.ImageLayerMethod, exception *ExceptionInfo) *Image {
	return MergeImageLayers(i, method, exception)
}

func (i *Image) Minify(exception *ExceptionInfo) *Image { return MinifyImage(i, exception) }

func (i *Image) Modulate(modulate string) bool { return ModulateImage(i, modulate) }

func (i *Image) MorphImages(numberFrames T.Size, exception *ExceptionInfo) *Image {
	return MorphImages(i, numberFrames, exception)
}

func (i *Image) MosaicImages(exception *ExceptionInfo) *Image { return MosaicImages(i, exception) }

func (i *Image) MotionBlur(radius, sigma, angle float64, exception *ExceptionInfo) *Image {
	return MotionBlurImage(i, radius, sigma, angle, exception)
}

func (i *Image) Negate(grayscale bool) bool { return NegateImage(i, grayscale) }

func (i *Image) NegateChannel(channel T.ChannelType, grayscale bool) bool {
	return NegateImageChannel(i, channel, grayscale)
}

func (i *Image) Normalize() bool { return NormalizeImage(i) }

func (i *Image) NormalizeChannel(channel T.ChannelType) bool {
	return NormalizeImageChannel(i, channel)
}

func (i *Image) OilPaint(radius float64, exception *ExceptionInfo) *Image {
	return OilPaintImage(i, radius, exception)
}

func (i *Image) Opaque(target, fill T.PixelPacket) bool { return OpaqueImage(i, target, fill) }

func (i *Image) OpaquePaint(target, fill *MagickPixelPacket, invert bool) bool {
	return OpaquePaintImage(i, target, fill, invert)
}

func (i *Image) OpaquePaintChannel(channel T.ChannelType, target, fill *MagickPixelPacket, invert bool) bool {
	return OpaquePaintImageChannel(i, channel, target, fill, invert)
}

func (i *Image) OpenCacheView() *CacheView { return OpenCacheView(i) }

func (i *Image) OptimizeLayers(exception *ExceptionInfo) *Image {
	return OptimizeImageLayers(i, exception)
}

func (i *Image) OptimizeTransparency(exception *ExceptionInfo) {
	OptimizeImageTransparency(i, exception)
}

func (i *Image) OptimizePlusLayers(exception *ExceptionInfo) *Image {
	return OptimizePlusImageLayers(i, exception)
}

func (i *Image) OrderedDither() bool { return OrderedDitherImage(i) }

func (i *Image) OrderedDitherChannel(channel T.ChannelType, exception *ExceptionInfo) bool {
	return OrderedDitherImageChannel(i, channel, exception)
}

func (i *Image) OrderedPosterize(thresholdMap string, exception *ExceptionInfo) bool {
	return OrderedPosterizeImage(i, thresholdMap, exception)
}

func (i *Image) OrderedPosterizeChannel(channel T.ChannelType, thresholdMap string, exception *ExceptionInfo) bool {
	return OrderedPosterizeImageChannel(i, channel, thresholdMap, exception)
}

func (i *Image) PackbitsEncode(length uint32, pixels *byte) bool {
	return PackbitsEncodeImage(i, length, pixels)
}

func (i *Image) PaintFloodfill(channel T.ChannelType, target *MagickPixelPacket, x, y T.Long, drawInfo *DrawInfo, method T.PaintMethod) bool {
	return PaintFloodfillImage(i, channel, target, x, y, drawInfo, method)
}

func (i *Image) PaintOpaque(target, fill *MagickPixelPacket) bool {
	return PaintOpaqueImage(i, target, fill)
}

func (i *Image) PaintOpaqueChannel(channel T.ChannelType, target, fill *MagickPixelPacket) bool {
	return PaintOpaqueImageChannel(i, channel, target, fill)
}

func (i *Image) PaintTransparentImage(target *MagickPixelPacket, opacity T.Quantum) bool {
	return PaintTransparentImage(i, target, opacity)
}

func (i *Image) ParseGravityGeometry(geometry string, regionInfo *T.RectangleInfo) T.MagickStatusType {
	return ParseGravityGeometry(i, geometry, regionInfo)
}

func (i *Image) ParsePageGeometry(geometry string, regionInfo *T.RectangleInfo) T.MagickStatusType {
	return ParsePageGeometry(i, geometry, regionInfo)
}

func (i *Image) PersistCache(filename string, attach bool, offset *T.MagickOffsetType, exception *ExceptionInfo) bool {
	return PersistCache(i, filename, attach, offset, exception)
}

func (i *Image) Plasma(segment *T.SegmentInfo, attenuate, depth T.Size) bool {
	return PlasmaImage(i, segment, attenuate, depth)
}

func (i *Image) Polaroid(drawInfo *DrawInfo, angle float64, exception *ExceptionInfo) *Image {
	return PolaroidImage(i, drawInfo, angle, exception)
}

func (i *Image) PopPixels(quantum T.QuantumType, destination *byte) bool {
	return PopImagePixels(i, quantum, destination)
}

func (i *Image) Posterize(levels T.Size, dither bool) bool {
	return PosterizeImage(i, levels, dither)
}

func (i *Image) Preview(preview T.PreviewType, exception *ExceptionInfo) *Image {
	return PreviewImage(i, preview, exception)
}

func (i *Image) ProfileImage(name string, datum *T.Void, length uint32, clone bool) bool {
	return ProfileImage(i, name, datum, length, clone)
}

func (i *Image) PushPixels(quantum T.QuantumType, source *byte) bool {
	return PushImagePixels(i, quantum, source)
}

func (i *Image) QuantizationError() uint { return QuantizationError(i) }

func (i *Image) QueryColorname(color *T.PixelPacket, compliance T.ComplianceType, name string, exception *ExceptionInfo) bool {
	return QueryColorname(i, color, compliance, name, exception)
}

func (i *Image) QueryMagickColorname(color *MagickPixelPacket, compliance T.ComplianceType, hex bool, name string, exception *ExceptionInfo) bool {
	return QueryMagickColorname(i, color, compliance, hex, name, exception)
}

func (i *Image) RadialBlur(angle float64, exception *ExceptionInfo) *Image {
	return RadialBlurImage(i, angle, exception)
}

func (i *Image) RadialBlurChannel(channel T.ChannelType, angle float64, exception *ExceptionInfo) *Image {
	return RadialBlurImageChannel(i, channel, angle, exception)
}

func (i *Image) RaiseImage(raiseInfo *T.RectangleInfo, raise bool) bool {
	return RaiseImage(i, raiseInfo, raise)
}

func (i *Image) RandomChannelThresholdImage(channel, thresholds string, exception *ExceptionInfo) uint {
	return RandomChannelThresholdImage(i, channel, thresholds, exception)
}

func (i *Image) RandomThresholdImage(thresholds string, exception *ExceptionInfo) bool {
	return RandomThresholdImage(i, thresholds, exception)
}

func (i *Image) RandomThresholdImageChannel(channel T.ChannelType, thresholds string, exception *ExceptionInfo) bool {
	return RandomThresholdImageChannel(i, channel, thresholds, exception)
}

func (i *Image) ReadBlob(length uint32, data *byte) int32 { return ReadBlob(i, length, data) }

func (i *Image) ReadBlobByte() int { return ReadBlobByte(i) }

func (i *Image) ReadBlobDouble() float64 { return ReadBlobDouble(i) }

func (i *Image) ReadBlobFloat() float32 { return ReadBlobFloat(i) }

func (i *Image) ReadBlobLong() T.Size { return ReadBlobLong(i) }

func (i *Image) ReadBlobLongLong() T.MagickSizeType { return ReadBlobLongLong(i) }

func (i *Image) ReadBlobLSBLong() T.Size { return ReadBlobLSBLong(i) }

func (i *Image) ReadBlobLSBShort() T.UnsignedShort { return ReadBlobLSBShort(i) }

func (i *Image) ReadBlobMSBLong() T.Size { return ReadBlobMSBLong(i) }

func (i *Image) ReadBlobMSBShort() T.UnsignedShort { return ReadBlobMSBShort(i) }

func (i *Image) ReadBlobShort() T.UnsignedShort { return ReadBlobShort(i) }

func (i *Image) ReadBlobString(str string) string { return ReadBlobString(i, str) }

func (i *Image) Recolor(order T.Size, colorMatrix *float64, exception *ExceptionInfo) *Image {
	return RecolorImage(i, order, colorMatrix, exception)
}

func (i *Image) ReduceNoise(radius float64, exception *ExceptionInfo) *Image {
	return ReduceNoiseImage(i, radius, exception)
}

func (i *Image) Reference() *Image { return ReferenceImage(i) }

func (i *Image) RemoveArtifact(artifact string) string { return RemoveImageArtifact(i, artifact) }

func (i *Image) RemoveProfile(name string) *StringInfo { return RemoveImageProfile(i, name) }

func (i *Image) RemoveProperty(property string) string { return RemoveImageProperty(i, property) }

func (i *Image) Resample(xResolution, yResolution float64, filter T.FilterTypes, blur float64, exception *ExceptionInfo) *Image {
	return ResampleImage(i, xResolution, yResolution, filter, blur, exception)
}

func (i *Image) ResetArtifactIterator() { ResetImageArtifactIterator(i) }

func (i *Image) ResetAttributeIterator() { ResetImageAttributeIterator(i) }

func (i *Image) ResetPage(page string) bool { return ResetImagePage(i, page) }

func (i *Image) ResetProfileIterator() { ResetImageProfileIterator(i) }

func (i *Image) ResetPropertyIterator() { ResetImagePropertyIterator(i) }

func (i *Image) Resize(columns, rows T.Size, filter T.FilterTypes, blur float64, exception *ExceptionInfo) *Image {
	return ResizeImage(i, columns, rows, filter, blur, exception)
}

func (i *Image) RGBTransform(colorspace T.ColorspaceType) bool {
	return RGBTransformImage(i, colorspace)
}

func (i *Image) Roll(xOffset, yOffset T.Long, exception *ExceptionInfo) *Image {
	return RollImage(i, xOffset, yOffset, exception)
}

func (i *Image) Rotate(degrees float64, exception *ExceptionInfo) *Image {
	return RotateImage(i, degrees, exception)
}

func (i *Image) Sample(columns, rows T.Size, exception *ExceptionInfo) *Image {
	return SampleImage(i, columns, rows, exception)
}

func (i *Image) Scale(columns, rows T.Size, exception *ExceptionInfo) *Image {
	return ScaleImage(i, columns, rows, exception)
}

func (i *Image) SeekBlob(offset T.MagickOffsetType, whence int) T.MagickOffsetType {
	return SeekBlob(i, offset, whence)
}

func (i *Image) Segment(colorspace T.ColorspaceType, verbose bool, clusterThreshold, smoothThreshold float64) bool {
	return SegmentImage(i, colorspace, verbose, clusterThreshold, smoothThreshold)
}

func (i *Image) SeparateChannel(channel T.ChannelType) bool {
	return SeparateImageChannel(i, channel)
}

func (i *Image) SeparateImages(channel T.ChannelType, exception *ExceptionInfo) *Image {
	return SeparateImages(i, channel, exception)
}

func (i *Image) SepiaTone(threshold float64, exception *ExceptionInfo) *Image {
	return SepiaToneImage(i, threshold, exception)
}

func (i *Image) SetBlobExempt(exempt bool) { SetBlobExempt(i, exempt) }

func (i *Image) SetGeometry(geometry *T.RectangleInfo) { SetGeometry(i, geometry) }

func (i *Image) SetImage(opacity T.Quantum) { SetImage(i, opacity) }

func (i *Image) SetAlphaChannel(alphaType T.AlphaChannelType) bool {
	return SetImageAlphaChannel(i, alphaType)
}

func (i *Image) SetArtifact(artifact, value string) bool {
	return SetImageArtifact(i, artifact, value)
}

func (i *Image) SetAttribute(key, value string) bool { return SetImageAttribute(i, key, value) }

func (i *Image) SetBackgroundColor() bool { return SetImageBackgroundColor(i) }

func (i *Image) SetChannelDepth(channel T.ChannelType, depth T.Size) bool {
	return SetImageChannelDepth(i, channel, depth)
}

func (i *Image) SetClipMask(clipMask *Image) bool { return SetImageClipMask(i, clipMask) }

func (i *Image) SetDepth(depth T.Size) bool { return SetImageDepth(i, depth) }

func (i *Image) SetExtent(columns, rows T.Size) bool {
	return SetImageExtent(i, columns, rows)
}

func (i *Image) SetMask(mask *Image) bool { return SetImageMask(i, mask) }

func (i *Image) SetOpacity(opacity T.Quantum) bool { return SetImageOpacity(i, opacity) }

func (i *Image) SetPixels(x, y T.Long, columns, rows T.Size) *T.PixelPacket {
	return SetImagePixels(i, x, y, columns, rows)
}

func (i *Image) SetProfile(name string, profile *StringInfo) bool {
	return SetImageProfile(i, name, profile)
}

func (i *Image) SetProgressMonitor(progressMonitor T.MagickProgressMonitor, clientData *T.Void) T.MagickProgressMonitor {
	return SetImageProgressMonitor(i, progressMonitor, clientData)
}

func (i *Image) SetProperty(property, value string) bool {
	return SetImageProperty(i, property, value)
}

func (i *Image) SetStorageClass(storageClass T.ClassType) bool {
	return SetImageStorageClass(i, storageClass)
}

func (i *Image) SetType(imageType T.ImageType) bool { return SetImageType(i, imageType) }

func (i *Image) SetVirtualPixelMethod(virtualPixelMethod T.VirtualPixelMethod) T.VirtualPixelMethod {
	return SetImageVirtualPixelMethod(i, virtualPixelMethod)
}

func (i *Image) Shade(gray bool, azimuth, elevation float64, exception *ExceptionInfo) *Image {
	return ShadeImage(i, gray, azimuth, elevation, exception)
}

func (i *Image) Shadow(opacity, sigma float64, xOffset, yOffset T.Long, exception *ExceptionInfo) *Image {
	return ShadowImage(i, opacity, sigma, xOffset, yOffset, exception)
}

func (i *Image) Sharpen(radius, sigma float64, exception *ExceptionInfo) *Image {
	return SharpenImage(i, radius, sigma, exception)
}

func (i *Image) SharpenChannel(channel T.ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return SharpenImageChannel(i, channel, radius, sigma, exception)
}

func (i *Image) Shave(shaveInfo *T.RectangleInfo, exception *ExceptionInfo) *Image {
	return ShaveImage(i, shaveInfo, exception)
}

func (i *Image) Shear(xShear, yShear float64, exception *ExceptionInfo) *Image {
	return ShearImage(i, xShear, yShear, exception)
}

func (i *Image) SigmoidalContrast(sharpen bool, levels string) bool {
	return SigmoidalContrastImage(i, sharpen, levels)
}

func (i *Image) SigmoidalContrastChannel(channel T.ChannelType, sharpen bool, contrast, midpoint float64) bool {
	return SigmoidalContrastImageChannel(i, channel, sharpen, contrast, midpoint)
}

func (i *Image) Signature() bool { return SignatureImage(i) }

func (i *Image) SizeBlob() T.MagickOffsetType { return SizeBlob(i) }

func (i *Image) Sketch(radius, sigma, angle float64, exception *ExceptionInfo) *Image {
	return SketchImage(i, radius, sigma, angle, exception)
}

func (i *Image) Solarize(threshold float64) bool { return SolarizeImage(i, threshold) }

func (i *Image) SortColormapByIntensity() bool { return SortColormapByIntensity(i) }

func (i *Image) Splice(geometry *T.RectangleInfo, exception *ExceptionInfo) *Image {
	return SpliceImage(i, geometry, exception)
}

func (i *Image) Spread(radius float64, exception *ExceptionInfo) *Image {
	return SpreadImage(i, radius, exception)
}

func (i *Image) Stegano(watermark *Image, exception *ExceptionInfo) *Image {
	return SteganoImage(i, watermark, exception)
}

func (i *Image) Stereo(offsetImage *Image, exception *ExceptionInfo) *Image {
	return StereoImage(i, offsetImage, exception)
}

func (i *Image) Strip() bool { return StripImage(i) }

func (i *Image) Swirl(degrees float64, exception *ExceptionInfo) *Image {
	return SwirlImage(i, degrees, exception)
}

func (i *Image) Sync() bool { return SyncImage(i) }

func (i *Image) SyncPixels() bool { return SyncImagePixels(i) }

func (i *Image) SyncProfiles() bool { return SyncImageProfiles(i) }

func (i *Image) TellBlob() T.MagickOffsetType { return TellBlob(i) }

func (i *Image) Texture(texture *Image) bool { return TextureImage(i, texture) }

func (i *Image) Threshold(threshold float64) uint { return ThresholdImage(i, threshold) }

func (i *Image) ThresholdChannel(threshold string) uint {
	return ThresholdImageChannel(i, threshold)
}

func (i *Image) Thumbnail(columns, rows T.Size, exception *ExceptionInfo) *Image {
	return ThumbnailImage(i, columns, rows, exception)
}

func (i *Image) Tint(opacity string, tint T.PixelPacket, exception *ExceptionInfo) *Image {
	return TintImage(i, opacity, tint, exception)
}

func (i *Image) TransformColorspace(colorspace T.ColorspaceType) uint {
	return TransformColorspace(i, colorspace)
}

func (i *Image) TransformImageColorspace(colorspace T.ColorspaceType) bool {
	return TransformImageColorspace(i, colorspace)
}

func (i *Image) TransformRGB(colorspace T.ColorspaceType) bool {
	return TransformRGBImage(i, colorspace)
}

func (i *Image) Transparent(target T.PixelPacket, opacity T.Quantum) bool {
	return TransparentImage(i, target, opacity)
}

func (i *Image) TransparentPaint(target *MagickPixelPacket, opacity T.Quantum, invert bool) bool {
	return TransparentPaintImage(i, target, opacity, invert)
}

func (i *Image) Transpose(exception *ExceptionInfo) *Image { return TransposeImage(i, exception) }

func (i *Image) Transverse(exception *ExceptionInfo) *Image {
	return TransverseImage(i, exception)
}

func (i *Image) Trim(exception *ExceptionInfo) *Image { return TrimImage(i, exception) }

func (i *Image) UniqueColors(exception *ExceptionInfo) *Image {
	return UniqueImageColors(i, exception)
}

func (i *Image) UnsharpMask(radius, sigma, amount, threshold float64, exception *ExceptionInfo) *Image {
	return UnsharpMaskImage(i, radius, sigma, amount, threshold, exception)
}

func (i *Image) UnsharpMaskChannel(channel T.ChannelType, radius, sigma, amount, threshold float64, exception *ExceptionInfo) *Image {
	return UnsharpMaskImageChannel(i, channel, radius, sigma, amount, threshold, exception)
}

func (i *Image) ValidateColormapIndex(index T.Size) T.IndexPacket {
	return ValidateColormapIndex(i, index)
}

func (i *Image) Vignette(radius, sigma float64, x, y T.Long, exception *ExceptionInfo) *Image {
	return VignetteImage(i, radius, sigma, x, y, exception)
}

func (i *Image) Wave(amplitude, waveLength float64, exception *ExceptionInfo) *Image {
	return WaveImage(i, amplitude, waveLength, exception)
}

func (i *Image) WhiteThreshold(threshold string) bool { return WhiteThresholdImage(i, threshold) }

func (i *Image) WriteBlob(length uint32, data *byte) int32 { return WriteBlob(i, length, data) }

func (i *Image) WriteBlobByte(value byte) int32 { return WriteBlobByte(i, value) }

func (i *Image) WriteBlobFloat(value float32) int32 { return WriteBlobFloat(i, value) }

func (i *Image) WriteBlobLong(value T.UnsignedLong) int32 { return WriteBlobLong(i, value) }

func (i *Image) WriteBlobLSBLong(value T.UnsignedLong) int32 { return WriteBlobLSBLong(i, value) }

func (i *Image) WriteBlobLSBShort(value T.UnsignedShort) int32 { return WriteBlobLSBShort(i, value) }

func (i *Image) WriteBlobMSBLong(value T.UnsignedLong) int32 { return WriteBlobMSBLong(i, value) }

func (i *Image) WriteBlobMSBShort(value T.UnsignedShort) int32 { return WriteBlobMSBShort(i, value) }

func (i *Image) WriteBlobShort(value T.UnsignedShort) int32 { return WriteBlobShort(i, value) }

func (i *Image) WriteBlobString(str string) int32 { return WriteBlobString(i, str) }

func (i *Image) ZLIBEncode(length uint32, pixels *byte) bool {
	return ZLIBEncodeImage(i, length, pixels)
}

func (i *Image) Zoom(columns, rows T.Size, exception *ExceptionInfo) *Image {
	return ZoomImage(i, columns, rows, exception)
}

func (i *Image) AcquireAuthenticCacheView(exception *ExceptionInfo) *CacheView {
	return AcquireAuthenticCacheView(i, exception)
}

func (i *Image) AcquireCacheView() *CacheView { return AcquireCacheView(i) }

func (i *Image) AcquireImageColormap(colors uint32) bool { return AcquireImageColormap(i, colors) }

func (i *Image) AcquirePixelCache(a1 T.VirtualPixelMethod, a2 T.Long, a3 T.Long, a4 T.Size, a5 T.Size, a6 *ExceptionInfo) *T.PixelPacket {
	return AcquirePixelCache(i, a1, a2, a3, a4, a5, a6)
}

func (i *Image) AcquirePixels() *T.PixelPacket { return AcquirePixels(i) }

func (i *Image) AcquireVirtualCacheView(exception *ExceptionInfo) *CacheView {
	return AcquireVirtualCacheView(i, exception)
}

func (i *Image) AutoGamma() bool { return AutoGammaImage(i) }

func (i *Image) AutoGammaChannel(channel T.ChannelType) bool {
	return AutoGammaImageChannel(i, channel)
}

func (i *Image) AutoLevel() bool { return AutoLevelImage(i) }

func (i *Image) AutoLevelChannel(channel T.ChannelType) bool {
	return AutoLevelImageChannel(i, channel)
}

func (i *Image) BlueShift(factor float64, exception *ExceptionInfo) *Image {
	return BlueShiftImage(i, factor, exception)
}

func (i *Image) BrightnessContrast(brightness, contrast float64) bool {
	return BrightnessContrastImage(i, brightness, contrast)
}

func (i *Image) BrightnessContrastChannel(channel T.ChannelType, brightness, contrast float64) bool {
	return BrightnessContrastImageChannel(i, channel, brightness, contrast)
}

func (i *Image) ColorDecisionList(colorCorrectionCollection string) bool {
	return ColorDecisionListImage(i, colorCorrectionCollection)
}

func (i *Image) ColorMatrix(colorMatrix *T.KernelInfo, exception *ExceptionInfo) *Image {
	return ColorMatrixImage(i, colorMatrix, exception)
}

func (i *Image) CropToHBITMAP(r *T.RectangleInfo, e *ExceptionInfo) *T.Void {
	return CropImageToHBITMAP(i, r, e)
}

func (i *Image) CropToTiles(cropGeometry *T.RectangleInfo, exception *ExceptionInfo) *Image {
	return CropImageToTiles(i, cropGeometry, exception)
}

func (i *Image) Decipher(passphrase string, exception *ExceptionInfo) bool {
	return DecipherImage(i, passphrase, exception)
}

func (i *Image) Deskew(threshold float64, exception *ExceptionInfo) *Image {
	return DeskewImage(i, threshold, exception)
}

func (i *Image) Encipher(passphrase string, exception *ExceptionInfo) bool {
	return EncipherImage(i, passphrase, exception)
}

func (i *Image) ExtractSubimage(reference *Image, exception *ExceptionInfo) *Image {
	return ExtractSubimageFromImage(i, reference, exception)
}

func (i *Image) Filter(kernel *T.KernelInfo, exception *ExceptionInfo) *Image {
	return FilterImage(i, kernel, exception)
}

func (i *Image) FilterChannel(channel T.ChannelType, kernel *T.KernelInfo, exception *ExceptionInfo) *Image {
	return FilterImageChannel(i, channel, kernel, exception)
}

func (i *Image) FFT(modulus bool, exception *ExceptionInfo) *Image {
	return ForwardFourierTransformImage(i, modulus, exception)
}

func (i *Image) Function(function T.MagickFunction, numberParameters int32, parameters *float64, exception *ExceptionInfo) bool {
	return FunctionImage(i, function, numberParameters, parameters, exception)
}

func (i *Image) FunctionChannel(channel T.ChannelType, function T.MagickFunction, numberParameters int32, argument *float64, exception *ExceptionInfo) bool {
	return FunctionImageChannel(i, channel, function, numberParameters, argument, exception)
}

func (i *Image) AuthenticIndexQueue() *T.IndexPacket { return GetAuthenticIndexQueue(i) }

func (i *Image) AuthenticPixelQueue() *T.PixelPacket { return GetAuthenticPixelQueue(i) }

func (i *Image) AuthenticPixels(x, y int32, columns, rows uint32, exception *ExceptionInfo) *T.PixelPacket {
	return GetAuthenticPixels(i, x, y, columns, rows, exception)
}

func (i *Image) BlobProperties() *stat { return GetBlobProperties(i) }

func (i *Image) AlphaChannel() bool { return GetImageAlphaChannel(i) }

func (i *Image) ChannelDistortions(reconstructImage *Image, metric T.MetricType, exception *ExceptionInfo) *float64 {
	return GetImageChannelDistortions(i, reconstructImage, metric, exception)
}

func (i *Image) ChannelFeatures(distance uint32, exception *ExceptionInfo) *T.ChannelFeatures {
	return GetImageChannelFeatures(i, distance, exception)
}

func (i *Image) ChannelKurtosis(channel T.ChannelType, kurtosis, kewness *float64, exception *ExceptionInfo) bool {
	return GetImageChannelKurtosis(i, channel, kurtosis, kewness, exception)
}

func (i *Image) Channels() uint32 { return GetImageChannels(i) }

func (i *Image) AuthenticPixel(x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) bool {
	return GetOneAuthenticPixel(i, x, y, pixel, exception)
}

func (i *Image) VirtualMagickPixel(x, y int32, pixel *MagickPixelPacket, exception *ExceptionInfo) bool {
	return GetOneVirtualMagickPixel(i, x, y, pixel, exception)
}

func (i *Image) VirtualMethodPixel(virtualPixelMethod T.VirtualPixelMethod, x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) {
	GetOneVirtualMethodPixel(i, virtualPixelMethod, x, y, pixel, exception)
}

func (i *Image) VirtualPixel(x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) bool {
	return GetOneVirtualPixel(i, x, y, pixel, exception)
}

func (i *Image) VirtualIndexQueue() *T.IndexPacket { return GetVirtualIndexQueue(i) }

func (i *Image) VirtualPixelQueue() *T.PixelPacket { return GetVirtualPixelQueue(i) }

func (i *Image) VirtualPixels(x, y int32, columns, rows uint32, exception *ExceptionInfo) *T.PixelPacket {
	return GetVirtualPixels(i, x, y, columns, rows, exception)
}

func (i *Image) HaldClut(haldImage *Image) bool { return HaldClutImage(i, haldImage) }

func (i *Image) HaldClutChannel(channel T.ChannelType, haldImage *Image) bool {
	return HaldClutImageChannel(i, channel, haldImage)
}

func (i *Image) ToHBITMAP() **T.Void { return ImageToHBITMAP(i) }

func (i *Image) IntegralRotate(rotations uint32, exception *ExceptionInfo) *Image {
	return IntegralRotateImage(i, rotations, exception)
}

func (i *Image) InverseFourierTransform(phaseImage *Image, modulus bool, exception *ExceptionInfo) *Image {
	return InverseFourierTransformImage(i, phaseImage, modulus, exception)
}

func (i *Image) LevelColors(blackColor, whiteColor *MagickPixelPacket, invert bool) bool {
	return LevelColorsImage(i, blackColor, whiteColor, invert)
}

func (i *Image) LevelColorsChannel(channel T.ChannelType, blackColor, whiteColor *MagickPixelPacket, invert bool) bool {
	return LevelColorsImageChannel(i, channel, blackColor, whiteColor, invert)
}

func (i *Image) LevelImageColors(channel T.ChannelType, blackColor, whiteColor *MagickPixelPacket, invert bool) bool {
	// Name conflict w/LevelColorsImage if shortened
	return LevelImageColors(i, channel, blackColor, whiteColor, invert)
}

func (i *Image) Levelize(blackPoint, whitePoint, gamma float64) bool {
	return LevelizeImage(i, blackPoint, whitePoint, gamma)
}

func (i *Image) LevelizeChannel(channel T.ChannelType, blackPoint, whitePoint, gamma float64) bool {
	return LevelizeImageChannel(i, channel, blackPoint, whitePoint, gamma)
}

func (i *Image) MinMaxStretch(channel T.ChannelType, blackAdjust, whiteAdjust float64) bool {
	return MinMaxStretchImage(i, channel, blackAdjust, whiteAdjust)
}

func (i *Image) Mode(radius float64, exception *ExceptionInfo) *Image {
	return ModeImage(i, radius, exception)
}

func (i *Image) MorphologyApply(method T.MorphologyMethod, channel T.ChannelType, iterations int32, kernel *T.KernelInfo, compose T.CompositeMethod, bias float64, exception *ExceptionInfo) *Image {
	return MorphologyApply(i, method, channel, iterations, kernel, compose, bias, exception)
}

func (i *Image) Morphology(method T.MorphologyMethod, iterations int32, kernel *T.KernelInfo, exception *ExceptionInfo) *Image {
	return MorphologyImage(i, method, iterations, kernel, exception)
}

func (i *Image) MorphologyChannel(channel T.ChannelType, method T.MorphologyMethod, iterations int32, kernel *T.KernelInfo, exception *ExceptionInfo) *Image {
	return MorphologyImageChannel(i, channel, method, iterations, kernel, exception)
}

func (i *Image) MotionBlurChannel(channel T.ChannelType, radius, sigma, angle float64, exception *ExceptionInfo) *Image {
	return MotionBlurImageChannel(i, channel, radius, sigma, angle, exception)
}

func (i *Image) ParseSizeGeometry(geometry string, regionInfo *T.RectangleInfo) T.MagickStatusType { // doc RectangeInfo
	return ParseSizeGeometry(i, geometry, regionInfo)
}

func (i *Image) PasskeyDecipher(passkey *StringInfo, exception *ExceptionInfo) bool {
	return PasskeyDecipherImage(i, passkey, exception)
}

func (i *Image) PasskeyEncipher(passkey *StringInfo, exception *ExceptionInfo) bool {
	return PasskeyEncipherImage(i, passkey, exception)
}

func (i *Image) PosterizeChannel(channel T.ChannelType, levels uint32, dither bool) bool {
	return PosterizeImageChannel(i, channel, levels, dither)
}

func (i *Image) QueueAuthenticPixels(x, y int32, columns, rows uint32, exception *ExceptionInfo) *T.PixelPacket {
	return QueueAuthenticPixels(i, x, y, columns, rows, exception)
}

func (i *Image) SelectiveBlur(radius, sigma, threshold float64, exception *ExceptionInfo) *Image {
	return SelectiveBlurImage(i, radius, sigma, threshold, exception)
}

func (i *Image) SelectiveBlurChannel(channel T.ChannelType, radius, sigma, threshold float64, exception *ExceptionInfo) *Image {
	return SelectiveBlurImageChannel(i, channel, radius, sigma, threshold, exception)
}

func (i *Image) SetChannels(channels uint32) bool { return SetImageChannels(i, channels) }

func (i *Image) SetColor(color *MagickPixelPacket) bool { return SetImageColor(i, color) }

func (i *Image) SetColorspace(colorspace T.ColorspaceType) {
	SetImageColorspace(i, colorspace)
}

func (i *Image) ShearRotate(degrees float64, exception *ExceptionInfo) *Image {
	return ShearRotateImage(i, degrees, exception)
}

func (i *Image) Similarity(reference *Image, offset *T.RectangleInfo, similarity *float64, exception *ExceptionInfo) *Image {
	return SimilarityImage(i, reference, offset, similarity, exception)
}

func (i *Image) SparseColor(channel T.ChannelType, method T.SparseColorMethod, numberArguments uint32, arguments *float64, exception *ExceptionInfo) *Image {
	return SparseColorImage(i, channel, method, numberArguments, arguments, exception)
}

func (i *Image) Statistic(type_ T.StatisticType, width, height uint32, exception *ExceptionInfo) *Image {
	return StatisticImage(i, type_, width, height, exception)
}

func (i *Image) StatisticChannel(channel T.ChannelType, type_ T.StatisticType, width, height uint32, exception *ExceptionInfo) *Image {
	return StatisticImageChannel(i, channel, type_, width, height, exception)
}

func (i *Image) StereoAnaglyph(rightImage *Image, xOffset, yOffset int32, exception *ExceptionInfo) *Image {
	return StereoAnaglyphImage(i, rightImage, xOffset, yOffset, exception)
}

func (i *Image) SyncAuthenticPixels(exception *ExceptionInfo) bool {
	return SyncAuthenticPixels(i, exception)
}

var CloneImageList func(images *Image, exception *ExceptionInfo) *Image

var CloneImages func(images *Image, scenes string, exception *ExceptionInfo) *Image

var ConsolidateCMYKImages func(images *Image, exception *ExceptionInfo) *Image

var DeconstructImages func(images *Image, exception *ExceptionInfo) *Image

var DeleteImageList func(images *Image, offset T.Long) uint

var DestroyImageList func(images *Image) *Image

var GetFirstImageInList func(images *Image) *Image

var GetImageFromList func(images *Image, index T.Long) *Image

var GetImageIndexInList func(images *Image) T.Long

var GetImageList func(images *Image, offset T.Long, exception *ExceptionInfo) *Image

var GetImageListIndex func(images *Image) T.Long

var GetImageListLength func(images *Image) T.Size

var GetImageListSize func(images *Image) T.Size

var GetLastImageInList func(images *Image) *Image

var GetNextImage func(images *Image) *Image

var GetNextImageInList func(images *Image) *Image

var GetPreviousImage func(images *Image) *Image

var GetPreviousImageInList func(images *Image) *Image

var ImageListToArray func(images *Image, exception *ExceptionInfo) []*Image

var MontageImages func(images *Image, montageInfo *T.MontageInfo, exception *ExceptionInfo) *Image

var SpliceImageList func(images *Image, offset T.Long, length T.Size, splices *Image, exception *ExceptionInfo) *Image

var SplitImageList func(images *Image) *Image

var SyncImageList func(images *Image)

var SyncNextImageInList func(images *Image) *Image

var AppendImageToList func(images **Image, image *Image)

var DeleteImageFromList func(images **Image)

var DeleteImages func(images **Image, scenes string, exception *ExceptionInfo)

var InsertImageInList func(images **Image, image *Image)

var PopImageList func(images **Image) *Image

var PrependImageToList func(images **Image, image *Image)

var PushImageList func(images **Image, image *Image, exception *ExceptionInfo) uint

var RemoveDuplicateLayers func(images **Image, exception *ExceptionInfo)

var RemoveFirstImageFromList func(images **Image) *Image

var RemoveImageFromList func(images **Image) *Image

var RemoveLastImageFromList func(images **Image) *Image

var RemoveZeroDelayLayers func(images **Image, exception *ExceptionInfo)

var ReplaceImageInList func(images **Image, image *Image)

var ReverseImageList func(images **Image)

var SetImageList func(images **Image, image *Image, offset T.Long, exception *ExceptionInfo) uint

var ShiftImageList func(images **Image) *Image

var SpliceImageIntoList func(images **Image, length T.Size, splice *Image) *Image

var TransformImages func(images **Image, cropGeometry, imageGeometry string) bool

var UnshiftImageList func(images **Image, image *Image, exception *ExceptionInfo) uint

var AcquireCacheViewIndexes func(c *CacheView) *T.IndexPacket

var AcquireCacheViewPixels func(c *CacheView, x, y T.Long, columns, rows T.Size, exception *ExceptionInfo) *T.PixelPacket

var AcquireOneCacheViewPixel func(c *CacheView, x, y T.Long, exception *ExceptionInfo) T.PixelPacket

var CloneCacheView func(c *CacheView) *CacheView

var CloseCacheView func(c *CacheView) *CacheView

var GetCacheView func(c *CacheView, x, y T.Long, columns, rows T.Size) *T.PixelPacket

var GetCacheViewColorspace func(c *CacheView) T.ColorspaceType

var GetCacheViewException func(c *CacheView) *ExceptionInfo

var GetCacheViewIndexes func(c *CacheView) *T.IndexPacket

var GetCacheViewPixels func(c *CacheView, x, y T.Long, columns, rows T.Size) *T.PixelPacket

var GetCacheViewStorageClass func(c *CacheView) T.ClassType

var SetCacheViewStorageClass func(c *CacheView, storageClass T.ClassType) bool

var SetCacheViewVirtualPixelMethod func(c *CacheView, virtualPixelMethod T.VirtualPixelMethod) bool

//sketchy docs and/or deprecated

var DestroyCacheView func(c *CacheView) *CacheView

var GetCacheViewAuthenticIndexQueue func(c *CacheView) *T.IndexPacket

var GetCacheViewAuthenticPixelQueue func(c *CacheView) *T.PixelPacket

var GetCacheViewAuthenticPixels func(c *CacheView, x, y int32, columns, rows uint32, exception *ExceptionInfo) *T.PixelPacket

var GetCacheViewChannels func(c *CacheView) uint32

var GetCacheViewVirtualIndexQueue func(c *CacheView) *T.IndexPacket

var GetCacheViewVirtualPixelQueue func(c *CacheView) *T.PixelPacket

var GetCacheViewVirtualPixels func(c *CacheView, x, y int32, columns, rows uint32, exception *ExceptionInfo) *T.PixelPacket

var GetOneCacheViewAuthenticPixel func(c *CacheView, x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) bool // doc Pixelpacket
var GetOneCacheViewVirtualMethodPixel func(c *CacheView, virtualPixelMethod T.VirtualPixelMethod, x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) bool

var GetOneCacheViewVirtualPixel func(c *CacheView, x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) bool

var QueueCacheViewAuthenticPixels func(c *CacheView, x, y int32, columns, rows uint32, exception *ExceptionInfo) *T.PixelPacket

var SetCacheViewPixels func(c *CacheView, x, y int32, columns, rows uint32) *T.PixelPacket

var SyncCacheView func(c *CacheView) bool

var SyncCacheViewAuthenticPixels func(c *CacheView, exception *ExceptionInfo) bool

var SyncCacheViewPixels func(c *CacheView) bool

func (c *CacheView) AcquireIndexes() *T.IndexPacket { return AcquireCacheViewIndexes(c) }

func (c *CacheView) AcquirePixels(x, y T.Long, columns, rows T.Size, exception *ExceptionInfo) *T.PixelPacket {
	return AcquireCacheViewPixels(c, x, y, columns, rows, exception)
}

func (c *CacheView) AcquirePixel(x, y T.Long, exception *ExceptionInfo) T.PixelPacket {
	return AcquireOneCacheViewPixel(c, x, y, exception)
}

func (c *CacheView) Clone() *CacheView { return CloneCacheView(c) }

func (c *CacheView) Close() *CacheView { return CloseCacheView(c) }

func (c *CacheView) Get(x, y T.Long, columns, rows T.Size) *T.PixelPacket {
	return GetCacheView(c, x, y, columns, rows)
}

func (c *CacheView) Colorspace() T.ColorspaceType { return GetCacheViewColorspace(c) }

func (c *CacheView) Exception() *ExceptionInfo { return GetCacheViewException(c) }

func (c *CacheView) ViewIndexes() *T.IndexPacket { return GetCacheViewIndexes(c) }

func (c *CacheView) Pixels(x, y T.Long, columns, rows T.Size) *T.PixelPacket {
	return GetCacheViewPixels(c, x, y, columns, rows)
}

func (c *CacheView) StorageClass() T.ClassType { return GetCacheViewStorageClass(c) }

func (c *CacheView) SetStorageClass(storageClass T.ClassType) bool {
	return SetCacheViewStorageClass(c, storageClass)
}

func (c *CacheView) SetVirtualPixelMethod(virtualPixelMethod T.VirtualPixelMethod) bool {
	return SetCacheViewVirtualPixelMethod(c, virtualPixelMethod)
}

func (c *CacheView) Destroy() *CacheView { return DestroyCacheView(c) }

func (c *CacheView) AuthenticIndexQueue() *T.IndexPacket { return GetCacheViewAuthenticIndexQueue(c) }

func (c *CacheView) AuthenticPixelQueue() *T.PixelPacket { return GetCacheViewAuthenticPixelQueue(c) }

func (c *CacheView) AuthenticPixels(x, y int32, columns, rows uint32, exception *ExceptionInfo) *T.PixelPacket {
	return GetCacheViewAuthenticPixels(c, x, y, columns, rows, exception)
}

func (c *CacheView) Channels() uint32 { return GetCacheViewChannels(c) }

func (c *CacheView) VirtualIndexQueue() *T.IndexPacket { return GetCacheViewVirtualIndexQueue(c) }

func (c *CacheView) VirtualPixelQueue() *T.PixelPacket { return GetCacheViewVirtualPixelQueue(c) }

func (c *CacheView) VirtualPixels(x, y int32, columns, rows uint32, exception *ExceptionInfo) *T.PixelPacket {
	return GetCacheViewVirtualPixels(c, x, y, columns, rows, exception)
}

func (c *CacheView) AuthenticPixel(x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) bool {
	return GetOneCacheViewAuthenticPixel(c, x, y, pixel, exception)
}

func (c *CacheView) VirtualMethodPixel(virtualPixelMethod T.VirtualPixelMethod, x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) bool {
	return GetOneCacheViewVirtualMethodPixel(c, virtualPixelMethod, x, y, pixel, exception)
}

func (c *CacheView) VirtualPixel(x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) bool {
	return GetOneCacheViewVirtualPixel(c, x, y, pixel, exception)
}

func (c *CacheView) QueueAuthenticPixels(x, y int32, columns, rows uint32, exception *ExceptionInfo) *T.PixelPacket {
	return QueueCacheViewAuthenticPixels(c, x, y, columns, rows, exception)
}

func (c *CacheView) SetPixels(x, y int32, columns, rows uint32) *T.PixelPacket {
	return SetCacheViewPixels(c, x, y, columns, rows)
}

func (c *CacheView) Sync() bool { return SyncCacheView(c) }

func (c *CacheView) SyncAuthenticPixels(exception *ExceptionInfo) bool {
	return SyncCacheViewAuthenticPixels(c, exception)
}

func (c *CacheView) SyncPixels() bool { return SyncCacheViewPixels(c) }

var GetImageDecoder func(m *MagickInfo) *T.DecodeImageHandler

var GetImageEncoder func(m *MagickInfo) *T.EncodeImageHandler

var GetMagickAdjoin func(m *MagickInfo) bool

var GetMagickBlobSupport func(m *MagickInfo) bool

var GetMagickDescription func(m *MagickInfo) string

var GetMagickEndianSupport func(m *MagickInfo) bool

var GetMagickSeekableStream func(m *MagickInfo) bool

var GetMagickThreadSupport func(m *MagickInfo) T.MagickStatusType

var RegisterMagickInfo func(m *MagickInfo) *MagickInfo

func (m *MagickInfo) Decoder() *T.DecodeImageHandler { return GetImageDecoder(m) }

func (m *MagickInfo) Encoder() *T.EncodeImageHandler { return GetImageEncoder(m) }

func (m *MagickInfo) Adjoin() bool { return GetMagickAdjoin(m) }

func (m *MagickInfo) BlobSupport() bool { return GetMagickBlobSupport(m) }

func (m *MagickInfo) Description() string { return GetMagickDescription(m) }

func (m *MagickInfo) EndianSupport() bool { return GetMagickEndianSupport(m) }

func (m *MagickInfo) SeekableStream() bool { return GetMagickSeekableStream(m) }

func (m *MagickInfo) ThreadSupport() T.MagickStatusType { return GetMagickThreadSupport(m) }

func (m *MagickInfo) Register() *MagickInfo { return RegisterMagickInfo(m) }

var XAnimateBackgroundImage func(d *Display, resourceInfo *T.XResourceInfo, images *Image)

var XAnimateImages func(d *Display, resourceInfo *T.XResourceInfo, argv []string, argc int, images *Image) *Image

var XAnnotateImage func(d *Display, pixel *T.XPixelInfo, annotateInfo *T.XAnnotateInfo, image *Image) bool

var XBestFont func(d *Display, resourceInfo *T.XResourceInfo, textFont bool) *T.XFontStruct

var XBestIconSize func(d *Display, window *T.XWindowInfo, image *Image)

var XBestPixel func(d *Display, colormap T.Colormap, colors *T.XColor, numberColors uint, color *T.XColor)

var XBestVisualInfo func(d *Display, mapInfo *T.XStandardColormap, resourceInfo *T.XResourceInfo) *T.XVisualInfo

var XCheckDefineCursor func(d *Display, window T.Window, cursor T.Cursor) int

var XCheckRefreshWindows func(d *Display, windows *T.XWindows)

var XClientMessage func(d *Display, window T.Window, protocol, reason T.Atom, timestamp T.Time)

var XColorBrowserWidget func(d *Display, windows *T.XWindows, action, reply string)

var XCommandWidget func(d *Display, windows *T.XWindows, selections []string, event *T.XEvent) int

var XConfigureImageColormap func(d *Display, resourceInfo *T.XResourceInfo, windows *T.XWindows, image *Image)

var XConfirmWidget func(d *Display, windows *T.XWindows, reason, description string) int

var XConstrainWindowPosition func(d *Display, windowInfo *T.XWindowInfo)

var XDelay func(d *Display, milliseconds T.Size)

var XDestroyWindowColors func(d *Display, window T.Window)

var XDialogWidget func(d *Display, windows *T.XWindows, action, query, reply string) int

var XDisplayBackgroundImage func(d *Display, resourceInfo *T.XResourceInfo, image *Image) bool

var XDisplayImage func(d *Display, resourceInfo *T.XResourceInfo, argv []string, argc int, image **Image, state *T.Size) *Image

var XDisplayImageInfo func(d *Display, resourceInfo *T.XResourceInfo, windows *T.XWindows, undoImage, image *Image)

var XDrawImage func(d *Display, pixel *T.XPixelInfo, drawInfo *T.XDrawInfo, image *Image) bool

var XError func(d *Display, err *T.XErrorEvent) int

var XFileBrowserWidget func(d *Display, windows *T.XWindows, action, reply string)

var XFontBrowserWidget func(d *Display, windows *T.XWindows, action, reply string)

var XFreeResources func(d *Display, visualInfo *T.XVisualInfo, mapInfo *T.XStandardColormap, pixel *T.XPixelInfo, fontInfo *T.XFontStruct, resourceInfo *T.XResourceInfo, windowInfo *T.XWindowInfo)

var XFreeStandardColormap func(d *Display, visualInfo *T.XVisualInfo, mapInfo *T.XStandardColormap, pixel *T.XPixelInfo)

var XGetPixelPacket func(d *Display, visualInfo *T.XVisualInfo, mapInfo *T.XStandardColormap, resourceInfo *T.XResourceInfo, image *Image, pixel *T.XPixelInfo)

var XGetResourceDatabase func(d *Display, clientName string) XrmDatabase

var XGetScreenDensity func(d *Display) string

var XGetWindowColor func(d *Display, windows *T.XWindows, name string) bool

var XGetWindowInfo func(d *Display, visualInfo *T.XVisualInfo, mapInfo *T.XStandardColormap, pixel *T.XPixelInfo, fontInfo *T.XFontStruct, resourceInfo *T.XResourceInfo, window *T.XWindowInfo)

var XHighlightEllipse func(d *Display, window T.Window, annotateContext T.GC, highlightInfo *T.RectangleInfo)

var XHighlightLine func(d *Display, window T.Window, annotateContext T.GC, highlightInfo *T.XSegment)

var XHighlightRectangle func(d *Display, window T.Window, annotateContext T.GC, highlightInfo *T.RectangleInfo)

var XInfoWidget func(d *Display, windows *T.XWindows, activity string)

var XInitializeWindows func(d *Display, resourceInfo *T.XResourceInfo) *T.XWindows

var XListBrowserWidget func(d *Display, windows *T.XWindows, windowInfo *T.XWindowInfo, list []string, action, query, reply string)

var XMakeCursor func(d *Display, window T.Window, colormap T.Colormap, backgroundColor, foregroundColor string) T.Cursor

var XMakeImage func(d *Display, resourceInfo *T.XResourceInfo, window *T.XWindowInfo, image *Image, width, height uint) bool

var XMakeMagnifyImage func(d *Display, windows *T.XWindows)

var XMakeStandardColormap func(d *Display, visualInfo *T.XVisualInfo, resourceInfo *T.XResourceInfo, image *Image, mapInfo *T.XStandardColormap, pixel *T.XPixelInfo)

var XMakeWindow func(d *Display, parent T.Window, argv []string, argc int, classHint *T.XClassHint, managerHints *T.XWMHints, windowInfo *T.XWindowInfo)

var XMenuWidget func(d *Display, windows *T.XWindows, title string, selections []string, item string) int

var XNoticeWidget func(d *Display, windows *T.XWindows, reason, description string)

var XPreferencesWidget func(d *Display, resourceInfo *T.XResourceInfo, windows *T.XWindows) bool

var XProgressMonitorWidget func(d *Display, windows *T.XWindows, task string, offset T.MagickOffsetType, span T.MagickSizeType)

var XQueryPosition func(d *Display, window T.Window, x, y *int)

var XRefreshWindow func(d *Display, window *T.XWindowInfo, event *T.XEvent)

var XRemoteCommand func(d *Display, window, filename string) bool

var XRetainWindowColors func(d *Display, window T.Window)

var XrmGetDatabase func(d *Display) XrmDatabase

var XSetCursorState func(d *Display, windows *T.XWindows, state T.MagickStatusType)

var XTextViewWidget func(d *Display, resourceInfo *T.XResourceInfo, windows *T.XWindows, mono bool, title string, textlist []string)

var XWindowByID func(d *Display, rootWindow T.Window, id T.Size) T.Window

var XWindowByName func(d *Display, rootWindow T.Window, name string) T.Window

var XWindowByProperty func(d *Display, window T.Window, property T.Atom) T.Window

func (d *Display) XAnimateBackgroundImage(resourceInfo *T.XResourceInfo, images *Image) {
	XAnimateBackgroundImage(d, resourceInfo, images)
}

func (d *Display) XAnimateImages(resourceInfo *T.XResourceInfo, argv []string, argc int, images *Image) *Image {
	return XAnimateImages(d, resourceInfo, argv, argc, images)
}

func (d *Display) XAnnotateImage(pixel *T.XPixelInfo, annotateInfo *T.XAnnotateInfo, image *Image) bool {
	return XAnnotateImage(d, pixel, annotateInfo, image)
}

func (d *Display) XBestFont(resourceInfo *T.XResourceInfo, textFont bool) *T.XFontStruct {
	return XBestFont(d, resourceInfo, textFont)
}

func (d *Display) XBestIconSize(window *T.XWindowInfo, image *Image) {
	XBestIconSize(d, window, image)
}

func (d *Display) XBestPixel(colormap T.Colormap, colors *T.XColor, numberColors uint, color *T.XColor) {
	XBestPixel(d, colormap, colors, numberColors, color)
}

func (d *Display) XBestVisualInfo(mapInfo *T.XStandardColormap, resourceInfo *T.XResourceInfo) *T.XVisualInfo {
	return XBestVisualInfo(d, mapInfo, resourceInfo)
}

func (d *Display) XCheckDefineCursor(window T.Window, cursor T.Cursor) int {
	return XCheckDefineCursor(d, window, cursor)
}

func (d *Display) XCheckRefreshWindows(windows *T.XWindows) { XCheckRefreshWindows(d, windows) }

func (d *Display) XClientMessage(window T.Window, protocol, reason T.Atom, timestamp T.Time) {
	XClientMessage(d, window, protocol, reason, timestamp)
}

func (d *Display) XColorBrowserWidget(windows *T.XWindows, action, reply string) {
	XColorBrowserWidget(d, windows, action, reply)
}

func (d *Display) XCommandWidget(windows *T.XWindows, selections []string, event *T.XEvent) int {
	return XCommandWidget(d, windows, selections, event)
}

func (d *Display) XConfigureImageColormap(resourceInfo *T.XResourceInfo, windows *T.XWindows, image *Image) {
	XConfigureImageColormap(d, resourceInfo, windows, image)
}

func (d *Display) XConfirmWidget(windows *T.XWindows, reason, description string) int {
	return XConfirmWidget(d, windows, reason, description)
}

func (d *Display) XConstrainWindowPosition(windowInfo *T.XWindowInfo) {
	XConstrainWindowPosition(d, windowInfo)
}

func (d *Display) XDelay(milliseconds T.Size) { XDelay(d, milliseconds) }

func (d *Display) XDestroyWindowColors(window T.Window) { XDestroyWindowColors(d, window) }

func (d *Display) XDialogWidget(windows *T.XWindows, action, query, reply string) int {
	return XDialogWidget(d, windows, action, query, reply)
}

func (d *Display) XBackgroundImage(resourceInfo *T.XResourceInfo, image *Image) bool {
	return XDisplayBackgroundImage(d, resourceInfo, image)
}

func (d *Display) XDisplayImage(resourceInfo *T.XResourceInfo, argv []string, argc int, image **Image, state *T.Size) *Image {
	return XDisplayImage(d, resourceInfo, argv, argc, image, state)
}

func (d *Display) XDisplayImageInfo(resourceInfo *T.XResourceInfo, windows *T.XWindows, undoImage, image *Image) {
	XDisplayImageInfo(d, resourceInfo, windows, undoImage, image)
}

func (d *Display) XDrawImage(pixel *T.XPixelInfo, drawInfo *T.XDrawInfo, image *Image) bool {
	return XDrawImage(d, pixel, drawInfo, image)
}

func (d *Display) XError(err *T.XErrorEvent) int { return XError(d, err) }

func (d *Display) XFileBrowserWidget(windows *T.XWindows, action, reply string) {
	XFileBrowserWidget(d, windows, action, reply)
}

func (d *Display) XFontBrowserWidget(windows *T.XWindows, action, reply string) {
	XFontBrowserWidget(d, windows, action, reply)
}

func (d *Display) XFreeResources(visualInfo *T.XVisualInfo, mapInfo *T.XStandardColormap, pixel *T.XPixelInfo, fontInfo *T.XFontStruct, resourceInfo *T.XResourceInfo, windowInfo *T.XWindowInfo) {
	XFreeResources(d, visualInfo, mapInfo, pixel, fontInfo, resourceInfo, windowInfo)
}

func (d *Display) XFreeStandardColormap(visualInfo *T.XVisualInfo, mapInfo *T.XStandardColormap, pixel *T.XPixelInfo) {
	XFreeStandardColormap(d, visualInfo, mapInfo, pixel)
}

func (d *Display) XGetPixelPacket(visualInfo *T.XVisualInfo, mapInfo *T.XStandardColormap, resourceInfo *T.XResourceInfo, image *Image, pixel *T.XPixelInfo) {
	XGetPixelPacket(d, visualInfo, mapInfo, resourceInfo, image, pixel)
}

func (d *Display) XGetResourceDatabase(clientName string) XrmDatabase {
	return XGetResourceDatabase(d, clientName)
}

func (d *Display) XGetScreenDensity() string { return XGetScreenDensity(d) }

func (d *Display) XGetWindowColor(windows *T.XWindows, name string) bool {
	return XGetWindowColor(d, windows, name)
}

func (d *Display) XGetWindowInfo(visualInfo *T.XVisualInfo, mapInfo *T.XStandardColormap, pixel *T.XPixelInfo, fontInfo *T.XFontStruct, resourceInfo *T.XResourceInfo, window *T.XWindowInfo) {
	XGetWindowInfo(d, visualInfo, mapInfo, pixel, fontInfo, resourceInfo, window)
}

func (d *Display) XHighlightEllipse(window T.Window, annotateContext T.GC, highlightInfo *T.RectangleInfo) {
	XHighlightEllipse(d, window, annotateContext, highlightInfo)
}

func (d *Display) XHighlightLine(window T.Window, annotateContext T.GC, highlightInfo *T.XSegment) {
	XHighlightLine(d, window, annotateContext, highlightInfo)
}

func (d *Display) XHighlightRectangle(window T.Window, annotateContext T.GC, highlightInfo *T.RectangleInfo) {
	XHighlightRectangle(d, window, annotateContext, highlightInfo)
}

func (d *Display) XInfoWidget(windows *T.XWindows, activity string) { XInfoWidget(d, windows, activity) }

func (d *Display) XInitializeWindows(resourceInfo *T.XResourceInfo) *T.XWindows {
	return XInitializeWindows(d, resourceInfo)
}

func (d *Display) XListBrowserWidget(windows *T.XWindows, windowInfo *T.XWindowInfo, list []string, action, query, reply string) {
	XListBrowserWidget(d, windows, windowInfo, list, action, query, reply)
}

func (d *Display) XMakeCursor(window T.Window, colormap T.Colormap, backgroundColor, foregroundColor string) T.Cursor {
	return XMakeCursor(d, window, colormap, backgroundColor, foregroundColor)
}

func (d *Display) XMakeImage(resourceInfo *T.XResourceInfo, window *T.XWindowInfo, image *Image, width, height uint) bool {
	return XMakeImage(d, resourceInfo, window, image, width, height)
}

func (d *Display) XMakeMagnifyImage(windows *T.XWindows) { XMakeMagnifyImage(d, windows) }

func (d *Display) XMakeStandardColormap(visualInfo *T.XVisualInfo, resourceInfo *T.XResourceInfo, image *Image, mapInfo *T.XStandardColormap, pixel *T.XPixelInfo) {
	XMakeStandardColormap(d, visualInfo, resourceInfo, image, mapInfo, pixel)
}

func (d *Display) XMakeWindow(parent T.Window, argv []string, argc int, classHint *T.XClassHint, managerHints *T.XWMHints, windowInfo *T.XWindowInfo) {
	XMakeWindow(d, parent, argv, argc, classHint, managerHints, windowInfo)
}

func (d *Display) XMenuWidget(windows *T.XWindows, title string, selections []string, item string) int {
	return XMenuWidget(d, windows, title, selections, item)
}

func (d *Display) XNoticeWidget(windows *T.XWindows, reason, description string) {
	XNoticeWidget(d, windows, reason, description)
}

func (d *Display) XPreferencesWidget(resourceInfo *T.XResourceInfo, windows *T.XWindows) bool {
	return XPreferencesWidget(d, resourceInfo, windows)
}

func (d *Display) XProgressMonitorWidget(windows *T.XWindows, task string, offset T.MagickOffsetType, span T.MagickSizeType) {
	XProgressMonitorWidget(d, windows, task, offset, span)
}

func (d *Display) XQueryPosition(window T.Window, x, y *int) { XQueryPosition(d, window, x, y) }

func (d *Display) XRefreshWindow(window *T.XWindowInfo, event *T.XEvent) {
	XRefreshWindow(d, window, event)
}

func (d *Display) XRemoteCommand(window, filename string) bool {
	return XRemoteCommand(d, window, filename)
}

func (d *Display) XRetainWindowColors(window T.Window) { XRetainWindowColors(d, window) }

func (d *Display) XrmGetDatabase() XrmDatabase { return XrmGetDatabase(d) }

func (d *Display) XSetCursorState(windows *T.XWindows, state T.MagickStatusType) {
	XSetCursorState(d, windows, state)
}

func (d *Display) XTextViewWidget(resourceInfo *T.XResourceInfo, windows *T.XWindows, mono bool, title string, textlist []string) {
	XTextViewWidget(d, resourceInfo, windows, mono, title, textlist)
}

func (d *Display) XWindowByID(rootWindow T.Window, id T.Size) T.Window {
	return XWindowByID(d, rootWindow, id)
}

func (d *Display) XWindowByName(rootWindow T.Window, name string) T.Window {
	return XWindowByName(d, rootWindow, name)
}

func (d *Display) XWindowByProperty(window T.Window, property T.Atom) T.Window {
	return XWindowByProperty(d, window, property)
}

var AcquireImageInfo func() *ImageInfo

var AcquireQuantizeInfo func(i *ImageInfo) *QuantizeInfo

var AcquireQuantumInfo func(i *ImageInfo) *T.QuantumInfo

var AcquireStreamInfo func(i *ImageInfo) *StreamInfo

var AllocateImage func(i *ImageInfo) *Image

var AllocateNextImage func(i *ImageInfo, image *Image)

var AnimateImages func(i *ImageInfo, images *Image) bool

var BlobToImage func(i *ImageInfo, blob *T.Void, length uint32, exception *ExceptionInfo) *Image

var CloneDrawInfo func(i *ImageInfo, drawInfo *DrawInfo) *DrawInfo

var CloneImageInfo func(i *ImageInfo) *ImageInfo

var CloneImageOptions func(i *ImageInfo, cloneInfo *ImageInfo) bool

var CloneMontageInfo func(i *ImageInfo, montageInfo *T.MontageInfo) *T.MontageInfo

var DefineImageOption func(i *ImageInfo, option string) bool

var DeleteImageOption func(i *ImageInfo, option string) bool

var DestroyImageInfo func(i *ImageInfo) *ImageInfo

var DestroyImageOptions func(i *ImageInfo)

var DisplayImages func(i *ImageInfo, images *Image) bool

var GetDelegateCommand func(i *ImageInfo, image *Image, decode, encode string, exception *ExceptionInfo) string

var GetDrawInfo func(i *ImageInfo, drawInfo *DrawInfo)

var GetImageInfo func(i *ImageInfo)

var GetImageOption func(i *ImageInfo, key string) string

var GetMontageInfo func(i *ImageInfo, montageInfo *T.MontageInfo)

var GetNextImageOption func(i *ImageInfo) string

var GetQuantumInfo func(i *ImageInfo, quantumInfo *T.QuantumInfo)

var HuffmanEncodeImage func(i *ImageInfo, image *Image) bool

var ImagesToBlob func(i *ImageInfo, images *Image, length *uint32, exception *ExceptionInfo) *byte

var ImageToBlob func(i *ImageInfo, image *Image, length *uint32, exception *ExceptionInfo) *byte

var InjectImageBlob func(i *ImageInfo, image *Image, format string) bool

var InterpretImageAttributes func(i *ImageInfo, image *Image, embedText string) string

var InterpretImageProperties func(i *ImageInfo, image *Image, embedText string) string

var InvokeDelegate func(i *ImageInfo, image *Image, decode, encode string, exception *ExceptionInfo) bool

var MontageImageList func(i *ImageInfo, montageInfo *T.MontageInfo, images *Image, exception *ExceptionInfo) *Image

var NewMagickImage func(i *ImageInfo, width, height T.Size, background *MagickPixelPacket) *Image

var OpenBlob func(i *ImageInfo, image *Image, mode T.BlobMode, exception *ExceptionInfo) bool

var OpenStream func(i *ImageInfo, streamInfo *StreamInfo, filename string, exception *ExceptionInfo) bool

var PingBlob func(i *ImageInfo, blob *T.Void, length uint32, exception *ExceptionInfo) *Image

var PingImage func(i *ImageInfo, exception *ExceptionInfo) *Image

var ReadImage func(i *ImageInfo, exception *ExceptionInfo) *Image

var ReadInlineImage func(i *ImageInfo, content string, exception *ExceptionInfo) *Image

var ReadStream func(i *ImageInfo, stream T.StreamHandler, exception *ExceptionInfo) *Image

var RemoteDisplayCommand func(i *ImageInfo, window, filename string, exception *ExceptionInfo) bool

var RemoveImageOption func(i *ImageInfo, option string) string

var ResetImageOptionIterator func(i *ImageInfo)

var SetImageInfo func(i *ImageInfo, rectify bool, exception *ExceptionInfo) bool

var SetImageInfoBlob func(i *ImageInfo, blob *T.Void, length uint32)

var SetImageInfoFile func(i *ImageInfo, file *FILE)

var SetImageInfoProgressMonitor func(i *ImageInfo, progressMonitor T.MagickProgressMonitor, clientData *T.Void) T.MagickProgressMonitor

var SetImageOption func(i *ImageInfo, option, value string) bool

var StreamImage func(i *ImageInfo, streamInfo *StreamInfo, exception *ExceptionInfo) *Image

var TranslateText func(i *ImageInfo, image *Image, embedText string) string

var WriteImage func(i *ImageInfo, image *Image) bool

var WriteImages func(i *ImageInfo, images *Image, filename string, exception *ExceptionInfo) bool

var WriteStream func(i *ImageInfo, image *Image, stream T.StreamHandler) bool

var XImportImage func(i *ImageInfo, ximageInfo *T.XImportInfo) *Image

var AcquireImage func(i *ImageInfo) *Image

var AcquireNextImage func(i *ImageInfo, image *Image)

var GetImageInfoFile func(i *ImageInfo) *FILE

var GetMagickProperty func(i *ImageInfo, image *Image, property string) string

var PingImages func(i *ImageInfo, exception *ExceptionInfo) *Image

var ReadImages func(i *ImageInfo, exception *ExceptionInfo) *Image

var SyncImageSettings func(i *ImageInfo, image *Image) bool

var SyncImagesSettings func(i *ImageInfo, image *Image) bool

func (i *ImageInfo) AcquireQuantizeInfo() *QuantizeInfo { return AcquireQuantizeInfo(i) }

func (i *ImageInfo) AcquireQuantumInfo() *T.QuantumInfo { return AcquireQuantumInfo(i) }

func (i *ImageInfo) AcquireStreamInfo() *StreamInfo { return AcquireStreamInfo(i) }

func (i *ImageInfo) AllocateImage() *Image { return AllocateImage(i) }

func (i *ImageInfo) AllocateNextImage(image *Image) { AllocateNextImage(i, image) }

func (i *ImageInfo) AnimateImages(images *Image) bool { return AnimateImages(i, images) }

func (i *ImageInfo) BlobToImage(blob *T.Void, length uint32, exception *ExceptionInfo) *Image {
	return BlobToImage(i, blob, length, exception)
}

func (i *ImageInfo) CloneDrawInfo(drawInfo *DrawInfo) *DrawInfo { return CloneDrawInfo(i, drawInfo) }

func (i *ImageInfo) CloneInfo() *ImageInfo { return CloneImageInfo(i) }

func (i *ImageInfo) CloneOptions(cloneInfo *ImageInfo) bool {
	return CloneImageOptions(i, cloneInfo)
}

func (i *ImageInfo) CloneMontageInfo(montageInfo *T.MontageInfo) *T.MontageInfo {
	return CloneMontageInfo(i, montageInfo)
}

func (i *ImageInfo) DefineOption(option string) bool { return DefineImageOption(i, option) }

func (i *ImageInfo) DeleteOption(option string) bool { return DeleteImageOption(i, option) }

func (i *ImageInfo) Destroy() *ImageInfo { return DestroyImageInfo(i) }

func (i *ImageInfo) DestroyOptions() { DestroyImageOptions(i) }

func (i *ImageInfo) DisplayImages(images *Image) bool { return DisplayImages(i, images) }

func (i *ImageInfo) DelegateCommand(image *Image, decode, encode string, exception *ExceptionInfo) string {
	return GetDelegateCommand(i, image, decode, encode, exception)
}

func (i *ImageInfo) DrawInfo(drawInfo *DrawInfo) { GetDrawInfo(i, drawInfo) }

func (i *ImageInfo) Get() { GetImageInfo(i) }

func (i *ImageInfo) Option(key string) string { return GetImageOption(i, key) }

func (i *ImageInfo) MontageInfo(montageInfo *T.MontageInfo) { GetMontageInfo(i, montageInfo) }

func (i *ImageInfo) NextOption() string { return GetNextImageOption(i) }

func (i *ImageInfo) QuantumInfo(quantumInfo *T.QuantumInfo) { GetQuantumInfo(i, quantumInfo) }

func (i *ImageInfo) HuffmanEncode(image *Image) bool { return HuffmanEncodeImage(i, image) }

func (i *ImageInfo) ImagesToBlob(images *Image, length *uint32, exception *ExceptionInfo) *byte {
	return ImagesToBlob(i, images, length, exception)
}

func (i *ImageInfo) ImageToBlob(image *Image, length *uint32, exception *ExceptionInfo) *byte {
	return ImageToBlob(i, image, length, exception)
}

func (i *ImageInfo) InjectImageBlob(image *Image, format string) bool {
	return InjectImageBlob(i, image, format)
}

func (i *ImageInfo) InterpretAttributes(image *Image, embedText string) string {
	return InterpretImageAttributes(i, image, embedText)
}

func (i *ImageInfo) InterpretProperties(image *Image, embedText string) string {
	return InterpretImageProperties(i, image, embedText)
}

func (i *ImageInfo) InvokeDelegate(image *Image, decode, encode string, exception *ExceptionInfo) bool {
	return InvokeDelegate(i, image, decode, encode, exception)
}

func (i *ImageInfo) MontageImageList(montageInfo *T.MontageInfo, images *Image, exception *ExceptionInfo) *Image {
	return MontageImageList(i, montageInfo, images, exception)
}

func (i *ImageInfo) NewMagickImage(width, height T.Size, background *MagickPixelPacket) *Image {
	return NewMagickImage(i, width, height, background)
}

func (i *ImageInfo) OpenBlob(image *Image, mode T.BlobMode, exception *ExceptionInfo) bool {
	return OpenBlob(i, image, mode, exception)
}

func (i *ImageInfo) OpenStream(streamInfo *StreamInfo, filename string, exception *ExceptionInfo) bool {
	return OpenStream(i, streamInfo, filename, exception)
}

func (i *ImageInfo) PingBlob(blob *T.Void, length uint32, exception *ExceptionInfo) *Image {
	return PingBlob(i, blob, length, exception)
}

func (i *ImageInfo) PingImage(exception *ExceptionInfo) *Image { return PingImage(i, exception) }

func (i *ImageInfo) ReadImage(exception *ExceptionInfo) *Image { return ReadImage(i, exception) }

func (i *ImageInfo) ReadInlineImage(content string, exception *ExceptionInfo) *Image {
	return ReadInlineImage(i, content, exception)
}

func (i *ImageInfo) ReadStream(stream T.StreamHandler, exception *ExceptionInfo) *Image {
	return ReadStream(i, stream, exception)
}

func (i *ImageInfo) RemoteDisplayCommand(window, filename string, exception *ExceptionInfo) bool {
	return RemoteDisplayCommand(i, window, filename, exception)
}

func (i *ImageInfo) RemoveImageOption(option string) string { return RemoveImageOption(i, option) }

func (i *ImageInfo) ResetImageOptionIterator() { ResetImageOptionIterator(i) }

func (i *ImageInfo) SetImageInfo(rectify bool, exception *ExceptionInfo) bool {
	return SetImageInfo(i, rectify, exception)
}

func (i *ImageInfo) SetBlob(blob *T.Void, length uint32) { SetImageInfoBlob(i, blob, length) }

func (i *ImageInfo) SetFile(file *FILE) { SetImageInfoFile(i, file) }

func (i *ImageInfo) SetProgressMonitor(progressMonitor T.MagickProgressMonitor, clientData *T.Void) T.MagickProgressMonitor {
	return SetImageInfoProgressMonitor(i, progressMonitor, clientData)
}

func (i *ImageInfo) SetOption(option, value string) bool {
	return SetImageOption(i, option, value)
}

func (i *ImageInfo) StreamImage(streamInfo *StreamInfo, exception *ExceptionInfo) *Image {
	return StreamImage(i, streamInfo, exception)
}

func (i *ImageInfo) TranslateText(image *Image, embedText string) string {
	return TranslateText(i, image, embedText)
}

func (i *ImageInfo) WriteImage(image *Image) bool { return WriteImage(i, image) }

func (i *ImageInfo) WriteImages(images *Image, filename string, exception *ExceptionInfo) bool {
	return WriteImages(i, images, filename, exception)
}

func (i *ImageInfo) WriteStream(image *Image, stream T.StreamHandler) bool {
	return WriteStream(i, image, stream)
}

func (i *ImageInfo) XImportImage(ximageInfo *T.XImportInfo) *Image {
	return XImportImage(i, ximageInfo)
}

func (i *ImageInfo) AcquireImage() *Image { return AcquireImage(i) }

func (i *ImageInfo) AcquireNextImage(image *Image) { AcquireNextImage(i, image) }

func (i *ImageInfo) GetImageInfoFile() *FILE { return GetImageInfoFile(i) }

func (i *ImageInfo) GetMagickProperty(image *Image, property string) string {
	return GetMagickProperty(i, image, property)
}

func (i *ImageInfo) PingImages(exception *ExceptionInfo) *Image { return PingImages(i, exception) }

func (i *ImageInfo) ReadImages(exception *ExceptionInfo) *Image { return ReadImages(i, exception) }

func (i *ImageInfo) SyncImageSettings(image *Image) bool { return SyncImageSettings(i, image) }

func (i *ImageInfo) SyncImagesSettings(image *Image) bool { return SyncImagesSettings(i, image) }

var AddChildToXMLTree func(x *XMLTreeInfo, tag string, offset uint32) *XMLTreeInfo

var AddPathToXMLTree func(x *XMLTreeInfo, path string, offset uint32) *XMLTreeInfo

var DestroyXMLTree func(x *XMLTreeInfo) *XMLTreeInfo

var GetNextXMLTreeTag func(x *XMLTreeInfo) *XMLTreeInfo

var GetXMLTreeAttribute func(x *XMLTreeInfo, tag string) string

var GetXMLTreeAttributes func(x *XMLTreeInfo, attributes *SplayTreeInfo) bool

var GetXMLTreeChild func(x *XMLTreeInfo, tag string) *XMLTreeInfo

var GetXMLTreeContent func(x *XMLTreeInfo) string

var GetXMLTreeOrdered func(x *XMLTreeInfo) *XMLTreeInfo

var GetXMLTreePath func(x *XMLTreeInfo, path string) *XMLTreeInfo

var GetXMLTreeProcessingInstructions func(x *XMLTreeInfo, target string) []string

var GetXMLTreeSibling func(x *XMLTreeInfo) *XMLTreeInfo

var GetXMLTreeTag func(x *XMLTreeInfo) string

var InsertTagIntoXMLTree func(x *XMLTreeInfo, child *XMLTreeInfo, offset uint32) *XMLTreeInfo

var PruneTagFromXMLTree func(x *XMLTreeInfo) *XMLTreeInfo

var SetXMLTreeAttribute func(x *XMLTreeInfo, tag, value string) *XMLTreeInfo

var SetXMLTreeContent func(x *XMLTreeInfo, content string) *XMLTreeInfo

var XMLTreeInfoToXML func(x *XMLTreeInfo) string

func (x *XMLTreeInfo) AddChild(tag string, offset uint32) *XMLTreeInfo {
	return AddChildToXMLTree(x, tag, offset)
}

func (x *XMLTreeInfo) AddPath(path string, offset uint32) *XMLTreeInfo {
	return AddPathToXMLTree(x, path, offset)
}

func (x *XMLTreeInfo) Destroy() *XMLTreeInfo { return DestroyXMLTree(x) }

func (x *XMLTreeInfo) NextTag() *XMLTreeInfo { return GetNextXMLTreeTag(x) }

func (x *XMLTreeInfo) Attribute(tag string) string { return GetXMLTreeAttribute(x, tag) }

func (x *XMLTreeInfo) Attributes(attributes *SplayTreeInfo) bool {
	return GetXMLTreeAttributes(x, attributes)
}

func (x *XMLTreeInfo) Child(tag string) *XMLTreeInfo { return GetXMLTreeChild(x, tag) }

func (x *XMLTreeInfo) Content() string { return GetXMLTreeContent(x) }

func (x *XMLTreeInfo) Ordered() *XMLTreeInfo { return GetXMLTreeOrdered(x) }

func (x *XMLTreeInfo) Path(path string) *XMLTreeInfo { return GetXMLTreePath(x, path) }

func (x *XMLTreeInfo) ProcessingInstructions(target string) []string {
	return GetXMLTreeProcessingInstructions(x, target)
}

func (x *XMLTreeInfo) Sibling() *XMLTreeInfo { return GetXMLTreeSibling(x) }

func (x *XMLTreeInfo) Tag() string { return GetXMLTreeTag(x) }

func (x *XMLTreeInfo) InsertTag(child *XMLTreeInfo, offset uint32) *XMLTreeInfo {
	return InsertTagIntoXMLTree(x, child, offset)
}

func (x *XMLTreeInfo) PruneTag() *XMLTreeInfo { return PruneTagFromXMLTree(x) }

func (x *XMLTreeInfo) SetAttribute(tag, value string) *XMLTreeInfo {
	return SetXMLTreeAttribute(x, tag, value)
}

func (x *XMLTreeInfo) SetContent(content string) *XMLTreeInfo {
	return SetXMLTreeContent(x, content)
}

func (x *XMLTreeInfo) XML() string { return XMLTreeInfoToXML(x) }

var NewLinkedList func(capacity T.Size) *LinkedListInfo

var DestroyConfigureOptions func(options *LinkedListInfo) *LinkedListInfo

var DestroyLocaleOptions func(messages *LinkedListInfo) *LinkedListInfo

var AppendValueToLinkedList func(l *LinkedListInfo, value *T.Void) bool

var ClearLinkedList func(l *LinkedListInfo, relinquishValue func(*T.Void) *T.Void)

var DestroyLinkedList func(l *LinkedListInfo, relinquishValue func(*T.Void) *T.Void)

var GetLastValueInLinkedList func(l *LinkedListInfo) *T.Void

var GetNextValueInLinkedList func(l *LinkedListInfo) *T.Void

var GetNumberOfElementsInLinkedList func(l *LinkedListInfo) T.Size

var GetValueFromLinkedList func(l *LinkedListInfo, index T.Size) *T.Void

var InsertValueInLinkedList func(l *LinkedListInfo, index T.Size, value *T.Void) bool

var InsertValueInSortedLinkedList func(l *LinkedListInfo, compare func(*T.Void, *T.Void) int, replace **T.Void, value *T.Void) bool

var IsLinkedListEmpty func(l *LinkedListInfo) bool

var LinkedListToArray func(l *LinkedListInfo, array **T.Void) bool

var RemoveElementByValueFromLinkedList func(l *LinkedListInfo, value *T.Void) *T.Void

var RemoveElementFromLinkedList func(l *LinkedListInfo, index T.Size) *T.Void

var RemoveLastElementFromLinkedList func(l *LinkedListInfo) *T.Void

var ResetLinkedListIterator func(l *LinkedListInfo)

func (l *LinkedListInfo) DestroyConfigureOptions() *LinkedListInfo { return DestroyConfigureOptions(l) }

func (l *LinkedListInfo) DestroyLocaleOptions() *LinkedListInfo { return DestroyLocaleOptions(l) }

func (l *LinkedListInfo) AppendValue(value *T.Void) bool {
	return AppendValueToLinkedList(l, value)
}

func (l *LinkedListInfo) Clear(relinquishValue func(*T.Void) *T.Void) {
	ClearLinkedList(l, relinquishValue)
}

func (l *LinkedListInfo) Destroy(relinquishValue func(*T.Void) *T.Void) {
	DestroyLinkedList(l, relinquishValue)
}

func (l *LinkedListInfo) LastValue() *T.Void { return GetLastValueInLinkedList(l) }

func (l *LinkedListInfo) NextValue() *T.Void { return GetNextValueInLinkedList(l) }

func (l *LinkedListInfo) NumberOfElements() T.Size { return GetNumberOfElementsInLinkedList(l) }

func (l *LinkedListInfo) Value(index T.Size) *T.Void { return GetValueFromLinkedList(l, index) }

func (l *LinkedListInfo) InsertValue(index T.Size, value *T.Void) bool {
	return InsertValueInLinkedList(l, index, value)
}

func (l *LinkedListInfo) InsertValueSorted(compare func(*T.Void, *T.Void) int, replace **T.Void, value *T.Void) bool {
	return InsertValueInSortedLinkedList(l, compare, replace, value)
}

func (l *LinkedListInfo) Empty() bool { return IsLinkedListEmpty(l) }

func (l *LinkedListInfo) Array(array **T.Void) bool { return LinkedListToArray(l, array) }

func (l *LinkedListInfo) RemoveElementByValue(value *T.Void) *T.Void {
	return RemoveElementByValueFromLinkedList(l, value)
}

func (l *LinkedListInfo) RemoveElement(index T.Size) *T.Void {
	return RemoveElementFromLinkedList(l, index)
}

func (l *LinkedListInfo) RemoveLastElement() *T.Void { return RemoveLastElementFromLinkedList(l) }

func (l *LinkedListInfo) ResetIterator() { ResetLinkedListIterator(l) }

var GetDelegateInfo func(decode, encode string, exception *ExceptionInfo) *DelegateInfo

var GetDelegateInfoList func(pattern string, numberDelegates *T.Size, exception *ExceptionInfo) **DelegateInfo

var GetDelegateCommands func(d *DelegateInfo) string

var GetDelegateMode func(d *DelegateInfo) T.Long

var GetDelegateThreadSupport func(d *DelegateInfo) bool

func (d *DelegateInfo) Commands() string { return GetDelegateCommands(d) }

func (d *DelegateInfo) Mode() T.Long { return GetDelegateMode(d) }

func (d *DelegateInfo) ThreadSupport() bool { return GetDelegateThreadSupport(d) }

var NewSplayTree func(compare func(*T.Void, *T.Void) int, relinquishKey, relinquishValue func(*T.Void) *T.Void) *SplayTreeInfo

var AddValueToSplayTree func(s *SplayTreeInfo, key, value *T.Void) bool

var CloneSplayTree func(s *SplayTreeInfo, cloneKey, cloneValue func(*T.Void) *T.Void) *SplayTreeInfo

var DeleteNodeByValueFromSplayTree func(s *SplayTreeInfo, value *T.Void) bool

var DeleteNodeFromSplayTree func(s *SplayTreeInfo, key *T.Void) bool

var DestroySplayTree func(s *SplayTreeInfo) *SplayTreeInfo

var GetNextKeyInSplayTree func(s *SplayTreeInfo) *T.Void

var GetNextValueInSplayTree func(s *SplayTreeInfo) *T.Void

var GetNumberOfNodesInSplayTree func(s *SplayTreeInfo) T.Size

var GetValueFromSplayTree func(s *SplayTreeInfo, key *T.Void) *T.Void

var RemoveNodeByValueFromSplayTree func(s *SplayTreeInfo, value *T.Void) *T.Void

var RemoveNodeFromSplayTree func(s *SplayTreeInfo, key *T.Void) *T.Void

var ResetSplayTreeIterator func(s *SplayTreeInfo)

func (s *SplayTreeInfo) AddValue(key, value *T.Void) bool { return AddValueToSplayTree(s, key, value) }

func (s *SplayTreeInfo) Clone(cloneKey, cloneValue func(*T.Void) *T.Void) *SplayTreeInfo {
	return CloneSplayTree(s, cloneKey, cloneValue)
}

func (s *SplayTreeInfo) DeleteNodeByValue(value *T.Void) bool {
	return DeleteNodeByValueFromSplayTree(s, value)
}

func (s *SplayTreeInfo) DeleteNode(key *T.Void) bool { return DeleteNodeFromSplayTree(s, key) }

func (s *SplayTreeInfo) Destroy() *SplayTreeInfo { return DestroySplayTree(s) }

func (s *SplayTreeInfo) NextKey() *T.Void { return GetNextKeyInSplayTree(s) }

func (s *SplayTreeInfo) NextValue() *T.Void { return GetNextValueInSplayTree(s) }

func (s *SplayTreeInfo) NumberOfNodes() T.Size { return GetNumberOfNodesInSplayTree(s) }

func (s *SplayTreeInfo) Value(key *T.Void) *T.Void { return GetValueFromSplayTree(s, key) }

func (s *SplayTreeInfo) RemoveNodeByValue(value *T.Void) *T.Void {
	return RemoveNodeByValueFromSplayTree(s, value)
}

func (s *SplayTreeInfo) RemoveNode(key *T.Void) *T.Void { return RemoveNodeFromSplayTree(s, key) }

func (s *SplayTreeInfo) ResetIterator() { ResetSplayTreeIterator(s) }

var DestroyHashmap func(h *HashmapInfo) *HashmapInfo

func (h *HashmapInfo) Destroy() *HashmapInfo { return DestroyHashmap(h) }

var GetNextKeyInHashmap func(h *HashmapInfo) *T.Void

func (h *HashmapInfo) NextKey() *T.Void { return GetNextKeyInHashmap(h) }

var GetNextValueInHashmap func(h *HashmapInfo) *T.Void

func (h *HashmapInfo) NextValue() *T.Void { return GetNextValueInHashmap(h) }

var GetNumberOfEntriesInHashmap func(h *HashmapInfo) T.Size

func (h *HashmapInfo) NumberOfEntries() T.Size { return GetNumberOfEntriesInHashmap(h) }

var GetValueFromHashmap func(h *HashmapInfo, key *T.Void) *T.Void

func (h *HashmapInfo) Value(key *T.Void) *T.Void { return GetValueFromHashmap(h, key) }

var IsHashmapEmpty func(h *HashmapInfo) bool

func (h *HashmapInfo) Empty() bool { return IsHashmapEmpty(h) }

var PutEntryInHashmap func(h *HashmapInfo, key, value *T.Void) bool

func (h *HashmapInfo) PutEntry(key, value *T.Void) bool { return PutEntryInHashmap(h, key, value) }

var RemoveEntryFromHashmap func(h *HashmapInfo, key *T.Void) *T.Void

func (h *HashmapInfo) RemoveEntry(key *T.Void) *T.Void { return RemoveEntryFromHashmap(h, key) }

var ResetHashmapIterator func(h *HashmapInfo)

func (h *HashmapInfo) ResetIterator() { ResetHashmapIterator(h) }

var CloneImageView func(i *ImageView) *ImageView

func (i *ImageView) Clone() *ImageView { return CloneImageView(i) }

var DestroyImageView func(i *ImageView) *ImageView

func (i *ImageView) Destroy() *ImageView { return DestroyImageView(i) }

var DuplexTransferImageViewIterator func(i *ImageView, duplex *ImageView, destination *ImageView, transfer T.DuplexTransferImageViewMethod, context *T.Void) bool

func (i *ImageView) DuplexTransferIterator(duplex *ImageView, destination *ImageView, transfer T.DuplexTransferImageViewMethod, context *T.Void) bool {
	return DuplexTransferImageViewIterator(i, duplex, destination, transfer, context)
}

var GetImageViewAuthenticIndexes func(i *ImageView) *T.IndexPacket

func (i *ImageView) AuthenticIndexes() *T.IndexPacket { return GetImageViewAuthenticIndexes(i) }

var GetImageViewAuthenticPixels func(i *ImageView) *T.PixelPacket

func (i *ImageView) AuthenticPixels() *T.PixelPacket { return GetImageViewAuthenticPixels(i) }

var GetImageViewExtent func(i *ImageView) *T.RectangleInfo // doc T.RectangleInfo
func (i *ImageView) Extent() *T.RectangleInfo              { return GetImageViewExtent(i) }

var GetImageViewImage func(i *ImageView) *Image

func (i *ImageView) Image() *Image { return GetImageViewImage(i) }

var GetImageViewIterator func(i *ImageView, get T.GetImageViewMethod, context *T.Void) bool

func (i *ImageView) Iterator(get T.GetImageViewMethod, context *T.Void) bool {
	return GetImageViewIterator(i, get, context)
}

var GetImageViewVirtualIndexes func(i *ImageView) *T.IndexPacket

func (i *ImageView) VirtualIndexes() *T.IndexPacket { return GetImageViewVirtualIndexes(i) }

var GetImageViewVirtualPixels func(i *ImageView) *T.PixelPacket

func (i *ImageView) VirtualPixels() *T.PixelPacket { return GetImageViewVirtualPixels(i) }

var IsImageView func(i *ImageView) bool

func (i *ImageView) IsImageView() bool { return IsImageView(i) }

var SetImageViewDescription func(i *ImageView, description string)

func (i *ImageView) SetDescription(description string) { SetImageViewDescription(i, description) }

var SetImageViewIterator func(i *ImageView, s SetImageViewMethod, context *T.Void) bool

func (i *ImageView) SetIterator(set SetImageViewMethod, context *T.Void) bool {
	return SetImageViewIterator(i, set, context)
}

var SetImageViewThreads func(i *ImageView, numberThreads uint32)

func (i *ImageView) SetThreads(numberThreads uint32) { SetImageViewThreads(i, numberThreads) }

var TransferImageViewIterator func(i *ImageView, destination *ImageView, transfer T.TransferImageViewMethod, context *T.Void) bool

func (i *ImageView) TransferIterator(destination *ImageView, transfer T.TransferImageViewMethod, context *T.Void) bool {
	return TransferImageViewIterator(i, destination, transfer, context)
}

var UpdateImageViewIterator func(i *ImageView, update T.UpdateImageViewMethod, context *T.Void) bool

func (i *ImageView) UpdateIterator(update T.UpdateImageViewMethod, context *T.Void) bool {
	return UpdateImageViewIterator(i, update, context)
}

var CloneQuantizeInfo func(q *QuantizeInfo) *QuantizeInfo

func (q *QuantizeInfo) Clone() *QuantizeInfo { return CloneQuantizeInfo(q) }

var DestroyQuantizeInfo func(q *QuantizeInfo) *QuantizeInfo

func (q *QuantizeInfo) Destroy() *QuantizeInfo { return DestroyQuantizeInfo(q) }

var GetQuantizeInfo func(q *QuantizeInfo)

func (q *QuantizeInfo) Get() { GetQuantizeInfo(q) }

var QuantizeImage func(q *QuantizeInfo, image *Image) bool

func (q *QuantizeInfo) QuantizeImage(image *Image) bool { return QuantizeImage(q, image) }

var QuantizeImages func(q *QuantizeInfo, images *Image) bool

func (q *QuantizeInfo) QuantizeImages(images *Image) bool { return QuantizeImages(q, images) }

var AffinityImage func(q *QuantizeInfo, image *Image, affinityImage *Image) bool

func (q *QuantizeInfo) AffinityImage(image *Image, affinityImage *Image) bool {
	return AffinityImage(q, image, affinityImage)
}

var AffinityImages func(q *QuantizeInfo, images *Image, affinityImage *Image) bool

func (q *QuantizeInfo) AffinityImages(images *Image, affinityImage *Image) bool {
	return AffinityImages(q, images, affinityImage)
}

var RemapImage func(q *QuantizeInfo, image *Image, remapImage *Image) bool

func (q *QuantizeInfo) RemapImage(image *Image, remapImage *Image) bool {
	return RemapImage(q, image, remapImage)
}

var RemapImages func(q *QuantizeInfo, images *Image, remapImage *Image) bool

func (q *QuantizeInfo) RemapImages(images *Image, remapImage *Image) bool {
	return RemapImages(q, images, remapImage)
}

var DestroyStreamInfo func(s *StreamInfo) *StreamInfo

func (s *StreamInfo) DestroyStreamInfo() *StreamInfo { return DestroyStreamInfo(s) }

var GetStreamInfoClientData func(s *StreamInfo) *T.Void

func (s *StreamInfo) GetStreamInfoClientData() *T.Void { return GetStreamInfoClientData(s) }

var SetStreamInfoClientData func(s *StreamInfo, clientData *T.Void)

func (s *StreamInfo) SetStreamInfoClientData(clientData *T.Void) {
	SetStreamInfoClientData(s, clientData)
}

var SetStreamInfoMap func(s *StreamInfo, map_ string)

func (s *StreamInfo) SetStreamInfoMap(map_ string) { SetStreamInfoMap(s, map_) }

var SetStreamInfoStorageType func(s *StreamInfo, storageType T.StorageType)

func (s *StreamInfo) SetStreamInfoStorageType(storageType T.StorageType) {
	SetStreamInfoStorageType(s, storageType)
}

var CloneStringInfo func(s *StringInfo) *StringInfo

func (s *StringInfo) Clone() *StringInfo { return CloneStringInfo(s) }

var CompareStringInfo func(s *StringInfo, source *StringInfo) int

func (s *StringInfo) Compare(source *StringInfo) int { return CompareStringInfo(s, source) }

var ConcatenateStringInfo func(s *StringInfo, source *StringInfo)

func (s *StringInfo) Concatenate(source *StringInfo) { ConcatenateStringInfo(s, source) }

var DestroyStringInfo func(s *StringInfo) *StringInfo

func (s *StringInfo) Destroy() *StringInfo { return DestroyStringInfo(s) }

var GetStringInfoDatum func(s *StringInfo) *byte

func (s *StringInfo) Datum() *byte { return GetStringInfoDatum(s) }

var GetStringInfoLength func(s *StringInfo) uint32

func (s *StringInfo) Length() uint32 { return GetStringInfoLength(s) }

var GetStringInfoPath func(s *StringInfo) string

func (s *StringInfo) Path() string { return GetStringInfoPath(s) }

var ResetStringInfo func(s *StringInfo)

func (s *StringInfo) Reset() { ResetStringInfo(s) }

var SetStringInfo func(s *StringInfo, source *StringInfo)

func (s *StringInfo) Set(source *StringInfo) { SetStringInfo(s, source) }

var SetStringInfoDatum func(s *StringInfo, source *byte)

func (s *StringInfo) SetDatum(source *byte) { SetStringInfoDatum(s, source) }

var SetStringInfoLength func(s *StringInfo, length uint32)

func (s *StringInfo) SetLength(length uint32) { SetStringInfoLength(s, length) }

var SetStringInfoPath func(s *StringInfo, path string)

func (s *StringInfo) SetPath(path string) { SetStringInfoPath(s, path) }

var SplitStringInfo func(s *StringInfo, offset uint32) *StringInfo

func (s *StringInfo) Split(offset uint32) *StringInfo { return SplitStringInfo(s, offset) }

var StringInfoToString func(s *StringInfo) string

func (s *StringInfo) String() string { return StringInfoToString(s) }

var ListCoderInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) CoderInfo(exception *ExceptionInfo) bool { return ListCoderInfo(f, exception) }

var ListColorInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) ColorInfo(exception *ExceptionInfo) bool { return ListColorInfo(f, exception) }

var ListConfigureInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) ConfigureInfo(exception *ExceptionInfo) bool { return ListConfigureInfo(f, exception) }

var ListDelegateInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) DelegateInfo(exception *ExceptionInfo) bool { return ListDelegateInfo(f, exception) }

var ListLocaleInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) LocaleInfo(exception *ExceptionInfo) bool { return ListLocaleInfo(f, exception) }

var ListLogInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) LogInfo(exception *ExceptionInfo) bool { return ListLogInfo(f, exception) }

var ListMagicInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) MagicInfo(exception *ExceptionInfo) bool { return ListMagicInfo(f, exception) }

var ListMagickInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) MagickInfo(exception *ExceptionInfo) bool { return ListMagickInfo(f, exception) }

var ListCommandOptions func(f *FILE, option CommandOption, exception *ExceptionInfo) bool

func (f *FILE) CommandOptions(option CommandOption, exception *ExceptionInfo) bool {
	return ListCommandOptions(f, option, exception)
}

var ListMagickResourceInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) MagickResourceInfo(exception *ExceptionInfo) bool {
	return ListMagickResourceInfo(f, exception)
}

var ListMimeInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) MimeInfo(exception *ExceptionInfo) bool { return ListMimeInfo(f, exception) }

var ListModuleInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) ModuleInfo(exception *ExceptionInfo) bool { return ListModuleInfo(f, exception) }

var ListThresholdMaps func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) ThresholdMaps(exception *ExceptionInfo) bool { return ListThresholdMaps(f, exception) }

var ListTypeInfo func(f *FILE, exception *ExceptionInfo) bool

func (f *FILE) TypeInfo(exception *ExceptionInfo) bool { return ListTypeInfo(f, exception) }

var AcquireExceptionInfo func() *ExceptionInfo

var CatchException func(e *ExceptionInfo)

func (e *ExceptionInfo) Catch() { CatchException(e) }

var ClearMagickException func(e *ExceptionInfo)

func (e *ExceptionInfo) Clear() { ClearMagickException(e) }

var DestroyExceptionInfo func(e *ExceptionInfo) *ExceptionInfo

func (e *ExceptionInfo) Destroy() *ExceptionInfo { return DestroyExceptionInfo(e) }

var GetExceptionInfo func(e *ExceptionInfo)

func (e *ExceptionInfo) Get() { GetExceptionInfo(e) }

var InheritException func(e *ExceptionInfo, relative *ExceptionInfo)

func (e *ExceptionInfo) Inherit(relative *ExceptionInfo) { InheritException(e, relative) }

var SetExceptionInfo func(e *ExceptionInfo, severity T.ExceptionType) bool

func (e *ExceptionInfo) Set(severity T.ExceptionType) bool { return SetExceptionInfo(e, severity) }

var ThrowException func(e *ExceptionInfo, severity T.ExceptionType, reason, description string) bool

func (e *ExceptionInfo) Throw(severity T.ExceptionType, reason, description string) bool {
	return ThrowException(e, severity, reason, description)
}

var ThrowMagickException func(e *ExceptionInfo, module, function string, line T.Size, severity T.ExceptionType, tag, format string, va ...VArg) bool

func (e *ExceptionInfo) ThrowMagick(module, function string, line T.Size, severity T.ExceptionType, tag, format string, va ...VArg) bool {
	return ThrowMagickException(e, module, function, line, severity, tag, format, va)
}

var ThrowMagickExceptionList func(e *ExceptionInfo, module, function string, line T.Size, severity T.ExceptionType, tag, format string, operands VAList) bool

func (e *ExceptionInfo) ThrowMagickList(module, function string, line T.Size, severity T.ExceptionType, tag, format string, operands VAList) bool {
	return ThrowMagickExceptionList(e, module, function, line, severity, tag, format, operands)
}

var XGetResourceClass func(x XrmDatabase, clientName, keyword, resourceDefault string) string

func (x XrmDatabase) ResourceClass(clientName, keyword, resourceDefault string) string {
	return XGetResourceClass(x, clientName, keyword, resourceDefault)
}

var XGetResourceInfo func(x XrmDatabase, clientName string, resourceInfo *T.XResourceInfo)

func (x XrmDatabase) ResourceInfo(clientName string, resourceInfo *T.XResourceInfo) {
	XGetResourceInfo(x, clientName, resourceInfo)
}

var XGetResourceInstance func(x XrmDatabase, clientName, keyword, resourceDefault string) string

func (x XrmDatabase) ResourceInstance(clientName, keyword, resourceDefault string) string {
	return XGetResourceInstance(x, clientName, keyword, resourceDefault)
}

var XrmCombineDatabase func(x XrmDatabase, target *XrmDatabase, override bool)

func (x XrmDatabase) Combine(target *XrmDatabase, override bool) {
	XrmCombineDatabase(x, target, override)
}

var AcquireMagickResource func(r ResourceType, size T.MagickSizeType) bool

func (r ResourceType) AcquireResource(size T.MagickSizeType) bool {
	return AcquireMagickResource(r, size)
}

var GetMagickResource func(r ResourceType) T.MagickSizeType

func (r ResourceType) Resource() T.MagickSizeType { return GetMagickResource(r) }

var GetMagickResourceLimit func(r ResourceType) T.MagickSizeType

func (r ResourceType) Limit() T.MagickSizeType { return GetMagickResourceLimit(r) }

var RelinquishMagickResource func(r ResourceType, size T.MagickSizeType)

func (r ResourceType) Relinquish(size T.MagickSizeType) { RelinquishMagickResource(r, size) }

var SetMagickResourceLimit func(r ResourceType, limit T.MagickSizeType) bool

func (r ResourceType) SetLimit(limit T.MagickSizeType) bool { return SetMagickResourceLimit(r, limit) }

var AttachBlob func(b *BlobInfo, blob *T.Void, length uint32)

func (b *BlobInfo) Attach(blob *T.Void, length uint32) { AttachBlob(b, blob, length) }

var CloneBlobInfo func(b *BlobInfo) *BlobInfo

func (b *BlobInfo) Clone() *BlobInfo { return CloneBlobInfo(b) }

var DetachBlob func(b *BlobInfo) *byte

func (b *BlobInfo) Detach() *byte { return DetachBlob(b) }

var GetBlobInfo func(b *BlobInfo)

func (b *BlobInfo) Get() { GetBlobInfo(b) }

var ReferenceBlob func(b *BlobInfo) *BlobInfo

func (b *BlobInfo) Reference() *BlobInfo { return ReferenceBlob(b) }

var ConcatenateColorComponent func(m *MagickPixelPacket, channel T.ChannelType, compliance T.ComplianceType, tuple string)

func (m *MagickPixelPacket) ConcatenateColorComponent(channel T.ChannelType, compliance T.ComplianceType, tuple string) {
	ConcatenateColorComponent(m, channel, compliance, tuple)
}

var GetColorTuple func(m *MagickPixelPacket, hex bool, tuple string)

func (m *MagickPixelPacket) ColorTuple(hex bool, tuple string) { GetColorTuple(m, hex, tuple) }

var IsMagickColorSimilar func(p *MagickPixelPacket, q *MagickPixelPacket) bool

func (p *MagickPixelPacket) ColorSimilar(q *MagickPixelPacket) bool { return IsMagickColorSimilar(p, q) }

var CloneMagickPixelPacket func(m *MagickPixelPacket) *MagickPixelPacket // doc c-static inline
func (m *MagickPixelPacket) Clone() *MagickPixelPacket                   { return CloneMagickPixelPacket(m) }

var AcquireTimerInfo func() *TimerInfo

var ContinueTimer func(t *TimerInfo) bool

func (t *TimerInfo) Continue() bool { return ContinueTimer(t) }

var DestroyTimerInfo func(t *TimerInfo) *TimerInfo

func (t *TimerInfo) Destroy() *TimerInfo { return DestroyTimerInfo(t) }

var GetElapsedTime func(t *TimerInfo) float64

func (t *TimerInfo) ElapsedTime() float64 { return GetElapsedTime(t) }

var GetTimerInfo func(t *TimerInfo)

func (t *TimerInfo) GetTimerInfo() { GetTimerInfo(t) }

var GetUserTime func(t *TimerInfo) float64

func (t *TimerInfo) UserTime() float64 { return GetUserTime(t) }

var ResetTimer func(t *TimerInfo)

func (t *TimerInfo) Reset() { ResetTimer(t) }

var StartTimer func(t *TimerInfo, reset bool)

func (t *TimerInfo) Start(reset bool) { StartTimer(t, reset) }

var AcquireSemaphoreInfo func(s **SemaphoreInfo)

var AllocateSemaphoreInfo func() *SemaphoreInfo

var LiberateSemaphoreInfo func(s **SemaphoreInfo)

var DestroySemaphoreInfo func(s *SemaphoreInfo) *SemaphoreInfo

func (s *SemaphoreInfo) Destroy() *SemaphoreInfo { return DestroySemaphoreInfo(s) }

var LockSemaphoreInfo func(s *SemaphoreInfo) bool

func (s *SemaphoreInfo) Lock() bool { return LockSemaphoreInfo(s) }

var RelinquishSemaphoreInfo func(s *SemaphoreInfo)

func (s *SemaphoreInfo) Relinquish() { RelinquishSemaphoreInfo(s) }

var UnlockSemaphoreInfo func(s *SemaphoreInfo) bool

func (s *SemaphoreInfo) Unlock() bool { return UnlockSemaphoreInfo(s) }

var DestroyResizeFilter func(r *ResizeFilter) *ResizeFilter

func (r *ResizeFilter) Destroy() *ResizeFilter { return DestroyResizeFilter(r) }

var GetResizeFilterSupport func(r *ResizeFilter) T.MagickRealType

func (r *ResizeFilter) Support() T.MagickRealType { return GetResizeFilterSupport(r) }

var GetResizeFilterWeight func(r *ResizeFilter, x T.MagickRealType) T.MagickRealType

func (r *ResizeFilter) Weight(x T.MagickRealType) T.MagickRealType { return GetResizeFilterWeight(r, x) }

var SetResizeFilterSupport func(r *ResizeFilter, support T.MagickRealType)

func (r *ResizeFilter) SetSupport(support T.MagickRealType) { SetResizeFilterSupport(r, support) }

var DestroyResampleFilter func(r *ResampleFilter) *ResampleFilter

var ResamplePixelColor func(r *ResampleFilter, u0, v0 float64) MagickPixelPacket

var ScaleResampleFilter func(r *ResampleFilter, dux, duy, dvx, dvy float64)

var SetResampleFilterInterpolateMethod func(r *ResampleFilter, method T.InterpolatePixelMethod) bool

var SetResampleFilterVirtualPixelMethod func(r *ResampleFilter, method T.VirtualPixelMethod) bool

func (r *ResampleFilter) Destroy() *ResampleFilter { return DestroyResampleFilter(r) }

func (r *ResampleFilter) PixelColor(u0, v0 float64) MagickPixelPacket {
	return ResamplePixelColor(r, u0, v0)
}

func (r *ResampleFilter) Scale(dux, duy, dvx, dvy float64) {
	ScaleResampleFilter(r, dux, duy, dvx, dvy)
}

func (r *ResampleFilter) InterpolateMethod(method T.InterpolatePixelMethod) bool {
	return SetResampleFilterInterpolateMethod(r, method)
}

func (r *ResampleFilter) VirtualPixelMethod(method T.VirtualPixelMethod) bool {
	return SetResampleFilterVirtualPixelMethod(r, method)
}

var AcquireDrawInfo func() *DrawInfo

var DestroyDrawInfo func(d *DrawInfo) *DrawInfo

func (d *DrawInfo) Destroy() *DrawInfo { return DestroyDrawInfo(d) }

var AcquireTokenInfo func() *TokenInfo

var DestroyTokenInfo func(t *TokenInfo) *TokenInfo

var Tokenizer func(t *TokenInfo, flag uint, token string, maxTokenLength uint32, line, white, breakSet, quote string, escape int8, breaker string, next *int, quoted string) int

func (t *TokenInfo) Destroy() *TokenInfo { return DestroyTokenInfo(t) }

func (t *TokenInfo) Tokenizer(flag uint, token string, maxTokenLength uint32, line, white, breakSet, quote string, escape int8, breaker string, next *int, quoted string) int {
	return Tokenizer(t, flag, token, maxTokenLength, line, white, breakSet, quote, escape, breaker, next, quoted)
}

var AcquireSignatureInfo func() *SignatureInfo

var DestroySignatureInfo func(s *SignatureInfo) *SignatureInfo

var FinalizeSignature func(s *SignatureInfo)

var GetSignatureInfo func(s *SignatureInfo)

var UpdateSignature func(s *SignatureInfo, message *byte, length uint32)

func (s *SignatureInfo) Destroy() *SignatureInfo { return DestroySignatureInfo(s) }

func (s *SignatureInfo) Finalize() { FinalizeSignature(s) }

func (s *SignatureInfo) Get() { GetSignatureInfo(s) }

func (s *SignatureInfo) Update(message *byte, length uint32) { UpdateSignature(s, message, length) }

var DestroyFxInfo func(f *FxInfo) *FxInfo

var FxEvaluateChannelExpression func(f *FxInfo, channel T.ChannelType, x, y T.Long, alpha *T.MagickRealType, exception *ExceptionInfo) bool

var FxEvaluateExpression func(f *FxInfo, alpha *T.MagickRealType, exception *ExceptionInfo) bool

func (f *FxInfo) Destroy() *FxInfo { return DestroyFxInfo(f) }

func (f *FxInfo) EvaluateChannelExpression(channel T.ChannelType, x, y T.Long, alpha *T.MagickRealType, exception *ExceptionInfo) bool {
	return FxEvaluateChannelExpression(f, channel, x, y, alpha, exception)
}

func (f *FxInfo) EvaluateExpression(alpha *T.MagickRealType, exception *ExceptionInfo) bool {
	return FxEvaluateExpression(f, alpha, exception)
}

var GetCommandOptions func(value CommandOption) []string

var CommandOptionToMnemonic func(option CommandOption, type_ T.Long) string

var ParseCommandOption func(option CommandOption, list bool, options string) T.Long

func (m CommandOption) Options() []string { return GetCommandOptions(m) }

func (m CommandOption) Mnemonic(type_ T.Long) string { return CommandOptionToMnemonic(m, type_) }

func (m CommandOption) Parse(list bool, options string) T.Long {
	return ParseCommandOption(m, list, options)
}

var AcquirePixelCachePixels /*IM*/ func(i *Image, m *T.MagickSizeType, e *ExceptionInfo) *T.Void

var AcquireRandomInfo func() *T.RandomInfo

var DestroyRandomInfo func(*T.RandomInfo) *T.RandomInfo

var AnnotateComponentGenesis func() bool

var AnnotateComponentTerminus func()

var AsynchronousResourceComponentTerminus func()

var BlackThresholdImageChannel func(*Image, T.ChannelType, string, *ExceptionInfo) bool

var BlobToStringInfo func(*T.Void, uint32) *StringInfo

var CacheComponentGenesis func() bool

var CacheComponentTerminus func()

var ClampImage func(*Image) bool

var ClampImageChannel func(*Image, T.ChannelType) bool

var ClonePixelCache func(Cache) Cache

var ClonePixelCacheMethods func(Cache, Cache)

var CoderComponentGenesis func() bool

var CoderComponentTerminus func()

var ColorComponentGenesis func() bool

var ColorComponentTerminus func()

var ConfigureComponentGenesis func() bool

var ConfigureComponentTerminus func()

var ConstituteComponentGenesis func() bool

var ConstituteComponentTerminus func()

var ConvertHCLToRGB func(float64, float64, float64, *T.Quantum, *T.Quantum, *T.Quantum)

var ConvertRGBToHCL func(T.Quantum, T.Quantum, T.Quantum, *float64, *float64, *float64)

var DelegateComponentGenesis func() bool

var DelegateComponentTerminus func()

var DestroyPixelCacheNexus func(**T.NexusInfo, uint32) **T.NexusInfo

var DiscardBlobBytes func(*Image, T.MagickSizeType) bool

var DistortResizeImage func(*Image, uint32, uint32, *ExceptionInfo) *Image

var DuplicateBlob func(*Image, *Image)

var FormatLocaleFile func(*FILE, string, ...VArg) int32

var FormatLocaleFileList func(*FILE, string, VAList) int32

var FormatLocaleString func(string, uint32, string, ...VArg) int32

var FormatLocaleStringList func(string, uint32, string, VAList) int32

var GenerateDifferentialNoise func(*T.RandomInfo, T.Quantum, T.NoiseType, T.MagickRealType) float64

var GetAuthenticPixelCacheNexus func(*Image, int32, int32, uint32, uint32, *T.NexusInfo, *ExceptionInfo) *T.PixelPacket

var GetCacheViewExtent func(*T.CacheView) T.MagickSizeType

var GetCommandOptionFlags func(CommandOption, bool, string) int32

var GetConfigureOption func(string) string

var GetImageExtent func(*Image) T.MagickSizeType

var GetImageKurtosis func(*Image, *float64, *float64, *ExceptionInfo) bool

var GetImagePixelCacheType func(*Image) T.CacheType

var GetImageReferenceCount func(*Image) int32

var GetMagickPageSize func() int32

var GetMagickRawSupport func(*MagickInfo) bool

var GetPathAttributes func(string, *T.Void) bool

var GetPixelCacheChannels func(Cache) uint32

var GetPixelCacheColorspace func(Cache) T.ColorspaceType

var GetPixelCacheMethods func(*T.CacheMethods)

var GetPixelCacheNexusExtent func(Cache, *T.NexusInfo) T.MagickSizeType

var GetPixelCachePixels func(*Image, *T.MagickSizeType, *ExceptionInfo) *T.Void

var GetPixelCacheStorageClass func(Cache) T.ClassType

var GetPixelCacheTileSize func(*Image, *uint32, *uint32)

var GetPixelCacheType func(*Image) T.CacheType

var GetPixelCacheVirtualMethod func(*Image) T.VirtualPixelMethod

var GetPolicyInfoList func(string, *uint32, *ExceptionInfo) []*PolicyInfo

var GetPolicyList func(string, *uint32, *ExceptionInfo) []string

var GetPolicyValue func(string) string

var GetPseudoRandomValue func(*T.RandomInfo) float64

var GetQuantumEndian func(*T.QuantumInfo) T.EndianType

var GetQuantumExtent func(*Image, *T.QuantumInfo, T.QuantumType) uint32

var GetQuantumFormat func(*T.QuantumInfo) T.QuantumFormatType

var GetQuantumPixels func(*T.QuantumInfo) *byte

var GetQuantumType func(*Image, *ExceptionInfo) T.QuantumType

var GetRandomSecretKey func(*T.RandomInfo) T.Size

var GetSignatureBlocksize func(*SignatureInfo) uint

var GetSignatureDigest func(*SignatureInfo) *StringInfo

var GetSignatureDigestsize func(*SignatureInfo) uint

var Gettimeofday func(*timeval, *timezone)

var GetVirtualIndexesFromNexus func(Cache, *T.NexusInfo) *T.IndexPacket

var GetVirtualPixelsFromNexus func(*Image, T.VirtualPixelMethod, int32, int32, uint32, uint32, *T.NexusInfo, *ExceptionInfo) *T.PixelPacket

var GetVirtualPixelsNexus func(Cache, *T.NexusInfo) *T.PixelPacket

var InitializeSignature func(*SignatureInfo)

var InterpolateMagickPixelPacket func(*Image, *CacheView, T.InterpolatePixelMethod, float64, float64, *MagickPixelPacket, *ExceptionInfo) bool

var InterpolativeResizeImage func(*Image, uint32, uint32, T.InterpolatePixelMethod, *ExceptionInfo) *Image

var InterpretLocaleValue func(string, **char) float64

var InterpretSiPrefixValue func(string, **char) float64

var InversesRGBCompandor func(T.MagickRealType) T.MagickRealType

var IsCommandOption func(string) bool

var IsPathAccessible func(string) bool

var IsRightsAuthorized func(T.PolicyDomain, T.PolicyRights, string) bool

var IsStringNotFalse func(string) bool

var IsStringTrue func(string) bool

var ListPolicyInfo func(*FILE, *ExceptionInfo) bool

var LocaleComponentGenesis func() bool

var LocaleComponentTerminus func()

var LogComponentGenesis func() bool

var LogComponentTerminus func()

var MagicComponentGenesis func() bool

var MagicComponentTerminus func()

var MagickComponentGenesis func() bool

var MagickComponentTerminus func()

var MagickCreateThreadKey func(*MagickThreadKey) bool

var MagickDelay func(T.MagickSizeType)

var MagickDeleteThreadKey func(MagickThreadKey) bool

var MagickGetThreadValue func(MagickThreadKey) *T.Void

var MagickSetThreadValue func(MagickThreadKey, *T.Void) bool

var MimeComponentGenesis func() bool

var MimeComponentTerminus func()

var ModuleComponentGenesis func() bool

var ModuleComponentTerminus func()

var NTArgvToUTF8 func(argc int, argv []*wchar_t) []string

var NTGatherRandomData func(uint32, *byte) bool

var ParseRegionGeometry func(*Image, string, *T.RectangleInfo, *ExceptionInfo) T.MagickStatusType

var PerceptibleImage func(*Image, float64) bool

var PerceptibleImageChannel func(*Image, T.ChannelType, float64) bool

var PersistPixelCache func(*Image, string, bool, *T.MagickOffsetType, *ExceptionInfo) bool

var PolicyComponentGenesis func() bool

var PolicyComponentTerminus func()

var PolynomialImage func(*Image, uint32, *float64, *ExceptionInfo) *Image

var PolynomialImageChannel func(*Image, T.ChannelType, uint32, *float64, *ExceptionInfo) *Image

var QueryMagickColorCompliance func(string, T.ComplianceType, *MagickPixelPacket, *ExceptionInfo) bool

var QueueAuthenticPixel func(*Image, int32, int32, uint32, uint32, bool, *T.NexusInfo, *ExceptionInfo) *T.PixelPacket

var QueueAuthenticPixelCacheNexus func(*Image, int32, int32, uint32, uint32, bool, *T.NexusInfo, *ExceptionInfo) *T.PixelPacket

var RandomComponentGenesis func() bool

var RandomComponentTerminus func()

var ReadBlobMSBLongLong func(*Image) T.MagickSizeType

var ReferencePixelCache func(Cache) Cache

var RegistryComponentGenesis func() bool

var ResetImageOptions func(*ImageInfo)

var ResetSplayTree func(*SplayTreeInfo)

var ResourceComponentGenesis func() bool

var ResourceComponentTerminus func()

var SeedPseudoRandomGenerator func(T.Size)

var SemaphoreComponentGenesis func() bool

var SemaphoreComponentTerminus func()

var SetBlobExtent func(*Image, T.MagickSizeType) bool

var SetPixelCacheMethods func(Cache, *T.CacheMethods)

var SetPixelCacheVirtualMethod func(*Image, T.VirtualPixelMethod) T.VirtualPixelMethod

var SetQuantumAlphaType func(*T.QuantumInfo, QuantumAlphaType)

var SetQuantumDepth func(*Image, *T.QuantumInfo, uint32) bool

var SetQuantumEndian func(*Image, *T.QuantumInfo, T.EndianType) bool

var SetQuantumFormat func(*Image, *T.QuantumInfo, T.QuantumFormatType) bool

var SetQuantumImageType func(*Image, T.QuantumType)

var SetQuantumMinIsWhite func(*T.QuantumInfo, bool)

var SetQuantumPack func(*T.QuantumInfo, bool)

var SetQuantumPad func(*Image, *T.QuantumInfo, uint32) bool

var SetQuantumQuantum func(*T.QuantumInfo, uint32)

var SetQuantumScale func(*T.QuantumInfo, float64)

var SetRandomKey func(*T.RandomInfo, uint32, *byte)

var SetRandomSecretKey func(T.Size)

var SetRandomTrueRandom func(bool)

var SetResampleFilter func(*ResampleFilter, T.FilterTypes, float64)

var SetSignatureDigest func(*SignatureInfo, *StringInfo)

var SimilarityMetricImage func(*Image, *Image, T.MetricType, *T.RectangleInfo, *float64, *ExceptionInfo) *Image

var SolarizeImageChannel func(*Image, T.ChannelType, float64, *ExceptionInfo) bool

var SRGBCompandor func(T.MagickRealType) T.MagickRealType

var StringInfoToHexString func(*StringInfo) string

var StringToArrayOfDoubles func(string, *int32, *ExceptionInfo) []float64

var SyncAuthenticPixelCacheNexus func(*Image, *T.NexusInfo, *ExceptionInfo) bool

var TransparentPaintImageChroma func(*Image, *T.MagickPixelPacket, *T.MagickPixelPacket, T.Quantum, bool) bool

var TypeComponentGenesis func() bool

var TypeComponentTerminus func()

var UnityAddKernelInfo func(*T.KernelInfo, float64)

var WhiteThresholdImageChannel func(*Image, T.ChannelType, string, *ExceptionInfo) bool

var WriteBlobMSBLongLong func(*Image, T.MagickSizeType) int32

var XComponentGenesis func() bool

var XComponentTerminus func()

var AcquireMagickMatrix func(nptrs, size T.Size) **float64

var AcquireMagickMemory func(size uint32) *T.Void

var AcquireMemory func(size uint32) *T.Void

var AcquireQuantumMemory func(count, quantum uint32) *T.Void

var AcquireString func(source string) string

var AcquireStringInfo func(length uint32) *StringInfo

var AcquireUniqueFilename func(path string) bool

var AcquireUniqueFileResource func(path string) int

var AcquireUniqueSymbolicLink func(source, destination string) bool

var AllocateString func(source string) string

var AppendImageFormat func(format, filename string)

var Base64Decode func(source string, length *uint32) *byte

var Base64Encode func(blob *byte, blobLength uint32, encodeLength *uint32) string

var BlobToFile func(filename string, blob *T.Void, length uint32, exception *ExceptionInfo) bool

var CanonicalXMLContent func(content string, pedantic bool) string

var ChopPathComponents func(path string, components T.Size)

var CloneMemory func(destination, source *T.Void, size uint32) *T.Void

var CloneString func(destination []string, source string) string

var CloseMagickLog func()

var CompareHashmapString func(target, source *T.Void) bool

var CompareHashmapStringInfo func(target, source *T.Void) bool

var CompareSplayTreeString func(target, source *T.Void) int

var CompareSplayTreeStringInfo func(target, source *T.Void) int

var ConcatenateMagickString func(destination, source string, length uint32) uint32

var ConcatenateString func(destination []string, source string) bool

var ConfigureFileToStringInfo func(filename string) *StringInfo

var ConstantString func(source string) string

var ConstituteImage func(columns, rows T.Size, map_ string, storage T.StorageType, pixels *T.Void, exception *ExceptionInfo) *Image

var ConvertHSBToRGB func(hue, saturation, brightness float64, red, green, blue *T.Quantum)

var ConvertHSLToRGB func(hue, saturation, luminosity float64, red, green, blue *T.Quantum)

var ConvertHWBToRGB func(hue, whiteness, blackness float64, red, green, blue *T.Quantum)

var ConvertRGBToHSB func(red, green, blue T.Quantum, hue, saturation, brightness *float64)

var ConvertRGBToHSL func(red, green, blue T.Quantum, hue, saturation, luminosity *float64)

var ConvertRGBToHWB func(red, green, blue T.Quantum, hue, whiteness, blackness *float64)

var CopyMagickMemory func(destination, source *T.Void, size uint32) *T.Void

var CopyMagickString func(destination, source string, length uint32) uint32

var DefineImageRegistry func(type_ T.RegistryType, option string, exception *ExceptionInfo) bool

var DeleteImageRegistry func(key string) bool

var DeleteMagickRegistry func(id T.Long) bool

var DestroyConstitute func()

var DestroyMagick func()

var DestroyMagickMemory func()

var DestroyMagickRegistry func()

var DestroyModuleList func()

var DestroyMontageInfo func(montageInfo *T.MontageInfo) *T.MontageInfo

var DestroyQuantumInfo func(quantumInfo *T.QuantumInfo) *T.QuantumInfo

var DestroyString func(str string) string

var DestroyStringList func(list []string) []string

var DestroyThresholdMap func(map_ *T.ThresholdMap) *T.ThresholdMap

var DestroyXResources func()

var DestroyXWidget func()

var EscapeString func(source string, escape int8) string

var ExpandAffine func(affine *T.AffineMatrix) float64

var ExpandFilename func(path string)

var ExpandFilenames func(argc *int, argv *[]string) bool

var FileToBlob func(filename string, extent uint32, length *uint32, exception *ExceptionInfo) *byte

var FileToString func(filename string, extent uint32, exception *ExceptionInfo) string

var FileToStringInfo func(filename string, extent uint32, exception *ExceptionInfo) *StringInfo

var FormatMagickSize func(size T.MagickSizeType, format string) T.Long

var FormatMagickString func(str string, length uint32, format string, va ...VArg) T.Long

var FormatMagickStringList func(str string, length uint32, format string, operands VAList) T.Long

var FormatMagickTime func(time time_t, length uint32, timestamp string) T.Long

var FormatString func(str, format string, va ...VArg)

var FormatStringList func(str, format string, operands VAList)

var FuzzyColorMatch func(p, q *T.PixelPacket, fuzz float64) uint

var GaussJordanElimination func(matrix, vectors **float64, rank, nvecs T.Size) bool

var GetAffineMatrix func(affineMatrix *T.AffineMatrix)

var GetClientName func() string

var GetClientPath func() string

var GetCoderInfo func(name string, exception *ExceptionInfo) *T.CoderInfo

var GetCoderInfoList func(pattern string, numberCoders *T.Size, exception *ExceptionInfo) **T.CoderInfo

var GetCoderList func(pattern string, numberCoders *T.Size, exception *ExceptionInfo) []string

var GetColorInfo func(name string, exception *ExceptionInfo) *T.ColorInfo

var GetColorInfoList func(pattern string, numberColors *T.Size, exception *ExceptionInfo) **T.ColorInfo

var GetColorList func(pattern string, numberColors *T.Size, exception *ExceptionInfo) []string

var GetConfigureBlob func(filename, path string, length *uint32, exception *ExceptionInfo) *T.Void

var GetConfigureInfo func(name string, exception *ExceptionInfo) *T.ConfigureInfo

var GetConfigureInfoList func(pattern string, numberOptions *T.Size, exception *ExceptionInfo) **T.ConfigureInfo

var GetConfigureList func(pattern string, numberOptions *T.Size, exception *ExceptionInfo) []string

var GetConfigureOptions func(filename string, exception *ExceptionInfo) *LinkedListInfo

var GetConfigurePaths func(filename string, exception *ExceptionInfo) *LinkedListInfo

var GetConfigureValue func(configureInfo *T.ConfigureInfo) string

var GetDelegateList func(pattern string, numberDelegates *T.Size, exception *ExceptionInfo) []string

var GetEnvironmentValue func(name string) string

var GetExceptionMessage func(errorCode int) string

var GetExecutionPath func(path string, extent uint32) bool

var GetGeometry func(geometry string, x, y *T.Long, width, height *T.Size) T.MagickStatusType

var GetImageFromMagickRegistry func(name string, id *T.Long, exception *ExceptionInfo) *Image

var GetImageMagick func(magick *byte, length uint32) string

var GetImageRegistry func(type_ T.RegistryType, key string, exception *ExceptionInfo) *T.Void

var GetLocaleExceptionMessage func(severity T.ExceptionType, tag string) string

var GetLocaleInfo func(tag string, exception *ExceptionInfo) *T.LocaleInfo

var GetLocaleInfoList func(pattern string, numberMessages *T.Size, exception *ExceptionInfo) []*T.LocaleInfo

var GetLocaleList func(pattern string, numberMessages *T.Size, exception *ExceptionInfo) []string

var GetLocaleMessage func(tag string) string

var GetLocaleOptions func(filename string, exception *ExceptionInfo) *LinkedListInfo

var GetLocaleValue func(localeInfo *T.LocaleInfo) string

var GetLogInfoList func(pattern string, numberPreferences *T.Size, exception *ExceptionInfo) []*LogInfo

var GetLogList func(pattern string, numberPreferences *T.Size, exception *ExceptionInfo) []string

var GetLogName func() string

var GetMagicInfo func(magic *byte, length uint32, exception *ExceptionInfo) *T.MagicInfo

var GetMagicInfoList func(pattern string, numberAliases *T.Size, exception *ExceptionInfo) []*T.MagicInfo

var GetMagickCopyright func() string

var GetMagickGeometry func(geometry string, x, y *T.Long, width, height *T.Size) uint

var GetMagickHomeURL func() string

var GetMagickInfo func(name string, exception *ExceptionInfo) *MagickInfo

var GetMagickInfoList func(pattern string, numberFormats *T.Size, exception *ExceptionInfo) []*MagickInfo

var GetMagickList func(pattern string, numberFormats *T.Size, exception *ExceptionInfo) []string

var GetMagickPackageName func() string

var GetMagickQuantumDepth func(depth *T.Size) string

var GetMagickQuantumRange func(range_ *T.Size) string

var GetMagickRegistry func(id T.Long, type_ *T.RegistryType, length *uint32, exception *ExceptionInfo) *T.Void

var GetMagickReleaseDate func() string

var GetMagickToken func(start string, end []string, token string)

var GetMagickPrecision func() int

var GetMagickVersion func(version *T.Size) string

var GetMagicList func(pattern string, numberAliases *T.Size, exception *ExceptionInfo) []string

var GetMagicName func(magicInfo *T.MagicInfo) string

var GetMimeDescription func(mimeInfo *MimeInfo) string

var GetMimeInfo func(filename string, magic *byte, length uint32, exception *ExceptionInfo) *MimeInfo

var GetMimeInfoList func(pattern string, numberAliases *T.Size, exception *ExceptionInfo) []*MimeInfo

var GetMimeList func(pattern string, numberAliases *T.Size, exception *ExceptionInfo) []string

var GetMimeType func(mimeInfo *MimeInfo) string

var GetModuleInfo func(tag string, exception *ExceptionInfo) *T.ModuleInfo

var GetModuleInfoList func(pattern string, numberModules *T.Size, exception *ExceptionInfo) **T.ModuleInfo

var GetModuleList func(pattern string, numberModules *T.Size, exception *ExceptionInfo) []string

var GetMonitorHandler func() T.MonitorHandler

var GetNextImageRegistry func() string

var GetOptimalKernelWidth func(radius, sigma float64) T.Size

var GetOptimalKernelWidth1D func(radius, sigma float64) T.Size

var GetOptimalKernelWidth2D func(radius, sigma float64) T.Size

var GetPageGeometry func(pageGeometry string) string

var GetPathComponent func(path string, type_ T.PathType, component string)

var GetPathComponents func(path string, numberComponents *T.Size) []string

var GetRandomKey func(key *byte, length uint32)

var GetRandomValue func() float64

var GetThresholdMap func(mapId string, exception *ExceptionInfo) *T.ThresholdMap

var GetThresholdMapFile func(xml, filename, mapId string, exception *ExceptionInfo) *T.ThresholdMap

var GetTypeInfo func(name string, exception *ExceptionInfo) *T.TypeInfo

var GetTypeInfoByFamily func(family string, style T.StyleType, stretch T.StretchType, weight T.Size, exception *ExceptionInfo) *T.TypeInfo

var GetTypeInfoList func(pattern string, numberFonts *T.Size, exception *ExceptionInfo) **T.TypeInfo

var GetTypeList func(pattern string, numberFonts *T.Size, exception *ExceptionInfo) []string

var GlobExpression func(expression, pattern string, caseInsensitive bool) bool

var GravityAdjustGeometry func(width, height T.Size, gravity T.GravityType, region *T.RectangleInfo)

var HashPointerType func(pointer *T.Void) uint32

var HashStringInfoType func(string *T.Void) uint32

var HashStringType func(string *T.Void) uint32

var IdentityAffine func(affine *T.AffineMatrix)

var InitializeMagick func(path string)

var InterpretImageFilename func(str string, length uint32, format string, value int) T.Long

var InvokeDynamicImageFilter func(tag string, images **Image, argc int, argv []string, exception *ExceptionInfo) bool

var InvokeStaticImageFilter func(tag string, image **Image, argc int, argv []string, exception *ExceptionInfo) bool

var IsEventLogging func() bool

var IsGeometry func(geometry string) bool

var IsGlob func(path string) bool

var IsMagickConflict func(magick string) bool

var IsMagickInstantiated func() bool

var IsMagickTrue func(value string) bool

var IsSceneGeometry func(geometry string, pedantic bool) bool

var IsSubimage func(geometry string, pedantic uint) uint

var LeastSquaresAddTerms func(matrix, vectors **float64, terms, results *float64, rank, nvecs T.Size)

var LiberateMemory func(memory **T.Void)

var ListFiles func(directory, pattern string, numberEntries *T.Size) []string

var LoadMimeLists func(filename string, exception *ExceptionInfo) bool

var LocaleCompare func(p, q string) T.Long

var LocaleLower func(str string)

var LocaleNCompare func(p, q string, length uint32) T.Long

var LocaleUpper func(str string)

var LogMagickEvent func(type_ T.LogEventType, module, function string, line T.Size, format string, va ...VArg) bool

var LogMagickEventList func(type_ T.LogEventType, module, function string, line T.Size, format string, operands VAList) bool

var MagickCoreGenesis func(path string, establishSignalHandlers bool)

var MagickCoreTerminus func()

var MagickError func(err T.ExceptionType, reason, description string)

var MagickFatalError func(err T.ExceptionType, reason, description string)

var MagickIncarnate func(path string)

var MagickMonitor func(text string, offset T.MagickOffsetType, span T.MagickSizeType, clientData *T.Void) bool

var MagickToMime func(magick string) string

var MagickWarning func(warning T.ExceptionType, reason, description string)

var MapBlob func(file int, mode T.MapMode, offset T.MagickOffsetType, length uint32) *byte

var MapImages func(images, mapImage *Image, dither bool) bool

var ModifyImage func(image **Image, exception *ExceptionInfo) bool

var MSBOrderLong func(buffer *byte, length uint32)

var MSBOrderShort func(p *byte, length uint32)

var MultilineCensus func(label string) T.Size

var NewHashmap func(capacity T.Size, hash func(*T.Void) uint32, compare func(*T.Void, *T.Void) *bool, relinquishKey, relinquishValue func(*T.Void) *T.Void) *HashmapInfo

var NewImageList func() *Image

var NewXMLTree func(xml string, exception *ExceptionInfo) *XMLTreeInfo

var NewXMLTreeTag func(tag string) *XMLTreeInfo

var OpenModule func(module string, exception *ExceptionInfo) bool

var OpenModules func(exception *ExceptionInfo) bool

var ParseAbsoluteGeometry func(geometry string, regionInfo *T.RectangleInfo) T.MagickStatusType

var ParseAffineGeometry func(geometry string, affineMatrix *T.AffineMatrix) T.MagickStatusType

var ParseChannelOption func(channels string) T.Long

var ParseGeometry func(geometry string, geometryInfo *T.GeometryInfo) T.MagickStatusType

var ParseImageGeometry func(geometry string, x, y *T.Long, width, height *T.Size) int

var ParseMetaGeometry func(geometry string, x, y *T.Long, width, height *T.Size) T.MagickStatusType

var PostscriptGeometry func(page string) string

var PrintStringInfo func(file *FILE, id string, stringInfo *StringInfo)

var QueryColorDatabase func(name string, color *T.PixelPacket, exception *ExceptionInfo) bool

var QueryMagickColor func(name string, color *MagickPixelPacket, exception *ExceptionInfo) bool

var ReacquireMemory func(memory **T.Void, size uint32)

var RegisterStaticModules func()

var RelinquishMagickMatrix func(matrix **float64, nptrs T.Size) **float64

var RelinquishMagickMemory func(memory *T.Void) *T.Void

var RelinquishUniqueFileResource func(path string) bool

var RemoveImageRegistry func(key string) *T.Void

var ResetImageRegistryIterator func()

var ResetMagickMemory func(memory *T.Void, byte int, size uint32) *T.Void

var ResizeMagickMemory func(memory *T.Void, size uint32) *T.Void

var ResizeQuantumMemory func(memory *T.Void, count, quantum uint32) *T.Void

var SetCacheThreshold func(size T.Size)

var SetClientName func(name string) string

var SetClientPath func(path string) string

var SetErrorHandler func(handler ErrorHandler) ErrorHandler

var SetFatalErrorHandler func(handler FatalErrorHandler) FatalErrorHandler

var SetGeometryInfo func(geometryInfo *T.GeometryInfo)

var SetImageRegistry func(type_ T.RegistryType, key string, value *T.Void, exception *ExceptionInfo) bool

var SetLogEventMask func(events string) T.LogEventType

var SetLogFormat func(format string)

var SetLogName func(name string) string

var SetMagickInfo func(name string) *MagickInfo

var SetMagickRegistry func(type_ T.RegistryType, blob *T.Void, length uint32, exception *ExceptionInfo) T.Long

var SetMonitorHandler func(handler T.MonitorHandler) T.MonitorHandler

var SetWarningHandler func(handler WarningHandler) WarningHandler

var StringToArgv func(text string, argc *int) []string

var StringToDouble func(str string, interval float64) float64

var StringToken func(delimiters string, string []string) string

var StringToList func(text string) []string

var StringToStringInfo func(str string) *StringInfo

var Strip func(message string)

var StripString func(message string)

var SubstituteString func(buffer []string, search, replace string) bool

var SystemCommand func(verbose bool, command string) int

var TemporaryFilename func(path string)

var TransformImage func(image **Image, cropGeometry, imageGeometry string) bool

var UnmapBlob func(map_ *T.Void, length uint32) bool

var UnregisterMagickInfo func(name string) bool

var UnregisterStaticModules func()

var XDestroyResourceInfo func(resourceInfo *T.XResourceInfo)

var XGetAnnotateInfo func(annotateInfo *T.XAnnotateInfo)

var XGetImportInfo func(ximageInfo *T.XImportInfo)

var XGetMapInfo func(visualInfo *T.XVisualInfo, colormap T.Colormap, mapInfo *T.XStandardColormap)

var XInitImage func(ximage *T.XImage) T.Status

var XMagickProgressMonitor func(tag string, quantum T.MagickOffsetType, span T.MagickSizeType, clientData *T.Void) bool

var XQueryColorDatabase func(target string, color *T.XColor) bool

var XrmCombineFileDatabase func(filename string, target *XrmDatabase, override bool) T.Status

var XSetLocaleModifiers func(modifiers string) string

var XSetWindows func(windowsInfo *T.XWindows) *T.XWindows

var XSupportsLocale func() bool

var XUserPreferences func(resourceInfo *T.XResourceInfo)

var XWarning func(warning T.ExceptionType, reason, description string)

var AcquireAlignedMemory func(count uint32, quantum uint32) *T.Void

var AcquireKernelBuiltIn func(type_ T.KernelInfoType, args T.GeometryInfo) *T.KernelInfo

var AcquireKernelInfo func(kernelString string) *T.KernelInfo

var AcquireModuleInfo func(path, tag string) *T.ModuleInfo

var AcquireOneCacheViewVirtualPixel func(cacheView *CacheView, virtualPixelMethod T.VirtualPixelMethod, x, y int32, pixel *T.PixelPacket, exception *ExceptionInfo) bool

var AcquirePixelCacheNexus func(numberThreads uint32) **T.NexusInfo

var CloneKernelInfo func(kernel *T.KernelInfo) *T.KernelInfo

var DestroyKernelInfo func(kernel *T.KernelInfo) *T.KernelInfo

var DestroyPixelCache func()

var DuplicateImages func(images *Image, numberDuplicates uint32, scenes string, exception *ExceptionInfo) *Image

var EvaluateImages func(images *Image, op T.MagickEvaluateOperator, value float64, exception *ExceptionInfo) bool

var GetImageViewException func(pixelImage *ImageView, severity *T.ExceptionType) string

var GetMagickFeatures func() string

var GetMagickMemoryMethods func(a *AcquireMemoryHandler, r *ResizeMemoryHandler, destroyMemoryHandler *DestroyMemoryHandler)

var HSLTransform func(hue, saturation, lightness float64, red, green, blue *T.Quantum)

var InitializeModuleList func(exception *ExceptionInfo)

var MaximumImages func(images *Image, exception *ExceptionInfo) *Image

var MinimumImages func(images *Image, exception *ExceptionInfo) *Image

var NewImageView func(i *Image) *ImageView

var NewImageViewRegion func(i *Image, x, y int32, width, height uint32) *ImageView

var OpenMagickStream func(path, mode string) *FILE

var QueryColorCompliance func(name string, compliance T.ComplianceType, color *T.PixelPacket, exception *ExceptionInfo) bool

var RegistryComponentTerminus func()

var RelinquishAlignedMemory func(memory *T.Void) *T.Void

var ReplaceImageInListReturnLast func(images **Image, replace *Image)

var ScaleGeometryKernelInfo func(kernel *T.KernelInfo, scalingFactor float64, normalizeFlags T.MagickStatusType)

var ScaleKernelInfo func(kernel *T.KernelInfo, scalingFactor float64, normalizeFlags T.MagickStatusType)

var SetMagickMemoryMethods func(a AcquireMemoryHandler, r ResizeMemoryHandler, d DestroyMemoryHandler)

var SetMagickPrecision func(precision int) int

var ShowKernelInfo func(kernel *T.KernelInfo)

var SmushImages func(images *Image, stack bool, exception *ExceptionInfo) *Image

var TransformHSL func(red, green, blue T.Quantum, hue, saturation, lightness *float64)

var ZeroKernelNans func(kernel *T.KernelInfo)

var Exit func(int) int

var IsWindows95 func() bool

var NTCloseDirectory func(*T.DIR) int

var NTCloseLibrary func(*T.Void) int

var NTControlHandler func() int

var NTElapsedTime func() float64

var NTErrorHandler func(T.ExceptionType, string, string) // Return
var NTExitLibrary func() int

var NTGetExecutionPath func(string, uint32) bool

var NTGetLastError func() string

var NTGetLibraryError func() string

var NTGetLibrarySymbol func(*T.Void, string) *T.Void //* Return?
var NTGetModulePath func(string, string) bool

var NTGhostscriptDLL func(string, int) int

var NTGhostscriptDLLVectors func() *T.GhostInfo

var NTGhostscriptEXE func(string, int) int

var NTGhostscriptFonts func(string, int) int

var NTGhostscriptLoadDLL func() int

var NTGhostscriptUnLoadDLL func() int

var NTInitializeLibrary func() int

var NTIsMagickConflict func(string) bool

var NTLoadTypeLists func(*SplayTreeInfo, *ExceptionInfo) bool

var NTMapMemory func(string, uint32, int, int, int, T.MagickOffsetType) **T.Void //** Return?
var NTOpenDirectory func(string) *T.DIR

var NTOpenLibrary func(string) ***T.Void //*** Return?
var NTReadDirectory func(*T.DIR) *dirent

var NTRegistryKeyLookup func(string) string

var NTReportEvent func(string, bool) bool

var NTResourceToBlob func(string) *byte

var NTSeekDirectory func(*T.DIR, T.Long) ***T.Void //*** Return?
var NTSetSearchPath func(string) int

var NTSyncMemory func(*T.Void, uint32, int) int

var NTSystemCommand func(string) int

var NTSystemConfiguration func(int) T.Long

var NTTellDirectory func(*T.DIR) T.Long

var NTTruncateFile func(int, off_t) int

var NTUnmapMemory func(*T.Void, uint32) int

var NTUserTime func() float64

var NTWarningHandler func(T.ExceptionType, string, string) ***T.Void //*** Return?

var PlasmaImageProxy func(image *Image, image_view *CacheView, random_info *T.RandomInfo, segment *T.SegmentInfo, attenuate, depth uint32) bool

func init() {
	AddDllApis(DLL, false, allApis)
	AddDllData(DLL, false, allData)
	allApis = nil
	allData = nil
}

var DLL = "CORE_RL_magick_.dll"

var allApis = Apis{
	{"AccelerateConvolveImage", &AccelerateConvolveImage},
	{"AcquireAlignedMemory", &AcquireAlignedMemory},
	{"AcquireAuthenticCacheView", &AcquireAuthenticCacheView},
	{"AcquireCacheView", &AcquireCacheView},
	{"AcquireCacheViewIndexes", &AcquireCacheViewIndexes},
	{"AcquireCacheViewPixels", &AcquireCacheViewPixels},
	{"AcquireDrawInfo", &AcquireDrawInfo},
	{"AcquireExceptionInfo", &AcquireExceptionInfo},
	{"AcquireFxInfo", &AcquireFxInfo},
	{"AcquireImage", &AcquireImage},
	{"AcquireImageColormap", &AcquireImageColormap},
	{"AcquireImageInfo", &AcquireImageInfo},
	{"AcquireImagePixels", &AcquireImagePixels},
	{"AcquireIndexes", &AcquireIndexes},
	{"AcquireKernelBuiltIn", &AcquireKernelBuiltIn},
	{"AcquireKernelInfo", &AcquireKernelInfo},
	{"AcquireMagickMatrix", &AcquireMagickMatrix},
	{"AcquireMagickMemory", &AcquireMagickMemory},
	{"AcquireMagickResource", &AcquireMagickResource},
	{"AcquireMemory", &AcquireMemory},
	{"AcquireModuleInfo", &AcquireModuleInfo},
	{"AcquireNextImage", &AcquireNextImage},
	{"AcquireOneCacheViewPixel", &AcquireOneCacheViewPixel},
	{"AcquireOneCacheViewVirtualPixel", &AcquireOneCacheViewVirtualPixel},
	{"AcquireOneMagickPixel", &AcquireOneMagickPixel},
	{"AcquireOnePixel", &AcquireOnePixel},
	{"AcquireOneVirtualPixel", &AcquireOneVirtualPixel},
	{"AcquirePixelCache", &AcquirePixelCache},
	{"AcquirePixelCacheNexus", &AcquirePixelCacheNexus},
	{"AcquirePixelCachePixels", &AcquirePixelCachePixels},
	{"AcquirePixels", &AcquirePixels},
	{"AcquireQuantizeInfo", &AcquireQuantizeInfo},
	{"AcquireQuantumInfo", &AcquireQuantumInfo},
	{"AcquireQuantumMemory", &AcquireQuantumMemory},
	{"AcquireRandomInfo", &AcquireRandomInfo},
	{"AcquireResampleFilter", &AcquireResampleFilter},
	{"AcquireResizeFilter", &AcquireResizeFilter},
	{"AcquireSemaphoreInfo", &AcquireSemaphoreInfo},
	{"AcquireSignatureInfo", &AcquireSignatureInfo},
	{"AcquireStreamInfo", &AcquireStreamInfo},
	{"AcquireString", &AcquireString},
	{"AcquireStringInfo", &AcquireStringInfo},
	{"AcquireTimerInfo", &AcquireTimerInfo},
	{"AcquireTokenInfo", &AcquireTokenInfo},
	{"AcquireUniqueFilename", &AcquireUniqueFilename},
	{"AcquireUniqueFileResource", &AcquireUniqueFileResource},
	{"AcquireUniqueSymbolicLink", &AcquireUniqueSymbolicLink},
	{"AcquireVirtualCacheView", &AcquireVirtualCacheView},
	{"AdaptiveBlurImage", &AdaptiveBlurImage},
	{"AdaptiveBlurImageChannel", &AdaptiveBlurImageChannel},
	{"AdaptiveResizeImage", &AdaptiveResizeImage},
	{"AdaptiveSharpenImage", &AdaptiveSharpenImage},
	{"AdaptiveSharpenImageChannel", &AdaptiveSharpenImageChannel},
	{"AdaptiveThresholdImage", &AdaptiveThresholdImage},
	{"AddChildToXMLTree", &AddChildToXMLTree},
	{"AddNoiseImage", &AddNoiseImage},
	{"AddNoiseImageChannel", &AddNoiseImageChannel},
	{"AddPathToXMLTree", &AddPathToXMLTree},
	{"AddValueToSplayTree", &AddValueToSplayTree},
	{"AffineTransformImage", &AffineTransformImage},
	{"AffinityImage", &AffinityImage},
	{"AffinityImages", &AffinityImages},
	{"AllocateImage", &AllocateImage},
	{"AllocateImageColormap", &AllocateImageColormap},
	{"AllocateNextImage", &AllocateNextImage},
	{"AllocateSemaphoreInfo", &AllocateSemaphoreInfo},
	{"AllocateString", &AllocateString},
	{"AnimateImages", &AnimateImages},
	{"AnnotateComponentGenesis", &AnnotateComponentGenesis},
	{"AnnotateComponentTerminus", &AnnotateComponentTerminus},
	{"AnnotateImage", &AnnotateImage},
	{"AppendImageFormat", &AppendImageFormat},
	{"AppendImages", &AppendImages},
	{"AppendImageToList", &AppendImageToList},
	{"AppendValueToLinkedList", &AppendValueToLinkedList},
	{"Ascii85Encode", &Ascii85Encode},
	{"Ascii85Flush", &Ascii85Flush},
	{"Ascii85Initialize", &Ascii85Initialize},
	{"AsynchronousResourceComponentTerminus", &AsynchronousResourceComponentTerminus},
	{"AttachBlob", &AttachBlob},
	{"AutoGammaImage", &AutoGammaImage},
	{"AutoGammaImageChannel", &AutoGammaImageChannel},
	{"AutoLevelImage", &AutoLevelImage},
	{"AutoLevelImageChannel", &AutoLevelImageChannel},
	{"AverageImages", &AverageImages},
	{"Base64Decode", &Base64Decode},
	{"Base64Encode", &Base64Encode},
	{"BilevelImage", &BilevelImage},
	{"BilevelImageChannel", &BilevelImageChannel},
	{"BlackThresholdImage", &BlackThresholdImage},
	{"BlackThresholdImageChannel", &BlackThresholdImageChannel},
	{"BlobToFile", &BlobToFile},
	{"BlobToImage", &BlobToImage},
	{"BlobToStringInfo", &BlobToStringInfo},
	{"BlueShiftImage", &BlueShiftImage},
	{"BlurImage", &BlurImage},
	{"BlurImageChannel", &BlurImageChannel},
	{"BorderImage", &BorderImage},
	{"BrightnessContrastImage", &BrightnessContrastImage},
	{"BrightnessContrastImageChannel", &BrightnessContrastImageChannel},
	{"CacheComponentGenesis", &CacheComponentGenesis},
	{"CacheComponentTerminus", &CacheComponentTerminus},
	{"CanonicalXMLContent", &CanonicalXMLContent},
	{"CatchException", &CatchException},
	{"CatchImageException", &CatchImageException},
	{"ChannelImage", &ChannelImage},
	{"ChannelThresholdImage", &ChannelThresholdImage},
	{"CharcoalImage", &CharcoalImage},
	{"ChopImage", &ChopImage},
	{"ChopPathComponents", &ChopPathComponents},
	{"ClampImage", &ClampImage},
	{"ClampImageChannel", &ClampImageChannel},
	{"ClearLinkedList", &ClearLinkedList},
	{"ClearMagickException", &ClearMagickException},
	{"ClipImage", &ClipImage},
	{"ClipImagePath", &ClipImagePath},
	{"ClipPathImage", &ClipPathImage},
	{"CloneBlobInfo", &CloneBlobInfo},
	{"CloneCacheView", &CloneCacheView},
	{"CloneDrawInfo", &CloneDrawInfo},
	{"CloneImage", &CloneImage},
	{"CloneImageArtifacts", &CloneImageArtifacts},
	{"CloneImageAttributes", &CloneImageAttributes},
	{"CloneImageInfo", &CloneImageInfo},
	{"CloneImageList", &CloneImageList},
	{"CloneImageOptions", &CloneImageOptions},
	{"CloneImageProfiles", &CloneImageProfiles},
	{"CloneImageProperties", &CloneImageProperties},
	{"CloneImages", &CloneImages},
	{"CloneImageView", &CloneImageView},
	{"CloneKernelInfo", &CloneKernelInfo},
	{"CloneMagickPixelPacket", &CloneMagickPixelPacket},
	{"CloneMemory", &CloneMemory},
	{"CloneMontageInfo", &CloneMontageInfo},
	{"ClonePixelCache", &ClonePixelCache},
	{"ClonePixelCacheMethods", &ClonePixelCacheMethods},
	{"CloneQuantizeInfo", &CloneQuantizeInfo},
	{"CloneSplayTree", &CloneSplayTree},
	{"CloneString", &CloneString},
	{"CloneStringInfo", &CloneStringInfo},
	{"CloseBlob", &CloseBlob},
	{"CloseCacheView", &CloseCacheView},
	{"CloseMagickLog", &CloseMagickLog},
	{"ClutImage", &ClutImage},
	{"ClutImageChannel", &ClutImageChannel},
	{"CoalesceImages", &CoalesceImages},
	{"CoderComponentGenesis", &CoderComponentGenesis},
	{"CoderComponentTerminus", &CoderComponentTerminus},
	{"ColorComponentGenesis", &ColorComponentGenesis},
	{"ColorComponentTerminus", &ColorComponentTerminus},
	{"ColorDecisionListImage", &ColorDecisionListImage},
	{"ColorFloodfillImage", &ColorFloodfillImage},
	{"ColorizeImage", &ColorizeImage},
	{"ColorMatrixImage", &ColorMatrixImage},
	{"CombineImages", &CombineImages},
	{"CommandOptionToMnemonic", &CommandOptionToMnemonic},
	{"CompareHashmapString", &CompareHashmapString},
	{"CompareHashmapStringInfo", &CompareHashmapStringInfo},
	{"CompareImageChannels", &CompareImageChannels},
	{"CompareImageLayers", &CompareImageLayers},
	{"CompareImages", &CompareImages},
	{"CompareSplayTreeString", &CompareSplayTreeString},
	{"CompareSplayTreeStringInfo", &CompareSplayTreeStringInfo},
	{"CompareStringInfo", &CompareStringInfo},
	{"CompositeImage", &CompositeImage},
	{"CompositeImageChannel", &CompositeImageChannel},
	{"CompositeLayers", &CompositeLayers},
	{"CompressImageColormap", &CompressImageColormap},
	{"ConcatenateColorComponent", &ConcatenateColorComponent},
	{"ConcatenateMagickString", &ConcatenateMagickString},
	{"ConcatenateString", &ConcatenateString},
	{"ConcatenateStringInfo", &ConcatenateStringInfo},
	{"ConfigureComponentGenesis", &ConfigureComponentGenesis},
	{"ConfigureComponentTerminus", &ConfigureComponentTerminus},
	{"ConfigureFileToStringInfo", &ConfigureFileToStringInfo},
	{"ConsolidateCMYKImages", &ConsolidateCMYKImages},
	{"ConstantString", &ConstantString},
	{"ConstituteComponentGenesis", &ConstituteComponentGenesis},
	{"ConstituteComponentTerminus", &ConstituteComponentTerminus},
	{"ConstituteImage", &ConstituteImage},
	{"ContinueTimer", &ContinueTimer},
	{"ContrastImage", &ContrastImage},
	{"ContrastStretchImage", &ContrastStretchImage},
	{"ContrastStretchImageChannel", &ContrastStretchImageChannel},
	{"ConvertHCLToRGB", &ConvertHCLToRGB},
	{"ConvertHSBToRGB", &ConvertHSBToRGB},
	{"ConvertHSLToRGB", &ConvertHSLToRGB},
	{"ConvertHWBToRGB", &ConvertHWBToRGB},
	{"ConvertRGBToHCL", &ConvertRGBToHCL},
	{"ConvertRGBToHSB", &ConvertRGBToHSB},
	{"ConvertRGBToHSL", &ConvertRGBToHSL},
	{"ConvertRGBToHWB", &ConvertRGBToHWB},
	{"ConvolveImage", &ConvolveImage},
	{"ConvolveImageChannel", &ConvolveImageChannel},
	{"CopyMagickMemory", &CopyMagickMemory},
	{"CopyMagickString", &CopyMagickString},
	{"CropImage", &CropImage},
	{"CropImageToHBITMAP", &CropImageToHBITMAP},
	{"CropImageToTiles", &CropImageToTiles},
	{"CycleColormapImage", &CycleColormapImage},
	{"DecipherImage", &DecipherImage},
	{"DeconstructImages", &DeconstructImages},
	{"DefineImageArtifact", &DefineImageArtifact},
	{"DefineImageOption", &DefineImageOption},
	{"DefineImageProperty", &DefineImageProperty},
	{"DefineImageRegistry", &DefineImageRegistry},
	{"DelegateComponentGenesis", &DelegateComponentGenesis},
	{"DelegateComponentTerminus", &DelegateComponentTerminus},
	{"DeleteImageArtifact", &DeleteImageArtifact},
	{"DeleteImageAttribute", &DeleteImageAttribute},
	{"DeleteImageFromList", &DeleteImageFromList},
	{"DeleteImageList", &DeleteImageList},
	{"DeleteImageOption", &DeleteImageOption},
	{"DeleteImageProfile", &DeleteImageProfile},
	{"DeleteImageProperty", &DeleteImageProperty},
	{"DeleteImageRegistry", &DeleteImageRegistry},
	{"DeleteImages", &DeleteImages},
	{"DeleteMagickRegistry", &DeleteMagickRegistry},
	{"DeleteNodeByValueFromSplayTree", &DeleteNodeByValueFromSplayTree},
	{"DeleteNodeFromSplayTree", &DeleteNodeFromSplayTree},
	{"DescribeImage", &DescribeImage},
	{"DeskewImage", &DeskewImage},
	{"DespeckleImage", &DespeckleImage},
	{"DestroyBlob", &DestroyBlob},
	{"DestroyCacheView", &DestroyCacheView},
	{"DestroyConfigureOptions", &DestroyConfigureOptions},
	{"DestroyConstitute", &DestroyConstitute},
	{"DestroyDrawInfo", &DestroyDrawInfo},
	{"DestroyExceptionInfo", &DestroyExceptionInfo},
	{"DestroyFxInfo", &DestroyFxInfo},
	{"DestroyHashmap", &DestroyHashmap},
	{"DestroyImage", &DestroyImage},
	{"DestroyImageArtifacts", &DestroyImageArtifacts},
	{"DestroyImageAttributes", &DestroyImageAttributes},
	{"DestroyImageInfo", &DestroyImageInfo},
	{"DestroyImageList", &DestroyImageList},
	{"DestroyImageOptions", &DestroyImageOptions},
	{"DestroyImagePixels", &DestroyImagePixels},
	{"DestroyImageProfiles", &DestroyImageProfiles},
	{"DestroyImageProperties", &DestroyImageProperties},
	{"DestroyImages", &DestroyImages},
	{"DestroyImageView", &DestroyImageView},
	{"DestroyKernelInfo", &DestroyKernelInfo},
	{"DestroyLinkedList", &DestroyLinkedList},
	{"DestroyLocaleOptions", &DestroyLocaleOptions},
	{"DestroyMagick", &DestroyMagick},
	{"DestroyMagickMemory", &DestroyMagickMemory},
	{"DestroyMagickRegistry", &DestroyMagickRegistry},
	{"DestroyModuleList", &DestroyModuleList},
	{"DestroyMontageInfo", &DestroyMontageInfo},
	{"DestroyPixelCache", &DestroyPixelCache},
	{"DestroyPixelCacheNexus", &DestroyPixelCacheNexus},
	{"DestroyQuantizeInfo", &DestroyQuantizeInfo},
	{"DestroyQuantumInfo", &DestroyQuantumInfo},
	{"DestroyRandomInfo", &DestroyRandomInfo},
	{"DestroyResampleFilter", &DestroyResampleFilter},
	{"DestroyResizeFilter", &DestroyResizeFilter},
	{"DestroySemaphoreInfo", &DestroySemaphoreInfo},
	{"DestroySignatureInfo", &DestroySignatureInfo},
	{"DestroySplayTree", &DestroySplayTree},
	{"DestroyStreamInfo", &DestroyStreamInfo},
	{"DestroyString", &DestroyString},
	{"DestroyStringInfo", &DestroyStringInfo},
	{"DestroyStringList", &DestroyStringList},
	{"DestroyThresholdMap", &DestroyThresholdMap},
	{"DestroyTimerInfo", &DestroyTimerInfo},
	{"DestroyTokenInfo", &DestroyTokenInfo},
	{"DestroyXMLTree", &DestroyXMLTree},
	{"DestroyXResources", &DestroyXResources},
	{"DestroyXWidget", &DestroyXWidget},
	{"DetachBlob", &DetachBlob},
	{"DisassociateImageStream", &DisassociateImageStream},
	{"DiscardBlobBytes", &DiscardBlobBytes},
	{"DispatchImage", &DispatchImage},
	{"DisplayImages", &DisplayImages},
	{"DisposeImages", &DisposeImages},
	{"DistortImage", &DistortImage},
	{"DistortResizeImage", &DistortResizeImage},
	{"DrawAffineImage", &DrawAffineImage},
	{"DrawClipPath", &DrawClipPath},
	{"DrawGradientImage", &DrawGradientImage},
	{"DrawImage", &DrawImage},
	{"DrawPatternPath", &DrawPatternPath},
	{"DrawPrimitive", &DrawPrimitive},
	{"DuplexTransferImageViewIterator", &DuplexTransferImageViewIterator},
	{"DuplicateBlob", &DuplicateBlob},
	{"DuplicateImages", &DuplicateImages},
	{"EdgeImage", &EdgeImage},
	{"EmbossImage", &EmbossImage},
	{"EncipherImage", &EncipherImage},
	{"EnhanceImage", &EnhanceImage},
	{"EOFBlob", &EOFBlob},
	{"EqualizeImage", &EqualizeImage},
	{"EqualizeImageChannel", &EqualizeImageChannel},
	{"EscapeString", &EscapeString},
	{"EvaluateImage", &EvaluateImage},
	{"EvaluateImageChannel", &EvaluateImageChannel},
	{"EvaluateImages", &EvaluateImages},
	{"ExcerptImage", &ExcerptImage},
	{"Exit", &Exit},
	{"ExpandAffine", &ExpandAffine},
	{"ExpandFilename", &ExpandFilename},
	{"ExpandFilenames", &ExpandFilenames},
	{"ExportImagePixels", &ExportImagePixels},
	{"ExportQuantumPixels", &ExportQuantumPixels},
	{"ExtentImage", &ExtentImage},
	{"ExtractSubimageFromImage", &ExtractSubimageFromImage},
	{"FileToBlob", &FileToBlob},
	{"FileToImage", &FileToImage},
	{"FileToString", &FileToString},
	{"FileToStringInfo", &FileToStringInfo},
	{"FilterImage", &FilterImage},
	{"FilterImageChannel", &FilterImageChannel},
	{"FinalizeSignature", &FinalizeSignature},
	{"FlattenImages", &FlattenImages},
	{"FlipImage", &FlipImage},
	{"FloodfillPaintImage", &FloodfillPaintImage},
	{"FlopImage", &FlopImage},
	{"FormatImageAttribute", &FormatImageAttribute},
	{"FormatImageAttributeList", &FormatImageAttributeList},
	{"FormatImageProperty", &FormatImageProperty},
	{"FormatImagePropertyList", &FormatImagePropertyList},
	{"FormatLocaleFile", &FormatLocaleFile},
	{"FormatLocaleFileList", &FormatLocaleFileList},
	{"FormatLocaleString", &FormatLocaleString},
	{"FormatLocaleStringList", &FormatLocaleStringList},
	{"FormatMagickCaption", &FormatMagickCaption},
	{"FormatMagickSize", &FormatMagickSize},
	{"FormatMagickString", &FormatMagickString},
	{"FormatMagickStringList", &FormatMagickStringList},
	{"FormatMagickTime", &FormatMagickTime},
	{"FormatString", &FormatString},
	{"FormatStringList", &FormatStringList},
	{"ForwardFourierTransformImage", &ForwardFourierTransformImage},
	{"FrameImage", &FrameImage},
	{"FunctionImage", &FunctionImage},
	{"FunctionImageChannel", &FunctionImageChannel},
	{"FuzzyColorCompare", &FuzzyColorCompare},
	{"FuzzyColorMatch", &FuzzyColorMatch},
	{"FuzzyOpacityCompare", &FuzzyOpacityCompare},
	{"FxEvaluateChannelExpression", &FxEvaluateChannelExpression},
	{"FxEvaluateExpression", &FxEvaluateExpression},
	{"FxImage", &FxImage},
	{"FxImageChannel", &FxImageChannel},
	// Undocumented {"FxPreprocessExpression", &FxPreprocessExpression},
	{"GammaImage", &GammaImage},
	{"GammaImageChannel", &GammaImageChannel},
	{"GaussianBlurImage", &GaussianBlurImage},
	{"GaussianBlurImageChannel", &GaussianBlurImageChannel},
	{"GaussJordanElimination", &GaussJordanElimination},
	{"GenerateDifferentialNoise", &GenerateDifferentialNoise},
	{"GetAffineMatrix", &GetAffineMatrix},
	{"GetAuthenticIndexQueue", &GetAuthenticIndexQueue},
	{"GetAuthenticPixelCacheNexus", &GetAuthenticPixelCacheNexus},
	{"GetAuthenticPixelQueue", &GetAuthenticPixelQueue},
	{"GetAuthenticPixels", &GetAuthenticPixels},
	{"GetBlobError", &GetBlobError},
	{"GetBlobFileHandle", &GetBlobFileHandle},
	{"GetBlobInfo", &GetBlobInfo},
	{"GetBlobProperties", &GetBlobProperties},
	{"GetBlobSize", &GetBlobSize},
	{"GetBlobStreamData", &GetBlobStreamData},
	{"GetBlobStreamHandler", &GetBlobStreamHandler},
	{"GetCacheView", &GetCacheView},
	{"GetCacheViewAuthenticIndexQueue", &GetCacheViewAuthenticIndexQueue},
	{"GetCacheViewAuthenticPixelQueue", &GetCacheViewAuthenticPixelQueue},
	{"GetCacheViewAuthenticPixels", &GetCacheViewAuthenticPixels},
	{"GetCacheViewChannels", &GetCacheViewChannels},
	{"GetCacheViewColorspace", &GetCacheViewColorspace},
	{"GetCacheViewException", &GetCacheViewException},
	{"GetCacheViewExtent", &GetCacheViewExtent},
	{"GetCacheViewIndexes", &GetCacheViewIndexes},
	{"GetCacheViewPixels", &GetCacheViewPixels},
	{"GetCacheViewStorageClass", &GetCacheViewStorageClass},
	{"GetCacheViewVirtualIndexQueue", &GetCacheViewVirtualIndexQueue},
	{"GetCacheViewVirtualPixelQueue", &GetCacheViewVirtualPixelQueue},
	{"GetCacheViewVirtualPixels", &GetCacheViewVirtualPixels},
	{"GetClientName", &GetClientName},
	{"GetClientPath", &GetClientPath},
	{"GetCoderInfo", &GetCoderInfo},
	{"GetCoderInfoList", &GetCoderInfoList},
	{"GetCoderList", &GetCoderList},
	// Undocumented {"GetColorCompliance", &GetColorCompliance},
	{"GetColorInfo", &GetColorInfo},
	{"GetColorInfoList", &GetColorInfoList},
	{"GetColorList", &GetColorList},
	{"GetColorTuple", &GetColorTuple},
	{"GetCommandOptionFlags", &GetCommandOptionFlags},
	{"GetCommandOptions", &GetCommandOptions},
	{"GetConfigureBlob", &GetConfigureBlob},
	{"GetConfigureInfo", &GetConfigureInfo},
	{"GetConfigureInfoList", &GetConfigureInfoList},
	{"GetConfigureList", &GetConfigureList},
	{"GetConfigureOption", &GetConfigureOption},
	{"GetConfigureOptions", &GetConfigureOptions},
	{"GetConfigurePaths", &GetConfigurePaths},
	{"GetConfigureValue", &GetConfigureValue},
	{"GetDelegateCommand", &GetDelegateCommand},
	{"GetDelegateCommands", &GetDelegateCommands},
	{"GetDelegateInfo", &GetDelegateInfo},
	{"GetDelegateInfoList", &GetDelegateInfoList},
	{"GetDelegateList", &GetDelegateList},
	{"GetDelegateMode", &GetDelegateMode},
	{"GetDelegateThreadSupport", &GetDelegateThreadSupport},
	{"GetDrawInfo", &GetDrawInfo},
	{"GetElapsedTime", &GetElapsedTime},
	{"GetEnvironmentValue", &GetEnvironmentValue},
	{"GetExceptionInfo", &GetExceptionInfo},
	{"GetExceptionMessage", &GetExceptionMessage},
	{"GetExecutionPath", &GetExecutionPath},
	{"GetFirstImageInList", &GetFirstImageInList},
	{"GetGeometry", &GetGeometry},
	{"GetImageAlphaChannel", &GetImageAlphaChannel},
	{"GetImageArtifact", &GetImageArtifact},
	{"GetImageAttribute", &GetImageAttribute},
	{"GetImageBoundingBox", &GetImageBoundingBox},
	{"GetImageChannelDepth", &GetImageChannelDepth},
	{"GetImageChannelDistortion", &GetImageChannelDistortion},
	{"GetImageChannelDistortions", &GetImageChannelDistortions},
	{"GetImageChannelExtrema", &GetImageChannelExtrema},
	{"GetImageChannelFeatures", &GetImageChannelFeatures},
	{"GetImageChannelKurtosis", &GetImageChannelKurtosis},
	{"GetImageChannelMean", &GetImageChannelMean},
	{"GetImageChannelRange", &GetImageChannelRange},
	{"GetImageChannels", &GetImageChannels},
	{"GetImageChannelStatistics", &GetImageChannelStatistics},
	{"GetImageClipMask", &GetImageClipMask},
	{"GetImageClippingPathAttribute", &GetImageClippingPathAttribute},
	{"GetImageDecoder", &GetImageDecoder},
	{"GetImageDepth", &GetImageDepth},
	{"GetImageDistortion", &GetImageDistortion},
	{"GetImageDynamicThreshold", &GetImageDynamicThreshold},
	{"GetImageEncoder", &GetImageEncoder},
	{"GetImageException", &GetImageException},
	{"GetImageExtent", &GetImageExtent},
	{"GetImageExtrema", &GetImageExtrema},
	{"GetImageFromList", &GetImageFromList},
	{"GetImageFromMagickRegistry", &GetImageFromMagickRegistry},
	{"GetImageGeometry", &GetImageGeometry},
	{"GetImageHistogram", &GetImageHistogram},
	{"GetImageIndexInList", &GetImageIndexInList},
	{"GetImageInfo", &GetImageInfo},
	{"GetImageInfoFile", &GetImageInfoFile},
	{"GetImageKurtosis", &GetImageKurtosis},
	{"GetImageList", &GetImageList},
	{"GetImageListIndex", &GetImageListIndex},
	{"GetImageListLength", &GetImageListLength},
	{"GetImageListSize", &GetImageListSize},
	{"GetImageMagick", &GetImageMagick},
	{"GetImageMask", &GetImageMask},
	{"GetImageMean", &GetImageMean},
	{"GetImageOption", &GetImageOption},
	{"GetImagePixelCacheType", &GetImagePixelCacheType},
	{"GetImagePixels", &GetImagePixels},
	{"GetImageProfile", &GetImageProfile},
	{"GetImageProperty", &GetImageProperty},
	{"GetImageQuantizeError", &GetImageQuantizeError},
	{"GetImageQuantumDepth", &GetImageQuantumDepth},
	{"GetImageRange", &GetImageRange},
	{"GetImageReferenceCount", &GetImageReferenceCount},
	{"GetImageRegistry", &GetImageRegistry},
	{"GetImageTotalInkDensity", &GetImageTotalInkDensity},
	{"GetImageType", &GetImageType},
	{"GetImageViewAuthenticIndexes", &GetImageViewAuthenticIndexes},
	{"GetImageViewAuthenticPixels", &GetImageViewAuthenticPixels},
	{"GetImageViewException", &GetImageViewException},
	{"GetImageViewExtent", &GetImageViewExtent},
	{"GetImageViewImage", &GetImageViewImage},
	{"GetImageViewIterator", &GetImageViewIterator},
	{"GetImageViewVirtualIndexes", &GetImageViewVirtualIndexes},
	{"GetImageViewVirtualPixels", &GetImageViewVirtualPixels},
	{"GetImageVirtualPixelMethod", &GetImageVirtualPixelMethod},
	{"GetIndexes", &GetIndexes},
	{"GetLastImageInList", &GetLastImageInList},
	{"GetLastValueInLinkedList", &GetLastValueInLinkedList},
	{"GetLocaleExceptionMessage", &GetLocaleExceptionMessage},
	{"GetLocaleInfo_", &GetLocaleInfo},
	{"GetLocaleInfoList", &GetLocaleInfoList},
	{"GetLocaleList", &GetLocaleList},
	{"GetLocaleMessage", &GetLocaleMessage},
	{"GetLocaleOptions", &GetLocaleOptions},
	{"GetLocaleValue", &GetLocaleValue},
	{"GetLogInfoList", &GetLogInfoList},
	{"GetLogList", &GetLogList},
	{"GetLogName", &GetLogName},
	{"GetMagicInfo", &GetMagicInfo},
	{"GetMagicInfoList", &GetMagicInfoList},
	{"GetMagickAdjoin", &GetMagickAdjoin},
	{"GetMagickBlobSupport", &GetMagickBlobSupport},
	{"GetMagickCopyright", &GetMagickCopyright},
	{"GetMagickDescription", &GetMagickDescription},
	{"GetMagickEndianSupport", &GetMagickEndianSupport},
	{"GetMagickFeatures", &GetMagickFeatures},
	{"GetMagickGeometry", &GetMagickGeometry},
	{"GetMagickHomeURL", &GetMagickHomeURL},
	{"GetMagickInfo", &GetMagickInfo},
	{"GetMagickInfoList", &GetMagickInfoList},
	{"GetMagickList", &GetMagickList},
	{"GetMagickMemoryMethods", &GetMagickMemoryMethods},
	{"GetMagickPackageName", &GetMagickPackageName},
	{"GetMagickPageSize", &GetMagickPageSize},
	{"GetMagickPixelPacket", &GetMagickPixelPacket},
	{"GetMagickPrecision", &GetMagickPrecision},
	{"GetMagickProperty", &GetMagickProperty},
	{"GetMagickQuantumDepth", &GetMagickQuantumDepth},
	{"GetMagickQuantumRange", &GetMagickQuantumRange},
	{"GetMagickRawSupport", &GetMagickRawSupport},
	{"GetMagickRegistry", &GetMagickRegistry},
	{"GetMagickReleaseDate", &GetMagickReleaseDate},
	{"GetMagickResource", &GetMagickResource},
	{"GetMagickResourceLimit", &GetMagickResourceLimit},
	{"GetMagickSeekableStream", &GetMagickSeekableStream},
	{"GetMagickThreadSupport", &GetMagickThreadSupport},
	{"GetMagickToken", &GetMagickToken},
	{"GetMagickVersion", &GetMagickVersion},
	{"GetMagicList", &GetMagicList},
	{"GetMagicName", &GetMagicName},
	{"GetMimeDescription", &GetMimeDescription},
	{"GetMimeInfo", &GetMimeInfo},
	{"GetMimeInfoList", &GetMimeInfoList},
	{"GetMimeList", &GetMimeList},
	{"GetMimeType", &GetMimeType},
	{"GetModuleInfo", &GetModuleInfo},
	{"GetModuleInfoList", &GetModuleInfoList},
	{"GetModuleList", &GetModuleList},
	{"GetMonitorHandler", &GetMonitorHandler},
	{"GetMontageInfo", &GetMontageInfo},
	{"GetMultilineTypeMetrics", &GetMultilineTypeMetrics},
	{"GetNextImage", &GetNextImage},
	{"GetNextImageArtifact", &GetNextImageArtifact},
	{"GetNextImageAttribute", &GetNextImageAttribute},
	{"GetNextImageInList", &GetNextImageInList},
	{"GetNextImageOption", &GetNextImageOption},
	{"GetNextImageProfile", &GetNextImageProfile},
	{"GetNextImageProperty", &GetNextImageProperty},
	{"GetNextImageRegistry", &GetNextImageRegistry},
	{"GetNextKeyInHashmap", &GetNextKeyInHashmap},
	{"GetNextKeyInSplayTree", &GetNextKeyInSplayTree},
	{"GetNextValueInHashmap", &GetNextValueInHashmap},
	{"GetNextValueInLinkedList", &GetNextValueInLinkedList},
	{"GetNextValueInSplayTree", &GetNextValueInSplayTree},
	{"GetNextXMLTreeTag", &GetNextXMLTreeTag},
	{"GetNumberColors", &GetNumberColors},
	{"GetNumberOfElementsInLinkedList", &GetNumberOfElementsInLinkedList},
	{"GetNumberOfEntriesInHashmap", &GetNumberOfEntriesInHashmap},
	{"GetNumberOfNodesInSplayTree", &GetNumberOfNodesInSplayTree},
	{"GetNumberScenes", &GetNumberScenes},
	{"GetOneAuthenticPixel", &GetOneAuthenticPixel},
	{"GetOneCacheViewAuthenticPixel", &GetOneCacheViewAuthenticPixel},
	{"GetOneCacheViewVirtualMethodPixel", &GetOneCacheViewVirtualMethodPixel},
	{"GetOneCacheViewVirtualPixel", &GetOneCacheViewVirtualPixel},
	{"GetOnePixel", &GetOnePixel},
	{"GetOneVirtualMagickPixel", &GetOneVirtualMagickPixel},
	{"GetOneVirtualMethodPixel", &GetOneVirtualMethodPixel},
	{"GetOneVirtualPixel", &GetOneVirtualPixel},
	{"GetOptimalKernelWidth", &GetOptimalKernelWidth},
	{"GetOptimalKernelWidth1D", &GetOptimalKernelWidth1D},
	{"GetOptimalKernelWidth2D", &GetOptimalKernelWidth2D},
	{"GetPageGeometry", &GetPageGeometry},
	{"GetPathAttributes", &GetPathAttributes},
	{"GetPathComponent", &GetPathComponent},
	{"GetPathComponents", &GetPathComponents},
	{"GetPixelCacheChannels", &GetPixelCacheChannels},
	{"GetPixelCacheColorspace", &GetPixelCacheColorspace},
	{"GetPixelCacheMethods", &GetPixelCacheMethods},
	{"GetPixelCacheNexusExtent", &GetPixelCacheNexusExtent},
	// Undocumented {"GetPixelCacheNexusIndexes", &GetPixelCacheNexusIndexes},
	// Undocumented {"GetPixelCacheNexusPixels", &GetPixelCacheNexusPixels},
	{"GetPixelCachePixels", &GetPixelCachePixels},
	{"GetPixelCacheStorageClass", &GetPixelCacheStorageClass},
	{"GetPixelCacheTileSize", &GetPixelCacheTileSize},
	{"GetPixelCacheType", &GetPixelCacheType},
	{"GetPixelCacheVirtualMethod", &GetPixelCacheVirtualMethod},
	{"GetPixels", &GetPixels},
	{"GetPolicyInfoList", &GetPolicyInfoList},
	{"GetPolicyList", &GetPolicyList},
	{"GetPolicyValue", &GetPolicyValue},
	{"GetPreviousImage", &GetPreviousImage},
	{"GetPreviousImageInList", &GetPreviousImageInList},
	{"GetPseudoRandomValue", &GetPseudoRandomValue},
	{"GetQuantizeInfo", &GetQuantizeInfo},
	{"GetQuantumEndian", &GetQuantumEndian},
	{"GetQuantumExtent", &GetQuantumExtent},
	{"GetQuantumFormat", &GetQuantumFormat},
	{"GetQuantumInfo", &GetQuantumInfo},
	{"GetQuantumPixels", &GetQuantumPixels},
	{"GetQuantumType", &GetQuantumType},
	{"GetRandomKey", &GetRandomKey},
	{"GetRandomSecretKey", &GetRandomSecretKey},
	{"GetRandomValue", &GetRandomValue},
	{"GetResizeFilterSupport", &GetResizeFilterSupport},
	{"GetResizeFilterWeight", &GetResizeFilterWeight},
	{"GetSignatureBlocksize", &GetSignatureBlocksize},
	{"GetSignatureDigest", &GetSignatureDigest},
	{"GetSignatureDigestsize", &GetSignatureDigestsize},
	{"GetStreamInfoClientData", &GetStreamInfoClientData},
	{"GetStringInfoDatum", &GetStringInfoDatum},
	{"GetStringInfoLength", &GetStringInfoLength},
	{"GetStringInfoPath", &GetStringInfoPath},
	{"GetThresholdMap", &GetThresholdMap},
	{"GetThresholdMapFile", &GetThresholdMapFile},
	{"gettimeofday", &Gettimeofday},
	{"GetTimerInfo", &GetTimerInfo},
	{"GetTypeInfo", &GetTypeInfo},
	{"GetTypeInfoByFamily", &GetTypeInfoByFamily},
	{"GetTypeInfoList", &GetTypeInfoList},
	{"GetTypeList", &GetTypeList},
	{"GetTypeMetrics", &GetTypeMetrics},
	{"GetUserTime", &GetUserTime},
	{"GetValueFromHashmap", &GetValueFromHashmap},
	{"GetValueFromLinkedList", &GetValueFromLinkedList},
	{"GetValueFromSplayTree", &GetValueFromSplayTree},
	{"GetVirtualIndexesFromNexus", &GetVirtualIndexesFromNexus},
	{"GetVirtualIndexQueue", &GetVirtualIndexQueue},
	{"GetVirtualPixelQueue", &GetVirtualPixelQueue},
	{"GetVirtualPixels", &GetVirtualPixels},
	{"GetVirtualPixelsFromNexus", &GetVirtualPixelsFromNexus},
	{"GetVirtualPixelsNexus", &GetVirtualPixelsNexus},
	{"GetXMLTreeAttribute", &GetXMLTreeAttribute},
	{"GetXMLTreeAttributes", &GetXMLTreeAttributes},
	{"GetXMLTreeChild", &GetXMLTreeChild},
	{"GetXMLTreeContent", &GetXMLTreeContent},
	{"GetXMLTreeOrdered", &GetXMLTreeOrdered},
	{"GetXMLTreePath", &GetXMLTreePath},
	{"GetXMLTreeProcessingInstructions", &GetXMLTreeProcessingInstructions},
	{"GetXMLTreeSibling", &GetXMLTreeSibling},
	{"GetXMLTreeTag", &GetXMLTreeTag},
	{"GlobExpression", &GlobExpression},
	{"GradientImage", &GradientImage},
	{"GravityAdjustGeometry", &GravityAdjustGeometry},
	{"HaldClutImage", &HaldClutImage},
	{"HaldClutImageChannel", &HaldClutImageChannel},
	{"HashPointerType", &HashPointerType},
	{"HashStringInfoType", &HashStringInfoType},
	{"HashStringType", &HashStringType},
	{"HSLTransform", &HSLTransform},
	{"HuffmanDecodeImage", &HuffmanDecodeImage},
	{"HuffmanEncodeImage", &HuffmanEncodeImage},
	{"IdentifyImage", &IdentifyImage},
	{"IdentityAffine", &IdentityAffine},
	{"ImageListToArray", &ImageListToArray},
	{"ImagesToBlob", &ImagesToBlob},
	{"ImageToBlob", &ImageToBlob},
	{"ImageToFile", &ImageToFile},
	{"ImageToHBITMAP", &ImageToHBITMAP},
	{"ImplodeImage", &ImplodeImage},
	{"ImportImagePixels", &ImportImagePixels},
	{"ImportQuantumPixels", &ImportQuantumPixels},
	{"InheritException", &InheritException},
	{"InitializeMagick", &InitializeMagick},
	{"InitializeModuleList", &InitializeModuleList},
	{"InitializeSignature", &InitializeSignature},
	{"InjectImageBlob", &InjectImageBlob},
	{"InsertImageInList", &InsertImageInList},
	{"InsertTagIntoXMLTree", &InsertTagIntoXMLTree},
	{"InsertValueInLinkedList", &InsertValueInLinkedList},
	{"InsertValueInSortedLinkedList", &InsertValueInSortedLinkedList},
	{"IntegralRotateImage", &IntegralRotateImage},
	{"InterpolateMagickPixelPacket", &InterpolateMagickPixelPacket},
	{"InterpolatePixelColor", &InterpolatePixelColor},
	{"InterpolativeResizeImage", &InterpolativeResizeImage},
	{"InterpretImageAttributes", &InterpretImageAttributes},
	{"InterpretImageFilename", &InterpretImageFilename},
	{"InterpretImageProperties", &InterpretImageProperties},
	{"InterpretLocaleValue", &InterpretLocaleValue},
	{"InterpretSiPrefixValue", &InterpretSiPrefixValue},
	{"InverseFourierTransformImage", &InverseFourierTransformImage},
	{"InversesRGBCompandor", &InversesRGBCompandor},
	{"InvokeDelegate", &InvokeDelegate},
	{"InvokeDynamicImageFilter", &InvokeDynamicImageFilter},
	{"InvokeStaticImageFilter", &InvokeStaticImageFilter},
	{"IsBlobExempt", &IsBlobExempt},
	{"IsBlobSeekable", &IsBlobSeekable},
	{"IsBlobTemporary", &IsBlobTemporary},
	{"IsColorSimilar", &IsColorSimilar},
	{"IsCommandOption", &IsCommandOption},
	{"IsEventLogging", &IsEventLogging},
	{"IsGeometry", &IsGeometry},
	{"IsGlob", &IsGlob},
	{"IsGrayImage", &IsGrayImage},
	{"IsHashmapEmpty", &IsHashmapEmpty},
	{"IsHighDynamicRangeImage", &IsHighDynamicRangeImage},
	{"IsHistogramImage", &IsHistogramImage},
	{"IsImageObject", &IsImageObject},
	{"IsImagesEqual", &IsImagesEqual},
	{"IsImageSimilar", &IsImageSimilar},
	{"IsImageView", &IsImageView},
	{"IsLinkedListEmpty", &IsLinkedListEmpty},
	{"IsMagickColorSimilar", &IsMagickColorSimilar},
	{"IsMagickConflict", &IsMagickConflict},
	{"IsMagickInstantiated", &IsMagickInstantiated},
	{"IsMagickTrue", &IsMagickTrue},
	{"IsMonochromeImage", &IsMonochromeImage},
	{"IsOpacitySimilar", &IsOpacitySimilar},
	{"IsOpaqueImage", &IsOpaqueImage},
	{"IsPaletteImage", &IsPaletteImage},
	{"IsPathAccessible", &IsPathAccessible},
	{"IsRightsAuthorized", &IsRightsAuthorized},
	{"IsSceneGeometry", &IsSceneGeometry},
	{"IsStringNotFalse", &IsStringNotFalse},
	{"IsStringTrue", &IsStringTrue},
	{"IsSubimage", &IsSubimage},
	{"IsTaintImage", &IsTaintImage},
	{"IsWindows95", &IsWindows95},
	{"LeastSquaresAddTerms", &LeastSquaresAddTerms},
	{"LevelColorsImage", &LevelColorsImage},
	{"LevelColorsImageChannel", &LevelColorsImageChannel},
	{"LevelImage", &LevelImage},
	{"LevelImageChannel", &LevelImageChannel},
	{"LevelImageColors", &LevelImageColors},
	{"LevelizeImage", &LevelizeImage},
	{"LevelizeImageChannel", &LevelizeImageChannel},
	{"LiberateMemory", &LiberateMemory},
	{"LiberateSemaphoreInfo", &LiberateSemaphoreInfo},
	{"LinearStretchImage", &LinearStretchImage},
	{"LinkedListToArray", &LinkedListToArray},
	{"LiquidRescaleImage", &LiquidRescaleImage},
	{"ListCoderInfo", &ListCoderInfo},
	{"ListColorInfo", &ListColorInfo},
	{"ListCommandOptions", &ListCommandOptions},
	{"ListConfigureInfo", &ListConfigureInfo},
	{"ListDelegateInfo", &ListDelegateInfo},
	{"ListFiles", &ListFiles},
	{"ListLocaleInfo", &ListLocaleInfo},
	{"ListLogInfo", &ListLogInfo},
	{"ListMagicInfo", &ListMagicInfo},
	{"ListMagickInfo", &ListMagickInfo},
	{"ListMagickResourceInfo", &ListMagickResourceInfo},
	{"ListMimeInfo", &ListMimeInfo},
	{"ListModuleInfo", &ListModuleInfo},
	{"ListPolicyInfo", &ListPolicyInfo},
	{"ListThresholdMaps", &ListThresholdMaps},
	{"ListTypeInfo", &ListTypeInfo},
	{"LoadMimeLists", &LoadMimeLists},
	{"LocaleCompare", &LocaleCompare},
	{"LocaleComponentGenesis", &LocaleComponentGenesis},
	{"LocaleComponentTerminus", &LocaleComponentTerminus},
	{"LocaleLower", &LocaleLower},
	{"LocaleNCompare", &LocaleNCompare},
	{"LocaleUpper", &LocaleUpper},
	{"LockSemaphoreInfo", &LockSemaphoreInfo},
	{"LogComponentGenesis", &LogComponentGenesis},
	{"LogComponentTerminus", &LogComponentTerminus},
	{"LogMagickEvent", &LogMagickEvent},
	{"LogMagickEventList", &LogMagickEventList},
	{"LZWEncodeImage", &LZWEncodeImage},
	{"MagicComponentGenesis", &MagicComponentGenesis},
	{"MagicComponentTerminus", &MagicComponentTerminus},
	{"MagickComponentGenesis", &MagickComponentGenesis},
	{"MagickComponentTerminus", &MagickComponentTerminus},
	{"MagickCoreGenesis", &MagickCoreGenesis},
	{"MagickCoreTerminus", &MagickCoreTerminus},
	{"MagickCreateThreadKey", &MagickCreateThreadKey},
	{"MagickDelay", &MagickDelay},
	{"MagickDeleteThreadKey", &MagickDeleteThreadKey},
	{"MagickError", &MagickError},
	{"MagickFatalError", &MagickFatalError},
	{"MagickGetThreadValue", &MagickGetThreadValue},
	{"MagickIncarnate", &MagickIncarnate},
	{"MagickMonitor", &MagickMonitor},
	{"MagickSetThreadValue", &MagickSetThreadValue},
	{"MagickToMime", &MagickToMime},
	{"MagickWarning", &MagickWarning},
	{"MagnifyImage", &MagnifyImage},
	{"MapBlob", &MapBlob},
	{"MapImage", &MapImage},
	{"MapImages", &MapImages},
	{"MatteFloodfillImage", &MatteFloodfillImage},
	{"MaximumImages", &MaximumImages},
	{"MedianFilterImage", &MedianFilterImage},
	{"MergeImageLayers", &MergeImageLayers},
	{"MimeComponentGenesis", &MimeComponentGenesis},
	{"MimeComponentTerminus", &MimeComponentTerminus},
	{"MinifyImage", &MinifyImage},
	{"MinimumImages", &MinimumImages},
	{"MinMaxStretchImage", &MinMaxStretchImage},
	{"ModeImage", &ModeImage},
	{"ModifyImage", &ModifyImage},
	{"ModulateImage", &ModulateImage},
	{"ModuleComponentGenesis", &ModuleComponentGenesis},
	{"ModuleComponentTerminus", &ModuleComponentTerminus},
	{"MontageImageList", &MontageImageList},
	{"MontageImages", &MontageImages},
	{"MorphImages", &MorphImages},
	{"MorphologyApply", &MorphologyApply},
	{"MorphologyImage", &MorphologyImage},
	{"MorphologyImageChannel", &MorphologyImageChannel},
	{"MosaicImages", &MosaicImages},
	{"MotionBlurImage", &MotionBlurImage},
	{"MotionBlurImageChannel", &MotionBlurImageChannel},
	{"MSBOrderLong", &MSBOrderLong},
	{"MSBOrderShort", &MSBOrderShort},
	{"MultilineCensus", &MultilineCensus},
	{"NegateImage", &NegateImage},
	{"NegateImageChannel", &NegateImageChannel},
	{"NewHashmap", &NewHashmap},
	{"NewImageList", &NewImageList},
	{"NewImageView", &NewImageView},
	{"NewImageViewRegion", &NewImageViewRegion},
	{"NewLinkedList", &NewLinkedList},
	{"NewMagickImage", &NewMagickImage},
	{"NewSplayTree", &NewSplayTree},
	{"NewXMLTree", &NewXMLTree},
	{"NewXMLTreeTag", &NewXMLTreeTag},
	{"NormalizeImage", &NormalizeImage},
	{"NormalizeImageChannel", &NormalizeImageChannel},
	{"NTArgvToUTF8", &NTArgvToUTF8},
	{"NTCloseDirectory", &NTCloseDirectory},
	{"NTCloseLibrary", &NTCloseLibrary},
	{"NTControlHandler", &NTControlHandler},
	{"NTElapsedTime", &NTElapsedTime},
	{"NTErrorHandler", &NTErrorHandler},
	{"NTExitLibrary", &NTExitLibrary},
	{"NTGatherRandomData", &NTGatherRandomData},
	{"NTGetExecutionPath", &NTGetExecutionPath},
	{"NTGetLastError", &NTGetLastError},
	{"NTGetLibraryError", &NTGetLibraryError},
	{"NTGetLibrarySymbol", &NTGetLibrarySymbol},
	{"NTGetModulePath", &NTGetModulePath},
	{"NTGhostscriptDLL", &NTGhostscriptDLL},
	{"NTGhostscriptDLLVectors", &NTGhostscriptDLLVectors},
	{"NTGhostscriptEXE", &NTGhostscriptEXE},
	{"NTGhostscriptFonts", &NTGhostscriptFonts},
	{"NTGhostscriptLoadDLL", &NTGhostscriptLoadDLL},
	{"NTGhostscriptUnLoadDLL", &NTGhostscriptUnLoadDLL},
	{"NTInitializeLibrary", &NTInitializeLibrary},
	{"NTIsMagickConflict", &NTIsMagickConflict},
	{"NTLoadTypeLists", &NTLoadTypeLists},
	{"NTMapMemory", &NTMapMemory},
	{"NTOpenDirectory", &NTOpenDirectory},
	{"NTOpenLibrary", &NTOpenLibrary},
	{"NTReadDirectory", &NTReadDirectory},
	{"NTRegistryKeyLookup", &NTRegistryKeyLookup},
	{"NTReportEvent", &NTReportEvent},
	{"NTResourceToBlob", &NTResourceToBlob},
	{"NTSeekDirectory", &NTSeekDirectory},
	{"NTSetSearchPath", &NTSetSearchPath},
	{"NTSyncMemory", &NTSyncMemory},
	{"NTSystemCommand", &NTSystemCommand},
	{"NTSystemConfiguration", &NTSystemConfiguration},
	{"NTTellDirectory", &NTTellDirectory},
	{"NTTruncateFile", &NTTruncateFile},
	{"NTUnmapMemory", &NTUnmapMemory},
	{"NTUserTime", &NTUserTime},
	{"NTWarningHandler", &NTWarningHandler},
	{"OilPaintImage", &OilPaintImage},
	{"OpaqueImage", &OpaqueImage},
	{"OpaquePaintImage", &OpaquePaintImage},
	{"OpaquePaintImageChannel", &OpaquePaintImageChannel},
	{"OpenBlob", &OpenBlob},
	{"OpenCacheView", &OpenCacheView},
	{"OpenMagickStream", &OpenMagickStream},
	{"OpenModule", &OpenModule},
	{"OpenModules", &OpenModules},
	{"OpenStream", &OpenStream},
	{"OptimizeImageLayers", &OptimizeImageLayers},
	{"OptimizeImageTransparency", &OptimizeImageTransparency},
	{"OptimizePlusImageLayers", &OptimizePlusImageLayers},
	{"OrderedDitherImage", &OrderedDitherImage},
	{"OrderedDitherImageChannel", &OrderedDitherImageChannel},
	{"OrderedPosterizeImage", &OrderedPosterizeImage},
	{"OrderedPosterizeImageChannel", &OrderedPosterizeImageChannel},
	{"PackbitsEncodeImage", &PackbitsEncodeImage},
	{"PaintFloodfillImage", &PaintFloodfillImage},
	{"PaintOpaqueImage", &PaintOpaqueImage},
	{"PaintOpaqueImageChannel", &PaintOpaqueImageChannel},
	{"PaintTransparentImage", &PaintTransparentImage},
	{"ParseAbsoluteGeometry", &ParseAbsoluteGeometry},
	{"ParseAffineGeometry", &ParseAffineGeometry},
	{"ParseChannelOption", &ParseChannelOption},
	{"ParseCommandOption", &ParseCommandOption},
	{"ParseGeometry", &ParseGeometry},
	{"ParseGravityGeometry", &ParseGravityGeometry},
	{"ParseImageGeometry", &ParseImageGeometry},
	{"ParseMetaGeometry", &ParseMetaGeometry},
	{"ParsePageGeometry", &ParsePageGeometry},
	{"ParseRegionGeometry", &ParseRegionGeometry},
	{"ParseSizeGeometry", &ParseSizeGeometry},
	{"PasskeyDecipherImage", &PasskeyDecipherImage},
	{"PasskeyEncipherImage", &PasskeyEncipherImage},
	{"PerceptibleImage", &PerceptibleImage},
	{"PerceptibleImageChannel", &PerceptibleImageChannel},
	{"PersistPixelCache", &PersistPixelCache},
	{"PingBlob", &PingBlob},
	{"PingImage", &PingImage},
	{"PingImages", &PingImages},
	{"PlasmaImage", &PlasmaImage},
	{"PlasmaImageProxy", &PlasmaImageProxy},
	{"PolaroidImage", &PolaroidImage},
	{"PolicyComponentGenesis", &PolicyComponentGenesis},
	{"PolicyComponentTerminus", &PolicyComponentTerminus},
	{"PolynomialImage", &PolynomialImage},
	{"PolynomialImageChannel", &PolynomialImageChannel},
	{"PopImageList", &PopImageList},
	{"PopImagePixels", &PopImagePixels},
	{"PosterizeImage", &PosterizeImage},
	{"PosterizeImageChannel", &PosterizeImageChannel},
	{"PostscriptGeometry", &PostscriptGeometry},
	{"PrependImageToList", &PrependImageToList},
	{"PreviewImage", &PreviewImage},
	{"PrintStringInfo", &PrintStringInfo},
	{"ProfileImage", &ProfileImage},
	{"PruneTagFromXMLTree", &PruneTagFromXMLTree},
	{"PushImageList", &PushImageList},
	{"PushImagePixels", &PushImagePixels},
	{"PutEntryInHashmap", &PutEntryInHashmap},
	{"QuantizationError", &QuantizationError},
	{"QuantizeImage", &QuantizeImage},
	{"QuantizeImages", &QuantizeImages},
	{"QueryColorCompliance", &QueryColorCompliance},
	{"QueryColorDatabase", &QueryColorDatabase},
	{"QueryColorname", &QueryColorname},
	{"QueryMagickColor", &QueryMagickColor},
	{"QueryMagickColorCompliance", &QueryMagickColorCompliance},
	{"QueryMagickColorname", &QueryMagickColorname},
	{"QueueAuthenticPixel", &QueueAuthenticPixel},
	{"QueueAuthenticPixelCacheNexus", &QueueAuthenticPixelCacheNexus},
	{"QueueAuthenticPixels", &QueueAuthenticPixels},
	{"QueueCacheViewAuthenticPixels", &QueueCacheViewAuthenticPixels},
	{"RadialBlurImage", &RadialBlurImage},
	{"RadialBlurImageChannel", &RadialBlurImageChannel},
	{"RaiseImage", &RaiseImage},
	{"RandomChannelThresholdImage", &RandomChannelThresholdImage},
	{"RandomComponentGenesis", &RandomComponentGenesis},
	{"RandomComponentTerminus", &RandomComponentTerminus},
	{"RandomThresholdImage", &RandomThresholdImage},
	{"RandomThresholdImageChannel", &RandomThresholdImageChannel},
	{"ReacquireMemory", &ReacquireMemory},
	{"ReadBlob", &ReadBlob},
	{"ReadBlobByte", &ReadBlobByte},
	{"ReadBlobDouble", &ReadBlobDouble},
	{"ReadBlobFloat", &ReadBlobFloat},
	{"ReadBlobLong", &ReadBlobLong},
	{"ReadBlobLongLong", &ReadBlobLongLong},
	{"ReadBlobLSBLong", &ReadBlobLSBLong},
	{"ReadBlobLSBShort", &ReadBlobLSBShort},
	{"ReadBlobMSBLong", &ReadBlobMSBLong},
	{"ReadBlobMSBLongLong", &ReadBlobMSBLongLong},
	{"ReadBlobMSBShort", &ReadBlobMSBShort},
	{"ReadBlobShort", &ReadBlobShort},
	{"ReadBlobString", &ReadBlobString},
	{"ReadImage", &ReadImage},
	{"ReadImages", &ReadImages},
	{"ReadInlineImage", &ReadInlineImage},
	{"ReadStream", &ReadStream},
	{"RecolorImage", &RecolorImage},
	{"ReduceNoiseImage", &ReduceNoiseImage},
	{"ReferenceBlob", &ReferenceBlob},
	{"ReferenceImage", &ReferenceImage},
	{"ReferencePixelCache", &ReferencePixelCache},
	{"RegisterMagickInfo", &RegisterMagickInfo},
	{"RegisterStaticModules", &RegisterStaticModules},
	{"RegistryComponentGenesis", &RegistryComponentGenesis},
	{"RegistryComponentTerminus", &RegistryComponentTerminus},
	{"RelinquishAlignedMemory", &RelinquishAlignedMemory},
	{"RelinquishMagickMatrix", &RelinquishMagickMatrix},
	{"RelinquishMagickMemory", &RelinquishMagickMemory},
	{"RelinquishMagickResource", &RelinquishMagickResource},
	{"RelinquishSemaphoreInfo", &RelinquishSemaphoreInfo},
	{"RelinquishUniqueFileResource", &RelinquishUniqueFileResource},
	{"RemapImage", &RemapImage},
	{"RemapImages", &RemapImages},
	{"RemoteDisplayCommand", &RemoteDisplayCommand},
	{"RemoveDuplicateLayers", &RemoveDuplicateLayers},
	{"RemoveElementByValueFromLinkedList", &RemoveElementByValueFromLinkedList},
	{"RemoveElementFromLinkedList", &RemoveElementFromLinkedList},
	{"RemoveEntryFromHashmap", &RemoveEntryFromHashmap},
	{"RemoveFirstImageFromList", &RemoveFirstImageFromList},
	{"RemoveImageArtifact", &RemoveImageArtifact},
	{"RemoveImageFromList", &RemoveImageFromList},
	{"RemoveImageOption", &RemoveImageOption},
	{"RemoveImageProfile", &RemoveImageProfile},
	{"RemoveImageProperty", &RemoveImageProperty},
	{"RemoveImageRegistry", &RemoveImageRegistry},
	{"RemoveLastElementFromLinkedList", &RemoveLastElementFromLinkedList},
	{"RemoveLastImageFromList", &RemoveLastImageFromList},
	{"RemoveNodeByValueFromSplayTree", &RemoveNodeByValueFromSplayTree},
	{"RemoveNodeFromSplayTree", &RemoveNodeFromSplayTree},
	{"RemoveZeroDelayLayers", &RemoveZeroDelayLayers},
	{"ReplaceImageInList", &ReplaceImageInList},
	{"ReplaceImageInListReturnLast", &ReplaceImageInListReturnLast},
	{"ResampleImage", &ResampleImage},
	{"ResamplePixelColor", &ResamplePixelColor},
	{"ResetHashmapIterator", &ResetHashmapIterator},
	{"ResetImageArtifactIterator", &ResetImageArtifactIterator},
	{"ResetImageAttributeIterator", &ResetImageAttributeIterator},
	{"ResetImageOptionIterator", &ResetImageOptionIterator},
	{"ResetImageOptions", &ResetImageOptions},
	{"ResetImagePage", &ResetImagePage},
	{"ResetImageProfileIterator", &ResetImageProfileIterator},
	{"ResetImagePropertyIterator", &ResetImagePropertyIterator},
	{"ResetImageRegistryIterator", &ResetImageRegistryIterator},
	{"ResetLinkedListIterator", &ResetLinkedListIterator},
	{"ResetMagickMemory", &ResetMagickMemory},
	{"ResetSplayTree", &ResetSplayTree},
	{"ResetSplayTreeIterator", &ResetSplayTreeIterator},
	{"ResetStringInfo", &ResetStringInfo},
	{"ResetTimer", &ResetTimer},
	{"ResizeImage", &ResizeImage},
	{"ResizeMagickMemory", &ResizeMagickMemory},
	{"ResizeQuantumMemory", &ResizeQuantumMemory},
	{"ResourceComponentGenesis", &ResourceComponentGenesis},
	{"ResourceComponentTerminus", &ResourceComponentTerminus},
	{"ReverseImageList", &ReverseImageList},
	{"RGBTransformImage", &RGBTransformImage},
	{"RollImage", &RollImage},
	{"RotateImage", &RotateImage},
	{"SampleImage", &SampleImage},
	{"ScaleGeometryKernelInfo", &ScaleGeometryKernelInfo},
	{"ScaleImage", &ScaleImage},
	{"ScaleKernelInfo", &ScaleKernelInfo},
	{"ScaleResampleFilter", &ScaleResampleFilter},
	{"SeedPseudoRandomGenerator", &SeedPseudoRandomGenerator},
	{"SeekBlob", &SeekBlob},
	{"SegmentImage", &SegmentImage},
	{"SelectiveBlurImage", &SelectiveBlurImage},
	{"SelectiveBlurImageChannel", &SelectiveBlurImageChannel},
	{"SemaphoreComponentGenesis", &SemaphoreComponentGenesis},
	{"SemaphoreComponentTerminus", &SemaphoreComponentTerminus},
	{"SeparateImageChannel", &SeparateImageChannel},
	{"SeparateImages", &SeparateImages},
	{"SepiaToneImage", &SepiaToneImage},
	{"SetBlobExempt", &SetBlobExempt},
	{"SetBlobExtent", &SetBlobExtent},
	{"SetCacheThreshold", &SetCacheThreshold},
	{"SetCacheViewPixels", &SetCacheViewPixels},
	{"SetCacheViewStorageClass", &SetCacheViewStorageClass},
	{"SetCacheViewVirtualPixelMethod", &SetCacheViewVirtualPixelMethod},
	{"SetClientName", &SetClientName},
	{"SetClientPath", &SetClientPath},
	{"SetErrorHandler", &SetErrorHandler},
	{"SetExceptionInfo", &SetExceptionInfo},
	{"SetFatalErrorHandler", &SetFatalErrorHandler},
	{"SetGeometry", &SetGeometry},
	{"SetGeometryInfo", &SetGeometryInfo},
	{"SetImage", &SetImage},
	{"SetImageAlphaChannel", &SetImageAlphaChannel},
	{"SetImageArtifact", &SetImageArtifact},
	{"SetImageAttribute", &SetImageAttribute},
	{"SetImageBackgroundColor", &SetImageBackgroundColor},
	{"SetImageChannelDepth", &SetImageChannelDepth},
	{"SetImageChannels", &SetImageChannels},
	{"SetImageClipMask", &SetImageClipMask},
	{"SetImageColor", &SetImageColor},
	{"SetImageColorspace", &SetImageColorspace},
	{"SetImageDepth", &SetImageDepth},
	{"SetImageExtent", &SetImageExtent},
	{"SetImageInfo", &SetImageInfo},
	{"SetImageInfoBlob", &SetImageInfoBlob},
	{"SetImageInfoFile", &SetImageInfoFile},
	{"SetImageInfoProgressMonitor", &SetImageInfoProgressMonitor},
	{"SetImageList", &SetImageList},
	{"SetImageMask", &SetImageMask},
	{"SetImageOpacity", &SetImageOpacity},
	{"SetImageOption", &SetImageOption},
	{"SetImagePixels", &SetImagePixels},
	{"SetImageProfile", &SetImageProfile},
	{"SetImageProgressMonitor", &SetImageProgressMonitor},
	{"SetImageProperty", &SetImageProperty},
	{"SetImageRegistry", &SetImageRegistry},
	{"SetImageStorageClass", &SetImageStorageClass},
	{"SetImageType", &SetImageType},
	{"SetImageViewDescription", &SetImageViewDescription},
	{"SetImageViewIterator", &SetImageViewIterator},
	{"SetImageViewThreads", &SetImageViewThreads},
	{"SetImageVirtualPixelMethod", &SetImageVirtualPixelMethod},
	{"SetLogEventMask", &SetLogEventMask},
	{"SetLogFormat", &SetLogFormat},
	{"SetLogName", &SetLogName},
	{"SetMagickInfo", &SetMagickInfo},
	{"SetMagickMemoryMethods", &SetMagickMemoryMethods},
	{"SetMagickPrecision", &SetMagickPrecision},
	{"SetMagickRegistry", &SetMagickRegistry},
	{"SetMagickResourceLimit", &SetMagickResourceLimit},
	{"SetMonitorHandler", &SetMonitorHandler},
	{"SetPixelCacheMethods", &SetPixelCacheMethods},
	{"SetPixelCacheVirtualMethod", &SetPixelCacheVirtualMethod},
	{"SetQuantumAlphaType", &SetQuantumAlphaType},
	{"SetQuantumDepth", &SetQuantumDepth},
	{"SetQuantumEndian", &SetQuantumEndian},
	{"SetQuantumFormat", &SetQuantumFormat},
	{"SetQuantumImageType", &SetQuantumImageType},
	{"SetQuantumMinIsWhite", &SetQuantumMinIsWhite},
	{"SetQuantumPack", &SetQuantumPack},
	{"SetQuantumPad", &SetQuantumPad},
	{"SetQuantumQuantum", &SetQuantumQuantum},
	{"SetQuantumScale", &SetQuantumScale},
	{"SetRandomKey", &SetRandomKey},
	{"SetRandomSecretKey", &SetRandomSecretKey},
	{"SetRandomTrueRandom", &SetRandomTrueRandom},
	{"SetResampleFilter", &SetResampleFilter},
	{"SetResampleFilterInterpolateMethod", &SetResampleFilterInterpolateMethod},
	{"SetResampleFilterVirtualPixelMethod", &SetResampleFilterVirtualPixelMethod},
	{"SetSignatureDigest", &SetSignatureDigest},
	{"SetStreamInfoClientData", &SetStreamInfoClientData},
	{"SetStreamInfoMap", &SetStreamInfoMap},
	{"SetStreamInfoStorageType", &SetStreamInfoStorageType},
	{"SetStringInfo", &SetStringInfo},
	{"SetStringInfoDatum", &SetStringInfoDatum},
	{"SetStringInfoLength", &SetStringInfoLength},
	{"SetStringInfoPath", &SetStringInfoPath},
	{"SetWarningHandler", &SetWarningHandler},
	{"SetXMLTreeAttribute", &SetXMLTreeAttribute},
	{"SetXMLTreeContent", &SetXMLTreeContent},
	{"ShadeImage", &ShadeImage},
	{"ShadowImage", &ShadowImage},
	{"SharpenImage", &SharpenImage},
	{"SharpenImageChannel", &SharpenImageChannel},
	{"ShaveImage", &ShaveImage},
	{"ShearImage", &ShearImage},
	{"ShearRotateImage", &ShearRotateImage},
	{"ShiftImageList", &ShiftImageList},
	{"ShowKernelInfo", &ShowKernelInfo},
	{"SigmoidalContrastImage", &SigmoidalContrastImage},
	{"SigmoidalContrastImageChannel", &SigmoidalContrastImageChannel},
	{"SignatureImage", &SignatureImage},
	{"SimilarityImage", &SimilarityImage},
	{"SimilarityMetricImage", &SimilarityMetricImage},
	{"SizeBlob", &SizeBlob},
	{"SketchImage", &SketchImage},
	{"SmushImages", &SmushImages},
	{"SolarizeImage", &SolarizeImage},
	{"SolarizeImageChannel", &SolarizeImageChannel},
	{"SortColormapByIntensity", &SortColormapByIntensity},
	{"SparseColorImage", &SparseColorImage},
	{"SpliceImage", &SpliceImage},
	{"SpliceImageIntoList", &SpliceImageIntoList},
	{"SpliceImageList", &SpliceImageList},
	{"SplitImageList", &SplitImageList},
	{"SplitStringInfo", &SplitStringInfo},
	{"SpreadImage", &SpreadImage},
	{"sRGBCompandor", &SRGBCompandor},
	{"StartTimer", &StartTimer},
	{"StatisticImage", &StatisticImage},
	{"StatisticImageChannel", &StatisticImageChannel},
	{"SteganoImage", &SteganoImage},
	{"StereoAnaglyphImage", &StereoAnaglyphImage},
	{"StereoImage", &StereoImage},
	{"StreamImage", &StreamImage},
	{"StringInfoToHexString", &StringInfoToHexString},
	{"StringInfoToString", &StringInfoToString},
	{"StringToArgv", &StringToArgv},
	{"StringToArrayOfDoubles", &StringToArrayOfDoubles},
	{"StringToken", &StringToken},
	{"StringToList", &StringToList},
	{"StringToStringInfo", &StringToStringInfo},
	{"Strip", &Strip},
	{"StripImage", &StripImage},
	{"StripString", &StripString},
	{"SubstituteString", &SubstituteString},
	{"SwirlImage", &SwirlImage},
	{"SyncAuthenticPixelCacheNexus", &SyncAuthenticPixelCacheNexus},
	{"SyncAuthenticPixels", &SyncAuthenticPixels},
	{"SyncCacheView", &SyncCacheView},
	{"SyncCacheViewAuthenticPixels", &SyncCacheViewAuthenticPixels},
	{"SyncCacheViewPixels", &SyncCacheViewPixels},
	{"SyncImage", &SyncImage},
	{"SyncImageList", &SyncImageList},
	{"SyncImagePixels", &SyncImagePixels},
	{"SyncImageProfiles", &SyncImageProfiles},
	{"SyncImageSettings", &SyncImageSettings},
	{"SyncImagesSettings", &SyncImagesSettings},
	{"SyncNextImageInList", &SyncNextImageInList},
	{"SystemCommand", &SystemCommand},
	{"TellBlob", &TellBlob},
	{"TemporaryFilename", &TemporaryFilename},
	{"TextureImage", &TextureImage},
	{"ThresholdImage", &ThresholdImage},
	{"ThresholdImageChannel", &ThresholdImageChannel},
	{"ThrowException", &ThrowException},
	{"ThrowMagickException", &ThrowMagickException},
	{"ThrowMagickExceptionList", &ThrowMagickExceptionList},
	{"ThumbnailImage", &ThumbnailImage},
	{"TintImage", &TintImage},
	{"Tokenizer", &Tokenizer},
	{"TransferImageViewIterator", &TransferImageViewIterator},
	{"TransformColorspace", &TransformColorspace},
	{"TransformHSL", &TransformHSL},
	{"TransformImage", &TransformImage},
	{"TransformImageColorspace", &TransformImageColorspace},
	{"TransformImages", &TransformImages},
	{"TransformRGBImage", &TransformRGBImage},
	{"TranslateText", &TranslateText},
	{"TransparentImage", &TransparentImage},
	{"TransparentPaintImage", &TransparentPaintImage},
	{"TransparentPaintImageChroma", &TransparentPaintImageChroma},
	{"TransposeImage", &TransposeImage},
	{"TransverseImage", &TransverseImage},
	{"TrimImage", &TrimImage},
	{"TypeComponentGenesis", &TypeComponentGenesis},
	{"TypeComponentTerminus", &TypeComponentTerminus},
	{"UniqueImageColors", &UniqueImageColors},
	{"UnityAddKernelInfo", &UnityAddKernelInfo},
	{"UnlockSemaphoreInfo", &UnlockSemaphoreInfo},
	{"UnmapBlob", &UnmapBlob},
	{"UnregisterMagickInfo", &UnregisterMagickInfo},
	{"UnregisterStaticModules", &UnregisterStaticModules},
	{"UnsharpMaskImage", &UnsharpMaskImage},
	{"UnsharpMaskImageChannel", &UnsharpMaskImageChannel},
	{"UnshiftImageList", &UnshiftImageList},
	{"UpdateImageViewIterator", &UpdateImageViewIterator},
	{"UpdateSignature", &UpdateSignature},
	{"ValidateColormapIndex", &ValidateColormapIndex},
	{"VignetteImage", &VignetteImage},
	{"WaveImage", &WaveImage},
	{"WhiteThresholdImage", &WhiteThresholdImage},
	{"WhiteThresholdImageChannel", &WhiteThresholdImageChannel},
	{"WriteBlob", &WriteBlob},
	{"WriteBlobByte", &WriteBlobByte},
	{"WriteBlobFloat", &WriteBlobFloat},
	{"WriteBlobLong", &WriteBlobLong},
	{"WriteBlobLSBLong", &WriteBlobLSBLong},
	{"WriteBlobLSBShort", &WriteBlobLSBShort},
	{"WriteBlobMSBLong", &WriteBlobMSBLong},
	{"WriteBlobMSBLongLong", &WriteBlobMSBLongLong},
	{"WriteBlobMSBShort", &WriteBlobMSBShort},
	{"WriteBlobShort", &WriteBlobShort},
	{"WriteBlobString", &WriteBlobString},
	{"WriteImage", &WriteImage},
	{"WriteImages", &WriteImages},
	{"WriteStream", &WriteStream},
	{"XAnimateBackgroundImage", &XAnimateBackgroundImage},
	{"XAnimateImages", &XAnimateImages},
	{"XAnnotateImage", &XAnnotateImage},
	{"XBestFont", &XBestFont},
	{"XBestIconSize", &XBestIconSize},
	{"XBestPixel", &XBestPixel},
	{"XBestVisualInfo", &XBestVisualInfo},
	{"XCheckDefineCursor", &XCheckDefineCursor},
	{"XCheckRefreshWindows", &XCheckRefreshWindows},
	{"XClientMessage", &XClientMessage},
	{"XColorBrowserWidget", &XColorBrowserWidget},
	{"XCommandWidget", &XCommandWidget},
	{"XComponentGenesis", &XComponentGenesis},
	{"XComponentTerminus", &XComponentTerminus},
	{"XConfigureImageColormap", &XConfigureImageColormap},
	{"XConfirmWidget", &XConfirmWidget},
	{"XConstrainWindowPosition", &XConstrainWindowPosition},
	{"XDelay", &XDelay},
	{"XDestroyResourceInfo", &XDestroyResourceInfo},
	{"XDestroyWindowColors", &XDestroyWindowColors},
	{"XDialogWidget", &XDialogWidget},
	{"XDisplayBackgroundImage", &XDisplayBackgroundImage},
	{"XDisplayImage", &XDisplayImage},
	{"XDisplayImageInfo", &XDisplayImageInfo},
	{"XDrawImage", &XDrawImage},
	{"XError", &XError},
	{"XFileBrowserWidget", &XFileBrowserWidget},
	{"XFontBrowserWidget", &XFontBrowserWidget},
	{"XFreeResources", &XFreeResources},
	{"XFreeStandardColormap", &XFreeStandardColormap},
	{"XGetAnnotateInfo", &XGetAnnotateInfo},
	{"XGetImportInfo", &XGetImportInfo},
	{"XGetMapInfo", &XGetMapInfo},
	{"XGetPixelPacket", &XGetPixelPacket},
	{"XGetResourceClass", &XGetResourceClass},
	{"XGetResourceDatabase", &XGetResourceDatabase},
	{"XGetResourceInfo", &XGetResourceInfo},
	{"XGetResourceInstance", &XGetResourceInstance},
	{"XGetScreenDensity", &XGetScreenDensity},
	{"XGetWindowColor", &XGetWindowColor},
	{"XGetWindowInfo", &XGetWindowInfo},
	{"XHighlightEllipse", &XHighlightEllipse},
	{"XHighlightLine", &XHighlightLine},
	{"XHighlightRectangle", &XHighlightRectangle},
	{"XImportImage", &XImportImage},
	{"XInfoWidget", &XInfoWidget},
	{"XInitializeWindows", &XInitializeWindows},
	{"XListBrowserWidget", &XListBrowserWidget},
	{"XMagickProgressMonitor", &XMagickProgressMonitor},
	{"XMakeCursor", &XMakeCursor},
	{"XMakeImage", &XMakeImage},
	{"XMakeMagnifyImage", &XMakeMagnifyImage},
	{"XMakeStandardColormap", &XMakeStandardColormap},
	{"XMakeWindow", &XMakeWindow},
	{"XMenuWidget", &XMenuWidget},
	{"XMLTreeInfoToXML", &XMLTreeInfoToXML},
	{"XNoticeWidget", &XNoticeWidget},
	{"XPreferencesWidget", &XPreferencesWidget},
	{"XProgressMonitorWidget", &XProgressMonitorWidget},
	{"XQueryColorDatabase", &XQueryColorDatabase},
	{"XQueryPosition", &XQueryPosition},
	{"XRefreshWindow", &XRefreshWindow},
	{"XRemoteCommand", &XRemoteCommand},
	{"XRetainWindowColors", &XRetainWindowColors},
	{"XSetCursorState", &XSetCursorState},
	{"XSetWindows", &XSetWindows},
	{"XTextViewWidget", &XTextViewWidget},
	{"XUserPreferences", &XUserPreferences},
	{"XWarning", &XWarning},
	{"XWindowByID", &XWindowByID},
	{"XWindowByName", &XWindowByName},
	{"XWindowByProperty", &XWindowByProperty},
	{"ZeroKernelNans", &ZeroKernelNans},
	{"ZLIBEncodeImage", &ZLIBEncodeImage},
	{"ZoomImage", &ZoomImage},
}

var allData = Data{
	{"BackgroundColor", (*byte)(nil)},
	{"BorderColor", (*byte)(nil)},
	{"DefaultTileFrame", (*byte)(nil)},
	{"DefaultTileGeometry", (*byte)(nil)},
	{"DefaultTileLabel", (*byte)(nil)},
	{"ForegroundColor", (*byte)(nil)},
	{"LoadImagesTag", (*byte)(nil)},
	{"LoadImageTag", (*byte)(nil)},
	{"MatteColor", (*byte)(nil)},
	{"PSDensityGeometry", (*byte)(nil)},
	{"PSPageGeometry", (*byte)(nil)},
	{"SaveImagesTag", (*byte)(nil)},
	{"SaveImageTag", (*byte)(nil)},
	{"DefaultResolution", (*float64)(nil)},
}
