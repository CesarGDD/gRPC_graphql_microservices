import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "http://localhost:8080/query",
  documents: "src/**/*.graphql",
  generates: {
    "src/graphql/generated/": {  // Specify a file, not a directory
      preset: "client",
      plugins: [
        "typescript-urql",
      ],
      config: {
        withHooks: true
      }
    }
  }
};

export default config;
