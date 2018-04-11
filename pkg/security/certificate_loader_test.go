// Copyright 2017 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package security_test

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/cockroachdb/cockroach/pkg/security"
	"github.com/cockroachdb/cockroach/pkg/security/securitytest"
	"github.com/cockroachdb/cockroach/pkg/testutils"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/timeutil"
)

func TestLoadEmbeddedCerts(t *testing.T) {
	defer leaktest.AfterTest(t)()
	cl := security.NewCertificateLoader(security.EmbeddedCertsDir)
	if err := cl.Load(); err != nil {
		t.Error(err)
	}

	assets, err := securitytest.AssetReadDir(security.EmbeddedCertsDir)
	if err != nil {
		t.Fatal(err)
	}

	// Check that we have "found pairs * 2 = num assets".
	certs := cl.Certificates()
	if act, exp := len(certs), len(assets); act*2 != exp {
		t.Errorf("found %d keypairs, but have %d embedded files", act, exp)
	}

	// Check that all non-CA pairs include a key.
	for _, c := range certs {
		if c.FileUsage == security.CAPem {
			if len(c.KeyFilename) != 0 {
				t.Errorf("CA key was loaded for CertInfo %+v", c)
			}
		} else if len(c.KeyFilename) == 0 {
			t.Errorf("no key found as part of CertInfo %+v", c)
		}
	}
}

func countLoadedCertificates(certsDir string) (int, error) {
	cl := security.NewCertificateLoader(certsDir)
	if err := cl.Load(); err != nil {
		return 0, nil
	}
	return len(cl.Certificates()), nil
}

// Generate a x509 cert with specific fields.
func makeTestCert(
	t *testing.T, commonName string, keyUsage x509.KeyUsage, extUsages []x509.ExtKeyUsage,
) []byte {
	// Make smallest rsa key possible: not saved.
	key, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		t.Fatalf("error on GenerateKey for CN=%s: %v", commonName, err)
	}

	// Specify the smallest possible set of fields.
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: commonName,
		},
		NotBefore: timeutil.Now().Add(-time.Hour),
		NotAfter:  timeutil.Now().Add(time.Hour),
		KeyUsage:  keyUsage,
	}

	template.ExtKeyUsage = extUsages

	certBytes, err := x509.CreateCertificate(rand.Reader, template, template, key.Public(), key)
	if err != nil {
		t.Fatalf("error on CreateCertificate for CN=%s: %v", commonName, err)
	}

	certBlock := &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}
	return pem.EncodeToMemory(certBlock)
}

func TestNamingScheme(t *testing.T) {
	defer leaktest.AfterTest(t)()

	fullKeyUsage := x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature
	// Build a few certificates. These are barebones since we only need to check our custom validation,
	// not chain verification.
	caCert := makeTestCert(t, "", 0, nil)

	goodNodeCert := makeTestCert(t, "node", fullKeyUsage, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth})
	badUserNodeCert := makeTestCert(t, "notnode", fullKeyUsage, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth})
	noServerAuthNodeCert := makeTestCert(t, "node", fullKeyUsage, []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth})
	noClientAuthNodeCert := makeTestCert(t, "node", fullKeyUsage, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth})
	noAuthNodeCert := makeTestCert(t, "node", fullKeyUsage, nil)
	noEnciphermentNodeCert := makeTestCert(t, "node", x509.KeyUsageDigitalSignature, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth})
	noSignatureNodeCert := makeTestCert(t, "node", x509.KeyUsageKeyEncipherment, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth})

	goodRootCert := makeTestCert(t, "root", fullKeyUsage, []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth})
	notRootCert := makeTestCert(t, "notroot", fullKeyUsage, []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth})
	noClientAuthRootCert := makeTestCert(t, "root", fullKeyUsage, nil)
	noEnciphermentRootCert := makeTestCert(t, "root", x509.KeyUsageDigitalSignature, []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth})
	noSignatureRootCert := makeTestCert(t, "root", x509.KeyUsageKeyEncipherment, []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth})

	// Do not use embedded certs.
	security.ResetAssetLoader()
	defer ResetTest()

	// Some test cases are skipped on windows due to non-UGO permissions.
	isWindows := runtime.GOOS == "windows"

	// Test non-existent directory.
	// If the directory exists, we still expect no failures, unless it happens to contain
	// valid filenames, so we don't need to try too hard to generate a unique name.
	if count, err := countLoadedCertificates("my_non_existent_directory-only_for_tests"); err != nil {
		t.Error(err)
	} else if exp := 0; exp != count {
		t.Errorf("found %d certificates, expected %d", count, exp)
	}

	// Create directory.
	certsDir, err := ioutil.TempDir("", "certs_test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(certsDir); err != nil {
			t.Fatal(err)
		}
	}()

	type testFile struct {
		name     string
		mode     os.FileMode
		contents []byte
	}

	testData := []struct {
		// Files to write (name and mode).
		files []testFile
		// Certinfos found. Ordered by cert base filename.
		certs []security.CertInfo
		// Set to true to skip permissions checks.
		skipChecks bool
		// Skip test case on windows as permissions are always ignored.
		skipWindows bool
	}{
		{
			// Empty directory.
			files: []testFile{},
			certs: []security.CertInfo{},
		},
		{
			// Test bad names, including ca/node certs with blobs in the middle, wrong separator.
			// We only need to test certs, if they're not loaded, neither will keys.
			files: []testFile{
				{"ca.foo.crt", 0777, []byte{}},
				{"cr..crt", 0777, []byte{}},
				{"node.foo.crt", 0777, []byte{}},
				{"node..crt", 0777, []byte{}},
				{"client.crt", 0777, []byte{}},
				{"client..crt", 0777, []byte{}},
			},
			certs: []security.CertInfo{},
		},
		{
			// Test proper names, but no key files, only the CA cert should be loaded without error.
			files: []testFile{
				{"ca.crt", 0777, caCert},
				{"node.crt", 0777, goodNodeCert},
				{"client.root.crt", 0777, goodRootCert},
			},
			certs: []security.CertInfo{
				{FileUsage: security.CAPem, Filename: "ca.crt", FileContents: caCert},
				{FileUsage: security.ClientPem, Filename: "client.root.crt", Name: "root",
					Error: errors.New(".* no such file or directory")},
				{FileUsage: security.NodePem, Filename: "node.crt",
					Error: errors.New(".* no such file or directory")},
			},
		},
		{
			// Key files, but wrong permissions.
			// We don't load CA keys here, so permissions for them don't matter.
			files: []testFile{
				{"ca.crt", 0777, caCert},
				{"ca.key", 0777, []byte{}},
				{"node.crt", 0777, goodNodeCert},
				{"node.key", 0704, []byte{}},
				{"client.root.crt", 0777, goodRootCert},
				{"client.root.key", 0740, []byte{}},
			},
			certs: []security.CertInfo{
				{FileUsage: security.CAPem, Filename: "ca.crt", FileContents: caCert},
				{FileUsage: security.ClientPem, Filename: "client.root.crt", Name: "root",
					Error: errors.New(".* exceeds -rwx------")},
				{FileUsage: security.NodePem, Filename: "node.crt",
					Error: errors.New(".* exceeds -rwx------")},
			},
			skipWindows: true,
		},
		{
			// Bad cert files.
			files: []testFile{
				{"ca.crt", 0777, []byte{}},
				{"ca.key", 0777, []byte{}},
				{"node.crt", 0777, []byte("foo")},
				{"node.key", 0700, []byte{}},
				{"client.root.crt", 0777, append(goodRootCert, []byte("-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----")...)},
				{"client.root.key", 0700, []byte{}},
			},
			certs: []security.CertInfo{
				{FileUsage: security.CAPem, Filename: "ca.crt",
					Error: errors.New("empty certificate file: ca.crt")},
				{FileUsage: security.ClientPem, Filename: "client.root.crt", Name: "root",
					Error: errors.New("failed to parse certificate 1 in file client.root.crt")},
				{FileUsage: security.NodePem, Filename: "node.crt",
					Error: errors.New("no certificates found in node.crt")},
			},
		},
		{
			// Bad CommonName.
			files: []testFile{
				{"node.crt", 0777, badUserNodeCert},
				{"node.key", 0700, []byte{}},
				{"client.root.crt", 0777, notRootCert},
				{"client.root.key", 0700, []byte{}},
			},
			certs: []security.CertInfo{
				{FileUsage: security.ClientPem, Filename: "client.root.crt", Name: "root",
					Error: errors.New("client certificate has Subject \"CN=notroot\", expected \"CN=root")},
				{FileUsage: security.NodePem, Filename: "node.crt",
					Error: errors.New("node certificate has Subject \"CN=notnode\", expected \"CN=node")},
			},
		},
		{
			// No ServerAuth key usage.
			files: []testFile{
				{"node.crt", 0777, noServerAuthNodeCert},
				{"node.key", 0700, []byte{}},
			},
			certs: []security.CertInfo{
				{FileUsage: security.NodePem, Filename: "node.crt",
					Error: errors.New("node certificate extended key usages: ServerAuth=false, ClientAuth=true, but both are needed")},
			},
		},
		{
			// No ClientAuth key usage.
			files: []testFile{
				{"node.crt", 0777, noClientAuthNodeCert},
				{"node.key", 0700, []byte{}},
				{"client.root.crt", 0777, noClientAuthRootCert},
				{"client.root.key", 0700, []byte{}},
			},
			certs: []security.CertInfo{
				{FileUsage: security.ClientPem, Filename: "client.root.crt", Name: "root",
					Error: errors.New("client certificate does not have ClientAuth extended key usage")},
				{FileUsage: security.NodePem, Filename: "node.crt",
					Error: errors.New("node certificate extended key usages: ServerAuth=true, ClientAuth=false, but both are needed")},
			},
		},
		{
			// No auth key usage.
			files: []testFile{
				{"node.crt", 0777, noAuthNodeCert},
				{"node.key", 0700, []byte{}},
			},
			certs: []security.CertInfo{
				{FileUsage: security.NodePem, Filename: "node.crt",
					Error: errors.New("node certificate extended key usages: ServerAuth=false, ClientAuth=false, but both are needed")},
			},
		},
		{
			// No KeyEncipherment key usage.
			files: []testFile{
				{"node.crt", 0777, noEnciphermentNodeCert},
				{"node.key", 0700, []byte{}},
				{"client.root.crt", 0777, noEnciphermentRootCert},
				{"client.root.key", 0700, []byte{}},
			},
			certs: []security.CertInfo{
				{FileUsage: security.NodePem, Filename: "client.root.crt",
					Error: errors.New("client certificate key usages: KeyEncipherment=false, DigitalSignature=true, but both are needed")},
				{FileUsage: security.NodePem, Filename: "node.crt",
					Error: errors.New("node certificate key usages: KeyEncipherment=false, DigitalSignature=true, but both are needed")},
			},
		},
		{
			// No DigitalSignature key usage.
			files: []testFile{
				{"node.crt", 0777, noSignatureNodeCert},
				{"node.key", 0700, []byte{}},
				{"client.root.crt", 0777, noSignatureRootCert},
				{"client.root.key", 0700, []byte{}},
			},
			certs: []security.CertInfo{
				{FileUsage: security.NodePem, Filename: "client.root.crt",
					Error: errors.New("client certificate key usages: KeyEncipherment=true, DigitalSignature=false, but both are needed")},
				{FileUsage: security.NodePem, Filename: "node.crt",
					Error: errors.New("node certificate key usages: KeyEncipherment=true, DigitalSignature=false, but both are needed")},
			},
		},
		{
			// Everything loads.
			files: []testFile{
				{"ca.crt", 0777, caCert},
				{"ca.key", 0700, []byte("ca.key")},
				{"node.crt", 0777, goodNodeCert},
				{"node.key", 0700, []byte("node.key")},
				{"client.root.crt", 0777, goodRootCert},
				{"client.root.key", 0700, []byte("client.root.key")},
			},
			certs: []security.CertInfo{
				{FileUsage: security.CAPem, Filename: "ca.crt", FileContents: caCert},
				{FileUsage: security.ClientPem, Filename: "client.root.crt", KeyFilename: "client.root.key",
					Name: "root", FileContents: goodRootCert, KeyFileContents: []byte("client.root.key")},
				{FileUsage: security.NodePem, Filename: "node.crt", KeyFilename: "node.key",
					FileContents: goodNodeCert, KeyFileContents: []byte("node.key")},
			},
		},
		{
			// Bad key permissions, but skip permissions checks.
			files: []testFile{
				{"ca.crt", 0777, caCert},
				{"ca.key", 0777, []byte("ca.key")},
				{"node.crt", 0777, goodNodeCert},
				{"node.key", 0777, []byte("node.key")},
				{"client.root.crt", 0777, goodRootCert},
				{"client.root.key", 0777, []byte("client.root.key")},
			},
			certs: []security.CertInfo{
				{FileUsage: security.CAPem, Filename: "ca.crt", FileContents: caCert},
				{FileUsage: security.ClientPem, Filename: "client.root.crt", KeyFilename: "client.root.key",
					Name: "root", FileContents: goodRootCert, KeyFileContents: []byte("client.root.key")},
				{FileUsage: security.NodePem, Filename: "node.crt", KeyFilename: "node.key",
					FileContents: goodNodeCert, KeyFileContents: []byte("node.key")},
			},
			skipChecks: true,
		},
	}

	for testNum, data := range testData {
		if data.skipWindows && isWindows {
			continue
		}

		// Write all files.
		for _, f := range data.files {
			n := f.name
			if err := ioutil.WriteFile(filepath.Join(certsDir, n), f.contents, f.mode); err != nil {
				t.Fatalf("#%d: could not write file %s: %v", testNum, n, err)
			}
		}

		// Load certs.
		cl := security.NewCertificateLoader(certsDir)
		if data.skipChecks {
			cl.TestDisablePermissionChecks()
		}
		if err := cl.Load(); err != nil {
			t.Errorf("#%d: unexpected error: %v", testNum, err)
		}

		// Check count of certificates.
		if expected, actual := len(data.certs), len(cl.Certificates()); expected != actual {
			t.Errorf("#%d: expected %d certificates, found %d", testNum, expected, actual)
		}

		// Check individual certificates.
		for i, actual := range cl.Certificates() {
			expected := data.certs[i]

			if expected.Error == nil {
				if actual.Error != nil {
					t.Errorf("#%d: expected success, got error: %+v", testNum, actual.Error)
					continue
				}
			} else {
				if !testutils.IsError(actual.Error, expected.Error.Error()) {
					t.Errorf("#%d: mismatched error, expected: %+v, got %+v", testNum, expected.Error, actual.Error)
				}
				continue
			}

			// Compare some fields.
			if actual.FileUsage != expected.FileUsage ||
				actual.Filename != expected.Filename ||
				actual.KeyFilename != expected.KeyFilename ||
				actual.Name != expected.Name {
				t.Errorf("#%d: mismatching CertInfo, expected: %+v, got %+v", testNum, expected, actual)
				continue
			}
			if actual.Filename != "" {
				if !bytes.Equal(actual.FileContents, expected.FileContents) {
					t.Errorf("#%d: bad file contents: expected %s, got %s", testNum, expected.FileContents, actual.FileContents)
					continue
				}
				if a, e := len(actual.ParsedCertificates), 1; a != e {
					t.Errorf("#%d: expected %d certificates, found: %d", testNum, e, a)
					continue
				}
			}
			if actual.KeyFilename != "" && !bytes.Equal(actual.KeyFileContents, expected.KeyFileContents) {
				t.Errorf("#%d: bad file contents: expected %s, got %s", testNum, expected.KeyFileContents, actual.KeyFileContents)
				continue
			}
		}

		// Wipe all files.
		for _, f := range data.files {
			n := f.name
			if err := os.Remove(filepath.Join(certsDir, n)); err != nil {
				t.Fatalf("#%d: could not delete file %s: %v", testNum, n, err)
			}
		}
	}
}
