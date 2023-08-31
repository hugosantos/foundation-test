// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package schema

import (
	"fmt"
)

func (e *Endpoint) GetServerOwnerPackage() PackageName {
	return PackageName(e.ServerOwner)
}

func (e *Endpoint) HasKind(str string) bool {
	for _, md := range e.ServiceMetadata {
		if md.GetKind() == str {
			return true
		}
	}
	return false
}

func (e *Endpoint) Address() string {
	if len(e.Ports) == 0 {
		panic("no ports")
	}

	return fmt.Sprintf("%s:%d", e.AllocatedName, e.Ports[0].GetPort().GetContainerPort())
}
