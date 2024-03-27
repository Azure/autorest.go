/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../../../codemodel.go/src/gocodemodel.js';
import { contentPreamble } from '../helpers.js';
import { ImportManager } from '../imports.js';

export class RequiredHelpers {
  getHeaderValue: boolean;
  getOptional: boolean;
  initServer: boolean;
  parseOptional: boolean;
  parseWithCast: boolean;
  readRequestBody: boolean;
  splitHelper: boolean;
  tracker: boolean;

  constructor() {
    this.getHeaderValue = false;
    this.getOptional = false;
    this.initServer = false;
    this.parseOptional = false;
    this.parseWithCast = false;
    this.readRequestBody = false;
    this.splitHelper = false;
    this.tracker = false;
  }
}

export function generateServerInternal(codeModel: go.CodeModel, requiredHelpers: RequiredHelpers): string {
  if (codeModel.clients.length === 0) {
    return '';
  }
  const text = contentPreamble(codeModel, 'fake');
  const imports = new ImportManager();
  let body = alwaysUsed;

  if (requiredHelpers.getHeaderValue) {
    body += emitGetHeaderValue(imports);
  }
  if (requiredHelpers.getOptional) {
    body += emitGetOptional(imports);
  }
  if (requiredHelpers.initServer) {
    body += emitInitServer(imports);
  }
  if (requiredHelpers.parseOptional) {
    body += emitParseOptional();
  }
  if (requiredHelpers.parseWithCast) {
    body += emitParseWithCast();
  }
  if (requiredHelpers.readRequestBody) {
    body += emitReadRequestBody(imports);
  }
  if (requiredHelpers.splitHelper) {
    body += emitSplitHelper(imports);
  }
  if (requiredHelpers.tracker) {
    body += emitTracker(imports);
  }

  return text + imports.text() + body;
}

// contains helpers that are used in all servers
const alwaysUsed = `
type nonRetriableError struct {
	error
}

func (nonRetriableError) NonRetriable() {
	// marker method
}

func contains[T comparable](s []T, v T) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}
`;

function emitGetOptional(imports: ImportManager): string {
  imports.add('reflect');
  return `
func getOptional[T any](v T) *T {
	if reflect.ValueOf(v).IsZero() {
		return nil
	}
	return &v
}
`;
}

function emitGetHeaderValue(imports: ImportManager): string {
  imports.add('net/http');
  return `
func getHeaderValue(h http.Header, k string) string {
	v := h[k]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}
`;
}

function emitInitServer(imports: ImportManager): string {
  imports.add('sync');
  return `
func initServer[T any](mu *sync.Mutex, dst **T, src func() *T) {
	mu.Lock()
	if *dst == nil {
		*dst = src()
	}
	mu.Unlock()
}
`;
}

function emitParseOptional(): string {
  return `
func parseOptional[T any](v string, parse func(v string) (T, error)) (*T, error) {
	if v == "" {
		return nil, nil
	}
	t, err := parse(v)
	if err != nil {
		return nil, err
	}
	return &t, err
}
`;
}

function emitParseWithCast(): string {
  return `
func parseWithCast[T any](v string, parse func(v string) (T, error)) (T, error) {
	t, err := parse(v)
	if err != nil {
		return *new(T), err
	}
	return t, err
}
`;
}

function emitReadRequestBody(imports: ImportManager): string {
  imports.add('net/http');
  imports.add('io');
  return `
func readRequestBody(req *http.Request) ([]byte, error) {
	if req.Body == nil {
		return nil, nil
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	req.Body.Close()
	return body, nil
}
`;
}

function emitSplitHelper(imports: ImportManager): string {
  imports.add('strings');
  return `
func splitHelper(s, sep string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, sep)
}
`;
}

function emitTracker(imports: ImportManager): string {
  imports.add('net/http');
  imports.add('sync');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server');
  return `
func newTracker[T any]() *tracker[T] {
	return &tracker[T]{
		items: map[string]*T{},
	}
}

type tracker[T any] struct {
	items map[string]*T
	mu sync.Mutex
}

func (p *tracker[T]) get(req *http.Request) *T {
	p.mu.Lock()
	defer p.mu.Unlock()
	if item, ok := p.items[server.SanitizePagerPollerPath(req.URL.Path)]; ok {
		return item
	}
	return nil
}

func (p *tracker[T]) add(req *http.Request, item *T) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.items[server.SanitizePagerPollerPath(req.URL.Path)] = item
}

func (p *tracker[T]) remove(req *http.Request) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.items, server.SanitizePagerPollerPath(req.URL.Path))
}
`;
}
