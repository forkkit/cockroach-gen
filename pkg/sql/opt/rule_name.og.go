// Code generated by optgen; DO NOT EDIT.

package opt

const (
	startAutoRule RuleName = iota + NumManualRuleNames

	// ------------------------------------------------------------
	// Normalize Rule Names
	// ------------------------------------------------------------
	EliminateAggDistinct
	EliminateEmptyFilters
	EliminateSingletonAndOr
	SimplifyAnd
	SimplifyOr
	SimplifyFilters
	FoldNullAndOr
	FoldNotTrue
	FoldNotFalse
	NegateComparison
	EliminateNot
	NegateAnd
	NegateOr
	ExtractRedundantClause
	ExtractRedundantSubclause
	CommuteVarInequality
	CommuteConstInequality
	NormalizeCmpPlusConst
	NormalizeCmpMinusConst
	NormalizeCmpConstMinus
	NormalizeTupleEquality
	FoldNullComparisonLeft
	FoldNullComparisonRight
	FoldIsNull
	FoldNonNullIsNull
	FoldIsNotNull
	FoldNonNullIsNotNull
	CommuteNullIs
	DecorrelateJoin
	TryDecorrelateSelect
	TryDecorrelateProject
	TryDecorrelateProjectSelect
	TryDecorrelateProjectInnerJoin
	TryDecorrelateInnerJoin
	TryDecorrelateInnerLeftJoin
	TryDecorrelateGroupBy
	TryDecorrelateScalarGroupBy
	TryDecorrelateSemiJoin
	TryDecorrelateLimitOne
	TryDecorrelateZip
	HoistSelectExists
	HoistSelectNotExists
	HoistSelectSubquery
	HoistProjectSubquery
	HoistJoinSubquery
	HoistValuesSubquery
	HoistZipSubquery
	NormalizeAnyFilter
	NormalizeNotAnyFilter
	FoldArray
	FoldBinary
	FoldUnary
	FoldComparison
	ConvertGroupByToDistinct
	EliminateDistinct
	EliminateGroupByProject
	ReduceGroupingCols
	EliminateAggDistinctForKeys
	PushSelectIntoInlinableProject
	InlineProjectInProject
	PushFilterIntoJoinLeftAndRight
	MapFilterIntoJoinLeft
	MapFilterIntoJoinRight
	PushFilterIntoJoinLeft
	PushFilterIntoJoinRight
	SimplifyLeftJoinWithoutFilters
	SimplifyRightJoinWithoutFilters
	SimplifyLeftJoinWithFilters
	SimplifyRightJoinWithFilters
	EliminateSemiJoin
	EliminateAntiJoin
	EliminateJoinNoColsLeft
	EliminateJoinNoColsRight
	HoistJoinProject
	SimplifyJoinNotNullEquality
	EliminateLimit
	PushLimitIntoProject
	PushOffsetIntoProject
	EliminateMax1Row
	FoldPlusZero
	FoldZeroPlus
	FoldMinusZero
	FoldMultOne
	FoldOneMult
	FoldDivOne
	InvertMinus
	EliminateUnaryMinus
	SimplifyLimitOrdering
	SimplifyOffsetOrdering
	SimplifyGroupByOrdering
	SimplifyRowNumberOrdering
	SimplifyExplainOrdering
	EliminateProject
	EliminateProjectProject
	PruneProjectCols
	PruneScanCols
	PruneSelectCols
	PruneLimitCols
	PruneOffsetCols
	PruneJoinLeftCols
	PruneJoinRightCols
	PruneAggCols
	PruneGroupByCols
	PruneValuesCols
	PruneRowNumberCols
	PruneExplainCols
	RejectNullsLeftJoin
	RejectNullsRightJoin
	RejectNullsGroupBy
	CommuteVar
	CommuteConst
	EliminateCoalesce
	SimplifyCoalesce
	EliminateCast
	FoldNullCast
	FoldNullUnary
	FoldNullBinaryLeft
	FoldNullBinaryRight
	FoldNullInNonEmpty
	FoldNullInEmpty
	FoldNullNotInEmpty
	NormalizeInConst
	FoldInNull
	EliminateExistsProject
	EliminateExistsGroupBy
	NormalizeJSONFieldAccess
	NormalizeJSONContains
	SimplifyCaseWhenConstValue
	EliminateSelect
	MergeSelects
	PushSelectIntoProject
	MergeSelectInnerJoin
	PushSelectCondLeftIntoJoinLeftAndRight
	PushSelectCondRightIntoJoinLeftAndRight
	PushSelectIntoJoinLeft
	PushSelectIntoJoinRight
	PushSelectIntoGroupBy
	RemoveNotNullCondition
	EliminateUnionAllLeft
	EliminateUnionAllRight

	// startExploreRule tracks the number of normalization rules;
	// all rules greater than this value are exploration rules.
	startExploreRule

	// ------------------------------------------------------------
	// Explore Rule Names
	// ------------------------------------------------------------
	ReplaceMinWithLimit
	ReplaceMaxWithLimit
	CommuteJoin
	CommuteLeftJoin
	CommuteRightJoin
	GenerateMergeJoins
	GenerateLookupJoins
	GenerateLookupJoinsWithFilter
	GenerateLimitedScans
	PushLimitIntoConstrainedScan
	PushLimitIntoIndexJoin
	GenerateIndexScans
	GenerateConstrainedScans
	GenerateInvertedIndexScans

	// NumRuleNames tracks the total count of rule names.
	NumRuleNames
)
