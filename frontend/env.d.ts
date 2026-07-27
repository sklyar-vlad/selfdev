/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_AUTH_HOST: string
  readonly VITE_API_HOST: string
  readonly VITE_REDIRECT_URI: string
  readonly VITE_CASDOOR_CLIENT_ID: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}

declare module '*.vue' {
  import { DefineComponent } from 'vue'

  const component: DefineComponent<{}, {}, any>

  export default component
}
