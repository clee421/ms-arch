import React, { ReactElement } from 'react';
import { Store } from 'redux';
import { Provider } from 'react-redux';
import App from './app';

interface Props {
  store: Store;
}

const Root = ({ store }: Props): ReactElement => {
  return (
    <Provider store={store}>
      <App />
    </Provider>
  );
};

export default Root;
