import { describe, it, expect } from 'vitest';
import { camelCase, pascalCase } from '../../../codegen.go/src/core/helpers.js';

describe('pascalCase', () => {
  it('preserves trailing duplicate words (does not remove duplicates)', () => {
    // regression: previously "AuthorizationServerServer" collapsed to
    // "AuthorizationServer", causing duplicate TransportInterceptor variables
    // in the generated fake package.
    expect(pascalCase('AuthorizationServerServer')).toEqual('AuthorizationServerServer');
    expect(pascalCase('authorization_server_server')).toEqual('AuthorizationServerServer');
    expect(pascalCase(['authorization', 'server', 'server'])).toEqual('AuthorizationServerServer');
  });

  it('capitalizes single word', () => {
    expect(pascalCase('foo')).toEqual('Foo');
  });

  it('returns empty for undefined or empty input', () => {
    expect(pascalCase(undefined as unknown as string)).toEqual('');
    expect(pascalCase('')).toEqual('');
    expect(pascalCase([])).toEqual('');
  });
});

describe('camelCase', () => {
  it('preserves trailing duplicate words (does not remove duplicates)', () => {
    expect(camelCase('AuthorizationServerServer')).toEqual('authorizationServerServer');
    expect(camelCase(['authorization', 'server', 'server'])).toEqual('authorizationServerServer');
  });

  it('uncapitalizes single word', () => {
    expect(camelCase('Foo')).toEqual('foo');
  });
});
