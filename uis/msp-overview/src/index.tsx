import React from 'react';
import ReactDOM from 'react-dom';

// CSS Reset
import '@csstools/normalize.css';
import './index.css';

// Root Component
import Root from './components/root';

// Store
import configureStore from './store';

// Service Worker
import * as serviceWorker from './service-worker';

document.addEventListener('DOMContentLoaded', (): void => {
  const rootElement = document.getElementById('root');
  const store = configureStore();

  ReactDOM.render(<Root store={store} />, rootElement);
});

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
