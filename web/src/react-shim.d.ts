// Temporary shim to satisfy TypeScript when node_modules are not installed.
// Provides minimal React surface for hooks/JSX typing.
declare module "react" {
  export function useState<T>(initial: T): [T, (value: T | ((prev: T) => T)) => void];
  export function useEffect(effect: () => void | (() => void), deps?: any[]): void;
  export function useMemo<T>(factory: () => T, deps?: any[]): T;
  export const StrictMode: any;
  export const Fragment: any;
  const React: any;
  export default React;
}

declare module "react/jsx-runtime" {
  const jsx: any;
  const jsxs: any;
  const Fragment: any;
  export { jsx, jsxs, Fragment };
}

declare module "react-dom/client" {
  export function createRoot(container: any): { render: (c: any) => void };
}

declare namespace JSX {
  interface IntrinsicElements {
    [elemName: string]: any;
  }
}
