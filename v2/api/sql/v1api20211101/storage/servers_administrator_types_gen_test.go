// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_ServersAdministrator_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersAdministrator via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersAdministrator, ServersAdministratorGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersAdministrator runs a test to see if a specific instance of ServersAdministrator round trips to JSON and back losslessly
func RunJSONSerializationTestForServersAdministrator(subject ServersAdministrator) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersAdministrator
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

// Generator of ServersAdministrator instances for property testing - lazily instantiated by
// ServersAdministratorGenerator()
var serversAdministratorGenerator gopter.Gen

// ServersAdministratorGenerator returns a generator of ServersAdministrator instances for property testing.
func ServersAdministratorGenerator() gopter.Gen {
	if serversAdministratorGenerator != nil {
		return serversAdministratorGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersAdministrator(generators)
	serversAdministratorGenerator = gen.Struct(reflect.TypeOf(ServersAdministrator{}), generators)

	return serversAdministratorGenerator
}

// AddRelatedPropertyGeneratorsForServersAdministrator is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersAdministrator(gens map[string]gopter.Gen) {
	gens["Spec"] = ServersAdministrator_SpecGenerator()
	gens["Status"] = ServersAdministrator_STATUSGenerator()
}

func Test_ServersAdministrator_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersAdministrator_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersAdministrator_STATUS, ServersAdministrator_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersAdministrator_STATUS runs a test to see if a specific instance of ServersAdministrator_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServersAdministrator_STATUS(subject ServersAdministrator_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersAdministrator_STATUS
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

// Generator of ServersAdministrator_STATUS instances for property testing - lazily instantiated by
// ServersAdministrator_STATUSGenerator()
var serversAdministrator_STATUSGenerator gopter.Gen

// ServersAdministrator_STATUSGenerator returns a generator of ServersAdministrator_STATUS instances for property testing.
func ServersAdministrator_STATUSGenerator() gopter.Gen {
	if serversAdministrator_STATUSGenerator != nil {
		return serversAdministrator_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersAdministrator_STATUS(generators)
	serversAdministrator_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersAdministrator_STATUS{}), generators)

	return serversAdministrator_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServersAdministrator_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersAdministrator_STATUS(gens map[string]gopter.Gen) {
	gens["AdministratorType"] = gen.PtrOf(gen.AlphaString())
	gens["AzureADOnlyAuthentication"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Login"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Sid"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersAdministrator_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersAdministrator_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersAdministrator_Spec, ServersAdministrator_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersAdministrator_Spec runs a test to see if a specific instance of ServersAdministrator_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersAdministrator_Spec(subject ServersAdministrator_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersAdministrator_Spec
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

// Generator of ServersAdministrator_Spec instances for property testing - lazily instantiated by
// ServersAdministrator_SpecGenerator()
var serversAdministrator_SpecGenerator gopter.Gen

// ServersAdministrator_SpecGenerator returns a generator of ServersAdministrator_Spec instances for property testing.
func ServersAdministrator_SpecGenerator() gopter.Gen {
	if serversAdministrator_SpecGenerator != nil {
		return serversAdministrator_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersAdministrator_Spec(generators)
	serversAdministrator_SpecGenerator = gen.Struct(reflect.TypeOf(ServersAdministrator_Spec{}), generators)

	return serversAdministrator_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersAdministrator_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersAdministrator_Spec(gens map[string]gopter.Gen) {
	gens["AdministratorType"] = gen.PtrOf(gen.AlphaString())
	gens["Login"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Sid"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}
