// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package wand provides API definitions for
//accessing CORE_RL_wand_.dll.
package wand

import (
	. "github.com/tHinqa/outside"
	I "github.com/tHinqa/outside-imagemagick"
	// . "github.com/tHinqa/outside/types"
)

// Opaque
type (
	CacheView   struct{}
	DrawingWand struct{}
	PixelView   struct{}
	Void        struct{}
)

type (
	DuplexTransferPixelViewMethod func(*PixelView, *PixelView, *PixelView, *Void) bool
	GetPixelViewMethod            func(*PixelView, *Void) bool
	GetWandViewMethod             func(*I.WandView, I.SSize, int, *Void) bool
	SetPixelViewMethod            func(*PixelView, *Void) bool
	SetWandViewMethod             func(*I.WandView, I.SSize, int, *Void) bool
	TransferPixelViewMethod       func(*PixelView, *PixelView, *Void) bool
	UpdatePixelViewMethod         func(*PixelView, *Void) bool
)

type (
	DrawInfo I.DrawInfo

	ExceptionInfo     I.ExceptionInfo
	FILE              I.FILE
	Image             I.Image
	ImageInfo         I.ImageInfo
	ImageView         struct{}
	MagickPixelPacket I.MagickPixelPacket
	MagickWand        struct{}
	PixelIterator     I.Fixme
	PixelWand         struct{}
	WandView          struct{}
)

var NewMagickWand func() *MagickWand

var NewPixelIterator func(m *MagickWand) *PixelIterator

var NewPixelRegionIterator func(m *MagickWand, x, y int32, width, height uint32) *PixelIterator

var NewWandView func(m *MagickWand) *WandView

var NewWandViewExtent func(m *MagickWand, x, y int32, width, height uint32) *WandView

var ClearMagickWand func(m *MagickWand)

func (m *MagickWand) Clear() { ClearMagickWand(m) }

var CloneMagickWand func(m *MagickWand) *MagickWand

func (m *MagickWand) Clone() *MagickWand { return CloneMagickWand(m) }

var DestroyMagickWand func(m *MagickWand)

func (m *MagickWand) Destroy() (_ *MagickWand) { DestroyMagickWand(m); return }

var GetImageFromMagickWand func(m *MagickWand) *Image

func (m *MagickWand) ImageFrom() *Image { return GetImageFromMagickWand(m) }

var IsMagickWand func(m *MagickWand) bool

func (m *MagickWand) IsMagickWand() bool { return IsMagickWand(m) }

var AdaptiveBlurImage func(m *MagickWand, radius, sigma float64) bool

func (m *MagickWand) AdaptiveBlur(radius, sigma float64) bool {
	return AdaptiveBlurImage(m, radius, sigma)
}

var AdaptiveBlurImageChannel func(m *MagickWand, channel I.ChannelType, radius, sigma float64) bool

func (m *MagickWand) AdaptiveBlurChannel(channel I.ChannelType, radius, sigma float64) bool {
	return AdaptiveBlurImageChannel(m, channel, radius, sigma)
}

var AdaptiveResizeImage func(m *MagickWand, columns, rows uint32) bool

func (m *MagickWand) AdaptiveResize(columns, rows uint32) bool {
	return AdaptiveResizeImage(m, columns, rows)
}

var AdaptiveSharpenImage func(m *MagickWand, radius, sigma float64) bool

func (m *MagickWand) AdaptiveSharpen(radius, sigma float64) bool {
	return AdaptiveSharpenImage(m, radius, sigma)
}

var AdaptiveSharpenImageChannel func(m *MagickWand, channel I.ChannelType, radius, sigma float64) bool

func (m *MagickWand) AdaptiveSharpenChannel(channel I.ChannelType, radius, sigma float64) bool {
	return AdaptiveSharpenImageChannel(m, channel, radius, sigma)
}

var AdaptiveThresholdImage func(m *MagickWand, width, height uint32, offset int32) bool

func (m *MagickWand) AdaptiveThreshold(width, height uint32, offset int32) bool {
	return AdaptiveThresholdImage(m, width, height, offset)
}

var AddImage func(m *MagickWand, addWand *MagickWand) bool

func (m *MagickWand) Add(addWand *MagickWand) bool { return AddImage(m, addWand) }

var AddNoiseImage func(m *MagickWand, noiseType I.NoiseType) bool

func (m *MagickWand) AddNoise(noiseType I.NoiseType) bool {
	return AddNoiseImage(m, noiseType)
}

var AddNoiseImageChannel func(m *MagickWand, channel I.ChannelType, noiseType I.NoiseType) bool

func (m *MagickWand) AddNoiseChannel(channel I.ChannelType, noiseType I.NoiseType) bool {
	return AddNoiseImageChannel(m, channel, noiseType)
}

var AffineTransformImage func(m *MagickWand, d *DrawingWand) bool

func (m *MagickWand) AffineTransform(d *DrawingWand) bool {
	return AffineTransformImage(m, d)
}

var AnimateImages func(m *MagickWand, xServerName string) bool

func (m *MagickWand) Animate(xServerName string) bool {
	return AnimateImages(m, xServerName)
}

var AnnotateImage func(m *MagickWand, d *DrawingWand, x, y, angle float64, text string) bool

func (m *MagickWand) Annotate(d *DrawingWand, x, y, angle float64, text string) bool {
	return AnnotateImage(m, d, x, y, angle, text)
}

var AppendImages func(m *MagickWand, stack bool) *MagickWand

func (m *MagickWand) Append(stack bool) *MagickWand { return AppendImages(m, stack) }

var AutoGammaImage func(m *MagickWand) bool

func (m *MagickWand) AutoGamma() bool { return AutoGammaImage(m) }

var AutoGammaImageChannel func(m *MagickWand, channel I.ChannelType) bool

func (m *MagickWand) AutoGammaChannel(channel I.ChannelType) bool {
	return AutoGammaImageChannel(m, channel)
}

var AutoLevelImage func(m *MagickWand) bool

func (m *MagickWand) AutoLevel() bool { return AutoLevelImage(m) }

var AutoLevelImageChannel func(m *MagickWand, channel I.ChannelType) bool

func (m *MagickWand) AutoLevelChannel(channel I.ChannelType) bool {
	return AutoLevelImageChannel(m, channel)
}

var BlackThresholdImage func(m *MagickWand, threshold *PixelWand) bool

func (m *MagickWand) BlackThreshold(threshold *PixelWand) bool {
	return BlackThresholdImage(m, threshold)
}

var BlueShiftImage func(m *MagickWand, factor float64) bool

func (m *MagickWand) BlueShift(factor float64) bool { return BlueShiftImage(m, factor) }

var BlurImage func(m *MagickWand, radius, sigma float64) bool

func (m *MagickWand) Blur(radius, sigma float64) bool { return BlurImage(m, radius, sigma) }

var BlurImageChannel func(m *MagickWand, channel I.ChannelType, radius, sigma float64) bool

func (m *MagickWand) BlurChannel(channel I.ChannelType, radius, sigma float64) bool {
	return BlurImageChannel(m, channel, radius, sigma)
}

var BorderImage func(m *MagickWand, bordercolor *PixelWand, width, height uint32) bool

func (m *MagickWand) Border(bordercolor *PixelWand, width, height uint32) bool {
	return BorderImage(m, bordercolor, width, height)
}

var BrightnessContrastImage func(m *MagickWand, brightness, contrast float64) bool

func (m *MagickWand) BrightnessContrast(brightness, contrast float64) bool {
	return BrightnessContrastImage(m, brightness, contrast)
}

var BrightnessContrastImageChannel func(m *MagickWand, channel I.ChannelType, brightness, contrast float64) bool

func (m *MagickWand) BrightnessContrastChannel(channel I.ChannelType, brightness, contrast float64) bool {
	return BrightnessContrastImageChannel(m, channel, brightness, contrast)
}

var CharcoalImage func(m *MagickWand, radius, sigma float64) bool

func (m *MagickWand) Charcoal(radius, sigma float64) bool {
	return CharcoalImage(m, radius, sigma)
}

var ChopImage func(m *MagickWand, width, height uint32, x, y int32) bool

func (m *MagickWand) Chop(width, height uint32, x, y int32) bool {
	return ChopImage(m, width, height, x, y)
}

var ClampImage func(m *MagickWand) bool

func (m *MagickWand) Clamp() bool { return ClampImage(m) }

var ClampImageChannel func(m *MagickWand, channel I.ChannelType) bool

func (m *MagickWand) ClampChannel(channel I.ChannelType) bool {
	return ClampImageChannel(m, channel)
}

var ClearException func(m *MagickWand) bool

func (m *MagickWand) ClearException() bool { return ClearException(m) }

var ClipImage func(m *MagickWand) bool

func (m *MagickWand) Clip() bool { return ClipImage(m) }

var ClipImagePath func(m *MagickWand, pathname string, inside bool) bool

func (m *MagickWand) ClipPath(pathname string, inside bool) bool {
	return ClipImagePath(m, pathname, inside)
}

var ClutImage func(m *MagickWand, clutWand *MagickWand) bool

func (m *MagickWand) Clut(clutWand *MagickWand) bool { return ClutImage(m, clutWand) }

var ClutImageChannel func(m *MagickWand, channel I.ChannelType, clutWand *MagickWand) bool

func (m *MagickWand) ClutChannel(channel I.ChannelType, clutWand *MagickWand) bool {
	return ClutImageChannel(m, channel, clutWand)
}

var CoalesceImages func(m *MagickWand) *MagickWand

func (m *MagickWand) Coalesce() *MagickWand { return CoalesceImages(m) }

var ColorDecisionListImage func(m *MagickWand, colorCorrectionCollection string) bool

func (m *MagickWand) ColorDecisionList(colorCorrectionCollection string) bool {
	return ColorDecisionListImage(m, colorCorrectionCollection)
}

var ColorizeImage func(m *MagickWand, colorize, opacity *PixelWand) bool

func (m *MagickWand) Colorize(colorize, opacity *PixelWand) bool {
	return ColorizeImage(m, colorize, opacity)
}

var ColorMatrixImage func(m *MagickWand, colorMatrix *I.KernelInfo) bool

func (m *MagickWand) ColorMatrix(colorMatrix *I.KernelInfo) bool {
	return ColorMatrixImage(m, colorMatrix)
}

var CombineImages func(m *MagickWand, channel I.ChannelType) *MagickWand

func (m *MagickWand) Combine(channel I.ChannelType) *MagickWand {
	return CombineImages(m, channel)
}

var CommentImage func(m *MagickWand, comment string) bool

func (m *MagickWand) Comment(comment string) bool { return CommentImage(m, comment) }

var CompareImageChannels func(m *MagickWand, reference *MagickWand, channel I.ChannelType, metric I.MetricType, distortion *float64) *MagickWand

func (m *MagickWand) CompareChannels(reference *MagickWand, channel I.ChannelType, metric I.MetricType, distortion *float64) *MagickWand {
	return CompareImageChannels(m, reference, channel, metric, distortion)
}

var CompareImageLayers func(m *MagickWand, method I.ImageLayerMethod) *MagickWand

func (m *MagickWand) CompareLayers(method I.ImageLayerMethod) *MagickWand {
	return CompareImageLayers(m, method)
}

var CompareImages func(m *MagickWand, reference *MagickWand, metric I.MetricType, distortion *float64) *MagickWand

func (m *MagickWand) Compare(reference *MagickWand, metric I.MetricType, distortion *float64) *MagickWand {
	return CompareImages(m, reference, metric, distortion)
}

var CompositeImage func(m *MagickWand, sourceWand *MagickWand, compose I.CompositeOperator, x, y int32) bool

func (m *MagickWand) Composite(sourceWand *MagickWand, compose I.CompositeOperator, x, y int32) bool {
	return CompositeImage(m, sourceWand, compose, x, y)
}

var CompositeImageChannel func(m *MagickWand, channel I.ChannelType, compositeWand *MagickWand, compose I.CompositeOperator, x, y int32) bool

func (m *MagickWand) CompositeChannel(channel I.ChannelType, compositeWand *MagickWand, compose I.CompositeOperator, x, y int32) bool {
	return CompositeImageChannel(m, channel, compositeWand, compose, x, y)
}

var CompositeLayers func(m *MagickWand, sourceWand *MagickWand, compose I.CompositeOperator, x, y int32) bool

func (m *MagickWand) CompositeLayers(sourceWand *MagickWand, compose I.CompositeOperator, x, y int32) bool {
	return CompositeLayers(m, sourceWand, compose, x, y)
}

var ConstituteImage func(m *MagickWand, columns, rows uint32, map_ string, storage I.StorageType, pixels *Void) bool

func (m *MagickWand) Constitute(columns, rows uint32, map_ string, storage I.StorageType, pixels *Void) bool {
	return ConstituteImage(m, columns, rows, map_, storage, pixels)
}

var ContrastImage func(m *MagickWand, sharpen bool) bool

func (m *MagickWand) Contrast(sharpen bool) bool { return ContrastImage(m, sharpen) }

var ContrastStretchImage func(m *MagickWand, blackPoint, whitePoint float64) bool

func (m *MagickWand) ContrastStretch(blackPoint, whitePoint float64) bool {
	return ContrastStretchImage(m, blackPoint, whitePoint)
}

var ContrastStretchImageChannel func(m *MagickWand, channel I.ChannelType, blackPoint, whitePoint float64) bool

func (m *MagickWand) ContrastStretchChannel(channel I.ChannelType, blackPoint, whitePoint float64) bool {
	return ContrastStretchImageChannel(m, channel, blackPoint, whitePoint)
}

var ConvolveImage func(m *MagickWand, order uint32, kernel []float64) bool

func (m *MagickWand) Convolve(order uint32, kernel []float64) bool {
	return ConvolveImage(m, order, kernel)
}

var ConvolveImageChannel func(m *MagickWand, channel I.ChannelType, order uint32, kernel *float64) bool

func (m *MagickWand) ConvolveChannel(channel I.ChannelType, order uint32, kernel *float64) bool {
	return ConvolveImageChannel(m, channel, order, kernel)
}

var CropImage func(m *MagickWand, width, height uint32, x, y int32) bool

func (m *MagickWand) Crop(width, height uint32, x, y int32) bool {
	return CropImage(m, width, height, x, y)
}

var CycleColormapImage func(m *MagickWand, displace int32) bool

func (m *MagickWand) CycleColormap(displace int32) bool {
	return CycleColormapImage(m, displace)
}

var DecipherImage func(m *MagickWand, passphrase string) bool

func (m *MagickWand) Decipher(passphrase string) bool { return DecipherImage(m, passphrase) }

var DeconstructImages func(m *MagickWand) *MagickWand

func (m *MagickWand) Deconstruct() *MagickWand { return DeconstructImages(m) }

var DeleteImageArtifact func(m *MagickWand, artifact string) bool

func (m *MagickWand) DeleteArtifact(artifact string) bool {
	return DeleteImageArtifact(m, artifact)
}

var DeleteImageProperty func(m *MagickWand, property string) bool

func (m *MagickWand) DeleteProperty(property string) bool {
	return DeleteImageProperty(m, property)
}

var DeleteOption func(m *MagickWand, option string) bool

func (m *MagickWand) DeleteOption(option string) bool { return DeleteOption(m, option) }

var DeskewImage func(m *MagickWand, threshold float64) bool

func (m *MagickWand) Deskew(threshold float64) bool { return DeskewImage(m, threshold) }

var DespeckleImage func(m *MagickWand) bool

func (m *MagickWand) Despeckle() bool { return DespeckleImage(m) }

var DisplayImage func(m *MagickWand, xServerName string) bool

func (m *MagickWand) Display(xServerName string) bool { return DisplayImage(m, xServerName) }

var DisplayImages func(m *MagickWand, xServerName string) bool

func (m *MagickWand) DisplayImages(xServerName string) bool {
	return DisplayImages(m, xServerName)
}

var DistortImage func(m *MagickWand, method I.DistortImageMethod, numberArguments uint32, arguments *float64, bestfit bool) bool

func (m *MagickWand) Distort(method I.DistortImageMethod, numberArguments uint32, arguments *float64, bestfit bool) bool {
	return DistortImage(m, method, numberArguments, arguments, bestfit)
}

var DrawImage func(m *MagickWand, d *DrawingWand) bool

func (m *MagickWand) Draw(d *DrawingWand) bool { return DrawImage(m, d) }

var EdgeImage func(m *MagickWand, radius float64) bool

func (m *MagickWand) Edge(radius float64) bool { return EdgeImage(m, radius) }

var EmbossImage func(m *MagickWand, radius, sigma float64) bool

func (m *MagickWand) Emboss(radius, sigma float64) bool {
	return EmbossImage(m, radius, sigma)
}

var EncipherImage func(m *MagickWand, passphrase string) bool

func (m *MagickWand) Encipher(passphrase string) bool { return EncipherImage(m, passphrase) }

var EnhanceImage func(m *MagickWand) bool

func (m *MagickWand) Enhance() bool { return EnhanceImage(m) }

var EqualizeImage func(m *MagickWand) bool

func (m *MagickWand) Equalize() bool { return EqualizeImage(m) }

var EqualizeImageChannel func(m *MagickWand, channel I.ChannelType) bool

func (m *MagickWand) EqualizeChannel(channel I.ChannelType) bool {
	return EqualizeImageChannel(m, channel)
}

var EvaluateImage func(m *MagickWand, operator I.MagickEvaluateOperator, value float64) bool

func (m *MagickWand) Evaluate(operator I.MagickEvaluateOperator, value float64) bool {
	return EvaluateImage(m, operator, value)
}

var EvaluateImageChannel func(m *MagickWand, channel I.ChannelType, op I.MagickEvaluateOperator, value float64) bool

func (m *MagickWand) EvaluateChannel(channel I.ChannelType, op I.MagickEvaluateOperator, value float64) bool {
	return EvaluateImageChannel(m, channel, op, value)
}

var EvaluateImages func(m *MagickWand, operator I.MagickEvaluateOperator) bool

func (m *MagickWand) EvaluateImages(operator I.MagickEvaluateOperator) bool {
	return EvaluateImages(m, operator)
}

var ExportImagePixels func(m *MagickWand, x, y int32, columns, rows uint32, map_ string, storage I.StorageType, pixels *Void) bool

func (m *MagickWand) ExportPixels(x, y int32, columns, rows uint32, map_ string, storage I.StorageType, pixels *Void) bool {
	return ExportImagePixels(m, x, y, columns, rows, map_, storage, pixels)
}

var ExtentImage func(m *MagickWand, width, height uint32, x, y int32) bool

func (m *MagickWand) Extent(width, height uint32, x, y int32) bool {
	return ExtentImage(m, width, height, x, y)
}

var FilterImage func(m *MagickWand, kernel *I.KernelInfo) bool

func (m *MagickWand) Filter(kernel *I.KernelInfo) bool { return FilterImage(m, kernel) }

var FilterImageChannel func(m *MagickWand, channel I.ChannelType, kernel *I.KernelInfo) bool

func (m *MagickWand) FilterChannel(channel I.ChannelType, kernel *I.KernelInfo) bool {
	return FilterImageChannel(m, channel, kernel)
}

var FlipImage func(m *MagickWand) bool

func (m *MagickWand) Flip() bool { return FlipImage(m) }

var FloodfillPaintImage func(m *MagickWand, channel I.ChannelType, fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y int32, invert bool) bool

func (m *MagickWand) FloodfillPaint(channel I.ChannelType, fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y int32, invert bool) bool {
	return FloodfillPaintImage(m, channel, fill, fuzz, bordercolor, x, y, invert)
}

var FlopImage func(m *MagickWand) bool

func (m *MagickWand) Flop() bool { return FlopImage(m) }

var ForwardFourierTransformImage func(m *MagickWand, magnitude bool) bool

func (m *MagickWand) ForwardFourierTransform(magnitude bool) bool {
	return ForwardFourierTransformImage(m, magnitude)
}

var FrameImage func(m *MagickWand, matteColor *PixelWand, width, height uint32, innerBevel, outerBevel int32) bool

func (m *MagickWand) Frame(matteColor *PixelWand, width, height uint32, innerBevel, outerBevel int32) bool {
	return FrameImage(m, matteColor, width, height, innerBevel, outerBevel)
}

var FunctionImage func(m *MagickWand, function I.MagickFunction, numberArguments uint32, arguments *float64) bool

func (m *MagickWand) Function(function I.MagickFunction, numberArguments uint32, arguments *float64) bool {
	return FunctionImage(m, function, numberArguments, arguments)
}

var FunctionImageChannel func(m *MagickWand, channel I.ChannelType, function I.MagickFunction, numberArguments uint32, arguments *float64) bool

func (m *MagickWand) FunctionChannel(channel I.ChannelType, function I.MagickFunction, numberArguments uint32, arguments *float64) bool {
	return FunctionImageChannel(m, channel, function, numberArguments, arguments)
}

var FxImage func(m *MagickWand, expression string) *MagickWand

func (m *MagickWand) Fx(expression string) *MagickWand { return FxImage(m, expression) }

var FxImageChannel func(m *MagickWand, channel I.ChannelType, expression string) *MagickWand

func (m *MagickWand) FxChannel(channel I.ChannelType, expression string) *MagickWand {
	return FxImageChannel(m, channel, expression)
}

var GammaImage func(m *MagickWand, gamma float64) bool

func (m *MagickWand) Gamma(gamma float64) bool { return GammaImage(m, gamma) }

var GammaImageChannel func(m *MagickWand, channel I.ChannelType, gamma float64) bool

func (m *MagickWand) GammaChannel(channel I.ChannelType, gamma float64) bool {
	return GammaImageChannel(m, channel, gamma)
}

var GaussianBlurImage func(m *MagickWand, radius, sigma float64) bool

func (m *MagickWand) GaussianBlur(radius, sigma float64) bool {
	return GaussianBlurImage(m, radius, sigma)
}

var GaussianBlurImageChannel func(m *MagickWand, channel I.ChannelType, radius, sigma float64) bool

func (m *MagickWand) GaussianBlurChannel(channel I.ChannelType, radius, sigma float64) bool {
	return GaussianBlurImageChannel(m, channel, radius, sigma)
}

var GetAntialias func(m *MagickWand) bool

func (m *MagickWand) Antialias() bool { return GetAntialias(m) }

var GetBackgroundColor func(m *MagickWand) *PixelWand

func (m *MagickWand) BackgroundColor() *PixelWand { return GetBackgroundColor(m) }

var GetColorspace func(m *MagickWand) I.ColorspaceType

func (m *MagickWand) Colorspace() I.ColorspaceType { return GetColorspace(m) }

var GetCompression func(m *MagickWand) I.CompressionType

func (m *MagickWand) Compression() I.CompressionType { return GetCompression(m) }

var GetCompressionQuality func(m *MagickWand) uint32

func (m *MagickWand) CompressionQuality() uint32 { return GetCompressionQuality(m) }

var GetException func(m *MagickWand, severity *I.ExceptionType) string

func (m *MagickWand) Exception(severity *I.ExceptionType) string {
	return GetException(m, severity)
}

var GetExceptionType func(m *MagickWand) I.ExceptionType

func (m *MagickWand) ExceptionType() I.ExceptionType { return GetExceptionType(m) }

var GetFilename func(m *MagickWand) string

func (m *MagickWand) Filename() string { return GetFilename(m) }

var GetFont func(m *MagickWand) string

func (m *MagickWand) Font() string { return GetFont(m) }

var GetFormat func(m *MagickWand) string

func (m *MagickWand) Format() string { return GetFormat(m) }

var GetGravity func(m *MagickWand) I.GravityType

func (m *MagickWand) Gravity() I.GravityType { return GetGravity(m) }

var GetImage func(m *MagickWand) *MagickWand

func (m *MagickWand) Image() *MagickWand { return GetImage(m) }

var GetImageAlphaChannel func(m *MagickWand) uint32

func (m *MagickWand) AlphaChannel() uint32 { return GetImageAlphaChannel(m) }

var GetImageArtifact func(m *MagickWand, artifact string) string

func (m *MagickWand) Artifact(artifact string) string { return GetImageArtifact(m, artifact) }

var GetImageArtifacts func(m *MagickWand, pattern, numberArtifacts *uint32) string

func (m *MagickWand) Artifacts(pattern, numberArtifacts *uint32) string {
	return GetImageArtifacts(m, pattern, numberArtifacts)
}

var GetImageBackgroundColor func(m *MagickWand, backgroundColor *PixelWand) bool

func (m *MagickWand) ImageBackgroundColor(backgroundColor *PixelWand) bool {
	return GetImageBackgroundColor(m, backgroundColor)
}

var GetImageBlob func(m *MagickWand, length *uint32) *byte

func (m *MagickWand) Blob(length *uint32) *byte { return GetImageBlob(m, length) }

var GetImageBluePrimary func(m *MagickWand, x, y *float64) bool

func (m *MagickWand) BluePrimary(x, y *float64) bool { return GetImageBluePrimary(m, x, y) }

var GetImageBorderColor func(m *MagickWand, borderColor *PixelWand) bool

func (m *MagickWand) BorderColor(borderColor *PixelWand) bool {
	return GetImageBorderColor(m, borderColor)
}

var GetImageChannelDepth func(m *MagickWand, channel I.ChannelType) uint32

func (m *MagickWand) ChannelDepth(channel I.ChannelType) uint32 {
	return GetImageChannelDepth(m, channel)
}

var GetImageChannelDistortion func(m *MagickWand, reference *MagickWand, channel I.ChannelType, metric I.MetricType, distortion *float64) bool

func (m *MagickWand) ChannelDistortion(reference *MagickWand, channel I.ChannelType, metric I.MetricType, distortion *float64) bool {
	return GetImageChannelDistortion(m, reference, channel, metric, distortion)
}

var GetImageChannelDistortions func(m *MagickWand, reference *MagickWand, metric I.MetricType) []float64

func (m *MagickWand) ChannelDistortions(reference *MagickWand, metric I.MetricType) []float64 {
	return GetImageChannelDistortions(m, reference, metric)
}

var GetImageChannelFeatures func(m *MagickWand, distance uint32) []I.ChannelFeatures

func (m *MagickWand) ChannelFeatures(distance uint32) []I.ChannelFeatures {
	return GetImageChannelFeatures(m, distance)
}

var GetImageChannelKurtosis func(m *MagickWand, channel I.ChannelType, kurtosis, skewness *float64) bool

func (m *MagickWand) ChannelKurtosis(channel I.ChannelType, kurtosis, skewness *float64) bool {
	return GetImageChannelKurtosis(m, channel, kurtosis, skewness)
}

var GetImageChannelMean func(m *MagickWand, channel I.ChannelType, mean, standardDeviation *float64) bool

func (m *MagickWand) ChannelMean(channel I.ChannelType, mean, standardDeviation *float64) bool {
	return GetImageChannelMean(m, channel, mean, standardDeviation)
}

var GetImageChannelRange func(m *MagickWand, channel I.ChannelType, minima, maxima *float64) bool

func (m *MagickWand) ChannelRange(channel I.ChannelType, minima, maxima *float64) bool {
	return GetImageChannelRange(m, channel, minima, maxima)
}

var GetImageChannelStatistics func(m *MagickWand) []I.ChannelStatistics

func (m *MagickWand) ChannelStatistics() []I.ChannelStatistics {
	return GetImageChannelStatistics(m)
}

var GetImageClipMask func(m *MagickWand) *MagickWand

func (m *MagickWand) ClipMask() *MagickWand { return GetImageClipMask(m) }

var GetImageColormapColor func(m *MagickWand, index uint32, color *PixelWand) bool

func (m *MagickWand) ColormapColor(index uint32, color *PixelWand) bool {
	return GetImageColormapColor(m, index, color)
}

var GetImageColors func(m *MagickWand) uint32

func (m *MagickWand) Colors() uint32 { return GetImageColors(m) }

var GetImageColorspace func(m *MagickWand) I.ColorspaceType

func (m *MagickWand) ImageColorspace() I.ColorspaceType { return GetImageColorspace(m) }

var GetImageCompose func(m *MagickWand) I.CompositeOperator

func (m *MagickWand) Compose() I.CompositeOperator { return GetImageCompose(m) }

var GetImageCompression func(m *MagickWand) I.CompressionType

func (m *MagickWand) ImageCompression() I.CompressionType { return GetImageCompression(m) }

var GetImageCompressionQuality func(m *MagickWand) uint32

func (m *MagickWand) ImageCompressionQuality() uint32 { return GetImageCompressionQuality(m) }

var GetImageDelay func(m *MagickWand) uint32

func (m *MagickWand) Delay() uint32 { return GetImageDelay(m) }

var GetImageDepth func(m *MagickWand) uint32

func (m *MagickWand) Depth() uint32 { return GetImageDepth(m) }

var GetImageDispose func(m *MagickWand) I.DisposeType

func (m *MagickWand) Dispose() I.DisposeType { return GetImageDispose(m) }

var GetImageDistortion func(m *MagickWand, reference *MagickWand, metric I.MetricType, distortion *float64) bool

func (m *MagickWand) Distortion(reference *MagickWand, metric I.MetricType, distortion *float64) bool {
	return GetImageDistortion(m, reference, metric, distortion)
}

var GetImageFilename func(m *MagickWand) string

func (m *MagickWand) ImageFilename() string { return GetImageFilename(m) }

var GetImageFormat func(m *MagickWand) string

func (m *MagickWand) ImageFormat() string { return GetImageFormat(m) }

var GetImageFuzz func(m *MagickWand) float64

func (m *MagickWand) Fuzz() float64 { return GetImageFuzz(m) }

var GetImageGamma func(m *MagickWand) float64

func (m *MagickWand) ImageGamma() float64 { return GetImageGamma(m) }

var GetImageGravity func(m *MagickWand) I.GravityType

func (m *MagickWand) ImageGravity() I.GravityType { return GetImageGravity(m) }

var GetImageGreenPrimary func(m *MagickWand, x, y *float64) bool

func (m *MagickWand) GreenPrimary(x, y *float64) bool { return GetImageGreenPrimary(m, x, y) }

var GetImageHeight func(m *MagickWand) uint32

func (m *MagickWand) Height() uint32 { return GetImageHeight(m) }

var GetImageHistogram func(m *MagickWand, numberColors *uint32) []*PixelWand

func (m *MagickWand) Histogram(numberColors *uint32) []*PixelWand {
	return GetImageHistogram(m, numberColors)
}

var GetImageInterlaceScheme func(m *MagickWand) I.InterlaceType

func (m *MagickWand) ImageInterlaceScheme() I.InterlaceType { return GetImageInterlaceScheme(m) }

var GetImageInterpolateMethod func(m *MagickWand) I.InterpolatePixelMethod

func (m *MagickWand) ImageInterpolateMethod() I.InterpolatePixelMethod {
	return GetImageInterpolateMethod(m)
}

var GetImageIterations func(m *MagickWand) uint32

func (m *MagickWand) Iterations() uint32 { return GetImageIterations(m) }

var GetImageLength func(m *MagickWand, length *I.MagickSizeType) bool

func (m *MagickWand) Length(length *I.MagickSizeType) bool {
	return GetImageLength(m, length)
}

var GetImageMatteColor func(m *MagickWand, matteColor *PixelWand) bool

func (m *MagickWand) MatteColor(matteColor *PixelWand) bool {
	return GetImageMatteColor(m, matteColor)
}

var GetImageOrientation func(m *MagickWand) I.OrientationType

func (m *MagickWand) ImageOrientation() I.OrientationType { return GetImageOrientation(m) }

var GetImagePage func(m *MagickWand, width, height *uint32, x, y *int32) bool

func (m *MagickWand) ImagePage(width, height *uint32, x, y *int32) bool {
	return GetImagePage(m, width, height, x, y)
}

var GetImagePixelColor func(m *MagickWand, x, y int32, color *PixelWand) bool

func (m *MagickWand) PixelColor(x, y int32, color *PixelWand) bool {
	return GetImagePixelColor(m, x, y, color)
}

var GetImageProfile func(m *MagickWand, name string, length *uint32) *byte

func (m *MagickWand) Profile(name string, length *uint32) *byte {
	return GetImageProfile(m, name, length)
}

var GetImageProfiles func(m *MagickWand, pattern, numberProfiles *uint32) string

func (m *MagickWand) Profiles(pattern, numberProfiles *uint32) string {
	return GetImageProfiles(m, pattern, numberProfiles)
}

var GetImageProperties func(m *MagickWand, pattern, numberProperties *uint32) string

func (m *MagickWand) Properties(pattern, numberProperties *uint32) string {
	return GetImageProperties(m, pattern, numberProperties)
}

var GetImageProperty func(m *MagickWand, property string) string

func (m *MagickWand) Property(property string) string { return GetImageProperty(m, property) }

var GetImageRedPrimary func(m *MagickWand, x, y *float64) bool

func (m *MagickWand) RedPrimary(x, y *float64) bool { return GetImageRedPrimary(m, x, y) }

var GetImageRegion func(m *MagickWand, width, height uint32, x, y int32) *MagickWand

func (m *MagickWand) Region(width, height uint32, x, y int32) *MagickWand {
	return GetImageRegion(m, width, height, x, y)
}

var GetImageRenderingIntent func(m *MagickWand) I.RenderingIntent

func (m *MagickWand) RenderingIntent() I.RenderingIntent { return GetImageRenderingIntent(m) }

var GetImageResolution func(m *MagickWand, x, y *float64) bool

func (m *MagickWand) ImageResolution(x, y *float64) bool { return GetImageResolution(m, x, y) }

var GetImagesBlob func(m *MagickWand, length *uint32) *byte

func (m *MagickWand) ImagesBlob(length *uint32) *byte { return GetImagesBlob(m, length) }

var GetImageScene func(m *MagickWand) uint32

func (m *MagickWand) ImageScene() uint32 { return GetImageScene(m) }

var GetImageSignature func(m *MagickWand) string

func (m *MagickWand) ImageSignature() string { return GetImageSignature(m) }

var GetImageTicksPerSecond func(m *MagickWand) uint32

func (m *MagickWand) TicksPerSecond() uint32 { return GetImageTicksPerSecond(m) }

var GetImageTotalInkDensity func(m *MagickWand) float64

func (m *MagickWand) TotalInkDensity() float64 { return GetImageTotalInkDensity(m) }

var GetImageType func(m *MagickWand) I.ImageType

func (m *MagickWand) ImageType() I.ImageType { return GetImageType(m) }

var GetImageUnits func(m *MagickWand) I.ResolutionType

func (m *MagickWand) ImageUnits() I.ResolutionType { return GetImageUnits(m) }

var GetImageVirtualPixelMethod func(m *MagickWand) I.VirtualPixelMethod

func (m *MagickWand) ImageVirtualPixelMethod() I.VirtualPixelMethod {
	return GetImageVirtualPixelMethod(m)
}

var GetImageWhitePoint func(m *MagickWand, x, y *float64) bool

func (m *MagickWand) ImageWhitePoint(x, y *float64) bool { return GetImageWhitePoint(m, x, y) }

var GetImageWidth func(m *MagickWand) uint32

func (m *MagickWand) ImageWidth() uint32 { return GetImageWidth(m) }

var GetInterlaceScheme func(m *MagickWand) I.InterlaceType

func (m *MagickWand) InterlaceScheme() I.InterlaceType { return GetInterlaceScheme(m) }

var GetInterpolateMethod func(m *MagickWand) I.InterpolatePixelMethod

func (m *MagickWand) InterpolateMethod() I.InterpolatePixelMethod {
	return GetInterpolateMethod(m)
}

var GetIteratorIndex func(m *MagickWand) int32

func (m *MagickWand) IteratorIndex() int32 { return GetIteratorIndex(m) }

var GetNumberImages func(m *MagickWand) uint32

func (m *MagickWand) NumberImages() uint32 { return GetNumberImages(m) }

var GetOption func(m *MagickWand, key string) string

func (m *MagickWand) Option(key string) string { return GetOption(m, key) }

var GetOptions func(m *MagickWand, pattern, numberOptions *uint32) string

func (m *MagickWand) Options(pattern, numberOptions *uint32) string {
	return GetOptions(m, pattern, numberOptions)
}

var GetOrientation func(m *MagickWand) I.OrientationType

func (m *MagickWand) Orientation() I.OrientationType { return GetOrientation(m) }

var GetPage func(m *MagickWand, width, height *uint32, x, y *int32) bool

func (m *MagickWand) Page(width, height *uint32, x, y *int32) bool {
	return GetPage(m, width, height, x, y)
}

var GetPointsize func(m *MagickWand) float64

func (m *MagickWand) Pointsize() float64 { return GetPointsize(m) }

var GetResolution func(m *MagickWand, x, y *float64) bool

func (m *MagickWand) Resolution(x, y *float64) bool { return GetResolution(m, x, y) }

var GetSamplingFactors func(m *MagickWand, numberFactors *uint32) []float64

func (m *MagickWand) SamplingFactors(numberFactors *uint32) []float64 {
	return GetSamplingFactors(m, numberFactors)
}

var GetSize func(m *MagickWand, columns, rows *uint32) bool

func (m *MagickWand) Size(columns, rows *uint32) bool { return GetSize(m, columns, rows) }

var GetSizeOffset func(m *MagickWand, offset *int32) bool

func (m *MagickWand) SizeOffset(offset *int32) bool { return GetSizeOffset(m, offset) }

var GetType func(m *MagickWand) I.ImageType

func (m *MagickWand) Type() I.ImageType { return GetType(m) }

var HaldClutImage func(m *MagickWand, haldWand *MagickWand) bool

func (m *MagickWand) HaldClut(haldWand *MagickWand) bool {
	return HaldClutImage(m, haldWand)
}

var HaldClutImageChannel func(m *MagickWand, channel I.ChannelType, haldWand *MagickWand) bool

func (m *MagickWand) HaldClutChannel(channel I.ChannelType, haldWand *MagickWand) bool {
	return HaldClutImageChannel(m, channel, haldWand)
}

var HasNextImage func(m *MagickWand) bool

func (m *MagickWand) HasNext() bool { return HasNextImage(m) }

var HasPreviousImage func(m *MagickWand) bool

func (m *MagickWand) HasPrevious() bool { return HasPreviousImage(m) }

var IdentifyImage func(m *MagickWand) string

func (m *MagickWand) Identify() string { return IdentifyImage(m) }

var ImplodeImage func(m *MagickWand, radius float64) bool

func (m *MagickWand) Implode(radius float64) bool { return ImplodeImage(m, radius) }

var ImportImagePixels func(m *MagickWand, x, y int32, columns, rows uint32, map_ string, storage I.StorageType, pixels *Void) bool

func (m *MagickWand) ImportPixels(x, y int32, columns, rows uint32, map_ string, storage I.StorageType, pixels *Void) bool {
	return ImportImagePixels(m, x, y, columns, rows, map_, storage, pixels)
}

var InverseFourierTransformImage func(m *MagickWand, phaseWand *MagickWand, magnitude bool) bool

func (m *MagickWand) InverseFourierTransform(phaseWand *MagickWand, magnitude bool) bool {
	return InverseFourierTransformImage(m, phaseWand, magnitude)
}

var LabelImage func(m *MagickWand, label string) bool

func (m *MagickWand) Label(label string) bool { return LabelImage(m, label) }

var LevelImage func(m *MagickWand, blackPoint, gamma, whitePoint float64) bool

func (m *MagickWand) Level(blackPoint, gamma, whitePoint float64) bool {
	return LevelImage(m, blackPoint, gamma, whitePoint)
}

var LevelImageChannel func(m *MagickWand, channel I.ChannelType, blackPoint, gamma, whitePoint float64) bool

func (m *MagickWand) LevelChannel(channel I.ChannelType, blackPoint, gamma, whitePoint float64) bool {
	return LevelImageChannel(m, channel, blackPoint, gamma, whitePoint)
}

var LinearStretchImage func(m *MagickWand, blackPoint, whitePoint float64) bool

func (m *MagickWand) LinearStretch(blackPoint, whitePoint float64) bool {
	return LinearStretchImage(m, blackPoint, whitePoint)
}

var MagnifyImage func(m *MagickWand) bool

func (m *MagickWand) Magnify() bool { return MagnifyImage(m) }

var MapImage func(m *MagickWand, mapWand *MagickWand, dither bool) bool

func (m *MagickWand) Map(mapWand *MagickWand, dither bool) bool {
	return MapImage(m, mapWand, dither)
}

var MergeImageLayers func(m *MagickWand, method I.ImageLayerMethod) *MagickWand

func (m *MagickWand) MergeLayers(method I.ImageLayerMethod) *MagickWand {
	return MergeImageLayers(m, method)
}

var MinifyImage func(m *MagickWand) bool

func (m *MagickWand) Minify() bool { return MinifyImage(m) }

var ModulateImage func(m *MagickWand, brightness, saturation, hue float64) bool

func (m *MagickWand) Modulate(brightness, saturation, hue float64) bool {
	return ModulateImage(m, brightness, saturation, hue)
}

var MontageImage func(m *MagickWand, drawingWand DrawingWand, tileGeometry, thumbnailGeometry string, mode I.MontageMode, frame string) *MagickWand

func (m *MagickWand) Montage(drawingWand DrawingWand, tileGeometry, thumbnailGeometry string, mode I.MontageMode, frame string) *MagickWand {
	return MontageImage(m, drawingWand, tileGeometry, thumbnailGeometry, mode, frame)
}

var MorphImages func(m *MagickWand, numberFrames uint32) *MagickWand

func (m *MagickWand) Morph(numberFrames uint32) *MagickWand {
	return MorphImages(m, numberFrames)
}

var MorphologyImage func(m *MagickWand, method I.MorphologyMethod, iterations int32, kernel *I.KernelInfo) bool

func (m *MagickWand) Morphology(method I.MorphologyMethod, iterations int32, kernel *I.KernelInfo) bool {
	return MorphologyImage(m, method, iterations, kernel)
}

var MorphologyImageChannel func(m *MagickWand, channel I.ChannelType, method I.MorphologyMethod, iterations int32, kernel *I.KernelInfo) bool

func (m *MagickWand) MorphologyChannel(channel I.ChannelType, method I.MorphologyMethod, iterations int32, kernel *I.KernelInfo) bool {
	return MorphologyImageChannel(m, channel, method, iterations, kernel)
}

var MotionBlurImage func(m *MagickWand, radius, sigma, angle float64) bool

func (m *MagickWand) MotionBlur(radius, sigma, angle float64) bool {
	return MotionBlurImage(m, radius, sigma, angle)
}

var MotionBlurImageChannel func(m *MagickWand, channel I.ChannelType, radius, sigma, angle float64) bool

func (m *MagickWand) MotionBlurChannel(channel I.ChannelType, radius, sigma, angle float64) bool {
	return MotionBlurImageChannel(m, channel, radius, sigma, angle)
}

var NegateImage func(m *MagickWand, gray bool) bool

func (m *MagickWand) Negate(gray bool) bool { return NegateImage(m, gray) }

var NegateImageChannel func(m *MagickWand, channel I.ChannelType, gray bool) bool

func (m *MagickWand) NegateChannel(channel I.ChannelType, gray bool) bool {
	return NegateImageChannel(m, channel, gray)
}

var NewImage func(m *MagickWand, columns, rows uint32, background *PixelWand) bool

func (m *MagickWand) New(columns, rows uint32, background *PixelWand) bool {
	return NewImage(m, columns, rows, background)
}

var NextImage func(m *MagickWand) bool

func (m *MagickWand) Next() bool { return NextImage(m) }

var NormalizeImage func(m *MagickWand) bool

func (m *MagickWand) Normalize() bool { return NormalizeImage(m) }

var NormalizeImageChannel func(m *MagickWand, channel I.ChannelType) bool

func (m *MagickWand) NormalizeChannel(channel I.ChannelType) bool {
	return NormalizeImageChannel(m, channel)
}

var OilPaintImage func(m *MagickWand, radius float64) bool

func (m *MagickWand) OilPaint(radius float64) bool { return OilPaintImage(m, radius) }

var OpaquePaintImage func(m *MagickWand, target, fill *PixelWand, fuzz float64, invert bool) bool

func (m *MagickWand) OpaquePaint(target, fill *PixelWand, fuzz float64, invert bool) bool {
	return OpaquePaintImage(m, target, fill, fuzz, invert)
}

var OpaquePaintImageChannel func(m *MagickWand, channel I.ChannelType, target, fill *PixelWand, fuzz float64, invert bool) bool

func (m *MagickWand) OpaquePaintChannel(channel I.ChannelType, target, fill *PixelWand, fuzz float64, invert bool) bool {
	return OpaquePaintImageChannel(m, channel, target, fill, fuzz, invert)
}

var OptimizeImageLayers func(m *MagickWand) *MagickWand

func (m *MagickWand) OptimizeLayers() *MagickWand { return OptimizeImageLayers(m) }

var OrderedPosterizeImage func(m *MagickWand, thresholdMap string) bool

func (m *MagickWand) OrderedPosterize(thresholdMap string) bool {
	return OrderedPosterizeImage(m, thresholdMap)
}

var OrderedPosterizeImageChannel func(m *MagickWand, channel I.ChannelType, thresholdMap string) bool

func (m *MagickWand) OrderedPosterizeChannel(channel I.ChannelType, thresholdMap string) bool {
	return OrderedPosterizeImageChannel(m, channel, thresholdMap)
}

var PaintOpaqueImageChannel func(m *MagickWand, channel I.ChannelType, target, fill *PixelWand, fuzz float64) bool

func (m *MagickWand) PaintOpaqueChannel(channel I.ChannelType, target, fill *PixelWand, fuzz float64) bool {
	return PaintOpaqueImageChannel(m, channel, target, fill, fuzz)
}

var PingImage func(m *MagickWand, filename string) bool

func (m *MagickWand) Ping(filename string) bool { return PingImage(m, filename) }

var PingImageBlob func(m *MagickWand, blob *Void, length uint32) bool

func (m *MagickWand) PingBlob(blob *Void, length uint32) bool {
	return PingImageBlob(m, blob, length)
}

var PingImageFile func(m *MagickWand, file *FILE) bool

func (m *MagickWand) PingFile(file *FILE) bool { return PingImageFile(m, file) }

var PolaroidImage func(m *MagickWand, d *DrawingWand, angle float64) bool

func (m *MagickWand) Polaroid(d *DrawingWand, angle float64) bool {
	return PolaroidImage(m, d, angle)
}

var PosterizeImage func(m *MagickWand, levels uint32, dither bool) bool

func (m *MagickWand) Posterize(levels uint32, dither bool) bool {
	return PosterizeImage(m, levels, dither)
}

var PreviewImages func(m *MagickWand, preview I.PreviewType) *MagickWand

func (m *MagickWand) PreviewImages(preview I.PreviewType) *MagickWand {
	return PreviewImages(m, preview)
}

var PreviousImage func(m *MagickWand) bool

func (m *MagickWand) Previous() bool { return PreviousImage(m) }

var ProfileImage func(m *MagickWand, name string, profile *Void, length uint32) bool

func (m *MagickWand) ProfileImage(name string, profile *Void, length uint32) bool {
	return ProfileImage(m, name, profile, length)
}

var QuantizeImage func(m *MagickWand, numberColors uint32, colorspace I.ColorspaceType, treedepth uint32, dither, measureError bool) bool

func (m *MagickWand) Quantize(numberColors uint32, colorspace I.ColorspaceType, treedepth uint32, dither, measureError bool) bool {
	return QuantizeImage(m, numberColors, colorspace, treedepth, dither, measureError)
}

var QuantizeImages func(m *MagickWand, numberColors uint32, colorspace I.ColorspaceType, treedepth uint32, dither, measureError bool) bool

func (m *MagickWand) QuantizeImages(numberColors uint32, colorspace I.ColorspaceType, treedepth uint32, dither, measureError bool) bool {
	return QuantizeImages(m, numberColors, colorspace, treedepth, dither, measureError)
}

var QueryFontMetrics func(m *MagickWand, drawingWand *DrawingWand, text string) []float64

func (m *MagickWand) QueryFontMetrics(drawingWand *DrawingWand, text string) []float64 {
	return QueryFontMetrics(m, drawingWand, text)
}

var QueryMultilineFontMetrics func(m *MagickWand, drawingWand *DrawingWand, text string) []float64

func (m *MagickWand) QueryMultilineFontMetrics(drawingWand *DrawingWand, text string) []float64 {
	return QueryMultilineFontMetrics(m, drawingWand, text)
}

var RadialBlurImage func(m *MagickWand, angle float64) bool

func (m *MagickWand) RadialBlur(angle float64) bool { return RadialBlurImage(m, angle) }

var RadialBlurImageChannel func(m *MagickWand, channel I.ChannelType, angle float64) bool

func (m *MagickWand) RadialBlurChannel(channel I.ChannelType, angle float64) bool {
	return RadialBlurImageChannel(m, channel, angle)
}

var RaiseImage func(m *MagickWand, width, height uint32, x, y int32, raise bool) bool

func (m *MagickWand) Raise(width, height uint32, x, y int32, raise bool) bool {
	return RaiseImage(m, width, height, x, y, raise)
}

var RandomThresholdImage func(m *MagickWand, low, high float64) bool

func (m *MagickWand) RandomThreshold(low, high float64) bool {
	return RandomThresholdImage(m, low, high)
}

var RandomThresholdImageChannel func(m *MagickWand, channel I.ChannelType, low, high float64) bool

func (m *MagickWand) RandomThresholdChannel(channel I.ChannelType, low, high float64) bool {
	return RandomThresholdImageChannel(m, channel, low, high)
}

var ReadImage func(m *MagickWand, filename string) bool

func (m *MagickWand) Read(filename string) bool { return ReadImage(m, filename) }

var ReadImageBlob func(m *MagickWand, blob *byte, length uint32) bool

func (m *MagickWand) ReadBlob(blob *byte, length uint32) bool {
	return ReadImageBlob(m, blob, length)
}

var ReadImageFile func(m *MagickWand, file *FILE) bool

func (m *MagickWand) ReadFile(file *FILE) bool { return ReadImageFile(m, file) }

var RemapImage func(m *MagickWand, remapWand *MagickWand, method I.DitherMethod) bool

func (m *MagickWand) Remap(remapWand *MagickWand, method I.DitherMethod) bool {
	return RemapImage(m, remapWand, method)
}

var RemoveImage func(m *MagickWand) bool

func (m *MagickWand) Remove() bool { return RemoveImage(m) }

var RemoveImageProfile func(m *MagickWand, name string, length *uint32) []byte

func (m *MagickWand) RemoveProfile(name string, length *uint32) []byte {
	return RemoveImageProfile(m, name, length)
}

var ResampleImage func(m *MagickWand, xResolution, yResolution float64, filter I.FilterTypes, blur float64) bool

func (m *MagickWand) Resample(xResolution, yResolution float64, filter I.FilterTypes, blur float64) bool {
	return ResampleImage(m, xResolution, yResolution, filter, blur)
}

var ResetImagePage func(m *MagickWand, page string) bool

func (m *MagickWand) ResetPage(page string) bool { return ResetImagePage(m, page) }

var ResetIterator func(m *MagickWand)

func (m *MagickWand) ResetIterator() { ResetIterator(m) }

var ResizeImage func(m *MagickWand, columns, rows uint32, filter I.FilterTypes, blur float64) bool

func (m *MagickWand) Resize(columns, rows uint32, filter I.FilterTypes, blur float64) bool {
	return ResizeImage(m, columns, rows, filter, blur)
}

var RollImage func(m *MagickWand, x int32, y uint32) bool

func (m *MagickWand) Roll(x int32, y uint32) bool { return RollImage(m, x, y) }

var RotateImage func(m *MagickWand, background *PixelWand, degrees float64) bool

func (m *MagickWand) Rotate(background *PixelWand, degrees float64) bool {
	return RotateImage(m, background, degrees)
}

var SampleImage func(m *MagickWand, columns, rows uint32) bool

func (m *MagickWand) Sample(columns, rows uint32) bool {
	return SampleImage(m, columns, rows)
}

var ScaleImage func(m *MagickWand, columns, rows uint32) bool

func (m *MagickWand) Scale(columns, rows uint32) bool { return ScaleImage(m, columns, rows) }

var SegmentImage func(m *MagickWand, colorspace I.ColorspaceType, verbose bool, clusterThreshold, smoothThreshold float64) bool

func (m *MagickWand) Segment(colorspace I.ColorspaceType, verbose bool, clusterThreshold, smoothThreshold float64) bool {
	return SegmentImage(m, colorspace, verbose, clusterThreshold, smoothThreshold)
}

var SelectiveBlurImage func(m *MagickWand, radius, sigma, threshold float64) bool

func (m *MagickWand) SelectiveBlur(radius, sigma, threshold float64) bool {
	return SelectiveBlurImage(m, radius, sigma, threshold)
}

var SelectiveBlurImageChannel func(m *MagickWand, channel I.ChannelType, radius, sigma, threshold float64) bool

func (m *MagickWand) SelectiveBlurChannel(channel I.ChannelType, radius, sigma, threshold float64) bool {
	return SelectiveBlurImageChannel(m, channel, radius, sigma, threshold)
}

var SeparateImageChannel func(m *MagickWand, channel I.ChannelType) bool

func (m *MagickWand) SeparateChannel(channel I.ChannelType) bool {
	return SeparateImageChannel(m, channel)
}

var SepiaToneImage func(m *MagickWand, threshold float64) bool

func (m *MagickWand) SepiaTone(threshold float64) bool {
	return SepiaToneImage(m, threshold)
}

var SetAntialias func(m *MagickWand, antialias bool) bool

func (m *MagickWand) SetAntialias(antialias bool) bool { return SetAntialias(m, antialias) }

var SetBackgroundColor func(m *MagickWand, background *PixelWand) bool

func (m *MagickWand) SetBackgroundColor(background *PixelWand) bool {
	return SetBackgroundColor(m, background)
}

var SetColorspace func(m *MagickWand, colorspace I.ColorspaceType) bool

func (m *MagickWand) SetColorspace(colorspace I.ColorspaceType) bool {
	return SetColorspace(m, colorspace)
}

var SetCompression func(m *MagickWand, compression I.CompressionType) bool

func (m *MagickWand) SetCompression(compression I.CompressionType) bool {
	return SetCompression(m, compression)
}

var SetCompressionQuality func(m *MagickWand, quality uint32) bool

func (m *MagickWand) SetCompressionQuality(quality uint32) bool {
	return SetCompressionQuality(m, quality)
}

var SetDepth func(m *MagickWand, depth uint32) bool

func (m *MagickWand) SetDepth(depth uint32) bool { return SetDepth(m, depth) }

var SetExtract func(m *MagickWand, geometry string) bool

func (m *MagickWand) SetExtract(geometry string) bool { return SetExtract(m, geometry) }

var SetFilename func(m *MagickWand, filename string) bool

func (m *MagickWand) SetFilename(filename string) bool { return SetFilename(m, filename) }

var SetFirstIterator func(m *MagickWand)

func (m *MagickWand) SetFirstIterator() { SetFirstIterator(m) }

var SetFont func(m *MagickWand, font string) bool

func (m *MagickWand) SetFont(font string) bool { return SetFont(m, font) }

var SetFormat func(m *MagickWand, format string) bool

func (m *MagickWand) SetFormat(format string) bool { return SetFormat(m, format) }

var SetGravity func(m *MagickWand, type_ I.GravityType) bool

func (m *MagickWand) SetGravity(type_ I.GravityType) bool { return SetGravity(m, type_) }

var SetImage func(m *MagickWand, setWand *MagickWand) bool

func (m *MagickWand) SetImage(setWand *MagickWand) bool { return SetImage(m, setWand) }

var SetImageAlphaChannel func(m *MagickWand, alphaType I.AlphaChannelType) bool

func (m *MagickWand) SetAlphaChannel(alphaType I.AlphaChannelType) bool {
	return SetImageAlphaChannel(m, alphaType)
}

var SetImageArtifact func(m *MagickWand, artifact, value string) bool

func (m *MagickWand) SetArtifact(artifact, value string) bool {
	return SetImageArtifact(m, artifact, value)
}

var SetImageBackgroundColor func(m *MagickWand, background *PixelWand) bool

func (m *MagickWand) SetImageBackgroundColor(background *PixelWand) bool {
	return SetImageBackgroundColor(m, background)
}

var SetImageBias func(m *MagickWand, bias float64) bool

func (m *MagickWand) SetBias(bias float64) bool { return SetImageBias(m, bias) }

var SetImageBluePrimary func(m *MagickWand, x, y float64) bool

func (m *MagickWand) SetBluePrimary(x, y float64) bool {
	return SetImageBluePrimary(m, x, y)
}

var SetImageBorderColor func(m *MagickWand, border *PixelWand) bool

func (m *MagickWand) SetBorderColor(border *PixelWand) bool {
	return SetImageBorderColor(m, border)
}

var SetImageChannelDepth func(m *MagickWand, channel I.ChannelType, depth uint32) bool

func (m *MagickWand) SetChannelDepth(channel I.ChannelType, depth uint32) bool {
	return SetImageChannelDepth(m, channel, depth)
}

var SetImageClipMask func(m *MagickWand, clipMask *MagickWand) bool

func (m *MagickWand) SetClipMask(clipMask *MagickWand) bool {
	return SetImageClipMask(m, clipMask)
}

var SetImageColor func(m *MagickWand, color *PixelWand) bool

func (m *MagickWand) SetColor(color *PixelWand) bool { return SetImageColor(m, color) }

var SetImageColormapColor func(m *MagickWand, index uint32, color *PixelWand) bool

func (m *MagickWand) SetColormapColor(index uint32, color *PixelWand) bool {
	return SetImageColormapColor(m, index, color)
}

var SetImageColorspace func(m *MagickWand, colorspace I.ColorspaceType) bool

func (m *MagickWand) SetImageColorspace(colorspace I.ColorspaceType) bool {
	return SetImageColorspace(m, colorspace)
}

var SetImageCompose func(m *MagickWand, compose I.CompositeOperator) bool

func (m *MagickWand) SetCompose(compose I.CompositeOperator) bool {
	return SetImageCompose(m, compose)
}

var SetImageCompression func(m *MagickWand, compression I.CompressionType) bool

func (m *MagickWand) SetImageCompression(compression I.CompressionType) bool {
	return SetImageCompression(m, compression)
}

var SetImageCompressionQuality func(m *MagickWand, quality uint32) bool

func (m *MagickWand) SetImageCompressionQuality(quality uint32) bool {
	return SetImageCompressionQuality(m, quality)
}

var SetImageDelay func(m *MagickWand, delay uint32) bool

func (m *MagickWand) SetDelay(delay uint32) bool { return SetImageDelay(m, delay) }

var SetImageDepth func(m *MagickWand, depth uint32) bool

func (m *MagickWand) SetImageDepth(depth uint32) bool { return SetImageDepth(m, depth) }

var SetImageDispose func(m *MagickWand, dispose I.DisposeType) bool

func (m *MagickWand) SetDispose(dispose I.DisposeType) bool {
	return SetImageDispose(m, dispose)
}

var SetImageEndian func(m *MagickWand, endian I.EndianType) bool

func (m *MagickWand) SetEndian(endian I.EndianType) bool { return SetImageEndian(m, endian) }

var SetImageExtent func(m *MagickWand, columns uint32, rows uint) bool

func (m *MagickWand) SetExtent(columns uint32, rows uint) bool {
	return SetImageExtent(m, columns, rows)
}

var SetImageFilename func(m *MagickWand, filename string) bool

func (m *MagickWand) SetImageFilename(filename string) bool {
	return SetImageFilename(m, filename)
}

var SetImageFormat func(m *MagickWand, format string) bool

func (m *MagickWand) SetImageFormat(format string) bool { return SetImageFormat(m, format) }

var SetImageFuzz func(m *MagickWand, fuzz float64) bool

func (m *MagickWand) SetFuzz(fuzz float64) bool { return SetImageFuzz(m, fuzz) }

var SetImageGamma func(m *MagickWand, gamma float64) bool

func (m *MagickWand) SetGamma(gamma float64) bool { return SetImageGamma(m, gamma) }

var SetImageGravity func(m *MagickWand, gravity I.GravityType) bool

func (m *MagickWand) SetImageGravity(gravity I.GravityType) bool {
	return SetImageGravity(m, gravity)
}

var SetImageGreenPrimary func(m *MagickWand, x, y float64) bool

func (m *MagickWand) SetGreenPrimary(x, y float64) bool {
	return SetImageGreenPrimary(m, x, y)
}

var SetImageInterlaceScheme func(m *MagickWand, interlace I.InterlaceType) bool

func (m *MagickWand) SetImageInterlaceScheme(interlace I.InterlaceType) bool {
	return SetImageInterlaceScheme(m, interlace)
}

var SetImageInterpolateMethod func(m *MagickWand, method I.InterpolatePixelMethod) bool

func (m *MagickWand) SetImageInterpolateMethod(method I.InterpolatePixelMethod) bool {
	return SetImageInterpolateMethod(m, method)
}

var SetImageIterations func(m *MagickWand, iterations uint32) bool

func (m *MagickWand) SetIterations(iterations uint32) bool {
	return SetImageIterations(m, iterations)
}

var SetImageMatte func(m *MagickWand, matte *bool) bool

func (m *MagickWand) SetMatte(matte *bool) bool { return SetImageMatte(m, matte) }

var SetImageMatteColor func(m *MagickWand, matte *PixelWand) bool

func (m *MagickWand) SetMatteColor(matte *PixelWand) bool {
	return SetImageMatteColor(m, matte)
}

var SetImageOpacity func(m *MagickWand, alpha float64) bool

func (m *MagickWand) SetOpacity(alpha float64) bool { return SetImageOpacity(m, alpha) }

var SetImageOrientation func(m *MagickWand, orientation I.OrientationType) bool

func (m *MagickWand) SetImageOrientation(orientation I.OrientationType) bool {
	return SetImageOrientation(m, orientation)
}

var SetImagePage func(m *MagickWand, width, height uint32, x, y int32) bool

func (m *MagickWand) SetImagePage(width, height uint32, x, y int32) bool {
	return SetImagePage(m, width, height, x, y)
}

var SetImageProfile func(m *MagickWand, name string, profile *Void, length uint32) bool

func (m *MagickWand) SetImageProfile(name string, profile *Void, length uint32) bool {
	return SetImageProfile(m, name, profile, length)
}

var SetImageProgressMonitor func(m *MagickWand, progressMonitor, clientData *Void) I.MagickProgressMonitor

func (m *MagickWand) SetImageProgressMonitor(progressMonitor, clientData *Void) I.MagickProgressMonitor {
	return SetImageProgressMonitor(m, progressMonitor, clientData)
}

var SetImageProperty func(m *MagickWand, property, value string) bool

func (m *MagickWand) SetImageProperty(property, value string) bool {
	return SetImageProperty(m, property, value)
}

var SetImageRedPrimary func(m *MagickWand, x, y float64) bool

func (m *MagickWand) SetRedPrimary(x, y float64) bool { return SetImageRedPrimary(m, x, y) }

var SetImageRenderingIntent func(m *MagickWand, renderingIntent I.RenderingIntent) bool

func (m *MagickWand) SetRenderingIntent(renderingIntent I.RenderingIntent) bool {
	return SetImageRenderingIntent(m, renderingIntent)
}

var SetImageResolution func(m *MagickWand, xResolution, yResolution float64) bool

func (m *MagickWand) SetImageResolution(xResolution, yResolution float64) bool {
	return SetResolution(m, xResolution, yResolution)
}

var SetImageScene func(m *MagickWand, scene uint32) bool

func (m *MagickWand) SetScene(scene uint32) bool { return SetImageScene(m, scene) }

var SetImageTicksPerSecond func(m *MagickWand, ticksPerSecond int32) bool

func (m *MagickWand) SetImageTicksPerSecond(ticksPerSecond int32) bool {
	return SetImageTicksPerSecond(m, ticksPerSecond)
}

var SetImageType func(m *MagickWand, imageType I.ImageType) bool

func (m *MagickWand) SetImageType(imageType I.ImageType) bool {
	return SetImageType(m, imageType)
}

var SetImageUnits func(m *MagickWand, units I.ResolutionType) bool

func (m *MagickWand) SetUnits(units I.ResolutionType) bool { return SetImageUnits(m, units) }

var SetImageVirtualPixelMethod func(m *MagickWand, method I.VirtualPixelMethod) I.VirtualPixelMethod

func (m *MagickWand) SetVirtualPixelMethod(method I.VirtualPixelMethod) I.VirtualPixelMethod {
	return SetImageVirtualPixelMethod(m, method)
}

var SetImageWhitePoint func(m *MagickWand, x, y float64) bool

func (m *MagickWand) SetImageWhitePoint(x, y float64) bool { return SetImageWhitePoint(m, x, y) }

var SetInterlaceScheme func(m *MagickWand, interlaceScheme I.InterlaceType) bool

func (m *MagickWand) SetInterlaceScheme(interlaceScheme I.InterlaceType) bool {
	return SetInterlaceScheme(m, interlaceScheme)
}

var SetInterpolateMethod func(m *MagickWand, method I.InterpolatePixelMethod) bool

func (m *MagickWand) SetInterpolateMethod(method I.InterpolatePixelMethod) bool {
	return SetInterpolateMethod(m, method)
}

var SetIteratorIndex func(m *MagickWand, index int32) bool

func (m *MagickWand) SetIteratorIndex(index int32) bool { return SetIteratorIndex(m, index) }

var SetLastIterator func(m *MagickWand)

func (m *MagickWand) SetLastIterator() { SetLastIterator(m) }

var SetOption func(m *MagickWand, key, value string) bool

func (m *MagickWand) SetOption(key, value string) bool { return SetOption(m, key, value) }

var SetOrientation func(m *MagickWand, orientation I.OrientationType) bool

func (m *MagickWand) SetOrientation(orientation I.OrientationType) bool {
	return SetOrientation(m, orientation)
}

var SetPage func(m *MagickWand, width, height uint32, x, y int32) bool

func (m *MagickWand) SetPage(width, height uint32, x, y int32) bool {
	return SetPage(m, width, height, x, y)
}

var SetPassphrase func(m *MagickWand, passphrase string) bool

func (m *MagickWand) SetPassphrase(passphrase string) bool { return SetPassphrase(m, passphrase) }

var SetPointsize func(m *MagickWand, pointsize float64) bool

func (m *MagickWand) SetPointsize(pointsize float64) bool { return SetPointsize(m, pointsize) }

var SetProgressMonitor func(m *MagickWand, progressMonitor, clientData *Void) I.MagickProgressMonitor

func (m *MagickWand) SetProgressMonitor(progressMonitor, clientData *Void) I.MagickProgressMonitor {
	return SetProgressMonitor(m, progressMonitor, clientData)
}

var SetResolution func(m *MagickWand, xResolution, yResolution float64) bool

func (m *MagickWand) SetResolution(xResolution, yResolution float64) bool {
	return SetResolution(m, xResolution, yResolution)
}

var SetSamplingFactors func(m *MagickWand, numberFactors uint32, samplingFactors *float64) bool

func (m *MagickWand) SetSamplingFactors(numberFactors uint32, samplingFactors *float64) bool {
	return SetSamplingFactors(m, numberFactors, samplingFactors)
}

var SetSize func(m *MagickWand, columns, rows uint32) bool

func (m *MagickWand) SetSize(columns, rows uint32) bool { return SetSize(m, columns, rows) }

var SetSizeOffset func(m *MagickWand, columns, rows uint32, offset int32) bool

func (m *MagickWand) SetSizeOffset(columns, rows uint32, offset int32) bool {
	return SetSizeOffset(m, columns, rows, offset)
}

var SetType func(m *MagickWand, imageType I.ImageType) bool

func (m *MagickWand) SetType(imageType I.ImageType) bool { return SetType(m, imageType) }

var ShadeImage func(m *MagickWand, gray bool, azimuth, elevation float64) bool

func (m *MagickWand) Shade(gray bool, azimuth, elevation float64) bool {
	return ShadeImage(m, gray, azimuth, elevation)
}

var ShadowImage func(m *MagickWand, opacity, sigma float64, x, y int32) bool

func (m *MagickWand) Shadow(opacity, sigma float64, x, y int32) bool {
	return ShadowImage(m, opacity, sigma, x, y)
}

var SharpenImage func(m *MagickWand, radius, sigma float64) bool

func (m *MagickWand) Sharpen(radius, sigma float64) bool {
	return SharpenImage(m, radius, sigma)
}

var SharpenImageChannel func(m *MagickWand, channel I.ChannelType, radius, sigma float64) bool

func (m *MagickWand) SharpenChannel(channel I.ChannelType, radius, sigma float64) bool {
	return SharpenImageChannel(m, channel, radius, sigma)
}

var ShaveImage func(m *MagickWand, columns, rows uint32) bool

func (m *MagickWand) Shave(columns, rows uint32) bool { return ShaveImage(m, columns, rows) }

var ShearImage func(m *MagickWand, background *PixelWand, xShear, yShear float64) bool

func (m *MagickWand) Shear(background *PixelWand, xShear, yShear float64) bool {
	return ShearImage(m, background, xShear, yShear)
}

var SigmoidalContrastImage func(m *MagickWand, sharpen bool, alpha, beta float64) bool

func (m *MagickWand) SigmoidalContrastImage(sharpen bool, alpha, beta float64) bool {
	return SigmoidalContrastImage(m, sharpen, alpha, beta)
}

var SigmoidalContrastImageChannel func(m *MagickWand, channel I.ChannelType, sharpen bool, alpha, beta float64) bool

func (m *MagickWand) SigmoidalContrastImageChannel(channel I.ChannelType, sharpen bool, alpha, beta float64) bool {
	return SigmoidalContrastImageChannel(m, channel, sharpen, alpha, beta)
}

var SimilarityImage func(m *MagickWand, reference *MagickWand, offset *I.RectangleInfo, similarity *float64) *MagickWand

func (m *MagickWand) SimilarityImage(reference *MagickWand, offset *I.RectangleInfo, similarity *float64) *MagickWand {
	return SimilarityImage(m, reference, offset, similarity)
}

var SketchImage func(m *MagickWand, radius, sigma, angle float64) bool

func (m *MagickWand) SketchImage(radius, sigma, angle float64) bool {
	return SketchImage(m, radius, sigma, angle)
}

var SmushImages func(m *MagickWand, stack bool, offset int32) *MagickWand

func (m *MagickWand) SmushImages(stack bool, offset int32) *MagickWand {
	return SmushImages(m, stack, offset)
}

var SolarizeImage func(m *MagickWand, threshold float64) bool

func (m *MagickWand) Solarize(threshold float64) bool { return SolarizeImage(m, threshold) }

var SparseColorImage func(m *MagickWand, channel I.ChannelType, method I.SparseColorMethod, numberArguments uint32, arguments *float64) bool

func (m *MagickWand) SparseColor(channel I.ChannelType, method I.SparseColorMethod, numberArguments uint32, arguments *float64) bool {
	return SparseColorImage(m, channel, method, numberArguments, arguments)
}

var SpliceImage func(m *MagickWand, width, height uint32, x, y int32) bool

func (m *MagickWand) Splice(width, height uint32, x, y int32) bool {
	return SpliceImage(m, width, height, x, y)
}

var SpreadImage func(m *MagickWand, radius float64) bool

func (m *MagickWand) Spread(radius float64) bool { return SpreadImage(m, radius) }

var StatisticImage func(m *MagickWand, type_ I.StatisticType, width float64, height uint32) bool

func (m *MagickWand) StatisticImage(type_ I.StatisticType, width float64, height uint32) bool {
	return StatisticImage(m, type_, width, height)
}

var SteganoImage func(m *MagickWand, watermarkWand, offset int32) *MagickWand

func (m *MagickWand) Stegano(watermarkWand, offset int32) *MagickWand {
	return SteganoImage(m, watermarkWand, offset)
}

var StereoImage func(m *MagickWand, offsetWand *MagickWand) *MagickWand

func (m *MagickWand) Stereo(offsetWand *MagickWand) *MagickWand {
	return StereoImage(m, offsetWand)
}

var StripImage func(m *MagickWand) bool

func (m *MagickWand) Strip() bool { return StripImage(m) }

var SwirlImage func(m *MagickWand, degrees float64) bool

func (m *MagickWand) Swirl(degrees float64) bool { return SwirlImage(m, degrees) }

var TextureImage func(m *MagickWand, textureWand *MagickWand) *MagickWand

func (m *MagickWand) Texture(textureWand *MagickWand) *MagickWand {
	return TextureImage(m, textureWand)
}

var ThresholdImage func(m *MagickWand, threshold float64) bool

func (m *MagickWand) Threshold(threshold float64) bool {
	return ThresholdImage(m, threshold)
}

var ThresholdImageChannel func(m *MagickWand, channel I.ChannelType, threshold float64) bool

func (m *MagickWand) ThresholdChannel(channel I.ChannelType, threshold float64) bool {
	return ThresholdImageChannel(m, channel, threshold)
}

var ThumbnailImage func(m *MagickWand, columns, rows uint32) bool

func (m *MagickWand) Thumbnail(columns, rows uint32) bool {
	return ThumbnailImage(m, columns, rows)
}

var TintImage func(m *MagickWand, tint, opacity *PixelWand) bool

func (m *MagickWand) Tint(tint, opacity *PixelWand) bool {
	return TintImage(m, tint, opacity)
}

var TransformImage func(m *MagickWand, crop, geometry string) *MagickWand

func (m *MagickWand) Transform(crop, geometry string) *MagickWand {
	return TransformImage(m, crop, geometry)
}

var TransformImageColorspace func(m *MagickWand, colorspace I.ColorspaceType) bool

func (m *MagickWand) TransformColorspace(colorspace I.ColorspaceType) bool {
	return TransformImageColorspace(m, colorspace)
}

var TransparentPaintImage func(m *MagickWand, target *PixelWand, alpha, fuzz float64, invert bool) bool

func (m *MagickWand) TransparentPaint(target *PixelWand, alpha, fuzz float64, invert bool) bool {
	return TransparentPaintImage(m, target, alpha, fuzz, invert)
}

var TransposeImage func(m *MagickWand) bool

func (m *MagickWand) Transpose() bool { return TransposeImage(m) }

var TransverseImage func(m *MagickWand) bool

var TrimImage func(m *MagickWand, fuzz float64) bool

func (m *MagickWand) Trim(fuzz float64) bool { return TrimImage(m, fuzz) }

var UniqueImageColors func(m *MagickWand) bool

func (m *MagickWand) UniqueColors() bool { return UniqueImageColors(m) }

var UnsharpMaskImage func(m *MagickWand, radius, sigma, amount, threshold float64) bool

func (m *MagickWand) UnsharpMask(radius, sigma, amount, threshold float64) bool {
	return UnsharpMaskImage(m, radius, sigma, amount, threshold)
}

var UnsharpMaskImageChannel func(m *MagickWand, channel I.ChannelType, radius, sigma, amount, threshold float64) bool

func (m *MagickWand) UnsharpMaskChannel(channel I.ChannelType, radius, sigma, amount, threshold float64) bool {
	return UnsharpMaskImageChannel(m, channel, radius, sigma, amount, threshold)
}

var VignetteImage func(m *MagickWand, blackPoint, whitePoint float64, x, y int32) bool

func (m *MagickWand) Vignette(blackPoint, whitePoint float64, x, y int32) bool {
	return VignetteImage(m, blackPoint, whitePoint, x, y)
}

var WaveImage func(m *MagickWand, amplitude, waveLength float64) bool

func (m *MagickWand) Wave(amplitude, waveLength float64) bool {
	return WaveImage(m, amplitude, waveLength)
}

var WhiteThresholdImage func(m *MagickWand, threshold *PixelWand) bool

func (m *MagickWand) WhiteThreshold(threshold *PixelWand) bool {
	return WhiteThresholdImage(m, threshold)
}

var WriteImage func(m *MagickWand, filename string) bool

func (m *MagickWand) Write(filename string) bool { return WriteImage(m, filename) }

var WriteImageFile func(m *MagickWand, file *FILE) bool

func (m *MagickWand) WriteFile(file *FILE) bool { return WriteImageFile(m, file) }

var WriteImages func(m *MagickWand, filename string, adjoin bool) bool

func (m *MagickWand) WriteImages(filename string, adjoin bool) bool {
	return WriteImages(m, filename, adjoin)
}

var WriteImagesFile func(m *MagickWand, file *FILE) bool

func (m *MagickWand) WriteImagesFile(file *FILE) bool { return WriteImagesFile(m, file) }

var NewDrawingWand func() *DrawingWand

var ClearDrawingWand func(d *DrawingWand)

func (d *DrawingWand) Clear() { ClearDrawingWand(d) }

var CloneDrawingWand func(d *DrawingWand) *DrawingWand

func (d *DrawingWand) Clone() *DrawingWand { return CloneDrawingWand(d) }

var DestroyDrawingWand func(d *DrawingWand) *DrawingWand

func (d *DrawingWand) Destroy() *DrawingWand { return DestroyDrawingWand(d) }

var IsDrawingWand func(d *DrawingWand) bool

func (d *DrawingWand) IsDrawingWand() bool { return IsDrawingWand(d) }

var PeekDrawingWand func(d *DrawingWand) *DrawInfo

func (d *DrawingWand) Peek() *DrawInfo { return PeekDrawingWand(d) }

var PopDrawingWand func(d *DrawingWand) bool

func (d *DrawingWand) Pop() bool { return PopDrawingWand(d) }

var PushDrawingWand func(d *DrawingWand) bool

func (d *DrawingWand) Push() bool { return PushDrawingWand(d) }

var DrawAffine func(d *DrawingWand, affine *I.AffineMatrix)

func (d *DrawingWand) Affine(affine *I.AffineMatrix) { DrawAffine(d, affine) }

var DrawAnnotation func(d *DrawingWand, x, y float64, text *uint8)

func (d *DrawingWand) Annotation(x, y float64, text *uint8) { DrawAnnotation(d, x, y, text) }

var DrawArc func(d *DrawingWand, sx, sy, ex, ey, sd, ed float64)

func (d *DrawingWand) Arc(sx, sy, ex, ey, sd, ed float64) { DrawArc(d, sx, sy, ex, ey, sd, ed) }

var DrawBezier func(d *DrawingWand, numberCoordinates uint32, coordinates *I.PointInfo)

func (d *DrawingWand) Bezier(numberCoordinates uint32, coordinates *I.PointInfo) {
	DrawBezier(d, numberCoordinates, coordinates)
}

var DrawCircle func(d *DrawingWand, ox, oy, px, py float64)

func (d *DrawingWand) Circle(ox, oy, px, py float64) { DrawCircle(d, ox, oy, px, py) }

var DrawClearException func(d *DrawingWand) bool

func (d *DrawingWand) ClearException() bool { return DrawClearException(d) }

var DrawColor func(d *DrawingWand, x, y float64, paintMethod I.PaintMethod)

func (d *DrawingWand) Color(x, y float64, paintMethod I.PaintMethod) {
	DrawColor(d, x, y, paintMethod)
}

var DrawComment func(d *DrawingWand, comment string)

func (d *DrawingWand) Comment(comment string) { DrawComment(d, comment) }

var DrawComposite func(d *DrawingWand, compose I.CompositeOperator, x, y, width, height float64, magickWand *MagickWand) bool

func (d *DrawingWand) Composite(compose I.CompositeOperator, x, y, width, height float64, magickWand *MagickWand) bool {
	return DrawComposite(d, compose, x, y, width, height, magickWand)
}

var DrawEllipse func(d *DrawingWand, ox, oy, rx, ry, start, end float64)

func (d *DrawingWand) Ellipse(ox, oy, rx, ry, start, end float64) {
	DrawEllipse(d, ox, oy, rx, ry, start, end)
}

var DrawGetBorderColor func(d *DrawingWand, borderColor *PixelWand)

func (d *DrawingWand) BorderColor(borderColor *PixelWand) { DrawGetBorderColor(d, borderColor) }

var DrawGetClipPath func(d *DrawingWand) string

func (d *DrawingWand) ClipPath() string { return DrawGetClipPath(d) }

var DrawGetClipUnits func(d *DrawingWand) I.ClipPathUnits

func (d *DrawingWand) ClipUnits() I.ClipPathUnits { return DrawGetClipUnits(d) }

var DrawGetException func(d *DrawingWand, severity *I.ExceptionType) string

func (d *DrawingWand) Exception(severity *I.ExceptionType) string {
	return DrawGetException(d, severity)
}

var DrawGetExceptionType func(d *DrawingWand) I.ExceptionType

func (d *DrawingWand) ExceptionType() I.ExceptionType { return DrawGetExceptionType(d) }

var DrawGetFillColor func(d *DrawingWand, fillColor *PixelWand)

func (d *DrawingWand) FillColor(fillColor *PixelWand) { DrawGetFillColor(d, fillColor) }

var DrawGetFillOpacity func(d *DrawingWand) float64

func (d *DrawingWand) FillOpacity() float64 { return DrawGetFillOpacity(d) }

var DrawGetFillRule func(d *DrawingWand) I.FillRule

func (d *DrawingWand) FillRule() I.FillRule { return DrawGetFillRule(d) }

var DrawGetFont func(d *DrawingWand) string

func (d *DrawingWand) Font() string { return DrawGetFont(d) }

var DrawGetFontFamily func(d *DrawingWand) string

func (d *DrawingWand) FontFamily() string { return DrawGetFontFamily(d) }

var DrawGetFontResolution func(d *DrawingWand, x, y *float64) bool

func (d *DrawingWand) FontResolution(x, y *float64) bool {
	return DrawGetFontResolution(d, x, y)
}

var DrawGetFontSize func(d *DrawingWand) float64

func (d *DrawingWand) FontSize() float64 { return DrawGetFontSize(d) }

var DrawGetFontStretch func(d *DrawingWand) I.StretchType

func (d *DrawingWand) FontStretch() I.StretchType { return DrawGetFontStretch(d) }

var DrawGetFontStyle func(d *DrawingWand) I.StyleType

func (d *DrawingWand) FontStyle() I.StyleType { return DrawGetFontStyle(d) }

var DrawGetFontWeight func(d *DrawingWand) uint32

func (d *DrawingWand) FontWeight() uint32 { return DrawGetFontWeight(d) }

var DrawGetGravity func(d *DrawingWand) I.GravityType

func (d *DrawingWand) Gravity() I.GravityType { return DrawGetGravity(d) }

var DrawGetOpacity func(d *DrawingWand) float64

func (d *DrawingWand) Opacity() float64 { return DrawGetOpacity(d) }

var DrawGetStrokeAntialias func(d *DrawingWand) bool

func (d *DrawingWand) StrokeAntialias() bool { return DrawGetStrokeAntialias(d) }

var DrawGetStrokeColor func(d *DrawingWand, strokeColor *PixelWand)

func (d *DrawingWand) StrokeColor(strokeColor *PixelWand) { DrawGetStrokeColor(d, strokeColor) }

var DrawGetStrokeDashArray func(d *DrawingWand, numberElements *uint32) *float64

func (d *DrawingWand) StrokeDashArray(numberElements *uint32) *float64 {
	return DrawGetStrokeDashArray(d, numberElements)
}

var DrawGetStrokeDashOffset func(d *DrawingWand) float64

func (d *DrawingWand) StrokeDashOffset() float64 { return DrawGetStrokeDashOffset(d) }

var DrawGetStrokeLineCap func(d *DrawingWand) I.LineCap

func (d *DrawingWand) StrokeLineCap() I.LineCap { return DrawGetStrokeLineCap(d) }

var DrawGetStrokeLineJoin func(d *DrawingWand) I.LineJoin

func (d *DrawingWand) StrokeLineJoin() I.LineJoin { return DrawGetStrokeLineJoin(d) }

var DrawGetStrokeMiterLimit func(d *DrawingWand) uint32

func (d *DrawingWand) StrokeMiterLimit() uint32 { return DrawGetStrokeMiterLimit(d) }

var DrawGetStrokeOpacity func(d *DrawingWand) float64

func (d *DrawingWand) StrokeOpacity() float64 { return DrawGetStrokeOpacity(d) }

var DrawGetStrokeWidth func(d *DrawingWand) float64

func (d *DrawingWand) StrokeWidth() float64 { return DrawGetStrokeWidth(d) }

var DrawGetTextAlignment func(d *DrawingWand) I.AlignType

func (d *DrawingWand) TextAlignment() I.AlignType { return DrawGetTextAlignment(d) }

var DrawGetTextAntialias func(d *DrawingWand) bool

func (d *DrawingWand) TextAntialias() bool { return DrawGetTextAntialias(d) }

var DrawGetTextDecoration func(d *DrawingWand) I.DecorationType

func (d *DrawingWand) TextDecoration() I.DecorationType { return DrawGetTextDecoration(d) }

var DrawGetTextEncoding func(d *DrawingWand) string

func (d *DrawingWand) TextEncoding() string { return DrawGetTextEncoding(d) }

var DrawGetTextInterwordSpacing func(d *DrawingWand) float64

func (d *DrawingWand) TextInterwordSpacing() float64 { return DrawGetTextInterwordSpacing(d) }

var DrawGetTextKerning func(d *DrawingWand) float64

func (d *DrawingWand) TextKerning() float64 { return DrawGetTextKerning(d) }

var DrawGetTextUnderColor func(d *DrawingWand, underColor *PixelWand)

func (d *DrawingWand) TextUnderColor(underColor *PixelWand) { DrawGetTextUnderColor(d, underColor) }

var DrawGetVectorGraphics func(d *DrawingWand) string

func (d *DrawingWand) VectorGraphics() string { return DrawGetVectorGraphics(d) }

var DrawLine func(d *DrawingWand, sx, sy, ex, ey float64)

func (d *DrawingWand) Line(sx, sy, ex, ey float64) { DrawLine(d, sx, sy, ex, ey) }

var DrawMatte func(d *DrawingWand, x, y float64, paintMethod I.PaintMethod)

func (d *DrawingWand) Matte(x, y float64, paintMethod I.PaintMethod) {
	DrawMatte(d, x, y, paintMethod)
}

var DrawPathClose func(d *DrawingWand)

func (d *DrawingWand) PathClose() { DrawPathClose(d) }

var DrawPathCurveToAbsolute func(d *DrawingWand, x1, y1, x2, y2, x, y float64)

func (d *DrawingWand) PathCurveToAbsolute(x1, y1, x2, y2, x, y float64) {
	DrawPathCurveToAbsolute(d, x1, y1, x2, y2, x, y)
}

var DrawPathCurveToQuadraticBezierAbsolute func(d *DrawingWand, x1, y1, x, y float64)

func (d *DrawingWand) PathCurveToQuadraticBezierAbsolute(x1, y1, x, y float64) {
	DrawPathCurveToQuadraticBezierAbsolute(d, x1, y1, x, y)
}

var DrawPathCurveToQuadraticBezierRelative func(d *DrawingWand, x1, y1, x, y float64)

func (d *DrawingWand) PathCurveToQuadraticBezierRelative(x1, y1, x, y float64) {
	DrawPathCurveToQuadraticBezierRelative(d, x1, y1, x, y)
}

var DrawPathCurveToQuadraticBezierSmoothAbsolute func(d *DrawingWand, x, y float64)

func (d *DrawingWand) PathCurveToQuadraticBezierSmoothAbsolute(x, y float64) {
	DrawPathCurveToQuadraticBezierSmoothAbsolute(d, x, y)
}

var DrawPathCurveToQuadraticBezierSmoothRelative func(d *DrawingWand, x, y float64)

func (d *DrawingWand) PathCurveToQuadraticBezierSmoothRelative(x, y float64) {
	DrawPathCurveToQuadraticBezierSmoothRelative(d, x, y)
}

var DrawPathCurveToRelative func(d *DrawingWand, x1, y1, x2, y2, x, y float64)

func (d *DrawingWand) PathCurveToRelative(x1, y1, x2, y2, x, y float64) {
	DrawPathCurveToRelative(d, x1, y1, x2, y2, x, y)
}

var DrawPathCurveToSmoothAbsolute func(d *DrawingWand, x2, y2, x, y float64)

func (d *DrawingWand) PathCurveToSmoothAbsolute(x2, y2, x, y float64) {
	DrawPathCurveToSmoothAbsolute(d, x2, y2, x, y)
}

var DrawPathCurveToSmoothRelative func(d *DrawingWand, x2, y2, x, y float64)

func (d *DrawingWand) PathCurveToSmoothRelative(x2, y2, x, y float64) {
	DrawPathCurveToSmoothRelative(d, x2, y2, x, y)
}

var DrawPathEllipticArcAbsolute func(d *DrawingWand, rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y float64)

func (d *DrawingWand) PathEllipticArcAbsolute(rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y float64) {
	DrawPathEllipticArcAbsolute(d, rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

var DrawPathEllipticArcRelative func(d *DrawingWand, rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y float64)

func (d *DrawingWand) PathEllipticArcRelative(rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y float64) {
	DrawPathEllipticArcRelative(d, rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

var DrawPathFinish func(d *DrawingWand)

func (d *DrawingWand) PathFinish() { DrawPathFinish(d) }

var DrawPathLineToAbsolute func(d *DrawingWand, x, y float64)

func (d *DrawingWand) PathLineToAbsolute(x, y float64) { DrawPathLineToAbsolute(d, x, y) }

var DrawPathLineToHorizontalAbsolute func(d *DrawingWand, x float64)

func (d *DrawingWand) PathLineToHorizontalAbsolute(x float64) {
	DrawPathLineToHorizontalAbsolute(d, x)
}

var DrawPathLineToHorizontalRelative func(d *DrawingWand, x float64)

func (d *DrawingWand) PathLineToHorizontalRelative(x float64) {
	DrawPathLineToHorizontalRelative(d, x)
}

var DrawPathLineToRelative func(d *DrawingWand, x, y float64)

func (d *DrawingWand) PathLineToRelative(x, y float64) { DrawPathLineToRelative(d, x, y) }

var DrawPathLineToVerticalAbsolute func(d *DrawingWand, y float64)

func (d *DrawingWand) PathLineToVerticalAbsolute(y float64) { DrawPathLineToVerticalAbsolute(d, y) }

var DrawPathLineToVerticalRelative func(d *DrawingWand, y float64)

func (d *DrawingWand) PathLineToVerticalRelative(y float64) { DrawPathLineToVerticalRelative(d, y) }

var DrawPathMoveToAbsolute func(d *DrawingWand, x, y float64)

func (d *DrawingWand) PathMoveToAbsolute(x, y float64) { DrawPathMoveToAbsolute(d, x, y) }

var DrawPathMoveToRelative func(d *DrawingWand, x, y float64)

func (d *DrawingWand) PathMoveToRelative(x, y float64) { DrawPathMoveToRelative(d, x, y) }

var DrawPathStart func(d *DrawingWand)

func (d *DrawingWand) PathStart() { DrawPathStart(d) }

var DrawPoint func(d *DrawingWand, x, y float64)

func (d *DrawingWand) Point(x, y float64) { DrawPoint(d, x, y) }

var DrawPolygon func(d *DrawingWand, numberCoordinates uint32, coordinates *I.PointInfo)

func (d *DrawingWand) Polygon(numberCoordinates uint32, coordinates *I.PointInfo) {
	DrawPolygon(d, numberCoordinates, coordinates)
}

var DrawPolyline func(d *DrawingWand, numberCoordinates uint32, coordinates *I.PointInfo)

func (d *DrawingWand) Polyline(numberCoordinates uint32, coordinates *I.PointInfo) {
	DrawPolyline(d, numberCoordinates, coordinates)
}

var DrawPopClipPath func(d *DrawingWand)

func (d *DrawingWand) PopClipPath() { DrawPopClipPath(d) }

var DrawPopDefs func(d *DrawingWand)

func (d *DrawingWand) PopDefs() { DrawPopDefs(d) }

var DrawPopPattern func(d *DrawingWand) bool

func (d *DrawingWand) PopPattern() bool { return DrawPopPattern(d) }

var DrawPushClipPath func(d *DrawingWand, clipMaskId string)

func (d *DrawingWand) PushClipPath(clipMaskId string) { DrawPushClipPath(d, clipMaskId) }

var DrawPushDefs func(d *DrawingWand)

func (d *DrawingWand) PushDefs() { DrawPushDefs(d) }

var DrawPushPattern func(d *DrawingWand, patternId string, x, y, width, height float64) bool

func (d *DrawingWand) PushPattern(patternId string, x, y, width, height float64) bool {
	return DrawPushPattern(d, patternId, x, y, width, height)
}

var DrawRectangle func(d *DrawingWand, x1, y1, x2, y2 float64)

func (d *DrawingWand) Rectangle(x1, y1, x2, y2 float64) { DrawRectangle(d, x1, y1, x2, y2) }

var DrawResetVectorGraphics func(d *DrawingWand)

func (d *DrawingWand) ResetVectorGraphics() { DrawResetVectorGraphics(d) }

var DrawRotate func(d *DrawingWand, degrees float64)

func (d *DrawingWand) Rotate(degrees float64) { DrawRotate(d, degrees) }

var DrawRoundRectangle func(d *DrawingWand, x1, y1, x2, y2, rx, ry float64)

func (d *DrawingWand) RoundRectangle(x1, y1, x2, y2, rx, ry float64) {
	DrawRoundRectangle(d, x1, y1, x2, y2, rx, ry)
}

var DrawScale func(d *DrawingWand, x, y float64)

func (d *DrawingWand) Scale(x, y float64) { DrawScale(d, x, y) }

var DrawSetBorderColor func(d *DrawingWand, borderWand *PixelWand)

func (d *DrawingWand) SetBorderColor(borderWand *PixelWand) { DrawSetBorderColor(d, borderWand) }

var DrawSetClipPath func(d *DrawingWand, clipMask string) bool

func (d *DrawingWand) SetClipPath(clipMask string) bool { return DrawSetClipPath(d, clipMask) }

var DrawSetClipRule func(d *DrawingWand, fillRule I.FillRule)

func (d *DrawingWand) SetClipRule(fillRule I.FillRule) { DrawSetClipRule(d, fillRule) }

var DrawSetClipUnits func(d *DrawingWand, clipUnits I.ClipPathUnits)

func (d *DrawingWand) SetClipUnits(clipUnits I.ClipPathUnits) { DrawSetClipUnits(d, clipUnits) }

var DrawSetFillColor func(d *DrawingWand, fillWand *PixelWand)

func (d *DrawingWand) SetFillColor(fillWand *PixelWand) { DrawSetFillColor(d, fillWand) }

var DrawSetFillOpacity func(d *DrawingWand, fillOpacity float64)

func (d *DrawingWand) SetFillOpacity(fillOpacity float64) { DrawSetFillOpacity(d, fillOpacity) }

var DrawSetFillPatternURL func(d *DrawingWand, fillUrl string) bool

func (d *DrawingWand) SetFillPatternURL(fillUrl string) bool {
	return DrawSetFillPatternURL(d, fillUrl)
}

var DrawSetFillRule func(d *DrawingWand, fillRule I.FillRule)

func (d *DrawingWand) SetFillRule(fillRule I.FillRule) { DrawSetFillRule(d, fillRule) }

var DrawSetFont func(d *DrawingWand, fontName string) bool

func (d *DrawingWand) SetFont(fontName string) bool { return DrawSetFont(d, fontName) }

var DrawSetFontFamily func(d *DrawingWand, fontFamily string) bool

func (d *DrawingWand) SetFontFamily(fontFamily string) bool {
	return DrawSetFontFamily(d, fontFamily)
}

var DrawSetFontResolution func(d *DrawingWand, xResolution, yResolution float64) bool

func (d *DrawingWand) SetFontResolution(xResolution, yResolution float64) bool {
	return DrawSetFontResolution(d, xResolution, yResolution)
}

var DrawSetFontSize func(d *DrawingWand, pointsize float64)

func (d *DrawingWand) SetFontSize(pointsize float64) { DrawSetFontSize(d, pointsize) }

var DrawSetFontStretch func(d *DrawingWand, fontStretch I.StretchType)

func (d *DrawingWand) SetFontStretch(fontStretch I.StretchType) { DrawSetFontStretch(d, fontStretch) }

var DrawSetFontStyle func(d *DrawingWand, style I.StyleType)

func (d *DrawingWand) SetFontStyle(style I.StyleType) { DrawSetFontStyle(d, style) }

var DrawSetFontWeight func(d *DrawingWand, fontWeight uint32)

func (d *DrawingWand) SetFontWeight(fontWeight uint32) { DrawSetFontWeight(d, fontWeight) }

var DrawSetGravity func(d *DrawingWand, gravity I.GravityType)

func (d *DrawingWand) SetGravity(gravity I.GravityType) { DrawSetGravity(d, gravity) }

var DrawSetOpacity func(d *DrawingWand, opacity float64)

func (d *DrawingWand) SetOpacity(opacity float64) { DrawSetOpacity(d, opacity) }

var DrawSetStrokeAntialias func(d *DrawingWand, strokeAntialias bool)

func (d *DrawingWand) SetStrokeAntialias(strokeAntialias bool) {
	DrawSetStrokeAntialias(d, strokeAntialias)
}

var DrawSetStrokeColor func(d *DrawingWand, strokeWand *PixelWand)

func (d *DrawingWand) SetStrokeColor(strokeWand *PixelWand) { DrawSetStrokeColor(d, strokeWand) }

var DrawSetStrokeDashArray func(d *DrawingWand, numberElements uint32, dashArray *float64) bool

func (d *DrawingWand) SetStrokeDashArray(numberElements uint32, dashArray *float64) bool {
	return DrawSetStrokeDashArray(d, numberElements, dashArray)
}

var DrawSetStrokeDashOffset func(d *DrawingWand, dashOffset float64)

func (d *DrawingWand) SetStrokeDashOffset(dashOffset float64) {
	DrawSetStrokeDashOffset(d, dashOffset)
}

var DrawSetStrokeLineCap func(d *DrawingWand, linecap I.LineCap)

func (d *DrawingWand) SetStrokeLineCap(linecap I.LineCap) { DrawSetStrokeLineCap(d, linecap) }

var DrawSetStrokeLineJoin func(d *DrawingWand, linejoin I.LineJoin)

func (d *DrawingWand) SetStrokeLineJoin(linejoin I.LineJoin) { DrawSetStrokeLineJoin(d, linejoin) }

var DrawSetStrokeMiterLimit func(d *DrawingWand, miterlimit uint32)

func (d *DrawingWand) SetStrokeMiterLimit(miterlimit uint32) { DrawSetStrokeMiterLimit(d, miterlimit) }

var DrawSetStrokeOpacity func(d *DrawingWand, strokeOpacity float64)

func (d *DrawingWand) SetStrokeOpacity(strokeOpacity float64) {
	DrawSetStrokeOpacity(d, strokeOpacity)
}

var DrawSetStrokePatternURL func(d *DrawingWand, strokeUrl string) bool

func (d *DrawingWand) SetStrokePatternURL(strokeUrl string) bool {
	return DrawSetStrokePatternURL(d, strokeUrl)
}

var DrawSetStrokeWidth func(d *DrawingWand, strokeWidth float64)

func (d *DrawingWand) SetStrokeWidth(strokeWidth float64) { DrawSetStrokeWidth(d, strokeWidth) }

var DrawSetTextAlignment func(d *DrawingWand, alignment I.AlignType)

func (d *DrawingWand) SetTextAlignment(alignment I.AlignType) { DrawSetTextAlignment(d, alignment) }

var DrawSetTextAntialias func(d *DrawingWand, textAntialias bool)

func (d *DrawingWand) SetTextAntialias(textAntialias bool) { DrawSetTextAntialias(d, textAntialias) }

var DrawSetTextDecoration func(d *DrawingWand, decoration I.DecorationType)

func (d *DrawingWand) SetTextDecoration(decoration I.DecorationType) {
	DrawSetTextDecoration(d, decoration)
}

var DrawSetTextEncoding func(d *DrawingWand, encoding string)

func (d *DrawingWand) SetTextEncoding(encoding string) { DrawSetTextEncoding(d, encoding) }

var DrawSetTextInterlineSpacing func(d *DrawingWand, interlineSpacing float64)

func (d *DrawingWand) SetTextInterlineSpacing(interlineSpacing float64) {
	DrawSetTextInterlineSpacing(d, interlineSpacing)
}

var DrawSetTextInterwordSpacing func(d *DrawingWand, interwordSpacing float64)

func (d *DrawingWand) SetTextInterwordSpacing(interwordSpacing float64) {
	DrawSetTextInterwordSpacing(d, interwordSpacing)
}

var DrawSetTextKerning func(d *DrawingWand, kerning float64)

func (d *DrawingWand) SetTextKerning(kerning float64) { DrawSetTextKerning(d, kerning) }

var DrawSetTextUnderColor func(d *DrawingWand, underWand *PixelWand)

func (d *DrawingWand) SetTextUnderColor(underWand *PixelWand) { DrawSetTextUnderColor(d, underWand) }

var DrawSetVectorGraphics func(d *DrawingWand, xml string) bool

func (d *DrawingWand) SetVectorGraphics(xml string) bool { return DrawSetVectorGraphics(d, xml) }

var DrawSetViewbox func(d *DrawingWand, x1, y1, x2, y2 int32)

func (d *DrawingWand) SetViewbox(x1, y1, x2, y2 int32) { DrawSetViewbox(d, x1, y1, x2, y2) }

var DrawSkewX func(d *DrawingWand, degrees float64)

func (d *DrawingWand) SkewX(degrees float64) { DrawSkewX(d, degrees) }

var DrawSkewY func(d *DrawingWand, degrees float64)

func (d *DrawingWand) SkewY(degrees float64) { DrawSkewY(d, degrees) }

var DrawTranslate func(d *DrawingWand, x, y float64)

func (d *DrawingWand) Translate(x, y float64) { DrawTranslate(d, x, y) }

var NewPixelWand func() *PixelWand

var NewPixelWands func(numberWands uint32) []*PixelWand

var ClonePixelWands func(wands []*PixelWand, numberWands uint32) []*PixelWand

var DestroyPixelWands func(wands []*PixelWand, numberWands uint32) []*PixelWand

var ClearPixelWand func(p *PixelWand)

func (p *PixelWand) Clear() { ClearPixelWand(p) }

var ClonePixelWand func(p *PixelWand) *PixelWand

func (p *PixelWand) Clone() *PixelWand { return ClonePixelWand(p) }

var DestroyPixelWand func(p *PixelWand) *PixelWand

func (p *PixelWand) Destroy() *PixelWand { return DestroyPixelWand(p) }

var IsPixelWand func(p *PixelWand) bool

func (p *PixelWand) IsPixelWand() bool { return IsPixelWand(p) }

var IsPixelWandSimilar func(p *PixelWand, q *PixelWand, fuzz float64) bool

func (p *PixelWand) IsSimilar(q *PixelWand, fuzz float64) bool {
	return IsPixelWandSimilar(p, q, fuzz)
}

var PixelClearException func(p *PixelWand) bool

func (p *PixelWand) ClearException() bool { return PixelClearException(p) }

var PixelGetAlpha func(p *PixelWand) float64

func (p *PixelWand) Alpha() float64 { return PixelGetAlpha(p) }

var PixelGetAlphaQuantum func(p *PixelWand) I.Quantum

func (p *PixelWand) AlphaQuantum() I.Quantum { return PixelGetAlphaQuantum(p) }

var PixelGetBlack func(p *PixelWand) float64

func (p *PixelWand) Black() float64 { return PixelGetBlack(p) }

var PixelGetBlackQuantum func(p *PixelWand) I.Quantum

func (p *PixelWand) BlackQuantum() I.Quantum { return PixelGetBlackQuantum(p) }

var PixelGetBlue func(p *PixelWand) float64

func (p *PixelWand) Blue() float64 { return PixelGetBlue(p) }

var PixelGetBlueQuantum func(p *PixelWand) I.Quantum

func (p *PixelWand) BlueQuantum() I.Quantum { return PixelGetBlueQuantum(p) }

var PixelGetColorAsNormalizedString func(p *PixelWand) string

func (p *PixelWand) ColorAsNormalizedString() string { return PixelGetColorAsNormalizedString(p) }

var PixelGetColorAsString func(p *PixelWand) string

func (p *PixelWand) ColorAsString() string { return PixelGetColorAsString(p) }

var PixelGetColorCount func(p *PixelWand) uint32

func (p *PixelWand) ColorCount() uint32 { return PixelGetColorCount(p) }

var PixelGetCyan func(p *PixelWand) float64

func (p *PixelWand) Cyan() float64 { return PixelGetCyan(p) }

var PixelGetCyanQuantum func(p *PixelWand) I.Quantum

func (p *PixelWand) CyanQuantum() I.Quantum { return PixelGetCyanQuantum(p) }

var PixelGetException func(p *PixelWand, severity *I.ExceptionType) string

func (p *PixelWand) Exception(severity *I.ExceptionType) string { return PixelGetException(p, severity) }

var PixelGetExceptionType func(p *PixelWand) I.ExceptionType

func (p *PixelWand) ExceptionType() I.ExceptionType { return PixelGetExceptionType(p) }

var PixelGetFuzz func(p *PixelWand) float64

func (p *PixelWand) Fuzz() float64 { return PixelGetFuzz(p) }

var PixelGetGreen func(p *PixelWand) float64

func (p *PixelWand) Green() float64 { return PixelGetGreen(p) }

var PixelGetGreenQuantum func(p *PixelWand) I.Quantum

func (p *PixelWand) GreenQuantum() I.Quantum { return PixelGetGreenQuantum(p) }

var PixelGetHSL func(p *PixelWand, hue, saturation, lightness *float64)

func (p *PixelWand) HSL(hue, saturation, lightness *float64) {
	PixelGetHSL(p, hue, saturation, lightness)
}

var PixelGetIndex func(p *PixelWand) I.IndexPacket

func (p *PixelWand) Index() I.IndexPacket { return PixelGetIndex(p) }

var PixelGetMagenta func(p *PixelWand) float64

func (p *PixelWand) Magenta() float64 { return PixelGetMagenta(p) }

var PixelGetMagentaQuantum func(p *PixelWand) I.Quantum

func (p *PixelWand) MagentaQuantum() I.Quantum { return PixelGetMagentaQuantum(p) }

var PixelGetMagickColor func(p *PixelWand, color *MagickPixelPacket)

func (p *PixelWand) MagickColor(color *MagickPixelPacket) { PixelGetMagickColor(p, color) }

var PixelGetOpacity func(p *PixelWand) float64

func (p *PixelWand) Opacity() float64 { return PixelGetOpacity(p) }

var PixelGetOpacityQuantum func(p *PixelWand) I.Quantum

func (p *PixelWand) OpacityQuantum() I.Quantum { return PixelGetOpacityQuantum(p) }

var PixelGetQuantumColor func(p *PixelWand, color *I.PixelPacket)

func (p *PixelWand) QuantumColor(color *I.PixelPacket) { PixelGetQuantumColor(p, color) }

var PixelGetRed func(p *PixelWand) float64

func (p *PixelWand) Red() float64 { return PixelGetRed(p) }

var PixelGetRedQuantum func(p *PixelWand) I.Quantum

func (p *PixelWand) RedQuantum() I.Quantum { return PixelGetRedQuantum(p) }

var PixelGetYellow func(p *PixelWand) float64

func (p *PixelWand) Yellow() float64 { return PixelGetYellow(p) }

var PixelGetYellowQuantum func(p *PixelWand) I.Quantum

func (p *PixelWand) YellowQuantum() I.Quantum { return PixelGetYellowQuantum(p) }

var PixelSetAlpha func(p *PixelWand, alpha float64)

func (p *PixelWand) SetAlpha(alpha float64) { PixelSetAlpha(p, alpha) }

var PixelSetAlphaQuantum func(p *PixelWand, opacity I.Quantum)

func (p *PixelWand) SetAlphaQuantum(opacity I.Quantum) { PixelSetAlphaQuantum(p, opacity) }

var PixelSetBlack func(p *PixelWand, black float64)

func (p *PixelWand) SetBlack(black float64) { PixelSetBlack(p, black) }

var PixelSetBlackQuantum func(p *PixelWand, black I.Quantum)

func (p *PixelWand) SetBlackQuantum(black I.Quantum) { PixelSetBlackQuantum(p, black) }

var PixelSetBlue func(p *PixelWand, blue float64)

func (p *PixelWand) SetBlue(blue float64) { PixelSetBlue(p, blue) }

var PixelSetBlueQuantum func(p *PixelWand, blue I.Quantum)

func (p *PixelWand) SetBlueQuantum(blue I.Quantum) { PixelSetBlueQuantum(p, blue) }

var PixelSetColor func(p *PixelWand, color string) bool

func (p *PixelWand) SetColor(color string) bool { return PixelSetColor(p, color) }

var PixelSetColorCount func(p *PixelWand, count uint32)

func (p *PixelWand) SetColorCount(count uint32) { PixelSetColorCount(p, count) }

var PixelSetColorFromWand func(p *PixelWand, color *PixelWand)

func (p *PixelWand) SetColorFromWand(color *PixelWand) { PixelSetColorFromWand(p, color) }

var PixelSetCyan func(p *PixelWand, cyan float64)

func (p *PixelWand) SetCyan(cyan float64) { PixelSetCyan(p, cyan) }

var PixelSetCyanQuantum func(p *PixelWand, cyan I.Quantum)

func (p *PixelWand) SetCyanQuantum(cyan I.Quantum) { PixelSetCyanQuantum(p, cyan) }

var PixelSetFuzz func(p *PixelWand, fuzz float64)

func (p *PixelWand) SetFuzz(fuzz float64) { PixelSetFuzz(p, fuzz) }

var PixelSetGreen func(p *PixelWand, green float64)

func (p *PixelWand) SetGreen(green float64) { PixelSetGreen(p, green) }

var PixelSetGreenQuantum func(p *PixelWand, green I.Quantum)

func (p *PixelWand) SetGreenQuantum(green I.Quantum) { PixelSetGreenQuantum(p, green) }

var PixelSetHSL func(p *PixelWand, hue, saturation, lightness float64)

func (p *PixelWand) SetHSL(hue, saturation, lightness float64) {
	PixelSetHSL(p, hue, saturation, lightness)
}

var PixelSetIndex func(p *PixelWand, index I.IndexPacket)

func (p *PixelWand) SetIndex(index I.IndexPacket) { PixelSetIndex(p, index) }

var PixelSetMagenta func(p *PixelWand, magenta float64)

func (p *PixelWand) SetMagenta(magenta float64) { PixelSetMagenta(p, magenta) }

var PixelSetMagentaQuantum func(p *PixelWand, magenta I.Quantum)

func (p *PixelWand) SetMagentaQuantum(magenta I.Quantum) { PixelSetMagentaQuantum(p, magenta) }

var PixelSetMagickColor func(p *PixelWand, color *MagickPixelPacket)

func (p *PixelWand) SetMagickColor(color *MagickPixelPacket) { PixelSetMagickColor(p, color) }

var PixelSetOpacity func(p *PixelWand, opacity float64)

func (p *PixelWand) SetOpacity(opacity float64) { PixelSetOpacity(p, opacity) }

var PixelSetOpacityQuantum func(p *PixelWand, opacity I.Quantum)

func (p *PixelWand) SetOpacityQuantum(opacity I.Quantum) { PixelSetOpacityQuantum(p, opacity) }

var PixelSetQuantumColor func(p *PixelWand, color *I.PixelPacket)

func (p *PixelWand) SetQuantumColor(color *I.PixelPacket) { PixelSetQuantumColor(p, color) }

var PixelSetRed func(p *PixelWand, red float64)

func (p *PixelWand) SetRed(red float64) { PixelSetRed(p, red) }

var PixelSetRedQuantum func(p *PixelWand, red I.Quantum)

func (p *PixelWand) SetRedQuantum(red I.Quantum) { PixelSetRedQuantum(p, red) }

var PixelSetYellow func(p *PixelWand, yellow float64)

func (p *PixelWand) SetYellow(yellow float64) { PixelSetYellow(p, yellow) }

var PixelSetYellowQuantum func(p *PixelWand, yellow I.Quantum)

func (p *PixelWand) SetYellowQuantum(yellow I.Quantum) { PixelSetYellowQuantum(p, yellow) }

var ClearPixelIterator func(p *PixelIterator)

func (p *PixelIterator) Clear() { ClearPixelIterator(p) }

var ClonePixelIterator func(p *PixelIterator) *PixelIterator

func (p *PixelIterator) Clone() *PixelIterator { return ClonePixelIterator(p) }

var DestroyPixelIterator func(p *PixelIterator) *PixelIterator

func (p *PixelIterator) Destroy() *PixelIterator { return DestroyPixelIterator(p) }

var IsPixelIterator func(p *PixelIterator) bool

func (p *PixelIterator) IsPixelIterator() bool { return IsPixelIterator(p) }

var PixelGetCurrentIteratorRow func(p *PixelIterator, numberWands *uint32) []*PixelWand

func (p *PixelIterator) CurrentRow(numberWands *uint32) []*PixelWand {
	return PixelGetCurrentIteratorRow(p, numberWands)
}

var PixelGetIteratorException func(p *PixelIterator, severity *I.ExceptionType) string

func (p *PixelIterator) IteratoException(severity *I.ExceptionType) string {
	return PixelGetIteratorException(p, severity)
}

var PixelGetIteratorExceptionType func(p *PixelIterator) I.ExceptionType

func (p *PixelIterator) ExceptionType() I.ExceptionType {
	return PixelGetIteratorExceptionType(p)
}

var PixelGetIteratorRow func(p *PixelIterator) bool

func (p *PixelIterator) Row() bool { return PixelGetIteratorRow(p) }

var PixelGetNextIteratorRow func(p *PixelIterator, numberWands *uint32) []*PixelWand

func (p *PixelIterator) NextRow(numberWands *uint32) []*PixelWand {
	return PixelGetNextIteratorRow(p, numberWands)
}

var PixelGetPreviousIteratorRow func(p *PixelIterator, numberWands *uint32) []*PixelWand

func (p *PixelIterator) PreviousRow(numberWands *uint32) []*PixelWand {
	return PixelGetPreviousIteratorRow(p, numberWands)
}

var PixelClearIteratorException func(p *PixelIterator) bool

func (p *PixelIterator) ClearException() bool { return PixelClearIteratorException(p) }

var PixelResetIterator func(p *PixelIterator)

func (p *PixelIterator) Reset() { PixelResetIterator(p) }

var PixelSetFirstIteratorRow func(p *PixelIterator)

func (p *PixelIterator) SetFirstIteratorRow() { PixelSetFirstIteratorRow(p) }

var PixelSetIteratorRow func(p *PixelIterator, row int32) bool

func (p *PixelIterator) SetIteratorRow(row int32) bool { return PixelSetIteratorRow(p, row) }

var PixelSetLastIteratorRow func(p *PixelIterator)

func (p *PixelIterator) SetLastIteratorRow() { PixelSetLastIteratorRow(p) }

var PixelSyncIterator func(p *PixelIterator) bool

func (p *PixelIterator) Sync() bool { return PixelSyncIterator(p) }

var CloneWandView func(w *WandView) *WandView

func (w *WandView) Clone() *WandView { return CloneWandView(w) }

var DestroyWandView func(w *WandView) *WandView

func (w *WandView) Destroy() *WandView { return DestroyWandView(w) }

var DuplexTransferWandViewIterator func(w *WandView, duplex, destination *WandView, transfer I.DuplexTransferWandViewMethod, context *Void) bool

func (w *WandView) DuplexTransferIterator(duplex, destination *WandView, transfer I.DuplexTransferWandViewMethod, context *Void) bool {
	return DuplexTransferWandViewIterator(w, duplex, destination, transfer, context)
}

var GetWandViewException func(w *WandView, severity *I.ExceptionType) string

func (w *WandView) Exception(severity *I.ExceptionType) string {
	return GetWandViewException(w, severity)
}

var GetWandViewExtent func(w *WandView) I.RectangleInfo

func (w *WandView) Extent() I.RectangleInfo { return GetWandViewExtent(w) }

var GetWandViewIterator func(w *WandView, get GetWandViewMethod, context *Void) bool

func (w *WandView) Iterator(get GetWandViewMethod, context *Void) bool {
	return GetWandViewIterator(w, get, context)
}

var GetWandViewPixels func(w *WandView) *PixelWand

func (w *WandView) Pixels() *PixelWand { return GetWandViewPixels(w) }

var GetWandViewWand func(w *WandView) *MagickWand

func (w *WandView) Wand() *MagickWand { return GetWandViewWand(w) }

var IsWandView func(w *WandView) bool

func (w *WandView) IsWandView() bool { return IsWandView(w) }

var SetWandViewDescription func(w *WandView, description string)

func (w *WandView) SetDescription(description string) { SetWandViewDescription(w, description) }

var SetWandViewIterator func(w *WandView, set SetWandViewMethod, context *Void) bool

func (w *WandView) SetIterator(set SetWandViewMethod, context *Void) bool {
	return SetWandViewIterator(w, set, context)
}

var SetWandViewThreads func(w *WandView, numberThreads uint32)

func (w *WandView) SetThreads(numberThreads uint32) { SetWandViewThreads(w, numberThreads) }

var TransferWandViewIterator func(w *WandView, destination *WandView, transfer I.TransferWandViewMethod, context *Void) bool

func (w *WandView) TransferIterator(destination *WandView, transfer I.TransferWandViewMethod, context *Void) bool {
	return TransferWandViewIterator(w, destination, transfer, context)
}

var UpdateWandViewIterator func(w *WandView, update I.UpdateWandViewMethod, context *Void) bool

func (w *WandView) UpdateIterator(update I.UpdateWandViewMethod, context *Void) bool {
	return UpdateWandViewIterator(w, update, context)
}

var AnimateImageCommand func(i *ImageInfo, argc int, argv, metadata []string, e *ExceptionInfo) bool
var CommandGenesis func(i *ImageInfo, command I.MagickCommand, argc int, argv, metadata []string, exception *ExceptionInfo) bool
var CompareImageCommand func(i *ImageInfo, argc int, argv, metadata []string, e *ExceptionInfo) bool
var CompositeImageCommand func(i *ImageInfo, argc int, argv, metadata []string, e *ExceptionInfo) bool
var ConjureImageCommand func(i *ImageInfo, argc int, argv, metadata []string, e *ExceptionInfo) bool
var ConvertImageCommand func(i *ImageInfo, argc int, argv, metadata []string, e *ExceptionInfo) bool
var DisplayImageCommand func(i *ImageInfo, argc int, argv, metadata []string, e *ExceptionInfo) bool
var IdentifyImageCommand func(i *ImageInfo, argc int, argv, metadata []string, e *ExceptionInfo) bool
var ImportImageCommand func(i *ImageInfo, argc int, argv, metadata []string, e *ExceptionInfo) bool
var MogrifyImage func(i *ImageInfo, argc int, argv []string, images []*Image, e *ExceptionInfo) bool
var MogrifyImageCommand func(i *ImageInfo, argc int, argv, metadata []string, e *ExceptionInfo) bool
var MogrifyImageInfo func(i *ImageInfo, argc int, argv []string, e *ExceptionInfo) bool
var MogrifyImageList func(i *ImageInfo, argc int, argv []string, images []*Image, e *ExceptionInfo) bool
var MogrifyImages func(i *ImageInfo, post bool, argc int, argv []string, images []*Image, e *ExceptionInfo) bool
var MontageImageCommand func(i *ImageInfo, argc int, argv, metadata []string, e *ExceptionInfo) bool
var StreamImageCommand func(i *ImageInfo, argc int, argv, metadata []string, e *ExceptionInfo) bool

func (i *ImageInfo) AnimateImageCommand(argc int, argv, metadata []string, e *ExceptionInfo) bool {
	return AnimateImageCommand(i, argc, argv, metadata, e)
}
func (i *ImageInfo) CommandGenesis(command I.MagickCommand, argc int, argv, metadata []string, exception *ExceptionInfo) bool {
	return CommandGenesis(i, command, argc, argv, metadata, exception)
}
func (i *ImageInfo) CompareImageCommand(argc int, argv, metadata []string, e *ExceptionInfo) bool {
	return CompareImageCommand(i, argc, argv, metadata, e)
}
func (i *ImageInfo) CompositeImageCommand(argc int, argv, metadata []string, e *ExceptionInfo) bool {
	return CompositeImageCommand(i, argc, argv, metadata, e)
}
func (i *ImageInfo) ConjureImageCommand(argc int, argv, metadata []string, e *ExceptionInfo) bool {
	return ConjureImageCommand(i, argc, argv, metadata, e)
}
func (i *ImageInfo) ConvertImageCommand(argc int, argv, metadata []string, e *ExceptionInfo) bool {
	return ConvertImageCommand(i, argc, argv, metadata, e)
}
func (i *ImageInfo) DisplayImageCommand(argc int, argv, metadata []string, e *ExceptionInfo) bool {
	return DisplayImageCommand(i, argc, argv, metadata, e)
}
func (i *ImageInfo) IdentifyImageCommand(argc int, argv, metadata []string, e *ExceptionInfo) bool {
	return IdentifyImageCommand(i, argc, argv, metadata, e)
}
func (i *ImageInfo) ImportImageCommand(argc int, argv, metadata []string, e *ExceptionInfo) bool {
	return ImportImageCommand(i, argc, argv, metadata, e)
}
func (i *ImageInfo) MogrifyImage(argc int, argv []string, images []*Image, e *ExceptionInfo) bool {
	return MogrifyImage(i, argc, argv, images, e)
}
func (i *ImageInfo) MogrifyImageCommand(argc int, argv, metadata []string, e *ExceptionInfo) bool {
	return MogrifyImageCommand(i, argc, argv, metadata, e)
}
func (i *ImageInfo) MogrifyImageInfo(argc int, argv []string, e *ExceptionInfo) bool {
	return MogrifyImageInfo(i, argc, argv, e)
}
func (i *ImageInfo) MogrifyImageList(argc int, argv []string, images []*Image, e *ExceptionInfo) bool {
	return MogrifyImageList(i, argc, argv, images, e)
}
func (i *ImageInfo) MogrifyImages(post bool, argc int, argv []string, images []*Image, e *ExceptionInfo) bool {
	return MogrifyImages(i, post, argc, argv, images, e)
}
func (i *ImageInfo) MontageImageCommand(argc int, argv, metadata []string, e *ExceptionInfo) bool {
	return MontageImageCommand(i, argc, argv, metadata, e)
}
func (i *ImageInfo) StreamImageCommand(argc int, argv, metadata []string, e *ExceptionInfo) bool {
	return StreamImageCommand(i, argc, argv, metadata, e)
}

var Resource func(r I.ResourceType) I.MagickSizeType

var ResourceLimit func(r I.ResourceType) I.MagickSizeType

var SetResourceLimit func(r I.ResourceType, limit I.MagickSizeType) bool

var DestroyImage func(i *Image) *Image

func (i *Image) Destroy() *Image { return DestroyImage(i) }

var GetCopyright func() string

var GetHomeURL func() string

var GetPackageName func() string

var GetQuantumDepth func(depth *uint32) string

var GetQuantumRange func(range_ *uint32) string

var GetReleaseDate func() string

var GetVersion func(version *uint32) string

var QueryConfigureOption func(option string) string

var QueryConfigureOptions func(pattern string, numberOptions *uint32) []string

var QueryFonts func(pattern string, numberFonts *uint32) []string

var QueryFormats func(pattern string, numberFormats *uint32) []string

var RelinquishMemory func(resource *Void) *Void

var WandGenesis func()

var WandTerminus func()

var NewMagickWandFromImage func(i *Image) *MagickWand

func (i *Image) NewMagickWand() *MagickWand { return NewMagickWandFromImage(i) }

var AcquireWandId func() uint32

var DestroyWandIds func()

var DrawAllocateWand func(*DrawInfo, *Image) *DrawingWand

var DrawGetClipRule func(d *DrawingWand) I.FillRule
var DrawGetTextInterlineSpacing func(d *DrawingWand) float64
var DrawRender func(d *DrawingWand) bool

func (d *DrawingWand) ClipRule() I.FillRule          { return DrawGetClipRule(d) }
func (d *DrawingWand) TextInterlineSpacing() float64 { return DrawGetTextInterlineSpacing(d) }
func (d *DrawingWand) Render() bool                  { return DrawRender(d) }

var GetImageChannelExtrema func(m *MagickWand, c I.ChannelType, minima, maxima *uint32) bool

func (m *MagickWand) GetImageChannelExtrema(c I.ChannelType, minima, maxima *uint32) bool {
	return GetImageChannelExtrema(m, c, minima, maxima)
}

var GetImageEndian func(*MagickWand) I.EndianType

var GetImageExtrema func(m *MagickWand, min, max *uint32) bool

func (m *MagickWand) Extrema(min, max *uint32) bool { return GetImageExtrema(m, min, max) }

var GetImageIndex func(m *MagickWand) int32

func (m *MagickWand) GetImageIndex() int32 { return GetImageIndex(m) }

var GetImageRange func(*MagickWand, *float64, *float64) bool

var LiquidRescaleImage func(m *MagickWand, columns, rows uint32, deltaX, rigidity float64) bool

func (m *MagickWand) LiquidRescaleImage(columns, rows uint32, deltaX, rigidity float64) bool {
	return LiquidRescaleImage(m, columns, rows, deltaX, rigidity)
}

var OptimizeImageTransparency func(*MagickWand) bool

var SolarizeImageChannel func(*MagickWand, I.ChannelType, float64) bool

var RelinquishWandId func(uint32)

var DLL = "CORE_RL_wand_.dll"

func init() {
	AddDllApis(DLL, false, allApis)
	allApis = nil
}

var allApis = Apis{
	{"AcquireWandId", &AcquireWandId},
	{"AnimateImageCommand", &AnimateImageCommand},
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
	{"CompareImageCommand", &CompareImageCommand},
	{"CompositeImageCommand", &CompositeImageCommand},
	{"ConjureImageCommand", &ConjureImageCommand},
	{"ConvertImageCommand", &ConvertImageCommand},
	{"DestroyDrawingWand", &DestroyDrawingWand},
	{"DestroyMagickWand", &DestroyMagickWand},
	{"DestroyPixelIterator", &DestroyPixelIterator},
	{"DestroyPixelView", &DestroyPixelView},
	{"DestroyPixelWand", &DestroyPixelWand},
	{"DestroyPixelWands", &DestroyPixelWands},
	{"DestroyWandIds", &DestroyWandIds},
	{"DestroyWandView", &DestroyWandView},
	{"DisplayImageCommand", &DisplayImageCommand},
	{"DrawAffine", &DrawAffine},
	{"DrawAllocateWand", &DrawAllocateWand},
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
	{"DrawGetClipRule", &DrawGetClipRule},
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
	{"DrawGetTextInterlineSpacing", &DrawGetTextInterlineSpacing},
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
	{"DrawRender", &DrawRender},
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
	{"IdentifyImageCommand", &IdentifyImageCommand},
	{"ImportImageCommand", &ImportImageCommand},
	{"IsDrawingWand", &IsDrawingWand},
	{"IsMagickWand", &IsMagickWand},
	{"IsPixelIterator", &IsPixelIterator},
	{"IsPixelView", &IsPixelView},
	{"IsPixelWand", &IsPixelWand},
	{"IsPixelWandSimilar", &IsPixelWandSimilar},
	{"IsWandView", &IsWandView},
	{"MagickAdaptiveBlurImage", &AdaptiveBlurImage},
	{"MagickAdaptiveBlurImageChannel", &AdaptiveBlurImageChannel},
	{"MagickAdaptiveResizeImage", &AdaptiveResizeImage},
	{"MagickAdaptiveSharpenImage", &AdaptiveSharpenImage},
	{"MagickAdaptiveSharpenImageChannel", &AdaptiveSharpenImageChannel},
	{"MagickAdaptiveThresholdImage", &AdaptiveThresholdImage},
	{"MagickAddImage", &AddImage},
	{"MagickAddNoiseImage", &AddNoiseImage},
	{"MagickAddNoiseImageChannel", &AddNoiseImageChannel},
	{"MagickAffineTransformImage", &AffineTransformImage},
	{"MagickAnimateImages", &AnimateImages},
	{"MagickAnnotateImage", &AnnotateImage},
	{"MagickAppendImages", &AppendImages},
	{"MagickAutoGammaImage", &AutoGammaImage},
	{"MagickAutoGammaImageChannel", &AutoGammaImageChannel},
	{"MagickAutoLevelImage", &AutoLevelImage},
	{"MagickAutoLevelImageChannel", &AutoLevelImageChannel},
	{"MagickAverageImages", &AverageImages},
	{"MagickBlackThresholdImage", &BlackThresholdImage},
	{"MagickBlueShiftImage", &BlueShiftImage},
	{"MagickBlurImage", &BlurImage},
	{"MagickBlurImageChannel", &BlurImageChannel},
	{"MagickBorderImage", &BorderImage},
	{"MagickBrightnessContrastImage", &BrightnessContrastImage},
	{"MagickBrightnessContrastImageChannel", &BrightnessContrastImageChannel},
	{"MagickCharcoalImage", &CharcoalImage},
	{"MagickChopImage", &ChopImage},
	{"MagickClampImage", &ClampImage},
	{"MagickClampImageChannel", &ClampImageChannel},
	{"MagickClearException", &ClearException},
	{"MagickClipImage", &ClipImage},
	{"MagickClipImagePath", &ClipImagePath},
	{"MagickClipPathImage", &ClipPathImage},
	{"MagickClutImage", &ClutImage},
	{"MagickClutImageChannel", &ClutImageChannel},
	{"MagickCoalesceImages", &CoalesceImages},
	{"MagickColorDecisionListImage", &ColorDecisionListImage},
	{"MagickColorFloodfillImage", &ColorFloodfillImage},
	{"MagickColorMatrixImage", &ColorMatrixImage},
	{"MagickColorizeImage", &ColorizeImage},
	{"MagickCombineImages", &CombineImages},
	{"MagickCommandGenesis", &CommandGenesis},
	{"MagickCommentImage", &CommentImage},
	{"MagickCompareImageChannels", &CompareImageChannels},
	{"MagickCompareImageLayers", &CompareImageLayers},
	{"MagickCompareImages", &CompareImages},
	{"MagickCompositeImage", &CompositeImage},
	{"MagickCompositeImageChannel", &CompositeImageChannel},
	{"MagickCompositeLayers", &CompositeLayers},
	{"MagickConstituteImage", &ConstituteImage},
	{"MagickContrastImage", &ContrastImage},
	{"MagickContrastStretchImage", &ContrastStretchImage},
	{"MagickContrastStretchImageChannel", &ContrastStretchImageChannel},
	{"MagickConvolveImage", &ConvolveImage},
	{"MagickConvolveImageChannel", &ConvolveImageChannel},
	{"MagickCropImage", &CropImage},
	{"MagickCycleColormapImage", &CycleColormapImage},
	{"MagickDecipherImage", &DecipherImage},
	{"MagickDeconstructImages", &DeconstructImages},
	{"MagickDeleteImageArtifact", &DeleteImageArtifact},
	{"MagickDeleteImageProperty", &DeleteImageProperty},
	{"MagickDeleteOption", &DeleteOption},
	{"MagickDescribeImage", &DescribeImage},
	{"MagickDeskewImage", &DeskewImage},
	{"MagickDespeckleImage", &DespeckleImage},
	{"MagickDestroyImage", &DestroyImage},
	{"MagickDisplayImage", &DisplayImage},
	{"MagickDisplayImages", &DisplayImages},
	{"MagickDistortImage", &DistortImage},
	{"MagickDrawImage", &DrawImage},
	{"MagickEdgeImage", &EdgeImage},
	{"MagickEmbossImage", &EmbossImage},
	{"MagickEncipherImage", &EncipherImage},
	{"MagickEnhanceImage", &EnhanceImage},
	{"MagickEqualizeImage", &EqualizeImage},
	{"MagickEqualizeImageChannel", &EqualizeImageChannel},
	{"MagickEvaluateImage", &EvaluateImage},
	{"MagickEvaluateImageChannel", &EvaluateImageChannel},
	{"MagickEvaluateImages", &EvaluateImages},
	{"MagickExportImagePixels", &ExportImagePixels},
	{"MagickExtentImage", &ExtentImage},
	{"MagickFilterImage", &FilterImage},
	{"MagickFilterImageChannel", &FilterImageChannel},
	{"MagickFlattenImages", &FlattenImages},
	{"MagickFlipImage", &FlipImage},
	{"MagickFloodfillPaintImage", &FloodfillPaintImage},
	{"MagickFlopImage", &FlopImage},
	{"MagickForwardFourierTransformImage", &ForwardFourierTransformImage},
	{"MagickFrameImage", &FrameImage},
	{"MagickFunctionImage", &FunctionImage},
	{"MagickFunctionImageChannel", &FunctionImageChannel},
	{"MagickFxImage", &FxImage},
	{"MagickFxImageChannel", &FxImageChannel},
	{"MagickGammaImage", &GammaImage},
	{"MagickGammaImageChannel", &GammaImageChannel},
	{"MagickGaussianBlurImage", &GaussianBlurImage},
	{"MagickGaussianBlurImageChannel", &GaussianBlurImageChannel},
	{"MagickGetAntialias", &GetAntialias},
	{"MagickGetBackgroundColor", &GetBackgroundColor},
	{"MagickGetColorspace", &GetColorspace},
	{"MagickGetCompression", &GetCompression},
	{"MagickGetCompressionQuality", &GetCompressionQuality},
	{"MagickGetCopyright", &GetCopyright},
	{"MagickGetException", &GetException},
	{"MagickGetExceptionType", &GetExceptionType},
	{"MagickGetFilename", &GetFilename},
	{"MagickGetFont", &GetFont},
	{"MagickGetFormat", &GetFormat},
	{"MagickGetGravity", &GetGravity},
	{"MagickGetHomeURL", &GetHomeURL},
	{"MagickGetImage", &GetImage},
	{"MagickGetImageAlphaChannel", &GetImageAlphaChannel},
	{"MagickGetImageArtifact", &GetImageArtifact},
	{"MagickGetImageArtifacts", &GetImageArtifacts},
	{"MagickGetImageAttribute", &GetImageAttribute},
	{"MagickGetImageBackgroundColor", &GetImageBackgroundColor},
	{"MagickGetImageBlob", &GetImageBlob},
	{"MagickGetImageBluePrimary", &GetImageBluePrimary},
	{"MagickGetImageBorderColor", &GetImageBorderColor},
	{"MagickGetImageChannelDepth", &GetImageChannelDepth},
	{"MagickGetImageChannelDistortion", &GetImageChannelDistortion},
	{"MagickGetImageChannelDistortions", &GetImageChannelDistortions},
	{"MagickGetImageChannelExtrema", &GetImageChannelExtrema},
	{"MagickGetImageChannelFeatures", &GetImageChannelFeatures},
	{"MagickGetImageChannelKurtosis", &GetImageChannelKurtosis},
	{"MagickGetImageChannelMean", &GetImageChannelMean},
	{"MagickGetImageChannelRange", &GetImageChannelRange},
	{"MagickGetImageChannelStatistics", &GetImageChannelStatistics},
	{"MagickGetImageClipMask", &GetImageClipMask},
	{"MagickGetImageColormapColor", &GetImageColormapColor},
	{"MagickGetImageColors", &GetImageColors},
	{"MagickGetImageColorspace", &GetImageColorspace},
	{"MagickGetImageCompose", &GetImageCompose},
	{"MagickGetImageCompression", &GetImageCompression},
	{"MagickGetImageCompressionQuality", &GetImageCompressionQuality},
	{"MagickGetImageDelay", &GetImageDelay},
	{"MagickGetImageDepth", &GetImageDepth},
	{"MagickGetImageDispose", &GetImageDispose},
	{"MagickGetImageDistortion", &GetImageDistortion},
	{"MagickGetImageEndian", &GetImageEndian},
	{"MagickGetImageExtrema", &GetImageExtrema},
	{"MagickGetImageFilename", &GetImageFilename},
	{"MagickGetImageFormat", &GetImageFormat},
	{"MagickGetImageFuzz", &GetImageFuzz},
	{"MagickGetImageGamma", &GetImageGamma},
	{"MagickGetImageGravity", &GetImageGravity},
	{"MagickGetImageGreenPrimary", &GetImageGreenPrimary},
	{"MagickGetImageHeight", &GetImageHeight},
	{"MagickGetImageHistogram", &GetImageHistogram},
	{"MagickGetImageIndex", &GetImageIndex},
	{"MagickGetImageInterlaceScheme", &GetImageInterlaceScheme},
	{"MagickGetImageInterpolateMethod", &GetImageInterpolateMethod},
	{"MagickGetImageIterations", &GetImageIterations},
	{"MagickGetImageLength", &GetImageLength},
	{"MagickGetImageMatte", &GetImageMatte},
	{"MagickGetImageMatteColor", &GetImageMatteColor},
	{"MagickGetImageOrientation", &GetImageOrientation},
	{"MagickGetImagePage", &GetImagePage},
	{"MagickGetImagePixelColor", &GetImagePixelColor},
	{"MagickGetImagePixels", &GetImagePixels},
	{"MagickGetImageProfile", &GetImageProfile},
	{"MagickGetImageProfiles", &GetImageProfiles},
	{"MagickGetImageProperties", &GetImageProperties},
	{"MagickGetImageProperty", &GetImageProperty},
	{"MagickGetImageRange", &GetImageRange},
	{"MagickGetImageRedPrimary", &GetImageRedPrimary},
	{"MagickGetImageRegion", &GetImageRegion},
	{"MagickGetImageRenderingIntent", &GetImageRenderingIntent},
	{"MagickGetImageResolution", &GetImageResolution},
	{"MagickGetImageScene", &GetImageScene},
	{"MagickGetImageSignature", &GetImageSignature},
	{"MagickGetImageSize", &GetImageSize},
	{"MagickGetImageTicksPerSecond", &GetImageTicksPerSecond},
	{"MagickGetImageTotalInkDensity", &GetImageTotalInkDensity},
	{"MagickGetImageType", &GetImageType},
	{"MagickGetImageUnits", &GetImageUnits},
	{"MagickGetImageVirtualPixelMethod", &GetImageVirtualPixelMethod},
	{"MagickGetImageWhitePoint", &GetImageWhitePoint},
	{"MagickGetImageWidth", &GetImageWidth},
	{"MagickGetImagesBlob", &GetImagesBlob},
	{"MagickGetInterlaceScheme", &GetInterlaceScheme},
	{"MagickGetInterpolateMethod", &GetInterpolateMethod},
	{"MagickGetIteratorIndex", &GetIteratorIndex},
	{"MagickGetNumberImages", &GetNumberImages},
	{"MagickGetOption", &GetOption},
	{"MagickGetOptions", &GetOptions},
	{"MagickGetOrientation", &GetOrientation},
	{"MagickGetPackageName", &GetPackageName},
	{"MagickGetPage", &GetPage},
	{"MagickGetPointsize", &GetPointsize},
	{"MagickGetQuantumDepth", &GetQuantumDepth},
	{"MagickGetQuantumRange", &GetQuantumRange},
	{"MagickGetReleaseDate", &GetReleaseDate},
	{"MagickGetResolution", &GetResolution},
	{"MagickGetResource", &Resource},
	{"MagickGetResourceLimit", &ResourceLimit},
	{"MagickGetSamplingFactors", &GetSamplingFactors},
	{"MagickGetSize", &GetSize},
	{"MagickGetSizeOffset", &GetSizeOffset},
	{"MagickGetType", &GetType},
	{"MagickGetVersion", &GetVersion},
	{"MagickHaldClutImage", &HaldClutImage},
	{"MagickHaldClutImageChannel", &HaldClutImageChannel},
	{"MagickHasNextImage", &HasNextImage},
	{"MagickHasPreviousImage", &HasPreviousImage},
	{"MagickIdentifyImage", &IdentifyImage},
	{"MagickImplodeImage", &ImplodeImage},
	{"MagickImportImagePixels", &ImportImagePixels},
	{"MagickInverseFourierTransformImage", &InverseFourierTransformImage},
	{"MagickLabelImage", &LabelImage},
	{"MagickLevelImage", &LevelImage},
	{"MagickLevelImageChannel", &LevelImageChannel},
	{"MagickLinearStretchImage", &LinearStretchImage},
	{"MagickLiquidRescaleImage", &LiquidRescaleImage},
	{"MagickMagnifyImage", &MagnifyImage},
	{"MagickMapImage", &MapImage},
	{"MagickMatteFloodfillImage", &MatteFloodfillImage},
	{"MagickMaximumImages", &MaximumImages},
	{"MagickMedianFilterImage", &MedianFilterImage},
	{"MagickMergeImageLayers", &MergeImageLayers},
	{"MagickMinifyImage", &MinifyImage},
	{"MagickMinimumImages", &MinimumImages},
	{"MagickModeImage", &ModeImage},
	{"MagickModulateImage", &ModulateImage},
	{"MagickMontageImage", &MontageImage},
	{"MagickMorphImages", &MorphImages},
	{"MagickMorphologyImage", &MorphologyImage},
	{"MagickMorphologyImageChannel", &MorphologyImageChannel},
	{"MagickMosaicImages", &MosaicImages},
	{"MagickMotionBlurImage", &MotionBlurImage},
	{"MagickMotionBlurImageChannel", &MotionBlurImageChannel},
	{"MagickNegateImage", &NegateImage},
	{"MagickNegateImageChannel", &NegateImageChannel},
	{"MagickNewImage", &NewImage},
	{"MagickNextImage", &NextImage},
	{"MagickNormalizeImage", &NormalizeImage},
	{"MagickNormalizeImageChannel", &NormalizeImageChannel},
	{"MagickOilPaintImage", &OilPaintImage},
	{"MagickOpaqueImage", &OpaqueImage},
	{"MagickOpaquePaintImage", &OpaquePaintImage},
	{"MagickOpaquePaintImageChannel", &OpaquePaintImageChannel},
	{"MagickOptimizeImageLayers", &OptimizeImageLayers},
	{"MagickOptimizeImageTransparency", &OptimizeImageTransparency},
	{"MagickOrderedPosterizeImage", &OrderedPosterizeImage},
	{"MagickOrderedPosterizeImageChannel", &OrderedPosterizeImageChannel},
	{"MagickPaintFloodfillImage", &PaintFloodfillImage},
	{"MagickPaintOpaqueImage", &PaintOpaqueImage},
	{"MagickPaintOpaqueImageChannel", &PaintOpaqueImageChannel},
	{"MagickPaintTransparentImage", &PaintTransparentImage},
	{"MagickPingImage", &PingImage},
	{"MagickPingImageBlob", &PingImageBlob},
	{"MagickPingImageFile", &PingImageFile},
	{"MagickPolaroidImage", &PolaroidImage},
	{"MagickPosterizeImage", &PosterizeImage},
	{"MagickPreviewImages", &PreviewImages},
	{"MagickPreviousImage", &PreviousImage},
	{"MagickProfileImage", &ProfileImage},
	{"MagickQuantizeImage", &QuantizeImage},
	{"MagickQuantizeImages", &QuantizeImages},
	{"MagickQueryConfigureOption", &QueryConfigureOption},
	{"MagickQueryConfigureOptions", &QueryConfigureOptions},
	{"MagickQueryFontMetrics", &QueryFontMetrics},
	{"MagickQueryFonts", &QueryFonts},
	{"MagickQueryFormats", &QueryFormats},
	{"MagickQueryMultilineFontMetrics", &QueryMultilineFontMetrics},
	{"MagickRadialBlurImage", &RadialBlurImage},
	{"MagickRadialBlurImageChannel", &RadialBlurImageChannel},
	{"MagickRaiseImage", &RaiseImage},
	{"MagickRandomThresholdImage", &RandomThresholdImage},
	{"MagickRandomThresholdImageChannel", &RandomThresholdImageChannel},
	{"MagickReadImage", &ReadImage},
	{"MagickReadImageBlob", &ReadImageBlob},
	{"MagickReadImageFile", &ReadImageFile},
	{"MagickRecolorImage", &RecolorImage},
	{"MagickReduceNoiseImage", &ReduceNoiseImage},
	{"MagickRegionOfInterestImage", &RegionOfInterestImage},
	{"MagickRelinquishMemory", &RelinquishMemory},
	{"MagickRemapImage", &RemapImage},
	{"MagickRemoveImage", &RemoveImage},
	{"MagickRemoveImageProfile", &RemoveImageProfile},
	{"MagickResampleImage", &ResampleImage},
	{"MagickResetImagePage", &ResetImagePage},
	{"MagickResetIterator", &ResetIterator},
	{"MagickResizeImage", &ResizeImage},
	{"MagickRollImage", &RollImage},
	{"MagickRotateImage", &RotateImage},
	{"MagickSampleImage", &SampleImage},
	{"MagickScaleImage", &ScaleImage},
	{"MagickSegmentImage", &SegmentImage},
	{"MagickSelectiveBlurImage", &SelectiveBlurImage},
	{"MagickSelectiveBlurImageChannel", &SelectiveBlurImageChannel},
	{"MagickSeparateImageChannel", &SeparateImageChannel},
	{"MagickSepiaToneImage", &SepiaToneImage},
	{"MagickSetAntialias", &SetAntialias},
	{"MagickSetBackgroundColor", &SetBackgroundColor},
	{"MagickSetColorspace", &SetColorspace},
	{"MagickSetCompression", &SetCompression},
	{"MagickSetCompressionQuality", &SetCompressionQuality},
	{"MagickSetDepth", &SetDepth},
	{"MagickSetExtract", &SetExtract},
	{"MagickSetFilename", &SetFilename},
	{"MagickSetFirstIterator", &SetFirstIterator},
	{"MagickSetFont", &SetFont},
	{"MagickSetFormat", &SetFormat},
	{"MagickSetGravity", &SetGravity},
	{"MagickSetImage", &SetImage},
	{"MagickSetImageAlphaChannel", &SetImageAlphaChannel},
	{"MagickSetImageArtifact", &SetImageArtifact},
	{"MagickSetImageAttribute", &SetImageAttribute},
	{"MagickSetImageBackgroundColor", &SetImageBackgroundColor},
	{"MagickSetImageBias", &SetImageBias},
	{"MagickSetImageBluePrimary", &SetImageBluePrimary},
	{"MagickSetImageBorderColor", &SetImageBorderColor},
	{"MagickSetImageChannelDepth", &SetImageChannelDepth},
	{"MagickSetImageClipMask", &SetImageClipMask},
	{"MagickSetImageColor", &SetImageColor},
	{"MagickSetImageColormapColor", &SetImageColormapColor},
	{"MagickSetImageColorspace", &SetImageColorspace},
	{"MagickSetImageCompose", &SetImageCompose},
	{"MagickSetImageCompression", &SetImageCompression},
	{"MagickSetImageCompressionQuality", &SetImageCompressionQuality},
	{"MagickSetImageDelay", &SetImageDelay},
	{"MagickSetImageDepth", &SetImageDepth},
	{"MagickSetImageDispose", &SetImageDispose},
	{"MagickSetImageEndian", &SetImageEndian},
	{"MagickSetImageExtent", &SetImageExtent},
	{"MagickSetImageFilename", &SetImageFilename},
	{"MagickSetImageFormat", &SetImageFormat},
	{"MagickSetImageFuzz", &SetImageFuzz},
	{"MagickSetImageGamma", &SetImageGamma},
	{"MagickSetImageGravity", &SetImageGravity},
	{"MagickSetImageGreenPrimary", &SetImageGreenPrimary},
	{"MagickSetImageIndex", &SetImageIndex},
	{"MagickSetImageInterlaceScheme", &SetImageInterlaceScheme},
	{"MagickSetImageInterpolateMethod", &SetImageInterpolateMethod},
	{"MagickSetImageIterations", &SetImageIterations},
	{"MagickSetImageMatte", &SetImageMatte},
	{"MagickSetImageMatteColor", &SetImageMatteColor},
	{"MagickSetImageOpacity", &SetImageOpacity},
	{"MagickSetImageOption", &SetImageOption},
	{"MagickSetImageOrientation", &SetImageOrientation},
	{"MagickSetImagePage", &SetImagePage},
	{"MagickSetImagePixels", &SetImagePixels},
	{"MagickSetImageProfile", &SetImageProfile},
	{"MagickSetImageProgressMonitor", &SetImageProgressMonitor},
	{"MagickSetImageProperty", &SetImageProperty},
	{"MagickSetImageRedPrimary", &SetImageRedPrimary},
	{"MagickSetImageRenderingIntent", &SetImageRenderingIntent},
	{"MagickSetImageResolution", &SetImageResolution},
	{"MagickSetImageScene", &SetImageScene},
	{"MagickSetImageTicksPerSecond", &SetImageTicksPerSecond},
	{"MagickSetImageType", &SetImageType},
	{"MagickSetImageUnits", &SetImageUnits},
	{"MagickSetImageVirtualPixelMethod", &SetImageVirtualPixelMethod},
	{"MagickSetImageWhitePoint", &SetImageWhitePoint},
	{"MagickSetInterlaceScheme", &SetInterlaceScheme},
	{"MagickSetInterpolateMethod", &SetInterpolateMethod},
	{"MagickSetIteratorIndex", &SetIteratorIndex},
	{"MagickSetLastIterator", &SetLastIterator},
	{"MagickSetOption", &SetOption},
	{"MagickSetOrientation", &SetOrientation},
	{"MagickSetPage", &SetPage},
	{"MagickSetPassphrase", &SetPassphrase},
	{"MagickSetPointsize", &SetPointsize},
	{"MagickSetProgressMonitor", &SetProgressMonitor},
	{"MagickSetResolution", &SetResolution},
	{"MagickSetResourceLimit", &SetResourceLimit},
	{"MagickSetSamplingFactors", &SetSamplingFactors},
	{"MagickSetSize", &SetSize},
	{"MagickSetSizeOffset", &SetSizeOffset},
	{"MagickSetType", &SetType},
	{"MagickShadeImage", &ShadeImage},
	{"MagickShadowImage", &ShadowImage},
	{"MagickSharpenImage", &SharpenImage},
	{"MagickSharpenImageChannel", &SharpenImageChannel},
	{"MagickShaveImage", &ShaveImage},
	{"MagickShearImage", &ShearImage},
	{"MagickSigmoidalContrastImage", &SigmoidalContrastImage},
	{"MagickSigmoidalContrastImageChannel", &SigmoidalContrastImageChannel},
	{"MagickSimilarityImage", &SimilarityImage},
	{"MagickSketchImage", &SketchImage},
	{"MagickSmushImages", &SmushImages},
	{"MagickSolarizeImage", &SolarizeImage},
	{"MagickSolarizeImageChannel", &SolarizeImageChannel},
	{"MagickSparseColorImage", &SparseColorImage},
	{"MagickSpliceImage", &SpliceImage},
	{"MagickSpreadImage", &SpreadImage},
	{"MagickStatisticImage", &StatisticImage},
	{"MagickSteganoImage", &SteganoImage},
	{"MagickStereoImage", &StereoImage},
	{"MagickStripImage", &StripImage},
	{"MagickSwirlImage", &SwirlImage},
	{"MagickTextureImage", &TextureImage},
	{"MagickThresholdImage", &ThresholdImage},
	{"MagickThresholdImageChannel", &ThresholdImageChannel},
	{"MagickThumbnailImage", &ThumbnailImage},
	{"MagickTintImage", &TintImage},
	{"MagickTransformImage", &TransformImage},
	{"MagickTransformImageColorspace", &TransformImageColorspace},
	{"MagickTransparentImage", &TransparentImage},
	{"MagickTransparentPaintImage", &TransparentPaintImage},
	{"MagickTransposeImage", &TransposeImage},
	{"MagickTransverseImage", &TransverseImage},
	{"MagickTrimImage", &TrimImage},
	{"MagickUniqueImageColors", &UniqueImageColors},
	{"MagickUnsharpMaskImage", &UnsharpMaskImage},
	{"MagickUnsharpMaskImageChannel", &UnsharpMaskImageChannel},
	{"MagickVignetteImage", &VignetteImage},
	{"MagickWandGenesis", &WandGenesis},
	{"MagickWandTerminus", &WandTerminus},
	{"MagickWaveImage", &WaveImage},
	{"MagickWhiteThresholdImage", &WhiteThresholdImage},
	{"MagickWriteImage", &WriteImage},
	{"MagickWriteImageBlob", &WriteImageBlob},
	{"MagickWriteImageFile", &WriteImageFile},
	{"MagickWriteImages", &WriteImages},
	{"MagickWriteImagesFile", &WriteImagesFile},
	{"MogrifyImage", &MogrifyImage},
	{"MogrifyImageCommand", &MogrifyImageCommand},
	{"MogrifyImageInfo", &MogrifyImageInfo},
	{"MogrifyImageList", &MogrifyImageList},
	{"MogrifyImages", &MogrifyImages},
	{"MontageImageCommand", &MontageImageCommand},
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
	// Undocumented {"PixelGetPreviousRow", &PixelGetPreviousRow},
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
	{"RelinquishWandId", &RelinquishWandId},
	{"SetPixelViewIterator", &SetPixelViewIterator},
	{"SetWandViewDescription", &SetWandViewDescription},
	{"SetWandViewIterator", &SetWandViewIterator},
	{"SetWandViewThreads", &SetWandViewThreads},
	{"StreamImageCommand", &StreamImageCommand},
	{"TransferPixelViewIterator", &TransferPixelViewIterator},
	{"TransferWandViewIterator", &TransferWandViewIterator},
	{"UpdatePixelViewIterator", &UpdatePixelViewIterator},
	{"UpdateWandViewIterator", &UpdateWandViewIterator},
}

// Deprecated.
var AverageImages func(m *MagickWand) *MagickWand

// Deprecated.
func (m *MagickWand) Average() *MagickWand { return AverageImages(m) }

// Deprecated.
var ColorFloodfillImage func(m *MagickWand, fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y int32) bool

// Deprecated.
func (m *MagickWand) ColorFloodfill(fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y int32) bool {
	return ColorFloodfillImage(m, fill, fuzz, bordercolor, x, y)
}

// Deprecated.
var DescribeImage func(m *MagickWand) string

// Deprecated.
func (m *MagickWand) Describe() string { return DescribeImage(m) }

// Deprecated.
var FlattenImages func(m *MagickWand) *MagickWand

// Deprecated.
func (m *MagickWand) Flatten() *MagickWand { return FlattenImages(m) }

// Deprecated.
var GetImageAttribute func(m *MagickWand, property string) string

// Deprecated.
func (m *MagickWand) Attribute(property string) string {
	return GetImageAttribute(m, property)
}

// Deprecated.
var GetImagePixels func(m *MagickWand, x, y int32, columns, rows uint32, map_ string, storage I.StorageType, pixels *Void) bool

// Deprecated.
func (m *MagickWand) Pixels(x, y int32, columns, rows uint32, map_ string, storage I.StorageType, pixels *Void) bool {
	return GetImagePixels(m, x, y, columns, rows, map_, storage, pixels)
}

// Deprecated.
var GetImageSize func(m *MagickWand, length *I.MagickSizeType) bool

// Deprecated.
func (m *MagickWand) ImageSize(length *I.MagickSizeType) bool { return GetImageSize(m, length) }

// Deprecated.
var MatteFloodfillImage func(m *MagickWand, alpha, fuzz float64, bordercolor *PixelWand, x, y int32) bool

// Deprecated.
func (m *MagickWand) MatteFloodfill(alpha, fuzz float64, bordercolor *PixelWand, x, y int32) bool {
	return MatteFloodfillImage(m, alpha, fuzz, bordercolor, x, y)
}

// Deprecated.
var MedianFilterImage func(m *MagickWand, radius float64) bool

// Deprecated.
func (m *MagickWand) MedianFilter(radius float64) bool {
	return MedianFilterImage(m, radius)
}

// Deprecated.
var MosaicImages func(m *MagickWand) *MagickWand

// Deprecated.
func (m *MagickWand) Mosaic() *MagickWand { return MosaicImages(m) }

// Deprecated.
var OpaqueImage func(m *MagickWand, target, fill *PixelWand, fuzz float64) bool

// Deprecated.
func (m *MagickWand) Opaque(target, fill *PixelWand, fuzz float64) bool {
	return OpaqueImage(m, target, fill, fuzz)
}

// Deprecated.
var ReduceNoiseImage func(m *MagickWand, radius float64) bool

// Deprecated.
func (m *MagickWand) ReduceNoise(radius float64) bool { return ReduceNoiseImage(m, radius) }

// Deprecated.
var SetImageAttribute func(m *MagickWand, property, value string) bool

// Deprecated.
func (m *MagickWand) SetAttribute(property, value string) bool {
	return SetImageAttribute(m, property, value)
}

// Deprecated.
var SetImageIndex func(m *MagickWand, index int32) bool

// Deprecated.
func (m *MagickWand) SetIndex(index int32) bool { return SetImageIndex(m, index) }

// Deprecated.
var SetImagePixels func(m *MagickWand, x, y int32, columns, rows uint32, map_ string, storage I.StorageType, pixels *Void) bool

// Deprecated.
func (m *MagickWand) SetImagePixels(x, y int32, columns, rows uint32, map_ string, storage I.StorageType, pixels *Void) bool {
	return SetImagePixels(m, x, y, columns, rows, map_, storage, pixels)
}

// Deprecated.
var TransparentImage func(m *MagickWand, target *PixelWand, alpha, fuzz float64) bool

// Deprecated.
func (m *MagickWand) Transparent(target *PixelWand, alpha, fuzz float64) bool {
	return TransparentImage(m, target, alpha, fuzz)
}

// Deprecated.
var WriteImageBlob func(m *MagickWand, length *uint32) *byte

// Deprecated.
func (m *MagickWand) WriteBlob(length *uint32) *byte {
	return WriteImageBlob(m, length)
}

// Deprecated.
var SetImageOption func(m *MagickWand, format, key, value string) bool

// Deprecated.
func (m *MagickWand) SetImageOption(format, key, value string) bool {
	return SetImageOption(m, format, key, value)
}

// Deprecated.
var NewPixelView func(m *MagickWand) *PixelView

// Deprecated.
var NewPixelViewRegion func(m *MagickWand, x, y int32, width, height uint32) *PixelView

// Deprecated.
var GetImageMatte func(m *MagickWand) uint32

// Deprecated.
func (m *MagickWand) Matte() uint32 { return GetImageMatte(m) }

// Deprecated.
var MaximumImages func(m *MagickWand) *MagickWand

// Deprecated.
func (m *MagickWand) MaximumImages() *MagickWand { return MaximumImages(m) }

// Deprecated.
var MinimumImages func(m *MagickWand) *MagickWand

// Deprecated.
func (m *MagickWand) MinimumImages() *MagickWand { return MinimumImages(m) }

// Deprecated.
var ModeImage func(m *MagickWand, radius float64) bool

// Deprecated.
func (m *MagickWand) Mode(radius float64) bool { return ModeImage(m, radius) }

// Deprecated.
var PaintFloodfillImage func(m *MagickWand, channel I.ChannelType, fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y int32) bool

// Deprecated.
func (m *MagickWand) PaintFloodfill(channel I.ChannelType, fill *PixelWand, fuzz float64, bordercolor *PixelWand, x, y int32) bool {
	return PaintFloodfillImage(m, channel, fill, fuzz, bordercolor, x, y)
}

// Deprecated.
var PaintOpaqueImage func(m *MagickWand, target, fill *PixelWand, fuzz float64) bool

// Deprecated.
func (m *MagickWand) PaintOpaque(target, fill *PixelWand, fuzz float64) bool {
	return PaintOpaqueImage(m, target, fill, fuzz)
}

// Deprecated.
var PaintTransparentImage func(m *MagickWand, target *PixelWand, alpha, fuzz float64) bool

// Deprecated.
func (m *MagickWand) PaintTransparent(target *PixelWand, alpha, fuzz float64) bool {
	return PaintTransparentImage(m, target, alpha, fuzz)
}

// Deprecated.
var RecolorImage func(m *MagickWand, order uint32, colorMatrix *float64) bool

// Deprecated.
func (m *MagickWand) Recolor(order uint32, colorMatrix *float64) bool {
	return RecolorImage(m, order, colorMatrix)
}

// Deprecated.
var RegionOfInterestImage func(m *MagickWand, width, height uint32, x, y int32) *MagickWand

// Deprecated.
func (m *MagickWand) RegionOfInterest(width, height uint32, x, y int32) *MagickWand {
	return RegionOfInterestImage(m, width, height, x, y)
}

// Deprecated.
var DrawGetFillAlpha func(d *DrawingWand) float64

// Deprecated.
func (d *DrawingWand) FillAlpha() float64 { return DrawGetFillAlpha(d) }

// Deprecated.
var DrawGetStrokeAlpha func(d *DrawingWand) float64

// Deprecated.
func (d *DrawingWand) StrokeAlpha() float64 { return DrawGetStrokeAlpha(d) }

// Deprecated.
var DrawPeekGraphicWand func(d *DrawingWand) *DrawInfo

// Deprecated.
func (d *DrawingWand) PeekGraphicWand() *DrawInfo { return DrawPeekGraphicWand(d) }

// Deprecated.
var DrawPopGraphicContext func(d *DrawingWand) bool

// Deprecated.
func (d *DrawingWand) PopGraphicContext() bool { return DrawPopGraphicContext(d) }

// Deprecated.
var DrawPushGraphicContext func(d *DrawingWand) bool

// Deprecated.
func (d *DrawingWand) PushGraphicContext() bool { return DrawPushGraphicContext(d) }

// Deprecated.
var DrawSetFillAlpha func(d *DrawingWand, fillAlpha float64)

// Deprecated.
func (d *DrawingWand) SetFillAlpha(fillAlpha float64) { DrawSetFillAlpha(d, fillAlpha) }

// Deprecated.
var DrawSetStrokeAlpha func(d *DrawingWand, strokeAlpha float64)

// Deprecated.
func (d *DrawingWand) SetStrokeAlpha(strokeAlpha float64) { DrawSetStrokeAlpha(d, strokeAlpha) }

// Deprecated.
var GetPixelViewException func(p *PixelWand, severity *I.ExceptionType) string

// Deprecated.
func (p *PixelWand) PixelViewException(severity *I.ExceptionType) string {
	return GetPixelViewException(p, severity)
}

// Deprecated.
var ClonePixelView func(pixelView *PixelView) *PixelView

// Deprecated.
var DestroyPixelView func(pixelView *PixelView, numberWands, numberThreads uint32) *PixelView

// Deprecated.
var DuplexTransferPixelViewIterator func(source, duplex, destination *PixelView, transfer DuplexTransferPixelViewMethod, context *Void) bool

// Deprecated.
func (p *PixelView) DuplexTransferIterator(duplex, destination *PixelView, transfer DuplexTransferPixelViewMethod, context *Void) bool {
	return DuplexTransferPixelViewIterator(p, duplex, destination, transfer, context)
}

// Deprecated.
var GetPixelViewHeight func(pixelView *PixelView) uint32

// Deprecated.
func (p *PixelView) Height() uint32 { return GetPixelViewHeight(p) }

// Deprecated.
var GetPixelViewIterator func(source *PixelView, get GetPixelViewMethod, context *Void) bool

// Deprecated.
func (p *PixelView) Iterator(get GetPixelViewMethod, context *Void) bool {
	return GetPixelViewIterator(p, get, context)
}

// Deprecated.
var GetPixelViewPixels func(pixelView *PixelView) *PixelWand

// Deprecated.
func (p *PixelView) Pixels() *PixelWand { return GetPixelViewPixels(p) }

// Deprecated.
var GetPixelViewWand func(pixelView *PixelView) *MagickWand

// Deprecated.
func (p *PixelView) Wand() *MagickWand { return GetPixelViewWand(p) }

// Deprecated.
var GetPixelViewWidth func(pixelView *PixelView) uint32

// Deprecated.
func (p *PixelView) Width() uint32 { return GetPixelViewWidth(p) }

// Deprecated.
var GetPixelViewX func(pixelView *PixelView) int32

// Deprecated.
func (p *PixelView) X() int32 { return GetPixelViewX(p) }

// Deprecated.
var GetPixelViewY func(pixelView *PixelView) int32

// Deprecated.
func (p *PixelView) Y() int32 { return GetPixelViewY(p) }

// Deprecated.
var IsPixelView func(pixelView *PixelView) bool

// Deprecated.
func (p *PixelView) IsPixelView() bool { return IsPixelView(p) }

// Deprecated.
var SetPixelViewIterator func(destination *PixelView, set SetPixelViewMethod, context *Void) bool

// Deprecated.
func (p *PixelView) SetIterator(set SetPixelViewMethod, context *Void) bool {
	return SetPixelViewIterator(p, set, context)
}

// Deprecated.
var TransferPixelViewIterator func(source, destination *PixelView, transfer TransferPixelViewMethod, context *Void) bool

// Deprecated.
func (p *PixelView) TransIterator(destination *PixelView, transfer TransferPixelViewMethod, context *Void) bool {
	return TransferPixelViewIterator(p, destination, transfer, context)
}

// Deprecated.
var UpdatePixelViewIterator func(source *PixelView, update UpdatePixelViewMethod, context *Void) bool

// Deprecated.
func (p *PixelView) UpdateIterator(update UpdatePixelViewMethod, context *Void) bool {
	return UpdatePixelViewIterator(p, update, context)
}

// Deprecated.
var PixelGetNextRow func(p *PixelIterator, numberWands *uint32) []*PixelWand

// Deprecated.
func (p *PixelIterator) GetNextRow(numberWands *uint32) []*PixelWand {
	return PixelGetNextRow(p, numberWands)
}

// Deprecated.
var PixelIteratorGetException func(p *PixelIterator, severity *I.ExceptionType) string

// Deprecated.
func (p *PixelIterator) Exception(severity *I.ExceptionType) string {
	return PixelIteratorGetException(p, severity)
}

// Deprecated.
var ClipPathImage func(m *MagickWand, pathname string, inside bool) bool
