// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imgui

// #include "cimgui_wrapper.h"
import "C"
import "unsafe"

type BitVector uintptr

func (data BitVector) handle() *C.ImBitVector {
	return (*C.ImBitVector)(unsafe.Pointer(data))
}

func (data BitVector) c() C.ImBitVector {
	return *(data.handle())
}

func newBitVectorFromC(cvalue C.ImBitVector) BitVector {
	return BitVector(unsafe.Pointer(&cvalue))
}

type DrawChannel uintptr

func (data DrawChannel) handle() *C.ImDrawChannel {
	return (*C.ImDrawChannel)(unsafe.Pointer(data))
}

func (data DrawChannel) c() C.ImDrawChannel {
	return *(data.handle())
}

func newDrawChannelFromC(cvalue C.ImDrawChannel) DrawChannel {
	return DrawChannel(unsafe.Pointer(&cvalue))
}

type DrawCmd uintptr

func (data DrawCmd) handle() *C.ImDrawCmd {
	return (*C.ImDrawCmd)(unsafe.Pointer(data))
}

func (data DrawCmd) c() C.ImDrawCmd {
	return *(data.handle())
}

func newDrawCmdFromC(cvalue C.ImDrawCmd) DrawCmd {
	return DrawCmd(unsafe.Pointer(&cvalue))
}

type DrawCmdHeader uintptr

func (data DrawCmdHeader) handle() *C.ImDrawCmdHeader {
	return (*C.ImDrawCmdHeader)(unsafe.Pointer(data))
}

func (data DrawCmdHeader) c() C.ImDrawCmdHeader {
	return *(data.handle())
}

func newDrawCmdHeaderFromC(cvalue C.ImDrawCmdHeader) DrawCmdHeader {
	return DrawCmdHeader(unsafe.Pointer(&cvalue))
}

type DrawData uintptr

func (data DrawData) handle() *C.ImDrawData {
	return (*C.ImDrawData)(unsafe.Pointer(data))
}

func (data DrawData) c() C.ImDrawData {
	return *(data.handle())
}

func newDrawDataFromC(cvalue C.ImDrawData) DrawData {
	return DrawData(unsafe.Pointer(&cvalue))
}

type DrawDataBuilder uintptr

func (data DrawDataBuilder) handle() *C.ImDrawDataBuilder {
	return (*C.ImDrawDataBuilder)(unsafe.Pointer(data))
}

func (data DrawDataBuilder) c() C.ImDrawDataBuilder {
	return *(data.handle())
}

func newDrawDataBuilderFromC(cvalue C.ImDrawDataBuilder) DrawDataBuilder {
	return DrawDataBuilder(unsafe.Pointer(&cvalue))
}

type DrawList uintptr

func (data DrawList) handle() *C.ImDrawList {
	return (*C.ImDrawList)(unsafe.Pointer(data))
}

func (data DrawList) c() C.ImDrawList {
	return *(data.handle())
}

func newDrawListFromC(cvalue C.ImDrawList) DrawList {
	return DrawList(unsafe.Pointer(&cvalue))
}

type DrawListSharedData uintptr

func (data DrawListSharedData) handle() *C.ImDrawListSharedData {
	return (*C.ImDrawListSharedData)(unsafe.Pointer(data))
}

func (data DrawListSharedData) c() C.ImDrawListSharedData {
	return *(data.handle())
}

func newDrawListSharedDataFromC(cvalue C.ImDrawListSharedData) DrawListSharedData {
	return DrawListSharedData(unsafe.Pointer(&cvalue))
}

type DrawListSplitter uintptr

func (data DrawListSplitter) handle() *C.ImDrawListSplitter {
	return (*C.ImDrawListSplitter)(unsafe.Pointer(data))
}

func (data DrawListSplitter) c() C.ImDrawListSplitter {
	return *(data.handle())
}

func newDrawListSplitterFromC(cvalue C.ImDrawListSplitter) DrawListSplitter {
	return DrawListSplitter(unsafe.Pointer(&cvalue))
}

type DrawVert uintptr

func (data DrawVert) handle() *C.ImDrawVert {
	return (*C.ImDrawVert)(unsafe.Pointer(data))
}

func (data DrawVert) c() C.ImDrawVert {
	return *(data.handle())
}

func newDrawVertFromC(cvalue C.ImDrawVert) DrawVert {
	return DrawVert(unsafe.Pointer(&cvalue))
}

type Font uintptr

func (data Font) handle() *C.ImFont {
	return (*C.ImFont)(unsafe.Pointer(data))
}

func (data Font) c() C.ImFont {
	return *(data.handle())
}

func newFontFromC(cvalue C.ImFont) Font {
	return Font(unsafe.Pointer(&cvalue))
}

type FontAtlas uintptr

func (data FontAtlas) handle() *C.ImFontAtlas {
	return (*C.ImFontAtlas)(unsafe.Pointer(data))
}

func (data FontAtlas) c() C.ImFontAtlas {
	return *(data.handle())
}

func newFontAtlasFromC(cvalue C.ImFontAtlas) FontAtlas {
	return FontAtlas(unsafe.Pointer(&cvalue))
}

type FontAtlasCustomRect uintptr

func (data FontAtlasCustomRect) handle() *C.ImFontAtlasCustomRect {
	return (*C.ImFontAtlasCustomRect)(unsafe.Pointer(data))
}

func (data FontAtlasCustomRect) c() C.ImFontAtlasCustomRect {
	return *(data.handle())
}

func newFontAtlasCustomRectFromC(cvalue C.ImFontAtlasCustomRect) FontAtlasCustomRect {
	return FontAtlasCustomRect(unsafe.Pointer(&cvalue))
}

type FontBuilderIO uintptr

func (data FontBuilderIO) handle() *C.ImFontBuilderIO {
	return (*C.ImFontBuilderIO)(unsafe.Pointer(data))
}

func (data FontBuilderIO) c() C.ImFontBuilderIO {
	return *(data.handle())
}

func newFontBuilderIOFromC(cvalue C.ImFontBuilderIO) FontBuilderIO {
	return FontBuilderIO(unsafe.Pointer(&cvalue))
}

type FontConfig uintptr

func (data FontConfig) handle() *C.ImFontConfig {
	return (*C.ImFontConfig)(unsafe.Pointer(data))
}

func (data FontConfig) c() C.ImFontConfig {
	return *(data.handle())
}

func newFontConfigFromC(cvalue C.ImFontConfig) FontConfig {
	return FontConfig(unsafe.Pointer(&cvalue))
}

type FontGlyph uintptr

func (data FontGlyph) handle() *C.ImFontGlyph {
	return (*C.ImFontGlyph)(unsafe.Pointer(data))
}

func (data FontGlyph) c() C.ImFontGlyph {
	return *(data.handle())
}

func newFontGlyphFromC(cvalue C.ImFontGlyph) FontGlyph {
	return FontGlyph(unsafe.Pointer(&cvalue))
}

type FontGlyphRangesBuilder uintptr

func (data FontGlyphRangesBuilder) handle() *C.ImFontGlyphRangesBuilder {
	return (*C.ImFontGlyphRangesBuilder)(unsafe.Pointer(data))
}

func (data FontGlyphRangesBuilder) c() C.ImFontGlyphRangesBuilder {
	return *(data.handle())
}

func newFontGlyphRangesBuilderFromC(cvalue C.ImFontGlyphRangesBuilder) FontGlyphRangesBuilder {
	return FontGlyphRangesBuilder(unsafe.Pointer(&cvalue))
}

type ColorMod uintptr

func (data ColorMod) handle() *C.ImGuiColorMod {
	return (*C.ImGuiColorMod)(unsafe.Pointer(data))
}

func (data ColorMod) c() C.ImGuiColorMod {
	return *(data.handle())
}

func newColorModFromC(cvalue C.ImGuiColorMod) ColorMod {
	return ColorMod(unsafe.Pointer(&cvalue))
}

type ComboPreviewData uintptr

func (data ComboPreviewData) handle() *C.ImGuiComboPreviewData {
	return (*C.ImGuiComboPreviewData)(unsafe.Pointer(data))
}

func (data ComboPreviewData) c() C.ImGuiComboPreviewData {
	return *(data.handle())
}

func newComboPreviewDataFromC(cvalue C.ImGuiComboPreviewData) ComboPreviewData {
	return ComboPreviewData(unsafe.Pointer(&cvalue))
}

type Context uintptr

func (data Context) handle() *C.ImGuiContext {
	return (*C.ImGuiContext)(unsafe.Pointer(data))
}

func (data Context) c() C.ImGuiContext {
	return *(data.handle())
}

func newContextFromC(cvalue C.ImGuiContext) Context {
	return Context(unsafe.Pointer(&cvalue))
}

type ContextHook uintptr

func (data ContextHook) handle() *C.ImGuiContextHook {
	return (*C.ImGuiContextHook)(unsafe.Pointer(data))
}

func (data ContextHook) c() C.ImGuiContextHook {
	return *(data.handle())
}

func newContextHookFromC(cvalue C.ImGuiContextHook) ContextHook {
	return ContextHook(unsafe.Pointer(&cvalue))
}

type DataTypeInfo uintptr

func (data DataTypeInfo) handle() *C.ImGuiDataTypeInfo {
	return (*C.ImGuiDataTypeInfo)(unsafe.Pointer(data))
}

func (data DataTypeInfo) c() C.ImGuiDataTypeInfo {
	return *(data.handle())
}

func newDataTypeInfoFromC(cvalue C.ImGuiDataTypeInfo) DataTypeInfo {
	return DataTypeInfo(unsafe.Pointer(&cvalue))
}

type DataTypeTempStorage uintptr

func (data DataTypeTempStorage) handle() *C.ImGuiDataTypeTempStorage {
	return (*C.ImGuiDataTypeTempStorage)(unsafe.Pointer(data))
}

func (data DataTypeTempStorage) c() C.ImGuiDataTypeTempStorage {
	return *(data.handle())
}

func newDataTypeTempStorageFromC(cvalue C.ImGuiDataTypeTempStorage) DataTypeTempStorage {
	return DataTypeTempStorage(unsafe.Pointer(&cvalue))
}

type DockContext uintptr

func (data DockContext) handle() *C.ImGuiDockContext {
	return (*C.ImGuiDockContext)(unsafe.Pointer(data))
}

func (data DockContext) c() C.ImGuiDockContext {
	return *(data.handle())
}

func newDockContextFromC(cvalue C.ImGuiDockContext) DockContext {
	return DockContext(unsafe.Pointer(&cvalue))
}

type DockNode uintptr

func (data DockNode) handle() *C.ImGuiDockNode {
	return (*C.ImGuiDockNode)(unsafe.Pointer(data))
}

func (data DockNode) c() C.ImGuiDockNode {
	return *(data.handle())
}

func newDockNodeFromC(cvalue C.ImGuiDockNode) DockNode {
	return DockNode(unsafe.Pointer(&cvalue))
}

type GroupData uintptr

func (data GroupData) handle() *C.ImGuiGroupData {
	return (*C.ImGuiGroupData)(unsafe.Pointer(data))
}

func (data GroupData) c() C.ImGuiGroupData {
	return *(data.handle())
}

func newGroupDataFromC(cvalue C.ImGuiGroupData) GroupData {
	return GroupData(unsafe.Pointer(&cvalue))
}

type IO uintptr

func (data IO) handle() *C.ImGuiIO {
	return (*C.ImGuiIO)(unsafe.Pointer(data))
}

func (data IO) c() C.ImGuiIO {
	return *(data.handle())
}

func newIOFromC(cvalue C.ImGuiIO) IO {
	return IO(unsafe.Pointer(&cvalue))
}

type InputEvent uintptr

func (data InputEvent) handle() *C.ImGuiInputEvent {
	return (*C.ImGuiInputEvent)(unsafe.Pointer(data))
}

func (data InputEvent) c() C.ImGuiInputEvent {
	return *(data.handle())
}

func newInputEventFromC(cvalue C.ImGuiInputEvent) InputEvent {
	return InputEvent(unsafe.Pointer(&cvalue))
}

type InputEventAppFocused uintptr

func (data InputEventAppFocused) handle() *C.ImGuiInputEventAppFocused {
	return (*C.ImGuiInputEventAppFocused)(unsafe.Pointer(data))
}

func (data InputEventAppFocused) c() C.ImGuiInputEventAppFocused {
	return *(data.handle())
}

func newInputEventAppFocusedFromC(cvalue C.ImGuiInputEventAppFocused) InputEventAppFocused {
	return InputEventAppFocused(unsafe.Pointer(&cvalue))
}

type InputEventKey uintptr

func (data InputEventKey) handle() *C.ImGuiInputEventKey {
	return (*C.ImGuiInputEventKey)(unsafe.Pointer(data))
}

func (data InputEventKey) c() C.ImGuiInputEventKey {
	return *(data.handle())
}

func newInputEventKeyFromC(cvalue C.ImGuiInputEventKey) InputEventKey {
	return InputEventKey(unsafe.Pointer(&cvalue))
}

type InputEventMouseButton uintptr

func (data InputEventMouseButton) handle() *C.ImGuiInputEventMouseButton {
	return (*C.ImGuiInputEventMouseButton)(unsafe.Pointer(data))
}

func (data InputEventMouseButton) c() C.ImGuiInputEventMouseButton {
	return *(data.handle())
}

func newInputEventMouseButtonFromC(cvalue C.ImGuiInputEventMouseButton) InputEventMouseButton {
	return InputEventMouseButton(unsafe.Pointer(&cvalue))
}

type InputEventMousePos uintptr

func (data InputEventMousePos) handle() *C.ImGuiInputEventMousePos {
	return (*C.ImGuiInputEventMousePos)(unsafe.Pointer(data))
}

func (data InputEventMousePos) c() C.ImGuiInputEventMousePos {
	return *(data.handle())
}

func newInputEventMousePosFromC(cvalue C.ImGuiInputEventMousePos) InputEventMousePos {
	return InputEventMousePos(unsafe.Pointer(&cvalue))
}

type InputEventMouseViewport uintptr

func (data InputEventMouseViewport) handle() *C.ImGuiInputEventMouseViewport {
	return (*C.ImGuiInputEventMouseViewport)(unsafe.Pointer(data))
}

func (data InputEventMouseViewport) c() C.ImGuiInputEventMouseViewport {
	return *(data.handle())
}

func newInputEventMouseViewportFromC(cvalue C.ImGuiInputEventMouseViewport) InputEventMouseViewport {
	return InputEventMouseViewport(unsafe.Pointer(&cvalue))
}

type InputEventMouseWheel uintptr

func (data InputEventMouseWheel) handle() *C.ImGuiInputEventMouseWheel {
	return (*C.ImGuiInputEventMouseWheel)(unsafe.Pointer(data))
}

func (data InputEventMouseWheel) c() C.ImGuiInputEventMouseWheel {
	return *(data.handle())
}

func newInputEventMouseWheelFromC(cvalue C.ImGuiInputEventMouseWheel) InputEventMouseWheel {
	return InputEventMouseWheel(unsafe.Pointer(&cvalue))
}

type InputEventText uintptr

func (data InputEventText) handle() *C.ImGuiInputEventText {
	return (*C.ImGuiInputEventText)(unsafe.Pointer(data))
}

func (data InputEventText) c() C.ImGuiInputEventText {
	return *(data.handle())
}

func newInputEventTextFromC(cvalue C.ImGuiInputEventText) InputEventText {
	return InputEventText(unsafe.Pointer(&cvalue))
}

type InputTextCallbackData uintptr

func (data InputTextCallbackData) handle() *C.ImGuiInputTextCallbackData {
	return (*C.ImGuiInputTextCallbackData)(unsafe.Pointer(data))
}

func (data InputTextCallbackData) c() C.ImGuiInputTextCallbackData {
	return *(data.handle())
}

func newInputTextCallbackDataFromC(cvalue C.ImGuiInputTextCallbackData) InputTextCallbackData {
	return InputTextCallbackData(unsafe.Pointer(&cvalue))
}

type InputTextState uintptr

func (data InputTextState) handle() *C.ImGuiInputTextState {
	return (*C.ImGuiInputTextState)(unsafe.Pointer(data))
}

func (data InputTextState) c() C.ImGuiInputTextState {
	return *(data.handle())
}

func newInputTextStateFromC(cvalue C.ImGuiInputTextState) InputTextState {
	return InputTextState(unsafe.Pointer(&cvalue))
}

type KeyData uintptr

func (data KeyData) handle() *C.ImGuiKeyData {
	return (*C.ImGuiKeyData)(unsafe.Pointer(data))
}

func (data KeyData) c() C.ImGuiKeyData {
	return *(data.handle())
}

func newKeyDataFromC(cvalue C.ImGuiKeyData) KeyData {
	return KeyData(unsafe.Pointer(&cvalue))
}

type LastItemData uintptr

func (data LastItemData) handle() *C.ImGuiLastItemData {
	return (*C.ImGuiLastItemData)(unsafe.Pointer(data))
}

func (data LastItemData) c() C.ImGuiLastItemData {
	return *(data.handle())
}

func newLastItemDataFromC(cvalue C.ImGuiLastItemData) LastItemData {
	return LastItemData(unsafe.Pointer(&cvalue))
}

type ListClipper uintptr

func (data ListClipper) handle() *C.ImGuiListClipper {
	return (*C.ImGuiListClipper)(unsafe.Pointer(data))
}

func (data ListClipper) c() C.ImGuiListClipper {
	return *(data.handle())
}

func newListClipperFromC(cvalue C.ImGuiListClipper) ListClipper {
	return ListClipper(unsafe.Pointer(&cvalue))
}

type ListClipperData uintptr

func (data ListClipperData) handle() *C.ImGuiListClipperData {
	return (*C.ImGuiListClipperData)(unsafe.Pointer(data))
}

func (data ListClipperData) c() C.ImGuiListClipperData {
	return *(data.handle())
}

func newListClipperDataFromC(cvalue C.ImGuiListClipperData) ListClipperData {
	return ListClipperData(unsafe.Pointer(&cvalue))
}

type ListClipperRange uintptr

func (data ListClipperRange) handle() *C.ImGuiListClipperRange {
	return (*C.ImGuiListClipperRange)(unsafe.Pointer(data))
}

func (data ListClipperRange) c() C.ImGuiListClipperRange {
	return *(data.handle())
}

func newListClipperRangeFromC(cvalue C.ImGuiListClipperRange) ListClipperRange {
	return ListClipperRange(unsafe.Pointer(&cvalue))
}

type MenuColumns uintptr

func (data MenuColumns) handle() *C.ImGuiMenuColumns {
	return (*C.ImGuiMenuColumns)(unsafe.Pointer(data))
}

func (data MenuColumns) c() C.ImGuiMenuColumns {
	return *(data.handle())
}

func newMenuColumnsFromC(cvalue C.ImGuiMenuColumns) MenuColumns {
	return MenuColumns(unsafe.Pointer(&cvalue))
}

type MetricsConfig uintptr

func (data MetricsConfig) handle() *C.ImGuiMetricsConfig {
	return (*C.ImGuiMetricsConfig)(unsafe.Pointer(data))
}

func (data MetricsConfig) c() C.ImGuiMetricsConfig {
	return *(data.handle())
}

func newMetricsConfigFromC(cvalue C.ImGuiMetricsConfig) MetricsConfig {
	return MetricsConfig(unsafe.Pointer(&cvalue))
}

type NavItemData uintptr

func (data NavItemData) handle() *C.ImGuiNavItemData {
	return (*C.ImGuiNavItemData)(unsafe.Pointer(data))
}

func (data NavItemData) c() C.ImGuiNavItemData {
	return *(data.handle())
}

func newNavItemDataFromC(cvalue C.ImGuiNavItemData) NavItemData {
	return NavItemData(unsafe.Pointer(&cvalue))
}

type NextItemData uintptr

func (data NextItemData) handle() *C.ImGuiNextItemData {
	return (*C.ImGuiNextItemData)(unsafe.Pointer(data))
}

func (data NextItemData) c() C.ImGuiNextItemData {
	return *(data.handle())
}

func newNextItemDataFromC(cvalue C.ImGuiNextItemData) NextItemData {
	return NextItemData(unsafe.Pointer(&cvalue))
}

type NextWindowData uintptr

func (data NextWindowData) handle() *C.ImGuiNextWindowData {
	return (*C.ImGuiNextWindowData)(unsafe.Pointer(data))
}

func (data NextWindowData) c() C.ImGuiNextWindowData {
	return *(data.handle())
}

func newNextWindowDataFromC(cvalue C.ImGuiNextWindowData) NextWindowData {
	return NextWindowData(unsafe.Pointer(&cvalue))
}

type OldColumnData uintptr

func (data OldColumnData) handle() *C.ImGuiOldColumnData {
	return (*C.ImGuiOldColumnData)(unsafe.Pointer(data))
}

func (data OldColumnData) c() C.ImGuiOldColumnData {
	return *(data.handle())
}

func newOldColumnDataFromC(cvalue C.ImGuiOldColumnData) OldColumnData {
	return OldColumnData(unsafe.Pointer(&cvalue))
}

type OldColumns uintptr

func (data OldColumns) handle() *C.ImGuiOldColumns {
	return (*C.ImGuiOldColumns)(unsafe.Pointer(data))
}

func (data OldColumns) c() C.ImGuiOldColumns {
	return *(data.handle())
}

func newOldColumnsFromC(cvalue C.ImGuiOldColumns) OldColumns {
	return OldColumns(unsafe.Pointer(&cvalue))
}

type OnceUponAFrame uintptr

func (data OnceUponAFrame) handle() *C.ImGuiOnceUponAFrame {
	return (*C.ImGuiOnceUponAFrame)(unsafe.Pointer(data))
}

func (data OnceUponAFrame) c() C.ImGuiOnceUponAFrame {
	return *(data.handle())
}

func newOnceUponAFrameFromC(cvalue C.ImGuiOnceUponAFrame) OnceUponAFrame {
	return OnceUponAFrame(unsafe.Pointer(&cvalue))
}

type Payload uintptr

func (data Payload) handle() *C.ImGuiPayload {
	return (*C.ImGuiPayload)(unsafe.Pointer(data))
}

func (data Payload) c() C.ImGuiPayload {
	return *(data.handle())
}

func newPayloadFromC(cvalue C.ImGuiPayload) Payload {
	return Payload(unsafe.Pointer(&cvalue))
}

type PlatformIO uintptr

func (data PlatformIO) handle() *C.ImGuiPlatformIO {
	return (*C.ImGuiPlatformIO)(unsafe.Pointer(data))
}

func (data PlatformIO) c() C.ImGuiPlatformIO {
	return *(data.handle())
}

func newPlatformIOFromC(cvalue C.ImGuiPlatformIO) PlatformIO {
	return PlatformIO(unsafe.Pointer(&cvalue))
}

type PlatformImeData uintptr

func (data PlatformImeData) handle() *C.ImGuiPlatformImeData {
	return (*C.ImGuiPlatformImeData)(unsafe.Pointer(data))
}

func (data PlatformImeData) c() C.ImGuiPlatformImeData {
	return *(data.handle())
}

func newPlatformImeDataFromC(cvalue C.ImGuiPlatformImeData) PlatformImeData {
	return PlatformImeData(unsafe.Pointer(&cvalue))
}

type PlatformMonitor uintptr

func (data PlatformMonitor) handle() *C.ImGuiPlatformMonitor {
	return (*C.ImGuiPlatformMonitor)(unsafe.Pointer(data))
}

func (data PlatformMonitor) c() C.ImGuiPlatformMonitor {
	return *(data.handle())
}

func newPlatformMonitorFromC(cvalue C.ImGuiPlatformMonitor) PlatformMonitor {
	return PlatformMonitor(unsafe.Pointer(&cvalue))
}

type PopupData uintptr

func (data PopupData) handle() *C.ImGuiPopupData {
	return (*C.ImGuiPopupData)(unsafe.Pointer(data))
}

func (data PopupData) c() C.ImGuiPopupData {
	return *(data.handle())
}

func newPopupDataFromC(cvalue C.ImGuiPopupData) PopupData {
	return PopupData(unsafe.Pointer(&cvalue))
}

type PtrOrIndex uintptr

func (data PtrOrIndex) handle() *C.ImGuiPtrOrIndex {
	return (*C.ImGuiPtrOrIndex)(unsafe.Pointer(data))
}

func (data PtrOrIndex) c() C.ImGuiPtrOrIndex {
	return *(data.handle())
}

func newPtrOrIndexFromC(cvalue C.ImGuiPtrOrIndex) PtrOrIndex {
	return PtrOrIndex(unsafe.Pointer(&cvalue))
}

type SettingsHandler uintptr

func (data SettingsHandler) handle() *C.ImGuiSettingsHandler {
	return (*C.ImGuiSettingsHandler)(unsafe.Pointer(data))
}

func (data SettingsHandler) c() C.ImGuiSettingsHandler {
	return *(data.handle())
}

func newSettingsHandlerFromC(cvalue C.ImGuiSettingsHandler) SettingsHandler {
	return SettingsHandler(unsafe.Pointer(&cvalue))
}

type ShrinkWidthItem uintptr

func (data ShrinkWidthItem) handle() *C.ImGuiShrinkWidthItem {
	return (*C.ImGuiShrinkWidthItem)(unsafe.Pointer(data))
}

func (data ShrinkWidthItem) c() C.ImGuiShrinkWidthItem {
	return *(data.handle())
}

func newShrinkWidthItemFromC(cvalue C.ImGuiShrinkWidthItem) ShrinkWidthItem {
	return ShrinkWidthItem(unsafe.Pointer(&cvalue))
}

type SizeCallbackData uintptr

func (data SizeCallbackData) handle() *C.ImGuiSizeCallbackData {
	return (*C.ImGuiSizeCallbackData)(unsafe.Pointer(data))
}

func (data SizeCallbackData) c() C.ImGuiSizeCallbackData {
	return *(data.handle())
}

func newSizeCallbackDataFromC(cvalue C.ImGuiSizeCallbackData) SizeCallbackData {
	return SizeCallbackData(unsafe.Pointer(&cvalue))
}

type StackLevelInfo uintptr

func (data StackLevelInfo) handle() *C.ImGuiStackLevelInfo {
	return (*C.ImGuiStackLevelInfo)(unsafe.Pointer(data))
}

func (data StackLevelInfo) c() C.ImGuiStackLevelInfo {
	return *(data.handle())
}

func newStackLevelInfoFromC(cvalue C.ImGuiStackLevelInfo) StackLevelInfo {
	return StackLevelInfo(unsafe.Pointer(&cvalue))
}

type StackSizes uintptr

func (data StackSizes) handle() *C.ImGuiStackSizes {
	return (*C.ImGuiStackSizes)(unsafe.Pointer(data))
}

func (data StackSizes) c() C.ImGuiStackSizes {
	return *(data.handle())
}

func newStackSizesFromC(cvalue C.ImGuiStackSizes) StackSizes {
	return StackSizes(unsafe.Pointer(&cvalue))
}

type StackTool uintptr

func (data StackTool) handle() *C.ImGuiStackTool {
	return (*C.ImGuiStackTool)(unsafe.Pointer(data))
}

func (data StackTool) c() C.ImGuiStackTool {
	return *(data.handle())
}

func newStackToolFromC(cvalue C.ImGuiStackTool) StackTool {
	return StackTool(unsafe.Pointer(&cvalue))
}

type Storage uintptr

func (data Storage) handle() *C.ImGuiStorage {
	return (*C.ImGuiStorage)(unsafe.Pointer(data))
}

func (data Storage) c() C.ImGuiStorage {
	return *(data.handle())
}

func newStorageFromC(cvalue C.ImGuiStorage) Storage {
	return Storage(unsafe.Pointer(&cvalue))
}

type StoragePair uintptr

func (data StoragePair) handle() *C.ImGuiStoragePair {
	return (*C.ImGuiStoragePair)(unsafe.Pointer(data))
}

func (data StoragePair) c() C.ImGuiStoragePair {
	return *(data.handle())
}

func newStoragePairFromC(cvalue C.ImGuiStoragePair) StoragePair {
	return StoragePair(unsafe.Pointer(&cvalue))
}

type Style uintptr

func (data Style) handle() *C.ImGuiStyle {
	return (*C.ImGuiStyle)(unsafe.Pointer(data))
}

func (data Style) c() C.ImGuiStyle {
	return *(data.handle())
}

func newStyleFromC(cvalue C.ImGuiStyle) Style {
	return Style(unsafe.Pointer(&cvalue))
}

type StyleMod uintptr

func (data StyleMod) handle() *C.ImGuiStyleMod {
	return (*C.ImGuiStyleMod)(unsafe.Pointer(data))
}

func (data StyleMod) c() C.ImGuiStyleMod {
	return *(data.handle())
}

func newStyleModFromC(cvalue C.ImGuiStyleMod) StyleMod {
	return StyleMod(unsafe.Pointer(&cvalue))
}

type TabBar uintptr

func (data TabBar) handle() *C.ImGuiTabBar {
	return (*C.ImGuiTabBar)(unsafe.Pointer(data))
}

func (data TabBar) c() C.ImGuiTabBar {
	return *(data.handle())
}

func newTabBarFromC(cvalue C.ImGuiTabBar) TabBar {
	return TabBar(unsafe.Pointer(&cvalue))
}

type TabItem uintptr

func (data TabItem) handle() *C.ImGuiTabItem {
	return (*C.ImGuiTabItem)(unsafe.Pointer(data))
}

func (data TabItem) c() C.ImGuiTabItem {
	return *(data.handle())
}

func newTabItemFromC(cvalue C.ImGuiTabItem) TabItem {
	return TabItem(unsafe.Pointer(&cvalue))
}

type Table uintptr

func (data Table) handle() *C.ImGuiTable {
	return (*C.ImGuiTable)(unsafe.Pointer(data))
}

func (data Table) c() C.ImGuiTable {
	return *(data.handle())
}

func newTableFromC(cvalue C.ImGuiTable) Table {
	return Table(unsafe.Pointer(&cvalue))
}

type TableCellData uintptr

func (data TableCellData) handle() *C.ImGuiTableCellData {
	return (*C.ImGuiTableCellData)(unsafe.Pointer(data))
}

func (data TableCellData) c() C.ImGuiTableCellData {
	return *(data.handle())
}

func newTableCellDataFromC(cvalue C.ImGuiTableCellData) TableCellData {
	return TableCellData(unsafe.Pointer(&cvalue))
}

type TableColumn uintptr

func (data TableColumn) handle() *C.ImGuiTableColumn {
	return (*C.ImGuiTableColumn)(unsafe.Pointer(data))
}

func (data TableColumn) c() C.ImGuiTableColumn {
	return *(data.handle())
}

func newTableColumnFromC(cvalue C.ImGuiTableColumn) TableColumn {
	return TableColumn(unsafe.Pointer(&cvalue))
}

type TableColumnSettings uintptr

func (data TableColumnSettings) handle() *C.ImGuiTableColumnSettings {
	return (*C.ImGuiTableColumnSettings)(unsafe.Pointer(data))
}

func (data TableColumnSettings) c() C.ImGuiTableColumnSettings {
	return *(data.handle())
}

func newTableColumnSettingsFromC(cvalue C.ImGuiTableColumnSettings) TableColumnSettings {
	return TableColumnSettings(unsafe.Pointer(&cvalue))
}

type TableColumnSortSpecs uintptr

func (data TableColumnSortSpecs) handle() *C.ImGuiTableColumnSortSpecs {
	return (*C.ImGuiTableColumnSortSpecs)(unsafe.Pointer(data))
}

func (data TableColumnSortSpecs) c() C.ImGuiTableColumnSortSpecs {
	return *(data.handle())
}

func newTableColumnSortSpecsFromC(cvalue C.ImGuiTableColumnSortSpecs) TableColumnSortSpecs {
	return TableColumnSortSpecs(unsafe.Pointer(&cvalue))
}

type TableInstanceData uintptr

func (data TableInstanceData) handle() *C.ImGuiTableInstanceData {
	return (*C.ImGuiTableInstanceData)(unsafe.Pointer(data))
}

func (data TableInstanceData) c() C.ImGuiTableInstanceData {
	return *(data.handle())
}

func newTableInstanceDataFromC(cvalue C.ImGuiTableInstanceData) TableInstanceData {
	return TableInstanceData(unsafe.Pointer(&cvalue))
}

type TableSettings uintptr

func (data TableSettings) handle() *C.ImGuiTableSettings {
	return (*C.ImGuiTableSettings)(unsafe.Pointer(data))
}

func (data TableSettings) c() C.ImGuiTableSettings {
	return *(data.handle())
}

func newTableSettingsFromC(cvalue C.ImGuiTableSettings) TableSettings {
	return TableSettings(unsafe.Pointer(&cvalue))
}

type TableSortSpecs uintptr

func (data TableSortSpecs) handle() *C.ImGuiTableSortSpecs {
	return (*C.ImGuiTableSortSpecs)(unsafe.Pointer(data))
}

func (data TableSortSpecs) c() C.ImGuiTableSortSpecs {
	return *(data.handle())
}

func newTableSortSpecsFromC(cvalue C.ImGuiTableSortSpecs) TableSortSpecs {
	return TableSortSpecs(unsafe.Pointer(&cvalue))
}

type TableTempData uintptr

func (data TableTempData) handle() *C.ImGuiTableTempData {
	return (*C.ImGuiTableTempData)(unsafe.Pointer(data))
}

func (data TableTempData) c() C.ImGuiTableTempData {
	return *(data.handle())
}

func newTableTempDataFromC(cvalue C.ImGuiTableTempData) TableTempData {
	return TableTempData(unsafe.Pointer(&cvalue))
}

type TextBuffer uintptr

func (data TextBuffer) handle() *C.ImGuiTextBuffer {
	return (*C.ImGuiTextBuffer)(unsafe.Pointer(data))
}

func (data TextBuffer) c() C.ImGuiTextBuffer {
	return *(data.handle())
}

func newTextBufferFromC(cvalue C.ImGuiTextBuffer) TextBuffer {
	return TextBuffer(unsafe.Pointer(&cvalue))
}

type TextFilter uintptr

func (data TextFilter) handle() *C.ImGuiTextFilter {
	return (*C.ImGuiTextFilter)(unsafe.Pointer(data))
}

func (data TextFilter) c() C.ImGuiTextFilter {
	return *(data.handle())
}

func newTextFilterFromC(cvalue C.ImGuiTextFilter) TextFilter {
	return TextFilter(unsafe.Pointer(&cvalue))
}

type TextRange uintptr

func (data TextRange) handle() *C.ImGuiTextRange {
	return (*C.ImGuiTextRange)(unsafe.Pointer(data))
}

func (data TextRange) c() C.ImGuiTextRange {
	return *(data.handle())
}

func newTextRangeFromC(cvalue C.ImGuiTextRange) TextRange {
	return TextRange(unsafe.Pointer(&cvalue))
}

type Viewport uintptr

func (data Viewport) handle() *C.ImGuiViewport {
	return (*C.ImGuiViewport)(unsafe.Pointer(data))
}

func (data Viewport) c() C.ImGuiViewport {
	return *(data.handle())
}

func newViewportFromC(cvalue C.ImGuiViewport) Viewport {
	return Viewport(unsafe.Pointer(&cvalue))
}

type ViewportP uintptr

func (data ViewportP) handle() *C.ImGuiViewportP {
	return (*C.ImGuiViewportP)(unsafe.Pointer(data))
}

func (data ViewportP) c() C.ImGuiViewportP {
	return *(data.handle())
}

func newViewportPFromC(cvalue C.ImGuiViewportP) ViewportP {
	return ViewportP(unsafe.Pointer(&cvalue))
}

type Window uintptr

func (data Window) handle() *C.ImGuiWindow {
	return (*C.ImGuiWindow)(unsafe.Pointer(data))
}

func (data Window) c() C.ImGuiWindow {
	return *(data.handle())
}

func newWindowFromC(cvalue C.ImGuiWindow) Window {
	return Window(unsafe.Pointer(&cvalue))
}

type WindowClass uintptr

func (data WindowClass) handle() *C.ImGuiWindowClass {
	return (*C.ImGuiWindowClass)(unsafe.Pointer(data))
}

func (data WindowClass) c() C.ImGuiWindowClass {
	return *(data.handle())
}

func newWindowClassFromC(cvalue C.ImGuiWindowClass) WindowClass {
	return WindowClass(unsafe.Pointer(&cvalue))
}

type WindowDockStyle uintptr

func (data WindowDockStyle) handle() *C.ImGuiWindowDockStyle {
	return (*C.ImGuiWindowDockStyle)(unsafe.Pointer(data))
}

func (data WindowDockStyle) c() C.ImGuiWindowDockStyle {
	return *(data.handle())
}

func newWindowDockStyleFromC(cvalue C.ImGuiWindowDockStyle) WindowDockStyle {
	return WindowDockStyle(unsafe.Pointer(&cvalue))
}

type WindowSettings uintptr

func (data WindowSettings) handle() *C.ImGuiWindowSettings {
	return (*C.ImGuiWindowSettings)(unsafe.Pointer(data))
}

func (data WindowSettings) c() C.ImGuiWindowSettings {
	return *(data.handle())
}

func newWindowSettingsFromC(cvalue C.ImGuiWindowSettings) WindowSettings {
	return WindowSettings(unsafe.Pointer(&cvalue))
}

type WindowStackData uintptr

func (data WindowStackData) handle() *C.ImGuiWindowStackData {
	return (*C.ImGuiWindowStackData)(unsafe.Pointer(data))
}

func (data WindowStackData) c() C.ImGuiWindowStackData {
	return *(data.handle())
}

func newWindowStackDataFromC(cvalue C.ImGuiWindowStackData) WindowStackData {
	return WindowStackData(unsafe.Pointer(&cvalue))
}

type WindowTempData uintptr

func (data WindowTempData) handle() *C.ImGuiWindowTempData {
	return (*C.ImGuiWindowTempData)(unsafe.Pointer(data))
}

func (data WindowTempData) c() C.ImGuiWindowTempData {
	return *(data.handle())
}

func newWindowTempDataFromC(cvalue C.ImGuiWindowTempData) WindowTempData {
	return WindowTempData(unsafe.Pointer(&cvalue))
}
