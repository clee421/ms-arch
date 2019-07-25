import { useDispatch } from 'react-redux';
import { Dispatch } from 'react';

type ActionDispatchFunction<T> = (dispatch: Dispatch<T>) => void;
type DispatchFunction<T> = (...args: T[]) => void;

export function useFunctionDispatch<T = unknown>(
  func: (...args: T[]) => ActionDispatchFunction<T>,
): DispatchFunction<T> {
  const dispatch = useDispatch();
  const dispatchFn = (...args: T[]): void => {
    func(...args)(dispatch);
  };

  return dispatchFn;
}
