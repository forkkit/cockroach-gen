// Code generated by "stringer -output=pkg/sql/opt/rule_name_string.go -type=RuleName pkg/sql/opt/rule_name.go pkg/sql/opt/rule_name.og.go"; DO NOT EDIT.

package opt

import "strconv"

const _RuleName_name = "InvalidRuleNameSimplifyProjectOrderingSimplifyRootOrderingPruneRootColsNumManualRuleNamesEliminateEmptyAndEliminateEmptyOrEliminateSingletonAndOrSimplifyAndSimplifyOrSimplifyFiltersFoldNullAndOrNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantClauseExtractRedundantSubclauseCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldIsNotNullFoldNonNullIsNotNullCommuteNullIsDecorrelateJoinTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateZipHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistZipSubqueryNormalizeAnyFilterNormalizeNotAnyFilterConvertGroupByToDistinctEliminateDistinctEliminateGroupByProjectReduceGroupingColsPushSelectIntoInlinableProjectInlineProjectInProjectEnsureJoinFiltersAndEnsureJoinFiltersPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinWithoutFiltersSimplifyRightJoinWithoutFiltersSimplifyLeftJoinWithFiltersSimplifyRightJoinWithFiltersEliminateSemiJoinEliminateAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightEliminateLimitPushLimitIntoProjectPushOffsetIntoProjectEliminateMax1RowFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusFoldUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyRowNumberOrderingSimplifyExplainOrderingEliminateProjectEliminateProjectProjectPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneRowNumberColsPruneExplainColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldNullInEmptyFoldNullNotInEmptyNormalizeInConstFoldInNullEliminateExistsProjectEliminateExistsGroupByNormalizeJSONFieldAccessNormalizeJSONContainsEliminateSelectEnsureSelectFiltersAndEnsureSelectFiltersMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectCondRightIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoJoinRightPushSelectIntoGroupByRemoveNotNullConditionstartExploreRuleReplaceMinWithLimitReplaceMaxWithLimitCommuteJoinCommuteLeftJoinCommuteRightJoinGenerateMergeJoinsGenerateLookupJoinGenerateLookupJoinWithFilterPushJoinThroughIndexJoinPushJoinThroughIndexJoinWithExtraFilterPushLimitIntoScanPushLimitIntoIndexJoinGenerateIndexScansConstrainScanPushFilterIntoIndexJoinNoRemainderPushFilterIntoIndexJoinConstrainIndexJoinScanGenerateInvertedIndexScansNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 38, 58, 71, 89, 106, 122, 145, 156, 166, 181, 194, 210, 222, 231, 239, 261, 286, 306, 328, 349, 371, 393, 415, 437, 460, 470, 487, 500, 520, 533, 548, 568, 589, 616, 646, 669, 696, 717, 744, 766, 783, 800, 820, 839, 859, 876, 895, 911, 929, 950, 974, 991, 1014, 1032, 1062, 1084, 1104, 1121, 1151, 1172, 1194, 1216, 1239, 1269, 1300, 1327, 1355, 1372, 1389, 1412, 1436, 1450, 1470, 1491, 1507, 1519, 1531, 1544, 1555, 1566, 1576, 1587, 1606, 1620, 1641, 1663, 1686, 1711, 1734, 1750, 1773, 1789, 1802, 1817, 1831, 1846, 1863, 1881, 1893, 1909, 1924, 1942, 1958, 1977, 1997, 2015, 2025, 2037, 2054, 2070, 2083, 2095, 2108, 2126, 2145, 2163, 2178, 2196, 2212, 2222, 2244, 2266, 2290, 2311, 2326, 2348, 2367, 2379, 2400, 2420, 2458, 2497, 2519, 2542, 2563, 2585, 2601, 2620, 2639, 2650, 2665, 2681, 2699, 2717, 2745, 2769, 2808, 2825, 2847, 2865, 2878, 2912, 2935, 2957, 2983, 2995}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
