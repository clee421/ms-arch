export interface SystemState {
  token: string;
  session: string;
}

export interface SystemPayload {
  token?: string;
  session?: string;
}

export const UPDATE_SYSTEM = 'UPDATE_SYSTEM';

interface UpdateSystemAction {
  type: typeof UPDATE_SYSTEM;
  payload: SystemPayload;
}

export type SystemActionTypes = UpdateSystemAction;
