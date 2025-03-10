// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imgui

// #include "extra_types.h"
// #include "cimnodes_structs_accessor.h"
// #include "cimnodes_wrapper.h"
import "C"

func NewEmulateThreeButtonMouse() *EmulateThreeButtonMouse {
	return newEmulateThreeButtonMouseFromC(C.EmulateThreeButtonMouse_EmulateThreeButtonMouse())
}

func (self *EmulateThreeButtonMouse) Destroy() {
	selfArg, selfFin := self.handle()
	C.EmulateThreeButtonMouse_destroy(selfArg)

	selfFin()
}

func NewNodesIO() *NodesIO {
	return newNodesIOFromC(C.ImNodesIO_ImNodesIO())
}

func (self *NodesIO) Destroy() {
	selfArg, selfFin := self.handle()
	C.ImNodesIO_destroy(selfArg)

	selfFin()
}

func NewNodesStyle() *NodesStyle {
	return newNodesStyleFromC(C.ImNodesStyle_ImNodesStyle())
}

func (self *NodesStyle) Destroy() {
	selfArg, selfFin := self.handle()
	C.ImNodesStyle_destroy(selfArg)

	selfFin()
}

func NewLinkDetachWithModifierClick() *LinkDetachWithModifierClick {
	return newLinkDetachWithModifierClickFromC(C.LinkDetachWithModifierClick_LinkDetachWithModifierClick())
}

func (self *LinkDetachWithModifierClick) Destroy() {
	selfArg, selfFin := self.handle()
	C.LinkDetachWithModifierClick_destroy(selfArg)

	selfFin()
}

func NewMultipleSelectModifier() *MultipleSelectModifier {
	return newMultipleSelectModifierFromC(C.MultipleSelectModifier_MultipleSelectModifier())
}

func (self *MultipleSelectModifier) Destroy() {
	selfArg, selfFin := self.handle()
	C.MultipleSelectModifier_destroy(selfArg)

	selfFin()
}

// imnodesBeginInputAttributeV parameter default value hint:
// shape: ImNodesPinShape_CircleFilled
func imnodesBeginInputAttributeV(id int32, shape NodesPinShape) {
	C.imnodes_BeginInputAttribute(C.int(id), C.ImNodesPinShape(shape))
}

func imnodesBeginNode(id int32) {
	C.imnodes_BeginNode(C.int(id))
}

func imnodesBeginNodeEditor() {
	C.imnodes_BeginNodeEditor()
}

func imnodesBeginNodeTitleBar() {
	C.imnodes_BeginNodeTitleBar()
}

// imnodesBeginOutputAttributeV parameter default value hint:
// shape: ImNodesPinShape_CircleFilled
func imnodesBeginOutputAttributeV(id int32, shape NodesPinShape) {
	C.imnodes_BeginOutputAttribute(C.int(id), C.ImNodesPinShape(shape))
}

func imnodesBeginStaticAttribute(id int32) {
	C.imnodes_BeginStaticAttribute(C.int(id))
}

func imnodesClearLinkSelectionInt(link_id int32) {
	C.imnodes_ClearLinkSelection_Int(C.int(link_id))
}

func imnodesClearLinkSelectionNil() {
	C.imnodes_ClearLinkSelection_Nil()
}

func imnodesClearNodeSelectionInt(node_id int32) {
	C.imnodes_ClearNodeSelection_Int(C.int(node_id))
}

func imnodesClearNodeSelectionNil() {
	C.imnodes_ClearNodeSelection_Nil()
}

func imnodesEditorContextGetPanning() Vec2 {
	pOut := new(Vec2)
	pOutArg, pOutFin := wrap[C.ImVec2, *Vec2](pOut)

	C.imnodes_EditorContextGetPanning(pOutArg)

	pOutFin()

	return *pOut
}

func imnodesEditorContextMoveToNode(node_id int32) {
	C.imnodes_EditorContextMoveToNode(C.int(node_id))
}

func imnodesEditorContextResetPanning(pos Vec2) {
	C.imnodes_EditorContextResetPanning(pos.toC())
}

func imnodesEndInputAttribute() {
	C.imnodes_EndInputAttribute()
}

func imnodesEndNode() {
	C.imnodes_EndNode()
}

func imnodesEndNodeEditor() {
	C.imnodes_EndNodeEditor()
}

func imnodesEndNodeTitleBar() {
	C.imnodes_EndNodeTitleBar()
}

func imnodesEndOutputAttribute() {
	C.imnodes_EndOutputAttribute()
}

func imnodesEndStaticAttribute() {
	C.imnodes_EndStaticAttribute()
}

func imnodesGetIO() *NodesIO {
	return newNodesIOFromC(C.imnodes_GetIO())
}

func imnodesGetNodeDimensions(id int32) Vec2 {
	pOut := new(Vec2)
	pOutArg, pOutFin := wrap[C.ImVec2, *Vec2](pOut)

	C.imnodes_GetNodeDimensions(pOutArg, C.int(id))

	pOutFin()

	return *pOut
}

func imnodesGetNodeEditorSpacePos(node_id int32) Vec2 {
	pOut := new(Vec2)
	pOutArg, pOutFin := wrap[C.ImVec2, *Vec2](pOut)

	C.imnodes_GetNodeEditorSpacePos(pOutArg, C.int(node_id))

	pOutFin()

	return *pOut
}

func imnodesGetNodeGridSpacePos(node_id int32) Vec2 {
	pOut := new(Vec2)
	pOutArg, pOutFin := wrap[C.ImVec2, *Vec2](pOut)

	C.imnodes_GetNodeGridSpacePos(pOutArg, C.int(node_id))

	pOutFin()

	return *pOut
}

func imnodesGetNodeScreenSpacePos(node_id int32) Vec2 {
	pOut := new(Vec2)
	pOutArg, pOutFin := wrap[C.ImVec2, *Vec2](pOut)

	C.imnodes_GetNodeScreenSpacePos(pOutArg, C.int(node_id))

	pOutFin()

	return *pOut
}

func imnodesGetSelectedLinks(link_ids *int32) {
	link_idsArg, link_idsFin := WrapNumberPtr[C.int, int32](link_ids)
	C.imnodes_GetSelectedLinks(link_idsArg)

	link_idsFin()
}

func imnodesGetSelectedNodes(node_ids *int32) {
	node_idsArg, node_idsFin := WrapNumberPtr[C.int, int32](node_ids)
	C.imnodes_GetSelectedNodes(node_idsArg)

	node_idsFin()
}

func imnodesGetStyle() *NodesStyle {
	return newNodesStyleFromC(C.imnodes_GetStyle())
}

// imnodesIsAnyAttributeActiveV parameter default value hint:
// attribute_id: NULL
func imnodesIsAnyAttributeActiveV(attribute_id *int32) bool {
	attribute_idArg, attribute_idFin := WrapNumberPtr[C.int, int32](attribute_id)

	defer func() {
		attribute_idFin()
	}()
	return C.imnodes_IsAnyAttributeActive(attribute_idArg) == C.bool(true)
}

func imnodesIsAttributeActive() bool {
	return C.imnodes_IsAttributeActive() == C.bool(true)
}

func imnodesIsEditorHovered() bool {
	return C.imnodes_IsEditorHovered() == C.bool(true)
}

// imnodesIsLinkCreatedBoolPtrV parameter default value hint:
// created_from_snap: NULL
func imnodesIsLinkCreatedBoolPtrV(started_at_attribute_id *int32, ended_at_attribute_id *int32, created_from_snap *bool) bool {
	started_at_attribute_idArg, started_at_attribute_idFin := WrapNumberPtr[C.int, int32](started_at_attribute_id)
	ended_at_attribute_idArg, ended_at_attribute_idFin := WrapNumberPtr[C.int, int32](ended_at_attribute_id)
	created_from_snapArg, created_from_snapFin := WrapBool(created_from_snap)

	defer func() {
		started_at_attribute_idFin()
		ended_at_attribute_idFin()
		created_from_snapFin()
	}()
	return C.imnodes_IsLinkCreated_BoolPtr(started_at_attribute_idArg, ended_at_attribute_idArg, created_from_snapArg) == C.bool(true)
}

// imnodesIsLinkCreatedIntPtrV parameter default value hint:
// created_from_snap: NULL
func imnodesIsLinkCreatedIntPtrV(started_at_node_id *int32, started_at_attribute_id *int32, ended_at_node_id *int32, ended_at_attribute_id *int32, created_from_snap *bool) bool {
	started_at_node_idArg, started_at_node_idFin := WrapNumberPtr[C.int, int32](started_at_node_id)
	started_at_attribute_idArg, started_at_attribute_idFin := WrapNumberPtr[C.int, int32](started_at_attribute_id)
	ended_at_node_idArg, ended_at_node_idFin := WrapNumberPtr[C.int, int32](ended_at_node_id)
	ended_at_attribute_idArg, ended_at_attribute_idFin := WrapNumberPtr[C.int, int32](ended_at_attribute_id)
	created_from_snapArg, created_from_snapFin := WrapBool(created_from_snap)

	defer func() {
		started_at_node_idFin()
		started_at_attribute_idFin()
		ended_at_node_idFin()
		ended_at_attribute_idFin()
		created_from_snapFin()
	}()
	return C.imnodes_IsLinkCreated_IntPtr(started_at_node_idArg, started_at_attribute_idArg, ended_at_node_idArg, ended_at_attribute_idArg, created_from_snapArg) == C.bool(true)
}

func imnodesIsLinkDestroyed(link_id *int32) bool {
	link_idArg, link_idFin := WrapNumberPtr[C.int, int32](link_id)

	defer func() {
		link_idFin()
	}()
	return C.imnodes_IsLinkDestroyed(link_idArg) == C.bool(true)
}

// imnodesIsLinkDroppedV parameter default value hint:
// started_at_attribute_id: NULL
// including_detached_links: true
func imnodesIsLinkDroppedV(started_at_attribute_id *int32, including_detached_links bool) bool {
	started_at_attribute_idArg, started_at_attribute_idFin := WrapNumberPtr[C.int, int32](started_at_attribute_id)

	defer func() {
		started_at_attribute_idFin()
	}()
	return C.imnodes_IsLinkDropped(started_at_attribute_idArg, C.bool(including_detached_links)) == C.bool(true)
}

func imnodesIsLinkHovered(link_id *int32) bool {
	link_idArg, link_idFin := WrapNumberPtr[C.int, int32](link_id)

	defer func() {
		link_idFin()
	}()
	return C.imnodes_IsLinkHovered(link_idArg) == C.bool(true)
}

func imnodesIsLinkSelected(link_id int32) bool {
	return C.imnodes_IsLinkSelected(C.int(link_id)) == C.bool(true)
}

func imnodesIsLinkStarted(started_at_attribute_id *int32) bool {
	started_at_attribute_idArg, started_at_attribute_idFin := WrapNumberPtr[C.int, int32](started_at_attribute_id)

	defer func() {
		started_at_attribute_idFin()
	}()
	return C.imnodes_IsLinkStarted(started_at_attribute_idArg) == C.bool(true)
}

func imnodesIsNodeHovered(node_id *int32) bool {
	node_idArg, node_idFin := WrapNumberPtr[C.int, int32](node_id)

	defer func() {
		node_idFin()
	}()
	return C.imnodes_IsNodeHovered(node_idArg) == C.bool(true)
}

func imnodesIsNodeSelected(node_id int32) bool {
	return C.imnodes_IsNodeSelected(C.int(node_id)) == C.bool(true)
}

func imnodesIsPinHovered(attribute_id *int32) bool {
	attribute_idArg, attribute_idFin := WrapNumberPtr[C.int, int32](attribute_id)

	defer func() {
		attribute_idFin()
	}()
	return C.imnodes_IsPinHovered(attribute_idArg) == C.bool(true)
}

func imnodesLink(id int32, start_attribute_id int32, end_attribute_id int32) {
	C.imnodes_Link(C.int(id), C.int(start_attribute_id), C.int(end_attribute_id))
}

func imnodesLoadCurrentEditorStateFromIniFile(file_name string) {
	file_nameArg, file_nameFin := WrapString(file_name)
	C.imnodes_LoadCurrentEditorStateFromIniFile(file_nameArg)

	file_nameFin()
}

func imnodesLoadCurrentEditorStateFromIniString(data string, data_size uint64) {
	dataArg, dataFin := WrapString(data)
	C.imnodes_LoadCurrentEditorStateFromIniString(dataArg, C.xulong(data_size))

	dataFin()
}

func imnodesNumSelectedLinks() int32 {
	return int32(C.imnodes_NumSelectedLinks())
}

func imnodesNumSelectedNodes() int32 {
	return int32(C.imnodes_NumSelectedNodes())
}

func imnodesPopAttributeFlag() {
	C.imnodes_PopAttributeFlag()
}

func imnodesPopColorStyle() {
	C.imnodes_PopColorStyle()
}

// imnodesPopStyleVarV parameter default value hint:
// count: 1
func imnodesPopStyleVarV(count int32) {
	C.imnodes_PopStyleVar(C.int(count))
}

func imnodesPushAttributeFlag(flag NodesAttributeFlags) {
	C.imnodes_PushAttributeFlag(C.ImNodesAttributeFlags(flag))
}

func imnodesPushColorStyle(item NodesCol, color uint32) {
	C.imnodes_PushColorStyle(C.ImNodesCol(item), C.uint(color))
}

func imnodesPushStyleVarFloat(style_item NodesStyleVar, value float32) {
	C.imnodes_PushStyleVar_Float(C.ImNodesStyleVar(style_item), C.float(value))
}

func imnodesPushStyleVarVec2(style_item NodesStyleVar, value Vec2) {
	C.imnodes_PushStyleVar_Vec2(C.ImNodesStyleVar(style_item), value.toC())
}

func imnodesSaveCurrentEditorStateToIniFile(file_name string) {
	file_nameArg, file_nameFin := WrapString(file_name)
	C.imnodes_SaveCurrentEditorStateToIniFile(file_nameArg)

	file_nameFin()
}

// imnodesSaveCurrentEditorStateToIniStringV parameter default value hint:
// data_size: NULL
func imnodesSaveCurrentEditorStateToIniStringV(data_size *uint64) string {
	return C.GoString(C.imnodes_SaveCurrentEditorStateToIniString((*C.xulong)(data_size)))
}

func imnodesSelectLink(link_id int32) {
	C.imnodes_SelectLink(C.int(link_id))
}

func imnodesSelectNode(node_id int32) {
	C.imnodes_SelectNode(C.int(node_id))
}

func imnodesSetImGuiContext(ctx *Context) {
	ctxArg, ctxFin := ctx.handle()
	C.imnodes_SetImGuiContext(ctxArg)

	ctxFin()
}

func imnodesSetNodeDraggable(node_id int32, draggable bool) {
	C.imnodes_SetNodeDraggable(C.int(node_id), C.bool(draggable))
}

func imnodesSetNodeEditorSpacePos(node_id int32, editor_space_pos Vec2) {
	C.imnodes_SetNodeEditorSpacePos(C.int(node_id), editor_space_pos.toC())
}

func imnodesSetNodeGridSpacePos(node_id int32, grid_pos Vec2) {
	C.imnodes_SetNodeGridSpacePos(C.int(node_id), grid_pos.toC())
}

func imnodesSetNodeScreenSpacePos(node_id int32, screen_space_pos Vec2) {
	C.imnodes_SetNodeScreenSpacePos(C.int(node_id), screen_space_pos.toC())
}

func imnodesSnapNodeToGrid(node_id int32) {
	C.imnodes_SnapNodeToGrid(C.int(node_id))
}

// imnodesStyleColorsClassicV parameter default value hint:
// dest: NULL
func imnodesStyleColorsClassicV(dest *NodesStyle) {
	destArg, destFin := dest.handle()
	C.imnodes_StyleColorsClassic(destArg)

	destFin()
}

// imnodesStyleColorsDarkV parameter default value hint:
// dest: NULL
func imnodesStyleColorsDarkV(dest *NodesStyle) {
	destArg, destFin := dest.handle()
	C.imnodes_StyleColorsDark(destArg)

	destFin()
}

// imnodesStyleColorsLightV parameter default value hint:
// dest: NULL
func imnodesStyleColorsLightV(dest *NodesStyle) {
	destArg, destFin := dest.handle()
	C.imnodes_StyleColorsLight(destArg)

	destFin()
}

func imnodesBeginInputAttribute(id int32) {
	C.wrap_imnodes_BeginInputAttribute(C.int(id))
}

func imnodesBeginOutputAttribute(id int32) {
	C.wrap_imnodes_BeginOutputAttribute(C.int(id))
}

func imnodesDestroyContext() {
	C.wrap_imnodes_DestroyContext()
}

func imnodesIsAnyAttributeActive() bool {
	return C.wrap_imnodes_IsAnyAttributeActive() == C.bool(true)
}

func imnodesIsLinkCreatedBoolPtr(started_at_attribute_id *int32, ended_at_attribute_id *int32) bool {
	started_at_attribute_idArg, started_at_attribute_idFin := WrapNumberPtr[C.int, int32](started_at_attribute_id)
	ended_at_attribute_idArg, ended_at_attribute_idFin := WrapNumberPtr[C.int, int32](ended_at_attribute_id)

	defer func() {
		started_at_attribute_idFin()
		ended_at_attribute_idFin()
	}()
	return C.wrap_imnodes_IsLinkCreated_BoolPtr(started_at_attribute_idArg, ended_at_attribute_idArg) == C.bool(true)
}

func imnodesIsLinkCreatedIntPtr(started_at_node_id *int32, started_at_attribute_id *int32, ended_at_node_id *int32, ended_at_attribute_id *int32) bool {
	started_at_node_idArg, started_at_node_idFin := WrapNumberPtr[C.int, int32](started_at_node_id)
	started_at_attribute_idArg, started_at_attribute_idFin := WrapNumberPtr[C.int, int32](started_at_attribute_id)
	ended_at_node_idArg, ended_at_node_idFin := WrapNumberPtr[C.int, int32](ended_at_node_id)
	ended_at_attribute_idArg, ended_at_attribute_idFin := WrapNumberPtr[C.int, int32](ended_at_attribute_id)

	defer func() {
		started_at_node_idFin()
		started_at_attribute_idFin()
		ended_at_node_idFin()
		ended_at_attribute_idFin()
	}()
	return C.wrap_imnodes_IsLinkCreated_IntPtr(started_at_node_idArg, started_at_attribute_idArg, ended_at_node_idArg, ended_at_attribute_idArg) == C.bool(true)
}

func imnodesIsLinkDropped() bool {
	return C.wrap_imnodes_IsLinkDropped() == C.bool(true)
}

func imnodesMiniMap() {
	C.wrap_imnodes_MiniMap()
}

func imnodesPopStyleVar() {
	C.wrap_imnodes_PopStyleVar()
}

func imnodesSaveCurrentEditorStateToIniString() string {
	return C.GoString(C.wrap_imnodes_SaveCurrentEditorStateToIniString())
}

func imnodesStyleColorsClassic() {
	C.wrap_imnodes_StyleColorsClassic()
}

func imnodesStyleColorsDark() {
	C.wrap_imnodes_StyleColorsDark()
}

func imnodesStyleColorsLight() {
	C.wrap_imnodes_StyleColorsLight()
}

func (self NodesIO) SetEmulateThreeButtonMouse(v EmulateThreeButtonMouse) {
	vArg, _ := v.c()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetEmulateThreeButtonMouse(selfArg, vArg)
}

func (self *NodesIO) EmulateThreeButtonMouse() EmulateThreeButtonMouse {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_ImNodesIO_GetEmulateThreeButtonMouse(selfArg)
	return *newEmulateThreeButtonMouseFromC(func() *C.EmulateThreeButtonMouse { result := result; return &result }())
}

func (self NodesIO) SetLinkDetachWithModifierClick(v LinkDetachWithModifierClick) {
	vArg, _ := v.c()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetLinkDetachWithModifierClick(selfArg, vArg)
}

func (self *NodesIO) LinkDetachWithModifierClick() LinkDetachWithModifierClick {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_ImNodesIO_GetLinkDetachWithModifierClick(selfArg)
	return *newLinkDetachWithModifierClickFromC(func() *C.LinkDetachWithModifierClick { result := result; return &result }())
}

func (self NodesIO) SetMultipleSelectModifier(v MultipleSelectModifier) {
	vArg, _ := v.c()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetMultipleSelectModifier(selfArg, vArg)
}

func (self *NodesIO) MultipleSelectModifier() MultipleSelectModifier {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_ImNodesIO_GetMultipleSelectModifier(selfArg)
	return *newMultipleSelectModifierFromC(func() *C.MultipleSelectModifier { result := result; return &result }())
}

func (self NodesIO) SetAltMouseButton(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetAltMouseButton(selfArg, C.int(v))
}

func (self *NodesIO) AltMouseButton() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_ImNodesIO_GetAltMouseButton(selfArg))
}

func (self NodesIO) SetAutoPanningSpeed(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesIO_SetAutoPanningSpeed(selfArg, C.float(v))
}

func (self *NodesIO) AutoPanningSpeed() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesIO_GetAutoPanningSpeed(selfArg))
}

func (self NodesStyle) SetGridSpacing(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetGridSpacing(selfArg, C.float(v))
}

func (self *NodesStyle) GridSpacing() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetGridSpacing(selfArg))
}

func (self NodesStyle) SetNodeCornerRounding(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetNodeCornerRounding(selfArg, C.float(v))
}

func (self *NodesStyle) NodeCornerRounding() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetNodeCornerRounding(selfArg))
}

func (self NodesStyle) SetNodePadding(v Vec2) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetNodePadding(selfArg, v.toC())
}

func (self *NodesStyle) NodePadding() Vec2 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *(&Vec2{}).fromC(C.wrap_ImNodesStyle_GetNodePadding(selfArg))
}

func (self NodesStyle) SetNodeBorderThickness(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetNodeBorderThickness(selfArg, C.float(v))
}

func (self *NodesStyle) NodeBorderThickness() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetNodeBorderThickness(selfArg))
}

func (self NodesStyle) SetLinkThickness(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetLinkThickness(selfArg, C.float(v))
}

func (self *NodesStyle) LinkThickness() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetLinkThickness(selfArg))
}

func (self NodesStyle) SetLinkLineSegmentsPerLength(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetLinkLineSegmentsPerLength(selfArg, C.float(v))
}

func (self *NodesStyle) LinkLineSegmentsPerLength() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetLinkLineSegmentsPerLength(selfArg))
}

func (self NodesStyle) SetLinkHoverDistance(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetLinkHoverDistance(selfArg, C.float(v))
}

func (self *NodesStyle) LinkHoverDistance() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetLinkHoverDistance(selfArg))
}

func (self NodesStyle) SetPinCircleRadius(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinCircleRadius(selfArg, C.float(v))
}

func (self *NodesStyle) PinCircleRadius() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinCircleRadius(selfArg))
}

func (self NodesStyle) SetPinQuadSideLength(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinQuadSideLength(selfArg, C.float(v))
}

func (self *NodesStyle) PinQuadSideLength() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinQuadSideLength(selfArg))
}

func (self NodesStyle) SetPinTriangleSideLength(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinTriangleSideLength(selfArg, C.float(v))
}

func (self *NodesStyle) PinTriangleSideLength() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinTriangleSideLength(selfArg))
}

func (self NodesStyle) SetPinLineThickness(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinLineThickness(selfArg, C.float(v))
}

func (self *NodesStyle) PinLineThickness() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinLineThickness(selfArg))
}

func (self NodesStyle) SetPinHoverRadius(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinHoverRadius(selfArg, C.float(v))
}

func (self *NodesStyle) PinHoverRadius() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinHoverRadius(selfArg))
}

func (self NodesStyle) SetPinOffset(v float32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetPinOffset(selfArg, C.float(v))
}

func (self *NodesStyle) PinOffset() float32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return float32(C.wrap_ImNodesStyle_GetPinOffset(selfArg))
}

func (self NodesStyle) SetMiniMapPadding(v Vec2) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetMiniMapPadding(selfArg, v.toC())
}

func (self *NodesStyle) MiniMapPadding() Vec2 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *(&Vec2{}).fromC(C.wrap_ImNodesStyle_GetMiniMapPadding(selfArg))
}

func (self NodesStyle) SetMiniMapOffset(v Vec2) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetMiniMapOffset(selfArg, v.toC())
}

func (self *NodesStyle) MiniMapOffset() Vec2 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *(&Vec2{}).fromC(C.wrap_ImNodesStyle_GetMiniMapOffset(selfArg))
}

func (self NodesStyle) SetFlags(v NodesStyleFlags) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_ImNodesStyle_SetFlags(selfArg, C.ImNodesStyleFlags(v))
}

func (self *NodesStyle) Flags() NodesStyleFlags {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return NodesStyleFlags(C.wrap_ImNodesStyle_GetFlags(selfArg))
}
