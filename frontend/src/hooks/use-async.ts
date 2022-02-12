import { useEffect } from 'react';
import useAsyncAction from './use-async-action';

/* eslint-disable-next-line @typescript-eslint/no-explicit-any */
export default function useAsync<T>(action: () => Promise<T>, dependencies: any[]) {
  const {
    perform, data, error, loading,
  } = useAsyncAction(action, [...dependencies, action]);

  useEffect(() => {
    perform();
  }, dependencies);

  return {
    data, error, loading, reload: perform,
  };
}
