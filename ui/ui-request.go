// Copyright (c) 2024  The Go-CoreUtils Authors
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

package ui

import (
	"fmt"

	"github.com/go-curses/cdk/log"

	editor "github.com/go-coreutils/etc-hosts-editor"
)

func (c *cUI) requestReload() {
	log.DebugF("reloading from: %v", c.SourceFile)
	var err error
	if c.HostFile, err = editor.ParseFile(c.SourceFile); err != nil {
		c.LastError = fmt.Errorf("error parsing %v: %v", c.SourceFile, err)
		log.Error(c.LastError)
		return
	}
	c.requestReloadContents()
}

func (c *cUI) requestReloadContents() {
	c.reloadEditor()
	c.focusEditor(nil)
}

func (c *cUI) requestSave() {
	log.DebugF("saving to: %v", c.SourceFile)
	if c.HostFile != nil {
		if err := c.HostFile.Save(); err != nil {
			log.Error(err)
		}
	}
	c.requestReload()
	c.QuitButton.GrabFocus()
}

func (c *cUI) requestQuit() {
	c.Display.RequestQuit()
}
