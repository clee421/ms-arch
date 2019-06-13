import React from 'react';
// import logo from './logo.svg';
import './app.scss';

const App: React.FC = () => {
  return (
    <div className="app-container">
      <div className="input-boxes">
        <label>
          username
          <input type="text" name="username" />
        </label>
        <label>
          password
          <input type="text" name="password" />
        </label>
        <label>
          token
          <input type="text" name="token" />
        </label>
      </div>
    </div>
  );
}

export default App;
