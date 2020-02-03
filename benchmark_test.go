// Lute - A structured markdown engine.
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under the Mulan PSL v1.
// You can use this software according to the terms and conditions of the Mulan PSL v1.
// You may obtain a copy of Mulan PSL v1 at:
//     http://license.coscl.org.cn/MulanPSL
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v1 for more details.

package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/88250/lute"
	gm "github.com/gomarkdown/markdown"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"gitlab.com/golang-commonmark/markdown"
	"gopkg.in/russross/blackfriday.v2"
)

const spec = "commonmark-spec"

func BenchmarkLute(b *testing.B) {
	bytes, err := ioutil.ReadFile(spec + ".md")
	if nil != err {
		b.Fatalf("read spec text failed: " + err.Error())
	}

	luteEngine := lute.New()
	luteEngine.GFMTaskListItem = true
	luteEngine.GFMTable = true
	luteEngine.GFMAutoLink = true
	luteEngine.GFMStrikethrough = true
	luteEngine.SoftBreak2HardBreak = false
	luteEngine.CodeSyntaxHighlight = false
	luteEngine.Footnotes = false
	luteEngine.ToC = false
	luteEngine.AutoSpace = false
	luteEngine.FixTermTypo = false
	luteEngine.ChinesePunct = false
	luteEngine.Emoji = false
	html, err := luteEngine.Markdown("spec text", bytes)
	if nil != err {
		b.Fatalf("unexpected: %s", err)
	}
	if err := ioutil.WriteFile(spec+".html", html, 0644); nil != err {
		b.Fatalf("write spec html failed: %s", err)
	}

	for i := 0; i < b.N; i++ {
		luteEngine.Markdown("spec text", bytes)
	}
}

func BenchmarkGolangCommonMark(b *testing.B) {
	bytes, err := ioutil.ReadFile(spec + ".md")
	if nil != err {
		b.Fatalf("read spec text failed: " + err.Error())
	}

	md := markdown.New(markdown.XHTMLOutput(true),
		markdown.Tables(true),
		markdown.Linkify(true),
		markdown.Typographer(false))
	for i := 0; i < b.N; i++ {
		md.RenderToString(bytes)
	}
}

func BenchmarkGoldMark(b *testing.B) {
	markdown, err := ioutil.ReadFile(spec + ".md")
	if nil != err {
		b.Fatalf("read spec text failed: " + err.Error())
	}

	goldmarkEngine := goldmark.New(
		goldmark.WithRendererOptions(html.WithXHTML()),
		goldmark.WithExtensions(
			extension.Table, extension.Strikethrough, extension.TaskList, extension.Linkify,
		),
	)

	var out bytes.Buffer
	for i := 0; i < b.N; i++ {
		out.Reset()
		if err := goldmarkEngine.Convert(markdown, &out); nil != err {
			panic(err)
		}
	}
}

func BenchmarkBlackFriday(b *testing.B) {
	markdown, err := ioutil.ReadFile(spec + ".md")
	if nil != err {
		b.Fatalf("read spec text failed: " + err.Error())
	}
	for i := 0; i < b.N; i++ {
		blackfriday.Run(markdown)
	}
}

func BenchmarkGoMarkdown(b *testing.B) {
	markdown, err := ioutil.ReadFile(spec + ".md")
	if nil != err {
		b.Fatalf("read spec text failed: " + err.Error())
	}
	for i := 0; i < b.N; i++ {
		gm.ToHTML(markdown, nil, nil)
	}
}
