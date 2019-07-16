export interface SystemState {
  token: string;
  session: string;
}

export const UPDATE_SYSTEM = 'UPDATE_SYSTEM';

interface UpdateSystemAction {
  type: typeof UPDATE_SYSTEM;
  payload: SystemState;
}

export type SystemActionTypes = UpdateSystemAction;
