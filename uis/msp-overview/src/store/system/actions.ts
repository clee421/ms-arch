import { SystemState, UPDATE_SYSTEM, SystemActionTypes } from './types';

export function updateSystem(newSystemState: SystemState): SystemActionTypes {
  return {
    type: UPDATE_SYSTEM,
    payload: newSystemState,
  };
}
