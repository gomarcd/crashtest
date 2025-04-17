import js from '@eslint/js';
import tseslint from 'typescript-eslint';
import pluginVue from 'eslint-plugin-vue';
import parserVue from 'vue-eslint-parser';
import globals from 'globals';

const vueRecommendedConfigs = (pluginVue && pluginVue.configs && Array.isArray(pluginVue.configs['flat/recommended']))
    ? pluginVue.configs['flat/recommended']
    : [];

if (vueRecommendedConfigs.length === 0) {
    console.warn("WARNING: Vue 'flat/recommended' config is not an array or is missing.");
}

export default [
    {
        ignores: ["dist/**", "node_modules/**", "wailsjs/**"],
    },
    js.configs.recommended,
    ...vueRecommendedConfigs,
    ...tseslint.configs.recommendedTypeChecked.map(config => ({
        ...config,
        files: ['src/**/*.{vue,ts,mts,cts}'],
    })),
    {
        files: ['src/**/*.vue'],
        plugins: {
            vue: pluginVue,
            '@typescript-eslint': tseslint.plugin,
        },
        languageOptions: {
            parser: parserVue,
            parserOptions: {
                parser: tseslint.parser,
                project: true,
                tsconfigRootDir: import.meta.dirname,
                extraFileExtensions: ['.vue'],
                sourceType: 'module',
                ecmaVersion: 'latest',
            },
            globals: { ...globals.browser },
        },
        rules: {
            'vue/multi-word-component-names': 'off',
            '@typescript-eslint/no-unused-vars': ['warn', { argsIgnorePattern: '^_' }],
            '@typescript-eslint/no-explicit-any': 'warn',
            '@typescript-eslint/no-unsafe-return': 'warn',
            '@typescript-eslint/no-unsafe-call': 'warn',
            '@typescript-eslint/no-unsafe-member-access': 'warn',
            '@typescript-eslint/no-floating-promises': 'warn',
        },
    },
    {
        files: ['src/**/*.{ts,mts,cts}'],
        plugins: {
            '@typescript-eslint': tseslint.plugin,
        },
        languageOptions: {
            parser: tseslint.parser,
            parserOptions: {
                project: true,
                tsconfigRootDir: import.meta.dirname,
            },
        },
        rules: {
            '@typescript-eslint/no-explicit-any': 'warn',
            '@typescript-eslint/no-namespace': 'off',
            '@typescript-eslint/no-empty-object-type': 'warn',
            '@typescript-eslint/no-floating-promises': 'error',
        },
    },
    {
        files: ['vite.config.ts', '*.config.js', '*.config.mjs', '*.config.cjs'],
        plugins: {
            '@typescript-eslint': tseslint.plugin,
        },
        languageOptions: {
            parser: tseslint.parser,
            parserOptions: {
                sourceType: 'module',
                ecmaVersion: 'latest',
            },
            globals: { ...globals.node },
        },
        rules: {
            '@typescript-eslint/no-unused-vars': ['warn', { argsIgnorePattern: '^_' }],
            '@typescript-eslint/no-explicit-any': 'off',
            '@typescript-eslint/await-thenable': 'off',
            '@typescript-eslint/no-floating-promises': 'off',
            '@typescript-eslint/no-misused-promises': 'off',
            '@typescript-eslint/no-unsafe-argument': 'off',
            '@typescript-eslint/no-unsafe-assignment': 'off',
            '@typescript-eslint/no-unsafe-call': 'off',
            '@typescript-eslint/no-unsafe-member-access': 'off',
            '@typescript-eslint/no-unsafe-return': 'off',
            '@typescript-eslint/restrict-plus-operands': 'off',
            '@typescript-eslint/restrict-template-expressions': 'off',
        },
    },
    {
        files: ['src/**/*.{js,mjs,cjs}'],
        languageOptions: {
            globals: { ...globals.browser, ...globals.node },
        },
        rules: {
        },
    },
];