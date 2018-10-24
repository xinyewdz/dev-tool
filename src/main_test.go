package main

import (
	"testing"
	"tool"
)

func TestRename(t *testing.T) {
	tool.Rename("D:/data/device/", "D:/data/device/device_new")
}
