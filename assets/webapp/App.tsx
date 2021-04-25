import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import MuiCssBaseline from '@material-ui/core/CssBaseline';
import { createGlobalStyle } from 'styled-components';
import { HashRouter as Router, Switch, Route } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from 'react-query';
import { ReactQueryDevtools } from 'react-query/devtools';

// Pages
import LoginPage from './pages/Login';
import SelectEnvironmentPage from './pages/SelectEnvironment';
import React from 'react';
import ApplicationsPage from './pages/Applications';
import ApplicationPage from './pages/Application';

// Styled Components Global
const GlobalStyle = createGlobalStyle`
  body {
    margin: 0;
    padding: 0;
    font-family: Open-Sans, Helvetica, Sans-Serif;
  }
`;

// Material UI Theme
const theme = createMuiTheme({
  palette: {
    primary: {
      main: '#0C2461',
    },
  },
});

// Setup React Query
const queryClient = new QueryClient();

const App = () => (
  <>
    <ThemeProvider theme={theme}>
      <QueryClientProvider client={queryClient}>
        <MuiCssBaseline />
        <GlobalStyle />
        <Router>
          <Switch>
            <Route path="/login" component={LoginPage} />
            <Route path="/select-environment" component={SelectEnvironmentPage} />
            <Route path="/environments/:envId/applications/:appId" component={ApplicationPage} />
            <Route path="/environments/:envId/applications" component={ApplicationsPage} />
          </Switch>
        </Router>
        {process.env.NODE_ENV === 'development' && <ReactQueryDevtools />}
      </QueryClientProvider>
    </ThemeProvider>
  </>
);

export default App;
