package utils

import (
	"fmt"
	"runtime"
)

var (
	Version	   string
	BuildDate  string
	Tag		   string
	ClosestTag string
)

type PkgVersion struct {
	Version	  string
	BuildType string
	BuildDate string
	Tag		  string
	GoVersion string
}

func CurrentVersion() PkgVersion {
	pkgVersion := PkgVersion{
		Version:   Version,
		BuildType: BuildType,
		BuildDate: BuildDate,
		Tag:	   Tag,
		GoVersion: runtime.Version(),
	}
	if pkgVersion.Tag == "" {
		pkgVersion.Version = fmt.Sprintf("untagged (%s)", ClosestTag)
	}
	return pkgVersion
}

func (s PkgVersion) String() string {
	if s.Tag == "" {
		return fmt.Sprintf("pkg version: %s (%s) %s %s", s.Version, s.BuildType, s.BuildDate, s.GoVersion)
	}
	return fmt.Sprintf("pkg version: %s (%s) (tagged as %q) %s %s", s.Version, s.BuildType, s.Tag, s.BuildDate, s.GoVersion)
}
