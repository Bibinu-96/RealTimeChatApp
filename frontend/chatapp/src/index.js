import React from 'react';
import { createRoot } from 'react-dom/client';
import App from './App';
import { SnackbarProvider} from 'notistack';
import * as serviceWorkerRegistration from './serviceWorkerRegistration';

const container = document.getElementById('root');
const root = createRoot(container);

root.render(
  <SnackbarProvider maxSnack={10}>

  <React.StrictMode>
    <App />
  </React.StrictMode>
  </SnackbarProvider>
);

// Register service worker
serviceWorkerRegistration.register();
