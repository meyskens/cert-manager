/*
Copyright 2019 The Jetstack cert-manager contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package validation

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"

	cmapi "github.com/jetstack/cert-manager/pkg/internal/apis/certmanager"
	"github.com/jetstack/cert-manager/pkg/util/pki"
)

func ValidateCertificateRequest(obj runtime.Object) field.ErrorList {
	cr := obj.(*cmapi.CertificateRequest)
	allErrs := ValidateCertificateRequestSpec(&cr.Spec, field.NewPath("spec"))
	return allErrs
}

func ValidateCertificateRequestSpec(crSpec *cmapi.CertificateRequestSpec, fldPath *field.Path) field.ErrorList {
	el := field.ErrorList{}

	el = append(el, validateIssuerRef(crSpec.IssuerRef, fldPath)...)

	_, err := pki.DecodeX509CertificateRequestBytes(crSpec.CSRPEM)
	if err != nil {
		el = append(el, field.Invalid(fldPath.Child("csr"), crSpec.CSRPEM, fmt.Sprintf("failed to decode csr: %s", err)))
	}

	return el
}
