// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package imagemagick provides API definitions for
//accessing CORE_RL_magick_.dll.
package core

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside/types"
)

type Quantum uint16         // byte uint16 uint32
type MagickRealType float32 // float64
const (
	MaxNumberFonts = 11
	MaxNumberPens  = 11
	MaxTextExtent  = 4096
)

type (
	CompositeMethod Fixme // Not in bin distribution .h
	DIR             Fixme
	Dirent          Fixme
	FILE            Fixme
	Stat            Fixme
	Time            Long // 32/64 dep
	WString         Fixme
)

// X
type (
	Atom           UnsignedLong
	Colormap       UnsignedLong
	Cursor         UnsignedLong
	Display        struct{}
	LinkedListInfo struct{}
	Pixmap         UnsignedLong
	PolicyInfo     struct{}
	Status         int
	// Visual            struct{}
	Window            UnsignedLong
	XClassHint        struct{}
	XColor            struct{}
	XErrorEvent       struct{}
	XEvent            struct{} // union
	XFontStruct       struct{}
	XImage            struct{}
	XPoint            struct{ x, y Short }
	XrmDatabase       struct{} // *struct{} in X - we make
	XSegment          struct{ x1, y1, x2, y2 Short }
	XStandardColormap struct{}
	XVisualInfo       struct{}
	XWMHints          struct{}
)
type (
	// Placeholder type to be rectified.
	Fixme            uintptr
	Char             byte // int8
	Enum             int
	IndexPacket      Quantum
	Long             int    // TODO(t):long on gcc vs msc 32 vs 64
	MagickOffsetType int64  // TODO(t):size on gcc vs msc 32 vs 64
	MagickSizeType   uint64 // TODO(t):size on gcc vs msc 32 vs 64
	MagickStatusType uint
	Off              Long
	Short            uint16 // TODO(t):short on gcc vs msc 32 vs 64
	Size             uint   // TODO(t):size_t on gcc vs msc 32 vs 64
	SSize            int    // TODO(t):ssize_t on gcc vs msc 32 vs 64
	UnsignedShort    uint16 // TODO(t):unsigned short on gcc vs msc 32 vs 64
	UnsignedLong     uint   // TODO(t):unsigned long on gcc vs msc 32 vs 64
	MagickThreadKey  *Size
)

// Opaque types
type (
	Ascii85Info    struct{}
	BlobInfo       struct{}
	Cache          *struct{}
	CacheView      struct{}
	FxInfo         struct{}
	GC             *struct{}
	HashmapInfo    struct{}
	ImageView      struct{}
	LogInfo        struct{}
	MimeInfo       struct{}
	RandomInfo     struct{}
	ResampleFilter struct{}
	ResizeFilter   struct{}
	SemaphoreInfo  struct{}
	SignatureInfo  struct{}
	SplayTreeInfo  struct{}
	StreamInfo     struct{}
	ThresholdMap   struct{}
	TokenInfo      struct{}
	Void           struct{}
	XMLTreeInfo    struct{}
)
type (
	AcquireMemoryHandler          func(Size) *Void
	DecodeImageHandler            func(*ImageInfo, *ExceptionInfo) *Image
	DestroyMemoryHandler          func(*Void)
	DuplexTransferImageViewMethod func(*ImageView, *ImageView, *ImageView, SSize, int, *Void) MagickBooleanType
	EncodeImageHandler            func(*ImageInfo, *Image) bool
	ErrorHandler                  func(ExceptionType, string, string)
	FatalErrorHandler             func(ExceptionType, string, string)
	GetImageViewMethod            func(*ImageView, SSize, int, *Void) MagickBooleanType
	IsImageFormatHandler          func(*byte, Size) MagickBooleanType
	MagickCommand                 func(*ImageInfo, int, []string, []string, *ExceptionInfo) bool
	MagickProgressMonitor         func(string, MagickOffsetType, MagickSizeType, *Void) bool
	MonitorHandler                func(string, MagickOffsetType, MagickSizeType, *ExceptionInfo) bool
	ResizeMemoryHandler           func(*Void, Size) *Void
	SetImageViewMethod            func(*ImageView, SSize, int, *Void) MagickBooleanType
	StreamHandler                 func(*Image, *Void, Size) Size
	TransferImageViewMethod       func(*ImageView, *ImageView, SSize, int, *Void) MagickBooleanType
	UpdateImageViewMethod         func(*ImageView, SSize, int, *Void) MagickBooleanType
	WarningHandler                func(ExceptionType, string, string)
)

// TODO(t): image *Image; images *Image; images **Image distinctions
var AccelerateConvolveImage func(*Image, *KernelInfo, *Image, *ExceptionInfo) bool

func (i *Image) AccelerateConvolve(k *KernelInfo, i2 *Image, e *ExceptionInfo) bool {
	return AccelerateConvolveImage(i, k, i2, e)
}

var AcquireFxInfo func(i *Image, expression string) *FxInfo

func (i *Image) AcquireFxInfo(expression string) *FxInfo { return AcquireFxInfo(i, expression) }

// Deprecated
var AcquireImagePixels func(i *Image, x, y SSize, columns, rows Size, exception *ExceptionInfo) *PixelPacket

// Deprecated
func (i *Image) AcquireImagePixels(x, y SSize, columns, rows Size, exception *ExceptionInfo) *PixelPacket {
	return AcquireImagePixels(i, x, y, columns, rows, exception)
}

// Deprecated
var AcquireIndexes func(i *Image) *IndexPacket

// Deprecated
func (i *Image) AcquireIndexes() *IndexPacket { return AcquireIndexes(i) }

// Deprecated
var AcquireOneMagickPixel func(i *Image, x, y SSize, exception *ExceptionInfo) MagickPixelPacket

// Deprecated
func (i *Image) AcquireOneMagickPixel(x, y SSize, exception *ExceptionInfo) MagickPixelPacket {
	return AcquireOneMagickPixel(i, x, y, exception)
}

// Deprecated
var AcquireOnePixel func(i *Image, x, y SSize, exception *ExceptionInfo) PixelPacket

// Deprecated
func (i *Image) AcquireOnePixel(x, y SSize, exception *ExceptionInfo) PixelPacket {
	return AcquireOnePixel(i, x, y, exception)
}

var AcquireOneVirtualPixel func(i *Image, virtualPixelMethod VirtualPixelMethod, x, y SSize, exception *ExceptionInfo) PixelPacket

func (i *Image) AcquireOneVirtualPixel(virtualPixelMethod VirtualPixelMethod, x, y SSize, exception *ExceptionInfo) PixelPacket {
	return AcquireOneVirtualPixel(i, virtualPixelMethod, x, y, exception)
}

var AcquireResampleFilter func(i *Image, exception *ExceptionInfo) *ResampleFilter

func (i *Image) AcquireResampleFilter(exception *ExceptionInfo) *ResampleFilter {
	return AcquireResampleFilter(i, exception)
}

var AcquireResizeFilter func(i *Image, filter FilterTypes, blur MagickRealType, cylindrical bool, exception *ExceptionInfo) *ResizeFilter

func (i *Image) AcquireResizeFilter(filter FilterTypes, blur MagickRealType, cylindrical bool, exception *ExceptionInfo) *ResizeFilter {
	return AcquireResizeFilter(i, filter, blur, cylindrical, exception)
}

var AdaptiveBlurImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveBlur(radius, sigma float64, exception *ExceptionInfo) *Image {
	return AdaptiveBlurImage(i, radius, sigma, exception)
}

var AdaptiveBlurImageChannel func(i *Image, channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveBlurChannel(channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return AdaptiveBlurImageChannel(i, channel, radius, sigma, exception)
}

var AdaptiveResizeImage func(i *Image, columns, rows Size, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveResize(columns, rows Size, exception *ExceptionInfo) *Image {
	return AdaptiveResizeImage(i, columns, rows, exception)
}

var AdaptiveSharpenImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveSharpen(radius, sigma float64, exception *ExceptionInfo) *Image {
	return AdaptiveSharpenImage(i, radius, sigma, exception)
}

var AdaptiveSharpenImageChannel func(i *Image, channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveSharpenChannel(channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return AdaptiveSharpenImageChannel(i, channel, radius, sigma, exception)
}

var AdaptiveThresholdImage func(i *Image, width, height Size, offset SSize, exception *ExceptionInfo) *Image

func (i *Image) AdaptiveThreshold(width, height Size, offset SSize, exception *ExceptionInfo) *Image {
	return AdaptiveThresholdImage(i, width, height, offset, exception)
}

var AddNoiseImage func(i *Image, noiseType NoiseType, exception *ExceptionInfo) *Image

func (i *Image) AddNoise(noiseType NoiseType, exception *ExceptionInfo) *Image {
	return AddNoiseImage(i, noiseType, exception)
}

var AddNoiseImageChannel func(i *Image, channel ChannelType, noiseType NoiseType, exception *ExceptionInfo) *Image

func (i *Image) AddNoiseChannel(channel ChannelType, noiseType NoiseType, exception *ExceptionInfo) *Image {
	return AddNoiseImageChannel(i, channel, noiseType, exception)
}

var AffineTransformImage func(i *Image, affineMatrix *AffineMatrix, exception *ExceptionInfo) *Image

func (i *Image) AffineTransform(affineMatrix *AffineMatrix, exception *ExceptionInfo) *Image {
	return AffineTransformImage(i, affineMatrix, exception)
}

// Deprecated
var AllocateImageColormap func(i *Image, colors Size) bool

// Deprecated
func (i *Image) AllocateColormap(colors Size) bool {
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
func (i *Image) Ascii85Flush()           { Ascii85Flush(i) }
func (i *Image) Ascii85Initialize()      { Ascii85Initialize(i) }

// Deprecated
func (i *Image) AverageImages(exception *ExceptionInfo) *Image { return AverageImages(i, exception) }
func (i *Image) Bilevel(threshold float64) bool                { return BilevelImage(i, threshold) }
func (i *Image) BilevelChannel(channel ChannelType, threshold float64) bool {
	return BilevelImageChannel(i, channel, threshold)
}
func (i *Image) BlackThreshold(threshold string) bool { return BlackThresholdImage(i, threshold) }
func (i *Image) Blur(radius, sigma float64, exception *ExceptionInfo) *Image {
	return BlurImage(i, radius, sigma, exception)
}
func (i *Image) BlurChannel(channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return BlurImageChannel(i, channel, radius, sigma, exception)
}
func (i *Image) Border(borderInfo *RectangleInfo, exception *ExceptionInfo) *Image {
	return BorderImage(i, borderInfo, exception)
}

var Ascii85Flush func(i *Image)
var Ascii85Initialize func(i *Image)

// Deprecated
var AverageImages func(i *Image, exception *ExceptionInfo) *Image
var BilevelImage func(i *Image, threshold float64) bool
var BilevelImageChannel func(i *Image, channel ChannelType, threshold float64) bool
var BlackThresholdImage func(i *Image, threshold string) bool
var BlurImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image
var BlurImageChannel func(i *Image, channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image
var BorderImage func(i *Image, borderInfo *RectangleInfo, exception *ExceptionInfo) *Image
var CatchImageException func(i *Image) ExceptionType
var ChannelImage func(i *Image, channel ChannelType) uint

// Deprecated
var ChannelThresholdImage func(i *Image, level string) uint
var CharcoalImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image
var ChopImage func(i *Image, chopInfo *RectangleInfo, exception *ExceptionInfo) *Image
var ClipImage func(i *Image) bool
var ClipImagePath func(i *Image, pathname string, inside bool) bool

// Deprecated
var ClipPathImage func(i *Image, pathname string, inside bool) bool
var CloneImage func(i *Image, columns, rows Size, orphan bool, exception *ExceptionInfo) *Image
var CloneImageArtifacts func(i *Image, cloneImage *Image) bool

// Deprecated
var CloneImageAttributes func(i *Image, cloneImage *Image) bool
var CloneImageProfiles func(i *Image, cloneImage *Image) bool
var CloneImageProperties func(i *Image, cloneImage *Image) bool
var CloseBlob func(i *Image) bool
var ClutImage func(i *Image, clutImage *Image) bool
var ClutImageChannel func(i *Image, channel ChannelType, clutImage *Image) bool
var CoalesceImages func(i *Image, exception *ExceptionInfo) *Image

// Deprecated
var ColorFloodfillImage func(i *Image, drawInfo *DrawInfo, target PixelPacket, xOffset, yOffset SSize, method PaintMethod) bool
var ColorizeImage func(i *Image, opacity string, colorize PixelPacket, exception *ExceptionInfo) *Image
var CombineImages func(i *Image, channel ChannelType, exception *ExceptionInfo) *Image
var CompareImageChannels func(i *Image, reconstructImage *Image, channel ChannelType, metric MetricType, distortion *float64, exception *ExceptionInfo) *Image
var CompareImageLayers func(i *Image, method ImageLayerMethod, exception *ExceptionInfo) *Image
var CompareImages func(i *Image, reconstructImage *Image, metric MetricType, distortion *float64, exception *ExceptionInfo) *Image
var CompositeImage func(i *Image, compose CompositeOperator, compositeImage *Image, xOffset, yOffset SSize) bool
var CompositeImageChannel func(i *Image, channel ChannelType, compose CompositeOperator, compositeImage *Image, xOffset, yOffset SSize) bool
var CompositeLayers func(i *Image, compose CompositeOperator, source *Image, xOffset, yOffset SSize, exception *ExceptionInfo)
var CompressImageColormap func(i *Image)
var ContrastImage func(i *Image, sharpen bool) bool
var ContrastStretchImage func(i *Image, levels string) bool
var ContrastStretchImageChannel func(i *Image, channel ChannelType, blackPoint, whitePoint float64) bool
var ConvolveImage func(i *Image, order Size, kernel *float64, exception *ExceptionInfo) *Image
var ConvolveImageChannel func(i *Image, channel ChannelType, order Size, kernel *float64, exception *ExceptionInfo) *Image
var CropImage func(i *Image, geometry *RectangleInfo, exception *ExceptionInfo) *Image
var CycleColormapImage func(i *Image, displace SSize) bool
var DefineImageArtifact func(i *Image, artifact string) bool
var DefineImageProperty func(i *Image, property string) bool
var DeleteImageArtifact func(i *Image, artifact string) bool

// Deprecated
var DeleteImageAttribute func(i *Image, key string) bool
var DeleteImageProfile func(i *Image, name string) bool
var DeleteImageProperty func(i *Image, property string) bool

// Deprecated
var DescribeImage func(i *Image, file *FILE, verbose bool) bool
var DespeckleImage func(i *Image, exception *ExceptionInfo) *Image
var DestroyBlob func(i *Image)
var DestroyImage func(i *Image) *Image
var DestroyImageArtifacts func(i *Image)

// Deprecated
var DestroyImageAttributes func(i *Image)
var DestroyImagePixels func(i *Image)
var DestroyImageProfiles func(i *Image)
var DestroyImageProperties func(i *Image)

// Deprecated
var DestroyImages func(i *Image)
var DisassociateImageStream func(i *Image)

// Deprecated
var DispatchImage func(i *Image, xOffset, yOffset SSize, columns, rows Size, map_ string, type_ StorageType, pixels *Void, exception *ExceptionInfo) uint
var DisposeImages func(i *Image, exception *ExceptionInfo) *Image
var DistortImage func(i *Image, method DistortImageMethod, numberArguments Size, arguments *float64, bestfit bool, exception *ExceptionInfo) *Image
var DrawAffineImage func(i *Image, source *Image, affine *AffineMatrix) bool
var DrawClipPath func(i *Image, drawInfo *DrawInfo, name string) bool
var DrawGradientImage func(i *Image, drawInfo *DrawInfo) bool
var DrawImage func(i *Image, drawInfo *DrawInfo) bool
var DrawPatternPath func(i *Image, drawInfo *DrawInfo, name string, pattern **Image) bool
var DrawPrimitive func(i *Image, drawInfo *DrawInfo, primitiveInfo *PrimitiveInfo) bool
var EdgeImage func(i *Image, radius float64, exception *ExceptionInfo) *Image
var EmbossImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image
var EnhanceImage func(i *Image, exception *ExceptionInfo) *Image
var EOFBlob func(i *Image) int
var EqualizeImage func(i *Image) bool
var EqualizeImageChannel func(i *Image, channel ChannelType) bool
var EvaluateImage func(i *Image, op MagickEvaluateOperator, value float64, exception *ExceptionInfo) bool
var EvaluateImageChannel func(i *Image, channel ChannelType, op MagickEvaluateOperator, value float64, exception *ExceptionInfo) bool
var ExcerptImage func(i *Image, geometry *RectangleInfo, exception *ExceptionInfo) *Image
var ExportImagePixels func(i *Image, xOffset, yOffset SSize, columns, rows Size, map_ string, type_ StorageType, pixels *Void, exception *ExceptionInfo) bool
var ExportQuantumPixels func(i *Image, quantumInfo *QuantumInfo, quantumType QuantumType, pixels *byte) bool
var ExtentImage func(i *Image, geometry *RectangleInfo, exception *ExceptionInfo) *Image
var FileToImage func(i *Image, filename string) bool

// Deprecated
var FlattenImages func(i *Image, exception *ExceptionInfo) *Image
var FlipImage func(i *Image, exception *ExceptionInfo) *Image
var FloodfillPaintImage func(i *Image, channel ChannelType, drawInfo *DrawInfo, target *MagickPixelPacket, xOffset, yOffset SSize, invert bool) bool
var FlopImage func(i *Image, exception *ExceptionInfo) *Image

// Deprecated
var FormatImageAttribute func(i *Image, key, format string, va ...VArg) bool
var FormatImageAttributeList func(i *Image, key, format string, operands VAList) bool
var FormatImageProperty func(i *Image, property, format string, va ...VArg) bool
var FormatImagePropertyList func(i *Image, property, format string, operands VAList) bool
var FormatMagickCaption func(i *Image, drawInfo *DrawInfo, caption string, metrics *TypeMetric) SSize
var FrameImage func(i *Image, frameInfo *FrameInfo, exception *ExceptionInfo) *Image
var FuzzyColorCompare func(i *Image, p, q *PixelPacket) bool
var FuzzyOpacityCompare func(i *Image, p, q *PixelPacket) bool
var FxImage func(i *Image, expression string, exception *ExceptionInfo) *Image
var FxImageChannel func(i *Image, channel ChannelType, expression string, exception *ExceptionInfo) *Image
var GammaImage func(i *Image, level string) bool
var GammaImageChannel func(i *Image, channel ChannelType, gamma float64) bool
var GaussianBlurImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image
var GaussianBlurImageChannel func(i *Image, channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image
var GetBlobError func(i *Image) bool
var GetBlobFileHandle func(i *Image) *FILE
var GetBlobSize func(i *Image) MagickSizeType
var GetBlobStreamData func(i *Image) *byte
var GetBlobStreamHandler func(i *Image) StreamHandler
var GetImageArtifact func(i *Image, artifact string) string

// Deprecated
var GetImageAttribute func(i *Image, key string) *ImageAttribute
var GetImageBoundingBox func(i *Image, exception *ExceptionInfo) RectangleInfo
var GetImageChannelDepth func(i *Image, channel ChannelType, exception *ExceptionInfo) Size
var GetImageChannelDistortion func(i *Image, reconstructImage *Image, channel ChannelType, metric MetricType, distortion *float64, exception *ExceptionInfo) bool
var GetImageChannelExtrema func(i *Image, channel ChannelType, minima, maxima *Size, exception *ExceptionInfo) bool
var GetImageChannelMean func(i *Image, channel ChannelType, mean, standardDeviation *float64, exception *ExceptionInfo) bool
var GetImageChannelRange func(i *Image, channel ChannelType, minima, maxima *float64, exception *ExceptionInfo) bool
var GetImageChannelStatistics func(i *Image, exception *ExceptionInfo) *ChannelStatistics
var GetImageClipMask func(i *Image, exception *ExceptionInfo) *Image

// Deprecated
var GetImageClippingPathAttribute func(i *Image) *ImageAttribute
var GetImageDepth func(i *Image, exception *ExceptionInfo) Size
var GetImageDistortion func(i *Image, reconstructImage *Image, metric MetricType, distortion *float64, exception *ExceptionInfo) bool
var GetImageDynamicThreshold func(i *Image, clusterThreshold, smoothThreshold float64, exception *ExceptionInfo) MagickPixelPacket
var GetImageException func(i *Image, exception *ExceptionInfo)
var GetImageExtrema func(i *Image, minima, maxima *Size, exception *ExceptionInfo) bool

// Deprecated
var GetImageGeometry func(i *Image, geometry string, sizeToFit uint, regionInfo *RectangleInfo) int
var GetImageHistogram func(i *Image, numberColors *Size, exception *ExceptionInfo) *ColorPacket
var GetImageMask func(i *Image, exception *ExceptionInfo) *Image
var GetImageMean func(i *Image, mean, standardDeviation *float64, exception *ExceptionInfo) bool

// Deprecated
var GetImagePixels func(i *Image, x, y SSize, columns, rows Size) *PixelPacket
var GetImageProfile func(i *Image, name string) *StringInfo
var GetImageProperty func(i *Image, property string) string
var GetImageQuantizeError func(i *Image) bool
var GetImageQuantumDepth func(i *Image, constrain bool) Size
var GetImageRange func(i *Image, minima, maxima *float64, exception *ExceptionInfo) bool
var GetImageTotalInkDensity func(i *Image) float64
var GetImageType func(i *Image, exception *ExceptionInfo) ImageType
var GetImageVirtualPixelMethod func(i *Image) VirtualPixelMethod

// Deprecated
var GetIndexes func(i *Image) *IndexPacket
var GetMagickPixelPacket func(i *Image, pixel *MagickPixelPacket)
var GetMultilineTypeMetrics func(i *Image, drawInfo *DrawInfo, metrics *TypeMetric) bool
var GetNextImageArtifact func(i *Image) string

// Deprecated
var GetNextImageAttribute func(i *Image) *ImageAttribute
var GetNextImageProfile func(i *Image) string
var GetNextImageProperty func(i *Image) string
var GetNumberColors func(i *Image, file *FILE, exception *ExceptionInfo) Size

// Deprecated
var GetNumberScenes func(i *Image) uint

// Deprecated
var GetOnePixel func(i *Image, x, y SSize) PixelPacket
var GetPixelCacheArea func(i *Image) MagickSizeType

// Deprecated
var GetPixels func(i *Image) *PixelPacket
var GetTypeMetrics func(i *Image, drawInfo *DrawInfo, metrics *TypeMetric) bool
var GradientImage func(i *Image, startColor, stopColor *PixelPacket) bool
var HuffmanDecodeImage func(i *Image) bool
var IdentifyImage func(i *Image, file *FILE, verbose bool) bool
var ImageToFile func(i *Image, filename string, exception *ExceptionInfo) bool
var ImplodeImage func(i *Image, amount float64, exception *ExceptionInfo) *Image
var ImportImagePixels func(i *Image, xOffset, yOffset SSize, columns, rows Size, map_ string, type_ StorageType, pixels *Void) bool
var ImportQuantumPixels func(i *Image, quantumInfo *QuantumInfo, quantumType QuantumType, pixels *byte) bool

// Deprecated
var InterpolatePixelColor func(i *Image, imageView *CacheView, method InterpolatePixelMethod, x, y float64, exception *ExceptionInfo) MagickPixelPacket
var IsBlobExempt func(i *Image) bool
var IsBlobSeekable func(i *Image) bool
var IsBlobTemporary func(i *Image) bool
var IsColorSimilar func(i *Image, p, q *PixelPacket) bool
var IsGrayImage func(i *Image, exception *ExceptionInfo) bool
var IsHighDynamicRangeImage func(i *Image, exception *ExceptionInfo) bool
var IsHistogramImage func(i *Image, exception *ExceptionInfo) bool
var IsImageObject func(i *Image) bool
var IsImagesEqual func(i *Image, reconstructImage *Image) bool
var IsImageSimilar func(i *Image, targetImage *Image, xOffset, yOffset *Long, exception *ExceptionInfo) bool
var IsMonochromeImage func(i *Image, exception *ExceptionInfo) bool
var IsOpacitySimilar func(i *Image, p, q *PixelPacket) bool
var IsOpaqueImage func(i *Image, exception *ExceptionInfo) bool
var IsPaletteImage func(i *Image, exception *ExceptionInfo) bool
var IsTaintImage func(i *Image) bool
var LevelImage func(i *Image, levels string) bool
var LevelImageChannel func(i *Image, channel ChannelType, blackPoint, whitePoint, gamma float64) bool
var LinearStretchImage func(i *Image, blackPoint, whitePoint float64) bool
var LiquidRescaleImage func(i *Image, columns, rows Size, deltaX, rigidity float64, exception *ExceptionInfo) *Image
var LZWEncodeImage func(i *Image, length uint32, pixels *byte) bool
var MagnifyImage func(i *Image, exception *ExceptionInfo) *Image

// Deprecated
var MapImage func(i *Image, mapImage *Image, dither bool) bool

// Deprecated
var MatteFloodfillImage func(i *Image, target PixelPacket, opacity Quantum, xOffset, yOffset SSize, method PaintMethod) bool

// Deprecated
var MedianFilterImage func(i *Image, radius float64, exception *ExceptionInfo) *Image
var MergeImageLayers func(i *Image, method ImageLayerMethod, exception *ExceptionInfo) *Image
var MinifyImage func(i *Image, exception *ExceptionInfo) *Image
var ModulateImage func(i *Image, modulate string) bool
var MorphImages func(i *Image, numberFrames Size, exception *ExceptionInfo) *Image

// Deprecated
var MosaicImages func(i *Image, exception *ExceptionInfo) *Image
var MotionBlurImage func(i *Image, radius, sigma, angle float64, exception *ExceptionInfo) *Image
var NegateImage func(i *Image, grayscale bool) bool
var NegateImageChannel func(i *Image, channel ChannelType, grayscale bool) bool
var NormalizeImage func(i *Image) bool
var NormalizeImageChannel func(i *Image, channel ChannelType) bool
var OilPaintImage func(i *Image, radius float64, exception *ExceptionInfo) *Image

// Deprecated
var OpaqueImage func(i *Image, target, fill PixelPacket) bool
var OpaquePaintImage func(i *Image, target, fill *MagickPixelPacket, invert bool) bool
var OpaquePaintImageChannel func(i *Image, channel ChannelType, target, fill *MagickPixelPacket, invert bool) bool

// Deprecated
var OpenCacheView func(i *Image) *CacheView
var OptimizeImageLayers func(i *Image, exception *ExceptionInfo) *Image
var OptimizeImageTransparency func(i *Image, exception *ExceptionInfo)
var OptimizePlusImageLayers func(i *Image, exception *ExceptionInfo) *Image
var OrderedDitherImage func(i *Image) bool
var OrderedDitherImageChannel func(i *Image, channel ChannelType, exception *ExceptionInfo) bool
var OrderedPosterizeImage func(i *Image, thresholdMap string, exception *ExceptionInfo) bool
var OrderedPosterizeImageChannel func(i *Image, channel ChannelType, thresholdMap string, exception *ExceptionInfo) bool
var PackbitsEncodeImage func(i *Image, length uint32, pixels *byte) bool

// Deprecated
var PaintFloodfillImage func(i *Image, channel ChannelType, target *MagickPixelPacket, x, y SSize, drawInfo *DrawInfo, method PaintMethod) bool

// Deprecated
var PaintOpaqueImage func(i *Image, target, fill *MagickPixelPacket) bool

// Deprecated
var PaintOpaqueImageChannel func(i *Image, channel ChannelType, target, fill *MagickPixelPacket) bool

// Deprecated
var PaintTransparentImage func(i *Image, target *MagickPixelPacket, opacity Quantum) bool
var ParseGravityGeometry func(i *Image, geometry string, regionInfo *RectangleInfo) MagickStatusType
var ParsePageGeometry func(i *Image, geometry string, regionInfo *RectangleInfo) MagickStatusType
var PersistCache func(i *Image, filename string, attach bool, offset *MagickOffsetType, exception *ExceptionInfo) bool
var PlasmaImage func(i *Image, segment *SegmentInfo, attenuate, depth Size) bool
var PolaroidImage func(i *Image, drawInfo *DrawInfo, angle float64, exception *ExceptionInfo) *Image

// Deprecated
var PopImagePixels func(i *Image, quantum QuantumType, destination *byte) bool
var PosterizeImage func(i *Image, levels Size, dither bool) bool
var PreviewImage func(i *Image, preview PreviewType, exception *ExceptionInfo) *Image
var ProfileImage func(i *Image, name string, datum *Void, length uint32, clone bool) bool

// Deprecated
var PushImagePixels func(i *Image, quantum QuantumType, source *byte) bool
var QuantizationError func(i *Image) uint
var QueryColorname func(i *Image, color *PixelPacket, compliance ComplianceType, name string, exception *ExceptionInfo) bool
var QueryMagickColorname func(i *Image, color *MagickPixelPacket, compliance ComplianceType, hex bool, name string, exception *ExceptionInfo) bool
var RadialBlurImage func(i *Image, angle float64, exception *ExceptionInfo) *Image
var RadialBlurImageChannel func(i *Image, channel ChannelType, angle float64, exception *ExceptionInfo) *Image
var RaiseImage func(i *Image, raiseInfo *RectangleInfo, raise bool) bool
var RandomChannelThresholdImage func(i *Image, channel, thresholds string, exception *ExceptionInfo) uint
var RandomThresholdImage func(i *Image, thresholds string, exception *ExceptionInfo) bool
var RandomThresholdImageChannel func(i *Image, channel ChannelType, thresholds string, exception *ExceptionInfo) bool
var ReadBlob func(i *Image, length uint32, data *byte) int32
var ReadBlobByte func(i *Image) int
var ReadBlobDouble func(i *Image) float64
var ReadBlobFloat func(i *Image) float32
var ReadBlobLong func(i *Image) Size
var ReadBlobLongLong func(i *Image) MagickSizeType
var ReadBlobLSBLong func(i *Image) Size
var ReadBlobLSBShort func(i *Image) UnsignedShort
var ReadBlobMSBLong func(i *Image) Size
var ReadBlobMSBShort func(i *Image) UnsignedShort
var ReadBlobShort func(i *Image) UnsignedShort
var ReadBlobString func(i *Image, str string) string

// Deprecated
var RecolorImage func(i *Image, order Size, colorMatrix *float64, exception *ExceptionInfo) *Image

// Deprecated
var ReduceNoiseImage func(i *Image, radius float64, exception *ExceptionInfo) *Image
var ReferenceImage func(i *Image) *Image
var RemoveImageArtifact func(i *Image, artifact string) string
var RemoveImageProfile func(i *Image, name string) *StringInfo
var RemoveImageProperty func(i *Image, property string) string
var ResampleImage func(i *Image, xResolution, yResolution float64, filter FilterTypes, blur float64, exception *ExceptionInfo) *Image
var ResetImageArtifactIterator func(i *Image)

// Deprecated
var ResetImageAttributeIterator func(i *Image)
var ResetImagePage func(i *Image, page string) bool
var ResetImageProfileIterator func(i *Image)
var ResetImagePropertyIterator func(i *Image)
var ResizeImage func(i *Image, columns, rows Size, filter FilterTypes, blur float64, exception *ExceptionInfo) *Image
var RGBTransformImage func(i *Image, colorspace ColorspaceType) bool
var RollImage func(i *Image, xOffset, yOffset SSize, exception *ExceptionInfo) *Image
var RotateImage func(i *Image, degrees float64, exception *ExceptionInfo) *Image
var SampleImage func(i *Image, columns, rows Size, exception *ExceptionInfo) *Image
var ScaleImage func(i *Image, columns, rows Size, exception *ExceptionInfo) *Image
var SeekBlob func(i *Image, offset MagickOffsetType, whence int) MagickOffsetType
var SegmentImage func(i *Image, colorspace ColorspaceType, verbose bool, clusterThreshold, smoothThreshold float64) bool
var SeparateImageChannel func(i *Image, channel ChannelType) bool
var SeparateImages func(i *Image, channel ChannelType, exception *ExceptionInfo) *Image
var SepiaToneImage func(i *Image, threshold float64, exception *ExceptionInfo) *Image
var SetBlobExempt func(i *Image, exempt bool)
var SetGeometry func(i *Image, geometry *RectangleInfo)

// Deprecated
var SetImage func(i *Image, opacity Quantum)
var SetImageAlphaChannel func(i *Image, alphaType AlphaChannelType) bool
var SetImageArtifact func(i *Image, artifact, value string) bool

// Deprecated
var SetImageAttribute func(i *Image, key, value string) bool
var SetImageBackgroundColor func(i *Image) bool
var SetImageChannelDepth func(i *Image, channel ChannelType, depth Size) bool
var SetImageClipMask func(i *Image, clipMask *Image) bool
var SetImageDepth func(i *Image, depth Size) bool
var SetImageExtent func(i *Image, columns, rows Size) bool
var SetImageMask func(i *Image, mask *Image) bool
var SetImageOpacity func(i *Image, opacity Quantum) bool

// Deprecated
var SetImagePixels func(i *Image, x, y SSize, columns, rows Size) *PixelPacket
var SetImageProfile func(i *Image, name string, profile *StringInfo) bool
var SetImageProgressMonitor func(i *Image, progressMonitor MagickProgressMonitor, clientData *Void) MagickProgressMonitor
var SetImageProperty func(i *Image, property, value string) bool
var SetImageStorageClass func(i *Image, storageClass ClassType) bool
var SetImageType func(i *Image, imageType ImageType) bool
var SetImageVirtualPixelMethod func(i *Image, virtualPixelMethod VirtualPixelMethod) VirtualPixelMethod
var ShadeImage func(i *Image, gray bool, azimuth, elevation float64, exception *ExceptionInfo) *Image
var ShadowImage func(i *Image, opacity, sigma float64, xOffset, yOffset SSize, exception *ExceptionInfo) *Image
var SharpenImage func(i *Image, radius, sigma float64, exception *ExceptionInfo) *Image
var SharpenImageChannel func(i *Image, channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image
var ShaveImage func(i *Image, shaveInfo *RectangleInfo, exception *ExceptionInfo) *Image
var ShearImage func(i *Image, xShear, yShear float64, exception *ExceptionInfo) *Image
var SigmoidalContrastImage func(i *Image, sharpen bool, levels string) bool
var SigmoidalContrastImageChannel func(i *Image, channel ChannelType, sharpen bool, contrast, midpoint float64) bool
var SignatureImage func(i *Image) bool
var SizeBlob func(i *Image) MagickOffsetType
var SketchImage func(i *Image, radius, sigma, angle float64, exception *ExceptionInfo) *Image
var SolarizeImage func(i *Image, threshold float64) bool
var SortColormapByIntensity func(i *Image) bool
var SpliceImage func(i *Image, geometry *RectangleInfo, exception *ExceptionInfo) *Image
var SpreadImage func(i *Image, radius float64, exception *ExceptionInfo) *Image
var SteganoImage func(i *Image, watermark *Image, exception *ExceptionInfo) *Image
var StereoImage func(i *Image, offsetImage *Image, exception *ExceptionInfo) *Image
var StripImage func(i *Image) bool
var SwirlImage func(i *Image, degrees float64, exception *ExceptionInfo) *Image
var SyncImage func(i *Image) bool

// Deprecated
var SyncImagePixels func(i *Image) bool
var SyncImageProfiles func(i *Image) bool
var TellBlob func(i *Image) MagickOffsetType
var TextureImage func(i *Image, texture *Image) bool

// Deprecated
var ThresholdImage func(i *Image, threshold float64) uint

// Deprecated
var ThresholdImageChannel func(i *Image, threshold string) uint
var ThumbnailImage func(i *Image, columns, rows Size, exception *ExceptionInfo) *Image
var TintImage func(i *Image, opacity string, tint PixelPacket, exception *ExceptionInfo) *Image
var TransformColorspace func(i *Image, colorspace ColorspaceType) uint
var TransformImageColorspace func(i *Image, colorspace ColorspaceType) bool
var TransformRGBImage func(i *Image, colorspace ColorspaceType) bool

// Deprecate
var TransparentImage func(i *Image, target PixelPacket, opacity Quantum) bool
var TransparentPaintImage func(i *Image, target *MagickPixelPacket, opacity Quantum, invert bool) bool
var TransposeImage func(i *Image, exception *ExceptionInfo) *Image
var TransverseImage func(i *Image, exception *ExceptionInfo) *Image
var TrimImage func(i *Image, exception *ExceptionInfo) *Image
var UniqueImageColors func(i *Image, exception *ExceptionInfo) *Image
var UnsharpMaskImage func(i *Image, radius, sigma, amount, threshold float64, exception *ExceptionInfo) *Image
var UnsharpMaskImageChannel func(i *Image, channel ChannelType, radius, sigma, amount, threshold float64, exception *ExceptionInfo) *Image
var ValidateColormapIndex func(i *Image, index Size) IndexPacket
var VignetteImage func(i *Image, radius, sigma float64, x, y SSize, exception *ExceptionInfo) *Image
var WaveImage func(i *Image, amplitude, waveLength float64, exception *ExceptionInfo) *Image
var WhiteThresholdImage func(i *Image, threshold string) bool
var WriteBlob func(i *Image, length uint32, data *byte) int32
var WriteBlobByte func(i *Image, value byte) int32
var WriteBlobFloat func(i *Image, value float32) int32
var WriteBlobLong func(i *Image, value UnsignedLong) int32
var WriteBlobLSBLong func(i *Image, value UnsignedLong) int32
var WriteBlobLSBShort func(i *Image, value UnsignedShort) int32
var WriteBlobMSBLong func(i *Image, value UnsignedLong) int32
var WriteBlobMSBShort func(i *Image, value UnsignedShort) int32
var WriteBlobShort func(i *Image, value UnsignedShort) int32
var WriteBlobString func(i *Image, str string) int32
var ZLIBEncodeImage func(i *Image, length uint32, pixels *byte) bool

// Deprecate
var ZoomImage func(i *Image, columns, rows Size, exception *ExceptionInfo) *Image
var AcquireAuthenticCacheView func(i *Image, exception *ExceptionInfo) *CacheView
var AcquireCacheView func(i *Image) *CacheView
var AcquireImageColormap func(i *Image, colors uint32) bool
var AcquirePixelCache func(*Image, VirtualPixelMethod, SSize, SSize, Size, Size, *ExceptionInfo) *PixelPacket

// Deprecated
var AcquirePixels func(i *Image) *PixelPacket // doc not ptr (i Image)
var AcquireVirtualCacheView func(i *Image, exception *ExceptionInfo) *CacheView
var AutoGammaImage func(i *Image) bool
var AutoGammaImageChannel func(i *Image, channel ChannelType) bool
var AutoLevelImage func(i *Image) bool
var AutoLevelImageChannel func(i *Image, channel ChannelType) bool
var BlueShiftImage func(i *Image, factor float64, exception *ExceptionInfo) *Image
var BrightnessContrastImage func(i *Image, brightness, contrast float64) bool
var BrightnessContrastImageChannel func(i *Image, channel ChannelType, brightness, contrast float64) bool
var ColorDecisionListImage func(i *Image, colorCorrectionCollection string) bool
var ColorMatrixImage func(i *Image, colorMatrix *KernelInfo, exception *ExceptionInfo) *Image
var CropImageToHBITMAP func(i *Image, r *RectangleInfo, e *ExceptionInfo) *Void
var CropImageToTiles func(i *Image, cropGeometry *RectangleInfo, exception *ExceptionInfo) *Image
var DecipherImage func(i *Image, passphrase string, exception *ExceptionInfo) bool
var DeskewImage func(i *Image, threshold float64, exception *ExceptionInfo) *Image
var EncipherImage func(i *Image, passphrase string, exception *ExceptionInfo) bool

// Deprecated
var ExtractSubimageFromImage func(i *Image, reference *Image, exception *ExceptionInfo) *Image
var FilterImage func(i *Image, kernel *KernelInfo, exception *ExceptionInfo) *Image
var FilterImageChannel func(i *Image, channel ChannelType, kernel *KernelInfo, exception *ExceptionInfo) *Image
var ForwardFourierTransformImage func(i *Image, modulus bool, exception *ExceptionInfo) *Image
var FunctionImage func(i *Image, function MagickFunction, numberParameters int32, parameters *float64, exception *ExceptionInfo) bool
var FunctionImageChannel func(i *Image, channel ChannelType, function MagickFunction, numberParameters int32, argument *float64, exception *ExceptionInfo) bool
var GetAuthenticIndexQueue func(i *Image) *IndexPacket
var GetAuthenticPixelQueue func(i *Image) *PixelPacket // doc not ptr (image Image)
var GetAuthenticPixels func(i *Image, x, y int32, columns, rows uint32, exception *ExceptionInfo) *PixelPacket
var GetBlobProperties func(i *Image) *Stat
var GetImageAlphaChannel func(i *Image) bool
var GetImageChannelDistortions func(i *Image, reconstructImage *Image, metric MetricType, exception *ExceptionInfo) *float64
var GetImageChannelFeatures func(i *Image, distance uint32, exception *ExceptionInfo) *ChannelFeatures
var GetImageChannelKurtosis func(i *Image, channel ChannelType, kurtosis, kewness *float64, exception *ExceptionInfo) bool
var GetImageChannels func(i *Image) uint32
var GetOneAuthenticPixel func(i *Image, x, y int32, pixel *PixelPacket, exception *ExceptionInfo) bool                                       // doc not ptr (image Image)
var GetOneVirtualMagickPixel func(i *Image, x, y int32, pixel *MagickPixelPacket, exception *ExceptionInfo) bool                             // doc not ptr (image Image...exception ExceptionInfo)
var GetOneVirtualMethodPixel func(i *Image, virtualPixelMethod VirtualPixelMethod, x, y int32, pixel *PixelPacket, exception *ExceptionInfo) // doc not ptr (image Image...pixel Pixelpacket...exception ExceptionInfo)
var GetOneVirtualPixel func(i *Image, x, y int32, pixel *PixelPacket, exception *ExceptionInfo) bool                                         // doc not ptr (image Image...exception ExceptionInfo)
var GetVirtualIndexQueue func(i *Image) *IndexPacket
var GetVirtualPixelQueue func(i *Image) *PixelPacket // doc not ptr (image Image)
var GetVirtualPixels func(i *Image, x, y int32, columns, rows uint32, exception *ExceptionInfo) *PixelPacket
var HaldClutImage func(i *Image, haldImage *Image) bool
var HaldClutImageChannel func(i *Image, channel ChannelType, haldImage *Image) bool
var ImageToHBITMAP func(i *Image) **Void
var IntegralRotateImage func(i *Image, rotations uint32, exception *ExceptionInfo) *Image
var InverseFourierTransformImage func(i *Image, phaseImage *Image, modulus bool, exception *ExceptionInfo) *Image
var LevelColorsImage func(i *Image, blackColor, whiteColor *MagickPixelPacket, invert bool) bool
var LevelColorsImageChannel func(i *Image, channel ChannelType, blackColor, whiteColor *MagickPixelPacket, invert bool) bool

// Deprecated
var LevelImageColors func(i *Image, channel ChannelType, blackColor, whiteColor *MagickPixelPacket, invert bool) bool
var LevelizeImage func(i *Image, blackPoint, whitePoint, gamma float64) bool
var LevelizeImageChannel func(i *Image, channel ChannelType, blackPoint, whitePoint, gamma float64) bool
var MinMaxStretchImage func(i *Image, channel ChannelType, blackAdjust, whiteAdjust float64) bool

// Deprecated
var ModeImage func(i *Image, radius float64, exception *ExceptionInfo) *Image
var MorphologyApply func(i *Image, method MorphologyMethod, channel ChannelType, iterations int32, kernel *KernelInfo, compose CompositeMethod, bias float64, exception *ExceptionInfo) *Image
var MorphologyImage func(i *Image, method MorphologyMethod, iterations int32, kernel *KernelInfo, exception *ExceptionInfo) *Image
var MorphologyImageChannel func(i *Image, channel ChannelType, method MorphologyMethod, iterations int32, kernel *KernelInfo, exception *ExceptionInfo) *Image
var MotionBlurImageChannel func(i *Image, channel ChannelType, radius, sigma, angle float64, exception *ExceptionInfo) *Image

// Deprecated
var ParseSizeGeometry func(i *Image, geometry string, regionInfo *RectangleInfo) MagickStatusType // doc not ptr (r RectangeInfo)
var PasskeyDecipherImage func(i *Image, passkey *StringInfo, exception *ExceptionInfo) bool
var PasskeyEncipherImage func(i *Image, passkey *StringInfo, exception *ExceptionInfo) bool
var PosterizeImageChannel func(i *Image, channel ChannelType, levels uint32, dither bool) bool
var QueueAuthenticPixels func(i *Image, x, y int32, columns, rows uint32, exception *ExceptionInfo) *PixelPacket
var SelectiveBlurImage func(i *Image, radius, sigma, threshold float64, exception *ExceptionInfo) *Image
var SelectiveBlurImageChannel func(i *Image, channel ChannelType, radius, sigma, threshold float64, exception *ExceptionInfo) *Image
var SetImageChannels func(i *Image, channels uint32) bool
var SetImageColor func(i *Image, color *MagickPixelPacket) bool
var SetImageColorspace func(i *Image, colorspace ColorspaceType) bool
var ShearRotateImage func(i *Image, degrees float64, exception *ExceptionInfo) *Image
var SimilarityImage func(i *Image, reference *Image, offset *RectangleInfo, similarity *float64, exception *ExceptionInfo) *Image
var SparseColorImage func(i *Image, channel ChannelType, method SparseColorMethod, numberArguments uint32, arguments *float64, exception *ExceptionInfo) *Image
var StatisticImage func(i *Image, type_ StatisticType, width, height uint32, exception *ExceptionInfo) *Image
var StatisticImageChannel func(i *Image, channel ChannelType, type_ StatisticType, width, height uint32, exception *ExceptionInfo) *Image
var StereoAnaglyphImage func(i *Image, rightImage *Image, xOffset, yOffset int32, exception *ExceptionInfo) *Image
var SyncAuthenticPixels func(i *Image, exception *ExceptionInfo) bool

func (i *Image) CatchException() ExceptionType    { return CatchImageException(i) }
func (i *Image) Channel(channel ChannelType) uint { return ChannelImage(i, channel) }

// Deprecated
func (i *Image) ChannelThreshold(level string) uint { return ChannelThresholdImage(i, level) }
func (i *Image) Charcoal(radius, sigma float64, exception *ExceptionInfo) *Image {
	return CharcoalImage(i, radius, sigma, exception)
}
func (i *Image) Chop(chopInfo *RectangleInfo, exception *ExceptionInfo) *Image {
	return ChopImage(i, chopInfo, exception)
}
func (i *Image) Clip() bool { return ClipImage(i) }
func (i *Image) ClipImagePath(pathname string, inside bool) bool {
	return ClipImagePath(i, pathname, inside)
}

// Deprecated
func (i *Image) ClipPathImage(pathname string, inside bool) bool {
	return ClipPathImage(i, pathname, inside)
}
func (i *Image) Clone(columns, rows Size, orphan bool, exception *ExceptionInfo) *Image {
	return CloneImage(i, columns, rows, orphan, exception)
}
func (i *Image) CloneArtifacts(cloneImage *Image) bool {
	return CloneImageArtifacts(i, cloneImage)
}

// Deprecated
func (i *Image) CloneAttributes(cloneImage *Image) bool {
	return CloneImageAttributes(i, cloneImage)
}
func (i *Image) CloneProfiles(cloneImage *Image) bool {
	return CloneImageProfiles(i, cloneImage)
}
func (i *Image) CloneImageProperties(cloneImage *Image) bool {
	return CloneImageProperties(i, cloneImage)
}
func (i *Image) CloseBlob() bool            { return CloseBlob(i) }
func (i *Image) Clut(clutImage *Image) bool { return ClutImage(i, clutImage) }
func (i *Image) ClutChannel(channel ChannelType, clutImage *Image) bool {
	return ClutImageChannel(i, channel, clutImage)
}
func (i *Image) CoalesceImages(exception *ExceptionInfo) *Image { return CoalesceImages(i, exception) }

// Deprecated
func (i *Image) ColorFloodfill(drawInfo *DrawInfo, target PixelPacket, xOffset, yOffset SSize, method PaintMethod) bool {
	return ColorFloodfillImage(i, drawInfo, target, xOffset, yOffset, method)
}
func (i *Image) Colorize(opacity string, colorize PixelPacket, exception *ExceptionInfo) *Image {
	return ColorizeImage(i, opacity, colorize, exception)
}
func (i *Image) CombineImages(channel ChannelType, exception *ExceptionInfo) *Image {
	return CombineImages(i, channel, exception)
}
func (i *Image) CompareChannels(reconstructImage *Image, channel ChannelType, metric MetricType, distortion *float64, exception *ExceptionInfo) *Image {
	return CompareImageChannels(i, reconstructImage, channel, metric, distortion, exception)
}
func (i *Image) CompareImageLayers(method ImageLayerMethod, exception *ExceptionInfo) *Image {
	return CompareImageLayers(i, method, exception)
}
func (i *Image) CompareImages(reconstructImage *Image, metric MetricType, distortion *float64, exception *ExceptionInfo) *Image {
	return CompareImages(i, reconstructImage, metric, distortion, exception)
}
func (i *Image) CompositeImage(compose CompositeOperator, compositeImage *Image, xOffset, yOffset SSize) bool {
	return CompositeImage(i, compose, compositeImage, xOffset, yOffset)
}
func (i *Image) CompositeChannel(channel ChannelType, compose CompositeOperator, compositeImage *Image, xOffset, yOffset SSize) bool {
	return CompositeImageChannel(i, channel, compose, compositeImage, xOffset, yOffset)
}
func (i *Image) CompositeLayers(compose CompositeOperator, source *Image, xOffset, yOffset SSize, exception *ExceptionInfo) {
	CompositeLayers(i, compose, source, xOffset, yOffset, exception)
}
func (i *Image) CompressColormap()                  { CompressImageColormap(i) }
func (i *Image) Contrast(sharpen bool) bool         { return ContrastImage(i, sharpen) }
func (i *Image) ContrastStretch(levels string) bool { return ContrastStretchImage(i, levels) }
func (i *Image) ContrastStretchChannel(channel ChannelType, blackPoint, whitePoint float64) bool {
	return ContrastStretchImageChannel(i, channel, blackPoint, whitePoint)
}
func (i *Image) Convolve(order Size, kernel *float64, exception *ExceptionInfo) *Image {
	return ConvolveImage(i, order, kernel, exception)
}
func (i *Image) ConvolveChannel(channel ChannelType, order Size, kernel *float64, exception *ExceptionInfo) *Image {
	return ConvolveImageChannel(i, channel, order, kernel, exception)
}
func (i *Image) CropImage(geometry *RectangleInfo, exception *ExceptionInfo) *Image {
	return CropImage(i, geometry, exception)
}
func (i *Image) CycleColormap(displace SSize) bool   { return CycleColormapImage(i, displace) }
func (i *Image) DefineArtifact(artifact string) bool { return DefineImageArtifact(i, artifact) }
func (i *Image) DefineProperty(property string) bool { return DefineImageProperty(i, property) }
func (i *Image) DeleteArtifact(artifact string) bool { return DeleteImageArtifact(i, artifact) }

// Deprecated
func (i *Image) DeleteAttribute(key string) bool     { return DeleteImageAttribute(i, key) }
func (i *Image) DeleteProfile(name string) bool      { return DeleteImageProfile(i, name) }
func (i *Image) DeleteProperty(property string) bool { return DeleteImageProperty(i, property) }

// Deprecated
func (i *Image) Describe(file *FILE, verbose bool) bool {
	return DescribeImage(i, file, verbose)
}
func (i *Image) Despeckle(exception *ExceptionInfo) *Image { return DespeckleImage(i, exception) }
func (i *Image) DestroyBlob()                              { DestroyBlob(i) }
func (i *Image) Destroy() *Image                           { return DestroyImage(i) }
func (i *Image) DestroyArtifacts()                         { DestroyImageArtifacts(i) }

// Deprecated
func (i *Image) DestroyAttributes() { DestroyImageAttributes(i) }
func (i *Image) DestroyPixels()     { DestroyImagePixels(i) }
func (i *Image) DestroyProfiles()   { DestroyImageProfiles(i) }
func (i *Image) DestroyProperties() { DestroyImageProperties(i) }

// Deprecated
func (i *Image) DestroyImages()      { DestroyImages(i) }
func (i *Image) DisassociateStream() { DisassociateImageStream(i) }

// Deprecated
func (i *Image) Dispatch(xOffset, yOffset SSize, columns, rows Size, map_ string, type_ StorageType, pixels *Void, exception *ExceptionInfo) uint {
	return DispatchImage(i, xOffset, yOffset, columns, rows, map_, type_, pixels, exception)
}
func (i *Image) Dispose(exception *ExceptionInfo) *Image { return DisposeImages(i, exception) }
func (i *Image) Distort(method DistortImageMethod, numberArguments Size, arguments *float64, bestfit bool, exception *ExceptionInfo) *Image {
	return DistortImage(i, method, numberArguments, arguments, bestfit, exception)
}
func (i *Image) DrawAffine(source *Image, affine *AffineMatrix) bool {
	return DrawAffineImage(i, source, affine)
}
func (i *Image) DrawClipPath(drawInfo *DrawInfo, name string) bool {
	return DrawClipPath(i, drawInfo, name)
}
func (i *Image) DrawGradient(drawInfo *DrawInfo) bool { return DrawGradientImage(i, drawInfo) }
func (i *Image) Draw(drawInfo *DrawInfo) bool         { return DrawImage(i, drawInfo) }
func (i *Image) DrawPatternPath(drawInfo *DrawInfo, name string, pattern **Image) bool {
	return DrawPatternPath(i, drawInfo, name, pattern)
}
func (i *Image) DrawPrimitive(drawInfo *DrawInfo, primitiveInfo *PrimitiveInfo) bool {
	return DrawPrimitive(i, drawInfo, primitiveInfo)
}
func (i *Image) Edge(radius float64, exception *ExceptionInfo) *Image {
	return EdgeImage(i, radius, exception)
}
func (i *Image) Emboss(radius, sigma float64, exception *ExceptionInfo) *Image {
	return EmbossImage(i, radius, sigma, exception)
}
func (i *Image) Enhance(exception *ExceptionInfo) *Image { return EnhanceImage(i, exception) }
func (i *Image) EOFBlob() int                            { return EOFBlob(i) }
func (i *Image) Equalize() bool                          { return EqualizeImage(i) }
func (i *Image) EqualizeChannel(channel ChannelType) bool {
	return EqualizeImageChannel(i, channel)
}
func (i *Image) Evaluate(op MagickEvaluateOperator, value float64, exception *ExceptionInfo) bool {
	return EvaluateImage(i, op, value, exception)
}
func (i *Image) EvaluateChannel(channel ChannelType, op MagickEvaluateOperator, value float64, exception *ExceptionInfo) bool {
	return EvaluateImageChannel(i, channel, op, value, exception)
}
func (i *Image) Excerpt(geometry *RectangleInfo, exception *ExceptionInfo) *Image {
	return ExcerptImage(i, geometry, exception)
}
func (i *Image) ExportPixels(xOffset, yOffset SSize, columns, rows Size, map_ string, type_ StorageType, pixels *Void, exception *ExceptionInfo) bool {
	return ExportImagePixels(i, xOffset, yOffset, columns, rows, map_, type_, pixels, exception)
}
func (i *Image) ExportQuantumPixels(quantumInfo *QuantumInfo, quantumType QuantumType, pixels *byte) bool {
	return ExportQuantumPixels(i, quantumInfo, quantumType, pixels)
}
func (i *Image) ExtentImage(geometry *RectangleInfo, exception *ExceptionInfo) *Image {
	return ExtentImage(i, geometry, exception)
}
func (i *Image) FromFile(filename string) bool { return FileToImage(i, filename) }

// Deprecated
func (i *Image) FlattenImages(exception *ExceptionInfo) *Image { return FlattenImages(i, exception) }
func (i *Image) Flip(exception *ExceptionInfo) *Image          { return FlipImage(i, exception) }
func (i *Image) FloodfillPaint(channel ChannelType, drawInfo *DrawInfo, target *MagickPixelPacket, xOffset, yOffset SSize, invert bool) bool {
	return FloodfillPaintImage(i, channel, drawInfo, target, xOffset, yOffset, invert)
}
func (i *Image) Flop(exception *ExceptionInfo) *Image { return FlopImage(i, exception) }

// Deprecated
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
func (i *Image) FormatMagickCaption(drawInfo *DrawInfo, caption string, metrics *TypeMetric) SSize {
	return FormatMagickCaption(i, drawInfo, caption, metrics)
}
func (i *Image) Frame(frameInfo *FrameInfo, exception *ExceptionInfo) *Image {
	return FrameImage(i, frameInfo, exception)
}
func (i *Image) FuzzyColorCompare(p, q *PixelPacket) bool   { return FuzzyColorCompare(i, p, q) }
func (i *Image) FuzzyOpacityCompare(p, q *PixelPacket) bool { return FuzzyOpacityCompare(i, p, q) }
func (i *Image) Fx(expression string, exception *ExceptionInfo) *Image {
	return FxImage(i, expression, exception)
}
func (i *Image) FxChannel(channel ChannelType, expression string, exception *ExceptionInfo) *Image {
	return FxImageChannel(i, channel, expression, exception)
}
func (i *Image) Gamma(level string) bool { return GammaImage(i, level) }
func (i *Image) GammaChannel(channel ChannelType, gamma float64) bool {
	return GammaImageChannel(i, channel, gamma)
}
func (i *Image) GaussianBlur(radius, sigma float64, exception *ExceptionInfo) *Image {
	return GaussianBlurImage(i, radius, sigma, exception)
}
func (i *Image) GaussianBlurChannel(channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return GaussianBlurImageChannel(i, channel, radius, sigma, exception)
}
func (i *Image) BlobError() bool                  { return GetBlobError(i) }
func (i *Image) BlobFileHandle() *FILE            { return GetBlobFileHandle(i) }
func (i *Image) BlobSize() MagickSizeType         { return GetBlobSize(i) }
func (i *Image) BlobStreamData() *byte            { return GetBlobStreamData(i) }
func (i *Image) BlobStreamHandler() StreamHandler { return GetBlobStreamHandler(i) }
func (i *Image) Artifact(artifact string) string  { return GetImageArtifact(i, artifact) }

// Deprecated
func (i *Image) Attribute(key string) *ImageAttribute { return GetImageAttribute(i, key) }
func (i *Image) BoundingBox(exception *ExceptionInfo) RectangleInfo {
	return GetImageBoundingBox(i, exception)
}
func (i *Image) ChannelDepth(channel ChannelType, exception *ExceptionInfo) Size {
	return GetImageChannelDepth(i, channel, exception)
}
func (i *Image) ChannelDistortion(reconstructImage *Image, channel ChannelType, metric MetricType, distortion *float64, exception *ExceptionInfo) bool {
	return GetImageChannelDistortion(i, reconstructImage, channel, metric, distortion, exception)
}
func (i *Image) ChannelExtrema(channel ChannelType, minima, maxima *Size, exception *ExceptionInfo) bool {
	return GetImageChannelExtrema(i, channel, minima, maxima, exception)
}
func (i *Image) ChannelMean(channel ChannelType, mean, standardDeviation *float64, exception *ExceptionInfo) bool {
	return GetImageChannelMean(i, channel, mean, standardDeviation, exception)
}
func (i *Image) ChannelRange(channel ChannelType, minima, maxima *float64, exception *ExceptionInfo) bool {
	return GetImageChannelRange(i, channel, minima, maxima, exception)
}
func (i *Image) ChannelStatistics(exception *ExceptionInfo) *ChannelStatistics {
	return GetImageChannelStatistics(i, exception)
}
func (i *Image) ClipMask(exception *ExceptionInfo) *Image {
	return GetImageClipMask(i, exception)
}

// Deprecated
func (i *Image) ClippingPathAttribute() *ImageAttribute {
	return GetImageClippingPathAttribute(i)
}
func (i *Image) Depth(exception *ExceptionInfo) Size {
	return GetImageDepth(i, exception)
}
func (i *Image) Distortion(reconstructImage *Image, metric MetricType, distortion *float64, exception *ExceptionInfo) bool {
	return GetImageDistortion(i, reconstructImage, metric, distortion, exception)
}
func (i *Image) DynamicThreshold(clusterThreshold, smoothThreshold float64, exception *ExceptionInfo) MagickPixelPacket {
	return GetImageDynamicThreshold(i, clusterThreshold, smoothThreshold, exception)
}
func (i *Image) Exception(exception *ExceptionInfo) { GetImageException(i, exception) }
func (i *Image) Extrema(minima, maxima *Size, exception *ExceptionInfo) bool {
	return GetImageExtrema(i, minima, maxima, exception)
}

// Deprecated
func (i *Image) Geometry(geometry string, sizeToFit uint, regionInfo *RectangleInfo) int {
	return GetImageGeometry(i, geometry, sizeToFit, regionInfo)
}
func (i *Image) Histogram(numberColors *Size, exception *ExceptionInfo) *ColorPacket {
	return GetImageHistogram(i, numberColors, exception)
}
func (i *Image) Mask(exception *ExceptionInfo) *Image { return GetImageMask(i, exception) }
func (i *Image) Mean(mean, standardDeviation *float64, exception *ExceptionInfo) bool {
	return GetImageMean(i, mean, standardDeviation, exception)
}

// Deprecated
func (i *Image) ImagePixels(x, y SSize, columns, rows Size) *PixelPacket {
	return GetImagePixels(i, x, y, columns, rows)
}
func (i *Image) Profile(name string) *StringInfo { return GetImageProfile(i, name) }
func (i *Image) Property(property string) string { return GetImageProperty(i, property) }
func (i *Image) QuantizeError() bool             { return GetImageQuantizeError(i) }
func (i *Image) QuantumDepth(constrain bool) Size {
	return GetImageQuantumDepth(i, constrain)
}
func (i *Image) Range(minima, maxima *float64, exception *ExceptionInfo) bool {
	return GetImageRange(i, minima, maxima, exception)
}
func (i *Image) TotalInkDensity() float64                { return GetImageTotalInkDensity(i) }
func (i *Image) Type(exception *ExceptionInfo) ImageType { return GetImageType(i, exception) }
func (i *Image) VirtualPixelMethod() VirtualPixelMethod  { return GetImageVirtualPixelMethod(i) }

// Deprecated
func (i *Image) Indexes() *IndexPacket                      { return GetIndexes(i) }
func (i *Image) MagickPixelPacket(pixel *MagickPixelPacket) { GetMagickPixelPacket(i, pixel) }
func (i *Image) MultilineTypeMetrics(drawInfo *DrawInfo, metrics *TypeMetric) bool {
	return GetMultilineTypeMetrics(i, drawInfo, metrics)
}
func (i *Image) NextArtifact() string { return GetNextImageArtifact(i) }

// Deprecated
func (i *Image) NextAttribute() *ImageAttribute { return GetNextImageAttribute(i) }
func (i *Image) NextProfile() string            { return GetNextImageProfile(i) }
func (i *Image) NextProperty() string           { return GetNextImageProperty(i) }
func (i *Image) NumberColors(file *FILE, exception *ExceptionInfo) Size {
	return GetNumberColors(i, file, exception)
}

// Deprecated
func (i *Image) NumberScenes() uint { return GetNumberScenes(i) }

// Deprecated
func (i *Image) Pixel(x, y SSize) PixelPacket   { return GetOnePixel(i, x, y) }
func (i *Image) PixelCacheArea() MagickSizeType { return GetPixelCacheArea(i) }

// Deprecated
func (i *Image) Pixels() *PixelPacket { return GetPixels(i) }
func (i *Image) TypeMetrics(drawInfo *DrawInfo, metrics *TypeMetric) bool {
	return GetTypeMetrics(i, drawInfo, metrics)
}
func (i *Image) Gradient(startColor, stopColor *PixelPacket) bool {
	return GradientImage(i, startColor, stopColor)
}
func (i *Image) HuffmanDecode() bool                    { return HuffmanDecodeImage(i) }
func (i *Image) Identify(file *FILE, verbose bool) bool { return IdentifyImage(i, file, verbose) }
func (i *Image) ToFile(filename string, exception *ExceptionInfo) bool {
	return ImageToFile(i, filename, exception)
}
func (i *Image) Implode(amount float64, exception *ExceptionInfo) *Image {
	return ImplodeImage(i, amount, exception)
}
func (i *Image) ImportPixels(xOffset, yOffset SSize, columns, rows Size, map_ string, type_ StorageType, pixels *Void) bool {
	return ImportImagePixels(i, xOffset, yOffset, columns, rows, map_, type_, pixels)
}
func (i *Image) ImportQuantumPixels(quantumInfo *QuantumInfo, quantumType QuantumType, pixels *byte) bool {
	return ImportQuantumPixels(i, quantumInfo, quantumType, pixels)
}

// Deprecated
func (i *Image) InterpolatePixelColor(imageView *CacheView, method InterpolatePixelMethod, x, y float64, exception *ExceptionInfo) MagickPixelPacket {
	return InterpolatePixelColor(i, imageView, method, x, y, exception)
}
func (i *Image) BlobExempt() bool                    { return IsBlobExempt(i) }
func (i *Image) BlobSeekable() bool                  { return IsBlobSeekable(i) }
func (i *Image) BlobTemporary() bool                 { return IsBlobTemporary(i) }
func (i *Image) ColorSimilar(p, q *PixelPacket) bool { return IsColorSimilar(i, p, q) }
func (i *Image) Gray(exception *ExceptionInfo) bool  { return IsGrayImage(i, exception) }
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
func (i *Image) Similar(targetImage *Image, xOffset, yOffset *Long, exception *ExceptionInfo) bool {
	return IsImageSimilar(i, targetImage, xOffset, yOffset, exception)
}
func (i *Image) Monochrome(exception *ExceptionInfo) bool {
	return IsMonochromeImage(i, exception)
}
func (i *Image) OpacitySimilar(p, q *PixelPacket) bool  { return IsOpacitySimilar(i, p, q) }
func (i *Image) IsOpaque(exception *ExceptionInfo) bool { return IsOpaqueImage(i, exception) }
func (i *Image) Palette(exception *ExceptionInfo) bool  { return IsPaletteImage(i, exception) }
func (i *Image) Taint() bool                            { return IsTaintImage(i) }
func (i *Image) Level(levels string) bool               { return LevelImage(i, levels) }
func (i *Image) LevelChannel(channel ChannelType, blackPoint, whitePoint, gamma float64) bool {
	return LevelImageChannel(i, channel, blackPoint, whitePoint, gamma)
}
func (i *Image) LinearStretch(blackPoint, whitePoint float64) bool {
	return LinearStretchImage(i, blackPoint, whitePoint)
}
func (i *Image) LiquidRescale(columns, rows Size, deltaX, rigidity float64, exception *ExceptionInfo) *Image {
	return LiquidRescaleImage(i, columns, rows, deltaX, rigidity, exception)
}
func (i *Image) LZWEncode(length uint32, pixels *byte) bool {
	return LZWEncodeImage(i, length, pixels)
}
func (i *Image) Magnify(exception *ExceptionInfo) *Image { return MagnifyImage(i, exception) }

// Deprecated
func (i *Image) Map(mapImage *Image, dither bool) bool { return MapImage(i, mapImage, dither) }
func (i *Image) MatteFloodfill(target PixelPacket, opacity Quantum, xOffset, yOffset SSize, method PaintMethod) bool {
	return MatteFloodfillImage(i, target, opacity, xOffset, yOffset, method)
}

// Deprecated
func (i *Image) MedianFilter(radius float64, exception *ExceptionInfo) *Image {
	return MedianFilterImage(i, radius, exception)
}
func (i *Image) MergeLayers(method ImageLayerMethod, exception *ExceptionInfo) *Image {
	return MergeImageLayers(i, method, exception)
}
func (i *Image) Minify(exception *ExceptionInfo) *Image { return MinifyImage(i, exception) }
func (i *Image) Modulate(modulate string) bool          { return ModulateImage(i, modulate) }
func (i *Image) Morph(numberFrames Size, exception *ExceptionInfo) *Image {
	return MorphImages(i, numberFrames, exception)
}
func (i *Image) Mosaic(exception *ExceptionInfo) *Image { return MosaicImages(i, exception) }
func (i *Image) MotionBlur(radius, sigma, angle float64, exception *ExceptionInfo) *Image {
	return MotionBlurImage(i, radius, sigma, angle, exception)
}
func (i *Image) Negate(grayscale bool) bool { return NegateImage(i, grayscale) }
func (i *Image) NegateChannel(channel ChannelType, grayscale bool) bool {
	return NegateImageChannel(i, channel, grayscale)
}
func (i *Image) Normalize() bool { return NormalizeImage(i) }
func (i *Image) NormalizeChannel(channel ChannelType) bool {
	return NormalizeImageChannel(i, channel)
}
func (i *Image) OilPaint(radius float64, exception *ExceptionInfo) *Image {
	return OilPaintImage(i, radius, exception)
}

// Deprecated
func (i *Image) Opaque(target, fill PixelPacket) bool { return OpaqueImage(i, target, fill) }
func (i *Image) OpaquePaint(target, fill *MagickPixelPacket, invert bool) bool {
	return OpaquePaintImage(i, target, fill, invert)
}
func (i *Image) OpaquePaintChannel(channel ChannelType, target, fill *MagickPixelPacket, invert bool) bool {
	return OpaquePaintImageChannel(i, channel, target, fill, invert)
}

// Deprecated
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
func (i *Image) OrderedDitherChannel(channel ChannelType, exception *ExceptionInfo) bool {
	return OrderedDitherImageChannel(i, channel, exception)
}
func (i *Image) OrderedPosterize(thresholdMap string, exception *ExceptionInfo) bool {
	return OrderedPosterizeImage(i, thresholdMap, exception)
}
func (i *Image) OrderedPosterizeChannel(channel ChannelType, thresholdMap string, exception *ExceptionInfo) bool {
	return OrderedPosterizeImageChannel(i, channel, thresholdMap, exception)
}
func (i *Image) PackbitsEncode(length uint32, pixels *byte) bool {
	return PackbitsEncodeImage(i, length, pixels)
}

// Deprecated
func (i *Image) PaintFloodfill(channel ChannelType, target *MagickPixelPacket, x, y SSize, drawInfo *DrawInfo, method PaintMethod) bool {
	return PaintFloodfillImage(i, channel, target, x, y, drawInfo, method)
}

// Deprecated
func (i *Image) PaintOpaque(target, fill *MagickPixelPacket) bool {
	return PaintOpaqueImage(i, target, fill)
}

// Deprecated
func (i *Image) PaintOpaqueChannel(channel ChannelType, target, fill *MagickPixelPacket) bool {
	return PaintOpaqueImageChannel(i, channel, target, fill)
}

// Deprecated
func (i *Image) PaintTransparent(target *MagickPixelPacket, opacity Quantum) bool {
	return PaintTransparentImage(i, target, opacity)
}
func (i *Image) ParseGravityGeometry(geometry string, regionInfo *RectangleInfo) MagickStatusType {
	return ParseGravityGeometry(i, geometry, regionInfo)
}
func (i *Image) ParsePageGeometry(geometry string, regionInfo *RectangleInfo) MagickStatusType {
	return ParsePageGeometry(i, geometry, regionInfo)
}
func (i *Image) PersistCache(filename string, attach bool, offset *MagickOffsetType, exception *ExceptionInfo) bool {
	return PersistCache(i, filename, attach, offset, exception)
}
func (i *Image) Plasma(segment *SegmentInfo, attenuate, depth Size) bool {
	return PlasmaImage(i, segment, attenuate, depth)
}
func (i *Image) Polaroid(drawInfo *DrawInfo, angle float64, exception *ExceptionInfo) *Image {
	return PolaroidImage(i, drawInfo, angle, exception)
}

// Deprecated
func (i *Image) PopPixels(quantum QuantumType, destination *byte) bool {
	return PopImagePixels(i, quantum, destination)
}
func (i *Image) Posterize(levels Size, dither bool) bool {
	return PosterizeImage(i, levels, dither)
}
func (i *Image) Preview(preview PreviewType, exception *ExceptionInfo) *Image {
	return PreviewImage(i, preview, exception)
}
func (i *Image) ProfileImage(name string, datum *Void, length uint32, clone bool) bool {
	return ProfileImage(i, name, datum, length, clone)
}

// Deprecated
func (i *Image) PushPixels(quantum QuantumType, source *byte) bool {
	return PushImagePixels(i, quantum, source)
}
func (i *Image) QuantizationError() uint { return QuantizationError(i) }
func (i *Image) QueryColorname(color *PixelPacket, compliance ComplianceType, name string, exception *ExceptionInfo) bool {
	return QueryColorname(i, color, compliance, name, exception)
}
func (i *Image) QueryMagickColorname(color *MagickPixelPacket, compliance ComplianceType, hex bool, name string, exception *ExceptionInfo) bool {
	return QueryMagickColorname(i, color, compliance, hex, name, exception)
}
func (i *Image) RadialBlur(angle float64, exception *ExceptionInfo) *Image {
	return RadialBlurImage(i, angle, exception)
}
func (i *Image) RadialBlurChannel(channel ChannelType, angle float64, exception *ExceptionInfo) *Image {
	return RadialBlurImageChannel(i, channel, angle, exception)
}
func (i *Image) RaiseImage(raiseInfo *RectangleInfo, raise bool) bool {
	return RaiseImage(i, raiseInfo, raise)
}
func (i *Image) RandomChannelThresholdImage(channel, thresholds string, exception *ExceptionInfo) uint {
	return RandomChannelThresholdImage(i, channel, thresholds, exception)
}
func (i *Image) RandomThresholdImage(thresholds string, exception *ExceptionInfo) bool {
	return RandomThresholdImage(i, thresholds, exception)
}
func (i *Image) RandomThresholdImageChannel(channel ChannelType, thresholds string, exception *ExceptionInfo) bool {
	return RandomThresholdImageChannel(i, channel, thresholds, exception)
}
func (i *Image) ReadBlob(length uint32, data *byte) int32 { return ReadBlob(i, length, data) }
func (i *Image) ReadBlobByte() int                        { return ReadBlobByte(i) }
func (i *Image) ReadBlobDouble() float64                  { return ReadBlobDouble(i) }
func (i *Image) ReadBlobFloat() float32                   { return ReadBlobFloat(i) }
func (i *Image) ReadBlobLong() Size                       { return ReadBlobLong(i) }
func (i *Image) ReadBlobLongLong() MagickSizeType         { return ReadBlobLongLong(i) }
func (i *Image) ReadBlobLSBLong() Size                    { return ReadBlobLSBLong(i) }
func (i *Image) ReadBlobLSBShort() UnsignedShort          { return ReadBlobLSBShort(i) }
func (i *Image) ReadBlobMSBLong() Size                    { return ReadBlobMSBLong(i) }
func (i *Image) ReadBlobMSBShort() UnsignedShort          { return ReadBlobMSBShort(i) }
func (i *Image) ReadBlobShort() UnsignedShort             { return ReadBlobShort(i) }
func (i *Image) ReadBlobString(str string) string         { return ReadBlobString(i, str) }

// Deprecated
func (i *Image) Recolor(order Size, colorMatrix *float64, exception *ExceptionInfo) *Image {
	return RecolorImage(i, order, colorMatrix, exception)
}

// Deprecated
func (i *Image) ReduceNoise(radius float64, exception *ExceptionInfo) *Image {
	return ReduceNoiseImage(i, radius, exception)
}
func (i *Image) Reference() *Image                     { return ReferenceImage(i) }
func (i *Image) RemoveArtifact(artifact string) string { return RemoveImageArtifact(i, artifact) }
func (i *Image) RemoveProfile(name string) *StringInfo { return RemoveImageProfile(i, name) }
func (i *Image) RemoveProperty(property string) string { return RemoveImageProperty(i, property) }
func (i *Image) Resample(xResolution, yResolution float64, filter FilterTypes, blur float64, exception *ExceptionInfo) *Image {
	return ResampleImage(i, xResolution, yResolution, filter, blur, exception)
}
func (i *Image) ResetArtifactIterator() { ResetImageArtifactIterator(i) }

// Deprecated
func (i *Image) ResetAttributeIterator()    { ResetImageAttributeIterator(i) }
func (i *Image) ResetPage(page string) bool { return ResetImagePage(i, page) }
func (i *Image) ResetProfileIterator()      { ResetImageProfileIterator(i) }
func (i *Image) ResetPropertyIterator()     { ResetImagePropertyIterator(i) }
func (i *Image) Resize(columns, rows Size, filter FilterTypes, blur float64, exception *ExceptionInfo) *Image {
	return ResizeImage(i, columns, rows, filter, blur, exception)
}
func (i *Image) RGBTransform(colorspace ColorspaceType) bool {
	return RGBTransformImage(i, colorspace)
}
func (i *Image) Roll(xOffset, yOffset SSize, exception *ExceptionInfo) *Image {
	return RollImage(i, xOffset, yOffset, exception)
}
func (i *Image) Rotate(degrees float64, exception *ExceptionInfo) *Image {
	return RotateImage(i, degrees, exception)
}
func (i *Image) Sample(columns, rows Size, exception *ExceptionInfo) *Image {
	return SampleImage(i, columns, rows, exception)
}
func (i *Image) Scale(columns, rows Size, exception *ExceptionInfo) *Image {
	return ScaleImage(i, columns, rows, exception)
}
func (i *Image) SeekBlob(offset MagickOffsetType, whence int) MagickOffsetType {
	return SeekBlob(i, offset, whence)
}
func (i *Image) Segment(colorspace ColorspaceType, verbose bool, clusterThreshold, smoothThreshold float64) bool {
	return SegmentImage(i, colorspace, verbose, clusterThreshold, smoothThreshold)
}
func (i *Image) SeparateChannel(channel ChannelType) bool {
	return SeparateImageChannel(i, channel)
}
func (i *Image) SeparateImages(channel ChannelType, exception *ExceptionInfo) *Image {
	return SeparateImages(i, channel, exception)
}
func (i *Image) SepiaTone(threshold float64, exception *ExceptionInfo) *Image {
	return SepiaToneImage(i, threshold, exception)
}
func (i *Image) SetBlobExempt(exempt bool)           { SetBlobExempt(i, exempt) }
func (i *Image) SetGeometry(geometry *RectangleInfo) { SetGeometry(i, geometry) }

// Deprecated
func (i *Image) SetImage(opacity Quantum) { SetImage(i, opacity) }
func (i *Image) SetAlphaChannel(alphaType AlphaChannelType) bool {
	return SetImageAlphaChannel(i, alphaType)
}
func (i *Image) SetArtifact(artifact, value string) bool {
	return SetImageArtifact(i, artifact, value)
}

// Deprecated
func (i *Image) SetAttribute(key, value string) bool { return SetImageAttribute(i, key, value) }
func (i *Image) SetBackgroundColor() bool            { return SetImageBackgroundColor(i) }
func (i *Image) SetChannelDepth(channel ChannelType, depth Size) bool {
	return SetImageChannelDepth(i, channel, depth)
}
func (i *Image) SetClipMask(clipMask *Image) bool { return SetImageClipMask(i, clipMask) }
func (i *Image) SetDepth(depth Size) bool         { return SetImageDepth(i, depth) }
func (i *Image) SetExtent(columns, rows Size) bool {
	return SetImageExtent(i, columns, rows)
}
func (i *Image) SetMask(mask *Image) bool        { return SetImageMask(i, mask) }
func (i *Image) SetOpacity(opacity Quantum) bool { return SetImageOpacity(i, opacity) }

// Deprecated
func (i *Image) SetPixels(x, y SSize, columns, rows Size) *PixelPacket {
	return SetImagePixels(i, x, y, columns, rows)
}
func (i *Image) SetProfile(name string, profile *StringInfo) bool {
	return SetImageProfile(i, name, profile)
}
func (i *Image) SetProgressMonitor(progressMonitor MagickProgressMonitor, clientData *Void) MagickProgressMonitor {
	return SetImageProgressMonitor(i, progressMonitor, clientData)
}
func (i *Image) SetProperty(property, value string) bool {
	return SetImageProperty(i, property, value)
}
func (i *Image) SetStorageClass(storageClass ClassType) bool {
	return SetImageStorageClass(i, storageClass)
}
func (i *Image) SetType(imageType ImageType) bool { return SetImageType(i, imageType) }
func (i *Image) SetVirtualPixelMethod(virtualPixelMethod VirtualPixelMethod) VirtualPixelMethod {
	return SetImageVirtualPixelMethod(i, virtualPixelMethod)
}
func (i *Image) Shade(gray bool, azimuth, elevation float64, exception *ExceptionInfo) *Image {
	return ShadeImage(i, gray, azimuth, elevation, exception)
}
func (i *Image) Shadow(opacity, sigma float64, xOffset, yOffset SSize, exception *ExceptionInfo) *Image {
	return ShadowImage(i, opacity, sigma, xOffset, yOffset, exception)
}
func (i *Image) Sharpen(radius, sigma float64, exception *ExceptionInfo) *Image {
	return SharpenImage(i, radius, sigma, exception)
}
func (i *Image) SharpenChannel(channel ChannelType, radius, sigma float64, exception *ExceptionInfo) *Image {
	return SharpenImageChannel(i, channel, radius, sigma, exception)
}
func (i *Image) Shave(shaveInfo *RectangleInfo, exception *ExceptionInfo) *Image {
	return ShaveImage(i, shaveInfo, exception)
}
func (i *Image) Shear(xShear, yShear float64, exception *ExceptionInfo) *Image {
	return ShearImage(i, xShear, yShear, exception)
}
func (i *Image) SigmoidalContrast(sharpen bool, levels string) bool {
	return SigmoidalContrastImage(i, sharpen, levels)
}
func (i *Image) SigmoidalContrastChannel(channel ChannelType, sharpen bool, contrast, midpoint float64) bool {
	return SigmoidalContrastImageChannel(i, channel, sharpen, contrast, midpoint)
}
func (i *Image) Signature() bool            { return SignatureImage(i) }
func (i *Image) SizeBlob() MagickOffsetType { return SizeBlob(i) }
func (i *Image) Sketch(radius, sigma, angle float64, exception *ExceptionInfo) *Image {
	return SketchImage(i, radius, sigma, angle, exception)
}
func (i *Image) Solarize(threshold float64) bool { return SolarizeImage(i, threshold) }
func (i *Image) SortColormapByIntensity() bool   { return SortColormapByIntensity(i) }
func (i *Image) Splice(geometry *RectangleInfo, exception *ExceptionInfo) *Image {
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

// Deprecated
func (i *Image) SyncPixels() bool            { return SyncImagePixels(i) }
func (i *Image) SyncProfiles() bool          { return SyncImageProfiles(i) }
func (i *Image) TellBlob() MagickOffsetType  { return TellBlob(i) }
func (i *Image) Texture(texture *Image) bool { return TextureImage(i, texture) }

// Deprecated
func (i *Image) Threshold(threshold float64) uint { return ThresholdImage(i, threshold) }

// Deprecated
func (i *Image) ThresholdChannel(threshold string) uint {
	return ThresholdImageChannel(i, threshold)
}
func (i *Image) Thumbnail(columns, rows Size, exception *ExceptionInfo) *Image {
	return ThumbnailImage(i, columns, rows, exception)
}
func (i *Image) Tint(opacity string, tint PixelPacket, exception *ExceptionInfo) *Image {
	return TintImage(i, opacity, tint, exception)
}
func (i *Image) TransformColorspace(colorspace ColorspaceType) uint {
	return TransformColorspace(i, colorspace)
}
func (i *Image) TransformImageColorspace(colorspace ColorspaceType) bool {
	return TransformImageColorspace(i, colorspace)
}
func (i *Image) TransformRGB(colorspace ColorspaceType) bool {
	return TransformRGBImage(i, colorspace)
}

// Deprecate
func (i *Image) Transparent(target PixelPacket, opacity Quantum) bool {
	return TransparentImage(i, target, opacity)
}
func (i *Image) TransparentPaint(target *MagickPixelPacket, opacity Quantum, invert bool) bool {
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
func (i *Image) UnsharpMaskChannel(channel ChannelType, radius, sigma, amount, threshold float64, exception *ExceptionInfo) *Image {
	return UnsharpMaskImageChannel(i, channel, radius, sigma, amount, threshold, exception)
}
func (i *Image) ValidateColormapIndex(index Size) IndexPacket {
	return ValidateColormapIndex(i, index)
}
func (i *Image) Vignette(radius, sigma float64, x, y SSize, exception *ExceptionInfo) *Image {
	return VignetteImage(i, radius, sigma, x, y, exception)
}
func (i *Image) Wave(amplitude, waveLength float64, exception *ExceptionInfo) *Image {
	return WaveImage(i, amplitude, waveLength, exception)
}
func (i *Image) WhiteThreshold(threshold string) bool        { return WhiteThresholdImage(i, threshold) }
func (i *Image) WriteBlob(length uint32, data *byte) int32   { return WriteBlob(i, length, data) }
func (i *Image) WriteBlobByte(value byte) int32              { return WriteBlobByte(i, value) }
func (i *Image) WriteBlobFloat(value float32) int32          { return WriteBlobFloat(i, value) }
func (i *Image) WriteBlobLong(value UnsignedLong) int32      { return WriteBlobLong(i, value) }
func (i *Image) WriteBlobLSBLong(value UnsignedLong) int32   { return WriteBlobLSBLong(i, value) }
func (i *Image) WriteBlobLSBShort(value UnsignedShort) int32 { return WriteBlobLSBShort(i, value) }
func (i *Image) WriteBlobMSBLong(value UnsignedLong) int32   { return WriteBlobMSBLong(i, value) }
func (i *Image) WriteBlobMSBShort(value UnsignedShort) int32 { return WriteBlobMSBShort(i, value) }
func (i *Image) WriteBlobShort(value UnsignedShort) int32    { return WriteBlobShort(i, value) }
func (i *Image) WriteBlobString(str string) int32            { return WriteBlobString(i, str) }
func (i *Image) ZLIBEncode(length uint32, pixels *byte) bool {
	return ZLIBEncodeImage(i, length, pixels)
}

// Deprecate
func (i *Image) Zoom(columns, rows Size, exception *ExceptionInfo) *Image {
	return ZoomImage(i, columns, rows, exception)
}
func (i *Image) AcquireAuthenticCacheView(exception *ExceptionInfo) *CacheView {
	return AcquireAuthenticCacheView(i, exception)
}
func (i *Image) AcquireCacheView() *CacheView            { return AcquireCacheView(i) }
func (i *Image) AcquireImageColormap(colors uint32) bool { return AcquireImageColormap(i, colors) }
func (i *Image) AcquirePixelCache(a1 VirtualPixelMethod, a2 SSize, a3 SSize, a4 Size, a5 Size, a6 *ExceptionInfo) *PixelPacket {
	return AcquirePixelCache(i, a1, a2, a3, a4, a5, a6)
}

// Deprecate
func (i *Image) AcquirePixels() *PixelPacket { return AcquirePixels(i) }
func (i *Image) AcquireVirtualCacheView(exception *ExceptionInfo) *CacheView {
	return AcquireVirtualCacheView(i, exception)
}
func (i *Image) AutoGamma() bool { return AutoGammaImage(i) }
func (i *Image) AutoGammaChannel(channel ChannelType) bool {
	return AutoGammaImageChannel(i, channel)
}
func (i *Image) AutoLevel() bool { return AutoLevelImage(i) }
func (i *Image) AutoLevelChannel(channel ChannelType) bool {
	return AutoLevelImageChannel(i, channel)
}
func (i *Image) BlueShift(factor float64, exception *ExceptionInfo) *Image {
	return BlueShiftImage(i, factor, exception)
}
func (i *Image) BrightnessContrast(brightness, contrast float64) bool {
	return BrightnessContrastImage(i, brightness, contrast)
}
func (i *Image) BrightnessContrastChannel(channel ChannelType, brightness, contrast float64) bool {
	return BrightnessContrastImageChannel(i, channel, brightness, contrast)
}
func (i *Image) ColorDecisionList(colorCorrectionCollection string) bool {
	return ColorDecisionListImage(i, colorCorrectionCollection)
}
func (i *Image) ColorMatrix(colorMatrix *KernelInfo, exception *ExceptionInfo) *Image {
	return ColorMatrixImage(i, colorMatrix, exception)
}
func (i *Image) CropToHBITMAP(r *RectangleInfo, e *ExceptionInfo) *Void {
	return CropImageToHBITMAP(i, r, e)
}
func (i *Image) CropToTiles(cropGeometry *RectangleInfo, exception *ExceptionInfo) *Image {
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

// Deprecated
func (i *Image) ExtractSubimage(reference *Image, exception *ExceptionInfo) *Image {
	return ExtractSubimageFromImage(i, reference, exception)
}
func (i *Image) Filter(kernel *KernelInfo, exception *ExceptionInfo) *Image {
	return FilterImage(i, kernel, exception)
}
func (i *Image) FilterChannel(channel ChannelType, kernel *KernelInfo, exception *ExceptionInfo) *Image {
	return FilterImageChannel(i, channel, kernel, exception)
}
func (i *Image) FFT(modulus bool, exception *ExceptionInfo) *Image {
	return ForwardFourierTransformImage(i, modulus, exception)
}
func (i *Image) Function(function MagickFunction, numberParameters int32, parameters *float64, exception *ExceptionInfo) bool {
	return FunctionImage(i, function, numberParameters, parameters, exception)
}
func (i *Image) FunctionChannel(channel ChannelType, function MagickFunction, numberParameters int32, argument *float64, exception *ExceptionInfo) bool {
	return FunctionImageChannel(i, channel, function, numberParameters, argument, exception)
}
func (i *Image) AuthenticIndexQueue() *IndexPacket { return GetAuthenticIndexQueue(i) }
func (i *Image) AuthenticPixelQueue() *PixelPacket { return GetAuthenticPixelQueue(i) }
func (i *Image) AuthenticPixels(x, y int32, columns, rows uint32, exception *ExceptionInfo) *PixelPacket {
	return GetAuthenticPixels(i, x, y, columns, rows, exception)
}
func (i *Image) BlobProperties() *Stat { return GetBlobProperties(i) }
func (i *Image) AlphaChannel() bool    { return GetImageAlphaChannel(i) }
func (i *Image) ChannelDistortions(reconstructImage *Image, metric MetricType, exception *ExceptionInfo) *float64 {
	return GetImageChannelDistortions(i, reconstructImage, metric, exception)
}
func (i *Image) ChannelFeatures(distance uint32, exception *ExceptionInfo) *ChannelFeatures {
	return GetImageChannelFeatures(i, distance, exception)
}
func (i *Image) ChannelKurtosis(channel ChannelType, kurtosis, kewness *float64, exception *ExceptionInfo) bool {
	return GetImageChannelKurtosis(i, channel, kurtosis, kewness, exception)
}
func (i *Image) Channels() uint32 { return GetImageChannels(i) }
func (i *Image) AuthenticPixel(x, y int32, pixel *PixelPacket, exception *ExceptionInfo) bool {
	return GetOneAuthenticPixel(i, x, y, pixel, exception)
}
func (i *Image) VirtualMagickPixel(x, y int32, pixel *MagickPixelPacket, exception *ExceptionInfo) bool {
	return GetOneVirtualMagickPixel(i, x, y, pixel, exception)
}
func (i *Image) VirtualMethodPixel(virtualPixelMethod VirtualPixelMethod, x, y int32, pixel *PixelPacket, exception *ExceptionInfo) {
	GetOneVirtualMethodPixel(i, virtualPixelMethod, x, y, pixel, exception)
}
func (i *Image) VirtualPixel(x, y int32, pixel *PixelPacket, exception *ExceptionInfo) bool {
	return GetOneVirtualPixel(i, x, y, pixel, exception)
}
func (i *Image) VirtualIndexQueue() *IndexPacket { return GetVirtualIndexQueue(i) }
func (i *Image) VirtualPixelQueue() *PixelPacket { return GetVirtualPixelQueue(i) }
func (i *Image) VirtualPixels(x, y int32, columns, rows uint32, exception *ExceptionInfo) *PixelPacket {
	return GetVirtualPixels(i, x, y, columns, rows, exception)
}
func (i *Image) HaldClut(haldImage *Image) bool { return HaldClutImage(i, haldImage) }
func (i *Image) HaldClutChannel(channel ChannelType, haldImage *Image) bool {
	return HaldClutImageChannel(i, channel, haldImage)
}
func (i *Image) ToHBITMAP() **Void { return ImageToHBITMAP(i) }
func (i *Image) IntegralRotate(rotations uint32, exception *ExceptionInfo) *Image {
	return IntegralRotateImage(i, rotations, exception)
}
func (i *Image) InverseFourierTransform(phaseImage *Image, modulus bool, exception *ExceptionInfo) *Image {
	return InverseFourierTransformImage(i, phaseImage, modulus, exception)
}
func (i *Image) LevelColors(blackColor, whiteColor *MagickPixelPacket, invert bool) bool {
	return LevelColorsImage(i, blackColor, whiteColor, invert)
}
func (i *Image) LevelColorsChannel(channel ChannelType, blackColor, whiteColor *MagickPixelPacket, invert bool) bool {
	return LevelColorsImageChannel(i, channel, blackColor, whiteColor, invert)
}

// Deprecated
func (i *Image) LevelImageColors(channel ChannelType, blackColor, whiteColor *MagickPixelPacket, invert bool) bool {
	// Name conflict w/LevelColorsImage if shortened
	return LevelImageColors(i, channel, blackColor, whiteColor, invert)
}
func (i *Image) Levelize(blackPoint, whitePoint, gamma float64) bool {
	return LevelizeImage(i, blackPoint, whitePoint, gamma)
}
func (i *Image) LevelizeChannel(channel ChannelType, blackPoint, whitePoint, gamma float64) bool {
	return LevelizeImageChannel(i, channel, blackPoint, whitePoint, gamma)
}
func (i *Image) MinMaxStretch(channel ChannelType, blackAdjust, whiteAdjust float64) bool {
	return MinMaxStretchImage(i, channel, blackAdjust, whiteAdjust)
}

// Deprecated
func (i *Image) Mode(radius float64, exception *ExceptionInfo) *Image {
	return ModeImage(i, radius, exception)
}
func (i *Image) MorphologyApply(method MorphologyMethod, channel ChannelType, iterations int32, kernel *KernelInfo, compose CompositeMethod, bias float64, exception *ExceptionInfo) *Image {
	return MorphologyApply(i, method, channel, iterations, kernel, compose, bias, exception)
}
func (i *Image) Morphology(method MorphologyMethod, iterations int32, kernel *KernelInfo, exception *ExceptionInfo) *Image {
	return MorphologyImage(i, method, iterations, kernel, exception)
}
func (i *Image) MorphologyChannel(channel ChannelType, method MorphologyMethod, iterations int32, kernel *KernelInfo, exception *ExceptionInfo) *Image {
	return MorphologyImageChannel(i, channel, method, iterations, kernel, exception)
}
func (i *Image) MotionBlurChannel(channel ChannelType, radius, sigma, angle float64, exception *ExceptionInfo) *Image {
	return MotionBlurImageChannel(i, channel, radius, sigma, angle, exception)
}

// Deprecated
func (i *Image) ParseSizeGeometry(geometry string, regionInfo *RectangleInfo) MagickStatusType { // doc not ptr (r RectangeInfo)
	return ParseSizeGeometry(i, geometry, regionInfo)
}
func (i *Image) PasskeyDecipher(passkey *StringInfo, exception *ExceptionInfo) bool {
	return PasskeyDecipherImage(i, passkey, exception)
}
func (i *Image) PasskeyEncipher(passkey *StringInfo, exception *ExceptionInfo) bool {
	return PasskeyEncipherImage(i, passkey, exception)
}
func (i *Image) PosterizeChannel(channel ChannelType, levels uint32, dither bool) bool {
	return PosterizeImageChannel(i, channel, levels, dither)
}
func (i *Image) QueueAuthenticPixels(x, y int32, columns, rows uint32, exception *ExceptionInfo) *PixelPacket {
	return QueueAuthenticPixels(i, x, y, columns, rows, exception)
}
func (i *Image) SelectiveBlur(radius, sigma, threshold float64, exception *ExceptionInfo) *Image {
	return SelectiveBlurImage(i, radius, sigma, threshold, exception)
}
func (i *Image) SelectiveBlurChannel(channel ChannelType, radius, sigma, threshold float64, exception *ExceptionInfo) *Image {
	return SelectiveBlurImageChannel(i, channel, radius, sigma, threshold, exception)
}
func (i *Image) SetChannels(channels uint32) bool       { return SetImageChannels(i, channels) }
func (i *Image) SetColor(color *MagickPixelPacket) bool { return SetImageColor(i, color) }
func (i *Image) SetColorspace(colorspace ColorspaceType) {
	SetImageColorspace(i, colorspace)
}
func (i *Image) ShearRotate(degrees float64, exception *ExceptionInfo) *Image {
	return ShearRotateImage(i, degrees, exception)
}
func (i *Image) Similarity(reference *Image, offset *RectangleInfo, similarity *float64, exception *ExceptionInfo) *Image {
	return SimilarityImage(i, reference, offset, similarity, exception)
}
func (i *Image) SparseColor(channel ChannelType, method SparseColorMethod, numberArguments uint32, arguments *float64, exception *ExceptionInfo) *Image {
	return SparseColorImage(i, channel, method, numberArguments, arguments, exception)
}
func (i *Image) Statistic(type_ StatisticType, width, height uint32, exception *ExceptionInfo) *Image {
	return StatisticImage(i, type_, width, height, exception)
}
func (i *Image) StatisticChannel(channel ChannelType, type_ StatisticType, width, height uint32, exception *ExceptionInfo) *Image {
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

// Deprecated
var DeleteImageList func(images *Image, offset SSize) uint
var DestroyImageList func(images *Image) *Image
var GetFirstImageInList func(images *Image) *Image
var GetImageFromList func(images *Image, index SSize) *Image
var GetImageIndexInList func(images *Image) SSize

// Deprecated
var GetImageList func(images *Image, offset SSize, exception *ExceptionInfo) *Image

// Deprecated
var GetImageListIndex func(images *Image) SSize
var GetImageListLength func(images *Image) Size

// Deprecated
var GetImageListSize func(images *Image) Size
var GetLastImageInList func(images *Image) *Image

// Deprecated
var GetNextImage func(images *Image) *Image
var GetNextImageInList func(images *Image) *Image

// Deprecated
var GetPreviousImage func(images *Image) *Image
var GetPreviousImageInList func(images *Image) *Image
var ImageListToArray func(images *Image, exception *ExceptionInfo) []*Image
var MontageImages func(images *Image, montageInfo *MontageInfo, exception *ExceptionInfo) *Image

// Deprecated
var SpliceImageList func(images *Image, offset SSize, length Size, splices *Image, exception *ExceptionInfo) *Image
var SplitImageList func(images *Image) *Image
var SyncImageList func(images *Image)
var SyncNextImageInList func(images *Image) *Image
var AppendImageToList func(images **Image, image *Image)
var DeleteImageFromList func(images **Image)
var DeleteImages func(images **Image, scenes string, exception *ExceptionInfo)
var InsertImageInList func(images **Image, image *Image)

// Deprecated
var PopImageList func(images **Image) *Image
var PrependImageToList func(images **Image, image *Image)

// Deprecated
var PushImageList func(images **Image, image *Image, exception *ExceptionInfo) uint
var RemoveDuplicateLayers func(images **Image, exception *ExceptionInfo)
var RemoveFirstImageFromList func(images **Image) *Image
var RemoveImageFromList func(images **Image) *Image
var RemoveLastImageFromList func(images **Image) *Image
var RemoveZeroDelayLayers func(images **Image, exception *ExceptionInfo)
var ReplaceImageInList func(images **Image, image *Image)
var ReverseImageList func(images **Image)

// Deprecated
var SetImageList func(images **Image, image *Image, offset SSize, exception *ExceptionInfo) uint

// Deprecated
var ShiftImageList func(images **Image) *Image
var SpliceImageIntoList func(images **Image, length Size, splice *Image) *Image
var TransformImages func(images **Image, cropGeometry, imageGeometry string) bool

// Deprecate
var UnshiftImageList func(images **Image, image *Image, exception *ExceptionInfo) uint

// Deprecated
var AcquireCacheViewIndexes func(c *CacheView) *IndexPacket

// Deprecated
var AcquireCacheViewPixels func(c *CacheView, x, y SSize, columns, rows Size, exception *ExceptionInfo) *PixelPacket

// Deprecated
var AcquireOneCacheViewPixel func(c *CacheView, x, y SSize, exception *ExceptionInfo) PixelPacket

// Deprecated
var CloneCacheView func(c *CacheView) *CacheView
var CloseCacheView func(c *CacheView) *CacheView

// Deprecated
var GetCacheView func(c *CacheView, x, y SSize, columns, rows Size) *PixelPacket
var GetCacheViewColorspace func(c *CacheView) ColorspaceType
var GetCacheViewException func(c *CacheView) *ExceptionInfo

// Deprecated
var GetCacheViewIndexes func(c *CacheView) *IndexPacket

// Deprecated
var GetCacheViewPixels func(c *CacheView, x, y SSize, columns, rows Size) *PixelPacket
var GetCacheViewStorageClass func(c *CacheView) ClassType
var SetCacheViewStorageClass func(c *CacheView, storageClass ClassType) bool
var SetCacheViewVirtualPixelMethod func(c *CacheView, virtualPixelMethod VirtualPixelMethod) bool

//sketchy docs and/or deprecated
var DestroyCacheView func(c *CacheView) *CacheView
var GetCacheViewAuthenticIndexQueue func(c *CacheView) *IndexPacket
var GetCacheViewAuthenticPixelQueue func(c *CacheView) *PixelPacket
var GetCacheViewAuthenticPixels func(c *CacheView, x, y int32, columns, rows uint32, exception *ExceptionInfo) *PixelPacket
var GetCacheViewChannels func(c *CacheView) uint32
var GetCacheViewVirtualIndexQueue func(c *CacheView) *IndexPacket
var GetCacheViewVirtualPixelQueue func(c *CacheView) *PixelPacket
var GetCacheViewVirtualPixels func(c *CacheView, x, y int32, columns, rows uint32, exception *ExceptionInfo) *PixelPacket
var GetOneCacheViewAuthenticPixel func(c *CacheView, x, y int32, pixel *PixelPacket, exception *ExceptionInfo) bool // doc not ptr (p Pixelpacket)
var GetOneCacheViewVirtualMethodPixel func(c *CacheView, virtualPixelMethod VirtualPixelMethod, x, y int32, pixel *PixelPacket, exception *ExceptionInfo) bool
var GetOneCacheViewVirtualPixel func(c *CacheView, x, y int32, pixel *PixelPacket, exception *ExceptionInfo) bool
var QueueCacheViewAuthenticPixels func(c *CacheView, x, y int32, columns, rows uint32, exception *ExceptionInfo) *PixelPacket

// Deprecated
var SetCacheViewPixels func(c *CacheView, x, y int32, columns, rows uint32) *PixelPacket

// Deprecated
var SyncCacheView func(c *CacheView) bool
var SyncCacheViewAuthenticPixels func(c *CacheView, exception *ExceptionInfo) bool

// Deprecated
var SyncCacheViewPixels func(c *CacheView) bool

// Deprecated
func (c *CacheView) AcquireIndexes() *IndexPacket { return AcquireCacheViewIndexes(c) }

// Deprecated
func (c *CacheView) AcquirePixels(x, y SSize, columns, rows Size, exception *ExceptionInfo) *PixelPacket {
	return AcquireCacheViewPixels(c, x, y, columns, rows, exception)
}

// Deprecated
func (c *CacheView) AcquirePixel(x, y SSize, exception *ExceptionInfo) PixelPacket {
	return AcquireOneCacheViewPixel(c, x, y, exception)
}

// Deprecated
func (c *CacheView) Clone() *CacheView { return CloneCacheView(c) }
func (c *CacheView) Close() *CacheView { return CloseCacheView(c) }

// Deprecated
func (c *CacheView) Get(x, y SSize, columns, rows Size) *PixelPacket {
	return GetCacheView(c, x, y, columns, rows)
}
func (c *CacheView) Colorspace() ColorspaceType { return GetCacheViewColorspace(c) }
func (c *CacheView) Exception() *ExceptionInfo  { return GetCacheViewException(c) }

// Deprecated
func (c *CacheView) Indexes() *IndexPacket { return GetCacheViewIndexes(c) }

// Deprecated
func (c *CacheView) Pixels(x, y SSize, columns, rows Size) *PixelPacket {
	return GetCacheViewPixels(c, x, y, columns, rows)
}
func (c *CacheView) StorageClass() ClassType { return GetCacheViewStorageClass(c) }
func (c *CacheView) SetStorageClass(storageClass ClassType) bool {
	return SetCacheViewStorageClass(c, storageClass)
}
func (c *CacheView) SetVirtualPixelMethod(virtualPixelMethod VirtualPixelMethod) bool {
	return SetCacheViewVirtualPixelMethod(c, virtualPixelMethod)
}
func (c *CacheView) Destroy() *CacheView               { return DestroyCacheView(c) }
func (c *CacheView) AuthenticIndexQueue() *IndexPacket { return GetCacheViewAuthenticIndexQueue(c) }
func (c *CacheView) AuthenticPixelQueue() *PixelPacket { return GetCacheViewAuthenticPixelQueue(c) }
func (c *CacheView) AuthenticPixels(x, y int32, columns, rows uint32, exception *ExceptionInfo) *PixelPacket {
	return GetCacheViewAuthenticPixels(c, x, y, columns, rows, exception)
}
func (c *CacheView) Channels() uint32                { return GetCacheViewChannels(c) }
func (c *CacheView) VirtualIndexQueue() *IndexPacket { return GetCacheViewVirtualIndexQueue(c) }
func (c *CacheView) VirtualPixelQueue() *PixelPacket { return GetCacheViewVirtualPixelQueue(c) }
func (c *CacheView) VirtualPixels(x, y int32, columns, rows uint32, exception *ExceptionInfo) *PixelPacket {
	return GetCacheViewVirtualPixels(c, x, y, columns, rows, exception)
}
func (c *CacheView) AuthenticPixel(x, y int32, pixel *PixelPacket, exception *ExceptionInfo) bool {
	return GetOneCacheViewAuthenticPixel(c, x, y, pixel, exception)
}
func (c *CacheView) VirtualMethodPixel(virtualPixelMethod VirtualPixelMethod, x, y int32, pixel *PixelPacket, exception *ExceptionInfo) bool {
	return GetOneCacheViewVirtualMethodPixel(c, virtualPixelMethod, x, y, pixel, exception)
}
func (c *CacheView) VirtualPixel(x, y int32, pixel *PixelPacket, exception *ExceptionInfo) bool {
	return GetOneCacheViewVirtualPixel(c, x, y, pixel, exception)
}
func (c *CacheView) QueueAuthenticPixels(x, y int32, columns, rows uint32, exception *ExceptionInfo) *PixelPacket {
	return QueueCacheViewAuthenticPixels(c, x, y, columns, rows, exception)
}

// Deprecated
func (c *CacheView) SetPixels(x, y int32, columns, rows uint32) *PixelPacket {
	return SetCacheViewPixels(c, x, y, columns, rows)
}

// Deprecated
func (c *CacheView) Sync() bool { return SyncCacheView(c) }
func (c *CacheView) SyncAuthenticPixels(exception *ExceptionInfo) bool {
	return SyncCacheViewAuthenticPixels(c, exception)
}

// Deprecated
func (c *CacheView) SyncPixels() bool { return SyncCacheViewPixels(c) }

var GetImageDecoder func(m *MagickInfo) *DecodeImageHandler
var GetImageEncoder func(m *MagickInfo) *EncodeImageHandler
var GetMagickAdjoin func(m *MagickInfo) bool
var GetMagickBlobSupport func(m *MagickInfo) bool
var GetMagickDescription func(m *MagickInfo) string
var GetMagickEndianSupport func(m *MagickInfo) bool
var GetMagickSeekableStream func(m *MagickInfo) bool
var GetMagickThreadSupport func(m *MagickInfo) MagickStatusType
var RegisterMagickInfo func(m *MagickInfo) *MagickInfo

func (m *MagickInfo) Decoder() *DecodeImageHandler    { return GetImageDecoder(m) }
func (m *MagickInfo) Encoder() *EncodeImageHandler    { return GetImageEncoder(m) }
func (m *MagickInfo) Adjoin() bool                    { return GetMagickAdjoin(m) }
func (m *MagickInfo) BlobSupport() bool               { return GetMagickBlobSupport(m) }
func (m *MagickInfo) Description() string             { return GetMagickDescription(m) }
func (m *MagickInfo) EndianSupport() bool             { return GetMagickEndianSupport(m) }
func (m *MagickInfo) SeekableStream() bool            { return GetMagickSeekableStream(m) }
func (m *MagickInfo) ThreadSupport() MagickStatusType { return GetMagickThreadSupport(m) }
func (m *MagickInfo) Register() *MagickInfo           { return RegisterMagickInfo(m) }

var XAnimateBackgroundImage func(d *Display, resourceInfo *XResourceInfo, images *Image)
var XAnimateImages func(d *Display, resourceInfo *XResourceInfo, argv []string, argc int, images *Image) *Image
var XAnnotateImage func(d *Display, pixel *XPixelInfo, annotateInfo *XAnnotateInfo, image *Image) bool
var XBestFont func(d *Display, resourceInfo *XResourceInfo, textFont bool) *XFontStruct
var XBestIconSize func(d *Display, window *XWindowInfo, image *Image)
var XBestPixel func(d *Display, colormap Colormap, colors *XColor, numberColors uint, color *XColor)
var XBestVisualInfo func(d *Display, mapInfo *XStandardColormap, resourceInfo *XResourceInfo) *XVisualInfo
var XCheckDefineCursor func(d *Display, window Window, cursor Cursor) int
var XCheckRefreshWindows func(d *Display, windows *XWindows)
var XClientMessage func(d *Display, window Window, protocol, reason Atom, timestamp Time)
var XColorBrowserWidget func(d *Display, windows *XWindows, action, reply string)
var XCommandWidget func(d *Display, windows *XWindows, selections []string, event *XEvent) int
var XConfigureImageColormap func(d *Display, resourceInfo *XResourceInfo, windows *XWindows, image *Image)
var XConfirmWidget func(d *Display, windows *XWindows, reason, description string) int
var XConstrainWindowPosition func(d *Display, windowInfo *XWindowInfo)
var XDelay func(d *Display, milliseconds Size)
var XDestroyWindowColors func(d *Display, window Window)
var XDialogWidget func(d *Display, windows *XWindows, action, query, reply string) int
var XDisplayBackgroundImage func(d *Display, resourceInfo *XResourceInfo, image *Image) bool
var XDisplayImage func(d *Display, resourceInfo *XResourceInfo, argv []string, argc int, image **Image, state *Size) *Image
var XDisplayImageInfo func(d *Display, resourceInfo *XResourceInfo, windows *XWindows, undoImage, image *Image)
var XDrawImage func(d *Display, pixel *XPixelInfo, drawInfo *XDrawInfo, image *Image) bool
var XError func(d *Display, err *XErrorEvent) int
var XFileBrowserWidget func(d *Display, windows *XWindows, action, reply string)
var XFontBrowserWidget func(d *Display, windows *XWindows, action, reply string)
var XFreeResources func(d *Display, visualInfo *XVisualInfo, mapInfo *XStandardColormap, pixel *XPixelInfo, fontInfo *XFontStruct, resourceInfo *XResourceInfo, windowInfo *XWindowInfo)
var XFreeStandardColormap func(d *Display, visualInfo *XVisualInfo, mapInfo *XStandardColormap, pixel *XPixelInfo)
var XGetPixelPacket func(d *Display, visualInfo *XVisualInfo, mapInfo *XStandardColormap, resourceInfo *XResourceInfo, image *Image, pixel *XPixelInfo)
var XGetResourceDatabase func(d *Display, clientName string) *XrmDatabase
var XGetScreenDensity func(d *Display) string
var XGetWindowColor func(d *Display, windows *XWindows, name string) bool
var XGetWindowInfo func(d *Display, visualInfo *XVisualInfo, mapInfo *XStandardColormap, pixel *XPixelInfo, fontInfo *XFontStruct, resourceInfo *XResourceInfo, window *XWindowInfo)
var XHighlightEllipse func(d *Display, window Window, annotateContext GC, highlightInfo *RectangleInfo)
var XHighlightLine func(d *Display, window Window, annotateContext GC, highlightInfo *XSegment)
var XHighlightRectangle func(d *Display, window Window, annotateContext GC, highlightInfo *RectangleInfo)
var XInfoWidget func(d *Display, windows *XWindows, activity string)
var XInitializeWindows func(d *Display, resourceInfo *XResourceInfo) *XWindows
var XListBrowserWidget func(d *Display, windows *XWindows, windowInfo *XWindowInfo, list []string, action, query, reply string)
var XMakeCursor func(d *Display, window Window, colormap Colormap, backgroundColor, foregroundColor string) Cursor
var XMakeImage func(d *Display, resourceInfo *XResourceInfo, window *XWindowInfo, image *Image, width, height uint) bool
var XMakeMagnifyImage func(d *Display, windows *XWindows)
var XMakeStandardColormap func(d *Display, visualInfo *XVisualInfo, resourceInfo *XResourceInfo, image *Image, mapInfo *XStandardColormap, pixel *XPixelInfo)
var XMakeWindow func(d *Display, parent Window, argv []string, argc int, classHint *XClassHint, managerHints *XWMHints, windowInfo *XWindowInfo)
var XMenuWidget func(d *Display, windows *XWindows, title string, selections []string, item string) int
var XNoticeWidget func(d *Display, windows *XWindows, reason, description string)
var XPreferencesWidget func(d *Display, resourceInfo *XResourceInfo, windows *XWindows) bool
var XProgressMonitorWidget func(d *Display, windows *XWindows, task string, offset MagickOffsetType, span MagickSizeType)
var XQueryPosition func(d *Display, window Window, x, y *int)
var XRefreshWindow func(d *Display, window *XWindowInfo, event *XEvent)
var XRemoteCommand func(d *Display, window, filename string) bool
var XRetainWindowColors func(d *Display, window Window)
var XrmGetDatabase func(d *Display) *XrmDatabase
var XSetCursorState func(d *Display, windows *XWindows, state MagickStatusType)
var XTextViewWidget func(d *Display, resourceInfo *XResourceInfo, windows *XWindows, mono bool, title string, textlist []string)
var XWindowByID func(d *Display, rootWindow Window, id Size) Window
var XWindowByName func(d *Display, rootWindow Window, name string) Window
var XWindowByProperty func(d *Display, window Window, property Atom) Window

func (d *Display) XAnimateBackgroundImage(resourceInfo *XResourceInfo, images *Image) {
	XAnimateBackgroundImage(d, resourceInfo, images)
}
func (d *Display) XAnimateImages(resourceInfo *XResourceInfo, argv []string, argc int, images *Image) *Image {
	return XAnimateImages(d, resourceInfo, argv, argc, images)
}
func (d *Display) XAnnotateImage(pixel *XPixelInfo, annotateInfo *XAnnotateInfo, image *Image) bool {
	return XAnnotateImage(d, pixel, annotateInfo, image)
}
func (d *Display) XBestFont(resourceInfo *XResourceInfo, textFont bool) *XFontStruct {
	return XBestFont(d, resourceInfo, textFont)
}
func (d *Display) XBestIconSize(window *XWindowInfo, image *Image) {
	XBestIconSize(d, window, image)
}
func (d *Display) XBestPixel(colormap Colormap, colors *XColor, numberColors uint, color *XColor) {
	XBestPixel(d, colormap, colors, numberColors, color)
}
func (d *Display) XBestVisualInfo(mapInfo *XStandardColormap, resourceInfo *XResourceInfo) *XVisualInfo {
	return XBestVisualInfo(d, mapInfo, resourceInfo)
}
func (d *Display) XCheckDefineCursor(window Window, cursor Cursor) int {
	return XCheckDefineCursor(d, window, cursor)
}
func (d *Display) XCheckRefreshWindows(windows *XWindows) { XCheckRefreshWindows(d, windows) }
func (d *Display) XClientMessage(window Window, protocol, reason Atom, timestamp Time) {
	XClientMessage(d, window, protocol, reason, timestamp)
}
func (d *Display) XColorBrowserWidget(windows *XWindows, action, reply string) {
	XColorBrowserWidget(d, windows, action, reply)
}
func (d *Display) XCommandWidget(windows *XWindows, selections []string, event *XEvent) int {
	return XCommandWidget(d, windows, selections, event)
}
func (d *Display) XConfigureImageColormap(resourceInfo *XResourceInfo, windows *XWindows, image *Image) {
	XConfigureImageColormap(d, resourceInfo, windows, image)
}
func (d *Display) XConfirmWidget(windows *XWindows, reason, description string) int {
	return XConfirmWidget(d, windows, reason, description)
}
func (d *Display) XConstrainWindowPosition(windowInfo *XWindowInfo) {
	XConstrainWindowPosition(d, windowInfo)
}
func (d *Display) XDelay(milliseconds Size)           { XDelay(d, milliseconds) }
func (d *Display) XDestroyWindowColors(window Window) { XDestroyWindowColors(d, window) }
func (d *Display) XDialogWidget(windows *XWindows, action, query, reply string) int {
	return XDialogWidget(d, windows, action, query, reply)
}
func (d *Display) XBackgroundImage(resourceInfo *XResourceInfo, image *Image) bool {
	return XDisplayBackgroundImage(d, resourceInfo, image)
}
func (d *Display) XDisplayImage(resourceInfo *XResourceInfo, argv []string, argc int, image **Image, state *Size) *Image {
	return XDisplayImage(d, resourceInfo, argv, argc, image, state)
}
func (d *Display) XDisplayImageInfo(resourceInfo *XResourceInfo, windows *XWindows, undoImage, image *Image) {
	XDisplayImageInfo(d, resourceInfo, windows, undoImage, image)
}
func (d *Display) XDrawImage(pixel *XPixelInfo, drawInfo *XDrawInfo, image *Image) bool {
	return XDrawImage(d, pixel, drawInfo, image)
}
func (d *Display) XError(err *XErrorEvent) int { return XError(d, err) }
func (d *Display) XFileBrowserWidget(windows *XWindows, action, reply string) {
	XFileBrowserWidget(d, windows, action, reply)
}
func (d *Display) XFontBrowserWidget(windows *XWindows, action, reply string) {
	XFontBrowserWidget(d, windows, action, reply)
}
func (d *Display) XFreeResources(visualInfo *XVisualInfo, mapInfo *XStandardColormap, pixel *XPixelInfo, fontInfo *XFontStruct, resourceInfo *XResourceInfo, windowInfo *XWindowInfo) {
	XFreeResources(d, visualInfo, mapInfo, pixel, fontInfo, resourceInfo, windowInfo)
}
func (d *Display) XFreeStandardColormap(visualInfo *XVisualInfo, mapInfo *XStandardColormap, pixel *XPixelInfo) {
	XFreeStandardColormap(d, visualInfo, mapInfo, pixel)
}
func (d *Display) XGetPixelPacket(visualInfo *XVisualInfo, mapInfo *XStandardColormap, resourceInfo *XResourceInfo, image *Image, pixel *XPixelInfo) {
	XGetPixelPacket(d, visualInfo, mapInfo, resourceInfo, image, pixel)
}
func (d *Display) XGetResourceDatabase(clientName string) *XrmDatabase {
	return XGetResourceDatabase(d, clientName)
}
func (d *Display) XGetScreenDensity() string { return XGetScreenDensity(d) }
func (d *Display) XGetWindowColor(windows *XWindows, name string) bool {
	return XGetWindowColor(d, windows, name)
}
func (d *Display) XGetWindowInfo(visualInfo *XVisualInfo, mapInfo *XStandardColormap, pixel *XPixelInfo, fontInfo *XFontStruct, resourceInfo *XResourceInfo, window *XWindowInfo) {
	XGetWindowInfo(d, visualInfo, mapInfo, pixel, fontInfo, resourceInfo, window)
}
func (d *Display) XHighlightEllipse(window Window, annotateContext GC, highlightInfo *RectangleInfo) {
	XHighlightEllipse(d, window, annotateContext, highlightInfo)
}
func (d *Display) XHighlightLine(window Window, annotateContext GC, highlightInfo *XSegment) {
	XHighlightLine(d, window, annotateContext, highlightInfo)
}
func (d *Display) XHighlightRectangle(window Window, annotateContext GC, highlightInfo *RectangleInfo) {
	XHighlightRectangle(d, window, annotateContext, highlightInfo)
}
func (d *Display) XInfoWidget(windows *XWindows, activity string) { XInfoWidget(d, windows, activity) }
func (d *Display) XInitializeWindows(resourceInfo *XResourceInfo) *XWindows {
	return XInitializeWindows(d, resourceInfo)
}
func (d *Display) XListBrowserWidget(windows *XWindows, windowInfo *XWindowInfo, list []string, action, query, reply string) {
	XListBrowserWidget(d, windows, windowInfo, list, action, query, reply)
}
func (d *Display) XMakeCursor(window Window, colormap Colormap, backgroundColor, foregroundColor string) Cursor {
	return XMakeCursor(d, window, colormap, backgroundColor, foregroundColor)
}
func (d *Display) XMakeImage(resourceInfo *XResourceInfo, window *XWindowInfo, image *Image, width, height uint) bool {
	return XMakeImage(d, resourceInfo, window, image, width, height)
}
func (d *Display) XMakeMagnifyImage(windows *XWindows) { XMakeMagnifyImage(d, windows) }
func (d *Display) XMakeStandardColormap(visualInfo *XVisualInfo, resourceInfo *XResourceInfo, image *Image, mapInfo *XStandardColormap, pixel *XPixelInfo) {
	XMakeStandardColormap(d, visualInfo, resourceInfo, image, mapInfo, pixel)
}
func (d *Display) XMakeWindow(parent Window, argv []string, argc int, classHint *XClassHint, managerHints *XWMHints, windowInfo *XWindowInfo) {
	XMakeWindow(d, parent, argv, argc, classHint, managerHints, windowInfo)
}
func (d *Display) XMenuWidget(windows *XWindows, title string, selections []string, item string) int {
	return XMenuWidget(d, windows, title, selections, item)
}
func (d *Display) XNoticeWidget(windows *XWindows, reason, description string) {
	XNoticeWidget(d, windows, reason, description)
}
func (d *Display) XPreferencesWidget(resourceInfo *XResourceInfo, windows *XWindows) bool {
	return XPreferencesWidget(d, resourceInfo, windows)
}
func (d *Display) XProgressMonitorWidget(windows *XWindows, task string, offset MagickOffsetType, span MagickSizeType) {
	XProgressMonitorWidget(d, windows, task, offset, span)
}
func (d *Display) XQueryPosition(window Window, x, y *int) { XQueryPosition(d, window, x, y) }
func (d *Display) XRefreshWindow(window *XWindowInfo, event *XEvent) {
	XRefreshWindow(d, window, event)
}
func (d *Display) XRemoteCommand(window, filename string) bool {
	return XRemoteCommand(d, window, filename)
}
func (d *Display) XRetainWindowColors(window Window) { XRetainWindowColors(d, window) }
func (d *Display) XrmGetDatabase() *XrmDatabase      { return XrmGetDatabase(d) }
func (d *Display) XSetCursorState(windows *XWindows, state MagickStatusType) {
	XSetCursorState(d, windows, state)
}
func (d *Display) XTextViewWidget(resourceInfo *XResourceInfo, windows *XWindows, mono bool, title string, textlist []string) {
	XTextViewWidget(d, resourceInfo, windows, mono, title, textlist)
}
func (d *Display) XWindowByID(rootWindow Window, id Size) Window {
	return XWindowByID(d, rootWindow, id)
}
func (d *Display) XWindowByName(rootWindow Window, name string) Window {
	return XWindowByName(d, rootWindow, name)
}
func (d *Display) XWindowByProperty(window Window, property Atom) Window {
	return XWindowByProperty(d, window, property)
}

var AcquireImageInfo func() *ImageInfo
var AcquireQuantizeInfo func(i *ImageInfo) *QuantizeInfo
var AcquireQuantumInfo func(i *ImageInfo) *QuantumInfo
var AcquireStreamInfo func(i *ImageInfo) *StreamInfo

// Deprecated
var AllocateImage func(i *ImageInfo) *Image

// Deprecated
var AllocateNextImage func(i *ImageInfo, image *Image)
var AnimateImages func(i *ImageInfo, images *Image) bool
var BlobToImage func(i *ImageInfo, blob *Void, length uint32, exception *ExceptionInfo) *Image
var CloneDrawInfo func(i *ImageInfo, drawInfo *DrawInfo) *DrawInfo
var CloneImageInfo func(i *ImageInfo) *ImageInfo
var CloneImageOptions func(i *ImageInfo, cloneInfo *ImageInfo) bool
var CloneMontageInfo func(i *ImageInfo, montageInfo *MontageInfo) *MontageInfo
var DefineImageOption func(i *ImageInfo, option string) bool
var DeleteImageOption func(i *ImageInfo, option string) bool
var DestroyImageInfo func(i *ImageInfo) *ImageInfo
var DestroyImageOptions func(i *ImageInfo)
var DisplayImages func(i *ImageInfo, images *Image) bool
var GetDelegateCommand func(i *ImageInfo, image *Image, decode, encode string, exception *ExceptionInfo) string
var GetDrawInfo func(i *ImageInfo, drawInfo *DrawInfo)
var GetImageInfo func(i *ImageInfo)
var GetImageOption func(i *ImageInfo, key string) string
var GetMontageInfo func(i *ImageInfo, montageInfo *MontageInfo)
var GetNextImageOption func(i *ImageInfo) string
var GetQuantumInfo func(i *ImageInfo, quantumInfo *QuantumInfo)
var HuffmanEncodeImage func(i *ImageInfo, image *Image) bool
var ImagesToBlob func(i *ImageInfo, images *Image, length *uint32, exception *ExceptionInfo) *byte
var ImageToBlob func(i *ImageInfo, image *Image, length *uint32, exception *ExceptionInfo) *byte
var InjectImageBlob func(i *ImageInfo, image *Image, format string) bool

// Deprecated
var InterpretImageAttributes func(i *ImageInfo, image *Image, embedText string) string
var InterpretImageProperties func(i *ImageInfo, image *Image, embedText string) string
var InvokeDelegate func(i *ImageInfo, image *Image, decode, encode string, exception *ExceptionInfo) bool
var MontageImageList func(i *ImageInfo, montageInfo *MontageInfo, images *Image, exception *ExceptionInfo) *Image
var NewMagickImage func(i *ImageInfo, width, height Size, background *MagickPixelPacket) *Image
var OpenBlob func(i *ImageInfo, image *Image, mode BlobMode, exception *ExceptionInfo) bool
var OpenStream func(i *ImageInfo, streamInfo *StreamInfo, filename string, exception *ExceptionInfo) bool
var PingBlob func(i *ImageInfo, blob *Void, length uint32, exception *ExceptionInfo) *Image
var PingImage func(i *ImageInfo, exception *ExceptionInfo) *Image
var ReadImage func(i *ImageInfo, exception *ExceptionInfo) *Image
var ReadInlineImage func(i *ImageInfo, content string, exception *ExceptionInfo) *Image
var ReadStream func(i *ImageInfo, stream StreamHandler, exception *ExceptionInfo) *Image
var RemoteDisplayCommand func(i *ImageInfo, window, filename string, exception *ExceptionInfo) bool
var RemoveImageOption func(i *ImageInfo, option string) string
var ResetImageOptionIterator func(i *ImageInfo)
var SetImageInfo func(i *ImageInfo, rectify bool, exception *ExceptionInfo) bool
var SetImageInfoBlob func(i *ImageInfo, blob *Void, length uint32)
var SetImageInfoFile func(i *ImageInfo, file *FILE)
var SetImageInfoProgressMonitor func(i *ImageInfo, progressMonitor MagickProgressMonitor, clientData *Void) MagickProgressMonitor
var SetImageOption func(i *ImageInfo, option, value string) bool
var StreamImage func(i *ImageInfo, streamInfo *StreamInfo, exception *ExceptionInfo) *Image

// Deprecated
var TranslateText func(i *ImageInfo, image *Image, embedText string) string
var WriteImage func(i *ImageInfo, image *Image) bool
var WriteImages func(i *ImageInfo, images *Image, filename string, exception *ExceptionInfo) bool
var WriteStream func(i *ImageInfo, image *Image, stream StreamHandler) bool
var XImportImage func(i *ImageInfo, ximageInfo *XImportInfo) *Image
var AcquireImage func(i *ImageInfo) *Image
var AcquireNextImage func(i *ImageInfo, image *Image)
var GetImageInfoFile func(i *ImageInfo) *FILE
var GetMagickProperty func(i *ImageInfo, image *Image, property string) string
var PingImages func(i *ImageInfo, exception *ExceptionInfo) *Image
var ReadImages func(i *ImageInfo, exception *ExceptionInfo) *Image
var SyncImageSettings func(i *ImageInfo, image *Image) bool
var SyncImagesSettings func(i *ImageInfo, image *Image) bool

func (i *ImageInfo) AcquireQuantizeInfo() *QuantizeInfo { return AcquireQuantizeInfo(i) }
func (i *ImageInfo) AcquireQuantumInfo() *QuantumInfo   { return AcquireQuantumInfo(i) }
func (i *ImageInfo) AcquireStreamInfo() *StreamInfo     { return AcquireStreamInfo(i) }

// Deprecated
func (i *ImageInfo) AllocateImage() *Image { return AllocateImage(i) }

// Deprecated
func (i *ImageInfo) AllocateNextImage(image *Image)   { AllocateNextImage(i, image) }
func (i *ImageInfo) AnimateImages(images *Image) bool { return AnimateImages(i, images) }
func (i *ImageInfo) BlobToImage(blob *Void, length uint32, exception *ExceptionInfo) *Image {
	return BlobToImage(i, blob, length, exception)
}
func (i *ImageInfo) CloneDrawInfo(drawInfo *DrawInfo) *DrawInfo { return CloneDrawInfo(i, drawInfo) }
func (i *ImageInfo) CloneInfo() *ImageInfo                      { return CloneImageInfo(i) }
func (i *ImageInfo) CloneOptions(cloneInfo *ImageInfo) bool {
	return CloneImageOptions(i, cloneInfo)
}
func (i *ImageInfo) CloneMontageInfo(montageInfo *MontageInfo) *MontageInfo {
	return CloneMontageInfo(i, montageInfo)
}
func (i *ImageInfo) DefineOption(option string) bool  { return DefineImageOption(i, option) }
func (i *ImageInfo) DeleteOption(option string) bool  { return DeleteImageOption(i, option) }
func (i *ImageInfo) Destroy() *ImageInfo              { return DestroyImageInfo(i) }
func (i *ImageInfo) DestroyOptions()                  { DestroyImageOptions(i) }
func (i *ImageInfo) DisplayImages(images *Image) bool { return DisplayImages(i, images) }
func (i *ImageInfo) DelegateCommand(image *Image, decode, encode string, exception *ExceptionInfo) string {
	return GetDelegateCommand(i, image, decode, encode, exception)
}
func (i *ImageInfo) DrawInfo(drawInfo *DrawInfo)          { GetDrawInfo(i, drawInfo) }
func (i *ImageInfo) Get()                                 { GetImageInfo(i) }
func (i *ImageInfo) Option(key string) string             { return GetImageOption(i, key) }
func (i *ImageInfo) MontageInfo(montageInfo *MontageInfo) { GetMontageInfo(i, montageInfo) }
func (i *ImageInfo) NextOption() string                   { return GetNextImageOption(i) }
func (i *ImageInfo) QuantumInfo(quantumInfo *QuantumInfo) { GetQuantumInfo(i, quantumInfo) }
func (i *ImageInfo) HuffmanEncode(image *Image) bool      { return HuffmanEncodeImage(i, image) }
func (i *ImageInfo) ImagesToBlob(images *Image, length *uint32, exception *ExceptionInfo) *byte {
	return ImagesToBlob(i, images, length, exception)
}
func (i *ImageInfo) ImageToBlob(image *Image, length *uint32, exception *ExceptionInfo) *byte {
	return ImageToBlob(i, image, length, exception)
}
func (i *ImageInfo) InjectImageBlob(image *Image, format string) bool {
	return InjectImageBlob(i, image, format)
}

// Deprecated
func (i *ImageInfo) InterpretAttributes(image *Image, embedText string) string {
	return InterpretImageAttributes(i, image, embedText)
}
func (i *ImageInfo) InterpretProperties(image *Image, embedText string) string {
	return InterpretImageProperties(i, image, embedText)
}
func (i *ImageInfo) InvokeDelegate(image *Image, decode, encode string, exception *ExceptionInfo) bool {
	return InvokeDelegate(i, image, decode, encode, exception)
}
func (i *ImageInfo) MontageImageList(montageInfo *MontageInfo, images *Image, exception *ExceptionInfo) *Image {
	return MontageImageList(i, montageInfo, images, exception)
}
func (i *ImageInfo) NewMagickImage(width, height Size, background *MagickPixelPacket) *Image {
	return NewMagickImage(i, width, height, background)
}
func (i *ImageInfo) OpenBlob(image *Image, mode BlobMode, exception *ExceptionInfo) bool {
	return OpenBlob(i, image, mode, exception)
}
func (i *ImageInfo) OpenStream(streamInfo *StreamInfo, filename string, exception *ExceptionInfo) bool {
	return OpenStream(i, streamInfo, filename, exception)
}
func (i *ImageInfo) PingBlob(blob *Void, length uint32, exception *ExceptionInfo) *Image {
	return PingBlob(i, blob, length, exception)
}
func (i *ImageInfo) PingImage(exception *ExceptionInfo) *Image { return PingImage(i, exception) }
func (i *ImageInfo) ReadImage(exception *ExceptionInfo) *Image { return ReadImage(i, exception) }
func (i *ImageInfo) ReadInlineImage(content string, exception *ExceptionInfo) *Image {
	return ReadInlineImage(i, content, exception)
}
func (i *ImageInfo) ReadStream(stream StreamHandler, exception *ExceptionInfo) *Image {
	return ReadStream(i, stream, exception)
}
func (i *ImageInfo) RemoteDisplayCommand(window, filename string, exception *ExceptionInfo) bool {
	return RemoteDisplayCommand(i, window, filename, exception)
}
func (i *ImageInfo) RemoveImageOption(option string) string { return RemoveImageOption(i, option) }
func (i *ImageInfo) ResetImageOptionIterator()              { ResetImageOptionIterator(i) }
func (i *ImageInfo) SetImageInfo(rectify bool, exception *ExceptionInfo) bool {
	return SetImageInfo(i, rectify, exception)
}
func (i *ImageInfo) SetBlob(blob *Void, length uint32) { SetImageInfoBlob(i, blob, length) }
func (i *ImageInfo) SetFile(file *FILE)                { SetImageInfoFile(i, file) }
func (i *ImageInfo) SetProgressMonitor(progressMonitor MagickProgressMonitor, clientData *Void) MagickProgressMonitor {
	return SetImageInfoProgressMonitor(i, progressMonitor, clientData)
}
func (i *ImageInfo) SetOption(option, value string) bool {
	return SetImageOption(i, option, value)
}
func (i *ImageInfo) StreamImage(streamInfo *StreamInfo, exception *ExceptionInfo) *Image {
	return StreamImage(i, streamInfo, exception)
}

// Deprecate
func (i *ImageInfo) TranslateText(image *Image, embedText string) string {
	return TranslateText(i, image, embedText)
}
func (i *ImageInfo) WriteImage(image *Image) bool { return WriteImage(i, image) }
func (i *ImageInfo) WriteImages(images *Image, filename string, exception *ExceptionInfo) bool {
	return WriteImages(i, images, filename, exception)
}
func (i *ImageInfo) WriteStream(image *Image, stream StreamHandler) bool {
	return WriteStream(i, image, stream)
}
func (i *ImageInfo) XImportImage(ximageInfo *XImportInfo) *Image {
	return XImportImage(i, ximageInfo)
}
func (i *ImageInfo) AcquireImage() *Image          { return AcquireImage(i) }
func (i *ImageInfo) AcquireNextImage(image *Image) { AcquireNextImage(i, image) }
func (i *ImageInfo) GetImageInfoFile() *FILE       { return GetImageInfoFile(i) }
func (i *ImageInfo) GetMagickProperty(image *Image, property string) string {
	return GetMagickProperty(i, image, property)
}
func (i *ImageInfo) PingImages(exception *ExceptionInfo) *Image { return PingImages(i, exception) }
func (i *ImageInfo) ReadImages(exception *ExceptionInfo) *Image { return ReadImages(i, exception) }
func (i *ImageInfo) SyncImageSettings(image *Image) bool        { return SyncImageSettings(i, image) }
func (i *ImageInfo) SyncImagesSettings(image *Image) bool       { return SyncImagesSettings(i, image) }

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
func (x *XMLTreeInfo) Destroy() *XMLTreeInfo       { return DestroyXMLTree(x) }
func (x *XMLTreeInfo) NextTag() *XMLTreeInfo       { return GetNextXMLTreeTag(x) }
func (x *XMLTreeInfo) Attribute(tag string) string { return GetXMLTreeAttribute(x, tag) }
func (x *XMLTreeInfo) Attributes(attributes *SplayTreeInfo) bool {
	return GetXMLTreeAttributes(x, attributes)
}
func (x *XMLTreeInfo) Child(tag string) *XMLTreeInfo { return GetXMLTreeChild(x, tag) }
func (x *XMLTreeInfo) Content() string               { return GetXMLTreeContent(x) }
func (x *XMLTreeInfo) Ordered() *XMLTreeInfo         { return GetXMLTreeOrdered(x) }
func (x *XMLTreeInfo) Path(path string) *XMLTreeInfo { return GetXMLTreePath(x, path) }
func (x *XMLTreeInfo) ProcessingInstructions(target string) []string {
	return GetXMLTreeProcessingInstructions(x, target)
}
func (x *XMLTreeInfo) Sibling() *XMLTreeInfo { return GetXMLTreeSibling(x) }
func (x *XMLTreeInfo) Tag() string           { return GetXMLTreeTag(x) }
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

var NewLinkedList func(capacity Size) *LinkedListInfo
var DestroyConfigureOptions func(options *LinkedListInfo) *LinkedListInfo
var DestroyLocaleOptions func(messages *LinkedListInfo) *LinkedListInfo
var AppendValueToLinkedList func(l *LinkedListInfo, value *Void) bool
var ClearLinkedList func(l *LinkedListInfo, relinquishValue func(*Void) *Void)
var DestroyLinkedList func(l *LinkedListInfo, relinquishValue func(*Void) *Void)
var GetLastValueInLinkedList func(l *LinkedListInfo) *Void
var GetNextValueInLinkedList func(l *LinkedListInfo) *Void
var GetNumberOfElementsInLinkedList func(l *LinkedListInfo) Size
var GetValueFromLinkedList func(l *LinkedListInfo, index Size) *Void
var InsertValueInLinkedList func(l *LinkedListInfo, index Size, value *Void) bool
var InsertValueInSortedLinkedList func(l *LinkedListInfo, compare func(*Void, *Void) int, replace **Void, value *Void) bool
var IsLinkedListEmpty func(l *LinkedListInfo) bool
var LinkedListToArray func(l *LinkedListInfo, array **Void) bool
var RemoveElementByValueFromLinkedList func(l *LinkedListInfo, value *Void) *Void
var RemoveElementFromLinkedList func(l *LinkedListInfo, index Size) *Void
var RemoveLastElementFromLinkedList func(l *LinkedListInfo) *Void
var ResetLinkedListIterator func(l *LinkedListInfo)

func (l *LinkedListInfo) DestroyConfigureOptions() *LinkedListInfo { return DestroyConfigureOptions(l) }
func (l *LinkedListInfo) DestroyLocaleOptions() *LinkedListInfo    { return DestroyLocaleOptions(l) }
func (l *LinkedListInfo) AppendValue(value *Void) bool {
	return AppendValueToLinkedList(l, value)
}
func (l *LinkedListInfo) Clear(relinquishValue func(*Void) *Void) {
	ClearLinkedList(l, relinquishValue)
}
func (l *LinkedListInfo) Destroy(relinquishValue func(*Void) *Void) {
	DestroyLinkedList(l, relinquishValue)
}
func (l *LinkedListInfo) LastValue() *Void       { return GetLastValueInLinkedList(l) }
func (l *LinkedListInfo) NextValue() *Void       { return GetNextValueInLinkedList(l) }
func (l *LinkedListInfo) NumberOfElements() Size { return GetNumberOfElementsInLinkedList(l) }
func (l *LinkedListInfo) Value(index Size) *Void { return GetValueFromLinkedList(l, index) }
func (l *LinkedListInfo) InsertValue(index Size, value *Void) bool {
	return InsertValueInLinkedList(l, index, value)
}
func (l *LinkedListInfo) InsertValueSorted(compare func(*Void, *Void) int, replace **Void, value *Void) bool {
	return InsertValueInSortedLinkedList(l, compare, replace, value)
}
func (l *LinkedListInfo) Empty() bool             { return IsLinkedListEmpty(l) }
func (l *LinkedListInfo) Array(array **Void) bool { return LinkedListToArray(l, array) }
func (l *LinkedListInfo) RemoveElementByValue(value *Void) *Void {
	return RemoveElementByValueFromLinkedList(l, value)
}
func (l *LinkedListInfo) RemoveElement(index Size) *Void {
	return RemoveElementFromLinkedList(l, index)
}
func (l *LinkedListInfo) RemoveLastElement() *Void { return RemoveLastElementFromLinkedList(l) }
func (l *LinkedListInfo) ResetIterator()           { ResetLinkedListIterator(l) }

var GetDelegateInfo func(decode, encode string, exception *ExceptionInfo) *DelegateInfo
var GetDelegateInfoList func(pattern string, numberDelegates *Size, exception *ExceptionInfo) **DelegateInfo
var GetDelegateCommands func(d *DelegateInfo) string
var GetDelegateMode func(d *DelegateInfo) SSize
var GetDelegateThreadSupport func(d *DelegateInfo) bool

func (d *DelegateInfo) Commands() string    { return GetDelegateCommands(d) }
func (d *DelegateInfo) Mode() SSize         { return GetDelegateMode(d) }
func (d *DelegateInfo) ThreadSupport() bool { return GetDelegateThreadSupport(d) }

var NewSplayTree func(compare func(*Void, *Void) int, relinquishKey, relinquishValue func(*Void) *Void) *SplayTreeInfo
var AddValueToSplayTree func(s *SplayTreeInfo, key, value *Void) bool
var CloneSplayTree func(s *SplayTreeInfo, cloneKey, cloneValue func(*Void) *Void) *SplayTreeInfo
var DeleteNodeByValueFromSplayTree func(s *SplayTreeInfo, value *Void) bool
var DeleteNodeFromSplayTree func(s *SplayTreeInfo, key *Void) bool
var DestroySplayTree func(s *SplayTreeInfo) *SplayTreeInfo
var GetNextKeyInSplayTree func(s *SplayTreeInfo) *Void
var GetNextValueInSplayTree func(s *SplayTreeInfo) *Void
var GetNumberOfNodesInSplayTree func(s *SplayTreeInfo) Size
var GetValueFromSplayTree func(s *SplayTreeInfo, key *Void) *Void
var RemoveNodeByValueFromSplayTree func(s *SplayTreeInfo, value *Void) *Void
var RemoveNodeFromSplayTree func(s *SplayTreeInfo, key *Void) *Void
var ResetSplayTreeIterator func(s *SplayTreeInfo)
var ResetSplayTree func(s *SplayTreeInfo)

func (s *SplayTreeInfo) Reset()                         { ResetSplayTree(s) }
func (s *SplayTreeInfo) AddValue(key, value *Void) bool { return AddValueToSplayTree(s, key, value) }
func (s *SplayTreeInfo) Clone(cloneKey, cloneValue func(*Void) *Void) *SplayTreeInfo {
	return CloneSplayTree(s, cloneKey, cloneValue)
}
func (s *SplayTreeInfo) DeleteNodeByValue(value *Void) bool {
	return DeleteNodeByValueFromSplayTree(s, value)
}
func (s *SplayTreeInfo) DeleteNode(key *Void) bool { return DeleteNodeFromSplayTree(s, key) }
func (s *SplayTreeInfo) Destroy() *SplayTreeInfo   { return DestroySplayTree(s) }
func (s *SplayTreeInfo) NextKey() *Void            { return GetNextKeyInSplayTree(s) }
func (s *SplayTreeInfo) NextValue() *Void          { return GetNextValueInSplayTree(s) }
func (s *SplayTreeInfo) NumberOfNodes() Size       { return GetNumberOfNodesInSplayTree(s) }
func (s *SplayTreeInfo) Value(key *Void) *Void     { return GetValueFromSplayTree(s, key) }
func (s *SplayTreeInfo) RemoveNodeByValue(value *Void) *Void {
	return RemoveNodeByValueFromSplayTree(s, value)
}
func (s *SplayTreeInfo) RemoveNode(key *Void) *Void { return RemoveNodeFromSplayTree(s, key) }
func (s *SplayTreeInfo) ResetIterator()             { ResetSplayTreeIterator(s) }

var DestroyHashmap func(h *HashmapInfo) *HashmapInfo

func (h *HashmapInfo) Destroy() *HashmapInfo { return DestroyHashmap(h) }

var GetNextKeyInHashmap func(h *HashmapInfo) *Void

func (h *HashmapInfo) NextKey() *Void { return GetNextKeyInHashmap(h) }

var GetNextValueInHashmap func(h *HashmapInfo) *Void

func (h *HashmapInfo) NextValue() *Void { return GetNextValueInHashmap(h) }

var GetNumberOfEntriesInHashmap func(h *HashmapInfo) Size

func (h *HashmapInfo) NumberOfEntries() Size { return GetNumberOfEntriesInHashmap(h) }

var GetValueFromHashmap func(h *HashmapInfo, key *Void) *Void

func (h *HashmapInfo) Value(key *Void) *Void { return GetValueFromHashmap(h, key) }

var IsHashmapEmpty func(h *HashmapInfo) bool

func (h *HashmapInfo) Empty() bool { return IsHashmapEmpty(h) }

var PutEntryInHashmap func(h *HashmapInfo, key, value *Void) bool

func (h *HashmapInfo) PutEntry(key, value *Void) bool { return PutEntryInHashmap(h, key, value) }

var RemoveEntryFromHashmap func(h *HashmapInfo, key *Void) *Void

func (h *HashmapInfo) RemoveEntry(key *Void) *Void { return RemoveEntryFromHashmap(h, key) }

var ResetHashmapIterator func(h *HashmapInfo)

func (h *HashmapInfo) ResetIterator() { ResetHashmapIterator(h) }

var CloneImageView func(i *ImageView) *ImageView

func (i *ImageView) Clone() *ImageView { return CloneImageView(i) }

var DestroyImageView func(i *ImageView) *ImageView

func (i *ImageView) Destroy() *ImageView { return DestroyImageView(i) }

var DuplexTransferImageViewIterator func(i *ImageView, duplex *ImageView, destination *ImageView, transfer DuplexTransferImageViewMethod, context *Void) bool

func (i *ImageView) DuplexTransferIterator(duplex *ImageView, destination *ImageView, transfer DuplexTransferImageViewMethod, context *Void) bool {
	return DuplexTransferImageViewIterator(i, duplex, destination, transfer, context)
}

var GetImageViewAuthenticIndexes func(i *ImageView) *IndexPacket

func (i *ImageView) AuthenticIndexes() *IndexPacket { return GetImageViewAuthenticIndexes(i) }

var GetImageViewAuthenticPixels func(i *ImageView) *PixelPacket

func (i *ImageView) AuthenticPixels() *PixelPacket { return GetImageViewAuthenticPixels(i) }

var GetImageViewExtent func(i *ImageView) *RectangleInfo // doc not ptr (RectangleInfo)
func (i *ImageView) Extent() *RectangleInfo              { return GetImageViewExtent(i) }

var GetImageViewImage func(i *ImageView) *Image

func (i *ImageView) Image() *Image { return GetImageViewImage(i) }

var GetImageViewIterator func(i *ImageView, get GetImageViewMethod, context *Void) bool

func (i *ImageView) Iterator(get GetImageViewMethod, context *Void) bool {
	return GetImageViewIterator(i, get, context)
}

var GetImageViewVirtualIndexes func(i *ImageView) *IndexPacket

func (i *ImageView) VirtualIndexes() *IndexPacket { return GetImageViewVirtualIndexes(i) }

var GetImageViewVirtualPixels func(i *ImageView) *PixelPacket

func (i *ImageView) VirtualPixels() *PixelPacket { return GetImageViewVirtualPixels(i) }

var IsImageView func(i *ImageView) bool

func (i *ImageView) IsImageView() bool { return IsImageView(i) }

var SetImageViewDescription func(i *ImageView, description string)

func (i *ImageView) SetDescription(description string) { SetImageViewDescription(i, description) }

var SetImageViewIterator func(i *ImageView, s SetImageViewMethod, context *Void) bool

func (i *ImageView) SetIterator(set SetImageViewMethod, context *Void) bool {
	return SetImageViewIterator(i, set, context)
}

var SetImageViewThreads func(i *ImageView, numberThreads uint32)

func (i *ImageView) SetThreads(numberThreads uint32) { SetImageViewThreads(i, numberThreads) }

var TransferImageViewIterator func(i *ImageView, destination *ImageView, transfer TransferImageViewMethod, context *Void) bool

func (i *ImageView) TransferIterator(destination *ImageView, transfer TransferImageViewMethod, context *Void) bool {
	return TransferImageViewIterator(i, destination, transfer, context)
}

var UpdateImageViewIterator func(i *ImageView, update UpdateImageViewMethod, context *Void) bool

func (i *ImageView) UpdateIterator(update UpdateImageViewMethod, context *Void) bool {
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

// Deprecated
var AffinityImage func(q *QuantizeInfo, image *Image, affinityImage *Image) bool

// Deprecated
func (q *QuantizeInfo) AffinityImage(image *Image, affinityImage *Image) bool {
	return AffinityImage(q, image, affinityImage)
}

// Deprecated
var AffinityImages func(q *QuantizeInfo, images *Image, affinityImage *Image) bool

// Deprecated
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

var GetStreamInfoClientData func(s *StreamInfo) *Void

func (s *StreamInfo) GetStreamInfoClientData() *Void { return GetStreamInfoClientData(s) }

var SetStreamInfoClientData func(s *StreamInfo, clientData *Void)

func (s *StreamInfo) SetStreamInfoClientData(clientData *Void) {
	SetStreamInfoClientData(s, clientData)
}

var SetStreamInfoMap func(s *StreamInfo, map_ string)

func (s *StreamInfo) SetStreamInfoMap(map_ string) { SetStreamInfoMap(s, map_) }

var SetStreamInfoStorageType func(s *StreamInfo, storageType StorageType)

func (s *StreamInfo) SetStreamInfoStorageType(storageType StorageType) {
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

// Deprecated
var SetExceptionInfo func(e *ExceptionInfo, severity ExceptionType) bool

// Deprecated
func (e *ExceptionInfo) Set(severity ExceptionType) bool { return SetExceptionInfo(e, severity) }

var ThrowException func(e *ExceptionInfo, severity ExceptionType, reason, description string) bool

func (e *ExceptionInfo) Throw(severity ExceptionType, reason, description string) bool {
	return ThrowException(e, severity, reason, description)
}

var ThrowMagickException func(e *ExceptionInfo, module, function string, line Size, severity ExceptionType, tag, format string, va ...VArg) bool

func (e *ExceptionInfo) ThrowMagick(module, function string, line Size, severity ExceptionType, tag, format string, va ...VArg) bool {
	return ThrowMagickException(e, module, function, line, severity, tag, format, va)
}

var ThrowMagickExceptionList func(e *ExceptionInfo, module, function string, line Size, severity ExceptionType, tag, format string, operands VAList) bool

func (e *ExceptionInfo) ThrowMagickList(module, function string, line Size, severity ExceptionType, tag, format string, operands VAList) bool {
	return ThrowMagickExceptionList(e, module, function, line, severity, tag, format, operands)
}

var XGetResourceClass func(x *XrmDatabase, clientName, keyword, resourceDefault string) string

func (x *XrmDatabase) ResourceClass(clientName, keyword, resourceDefault string) string {
	return XGetResourceClass(x, clientName, keyword, resourceDefault)
}

var XGetResourceInfo func(x *XrmDatabase, clientName string, resourceInfo *XResourceInfo)

func (x *XrmDatabase) ResourceInfo(clientName string, resourceInfo *XResourceInfo) {
	XGetResourceInfo(x, clientName, resourceInfo)
}

var XGetResourceInstance func(x *XrmDatabase, clientName, keyword, resourceDefault string) string

func (x *XrmDatabase) ResourceInstance(clientName, keyword, resourceDefault string) string {
	return XGetResourceInstance(x, clientName, keyword, resourceDefault)
}

var XrmCombineDatabase func(x *XrmDatabase, target **XrmDatabase, override bool)

func (x *XrmDatabase) Combine(target **XrmDatabase, override bool) {
	XrmCombineDatabase(x, target, override)
}

var AcquireMagickResource func(r ResourceType, size MagickSizeType) bool

func (r ResourceType) AcquireResource(size MagickSizeType) bool {
	return AcquireMagickResource(r, size)
}

var GetMagickResource func(r ResourceType) MagickSizeType

func (r ResourceType) Resource() MagickSizeType { return GetMagickResource(r) }

var GetMagickResourceLimit func(r ResourceType) MagickSizeType

func (r ResourceType) Limit() MagickSizeType { return GetMagickResourceLimit(r) }

var RelinquishMagickResource func(r ResourceType, size MagickSizeType)

func (r ResourceType) Relinquish(size MagickSizeType) { RelinquishMagickResource(r, size) }

var SetMagickResourceLimit func(r ResourceType, limit MagickSizeType) bool

func (r ResourceType) SetLimit(limit MagickSizeType) bool { return SetMagickResourceLimit(r, limit) }

var AttachBlob func(b *BlobInfo, blob *Void, length uint32)

func (b *BlobInfo) Attach(blob *Void, length uint32) { AttachBlob(b, blob, length) }

var CloneBlobInfo func(b *BlobInfo) *BlobInfo

func (b *BlobInfo) Clone() *BlobInfo { return CloneBlobInfo(b) }

var DetachBlob func(b *BlobInfo) *byte

func (b *BlobInfo) Detach() *byte { return DetachBlob(b) }

var GetBlobInfo func(b *BlobInfo)

func (b *BlobInfo) Get() { GetBlobInfo(b) }

var ReferenceBlob func(b *BlobInfo) *BlobInfo

func (b *BlobInfo) Reference() *BlobInfo { return ReferenceBlob(b) }

var ConcatenateColorComponent func(m *MagickPixelPacket, channel ChannelType, compliance ComplianceType, tuple string)

func (m *MagickPixelPacket) ConcatenateColorComponent(channel ChannelType, compliance ComplianceType, tuple string) {
	ConcatenateColorComponent(m, channel, compliance, tuple)
}

var GetColorTuple func(m *MagickPixelPacket, hex bool, tuple string)

func (m *MagickPixelPacket) ColorTuple(hex bool, tuple string) { GetColorTuple(m, hex, tuple) }

var IsMagickColorSimilar func(p *MagickPixelPacket, q *MagickPixelPacket) bool

func (p *MagickPixelPacket) ColorSimilar(q *MagickPixelPacket) bool { return IsMagickColorSimilar(p, q) }

var CloneMagickPixelPacket func(m *MagickPixelPacket) *MagickPixelPacket

func (m *MagickPixelPacket) Clone() *MagickPixelPacket { return CloneMagickPixelPacket(m) }

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

// Deprecated
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

var GetResizeFilterSupport func(r *ResizeFilter) MagickRealType

func (r *ResizeFilter) Support() MagickRealType { return GetResizeFilterSupport(r) }

var GetResizeFilterWeight func(r *ResizeFilter, x MagickRealType) MagickRealType

func (r *ResizeFilter) Weight(x MagickRealType) MagickRealType { return GetResizeFilterWeight(r, x) }

var SetResizeFilterSupport func(r *ResizeFilter, support MagickRealType)

func (r *ResizeFilter) SetSupport(support MagickRealType) { SetResizeFilterSupport(r, support) }

var DestroyResampleFilter func(r *ResampleFilter) *ResampleFilter
var ResamplePixelColor func(r *ResampleFilter, u0, v0 float64) MagickPixelPacket
var ScaleResampleFilter func(r *ResampleFilter, dux, duy, dvx, dvy float64)
var SetResampleFilterInterpolateMethod func(r *ResampleFilter, method InterpolatePixelMethod) bool
var SetResampleFilterVirtualPixelMethod func(r *ResampleFilter, method VirtualPixelMethod) bool

func (r *ResampleFilter) Destroy() *ResampleFilter { return DestroyResampleFilter(r) }
func (r *ResampleFilter) PixelColor(u0, v0 float64) MagickPixelPacket {
	return ResamplePixelColor(r, u0, v0)
}
func (r *ResampleFilter) Scale(dux, duy, dvx, dvy float64) {
	ScaleResampleFilter(r, dux, duy, dvx, dvy)
}
func (r *ResampleFilter) InterpolateMethod(method InterpolatePixelMethod) bool {
	return SetResampleFilterInterpolateMethod(r, method)
}
func (r *ResampleFilter) VirtualPixelMethod(method VirtualPixelMethod) bool {
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

func (s *SignatureInfo) Destroy() *SignatureInfo             { return DestroySignatureInfo(s) }
func (s *SignatureInfo) Finalize()                           { FinalizeSignature(s) }
func (s *SignatureInfo) Get()                                { GetSignatureInfo(s) }
func (s *SignatureInfo) Update(message *byte, length uint32) { UpdateSignature(s, message, length) }

var DestroyFxInfo func(f *FxInfo) *FxInfo
var FxEvaluateChannelExpression func(f *FxInfo, channel ChannelType, x, y SSize, alpha *MagickRealType, exception *ExceptionInfo) bool
var FxEvaluateExpression func(f *FxInfo, alpha *MagickRealType, exception *ExceptionInfo) bool

func (f *FxInfo) Destroy() *FxInfo { return DestroyFxInfo(f) }
func (f *FxInfo) EvaluateChannelExpression(channel ChannelType, x, y SSize, alpha *MagickRealType, exception *ExceptionInfo) bool {
	return FxEvaluateChannelExpression(f, channel, x, y, alpha, exception)
}
func (f *FxInfo) EvaluateExpression(alpha *MagickRealType, exception *ExceptionInfo) bool {
	return FxEvaluateExpression(f, alpha, exception)
}

var GetCommandOptions func(value CommandOption) []string
var CommandOptionToMnemonic func(option CommandOption, type_ SSize) string
var ParseCommandOption func(option CommandOption, list bool, options string) SSize

func (m CommandOption) Options() []string           { return GetCommandOptions(m) }
func (m CommandOption) Mnemonic(type_ SSize) string { return CommandOptionToMnemonic(m, type_) }
func (m CommandOption) Parse(list bool, options string) SSize {
	return ParseCommandOption(m, list, options)
}

var AcquirePixelCachePixels func(i *Image, m *MagickSizeType, e *ExceptionInfo) *Void

func (i *Image) AcquirePixelCachePixels(m *MagickSizeType, e *ExceptionInfo) *Void {
	return AcquirePixelCachePixels(i, m, e)
}

var AcquireRandomInfo func() *RandomInfo
var DestroyRandomInfo func(*RandomInfo) *RandomInfo
var AnnotateComponentGenesis func() bool
var AnnotateComponentTerminus func()
var AsynchronousResourceComponentTerminus func()
var BlackThresholdImageChannel func(*Image, ChannelType, string, *ExceptionInfo) bool

func (i *Image) BlackThresholdChannel(c ChannelType, threshold string, e *ExceptionInfo) bool {
	return BlackThresholdImageChannel(i, c, threshold, e)
}

var BlobToStringInfo func(*Void, uint32) *StringInfo
var CacheComponentGenesis func() bool
var CacheComponentTerminus func()
var ClampImage func(i *Image) bool
var ClampImageChannel func(i *Image, c ChannelType) bool

func (i *Image) ClampImage() bool                     { return ClampImage(i) }
func (i *Image) ClampImageChannel(c ChannelType) bool { return ClampImageChannel(i, c) }

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
var ConvertHCLToRGB func(float64, float64, float64, *Quantum, *Quantum, *Quantum)
var ConvertRGBToHCL func(Quantum, Quantum, Quantum, *float64, *float64, *float64)
var DelegateComponentGenesis func() bool
var DelegateComponentTerminus func()
var DestroyPixelCacheNexus func(**NexusInfo, uint32) **NexusInfo
var DiscardBlobBytes func(*Image, MagickSizeType) bool
var DistortResizeImage func(*Image, uint32, uint32, *ExceptionInfo) *Image
var DuplicateBlob func(*Image, *Image)

func (i *Image) DiscardBlobBytes(length MagickSizeType) bool { return DiscardBlobBytes(i, length) }
func (i *Image) DistortResizeImage(columns, rows uint32, e *ExceptionInfo) *Image {
	return DistortResizeImage(i, columns, rows, e)
}
func (i *Image) DuplicateBlob(duplicate *Image) { DuplicateBlob(i, duplicate) }

var FormatLocaleFile func(f *FILE, format string, va ...VArg) int32
var FormatLocaleFileList func(f *FILE, format string, va VAList) int32

func (f *FILE) FormatLocaleFile(format string, va ...VArg) int32 {
	return FormatLocaleFile(f, format, va)
}
func (f *FILE) FormatLocaleFileList(format string, va VAList) int32 {
	return FormatLocaleFileList(f, format, va)
}

var FormatLocaleString func(string, uint32, string, ...VArg) int32
var FormatLocaleStringList func(string, uint32, string, VAList) int32
var GenerateDifferentialNoise func(*RandomInfo, Quantum, NoiseType, MagickRealType) float64
var GetAuthenticPixelCacheNexus func(*Image, int32, int32, uint32, uint32, *NexusInfo, *ExceptionInfo) *PixelPacket
var GetCacheViewExtent func(*CacheView) MagickSizeType
var GetCommandOptionFlags func(CommandOption, bool, string) int32
var GetConfigureOption func(string) string
var GetImageExtent func(i *Image) MagickSizeType
var GetImageKurtosis func(i *Image, kurtosis, skewness *float64, e *ExceptionInfo) bool
var GetImagePixelCacheType func(i *Image) CacheType
var GetImageReferenceCount func(i *Image) int32

func (i *Image) Extent() MagickSizeType { return GetImageExtent(i) }
func (i *Image) Kurtosis(kurtosis, skewness *float64, e *ExceptionInfo) bool {
	return GetImageKurtosis(i, kurtosis, skewness, e)
}
func (i *Image) ImagePixelCacheType() CacheType { return GetImagePixelCacheType(i) }
func (i *Image) ReferenceCount() int32          { return GetImageReferenceCount(i) }

var GetMagickPageSize func() int32
var GetMagickRawSupport func(*MagickInfo) bool
var GetPathAttributes func(string, *Void) bool
var GetPixelCacheChannels func(Cache) uint32
var GetPixelCacheColorspace func(Cache) ColorspaceType
var GetPixelCacheMethods func(*CacheMethods)
var GetPixelCacheNexusExtent func(Cache, *NexusInfo) MagickSizeType
var GetPixelCachePixels func(i *Image, length *MagickSizeType, e *ExceptionInfo) *Void

func (i *Image) CachePixels(length *MagickSizeType, e *ExceptionInfo) *Void {
	return GetPixelCachePixels(i, length, e)
}

var GetPixelCacheStorageClass func(Cache) ClassType
var GetPixelCacheTileSize func(i *Image, width, height *uint32)
var GetPixelCacheType func(i *Image) CacheType
var GetPixelCacheVirtualMethod func(i *Image) VirtualPixelMethod

func (i *Image) CacheTileSize(width, height *uint32)    { GetPixelCacheTileSize(i, width, height) }
func (i *Image) CacheType() CacheType                   { return GetPixelCacheType(i) }
func (i *Image) CacheVirtualMethod() VirtualPixelMethod { return GetPixelCacheVirtualMethod(i) }

var GetPolicyInfoList func(string, *uint32, *ExceptionInfo) []*PolicyInfo
var GetPolicyList func(string, *uint32, *ExceptionInfo) []string
var GetPolicyValue func(string) string
var GetPseudoRandomValue func(*RandomInfo) float64
var GetQuantumEndian func(*QuantumInfo) EndianType
var GetQuantumExtent func(*Image, *QuantumInfo, QuantumType) uint32
var GetQuantumFormat func(*QuantumInfo) QuantumFormatType
var GetQuantumPixels func(*QuantumInfo) *byte
var GetQuantumType func(i *Image, e *ExceptionInfo) QuantumType

func (i *Image) QuantumType(e *ExceptionInfo) QuantumType { return GetQuantumType(i, e) }

var GetRandomSecretKey func(*RandomInfo) Size
var GetSignatureBlocksize func(*SignatureInfo) uint
var GetSignatureDigest func(*SignatureInfo) *StringInfo
var GetSignatureDigestsize func(*SignatureInfo) uint
var Gettimeofday func(*Timeval, *Timezone)
var GetVirtualIndexesFromNexus func(Cache, *NexusInfo) *IndexPacket
var GetVirtualPixelsFromNexus func(i *Image, v VirtualPixelMethod, x, y int32, columns, rows uint32, n *NexusInfo, e *ExceptionInfo) *PixelPacket

func (i *Image) VirtualPixelsFromNexus(m VirtualPixelMethod, x, y int32, columns, rows uint32, n *NexusInfo, e *ExceptionInfo) *PixelPacket {
	return GetVirtualPixelsFromNexus(i, m, x, y, columns, rows, n, e)
}

var GetVirtualPixelsNexus func(Cache, *NexusInfo) *PixelPacket
var InitializeSignature func(*SignatureInfo)
var InterpolateMagickPixelPacket func(i *Image, v *CacheView, m InterpolatePixelMethod, x, y float64, p *MagickPixelPacket, e *ExceptionInfo) bool
var InterpolativeResizeImage func(i *Image, columns, rows uint32, m InterpolatePixelMethod, e *ExceptionInfo) *Image

func (i *Image) InterpolatePacket(v *CacheView, m InterpolatePixelMethod, x, y float64, p *MagickPixelPacket, e *ExceptionInfo) bool {
	return InterpolateMagickPixelPacket(i, v, m, x, y, p, e)
}
func (i *Image) InterpolativeResize(columns, rows uint32, m InterpolatePixelMethod, e *ExceptionInfo) *Image {
	return InterpolativeResizeImage(i, columns, rows, m, e)
}

var InterpretLocaleValue func(string, []string) float64
var InterpretSiPrefixValue func(string, []string) float64
var InversesRGBCompandor func(MagickRealType) MagickRealType
var IsCommandOption func(string) bool
var IsPathAccessible func(string) bool
var IsRightsAuthorized func(PolicyDomain, PolicyRights, string) bool
var IsStringNotFalse func(string) bool
var IsStringTrue func(string) bool
var ListPolicyInfo func(f *FILE, e *ExceptionInfo) bool

func (f *FILE) ListPolicyInfo(e *ExceptionInfo) bool { return ListPolicyInfo(f, e) }

var LocaleComponentGenesis func() bool
var LocaleComponentTerminus func()
var LogComponentGenesis func() bool
var LogComponentTerminus func()
var MagicComponentGenesis func() bool
var MagicComponentTerminus func()
var MagickComponentGenesis func() bool
var MagickComponentTerminus func()
var MagickCreateThreadKey func(*MagickThreadKey) bool
var MagickDelay func(MagickSizeType)
var MagickDeleteThreadKey func(MagickThreadKey) bool
var MagickGetThreadValue func(MagickThreadKey) *Void
var MagickSetThreadValue func(MagickThreadKey, *Void) bool
var MimeComponentGenesis func() bool
var MimeComponentTerminus func()
var ModuleComponentGenesis func() bool
var ModuleComponentTerminus func()
var NTArgvToUTF8 func(argc int, argv []WString) []string
var NTGatherRandomData func(uint32, *byte) bool
var ParseRegionGeometry func(i *Image, geometry string, r *RectangleInfo, e *ExceptionInfo) MagickStatusType
var PerceptibleImage func(i *Image, epsilon float64) bool
var PerceptibleImageChannel func(i *Image, c ChannelType, epsilon float64) bool
var PersistPixelCache func(i *Image, filename string, attach bool, o *MagickOffsetType, e *ExceptionInfo) bool

func (i *Image) ParseRegionGeometry(geometry string, r *RectangleInfo, e *ExceptionInfo) MagickStatusType {
	return ParseRegionGeometry(i, geometry, r, e)
}
func (i *Image) Perceptible(epsilon float64) bool { return PerceptibleImage(i, epsilon) }
func (i *Image) PerceptibleChannel(c ChannelType, epsilon float64) bool {
	return PerceptibleImageChannel(i, c, epsilon)
}
func (i *Image) PersistPixelCache(filename string, attach bool, o *MagickOffsetType, e *ExceptionInfo) bool {
	return PersistPixelCache(i, filename, attach, o, e)
}

var PolicyComponentGenesis func() bool
var PolicyComponentTerminus func()
var PolynomialImage func(i *Image, nTerms uint32, terms []float64, e *ExceptionInfo) *Image
var PolynomialImageChannel func(i *Image, c ChannelType, nTerms uint32, terms []float64, e *ExceptionInfo) *Image

func (i *Image) Polynomial(nTerms uint32, terms []float64, e *ExceptionInfo) *Image {
	return PolynomialImage(i, nTerms, terms, e)
}
func (i *Image) PolynomialChannel(c ChannelType, nTerms uint32, terms []float64, e *ExceptionInfo) *Image {
	return PolynomialImageChannel(i, c, nTerms, terms, e)
}

var QueryMagickColorCompliance func(string, ComplianceType, *MagickPixelPacket, *ExceptionInfo) bool
var QueueAuthenticPixel func(i *Image, x, y int32, columns, rows uint32, clone bool, n *NexusInfo, e *ExceptionInfo) *PixelPacket
var QueueAuthenticPixelCacheNexus func(i *Image, x, y int32, columns, rows uint32, clone bool, n *NexusInfo, e *ExceptionInfo) *PixelPacket

func (i *Image) QueueAuthentic(x, y int32, columns, rows uint32, clone bool, n *NexusInfo, e *ExceptionInfo) *PixelPacket {
	return QueueAuthenticPixel(i, x, y, columns, rows, clone, n, e)
}
func (i *Image) QueueAuthenticNexus(x, y int32, columns, rows uint32, clone bool, n *NexusInfo, e *ExceptionInfo) *PixelPacket {
	return QueueAuthenticPixelCacheNexus(i, x, y, columns, rows, clone, n, e)
}

var RandomComponentGenesis func() bool
var RandomComponentTerminus func()
var ReadBlobMSBLongLong func(i *Image) MagickSizeType

func (i *Image) ReadBlobMSBLongLong() MagickSizeType { return ReadBlobMSBLongLong(i) }

var ReferencePixelCache func(Cache) Cache
var RegistryComponentGenesis func() bool
var ResetImageOptions func(*ImageInfo)
var ResourceComponentGenesis func() bool
var ResourceComponentTerminus func()
var SeedPseudoRandomGenerator func(Size)
var SemaphoreComponentGenesis func() bool
var SemaphoreComponentTerminus func()
var SetBlobExtent func(i *Image, extent MagickSizeType) bool

func (i *Image) SetBlobExtent(extent MagickSizeType) bool { return SetBlobExtent(i, extent) }

var SetPixelCacheMethods func(Cache, *CacheMethods)
var SetPixelCacheVirtualMethod func(i *Image, m VirtualPixelMethod) VirtualPixelMethod

func (i *Image) SetPixelCacheVirtualMethod(m VirtualPixelMethod) VirtualPixelMethod {
	return SetPixelCacheVirtualMethod(i, m)
}

var SetQuantumAlphaType func(*QuantumInfo, QuantumAlphaType)
var SetQuantumDepth func(i *Image, q *QuantumInfo, depth uint32) bool
var SetQuantumEndian func(i *Image, q *QuantumInfo, e EndianType) bool
var SetQuantumFormat func(i *Image, q *QuantumInfo, f QuantumFormatType) bool
var SetQuantumImageType func(i *Image, q QuantumType)

func (i *Image) SetQuantumDepth(q *QuantumInfo, depth uint32) bool {
	return SetQuantumDepth(i, q, depth)
}
func (i *Image) SetQuantumEndian(q *QuantumInfo, e EndianType) bool { return SetQuantumEndian(i, q, e) }
func (i *Image) SetQuantumFormat(q *QuantumInfo, f QuantumFormatType) bool {
	return SetQuantumFormat(i, q, f)
}
func (i *Image) SetQuantumType(q QuantumType) { SetQuantumImageType(i, q) }

var SetQuantumMinIsWhite func(*QuantumInfo, bool)
var SetQuantumPack func(*QuantumInfo, bool)
var SetQuantumPad func(i *Image, q *QuantumInfo, pad uint32) bool

func (i *Image) SetQuantumPad(q *QuantumInfo, pad uint32) bool { return SetQuantumPad(i, q, pad) }

var SetQuantumQuantum func(*QuantumInfo, uint32)
var SetQuantumScale func(*QuantumInfo, float64)
var SetRandomKey func(*RandomInfo, uint32, *byte)
var SetRandomSecretKey func(Size)
var SetRandomTrueRandom func(bool)
var SetResampleFilter func(*ResampleFilter, FilterTypes, float64)
var SetSignatureDigest func(*SignatureInfo, *StringInfo)
var SimilarityMetricImage func(*Image, *Image, MetricType, *RectangleInfo, *float64, *ExceptionInfo) *Image
var SolarizeImageChannel func(*Image, ChannelType, float64, *ExceptionInfo) bool

func (i *Image) SimilarityMetricImage(ref *Image, m MetricType, r *RectangleInfo, similarity *float64, e *ExceptionInfo) *Image {
	return SimilarityMetricImage(i, ref, m, r, similarity, e)
}
func (i *Image) SolarizeImageChannel(c ChannelType, threshold float64, e *ExceptionInfo) bool {
	return SolarizeImageChannel(i, c, threshold, e)
}

var SRGBCompandor func(MagickRealType) MagickRealType
var StringInfoToHexString func(*StringInfo) string
var StringToArrayOfDoubles func(string, *int32, *ExceptionInfo) []float64
var SyncAuthenticPixelCacheNexus func(i *Image, n *NexusInfo, e *ExceptionInfo) bool
var TransparentPaintImageChroma func(i *Image, low, high *MagickPixelPacket, opacity Quantum, invert bool) bool

func (i *Image) SyncAuthenticPixelCacheNexus(n *NexusInfo, e *ExceptionInfo) bool {
	return SyncAuthenticPixelCacheNexus(i, n, e)
}
func (i *Image) TransparentPaintImageChroma(low, high *MagickPixelPacket, opacity Quantum, invert bool) bool {
	return TransparentPaintImageChroma(i, low, high, opacity, invert)
}

var TypeComponentGenesis func() bool
var TypeComponentTerminus func()
var UnityAddKernelInfo func(*KernelInfo, float64)
var WhiteThresholdImageChannel func(i *Image, c ChannelType, threshold string, e *ExceptionInfo) bool
var WriteBlobMSBLongLong func(i *Image, value MagickSizeType) int32

func (i *Image) WhiteThresholdImageChannel(c ChannelType, threshold string, e *ExceptionInfo) bool {
	return WhiteThresholdImageChannel(i, c, threshold, e)
}
func (i *Image) WriteBlobMSBLongLong(value MagickSizeType) int32 {
	return WriteBlobMSBLongLong(i, value)
}

var XComponentGenesis func() bool
var XComponentTerminus func()
var AcquireMagickMatrix func(nptrs, size Size) **float64
var AcquireMagickMemory func(size uint32) *Void

// Deprecated
var AcquireMemory func(size uint32) *Void
var AcquireQuantumMemory func(count, quantum uint32) *Void
var AcquireString func(source string) string
var AcquireStringInfo func(length uint32) *StringInfo
var AcquireUniqueFilename func(path string) bool
var AcquireUniqueFileResource func(path string) int
var AcquireUniqueSymbolicLink func(source, destination string) bool

// Deprecated
var AllocateString func(source string) string
var AppendImageFormat func(format, filename string)
var Base64Decode func(source string, length *uint32) *byte
var Base64Encode func(blob *byte, blobLength uint32, encodeLength *uint32) string
var BlobToFile func(filename string, blob *Void, length uint32, exception *ExceptionInfo) bool
var CanonicalXMLContent func(content string, pedantic bool) string
var ChopPathComponents func(path string, components Size)

// Deprecated
var CloneMemory func(destination, source *Void, size uint32) *Void
var CloneString func(destination []string, source string) string
var CloseMagickLog func()
var CompareHashmapString func(target, source *Void) bool
var CompareHashmapStringInfo func(target, source *Void) bool
var CompareSplayTreeString func(target, source *Void) int
var CompareSplayTreeStringInfo func(target, source *Void) int
var ConcatenateMagickString func(destination, source string, length uint32) uint32
var ConcatenateString func(destination []string, source string) bool
var ConfigureFileToStringInfo func(filename string) *StringInfo
var ConstantString func(source string) string
var ConstituteImage func(columns, rows Size, map_ string, storage StorageType, pixels *Void, exception *ExceptionInfo) *Image
var ConvertHSBToRGB func(hue, saturation, brightness float64, red, green, blue *Quantum)
var ConvertHSLToRGB func(hue, saturation, luminosity float64, red, green, blue *Quantum)
var ConvertHWBToRGB func(hue, whiteness, blackness float64, red, green, blue *Quantum)
var ConvertRGBToHSB func(red, green, blue Quantum, hue, saturation, brightness *float64)
var ConvertRGBToHSL func(red, green, blue Quantum, hue, saturation, luminosity *float64)
var ConvertRGBToHWB func(red, green, blue Quantum, hue, whiteness, blackness *float64)
var CopyMagickMemory func(destination, source *Void, size uint32) *Void
var CopyMagickString func(destination, source string, length uint32) uint32
var DefineImageRegistry func(type_ RegistryType, option string, exception *ExceptionInfo) bool
var DeleteImageRegistry func(key string) bool

// Deprecated
var DeleteMagickRegistry func(id SSize) bool
var DestroyConstitute func()

// Deprecated
var DestroyMagick func()
var DestroyMagickMemory func()
var DestroyMagickRegistry func()
var DestroyModuleList func()
var DestroyMontageInfo func(montageInfo *MontageInfo) *MontageInfo
var DestroyQuantumInfo func(quantumInfo *QuantumInfo) *QuantumInfo
var DestroyString func(str string) string
var DestroyStringList func(list []string) []string
var DestroyThresholdMap func(map_ *ThresholdMap) *ThresholdMap
var DestroyXResources func()
var DestroyXWidget func()
var EscapeString func(source string, escape int8) string
var ExpandAffine func(affine *AffineMatrix) float64
var ExpandFilename func(path string)
var ExpandFilenames func(argc *int, argv *[]string) bool
var FileToBlob func(filename string, extent uint32, length *uint32, exception *ExceptionInfo) *byte
var FileToString func(filename string, extent uint32, exception *ExceptionInfo) string
var FileToStringInfo func(filename string, extent uint32, exception *ExceptionInfo) *StringInfo
var FormatMagickSize func(size MagickSizeType, format string) SSize

// Deprecated
var FormatMagickString func(str string, length uint32, format string, va ...VArg) SSize
var FormatMagickStringList func(str string, length uint32, format string, operands VAList) SSize
var FormatMagickTime func(time Time, length uint32, timestamp string) SSize

// Deprecated
var FormatString func(str, format string, va ...VArg)
var FormatStringList func(str, format string, operands VAList)
var FuzzyColorMatch func(p, q *PixelPacket, fuzz float64) uint
var GaussJordanElimination func(matrix, vectors **float64, rank, nvecs Size) bool
var GetAffineMatrix func(affineMatrix *AffineMatrix)
var GetClientName func() string
var GetClientPath func() string
var GetCoderInfo func(name string, exception *ExceptionInfo) *CoderInfo
var GetCoderInfoList func(pattern string, numberCoders *Size, exception *ExceptionInfo) **CoderInfo
var GetCoderList func(pattern string, numberCoders *Size, exception *ExceptionInfo) []string
var GetColorInfo func(name string, exception *ExceptionInfo) *ColorInfo
var GetColorInfoList func(pattern string, numberColors *Size, exception *ExceptionInfo) **ColorInfo
var GetColorList func(pattern string, numberColors *Size, exception *ExceptionInfo) []string

// Deprecated
var GetConfigureBlob func(filename, path string, length *uint32, exception *ExceptionInfo) *Void
var GetConfigureInfo func(name string, exception *ExceptionInfo) *ConfigureInfo
var GetConfigureInfoList func(pattern string, numberOptions *Size, exception *ExceptionInfo) **ConfigureInfo
var GetConfigureList func(pattern string, numberOptions *Size, exception *ExceptionInfo) []string
var GetConfigureOptions func(filename string, exception *ExceptionInfo) *LinkedListInfo
var GetConfigurePaths func(filename string, exception *ExceptionInfo) *LinkedListInfo
var GetConfigureValue func(configureInfo *ConfigureInfo) string
var GetDelegateList func(pattern string, numberDelegates *Size, exception *ExceptionInfo) []string
var GetEnvironmentValue func(name string) string
var GetExceptionMessage func(errorCode int) string
var GetExecutionPath func(path string, extent uint32) bool
var GetGeometry func(geometry string, x, y *Long, width, height *Size) MagickStatusType

// Deprecated
var GetImageFromMagickRegistry func(name string, id *Long, exception *ExceptionInfo) *Image
var GetImageMagick func(magick *byte, length uint32, format string) string
var GetImageRegistry func(type_ RegistryType, key string, exception *ExceptionInfo) *Void
var GetLocaleExceptionMessage func(severity ExceptionType, tag string) string
var GetLocaleInfo func(tag string, exception *ExceptionInfo) *LocaleInfo
var GetLocaleInfoList func(pattern string, numberMessages *Size, exception *ExceptionInfo) []*LocaleInfo
var GetLocaleList func(pattern string, numberMessages *Size, exception *ExceptionInfo) []string
var GetLocaleMessage func(tag string) string
var GetLocaleOptions func(filename string, exception *ExceptionInfo) *LinkedListInfo
var GetLocaleValue func(localeInfo *LocaleInfo) string
var GetLogInfoList func(pattern string, numberPreferences *Size, exception *ExceptionInfo) []*LogInfo
var GetLogList func(pattern string, numberPreferences *Size, exception *ExceptionInfo) []string
var GetLogName func() string
var GetMagicInfo func(magic *byte, length uint32, exception *ExceptionInfo) *MagicInfo
var GetMagicInfoList func(pattern string, numberAliases *Size, exception *ExceptionInfo) []*MagicInfo
var GetMagickCopyright func() string
var GetMagickGeometry func(geometry string, x, y *Long, width, height *Size) uint
var GetMagickHomeURL func() string
var GetMagickInfo func(name string, exception *ExceptionInfo) *MagickInfo
var GetMagickInfoList func(pattern string, numberFormats *Size, exception *ExceptionInfo) []*MagickInfo
var GetMagickList func(pattern string, numberFormats *Size, exception *ExceptionInfo) []string
var GetMagickPackageName func() string
var GetMagickQuantumDepth func(depth *Size) string
var GetMagickQuantumRange func(range_ *Size) string

// Deprecated
var GetMagickRegistry func(id SSize, type_ *RegistryType, length *uint32, exception *ExceptionInfo) *Void
var GetMagickReleaseDate func() string
var GetMagickToken func(start string, end []string, token string)
var GetMagickPrecision func() int
var GetMagickVersion func(version *Size) string
var GetMagicList func(pattern string, numberAliases *Size, exception *ExceptionInfo) []string
var GetMagicName func(magicInfo *MagicInfo) string
var GetMimeDescription func(m *MimeInfo) string

func (m *MimeInfo) Description() string { return GetMimeDescription(m) }

var GetMimeInfo func(filename string, magic *byte, length uint32, exception *ExceptionInfo) *MimeInfo
var GetMimeInfoList func(pattern string, numberAliases *Size, exception *ExceptionInfo) []*MimeInfo
var GetMimeList func(pattern string, numberAliases *Size, exception *ExceptionInfo) []string
var GetMimeType func(m *MimeInfo) string

func (m *MimeInfo) Type() string { return GetMimeType(m) }

var GetModuleInfo func(tag string, exception *ExceptionInfo) *ModuleInfo
var GetModuleInfoList func(pattern string, numberModules *Size, exception *ExceptionInfo) **ModuleInfo
var GetModuleList func(pattern string, numberModules *Size, exception *ExceptionInfo) []string
var GetMonitorHandler func() MonitorHandler
var GetNextImageRegistry func() string
var GetOptimalKernelWidth func(radius, sigma float64) Size
var GetOptimalKernelWidth1D func(radius, sigma float64) Size
var GetOptimalKernelWidth2D func(radius, sigma float64) Size
var GetPageGeometry func(pageGeometry string) string
var GetPathComponent func(path string, type_ PathType, component string)
var GetPathComponents func(path string, numberComponents *Size) []string
var GetRandomKey func(key *byte, length uint32)
var GetRandomValue func() float64
var GetThresholdMap func(mapId string, exception *ExceptionInfo) *ThresholdMap
var GetThresholdMapFile func(xml, filename, mapId string, exception *ExceptionInfo) *ThresholdMap
var GetTypeInfo func(name string, exception *ExceptionInfo) *TypeInfo
var GetTypeInfoByFamily func(family string, style StyleType, stretch StretchType, weight Size, exception *ExceptionInfo) *TypeInfo
var GetTypeInfoList func(pattern string, numberFonts *Size, exception *ExceptionInfo) **TypeInfo
var GetTypeList func(pattern string, numberFonts *Size, exception *ExceptionInfo) []string
var GlobExpression func(expression, pattern string, caseInsensitive bool) bool
var GravityAdjustGeometry func(width, height Size, gravity GravityType, region *RectangleInfo)
var HashPointerType func(pointer *Void) uint32
var HashStringInfoType func(string *Void) uint32
var HashStringType func(string *Void) uint32

// Deprecated
var IdentityAffine func(affine *AffineMatrix)

// Deprecated
var InitializeMagick func(path string)
var InterpretImageFilename func(str string, length uint32, format string, value int) SSize
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
var LeastSquaresAddTerms func(matrix, vectors **float64, terms, results *float64, rank, nvecs Size)

// Deprecated
var LiberateMemory func(memory **Void)
var ListFiles func(directory, pattern string, numberEntries *Size) []string
var LoadMimeLists func(filename string, exception *ExceptionInfo) bool
var LocaleCompare func(p, q string) SSize
var LocaleLower func(str string)
var LocaleNCompare func(p, q string, length uint32) SSize
var LocaleUpper func(str string)
var LogMagickEvent func(type_ LogEventType, module, function string, line Size, format string, va ...VArg) bool
var LogMagickEventList func(type_ LogEventType, module, function string, line Size, format string, operands VAList) bool
var MagickCoreGenesis func(path string, establishSignalHandlers bool)
var MagickCoreTerminus func()
var MagickError func(err ExceptionType, reason, description string)
var MagickFatalError func(err ExceptionType, reason, description string)

// Deprecated
var MagickIncarnate func(path string)

// Deprecated
var MagickMonitor func(text string, offset MagickOffsetType, span MagickSizeType, clientData *Void) bool
var MagickToMime func(magick string) string
var MagickWarning func(warning ExceptionType, reason, description string)
var MapBlob func(file int, mode MapMode, offset MagickOffsetType, length uint32) *byte

// Deprecated
var MapImages func(images, mapImage *Image, dither bool) bool
var ModifyImage func(image **Image, exception *ExceptionInfo) bool
var MSBOrderLong func(buffer *byte, length uint32)
var MSBOrderShort func(p *byte, length uint32)
var MultilineCensus func(label string) Size
var NewHashmap func(capacity Size, hash func(*Void) uint32, compare func(*Void, *Void) *bool, relinquishKey, relinquishValue func(*Void) *Void) *HashmapInfo
var NewImageList func() *Image
var NewXMLTree func(xml string, exception *ExceptionInfo) *XMLTreeInfo
var NewXMLTreeTag func(tag string) *XMLTreeInfo
var OpenModule func(module string, exception *ExceptionInfo) bool
var OpenModules func(exception *ExceptionInfo) bool
var ParseAbsoluteGeometry func(geometry string, regionInfo *RectangleInfo) MagickStatusType
var ParseAffineGeometry func(geometry string, affineMatrix *AffineMatrix) MagickStatusType
var ParseChannelOption func(channels string) SSize
var ParseGeometry func(geometry string, geometryInfo *GeometryInfo) MagickStatusType
var ParseImageGeometry func(geometry string, x, y *Long, width, height *Size) int
var ParseMetaGeometry func(geometry string, x, y *Long, width, height *Size) MagickStatusType

// Deprecated
var PostscriptGeometry func(page string) string
var PrintStringInfo func(file *FILE, id string, stringInfo *StringInfo)

func (f *FILE) PrintStringInfo(id string, stringInfo *StringInfo) {
	PrintStringInfo(f, id, stringInfo)
}

var QueryColorDatabase func(name string, color *PixelPacket, exception *ExceptionInfo) bool
var QueryMagickColor func(name string, color *MagickPixelPacket, exception *ExceptionInfo) bool
var ReacquireMemory func(memory **Void, size uint32)
var RegisterStaticModules func()
var RelinquishMagickMatrix func(matrix **float64, nptrs Size) **float64
var RelinquishMagickMemory func(memory *Void) *Void
var RelinquishUniqueFileResource func(path string) bool
var RemoveImageRegistry func(key string) *Void
var ResetImageRegistryIterator func()
var ResetMagickMemory func(memory *Void, byte int, size uint32) *Void
var ResizeMagickMemory func(memory *Void, size uint32) *Void
var ResizeQuantumMemory func(memory *Void, count, quantum uint32) *Void
var SetCacheThreshold func(size Size)
var SetClientName func(name string) string
var SetClientPath func(path string) string
var SetErrorHandler func(handler ErrorHandler) ErrorHandler
var SetFatalErrorHandler func(handler FatalErrorHandler) FatalErrorHandler
var SetGeometryInfo func(geometryInfo *GeometryInfo)
var SetImageRegistry func(type_ RegistryType, key string, value *Void, exception *ExceptionInfo) bool
var SetLogEventMask func(events string) LogEventType
var SetLogFormat func(format string)
var SetLogName func(name string) string
var SetMagickInfo func(name string) *MagickInfo

// Deprecated
var SetMagickRegistry func(type_ RegistryType, blob *Void, length uint32, exception *ExceptionInfo) SSize

// Deprecated
var SetMonitorHandler func(handler MonitorHandler) MonitorHandler
var SetWarningHandler func(handler WarningHandler) WarningHandler
var StringToArgv func(text string, argc *int) []string
var StringToDouble func(str string, interval float64) float64
var StringToken func(delimiters string, string []string) string
var StringToList func(text string) []string
var StringToStringInfo func(str string) *StringInfo

// Deprecated
var Strip func(message string)
var StripString func(message string)
var SubstituteString func(buffer []string, search, replace string) bool
var SystemCommand func(verbose bool, command string) int

// Deprecated
var TemporaryFilename func(path string)
var TransformImage func(image **Image, cropGeometry, imageGeometry string) bool
var UnmapBlob func(map_ *Void, length uint32) bool
var UnregisterMagickInfo func(name string) bool
var UnregisterStaticModules func()
var XDestroyResourceInfo func(resourceInfo *XResourceInfo)
var XGetAnnotateInfo func(annotateInfo *XAnnotateInfo)
var XGetImportInfo func(ximageInfo *XImportInfo)
var XGetMapInfo func(visualInfo *XVisualInfo, colormap Colormap, mapInfo *XStandardColormap)
var XInitImage func(ximage *XImage) Status
var XMagickProgressMonitor func(tag string, quantum MagickOffsetType, span MagickSizeType, clientData *Void) bool
var XQueryColorDatabase func(target string, color *XColor) bool
var XrmCombineFileDatabase func(filename string, target **XrmDatabase, override bool) Status
var XSetLocaleModifiers func(modifiers string) string
var XSetWindows func(windowsInfo *XWindows) *XWindows
var XSupportsLocale func() bool
var XUserPreferences func(resourceInfo *XResourceInfo)
var XWarning func(warning ExceptionType, reason, description string)
var AcquireAlignedMemory func(count uint32, quantum uint32) *Void
var AcquireKernelBuiltIn func(type_ KernelInfoType, args GeometryInfo) *KernelInfo
var AcquireKernelInfo func(kernelString string) *KernelInfo
var AcquireModuleInfo func(path, tag string) *ModuleInfo

// Deprecated
var AcquireOneCacheViewVirtualPixel func(cacheView *CacheView, virtualPixelMethod VirtualPixelMethod, x, y int32, pixel *PixelPacket, exception *ExceptionInfo) bool
var AcquirePixelCacheNexus func(numberThreads uint32) **NexusInfo
var CloneKernelInfo func(kernel *KernelInfo) *KernelInfo
var DestroyKernelInfo func(kernel *KernelInfo) *KernelInfo
var DestroyPixelCache func()
var DuplicateImages func(images *Image, numberDuplicates uint32, scenes string, exception *ExceptionInfo) *Image
var EvaluateImages func(images *Image, op MagickEvaluateOperator, value float64, exception *ExceptionInfo) bool
var GetImageViewException func(pixelImage *ImageView, severity *ExceptionType) string
var GetMagickFeatures func() string
var GetMagickMemoryMethods func(a *AcquireMemoryHandler, r *ResizeMemoryHandler, destroyMemoryHandler *DestroyMemoryHandler)

// Deprecated
var HSLTransform func(hue, saturation, lightness float64, red, green, blue *Quantum)
var InitializeModuleList func(exception *ExceptionInfo)

// Deprecated
var MaximumImages func(images *Image, exception *ExceptionInfo) *Image

// Deprecated
var MinimumImages func(images *Image, exception *ExceptionInfo) *Image
var NewImageView func(i *Image) *ImageView

func (i *Image) NewView() *ImageView { return NewImageView(i) }

var NewImageViewRegion func(i *Image, x, y int32, width, height uint32) *ImageView

// Deprecated
var OpenMagickStream func(path, mode string) *FILE
var QueryColorCompliance func(name string, compliance ComplianceType, color *PixelPacket, exception *ExceptionInfo) bool
var RegistryComponentTerminus func()
var RelinquishAlignedMemory func(memory *Void) *Void
var ReplaceImageInListReturnLast func(images **Image, replace *Image)
var ScaleGeometryKernelInfo func(kernel *KernelInfo, scalingFactor float64, normalizeFlags MagickStatusType)
var ScaleKernelInfo func(kernel *KernelInfo, scalingFactor float64, normalizeFlags MagickStatusType)
var SetMagickMemoryMethods func(a AcquireMemoryHandler, r ResizeMemoryHandler, d DestroyMemoryHandler)
var SetMagickPrecision func(precision int) int
var ShowKernelInfo func(kernel *KernelInfo)
var SmushImages func(images *Image, stack bool, exception *ExceptionInfo) *Image

// Deprecated
var TransformHSL func(red, green, blue Quantum, hue, saturation, lightness *float64)
var ZeroKernelNans func(kernel *KernelInfo)
var Exit func(int) int
var IsWindows95 func() bool
var NTCloseDirectory func(*DIR) int
var NTCloseLibrary func(*Void) int
var NTControlHandler func() int
var NTElapsedTime func() float64
var NTErrorHandler func(severity ExceptionType, reason string, description string)
var NTExitLibrary func() int
var NTGetExecutionPath func(string, uint32) bool
var NTGetLastError func() string
var NTGetLibraryError func() string
var NTGetLibrarySymbol func(handle *Void, name string) *Void
var NTGetModulePath func(string, string) bool
var NTGhostscriptDLL func(string, int) int
var NTGhostscriptDLLVectors func() *GhostInfo
var NTGhostscriptEXE func(string, int) int
var NTGhostscriptFonts func(string, int) int
var NTGhostscriptLoadDLL func() int
var NTGhostscriptUnLoadDLL func() int
var NTInitializeLibrary func() int
var NTIsMagickConflict func(string) bool
var NTLoadTypeLists func(*SplayTreeInfo, *ExceptionInfo) bool
var NTMapMemory func(address string, length uint32, protection int, access int, file int, offset MagickOffsetType) *Void
var NTOpenDirectory func(string) *DIR
var NTOpenLibrary func(filename string) *Void
var NTReadDirectory func(*DIR) *Dirent
var NTRegistryKeyLookup func(string) string
var NTReportEvent func(string, bool) bool
var NTResourceToBlob func(string) *byte
var NTSeekDirectory func(entry *DIR, position SSize)
var NTSetSearchPath func(string) int
var NTSyncMemory func(*Void, uint32, int) int
var NTSystemCommand func(string) int
var NTSystemConfiguration func(int) SSize
var NTTellDirectory func(*DIR) SSize
var NTTruncateFile func(int, Off) int
var NTUnmapMemory func(*Void, uint32) int
var NTUserTime func() float64
var NTWarningHandler func(severity ExceptionType, reason string, description string)
var PlasmaImageProxy func(i *Image, imageView *CacheView, randomInfo *RandomInfo, segment *SegmentInfo, attenuate, depth uint32) bool

func (i *Image) PlasmaProxy(imageView *CacheView, randomInfo *RandomInfo, segment *SegmentInfo, attenuate, depth uint32) bool {
	return PlasmaImageProxy(i, imageView, randomInfo, segment, attenuate, depth)
}

var GetColorCompliance func(name string, compliance ComplianceType, exception *ExceptionInfo) *ColorInfo

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
	{"GetColorCompliance", &GetColorCompliance},
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

type CacheInfo struct {
	// StorageClass             ClassType
	// Colorspace               ColorspaceType
	// Channels                 Size
	// Type                     CacheType
	// Mode                     MapMode
	// Mapped                   MagickBooleanType
	// Columns, Rows            Size
	// Offset                   MagickOffsetType
	// Length                   MagickSizeType
	// VirtualPixelMethod       VirtualPixelMethod
	// VirtualPixelColor        MagickPixelPacket
	// NumberThreads            Size
	// NexusInfo                **NexusInfo
	// Pixels                   *PixelPacket
	// Indexes                  *IndexPacket
	// ActiveIndexChannel       MagickBooleanType
	// File                     int
	// Filename                 [MaxTextExtent]Char
	// CacheFilename            [MaxTextExtent]Char
	// Methods                  CacheMethods
	// RandomInfo               *RandomInfo
	// NumberConnections        Size
	// ServerInfo               *Void
	// Synchronize, Debug       MagickBooleanType
	// Id                       MagickThreadType
	// ReferenceCount           SSize
	// Semaphore, FileSemaphore *SemaphoreInfo
	// Timestamp                Time
	// Signature                Size
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
	// Name             *PVString
	// Geometry         *PVString
	// IconName         *PVString
	// IconGeometry     *PVString
	// CropGeometry     *PVString
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
	// Mask             SSize
	// Orphan           MagickBooleanType
	// Mapped           MagickBooleanType
	// Stasis           MagickBooleanType
	// Image            *Image
	// Destroy          MagickBooleanType
}
type XResourceInfo struct {
	ResourceDatabase *XrmDatabase
	ImageInfo        *ImageInfo
	QuantizeInfo     *QuantizeInfo
	Colors           Size
	CloseServer      MagickBooleanType
	Backdrop         MagickBooleanType
	BackgroundColor  PVString
	BorderColor      PVString
	ClientName       PVString
	Colormap         XColormapType
	BorderWidth      uint
	Delay            Size
	ColorRecovery    MagickBooleanType
	ConfirmExit      MagickBooleanType
	ConfirmEdit      MagickBooleanType
	DisplayGamma     PVString
	Font             PVString
	FontName         [MaxNumberFonts]*Char //TODO(t): handle
	ForegroundColor  PVString
	DisplayWarnings  MagickBooleanType
	GammaCorrect     MagickBooleanType
	IconGeometry     PVString
	Iconic           MagickBooleanType
	Immutable        MagickBooleanType
	ImageGeometry    PVString
	MapType          PVString
	MatteColor       PVString
	Name             PVString
	Magnify          uint
	Pause            uint
	PenColors        [MaxNumberPens]*Char //TODO(t): handle
	TextFont         PVString
	Title            PVString
	Quantum          int
	Update           uint
	UsePixmap        MagickBooleanType
	UseSharedMemory  MagickBooleanType
	UndoCache        Size
	VisualType       PVString
	WindowGroup      PVString
	WindowId         PVString
	WriteFilename    PVString
	CopyImage        *Image
	Gravity          int
	HomeDirectory    [MaxTextExtent]Char
}
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
	DefaultChannels               = ((AllChannels | SyncChannels) & ^OpacityChannel)
	UndefinedChannel  ChannelType = 0
	GrayChannel                   = RedChannel
	CyanChannel                   = RedChannel
	MagentaChannel                = GreenChannel
	YellowChannel                 = BlueChannel
	OpacityChannel                = AlphaChannel
	matteChannel                  = AlphaChannel // Deprecated
	IndexChannel                  = BlackChannel
	GrayChannels                  = RGBChannels
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
	Path      PVString
	Magick    PVString
	Name      PVString
	Exempt    MagickBooleanType
	Stealth   MagickBooleanType
	Previous  *CoderInfo
	next      *CoderInfo // Deprecated
	Signature Size
}
type ColorInfo struct {
	Path       PVString
	Name       PVString
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
	Path      PVString
	Name      PVString
	Value     PVString
	Exempt    MagickBooleanType
	Stealth   MagickBooleanType
	Previous  *ConfigureInfo
	next      *ConfigureInfo // Deprecated
	Signature Size
}
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
	Primitive        PVString
	Geometry         PVString
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
	Text             PVString
	Face             Size
	Font             PVString
	Metrics          PVString
	Family           PVString
	Style            StyleType
	Stretch          StretchType
	Weight           Size
	Encoding         PVString
	Pointsize        float64
	Density          PVString
	Align            AlignType
	Undercolor       PixelPacket
	BorderColor      PixelPacket
	ServerName       PVString
	DashPattern      *float64
	ClipMask         PVString
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
type ElementReference struct {
	Id        PVString
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
	Reason      POVString
	Description POVString
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
type GhostInfo struct {
	exit           func(*struct{}) int
	initWithArgs   func(*struct{}, int, **Char) int
	newInstance    func(**struct{}, *Void) int
	runString      func(*struct{}, string, int, *int) int
	deleteInstance func(*struct{})
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
	Montage                *CString // NOTE(t): If PVString crash on i.Destroy()
	Directory              *CString
	Geometry_              *CString
	Offset                 SSize
	_                      int32 // padding
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
	_                      int32 // padding
	Timer                  TimerInfo
	ProgressMonitor        *MagickProgressMonitor //TODO(T): PURE FUNC CAUSES DESTRUCT TO CRASH
	ClientData             *Void
	Cache                  *Void
	Attributes             *Void
	Ascii85                *Ascii85Info
	_                      int32 // padding
	Filename               [MaxTextExtent]Char
	MagickFilename         [MaxTextExtent]Char
	Magick                 [MaxTextExtent]Char
	MagickColumns          Size
	MagickRows             Size
	Exception_             ExceptionInfo
	Debug                  MagickBooleanType
	ReferenceCount_        SSize
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
	Intensity              PixelIntensityMethod // Weird content
}
type ImageAttribute struct {
	Key         PVString
	Value       PVString
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
	Size               PVString
	Extract            PVString
	Page               PVString
	Scenes             PVString
	Scene              Size
	NumberScenes       Size
	Depth              Size
	Interlace          InterlaceType
	Endian             EndianType
	Units              ResolutionType
	Quality            Size
	SamplingFactor     PVString
	ServerName         PVString
	Font               PVString
	Texture            PVString
	Density            PVString
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
	View               PVString
	Authenticate       PVString
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
	Tile               PVString
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
	Path      PVString
	Tag       PVString
	Message   PVString
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

// Changed to bool for uses other than in structs
type MagickBooleanType Enum

const (
	MagickFalse MagickBooleanType = iota
	MagickTrue
)

type MagickFunction Enum

const (
	UndefinedFunction MagickFunction = iota
	PolynomialFunction
	SinusoidFunction
	ArcsinFunction
	ArctanFunction
)

type MagicInfo struct {
	Path      PVString
	Name      PVString
	Target    PVString
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
	Path             PVString
	Tag              PVString
	Handle           *Void
	UnregisterModule func()
	RegisterModule   func() Size
	LoadTime         Time
	Stealth          MagickBooleanType
	Previous         *ModuleInfo
	next             *ModuleInfo // Deprecated
	Signature        Size
}
type MontageInfo struct {
	Geometry        PVString
	Tile            PVString
	Title           PVString
	Frame           PVString
	Texture         PVString
	Font            PVString
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
	Text        PVString
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
	Name      PVString
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
	X, Y          SSize
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

type TypeInfo struct {
	Face        Size
	Path        PVString
	Name        PVString
	Description PVString
	Family      PVString
	Style       StyleType
	Stretch     StretchType
	Weight      Size
	Encoding    PVString
	Foundry     PVString
	Format      PVString
	Metrics     PVString
	Glyphs      PVString
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

type XAnnotateInfo struct {
	X        int
	Y        int
	Width    uint
	Height   uint
	Degrees  float64
	FontInfo *XFontStruct
	Text     PVString
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
type DelegateInfo struct {
	Path           PVString
	Decode         PVString
	Encode         PVString
	Commands_      PVString
	Mode_          SSize
	ThreadSupport_ MagickBooleanType
	Spawn          MagickBooleanType
	Stealth        MagickBooleanType
	Previous       *DelegateInfo
	next           *DelegateInfo // Deprecated
	Signature      Size
}
type MagickInfo struct {
	Name            PVString
	Description_    PVString
	Version         PVString
	Note            PVString
	Module          PVString
	ImageInfo       *ImageInfo
	Decoder_        *DecodeImageHandler
	Encoder_        *EncodeImageHandler
	Magick          *IsImageFormatHandler
	ClientData      *Void
	Adjoin_         MagickBooleanType
	Raw             MagickBooleanType
	EndianSupport_  MagickBooleanType
	BlobSupport_    MagickBooleanType
	SeekableStream_ MagickBooleanType
	FormatType      MagickFormatType
	ThreadSupport_  MagickStatusType
	Stealth         MagickBooleanType
	Previous        *MagickInfo
	next            *MagickInfo // Deprecated
	Signature       Size
	MimeType        PVString
}
type MagickFormatType Enum

const (
	UndefinedFormatType MagickFormatType = iota
	ImplicitFormatType
	ExplicitFormatType
)

type PixelIntensityMethod Enum

const (
	UndefinedPixelIntensityMethod PixelIntensityMethod = iota
	AveragePixelIntensityMethod
	BrightnessPixelIntensityMethod
	LightnessPixelIntensityMethod
	Rec601LumaPixelIntensityMethod
	Rec601LuminancePixelIntensityMethod
	Rec709LumaPixelIntensityMethod
	Rec709LuminancePixelIntensityMethod
	RMSPixelIntensityMethod
	MSPixelIntensityMethod
)

type Timezone struct{ MinutesWest, DstTime int }

type Timeval struct{ Sec, USec Long }
