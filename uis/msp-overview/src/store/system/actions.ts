import { Dispatch } from 'react';

// Request
import { post } from '../../util/request';

// Types
import { SystemPayload, UPDATE_SYSTEM, SystemActionTypes } from './types';
// type DispatchFn = (action: SystemActionTypes) => Dispatch<any>;

export function updateSystemAction(
  newSystemState: SystemPayload,
): SystemActionTypes {
  return {
    type: UPDATE_SYSTEM,
    payload: newSystemState,
  };
}

/* eslint-disable no-console*/
export const login = (username: string, password: string): any => (
  dispatch: Dispatch<SystemActionTypes>,
): any => {
  return post({ username, password })
    .then((response: Response): Promise<{ token: string }> => {
      return response.json();
    })
    .then((data: { token: string }): string => {
      console.log(data);
      dispatch(updateSystemAction({ token: data.token }));
      return data.token;
    })
    .catch((error: any): any => console.log(error));
};
/* eslint-enable no-console*/
