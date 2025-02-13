// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_CompositePath_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CompositePath_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCompositePath_STATUS_ARM, CompositePath_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCompositePath_STATUS_ARM runs a test to see if a specific instance of CompositePath_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCompositePath_STATUS_ARM(subject CompositePath_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CompositePath_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of CompositePath_STATUS_ARM instances for property testing - lazily instantiated by
// CompositePath_STATUS_ARMGenerator()
var compositePath_STATUS_ARMGenerator gopter.Gen

// CompositePath_STATUS_ARMGenerator returns a generator of CompositePath_STATUS_ARM instances for property testing.
func CompositePath_STATUS_ARMGenerator() gopter.Gen {
	if compositePath_STATUS_ARMGenerator != nil {
		return compositePath_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCompositePath_STATUS_ARM(generators)
	compositePath_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(CompositePath_STATUS_ARM{}), generators)

	return compositePath_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCompositePath_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCompositePath_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Order"] = gen.PtrOf(gen.OneConstOf(CompositePath_Order_STATUS_ARM_Ascending, CompositePath_Order_STATUS_ARM_Descending))
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

func Test_ConflictResolutionPolicy_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConflictResolutionPolicy_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConflictResolutionPolicy_STATUS_ARM, ConflictResolutionPolicy_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConflictResolutionPolicy_STATUS_ARM runs a test to see if a specific instance of ConflictResolutionPolicy_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConflictResolutionPolicy_STATUS_ARM(subject ConflictResolutionPolicy_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConflictResolutionPolicy_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ConflictResolutionPolicy_STATUS_ARM instances for property testing - lazily instantiated by
// ConflictResolutionPolicy_STATUS_ARMGenerator()
var conflictResolutionPolicy_STATUS_ARMGenerator gopter.Gen

// ConflictResolutionPolicy_STATUS_ARMGenerator returns a generator of ConflictResolutionPolicy_STATUS_ARM instances for property testing.
func ConflictResolutionPolicy_STATUS_ARMGenerator() gopter.Gen {
	if conflictResolutionPolicy_STATUS_ARMGenerator != nil {
		return conflictResolutionPolicy_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConflictResolutionPolicy_STATUS_ARM(generators)
	conflictResolutionPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ConflictResolutionPolicy_STATUS_ARM{}), generators)

	return conflictResolutionPolicy_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForConflictResolutionPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConflictResolutionPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ConflictResolutionPath"] = gen.PtrOf(gen.AlphaString())
	gens["ConflictResolutionProcedure"] = gen.PtrOf(gen.AlphaString())
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(ConflictResolutionPolicy_Mode_STATUS_ARM_Custom, ConflictResolutionPolicy_Mode_STATUS_ARM_LastWriterWins))
}

func Test_ContainerPartitionKey_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContainerPartitionKey_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContainerPartitionKey_STATUS_ARM, ContainerPartitionKey_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContainerPartitionKey_STATUS_ARM runs a test to see if a specific instance of ContainerPartitionKey_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForContainerPartitionKey_STATUS_ARM(subject ContainerPartitionKey_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContainerPartitionKey_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ContainerPartitionKey_STATUS_ARM instances for property testing - lazily instantiated by
// ContainerPartitionKey_STATUS_ARMGenerator()
var containerPartitionKey_STATUS_ARMGenerator gopter.Gen

// ContainerPartitionKey_STATUS_ARMGenerator returns a generator of ContainerPartitionKey_STATUS_ARM instances for property testing.
func ContainerPartitionKey_STATUS_ARMGenerator() gopter.Gen {
	if containerPartitionKey_STATUS_ARMGenerator != nil {
		return containerPartitionKey_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerPartitionKey_STATUS_ARM(generators)
	containerPartitionKey_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ContainerPartitionKey_STATUS_ARM{}), generators)

	return containerPartitionKey_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForContainerPartitionKey_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContainerPartitionKey_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(ContainerPartitionKey_Kind_STATUS_ARM_Hash, ContainerPartitionKey_Kind_STATUS_ARM_MultiHash, ContainerPartitionKey_Kind_STATUS_ARM_Range))
	gens["Paths"] = gen.SliceOf(gen.AlphaString())
	gens["SystemKey"] = gen.PtrOf(gen.Bool())
	gens["Version"] = gen.PtrOf(gen.Int())
}

func Test_ExcludedPath_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExcludedPath_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExcludedPath_STATUS_ARM, ExcludedPath_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExcludedPath_STATUS_ARM runs a test to see if a specific instance of ExcludedPath_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExcludedPath_STATUS_ARM(subject ExcludedPath_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExcludedPath_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ExcludedPath_STATUS_ARM instances for property testing - lazily instantiated by
// ExcludedPath_STATUS_ARMGenerator()
var excludedPath_STATUS_ARMGenerator gopter.Gen

// ExcludedPath_STATUS_ARMGenerator returns a generator of ExcludedPath_STATUS_ARM instances for property testing.
func ExcludedPath_STATUS_ARMGenerator() gopter.Gen {
	if excludedPath_STATUS_ARMGenerator != nil {
		return excludedPath_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExcludedPath_STATUS_ARM(generators)
	excludedPath_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ExcludedPath_STATUS_ARM{}), generators)

	return excludedPath_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForExcludedPath_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExcludedPath_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

func Test_IncludedPath_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IncludedPath_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIncludedPath_STATUS_ARM, IncludedPath_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIncludedPath_STATUS_ARM runs a test to see if a specific instance of IncludedPath_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIncludedPath_STATUS_ARM(subject IncludedPath_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IncludedPath_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of IncludedPath_STATUS_ARM instances for property testing - lazily instantiated by
// IncludedPath_STATUS_ARMGenerator()
var includedPath_STATUS_ARMGenerator gopter.Gen

// IncludedPath_STATUS_ARMGenerator returns a generator of IncludedPath_STATUS_ARM instances for property testing.
// We first initialize includedPath_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IncludedPath_STATUS_ARMGenerator() gopter.Gen {
	if includedPath_STATUS_ARMGenerator != nil {
		return includedPath_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIncludedPath_STATUS_ARM(generators)
	includedPath_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(IncludedPath_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIncludedPath_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForIncludedPath_STATUS_ARM(generators)
	includedPath_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(IncludedPath_STATUS_ARM{}), generators)

	return includedPath_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIncludedPath_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIncludedPath_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForIncludedPath_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIncludedPath_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(Indexes_STATUS_ARMGenerator())
}

func Test_Indexes_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Indexes_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIndexes_STATUS_ARM, Indexes_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIndexes_STATUS_ARM runs a test to see if a specific instance of Indexes_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIndexes_STATUS_ARM(subject Indexes_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Indexes_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Indexes_STATUS_ARM instances for property testing - lazily instantiated by Indexes_STATUS_ARMGenerator()
var indexes_STATUS_ARMGenerator gopter.Gen

// Indexes_STATUS_ARMGenerator returns a generator of Indexes_STATUS_ARM instances for property testing.
func Indexes_STATUS_ARMGenerator() gopter.Gen {
	if indexes_STATUS_ARMGenerator != nil {
		return indexes_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexes_STATUS_ARM(generators)
	indexes_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Indexes_STATUS_ARM{}), generators)

	return indexes_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIndexes_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIndexes_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DataType"] = gen.PtrOf(gen.OneConstOf(
		Indexes_DataType_STATUS_ARM_LineString,
		Indexes_DataType_STATUS_ARM_MultiPolygon,
		Indexes_DataType_STATUS_ARM_Number,
		Indexes_DataType_STATUS_ARM_Point,
		Indexes_DataType_STATUS_ARM_Polygon,
		Indexes_DataType_STATUS_ARM_String))
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(Indexes_Kind_STATUS_ARM_Hash, Indexes_Kind_STATUS_ARM_Range, Indexes_Kind_STATUS_ARM_Spatial))
	gens["Precision"] = gen.PtrOf(gen.Int())
}

func Test_IndexingPolicy_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IndexingPolicy_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIndexingPolicy_STATUS_ARM, IndexingPolicy_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIndexingPolicy_STATUS_ARM runs a test to see if a specific instance of IndexingPolicy_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIndexingPolicy_STATUS_ARM(subject IndexingPolicy_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IndexingPolicy_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of IndexingPolicy_STATUS_ARM instances for property testing - lazily instantiated by
// IndexingPolicy_STATUS_ARMGenerator()
var indexingPolicy_STATUS_ARMGenerator gopter.Gen

// IndexingPolicy_STATUS_ARMGenerator returns a generator of IndexingPolicy_STATUS_ARM instances for property testing.
// We first initialize indexingPolicy_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IndexingPolicy_STATUS_ARMGenerator() gopter.Gen {
	if indexingPolicy_STATUS_ARMGenerator != nil {
		return indexingPolicy_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexingPolicy_STATUS_ARM(generators)
	indexingPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(IndexingPolicy_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexingPolicy_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForIndexingPolicy_STATUS_ARM(generators)
	indexingPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(IndexingPolicy_STATUS_ARM{}), generators)

	return indexingPolicy_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIndexingPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIndexingPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Automatic"] = gen.PtrOf(gen.Bool())
	gens["IndexingMode"] = gen.PtrOf(gen.OneConstOf(IndexingPolicy_IndexingMode_STATUS_ARM_Consistent, IndexingPolicy_IndexingMode_STATUS_ARM_Lazy, IndexingPolicy_IndexingMode_STATUS_ARM_None))
}

// AddRelatedPropertyGeneratorsForIndexingPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIndexingPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CompositeIndexes"] = gen.SliceOf(gen.SliceOf(CompositePath_STATUS_ARMGenerator()))
	gens["ExcludedPaths"] = gen.SliceOf(ExcludedPath_STATUS_ARMGenerator())
	gens["IncludedPaths"] = gen.SliceOf(IncludedPath_STATUS_ARMGenerator())
	gens["SpatialIndexes"] = gen.SliceOf(SpatialSpec_STATUS_ARMGenerator())
}

func Test_SpatialSpec_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SpatialSpec_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSpatialSpec_STATUS_ARM, SpatialSpec_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSpatialSpec_STATUS_ARM runs a test to see if a specific instance of SpatialSpec_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSpatialSpec_STATUS_ARM(subject SpatialSpec_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SpatialSpec_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SpatialSpec_STATUS_ARM instances for property testing - lazily instantiated by
// SpatialSpec_STATUS_ARMGenerator()
var spatialSpec_STATUS_ARMGenerator gopter.Gen

// SpatialSpec_STATUS_ARMGenerator returns a generator of SpatialSpec_STATUS_ARM instances for property testing.
func SpatialSpec_STATUS_ARMGenerator() gopter.Gen {
	if spatialSpec_STATUS_ARMGenerator != nil {
		return spatialSpec_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSpatialSpec_STATUS_ARM(generators)
	spatialSpec_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SpatialSpec_STATUS_ARM{}), generators)

	return spatialSpec_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSpatialSpec_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSpatialSpec_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
	gens["Types"] = gen.SliceOf(gen.OneConstOf(
		SpatialType_STATUS_ARM_LineString,
		SpatialType_STATUS_ARM_MultiPolygon,
		SpatialType_STATUS_ARM_Point,
		SpatialType_STATUS_ARM_Polygon))
}

func Test_SqlContainerGetProperties_Resource_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerGetProperties_Resource_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerGetProperties_Resource_STATUS_ARM, SqlContainerGetProperties_Resource_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerGetProperties_Resource_STATUS_ARM runs a test to see if a specific instance of SqlContainerGetProperties_Resource_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerGetProperties_Resource_STATUS_ARM(subject SqlContainerGetProperties_Resource_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerGetProperties_Resource_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SqlContainerGetProperties_Resource_STATUS_ARM instances for property testing - lazily instantiated by
// SqlContainerGetProperties_Resource_STATUS_ARMGenerator()
var sqlContainerGetProperties_Resource_STATUS_ARMGenerator gopter.Gen

// SqlContainerGetProperties_Resource_STATUS_ARMGenerator returns a generator of SqlContainerGetProperties_Resource_STATUS_ARM instances for property testing.
// We first initialize sqlContainerGetProperties_Resource_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlContainerGetProperties_Resource_STATUS_ARMGenerator() gopter.Gen {
	if sqlContainerGetProperties_Resource_STATUS_ARMGenerator != nil {
		return sqlContainerGetProperties_Resource_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerGetProperties_Resource_STATUS_ARM(generators)
	sqlContainerGetProperties_Resource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerGetProperties_Resource_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerGetProperties_Resource_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForSqlContainerGetProperties_Resource_STATUS_ARM(generators)
	sqlContainerGetProperties_Resource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerGetProperties_Resource_STATUS_ARM{}), generators)

	return sqlContainerGetProperties_Resource_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlContainerGetProperties_Resource_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlContainerGetProperties_Resource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["DefaultTtl"] = gen.PtrOf(gen.Int())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

// AddRelatedPropertyGeneratorsForSqlContainerGetProperties_Resource_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerGetProperties_Resource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ConflictResolutionPolicy"] = gen.PtrOf(ConflictResolutionPolicy_STATUS_ARMGenerator())
	gens["IndexingPolicy"] = gen.PtrOf(IndexingPolicy_STATUS_ARMGenerator())
	gens["PartitionKey"] = gen.PtrOf(ContainerPartitionKey_STATUS_ARMGenerator())
	gens["UniqueKeyPolicy"] = gen.PtrOf(UniqueKeyPolicy_STATUS_ARMGenerator())
}

func Test_SqlContainerGetProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerGetProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerGetProperties_STATUS_ARM, SqlContainerGetProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerGetProperties_STATUS_ARM runs a test to see if a specific instance of SqlContainerGetProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerGetProperties_STATUS_ARM(subject SqlContainerGetProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerGetProperties_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SqlContainerGetProperties_STATUS_ARM instances for property testing - lazily instantiated by
// SqlContainerGetProperties_STATUS_ARMGenerator()
var sqlContainerGetProperties_STATUS_ARMGenerator gopter.Gen

// SqlContainerGetProperties_STATUS_ARMGenerator returns a generator of SqlContainerGetProperties_STATUS_ARM instances for property testing.
func SqlContainerGetProperties_STATUS_ARMGenerator() gopter.Gen {
	if sqlContainerGetProperties_STATUS_ARMGenerator != nil {
		return sqlContainerGetProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlContainerGetProperties_STATUS_ARM(generators)
	sqlContainerGetProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerGetProperties_STATUS_ARM{}), generators)

	return sqlContainerGetProperties_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlContainerGetProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerGetProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResource_STATUS_ARMGenerator())
	gens["Resource"] = gen.PtrOf(SqlContainerGetProperties_Resource_STATUS_ARMGenerator())
}

func Test_SqlDatabaseContainer_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainer_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainer_STATUS_ARM, SqlDatabaseContainer_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainer_STATUS_ARM runs a test to see if a specific instance of SqlDatabaseContainer_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainer_STATUS_ARM(subject SqlDatabaseContainer_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainer_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SqlDatabaseContainer_STATUS_ARM instances for property testing - lazily instantiated by
// SqlDatabaseContainer_STATUS_ARMGenerator()
var sqlDatabaseContainer_STATUS_ARMGenerator gopter.Gen

// SqlDatabaseContainer_STATUS_ARMGenerator returns a generator of SqlDatabaseContainer_STATUS_ARM instances for property testing.
// We first initialize sqlDatabaseContainer_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseContainer_STATUS_ARMGenerator() gopter.Gen {
	if sqlDatabaseContainer_STATUS_ARMGenerator != nil {
		return sqlDatabaseContainer_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainer_STATUS_ARM(generators)
	sqlDatabaseContainer_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainer_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainer_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainer_STATUS_ARM(generators)
	sqlDatabaseContainer_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainer_STATUS_ARM{}), generators)

	return sqlDatabaseContainer_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseContainer_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseContainer_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainer_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainer_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlContainerGetProperties_STATUS_ARMGenerator())
}

func Test_UniqueKeyPolicy_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UniqueKeyPolicy_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUniqueKeyPolicy_STATUS_ARM, UniqueKeyPolicy_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUniqueKeyPolicy_STATUS_ARM runs a test to see if a specific instance of UniqueKeyPolicy_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUniqueKeyPolicy_STATUS_ARM(subject UniqueKeyPolicy_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UniqueKeyPolicy_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UniqueKeyPolicy_STATUS_ARM instances for property testing - lazily instantiated by
// UniqueKeyPolicy_STATUS_ARMGenerator()
var uniqueKeyPolicy_STATUS_ARMGenerator gopter.Gen

// UniqueKeyPolicy_STATUS_ARMGenerator returns a generator of UniqueKeyPolicy_STATUS_ARM instances for property testing.
func UniqueKeyPolicy_STATUS_ARMGenerator() gopter.Gen {
	if uniqueKeyPolicy_STATUS_ARMGenerator != nil {
		return uniqueKeyPolicy_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForUniqueKeyPolicy_STATUS_ARM(generators)
	uniqueKeyPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UniqueKeyPolicy_STATUS_ARM{}), generators)

	return uniqueKeyPolicy_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForUniqueKeyPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUniqueKeyPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["UniqueKeys"] = gen.SliceOf(UniqueKey_STATUS_ARMGenerator())
}

func Test_UniqueKey_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UniqueKey_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUniqueKey_STATUS_ARM, UniqueKey_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUniqueKey_STATUS_ARM runs a test to see if a specific instance of UniqueKey_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUniqueKey_STATUS_ARM(subject UniqueKey_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UniqueKey_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UniqueKey_STATUS_ARM instances for property testing - lazily instantiated by
// UniqueKey_STATUS_ARMGenerator()
var uniqueKey_STATUS_ARMGenerator gopter.Gen

// UniqueKey_STATUS_ARMGenerator returns a generator of UniqueKey_STATUS_ARM instances for property testing.
func UniqueKey_STATUS_ARMGenerator() gopter.Gen {
	if uniqueKey_STATUS_ARMGenerator != nil {
		return uniqueKey_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUniqueKey_STATUS_ARM(generators)
	uniqueKey_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UniqueKey_STATUS_ARM{}), generators)

	return uniqueKey_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUniqueKey_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUniqueKey_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Paths"] = gen.SliceOf(gen.AlphaString())
}
