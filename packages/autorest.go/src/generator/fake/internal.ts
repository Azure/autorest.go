/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { values } from '@azure-tools/linq';
import { GoCodeModel, isLROMethod, isPageableMethod } from '../../gocodemodel/gocodemodel';
import { contentPreamble } from '../helpers';
import { ImportManager } from '../imports';

export async function generateServerInternal(codeModel: GoCodeModel): Promise<string> {
  if (codeModel.clients.length === 0) {
    return '';
  }
  const text = contentPreamble(codeModel, 'fake');
  const imports = new ImportManager();
  imports.add('io');
  imports.add('net/http');
  imports.add('reflect');
  let body = content;
  // only generate the tracker content if required
  let needsTracker = false;
  for (const client of values(codeModel.clients)) {
    for (const method of values(client.methods)) {
      if (isLROMethod(method) || isPageableMethod(method)) {
        needsTracker = true;
        break;
      }
    }
    if (needsTracker) {
      break;
    }
  }
  if (needsTracker) {
    imports.add('sync');
    imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server');
    body += tracker;
  }
  return text + imports.text() + body;
}

const content = `
type nonRetriableError struct {
	error
}

func (nonRetriableError) NonRetriable() {
	// marker method
}

func getOptional[T any](v T) *T {
	if reflect.ValueOf(v).IsZero() {
		return nil
	}
	return &v
}

func getHeaderValue(h http.Header, k string) string {
	v := h[k]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

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

func parseWithCast[T any](v string, parse func(v string) (T, error)) (T, error) {
	t, err := parse(v)
	if err != nil {
		return *new(T), err
	}
	return t, err
}

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

func contains[T comparable](s []T, v T) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}
`;

const tracker = `
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
