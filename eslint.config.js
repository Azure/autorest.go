// @ts-check

import eslint from '@eslint/js';
import tseslint from 'typescript-eslint';

export default tseslint.config([
  {
    files: ['**/packages/*/src/**/*.ts', '**/packages/*/test/**/*.ts'],
    extends: [
      eslint.configs.recommended,
      ...tseslint.configs.recommendedTypeChecked,
      {
        languageOptions: {
          parserOptions: {
            projectService: true,
            tsconfigRootDir: import.meta.dirname,
          },
        },
      },
    ],
    rules: {
      indent: [
        'warn',
        2,
        {
          SwitchCase: 1,
        },
      ],
      '@typescript-eslint/no-floating-promises': 'error',
      '@typescript-eslint/no-unsafe-declaration-merging': 'off',
    },
  },
  {
    ignores: ['**/dist/**/*'],
  },
]);
