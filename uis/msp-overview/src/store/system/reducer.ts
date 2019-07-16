import { UPDATE_SYSTEM, SystemState, SystemActionTypes } from './types';

const initialState: SystemState = {
  token: 'fake-token',
  session: '',
};

export function systemReducer(
  state = initialState,
  action: SystemActionTypes,
): SystemState {
  switch (action.type) {
    case UPDATE_SYSTEM: {
      return {
        ...state,
        ...action.payload,
      };
    }

    default:
      return state;
  }
}
