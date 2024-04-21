import React from 'react';
import ReactDOM from 'react-dom';
import { Provider as URQLProvider, Client, cacheExchange, fetchExchange } from 'urql';
import { Provider as ReduxProvider } from 'react-redux';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { store } from './store/store';

// Create an instance of urql client
const client = new Client({
  url: 'http://localhost:8080/query',
  exchanges: [cacheExchange, fetchExchange],
});

ReactDOM.render(
  <React.StrictMode>
    <ReduxProvider store={store}>
      <URQLProvider value={client}>
        <App />
      </URQLProvider>
    </ReduxProvider>
  </React.StrictMode>,
  document.getElementById('root')
);

// Optional: setup for measuring performance in your app
reportWebVitals();
