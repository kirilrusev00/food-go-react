import {
  useCallback, useRef, useState, useEffect,
} from 'react';

interface AsyncState<T> {
  data: T | undefined,
  /* eslint-disable-next-line @typescript-eslint/no-explicit-any */
  error: any,
  loading: boolean
}

/* eslint-disable-next-line @typescript-eslint/no-explicit-any */
export default function useAsyncAction<T>(action: () => Promise<T>, dependencies: any[]) {
  const [state, setState] = useState<AsyncState<T>>({
    data: undefined,
    loading: false,
    error: undefined,
  });

  const isCancelled = useRef(false);

  const perform = useCallback(() => {
    setState({ data: undefined, loading: true, error: undefined });

    (async function tryAction() {
      try {
        const data = await action();

        if (!isCancelled.current) {
          setState({ data, loading: false, error: undefined });
        }
      } catch (error) {
        if (!isCancelled.current) {
          setState({ data: undefined, loading: false, error });
        }
      }
    }());
  }, dependencies);

  useEffect(() => () => {
    isCancelled.current = true;
  }, []);

  return { ...state, perform };
}
