// Copyright 2017, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package trace contains types for representing trace information, and
functions for global configuration of tracing.

The following assumes a basic familiarity with OpenCensus concepts.
See http://opencensus.io.


Enabling Tracing for a Program

To use OpenCensus tracing, register at least one Exporter. You can use
one of the provided exporters or write your own.

    trace.RegisterExporter(anExporter)

By default, traces will be sampled relatively rarely. To change the sampling
frequency for your entire program, call SetDefaultSampler. Use a ProbabilitySampler
to sample a subset of traces, or use AlwaysSample to collect a trace on every run:

    trace.SetDefaultSampler(trace.AlwaysSample())


Adding Spans to a Trace

A trace consists of a tree of spans. In Go, the current span is carried in a
context.Context.

It is common to want to capture all the activity of a function call in a span. For
this to work, the function must take a context.Context as a parameter. Add these two
lines to the top of the function:

    ctx = trace.StartSpan(ctx, "your choice of name")
    defer trace.EndSpan(ctx)

StartSpan will create a child span if one already exists, and will create a
new top-level span otherwise.

As a suggestion, use the fully-qualified function name as the span name, e.g.
"github.com/me/mypackage.Run".
*/
package trace // import "go.opencensus.io/trace"
